// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/docker/go-connections/nat"
)

type FakePortBindingsFactory struct {
	ConstructStub        func(map[string]string) (nat.PortMap, error)
	constructMutex       sync.RWMutex
	constructArgsForCall []struct {
		arg1 map[string]string
	}
	constructReturns struct {
		result1 nat.PortMap
		result2 error
	}
	constructReturnsOnCall map[int]struct {
		result1 nat.PortMap
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePortBindingsFactory) Construct(arg1 map[string]string) (nat.PortMap, error) {
	fake.constructMutex.Lock()
	ret, specificReturn := fake.constructReturnsOnCall[len(fake.constructArgsForCall)]
	fake.constructArgsForCall = append(fake.constructArgsForCall, struct {
		arg1 map[string]string
	}{arg1})
	fake.recordInvocation("Construct", []interface{}{arg1})
	fake.constructMutex.Unlock()
	if fake.ConstructStub != nil {
		return fake.ConstructStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.constructReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePortBindingsFactory) ConstructCallCount() int {
	fake.constructMutex.RLock()
	defer fake.constructMutex.RUnlock()
	return len(fake.constructArgsForCall)
}

func (fake *FakePortBindingsFactory) ConstructCalls(stub func(map[string]string) (nat.PortMap, error)) {
	fake.constructMutex.Lock()
	defer fake.constructMutex.Unlock()
	fake.ConstructStub = stub
}

func (fake *FakePortBindingsFactory) ConstructArgsForCall(i int) map[string]string {
	fake.constructMutex.RLock()
	defer fake.constructMutex.RUnlock()
	argsForCall := fake.constructArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakePortBindingsFactory) ConstructReturns(result1 nat.PortMap, result2 error) {
	fake.constructMutex.Lock()
	defer fake.constructMutex.Unlock()
	fake.ConstructStub = nil
	fake.constructReturns = struct {
		result1 nat.PortMap
		result2 error
	}{result1, result2}
}

func (fake *FakePortBindingsFactory) ConstructReturnsOnCall(i int, result1 nat.PortMap, result2 error) {
	fake.constructMutex.Lock()
	defer fake.constructMutex.Unlock()
	fake.ConstructStub = nil
	if fake.constructReturnsOnCall == nil {
		fake.constructReturnsOnCall = make(map[int]struct {
			result1 nat.PortMap
			result2 error
		})
	}
	fake.constructReturnsOnCall[i] = struct {
		result1 nat.PortMap
		result2 error
	}{result1, result2}
}

func (fake *FakePortBindingsFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.constructMutex.RLock()
	defer fake.constructMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePortBindingsFactory) recordInvocation(key string, args []interface{}) {
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