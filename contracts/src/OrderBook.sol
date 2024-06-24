pragma solidity ^0.8.12;

import "@openzeppelin-upgrades/contracts/proxy/utils/Initializable.sol";
import "@openzeppelin-upgrades/contracts/access/OwnableUpgradeable.sol";
import "@eigenlayer/contracts/permissions/Pausable.sol";
import "../lib/eigenlayer-middleware/src/interfaces/IRegistryCoordinator.sol";
import {BLSSignatureChecker, IRegistryCoordinator} from "@eigenlayer-middleware/src/BLSSignatureChecker.sol";
import {OperatorStateRetriever} from "@eigenlayer-middleware/src/OperatorStateRetriever.sol";
import "@openzeppelin/contracts/utils/cryptography/SignatureChecker.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "./IOrderBook.sol";
import "./OrderEIP712.sol";


contract OrderBook is
IOrderBook,
Initializable,
OwnableUpgradeable,
Pausable,
BLSSignatureChecker,
OperatorStateRetriever,
OrderEIP712
{
    using BN254 for BN254.G1Point;
    /* STORAGE */
    // mapping of order indices to all order hashes
    // when an order is fulfilled, order hash is stored here,
    // and responses need to pass the actual order,
    // which is hashed onchain and checked against this mapping
    mapping(uint32 => bytes32) public allOrderHashes;
    mapping(uint32 => Order) public allOrderDetails;
    // mapping of order indices to hash of abi.encode(taskResponse, taskResponseMetadata)
    mapping(uint32 => bytes32) public allOrderResponses;

    /* CONSTANT */
    uint256 internal constant _THRESHOLD_DENOMINATOR = 100;
    address public aggregator;

    /* MODIFIERS */
    modifier onlyAggregator() {
        require(msg.sender == aggregator, "Aggregator must be the caller");
        _;
    }

    constructor(
        IRegistryCoordinator _registryCoordinator
    ) BLSSignatureChecker(_registryCoordinator) OrderEIP712("OrderBook", "1.0", address(0)){
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

    function createOrder(Order memory order, bytes calldata sig) public {
        // check if order already exists
        uint32 orderId = order.orderId;
        if (allOrderHashes[orderId] != bytes32(0)) revert OrderExists();
        if (order.taker != address(0) && order.taker != msg.sender) revert TakerMismatch();
        if (order.expiry < block.timestamp) revert OrderExpired();
        // check signature
        if (!verify(order, sig)) revert InvalidSignature();
        // transfer the tokens from maker to this contract
        IERC20(order.inputToken).transferFrom(order.maker, address(this), order.inputAmount);
        order.taker = msg.sender;
        allOrderHashes[orderId] = keccak256(abi.encode(order));
        allOrderDetails[orderId] = order;
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
        IERC20(order.inputToken).transferFrom(address(this), order.taker, order.inputAmount);
        // emitting event
        emit OrderRespondedEvent(orderResponse, orderResponseMetadata);
    }
}
