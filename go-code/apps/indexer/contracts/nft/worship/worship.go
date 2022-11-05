// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nft

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
)

// NftMetaData contains all meta data concerning the Nft contract.
var NftMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"votary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"releaseTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"}],\"name\":\"TokenOffered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"}],\"name\":\"TokenRedeemed\",\"type\":\"event\"}]",
}

// NftABI is the input ABI used to generate the binding from.
// Deprecated: Use NftMetaData.ABI instead.
var NftABI = NftMetaData.ABI

// Nft is an auto generated Go binding around an Ethereum contract.
type Nft struct {
	NftCaller     // Read-only binding to the contract
	NftTransactor // Write-only binding to the contract
	NftFilterer   // Log filterer for contract events
}

// NftCaller is an auto generated read-only Go binding around an Ethereum contract.
type NftCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NftTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NftTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NftFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NftFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NftSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NftSession struct {
	Contract     *Nft              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NftCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NftCallerSession struct {
	Contract *NftCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// NftTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NftTransactorSession struct {
	Contract     *NftTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NftRaw is an auto generated low-level Go binding around an Ethereum contract.
type NftRaw struct {
	Contract *Nft // Generic contract binding to access the raw methods on
}

// NftCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NftCallerRaw struct {
	Contract *NftCaller // Generic read-only contract binding to access the raw methods on
}

// NftTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NftTransactorRaw struct {
	Contract *NftTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNft creates a new instance of Nft, bound to a specific deployed contract.
func NewNft(address common.Address, backend bind.ContractBackend) (*Nft, error) {
	contract, err := bindNft(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Nft{NftCaller: NftCaller{contract: contract}, NftTransactor: NftTransactor{contract: contract}, NftFilterer: NftFilterer{contract: contract}}, nil
}

// NewNftCaller creates a new read-only instance of Nft, bound to a specific deployed contract.
func NewNftCaller(address common.Address, caller bind.ContractCaller) (*NftCaller, error) {
	contract, err := bindNft(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NftCaller{contract: contract}, nil
}

// NewNftTransactor creates a new write-only instance of Nft, bound to a specific deployed contract.
func NewNftTransactor(address common.Address, transactor bind.ContractTransactor) (*NftTransactor, error) {
	contract, err := bindNft(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NftTransactor{contract: contract}, nil
}

// NewNftFilterer creates a new log filterer instance of Nft, bound to a specific deployed contract.
func NewNftFilterer(address common.Address, filterer bind.ContractFilterer) (*NftFilterer, error) {
	contract, err := bindNft(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NftFilterer{contract: contract}, nil
}

// bindNft binds a generic wrapper to an already deployed contract.
func bindNft(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NftABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nft *NftRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Nft.Contract.NftCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nft *NftRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nft.Contract.NftTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nft *NftRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nft.Contract.NftTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Nft *NftCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Nft.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Nft *NftTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Nft.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Nft *NftTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Nft.Contract.contract.Transact(opts, method, params...)
}

// NftTokenOfferedIterator is returned from FilterTokenOffered and is used to iterate over the raw logs and unpacked data for TokenOffered events raised by the Nft contract.
type NftTokenOfferedIterator struct {
	Event *NftTokenOffered // Event containing the contract specifics and raw log

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
func (it *NftTokenOfferedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftTokenOffered)
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
		it.Event = new(NftTokenOffered)
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
func (it *NftTokenOfferedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftTokenOfferedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftTokenOffered represents a TokenOffered event raised by the Nft contract.
type NftTokenOffered struct {
	Nft              common.Address
	TokenId          *big.Int
	Votary           common.Address
	ReleaseTimestamp *big.Int
	Redeemer         common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterTokenOffered is a free log retrieval operation binding the contract event 0x76907a644301f3225d8bf001dbdb7e0a31054951df9bb96db68f9bb6e4f56735.
//
// Solidity: event TokenOffered(address indexed nft, uint256 indexed tokenId, address indexed votary, uint256 releaseTimestamp, address redeemer)
func (_Nft *NftFilterer) FilterTokenOffered(opts *bind.FilterOpts, nft []common.Address, tokenId []*big.Int, votary []common.Address) (*NftTokenOfferedIterator, error) {

	var nftRule []interface{}
	for _, nftItem := range nft {
		nftRule = append(nftRule, nftItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var votaryRule []interface{}
	for _, votaryItem := range votary {
		votaryRule = append(votaryRule, votaryItem)
	}

	logs, sub, err := _Nft.contract.FilterLogs(opts, "TokenOffered", nftRule, tokenIdRule, votaryRule)
	if err != nil {
		return nil, err
	}
	return &NftTokenOfferedIterator{contract: _Nft.contract, event: "TokenOffered", logs: logs, sub: sub}, nil
}

// WatchTokenOffered is a free log subscription operation binding the contract event 0x76907a644301f3225d8bf001dbdb7e0a31054951df9bb96db68f9bb6e4f56735.
//
// Solidity: event TokenOffered(address indexed nft, uint256 indexed tokenId, address indexed votary, uint256 releaseTimestamp, address redeemer)
func (_Nft *NftFilterer) WatchTokenOffered(opts *bind.WatchOpts, sink chan<- *NftTokenOffered, nft []common.Address, tokenId []*big.Int, votary []common.Address) (event.Subscription, error) {

	var nftRule []interface{}
	for _, nftItem := range nft {
		nftRule = append(nftRule, nftItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var votaryRule []interface{}
	for _, votaryItem := range votary {
		votaryRule = append(votaryRule, votaryItem)
	}

	logs, sub, err := _Nft.contract.WatchLogs(opts, "TokenOffered", nftRule, tokenIdRule, votaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftTokenOffered)
				if err := _Nft.contract.UnpackLog(event, "TokenOffered", log); err != nil {
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

// ParseTokenOffered is a log parse operation binding the contract event 0x76907a644301f3225d8bf001dbdb7e0a31054951df9bb96db68f9bb6e4f56735.
//
// Solidity: event TokenOffered(address indexed nft, uint256 indexed tokenId, address indexed votary, uint256 releaseTimestamp, address redeemer)
func (_Nft *NftFilterer) ParseTokenOffered(log types.Log) (*NftTokenOffered, error) {
	event := new(NftTokenOffered)
	if err := _Nft.contract.UnpackLog(event, "TokenOffered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftTokenRedeemedIterator is returned from FilterTokenRedeemed and is used to iterate over the raw logs and unpacked data for TokenRedeemed events raised by the Nft contract.
type NftTokenRedeemedIterator struct {
	Event *NftTokenRedeemed // Event containing the contract specifics and raw log

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
func (it *NftTokenRedeemedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftTokenRedeemed)
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
		it.Event = new(NftTokenRedeemed)
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
func (it *NftTokenRedeemedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftTokenRedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftTokenRedeemed represents a TokenRedeemed event raised by the Nft contract.
type NftTokenRedeemed struct {
	Nft      common.Address
	TokenId  *big.Int
	Redeemer common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTokenRedeemed is a free log retrieval operation binding the contract event 0x675a099843d220e2829d283422deacbf7d1692118a7569c329dd638932b6ae49.
//
// Solidity: event TokenRedeemed(address indexed nft, uint256 indexed tokenId, address indexed redeemer)
func (_Nft *NftFilterer) FilterTokenRedeemed(opts *bind.FilterOpts, nft []common.Address, tokenId []*big.Int, redeemer []common.Address) (*NftTokenRedeemedIterator, error) {

	var nftRule []interface{}
	for _, nftItem := range nft {
		nftRule = append(nftRule, nftItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var redeemerRule []interface{}
	for _, redeemerItem := range redeemer {
		redeemerRule = append(redeemerRule, redeemerItem)
	}

	logs, sub, err := _Nft.contract.FilterLogs(opts, "TokenRedeemed", nftRule, tokenIdRule, redeemerRule)
	if err != nil {
		return nil, err
	}
	return &NftTokenRedeemedIterator{contract: _Nft.contract, event: "TokenRedeemed", logs: logs, sub: sub}, nil
}

// WatchTokenRedeemed is a free log subscription operation binding the contract event 0x675a099843d220e2829d283422deacbf7d1692118a7569c329dd638932b6ae49.
//
// Solidity: event TokenRedeemed(address indexed nft, uint256 indexed tokenId, address indexed redeemer)
func (_Nft *NftFilterer) WatchTokenRedeemed(opts *bind.WatchOpts, sink chan<- *NftTokenRedeemed, nft []common.Address, tokenId []*big.Int, redeemer []common.Address) (event.Subscription, error) {

	var nftRule []interface{}
	for _, nftItem := range nft {
		nftRule = append(nftRule, nftItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var redeemerRule []interface{}
	for _, redeemerItem := range redeemer {
		redeemerRule = append(redeemerRule, redeemerItem)
	}

	logs, sub, err := _Nft.contract.WatchLogs(opts, "TokenRedeemed", nftRule, tokenIdRule, redeemerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftTokenRedeemed)
				if err := _Nft.contract.UnpackLog(event, "TokenRedeemed", log); err != nil {
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

// ParseTokenRedeemed is a log parse operation binding the contract event 0x675a099843d220e2829d283422deacbf7d1692118a7569c329dd638932b6ae49.
//
// Solidity: event TokenRedeemed(address indexed nft, uint256 indexed tokenId, address indexed redeemer)
func (_Nft *NftFilterer) ParseTokenRedeemed(log types.Log) (*NftTokenRedeemed, error) {
	event := new(NftTokenRedeemed)
	if err := _Nft.contract.UnpackLog(event, "TokenRedeemed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
