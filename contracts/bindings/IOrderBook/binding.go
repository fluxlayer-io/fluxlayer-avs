// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractIOrderBook

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

// IOrderBookOrder is an auto generated low-level Go binding around an user-defined struct.
type IOrderBookOrder struct {
	OrderId             uint32
	Maker               common.Address
	Taker               common.Address
	InputToken          common.Address
	InputAmount         *big.Int
	OutputToken         common.Address
	OutputAmount        *big.Int
	Expiry              *big.Int
	TargetNetworkNumber uint32
}

// IOrderBookOrderResponse is an auto generated low-level Go binding around an user-defined struct.
type IOrderBookOrderResponse struct {
	Recipient           common.Address
	ReferenceOrderIndex uint32
}

// IOrderBookOrderResponseMetadata is an auto generated low-level Go binding around an user-defined struct.
type IOrderBookOrderResponseMetadata struct {
	OrderResponsedBlock uint32
	HashOfNonSigners    [32]byte
}

// ContractIOrderBookMetaData contains all meta data concerning the ContractIOrderBook contract.
var ContractIOrderBookMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"event\",\"name\":\"FulfillEvent\",\"inputs\":[{\"name\":\"sig\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"order\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIOrderBook.Order\",\"components\":[{\"name\":\"orderId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"taker\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"outputToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"outputAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"expiry\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"targetNetworkNumber\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"quorumThresholdPercentage\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"quorumNumbers\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OrderRespondedEvent\",\"inputs\":[{\"name\":\"orderResponse\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIOrderBook.OrderResponse\",\"components\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"referenceOrderIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"OrderResponseMetadata\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structIOrderBook.OrderResponseMetadata\",\"components\":[{\"name\":\"orderResponsedBlock\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"hashOfNonSigners\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidChain\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOrder\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OrderExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OrderExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OrderFulfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TakerMismatch\",\"inputs\":[]}]",
	Bin: "0x6080604052348015600f57600080fd5b50603f80601d6000396000f3fe6080604052600080fdfea2646970667358221220a768aa325d24741f01c9c8880d01a3dd6559139994d69ec16d98339d3b64da8864736f6c634300080c0033",
}

// ContractIOrderBookABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractIOrderBookMetaData.ABI instead.
var ContractIOrderBookABI = ContractIOrderBookMetaData.ABI

// ContractIOrderBookBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractIOrderBookMetaData.Bin instead.
var ContractIOrderBookBin = ContractIOrderBookMetaData.Bin

// DeployContractIOrderBook deploys a new Ethereum contract, binding an instance of ContractIOrderBook to it.
func DeployContractIOrderBook(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ContractIOrderBook, error) {
	parsed, err := ContractIOrderBookMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractIOrderBookBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractIOrderBook{ContractIOrderBookCaller: ContractIOrderBookCaller{contract: contract}, ContractIOrderBookTransactor: ContractIOrderBookTransactor{contract: contract}, ContractIOrderBookFilterer: ContractIOrderBookFilterer{contract: contract}}, nil
}

// ContractIOrderBook is an auto generated Go binding around an Ethereum contract.
type ContractIOrderBook struct {
	ContractIOrderBookCaller     // Read-only binding to the contract
	ContractIOrderBookTransactor // Write-only binding to the contract
	ContractIOrderBookFilterer   // Log filterer for contract events
}

// ContractIOrderBookCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractIOrderBookCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIOrderBookTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractIOrderBookTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIOrderBookFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractIOrderBookFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractIOrderBookSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractIOrderBookSession struct {
	Contract     *ContractIOrderBook // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractIOrderBookCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractIOrderBookCallerSession struct {
	Contract *ContractIOrderBookCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ContractIOrderBookTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractIOrderBookTransactorSession struct {
	Contract     *ContractIOrderBookTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ContractIOrderBookRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractIOrderBookRaw struct {
	Contract *ContractIOrderBook // Generic contract binding to access the raw methods on
}

// ContractIOrderBookCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractIOrderBookCallerRaw struct {
	Contract *ContractIOrderBookCaller // Generic read-only contract binding to access the raw methods on
}

// ContractIOrderBookTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractIOrderBookTransactorRaw struct {
	Contract *ContractIOrderBookTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractIOrderBook creates a new instance of ContractIOrderBook, bound to a specific deployed contract.
