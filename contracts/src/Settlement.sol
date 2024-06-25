pragma solidity ^0.8.12;

import "./IOrderBook.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {OrderEIP712} from "./OrderEIP712.sol";


contract Settlement is IOrderBook, OrderEIP712 {
    mapping(uint32 => OrderExecution) public allOrderExecutions;
    constructor(
        address orderBookAddr
    ) OrderEIP712("OrderBook", "1.0", orderBookAddr){
    }

    function fulfill(Order memory order, uint32 quorumThresholdPercentage, bytes calldata quorumNumbers, bytes calldata sig) public {
        // check if order already fulfilled
        if (allOrderExecutions[order.orderId].createdBlock != 0) revert OrderFulfilled();
        // check chainid
        if (order.targetNetworkNumber != uint32(block.chainid)) revert InvalidChain();
        if (order.expiry < block.timestamp) revert OrderExpired();
        // check that the signer is the maker
        if (!verify(order, sig)) revert InvalidSignature();
        // transfer the tokens from taker to maker
        IERC20(order.outputToken).transferFrom(msg.sender, order.maker, order.outputAmount);
        allOrderExecutions[order.orderId] = OrderExecution(quorumThresholdPercentage, quorumNumbers, uint32(block.number));
        emit FulfillEvent(sig, order, quorumThresholdPercentage, quorumNumbers, msg.sender);
    }
}
