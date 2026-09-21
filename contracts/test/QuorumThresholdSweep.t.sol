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

// Parameter-sensitivity sweep over the DVN quorum threshold tau, run against the FIXED
// OrderBook.respondToFulfill() (see src/OrderBook.sol's fix and test/OrderBook.t.sol for the
// bug it closes). With n=4 equal-stake mock operators, sweeps tau in {25,50,75,100}% against
// actual signed stake in {25,50,75,100}% (1..4 of 4 operators signing) and records, per cell,
// whether respondToFulfill accepts or rejects the response and its measured gas cost.
// Expected AND asserted outcome: acceptance iff signedPct >= tau -- exactly the enforcement
// the fix added, and exactly what Section 4.13's planned tau parameter-sensitivity study
// needs to demonstrate.
//
// Run with: forge test --match-contract QuorumThresholdSweepTest -vv
// and read the SWEEP-prefixed log lines for the raw per-cell data (tau, signedPct, result,
// gas). Each test function is one grid cell; Foundry gives each its own fresh EVM/registry
// state via setUp(), which is required here since the mock operator registration helper
// always registers the same deterministic operator addresses and can't be called twice
// against the same state.
contract QuorumThresholdSweepTest is BLSMockAVSDeployer, SignUtils {
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

    function setUp() public {
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

    // Runs one (tau, numNonSigners) grid cell: attempts respondToFulfill and logs + asserts
    // whether it succeeded and how much gas it used (0 if it reverted). pseudoRandomNumber
    // only needs to differ per cell to vary underlying key material; it does not affect the
    // pass/fail outcome.
    function _runCell(uint256 pseudoRandomNumber, uint256 numNonSigners, uint32 thresholdPercentage) internal {
        (uint32 referenceBlockNumber, BLSSignatureChecker.NonSignerStakesAndSignature memory nonSignerStakesAndSignature)
        = _registerSignatoriesAndGetNonSignerStakeAndSignatureRandom(pseudoRandomNumber, numNonSigners, quorumBitmap);

        uint256 signedPct = ((maxOperatorsToRegister - numNonSigners) * 100) / maxOperatorsToRegister;

        vm.prank(aggregator);
        uint256 gasBefore = gasleft();
        bool ok;
        try orderBook.respondToFulfill(
            quorumNumbers, thresholdPercentage, referenceBlockNumber, _buildResponse(), nonSignerStakesAndSignature
        ) {
            ok = true;
        } catch {
            ok = false;
        }
        uint256 gasUsed = ok ? gasBefore - gasleft() : 0;

        console2.log(
            string.concat(
                "SWEEP tau=", vm.toString(thresholdPercentage),
                " signedPct=", vm.toString(signedPct),
                " result=", ok ? "ACCEPT" : "REJECT",
                " gas=", vm.toString(gasUsed)
            )
        );

        assertEq(ok, signedPct >= uint256(thresholdPercentage), "acceptance must match signedPct >= tau exactly");
    }

    // signedPct = 25 (1 of 4 operators signs, numNonSigners = 3)
    function test_Sweep_Tau25_Signed25() public { _runCell(101, 3, 25); }
    function test_Sweep_Tau50_Signed25() public { _runCell(102, 3, 50); }
    function test_Sweep_Tau75_Signed25() public { _runCell(103, 3, 75); }
    function test_Sweep_Tau100_Signed25() public { _runCell(104, 3, 100); }

    // signedPct = 50 (2 of 4, numNonSigners = 2)
    function test_Sweep_Tau25_Signed50() public { _runCell(201, 2, 25); }
    function test_Sweep_Tau50_Signed50() public { _runCell(202, 2, 50); }
    function test_Sweep_Tau75_Signed50() public { _runCell(203, 2, 75); }
    function test_Sweep_Tau100_Signed50() public { _runCell(204, 2, 100); }

    // signedPct = 75 (3 of 4, numNonSigners = 1)
    function test_Sweep_Tau25_Signed75() public { _runCell(301, 1, 25); }
    function test_Sweep_Tau50_Signed75() public { _runCell(302, 1, 50); }
    function test_Sweep_Tau75_Signed75() public { _runCell(303, 1, 75); }
    function test_Sweep_Tau100_Signed75() public { _runCell(304, 1, 100); }

    // signedPct = 100 (4 of 4, numNonSigners = 0)
    function test_Sweep_Tau25_Signed100() public { _runCell(401, 0, 25); }
    function test_Sweep_Tau50_Signed100() public { _runCell(402, 0, 50); }
    function test_Sweep_Tau75_Signed100() public { _runCell(403, 0, 75); }
    function test_Sweep_Tau100_Signed100() public { _runCell(404, 0, 100); }
}
