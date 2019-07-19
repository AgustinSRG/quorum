// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package permission

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// VoterManagerABI is the input ABI used to generate the binding from.
const VoterManagerABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"getPendingOpDetails\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_vAccount\",\"type\":\"address\"}],\"name\":\"addVoter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_vAccount\",\"type\":\"address\"}],\"name\":\"deleteVoter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_authOrg\",\"type\":\"string\"},{\"name\":\"_vAccount\",\"type\":\"address\"},{\"name\":\"_pendingOp\",\"type\":\"uint256\"}],\"name\":\"processVote\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_authOrg\",\"type\":\"string\"},{\"name\":\"_orgId\",\"type\":\"string\"},{\"name\":\"_enodeId\",\"type\":\"string\"},{\"name\":\"_account\",\"type\":\"address\"},{\"name\":\"_pendingOp\",\"type\":\"uint256\"}],\"name\":\"addVotingItem\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_permUpgradable\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_vAccount\",\"type\":\"address\"}],\"name\":\"VoterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_vAccount\",\"type\":\"address\"}],\"name\":\"VoterDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"VotingItemAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_orgId\",\"type\":\"string\"}],\"name\":\"VoteProcessed\",\"type\":\"event\"}]"

// VoterManager is an auto generated Go binding around an Ethereum contract.
type VoterManager struct {
	VoterManagerCaller     // Read-only binding to the contract
	VoterManagerTransactor // Write-only binding to the contract
	VoterManagerFilterer   // Log filterer for contract events
}

// VoterManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type VoterManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoterManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VoterManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoterManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VoterManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoterManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VoterManagerSession struct {
	Contract     *VoterManager     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VoterManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VoterManagerCallerSession struct {
	Contract *VoterManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// VoterManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VoterManagerTransactorSession struct {
	Contract     *VoterManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// VoterManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type VoterManagerRaw struct {
	Contract *VoterManager // Generic contract binding to access the raw methods on
}

// VoterManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VoterManagerCallerRaw struct {
	Contract *VoterManagerCaller // Generic read-only contract binding to access the raw methods on
}

// VoterManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VoterManagerTransactorRaw struct {
	Contract *VoterManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVoterManager creates a new instance of VoterManager, bound to a specific deployed contract.
func NewVoterManager(address common.Address, backend bind.ContractBackend) (*VoterManager, error) {
	contract, err := bindVoterManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VoterManager{VoterManagerCaller: VoterManagerCaller{contract: contract}, VoterManagerTransactor: VoterManagerTransactor{contract: contract}, VoterManagerFilterer: VoterManagerFilterer{contract: contract}}, nil
}

// NewVoterManagerCaller creates a new read-only instance of VoterManager, bound to a specific deployed contract.
func NewVoterManagerCaller(address common.Address, caller bind.ContractCaller) (*VoterManagerCaller, error) {
	contract, err := bindVoterManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VoterManagerCaller{contract: contract}, nil
}

// NewVoterManagerTransactor creates a new write-only instance of VoterManager, bound to a specific deployed contract.
func NewVoterManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*VoterManagerTransactor, error) {
	contract, err := bindVoterManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VoterManagerTransactor{contract: contract}, nil
}

// NewVoterManagerFilterer creates a new log filterer instance of VoterManager, bound to a specific deployed contract.
func NewVoterManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*VoterManagerFilterer, error) {
	contract, err := bindVoterManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VoterManagerFilterer{contract: contract}, nil
}

// bindVoterManager binds a generic wrapper to an already deployed contract.
func bindVoterManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VoterManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VoterManager *VoterManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VoterManager.Contract.VoterManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VoterManager *VoterManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VoterManager.Contract.VoterManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VoterManager *VoterManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VoterManager.Contract.VoterManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VoterManager *VoterManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VoterManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VoterManager *VoterManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VoterManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VoterManager *VoterManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VoterManager.Contract.contract.Transact(opts, method, params...)
}

