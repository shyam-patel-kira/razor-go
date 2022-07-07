// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	big "math/big"
	bindings "razor/pkg/bindings"

	bind "github.com/ethereum/go-ethereum/accounts/abi/bind"

	common "github.com/ethereum/go-ethereum/common"

	ethclient "github.com/ethereum/go-ethereum/ethclient"

	mock "github.com/stretchr/testify/mock"

	pflag "github.com/spf13/pflag"

	types "razor/core/types"
)

// UtilsInterface is an autogenerated mock type for the UtilsInterface type
type UtilsInterface struct {
	mock.Mock
}

// AddJobToJSON provides a mock function with given fields: s, job
func (_m *UtilsInterface) AddJobToJSON(s string, job *types.StructsJob) error {
	ret := _m.Called(s, job)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *types.StructsJob) error); ok {
		r0 = rf(s, job)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AllZero provides a mock function with given fields: bytesValue
func (_m *UtilsInterface) AllZero(bytesValue [32]byte) bool {
	ret := _m.Called(bytesValue)

	var r0 bool
	if rf, ok := ret.Get(0).(func([32]byte) bool); ok {
		r0 = rf(bytesValue)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// AssignLogFile provides a mock function with given fields: flagSet
func (_m *UtilsInterface) AssignLogFile(flagSet *pflag.FlagSet) {
	_m.Called(flagSet)
}

// AssignPassword provides a mock function with given fields:
func (_m *UtilsInterface) AssignPassword() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// AssignStakerId provides a mock function with given fields: flagSet, client, address
func (_m *UtilsInterface) AssignStakerId(flagSet *pflag.FlagSet, client *ethclient.Client, address string) (uint32, error) {
	ret := _m.Called(flagSet, client, address)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet, *ethclient.Client, string) uint32); ok {
		r0 = rf(flagSet, client, address)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet, *ethclient.Client, string) error); ok {
		r1 = rf(flagSet, client, address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CalculateBlockTime provides a mock function with given fields: client
func (_m *UtilsInterface) CalculateBlockTime(client *ethclient.Client) int64 {
	ret := _m.Called(client)

	var r0 int64
	if rf, ok := ret.Get(0).(func(*ethclient.Client) int64); ok {
		r0 = rf(client)
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// CheckAmountAndBalance provides a mock function with given fields: amountInWei, balance
func (_m *UtilsInterface) CheckAmountAndBalance(amountInWei *big.Int, balance *big.Int) *big.Int {
	ret := _m.Called(amountInWei, balance)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*big.Int, *big.Int) *big.Int); ok {
		r0 = rf(amountInWei, balance)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	return r0
}

// CheckEthBalanceIsZero provides a mock function with given fields: client, address
func (_m *UtilsInterface) CheckEthBalanceIsZero(client *ethclient.Client, address string) {
	_m.Called(client, address)
}

// ConnectToClient provides a mock function with given fields: provider
func (_m *UtilsInterface) ConnectToClient(provider string) *ethclient.Client {
	ret := _m.Called(provider)

	var r0 *ethclient.Client
	if rf, ok := ret.Get(0).(func(string) *ethclient.Client); ok {
		r0 = rf(provider)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ethclient.Client)
		}
	}

	return r0
}

// ConvertRZRToSRZR provides a mock function with given fields: sAmount, currentStake, totalSupply
func (_m *UtilsInterface) ConvertRZRToSRZR(sAmount *big.Int, currentStake *big.Int, totalSupply *big.Int) (*big.Int, error) {
	ret := _m.Called(sAmount, currentStake, totalSupply)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*big.Int, *big.Int, *big.Int) *big.Int); ok {
		r0 = rf(sAmount, currentStake, totalSupply)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*big.Int, *big.Int, *big.Int) error); ok {
		r1 = rf(sAmount, currentStake, totalSupply)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ConvertSRZRToRZR provides a mock function with given fields: sAmount, currentStake, totalSupply
func (_m *UtilsInterface) ConvertSRZRToRZR(sAmount *big.Int, currentStake *big.Int, totalSupply *big.Int) *big.Int {
	ret := _m.Called(sAmount, currentStake, totalSupply)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*big.Int, *big.Int, *big.Int) *big.Int); ok {
		r0 = rf(sAmount, currentStake, totalSupply)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	return r0
}

