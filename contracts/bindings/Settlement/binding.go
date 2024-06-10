// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractSettlement

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ContractSettlementMetaData contains all meta data concerning the ContractSettlement contract.
var ContractSettlementMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"fulfill\",\"inputs\":[{\"name\":\"inputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"outputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"outputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"fulfillEvent\",\"inputs\":[{\"name\":\"swapper\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"inputToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"inputAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"outputToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"outputAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x608060405234801561001057600080fd5b5061012b806100206000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c80635b7a7efe14602d575b600080fd5b603c603836600460b6565b603e565b005b604080513381526001600160a01b038681166020830152818301869052841660608201526080810183905290517f4c0264014410e86ffc99ec812823c6a5ac7ae71d8633ab05f5b5307ced070e619181900360a00190a150505050565b80356001600160a01b038116811460b157600080fd5b919050565b6000806000806080858703121560cb57600080fd5b60d285609b565b93506020850135925060e560408601609b565b939692955092936060013592505056fea2646970667358221220c587a55c2e83322f67ad171cf101cb159446129d50100069278b7b90e718476664736f6c634300080c0033",
}

// ContractSettlementABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractSettlementMetaData.ABI instead.
var ContractSettlementABI = ContractSettlementMetaData.ABI

// ContractSettlementBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractSettlementMetaData.Bin instead.
var ContractSettlementBin = ContractSettlementMetaData.Bin

// DeployContractSettlement deploys a new Ethereum contract, binding an instance of ContractSettlement to it.
func DeployContractSettlement(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ContractSettlement, error) {
	parsed, err := ContractSettlementMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractSettlementBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractSettlement{ContractSettlementCaller: ContractSettlementCaller{contract: contract}, ContractSettlementTransactor: ContractSettlementTransactor{contract: contract}, ContractSettlementFilterer: ContractSettlementFilterer{contract: contract}}, nil
}

// ContractSettlement is an auto generated Go binding around an Ethereum contract.
type ContractSettlement struct {
	ContractSettlementCaller     // Read-only binding to the contract
	ContractSettlementTransactor // Write-only binding to the contract
	ContractSettlementFilterer   // Log filterer for contract events
}

// ContractSettlementCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractSettlementCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSettlementTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractSettlementTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSettlementFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractSettlementFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSettlementSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSettlementSession struct {
	Contract     *ContractSettlement // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractSettlementCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractSettlementCallerSession struct {
	Contract *ContractSettlementCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ContractSettlementTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractSettlementTransactorSession struct {
	Contract     *ContractSettlementTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ContractSettlementRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractSettlementRaw struct {
	Contract *ContractSettlement // Generic contract binding to access the raw methods on
}

// ContractSettlementCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractSettlementCallerRaw struct {
	Contract *ContractSettlementCaller // Generic read-only contract binding to access the raw methods on
}

// ContractSettlementTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractSettlementTransactorRaw struct {
	Contract *ContractSettlementTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractSettlement creates a new instance of ContractSettlement, bound to a specific deployed contract.
