// Copyright (c) Ultraviolet
// SPDX-License-Identifier: Apache-2.0

// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	cvms "github.com/ultravioletrs/cocos/agent/cvms"

	storage "github.com/ultravioletrs/cocos/agent/cvms/api/grpc/storage"
)

// Storage is an autogenerated mock type for the Storage type
type Storage struct {
	mock.Mock
}

type Storage_Expecter struct {
	mock *mock.Mock
}

func (_m *Storage) EXPECT() *Storage_Expecter {
	return &Storage_Expecter{mock: &_m.Mock}
}

// Add provides a mock function with given fields: msg
func (_m *Storage) Add(msg *cvms.ClientStreamMessage) error {
	ret := _m.Called(msg)

	if len(ret) == 0 {
		panic("no return value specified for Add")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*cvms.ClientStreamMessage) error); ok {
		r0 = rf(msg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Storage_Add_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Add'
type Storage_Add_Call struct {
	*mock.Call
}

// Add is a helper method to define mock.On call
//   - msg *cvms.ClientStreamMessage
func (_e *Storage_Expecter) Add(msg interface{}) *Storage_Add_Call {
	return &Storage_Add_Call{Call: _e.mock.On("Add", msg)}
}

func (_c *Storage_Add_Call) Run(run func(msg *cvms.ClientStreamMessage)) *Storage_Add_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*cvms.ClientStreamMessage))
	})
	return _c
}

func (_c *Storage_Add_Call) Return(_a0 error) *Storage_Add_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Storage_Add_Call) RunAndReturn(run func(*cvms.ClientStreamMessage) error) *Storage_Add_Call {
	_c.Call.Return(run)
	return _c
}

// Clear provides a mock function with no fields
func (_m *Storage) Clear() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Clear")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Storage_Clear_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Clear'
type Storage_Clear_Call struct {
	*mock.Call
}

// Clear is a helper method to define mock.On call
func (_e *Storage_Expecter) Clear() *Storage_Clear_Call {
	return &Storage_Clear_Call{Call: _e.mock.On("Clear")}
}

func (_c *Storage_Clear_Call) Run(run func()) *Storage_Clear_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Storage_Clear_Call) Return(_a0 error) *Storage_Clear_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Storage_Clear_Call) RunAndReturn(run func() error) *Storage_Clear_Call {
	_c.Call.Return(run)
	return _c
}

// Load provides a mock function with no fields
func (_m *Storage) Load() ([]storage.Message, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Load")
	}

	var r0 []storage.Message
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]storage.Message, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []storage.Message); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]storage.Message)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Storage_Load_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Load'
type Storage_Load_Call struct {
	*mock.Call
}

// Load is a helper method to define mock.On call
func (_e *Storage_Expecter) Load() *Storage_Load_Call {
	return &Storage_Load_Call{Call: _e.mock.On("Load")}
}

func (_c *Storage_Load_Call) Run(run func()) *Storage_Load_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Storage_Load_Call) Return(_a0 []storage.Message, _a1 error) *Storage_Load_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Storage_Load_Call) RunAndReturn(run func() ([]storage.Message, error)) *Storage_Load_Call {
	_c.Call.Return(run)
	return _c
}

// Save provides a mock function with given fields: messages
func (_m *Storage) Save(messages []storage.Message) error {
	ret := _m.Called(messages)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func([]storage.Message) error); ok {
		r0 = rf(messages)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Storage_Save_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Save'
type Storage_Save_Call struct {
	*mock.Call
}

// Save is a helper method to define mock.On call
//   - messages []storage.Message
func (_e *Storage_Expecter) Save(messages interface{}) *Storage_Save_Call {
	return &Storage_Save_Call{Call: _e.mock.On("Save", messages)}
}

func (_c *Storage_Save_Call) Run(run func(messages []storage.Message)) *Storage_Save_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]storage.Message))
	})
	return _c
}

func (_c *Storage_Save_Call) Return(_a0 error) *Storage_Save_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Storage_Save_Call) RunAndReturn(run func([]storage.Message) error) *Storage_Save_Call {
	_c.Call.Return(run)
	return _c
}

// NewStorage creates a new instance of Storage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStorage(t interface {
	mock.TestingT
	Cleanup(func())
}) *Storage {
	mock := &Storage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
