// Copyright (c) Ultraviolet
// SPDX-License-Identifier: Apache-2.0

// Code generated by mockery v2.53.3. DO NOT EDIT.

package mocks

import (
	agent "github.com/ultravioletrs/cocos/agent"
	attestation "github.com/ultravioletrs/cocos/pkg/attestation"

	context "context"

	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

type Service_Expecter struct {
	mock *mock.Mock
}

func (_m *Service) EXPECT() *Service_Expecter {
	return &Service_Expecter{mock: &_m.Mock}
}

// Algo provides a mock function with given fields: ctx, algorithm
func (_m *Service) Algo(ctx context.Context, algorithm agent.Algorithm) error {
	ret := _m.Called(ctx, algorithm)

	if len(ret) == 0 {
		panic("no return value specified for Algo")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, agent.Algorithm) error); ok {
		r0 = rf(ctx, algorithm)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Service_Algo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Algo'
type Service_Algo_Call struct {
	*mock.Call
}

// Algo is a helper method to define mock.On call
//   - ctx context.Context
//   - algorithm agent.Algorithm
func (_e *Service_Expecter) Algo(ctx interface{}, algorithm interface{}) *Service_Algo_Call {
	return &Service_Algo_Call{Call: _e.mock.On("Algo", ctx, algorithm)}
}

func (_c *Service_Algo_Call) Run(run func(ctx context.Context, algorithm agent.Algorithm)) *Service_Algo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(agent.Algorithm))
	})
	return _c
}

func (_c *Service_Algo_Call) Return(_a0 error) *Service_Algo_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Service_Algo_Call) RunAndReturn(run func(context.Context, agent.Algorithm) error) *Service_Algo_Call {
	_c.Call.Return(run)
	return _c
}

// Attestation provides a mock function with given fields: ctx, reportData, nonce, attType
func (_m *Service) Attestation(ctx context.Context, reportData [64]byte, nonce [32]byte, attType attestation.PlatformType) ([]byte, error) {
	ret := _m.Called(ctx, reportData, nonce, attType)

	if len(ret) == 0 {
		panic("no return value specified for Attestation")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, [64]byte, [32]byte, attestation.PlatformType) ([]byte, error)); ok {
		return rf(ctx, reportData, nonce, attType)
	}
	if rf, ok := ret.Get(0).(func(context.Context, [64]byte, [32]byte, attestation.PlatformType) []byte); ok {
		r0 = rf(ctx, reportData, nonce, attType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, [64]byte, [32]byte, attestation.PlatformType) error); ok {
		r1 = rf(ctx, reportData, nonce, attType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Service_Attestation_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Attestation'
type Service_Attestation_Call struct {
	*mock.Call
}

// Attestation is a helper method to define mock.On call
//   - ctx context.Context
//   - reportData [64]byte
//   - nonce [32]byte
//   - attType attestation.PlatformType
func (_e *Service_Expecter) Attestation(ctx interface{}, reportData interface{}, nonce interface{}, attType interface{}) *Service_Attestation_Call {
	return &Service_Attestation_Call{Call: _e.mock.On("Attestation", ctx, reportData, nonce, attType)}
}

func (_c *Service_Attestation_Call) Run(run func(ctx context.Context, reportData [64]byte, nonce [32]byte, attType attestation.PlatformType)) *Service_Attestation_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([64]byte), args[2].([32]byte), args[3].(attestation.PlatformType))
	})
	return _c
}

func (_c *Service_Attestation_Call) Return(_a0 []byte, _a1 error) *Service_Attestation_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Service_Attestation_Call) RunAndReturn(run func(context.Context, [64]byte, [32]byte, attestation.PlatformType) ([]byte, error)) *Service_Attestation_Call {
	_c.Call.Return(run)
	return _c
}

// AttestationResult provides a mock function with given fields: ctx, nonce, attType
func (_m *Service) AttestationResult(ctx context.Context, nonce [32]byte, attType attestation.PlatformType) ([]byte, error) {
	ret := _m.Called(ctx, nonce, attType)

	if len(ret) == 0 {
		panic("no return value specified for AttestationResult")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, [32]byte, attestation.PlatformType) ([]byte, error)); ok {
		return rf(ctx, nonce, attType)
	}
	if rf, ok := ret.Get(0).(func(context.Context, [32]byte, attestation.PlatformType) []byte); ok {
		r0 = rf(ctx, nonce, attType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, [32]byte, attestation.PlatformType) error); ok {
		r1 = rf(ctx, nonce, attType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Service_AttestationResult_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AttestationResult'
type Service_AttestationResult_Call struct {
	*mock.Call
}

// AttestationResult is a helper method to define mock.On call
//   - ctx context.Context
//   - nonce [32]byte
//   - attType attestation.PlatformType
func (_e *Service_Expecter) AttestationResult(ctx interface{}, nonce interface{}, attType interface{}) *Service_AttestationResult_Call {
	return &Service_AttestationResult_Call{Call: _e.mock.On("AttestationResult", ctx, nonce, attType)}
}

func (_c *Service_AttestationResult_Call) Run(run func(ctx context.Context, nonce [32]byte, attType attestation.PlatformType)) *Service_AttestationResult_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([32]byte), args[2].(attestation.PlatformType))
	})
	return _c
}

func (_c *Service_AttestationResult_Call) Return(_a0 []byte, _a1 error) *Service_AttestationResult_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Service_AttestationResult_Call) RunAndReturn(run func(context.Context, [32]byte, attestation.PlatformType) ([]byte, error)) *Service_AttestationResult_Call {
	_c.Call.Return(run)
	return _c
}

