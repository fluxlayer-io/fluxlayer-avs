import "forge-std/Script.sol";
import {Utils} from "./utils/Utils.sol";
import "../src/Settlement.sol";

contract OrderResponse is Script, Utils {
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

        bytes32 response = settlement.allOrderResponses(settlement.latestOrderNum() - 1);
        console.logBytes32(response);
    }
}
