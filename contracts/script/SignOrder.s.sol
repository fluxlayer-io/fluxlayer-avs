import "./utils/SampleData.sol";
import "./utils/SignUtils.sol";

contract SignOrder is SignUtils, SampleData {
    address defaultTaker = address(0);

    function run() external {
        IOrderBook.Order memory order = createSampleOrder(
            orderId,
            maker,
            defaultTaker
        );

        bytes memory sig = signHash(makerPk, eip712Utils.getTypedDataHash(order));
        console.log("ORDER_SIGNATURE");
        console.logBytes(sig);
    }
}