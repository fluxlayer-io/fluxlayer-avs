import "../src/OrderBook.sol";
import "../src/IOrderBook.sol";
import "../src/ERC20Mock.sol";
import "forge-std/Test.sol";
import "../lib/eigenlayer-middleware/test/utils/BLSMockAVSDeployer.sol";
import {TransparentUpgradeableProxy} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "../script/utils/SignUtils.sol";
import "../script/utils/EIP712Utils.sol";
import "forge-std/console.sol";
import "../src/Settlement.sol";

contract SettlementTest is BLSMockAVSDeployer, SignUtils {
    OrderBook public orderBook;
    Settlement public settlement;
    EIP712Utils public eip712Utils;
    uint256 pk = vm.envUint("PRIVATE_KEY");
    uint32 orderId = 1;
    address taker = vm.addr(pk);
    address aggregator =
    address(uint160(uint256(keccak256(abi.encodePacked("aggregator")))));
    uint256 makerPk = vm.envUint("PRIVATE_KEY");
    uint256 fakeMakerPk = 2;
    address maker = vm.addr(makerPk);
    address inputToken;
    uint256 inputAmount = 100;
    address outputToken;
    uint256 outputAmount = 100;
    uint256 expiry = 9999999999;
    uint32 holesky = 17000;
    uint32 sepolia = 11155111;
    IOrderBook.Order order;
    uint32 quorumThresholdPercentage = 100;
    bytes quorumNumbers = new bytes(0);
    bytes sig;
    bytes invalidSig;

    ERC20Mock public inputErc20;
    ERC20Mock public outputErc20;
    uint32 signChainId = 17000;

    function setUp() public {
        _setUpBLSMockAVSDeployer();
        OrderBook orderBookImp = new OrderBook(IRegistryCoordinator(address(registryCoordinator)), signChainId);
        orderBook = OrderBook(
            address(
                new TransparentUpgradeableProxy(
                    address(orderBookImp),
                    address(proxyAdmin),
                    abi.encodeWithSelector(
                        orderBook.initialize.selector,
                        pauserRegistry,
                        registryCoordinatorOwner,
                        aggregator
                    )
                )
            )
        );
        eip712Utils = new EIP712Utils("OrderBook", "1.0", signChainId, address(orderBook));
        inputErc20 = new ERC20Mock();
        outputErc20 = new ERC20Mock();
        inputToken = address(inputErc20);
        outputToken = address(outputErc20);
        order = IOrderBook.Order(orderId, maker, address(0), inputToken, inputAmount, outputToken, outputAmount, expiry, sepolia);
        sig = signHash(makerPk, eip712Utils.getTypedDataHash(order));
        console.logBytes(sig);
        invalidSig = signHash(fakeMakerPk, eip712Utils.getTypedDataHash(order));
        // mint mock tokens to maker and taker
        inputErc20.mint(maker, inputAmount);
        outputErc20.mint(taker, outputAmount);
        // taker create order
        vm.startPrank(taker);
        orderBook.createOrder(order, sig);
        vm.stopPrank();
        vm.createSelectFork("sepolia_fork");
        settlement = new Settlement(address(orderBook), signChainId);
    }

    function testFulfillRevertWhenTakerIsNotMsgSender() public {
        // expect revert
        vm.expectRevert();
        // set wrong maker
        address wrongTaker = address(4);
        order.taker = wrongTaker;
        settlement.fulfill(order, quorumThresholdPercentage, quorumNumbers, sig);
    }

    function testFulFillRevertWhenExpiry() public {
        vm.expectRevert();
        // set invalid expiry
        uint32 invalidExpiry = 0;
        order.expiry = invalidExpiry;
        settlement.fulfill(order, quorumThresholdPercentage, quorumNumbers, sig);
    }

    function testFulFillRevertWhenInvalidSig() public {
        vm.expectRevert();
        settlement.fulfill(order, quorumThresholdPercentage, quorumNumbers, invalidSig);
    }

    function testFulFill() public {
        vm.prank(taker);
        settlement.fulfill(order, quorumThresholdPercentage, quorumNumbers, sig);
    }
}