// ConvertUint32ArrayToBigIntArray provides a mock function with given fields: uint32Array
func (_m *UtilsInterface) ConvertUint32ArrayToBigIntArray(uint32Array []uint32) []*big.Int {
	ret := _m.Called(uint32Array)

	var r0 []*big.Int
	if rf, ok := ret.Get(0).(func([]uint32) []*big.Int); ok {
		r0 = rf(uint32Array)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*big.Int)
		}
	}

	return r0
}

// ConvertUintArrayToUint16Array provides a mock function with given fields: uintArr
func (_m *UtilsInterface) ConvertUintArrayToUint16Array(uintArr []uint) []uint16 {
	ret := _m.Called(uintArr)

	var r0 []uint16
	if rf, ok := ret.Get(0).(func([]uint) []uint16); ok {
		r0 = rf(uintArr)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]uint16)
		}
	}

	return r0
}

// ConvertWeiToEth provides a mock function with given fields: data
func (_m *UtilsInterface) ConvertWeiToEth(data *big.Int) (*big.Float, error) {
	ret := _m.Called(data)

	var r0 *big.Float
	if rf, ok := ret.Get(0).(func(*big.Int) *big.Float); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Float)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*big.Int) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteJobFromJSON provides a mock function with given fields: s, jobId
func (_m *UtilsInterface) DeleteJobFromJSON(s string, jobId string) error {
	ret := _m.Called(s, jobId)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(s, jobId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FetchBalance provides a mock function with given fields: client, accountAddress
func (_m *UtilsInterface) FetchBalance(client *ethclient.Client, accountAddress string) (*big.Int, error) {
	ret := _m.Called(client, accountAddress)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*ethclient.Client, string) *big.Int); ok {
		r0 = rf(client, accountAddress)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, string) error); ok {
		r1 = rf(client, accountAddress)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetActiveCollections provides a mock function with given fields: client
func (_m *UtilsInterface) GetActiveCollections(client *ethclient.Client) ([]uint16, error) {
	ret := _m.Called(client)

	var r0 []uint16
	if rf, ok := ret.Get(0).(func(*ethclient.Client) []uint16); ok {
		r0 = rf(client)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]uint16)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client) error); ok {
		r1 = rf(client)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAggregatedDataOfCollection provides a mock function with given fields: client, collectionId, epoch
func (_m *UtilsInterface) GetAggregatedDataOfCollection(client *ethclient.Client, collectionId uint16, epoch uint32) (*big.Int, error) {
	ret := _m.Called(client, collectionId, epoch)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*ethclient.Client, uint16, uint32) *big.Int); ok {
		r0 = rf(client, collectionId, epoch)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, uint16, uint32) error); ok {
		r1 = rf(client, collectionId, epoch)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAmountInDecimal provides a mock function with given fields: amountInWei
func (_m *UtilsInterface) GetAmountInDecimal(amountInWei *big.Int) *big.Float {
	ret := _m.Called(amountInWei)

	var r0 *big.Float
	if rf, ok := ret.Get(0).(func(*big.Int) *big.Float); ok {
		r0 = rf(amountInWei)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Float)
		}
	}

	return r0
}

// GetAmountInWei provides a mock function with given fields: amount
func (_m *UtilsInterface) GetAmountInWei(amount *big.Int) *big.Int {
	ret := _m.Called(amount)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*big.Int) *big.Int); ok {
		r0 = rf(amount)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	return r0
}

