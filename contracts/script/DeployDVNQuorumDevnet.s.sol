// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

// Deploys a live, single-quorum devnet equivalent of the vendored BLSMockAVSDeployer /
// MockAVSDeployer Foundry test harness (lib/eigenlayer-middleware/test/utils/), broadcasting
// real transactions to a live Anvil RPC instead of using Foundry's in-test cheatcode-only
// deployment. This is the "registry-only, live two-chain" devnet: it deploys the same mock
// EigenLayer-core dependencies (DelegationMock, AVSDirectoryMock, PaymentCoordinatorMock) and
// the same real middleware registries (RegistryCoordinatorHarness, StakeRegistryHarness,
// BLSApkRegistryHarness, IndexRegistry) the vendored harness uses, and the same harness-only
// setBLSPublicKey/operator-share shortcuts it uses to register operators without a real staked
// EigenLayer core deployment -- but it does so via vm.startBroadcast, so every deployment and
// setup call here is a genuine on-chain transaction against the target RPC, not a Foundry-only
// simulation. Two components the vendored harness also deploys are intentionally omitted here
// because no code path this benchmark exercises (registration, checkSignatures, vault release)
// ever calls into them: the real (non-mock) Slasher/AVSDirectory/PaymentCoordinator proxies, and
// EigenPodManagerMock/StrategyManagerMock (stake here is set directly via DelegationMock.
// setOperatorShares, exactly as the vendored harness's own _setOperatorWeight helper does,
// never through a strategy manager).
import "forge-std/Script.sol";
import "forge-std/console2.sol";

import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import {PauserRegistry} from "eigenlayer-contracts/src/contracts/permissions/PauserRegistry.sol";
import {IStrategy} from "eigenlayer-contracts/src/contracts/interfaces/IStrategy.sol";
import {IPaymentCoordinator} from "eigenlayer-contracts/src/contracts/interfaces/IPaymentCoordinator.sol";
import {EmptyContract} from "eigenlayer-contracts/src/test/mocks/EmptyContract.sol";

import {BN254} from "@eigenlayer-middleware/src/libraries/BN254.sol";
import {OperatorStateRetriever} from "@eigenlayer-middleware/src/OperatorStateRetriever.sol";
import {RegistryCoordinator} from "@eigenlayer-middleware/src/RegistryCoordinator.sol";
import {RegistryCoordinatorHarness} from "@eigenlayer-middleware/test/harnesses/RegistryCoordinatorHarness.t.sol";
import {ServiceManagerMock} from "@eigenlayer-middleware/test/mocks/ServiceManagerMock.sol";
import {IndexRegistry} from "@eigenlayer-middleware/src/IndexRegistry.sol";
import {IStakeRegistry} from "@eigenlayer-middleware/src/interfaces/IStakeRegistry.sol";
import {IRegistryCoordinator} from "@eigenlayer-middleware/src/interfaces/IRegistryCoordinator.sol";
import {AVSDirectoryMock} from "@eigenlayer-middleware/test/mocks/AVSDirectoryMock.sol";
import {DelegationMock} from "@eigenlayer-middleware/test/mocks/DelegationMock.sol";
import {PaymentCoordinatorMock} from "@eigenlayer-middleware/test/mocks/PaymentCoordinatorMock.sol";
import {BLSApkRegistryHarness} from "@eigenlayer-middleware/test/harnesses/BLSApkRegistryHarness.sol";
import {StakeRegistryHarness} from "@eigenlayer-middleware/test/harnesses/StakeRegistryHarness.sol";

import {CrossChainLegVaultDVN} from "../src/CrossChainLegVaultDVN.sol";

