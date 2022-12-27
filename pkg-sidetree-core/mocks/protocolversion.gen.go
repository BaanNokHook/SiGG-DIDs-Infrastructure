// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"

	"github.com/trustbloc/sidetree-core-go/pkg/api/protocol"
)

type ProtocolVersion struct {
	DocumentComposerStub        func() protocol.DocumentComposer
	documentComposerMutex       sync.RWMutex
	documentComposerArgsForCall []struct {
	}
	documentComposerReturns struct {
		result1 protocol.DocumentComposer
	}
	documentComposerReturnsOnCall map[int]struct {
		result1 protocol.DocumentComposer
	}
	DocumentTransformerStub        func() protocol.DocumentTransformer
	documentTransformerMutex       sync.RWMutex
	documentTransformerArgsForCall []struct {
	}
	documentTransformerReturns struct {
		result1 protocol.DocumentTransformer
	}
	documentTransformerReturnsOnCall map[int]struct {
		result1 protocol.DocumentTransformer
	}
	DocumentValidatorStub        func() protocol.DocumentValidator
	documentValidatorMutex       sync.RWMutex
	documentValidatorArgsForCall []struct {
	}
	documentValidatorReturns struct {
		result1 protocol.DocumentValidator
	}
	documentValidatorReturnsOnCall map[int]struct {
		result1 protocol.DocumentValidator
	}
	OperationApplierStub        func() protocol.OperationApplier
	operationApplierMutex       sync.RWMutex
	operationApplierArgsForCall []struct {
	}
	operationApplierReturns struct {
		result1 protocol.OperationApplier
	}
	operationApplierReturnsOnCall map[int]struct {
		result1 protocol.OperationApplier
	}
	OperationHandlerStub        func() protocol.OperationHandler
	operationHandlerMutex       sync.RWMutex
	operationHandlerArgsForCall []struct {
	}
	operationHandlerReturns struct {
		result1 protocol.OperationHandler
	}
	operationHandlerReturnsOnCall map[int]struct {
		result1 protocol.OperationHandler
	}
	OperationParserStub        func() protocol.OperationParser
	operationParserMutex       sync.RWMutex
	operationParserArgsForCall []struct {
	}
	operationParserReturns struct {
		result1 protocol.OperationParser
	}
	operationParserReturnsOnCall map[int]struct {
		result1 protocol.OperationParser
	}
	OperationProviderStub        func() protocol.OperationProvider
	operationProviderMutex       sync.RWMutex
	operationProviderArgsForCall []struct {
	}
	operationProviderReturns struct {
		result1 protocol.OperationProvider
	}
	operationProviderReturnsOnCall map[int]struct {
		result1 protocol.OperationProvider
	}
	ProtocolStub        func() protocol.Protocol
	protocolMutex       sync.RWMutex
	protocolArgsForCall []struct {
	}
	protocolReturns struct {
		result1 protocol.Protocol
	}
	protocolReturnsOnCall map[int]struct {
		result1 protocol.Protocol
	}
	TransactionProcessorStub        func() protocol.TxnProcessor
	transactionProcessorMutex       sync.RWMutex
	transactionProcessorArgsForCall []struct {
	}
	transactionProcessorReturns struct {
		result1 protocol.TxnProcessor
	}
	transactionProcessorReturnsOnCall map[int]struct {
		result1 protocol.TxnProcessor
	}
	VersionStub        func() string
	versionMutex       sync.RWMutex
	versionArgsForCall []struct {
	}
	versionReturns struct {
		result1 string
	}
	versionReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ProtocolVersion) DocumentComposer() protocol.DocumentComposer {
	fake.documentComposerMutex.Lock()
	ret, specificReturn := fake.documentComposerReturnsOnCall[len(fake.documentComposerArgsForCall)]
	fake.documentComposerArgsForCall = append(fake.documentComposerArgsForCall, struct {
	}{})
	fake.recordInvocation("DocumentComposer", []interface{}{})
	fake.documentComposerMutex.Unlock()
	if fake.DocumentComposerStub != nil {
		return fake.DocumentComposerStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.documentComposerReturns
	return fakeReturns.result1
}

func (fake *ProtocolVersion) DocumentComposerCallCount() int {
	fake.documentComposerMutex.RLock()
	defer fake.documentComposerMutex.RUnlock()
	return len(fake.documentComposerArgsForCall)
}

func (fake *ProtocolVersion) DocumentComposerCalls(stub func() protocol.DocumentComposer) {
	fake.documentComposerMutex.Lock()
	defer fake.documentComposerMutex.Unlock()
	fake.DocumentComposerStub = stub
}

func (fake *ProtocolVersion) DocumentComposerReturns(result1 protocol.DocumentComposer) {
	fake.documentComposerMutex.Lock()
	defer fake.documentComposerMutex.Unlock()
	fake.DocumentComposerStub = nil
	fake.documentComposerReturns = struct {
		result1 protocol.DocumentComposer
	}{result1}
}

func (fake *ProtocolVersion) DocumentComposerReturnsOnCall(i int, result1 protocol.DocumentComposer) {
	fake.documentComposerMutex.Lock()
	defer fake.documentComposerMutex.Unlock()
	fake.DocumentComposerStub = nil
	if fake.documentComposerReturnsOnCall == nil {
		fake.documentComposerReturnsOnCall = make(map[int]struct {
			result1 protocol.DocumentComposer
		})
	}
	fake.documentComposerReturnsOnCall[i] = struct {
		result1 protocol.DocumentComposer
	}{result1}
}

func (fake *ProtocolVersion) DocumentTransformer() protocol.DocumentTransformer {
	fake.documentTransformerMutex.Lock()
	ret, specificReturn := fake.documentTransformerReturnsOnCall[len(fake.documentTransformerArgsForCall)]
	fake.documentTransformerArgsForCall = append(fake.documentTransformerArgsForCall, struct {
	}{})
	fake.recordInvocation("DocumentTransformer", []interface{}{})
	fake.documentTransformerMutex.Unlock()
	if fake.DocumentTransformerStub != nil {
		return fake.DocumentTransformerStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.documentTransformerReturns
	return fakeReturns.result1
}

func (fake *ProtocolVersion) DocumentTransformerCallCount() int {
	fake.documentTransformerMutex.RLock()
	defer fake.documentTransformerMutex.RUnlock()
	return len(fake.documentTransformerArgsForCall)
}

func (fake *ProtocolVersion) DocumentTransformerCalls(stub func() protocol.DocumentTransformer) {
	fake.documentTransformerMutex.Lock()
	defer fake.documentTransformerMutex.Unlock()
	fake.DocumentTransformerStub = stub
}

func (fake *ProtocolVersion) DocumentTransformerReturns(result1 protocol.DocumentTransformer) {
	fake.documentTransformerMutex.Lock()
	defer fake.documentTransformerMutex.Unlock()
	fake.DocumentTransformerStub = nil
	fake.documentTransformerReturns = struct {
		result1 protocol.DocumentTransformer
	}{result1}
}

func (fake *ProtocolVersion) DocumentTransformerReturnsOnCall(i int, result1 protocol.DocumentTransformer) {
	fake.documentTransformerMutex.Lock()
	defer fake.documentTransformerMutex.Unlock()
	fake.DocumentTransformerStub = nil
	if fake.documentTransformerReturnsOnCall == nil {
		fake.documentTransformerReturnsOnCall = make(map[int]struct {
			result1 protocol.DocumentTransformer
		})
	}
	fake.documentTransformerReturnsOnCall[i] = struct {
		result1 protocol.DocumentTransformer
	}{result1}
}

func (fake *ProtocolVersion) DocumentValidator() protocol.DocumentValidator {
	fake.documentValidatorMutex.Lock()
	ret, specificReturn := fake.documentValidatorReturnsOnCall[len(fake.documentValidatorArgsForCall)]
	fake.documentValidatorArgsForCall = append(fake.documentValidatorArgsForCall, struct {
	}{})
	fake.recordInvocation("DocumentValidator", []interface{}{})
	fake.documentValidatorMutex.Unlock()
	if fake.DocumentValidatorStub != nil {
		return fake.DocumentValidatorStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.documentValidatorReturns
	return fakeReturns.result1
}

func (fake *ProtocolVersion) DocumentValidatorCallCount() int {
	fake.documentValidatorMutex.RLock()
	defer fake.documentValidatorMutex.RUnlock()
	return len(fake.documentValidatorArgsForCall)
}

func (fake *ProtocolVersion) DocumentValidatorCalls(stub func() protocol.DocumentValidator) {
	fake.documentValidatorMutex.Lock()
	defer fake.documentValidatorMutex.Unlock()
	fake.DocumentValidatorStub = stub
}

func (fake *ProtocolVersion) DocumentValidatorReturns(result1 protocol.DocumentValidator) {
	fake.documentValidatorMutex.Lock()
	defer fake.documentValidatorMutex.Unlock()
	fake.DocumentValidatorStub = nil
	fake.documentValidatorReturns = struct {
		result1 protocol.DocumentValidator
	}{result1}
}

func (fake *ProtocolVersion) DocumentValidatorReturnsOnCall(i int, result1 protocol.DocumentValidator) {
	fake.documentValidatorMutex.Lock()
	defer fake.documentValidatorMutex.Unlock()
	fake.DocumentValidatorStub = nil
	if fake.documentValidatorReturnsOnCall == nil {
		fake.documentValidatorReturnsOnCall = make(map[int]struct {
			result1 protocol.DocumentValidator
		})
	}
	fake.documentValidatorReturnsOnCall[i] = struct {
		result1 protocol.DocumentValidator
	}{result1}
}

func (fake *ProtocolVersion) OperationApplier() protocol.OperationApplier {
	fake.operationApplierMutex.Lock()
	ret, specificReturn := fake.operationApplierReturnsOnCall[len(fake.operationApplierArgsForCall)]
	fake.operationApplierArgsForCall = append(fake.operationApplierArgsForCall, struct {
	}{})
	fake.recordInvocation("OperationApplier", []interface{}{})
	fake.operationApplierMutex.Unlock()
	if fake.OperationApplierStub != nil {
		return fake.OperationApplierStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.operationApplierReturns
	return fakeReturns.result1
}

func (fake *ProtocolVersion) OperationApplierCallCount() int {
	fake.operationApplierMutex.RLock()
	defer fake.operationApplierMutex.RUnlock()
	return len(fake.operationApplierArgsForCall)
}

func (fake *ProtocolVersion) OperationApplierCalls(stub func() protocol.OperationApplier) {
	fake.operationApplierMutex.Lock()
	defer fake.operationApplierMutex.Unlock()
	fake.OperationApplierStub = stub
}

func (fake *ProtocolVersion) OperationApplierReturns(result1 protocol.OperationApplier) {
	fake.operationApplierMutex.Lock()
	defer fake.operationApplierMutex.Unlock()
	fake.OperationApplierStub = nil
	fake.operationApplierReturns = struct {
		result1 protocol.OperationApplier
	}{result1}
}

func (fake *ProtocolVersion) OperationApplierReturnsOnCall(i int, result1 protocol.OperationApplier) {
	fake.operationApplierMutex.Lock()
	defer fake.operationApplierMutex.Unlock()
	fake.OperationApplierStub = nil
	if fake.operationApplierReturnsOnCall == nil {
		fake.operationApplierReturnsOnCall = make(map[int]struct {
			result1 protocol.OperationApplier
		})
	}
	fake.operationApplierReturnsOnCall[i] = struct {
		result1 protocol.OperationApplier
	}{result1}
}

func (fake *ProtocolVersion) OperationHandler() protocol.OperationHandler {
	fake.operationHandlerMutex.Lock()
	ret, specificReturn := fake.operationHandlerReturnsOnCall[len(fake.operationHandlerArgsForCall)]
	fake.operationHandlerArgsForCall = append(fake.operationHandlerArgsForCall, struct {
	}{})
	fake.recordInvocation("OperationHandler", []interface{}{})
	fake.operationHandlerMutex.Unlock()
	if fake.OperationHandlerStub != nil {
		return fake.OperationHandlerStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.operationHandlerReturns
	return fakeReturns.result1
}

func (fake *ProtocolVersion) OperationHandlerCallCount() int {
	fake.operationHandlerMutex.RLock()
	defer fake.operationHandlerMutex.RUnlock()
	return len(fake.operationHandlerArgsForCall)
}

func (fake *ProtocolVersion) OperationHandlerCalls(stub func() protocol.OperationHandler) {
	fake.operationHandlerMutex.Lock()
	defer fake.operationHandlerMutex.Unlock()
	fake.OperationHandlerStub = stub
}

func (fake *ProtocolVersion) OperationHandlerReturns(result1 protocol.OperationHandler) {
	fake.operationHandlerMutex.Lock()
	defer fake.operationHandlerMutex.Unlock()
	fake.OperationHandlerStub = nil
	fake.operationHandlerReturns = struct {
		result1 protocol.OperationHandler
	}{result1}
}

func (fake *ProtocolVersion) OperationHandlerReturnsOnCall(i int, result1 protocol.OperationHandler) {
	fake.operationHandlerMutex.Lock()
	defer fake.operationHandlerMutex.Unlock()
	fake.OperationHandlerStub = nil
	if fake.operationHandlerReturnsOnCall == nil {
		fake.operationHandlerReturnsOnCall = make(map[int]struct {
			result1 protocol.OperationHandler
		})
	}
	fake.operationHandlerReturnsOnCall[i] = struct {
		result1 protocol.OperationHandler
	}{result1}
}

func (fake *ProtocolVersion) OperationParser() protocol.OperationParser {
	fake.operationParserMutex.Lock()
	ret, specificReturn := fake.operationParserReturnsOnCall[len(fake.operationParserArgsForCall)]
	fake.operationParserArgsForCall = append(fake.operationParserArgsForCall, struct {
	}{})
	fake.recordInvocation("OperationParser", []interface{}{})
	fake.operationParserMutex.Unlock()
	if fake.OperationParserStub != nil {
		return fake.OperationParserStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.operationParserReturns
	return fakeReturns.result1
}

func (fake *ProtocolVersion) OperationParserCallCount() int {
	fake.operationParserMutex.RLock()
	defer fake.operationParserMutex.RUnlock()
	return len(fake.operationParserArgsForCall)
}

func (fake *ProtocolVersion) OperationParserCalls(stub func() protocol.OperationParser) {
	fake.operationParserMutex.Lock()
	defer fake.operationParserMutex.Unlock()
	fake.OperationParserStub = stub
}

func (fake *ProtocolVersion) OperationParserReturns(result1 protocol.OperationParser) {
	fake.operationParserMutex.Lock()
	defer fake.operationParserMutex.Unlock()
	fake.OperationParserStub = nil
	fake.operationParserReturns = struct {
		result1 protocol.OperationParser
	}{result1}
}

func (fake *ProtocolVersion) OperationParserReturnsOnCall(i int, result1 protocol.OperationParser) {
	fake.operationParserMutex.Lock()
	defer fake.operationParserMutex.Unlock()
	fake.OperationParserStub = nil
	if fake.operationParserReturnsOnCall == nil {
		fake.operationParserReturnsOnCall = make(map[int]struct {
			result1 protocol.OperationParser
		})
	}
	fake.operationParserReturnsOnCall[i] = struct {
		result1 protocol.OperationParser
	}{result1}
}

func (fake *ProtocolVersion) OperationProvider() protocol.OperationProvider {
	fake.operationProviderMutex.Lock()
	ret, specificReturn := fake.operationProviderReturnsOnCall[len(fake.operationProviderArgsForCall)]
	fake.operationProviderArgsForCall = append(fake.operationProviderArgsForCall, struct {
	}{})
	fake.recordInvocation("OperationProvider", []interface{}{})
	fake.operationProviderMutex.Unlock()
	if fake.OperationProviderStub != nil {
		return fake.OperationProviderStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.operationProviderReturns
	return fakeReturns.result1
}

func (fake *ProtocolVersion) OperationProviderCallCount() int {
	fake.operationProviderMutex.RLock()
	defer fake.operationProviderMutex.RUnlock()
	return len(fake.operationProviderArgsForCall)
}

func (fake *ProtocolVersion) OperationProviderCalls(stub func() protocol.OperationProvider) {
	fake.operationProviderMutex.Lock()
	defer fake.operationProviderMutex.Unlock()
	fake.OperationProviderStub = stub
}

func (fake *ProtocolVersion) OperationProviderReturns(result1 protocol.OperationProvider) {
	fake.operationProviderMutex.Lock()
	defer fake.operationProviderMutex.Unlock()
	fake.OperationProviderStub = nil
	fake.operationProviderReturns = struct {
		result1 protocol.OperationProvider
	}{result1}
}

func (fake *ProtocolVersion) OperationProviderReturnsOnCall(i int, result1 protocol.OperationProvider) {
	fake.operationProviderMutex.Lock()
	defer fake.operationProviderMutex.Unlock()
	fake.OperationProviderStub = nil
	if fake.operationProviderReturnsOnCall == nil {
		fake.operationProviderReturnsOnCall = make(map[int]struct {
			result1 protocol.OperationProvider
		})
	}
	fake.operationProviderReturnsOnCall[i] = struct {
		result1 protocol.OperationProvider
	}{result1}
}

func (fake *ProtocolVersion) Protocol() protocol.Protocol {
	fake.protocolMutex.Lock()
	ret, specificReturn := fake.protocolReturnsOnCall[len(fake.protocolArgsForCall)]
	fake.protocolArgsForCall = append(fake.protocolArgsForCall, struct {
	}{})
	fake.recordInvocation("Protocol", []interface{}{})
	fake.protocolMutex.Unlock()
	if fake.ProtocolStub != nil {
		return fake.ProtocolStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.protocolReturns
	return fakeReturns.result1
}

func (fake *ProtocolVersion) ProtocolCallCount() int {
	fake.protocolMutex.RLock()
	defer fake.protocolMutex.RUnlock()
	return len(fake.protocolArgsForCall)
}

func (fake *ProtocolVersion) ProtocolCalls(stub func() protocol.Protocol) {
	fake.protocolMutex.Lock()
	defer fake.protocolMutex.Unlock()
	fake.ProtocolStub = stub
}

func (fake *ProtocolVersion) ProtocolReturns(result1 protocol.Protocol) {
	fake.protocolMutex.Lock()
	defer fake.protocolMutex.Unlock()
	fake.ProtocolStub = nil
	fake.protocolReturns = struct {
		result1 protocol.Protocol
	}{result1}
}

func (fake *ProtocolVersion) ProtocolReturnsOnCall(i int, result1 protocol.Protocol) {
	fake.protocolMutex.Lock()
	defer fake.protocolMutex.Unlock()
	fake.ProtocolStub = nil
	if fake.protocolReturnsOnCall == nil {
		fake.protocolReturnsOnCall = make(map[int]struct {
			result1 protocol.Protocol
		})
	}
	fake.protocolReturnsOnCall[i] = struct {
		result1 protocol.Protocol
	}{result1}
}

func (fake *ProtocolVersion) TransactionProcessor() protocol.TxnProcessor {
	fake.transactionProcessorMutex.Lock()
	ret, specificReturn := fake.transactionProcessorReturnsOnCall[len(fake.transactionProcessorArgsForCall)]
	fake.transactionProcessorArgsForCall = append(fake.transactionProcessorArgsForCall, struct {
	}{})
	fake.recordInvocation("TransactionProcessor", []interface{}{})
	fake.transactionProcessorMutex.Unlock()
	if fake.TransactionProcessorStub != nil {
		return fake.TransactionProcessorStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.transactionProcessorReturns
	return fakeReturns.result1
}

func (fake *ProtocolVersion) TransactionProcessorCallCount() int {
	fake.transactionProcessorMutex.RLock()
	defer fake.transactionProcessorMutex.RUnlock()
	return len(fake.transactionProcessorArgsForCall)
}

func (fake *ProtocolVersion) TransactionProcessorCalls(stub func() protocol.TxnProcessor) {
	fake.transactionProcessorMutex.Lock()
	defer fake.transactionProcessorMutex.Unlock()
	fake.TransactionProcessorStub = stub
}

func (fake *ProtocolVersion) TransactionProcessorReturns(result1 protocol.TxnProcessor) {
	fake.transactionProcessorMutex.Lock()
	defer fake.transactionProcessorMutex.Unlock()
	fake.TransactionProcessorStub = nil
	fake.transactionProcessorReturns = struct {
		result1 protocol.TxnProcessor
	}{result1}
}

func (fake *ProtocolVersion) TransactionProcessorReturnsOnCall(i int, result1 protocol.TxnProcessor) {
	fake.transactionProcessorMutex.Lock()
	defer fake.transactionProcessorMutex.Unlock()
	fake.TransactionProcessorStub = nil
	if fake.transactionProcessorReturnsOnCall == nil {
		fake.transactionProcessorReturnsOnCall = make(map[int]struct {
			result1 protocol.TxnProcessor
		})
	}
	fake.transactionProcessorReturnsOnCall[i] = struct {
		result1 protocol.TxnProcessor
	}{result1}
}

func (fake *ProtocolVersion) Version() string {
	fake.versionMutex.Lock()
	ret, specificReturn := fake.versionReturnsOnCall[len(fake.versionArgsForCall)]
	fake.versionArgsForCall = append(fake.versionArgsForCall, struct {
	}{})
	fake.recordInvocation("Version", []interface{}{})
	fake.versionMutex.Unlock()
	if fake.VersionStub != nil {
		return fake.VersionStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.versionReturns
	return fakeReturns.result1
}

func (fake *ProtocolVersion) VersionCallCount() int {
	fake.versionMutex.RLock()
	defer fake.versionMutex.RUnlock()
	return len(fake.versionArgsForCall)
}

func (fake *ProtocolVersion) VersionCalls(stub func() string) {
	fake.versionMutex.Lock()
	defer fake.versionMutex.Unlock()
	fake.VersionStub = stub
}

func (fake *ProtocolVersion) VersionReturns(result1 string) {
	fake.versionMutex.Lock()
	defer fake.versionMutex.Unlock()
	fake.VersionStub = nil
	fake.versionReturns = struct {
		result1 string
	}{result1}
}

func (fake *ProtocolVersion) VersionReturnsOnCall(i int, result1 string) {
	fake.versionMutex.Lock()
	defer fake.versionMutex.Unlock()
	fake.VersionStub = nil
	if fake.versionReturnsOnCall == nil {
		fake.versionReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.versionReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *ProtocolVersion) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.documentComposerMutex.RLock()
	defer fake.documentComposerMutex.RUnlock()
	fake.documentTransformerMutex.RLock()
	defer fake.documentTransformerMutex.RUnlock()
	fake.documentValidatorMutex.RLock()
	defer fake.documentValidatorMutex.RUnlock()
	fake.operationApplierMutex.RLock()
	defer fake.operationApplierMutex.RUnlock()
	fake.operationHandlerMutex.RLock()
	defer fake.operationHandlerMutex.RUnlock()
	fake.operationParserMutex.RLock()
	defer fake.operationParserMutex.RUnlock()
	fake.operationProviderMutex.RLock()
	defer fake.operationProviderMutex.RUnlock()
	fake.protocolMutex.RLock()
	defer fake.protocolMutex.RUnlock()
	fake.transactionProcessorMutex.RLock()
	defer fake.transactionProcessorMutex.RUnlock()
	fake.versionMutex.RLock()
	defer fake.versionMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ProtocolVersion) recordInvocation(key string, args []interface{}) {
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

var _ protocol.Version = new(ProtocolVersion)
