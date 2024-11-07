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

import (
	protobuf "github.com/growingio/growingio-sdk-go/internal/protobuf"
	mock "github.com/stretchr/testify/mock"
)

// MockBatch is an autogenerated mock type for the Batch type
type MockBatch struct {
	mock.Mock
}

type MockBatch_Expecter struct {
	mock *mock.Mock
}

func (_m *MockBatch) EXPECT() *MockBatch_Expecter {
	return &MockBatch_Expecter{mock: &_m.Mock}
}

// pop provides a mock function with given fields:
func (_m *MockBatch) pop() {
	_m.Called()
}

// MockBatch_pop_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'pop'
type MockBatch_pop_Call struct {
	*mock.Call
}

// pop is a helper method to define mock.On call
func (_e *MockBatch_Expecter) pop() *MockBatch_pop_Call {
	return &MockBatch_pop_Call{Call: _e.mock.On("pop")}
}

func (_c *MockBatch_pop_Call) Run(run func()) *MockBatch_pop_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockBatch_pop_Call) Return() *MockBatch_pop_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockBatch_pop_Call) RunAndReturn(run func()) *MockBatch_pop_Call {
	_c.Call.Return(run)
	return _c
}

// pushEvent provides a mock function with given fields: event
func (_m *MockBatch) pushEvent(event *protobuf.EventV3Dto) {
	_m.Called(event)
}

// MockBatch_pushEvent_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'pushEvent'
type MockBatch_pushEvent_Call struct {
	*mock.Call
}

// pushEvent is a helper method to define mock.On call
//   - event *protobuf.EventV3Dto
func (_e *MockBatch_Expecter) pushEvent(event interface{}) *MockBatch_pushEvent_Call {
	return &MockBatch_pushEvent_Call{Call: _e.mock.On("pushEvent", event)}
}

func (_c *MockBatch_pushEvent_Call) Run(run func(event *protobuf.EventV3Dto)) *MockBatch_pushEvent_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*protobuf.EventV3Dto))
	})
	return _c
}

func (_c *MockBatch_pushEvent_Call) Return() *MockBatch_pushEvent_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockBatch_pushEvent_Call) RunAndReturn(run func(*protobuf.EventV3Dto)) *MockBatch_pushEvent_Call {
	_c.Call.Return(run)
	return _c
}

// pushItem provides a mock function with given fields: item
func (_m *MockBatch) pushItem(item *protobuf.ItemDto) {
	_m.Called(item)
}

// MockBatch_pushItem_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'pushItem'
type MockBatch_pushItem_Call struct {
	*mock.Call
}

// pushItem is a helper method to define mock.On call
//   - item *protobuf.ItemDto
func (_e *MockBatch_Expecter) pushItem(item interface{}) *MockBatch_pushItem_Call {
	return &MockBatch_pushItem_Call{Call: _e.mock.On("pushItem", item)}
}

func (_c *MockBatch_pushItem_Call) Run(run func(item *protobuf.ItemDto)) *MockBatch_pushItem_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*protobuf.ItemDto))
	})
	return _c
}

func (_c *MockBatch_pushItem_Call) Return() *MockBatch_pushItem_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockBatch_pushItem_Call) RunAndReturn(run func(*protobuf.ItemDto)) *MockBatch_pushItem_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockBatch creates a new instance of MockBatch. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockBatch(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockBatch {
	mock := &MockBatch{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
