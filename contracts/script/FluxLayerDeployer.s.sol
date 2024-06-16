// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.9;

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

import "@eigenlayer/contracts/permissions/PauserRegistry.sol";
import {IDelegationManager} from "@eigenlayer/contracts/interfaces/IDelegationManager.sol";
import {IAVSDirectory} from "@eigenlayer/contracts/interfaces/IAVSDirectory.sol";
import {IStrategyManager, IStrategy} from "@eigenlayer/contracts/interfaces/IStrategyManager.sol";
import {ISlasher} from "@eigenlayer/contracts/interfaces/ISlasher.sol";
import {StrategyBaseTVLLimits} from "@eigenlayer/contracts/strategies/StrategyBaseTVLLimits.sol";
import "@eigenlayer/test/mocks/EmptyContract.sol";

import "@eigenlayer-middleware/src/RegistryCoordinator.sol" as regcoord;
import {IBLSApkRegistry, IIndexRegistry, IStakeRegistry} from "@eigenlayer-middleware/src/RegistryCoordinator.sol";
import {BLSApkRegistry} from "@eigenlayer-middleware/src/BLSApkRegistry.sol";
import {IndexRegistry} from "@eigenlayer-middleware/src/IndexRegistry.sol";
import {StakeRegistry} from "@eigenlayer-middleware/src/StakeRegistry.sol";
import "@eigenlayer-middleware/src/OperatorStateRetriever.sol";

import "../src/ERC20Mock.sol";

import {Utils} from "./utils/Utils.sol";

import "forge-std/Test.sol";
import "forge-std/Script.sol";
import "forge-std/StdJson.sol";
import "forge-std/console.sol";
import "../src/Settlement.sol";
import "../src/FluxLayerServiceManager.sol";

