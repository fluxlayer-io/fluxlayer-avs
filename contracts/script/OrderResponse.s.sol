pragma solidity ^0.8.12;

import "forge-std/Script.sol";
import {Utils} from "./utils/Utils.sol";
import "../src/OrderBook.sol";

contract OrderResponse is Script, Utils {
    OrderBook orderBook;

    function run() external {
        string memory avsDeploymentOutput = readOutput(
            "flux_layer_avs_deployment_output"
        );
        orderBook = OrderBook(
            stdJson.readAddress(
                avsDeploymentOutput,
                ".addresses.orderBook"
            )
        );

        bytes32 response = orderBook.allOrderResponses(1);
        console.logBytes32(response);
    }
}
