// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	"github.com/hyperledger/fabric/orderer/consensus/hcs/protodef"
)

type AppMsgProcessor struct {
	ExpireByAgeStub        func(uint64) int
	expireByAgeMutex       sync.RWMutex
	expireByAgeArgsForCall []struct {
		arg1 uint64
	}
	expireByAgeReturns struct {
		result1 int
	}
	expireByAgeReturnsOnCall map[int]struct {
		result1 int
	}
	ExpireByAppIDStub        func([]byte) (int, error)
	expireByAppIDMutex       sync.RWMutex
	expireByAppIDArgsForCall []struct {
		arg1 []byte
	}
	expireByAppIDReturns struct {
		result1 int
		result2 error
	}
	expireByAppIDReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	IsPendingStub        func() bool
	isPendingMutex       sync.RWMutex
	isPendingArgsForCall []struct {
	}
	isPendingReturns struct {
		result1 bool
	}
	isPendingReturnsOnCall map[int]struct {
		result1 bool
	}
	ReassembleStub        func(*protodef.ApplicationMessageChunk) ([]byte, error)
	reassembleMutex       sync.RWMutex
	reassembleArgsForCall []struct {
		arg1 *protodef.ApplicationMessageChunk
	}
	reassembleReturns struct {
		result1 []byte
		result2 error
	}
	reassembleReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	SplitStub        func([]byte) ([]*protodef.ApplicationMessageChunk, error)
	splitMutex       sync.RWMutex
	splitArgsForCall []struct {
		arg1 []byte
	}
	splitReturns struct {
		result1 []*protodef.ApplicationMessageChunk
		result2 error
	}
	splitReturnsOnCall map[int]struct {
		result1 []*protodef.ApplicationMessageChunk
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *AppMsgProcessor) ExpireByAge(arg1 uint64) int {
	fake.expireByAgeMutex.Lock()
	ret, specificReturn := fake.expireByAgeReturnsOnCall[len(fake.expireByAgeArgsForCall)]
	fake.expireByAgeArgsForCall = append(fake.expireByAgeArgsForCall, struct {
		arg1 uint64
	}{arg1})
	fake.recordInvocation("ExpireByAge", []interface{}{arg1})
	fake.expireByAgeMutex.Unlock()
	if fake.ExpireByAgeStub != nil {
		return fake.ExpireByAgeStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.expireByAgeReturns
	return fakeReturns.result1
}

func (fake *AppMsgProcessor) ExpireByAgeCallCount() int {
	fake.expireByAgeMutex.RLock()
	defer fake.expireByAgeMutex.RUnlock()
	return len(fake.expireByAgeArgsForCall)
}

func (fake *AppMsgProcessor) ExpireByAgeCalls(stub func(uint64) int) {
	fake.expireByAgeMutex.Lock()
	defer fake.expireByAgeMutex.Unlock()
	fake.ExpireByAgeStub = stub
}

func (fake *AppMsgProcessor) ExpireByAgeArgsForCall(i int) uint64 {
	fake.expireByAgeMutex.RLock()
	defer fake.expireByAgeMutex.RUnlock()
	argsForCall := fake.expireByAgeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *AppMsgProcessor) ExpireByAgeReturns(result1 int) {
	fake.expireByAgeMutex.Lock()
	defer fake.expireByAgeMutex.Unlock()
	fake.ExpireByAgeStub = nil
	fake.expireByAgeReturns = struct {
		result1 int
	}{result1}
}

func (fake *AppMsgProcessor) ExpireByAgeReturnsOnCall(i int, result1 int) {
	fake.expireByAgeMutex.Lock()
	defer fake.expireByAgeMutex.Unlock()
	fake.ExpireByAgeStub = nil
	if fake.expireByAgeReturnsOnCall == nil {
		fake.expireByAgeReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.expireByAgeReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *AppMsgProcessor) ExpireByAppID(arg1 []byte) (int, error) {
	var arg1Copy []byte
	if arg1 != nil {
		arg1Copy = make([]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.expireByAppIDMutex.Lock()
	ret, specificReturn := fake.expireByAppIDReturnsOnCall[len(fake.expireByAppIDArgsForCall)]
	fake.expireByAppIDArgsForCall = append(fake.expireByAppIDArgsForCall, struct {
		arg1 []byte
	}{arg1Copy})
	fake.recordInvocation("ExpireByAppID", []interface{}{arg1Copy})
	fake.expireByAppIDMutex.Unlock()
	if fake.ExpireByAppIDStub != nil {
		return fake.ExpireByAppIDStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.expireByAppIDReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *AppMsgProcessor) ExpireByAppIDCallCount() int {
	fake.expireByAppIDMutex.RLock()
	defer fake.expireByAppIDMutex.RUnlock()
	return len(fake.expireByAppIDArgsForCall)
}

func (fake *AppMsgProcessor) ExpireByAppIDCalls(stub func([]byte) (int, error)) {
	fake.expireByAppIDMutex.Lock()
	defer fake.expireByAppIDMutex.Unlock()
	fake.ExpireByAppIDStub = stub
}

func (fake *AppMsgProcessor) ExpireByAppIDArgsForCall(i int) []byte {
	fake.expireByAppIDMutex.RLock()
	defer fake.expireByAppIDMutex.RUnlock()
	argsForCall := fake.expireByAppIDArgsForCall[i]
	return argsForCall.arg1
}

func (fake *AppMsgProcessor) ExpireByAppIDReturns(result1 int, result2 error) {
	fake.expireByAppIDMutex.Lock()
	defer fake.expireByAppIDMutex.Unlock()
	fake.ExpireByAppIDStub = nil
	fake.expireByAppIDReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *AppMsgProcessor) ExpireByAppIDReturnsOnCall(i int, result1 int, result2 error) {
	fake.expireByAppIDMutex.Lock()
	defer fake.expireByAppIDMutex.Unlock()
	fake.ExpireByAppIDStub = nil
	if fake.expireByAppIDReturnsOnCall == nil {
		fake.expireByAppIDReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.expireByAppIDReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *AppMsgProcessor) IsPending() bool {
	fake.isPendingMutex.Lock()
	ret, specificReturn := fake.isPendingReturnsOnCall[len(fake.isPendingArgsForCall)]
	fake.isPendingArgsForCall = append(fake.isPendingArgsForCall, struct {
	}{})
	fake.recordInvocation("IsPending", []interface{}{})
	fake.isPendingMutex.Unlock()
	if fake.IsPendingStub != nil {
		return fake.IsPendingStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.isPendingReturns
	return fakeReturns.result1
}

func (fake *AppMsgProcessor) IsPendingCallCount() int {
	fake.isPendingMutex.RLock()
	defer fake.isPendingMutex.RUnlock()
	return len(fake.isPendingArgsForCall)
}

func (fake *AppMsgProcessor) IsPendingCalls(stub func() bool) {
	fake.isPendingMutex.Lock()
	defer fake.isPendingMutex.Unlock()
	fake.IsPendingStub = stub
}

func (fake *AppMsgProcessor) IsPendingReturns(result1 bool) {
	fake.isPendingMutex.Lock()
	defer fake.isPendingMutex.Unlock()
	fake.IsPendingStub = nil
	fake.isPendingReturns = struct {
		result1 bool
	}{result1}
}

func (fake *AppMsgProcessor) IsPendingReturnsOnCall(i int, result1 bool) {
	fake.isPendingMutex.Lock()
	defer fake.isPendingMutex.Unlock()
	fake.IsPendingStub = nil
	if fake.isPendingReturnsOnCall == nil {
		fake.isPendingReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isPendingReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *AppMsgProcessor) Reassemble(arg1 *protodef.ApplicationMessageChunk) ([]byte, error) {
	fake.reassembleMutex.Lock()
	ret, specificReturn := fake.reassembleReturnsOnCall[len(fake.reassembleArgsForCall)]
	fake.reassembleArgsForCall = append(fake.reassembleArgsForCall, struct {
		arg1 *protodef.ApplicationMessageChunk
	}{arg1})
	fake.recordInvocation("Reassemble", []interface{}{arg1})
	fake.reassembleMutex.Unlock()
	if fake.ReassembleStub != nil {
		return fake.ReassembleStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.reassembleReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *AppMsgProcessor) ReassembleCallCount() int {
	fake.reassembleMutex.RLock()
	defer fake.reassembleMutex.RUnlock()
	return len(fake.reassembleArgsForCall)
}

func (fake *AppMsgProcessor) ReassembleCalls(stub func(*protodef.ApplicationMessageChunk) ([]byte, error)) {
	fake.reassembleMutex.Lock()
	defer fake.reassembleMutex.Unlock()
	fake.ReassembleStub = stub
}

func (fake *AppMsgProcessor) ReassembleArgsForCall(i int) *protodef.ApplicationMessageChunk {
	fake.reassembleMutex.RLock()
	defer fake.reassembleMutex.RUnlock()
	argsForCall := fake.reassembleArgsForCall[i]
	return argsForCall.arg1
}

func (fake *AppMsgProcessor) ReassembleReturns(result1 []byte, result2 error) {
	fake.reassembleMutex.Lock()
	defer fake.reassembleMutex.Unlock()
	fake.ReassembleStub = nil
	fake.reassembleReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *AppMsgProcessor) ReassembleReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.reassembleMutex.Lock()
	defer fake.reassembleMutex.Unlock()
	fake.ReassembleStub = nil
	if fake.reassembleReturnsOnCall == nil {
		fake.reassembleReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.reassembleReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *AppMsgProcessor) Split(arg1 []byte) ([]*protodef.ApplicationMessageChunk, error) {
	var arg1Copy []byte
	if arg1 != nil {
		arg1Copy = make([]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.splitMutex.Lock()
	ret, specificReturn := fake.splitReturnsOnCall[len(fake.splitArgsForCall)]
	fake.splitArgsForCall = append(fake.splitArgsForCall, struct {
		arg1 []byte
	}{arg1Copy})
	fake.recordInvocation("Split", []interface{}{arg1Copy})
	fake.splitMutex.Unlock()
	if fake.SplitStub != nil {
		return fake.SplitStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.splitReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *AppMsgProcessor) SplitCallCount() int {
	fake.splitMutex.RLock()
	defer fake.splitMutex.RUnlock()
	return len(fake.splitArgsForCall)
}

func (fake *AppMsgProcessor) SplitCalls(stub func([]byte) ([]*protodef.ApplicationMessageChunk, error)) {
	fake.splitMutex.Lock()
	defer fake.splitMutex.Unlock()
	fake.SplitStub = stub
}

func (fake *AppMsgProcessor) SplitArgsForCall(i int) []byte {
	fake.splitMutex.RLock()
	defer fake.splitMutex.RUnlock()
	argsForCall := fake.splitArgsForCall[i]
	return argsForCall.arg1
}

func (fake *AppMsgProcessor) SplitReturns(result1 []*protodef.ApplicationMessageChunk, result2 error) {
	fake.splitMutex.Lock()
	defer fake.splitMutex.Unlock()
	fake.SplitStub = nil
	fake.splitReturns = struct {
		result1 []*protodef.ApplicationMessageChunk
		result2 error
	}{result1, result2}
}

func (fake *AppMsgProcessor) SplitReturnsOnCall(i int, result1 []*protodef.ApplicationMessageChunk, result2 error) {
	fake.splitMutex.Lock()
	defer fake.splitMutex.Unlock()
	fake.SplitStub = nil
	if fake.splitReturnsOnCall == nil {
		fake.splitReturnsOnCall = make(map[int]struct {
			result1 []*protodef.ApplicationMessageChunk
			result2 error
		})
	}
	fake.splitReturnsOnCall[i] = struct {
		result1 []*protodef.ApplicationMessageChunk
		result2 error
	}{result1, result2}
}

func (fake *AppMsgProcessor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.expireByAgeMutex.RLock()
	defer fake.expireByAgeMutex.RUnlock()
	fake.expireByAppIDMutex.RLock()
	defer fake.expireByAppIDMutex.RUnlock()
	fake.isPendingMutex.RLock()
	defer fake.isPendingMutex.RUnlock()
	fake.reassembleMutex.RLock()
	defer fake.reassembleMutex.RUnlock()
	fake.splitMutex.RLock()
	defer fake.splitMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *AppMsgProcessor) recordInvocation(key string, args []interface{}) {
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