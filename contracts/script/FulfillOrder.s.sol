import "forge-std/Script.sol";
import {Utils} from "./utils/Utils.sol";
import "../src/Settlement.sol";
import {SignUtils} from "./utils/SignUtils.sol";


contract FulfillOrder is SignUtils, Utils {
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
        uint32 orderId = 1;
        // get current pk
        uint256 pk = vm.envUint("PRIVATE_KEY");
        address maker = vm.addr(pk);
        address taker = address(0);
        address inputToken = address(1);
        uint256 inputAmount = 100;
        address outputToken = address(2);
        uint256 outputAmount = 200;
        bytes memory quorumNumbers = new bytes(1);
        uint256 expiry = block.timestamp + 1000;
        uint32 targetNetworkNumber = 17000;
        bytes memory content = abi.encodePacked(orderId, maker, taker, inputToken, inputAmount, outputToken, outputAmount, expiry, targetNetworkNumber);
        bytes memory sig = signContent(pk, content);
        settlement.fulfill(Settlement.Fulfill(orderId, maker, taker, inputToken, inputAmount, outputToken, outputAmount, 100, quorumNumbers, expiry, targetNetworkNumber, sig));
        vm.stopBroadcast();
    }
}