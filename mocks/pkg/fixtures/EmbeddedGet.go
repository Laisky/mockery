// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	constraints "github.com/Laisky/mockery/v2/pkg/fixtures/constraints"
	mock "github.com/Laisky/testify/mock"
)

// EmbeddedGet is an autogenerated mock type for the EmbeddedGet type
type EmbeddedGet[T constraints.Signed] struct {
	mock.Mock
}

// Get provides a mock function with given fields:
func (_m *EmbeddedGet[T]) Get() T {
	ret := _m.Called()

	var r0 T
	if rf, ok := ret.Get(0).(func() T); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(T)
	}

	return r0
}

type mockConstructorTestingTNewEmbeddedGet interface {
	mock.TestingT
	Cleanup(func())
}

// NewEmbeddedGet creates a new instance of EmbeddedGet. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewEmbeddedGet[T constraints.Signed](t mockConstructorTestingTNewEmbeddedGet) *EmbeddedGet[T] {
	mock := &EmbeddedGet[T]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
