// Code generated by counterfeiter. DO NOT EDIT.
package create

import (
	"sync"

	"github.com/opctl/opctl/cli/model"
)

type FakeInvoker struct {
	InvokeStub        func(model.NodeCreateOpts)
	invokeMutex       sync.RWMutex
	invokeArgsForCall []struct {
		arg1 model.NodeCreateOpts
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeInvoker) Invoke(arg1 model.NodeCreateOpts) {
	fake.invokeMutex.Lock()
	fake.invokeArgsForCall = append(fake.invokeArgsForCall, struct {
		arg1 model.NodeCreateOpts
	}{arg1})
	fake.recordInvocation("Invoke", []interface{}{arg1})
	fake.invokeMutex.Unlock()
	if fake.InvokeStub != nil {
		fake.InvokeStub(arg1)
	}
}

func (fake *FakeInvoker) InvokeCallCount() int {
	fake.invokeMutex.RLock()
	defer fake.invokeMutex.RUnlock()
	return len(fake.invokeArgsForCall)
}

func (fake *FakeInvoker) InvokeCalls(stub func(model.NodeCreateOpts)) {
	fake.invokeMutex.Lock()
	defer fake.invokeMutex.Unlock()
	fake.InvokeStub = stub
}

func (fake *FakeInvoker) InvokeArgsForCall(i int) model.NodeCreateOpts {
	fake.invokeMutex.RLock()
	defer fake.invokeMutex.RUnlock()
	argsForCall := fake.invokeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeInvoker) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.invokeMutex.RLock()
	defer fake.invokeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeInvoker) recordInvocation(key string, args []interface{}) {
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

var _ Invoker = new(FakeInvoker)