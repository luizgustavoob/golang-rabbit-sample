// Code generated by mockery v2.50.4. DO NOT EDIT.

package mocks

import (
	api "github.com/golang-rabbit-sample/api-producer/internal/api"
	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// AddPerson provides a mock function with given fields: person
func (_m *Service) AddPerson(person *api.Person) (*api.Person, error) {
	ret := _m.Called(person)

	if len(ret) == 0 {
		panic("no return value specified for AddPerson")
	}

	var r0 *api.Person
	var r1 error
	if rf, ok := ret.Get(0).(func(*api.Person) (*api.Person, error)); ok {
		return rf(person)
	}
	if rf, ok := ret.Get(0).(func(*api.Person) *api.Person); ok {
		r0 = rf(person)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*api.Person)
		}
	}

	if rf, ok := ret.Get(1).(func(*api.Person) error); ok {
		r1 = rf(person)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewService(t interface {
	mock.TestingT
	Cleanup(func())
}) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
