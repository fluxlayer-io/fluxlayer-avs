// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

import "../src/OrderBook.sol";
import "../src/IOrderBook.sol";
import "../src/ERC20Mock.sol";
import "forge-std/Test.sol";
import "../lib/eigenlayer-middleware/test/utils/BLSMockAVSDeployer.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "../script/utils/SignUtils.sol";
import "../script/utils/EIP712Utils.sol";
import {BitmapUtils} from "../lib/eigenlayer-middleware/src/libraries/BitmapUtils.sol";
import {BN254} from "../lib/eigenlayer-middleware/src/libraries/BN254.sol";

// Regression tests for the DVN quorum-threshold check in OrderBook.respondToFulfill().
//
// checkSignatures() (inherited from BLSSignatureChecker) only proves that the aggregate
// BLS signature supplied is valid for whichever operators actually signed -- it makes no
// claim about what fraction of a quorum's total stake that represents. It is the caller's
// job to compare the returned QuorumStakeTotals against the requested
// quorumThresholdPercentage and reject responses that fall short.
//
// The original respondToFulfill() computed quorumStakeTotals but never used it: the
// quorumThresholdPercentage argument was accepted, recorded nowhere, and never checked.
// That means an aggregator response backed by an arbitrarily small slice of quorum stake
// (e.g. a single low-stake operator) was accepted and released the maker's escrowed
// tokens, regardless of the threshold requested -- the exact guarantee the DVN
// quorum-threshold parameter (tau) is supposed to provide in the thesis's Layer I design.
//
// testRespondToFulfill_LowStakeSignatureIsRejectedAt100PercentThreshold fails against the
// original contract and passes once the missing check (see src/OrderBook.sol) is added.
// testRespondToFulfill_HonestFullQuorumSucceeds guards against the fix breaking the
// legitimate full-quorum path.
contract OrderBookQuorumThresholdTest is BLSMockAVSDeployer, SignUtils {
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
    uint32 signChainId = 17000; // holesky, matches how OrderBook is deployed in Settlement.t.sol
    uint32 targetNetworkNumber = 11155111; // sepolia
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

        // Point the BLS mock harness's aggregate test signature at the hash of the
        // OrderResponse this "aggregator" will submit (instead of the harness's default
        // fixed "hello world" message), so the mock signature actually verifies against a
        // real response for this order.
        IOrderBook.OrderResponse memory orderResponse = _buildResponse();
        msgHash = keccak256(abi.encode(orderResponse));
        sigma = BN254.hashToG1(msgHash).scalar_mul(aggSignerPrivKey);
    }

    function _buildResponse() internal view returns (IOrderBook.OrderResponse memory) {
        return IOrderBook.OrderResponse({recipient: taker, referenceOrderIndex: orderId});
    }

    function testRespondToFulfill_HonestFullQuorumSucceeds() public {
        // All maxOperatorsToRegister (4) equal-stake operators sign: 100% of quorum stake.
        (uint32 referenceBlockNumber, BLSSignatureChecker.NonSignerStakesAndSignature memory nonSignerStakesAndSignature)
        = _registerSignatoriesAndGetNonSignerStakeAndSignatureRandom(1, 0, quorumBitmap);

        vm.prank(aggregator);
        orderBook.respondToFulfill(
            quorumNumbers, 100, referenceBlockNumber, _buildResponse(), nonSignerStakesAndSignature
        );

        assertEq(
            inputErc20.balanceOf(taker),
            inputAmount,
            "taker should receive the escrowed input tokens once the full quorum signs off"
        );
    }

    function testRespondToFulfill_LowStakeSignatureIsRejectedAt100PercentThreshold() public {
        // Only 1 of the 4 equal-stake mock operators actually signs (3 non-signers), i.e.
        // ~25% of quorum stake -- yet the aggregator requests a 100% threshold.
        (uint32 referenceBlockNumber, BLSSignatureChecker.NonSignerStakesAndSignature memory nonSignerStakesAndSignature)
        = _registerSignatoriesAndGetNonSignerStakeAndSignatureRandom(2, maxOperatorsToRegister - 1, quorumBitmap);

        vm.prank(aggregator);
        vm.expectRevert(bytes("OrderBook.respondToFulfill: signed stake fails quorum threshold check"));
        orderBook.respondToFulfill(
            quorumNumbers, 100, referenceBlockNumber, _buildResponse(), nonSignerStakesAndSignature
        );

        assertEq(
            inputErc20.balanceOf(taker), 0, "escrowed tokens must not move when the quorum threshold is not met"
        );
    }
}
