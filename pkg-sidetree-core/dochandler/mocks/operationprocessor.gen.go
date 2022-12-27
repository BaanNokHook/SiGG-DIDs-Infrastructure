// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"github.com/trustbloc/sidetree-core-go/pkg/document"
	"sync"

	"github.com/trustbloc/sidetree-core-go/pkg/api/protocol"
)

type OperationProcessor struct {
	ResolveStub        func(string, ...document.ResolutionOption) (*protocol.ResolutionModel, error)
	resolveMutex       sync.RWMutex
	resolveArgsForCall []struct {
		arg1 string
		arg2 []document.ResolutionOption
	}
	resolveReturns struct {
		result1 *protocol.ResolutionModel
		result2 error
	}
	resolveReturnsOnCall map[int]struct {
		result1 *protocol.ResolutionModel
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *OperationProcessor) Resolve(arg1 string, arg2 ...document.ResolutionOption) (*protocol.ResolutionModel, error) {
	fake.resolveMutex.Lock()
	ret, specificReturn := fake.resolveReturnsOnCall[len(fake.resolveArgsForCall)]
	fake.resolveArgsForCall = append(fake.resolveArgsForCall, struct {
		arg1 string
		arg2 []document.ResolutionOption
	}{arg1, arg2})
	fake.recordInvocation("Resolve", []interface{}{arg1, arg2})
	fake.resolveMutex.Unlock()
	if fake.ResolveStub != nil {
		return fake.ResolveStub(arg1, arg2...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.resolveReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *OperationProcessor) ResolveCallCount() int {
	fake.resolveMutex.RLock()
	defer fake.resolveMutex.RUnlock()
	return len(fake.resolveArgsForCall)
}

func (fake *OperationProcessor) ResolveCalls(stub func(string, ...document.ResolutionOption) (*protocol.ResolutionModel, error)) {
	fake.resolveMutex.Lock()
	defer fake.resolveMutex.Unlock()
	fake.ResolveStub = stub
}

func (fake *OperationProcessor) ResolveArgsForCall(i int) (string, []document.ResolutionOption) {
	fake.resolveMutex.RLock()
	defer fake.resolveMutex.RUnlock()
	argsForCall := fake.resolveArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *OperationProcessor) ResolveReturns(result1 *protocol.ResolutionModel, result2 error) {
	fake.resolveMutex.Lock()
	defer fake.resolveMutex.Unlock()
	fake.ResolveStub = nil
	fake.resolveReturns = struct {
		result1 *protocol.ResolutionModel
		result2 error
	}{result1, result2}
}

func (fake *OperationProcessor) ResolveReturnsOnCall(i int, result1 *protocol.ResolutionModel, result2 error) {
	fake.resolveMutex.Lock()
	defer fake.resolveMutex.Unlock()
	fake.ResolveStub = nil
	if fake.resolveReturnsOnCall == nil {
		fake.resolveReturnsOnCall = make(map[int]struct {
			result1 *protocol.ResolutionModel
			result2 error
		})
	}
	fake.resolveReturnsOnCall[i] = struct {
		result1 *protocol.ResolutionModel
		result2 error
	}{result1, result2}
}

func (fake *OperationProcessor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.resolveMutex.RLock()
	defer fake.resolveMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *OperationProcessor) recordInvocation(key string, args []interface{}) {
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