// Data provides a mock function with given fields: ctx, dataset
func (_m *Service) Data(ctx context.Context, dataset agent.Dataset) error {
	ret := _m.Called(ctx, dataset)

	if len(ret) == 0 {
		panic("no return value specified for Data")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, agent.Dataset) error); ok {
		r0 = rf(ctx, dataset)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Service_Data_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Data'
type Service_Data_Call struct {
	*mock.Call
}

// Data is a helper method to define mock.On call
//   - ctx context.Context
//   - dataset agent.Dataset
func (_e *Service_Expecter) Data(ctx interface{}, dataset interface{}) *Service_Data_Call {
	return &Service_Data_Call{Call: _e.mock.On("Data", ctx, dataset)}
}

func (_c *Service_Data_Call) Run(run func(ctx context.Context, dataset agent.Dataset)) *Service_Data_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(agent.Dataset))
	})
	return _c
}

func (_c *Service_Data_Call) Return(_a0 error) *Service_Data_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Service_Data_Call) RunAndReturn(run func(context.Context, agent.Dataset) error) *Service_Data_Call {
	_c.Call.Return(run)
	return _c
}

// IMAMeasurements provides a mock function with given fields: ctx
func (_m *Service) IMAMeasurements(ctx context.Context) ([]byte, []byte, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for IMAMeasurements")
	}

	var r0 []byte
	var r1 []byte
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]byte, []byte, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []byte); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) []byte); ok {
		r1 = rf(ctx)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]byte)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context) error); ok {
		r2 = rf(ctx)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Service_IMAMeasurements_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IMAMeasurements'
type Service_IMAMeasurements_Call struct {
	*mock.Call
}

// IMAMeasurements is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Service_Expecter) IMAMeasurements(ctx interface{}) *Service_IMAMeasurements_Call {
	return &Service_IMAMeasurements_Call{Call: _e.mock.On("IMAMeasurements", ctx)}
}

func (_c *Service_IMAMeasurements_Call) Run(run func(ctx context.Context)) *Service_IMAMeasurements_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Service_IMAMeasurements_Call) Return(_a0 []byte, _a1 []byte, _a2 error) *Service_IMAMeasurements_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *Service_IMAMeasurements_Call) RunAndReturn(run func(context.Context) ([]byte, []byte, error)) *Service_IMAMeasurements_Call {
	_c.Call.Return(run)
	return _c
}

// InitComputation provides a mock function with given fields: ctx, cmp
func (_m *Service) InitComputation(ctx context.Context, cmp agent.Computation) error {
	ret := _m.Called(ctx, cmp)

	if len(ret) == 0 {
		panic("no return value specified for InitComputation")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, agent.Computation) error); ok {
		r0 = rf(ctx, cmp)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Service_InitComputation_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InitComputation'
type Service_InitComputation_Call struct {
	*mock.Call
}

// InitComputation is a helper method to define mock.On call
//   - ctx context.Context
//   - cmp agent.Computation
func (_e *Service_Expecter) InitComputation(ctx interface{}, cmp interface{}) *Service_InitComputation_Call {
	return &Service_InitComputation_Call{Call: _e.mock.On("InitComputation", ctx, cmp)}
}

func (_c *Service_InitComputation_Call) Run(run func(ctx context.Context, cmp agent.Computation)) *Service_InitComputation_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(agent.Computation))
	})
	return _c
}

func (_c *Service_InitComputation_Call) Return(_a0 error) *Service_InitComputation_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Service_InitComputation_Call) RunAndReturn(run func(context.Context, agent.Computation) error) *Service_InitComputation_Call {
	_c.Call.Return(run)
	return _c
}

// Result provides a mock function with given fields: ctx
func (_m *Service) Result(ctx context.Context) ([]byte, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Result")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]byte, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []byte); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Service_Result_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Result'
type Service_Result_Call struct {
	*mock.Call
}

// Result is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Service_Expecter) Result(ctx interface{}) *Service_Result_Call {
	return &Service_Result_Call{Call: _e.mock.On("Result", ctx)}
}

func (_c *Service_Result_Call) Run(run func(ctx context.Context)) *Service_Result_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Service_Result_Call) Return(_a0 []byte, _a1 error) *Service_Result_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Service_Result_Call) RunAndReturn(run func(context.Context) ([]byte, error)) *Service_Result_Call {
	_c.Call.Return(run)
	return _c
}

// State provides a mock function with no fields
func (_m *Service) State() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for State")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Service_State_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'State'
type Service_State_Call struct {
	*mock.Call
}

// State is a helper method to define mock.On call
func (_e *Service_Expecter) State() *Service_State_Call {
	return &Service_State_Call{Call: _e.mock.On("State")}
}

func (_c *Service_State_Call) Run(run func()) *Service_State_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Service_State_Call) Return(_a0 string) *Service_State_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Service_State_Call) RunAndReturn(run func() string) *Service_State_Call {
	_c.Call.Return(run)
	return _c
}

// StopComputation provides a mock function with given fields: ctx
func (_m *Service) StopComputation(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for StopComputation")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Service_StopComputation_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StopComputation'
type Service_StopComputation_Call struct {
	*mock.Call
}

// StopComputation is a helper method to define mock.On call
//   - ctx context.Context
func (_e *Service_Expecter) StopComputation(ctx interface{}) *Service_StopComputation_Call {
	return &Service_StopComputation_Call{Call: _e.mock.On("StopComputation", ctx)}
}

func (_c *Service_StopComputation_Call) Run(run func(ctx context.Context)) *Service_StopComputation_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *Service_StopComputation_Call) Return(_a0 error) *Service_StopComputation_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Service_StopComputation_Call) RunAndReturn(run func(context.Context) error) *Service_StopComputation_Call {
	_c.Call.Return(run)
	return _c
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
