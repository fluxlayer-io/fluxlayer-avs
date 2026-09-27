// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import {BLSSignatureChecker, IRegistryCoordinator} from "@eigenlayer-middleware/src/BLSSignatureChecker.sol";

/// @title CrossChainLegVaultDVN
/// @notice A single-leg escrow, deployed independently on each chain a swap's two legs live on --
/// the same shape as CrossChainLegVault.sol (Section 4.13.6) -- but with a different release-
/// authorisation mechanism. CrossChainLegVault.release() is gated by the individual depositor's
/// own EIP-712 signature, collected and relayed by a trusted keeper. This contract instead gates
/// release() on a stake-weighted BLS quorum attestation from Layer I's own DVN operator set,
/// reusing the exact quorum-threshold-checking pattern OrderBook.respondToFulfill already uses
/// (BLSSignatureChecker.checkSignatures() plus an explicit signed-stake-vs-threshold check,
/// since checkSignatures() alone only proves a valid aggregate signature for whichever operators
/// signed, not what fraction of quorum stake that represents).
///
/// The attestation covers BOTH legs' exact terms jointly (see CrossChainRelease below), not just
/// this chain's own leg: the DVN quorum signs one message asserting "these are the exact terms of
/// both the maker's deposit and the taker's deposit for this order", and that same signed message
/// (and the same aggregate signature) is submitted to both chains' vaults, each independently
/// verifying it against its own local copy of the operator/stake registry. This is what lets the
/// release be described as quorum-confirmed on both legs at once, rather than each chain's release
/// being separately and independently authorised. It is still not single-transaction atomicity --
/// no vault design lets one transaction touch two different chains, so two separate release
/// transactions are still submitted, one per chain -- but neither release depends on any single
/// party's honesty or liveness the way CrossChainLegVault's keeper-relay design does: anyone
/// holding the quorum-signed attestation can submit either release, and the attestation itself
/// only exists once the DVN operators have jointly signed off on both legs' exact terms.
///
/// releaseWithQuorum() is deliberately NOT keeper-only (unlike CrossChainLegVault.release()):
/// authorisation comes entirely from the quorum attestation, not from the caller's identity, so
/// there is no keeper-liveness assumption left to state for this design -- the residual trust
/// assumption moves from "the keeper is honest and live" to "fewer than the quorum threshold's
/// worth of DVN operator stake is faulty", the same Byzantine-fault-tolerance assumption Layer I's
/// own quorum mechanism already rests on (Theorem 5.2).
contract CrossChainLegVaultDVN is BLSSignatureChecker {
    using SafeERC20 for IERC20;

    uint256 internal constant _THRESHOLD_DENOMINATOR = 100;

    enum LegStatus { Empty, Deposited, Released }

    struct Leg {
        address depositor;
        address token;
        uint256 amount;
        LegStatus status;
    }

    /// @notice One leg's exact terms: who deposited what, and who it releases to.
    struct LegTerms {
        address depositor;
        address token;
        uint256 amount;
        address recipient;
    }

    /// @notice The joint message the DVN quorum signs: one order's two legs, in full.
    struct CrossChainRelease {
        bytes32 orderId;
        LegTerms legA;
        LegTerms legB;
    }

    mapping(bytes32 => Leg) public legs;

    error LegAlreadyDeposited();
    error LegNotDeposited();
    error QuorumThresholdNotMet();
    error ReleaseTermsDoNotMatchEscrow();

    event LegDeposited(bytes32 indexed orderId, address depositor, address token, uint256 amount);
    event LegReleasedByQuorum(bytes32 indexed orderId, address recipient, bytes32 releaseMessageHash);

    constructor(IRegistryCoordinator _registryCoordinator) BLSSignatureChecker(_registryCoordinator) {}

    function deposit(bytes32 orderId, address token, uint256 amount) external {
        if (legs[orderId].status != LegStatus.Empty) revert LegAlreadyDeposited();
        IERC20(token).safeTransferFrom(msg.sender, address(this), amount);
        legs[orderId] = Leg(msg.sender, token, amount, LegStatus.Deposited);
        emit LegDeposited(orderId, msg.sender, token, amount);
    }

    function hashCrossChainRelease(CrossChainRelease calldata release) public pure returns (bytes32) {
        return keccak256(abi.encode(release));
    }

    /// @notice Releases this chain's leg of release.orderId, gated entirely by a DVN stake-quorum
    /// attestation over BOTH legs' exact terms -- not by the caller's identity, and not by either
    /// depositor's own signature. Callable by anyone holding a valid attestation.
    function releaseWithQuorum(
        CrossChainRelease calldata release,
        bytes calldata quorumNumbers,
        uint32 quorumThresholdPercentage,
        uint32 referenceBlockNumber,
        NonSignerStakesAndSignature memory nonSignerStakesAndSignature
    ) external returns (bytes32 releaseMessageHash) {
        releaseMessageHash = hashCrossChainRelease(release);

        (QuorumStakeTotals memory quorumStakeTotals, ) =
            checkSignatures(releaseMessageHash, quorumNumbers, referenceBlockNumber, nonSignerStakesAndSignature);

        // checkSignatures() only proves a valid aggregate BLS signature for whichever operators
        // signed; it does not by itself enforce that this represents quorumThresholdPercentage of
        // quorum stake. Enforce that explicitly, exactly as OrderBook.respondToFulfill does.
        for (uint256 i = 0; i < quorumNumbers.length; i++) {
            if (
                uint256(quorumStakeTotals.signedStakeForQuorum[i]) * _THRESHOLD_DENOMINATOR <
                uint256(quorumStakeTotals.totalStakeForQuorum[i]) * uint256(quorumThresholdPercentage)
            ) revert QuorumThresholdNotMet();
        }

        Leg storage leg = legs[release.orderId];
        if (leg.status != LegStatus.Deposited) revert LegNotDeposited();

        LegTerms memory matched = _matchEscrowedLeg(leg, release);

        leg.status = LegStatus.Released;
        IERC20(leg.token).safeTransfer(matched.recipient, leg.amount);
        emit LegReleasedByQuorum(release.orderId, matched.recipient, releaseMessageHash);
    }

    /// @dev Determines whether this vault's own escrowed leg is the release's legA or legB side
    /// (each CrossChainLegVaultDVN instance only ever holds one side of a given order), so the
    /// same joint CrossChainRelease struct can be submitted, unmodified, to both chains' vaults.
    function _matchEscrowedLeg(Leg storage leg, CrossChainRelease calldata release)
        internal
        view
        returns (LegTerms memory)
    {
        if (
            leg.depositor == release.legA.depositor &&
            leg.token == release.legA.token &&
            leg.amount == release.legA.amount
        ) return release.legA;
        if (
            leg.depositor == release.legB.depositor &&
            leg.token == release.legB.token &&
            leg.amount == release.legB.amount
        ) return release.legB;
        revert ReleaseTermsDoNotMatchEscrow();
    }

    function getLeg(bytes32 orderId) external view returns (Leg memory) {
        return legs[orderId];
    }
}
