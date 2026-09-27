// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

import "forge-std/Test.sol";
import "forge-std/console2.sol";
import "../src/CrossChainLegVaultDVN.sol";
import "../src/ERC20Mock.sol";
import "../lib/eigenlayer-middleware/test/utils/BLSMockAVSDeployer.sol";
import {IRegistryCoordinator} from "@eigenlayer-middleware/src/interfaces/IRegistryCoordinator.sol";
import {BLSSignatureChecker} from "@eigenlayer-middleware/src/BLSSignatureChecker.sol";
import {BN254} from "@eigenlayer-middleware/src/libraries/BN254.sol";
import {BitmapUtils} from "@eigenlayer-middleware/src/libraries/BitmapUtils.sol";

/// @dev One independent "chain": BLSMockAVSDeployer._setUpBLSMockAVSDeployer() deploys a fresh,
/// wholly self-contained set of mock EigenLayer core + real middleware registry contracts on
/// every call (they are ordinary `new` deployments assigned to this instance's own state
/// variables, not a shared singleton), so two ChainQuorumHarness instances are two genuinely
/// independent on-chain registries -- modelling two independent chains within one Foundry test,
/// consistent with how Layer I's own existing quorum benchmarks (Table 4.9/4.10, tau-sweep and
/// validator-count sweep) are already reported: a Foundry sweep, not a live two-process Anvil
/// deployment. Each harness also deploys its own CrossChainLegVaultDVN wired to its own
/// RegistryCoordinator.
contract ChainQuorumHarness is BLSMockAVSDeployer {
    using BN254 for BN254.G1Point;

    CrossChainLegVaultDVN public vault;

    function init() external {
        _setUpBLSMockAVSDeployer();
        vault = new CrossChainLegVaultDVN(IRegistryCoordinator(address(registryCoordinator)));
    }

    /// @dev Registers maxOperatorsToRegister deterministic operators (identical addresses and BLS
    /// keys across any two ChainQuorumHarness instances called with the same pseudoRandomNumber,
    /// since _generateSignerAndNonSignerPrivateKeys is a pure function of that seed and
    /// aggSignerPrivKey is a fixed constant in BLSMockAVSDeployer), then overrides the harness's
    /// signed message to releaseMessageHash before computing the aggregate signature -- the same
    /// msgHash/sigma override technique OrderBookQuorumThresholdTest.setUp() already uses
    /// elsewhere in this codebase -- so the returned attestation is genuinely over this specific
    /// cross-chain release, not the harness's placeholder "hello world" message.
    function attestRelease(bytes32 releaseMessageHash, uint256 pseudoRandomNumber, uint256 numNonSigners, uint256 quorumBitmap)
        external
        returns (uint32 referenceBlockNumber, BLSSignatureChecker.NonSignerStakesAndSignature memory nonSignerStakesAndSignature)
    {
        msgHash = releaseMessageHash;
        sigma = BN254.hashToG1(msgHash).scalar_mul(aggSignerPrivKey);
        return _registerSignatoriesAndGetNonSignerStakeAndSignatureRandom(pseudoRandomNumber, numNonSigners, quorumBitmap);
    }
}

