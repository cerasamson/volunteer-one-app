// Code generated by mockery v2.23.1. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"
)

// LikesController is an autogenerated mock type for the LikesController type
type LikesController struct {
	mock.Mock
}

// AllLikes provides a mock function with given fields: c
func (_m *LikesController) AllLikes(c *gin.Context) {
	_m.Called(c)
}

// CreateLike provides a mock function with given fields: c
func (_m *LikesController) CreateLike(c *gin.Context) {
	_m.Called(c)
}

// DeleteLike provides a mock function with given fields: c
func (_m *LikesController) DeleteLike(c *gin.Context) {
	_m.Called(c)
}

// FindLike provides a mock function with given fields: c
func (_m *LikesController) FindLike(c *gin.Context) {
	_m.Called(c)
}

// GetLikes provides a mock function with given fields: c
func (_m *LikesController) GetLikes(c *gin.Context) {
	_m.Called(c)
}

type mockConstructorTestingTNewLikesController interface {
	mock.TestingT
	Cleanup(func())
}

// NewLikesController creates a new instance of LikesController. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewLikesController(t mockConstructorTestingTNewLikesController) *LikesController {
	mock := &LikesController{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
