// Copyright (c) Ultraviolet
// SPDX-License-Identifier: Apache-2.0

// Code generated by mockery v2.53.2. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"

	manager "github.com/ultravioletrs/cocos/manager"

	mock "github.com/stretchr/testify/mock"
)

// ManagerServiceClient is an autogenerated mock type for the ManagerServiceClient type
type ManagerServiceClient struct {
	mock.Mock
}

type ManagerServiceClient_Expecter struct {
	mock *mock.Mock
}

func (_m *ManagerServiceClient) EXPECT() *ManagerServiceClient_Expecter {
	return &ManagerServiceClient_Expecter{mock: &_m.Mock}
}

// AttestationPolicy provides a mock function with given fields: ctx, in, opts
func (_m *ManagerServiceClient) AttestationPolicy(ctx context.Context, in *manager.AttestationPolicyReq, opts ...grpc.CallOption) (*manager.AttestationPolicyRes, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for AttestationPolicy")
	}

	var r0 *manager.AttestationPolicyRes
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *manager.AttestationPolicyReq, ...grpc.CallOption) (*manager.AttestationPolicyRes, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *manager.AttestationPolicyReq, ...grpc.CallOption) *manager.AttestationPolicyRes); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*manager.AttestationPolicyRes)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *manager.AttestationPolicyReq, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ManagerServiceClient_AttestationPolicy_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AttestationPolicy'
type ManagerServiceClient_AttestationPolicy_Call struct {
	*mock.Call
}

// AttestationPolicy is a helper method to define mock.On call
//   - ctx context.Context
//   - in *manager.AttestationPolicyReq
//   - opts ...grpc.CallOption
func (_e *ManagerServiceClient_Expecter) AttestationPolicy(ctx interface{}, in interface{}, opts ...interface{}) *ManagerServiceClient_AttestationPolicy_Call {
	return &ManagerServiceClient_AttestationPolicy_Call{Call: _e.mock.On("AttestationPolicy",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *ManagerServiceClient_AttestationPolicy_Call) Run(run func(ctx context.Context, in *manager.AttestationPolicyReq, opts ...grpc.CallOption)) *ManagerServiceClient_AttestationPolicy_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*manager.AttestationPolicyReq), variadicArgs...)
	})
	return _c
}

func (_c *ManagerServiceClient_AttestationPolicy_Call) Return(_a0 *manager.AttestationPolicyRes, _a1 error) *ManagerServiceClient_AttestationPolicy_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ManagerServiceClient_AttestationPolicy_Call) RunAndReturn(run func(context.Context, *manager.AttestationPolicyReq, ...grpc.CallOption) (*manager.AttestationPolicyRes, error)) *ManagerServiceClient_AttestationPolicy_Call {
	_c.Call.Return(run)
	return _c
}

// CreateVm provides a mock function with given fields: ctx, in, opts
func (_m *ManagerServiceClient) CreateVm(ctx context.Context, in *manager.CreateReq, opts ...grpc.CallOption) (*manager.CreateRes, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateVm")
	}

	var r0 *manager.CreateRes
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *manager.CreateReq, ...grpc.CallOption) (*manager.CreateRes, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *manager.CreateReq, ...grpc.CallOption) *manager.CreateRes); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*manager.CreateRes)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *manager.CreateReq, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ManagerServiceClient_CreateVm_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateVm'
type ManagerServiceClient_CreateVm_Call struct {
	*mock.Call
}

// CreateVm is a helper method to define mock.On call
//   - ctx context.Context
//   - in *manager.CreateReq
//   - opts ...grpc.CallOption
func (_e *ManagerServiceClient_Expecter) CreateVm(ctx interface{}, in interface{}, opts ...interface{}) *ManagerServiceClient_CreateVm_Call {
	return &ManagerServiceClient_CreateVm_Call{Call: _e.mock.On("CreateVm",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *ManagerServiceClient_CreateVm_Call) Run(run func(ctx context.Context, in *manager.CreateReq, opts ...grpc.CallOption)) *ManagerServiceClient_CreateVm_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*manager.CreateReq), variadicArgs...)
	})
	return _c
}

