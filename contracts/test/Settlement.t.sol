import "../src/Settlement.sol";
import "../src/ERC20Mock.sol";
import "forge-std/Test.sol";
import "../lib/eigenlayer-middleware/test/utils/BLSMockAVSDeployer.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "../script/utils/SignUtils.sol";

contract SettlementTest is BLSMockAVSDeployer, SignUtils {
    Settlement public settlement;
    uint256 pk = vm.envUint("PRIVATE_KEY");
    address taker = vm.addr(pk);
    address aggregator =
    address(uint160(uint256(keccak256(abi.encodePacked("aggregator")))));
    uint32 orderId = 1;
    uint256 makerPk = 1;
    uint256 fakeMakerPk = 2;
    address maker = vm.addr(makerPk);
    address inputToken;
    uint256 inputAmount = 100;
    address outputToken;
    uint256 outputAmount = 100;
    uint256 expiry = block.timestamp + 1;
    uint32 targetNetworkNumber = 17000;
    bytes quorumNumbers = new bytes(0);
    bytes sig;
    bytes invalidSig;

    ERC20Mock public inputErc20;
    ERC20Mock public outputErc20;

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
        inputErc20 = new ERC20Mock();
        outputErc20 = new ERC20Mock();
        inputToken = address(inputErc20);
        outputToken = address(outputErc20);
        sig = signOrder(makerPk);
        invalidSig = signOrder(fakeMakerPk);
        // mint mock tokens to maker and taker
        inputErc20.mint(maker, inputAmount);
        outputErc20.mint(taker, outputAmount);
    }

    function signOrder(uint256 pk) public returns (bytes memory) {
        bytes memory content = abi.encodePacked(orderId, maker, taker, inputToken, inputAmount, outputToken, outputAmount, expiry, targetNetworkNumber);
        return signContent(pk, content);
    }

    function testFulfillRevertWhenTakerIsNotMsgSender() public {
        // expect revert
        vm.expectRevert();
        address wrongTaker = address(4);
        settlement.fulfill(Settlement.Order(orderId, maker, wrongTaker, inputToken, inputAmount, outputToken, outputAmount, 100, quorumNumbers, expiry, targetNetworkNumber, 0), sig);
    }

    function testFulFillRevertWhenExpiry() public {
        vm.expectRevert();
        settlement.fulfill(Settlement.Order(orderId, maker, taker, inputToken, inputAmount, outputToken, outputAmount, 100, quorumNumbers, 0, targetNetworkNumber, 0), sig);
    }

    function testFulFillRevertWhenInvalidSig() public {
        vm.expectRevert();
        settlement.fulfill(Settlement.Order(orderId, maker, taker, inputToken, inputAmount, outputToken, outputAmount, 100, quorumNumbers, expiry, targetNetworkNumber, 0), invalidSig);
    }

    function testFulFill() public {
        vm.prank(taker);
        settlement.fulfill(Settlement.Order(orderId, maker, taker, inputToken, inputAmount, outputToken, outputAmount, 100, quorumNumbers, expiry, targetNetworkNumber, 0), sig);
    }
}
