// This file was generated by counterfeiter
package core

import (
	"sync"

	"github.com/dev-op-spec/engine/core/models"
)

type fakeOpRunner struct {
	RunStub        func(correlationId string, opUrl *models.Url, ancestorOpRunStartedEvents []models.OpRunStartedEvent) (opRunId string, err error)
	runMutex       sync.RWMutex
	runArgsForCall []struct {
		correlationId              string
		opUrl                      *models.Url
		ancestorOpRunStartedEvents []models.OpRunStartedEvent
	}
	runReturns struct {
		result1 string
		result2 error
	}
	KillStub        func(correlationId string, opRunId string) (err error)
	killMutex       sync.RWMutex
	killArgsForCall []struct {
		correlationId string
		opRunId       string
	}
	killReturns struct {
		result1 error
	}
}

func (fake *fakeOpRunner) Run(correlationId string, opUrl *models.Url, ancestorOpRunStartedEvents []models.OpRunStartedEvent) (opRunId string, err error) {
	ancestorOpRunStartedEventsCopy := make([]models.OpRunStartedEvent, len(ancestorOpRunStartedEvents))
	copy(ancestorOpRunStartedEventsCopy, ancestorOpRunStartedEvents)
	fake.runMutex.Lock()
	fake.runArgsForCall = append(fake.runArgsForCall, struct {
		correlationId              string
		opUrl                      *models.Url
		ancestorOpRunStartedEvents []models.OpRunStartedEvent
	}{correlationId, opUrl, ancestorOpRunStartedEventsCopy})
	fake.runMutex.Unlock()
	if fake.RunStub != nil {
		return fake.RunStub(correlationId, opUrl, ancestorOpRunStartedEvents)
	} else {
		return fake.runReturns.result1, fake.runReturns.result2
	}
}

func (fake *fakeOpRunner) RunCallCount() int {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return len(fake.runArgsForCall)
}

func (fake *fakeOpRunner) RunArgsForCall(i int) (string, *models.Url, []models.OpRunStartedEvent) {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return fake.runArgsForCall[i].correlationId, fake.runArgsForCall[i].opUrl, fake.runArgsForCall[i].ancestorOpRunStartedEvents
}

func (fake *fakeOpRunner) RunReturns(result1 string, result2 error) {
	fake.RunStub = nil
	fake.runReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *fakeOpRunner) Kill(correlationId string, opRunId string) (err error) {
	fake.killMutex.Lock()
	fake.killArgsForCall = append(fake.killArgsForCall, struct {
		correlationId string
		opRunId       string
	}{correlationId, opRunId})
	fake.killMutex.Unlock()
	if fake.KillStub != nil {
		return fake.KillStub(correlationId, opRunId)
	} else {
		return fake.killReturns.result1
	}
}

func (fake *fakeOpRunner) KillCallCount() int {
	fake.killMutex.RLock()
	defer fake.killMutex.RUnlock()
	return len(fake.killArgsForCall)
}

func (fake *fakeOpRunner) KillArgsForCall(i int) (string, string) {
	fake.killMutex.RLock()
	defer fake.killMutex.RUnlock()
	return fake.killArgsForCall[i].correlationId, fake.killArgsForCall[i].opRunId
}

func (fake *fakeOpRunner) KillReturns(result1 error) {
	fake.KillStub = nil
	fake.killReturns = struct {
		result1 error
	}{result1}
}
