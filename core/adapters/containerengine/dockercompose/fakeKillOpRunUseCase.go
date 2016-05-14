// This file was generated by counterfeiter
package dockercompose

import (
	"sync"

	"github.com/dev-op-spec/engine/core/logging"
)

type fakeKillOpRunUseCase struct {
	ExecuteStub        func(correlationId string, pathToOpDir string, logger logging.Logger) (err error)
	executeMutex       sync.RWMutex
	executeArgsForCall []struct {
		correlationId string
		pathToOpDir   string
		logger        logging.Logger
	}
	executeReturns struct {
		result1 error
	}
}

func (fake *fakeKillOpRunUseCase) Execute(correlationId string, pathToOpDir string, logger logging.Logger) (err error) {
	fake.executeMutex.Lock()
	fake.executeArgsForCall = append(fake.executeArgsForCall, struct {
		correlationId string
		pathToOpDir   string
		logger        logging.Logger
	}{correlationId, pathToOpDir, logger})
	fake.executeMutex.Unlock()
	if fake.ExecuteStub != nil {
		return fake.ExecuteStub(correlationId, pathToOpDir, logger)
	} else {
		return fake.executeReturns.result1
	}
}

func (fake *fakeKillOpRunUseCase) ExecuteCallCount() int {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return len(fake.executeArgsForCall)
}

func (fake *fakeKillOpRunUseCase) ExecuteArgsForCall(i int) (string, string, logging.Logger) {
	fake.executeMutex.RLock()
	defer fake.executeMutex.RUnlock()
	return fake.executeArgsForCall[i].correlationId, fake.executeArgsForCall[i].pathToOpDir, fake.executeArgsForCall[i].logger
}

func (fake *fakeKillOpRunUseCase) ExecuteReturns(result1 error) {
	fake.ExecuteStub = nil
	fake.executeReturns = struct {
		result1 error
	}{result1}
}
