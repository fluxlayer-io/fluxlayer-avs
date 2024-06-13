import "../src/Settlement.sol";
import "../lib/forge-std/src/Test.sol";
import "../lib/eigenlayer-middleware/test/utils/BLSMockAVSDeployer.sol";

contract SettlementTest is BLSMockAVSDeployer {
    Settlement settlement;

    function setUp() public {
        settlement = new Settlement(IRegistryCoordinator(address(0)));
    }

    function testFulfill() public {
        address inputToken = address(1);
        uint256 inputAmount = 100;
        address outputToken = address(2);
        uint256 outputAmount = 100;
        bytes memory quorumNumbers = new bytes(0);
        settlement.fulfill(inputToken, inputAmount, outputToken, outputAmount, 100, quorumNumbers);
    }
}
