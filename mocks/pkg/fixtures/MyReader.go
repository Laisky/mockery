// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	testing "testing"
)

// MyReader is an autogenerated mock type for the MyReader type
type MyReader struct {
	mock.Mock
}

// Read provides a mock function with given fields: p
func (_m *MyReader) Read(p []byte) (int, error) {
	ret := _m.Called(p)

	var r0 int
	if rf, ok := ret.Get(0).(func([]byte) int); ok {
		r0 = rf(p)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMyReader creates a new instance of MyReader. It also registers a cleanup function to assert the mocks expectations.
func NewMyReader(t testing.TB) *MyReader {
	mock := &MyReader{}

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
