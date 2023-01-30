// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	io "io"

	mock "github.com/Laisky/testify/mock"
)

// RequesterVariadicOneArgument is an autogenerated mock type for the RequesterVariadic type
type RequesterVariadicOneArgument struct {
	mock.Mock
}

// Get provides a mock function with given fields: values
func (_m *RequesterVariadicOneArgument) Get(values ...string) bool {
	ret := _m.Called(values)

	var r0 bool
	if rf, ok := ret.Get(0).(func(...string) bool); ok {
		r0 = rf(values...)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MultiWriteToFile provides a mock function with given fields: filename, w
func (_m *RequesterVariadicOneArgument) MultiWriteToFile(filename string, w ...io.Writer) string {
	ret := _m.Called(filename, w)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, ...io.Writer) string); ok {
		r0 = rf(filename, w...)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// OneInterface provides a mock function with given fields: a
func (_m *RequesterVariadicOneArgument) OneInterface(a ...interface{}) bool {
	ret := _m.Called(a)

	var r0 bool
	if rf, ok := ret.Get(0).(func(...interface{}) bool); ok {
		r0 = rf(a...)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Sprintf provides a mock function with given fields: format, a
func (_m *RequesterVariadicOneArgument) Sprintf(format string, a ...interface{}) string {
	ret := _m.Called(format, a)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, ...interface{}) string); ok {
		r0 = rf(format, a...)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type mockConstructorTestingTNewRequesterVariadicOneArgument interface {
	mock.TestingT
	Cleanup(func())
}

// NewRequesterVariadicOneArgument creates a new instance of RequesterVariadicOneArgument. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRequesterVariadicOneArgument(t mockConstructorTestingTNewRequesterVariadicOneArgument) *RequesterVariadicOneArgument {
	mock := &RequesterVariadicOneArgument{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
