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
    uint32 _quorumThresholdPercentage = 100;
    bytes _quorumNumbers = new bytes(1);
    string _root = "./";

    string _path = string.concat(_root, "orders.json");
    string _json = vm.readFile(_path);
    bytes _data = vm.parseJson(_json);

    // STRUCTS
    struct Order {
        uint32 orderId;
        address from;
        address receiver;
        address sellToken;
        uint256 sellAmount;
        address buyToken;
        uint256 buyAmount;
        bool isFulfilled;
        bytes tx;
        uint256 validTo;
        uint32 targetNetworkNumber;
        bytes signature;
    }

    function run() external {
        console.log("here0?");
        console.log("_json: ", _json);
        // IOrderBook.Order[] memory orders = abi.decode(_data, (IOrderBook.Order[]));
        Order[] memory orders = abi.decode(_data, (Order[]));
        console.log("here1?");
        Order memory order = orders[0];
        console.log("order: ", order.orderId);
        console.log("there?");
        // IOrderBook.Order memory order = createSampleOrder(
        //     orderId,
        //     maker,
        //     defaultTaker
        // );

        // bytes memory sig = signHash(makerPk, eip712Utils.getTypedDataHash(order));
        // console.log("signature");
        // console.logBytes(sig);

        vm.createSelectFork(vm.rpcUrl("holesky"));
        vm.startBroadcast(makerPk);
        IOrderBook.Order memory submitOrder = IOrderBook.Order(
            order.orderId,
            order.from,
            order.receiver,
            order.buyToken,
            order.buyAmount,
            order.sellToken,
            order.sellAmount,
            order.validTo,
            order.targetNetworkNumber
        );
        orderBook.createOrder(submitOrder, order.signature);
        vm.stopBroadcast();

        vm.createSelectFork(vm.rpcUrl("sepolia"));
        vm.startBroadcast(takerPk);
        settlement.fulfill(submitOrder, _quorumThresholdPercentage, _quorumNumbers, order.signature);
        vm.stopBroadcast();
    }
}