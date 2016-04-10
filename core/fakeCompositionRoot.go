// This file was generated by counterfeiter
package core

import (
  "sync"
)

type fakeCompositionRoot struct {
  AddOpUseCaseStub                     func() addOpUseCase
  addOpUseCaseMutex                    sync.RWMutex
  addOpUseCaseArgsForCall              []struct{}
  addOpUseCaseReturns                  struct {
                                         result1 addOpUseCase
                                       }
  AddSubOpUseCaseStub                  func() addSubOpUseCase
  addSubOpUseCaseMutex                 sync.RWMutex
  addSubOpUseCaseArgsForCall           []struct{}
  addSubOpUseCaseReturns               struct {
                                         result1 addSubOpUseCase
                                       }
  GetEventStreamUseCaseStub            func() getEventStreamUseCase
  getEventStreamUseCaseMutex           sync.RWMutex
  getEventStreamUseCaseArgsForCall     []struct{}
  getEventStreamUseCaseReturns         struct {
                                         result1 getEventStreamUseCase
                                       }
  ListOpsUseCaseStub                   func() listOpsUseCase
  listOpsUseCaseMutex                  sync.RWMutex
  listOpsUseCaseArgsForCall            []struct{}
  listOpsUseCaseReturns                struct {
                                         result1 listOpsUseCase
                                       }
  RunOpUseCaseStub                     func() addOpUseCase
  runOpUseCaseMutex                    sync.RWMutex
  runOpUseCaseArgsForCall              []struct{}
  runOpUseCaseReturns                  struct {
                                         result1 addOpUseCase
                                       }
  SetDescriptionOfOpUseCaseStub        func() setDescriptionOfOpUseCase
  setDescriptionOfOpUseCaseMutex       sync.RWMutex
  setDescriptionOfOpUseCaseArgsForCall []struct{}
  setDescriptionOfOpUseCaseReturns     struct {
                                         result1 setDescriptionOfOpUseCase
                                       }
}

func (fake *fakeCompositionRoot) AddOpUseCase() addOpUseCase {
  fake.addOpUseCaseMutex.Lock()
  fake.addOpUseCaseArgsForCall = append(fake.addOpUseCaseArgsForCall, struct{}{})
  fake.addOpUseCaseMutex.Unlock()
  if fake.AddOpUseCaseStub != nil {
    return fake.AddOpUseCaseStub()
  } else {
    return fake.addOpUseCaseReturns.result1
  }
}

func (fake *fakeCompositionRoot) AddOpUseCaseCallCount() int {
  fake.addOpUseCaseMutex.RLock()
  defer fake.addOpUseCaseMutex.RUnlock()
  return len(fake.addOpUseCaseArgsForCall)
}

func (fake *fakeCompositionRoot) AddOpUseCaseReturns(result1 addOpUseCase) {
  fake.AddOpUseCaseStub = nil
  fake.addOpUseCaseReturns = struct {
    result1 addOpUseCase
  }{result1}
}

func (fake *fakeCompositionRoot) AddSubOpUseCase() addSubOpUseCase {
  fake.addSubOpUseCaseMutex.Lock()
  fake.addSubOpUseCaseArgsForCall = append(fake.addSubOpUseCaseArgsForCall, struct{}{})
  fake.addSubOpUseCaseMutex.Unlock()
  if fake.AddSubOpUseCaseStub != nil {
    return fake.AddSubOpUseCaseStub()
  } else {
    return fake.addSubOpUseCaseReturns.result1
  }
}

func (fake *fakeCompositionRoot) AddSubOpUseCaseCallCount() int {
  fake.addSubOpUseCaseMutex.RLock()
  defer fake.addSubOpUseCaseMutex.RUnlock()
  return len(fake.addSubOpUseCaseArgsForCall)
}

func (fake *fakeCompositionRoot) AddSubOpUseCaseReturns(result1 addSubOpUseCase) {
  fake.AddSubOpUseCaseStub = nil
  fake.addSubOpUseCaseReturns = struct {
    result1 addSubOpUseCase
  }{result1}
}

