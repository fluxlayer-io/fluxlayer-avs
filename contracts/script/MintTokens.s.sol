import "./utils/SampleData.sol";

contract MintTokens is SampleData {
    function run() external {
        vm.createSelectFork(vm.rpcUrl("holesky"));
        ERC20Mock(inputToken).mint(maker, 100 ether);
        ERC20Mock(inputToken).mint(taker, 100 ether);
        vm.createSelectFork(vm.rpcUrl("sepolia"));
        ERC20Mock(outputToken).mint(maker, 100 ether);
        ERC20Mock(outputToken).mint(taker, 100 ether);
    }
}