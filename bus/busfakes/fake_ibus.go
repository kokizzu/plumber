// Code generated by counterfeiter. DO NOT EDIT.
package busfakes

import (
	"context"
	"sync"

	"github.com/batchcorp/plumber-schemas/build/go/protos/opts"
	"github.com/batchcorp/plumber/bus"
)

type FakeIBus struct {
	PublishCreateConnectionStub        func(context.Context, *opts.ConnectionOptions) error
	publishCreateConnectionMutex       sync.RWMutex
	publishCreateConnectionArgsForCall []struct {
		arg1 context.Context
		arg2 *opts.ConnectionOptions
	}
	publishCreateConnectionReturns struct {
		result1 error
	}
	publishCreateConnectionReturnsOnCall map[int]struct {
		result1 error
	}
	PublishCreateDynamicStub        func(context.Context, *opts.DynamicOptions) error
	publishCreateDynamicMutex       sync.RWMutex
	publishCreateDynamicArgsForCall []struct {
		arg1 context.Context
		arg2 *opts.DynamicOptions
	}
	publishCreateDynamicReturns struct {
		result1 error
	}
	publishCreateDynamicReturnsOnCall map[int]struct {
		result1 error
	}
	PublishCreateRelayStub        func(context.Context, *opts.RelayOptions) error
	publishCreateRelayMutex       sync.RWMutex
	publishCreateRelayArgsForCall []struct {
		arg1 context.Context
		arg2 *opts.RelayOptions
	}
	publishCreateRelayReturns struct {
		result1 error
	}
	publishCreateRelayReturnsOnCall map[int]struct {
		result1 error
	}
	PublishDeleteConnectionStub        func(context.Context, *opts.ConnectionOptions) error
	publishDeleteConnectionMutex       sync.RWMutex
	publishDeleteConnectionArgsForCall []struct {
		arg1 context.Context
		arg2 *opts.ConnectionOptions
	}
	publishDeleteConnectionReturns struct {
		result1 error
	}
	publishDeleteConnectionReturnsOnCall map[int]struct {
		result1 error
	}
	PublishDeleteDynamicStub        func(context.Context, *opts.DynamicOptions) error
	publishDeleteDynamicMutex       sync.RWMutex
	publishDeleteDynamicArgsForCall []struct {
		arg1 context.Context
		arg2 *opts.DynamicOptions
	}
	publishDeleteDynamicReturns struct {
		result1 error
	}
	publishDeleteDynamicReturnsOnCall map[int]struct {
		result1 error
	}
	PublishDeleteRelayStub        func(context.Context, *opts.RelayOptions) error
	publishDeleteRelayMutex       sync.RWMutex
	publishDeleteRelayArgsForCall []struct {
		arg1 context.Context
		arg2 *opts.RelayOptions
	}
	publishDeleteRelayReturns struct {
		result1 error
	}
	publishDeleteRelayReturnsOnCall map[int]struct {
		result1 error
	}
	PublishResumeDynamicStub        func(context.Context, *opts.DynamicOptions) error
	publishResumeDynamicMutex       sync.RWMutex
	publishResumeDynamicArgsForCall []struct {
		arg1 context.Context
		arg2 *opts.DynamicOptions
	}
	publishResumeDynamicReturns struct {
		result1 error
	}
	publishResumeDynamicReturnsOnCall map[int]struct {
		result1 error
	}
	PublishResumeRelayStub        func(context.Context, *opts.RelayOptions) error
	publishResumeRelayMutex       sync.RWMutex
	publishResumeRelayArgsForCall []struct {
		arg1 context.Context
		arg2 *opts.RelayOptions
	}
	publishResumeRelayReturns struct {
		result1 error
	}
	publishResumeRelayReturnsOnCall map[int]struct {
		result1 error
	}
	PublishStopDynamicStub        func(context.Context, *opts.DynamicOptions) error
	publishStopDynamicMutex       sync.RWMutex
	publishStopDynamicArgsForCall []struct {
		arg1 context.Context
		arg2 *opts.DynamicOptions
	}
	publishStopDynamicReturns struct {
		result1 error
	}
	publishStopDynamicReturnsOnCall map[int]struct {
		result1 error
	}
	PublishStopRelayStub        func(context.Context, *opts.RelayOptions) error
	publishStopRelayMutex       sync.RWMutex
	publishStopRelayArgsForCall []struct {
		arg1 context.Context
		arg2 *opts.RelayOptions
	}
	publishStopRelayReturns struct {
		result1 error
	}
	publishStopRelayReturnsOnCall map[int]struct {
		result1 error
	}
	PublishUpdateConnectionStub        func(context.Context, *opts.ConnectionOptions) error
	publishUpdateConnectionMutex       sync.RWMutex
	publishUpdateConnectionArgsForCall []struct {
		arg1 context.Context
		arg2 *opts.ConnectionOptions
	}
	publishUpdateConnectionReturns struct {
		result1 error
	}
	publishUpdateConnectionReturnsOnCall map[int]struct {
		result1 error
	}
	PublishUpdateDynamicStub        func(context.Context, *opts.DynamicOptions) error
	publishUpdateDynamicMutex       sync.RWMutex
	publishUpdateDynamicArgsForCall []struct {
		arg1 context.Context
		arg2 *opts.DynamicOptions
	}
	publishUpdateDynamicReturns struct {
		result1 error
	}
	publishUpdateDynamicReturnsOnCall map[int]struct {
		result1 error
	}
	PublishUpdateRelayStub        func(context.Context, *opts.RelayOptions) error
	publishUpdateRelayMutex       sync.RWMutex
	publishUpdateRelayArgsForCall []struct {
		arg1 context.Context
		arg2 *opts.RelayOptions
	}
	publishUpdateRelayReturns struct {
		result1 error
	}
	publishUpdateRelayReturnsOnCall map[int]struct {
		result1 error
	}
	StartStub        func(context.Context) error
	startMutex       sync.RWMutex
	startArgsForCall []struct {
		arg1 context.Context
	}
	startReturns struct {
		result1 error
	}
	startReturnsOnCall map[int]struct {
		result1 error
	}
	StopStub        func() error
	stopMutex       sync.RWMutex
	stopArgsForCall []struct {
	}
	stopReturns struct {
		result1 error
	}
	stopReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeIBus) PublishCreateConnection(arg1 context.Context, arg2 *opts.ConnectionOptions) error {
	fake.publishCreateConnectionMutex.Lock()
	ret, specificReturn := fake.publishCreateConnectionReturnsOnCall[len(fake.publishCreateConnectionArgsForCall)]
	fake.publishCreateConnectionArgsForCall = append(fake.publishCreateConnectionArgsForCall, struct {
		arg1 context.Context
		arg2 *opts.ConnectionOptions
	}{arg1, arg2})
	stub := fake.PublishCreateConnectionStub
	fakeReturns := fake.publishCreateConnectionReturns
	fake.recordInvocation("PublishCreateConnection", []interface{}{arg1, arg2})
	fake.publishCreateConnectionMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIBus) PublishCreateConnectionCallCount() int {
	fake.publishCreateConnectionMutex.RLock()
	defer fake.publishCreateConnectionMutex.RUnlock()
	return len(fake.publishCreateConnectionArgsForCall)
}

