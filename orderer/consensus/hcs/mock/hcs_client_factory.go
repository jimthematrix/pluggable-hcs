// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	hedera "github.com/hashgraph/hedera-sdk-go/v2"
	"github.com/hyperledger/fabric/orderer/consensus/hcs/factory"
)

type HcsClientFactory struct {
	GetHcsClientStub        func(map[string]hedera.AccountID, string, *hedera.AccountID, *hedera.PrivateKey) (factory.HcsClient, error)
	getHcsClientMutex       sync.RWMutex
	getHcsClientArgsForCall []struct {
		arg1 map[string]hedera.AccountID
		arg2 string
		arg3 *hedera.AccountID
		arg4 *hedera.PrivateKey
	}
	getHcsClientReturns struct {
		result1 factory.HcsClient
		result2 error
	}
	getHcsClientReturnsOnCall map[int]struct {
		result1 factory.HcsClient
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *HcsClientFactory) GetHcsClient(arg1 map[string]hedera.AccountID, arg2 string, arg3 *hedera.AccountID, arg4 *hedera.PrivateKey) (factory.HcsClient, error) {
	fake.getHcsClientMutex.Lock()
	ret, specificReturn := fake.getHcsClientReturnsOnCall[len(fake.getHcsClientArgsForCall)]
	fake.getHcsClientArgsForCall = append(fake.getHcsClientArgsForCall, struct {
		arg1 map[string]hedera.AccountID
		arg2 string
		arg3 *hedera.AccountID
		arg4 *hedera.PrivateKey
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("GetHcsClient", []interface{}{arg1, arg2, arg3, arg4})
	fake.getHcsClientMutex.Unlock()
	if fake.GetHcsClientStub != nil {
		return fake.GetHcsClientStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getHcsClientReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *HcsClientFactory) GetHcsClientCallCount() int {
	fake.getHcsClientMutex.RLock()
	defer fake.getHcsClientMutex.RUnlock()
	return len(fake.getHcsClientArgsForCall)
}

func (fake *HcsClientFactory) GetHcsClientCalls(stub func(map[string]hedera.AccountID, string, *hedera.AccountID, *hedera.PrivateKey) (factory.HcsClient, error)) {
	fake.getHcsClientMutex.Lock()
	defer fake.getHcsClientMutex.Unlock()
	fake.GetHcsClientStub = stub
}

func (fake *HcsClientFactory) GetHcsClientArgsForCall(i int) (map[string]hedera.AccountID, string, *hedera.AccountID, *hedera.PrivateKey) {
	fake.getHcsClientMutex.RLock()
	defer fake.getHcsClientMutex.RUnlock()
	argsForCall := fake.getHcsClientArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *HcsClientFactory) GetHcsClientReturns(result1 factory.HcsClient, result2 error) {
	fake.getHcsClientMutex.Lock()
	defer fake.getHcsClientMutex.Unlock()
	fake.GetHcsClientStub = nil
	fake.getHcsClientReturns = struct {
		result1 factory.HcsClient
		result2 error
	}{result1, result2}
}

func (fake *HcsClientFactory) GetHcsClientReturnsOnCall(i int, result1 factory.HcsClient, result2 error) {
	fake.getHcsClientMutex.Lock()
	defer fake.getHcsClientMutex.Unlock()
	fake.GetHcsClientStub = nil
	if fake.getHcsClientReturnsOnCall == nil {
		fake.getHcsClientReturnsOnCall = make(map[int]struct {
			result1 factory.HcsClient
			result2 error
		})
	}
	fake.getHcsClientReturnsOnCall[i] = struct {
		result1 factory.HcsClient
		result2 error
	}{result1, result2}
}

func (fake *HcsClientFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getHcsClientMutex.RLock()
	defer fake.getHcsClientMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *HcsClientFactory) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}