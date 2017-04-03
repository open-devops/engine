// This file was generated by counterfeiter
package core

import (
	"sync"

	"github.com/opspec-io/sdk-golang/model"
)

type fakeParallelCaller struct {
	CallStub        func(callId string, inboundScope map[string]*model.Data, rootOpId string, pkgRef string, scgParallelCall []*model.SCG) (err error)
	callMutex       sync.RWMutex
	callArgsForCall []struct {
		callId          string
		inboundScope    map[string]*model.Data
		rootOpId        string
		pkgRef          string
		scgParallelCall []*model.SCG
	}
	callReturns struct {
		result1 error
	}
	callReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *fakeParallelCaller) Call(callId string, inboundScope map[string]*model.Data, rootOpId string, pkgRef string, scgParallelCall []*model.SCG) (err error) {
	var scgParallelCallCopy []*model.SCG
	if scgParallelCall != nil {
		scgParallelCallCopy = make([]*model.SCG, len(scgParallelCall))
		copy(scgParallelCallCopy, scgParallelCall)
	}
	fake.callMutex.Lock()
	ret, specificReturn := fake.callReturnsOnCall[len(fake.callArgsForCall)]
	fake.callArgsForCall = append(fake.callArgsForCall, struct {
		callId          string
		inboundScope    map[string]*model.Data
		rootOpId        string
		pkgRef          string
		scgParallelCall []*model.SCG
	}{callId, inboundScope, rootOpId, pkgRef, scgParallelCallCopy})
	fake.recordInvocation("Call", []interface{}{callId, inboundScope, rootOpId, pkgRef, scgParallelCallCopy})
	fake.callMutex.Unlock()
	if fake.CallStub != nil {
		return fake.CallStub(callId, inboundScope, rootOpId, pkgRef, scgParallelCall)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.callReturns.result1
}

func (fake *fakeParallelCaller) CallCallCount() int {
	fake.callMutex.RLock()
	defer fake.callMutex.RUnlock()
	return len(fake.callArgsForCall)
}

func (fake *fakeParallelCaller) CallArgsForCall(i int) (string, map[string]*model.Data, string, string, []*model.SCG) {
	fake.callMutex.RLock()
	defer fake.callMutex.RUnlock()
	return fake.callArgsForCall[i].callId, fake.callArgsForCall[i].inboundScope, fake.callArgsForCall[i].rootOpId, fake.callArgsForCall[i].pkgRef, fake.callArgsForCall[i].scgParallelCall
}

func (fake *fakeParallelCaller) CallReturns(result1 error) {
	fake.CallStub = nil
	fake.callReturns = struct {
		result1 error
	}{result1}
}

func (fake *fakeParallelCaller) CallReturnsOnCall(i int, result1 error) {
	fake.CallStub = nil
	if fake.callReturnsOnCall == nil {
		fake.callReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.callReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *fakeParallelCaller) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.callMutex.RLock()
	defer fake.callMutex.RUnlock()
	return fake.invocations
}

func (fake *fakeParallelCaller) recordInvocation(key string, args []interface{}) {
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