func (_c *ManagerServiceClient_CreateVm_Call) Return(_a0 *manager.CreateRes, _a1 error) *ManagerServiceClient_CreateVm_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ManagerServiceClient_CreateVm_Call) RunAndReturn(run func(context.Context, *manager.CreateReq, ...grpc.CallOption) (*manager.CreateRes, error)) *ManagerServiceClient_CreateVm_Call {
	_c.Call.Return(run)
	return _c
}

// RemoveVm provides a mock function with given fields: ctx, in, opts
func (_m *ManagerServiceClient) RemoveVm(ctx context.Context, in *manager.RemoveReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for RemoveVm")
	}

	var r0 *emptypb.Empty
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *manager.RemoveReq, ...grpc.CallOption) (*emptypb.Empty, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *manager.RemoveReq, ...grpc.CallOption) *emptypb.Empty); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*emptypb.Empty)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *manager.RemoveReq, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ManagerServiceClient_RemoveVm_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveVm'
type ManagerServiceClient_RemoveVm_Call struct {
	*mock.Call
}

// RemoveVm is a helper method to define mock.On call
//   - ctx context.Context
//   - in *manager.RemoveReq
//   - opts ...grpc.CallOption
func (_e *ManagerServiceClient_Expecter) RemoveVm(ctx interface{}, in interface{}, opts ...interface{}) *ManagerServiceClient_RemoveVm_Call {
	return &ManagerServiceClient_RemoveVm_Call{Call: _e.mock.On("RemoveVm",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *ManagerServiceClient_RemoveVm_Call) Run(run func(ctx context.Context, in *manager.RemoveReq, opts ...grpc.CallOption)) *ManagerServiceClient_RemoveVm_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*manager.RemoveReq), variadicArgs...)
	})
	return _c
}

func (_c *ManagerServiceClient_RemoveVm_Call) Return(_a0 *emptypb.Empty, _a1 error) *ManagerServiceClient_RemoveVm_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ManagerServiceClient_RemoveVm_Call) RunAndReturn(run func(context.Context, *manager.RemoveReq, ...grpc.CallOption) (*emptypb.Empty, error)) *ManagerServiceClient_RemoveVm_Call {
	_c.Call.Return(run)
	return _c
}

// SVMInfo provides a mock function with given fields: ctx, in, opts
func (_m *ManagerServiceClient) SVMInfo(ctx context.Context, in *manager.SVMInfoReq, opts ...grpc.CallOption) (*manager.SVMInfoRes, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for SVMInfo")
	}

	var r0 *manager.SVMInfoRes
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *manager.SVMInfoReq, ...grpc.CallOption) (*manager.SVMInfoRes, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *manager.SVMInfoReq, ...grpc.CallOption) *manager.SVMInfoRes); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*manager.SVMInfoRes)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *manager.SVMInfoReq, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ManagerServiceClient_SVMInfo_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SVMInfo'
type ManagerServiceClient_SVMInfo_Call struct {
	*mock.Call
}

// SVMInfo is a helper method to define mock.On call
//   - ctx context.Context
//   - in *manager.SVMInfoReq
//   - opts ...grpc.CallOption
func (_e *ManagerServiceClient_Expecter) SVMInfo(ctx interface{}, in interface{}, opts ...interface{}) *ManagerServiceClient_SVMInfo_Call {
	return &ManagerServiceClient_SVMInfo_Call{Call: _e.mock.On("SVMInfo",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *ManagerServiceClient_SVMInfo_Call) Run(run func(ctx context.Context, in *manager.SVMInfoReq, opts ...grpc.CallOption)) *ManagerServiceClient_SVMInfo_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*manager.SVMInfoReq), variadicArgs...)
	})
	return _c
}

func (_c *ManagerServiceClient_SVMInfo_Call) Return(_a0 *manager.SVMInfoRes, _a1 error) *ManagerServiceClient_SVMInfo_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ManagerServiceClient_SVMInfo_Call) RunAndReturn(run func(context.Context, *manager.SVMInfoReq, ...grpc.CallOption) (*manager.SVMInfoRes, error)) *ManagerServiceClient_SVMInfo_Call {
	_c.Call.Return(run)
	return _c
}

// NewManagerServiceClient creates a new instance of ManagerServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewManagerServiceClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *ManagerServiceClient {
	mock := &ManagerServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
