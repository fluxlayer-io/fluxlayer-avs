contract Settlement {
    // fulfill event
    event fulfillEvent(address swapper, address inputToken, uint256 inputAmount, address outputToken, uint256 outputAmount);

    function fulfill(address inputToken,
        uint256 inputAmount,
        address outputToken,
        uint256 outputAmount) public {
        emit fulfillEvent(msg.sender, inputToken, inputAmount, outputToken, outputAmount);
    }
}