// GetPendingOpDetails is a free data retrieval call binding the contract method 0x014e6acc.
//
// Solidity: function getPendingOpDetails(_orgId string) constant returns(string, string, address, uint256)
func (_VoterManager *VoterManagerCaller) GetPendingOpDetails(opts *bind.CallOpts, _orgId string) (string, string, common.Address, *big.Int, error) {
	var (
		ret0 = new(string)
		ret1 = new(string)
		ret2 = new(common.Address)
		ret3 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _VoterManager.contract.Call(opts, out, "getPendingOpDetails", _orgId)
	return *ret0, *ret1, *ret2, *ret3, err
}

// GetPendingOpDetails is a free data retrieval call binding the contract method 0x014e6acc.
//
// Solidity: function getPendingOpDetails(_orgId string) constant returns(string, string, address, uint256)
func (_VoterManager *VoterManagerSession) GetPendingOpDetails(_orgId string) (string, string, common.Address, *big.Int, error) {
	return _VoterManager.Contract.GetPendingOpDetails(&_VoterManager.CallOpts, _orgId)
}

// GetPendingOpDetails is a free data retrieval call binding the contract method 0x014e6acc.
//
// Solidity: function getPendingOpDetails(_orgId string) constant returns(string, string, address, uint256)
func (_VoterManager *VoterManagerCallerSession) GetPendingOpDetails(_orgId string) (string, string, common.Address, *big.Int, error) {
	return _VoterManager.Contract.GetPendingOpDetails(&_VoterManager.CallOpts, _orgId)
}

// AddVoter is a paid mutator transaction binding the contract method 0x5607395b.
//
// Solidity: function addVoter(_orgId string, _vAccount address) returns()
func (_VoterManager *VoterManagerTransactor) AddVoter(opts *bind.TransactOpts, _orgId string, _vAccount common.Address) (*types.Transaction, error) {
	return _VoterManager.contract.Transact(opts, "addVoter", _orgId, _vAccount)
}

// AddVoter is a paid mutator transaction binding the contract method 0x5607395b.
//
// Solidity: function addVoter(_orgId string, _vAccount address) returns()
func (_VoterManager *VoterManagerSession) AddVoter(_orgId string, _vAccount common.Address) (*types.Transaction, error) {
	return _VoterManager.Contract.AddVoter(&_VoterManager.TransactOpts, _orgId, _vAccount)
}

// AddVoter is a paid mutator transaction binding the contract method 0x5607395b.
//
// Solidity: function addVoter(_orgId string, _vAccount address) returns()
func (_VoterManager *VoterManagerTransactorSession) AddVoter(_orgId string, _vAccount common.Address) (*types.Transaction, error) {
	return _VoterManager.Contract.AddVoter(&_VoterManager.TransactOpts, _orgId, _vAccount)
}

// AddVotingItem is a paid mutator transaction binding the contract method 0xe98ac22d.
//
// Solidity: function addVotingItem(_authOrg string, _orgId string, _enodeId string, _account address, _pendingOp uint256) returns()
func (_VoterManager *VoterManagerTransactor) AddVotingItem(opts *bind.TransactOpts, _authOrg string, _orgId string, _enodeId string, _account common.Address, _pendingOp *big.Int) (*types.Transaction, error) {
	return _VoterManager.contract.Transact(opts, "addVotingItem", _authOrg, _orgId, _enodeId, _account, _pendingOp)
}

// AddVotingItem is a paid mutator transaction binding the contract method 0xe98ac22d.
//
// Solidity: function addVotingItem(_authOrg string, _orgId string, _enodeId string, _account address, _pendingOp uint256) returns()
func (_VoterManager *VoterManagerSession) AddVotingItem(_authOrg string, _orgId string, _enodeId string, _account common.Address, _pendingOp *big.Int) (*types.Transaction, error) {
	return _VoterManager.Contract.AddVotingItem(&_VoterManager.TransactOpts, _authOrg, _orgId, _enodeId, _account, _pendingOp)
}

// AddVotingItem is a paid mutator transaction binding the contract method 0xe98ac22d.
//
// Solidity: function addVotingItem(_authOrg string, _orgId string, _enodeId string, _account address, _pendingOp uint256) returns()
func (_VoterManager *VoterManagerTransactorSession) AddVotingItem(_authOrg string, _orgId string, _enodeId string, _account common.Address, _pendingOp *big.Int) (*types.Transaction, error) {
	return _VoterManager.Contract.AddVotingItem(&_VoterManager.TransactOpts, _authOrg, _orgId, _enodeId, _account, _pendingOp)
}

// DeleteVoter is a paid mutator transaction binding the contract method 0x59cbd6fe.
//
// Solidity: function deleteVoter(_orgId string, _vAccount address) returns()
func (_VoterManager *VoterManagerTransactor) DeleteVoter(opts *bind.TransactOpts, _orgId string, _vAccount common.Address) (*types.Transaction, error) {
	return _VoterManager.contract.Transact(opts, "deleteVoter", _orgId, _vAccount)
}

// DeleteVoter is a paid mutator transaction binding the contract method 0x59cbd6fe.
//
// Solidity: function deleteVoter(_orgId string, _vAccount address) returns()
func (_VoterManager *VoterManagerSession) DeleteVoter(_orgId string, _vAccount common.Address) (*types.Transaction, error) {
	return _VoterManager.Contract.DeleteVoter(&_VoterManager.TransactOpts, _orgId, _vAccount)
}

// DeleteVoter is a paid mutator transaction binding the contract method 0x59cbd6fe.
//
// Solidity: function deleteVoter(_orgId string, _vAccount address) returns()
func (_VoterManager *VoterManagerTransactorSession) DeleteVoter(_orgId string, _vAccount common.Address) (*types.Transaction, error) {
	return _VoterManager.Contract.DeleteVoter(&_VoterManager.TransactOpts, _orgId, _vAccount)
}

// ProcessVote is a paid mutator transaction binding the contract method 0xb0213864.
//
// Solidity: function processVote(_authOrg string, _vAccount address, _pendingOp uint256) returns(bool)
func (_VoterManager *VoterManagerTransactor) ProcessVote(opts *bind.TransactOpts, _authOrg string, _vAccount common.Address, _pendingOp *big.Int) (*types.Transaction, error) {
	return _VoterManager.contract.Transact(opts, "processVote", _authOrg, _vAccount, _pendingOp)
}

// ProcessVote is a paid mutator transaction binding the contract method 0xb0213864.
//
// Solidity: function processVote(_authOrg string, _vAccount address, _pendingOp uint256) returns(bool)
func (_VoterManager *VoterManagerSession) ProcessVote(_authOrg string, _vAccount common.Address, _pendingOp *big.Int) (*types.Transaction, error) {
	return _VoterManager.Contract.ProcessVote(&_VoterManager.TransactOpts, _authOrg, _vAccount, _pendingOp)
}

// ProcessVote is a paid mutator transaction binding the contract method 0xb0213864.
//
// Solidity: function processVote(_authOrg string, _vAccount address, _pendingOp uint256) returns(bool)
func (_VoterManager *VoterManagerTransactorSession) ProcessVote(_authOrg string, _vAccount common.Address, _pendingOp *big.Int) (*types.Transaction, error) {
	return _VoterManager.Contract.ProcessVote(&_VoterManager.TransactOpts, _authOrg, _vAccount, _pendingOp)
}

// VoterManagerVoteProcessedIterator is returned from FilterVoteProcessed and is used to iterate over the raw logs and unpacked data for VoteProcessed events raised by the VoterManager contract.
type VoterManagerVoteProcessedIterator struct {
	Event *VoterManagerVoteProcessed // Event containing the contract specifics and raw log

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
func (it *VoterManagerVoteProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoterManagerVoteProcessed)
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
		it.Event = new(VoterManagerVoteProcessed)
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
func (it *VoterManagerVoteProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoterManagerVoteProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoterManagerVoteProcessed represents a VoteProcessed event raised by the VoterManager contract.
type VoterManagerVoteProcessed struct {
	OrgId string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterVoteProcessed is a free log retrieval operation binding the contract event 0x87999b54e45aa02834a1265e356d7bcdceb72b8cbb4396ebaeba32a103b43508.
//
// Solidity: e VoteProcessed(_orgId string)
func (_VoterManager *VoterManagerFilterer) FilterVoteProcessed(opts *bind.FilterOpts) (*VoterManagerVoteProcessedIterator, error) {

	logs, sub, err := _VoterManager.contract.FilterLogs(opts, "VoteProcessed")
	if err != nil {
		return nil, err
	}
	return &VoterManagerVoteProcessedIterator{contract: _VoterManager.contract, event: "VoteProcessed", logs: logs, sub: sub}, nil
}

// WatchVoteProcessed is a free log subscription operation binding the contract event 0x87999b54e45aa02834a1265e356d7bcdceb72b8cbb4396ebaeba32a103b43508.
//
// Solidity: e VoteProcessed(_orgId string)
func (_VoterManager *VoterManagerFilterer) WatchVoteProcessed(opts *bind.WatchOpts, sink chan<- *VoterManagerVoteProcessed) (event.Subscription, error) {

	logs, sub, err := _VoterManager.contract.WatchLogs(opts, "VoteProcessed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoterManagerVoteProcessed)
				if err := _VoterManager.contract.UnpackLog(event, "VoteProcessed", log); err != nil {
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

// VoterManagerVoterAddedIterator is returned from FilterVoterAdded and is used to iterate over the raw logs and unpacked data for VoterAdded events raised by the VoterManager contract.
type VoterManagerVoterAddedIterator struct {
	Event *VoterManagerVoterAdded // Event containing the contract specifics and raw log

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
func (it *VoterManagerVoterAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoterManagerVoterAdded)
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
		it.Event = new(VoterManagerVoterAdded)
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
func (it *VoterManagerVoterAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoterManagerVoterAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoterManagerVoterAdded represents a VoterAdded event raised by the VoterManager contract.
type VoterManagerVoterAdded struct {
	OrgId    string
	VAccount common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterVoterAdded is a free log retrieval operation binding the contract event 0x424f3ad05c61ea35cad66f22b70b1fad7250d8229921238078c401db36d34574.
//
// Solidity: e VoterAdded(_orgId string, _vAccount address)
func (_VoterManager *VoterManagerFilterer) FilterVoterAdded(opts *bind.FilterOpts) (*VoterManagerVoterAddedIterator, error) {

	logs, sub, err := _VoterManager.contract.FilterLogs(opts, "VoterAdded")
	if err != nil {
		return nil, err
	}
	return &VoterManagerVoterAddedIterator{contract: _VoterManager.contract, event: "VoterAdded", logs: logs, sub: sub}, nil
}

// WatchVoterAdded is a free log subscription operation binding the contract event 0x424f3ad05c61ea35cad66f22b70b1fad7250d8229921238078c401db36d34574.
//
// Solidity: e VoterAdded(_orgId string, _vAccount address)
func (_VoterManager *VoterManagerFilterer) WatchVoterAdded(opts *bind.WatchOpts, sink chan<- *VoterManagerVoterAdded) (event.Subscription, error) {

	logs, sub, err := _VoterManager.contract.WatchLogs(opts, "VoterAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoterManagerVoterAdded)
				if err := _VoterManager.contract.UnpackLog(event, "VoterAdded", log); err != nil {
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

// VoterManagerVoterDeletedIterator is returned from FilterVoterDeleted and is used to iterate over the raw logs and unpacked data for VoterDeleted events raised by the VoterManager contract.
type VoterManagerVoterDeletedIterator struct {
	Event *VoterManagerVoterDeleted // Event containing the contract specifics and raw log

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
func (it *VoterManagerVoterDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoterManagerVoterDeleted)
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
		it.Event = new(VoterManagerVoterDeleted)
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
func (it *VoterManagerVoterDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoterManagerVoterDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoterManagerVoterDeleted represents a VoterDeleted event raised by the VoterManager contract.
type VoterManagerVoterDeleted struct {
	OrgId    string
	VAccount common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterVoterDeleted is a free log retrieval operation binding the contract event 0x654cd85d9b2abaf3affef0a047625d088e6e4d0448935c9b5016b5f5aa0ca3b6.
//
// Solidity: e VoterDeleted(_orgId string, _vAccount address)
func (_VoterManager *VoterManagerFilterer) FilterVoterDeleted(opts *bind.FilterOpts) (*VoterManagerVoterDeletedIterator, error) {

	logs, sub, err := _VoterManager.contract.FilterLogs(opts, "VoterDeleted")
	if err != nil {
		return nil, err
	}
	return &VoterManagerVoterDeletedIterator{contract: _VoterManager.contract, event: "VoterDeleted", logs: logs, sub: sub}, nil
}

// WatchVoterDeleted is a free log subscription operation binding the contract event 0x654cd85d9b2abaf3affef0a047625d088e6e4d0448935c9b5016b5f5aa0ca3b6.
//
// Solidity: e VoterDeleted(_orgId string, _vAccount address)
func (_VoterManager *VoterManagerFilterer) WatchVoterDeleted(opts *bind.WatchOpts, sink chan<- *VoterManagerVoterDeleted) (event.Subscription, error) {

	logs, sub, err := _VoterManager.contract.WatchLogs(opts, "VoterDeleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoterManagerVoterDeleted)
				if err := _VoterManager.contract.UnpackLog(event, "VoterDeleted", log); err != nil {
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

// VoterManagerVotingItemAddedIterator is returned from FilterVotingItemAdded and is used to iterate over the raw logs and unpacked data for VotingItemAdded events raised by the VoterManager contract.
type VoterManagerVotingItemAddedIterator struct {
	Event *VoterManagerVotingItemAdded // Event containing the contract specifics and raw log

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
func (it *VoterManagerVotingItemAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoterManagerVotingItemAdded)
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
		it.Event = new(VoterManagerVotingItemAdded)
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
func (it *VoterManagerVotingItemAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoterManagerVotingItemAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoterManagerVotingItemAdded represents a VotingItemAdded event raised by the VoterManager contract.
type VoterManagerVotingItemAdded struct {
	OrgId string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterVotingItemAdded is a free log retrieval operation binding the contract event 0x5bfaebb5931145594f63236d2a59314c4dc6035b65d0ca4cee9c7298e2f06ca3.
//
// Solidity: e VotingItemAdded(_orgId string)
func (_VoterManager *VoterManagerFilterer) FilterVotingItemAdded(opts *bind.FilterOpts) (*VoterManagerVotingItemAddedIterator, error) {

	logs, sub, err := _VoterManager.contract.FilterLogs(opts, "VotingItemAdded")
	if err != nil {
		return nil, err
	}
	return &VoterManagerVotingItemAddedIterator{contract: _VoterManager.contract, event: "VotingItemAdded", logs: logs, sub: sub}, nil
}

// WatchVotingItemAdded is a free log subscription operation binding the contract event 0x5bfaebb5931145594f63236d2a59314c4dc6035b65d0ca4cee9c7298e2f06ca3.
//
// Solidity: e VotingItemAdded(_orgId string)
func (_VoterManager *VoterManagerFilterer) WatchVotingItemAdded(opts *bind.WatchOpts, sink chan<- *VoterManagerVotingItemAdded) (event.Subscription, error) {

	logs, sub, err := _VoterManager.contract.WatchLogs(opts, "VotingItemAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoterManagerVotingItemAdded)
				if err := _VoterManager.contract.UnpackLog(event, "VotingItemAdded", log); err != nil {
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
