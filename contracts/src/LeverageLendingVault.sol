// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

/// @title LeverageLendingVault
/// @notice Layer III of FluxLayer: an under-collateralised leverage lending vault (V_L) that
/// funds a maker's arbitrage order with internally-contained loan capital, subject to a
/// deterministic liquidation window. This contract is a direct implementation of the formal
/// model in the thesis (Section "Layer III: Under-Collateralised Leverage Lending Vault"):
///
///   C            = collateral value                                (Eq. 8)
///   L            = loan amount                                     (Eq. 9)
///   lambda = L/C = leverage ratio                                  (Eq. 10)
///   L <= lambda_max * C, lambda_max < infinity                     (Eq. 11)
///   for all t, f_borrow(t) is in P_internal (fund containment)     (Eq. 12)
///   position opened at t0 is liquidated for all t > t0 + Delta     (Eq. 13)
///   solvency invariant at liquidation: C_i(tL) >= L_i(tL) - R_i(tL) (Theorem 5.2's proof)
///
/// A single vault instance lends a single ERC20 asset (the thesis's "high-liquidity collateral,
/// e.g. USDT, USDC" assumption underlying lambda_max = 10 and Delta = 15 minutes). Both
/// collateral and the loan are denominated in that same asset, matching the deterministic
/// price-stability assumption (A4) the leverage bound relies on.
///
/// Fund containment (Eq. 12) is approximated on-chain by requiring the loan to be disbursed
/// only to an address the vault owner (the FluxLayer keeper operating fluxlayer-mpc) has
/// pre-registered as an internal execution address for that specific order, rather than to an
/// address chosen freely by the borrower.
contract LeverageLendingVault is Ownable {
    using SafeERC20 for IERC20;

    /// @notice The single asset this vault lends and accepts as collateral.
    IERC20 public immutable asset;

    /// @notice Leverage cap lambda_max (Eq. 11), scaled by LEVERAGE_DENOMINATOR.
    uint256 public immutable leverageMaxNumerator;
    uint256 public constant LEVERAGE_DENOMINATOR = 100;

    /// @notice Deterministic liquidation window Delta (Eq. 13), in seconds.
    uint256 public immutable liquidationWindow;

    enum PositionStatus {
        None,
        Open,
        Repaid,
        Liquidated
    }

    struct Position {
        address maker;
        uint256 collateral; // C, held by the vault separately from LP-withdrawable liquidity
        uint256 loanAmount; // L
        uint256 repaidAmount; // R, accumulated via repay()
        uint256 openedAt; // t0
        uint256 deadline; // t0 + Delta
        PositionStatus status;
    }

    /// @notice LP capital not currently lent out. Excludes borrower collateral, which is held
    /// by the vault but tracked per-position (`Position.collateral`) and is never withdrawable
    /// by LPs. openPosition() debits a loan from this pool; repay() and the seized portion of a
    /// liquidated position's collateral credit it back.
    uint256 public availableLiquidity;

    /// @notice LP share ledger (a simplified, non-tokenised vault: 1 unit deposited = 1 unit of
    /// claim, since this contract does not yet track accrued yield/fees per share).
    mapping(address => uint256) public lpBalance;
    uint256 public totalLpDeposits;

    /// @notice Addresses the owner has designated as valid loan-disbursement targets for a given
    /// order, standing in for Eq. 12's "internal execution environment" containment property.
    mapping(bytes32 => address) public internalExecutionAddress;

    mapping(bytes32 => Position) public positions;

    event LiquidityDeposited(address indexed lp, uint256 amount);
    event LiquidityWithdrawn(address indexed lp, uint256 amount);
    event InternalExecutionAddressRegistered(bytes32 indexed orderId, address indexed execAddress);
    event PositionOpened(
        bytes32 indexed orderId, address indexed maker, uint256 collateral, uint256 loanAmount, uint256 deadline
    );
    event PositionRepaid(bytes32 indexed orderId, uint256 repaidAmount, bool closed);
    event PositionLiquidated(bytes32 indexed orderId, uint256 collateralSeized, uint256 outstanding);

    error LeverageExceeded(uint256 loanAmount, uint256 maxLoanAmount);
    error InsufficientLiquidity(uint256 requested, uint256 available);
    error UnregisteredExecutionAddress();
    error PositionNotOpen(bytes32 orderId);
    error DeadlinePassed(bytes32 orderId, uint256 deadline);
    error DeadlineNotPassed(bytes32 orderId, uint256 deadline);
    error RepayExceedsOutstanding(uint256 repayAmount, uint256 outstanding);

    constructor(IERC20 _asset, uint256 _leverageMaxNumerator, uint256 _liquidationWindow) Ownable() {
        asset = _asset;
        leverageMaxNumerator = _leverageMaxNumerator;
        liquidationWindow = _liquidationWindow;
    }

    /// @notice LPs supply the internal liquidity pool that funds maker loans.
    function deposit(uint256 amount) external {
        asset.safeTransferFrom(msg.sender, address(this), amount);
        lpBalance[msg.sender] += amount;
        totalLpDeposits += amount;
        availableLiquidity += amount;
        emit LiquidityDeposited(msg.sender, amount);
    }

    /// @notice LPs withdraw liquidity that is not currently lent out. Borrower collateral is
    /// never part of `availableLiquidity`, so LP withdrawals cannot touch it.
    function withdraw(uint256 amount) external {
        require(lpBalance[msg.sender] >= amount, "insufficient LP balance");
        require(availableLiquidity >= amount, "insufficient available liquidity");
        lpBalance[msg.sender] -= amount;
        totalLpDeposits -= amount;
        availableLiquidity -= amount;
        asset.safeTransfer(msg.sender, amount);
        emit LiquidityWithdrawn(msg.sender, amount);
    }

    /// @notice The vault owner (the FluxLayer keeper) registers, for a given order, the address
    /// that a loan for that order may be disbursed to -- e.g. the maker's Cobo-custodied
    /// source-chain MPC wallet address that fluxlayer-mpc already controls for that order.
    function registerInternalExecutionAddress(bytes32 orderId, address execAddress) external onlyOwner {
        internalExecutionAddress[orderId] = execAddress;
        emit InternalExecutionAddressRegistered(orderId, execAddress);
    }

    /// @notice Opens a leveraged position for `orderId`: pulls `collateral` from the maker
    /// (held separately from LP liquidity), checks the leverage constraint (Eq. 11) and
    /// available LP liquidity, then disburses `loanAmount` out of the LP pool to the order's
    /// registered internal execution address (Eq. 12).
    function openPosition(bytes32 orderId, address maker, uint256 collateral, uint256 loanAmount)
        external
        onlyOwner
        returns (uint256 deadline)
    {
        if (positions[orderId].status != PositionStatus.None) revert PositionNotOpen(orderId);
        address execAddress = internalExecutionAddress[orderId];
        if (execAddress == address(0)) revert UnregisteredExecutionAddress();

        uint256 maxLoanAmount = (collateral * leverageMaxNumerator) / LEVERAGE_DENOMINATOR;
        if (loanAmount > maxLoanAmount) revert LeverageExceeded(loanAmount, maxLoanAmount);
        if (loanAmount > availableLiquidity) revert InsufficientLiquidity(loanAmount, availableLiquidity);

        asset.safeTransferFrom(maker, address(this), collateral);
        availableLiquidity -= loanAmount;

        deadline = block.timestamp + liquidationWindow;
        positions[orderId] = Position({
            maker: maker,
            collateral: collateral,
            loanAmount: loanAmount,
            repaidAmount: 0,
            openedAt: block.timestamp,
            deadline: deadline,
            status: PositionStatus.Open
        });

        asset.safeTransfer(execAddress, loanAmount);
        emit PositionOpened(orderId, maker, collateral, loanAmount, deadline);
    }

    /// @notice Repays some or all of an open position's loan, before its deadline. Funds come
    /// from the settlement proceeds the keeper has swept back from the order's execution.
    /// Repaid amounts are credited back to the LP-withdrawable pool immediately. Once the
    /// accumulated repayment covers the loan, the position closes and its full collateral is
    /// returned to the maker (Theorem 5.2's solvency invariant, satisfied with room to spare
    /// rather than merely met, once R >= L).
    function repay(bytes32 orderId, uint256 amount) external {
        Position storage position = positions[orderId];
        if (position.status != PositionStatus.Open) revert PositionNotOpen(orderId);
        if (block.timestamp > position.deadline) revert DeadlinePassed(orderId, position.deadline);

        uint256 outstanding = position.loanAmount - position.repaidAmount;
        if (amount > outstanding) revert RepayExceedsOutstanding(amount, outstanding);

        asset.safeTransferFrom(msg.sender, address(this), amount);
        position.repaidAmount += amount;
        availableLiquidity += amount;

        bool closed = position.repaidAmount == position.loanAmount;
        if (closed) {
            position.status = PositionStatus.Repaid;
            uint256 collateralReturned = position.collateral;
            position.collateral = 0;
            asset.safeTransfer(position.maker, collateralReturned);
        }
        emit PositionRepaid(orderId, amount, closed);
    }

    /// @notice Forcibly closes a position once its deadline has passed without full repayment
    /// (Eq. 13). Seizes up to `outstanding` of the position's collateral to make the LP pool
    /// whole; the solvency invariant C >= L - R (Theorem 5.2's proof) is what the thesis's
    /// Delta = 15 minutes / high-liquidity-collateral assumption (A4) is relied on to guarantee,
    /// so in the modelled honest/no-price-shock case `seized == outstanding` and the position
    /// closes with no loss to LPs. Any collateral left over after covering the outstanding
    /// balance is returned to the maker.
    function liquidate(bytes32 orderId) external {
        Position storage position = positions[orderId];
        if (position.status != PositionStatus.Open) revert PositionNotOpen(orderId);
        if (block.timestamp <= position.deadline) revert DeadlineNotPassed(orderId, position.deadline);

        uint256 outstanding = position.loanAmount - position.repaidAmount;
        uint256 collateral = position.collateral;
        position.collateral = 0;
        position.status = PositionStatus.Liquidated;

        uint256 seized = outstanding < collateral ? outstanding : collateral;
        uint256 surplus = collateral - seized;
        availableLiquidity += seized;
        if (surplus > 0) {
            asset.safeTransfer(position.maker, surplus);
        }
        emit PositionLiquidated(orderId, seized, outstanding);
    }

    function getPosition(bytes32 orderId) external view returns (Position memory) {
        return positions[orderId];
    }
}
