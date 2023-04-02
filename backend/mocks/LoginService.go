// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	models "github.com/VolunteerOne/volunteer-one-app/backend/models"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// LoginService is an autogenerated mock type for the LoginService type
type LoginService struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: _a0
func (_m *LoginService) CreateUser(_a0 models.Users) (models.Users, error) {
	ret := _m.Called(_a0)

	var r0 models.Users
	var r1 error
	if rf, ok := ret.Get(0).(func(models.Users) (models.Users, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(models.Users) models.Users); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(models.Users)
	}

	if rf, ok := ret.Get(1).(func(models.Users) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindUserFromEmail provides a mock function with given fields: _a0, _a1
func (_m *LoginService) FindUserFromEmail(_a0 string, _a1 models.Users) (models.Users, error) {
	ret := _m.Called(_a0, _a1)

	var r0 models.Users
	var r1 error
	if rf, ok := ret.Get(0).(func(string, models.Users) (models.Users, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(string, models.Users) models.Users); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(models.Users)
	}

	if rf, ok := ret.Get(1).(func(string, models.Users) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HashPassword provides a mock function with given fields: _a0
func (_m *LoginService) HashPassword(_a0 []byte) ([]byte, error) {
	ret := _m.Called(_a0)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte) ([]byte, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func([]byte) []byte); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveResetCodeToUser provides a mock function with given fields: _a0, _a1
func (_m *LoginService) SaveResetCodeToUser(_a0 uuid.UUID, _a1 models.Users) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(uuid.UUID, models.Users) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewLoginService interface {
	mock.TestingT
	Cleanup(func())
}

// NewLoginService creates a new instance of LoginService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewLoginService(t mockConstructorTestingTNewLoginService) *LoginService {
	mock := &LoginService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}