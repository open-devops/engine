// This file was generated by counterfeiter
package pkg

import (
	"sync"
)

type fakeRefResolver struct {
	ResolveStub        func(pkgRef string) string
	resolveMutex       sync.RWMutex
	resolveArgsForCall []struct {
		pkgRef string
	}
	resolveReturns struct {
		result1 string
	}
	resolveReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *fakeRefResolver) Resolve(pkgRef string) string {
	fake.resolveMutex.Lock()
	ret, specificReturn := fake.resolveReturnsOnCall[len(fake.resolveArgsForCall)]
	fake.resolveArgsForCall = append(fake.resolveArgsForCall, struct {
		pkgRef string
	}{pkgRef})
	fake.recordInvocation("Resolve", []interface{}{pkgRef})
	fake.resolveMutex.Unlock()
	if fake.ResolveStub != nil {
		return fake.ResolveStub(pkgRef)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.resolveReturns.result1
}

func (fake *fakeRefResolver) ResolveCallCount() int {
	fake.resolveMutex.RLock()
	defer fake.resolveMutex.RUnlock()
	return len(fake.resolveArgsForCall)
}

func (fake *fakeRefResolver) ResolveArgsForCall(i int) string {
	fake.resolveMutex.RLock()
	defer fake.resolveMutex.RUnlock()
	return fake.resolveArgsForCall[i].pkgRef
}

func (fake *fakeRefResolver) ResolveReturns(result1 string) {
	fake.ResolveStub = nil
	fake.resolveReturns = struct {
		result1 string
	}{result1}
}

func (fake *fakeRefResolver) ResolveReturnsOnCall(i int, result1 string) {
	fake.ResolveStub = nil
	if fake.resolveReturnsOnCall == nil {
		fake.resolveReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.resolveReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *fakeRefResolver) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.resolveMutex.RLock()
	defer fake.resolveMutex.RUnlock()
	return fake.invocations
}

func (fake *fakeRefResolver) recordInvocation(key string, args []interface{}) {
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
