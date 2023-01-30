// Code generated by mockery. DO NOT EDIT.

package mocks

import mock "github.com/Laisky/testify/mock"

// ConsulLock is an autogenerated mock type for the ConsulLock type
type ConsulLock struct {
	mock.Mock
}

// Lock provides a mock function with given fields: _a0
func (_m *ConsulLock) Lock(_a0 <-chan struct{}) (<-chan struct{}, error) {
	ret := _m.Called(_a0)

	var r0 <-chan struct{}
	if rf, ok := ret.Get(0).(func(<-chan struct{}) <-chan struct{}); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan struct{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(<-chan struct{}) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Unlock provides a mock function with given fields:
func (_m *ConsulLock) Unlock() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewConsulLock interface {
	mock.TestingT
	Cleanup(func())
}

// NewConsulLock creates a new instance of ConsulLock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewConsulLock(t mockConstructorTestingTNewConsulLock) *ConsulLock {
	mock := &ConsulLock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
