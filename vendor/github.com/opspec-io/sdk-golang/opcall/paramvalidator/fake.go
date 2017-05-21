// Code generated by counterfeiter. DO NOT EDIT.
package paramvalidator

import (
	"sync"

	"github.com/opspec-io/sdk-golang/model"
)

type Fake struct {
	ValidateStub        func(value *model.Data, param *model.Param) (errors []error)
	validateMutex       sync.RWMutex
	validateArgsForCall []struct {
		value *model.Data
		param *model.Param
	}
	validateReturns struct {
		result1 []error
	}
	validateReturnsOnCall map[int]struct {
		result1 []error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Fake) Validate(value *model.Data, param *model.Param) (errors []error) {
	fake.validateMutex.Lock()
	ret, specificReturn := fake.validateReturnsOnCall[len(fake.validateArgsForCall)]
	fake.validateArgsForCall = append(fake.validateArgsForCall, struct {
		value *model.Data
		param *model.Param
	}{value, param})
	fake.recordInvocation("Validate", []interface{}{value, param})
	fake.validateMutex.Unlock()
	if fake.ValidateStub != nil {
		return fake.ValidateStub(value, param)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.validateReturns.result1
}

func (fake *Fake) ValidateCallCount() int {
	fake.validateMutex.RLock()
	defer fake.validateMutex.RUnlock()
	return len(fake.validateArgsForCall)
}

func (fake *Fake) ValidateArgsForCall(i int) (*model.Data, *model.Param) {
	fake.validateMutex.RLock()
	defer fake.validateMutex.RUnlock()
	return fake.validateArgsForCall[i].value, fake.validateArgsForCall[i].param
}

func (fake *Fake) ValidateReturns(result1 []error) {
	fake.ValidateStub = nil
	fake.validateReturns = struct {
		result1 []error
	}{result1}
}

func (fake *Fake) ValidateReturnsOnCall(i int, result1 []error) {
	fake.ValidateStub = nil
	if fake.validateReturnsOnCall == nil {
		fake.validateReturnsOnCall = make(map[int]struct {
			result1 []error
		})
	}
	fake.validateReturnsOnCall[i] = struct {
		result1 []error
	}{result1}
}

func (fake *Fake) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.validateMutex.RLock()
	defer fake.validateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Fake) recordInvocation(key string, args []interface{}) {
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

var _ ParamValidator = new(Fake)
