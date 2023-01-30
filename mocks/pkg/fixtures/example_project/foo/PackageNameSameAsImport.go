// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	foo "github.com/Laisky/mockery/v2/pkg/fixtures/example_project/bar/foo"
	mock "github.com/Laisky/testify/mock"
)

// PackageNameSameAsImport is an autogenerated mock type for the PackageNameSameAsImport type
type PackageNameSameAsImport struct {
	mock.Mock
}

// NewClient provides a mock function with given fields:
func (_m *PackageNameSameAsImport) NewClient() foo.Client {
	ret := _m.Called()

	var r0 foo.Client
	if rf, ok := ret.Get(0).(func() foo.Client); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(foo.Client)
		}
	}

	return r0
}

type mockConstructorTestingTNewPackageNameSameAsImport interface {
	mock.TestingT
	Cleanup(func())
}

// NewPackageNameSameAsImport creates a new instance of PackageNameSameAsImport. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPackageNameSameAsImport(t mockConstructorTestingTNewPackageNameSameAsImport) *PackageNameSameAsImport {
	mock := &PackageNameSameAsImport{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
