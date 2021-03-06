// Code generated by mockery v2.7.4. DO NOT EDIT.

package mocks

import (
	context "context"

	repository "github.com/aryannr97/data-server/pkg/repository"
	mock "github.com/stretchr/testify/mock"
)

// MockUserRepository is an autogenerated mock type for the IUserRepository type
type MockUserRepository struct {
	mock.Mock
}

// AddUser provides a mock function with given fields: _a0, _a1
func (_m *MockUserRepository) AddUser(_a0 context.Context, _a1 repository.UserEntity) (string, error) {
	ret := _m.Called(_a0, _a1)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, repository.UserEntity) string); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, repository.UserEntity) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteUser provides a mock function with given fields: _a0, _a1
func (_m *MockUserRepository) DeleteUser(_a0 context.Context, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetUser provides a mock function with given fields: _a0, _a1
func (_m *MockUserRepository) GetUser(_a0 context.Context, _a1 string) (repository.UserDocument, error) {
	ret := _m.Called(_a0, _a1)

	var r0 repository.UserDocument
	if rf, ok := ret.Get(0).(func(context.Context, string) repository.UserDocument); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(repository.UserDocument)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: _a0, _a1, _a2
func (_m *MockUserRepository) UpdateUser(_a0 context.Context, _a1 string, _a2 repository.UserEntity) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, repository.UserEntity) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
