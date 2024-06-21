import "../src/ERC20Mock.sol";
import "forge-std/Script.sol";
import {Utils} from "./utils/Utils.sol";
import "../src/Settlement.sol";
import {SignUtils} from "./utils/SignUtils.sol";
import "./utils/EIP712Utils.sol";


contract FulfillOrder is SignUtils, Utils {
    Settlement settlement;
    // get current pk
    uint256 pk = vm.envUint("PRIVATE_KEY");
    address sender = vm.addr(pk);
    address taker = address(0);
    uint256 makerPk = pk;
    address maker = vm.addr(pk);
    address inputToken;
    uint256 inputAmount = 100;
    address outputToken;
    uint256 outputAmount = 100;
    uint32 quorumThresholdPercentage = 100;
    bytes quorumNumbers = new bytes(1);
    uint256 expiry = block.timestamp + 1000;
    uint32 targetNetworkNumber = 17000;

    function run() external {
        console.log("current chain id", block.chainid);
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

        EIP712Utils eip712Utils = new EIP712Utils("Settlement", "1.0", address(settlement));
        ISettlement.Order memory order = ISettlement.Order(
            maker,
            taker,
            inputToken,
            inputAmount,
            outputToken,
            outputAmount,
            expiry,
            targetNetworkNumber
        );

        vm.startBroadcast(pk);
        // mint inputToken, outputToken to maker, taker
        ERC20Mock(inputToken).mint(maker, inputAmount);
        ERC20Mock(outputToken).mint(sender, outputAmount);
        bytes memory sig = signHash(makerPk, eip712Utils.getTypedDataHash(order));
        console.log("sig");
        console.logBytes(sig);
        settlement.fulfill(order, quorumThresholdPercentage, quorumNumbers, sig);
        vm.stopBroadcast();
    }
}