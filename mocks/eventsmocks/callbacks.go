// Code generated by mockery v2.42.0. DO NOT EDIT.

package eventsmocks

import (
	core "github.com/hyperledger/firefly/pkg/core"
	events "github.com/hyperledger/firefly/pkg/events"

	mock "github.com/stretchr/testify/mock"
)

// Callbacks is an autogenerated mock type for the Callbacks type
type Callbacks struct {
	mock.Mock
}

// ConnectionClosed provides a mock function with given fields: connID
func (_m *Callbacks) ConnectionClosed(connID string) {
	_m.Called(connID)
}

// DeliveryResponse provides a mock function with given fields: connID, inflight
func (_m *Callbacks) DeliveryResponse(connID string, inflight *core.EventDeliveryResponse) {
	_m.Called(connID, inflight)
}

// EphemeralSubscription provides a mock function with given fields: connID, namespace, filter, options
func (_m *Callbacks) EphemeralSubscription(connID string, namespace string, filter *core.SubscriptionFilter, options *core.SubscriptionOptions) error {
	ret := _m.Called(connID, namespace, filter, options)

	if len(ret) == 0 {
		panic("no return value specified for EphemeralSubscription")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, *core.SubscriptionFilter, *core.SubscriptionOptions) error); ok {
		r0 = rf(connID, namespace, filter, options)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RegisterConnection provides a mock function with given fields: connID, matcher
func (_m *Callbacks) RegisterConnection(connID string, matcher events.SubscriptionMatcher) error {
	ret := _m.Called(connID, matcher)

	if len(ret) == 0 {
		panic("no return value specified for RegisterConnection")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, events.SubscriptionMatcher) error); ok {
		r0 = rf(connID, matcher)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewCallbacks creates a new instance of Callbacks. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewCallbacks(t interface {
	mock.TestingT
	Cleanup(func())
}) *Callbacks {
	mock := &Callbacks{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
