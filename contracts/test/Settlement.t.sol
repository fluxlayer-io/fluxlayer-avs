import "../src/Settlement.sol";
import "../lib/forge-std/src/Test.sol";

contract SettlementTest is Test {
    Settlement settlement;

    function setUp() public {
        settlement = new Settlement();
    }

    function testFulfill() public {
        address inputToken = address(1);
        uint256 inputAmount = 100;
        address outputToken = address(2);
        uint256 outputAmount = 100;
        settlement.fulfill(inputToken, inputAmount, outputToken, outputAmount);
    }
}
