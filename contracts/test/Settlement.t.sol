import "../src/Settlement.sol";
import "../lib/forge-std/src/Test.sol";
import "../lib/eigenlayer-middleware/test/utils/BLSMockAVSDeployer.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "../script/utils/SignUtils.sol";

contract SettlementTest is BLSMockAVSDeployer, SignUtils {
    Settlement public settlement;
    address aggregator =
    address(uint160(uint256(keccak256(abi.encodePacked("aggregator")))));
    uint32 orderId = 1;
    uint256 makerPk = 1;
    uint256 fakeMakerPk = 2;
    address maker = vm.addr(makerPk);
    address taker = address(0);
    address inputToken = address(1);
    uint256 inputAmount = 100;
    address outputToken = address(2);
    uint256 outputAmount = 100;
    uint256 expiry = block.timestamp + 1;
    bytes quorumNumbers = new bytes(0);
    bytes sig;
    bytes invalidSig;

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
        sig = signOrder(makerPk);
        invalidSig = signOrder(fakeMakerPk);
    }

    function signOrder(uint256 pk) public returns (bytes memory) {
        bytes memory content = abi.encodePacked(orderId, maker, taker, inputToken, inputAmount, outputToken, outputAmount, expiry);
        return signContent(pk, content);
    }

    function testFulfillRevertWhenTakerIsNotMsgSender() public {
        // expect revert
        vm.expectRevert("taker is not msg sender");
        settlement.fulfill(Settlement.Fulfill(orderId, maker, address(4), inputToken, inputAmount, outputToken, outputAmount, 100, quorumNumbers, expiry, sig));
    }

    function testFulFillRevertWhenExpiry() public {
        vm.expectRevert("order is expired");
        settlement.fulfill(Settlement.Fulfill(orderId, maker, taker, inputToken, inputAmount, outputToken, outputAmount, 100, quorumNumbers, 0, sig));
    }

    function testFulFillRevertWhenInvalidSig() public {
        vm.expectRevert("invalid signature");
        settlement.fulfill(Settlement.Fulfill(orderId, maker, taker, inputToken, inputAmount, outputToken, outputAmount, 100, quorumNumbers, expiry, invalidSig));
    }

    function testFulFill() public {
        settlement.fulfill(Settlement.Fulfill(orderId, maker, taker, inputToken, inputAmount, outputToken, outputAmount, 100, quorumNumbers, expiry, sig));
    }
}
