// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"sync"
	"time"

	"github.com/zebfross/syncthing-mobile/lib/events"
)

type BufferedSubscription struct {
	MaskStub        func() events.EventType
	maskMutex       sync.RWMutex
	maskArgsForCall []struct {
	}
	maskReturns struct {
		result1 events.EventType
	}
	maskReturnsOnCall map[int]struct {
		result1 events.EventType
	}
	SinceStub        func(int, []events.Event, time.Duration) []events.Event
	sinceMutex       sync.RWMutex
	sinceArgsForCall []struct {
		arg1 int
		arg2 []events.Event
		arg3 time.Duration
	}
	sinceReturns struct {
		result1 []events.Event
	}
	sinceReturnsOnCall map[int]struct {
		result1 []events.Event
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *BufferedSubscription) Mask() events.EventType {
	fake.maskMutex.Lock()
	ret, specificReturn := fake.maskReturnsOnCall[len(fake.maskArgsForCall)]
	fake.maskArgsForCall = append(fake.maskArgsForCall, struct {
	}{})
	stub := fake.MaskStub
	fakeReturns := fake.maskReturns
	fake.recordInvocation("Mask", []interface{}{})
	fake.maskMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *BufferedSubscription) MaskCallCount() int {
	fake.maskMutex.RLock()
	defer fake.maskMutex.RUnlock()
	return len(fake.maskArgsForCall)
}

func (fake *BufferedSubscription) MaskCalls(stub func() events.EventType) {
	fake.maskMutex.Lock()
	defer fake.maskMutex.Unlock()
	fake.MaskStub = stub
}

func (fake *BufferedSubscription) MaskReturns(result1 events.EventType) {
	fake.maskMutex.Lock()
	defer fake.maskMutex.Unlock()
	fake.MaskStub = nil
	fake.maskReturns = struct {
		result1 events.EventType
	}{result1}
}

func (fake *BufferedSubscription) MaskReturnsOnCall(i int, result1 events.EventType) {
	fake.maskMutex.Lock()
	defer fake.maskMutex.Unlock()
	fake.MaskStub = nil
	if fake.maskReturnsOnCall == nil {
		fake.maskReturnsOnCall = make(map[int]struct {
			result1 events.EventType
		})
	}
	fake.maskReturnsOnCall[i] = struct {
		result1 events.EventType
	}{result1}
}

func (fake *BufferedSubscription) Since(arg1 int, arg2 []events.Event, arg3 time.Duration) []events.Event {
	var arg2Copy []events.Event
	if arg2 != nil {
		arg2Copy = make([]events.Event, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.sinceMutex.Lock()
	ret, specificReturn := fake.sinceReturnsOnCall[len(fake.sinceArgsForCall)]
	fake.sinceArgsForCall = append(fake.sinceArgsForCall, struct {
		arg1 int
		arg2 []events.Event
		arg3 time.Duration
	}{arg1, arg2Copy, arg3})
	stub := fake.SinceStub
	fakeReturns := fake.sinceReturns
	fake.recordInvocation("Since", []interface{}{arg1, arg2Copy, arg3})
	fake.sinceMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *BufferedSubscription) SinceCallCount() int {
	fake.sinceMutex.RLock()
	defer fake.sinceMutex.RUnlock()
	return len(fake.sinceArgsForCall)
}

func (fake *BufferedSubscription) SinceCalls(stub func(int, []events.Event, time.Duration) []events.Event) {
	fake.sinceMutex.Lock()
	defer fake.sinceMutex.Unlock()
	fake.SinceStub = stub
}

func (fake *BufferedSubscription) SinceArgsForCall(i int) (int, []events.Event, time.Duration) {
	fake.sinceMutex.RLock()
	defer fake.sinceMutex.RUnlock()
	argsForCall := fake.sinceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *BufferedSubscription) SinceReturns(result1 []events.Event) {
	fake.sinceMutex.Lock()
	defer fake.sinceMutex.Unlock()
	fake.SinceStub = nil
	fake.sinceReturns = struct {
		result1 []events.Event
	}{result1}
}

func (fake *BufferedSubscription) SinceReturnsOnCall(i int, result1 []events.Event) {
	fake.sinceMutex.Lock()
	defer fake.sinceMutex.Unlock()
	fake.SinceStub = nil
	if fake.sinceReturnsOnCall == nil {
		fake.sinceReturnsOnCall = make(map[int]struct {
			result1 []events.Event
		})
	}
	fake.sinceReturnsOnCall[i] = struct {
		result1 []events.Event
	}{result1}
}

func (fake *BufferedSubscription) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.maskMutex.RLock()
	defer fake.maskMutex.RUnlock()
	fake.sinceMutex.RLock()
	defer fake.sinceMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *BufferedSubscription) recordInvocation(key string, args []interface{}) {
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

var _ events.BufferedSubscription = new(BufferedSubscription)