func (fake *FakeIBus) PublishCreateConnectionCalls(stub func(context.Context, *opts.ConnectionOptions) error) {
	fake.publishCreateConnectionMutex.Lock()
	defer fake.publishCreateConnectionMutex.Unlock()
	fake.PublishCreateConnectionStub = stub
}

func (fake *FakeIBus) PublishCreateConnectionArgsForCall(i int) (context.Context, *opts.ConnectionOptions) {
	fake.publishCreateConnectionMutex.RLock()
	defer fake.publishCreateConnectionMutex.RUnlock()
	argsForCall := fake.publishCreateConnectionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeIBus) PublishCreateConnectionReturns(result1 error) {
	fake.publishCreateConnectionMutex.Lock()
	defer fake.publishCreateConnectionMutex.Unlock()
	fake.PublishCreateConnectionStub = nil
	fake.publishCreateConnectionReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBus) PublishCreateConnectionReturnsOnCall(i int, result1 error) {
	fake.publishCreateConnectionMutex.Lock()
	defer fake.publishCreateConnectionMutex.Unlock()
	fake.PublishCreateConnectionStub = nil
	if fake.publishCreateConnectionReturnsOnCall == nil {
		fake.publishCreateConnectionReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.publishCreateConnectionReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBus) PublishCreateDynamic(arg1 context.Context, arg2 *opts.DynamicOptions) error {
	fake.publishCreateDynamicMutex.Lock()
	ret, specificReturn := fake.publishCreateDynamicReturnsOnCall[len(fake.publishCreateDynamicArgsForCall)]
	fake.publishCreateDynamicArgsForCall = append(fake.publishCreateDynamicArgsForCall, struct {
		arg1 context.Context
		arg2 *opts.DynamicOptions
	}{arg1, arg2})
	stub := fake.PublishCreateDynamicStub
	fakeReturns := fake.publishCreateDynamicReturns
	fake.recordInvocation("PublishCreateDynamic", []interface{}{arg1, arg2})
	fake.publishCreateDynamicMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIBus) PublishCreateDynamicCallCount() int {
	fake.publishCreateDynamicMutex.RLock()
	defer fake.publishCreateDynamicMutex.RUnlock()
	return len(fake.publishCreateDynamicArgsForCall)
}

func (fake *FakeIBus) PublishCreateDynamicCalls(stub func(context.Context, *opts.DynamicOptions) error) {
	fake.publishCreateDynamicMutex.Lock()
	defer fake.publishCreateDynamicMutex.Unlock()
	fake.PublishCreateDynamicStub = stub
}

func (fake *FakeIBus) PublishCreateDynamicArgsForCall(i int) (context.Context, *opts.DynamicOptions) {
	fake.publishCreateDynamicMutex.RLock()
	defer fake.publishCreateDynamicMutex.RUnlock()
	argsForCall := fake.publishCreateDynamicArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeIBus) PublishCreateDynamicReturns(result1 error) {
	fake.publishCreateDynamicMutex.Lock()
	defer fake.publishCreateDynamicMutex.Unlock()
	fake.PublishCreateDynamicStub = nil
	fake.publishCreateDynamicReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBus) PublishCreateDynamicReturnsOnCall(i int, result1 error) {
	fake.publishCreateDynamicMutex.Lock()
	defer fake.publishCreateDynamicMutex.Unlock()
	fake.PublishCreateDynamicStub = nil
	if fake.publishCreateDynamicReturnsOnCall == nil {
		fake.publishCreateDynamicReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.publishCreateDynamicReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBus) PublishCreateRelay(arg1 context.Context, arg2 *opts.RelayOptions) error {
	fake.publishCreateRelayMutex.Lock()
	ret, specificReturn := fake.publishCreateRelayReturnsOnCall[len(fake.publishCreateRelayArgsForCall)]
	fake.publishCreateRelayArgsForCall = append(fake.publishCreateRelayArgsForCall, struct {
		arg1 context.Context
		arg2 *opts.RelayOptions
	}{arg1, arg2})
	stub := fake.PublishCreateRelayStub
	fakeReturns := fake.publishCreateRelayReturns
	fake.recordInvocation("PublishCreateRelay", []interface{}{arg1, arg2})
	fake.publishCreateRelayMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIBus) PublishCreateRelayCallCount() int {
	fake.publishCreateRelayMutex.RLock()
	defer fake.publishCreateRelayMutex.RUnlock()
	return len(fake.publishCreateRelayArgsForCall)
}

func (fake *FakeIBus) PublishCreateRelayCalls(stub func(context.Context, *opts.RelayOptions) error) {
	fake.publishCreateRelayMutex.Lock()
	defer fake.publishCreateRelayMutex.Unlock()
	fake.PublishCreateRelayStub = stub
}

func (fake *FakeIBus) PublishCreateRelayArgsForCall(i int) (context.Context, *opts.RelayOptions) {
	fake.publishCreateRelayMutex.RLock()
	defer fake.publishCreateRelayMutex.RUnlock()
	argsForCall := fake.publishCreateRelayArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeIBus) PublishCreateRelayReturns(result1 error) {
	fake.publishCreateRelayMutex.Lock()
	defer fake.publishCreateRelayMutex.Unlock()
	fake.PublishCreateRelayStub = nil
	fake.publishCreateRelayReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBus) PublishCreateRelayReturnsOnCall(i int, result1 error) {
	fake.publishCreateRelayMutex.Lock()
	defer fake.publishCreateRelayMutex.Unlock()
	fake.PublishCreateRelayStub = nil
	if fake.publishCreateRelayReturnsOnCall == nil {
		fake.publishCreateRelayReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.publishCreateRelayReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBus) PublishDeleteConnection(arg1 context.Context, arg2 *opts.ConnectionOptions) error {
	fake.publishDeleteConnectionMutex.Lock()
	ret, specificReturn := fake.publishDeleteConnectionReturnsOnCall[len(fake.publishDeleteConnectionArgsForCall)]
	fake.publishDeleteConnectionArgsForCall = append(fake.publishDeleteConnectionArgsForCall, struct {
		arg1 context.Context
		arg2 *opts.ConnectionOptions
	}{arg1, arg2})
	stub := fake.PublishDeleteConnectionStub
	fakeReturns := fake.publishDeleteConnectionReturns
	fake.recordInvocation("PublishDeleteConnection", []interface{}{arg1, arg2})
	fake.publishDeleteConnectionMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIBus) PublishDeleteConnectionCallCount() int {
	fake.publishDeleteConnectionMutex.RLock()
	defer fake.publishDeleteConnectionMutex.RUnlock()
	return len(fake.publishDeleteConnectionArgsForCall)
}

func (fake *FakeIBus) PublishDeleteConnectionCalls(stub func(context.Context, *opts.ConnectionOptions) error) {
	fake.publishDeleteConnectionMutex.Lock()
	defer fake.publishDeleteConnectionMutex.Unlock()
	fake.PublishDeleteConnectionStub = stub
}

func (fake *FakeIBus) PublishDeleteConnectionArgsForCall(i int) (context.Context, *opts.ConnectionOptions) {
	fake.publishDeleteConnectionMutex.RLock()
	defer fake.publishDeleteConnectionMutex.RUnlock()
	argsForCall := fake.publishDeleteConnectionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeIBus) PublishDeleteConnectionReturns(result1 error) {
	fake.publishDeleteConnectionMutex.Lock()
	defer fake.publishDeleteConnectionMutex.Unlock()
	fake.PublishDeleteConnectionStub = nil
	fake.publishDeleteConnectionReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBus) PublishDeleteConnectionReturnsOnCall(i int, result1 error) {
	fake.publishDeleteConnectionMutex.Lock()
	defer fake.publishDeleteConnectionMutex.Unlock()
	fake.PublishDeleteConnectionStub = nil
	if fake.publishDeleteConnectionReturnsOnCall == nil {
		fake.publishDeleteConnectionReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.publishDeleteConnectionReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBus) PublishDeleteDynamic(arg1 context.Context, arg2 *opts.DynamicOptions) error {
	fake.publishDeleteDynamicMutex.Lock()
	ret, specificReturn := fake.publishDeleteDynamicReturnsOnCall[len(fake.publishDeleteDynamicArgsForCall)]
	fake.publishDeleteDynamicArgsForCall = append(fake.publishDeleteDynamicArgsForCall, struct {
		arg1 context.Context
		arg2 *opts.DynamicOptions
	}{arg1, arg2})
	stub := fake.PublishDeleteDynamicStub
	fakeReturns := fake.publishDeleteDynamicReturns
	fake.recordInvocation("PublishDeleteDynamic", []interface{}{arg1, arg2})
	fake.publishDeleteDynamicMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIBus) PublishDeleteDynamicCallCount() int {
	fake.publishDeleteDynamicMutex.RLock()
	defer fake.publishDeleteDynamicMutex.RUnlock()
	return len(fake.publishDeleteDynamicArgsForCall)
}

func (fake *FakeIBus) PublishDeleteDynamicCalls(stub func(context.Context, *opts.DynamicOptions) error) {
	fake.publishDeleteDynamicMutex.Lock()
	defer fake.publishDeleteDynamicMutex.Unlock()
	fake.PublishDeleteDynamicStub = stub
}

func (fake *FakeIBus) PublishDeleteDynamicArgsForCall(i int) (context.Context, *opts.DynamicOptions) {
	fake.publishDeleteDynamicMutex.RLock()
	defer fake.publishDeleteDynamicMutex.RUnlock()
	argsForCall := fake.publishDeleteDynamicArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeIBus) PublishDeleteDynamicReturns(result1 error) {
	fake.publishDeleteDynamicMutex.Lock()
	defer fake.publishDeleteDynamicMutex.Unlock()
	fake.PublishDeleteDynamicStub = nil
	fake.publishDeleteDynamicReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBus) PublishDeleteDynamicReturnsOnCall(i int, result1 error) {
	fake.publishDeleteDynamicMutex.Lock()
	defer fake.publishDeleteDynamicMutex.Unlock()
	fake.PublishDeleteDynamicStub = nil
	if fake.publishDeleteDynamicReturnsOnCall == nil {
		fake.publishDeleteDynamicReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.publishDeleteDynamicReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBus) PublishDeleteRelay(arg1 context.Context, arg2 *opts.RelayOptions) error {
	fake.publishDeleteRelayMutex.Lock()
	ret, specificReturn := fake.publishDeleteRelayReturnsOnCall[len(fake.publishDeleteRelayArgsForCall)]
	fake.publishDeleteRelayArgsForCall = append(fake.publishDeleteRelayArgsForCall, struct {
		arg1 context.Context
		arg2 *opts.RelayOptions
	}{arg1, arg2})
	stub := fake.PublishDeleteRelayStub
	fakeReturns := fake.publishDeleteRelayReturns
	fake.recordInvocation("PublishDeleteRelay", []interface{}{arg1, arg2})
	fake.publishDeleteRelayMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIBus) PublishDeleteRelayCallCount() int {
	fake.publishDeleteRelayMutex.RLock()
	defer fake.publishDeleteRelayMutex.RUnlock()
	return len(fake.publishDeleteRelayArgsForCall)
}

func (fake *FakeIBus) PublishDeleteRelayCalls(stub func(context.Context, *opts.RelayOptions) error) {
	fake.publishDeleteRelayMutex.Lock()
	defer fake.publishDeleteRelayMutex.Unlock()
	fake.PublishDeleteRelayStub = stub
}

func (fake *FakeIBus) PublishDeleteRelayArgsForCall(i int) (context.Context, *opts.RelayOptions) {
	fake.publishDeleteRelayMutex.RLock()
	defer fake.publishDeleteRelayMutex.RUnlock()
	argsForCall := fake.publishDeleteRelayArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeIBus) PublishDeleteRelayReturns(result1 error) {
	fake.publishDeleteRelayMutex.Lock()
	defer fake.publishDeleteRelayMutex.Unlock()
	fake.PublishDeleteRelayStub = nil
	fake.publishDeleteRelayReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBus) PublishDeleteRelayReturnsOnCall(i int, result1 error) {
	fake.publishDeleteRelayMutex.Lock()
	defer fake.publishDeleteRelayMutex.Unlock()
	fake.PublishDeleteRelayStub = nil
	if fake.publishDeleteRelayReturnsOnCall == nil {
		fake.publishDeleteRelayReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.publishDeleteRelayReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBus) PublishResumeDynamic(arg1 context.Context, arg2 *opts.DynamicOptions) error {
	fake.publishResumeDynamicMutex.Lock()
	ret, specificReturn := fake.publishResumeDynamicReturnsOnCall[len(fake.publishResumeDynamicArgsForCall)]
	fake.publishResumeDynamicArgsForCall = append(fake.publishResumeDynamicArgsForCall, struct {
		arg1 context.Context
		arg2 *opts.DynamicOptions
	}{arg1, arg2})
	stub := fake.PublishResumeDynamicStub
	fakeReturns := fake.publishResumeDynamicReturns
	fake.recordInvocation("PublishResumeDynamic", []interface{}{arg1, arg2})
	fake.publishResumeDynamicMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIBus) PublishResumeDynamicCallCount() int {
	fake.publishResumeDynamicMutex.RLock()
	defer fake.publishResumeDynamicMutex.RUnlock()
	return len(fake.publishResumeDynamicArgsForCall)
}

func (fake *FakeIBus) PublishResumeDynamicCalls(stub func(context.Context, *opts.DynamicOptions) error) {
	fake.publishResumeDynamicMutex.Lock()
	defer fake.publishResumeDynamicMutex.Unlock()
	fake.PublishResumeDynamicStub = stub
}

func (fake *FakeIBus) PublishResumeDynamicArgsForCall(i int) (context.Context, *opts.DynamicOptions) {
	fake.publishResumeDynamicMutex.RLock()
	defer fake.publishResumeDynamicMutex.RUnlock()
	argsForCall := fake.publishResumeDynamicArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeIBus) PublishResumeDynamicReturns(result1 error) {
	fake.publishResumeDynamicMutex.Lock()
	defer fake.publishResumeDynamicMutex.Unlock()
	fake.PublishResumeDynamicStub = nil
	fake.publishResumeDynamicReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBus) PublishResumeDynamicReturnsOnCall(i int, result1 error) {
	fake.publishResumeDynamicMutex.Lock()
	defer fake.publishResumeDynamicMutex.Unlock()
	fake.PublishResumeDynamicStub = nil
	if fake.publishResumeDynamicReturnsOnCall == nil {
		fake.publishResumeDynamicReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.publishResumeDynamicReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBus) PublishResumeRelay(arg1 context.Context, arg2 *opts.RelayOptions) error {
	fake.publishResumeRelayMutex.Lock()
	ret, specificReturn := fake.publishResumeRelayReturnsOnCall[len(fake.publishResumeRelayArgsForCall)]
	fake.publishResumeRelayArgsForCall = append(fake.publishResumeRelayArgsForCall, struct {
		arg1 context.Context
		arg2 *opts.RelayOptions
	}{arg1, arg2})
	stub := fake.PublishResumeRelayStub
	fakeReturns := fake.publishResumeRelayReturns
	fake.recordInvocation("PublishResumeRelay", []interface{}{arg1, arg2})
	fake.publishResumeRelayMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIBus) PublishResumeRelayCallCount() int {
	fake.publishResumeRelayMutex.RLock()
	defer fake.publishResumeRelayMutex.RUnlock()
	return len(fake.publishResumeRelayArgsForCall)
}

func (fake *FakeIBus) PublishResumeRelayCalls(stub func(context.Context, *opts.RelayOptions) error) {
	fake.publishResumeRelayMutex.Lock()
	defer fake.publishResumeRelayMutex.Unlock()
	fake.PublishResumeRelayStub = stub
}

func (fake *FakeIBus) PublishResumeRelayArgsForCall(i int) (context.Context, *opts.RelayOptions) {
	fake.publishResumeRelayMutex.RLock()
	defer fake.publishResumeRelayMutex.RUnlock()
	argsForCall := fake.publishResumeRelayArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeIBus) PublishResumeRelayReturns(result1 error) {
	fake.publishResumeRelayMutex.Lock()
	defer fake.publishResumeRelayMutex.Unlock()
	fake.PublishResumeRelayStub = nil
	fake.publishResumeRelayReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBus) PublishResumeRelayReturnsOnCall(i int, result1 error) {
	fake.publishResumeRelayMutex.Lock()
	defer fake.publishResumeRelayMutex.Unlock()
	fake.PublishResumeRelayStub = nil
	if fake.publishResumeRelayReturnsOnCall == nil {
		fake.publishResumeRelayReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.publishResumeRelayReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBus) PublishStopDynamic(arg1 context.Context, arg2 *opts.DynamicOptions) error {
	fake.publishStopDynamicMutex.Lock()
	ret, specificReturn := fake.publishStopDynamicReturnsOnCall[len(fake.publishStopDynamicArgsForCall)]
	fake.publishStopDynamicArgsForCall = append(fake.publishStopDynamicArgsForCall, struct {
		arg1 context.Context
		arg2 *opts.DynamicOptions
	}{arg1, arg2})
	stub := fake.PublishStopDynamicStub
	fakeReturns := fake.publishStopDynamicReturns
	fake.recordInvocation("PublishStopDynamic", []interface{}{arg1, arg2})
	fake.publishStopDynamicMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIBus) PublishStopDynamicCallCount() int {
	fake.publishStopDynamicMutex.RLock()
	defer fake.publishStopDynamicMutex.RUnlock()
	return len(fake.publishStopDynamicArgsForCall)
}

func (fake *FakeIBus) PublishStopDynamicCalls(stub func(context.Context, *opts.DynamicOptions) error) {
	fake.publishStopDynamicMutex.Lock()
	defer fake.publishStopDynamicMutex.Unlock()
	fake.PublishStopDynamicStub = stub
}

func (fake *FakeIBus) PublishStopDynamicArgsForCall(i int) (context.Context, *opts.DynamicOptions) {
	fake.publishStopDynamicMutex.RLock()
	defer fake.publishStopDynamicMutex.RUnlock()
	argsForCall := fake.publishStopDynamicArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeIBus) PublishStopDynamicReturns(result1 error) {
	fake.publishStopDynamicMutex.Lock()
	defer fake.publishStopDynamicMutex.Unlock()
	fake.PublishStopDynamicStub = nil
	fake.publishStopDynamicReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBus) PublishStopDynamicReturnsOnCall(i int, result1 error) {
	fake.publishStopDynamicMutex.Lock()
	defer fake.publishStopDynamicMutex.Unlock()
	fake.PublishStopDynamicStub = nil
	if fake.publishStopDynamicReturnsOnCall == nil {
		fake.publishStopDynamicReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.publishStopDynamicReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBus) PublishStopRelay(arg1 context.Context, arg2 *opts.RelayOptions) error {
	fake.publishStopRelayMutex.Lock()
	ret, specificReturn := fake.publishStopRelayReturnsOnCall[len(fake.publishStopRelayArgsForCall)]
	fake.publishStopRelayArgsForCall = append(fake.publishStopRelayArgsForCall, struct {
		arg1 context.Context
		arg2 *opts.RelayOptions
	}{arg1, arg2})
	stub := fake.PublishStopRelayStub
	fakeReturns := fake.publishStopRelayReturns
	fake.recordInvocation("PublishStopRelay", []interface{}{arg1, arg2})
	fake.publishStopRelayMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIBus) PublishStopRelayCallCount() int {
	fake.publishStopRelayMutex.RLock()
	defer fake.publishStopRelayMutex.RUnlock()
	return len(fake.publishStopRelayArgsForCall)
}

func (fake *FakeIBus) PublishStopRelayCalls(stub func(context.Context, *opts.RelayOptions) error) {
	fake.publishStopRelayMutex.Lock()
	defer fake.publishStopRelayMutex.Unlock()
	fake.PublishStopRelayStub = stub
}

func (fake *FakeIBus) PublishStopRelayArgsForCall(i int) (context.Context, *opts.RelayOptions) {
	fake.publishStopRelayMutex.RLock()
	defer fake.publishStopRelayMutex.RUnlock()
	argsForCall := fake.publishStopRelayArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeIBus) PublishStopRelayReturns(result1 error) {
	fake.publishStopRelayMutex.Lock()
	defer fake.publishStopRelayMutex.Unlock()
	fake.PublishStopRelayStub = nil
	fake.publishStopRelayReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBus) PublishStopRelayReturnsOnCall(i int, result1 error) {
	fake.publishStopRelayMutex.Lock()
	defer fake.publishStopRelayMutex.Unlock()
	fake.PublishStopRelayStub = nil
	if fake.publishStopRelayReturnsOnCall == nil {
		fake.publishStopRelayReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.publishStopRelayReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBus) PublishUpdateConnection(arg1 context.Context, arg2 *opts.ConnectionOptions) error {
	fake.publishUpdateConnectionMutex.Lock()
	ret, specificReturn := fake.publishUpdateConnectionReturnsOnCall[len(fake.publishUpdateConnectionArgsForCall)]
	fake.publishUpdateConnectionArgsForCall = append(fake.publishUpdateConnectionArgsForCall, struct {
		arg1 context.Context
		arg2 *opts.ConnectionOptions
	}{arg1, arg2})
	stub := fake.PublishUpdateConnectionStub
	fakeReturns := fake.publishUpdateConnectionReturns
	fake.recordInvocation("PublishUpdateConnection", []interface{}{arg1, arg2})
	fake.publishUpdateConnectionMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIBus) PublishUpdateConnectionCallCount() int {
	fake.publishUpdateConnectionMutex.RLock()
	defer fake.publishUpdateConnectionMutex.RUnlock()
	return len(fake.publishUpdateConnectionArgsForCall)
}

func (fake *FakeIBus) PublishUpdateConnectionCalls(stub func(context.Context, *opts.ConnectionOptions) error) {
	fake.publishUpdateConnectionMutex.Lock()
	defer fake.publishUpdateConnectionMutex.Unlock()
	fake.PublishUpdateConnectionStub = stub
}

func (fake *FakeIBus) PublishUpdateConnectionArgsForCall(i int) (context.Context, *opts.ConnectionOptions) {
	fake.publishUpdateConnectionMutex.RLock()
	defer fake.publishUpdateConnectionMutex.RUnlock()
	argsForCall := fake.publishUpdateConnectionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeIBus) PublishUpdateConnectionReturns(result1 error) {
	fake.publishUpdateConnectionMutex.Lock()
	defer fake.publishUpdateConnectionMutex.Unlock()
	fake.PublishUpdateConnectionStub = nil
	fake.publishUpdateConnectionReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBus) PublishUpdateConnectionReturnsOnCall(i int, result1 error) {
	fake.publishUpdateConnectionMutex.Lock()
	defer fake.publishUpdateConnectionMutex.Unlock()
	fake.PublishUpdateConnectionStub = nil
	if fake.publishUpdateConnectionReturnsOnCall == nil {
		fake.publishUpdateConnectionReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.publishUpdateConnectionReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBus) PublishUpdateDynamic(arg1 context.Context, arg2 *opts.DynamicOptions) error {
	fake.publishUpdateDynamicMutex.Lock()
	ret, specificReturn := fake.publishUpdateDynamicReturnsOnCall[len(fake.publishUpdateDynamicArgsForCall)]
	fake.publishUpdateDynamicArgsForCall = append(fake.publishUpdateDynamicArgsForCall, struct {
		arg1 context.Context
		arg2 *opts.DynamicOptions
	}{arg1, arg2})
	stub := fake.PublishUpdateDynamicStub
	fakeReturns := fake.publishUpdateDynamicReturns
	fake.recordInvocation("PublishUpdateDynamic", []interface{}{arg1, arg2})
	fake.publishUpdateDynamicMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIBus) PublishUpdateDynamicCallCount() int {
	fake.publishUpdateDynamicMutex.RLock()
	defer fake.publishUpdateDynamicMutex.RUnlock()
	return len(fake.publishUpdateDynamicArgsForCall)
}

func (fake *FakeIBus) PublishUpdateDynamicCalls(stub func(context.Context, *opts.DynamicOptions) error) {
	fake.publishUpdateDynamicMutex.Lock()
	defer fake.publishUpdateDynamicMutex.Unlock()
	fake.PublishUpdateDynamicStub = stub
}

func (fake *FakeIBus) PublishUpdateDynamicArgsForCall(i int) (context.Context, *opts.DynamicOptions) {
	fake.publishUpdateDynamicMutex.RLock()
	defer fake.publishUpdateDynamicMutex.RUnlock()
	argsForCall := fake.publishUpdateDynamicArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeIBus) PublishUpdateDynamicReturns(result1 error) {
	fake.publishUpdateDynamicMutex.Lock()
	defer fake.publishUpdateDynamicMutex.Unlock()
	fake.PublishUpdateDynamicStub = nil
	fake.publishUpdateDynamicReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBus) PublishUpdateDynamicReturnsOnCall(i int, result1 error) {
	fake.publishUpdateDynamicMutex.Lock()
	defer fake.publishUpdateDynamicMutex.Unlock()
	fake.PublishUpdateDynamicStub = nil
	if fake.publishUpdateDynamicReturnsOnCall == nil {
		fake.publishUpdateDynamicReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.publishUpdateDynamicReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBus) PublishUpdateRelay(arg1 context.Context, arg2 *opts.RelayOptions) error {
	fake.publishUpdateRelayMutex.Lock()
	ret, specificReturn := fake.publishUpdateRelayReturnsOnCall[len(fake.publishUpdateRelayArgsForCall)]
	fake.publishUpdateRelayArgsForCall = append(fake.publishUpdateRelayArgsForCall, struct {
		arg1 context.Context
		arg2 *opts.RelayOptions
	}{arg1, arg2})
	stub := fake.PublishUpdateRelayStub
	fakeReturns := fake.publishUpdateRelayReturns
	fake.recordInvocation("PublishUpdateRelay", []interface{}{arg1, arg2})
	fake.publishUpdateRelayMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIBus) PublishUpdateRelayCallCount() int {
	fake.publishUpdateRelayMutex.RLock()
	defer fake.publishUpdateRelayMutex.RUnlock()
	return len(fake.publishUpdateRelayArgsForCall)
}

func (fake *FakeIBus) PublishUpdateRelayCalls(stub func(context.Context, *opts.RelayOptions) error) {
	fake.publishUpdateRelayMutex.Lock()
	defer fake.publishUpdateRelayMutex.Unlock()
	fake.PublishUpdateRelayStub = stub
}

func (fake *FakeIBus) PublishUpdateRelayArgsForCall(i int) (context.Context, *opts.RelayOptions) {
	fake.publishUpdateRelayMutex.RLock()
	defer fake.publishUpdateRelayMutex.RUnlock()
	argsForCall := fake.publishUpdateRelayArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeIBus) PublishUpdateRelayReturns(result1 error) {
	fake.publishUpdateRelayMutex.Lock()
	defer fake.publishUpdateRelayMutex.Unlock()
	fake.PublishUpdateRelayStub = nil
	fake.publishUpdateRelayReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBus) PublishUpdateRelayReturnsOnCall(i int, result1 error) {
	fake.publishUpdateRelayMutex.Lock()
	defer fake.publishUpdateRelayMutex.Unlock()
	fake.PublishUpdateRelayStub = nil
	if fake.publishUpdateRelayReturnsOnCall == nil {
		fake.publishUpdateRelayReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.publishUpdateRelayReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBus) Start(arg1 context.Context) error {
	fake.startMutex.Lock()
	ret, specificReturn := fake.startReturnsOnCall[len(fake.startArgsForCall)]
	fake.startArgsForCall = append(fake.startArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.StartStub
	fakeReturns := fake.startReturns
	fake.recordInvocation("Start", []interface{}{arg1})
	fake.startMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIBus) StartCallCount() int {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	return len(fake.startArgsForCall)
}

func (fake *FakeIBus) StartCalls(stub func(context.Context) error) {
	fake.startMutex.Lock()
	defer fake.startMutex.Unlock()
	fake.StartStub = stub
}

func (fake *FakeIBus) StartArgsForCall(i int) context.Context {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	argsForCall := fake.startArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeIBus) StartReturns(result1 error) {
	fake.startMutex.Lock()
	defer fake.startMutex.Unlock()
	fake.StartStub = nil
	fake.startReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBus) StartReturnsOnCall(i int, result1 error) {
	fake.startMutex.Lock()
	defer fake.startMutex.Unlock()
	fake.StartStub = nil
	if fake.startReturnsOnCall == nil {
		fake.startReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.startReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBus) Stop() error {
	fake.stopMutex.Lock()
	ret, specificReturn := fake.stopReturnsOnCall[len(fake.stopArgsForCall)]
	fake.stopArgsForCall = append(fake.stopArgsForCall, struct {
	}{})
	stub := fake.StopStub
	fakeReturns := fake.stopReturns
	fake.recordInvocation("Stop", []interface{}{})
	fake.stopMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIBus) StopCallCount() int {
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	return len(fake.stopArgsForCall)
}

func (fake *FakeIBus) StopCalls(stub func() error) {
	fake.stopMutex.Lock()
	defer fake.stopMutex.Unlock()
	fake.StopStub = stub
}

func (fake *FakeIBus) StopReturns(result1 error) {
	fake.stopMutex.Lock()
	defer fake.stopMutex.Unlock()
	fake.StopStub = nil
	fake.stopReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBus) StopReturnsOnCall(i int, result1 error) {
	fake.stopMutex.Lock()
	defer fake.stopMutex.Unlock()
	fake.StopStub = nil
	if fake.stopReturnsOnCall == nil {
		fake.stopReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.stopReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIBus) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.publishCreateConnectionMutex.RLock()
	defer fake.publishCreateConnectionMutex.RUnlock()
	fake.publishCreateDynamicMutex.RLock()
	defer fake.publishCreateDynamicMutex.RUnlock()
	fake.publishCreateRelayMutex.RLock()
	defer fake.publishCreateRelayMutex.RUnlock()
	fake.publishDeleteConnectionMutex.RLock()
	defer fake.publishDeleteConnectionMutex.RUnlock()
	fake.publishDeleteDynamicMutex.RLock()
	defer fake.publishDeleteDynamicMutex.RUnlock()
	fake.publishDeleteRelayMutex.RLock()
	defer fake.publishDeleteRelayMutex.RUnlock()
	fake.publishResumeDynamicMutex.RLock()
	defer fake.publishResumeDynamicMutex.RUnlock()
	fake.publishResumeRelayMutex.RLock()
	defer fake.publishResumeRelayMutex.RUnlock()
	fake.publishStopDynamicMutex.RLock()
	defer fake.publishStopDynamicMutex.RUnlock()
	fake.publishStopRelayMutex.RLock()
	defer fake.publishStopRelayMutex.RUnlock()
	fake.publishUpdateConnectionMutex.RLock()
	defer fake.publishUpdateConnectionMutex.RUnlock()
	fake.publishUpdateDynamicMutex.RLock()
	defer fake.publishUpdateDynamicMutex.RUnlock()
	fake.publishUpdateRelayMutex.RLock()
	defer fake.publishUpdateRelayMutex.RUnlock()
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeIBus) recordInvocation(key string, args []interface{}) {
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

var _ bus.IBus = new(FakeIBus)
