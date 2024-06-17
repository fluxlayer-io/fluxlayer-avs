import "forge-std/Script.sol";
import {Utils} from "./utils/Utils.sol";
import "../src/Settlement.sol";


contract FulfillOrder is Script, Utils {
    Settlement settlement;

    function run() external {
        string memory avsDeploymentOutput = readOutput(
            "flux_layer_avs_deployment_output"
        );
        settlement = Settlement(
            stdJson.readAddress(
                avsDeploymentOutput,
                ".addresses.settlement"
            )
        );

        vm.startBroadcast();
        bytes memory quorumNumbers = new bytes(1);
        settlement.fulfill(address(1), 100, address(2), 200, 100, quorumNumbers);
        vm.stopBroadcast();
    }
}