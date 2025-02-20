// Copyright (c) Ultraviolet
// SPDX-License-Identifier: Apache-2.0

// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	sevsnp "github.com/google/go-sev-guest/proto/sevsnp"
	mock "github.com/stretchr/testify/mock"
)

// QuoteProvider is an autogenerated mock type for the QuoteProvider type
type QuoteProvider struct {
	mock.Mock
}

type QuoteProvider_Expecter struct {
	mock *mock.Mock
}

func (_m *QuoteProvider) EXPECT() *QuoteProvider_Expecter {
	return &QuoteProvider_Expecter{mock: &_m.Mock}
}

// GetRawQuote provides a mock function with given fields: reportData
func (_m *QuoteProvider) GetRawQuote(reportData [64]byte) ([]uint8, error) {
	ret := _m.Called(reportData)

	if len(ret) == 0 {
		panic("no return value specified for GetRawQuote")
	}

	var r0 []uint8
	var r1 error
	if rf, ok := ret.Get(0).(func([64]byte) ([]uint8, error)); ok {
		return rf(reportData)
	}
	if rf, ok := ret.Get(0).(func([64]byte) []uint8); ok {
		r0 = rf(reportData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]uint8)
		}
	}

	if rf, ok := ret.Get(1).(func([64]byte) error); ok {
		r1 = rf(reportData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QuoteProvider_GetRawQuote_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRawQuote'
type QuoteProvider_GetRawQuote_Call struct {
	*mock.Call
}

// GetRawQuote is a helper method to define mock.On call
//   - reportData [64]byte
func (_e *QuoteProvider_Expecter) GetRawQuote(reportData interface{}) *QuoteProvider_GetRawQuote_Call {
	return &QuoteProvider_GetRawQuote_Call{Call: _e.mock.On("GetRawQuote", reportData)}
}

func (_c *QuoteProvider_GetRawQuote_Call) Run(run func(reportData [64]byte)) *QuoteProvider_GetRawQuote_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([64]byte))
	})
	return _c
}

func (_c *QuoteProvider_GetRawQuote_Call) Return(_a0 []uint8, _a1 error) *QuoteProvider_GetRawQuote_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *QuoteProvider_GetRawQuote_Call) RunAndReturn(run func([64]byte) ([]uint8, error)) *QuoteProvider_GetRawQuote_Call {
	_c.Call.Return(run)
	return _c
}

// IsSupported provides a mock function with given fields:
func (_m *QuoteProvider) IsSupported() bool {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for IsSupported")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// QuoteProvider_IsSupported_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsSupported'
type QuoteProvider_IsSupported_Call struct {
	*mock.Call
}

// IsSupported is a helper method to define mock.On call
func (_e *QuoteProvider_Expecter) IsSupported() *QuoteProvider_IsSupported_Call {
	return &QuoteProvider_IsSupported_Call{Call: _e.mock.On("IsSupported")}
}

func (_c *QuoteProvider_IsSupported_Call) Run(run func()) *QuoteProvider_IsSupported_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *QuoteProvider_IsSupported_Call) Return(_a0 bool) *QuoteProvider_IsSupported_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *QuoteProvider_IsSupported_Call) RunAndReturn(run func() bool) *QuoteProvider_IsSupported_Call {
	_c.Call.Return(run)
	return _c
}

// Product provides a mock function with given fields:
func (_m *QuoteProvider) Product() *sevsnp.SevProduct {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Product")
	}

	var r0 *sevsnp.SevProduct
	if rf, ok := ret.Get(0).(func() *sevsnp.SevProduct); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sevsnp.SevProduct)
		}
	}

	return r0
}

// QuoteProvider_Product_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Product'
type QuoteProvider_Product_Call struct {
	*mock.Call
}

// Product is a helper method to define mock.On call
func (_e *QuoteProvider_Expecter) Product() *QuoteProvider_Product_Call {
	return &QuoteProvider_Product_Call{Call: _e.mock.On("Product")}
}

func (_c *QuoteProvider_Product_Call) Run(run func()) *QuoteProvider_Product_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *QuoteProvider_Product_Call) Return(_a0 *sevsnp.SevProduct) *QuoteProvider_Product_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *QuoteProvider_Product_Call) RunAndReturn(run func() *sevsnp.SevProduct) *QuoteProvider_Product_Call {
	_c.Call.Return(run)
	return _c
}

// NewQuoteProvider creates a new instance of QuoteProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewQuoteProvider(t interface {
	mock.TestingT
	Cleanup(func())
}) *QuoteProvider {
	mock := &QuoteProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
