// This file was generated by counterfeiter
package gometricsfakes

import (
	"sync"

	"github.com/bountylabs/go-metrics"
)

type FakeRegistry struct {
	EachStub        func(func(string, interface{}))
	eachMutex       sync.RWMutex
	eachArgsForCall []struct {
		arg1 func(string, interface{})
	}
	GetStub        func(string) interface{}
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		arg1 string
	}
	getReturns struct {
		result1 interface{}
	}
	GetOrRegisterStub        func(string, interface{}) interface{}
	getOrRegisterMutex       sync.RWMutex
	getOrRegisterArgsForCall []struct {
		arg1 string
		arg2 interface{}
	}
	getOrRegisterReturns struct {
		result1 interface{}
	}
	RegisterStub        func(string, interface{}) error
	registerMutex       sync.RWMutex
	registerArgsForCall []struct {
		arg1 string
		arg2 interface{}
	}
	registerReturns struct {
		result1 error
	}
	RunHealthchecksStub        func()
	runHealthchecksMutex       sync.RWMutex
	runHealthchecksArgsForCall []struct{}
	UnregisterStub             func(string)
	unregisterMutex            sync.RWMutex
	unregisterArgsForCall      []struct {
		arg1 string
	}
	UnregisterAllStub        func()
	unregisterAllMutex       sync.RWMutex
	unregisterAllArgsForCall []struct{}
	invocations              map[string][][]interface{}
	invocationsMutex         sync.RWMutex
}

func (fake *FakeRegistry) Each(arg1 func(string, interface{})) {
	fake.eachMutex.Lock()
	fake.eachArgsForCall = append(fake.eachArgsForCall, struct {
		arg1 func(string, interface{})
	}{arg1})
	fake.recordInvocation("Each", []interface{}{arg1})
	fake.eachMutex.Unlock()
	if fake.EachStub != nil {
		fake.EachStub(arg1)
	}
}

func (fake *FakeRegistry) EachCallCount() int {
	fake.eachMutex.RLock()
	defer fake.eachMutex.RUnlock()
	return len(fake.eachArgsForCall)
}

func (fake *FakeRegistry) EachArgsForCall(i int) func(string, interface{}) {
	fake.eachMutex.RLock()
	defer fake.eachMutex.RUnlock()
	return fake.eachArgsForCall[i].arg1
}

func (fake *FakeRegistry) Get(arg1 string) interface{} {
	fake.getMutex.Lock()
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Get", []interface{}{arg1})
	fake.getMutex.Unlock()
	if fake.GetStub != nil {
		return fake.GetStub(arg1)
	}
	return fake.getReturns.result1
}

func (fake *FakeRegistry) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeRegistry) GetArgsForCall(i int) string {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return fake.getArgsForCall[i].arg1
}

func (fake *FakeRegistry) GetReturns(result1 interface{}) {
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 interface{}
	}{result1}
}

func (fake *FakeRegistry) GetOrRegister(arg1 string, arg2 interface{}) interface{} {
	fake.getOrRegisterMutex.Lock()
	fake.getOrRegisterArgsForCall = append(fake.getOrRegisterArgsForCall, struct {
		arg1 string
		arg2 interface{}
	}{arg1, arg2})
	fake.recordInvocation("GetOrRegister", []interface{}{arg1, arg2})
	fake.getOrRegisterMutex.Unlock()
	if fake.GetOrRegisterStub != nil {
		return fake.GetOrRegisterStub(arg1, arg2)
	}
	return fake.getOrRegisterReturns.result1
}

func (fake *FakeRegistry) GetOrRegisterCallCount() int {
	fake.getOrRegisterMutex.RLock()
	defer fake.getOrRegisterMutex.RUnlock()
	return len(fake.getOrRegisterArgsForCall)
}

func (fake *FakeRegistry) GetOrRegisterArgsForCall(i int) (string, interface{}) {
	fake.getOrRegisterMutex.RLock()
	defer fake.getOrRegisterMutex.RUnlock()
	return fake.getOrRegisterArgsForCall[i].arg1, fake.getOrRegisterArgsForCall[i].arg2
}