func NewContractIOrderBook(address common.Address, backend bind.ContractBackend) (*ContractIOrderBook, error) {
	contract, err := bindContractIOrderBook(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractIOrderBook{ContractIOrderBookCaller: ContractIOrderBookCaller{contract: contract}, ContractIOrderBookTransactor: ContractIOrderBookTransactor{contract: contract}, ContractIOrderBookFilterer: ContractIOrderBookFilterer{contract: contract}}, nil
}

// NewContractIOrderBookCaller creates a new read-only instance of ContractIOrderBook, bound to a specific deployed contract.
func NewContractIOrderBookCaller(address common.Address, caller bind.ContractCaller) (*ContractIOrderBookCaller, error) {
	contract, err := bindContractIOrderBook(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractIOrderBookCaller{contract: contract}, nil
}

// NewContractIOrderBookTransactor creates a new write-only instance of ContractIOrderBook, bound to a specific deployed contract.
func NewContractIOrderBookTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractIOrderBookTransactor, error) {
	contract, err := bindContractIOrderBook(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractIOrderBookTransactor{contract: contract}, nil
}

// NewContractIOrderBookFilterer creates a new log filterer instance of ContractIOrderBook, bound to a specific deployed contract.
func NewContractIOrderBookFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractIOrderBookFilterer, error) {
	contract, err := bindContractIOrderBook(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractIOrderBookFilterer{contract: contract}, nil
}

// bindContractIOrderBook binds a generic wrapper to an already deployed contract.
func bindContractIOrderBook(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractIOrderBookMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractIOrderBook *ContractIOrderBookRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractIOrderBook.Contract.ContractIOrderBookCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractIOrderBook *ContractIOrderBookRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIOrderBook.Contract.ContractIOrderBookTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractIOrderBook *ContractIOrderBookRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractIOrderBook.Contract.ContractIOrderBookTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractIOrderBook *ContractIOrderBookCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractIOrderBook.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractIOrderBook *ContractIOrderBookTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractIOrderBook.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractIOrderBook *ContractIOrderBookTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractIOrderBook.Contract.contract.Transact(opts, method, params...)
}

// ContractIOrderBookFulfillEventIterator is returned from FilterFulfillEvent and is used to iterate over the raw logs and unpacked data for FulfillEvent events raised by the ContractIOrderBook contract.
type ContractIOrderBookFulfillEventIterator struct {
	Event *ContractIOrderBookFulfillEvent // Event containing the contract specifics and raw log

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
func (it *ContractIOrderBookFulfillEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIOrderBookFulfillEvent)
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
		it.Event = new(ContractIOrderBookFulfillEvent)
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
func (it *ContractIOrderBookFulfillEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIOrderBookFulfillEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIOrderBookFulfillEvent represents a FulfillEvent event raised by the ContractIOrderBook contract.
type ContractIOrderBookFulfillEvent struct {
	Sig                       []byte
	Order                     IOrderBookOrder
	QuorumThresholdPercentage uint32
	QuorumNumbers             []byte
	Recipient                 common.Address
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterFulfillEvent is a free log retrieval operation binding the contract event 0x712dd31697b327ff70b0442a703553fe27c70cce6ec06c292a9cbba1ada6d780.
//
// Solidity: event FulfillEvent(bytes sig, (uint32,address,address,address,uint256,address,uint256,uint256,uint32) order, uint32 quorumThresholdPercentage, bytes quorumNumbers, address recipient)
func (_ContractIOrderBook *ContractIOrderBookFilterer) FilterFulfillEvent(opts *bind.FilterOpts) (*ContractIOrderBookFulfillEventIterator, error) {

	logs, sub, err := _ContractIOrderBook.contract.FilterLogs(opts, "FulfillEvent")
	if err != nil {
		return nil, err
	}
	return &ContractIOrderBookFulfillEventIterator{contract: _ContractIOrderBook.contract, event: "FulfillEvent", logs: logs, sub: sub}, nil
}

// WatchFulfillEvent is a free log subscription operation binding the contract event 0x712dd31697b327ff70b0442a703553fe27c70cce6ec06c292a9cbba1ada6d780.
//
// Solidity: event FulfillEvent(bytes sig, (uint32,address,address,address,uint256,address,uint256,uint256,uint32) order, uint32 quorumThresholdPercentage, bytes quorumNumbers, address recipient)
func (_ContractIOrderBook *ContractIOrderBookFilterer) WatchFulfillEvent(opts *bind.WatchOpts, sink chan<- *ContractIOrderBookFulfillEvent) (event.Subscription, error) {

	logs, sub, err := _ContractIOrderBook.contract.WatchLogs(opts, "FulfillEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIOrderBookFulfillEvent)
				if err := _ContractIOrderBook.contract.UnpackLog(event, "FulfillEvent", log); err != nil {
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

// ParseFulfillEvent is a log parse operation binding the contract event 0x712dd31697b327ff70b0442a703553fe27c70cce6ec06c292a9cbba1ada6d780.
//
// Solidity: event FulfillEvent(bytes sig, (uint32,address,address,address,uint256,address,uint256,uint256,uint32) order, uint32 quorumThresholdPercentage, bytes quorumNumbers, address recipient)
func (_ContractIOrderBook *ContractIOrderBookFilterer) ParseFulfillEvent(log types.Log) (*ContractIOrderBookFulfillEvent, error) {
	event := new(ContractIOrderBookFulfillEvent)
	if err := _ContractIOrderBook.contract.UnpackLog(event, "FulfillEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractIOrderBookOrderRespondedEventIterator is returned from FilterOrderRespondedEvent and is used to iterate over the raw logs and unpacked data for OrderRespondedEvent events raised by the ContractIOrderBook contract.
type ContractIOrderBookOrderRespondedEventIterator struct {
	Event *ContractIOrderBookOrderRespondedEvent // Event containing the contract specifics and raw log

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
func (it *ContractIOrderBookOrderRespondedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractIOrderBookOrderRespondedEvent)
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
		it.Event = new(ContractIOrderBookOrderRespondedEvent)
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
func (it *ContractIOrderBookOrderRespondedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractIOrderBookOrderRespondedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractIOrderBookOrderRespondedEvent represents a OrderRespondedEvent event raised by the ContractIOrderBook contract.
type ContractIOrderBookOrderRespondedEvent struct {
	OrderResponse         IOrderBookOrderResponse
	OrderResponseMetadata IOrderBookOrderResponseMetadata
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterOrderRespondedEvent is a free log retrieval operation binding the contract event 0xa511545e83c0068efbd19fea5ef3009b8471cf6ceb74e8e9092300fe380655fd.
//
// Solidity: event OrderRespondedEvent((address,uint32) orderResponse, (uint32,bytes32) OrderResponseMetadata)
func (_ContractIOrderBook *ContractIOrderBookFilterer) FilterOrderRespondedEvent(opts *bind.FilterOpts) (*ContractIOrderBookOrderRespondedEventIterator, error) {

	logs, sub, err := _ContractIOrderBook.contract.FilterLogs(opts, "OrderRespondedEvent")
	if err != nil {
		return nil, err
	}
	return &ContractIOrderBookOrderRespondedEventIterator{contract: _ContractIOrderBook.contract, event: "OrderRespondedEvent", logs: logs, sub: sub}, nil
}

// WatchOrderRespondedEvent is a free log subscription operation binding the contract event 0xa511545e83c0068efbd19fea5ef3009b8471cf6ceb74e8e9092300fe380655fd.
//
// Solidity: event OrderRespondedEvent((address,uint32) orderResponse, (uint32,bytes32) OrderResponseMetadata)
func (_ContractIOrderBook *ContractIOrderBookFilterer) WatchOrderRespondedEvent(opts *bind.WatchOpts, sink chan<- *ContractIOrderBookOrderRespondedEvent) (event.Subscription, error) {

	logs, sub, err := _ContractIOrderBook.contract.WatchLogs(opts, "OrderRespondedEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractIOrderBookOrderRespondedEvent)
				if err := _ContractIOrderBook.contract.UnpackLog(event, "OrderRespondedEvent", log); err != nil {
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

// ParseOrderRespondedEvent is a log parse operation binding the contract event 0xa511545e83c0068efbd19fea5ef3009b8471cf6ceb74e8e9092300fe380655fd.
//
// Solidity: event OrderRespondedEvent((address,uint32) orderResponse, (uint32,bytes32) OrderResponseMetadata)
func (_ContractIOrderBook *ContractIOrderBookFilterer) ParseOrderRespondedEvent(log types.Log) (*ContractIOrderBookOrderRespondedEvent, error) {
	event := new(ContractIOrderBookOrderRespondedEvent)
	if err := _ContractIOrderBook.contract.UnpackLog(event, "OrderRespondedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
