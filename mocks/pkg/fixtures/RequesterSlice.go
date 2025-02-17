// Code generated by mockery. DO NOT EDIT.

package mocks

import mock "github.com/Laisky/testify/mock"

// RequesterSlice is an autogenerated mock type for the RequesterSlice type
type RequesterSlice struct {
	mock.Mock
}

// Get provides a mock function with given fields: path
func (_m *RequesterSlice) Get(path string) ([]string, error) {
	ret := _m.Called(path)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(path)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRequesterSlice interface {
	mock.TestingT
	Cleanup(func())
}

// NewRequesterSlice creates a new instance of RequesterSlice. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRequesterSlice(t mockConstructorTestingTNewRequesterSlice) *RequesterSlice {
	mock := &RequesterSlice{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