func (fake *FakeRegistry) GetOrRegisterReturns(result1 interface{}) {
	fake.GetOrRegisterStub = nil
	fake.getOrRegisterReturns = struct {
		result1 interface{}
	}{result1}
}

func (fake *FakeRegistry) Register(arg1 string, arg2 interface{}) error {
	fake.registerMutex.Lock()
	fake.registerArgsForCall = append(fake.registerArgsForCall, struct {
		arg1 string
		arg2 interface{}
	}{arg1, arg2})
	fake.recordInvocation("Register", []interface{}{arg1, arg2})
	fake.registerMutex.Unlock()
	if fake.RegisterStub != nil {
		return fake.RegisterStub(arg1, arg2)
	}
	return fake.registerReturns.result1
}

func (fake *FakeRegistry) RegisterCallCount() int {
	fake.registerMutex.RLock()
	defer fake.registerMutex.RUnlock()
	return len(fake.registerArgsForCall)
}

func (fake *FakeRegistry) RegisterArgsForCall(i int) (string, interface{}) {
	fake.registerMutex.RLock()
	defer fake.registerMutex.RUnlock()
	return fake.registerArgsForCall[i].arg1, fake.registerArgsForCall[i].arg2
}

func (fake *FakeRegistry) RegisterReturns(result1 error) {
	fake.RegisterStub = nil
	fake.registerReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRegistry) RunHealthchecks() {
	fake.runHealthchecksMutex.Lock()
	fake.runHealthchecksArgsForCall = append(fake.runHealthchecksArgsForCall, struct{}{})
	fake.recordInvocation("RunHealthchecks", []interface{}{})
	fake.runHealthchecksMutex.Unlock()
	if fake.RunHealthchecksStub != nil {
		fake.RunHealthchecksStub()
	}
}

func (fake *FakeRegistry) RunHealthchecksCallCount() int {
	fake.runHealthchecksMutex.RLock()
	defer fake.runHealthchecksMutex.RUnlock()
	return len(fake.runHealthchecksArgsForCall)
}

func (fake *FakeRegistry) Unregister(arg1 string) {
	fake.unregisterMutex.Lock()
	fake.unregisterArgsForCall = append(fake.unregisterArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Unregister", []interface{}{arg1})
	fake.unregisterMutex.Unlock()
	if fake.UnregisterStub != nil {
		fake.UnregisterStub(arg1)
	}
}

func (fake *FakeRegistry) UnregisterCallCount() int {
	fake.unregisterMutex.RLock()
	defer fake.unregisterMutex.RUnlock()
	return len(fake.unregisterArgsForCall)
}

func (fake *FakeRegistry) UnregisterArgsForCall(i int) string {
	fake.unregisterMutex.RLock()
	defer fake.unregisterMutex.RUnlock()
	return fake.unregisterArgsForCall[i].arg1
}

func (fake *FakeRegistry) UnregisterAll() {
	fake.unregisterAllMutex.Lock()
	fake.unregisterAllArgsForCall = append(fake.unregisterAllArgsForCall, struct{}{})
	fake.recordInvocation("UnregisterAll", []interface{}{})
	fake.unregisterAllMutex.Unlock()
	if fake.UnregisterAllStub != nil {
		fake.UnregisterAllStub()
	}
}

func (fake *FakeRegistry) UnregisterAllCallCount() int {
	fake.unregisterAllMutex.RLock()
	defer fake.unregisterAllMutex.RUnlock()
	return len(fake.unregisterAllArgsForCall)
}

func (fake *FakeRegistry) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.eachMutex.RLock()
	defer fake.eachMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.getOrRegisterMutex.RLock()
	defer fake.getOrRegisterMutex.RUnlock()
	fake.registerMutex.RLock()
	defer fake.registerMutex.RUnlock()
	fake.runHealthchecksMutex.RLock()
	defer fake.runHealthchecksMutex.RUnlock()
	fake.unregisterMutex.RLock()
	defer fake.unregisterMutex.RUnlock()
	fake.unregisterAllMutex.RLock()
	defer fake.unregisterAllMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeRegistry) recordInvocation(key string, args []interface{}) {
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

var _ metrics.Registry = new(FakeRegistry)
