# Intro

Assets:
- Token A (Holesky)
- Token B (Sepolia)

Contracts:
- Order (Holesky)
- Settlement (Sepolia)

Scenario: bridging asset A from Holesky to asset B on Sepolia

## TODOs
Holesky:
- [x] As a maker, I can create a cross-chain swap order with my signature to swap token A from Holesky to token B in Sepolia through the UI.
- [x] Order would be created off-chain on AVS
- [x] As a taker, I can read the order and bring the maker's order on-chain so that only I can fulfil it; token A will be taken simultaneously from the maker to the contract.

Sepolia:
- [x] As a taker, I can fulfil the order with token B on the Settlement contract; token B will be sent to the maker's address

Holesky:
- [x] AVS verify the taker's fulfilment transaction and uploads proof to the Order contract; token A will be sent to the taker's address simultaneously. 

## Run
1. Run local node: `anvil --block-time 1 -f https://ethereum-holesky-rpc.publicnode.com`
2. Deploy contract: `cd contracts
forge script script/FluxLayerDeployer.s.sol:FluxLayerDeployer --rpc-url http://127.0.0.1:8545  --private-key ${PRIVATE_KEY} --legacy --broadcast`
3. Start aggregator: `make start-aggregator`
4. Update contract address and start operator: `make start-operator`
