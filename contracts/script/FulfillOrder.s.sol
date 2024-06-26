pragma solidity ^0.8.12;

import "../src/ERC20Mock.sol";
import "forge-std/Script.sol";
import {Utils} from "./utils/Utils.sol";
import "../src/OrderBook.sol";
import {SignUtils} from "./utils/SignUtils.sol";
import "./utils/EIP712Utils.sol";
import {Settlement} from "../src/Settlement.sol";
import "./utils/SampleData.sol";


contract FulfillOrder is SignUtils, Utils, SampleData {
    address defaultTaker = address(0);
    uint32 quorumThresholdPercentage = 100;
    bytes quorumNumbers = new bytes(1);

    function run() external {
        IOrderBook.Order memory order = createSampleOrder(
            orderId,
            maker,
            defaultTaker
        );

        bytes memory sig = signHash(makerPk, eip712Utils.getTypedDataHash(order));
        console.log("signature");
        console.logBytes(sig);

        vm.createSelectFork(vm.rpcUrl("holesky"));
        vm.startBroadcast(takerPk);
        orderBook.createOrder(order, sig);
        vm.stopBroadcast();

        vm.createSelectFork(vm.rpcUrl("sepolia"));
        vm.startBroadcast(takerPk);
        settlement.fulfill(order, quorumThresholdPercentage, quorumNumbers, sig);
        vm.stopBroadcast();
    }
}