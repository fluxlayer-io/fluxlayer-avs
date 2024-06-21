import "../src/ERC20Mock.sol";
import "forge-std/Script.sol";
import {Utils} from "./utils/Utils.sol";
import "../src/Settlement.sol";
import {SignUtils} from "./utils/SignUtils.sol";


contract FulfillOrder is SignUtils, Utils {
    Settlement settlement;
    uint32 orderId = 1;
    // get current pk
    uint256 pk = vm.envUint("PRIVATE_KEY");
    address taker = vm.addr(pk);
    uint256 makerPk = 1;
    address maker = vm.addr(makerPk);
    address inputToken;
    uint256 inputAmount = 100;
    address outputToken;
    uint256 outputAmount = 200;
    bytes quorumNumbers = new bytes(1);
    uint256 expiry = block.timestamp + 1000;
    uint32 targetNetworkNumber = 17000;

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
        inputToken = address(
            stdJson.readAddress(
                avsDeploymentOutput,
                ".addresses.inputToken"
            )
        );
        outputToken = address(
            stdJson.readAddress(
                avsDeploymentOutput,
                ".addresses.outputToken"
            )
        );

        vm.startBroadcast(pk);
        // mint inputToken, outputToken to maker, taker
        ERC20Mock(inputToken).mint(maker, inputAmount);
        ERC20Mock(outputToken).mint(taker, outputAmount);
        bytes memory content = abi.encodePacked(orderId, maker, address(0), inputToken, inputAmount, outputToken, outputAmount, expiry, targetNetworkNumber);
        bytes memory sig = signContent(makerPk, content);
        settlement.fulfill(Settlement.Order(orderId, maker, address(0), inputToken, inputAmount, outputToken, outputAmount, 100, quorumNumbers, expiry, targetNetworkNumber, 0), sig);
        vm.stopBroadcast();
    }
}