// # To deploy and verify our contract
// anvil --block-time 5 -f https://ethereum-holesky-rpc.publicnode.com
// forge script script/FluxLayerDeployer.s.sol:FluxLayerDeployer --rpc-url http://127.0.0.1:8545  --private-key 0x3b5d8eba232abf04faa6b3cdd0f72148ccd8b52ab4881aaf21a8bf378260e45a --legacy --broadcast
contract FluxLayerDeployer is Script, Utils {
    // DEPLOYMENT CONSTANTS
    uint256 public constant QUORUM_THRESHOLD_PERCENTAGE = 100;
    uint32 public constant TASK_RESPONSE_WINDOW_BLOCK = 30;
    uint32 public constant TASK_DURATION_BLOCKS = 0;
    // TODO: right now hardcoding these (this address is anvil's default address 9)
    address public constant AGGREGATOR_ADDR =
    0xa0Ee7A142d267C1f36714E4a8F75612F20a79720;

    // ERC20 and Strategy: we need to deploy this erc20, create a strategy for it, and whitelist this strategy in the strategymanager

    ERC20Mock public erc20Mock;
    StrategyBaseTVLLimits public erc20MockStrategy;

    // Credible Squaring contracts
    ProxyAdmin public fluxLayerProxyAdmin;
    PauserRegistry public fluxLayerPauserReg;

    regcoord.RegistryCoordinator public registryCoordinator;
    regcoord.IRegistryCoordinator public registryCoordinatorImplementation;

    IBLSApkRegistry public blsApkRegistry;
    IBLSApkRegistry public blsApkRegistryImplementation;

    IIndexRegistry public indexRegistry;
    IIndexRegistry public indexRegistryImplementation;

    IStakeRegistry public stakeRegistry;
    IStakeRegistry public stakeRegistryImplementation;

    OperatorStateRetriever public operatorStateRetriever;

    FluxLayerServiceManager public fluxLayerServiceManager;
    FluxLayerServiceManager public fluxLayerServiceManagerImplementation;

    Settlement public settlement;
    Settlement public settlementImplementation;

    function run() external {
        // Eigenlayer contracts
        string memory eigenlayerDeployedContracts = readOutput(
            "eigenlayer_deployment_output"
        );
        IStrategyManager strategyManager = IStrategyManager(
            stdJson.readAddress(
                eigenlayerDeployedContracts,
                ".addresses.strategyManager"
            )
        );
        IDelegationManager delegationManager = IDelegationManager(
            stdJson.readAddress(
                eigenlayerDeployedContracts,
                ".addresses.delegation"
            )
        );
        IAVSDirectory avsDirectory = IAVSDirectory(
            stdJson.readAddress(
                eigenlayerDeployedContracts,
                ".addresses.avsDirectory"
            )
        );
        ProxyAdmin eigenLayerProxyAdmin = ProxyAdmin(
            stdJson.readAddress(
                eigenlayerDeployedContracts,
                ".addresses.eigenLayerProxyAdmin"
            )
        );
        PauserRegistry eigenLayerPauserReg = PauserRegistry(
            stdJson.readAddress(
                eigenlayerDeployedContracts,
                ".addresses.eigenLayerPauserReg"
            )
        );
        StrategyBaseTVLLimits baseStrategyImplementation = StrategyBaseTVLLimits(
            stdJson.readAddress(
                eigenlayerDeployedContracts,
                ".addresses.baseStrategyImplementation"
            )
        );

        address credibleSquaringCommunityMultisig = msg.sender;
        address credibleSquaringPauser = msg.sender;

        vm.startBroadcast();
        _deployErc20AndStrategyAndWhitelistStrategy(
            eigenLayerProxyAdmin,
            eigenLayerPauserReg,
            baseStrategyImplementation,
            strategyManager
        );
        _deployCredibleSquaringContracts(
            delegationManager,
            avsDirectory,
            erc20MockStrategy,
            credibleSquaringCommunityMultisig,
            credibleSquaringPauser
        );
        vm.stopBroadcast();
    }

    function _deployErc20AndStrategyAndWhitelistStrategy(
        ProxyAdmin eigenLayerProxyAdmin,
        PauserRegistry eigenLayerPauserReg,
        StrategyBaseTVLLimits baseStrategyImplementation,
        IStrategyManager strategyManager
    ) internal {
        erc20Mock = new ERC20Mock();
        // TODO(samlaf): any reason why we are using the strategybase with tvl limits instead of just using strategybase?
        // the maxPerDeposit and maxDeposits below are just arbitrary values.
        erc20MockStrategy = StrategyBaseTVLLimits(
            address(
                new TransparentUpgradeableProxy(
                    address(baseStrategyImplementation),
                    address(eigenLayerProxyAdmin),
                    abi.encodeWithSelector(
                        StrategyBaseTVLLimits.initialize.selector,
                        1 ether, // maxPerDeposit
                        100 ether, // maxDeposits
                        IERC20(erc20Mock),
                        eigenLayerPauserReg
                    )
                )
            )
        );
        IStrategy[] memory strats = new IStrategy[](1);
        strats[0] = erc20MockStrategy;
        bool[] memory thirdPartyTransfersForbiddenValues = new bool[](1);
        thirdPartyTransfersForbiddenValues[0] = false;
        strategyManager.addStrategiesToDepositWhitelist(
            strats,
            thirdPartyTransfersForbiddenValues
        );
    }

    function _deployCredibleSquaringContracts(
        IDelegationManager delegationManager,
        IAVSDirectory avsDirectory,
        IStrategy strat,
        address fluxLayerCommunityMultisig,
        address fluxLayerPauser
    ) internal {
        // Adding this as a temporary fix to make the rest of the script work with a single strategy
        // since it was originally written to work with an array of strategies
        IStrategy[1] memory deployedStrategyArray = [strat];
        uint numStrategies = deployedStrategyArray.length;

        // deploy proxy admin for ability to upgrade proxy contracts
        fluxLayerProxyAdmin = new ProxyAdmin();

        // deploy pauser registry
        {
            address[] memory pausers = new address[](2);
            pausers[0] = fluxLayerPauser;
            pausers[1] = fluxLayerCommunityMultisig;
            fluxLayerPauserReg = new PauserRegistry(
                pausers,
                fluxLayerCommunityMultisig
            );
        }

        EmptyContract emptyContract = new EmptyContract();

        // hard-coded inputs

        /**
         * First, deploy upgradeable proxy contracts that **will point** to the implementations. Since the implementation contracts are
         * not yet deployed, we give these proxies an empty contract as the initial implementation, to act as if they have no code.
         */
        fluxLayerServiceManager = FluxLayerServiceManager(
            address(
                new TransparentUpgradeableProxy(
                    address(emptyContract),
                    address(fluxLayerProxyAdmin),
                    ""
                )
            )
        );
        settlement = Settlement(
            address(
                new TransparentUpgradeableProxy(
                    address(emptyContract),
                    address(fluxLayerProxyAdmin),
                    ""
                )
            )
        );
        registryCoordinator = regcoord.RegistryCoordinator(
            address(
                new TransparentUpgradeableProxy(
                    address(emptyContract),
                    address(fluxLayerProxyAdmin),
                    ""
                )
            )
        );
        blsApkRegistry = IBLSApkRegistry(
            address(
                new TransparentUpgradeableProxy(
                    address(emptyContract),
                    address(fluxLayerProxyAdmin),
                    ""
                )
            )
        );
        indexRegistry = IIndexRegistry(
            address(
                new TransparentUpgradeableProxy(
                    address(emptyContract),
                    address(fluxLayerProxyAdmin),
                    ""
                )
            )
        );
        stakeRegistry = IStakeRegistry(
            address(
                new TransparentUpgradeableProxy(
                    address(emptyContract),
                    address(fluxLayerProxyAdmin),
                    ""
                )
            )
        );

        operatorStateRetriever = new OperatorStateRetriever();

        // Second, deploy the *implementation* contracts, using the *proxy contracts* as inputs
        {
            stakeRegistryImplementation = new StakeRegistry(
                registryCoordinator,
                delegationManager
            );

            fluxLayerProxyAdmin.upgrade(
                TransparentUpgradeableProxy(payable(address(stakeRegistry))),
                address(stakeRegistryImplementation)
            );

            blsApkRegistryImplementation = new BLSApkRegistry(
                registryCoordinator
            );

            fluxLayerProxyAdmin.upgrade(
                TransparentUpgradeableProxy(payable(address(blsApkRegistry))),
                address(blsApkRegistryImplementation)
            );

            indexRegistryImplementation = new IndexRegistry(
                registryCoordinator
            );

            fluxLayerProxyAdmin.upgrade(
                TransparentUpgradeableProxy(payable(address(indexRegistry))),
                address(indexRegistryImplementation)
            );
        }

        registryCoordinatorImplementation = new regcoord.RegistryCoordinator(
            fluxLayerServiceManager,
            regcoord.IStakeRegistry(address(stakeRegistry)),
            regcoord.IBLSApkRegistry(address(blsApkRegistry)),
            regcoord.IIndexRegistry(address(indexRegistry))
        );

        {
            uint numQuorums = 1;
            // for each quorum to setup, we need to define
            // QuorumOperatorSetParam, minimumStakeForQuorum, and strategyParams
            regcoord.IRegistryCoordinator.OperatorSetParam[]
            memory quorumsOperatorSetParams = new regcoord.IRegistryCoordinator.OperatorSetParam[](
                numQuorums
            );
            for (uint i = 0; i < numQuorums; i++) {
                // hard code these for now
                quorumsOperatorSetParams[i] = regcoord
                    .IRegistryCoordinator
                    .OperatorSetParam({
                    maxOperatorCount: 10000,
                    kickBIPsOfOperatorStake: 15000,
                    kickBIPsOfTotalStake: 100
                });
            }
            // set to 0 for every quorum
            uint96[] memory quorumsMinimumStake = new uint96[](numQuorums);
            IStakeRegistry.StrategyParams[][]
            memory quorumsStrategyParams = new IStakeRegistry.StrategyParams[][](
                numQuorums
            );
            for (uint i = 0; i < numQuorums; i++) {
                quorumsStrategyParams[i] = new IStakeRegistry.StrategyParams[](
                    numStrategies
                );
                for (uint j = 0; j < numStrategies; j++) {
                    quorumsStrategyParams[i][j] = IStakeRegistry
                        .StrategyParams({
                        strategy: deployedStrategyArray[j],
                    // setting this to 1 ether since the divisor is also 1 ether
                    // therefore this allows an operator to register with even just 1 token
                    // see https://github.com/Layr-Labs/eigenlayer-middleware/blob/m2-mainnet/src/StakeRegistry.sol#L484
                    //    weight += uint96(sharesAmount * strategyAndMultiplier.multiplier / WEIGHTING_DIVISOR);
                        multiplier: 1 ether
                    });
                }
            }
            fluxLayerProxyAdmin.upgradeAndCall(
                TransparentUpgradeableProxy(
                    payable(address(registryCoordinator))
                ),
                address(registryCoordinatorImplementation),
                abi.encodeWithSelector(
                    regcoord.RegistryCoordinator.initialize.selector,
                    // we set churnApprover and ejector to communityMultisig because we don't need them
                    fluxLayerCommunityMultisig,
                    fluxLayerCommunityMultisig,
                    fluxLayerCommunityMultisig,
                    fluxLayerPauserReg,
                    0, // 0 initialPausedStatus means everything unpaused
                    quorumsOperatorSetParams,
                    quorumsMinimumStake,
                    quorumsStrategyParams
                )
            );
        }

        fluxLayerServiceManagerImplementation = new FluxLayerServiceManager(
            avsDirectory,
            registryCoordinator,
            stakeRegistry,
            settlement
        );
        // Third, upgrade the proxy contracts to use the correct implementation contracts and initialize them.
        fluxLayerProxyAdmin.upgrade(
            TransparentUpgradeableProxy(
                payable(address(fluxLayerServiceManager))
            ),
            address(fluxLayerServiceManagerImplementation)
        );

        settlementImplementation = new Settlement(registryCoordinator);
        // Third, upgrade the proxy contracts to use the correct implementation contracts and initialize them.
        fluxLayerProxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(
                payable(address(settlement))
            ),
            address(settlementImplementation),
            abi.encodeWithSelector(
                settlement.initialize.selector,
                fluxLayerPauserReg,
                fluxLayerCommunityMultisig,
                AGGREGATOR_ADDR
            )
        );

        // WRITE JSON DATA
        string memory parent_object = "parent object";

        string memory deployed_addresses = "addresses";
        vm.serializeAddress(
            deployed_addresses,
            "erc20Mock",
            address(erc20Mock)
        );
        vm.serializeAddress(
            deployed_addresses,
            "erc20MockStrategy",
            address(erc20MockStrategy)
        );
        vm.serializeAddress(
            deployed_addresses,
            "fluxLayerServiceManager",
            address(fluxLayerServiceManager)
        );
        vm.serializeAddress(
            deployed_addresses,
            "fluxLayerServiceManagerImplementation",
            address(fluxLayerServiceManagerImplementation)
        );
        vm.serializeAddress(
            deployed_addresses,
            "registryCoordinator",
            address(registryCoordinator)
        );
        vm.serializeAddress(
            deployed_addresses,
            "registryCoordinatorImplementation",
            address(registryCoordinatorImplementation)
        );
        // settlement
        vm.serializeAddress(
            deployed_addresses,
            "settlement",
            address(settlement)
        );
        vm.serializeAddress(
            deployed_addresses,
            "settlementImplementation",
            address(settlementImplementation)
        );
        string memory deployed_addresses_output = vm.serializeAddress(
            deployed_addresses,
            "operatorStateRetriever",
            address(operatorStateRetriever)
        );

        // serialize all the data
        string memory finalJson = vm.serializeString(
            parent_object,
            deployed_addresses,
            deployed_addresses_output
        );

        writeOutput(finalJson, "flux_layer_avs_deployment_output");
    }
}