// GetBlockManager provides a mock function with given fields: client
func (_m *UtilsInterface) GetBlockManager(client *ethclient.Client) *bindings.BlockManager {
	ret := _m.Called(client)

	var r0 *bindings.BlockManager
	if rf, ok := ret.Get(0).(func(*ethclient.Client) *bindings.BlockManager); ok {
		r0 = rf(client)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bindings.BlockManager)
		}
	}

	return r0
}

// GetCollections provides a mock function with given fields: client
func (_m *UtilsInterface) GetCollections(client *ethclient.Client) ([]bindings.StructsCollection, error) {
	ret := _m.Called(client)

	var r0 []bindings.StructsCollection
	if rf, ok := ret.Get(0).(func(*ethclient.Client) []bindings.StructsCollection); ok {
		r0 = rf(client)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]bindings.StructsCollection)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client) error); ok {
		r1 = rf(client)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCommitDataFileName provides a mock function with given fields: address
func (_m *UtilsInterface) GetCommitDataFileName(address string) (string, error) {
	ret := _m.Called(address)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(address)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCommitments provides a mock function with given fields: client, address
func (_m *UtilsInterface) GetCommitments(client *ethclient.Client, address string) ([32]byte, error) {
	ret := _m.Called(client, address)

	var r0 [32]byte
	if rf, ok := ret.Get(0).(func(*ethclient.Client, string) [32]byte); ok {
		r0 = rf(client, address)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([32]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, string) error); ok {
		r1 = rf(client, address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetConfigFilePath provides a mock function with given fields:
func (_m *UtilsInterface) GetConfigFilePath() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDefaultPath provides a mock function with given fields:
func (_m *UtilsInterface) GetDefaultPath() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDelayedState provides a mock function with given fields: client, buffer
func (_m *UtilsInterface) GetDelayedState(client *ethclient.Client, buffer int32) (int64, error) {
	ret := _m.Called(client, buffer)

	var r0 int64
	if rf, ok := ret.Get(0).(func(*ethclient.Client, int32) int64); ok {
		r0 = rf(client, buffer)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, int32) error); ok {
		r1 = rf(client, buffer)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDisputeDataFileName provides a mock function with given fields: address
func (_m *UtilsInterface) GetDisputeDataFileName(address string) (string, error) {
	ret := _m.Called(address)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(address)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEpoch provides a mock function with given fields: client
func (_m *UtilsInterface) GetEpoch(client *ethclient.Client) (uint32, error) {
	ret := _m.Called(client)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(*ethclient.Client) uint32); ok {
		r0 = rf(client)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client) error); ok {
		r1 = rf(client)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEpochLastCommitted provides a mock function with given fields: client, stakerId
func (_m *UtilsInterface) GetEpochLastCommitted(client *ethclient.Client, stakerId uint32) (uint32, error) {
	ret := _m.Called(client, stakerId)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(*ethclient.Client, uint32) uint32); ok {
		r0 = rf(client, stakerId)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, uint32) error); ok {
		r1 = rf(client, stakerId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEpochLastRevealed provides a mock function with given fields: client, stakerId
func (_m *UtilsInterface) GetEpochLastRevealed(client *ethclient.Client, stakerId uint32) (uint32, error) {
	ret := _m.Called(client, stakerId)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(*ethclient.Client, uint32) uint32); ok {
		r0 = rf(client, stakerId)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, uint32) error); ok {
		r1 = rf(client, stakerId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEpochLimitForUpdateCommission provides a mock function with given fields: client
func (_m *UtilsInterface) GetEpochLimitForUpdateCommission(client *ethclient.Client) (uint16, error) {
	ret := _m.Called(client)

	var r0 uint16
	if rf, ok := ret.Get(0).(func(*ethclient.Client) uint16); ok {
		r0 = rf(client)
	} else {
		r0 = ret.Get(0).(uint16)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client) error); ok {
		r1 = rf(client)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetInfluenceSnapshot provides a mock function with given fields: client, stakerId, epoch
func (_m *UtilsInterface) GetInfluenceSnapshot(client *ethclient.Client, stakerId uint32, epoch uint32) (*big.Int, error) {
	ret := _m.Called(client, stakerId, epoch)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*ethclient.Client, uint32, uint32) *big.Int); ok {
		r0 = rf(client, stakerId, epoch)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, uint32, uint32) error); ok {
		r1 = rf(client, stakerId, epoch)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetJobFilePath provides a mock function with given fields:
func (_m *UtilsInterface) GetJobFilePath() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetJobs provides a mock function with given fields: client
func (_m *UtilsInterface) GetJobs(client *ethclient.Client) ([]bindings.StructsJob, error) {
	ret := _m.Called(client)

	var r0 []bindings.StructsJob
	if rf, ok := ret.Get(0).(func(*ethclient.Client) []bindings.StructsJob); ok {
		r0 = rf(client)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]bindings.StructsJob)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client) error); ok {
		r1 = rf(client)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetLock provides a mock function with given fields: client, address, stakerId, lockType
func (_m *UtilsInterface) GetLock(client *ethclient.Client, address string, stakerId uint32, lockType uint8) (types.Locks, error) {
	ret := _m.Called(client, address, stakerId, lockType)

	var r0 types.Locks
	if rf, ok := ret.Get(0).(func(*ethclient.Client, string, uint32, uint8) types.Locks); ok {
		r0 = rf(client, address, stakerId, lockType)
	} else {
		r0 = ret.Get(0).(types.Locks)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, string, uint32, uint8) error); ok {
		r1 = rf(client, address, stakerId, lockType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMaxAltBlocks provides a mock function with given fields: client
func (_m *UtilsInterface) GetMaxAltBlocks(client *ethclient.Client) (uint8, error) {
	ret := _m.Called(client)

	var r0 uint8
	if rf, ok := ret.Get(0).(func(*ethclient.Client) uint8); ok {
		r0 = rf(client)
	} else {
		r0 = ret.Get(0).(uint8)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client) error); ok {
		r1 = rf(client)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMaxCommission provides a mock function with given fields: client
func (_m *UtilsInterface) GetMaxCommission(client *ethclient.Client) (uint8, error) {
	ret := _m.Called(client)

	var r0 uint8
	if rf, ok := ret.Get(0).(func(*ethclient.Client) uint8); ok {
		r0 = rf(client)
	} else {
		r0 = ret.Get(0).(uint8)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client) error); ok {
		r1 = rf(client)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNumActiveCollections provides a mock function with given fields: client
func (_m *UtilsInterface) GetNumActiveCollections(client *ethclient.Client) (uint16, error) {
	ret := _m.Called(client)

	var r0 uint16
	if rf, ok := ret.Get(0).(func(*ethclient.Client) uint16); ok {
		r0 = rf(client)
	} else {
		r0 = ret.Get(0).(uint16)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client) error); ok {
		r1 = rf(client)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNumberOfProposedBlocks provides a mock function with given fields: client, epoch
func (_m *UtilsInterface) GetNumberOfProposedBlocks(client *ethclient.Client, epoch uint32) (uint8, error) {
	ret := _m.Called(client, epoch)

	var r0 uint8
	if rf, ok := ret.Get(0).(func(*ethclient.Client, uint32) uint8); ok {
		r0 = rf(client, epoch)
	} else {
		r0 = ret.Get(0).(uint8)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, uint32) error); ok {
		r1 = rf(client, epoch)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNumberOfStakers provides a mock function with given fields: client
func (_m *UtilsInterface) GetNumberOfStakers(client *ethclient.Client) (uint32, error) {
	ret := _m.Called(client)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(*ethclient.Client) uint32); ok {
		r0 = rf(client)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client) error); ok {
		r1 = rf(client)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOptions provides a mock function with given fields:
func (_m *UtilsInterface) GetOptions() bind.CallOpts {
	ret := _m.Called()

	var r0 bind.CallOpts
	if rf, ok := ret.Get(0).(func() bind.CallOpts); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bind.CallOpts)
	}

	return r0
}

// GetProposeDataFileName provides a mock function with given fields: address
func (_m *UtilsInterface) GetProposeDataFileName(address string) (string, error) {
	ret := _m.Called(address)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(address)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProposedBlock provides a mock function with given fields: client, epoch, proposedBlockId
func (_m *UtilsInterface) GetProposedBlock(client *ethclient.Client, epoch uint32, proposedBlockId uint32) (bindings.StructsBlock, error) {
	ret := _m.Called(client, epoch, proposedBlockId)

	var r0 bindings.StructsBlock
	if rf, ok := ret.Get(0).(func(*ethclient.Client, uint32, uint32) bindings.StructsBlock); ok {
		r0 = rf(client, epoch, proposedBlockId)
	} else {
		r0 = ret.Get(0).(bindings.StructsBlock)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, uint32, uint32) error); ok {
		r1 = rf(client, epoch, proposedBlockId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRogueRandomMedianValue provides a mock function with given fields:
func (_m *UtilsInterface) GetRogueRandomMedianValue() uint32 {
	ret := _m.Called()

	var r0 uint32
	if rf, ok := ret.Get(0).(func() uint32); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(uint32)
	}

	return r0
}

// GetRogueRandomValue provides a mock function with given fields: value
func (_m *UtilsInterface) GetRogueRandomValue(value int) *big.Int {
	ret := _m.Called(value)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(int) *big.Int); ok {
		r0 = rf(value)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	return r0
}

// GetSortedProposedBlockIds provides a mock function with given fields: client, epoch
func (_m *UtilsInterface) GetSortedProposedBlockIds(client *ethclient.Client, epoch uint32) ([]uint32, error) {
	ret := _m.Called(client, epoch)

	var r0 []uint32
	if rf, ok := ret.Get(0).(func(*ethclient.Client, uint32) []uint32); ok {
		r0 = rf(client, epoch)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]uint32)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, uint32) error); ok {
		r1 = rf(client, epoch)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStake provides a mock function with given fields: client, stakerId
func (_m *UtilsInterface) GetStake(client *ethclient.Client, stakerId uint32) (*big.Int, error) {
	ret := _m.Called(client, stakerId)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*ethclient.Client, uint32) *big.Int); ok {
		r0 = rf(client, stakerId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, uint32) error); ok {
		r1 = rf(client, stakerId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStakeSnapshot provides a mock function with given fields: client, stakerId, epoch
func (_m *UtilsInterface) GetStakeSnapshot(client *ethclient.Client, stakerId uint32, epoch uint32) (*big.Int, error) {
	ret := _m.Called(client, stakerId, epoch)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*ethclient.Client, uint32, uint32) *big.Int); ok {
		r0 = rf(client, stakerId, epoch)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, uint32, uint32) error); ok {
		r1 = rf(client, stakerId, epoch)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStakedToken provides a mock function with given fields: client, address
func (_m *UtilsInterface) GetStakedToken(client *ethclient.Client, address common.Address) *bindings.StakedToken {
	ret := _m.Called(client, address)

	var r0 *bindings.StakedToken
	if rf, ok := ret.Get(0).(func(*ethclient.Client, common.Address) *bindings.StakedToken); ok {
		r0 = rf(client, address)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bindings.StakedToken)
		}
	}

	return r0
}

// GetStaker provides a mock function with given fields: client, stakerId
func (_m *UtilsInterface) GetStaker(client *ethclient.Client, stakerId uint32) (bindings.StructsStaker, error) {
	ret := _m.Called(client, stakerId)

	var r0 bindings.StructsStaker
	if rf, ok := ret.Get(0).(func(*ethclient.Client, uint32) bindings.StructsStaker); ok {
		r0 = rf(client, stakerId)
	} else {
		r0 = ret.Get(0).(bindings.StructsStaker)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, uint32) error); ok {
		r1 = rf(client, stakerId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStakerId provides a mock function with given fields: client, address
func (_m *UtilsInterface) GetStakerId(client *ethclient.Client, address string) (uint32, error) {
	ret := _m.Called(client, address)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(*ethclient.Client, string) uint32); ok {
		r0 = rf(client, address)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, string) error); ok {
		r1 = rf(client, address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStakerSRZRBalance provides a mock function with given fields: client, staker
func (_m *UtilsInterface) GetStakerSRZRBalance(client *ethclient.Client, staker bindings.StructsStaker) (*big.Int, error) {
	ret := _m.Called(client, staker)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*ethclient.Client, bindings.StructsStaker) *big.Int); ok {
		r0 = rf(client, staker)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, bindings.StructsStaker) error); ok {
		r1 = rf(client, staker)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStringAddress provides a mock function with given fields: flagSet
func (_m *UtilsInterface) GetStringAddress(flagSet *pflag.FlagSet) (string, error) {
	ret := _m.Called(flagSet)

	var r0 string
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) string); ok {
		r0 = rf(flagSet)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTotalInfluenceRevealed provides a mock function with given fields: client, epoch, medianIndex
func (_m *UtilsInterface) GetTotalInfluenceRevealed(client *ethclient.Client, epoch uint32, medianIndex uint16) (*big.Int, error) {
	ret := _m.Called(client, epoch, medianIndex)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*ethclient.Client, uint32, uint16) *big.Int); ok {
		r0 = rf(client, epoch, medianIndex)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, uint32, uint16) error); ok {
		r1 = rf(client, epoch, medianIndex)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTxnOpts provides a mock function with given fields: transactionData
func (_m *UtilsInterface) GetTxnOpts(transactionData types.TransactionOptions) *bind.TransactOpts {
	ret := _m.Called(transactionData)

	var r0 *bind.TransactOpts
	if rf, ok := ret.Get(0).(func(types.TransactionOptions) *bind.TransactOpts); ok {
		r0 = rf(transactionData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bind.TransactOpts)
		}
	}

	return r0
}

// GetUint32BountyId provides a mock function with given fields: flagSet
func (_m *UtilsInterface) GetUint32BountyId(flagSet *pflag.FlagSet) (uint32, error) {
	ret := _m.Called(flagSet)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) uint32); ok {
		r0 = rf(flagSet)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUpdatedStaker provides a mock function with given fields: client, stakerId
func (_m *UtilsInterface) GetUpdatedStaker(client *ethclient.Client, stakerId uint32) (bindings.StructsStaker, error) {
	ret := _m.Called(client, stakerId)

	var r0 bindings.StructsStaker
	if rf, ok := ret.Get(0).(func(*ethclient.Client, uint32) bindings.StructsStaker); ok {
		r0 = rf(client, stakerId)
	} else {
		r0 = ret.Get(0).(bindings.StructsStaker)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, uint32) error); ok {
		r1 = rf(client, stakerId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetVoteValue provides a mock function with given fields: client, epoch, stakerId, medianIndex
func (_m *UtilsInterface) GetVoteValue(client *ethclient.Client, epoch uint32, stakerId uint32, medianIndex uint16) (*big.Int, error) {
	ret := _m.Called(client, epoch, stakerId, medianIndex)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*ethclient.Client, uint32, uint32, uint16) *big.Int); ok {
		r0 = rf(client, epoch, stakerId, medianIndex)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, uint32, uint32, uint16) error); ok {
		r1 = rf(client, epoch, stakerId, medianIndex)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetWithdrawInitiationPeriod provides a mock function with given fields: client
func (_m *UtilsInterface) GetWithdrawInitiationPeriod(client *ethclient.Client) (uint8, error) {
	ret := _m.Called(client)

	var r0 uint8
	if rf, ok := ret.Get(0).(func(*ethclient.Client) uint8); ok {
		r0 = rf(client)
	} else {
		r0 = ret.Get(0).(uint8)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client) error); ok {
		r1 = rf(client)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsFlagPassed provides a mock function with given fields: name
func (_m *UtilsInterface) IsFlagPassed(name string) bool {
	ret := _m.Called(name)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// PasswordPrompt provides a mock function with given fields:
func (_m *UtilsInterface) PasswordPrompt() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// PrivateKeyPrompt provides a mock function with given fields:
func (_m *UtilsInterface) PrivateKeyPrompt() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ReadFromCommitJsonFile provides a mock function with given fields: filePath
func (_m *UtilsInterface) ReadFromCommitJsonFile(filePath string) (types.CommitFileData, error) {
	ret := _m.Called(filePath)

	var r0 types.CommitFileData
	if rf, ok := ret.Get(0).(func(string) types.CommitFileData); ok {
		r0 = rf(filePath)
	} else {
		r0 = ret.Get(0).(types.CommitFileData)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(filePath)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadFromDisputeJsonFile provides a mock function with given fields: filePath
func (_m *UtilsInterface) ReadFromDisputeJsonFile(filePath string) (types.DisputeFileData, error) {
	ret := _m.Called(filePath)

	var r0 types.DisputeFileData
	if rf, ok := ret.Get(0).(func(string) types.DisputeFileData); ok {
		r0 = rf(filePath)
	} else {
		r0 = ret.Get(0).(types.DisputeFileData)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(filePath)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadFromProposeJsonFile provides a mock function with given fields: filePath
func (_m *UtilsInterface) ReadFromProposeJsonFile(filePath string) (types.ProposeFileData, error) {
	ret := _m.Called(filePath)

	var r0 types.ProposeFileData
	if rf, ok := ret.Get(0).(func(string) types.ProposeFileData); ok {
		r0 = rf(filePath)
	} else {
		r0 = ret.Get(0).(types.ProposeFileData)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(filePath)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveDataToCommitJsonFile provides a mock function with given fields: flePath, epoch, commitFileData
func (_m *UtilsInterface) SaveDataToCommitJsonFile(flePath string, epoch uint32, commitFileData types.CommitData) error {
	ret := _m.Called(flePath, epoch, commitFileData)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, uint32, types.CommitData) error); ok {
		r0 = rf(flePath, epoch, commitFileData)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveDataToDisputeJsonFile provides a mock function with given fields: filePath, bountyIdQueue
func (_m *UtilsInterface) SaveDataToDisputeJsonFile(filePath string, bountyIdQueue []uint32) error {
	ret := _m.Called(filePath, bountyIdQueue)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []uint32) error); ok {
		r0 = rf(filePath, bountyIdQueue)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveDataToProposeJsonFile provides a mock function with given fields: flePath, epoch, proposeFileData
func (_m *UtilsInterface) SaveDataToProposeJsonFile(flePath string, epoch uint32, proposeFileData types.ProposeData) error {
	ret := _m.Called(flePath, epoch, proposeFileData)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, uint32, types.ProposeData) error); ok {
		r0 = rf(flePath, epoch, proposeFileData)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SecondsToReadableTime provides a mock function with given fields: time
func (_m *UtilsInterface) SecondsToReadableTime(time int) string {
	ret := _m.Called(time)

	var r0 string
	if rf, ok := ret.Get(0).(func(int) string); ok {
		r0 = rf(time)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// WaitForBlockCompletion provides a mock function with given fields: client, hashToRead
func (_m *UtilsInterface) WaitForBlockCompletion(client *ethclient.Client, hashToRead string) error {
	ret := _m.Called(client, hashToRead)

	var r0 error
	if rf, ok := ret.Get(0).(func(*ethclient.Client, string) error); ok {
		r0 = rf(client, hashToRead)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitTillNextNSecs provides a mock function with given fields: seconds
func (_m *UtilsInterface) WaitTillNextNSecs(seconds int32) {
	_m.Called(seconds)
}

type mockConstructorTestingTNewUtilsInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewUtilsInterface creates a new instance of UtilsInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUtilsInterface(t mockConstructorTestingTNewUtilsInterface) *UtilsInterface {
	mock := &UtilsInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
