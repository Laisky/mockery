// Code generated by mockery. DO NOT EDIT.

package mocks

import mock "github.com/Laisky/testify/mock"

// Signed is an autogenerated mock type for the Signed type
type Signed struct {
	mock.Mock
}

type mockConstructorTestingTNewSigned interface {
	mock.TestingT
	Cleanup(func())
}

// NewSigned creates a new instance of Signed. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSigned(t mockConstructorTestingTNewSigned) *Signed {
	mock := &Signed{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
