// Code generated by mockery. DO NOT EDIT.

package mocks

import mock "github.com/Laisky/testify/mock"

// IfaceWithBuildTagInFilename is an autogenerated mock type for the IfaceWithBuildTagInFilename type
type IfaceWithBuildTagInFilename struct {
	mock.Mock
}

// Sprintf provides a mock function with given fields: format, a
func (_m *IfaceWithBuildTagInFilename) Sprintf(format string, a ...interface{}) string {
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

type mockConstructorTestingTNewIfaceWithBuildTagInFilename interface {
	mock.TestingT
	Cleanup(func())
}

// NewIfaceWithBuildTagInFilename creates a new instance of IfaceWithBuildTagInFilename. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIfaceWithBuildTagInFilename(t mockConstructorTestingTNewIfaceWithBuildTagInFilename) *IfaceWithBuildTagInFilename {
	mock := &IfaceWithBuildTagInFilename{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
