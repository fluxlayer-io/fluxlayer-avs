pragma solidity ^0.8.12;

import "../src/ERC20Mock.sol";
import "forge-std/Script.sol";
import {Utils} from "./utils/Utils.sol";
import "../src/OrderBook.sol";
import {SignUtils} from "./utils/SignUtils.sol";
import "./utils/EIP712Utils.sol";
import "../src/Settlement.sol";

contract SettlementDeployer is SignUtils, Utils {
    uint256 public PK = vm.envUint("PRIVATE_KEY");
    Settlement settlement;

    function run() external {
        // read orderBook address
        string memory avsDeploymentOutput = readOutput(
            "flux_layer_avs_deployment_output"
        );
        OrderBook orderBook = OrderBook(
            stdJson.readAddress(
                avsDeploymentOutput,
                ".addresses.orderBook"
            )
        );
        settlement = new Settlement(address(orderBook));
        string memory deployed_addresses = "addresses";
        string memory deployed_addresses_output = vm.serializeAddress(
            deployed_addresses,
            "settlement",
            address(settlement)
        );
        // WRITE JSON DATA
        string memory parent_object = "parent object";
        // serialize all the data
        string memory finalJson = vm.serializeString(
            parent_object,
            deployed_addresses,
            deployed_addresses_output
        );
        writeOutput(finalJson, "flux_layer_avs_deployment_output");
    }
}
