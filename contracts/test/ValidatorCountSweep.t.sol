// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

import "../src/OrderBook.sol";
import "../src/IOrderBook.sol";
import "../src/ERC20Mock.sol";
import "forge-std/Test.sol";
import "forge-std/console2.sol";
import "../lib/eigenlayer-middleware/test/utils/BLSMockAVSDeployer.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "../script/utils/SignUtils.sol";
import "../script/utils/EIP712Utils.sol";
import {BitmapUtils} from "../lib/eigenlayer-middleware/src/libraries/BitmapUtils.sol";
import {BN254} from "../lib/eigenlayer-middleware/src/libraries/BN254.sol";

// Parameter-sensitivity sweep over the DVN validator (operator) count n, run against the FIXED
// OrderBook.respondToFulfill() -- the companion sweep to test/QuorumThresholdSweep.t.sol, which
// fixes n=4 and sweeps tau and signed stake. Here tau is fixed at 67% and every registered
// operator signs (numNonSigners=0, signedPct=100%always >= tau), so the only independent
// variable across cells is n itself, and the only thing measured is how respondToFulfill's gas
// cost scales as the registered operator set grows.
//
// Reuses the same BLSMockAVSDeployer/MockAVSDeployer harness as QuorumThresholdSweep.t.sol, but
// overrides maxOperatorsToRegister and the registry coordinator's per-quorum operator cap
// (defaultMaxOperatorCount) per contract instead of using the harness's hardcoded n=4. Each n
// value is its own contract (Foundry gives each contract, and each test within it, fresh
// EVM/registry state), since maxOperatorsToRegister must be set before _setUpBLSMockAVSDeployer()
// runs inside setUp().
//
// Run with: forge test --match-contract ValidatorCountSweep -vv
// and read the SWEEP-prefixed log lines for the raw per-cell data (n, result, gas).
abstract contract ValidatorCountSweepBase is BLSMockAVSDeployer, SignUtils {
    using BN254 for BN254.G1Point;

    OrderBook public orderBook;
    EIP712Utils public eip712Utils;

    address aggregator = address(uint160(uint256(keccak256(abi.encodePacked("aggregator")))));
    uint256 makerPk = 1;
    address maker = vm.addr(makerPk);
    address taker = address(0xBEEF);

    address inputToken;
    uint256 inputAmount = 100 ether;
    address outputToken = address(0xCAFE);
    uint256 outputAmount = 100 ether;
    uint256 expiry = 9999999999;
    uint32 signChainId = 17000;
    uint32 targetNetworkNumber = 11155111;
    uint32 orderId = 1;

    IOrderBook.Order order;
    uint256 quorumBitmap = 1;
    bytes quorumNumbers;

    ERC20Mock public inputErc20;

    uint32 constant FIXED_TAU_PCT = 67;

    function _setN(uint256 n) internal {
        maxOperatorsToRegister = n;
        defaultMaxOperatorCount = uint32(n);
    }

    function setUp() public virtual {
        _setUpBLSMockAVSDeployer();
        quorumNumbers = BitmapUtils.bitmapToBytesArray(quorumBitmap);

        OrderBook orderBookImp = new OrderBook(IRegistryCoordinator(address(registryCoordinator)), signChainId);
        orderBook = OrderBook(
            address(
                new TransparentUpgradeableProxy(
                    address(orderBookImp),
                    address(proxyAdmin),
                    abi.encodeWithSelector(
                        orderBookImp.initialize.selector,
                        pauserRegistry,
                        registryCoordinatorOwner,
                        aggregator
                    )
                )
            )
        );

        eip712Utils = new EIP712Utils("OrderBook", "1.0", signChainId, address(orderBook));
        inputErc20 = new ERC20Mock();
        inputToken = address(inputErc20);

        order = IOrderBook.Order(
            orderId, maker, taker, inputToken, inputAmount, outputToken, outputAmount, expiry, targetNetworkNumber
        );
        bytes memory sig = signHash(makerPk, eip712Utils.getTypedDataHash(order));

        inputErc20.mint(maker, inputAmount);
        vm.prank(taker);
        orderBook.createOrder(order, sig);

        IOrderBook.OrderResponse memory orderResponse = _buildResponse();
        msgHash = keccak256(abi.encode(orderResponse));
        sigma = BN254.hashToG1(msgHash).scalar_mul(aggSignerPrivKey);
    }

    function _buildResponse() internal view returns (IOrderBook.OrderResponse memory) {
        return IOrderBook.OrderResponse({recipient: taker, referenceOrderIndex: orderId});
    }

    function _runCell(uint256 pseudoRandomNumber) internal {
        (uint32 referenceBlockNumber, BLSSignatureChecker.NonSignerStakesAndSignature memory nonSignerStakesAndSignature)
        = _registerSignatoriesAndGetNonSignerStakeAndSignatureRandom(pseudoRandomNumber, 0, quorumBitmap);

        vm.prank(aggregator);
        uint256 gasBefore = gasleft();
        bool ok;
        try orderBook.respondToFulfill(
            quorumNumbers, FIXED_TAU_PCT, referenceBlockNumber, _buildResponse(), nonSignerStakesAndSignature
        ) {
            ok = true;
        } catch {
            ok = false;
        }
        uint256 gasUsed = ok ? gasBefore - gasleft() : 0;

        console2.log(
            string.concat(
                "SWEEP n=", vm.toString(maxOperatorsToRegister),
                " tau=", vm.toString(FIXED_TAU_PCT),
                " result=", ok ? "ACCEPT" : "REJECT",
                " gas=", vm.toString(gasUsed)
            )
        );

        assertTrue(ok, "full participation must always accept at any n");
    }
}

contract ValidatorCountSweep_N4Test is ValidatorCountSweepBase {
    function setUp() public override {
        _setN(4);
        super.setUp();
    }

    function test_Sweep_N4() public {
        _runCell(501);
    }
}

contract ValidatorCountSweep_N8Test is ValidatorCountSweepBase {
    function setUp() public override {
        _setN(8);
        super.setUp();
    }

    function test_Sweep_N8() public {
        _runCell(502);
    }
}

contract ValidatorCountSweep_N16Test is ValidatorCountSweepBase {
    function setUp() public override {
        _setN(16);
        super.setUp();
    }

    function test_Sweep_N16() public {
        _runCell(503);
    }
}

contract ValidatorCountSweep_N32Test is ValidatorCountSweepBase {
    function setUp() public override {
        _setN(32);
        super.setUp();
    }

    function test_Sweep_N32() public {
        _runCell(504);
    }
}