contract DeployDVNQuorumDevnet is Script {
    uint256 public constant WEIGHTING_DIVISOR = 1e18;
    uint96 public constant DEFAULT_STAKE = 1 ether;

    ProxyAdmin public proxyAdmin;
    PauserRegistry public pauserRegistry;
    EmptyContract public emptyContract;

    DelegationMock public delegationMock;
    AVSDirectoryMock public avsDirectoryMock;
    PaymentCoordinatorMock public paymentCoordinatorMock;

    RegistryCoordinatorHarness public registryCoordinator;
    StakeRegistryHarness public stakeRegistry;
    BLSApkRegistryHarness public blsApkRegistry;
    IndexRegistry public indexRegistry;
    ServiceManagerMock public serviceManager;
    OperatorStateRetriever public operatorStateRetriever;

    CrossChainLegVaultDVN public vault;

    address public deployer;

    function run() external {
        uint256 deployerPk = vm.envUint("PRIVATE_KEY");
        deployer = vm.addr(deployerPk);
        string memory operatorsJsonPath = vm.envString("OPERATORS_JSON_PATH");
        string memory outputJsonPath = vm.envString("OUTPUT_JSON_PATH");
        string memory json = vm.readFile(operatorsJsonPath);
        uint256 numOperators = vm.parseJsonUint(json, ".numOperators");

        vm.startBroadcast(deployerPk);
        _deployMockCoreAndEmptyProxies();
        _deployRegistryImplementationsAndUpgrade();
        _registerOperators(json, numOperators);
        _initializeRegistryCoordinator();
        operatorStateRetriever = new OperatorStateRetriever();
        vault = new CrossChainLegVaultDVN(IRegistryCoordinator(address(registryCoordinator)));
        vm.stopBroadcast();

        _writeOutput(outputJsonPath);
    }

    function _deployMockCoreAndEmptyProxies() internal {
        emptyContract = new EmptyContract();

        address[] memory pausers = new address[](1);
        pausers[0] = deployer;
        pauserRegistry = new PauserRegistry(pausers, deployer);

        proxyAdmin = new ProxyAdmin();

        delegationMock = new DelegationMock();
        avsDirectoryMock = new AVSDirectoryMock();
        paymentCoordinatorMock = new PaymentCoordinatorMock();

        registryCoordinator = RegistryCoordinatorHarness(address(
            new TransparentUpgradeableProxy(address(emptyContract), address(proxyAdmin), "")
        ));
        stakeRegistry = StakeRegistryHarness(address(
            new TransparentUpgradeableProxy(address(emptyContract), address(proxyAdmin), "")
        ));
        indexRegistry = IndexRegistry(address(
            new TransparentUpgradeableProxy(address(emptyContract), address(proxyAdmin), "")
        ));
        blsApkRegistry = BLSApkRegistryHarness(address(
            new TransparentUpgradeableProxy(address(emptyContract), address(proxyAdmin), "")
        ));
        serviceManager = ServiceManagerMock(address(
            new TransparentUpgradeableProxy(address(emptyContract), address(proxyAdmin), "")
        ));
    }

    function _deployRegistryImplementationsAndUpgrade() internal {
        StakeRegistryHarness stakeRegistryImpl = new StakeRegistryHarness(
            IRegistryCoordinator(address(registryCoordinator)),
            delegationMock
        );
        proxyAdmin.upgrade(
            TransparentUpgradeableProxy(payable(address(stakeRegistry))),
            address(stakeRegistryImpl)
        );

        BLSApkRegistryHarness blsApkRegistryImpl = new BLSApkRegistryHarness(registryCoordinator);
        proxyAdmin.upgrade(
            TransparentUpgradeableProxy(payable(address(blsApkRegistry))),
            address(blsApkRegistryImpl)
        );

        IndexRegistry indexRegistryImpl = new IndexRegistry(registryCoordinator);
        proxyAdmin.upgrade(
            TransparentUpgradeableProxy(payable(address(indexRegistry))),
            address(indexRegistryImpl)
        );

        ServiceManagerMock serviceManagerImpl = new ServiceManagerMock(
            avsDirectoryMock,
            IPaymentCoordinator(address(paymentCoordinatorMock)),
            registryCoordinator,
            stakeRegistry
        );
        proxyAdmin.upgrade(
            TransparentUpgradeableProxy(payable(address(serviceManager))),
            address(serviceManagerImpl)
        );
        serviceManager.initialize({initialOwner: deployer});
    }

    // Registers each operator's BLS G1 pubkey via the harness shortcut, and sets their stake
    // directly via DelegationMock, exactly mirroring MockAVSDeployer._setOperatorWeight and
    // BLSMockAVSDeployer's own registration technique (setBLSPublicKey bypasses the normal
    // proof-of-possession check; the real registerOperator() call, sent later by Go using each
    // operator's own key, still runs unmodified).
    function _registerOperators(string memory json, uint256 numOperators) internal {
        for (uint256 i = 0; i < numOperators; i++) {
            string memory base = string.concat(".operators[", vm.toString(i), "]");
            address operatorAddr = vm.parseJsonAddress(json, string.concat(base, ".addr"));
            uint256 pkX = vm.parseJsonUint(json, string.concat(base, ".pubkeyX"));
            uint256 pkY = vm.parseJsonUint(json, string.concat(base, ".pubkeyY"));

            blsApkRegistry.setBLSPublicKey(operatorAddr, BN254.G1Point(pkX, pkY));
            delegationMock.setOperatorShares(operatorAddr, IStrategy(address(0)), DEFAULT_STAKE);

            (bool sent, ) = operatorAddr.call{value: 1 ether}("");
            require(sent, "funding operator failed");
        }
    }

    function _initializeRegistryCoordinator() internal {
        IRegistryCoordinator.OperatorSetParam[] memory operatorSetParams =
            new IRegistryCoordinator.OperatorSetParam[](1);
        operatorSetParams[0] = IRegistryCoordinator.OperatorSetParam({
            maxOperatorCount: 10,
            kickBIPsOfOperatorStake: 15000,
            kickBIPsOfTotalStake: 150
        });

        uint96[] memory minimumStakeForQuorum = new uint96[](1);
        minimumStakeForQuorum[0] = 1;

        IStakeRegistry.StrategyParams[][] memory quorumStrategies =
            new IStakeRegistry.StrategyParams[][](1);
        quorumStrategies[0] = new IStakeRegistry.StrategyParams[](1);
        quorumStrategies[0][0] = IStakeRegistry.StrategyParams(
            IStrategy(address(0)),
            uint96(WEIGHTING_DIVISOR)
        );

        RegistryCoordinatorHarness registryCoordinatorImpl = new RegistryCoordinatorHarness(
            serviceManager,
            stakeRegistry,
            blsApkRegistry,
            indexRegistry
        );
        proxyAdmin.upgradeAndCall(
            TransparentUpgradeableProxy(payable(address(registryCoordinator))),
            address(registryCoordinatorImpl),
            abi.encodeWithSelector(
                RegistryCoordinator.initialize.selector,
                deployer,
                deployer,
                deployer,
                pauserRegistry,
                uint256(0),
                operatorSetParams,
                minimumStakeForQuorum,
                quorumStrategies
            )
        );
    }

    function _writeOutput(string memory outputJsonPath) internal {
        string memory out = "out";
        vm.serializeAddress(out, "registryCoordinator", address(registryCoordinator));
        vm.serializeAddress(out, "stakeRegistry", address(stakeRegistry));
        vm.serializeAddress(out, "blsApkRegistry", address(blsApkRegistry));
        vm.serializeAddress(out, "indexRegistry", address(indexRegistry));
        vm.serializeAddress(out, "operatorStateRetriever", address(operatorStateRetriever));
        string memory finalJson = vm.serializeAddress(out, "vault", address(vault));
        vm.writeJson(finalJson, outputJsonPath);

        console2.log("Deployed vault at", address(vault));
        console2.log("Deployed registryCoordinator at", address(registryCoordinator));
    }
}
