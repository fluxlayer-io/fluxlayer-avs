pragma solidity ^0.8.12;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";

/// @title CrossChainLegVault
/// @notice A single-leg escrow, deployed independently on each chain a swap's two legs live on:
/// one instance on the maker's chain holding the maker's deposit, a separate instance on the
/// taker's chain holding the taker's deposit. Each depositor pre-authorises release of their own
/// leg -- to an exact recipient, for an exact amount -- with an EIP-712 signature; only
/// FluxLayer's keeper can call release(), and the keeper only ever does so once it holds BOTH
/// parties' signatures for the paired order, collected off-chain before either release is
/// submitted on either chain.
///
/// This is the general cross-chain case, and it is NOT single-transaction atomicity: no vault
/// design lets one transaction touch two different chains, so the keeper necessarily submits two
/// separate release transactions, one per chain, back-to-back rather than as a single atomic
/// unit. What it gives instead is dual-signature-gated, keeper-coordinated release: neither leg
/// can ever be redirected to a recipient its own depositor did not sign for, so the keeper alone
/// cannot steal or misdirect either side -- but there is a real, if narrow, window between the two
/// release transactions during which only one of them has confirmed, whose width this thesis
/// measures directly rather than asserting away. A keeper that never submits either release
/// leaves both legs safely un-released in their own escrow, refundable by whatever timeout policy
/// the integrating order logic chooses (not modelled in this contract itself).
///
/// When both legs happen to share the same chain, they can instead be escrowed in a single
/// TwoOfTwoSettlementVault instance and released together in one transaction, for genuine
/// same-chain atomicity (see TwoOfTwoSettlementVault.sol) -- a special case this general,
/// per-chain vault does not need and does not attempt to reproduce.
contract CrossChainLegVault {
    using ECDSA for bytes32;
    using SafeERC20 for IERC20;

    enum LegStatus { Empty, Deposited, Released }

    struct Leg {
        address depositor;
        address token;
        uint256 amount;
        LegStatus status;
    }

    mapping(bytes32 => Leg) public legs;

    address public immutable keeper;
    bytes32 private immutable _DOMAIN_SEPARATOR;
    bytes32 private constant _RELEASE_TYPEHASH =
        keccak256("Release(bytes32 orderId,address depositor,address token,uint256 amount,address recipient)");

    error NotKeeper();
    error LegAlreadyDeposited();
    error LegNotDeposited();
    error InvalidSignature();

    event LegDeposited(bytes32 indexed orderId, address depositor, address token, uint256 amount);
    event LegReleased(bytes32 indexed orderId, address recipient);

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

    function deposit(bytes32 orderId, address token, uint256 amount) external {
        if (legs[orderId].status != LegStatus.Empty) revert LegAlreadyDeposited();
        IERC20(token).safeTransferFrom(msg.sender, address(this), amount);
        legs[orderId] = Leg(msg.sender, token, amount, LegStatus.Deposited);
        emit LegDeposited(orderId, msg.sender, token, amount);
    }

    function hashRelease(bytes32 orderId, address depositor, address token, uint256 amount, address recipient)
        public
        view
        returns (bytes32)
    {
        bytes32 structHash = keccak256(abi.encode(_RELEASE_TYPEHASH, orderId, depositor, token, amount, recipient));
        return keccak256(abi.encodePacked("\x19\x01", _DOMAIN_SEPARATOR, structHash));
    }

    /// @notice Releases this chain's leg of orderId to recipient. Callable only by the keeper, and
    /// only with the depositor's own signature over these exact terms -- the keeper alone cannot
    /// redirect funds to any recipient the depositor did not sign for.
    function release(bytes32 orderId, address recipient, bytes calldata depositorSig) external {
        if (msg.sender != keeper) revert NotKeeper();
        Leg storage leg = legs[orderId];
        if (leg.status != LegStatus.Deposited) revert LegNotDeposited();

        bytes32 h = hashRelease(orderId, leg.depositor, leg.token, leg.amount, recipient);
        if (h.recover(depositorSig) != leg.depositor) revert InvalidSignature();

        leg.status = LegStatus.Released;
        IERC20(leg.token).safeTransfer(recipient, leg.amount);
        emit LegReleased(orderId, recipient);
    }

    function getLeg(bytes32 orderId) external view returns (Leg memory) {
        return legs[orderId];
    }
}
