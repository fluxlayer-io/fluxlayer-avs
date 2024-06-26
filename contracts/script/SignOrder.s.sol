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
        console.log(vm.addr(0x2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6));
    }
}