func NewContractSettlement(address common.Address, backend bind.ContractBackend) (*ContractSettlement, error) {
	contract, err := bindContractSettlement(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractSettlement{ContractSettlementCaller: ContractSettlementCaller{contract: contract}, ContractSettlementTransactor: ContractSettlementTransactor{contract: contract}, ContractSettlementFilterer: ContractSettlementFilterer{contract: contract}}, nil
}

// NewContractSettlementCaller creates a new read-only instance of ContractSettlement, bound to a specific deployed contract.
func NewContractSettlementCaller(address common.Address, caller bind.ContractCaller) (*ContractSettlementCaller, error) {
	contract, err := bindContractSettlement(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractSettlementCaller{contract: contract}, nil
}

// NewContractSettlementTransactor creates a new write-only instance of ContractSettlement, bound to a specific deployed contract.
func NewContractSettlementTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractSettlementTransactor, error) {
	contract, err := bindContractSettlement(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractSettlementTransactor{contract: contract}, nil
}

// NewContractSettlementFilterer creates a new log filterer instance of ContractSettlement, bound to a specific deployed contract.
func NewContractSettlementFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractSettlementFilterer, error) {
	contract, err := bindContractSettlement(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractSettlementFilterer{contract: contract}, nil
}

// bindContractSettlement binds a generic wrapper to an already deployed contract.
func bindContractSettlement(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractSettlementMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractSettlement *ContractSettlementRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractSettlement.Contract.ContractSettlementCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractSettlement *ContractSettlementRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractSettlement.Contract.ContractSettlementTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractSettlement *ContractSettlementRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractSettlement.Contract.ContractSettlementTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractSettlement *ContractSettlementCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractSettlement.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractSettlement *ContractSettlementTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractSettlement.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractSettlement *ContractSettlementTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractSettlement.Contract.contract.Transact(opts, method, params...)
}

// Fulfill is a paid mutator transaction binding the contract method 0x5b7a7efe.
//
// Solidity: function fulfill(address inputToken, uint256 inputAmount, address outputToken, uint256 outputAmount) returns()
func (_ContractSettlement *ContractSettlementTransactor) Fulfill(opts *bind.TransactOpts, inputToken common.Address, inputAmount *big.Int, outputToken common.Address, outputAmount *big.Int) (*types.Transaction, error) {
	return _ContractSettlement.contract.Transact(opts, "fulfill", inputToken, inputAmount, outputToken, outputAmount)
}

// Fulfill is a paid mutator transaction binding the contract method 0x5b7a7efe.
//
// Solidity: function fulfill(address inputToken, uint256 inputAmount, address outputToken, uint256 outputAmount) returns()
func (_ContractSettlement *ContractSettlementSession) Fulfill(inputToken common.Address, inputAmount *big.Int, outputToken common.Address, outputAmount *big.Int) (*types.Transaction, error) {
	return _ContractSettlement.Contract.Fulfill(&_ContractSettlement.TransactOpts, inputToken, inputAmount, outputToken, outputAmount)
}

// Fulfill is a paid mutator transaction binding the contract method 0x5b7a7efe.
//
// Solidity: function fulfill(address inputToken, uint256 inputAmount, address outputToken, uint256 outputAmount) returns()
func (_ContractSettlement *ContractSettlementTransactorSession) Fulfill(inputToken common.Address, inputAmount *big.Int, outputToken common.Address, outputAmount *big.Int) (*types.Transaction, error) {
	return _ContractSettlement.Contract.Fulfill(&_ContractSettlement.TransactOpts, inputToken, inputAmount, outputToken, outputAmount)
}

// ContractSettlementFulfillEventIterator is returned from FilterFulfillEvent and is used to iterate over the raw logs and unpacked data for FulfillEvent events raised by the ContractSettlement contract.
type ContractSettlementFulfillEventIterator struct {
	Event *ContractSettlementFulfillEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractSettlementFulfillEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSettlementFulfillEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractSettlementFulfillEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractSettlementFulfillEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSettlementFulfillEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSettlementFulfillEvent represents a FulfillEvent event raised by the ContractSettlement contract.
type ContractSettlementFulfillEvent struct {
	Swapper      common.Address
	InputToken   common.Address
	InputAmount  *big.Int
	OutputToken  common.Address
	OutputAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFulfillEvent is a free log retrieval operation binding the contract event 0x4c0264014410e86ffc99ec812823c6a5ac7ae71d8633ab05f5b5307ced070e61.
//
// Solidity: event fulfillEvent(address swapper, address inputToken, uint256 inputAmount, address outputToken, uint256 outputAmount)
func (_ContractSettlement *ContractSettlementFilterer) FilterFulfillEvent(opts *bind.FilterOpts) (*ContractSettlementFulfillEventIterator, error) {

	logs, sub, err := _ContractSettlement.contract.FilterLogs(opts, "fulfillEvent")
	if err != nil {
		return nil, err
	}
	return &ContractSettlementFulfillEventIterator{contract: _ContractSettlement.contract, event: "fulfillEvent", logs: logs, sub: sub}, nil
}

// WatchFulfillEvent is a free log subscription operation binding the contract event 0x4c0264014410e86ffc99ec812823c6a5ac7ae71d8633ab05f5b5307ced070e61.
//
// Solidity: event fulfillEvent(address swapper, address inputToken, uint256 inputAmount, address outputToken, uint256 outputAmount)
func (_ContractSettlement *ContractSettlementFilterer) WatchFulfillEvent(opts *bind.WatchOpts, sink chan<- *ContractSettlementFulfillEvent) (event.Subscription, error) {

	logs, sub, err := _ContractSettlement.contract.WatchLogs(opts, "fulfillEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSettlementFulfillEvent)
				if err := _ContractSettlement.contract.UnpackLog(event, "fulfillEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFulfillEvent is a log parse operation binding the contract event 0x4c0264014410e86ffc99ec812823c6a5ac7ae71d8633ab05f5b5307ced070e61.
//
// Solidity: event fulfillEvent(address swapper, address inputToken, uint256 inputAmount, address outputToken, uint256 outputAmount)
func (_ContractSettlement *ContractSettlementFilterer) ParseFulfillEvent(log types.Log) (*ContractSettlementFulfillEvent, error) {
	event := new(ContractSettlementFulfillEvent)
	if err := _ContractSettlement.contract.UnpackLog(event, "fulfillEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