func (fake *fakeCompositionRoot) GetEventStreamUseCase() getEventStreamUseCase {
  fake.getEventStreamUseCaseMutex.Lock()
  fake.getEventStreamUseCaseArgsForCall = append(fake.getEventStreamUseCaseArgsForCall, struct{}{})
  fake.getEventStreamUseCaseMutex.Unlock()
  if fake.GetEventStreamUseCaseStub != nil {
    return fake.GetEventStreamUseCaseStub()
  } else {
    return fake.getEventStreamUseCaseReturns.result1
  }
}

func (fake *fakeCompositionRoot) GetEventStreamUseCaseCallCount() int {
  fake.getEventStreamUseCaseMutex.RLock()
  defer fake.getEventStreamUseCaseMutex.RUnlock()
  return len(fake.getEventStreamUseCaseArgsForCall)
}

func (fake *fakeCompositionRoot) GetEventStreamUseCaseReturns(result1 getEventStreamUseCase) {
  fake.GetEventStreamUseCaseStub = nil
  fake.getEventStreamUseCaseReturns = struct {
    result1 getEventStreamUseCase
  }{result1}
}

func (fake *fakeCompositionRoot) ListOpsUseCase() listOpsUseCase {
  fake.listOpsUseCaseMutex.Lock()
  fake.listOpsUseCaseArgsForCall = append(fake.listOpsUseCaseArgsForCall, struct{}{})
  fake.listOpsUseCaseMutex.Unlock()
  if fake.ListOpsUseCaseStub != nil {
    return fake.ListOpsUseCaseStub()
  } else {
    return fake.listOpsUseCaseReturns.result1
  }
}

func (fake *fakeCompositionRoot) ListOpsUseCaseCallCount() int {
  fake.listOpsUseCaseMutex.RLock()
  defer fake.listOpsUseCaseMutex.RUnlock()
  return len(fake.listOpsUseCaseArgsForCall)
}

func (fake *fakeCompositionRoot) ListOpsUseCaseReturns(result1 listOpsUseCase) {
  fake.ListOpsUseCaseStub = nil
  fake.listOpsUseCaseReturns = struct {
    result1 listOpsUseCase
  }{result1}
}

func (fake *fakeCompositionRoot) RunOpUseCase() addOpUseCase {
  fake.runOpUseCaseMutex.Lock()
  fake.runOpUseCaseArgsForCall = append(fake.runOpUseCaseArgsForCall, struct{}{})
  fake.runOpUseCaseMutex.Unlock()
  if fake.RunOpUseCaseStub != nil {
    return fake.RunOpUseCaseStub()
  } else {
    return fake.runOpUseCaseReturns.result1
  }
}

func (fake *fakeCompositionRoot) RunOpUseCaseCallCount() int {
  fake.runOpUseCaseMutex.RLock()
  defer fake.runOpUseCaseMutex.RUnlock()
  return len(fake.runOpUseCaseArgsForCall)
}

func (fake *fakeCompositionRoot) RunOpUseCaseReturns(result1 addOpUseCase) {
  fake.RunOpUseCaseStub = nil
  fake.runOpUseCaseReturns = struct {
    result1 addOpUseCase
  }{result1}
}

func (fake *fakeCompositionRoot) SetDescriptionOfOpUseCase() setDescriptionOfOpUseCase {
  fake.setDescriptionOfOpUseCaseMutex.Lock()
  fake.setDescriptionOfOpUseCaseArgsForCall = append(fake.setDescriptionOfOpUseCaseArgsForCall, struct{}{})
  fake.setDescriptionOfOpUseCaseMutex.Unlock()
  if fake.SetDescriptionOfOpUseCaseStub != nil {
    return fake.SetDescriptionOfOpUseCaseStub()
  } else {
    return fake.setDescriptionOfOpUseCaseReturns.result1
  }
}

func (fake *fakeCompositionRoot) SetDescriptionOfOpUseCaseCallCount() int {
  fake.setDescriptionOfOpUseCaseMutex.RLock()
  defer fake.setDescriptionOfOpUseCaseMutex.RUnlock()
  return len(fake.setDescriptionOfOpUseCaseArgsForCall)
}

func (fake *fakeCompositionRoot) SetDescriptionOfOpUseCaseReturns(result1 setDescriptionOfOpUseCase) {
  fake.SetDescriptionOfOpUseCaseStub = nil
  fake.setDescriptionOfOpUseCaseReturns = struct {
    result1 setDescriptionOfOpUseCase
  }{result1}
}
