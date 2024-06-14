import "../src/Settlement.sol";
import "../lib/forge-std/src/Test.sol";
import "../lib/eigenlayer-middleware/test/utils/BLSMockAVSDeployer.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

contract SettlementTest is BLSMockAVSDeployer {
    Settlement public settlement;
    address aggregator =
    address(uint160(uint256(keccak256(abi.encodePacked("aggregator")))));

    function setUp() public {
        _setUpBLSMockAVSDeployer();
        Settlement settlementImp = new Settlement(IRegistryCoordinator(address(registryCoordinator)));
        settlement = Settlement(
            address(
                new TransparentUpgradeableProxy(
                    address(settlementImp),
                    address(proxyAdmin),
                    abi.encodeWithSelector(
                        settlement.initialize.selector,
                        pauserRegistry,
                        registryCoordinatorOwner,
                        aggregator
                    )
                )
            )
        );
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
