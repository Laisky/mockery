// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	testing "testing"

	mock "github.com/stretchr/testify/mock"
)

// IfaceWithBuildTagInComment is an autogenerated mock type for the IfaceWithBuildTagInComment type
type IfaceWithBuildTagInComment struct {
	mock.Mock
}

// Sprintf provides a mock function with given fields: format, a
func (_m *IfaceWithBuildTagInComment) Sprintf(format string, a ...interface{}) string {
	var _ca []interface{}
	_ca = append(_ca, format)
	_ca = append(_ca, a...)
	ret := _m.Called(_ca...)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, ...interface{}) string); ok {
		r0 = rf(format, a...)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// NewIfaceWithBuildTagInComment creates a new instance of IfaceWithBuildTagInComment. It also registers a cleanup function to assert the mocks expectations.
func NewIfaceWithBuildTagInComment(t testing.TB) *IfaceWithBuildTagInComment {
	mock := &IfaceWithBuildTagInComment{}

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
