pragma solidity ^0.8.12;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";

/// @title TwoOfTwoSettlementVault
/// @notice Each order gets two independent 2-of-2 legs: a maker leg (maker + FluxLayer keeper)
/// and a taker leg (taker + FluxLayer keeper). Each depositor pre-authorises release of their
/// own leg with an EIP-712 signature over the exact swap terms; the keeper is the only account
/// that can ever call settle(), and settle() releases BOTH legs inside a single transaction --
/// if either transfer fails for any reason, the whole call reverts and neither leg moves.
///
/// This is what "atomic" means here: not a single transaction spanning two different chains
/// (that is a physical impossibility no vault design changes), but a single transaction, on the
/// chain both legs already share, that cannot partially succeed. Both legs must live on the
/// same chain for settle() to be callable at all; a genuinely cross-chain order (maker and
/// taker legs on different chains) still needs the DVN-quorum-gated release this replaces on
/// neither, since no single EVM transaction can touch two chains.
contract TwoOfTwoSettlementVault {
    using ECDSA for bytes32;
    using SafeERC20 for IERC20;

    enum LegStatus {
        Empty,
        Deposited,
        Released
    }

    struct Leg {
        address depositor;
        address token;
        uint256 amount;
        LegStatus status;
    }

    // orderId => maker leg / taker leg
    mapping(bytes32 => Leg) public makerLegs;
    mapping(bytes32 => Leg) public takerLegs;

    address public immutable keeper;
    bytes32 private immutable _DOMAIN_SEPARATOR;
    bytes32 private constant _RELEASE_TYPEHASH =
        keccak256("Release(bytes32 orderId,address depositor,address token,uint256 amount,address recipient)");

    error NotKeeper();
    error LegAlreadyDeposited();
    error LegNotDeposited();
    error InvalidSignature();

    event LegDeposited(bytes32 indexed orderId, bool isMakerLeg, address depositor, address token, uint256 amount);
    event Settled(bytes32 indexed orderId);

    constructor(address _keeper, string memory name, string memory version) {
        keeper = _keeper;
        _DOMAIN_SEPARATOR = keccak256(
            abi.encode(
                keccak256("EIP712Domain(string name,string version,uint256 chainId,address verifyingContract)"),
                keccak256(bytes(name)),
                keccak256(bytes(version)),
                block.chainid,
                address(this)
            )
        );
    }

    function depositMakerLeg(bytes32 orderId, address token, uint256 amount) external {
        _deposit(makerLegs, orderId, token, amount, true);
    }

    function depositTakerLeg(bytes32 orderId, address token, uint256 amount) external {
        _deposit(takerLegs, orderId, token, amount, false);
    }

    function _deposit(mapping(bytes32 => Leg) storage legs, bytes32 orderId, address token, uint256 amount, bool isMaker)
        internal
    {
        if (legs[orderId].status != LegStatus.Empty) revert LegAlreadyDeposited();
        IERC20(token).safeTransferFrom(msg.sender, address(this), amount);
        legs[orderId] = Leg(msg.sender, token, amount, LegStatus.Deposited);
        emit LegDeposited(orderId, isMaker, msg.sender, token, amount);
    }

    function hashRelease(bytes32 orderId, address depositor, address token, uint256 amount, address recipient)
        public
        view
        returns (bytes32)
    {
        bytes32 structHash = keccak256(abi.encode(_RELEASE_TYPEHASH, orderId, depositor, token, amount, recipient));
        return keccak256(abi.encodePacked("\x19\x01", _DOMAIN_SEPARATOR, structHash));
    }

    /// @notice Atomically releases both legs of orderId: the maker's deposited token to the
    /// taker, and the taker's deposited token to the maker, in the same transaction. Requires
    /// the keeper to call it (so no depositor can unilaterally trigger the other side's
    /// release) AND requires each depositor's own signature authorising release to the
    /// counterparty (so the keeper alone cannot move either leg without that depositor's
    /// consent) -- a 2-of-2 requirement on each leg. If either transfer reverts, the whole call
    /// reverts, so it is never possible for one leg to move without the other.
    function settle(bytes32 orderId, bytes calldata makerSig, bytes calldata takerSig) external {
        if (msg.sender != keeper) revert NotKeeper();
        Leg storage makerLeg = makerLegs[orderId];
        Leg storage takerLeg = takerLegs[orderId];
        if (makerLeg.status != LegStatus.Deposited) revert LegNotDeposited();
        if (takerLeg.status != LegStatus.Deposited) revert LegNotDeposited();

        bytes32 makerHash = hashRelease(orderId, makerLeg.depositor, makerLeg.token, makerLeg.amount, takerLeg.depositor);
        if (makerHash.recover(makerSig) != makerLeg.depositor) revert InvalidSignature();
        bytes32 takerHash = hashRelease(orderId, takerLeg.depositor, takerLeg.token, takerLeg.amount, makerLeg.depositor);
        if (takerHash.recover(takerSig) != takerLeg.depositor) revert InvalidSignature();

        makerLeg.status = LegStatus.Released;
        takerLeg.status = LegStatus.Released;

        IERC20(makerLeg.token).safeTransfer(takerLeg.depositor, makerLeg.amount);
        IERC20(takerLeg.token).safeTransfer(makerLeg.depositor, takerLeg.amount);

        emit Settled(orderId);
    }

    function getMakerLeg(bytes32 orderId) external view returns (Leg memory) {
        return makerLegs[orderId];
    }

    function getTakerLeg(bytes32 orderId) external view returns (Leg memory) {
        return takerLegs[orderId];
    }
}
