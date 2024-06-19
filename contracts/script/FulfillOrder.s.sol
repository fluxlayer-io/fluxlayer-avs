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
        uint32 orderId = 1;
        address maker = address(3);
        address taker = address(0);
        uint256 expiry = 0;
        bytes memory sig = new bytes(0);
        settlement.fulfill(Settlement.Fulfill(orderId, maker, taker, address(1), 100, address(2), 200, 100, quorumNumbers, expiry, sig));
        vm.stopBroadcast();
    }
}