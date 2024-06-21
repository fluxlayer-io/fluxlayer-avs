import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@eigenlayer/contracts/permissions/Pausable.sol";
import "../lib/eigenlayer-middleware/src/interfaces/IRegistryCoordinator.sol";
import {BLSSignatureChecker, IRegistryCoordinator} from "@eigenlayer-middleware/src/BLSSignatureChecker.sol";
import {OperatorStateRetriever} from "@eigenlayer-middleware/src/OperatorStateRetriever.sol";
import "@openzeppelin/contracts/utils/cryptography/SignatureChecker.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/utils/cryptography/draft-EIP712.sol";
import "./ISettlement.sol";


contract Settlement is
ISettlement,
Initializable,
OwnableUpgradeable,
Pausable,
BLSSignatureChecker,
OperatorStateRetriever,
EIP712
{
    using BN254 for BN254.G1Point;
    using ECDSA for bytes32;
    /* STORAGE */
    // mapping of order indices to all order hashes
    // when an order is fulfilled, order hash is stored here,
    // and responses need to pass the actual order,
    uint32 orderId;
    // which is hashed onchain and checked against this mapping
    mapping(uint32 => bytes32) public allOrderHashes;
    mapping(uint32 => Order) public allOrderDetails;
    mapping(uint32 => OrderExecution) public allOrderExecutions;
    // mapping of order indices to hash of abi.encode(taskResponse, taskResponseMetadata)
    mapping(uint32 => bytes32) public allOrderResponses;

    /* CONSTANT */
    bytes32 private constant _TYPE_HASH = keccak256("Order(address maker,address taker,address inputToken,uint256 inputAmount,address outputToken,uint256 outputAmount,uint256 expiry,uint32 targetNetworkNumber)");
    uint256 internal constant _THRESHOLD_DENOMINATOR = 100;
    address public aggregator;

    /* MODIFIERS */
    modifier onlyAggregator() {
        require(msg.sender == aggregator, "Aggregator must be the caller");
        _;
    }

    constructor(
        IRegistryCoordinator _registryCoordinator
    ) BLSSignatureChecker(_registryCoordinator) EIP712("Settlement", "1.0"){
    }

    function initialize(
        IPauserRegistry _pauserRegistry,
        address initialOwner,
        address _aggregator
    ) public initializer {
        _initializePauser(_pauserRegistry, UNPAUSE_ALL);
        _transferOwnership(initialOwner);
        aggregator = _aggregator;
    }

    function hashOrder(Order calldata order) public view returns (bytes32) {
        return _hashTypedDataV4(keccak256(abi.encode(
            _TYPE_HASH,
            order.maker,
            order.taker,
            order.inputToken,
            order.inputAmount,
            order.outputToken,
            order.outputAmount,
            order.expiry,
            order.targetNetworkNumber
        )));
    }

    function verify(Order calldata order, bytes calldata signature) public view returns (bool) {
        address signer = hashOrder(order).recover(signature);
        return signer == order.maker;
    }

    function fulfill(Order calldata order, uint32 quorumThresholdPercentage, bytes calldata quorumNumbers, bytes calldata sig) public {
        // check taker address
        if (order.taker != address(0) && msg.sender != order.taker) revert TakerMismatch();
        if (order.expiry < block.timestamp) revert OrderExpired();
        // check that the signer is the maker
        if (!verify(order, sig)) revert InvalidSignature();
        // transfer the tokens from maker to this contract
        IERC20(order.inputToken).transferFrom(order.maker, address(this), order.inputAmount);
        // transfer the tokens from taker to maker
        IERC20(order.outputToken).transferFrom(msg.sender, order.maker, order.outputAmount);
        allOrderExecutions[orderId] = OrderExecution(quorumThresholdPercentage, quorumNumbers, uint32(block.number));
        allOrderHashes[orderId] = keccak256(abi.encode(order));
        allOrderDetails[orderId] = order;
        emit FulfillEvent(sig, orderId, order, quorumThresholdPercentage, quorumNumbers, msg.sender);
        orderId += 1;
    }

    // NOTE: this function responds to existing tasks.
    function respondToFulfill(
        bytes calldata quorumNumbers,
        uint32 quorumThresholdPercentage,
        uint32 createdBlock,
        OrderResponse calldata orderResponse,
        NonSignerStakesAndSignature memory nonSignerStakesAndSignature
    ) external onlyAggregator {
        // some logical checks
        require(
            allOrderResponses[orderResponse.referenceOrderIndex] == bytes32(0),
            "Aggregator has already responded to the order"
        );
        /* CHECKING SIGNATURES & WHETHER THRESHOLD IS MET OR NOT */
        // calculate message which operators signed
        bytes32 message = keccak256(abi.encode(orderResponse));

        // check the BLS signature
        (
            QuorumStakeTotals memory quorumStakeTotals,
            bytes32 hashOfNonSigners
        ) = checkSignatures(
            message,
            quorumNumbers,
            createdBlock,
            nonSignerStakesAndSignature
        );

        OrderResponseMetadata memory orderResponseMetadata = OrderResponseMetadata(
            uint32(block.number),
            hashOfNonSigners
        );
        // updating the storage with order response
        allOrderResponses[orderResponse.referenceOrderIndex] = keccak256(
            abi.encode(orderResponse, orderResponseMetadata)
        );

        // get order
        Order memory order = allOrderDetails[orderResponse.referenceOrderIndex];
        IERC20(order.inputToken).transferFrom(address(this), orderResponse.recipient, order.inputAmount);
        // emitting event
        emit OrderRespondedEvent(orderResponse, orderResponseMetadata);
    }

}
