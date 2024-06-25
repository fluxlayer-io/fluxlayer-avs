package config

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/urfave/cli"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/wallet"
	"github.com/Layr-Labs/eigensdk-go/chainio/txmgr"
	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	sdklogging "github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/signerv2"

	sdkutils "github.com/Layr-Labs/eigensdk-go/utils"
)

// Config contains all of the configuration information for a credible squaring aggregators and challengers.
// Operators use a separate config. (see config-files/operator.anvil.yaml)
type Config struct {
	EcdsaPrivateKey           *ecdsa.PrivateKey
	BlsPrivateKey             *bls.PrivateKey
	Logger                    sdklogging.Logger
	EigenMetricsIpPortAddress string
	// we need the url for the eigensdk currently... eventually standardize api so as to
	// only take an ethclient or an rpcUrl (and build the ethclient at each constructor site)
	EthHttpRpcUrl                    string
	EthWsRpcUrl                      string
	EthHttpRpcUrl2                   string
	EthWsRpcUrl2                     string
	EthHttpClient                    eth.Client
	EthWsClient                      eth.Client
	EthHttpClient2                   eth.Client
	EthWsClient2                     eth.Client
	OperatorStateRetrieverAddr       common.Address
	FluxLayerRegistryCoordinatorAddr common.Address
	OrderBookAddr                    common.Address
	SettlementAddr                   common.Address
	AggregatorServerIpPortAddr       string
	RegisterOperatorOnStartup        bool
	// json:"-" skips this field when marshaling (only used for logging to stdout), since SignerFn doesnt implement marshalJson
	SignerFn          signerv2.SignerFn `json:"-"`
	TxMgr             txmgr.TxManager
	AggregatorAddress common.Address
}

// These are read from ConfigFileFlag
type ConfigRaw struct {
	Environment                sdklogging.LogLevel `yaml:"environment"`
	EthRpcUrl                  string              `yaml:"eth_rpc_url"`
	EthWsUrl                   string              `yaml:"eth_ws_url"`
	EthRpcUrl2                 string              `yaml:"eth_rpc_url2"`
	EthWsUrl2                  string              `yaml:"eth_ws_url2"`
	AggregatorServerIpPortAddr string              `yaml:"aggregator_server_ip_port_address"`
	RegisterOperatorOnStartup  bool                `yaml:"register_operator_on_startup"`
}

// These are read from FluxLayerDeploymentFileFlag
type FluxLayerDeploymentRaw struct {
	Addresses FluxLayerContractsRaw `json:"addresses"`
}
type FluxLayerContractsRaw struct {
	RegistryCoordinatorAddr    string `json:"registryCoordinator"`
	OperatorStateRetrieverAddr string `json:"operatorStateRetriever"`
	OrderBookAddr              string `json:"orderBook"`
	SettlementAddr             string `json:"settlement"`
}

