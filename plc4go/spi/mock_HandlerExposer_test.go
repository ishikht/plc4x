/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

// Code generated by mockery v2.32.4. DO NOT EDIT.

package spi

import mock "github.com/stretchr/testify/mock"

// MockHandlerExposer is an autogenerated mock type for the HandlerExposer type
type MockHandlerExposer struct {
	mock.Mock
}

type MockHandlerExposer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockHandlerExposer) EXPECT() *MockHandlerExposer_Expecter {
	return &MockHandlerExposer_Expecter{mock: &_m.Mock}
}

// GetPlcTagHandler provides a mock function with given fields:
func (_m *MockHandlerExposer) GetPlcTagHandler() PlcTagHandler {
	ret := _m.Called()

	var r0 PlcTagHandler
	if rf, ok := ret.Get(0).(func() PlcTagHandler); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(PlcTagHandler)
		}
	}

	return r0
}

// MockHandlerExposer_GetPlcTagHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPlcTagHandler'
type MockHandlerExposer_GetPlcTagHandler_Call struct {
	*mock.Call
}

// GetPlcTagHandler is a helper method to define mock.On call
func (_e *MockHandlerExposer_Expecter) GetPlcTagHandler() *MockHandlerExposer_GetPlcTagHandler_Call {
	return &MockHandlerExposer_GetPlcTagHandler_Call{Call: _e.mock.On("GetPlcTagHandler")}
}

func (_c *MockHandlerExposer_GetPlcTagHandler_Call) Run(run func()) *MockHandlerExposer_GetPlcTagHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockHandlerExposer_GetPlcTagHandler_Call) Return(_a0 PlcTagHandler) *MockHandlerExposer_GetPlcTagHandler_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockHandlerExposer_GetPlcTagHandler_Call) RunAndReturn(run func() PlcTagHandler) *MockHandlerExposer_GetPlcTagHandler_Call {
	_c.Call.Return(run)
	return _c
}

// GetPlcValueHandler provides a mock function with given fields:
func (_m *MockHandlerExposer) GetPlcValueHandler() PlcValueHandler {
	ret := _m.Called()

	var r0 PlcValueHandler
	if rf, ok := ret.Get(0).(func() PlcValueHandler); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(PlcValueHandler)
		}
	}

	return r0
}

// MockHandlerExposer_GetPlcValueHandler_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPlcValueHandler'
type MockHandlerExposer_GetPlcValueHandler_Call struct {
	*mock.Call
}

// GetPlcValueHandler is a helper method to define mock.On call
func (_e *MockHandlerExposer_Expecter) GetPlcValueHandler() *MockHandlerExposer_GetPlcValueHandler_Call {
	return &MockHandlerExposer_GetPlcValueHandler_Call{Call: _e.mock.On("GetPlcValueHandler")}
}

func (_c *MockHandlerExposer_GetPlcValueHandler_Call) Run(run func()) *MockHandlerExposer_GetPlcValueHandler_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockHandlerExposer_GetPlcValueHandler_Call) Return(_a0 PlcValueHandler) *MockHandlerExposer_GetPlcValueHandler_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockHandlerExposer_GetPlcValueHandler_Call) RunAndReturn(run func() PlcValueHandler) *MockHandlerExposer_GetPlcValueHandler_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockHandlerExposer creates a new instance of MockHandlerExposer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockHandlerExposer(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockHandlerExposer {
	mock := &MockHandlerExposer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
