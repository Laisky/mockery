// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	mock "github.com/Laisky/testify/mock"
	constraints "github.com/Laisky/mockery/v2/pkg/fixtures/constraints"
)

// GetGeneric is an autogenerated mock type for the GetGeneric type
type GetGeneric[T constraints.Integer] struct {
	mock.Mock
}

// Get provides a mock function with given fields:
func (_m *GetGeneric[T]) Get() T {
	ret := _m.Called()

	var r0 T
	if rf, ok := ret.Get(0).(func() T); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(T)
	}

	return r0
}

type mockConstructorTestingTNewGetGeneric interface {
	mock.TestingT
	Cleanup(func())
}

// NewGetGeneric creates a new instance of GetGeneric. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGetGeneric[T constraints.Integer](t mockConstructorTestingTNewGetGeneric) *GetGeneric[T] {
	mock := &GetGeneric[T]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