// NewConfig parses config file to read from from flags or environment variables
// Note: This config is shared by challenger and aggregator and so we put in the core.
// Operator has a different config and is meant to be used by the operator CLI.
func NewConfig(ctx *cli.Context) (*Config, error) {

	var configRaw ConfigRaw
	configFilePath := ctx.GlobalString(ConfigFileFlag.Name)
	if configFilePath != "" {
		sdkutils.ReadYamlConfig(configFilePath, &configRaw)
	}

	var fluxLayerDeploymentRaw FluxLayerDeploymentRaw
	fluxLayerDeploymentFilePath := ctx.GlobalString(FluxLayerDeploymentFileFlag.Name)
	if _, err := os.Stat(fluxLayerDeploymentFilePath); errors.Is(err, os.ErrNotExist) {
		panic("Path " + fluxLayerDeploymentFilePath + " does not exist")
	}
	sdkutils.ReadJsonConfig(fluxLayerDeploymentFilePath, &fluxLayerDeploymentRaw)

	logger, err := sdklogging.NewZapLogger(configRaw.Environment)
	if err != nil {
		return nil, err
	}

	ethRpcClient, err := eth.NewClient(configRaw.EthRpcUrl)
	if err != nil {
		logger.Errorf("Cannot create http ethclient", "err", err)
		return nil, err
	}

	ethWsClient, err := eth.NewClient(configRaw.EthWsUrl)
	if err != nil {
		logger.Errorf("Cannot create ws ethclient", "err", err)
		return nil, err
	}

	ethRpcClient2, err := eth.NewClient(configRaw.EthRpcUrl2)
	if err != nil {
		logger.Errorf("Cannot create http ethclient2", "err", err)
		return nil, err
	}

	ethWsClient2, err := eth.NewClient(configRaw.EthWsUrl2)
	if err != nil {
		logger.Errorf("Cannot create ws ethclient2", "err", err)
		return nil, err
	}

	ecdsaPrivateKeyString := ctx.GlobalString(EcdsaPrivateKeyFlag.Name)
	if ecdsaPrivateKeyString[:2] == "0x" {
		ecdsaPrivateKeyString = ecdsaPrivateKeyString[2:]
	}
	ecdsaPrivateKey, err := crypto.HexToECDSA(ecdsaPrivateKeyString)
	if err != nil {
		logger.Errorf("Cannot parse ecdsa private key", "err", err)
		return nil, err
	}

	aggregatorAddr, err := sdkutils.EcdsaPrivateKeyToAddress(ecdsaPrivateKey)
	if err != nil {
		logger.Error("Cannot get operator address", "err", err)
		return nil, err
	}

	chainId, err := ethRpcClient.ChainID(context.Background())
	if err != nil {
		logger.Error("Cannot get chainId", "err", err)
		return nil, err
	}

	signerV2, _, err := signerv2.SignerFromConfig(signerv2.Config{PrivateKey: ecdsaPrivateKey}, chainId)
	if err != nil {
		panic(err)
	}
	skWallet, err := wallet.NewPrivateKeyWallet(ethRpcClient, signerV2, aggregatorAddr, logger)
	if err != nil {
		panic(err)
	}
	// log sender
	logger.Info("txMgr Sender Address", "address", aggregatorAddr.String())
	txMgr := txmgr.NewSimpleTxManager(skWallet, ethRpcClient, logger, aggregatorAddr)

	config := &Config{
		EcdsaPrivateKey:                  ecdsaPrivateKey,
		Logger:                           logger,
		EthWsRpcUrl:                      configRaw.EthWsUrl,
		EthHttpRpcUrl:                    configRaw.EthRpcUrl,
		EthWsRpcUrl2:                     configRaw.EthWsUrl2,
		EthHttpRpcUrl2:                   configRaw.EthRpcUrl2,
		EthHttpClient:                    ethRpcClient,
		EthWsClient:                      ethWsClient,
		EthHttpClient2:                   ethRpcClient2,
		EthWsClient2:                     ethWsClient2,
		OperatorStateRetrieverAddr:       common.HexToAddress(fluxLayerDeploymentRaw.Addresses.OperatorStateRetrieverAddr),
		FluxLayerRegistryCoordinatorAddr: common.HexToAddress(fluxLayerDeploymentRaw.Addresses.RegistryCoordinatorAddr),
		OrderBookAddr:                    common.HexToAddress(fluxLayerDeploymentRaw.Addresses.OrderBookAddr),
		SettlementAddr:                   common.HexToAddress(fluxLayerDeploymentRaw.Addresses.SettlementAddr),
		AggregatorServerIpPortAddr:       configRaw.AggregatorServerIpPortAddr,
		RegisterOperatorOnStartup:        configRaw.RegisterOperatorOnStartup,
		SignerFn:                         signerV2,
		TxMgr:                            txMgr,
		AggregatorAddress:                aggregatorAddr,
	}
	config.validate()
	return config, nil
}

func (c *Config) validate() {
	// TODO: make sure every pointer is non-nil
	if c.OperatorStateRetrieverAddr == common.HexToAddress("") {
		panic("Config: BLSOperatorStateRetrieverAddr is required")
	}
	if c.FluxLayerRegistryCoordinatorAddr == common.HexToAddress("") {
		panic("Config: FluxLayerRegistryCoordinatorAddr is required")
	}
}

var (
	/* Required Flags */
	ConfigFileFlag = cli.StringFlag{
		Name:     "config",
		Required: true,
		Usage:    "Load configuration from `FILE`",
	}
	FluxLayerDeploymentFileFlag = cli.StringFlag{
		Name:     "flux-layer-deployment",
		Required: true,
		Usage:    "Load flux layer contract addresses from `FILE`",
	}
	EcdsaPrivateKeyFlag = cli.StringFlag{
		Name:     "ecdsa-private-key",
		Usage:    "Ethereum private key",
		Required: true,
		EnvVar:   "ECDSA_PRIVATE_KEY",
	}
	/* Optional Flags */
)

var requiredFlags = []cli.Flag{
	ConfigFileFlag,
	FluxLayerDeploymentFileFlag,
	EcdsaPrivateKeyFlag,
}

var optionalFlags = []cli.Flag{}

func init() {
	Flags = append(requiredFlags, optionalFlags...)
}

// Flags contains the list of configuration options available to the binary.
var Flags []cli.Flag