/// Live-measured (real BLS math, real on-chain quorum verification, real gas), Foundry-simulated
/// two-chain benchmark for CrossChainLegVaultDVN -- the DVN-quorum-gated counterpart to
/// crosschain_leg_benchmark_test.go's keeper-based CrossChainLegVault measurement
/// (Section 4.13.6/4.13.8). "Chain A" is the maker's chain, "chain B" is the taker's chain,
/// mirroring that same benchmark's topology exactly, so the two designs are compared like for
/// like: same swap shape, different release-authorisation mechanism.
contract CrossChainLegVaultDVNTest is Test {
    using BN254 for BN254.G1Point;

    ChainQuorumHarness chainA;
    ChainQuorumHarness chainB;
    // Cached directly after init(), not read back through chainA.vault()/chainB.vault()'s own
    // external getter call at each use site: reading a public state variable through another
    // contract is itself an external call, which would silently consume a preceding vm.prank
    // meant for the call that follows it instead.
    CrossChainLegVaultDVN vaultA;
    CrossChainLegVaultDVN vaultB;

    ERC20Mock tokenA;
    ERC20Mock tokenB;

    address makerAddrA = address(uint160(uint256(keccak256("makerAddrA"))));
    address takerAddrA = address(uint160(uint256(keccak256("takerAddrA"))));
    address makerAddrB = address(uint160(uint256(keccak256("makerAddrB"))));
    address takerAddrB = address(uint160(uint256(keccak256("takerAddrB"))));

    uint256 constant AMOUNT = 100 ether;
    bytes32 orderId = keccak256("dvn-crosschain-order-1");
    uint256 quorumBitmap = 1;
    bytes quorumNumbers;

    CrossChainLegVaultDVN.CrossChainRelease release;

    function setUp() public {
        quorumNumbers = BitmapUtils.bitmapToBytesArray(quorumBitmap);

        chainA = new ChainQuorumHarness();
        chainA.init();
        chainB = new ChainQuorumHarness();
        chainB.init();
        vaultA = chainA.vault();
        vaultB = chainB.vault();

        tokenA = new ERC20Mock();
        tokenB = new ERC20Mock();

        tokenA.mint(makerAddrA, AMOUNT);
        tokenB.mint(takerAddrB, AMOUNT);

        vm.prank(makerAddrA);
        tokenA.approve(address(vaultA), AMOUNT);
        vm.prank(makerAddrA);
        vaultA.deposit(orderId, address(tokenA), AMOUNT);

        vm.prank(takerAddrB);
        tokenB.approve(address(vaultB), AMOUNT);
        vm.prank(takerAddrB);
        vaultB.deposit(orderId, address(tokenB), AMOUNT);

        release = CrossChainLegVaultDVN.CrossChainRelease({
            orderId: orderId,
            legA: CrossChainLegVaultDVN.LegTerms({
                depositor: makerAddrA, token: address(tokenA), amount: AMOUNT, recipient: takerAddrA
            }),
            legB: CrossChainLegVaultDVN.LegTerms({
                depositor: takerAddrB, token: address(tokenB), amount: AMOUNT, recipient: makerAddrB
            })
        });
    }

    function _releaseMessageHash() internal view returns (bytes32) {
        return vaultA.hashCrossChainRelease(release);
    }

    /// @notice Full quorum (all 4 mock operators, registered identically and independently on
    /// both chains) signs the joint release message once; the SAME message hash and SAME
    /// aggregate signature (sigma) is then submitted, together with each chain's own locally
    /// computed NonSignerStakesAndSignature, to both chains' vaults -- proving one joint DVN
    /// quorum decision genuinely gates both releases, not two separately-authorised ones.
    function testReleaseWithQuorum_FullQuorumReleasesBothLegs() public {
        bytes32 releaseMessageHash = _releaseMessageHash();

        (uint32 refBlockA, BLSSignatureChecker.NonSignerStakesAndSignature memory attA) =
            chainA.attestRelease(releaseMessageHash, 1, 0, quorumBitmap);
        (uint32 refBlockB, BLSSignatureChecker.NonSignerStakesAndSignature memory attB) =
            chainB.attestRelease(releaseMessageHash, 1, 0, quorumBitmap);

        uint256 gasBeforeA = gasleft();
        vaultA.releaseWithQuorum(release, quorumNumbers, 100, refBlockA, attA);
        uint256 gasUsedA = gasBeforeA - gasleft();
        console2.log("BENCH case=dvn_quorum_crosschain chain=A releaseWithQuorum_gas", gasUsedA);

        uint256 gasBeforeB = gasleft();
        vaultB.releaseWithQuorum(release, quorumNumbers, 100, refBlockB, attB);
        uint256 gasUsedB = gasBeforeB - gasleft();
        console2.log("BENCH case=dvn_quorum_crosschain chain=B releaseWithQuorum_gas", gasUsedB);

        assertEq(tokenA.balanceOf(takerAddrA), AMOUNT, "taker should receive the maker's chain-A leg once the quorum signs off");
        assertEq(tokenB.balanceOf(makerAddrB), AMOUNT, "maker should receive the taker's chain-B leg once the quorum signs off");
        assertEq(uint8(vaultA.getLeg(orderId).status), uint8(CrossChainLegVaultDVN.LegStatus.Released));
        assertEq(uint8(vaultB.getLeg(orderId).status), uint8(CrossChainLegVaultDVN.LegStatus.Released));
    }

    /// @notice Regression test mirroring OrderBookQuorumThresholdTest.
    /// testRespondToFulfill_LowStakeSignatureIsRejectedAt100PercentThreshold: only 1 of 4
    /// equal-stake operators signs (~25% of quorum stake) against a 100% threshold requirement.
    /// Both chains must independently reject the release and leave their own escrow untouched --
    /// confirming the quorum-threshold check is real, not merely a valid-signature check.
    function testReleaseWithQuorum_LowStakeSignatureIsRejectedAt100PercentThreshold() public {
        bytes32 releaseMessageHash = _releaseMessageHash();

        (uint32 refBlockA, BLSSignatureChecker.NonSignerStakesAndSignature memory attA) =
            chainA.attestRelease(releaseMessageHash, 2, 3, quorumBitmap); // 4 operators, 3 non-signers
        (uint32 refBlockB, BLSSignatureChecker.NonSignerStakesAndSignature memory attB) =
            chainB.attestRelease(releaseMessageHash, 2, 3, quorumBitmap);

        vm.expectRevert(CrossChainLegVaultDVN.QuorumThresholdNotMet.selector);
        vaultA.releaseWithQuorum(release, quorumNumbers, 100, refBlockA, attA);

        vm.expectRevert(CrossChainLegVaultDVN.QuorumThresholdNotMet.selector);
        vaultB.releaseWithQuorum(release, quorumNumbers, 100, refBlockB, attB);

        assertEq(tokenA.balanceOf(takerAddrA), 0, "chain A leg must remain escrowed, unchanged, when quorum threshold is not met");
        assertEq(tokenB.balanceOf(makerAddrB), 0, "chain B leg must remain escrowed, unchanged, when quorum threshold is not met");
        assertEq(uint8(vaultA.getLeg(orderId).status), uint8(CrossChainLegVaultDVN.LegStatus.Deposited));
        assertEq(uint8(vaultB.getLeg(orderId).status), uint8(CrossChainLegVaultDVN.LegStatus.Deposited));
    }
}
