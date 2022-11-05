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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"offerer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"increasedNonce\",\"type\":\"uint256\"}],\"name\":\"AllOrdersCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"assets_struct_hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nft_id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nft_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"ft\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ft_amount\",\"type\":\"uint256\"}],\"name\":\"AssetBytesInfo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"orderBytes\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"assetsBytes\",\"type\":\"bytes\"}],\"name\":\"FixedPriceOrderMatched\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"conduit\",\"type\":\"address\"}],\"name\":\"NewConduit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"order_struct_hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"royalty_recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"royalty_rate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"start_at\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"expire_at\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"maker_nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"taker_get_nft\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"assets_hash\",\"type\":\"bytes32\"}],\"name\":\"OrderBytesInfo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"OrderCancelled\",\"type\":\"event\"}]",
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

// NftAllOrdersCancelledIterator is returned from FilterAllOrdersCancelled and is used to iterate over the raw logs and unpacked data for AllOrdersCancelled events raised by the Nft contract.
type NftAllOrdersCancelledIterator struct {
	Event *NftAllOrdersCancelled // Event containing the contract specifics and raw log

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
func (it *NftAllOrdersCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftAllOrdersCancelled)
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
		it.Event = new(NftAllOrdersCancelled)
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
func (it *NftAllOrdersCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftAllOrdersCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftAllOrdersCancelled represents a AllOrdersCancelled event raised by the Nft contract.
type NftAllOrdersCancelled struct {
	Offerer        common.Address
	IncreasedNonce *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterAllOrdersCancelled is a free log retrieval operation binding the contract event 0x83a782ac7424737a1190d4668474e765f07d603de0485a081dbc343ac1b02099.
//
// Solidity: event AllOrdersCancelled(address indexed offerer, uint256 increasedNonce)
func (_Nft *NftFilterer) FilterAllOrdersCancelled(opts *bind.FilterOpts, offerer []common.Address) (*NftAllOrdersCancelledIterator, error) {

	var offererRule []interface{}
	for _, offererItem := range offerer {
		offererRule = append(offererRule, offererItem)
	}

	logs, sub, err := _Nft.contract.FilterLogs(opts, "AllOrdersCancelled", offererRule)
	if err != nil {
		return nil, err
	}
	return &NftAllOrdersCancelledIterator{contract: _Nft.contract, event: "AllOrdersCancelled", logs: logs, sub: sub}, nil
}

// WatchAllOrdersCancelled is a free log subscription operation binding the contract event 0x83a782ac7424737a1190d4668474e765f07d603de0485a081dbc343ac1b02099.
//
// Solidity: event AllOrdersCancelled(address indexed offerer, uint256 increasedNonce)
func (_Nft *NftFilterer) WatchAllOrdersCancelled(opts *bind.WatchOpts, sink chan<- *NftAllOrdersCancelled, offerer []common.Address) (event.Subscription, error) {

	var offererRule []interface{}
	for _, offererItem := range offerer {
		offererRule = append(offererRule, offererItem)
	}

	logs, sub, err := _Nft.contract.WatchLogs(opts, "AllOrdersCancelled", offererRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftAllOrdersCancelled)
				if err := _Nft.contract.UnpackLog(event, "AllOrdersCancelled", log); err != nil {
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

// ParseAllOrdersCancelled is a log parse operation binding the contract event 0x83a782ac7424737a1190d4668474e765f07d603de0485a081dbc343ac1b02099.
//
// Solidity: event AllOrdersCancelled(address indexed offerer, uint256 increasedNonce)
func (_Nft *NftFilterer) ParseAllOrdersCancelled(log types.Log) (*NftAllOrdersCancelled, error) {
	event := new(NftAllOrdersCancelled)
	if err := _Nft.contract.UnpackLog(event, "AllOrdersCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftAssetBytesInfoIterator is returned from FilterAssetBytesInfo and is used to iterate over the raw logs and unpacked data for AssetBytesInfo events raised by the Nft contract.
type NftAssetBytesInfoIterator struct {
	Event *NftAssetBytesInfo // Event containing the contract specifics and raw log

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
func (it *NftAssetBytesInfoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftAssetBytesInfo)
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
		it.Event = new(NftAssetBytesInfo)
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
func (it *NftAssetBytesInfoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftAssetBytesInfoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftAssetBytesInfo represents a AssetBytesInfo event raised by the Nft contract.
type NftAssetBytesInfo struct {
	AssetsStructHash [32]byte
	Nft              common.Address
	NftId            *big.Int
	NftAmount        *big.Int
	Ft               common.Address
	FtAmount         *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterAssetBytesInfo is a free log retrieval operation binding the contract event 0xb172a07c0f99cb970d22581f6a1015978111c9ae944eb60390b7f5753dd430f9.
//
// Solidity: event AssetBytesInfo(bytes32 assets_struct_hash, address nft, uint256 nft_id, uint256 nft_amount, address ft, uint256 ft_amount)
func (_Nft *NftFilterer) FilterAssetBytesInfo(opts *bind.FilterOpts) (*NftAssetBytesInfoIterator, error) {

	logs, sub, err := _Nft.contract.FilterLogs(opts, "AssetBytesInfo")
	if err != nil {
		return nil, err
	}
	return &NftAssetBytesInfoIterator{contract: _Nft.contract, event: "AssetBytesInfo", logs: logs, sub: sub}, nil
}

// WatchAssetBytesInfo is a free log subscription operation binding the contract event 0xb172a07c0f99cb970d22581f6a1015978111c9ae944eb60390b7f5753dd430f9.
//
// Solidity: event AssetBytesInfo(bytes32 assets_struct_hash, address nft, uint256 nft_id, uint256 nft_amount, address ft, uint256 ft_amount)
func (_Nft *NftFilterer) WatchAssetBytesInfo(opts *bind.WatchOpts, sink chan<- *NftAssetBytesInfo) (event.Subscription, error) {

	logs, sub, err := _Nft.contract.WatchLogs(opts, "AssetBytesInfo")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftAssetBytesInfo)
				if err := _Nft.contract.UnpackLog(event, "AssetBytesInfo", log); err != nil {
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

// ParseAssetBytesInfo is a log parse operation binding the contract event 0xb172a07c0f99cb970d22581f6a1015978111c9ae944eb60390b7f5753dd430f9.
//
// Solidity: event AssetBytesInfo(bytes32 assets_struct_hash, address nft, uint256 nft_id, uint256 nft_amount, address ft, uint256 ft_amount)
func (_Nft *NftFilterer) ParseAssetBytesInfo(log types.Log) (*NftAssetBytesInfo, error) {
	event := new(NftAssetBytesInfo)
	if err := _Nft.contract.UnpackLog(event, "AssetBytesInfo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftFixedPriceOrderMatchedIterator is returned from FilterFixedPriceOrderMatched and is used to iterate over the raw logs and unpacked data for FixedPriceOrderMatched events raised by the Nft contract.
type NftFixedPriceOrderMatchedIterator struct {
	Event *NftFixedPriceOrderMatched // Event containing the contract specifics and raw log

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
func (it *NftFixedPriceOrderMatchedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftFixedPriceOrderMatched)
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
		it.Event = new(NftFixedPriceOrderMatched)
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
func (it *NftFixedPriceOrderMatchedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftFixedPriceOrderMatchedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftFixedPriceOrderMatched represents a FixedPriceOrderMatched event raised by the Nft contract.
type NftFixedPriceOrderMatched struct {
	Maker       common.Address
	Taker       common.Address
	OrderHash   [32]byte
	OrderBytes  []byte
	AssetsBytes []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFixedPriceOrderMatched is a free log retrieval operation binding the contract event 0xf98e5acebea57eaddbc56ef62e1ad0c1fb3f8b3b06e16d83ef8e4ad578bdad52.
//
// Solidity: event FixedPriceOrderMatched(address indexed maker, address indexed taker, bytes32 indexed orderHash, bytes orderBytes, bytes assetsBytes)
func (_Nft *NftFilterer) FilterFixedPriceOrderMatched(opts *bind.FilterOpts, maker []common.Address, taker []common.Address, orderHash [][32]byte) (*NftFixedPriceOrderMatchedIterator, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}
	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}
	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _Nft.contract.FilterLogs(opts, "FixedPriceOrderMatched", makerRule, takerRule, orderHashRule)
	if err != nil {
		return nil, err
	}
	return &NftFixedPriceOrderMatchedIterator{contract: _Nft.contract, event: "FixedPriceOrderMatched", logs: logs, sub: sub}, nil
}

// WatchFixedPriceOrderMatched is a free log subscription operation binding the contract event 0xf98e5acebea57eaddbc56ef62e1ad0c1fb3f8b3b06e16d83ef8e4ad578bdad52.
//
// Solidity: event FixedPriceOrderMatched(address indexed maker, address indexed taker, bytes32 indexed orderHash, bytes orderBytes, bytes assetsBytes)
func (_Nft *NftFilterer) WatchFixedPriceOrderMatched(opts *bind.WatchOpts, sink chan<- *NftFixedPriceOrderMatched, maker []common.Address, taker []common.Address, orderHash [][32]byte) (event.Subscription, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}
	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}
	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _Nft.contract.WatchLogs(opts, "FixedPriceOrderMatched", makerRule, takerRule, orderHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftFixedPriceOrderMatched)
				if err := _Nft.contract.UnpackLog(event, "FixedPriceOrderMatched", log); err != nil {
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

// ParseFixedPriceOrderMatched is a log parse operation binding the contract event 0xf98e5acebea57eaddbc56ef62e1ad0c1fb3f8b3b06e16d83ef8e4ad578bdad52.
//
// Solidity: event FixedPriceOrderMatched(address indexed maker, address indexed taker, bytes32 indexed orderHash, bytes orderBytes, bytes assetsBytes)
func (_Nft *NftFilterer) ParseFixedPriceOrderMatched(log types.Log) (*NftFixedPriceOrderMatched, error) {
	event := new(NftFixedPriceOrderMatched)
	if err := _Nft.contract.UnpackLog(event, "FixedPriceOrderMatched", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftNewConduitIterator is returned from FilterNewConduit and is used to iterate over the raw logs and unpacked data for NewConduit events raised by the Nft contract.
type NftNewConduitIterator struct {
	Event *NftNewConduit // Event containing the contract specifics and raw log

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
func (it *NftNewConduitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftNewConduit)
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
		it.Event = new(NftNewConduit)
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
func (it *NftNewConduitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftNewConduitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftNewConduit represents a NewConduit event raised by the Nft contract.
type NftNewConduit struct {
	User    common.Address
	Conduit common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewConduit is a free log retrieval operation binding the contract event 0x79f3cef39f5f08f8f2e396ce836c7c5f58736a27400595b231f3e91afb489f2d.
//
// Solidity: event NewConduit(address user, address conduit)
func (_Nft *NftFilterer) FilterNewConduit(opts *bind.FilterOpts) (*NftNewConduitIterator, error) {

	logs, sub, err := _Nft.contract.FilterLogs(opts, "NewConduit")
	if err != nil {
		return nil, err
	}
	return &NftNewConduitIterator{contract: _Nft.contract, event: "NewConduit", logs: logs, sub: sub}, nil
}

// WatchNewConduit is a free log subscription operation binding the contract event 0x79f3cef39f5f08f8f2e396ce836c7c5f58736a27400595b231f3e91afb489f2d.
//
// Solidity: event NewConduit(address user, address conduit)
func (_Nft *NftFilterer) WatchNewConduit(opts *bind.WatchOpts, sink chan<- *NftNewConduit) (event.Subscription, error) {

	logs, sub, err := _Nft.contract.WatchLogs(opts, "NewConduit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftNewConduit)
				if err := _Nft.contract.UnpackLog(event, "NewConduit", log); err != nil {
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

// ParseNewConduit is a log parse operation binding the contract event 0x79f3cef39f5f08f8f2e396ce836c7c5f58736a27400595b231f3e91afb489f2d.
//
// Solidity: event NewConduit(address user, address conduit)
func (_Nft *NftFilterer) ParseNewConduit(log types.Log) (*NftNewConduit, error) {
	event := new(NftNewConduit)
	if err := _Nft.contract.UnpackLog(event, "NewConduit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftOrderBytesInfoIterator is returned from FilterOrderBytesInfo and is used to iterate over the raw logs and unpacked data for OrderBytesInfo events raised by the Nft contract.
type NftOrderBytesInfoIterator struct {
	Event *NftOrderBytesInfo // Event containing the contract specifics and raw log

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
func (it *NftOrderBytesInfoIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftOrderBytesInfo)
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
		it.Event = new(NftOrderBytesInfo)
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
func (it *NftOrderBytesInfoIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftOrderBytesInfoIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftOrderBytesInfo represents a OrderBytesInfo event raised by the Nft contract.
type NftOrderBytesInfo struct {
	OrderStructHash  [32]byte
	Maker            common.Address
	Taker            common.Address
	RoyaltyRecipient common.Address
	RoyaltyRate      *big.Int
	StartAt          uint64
	ExpireAt         uint64
	MakerNonce       uint64
	TakerGetNft      bool
	AssetsHash       [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOrderBytesInfo is a free log retrieval operation binding the contract event 0x8c25274764e0998cb9a4c1405cbb44678d1a345dc8d5524a4310a1b066f81297.
//
// Solidity: event OrderBytesInfo(bytes32 order_struct_hash, address maker, address taker, address royalty_recipient, uint256 royalty_rate, uint64 start_at, uint64 expire_at, uint64 maker_nonce, bool taker_get_nft, bytes32 assets_hash)
func (_Nft *NftFilterer) FilterOrderBytesInfo(opts *bind.FilterOpts) (*NftOrderBytesInfoIterator, error) {

	logs, sub, err := _Nft.contract.FilterLogs(opts, "OrderBytesInfo")
	if err != nil {
		return nil, err
	}
	return &NftOrderBytesInfoIterator{contract: _Nft.contract, event: "OrderBytesInfo", logs: logs, sub: sub}, nil
}

// WatchOrderBytesInfo is a free log subscription operation binding the contract event 0x8c25274764e0998cb9a4c1405cbb44678d1a345dc8d5524a4310a1b066f81297.
//
// Solidity: event OrderBytesInfo(bytes32 order_struct_hash, address maker, address taker, address royalty_recipient, uint256 royalty_rate, uint64 start_at, uint64 expire_at, uint64 maker_nonce, bool taker_get_nft, bytes32 assets_hash)
func (_Nft *NftFilterer) WatchOrderBytesInfo(opts *bind.WatchOpts, sink chan<- *NftOrderBytesInfo) (event.Subscription, error) {

	logs, sub, err := _Nft.contract.WatchLogs(opts, "OrderBytesInfo")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftOrderBytesInfo)
				if err := _Nft.contract.UnpackLog(event, "OrderBytesInfo", log); err != nil {
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

// ParseOrderBytesInfo is a log parse operation binding the contract event 0x8c25274764e0998cb9a4c1405cbb44678d1a345dc8d5524a4310a1b066f81297.
//
// Solidity: event OrderBytesInfo(bytes32 order_struct_hash, address maker, address taker, address royalty_recipient, uint256 royalty_rate, uint64 start_at, uint64 expire_at, uint64 maker_nonce, bool taker_get_nft, bytes32 assets_hash)
func (_Nft *NftFilterer) ParseOrderBytesInfo(log types.Log) (*NftOrderBytesInfo, error) {
	event := new(NftOrderBytesInfo)
	if err := _Nft.contract.UnpackLog(event, "OrderBytesInfo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NftOrderCancelledIterator is returned from FilterOrderCancelled and is used to iterate over the raw logs and unpacked data for OrderCancelled events raised by the Nft contract.
type NftOrderCancelledIterator struct {
	Event *NftOrderCancelled // Event containing the contract specifics and raw log

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
func (it *NftOrderCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NftOrderCancelled)
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
		it.Event = new(NftOrderCancelled)
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
func (it *NftOrderCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NftOrderCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NftOrderCancelled represents a OrderCancelled event raised by the Nft contract.
type NftOrderCancelled struct {
	Maker     common.Address
	OrderHash [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOrderCancelled is a free log retrieval operation binding the contract event 0x35974c4230d53fb4c6e8553fd900c88ba92747dbc689a79bcd6ba755cb936985.
//
// Solidity: event OrderCancelled(address indexed maker, bytes32 indexed orderHash)
func (_Nft *NftFilterer) FilterOrderCancelled(opts *bind.FilterOpts, maker []common.Address, orderHash [][32]byte) (*NftOrderCancelledIterator, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}
	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _Nft.contract.FilterLogs(opts, "OrderCancelled", makerRule, orderHashRule)
	if err != nil {
		return nil, err
	}
	return &NftOrderCancelledIterator{contract: _Nft.contract, event: "OrderCancelled", logs: logs, sub: sub}, nil
}

// WatchOrderCancelled is a free log subscription operation binding the contract event 0x35974c4230d53fb4c6e8553fd900c88ba92747dbc689a79bcd6ba755cb936985.
//
// Solidity: event OrderCancelled(address indexed maker, bytes32 indexed orderHash)
func (_Nft *NftFilterer) WatchOrderCancelled(opts *bind.WatchOpts, sink chan<- *NftOrderCancelled, maker []common.Address, orderHash [][32]byte) (event.Subscription, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}
	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _Nft.contract.WatchLogs(opts, "OrderCancelled", makerRule, orderHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NftOrderCancelled)
				if err := _Nft.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
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

// ParseOrderCancelled is a log parse operation binding the contract event 0x35974c4230d53fb4c6e8553fd900c88ba92747dbc689a79bcd6ba755cb936985.
//
// Solidity: event OrderCancelled(address indexed maker, bytes32 indexed orderHash)
func (_Nft *NftFilterer) ParseOrderCancelled(log types.Log) (*NftOrderCancelled, error) {
	event := new(NftOrderCancelled)
	if err := _Nft.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
