import "forge-std/Script.sol";
import "../../src/IOrderBook.sol";
import "./Utils.sol";
import "./EIP712Utils.sol";
import "../../src/OrderBook.sol";
import {Settlement} from "../../src/Settlement.sol";

contract SampleData is Utils {
    // get current pk
    uint32 orderId = 6;
    uint256 takerPk = vm.envUint("PRIVATE_KEY");
    address taker = vm.addr(takerPk);
    uint256 makerPk = vm.envUint("MAKER_PRIVATE_KEY");
    address maker = vm.addr(makerPk);
    uint256 expiry = 9999999999;
    uint32 targetNetworkNumber = 11155111;
    uint256 inputAmount = 0.1 ether;
    uint256 outputAmount = 0.1 ether;
    uint32 signChainId = 17000;
    address inputToken;
    address outputToken;
    OrderBook orderBook;
    Settlement settlement;
    EIP712Utils eip712Utils;

    constructor() {
        string memory avsDeploymentOutput = readOutput(
            "flux_layer_avs_deployment_output"
        );
        orderBook = OrderBook(
            stdJson.readAddress(
                avsDeploymentOutput,
                ".addresses.orderBook"
            )
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
        eip712Utils = new EIP712Utils("OrderBook", "1.0", signChainId, address(orderBook));
    }

    function createSampleOrder(uint32 orderId, address maker, address taker) public returns (IOrderBook.Order memory) {
        return IOrderBook.Order(
            orderId,
            maker,
            taker,
            inputToken,
            inputAmount,
            outputToken,
            outputAmount,
            expiry,
            targetNetworkNumber
        );
    }
}