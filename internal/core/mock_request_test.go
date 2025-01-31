//
// @license
// Copyright (C) 2024 Beijing Yishu Technology Co., Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Code generated by mockery v2.46.3. DO NOT EDIT.

package core

import mock "github.com/stretchr/testify/mock"

// MockRequest is an autogenerated mock type for the Request type
type MockRequest struct {
	mock.Mock
}

type MockRequest_Expecter struct {
	mock *mock.Mock
}

func (_m *MockRequest) EXPECT() *MockRequest_Expecter {
	return &MockRequest_Expecter{mock: &_m.Mock}
}

// prepare provides a mock function with given fields:
func (_m *MockRequest) prepare() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for prepare")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRequest_prepare_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'prepare'
type MockRequest_prepare_Call struct {
	*mock.Call
}

// prepare is a helper method to define mock.On call
func (_e *MockRequest_Expecter) prepare() *MockRequest_prepare_Call {
	return &MockRequest_prepare_Call{Call: _e.mock.On("prepare")}
}

func (_c *MockRequest_prepare_Call) Run(run func()) *MockRequest_prepare_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockRequest_prepare_Call) Return(_a0 error) *MockRequest_prepare_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRequest_prepare_Call) RunAndReturn(run func() error) *MockRequest_prepare_Call {
	_c.Call.Return(run)
	return _c
}

// send provides a mock function with given fields:
func (_m *MockRequest) send() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for send")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockRequest_send_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'send'
type MockRequest_send_Call struct {
	*mock.Call
}

// send is a helper method to define mock.On call
func (_e *MockRequest_Expecter) send() *MockRequest_send_Call {
	return &MockRequest_send_Call{Call: _e.mock.On("send")}
}

func (_c *MockRequest_send_Call) Run(run func()) *MockRequest_send_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockRequest_send_Call) Return(_a0 error) *MockRequest_send_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockRequest_send_Call) RunAndReturn(run func() error) *MockRequest_send_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockRequest creates a new instance of MockRequest. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockRequest(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockRequest {
	mock := &MockRequest{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
