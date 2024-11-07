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

package core

import (
	"reflect"
	"testing"
	"time"

	"github.com/growingio/growingio-sdk-go/internal/protobuf"
	"github.com/stretchr/testify/assert"
)

func TestSendOrStoreEvent(t *testing.T) {
	event := &protobuf.EventV3Dto{}
	mockBatch := NewMockBatch(t)

	BatchEnable = true
	mockBatch.EXPECT().pushEvent(event)
	sendOrStoreEvent(event)
	mockBatch.AssertExpectations(t)
}

func TestSendOrStoreItem(t *testing.T) {
	event := &protobuf.ItemDto{}
	mockBatch := NewMockBatch(t)

	BatchEnable = true
	mockBatch.EXPECT().pushItem(event)
	sendOrStoreItem(event)
	mockBatch.AssertExpectations(t)
}

func TestGetEventTimeWhenSet(t *testing.T) {
	eventTime := int64(1234567890)
	e := &EventBuilder{EventTime: eventTime}
	result := e.getEventTime()
	assert.Equal(t, eventTime, result)
}

func TestGetEventTimeWhenNoSet(t *testing.T) {
	e := &EventBuilder{}
	before := time.Now().UnixMilli()
	result := e.getEventTime()
	after := time.Now().UnixMilli()
	assert.GreaterOrEqual(t, result, before)
	assert.LessOrEqual(t, result, after)
}

func TestGetAnonymousIdWhenSet(t *testing.T) {
	anonymousId := "F06A848D-C1DD-4999-B646-82EB7578BBBB"
	e := &EventBuilder{AnonymousId: anonymousId}
	result := e.getAnonymousId()
	assert.Equal(t, anonymousId, result)
}

func TestGetAnonymousIdWhenNoSet(t *testing.T) {
	e := &EventBuilder{}
	result := e.getAnonymousId()
	assert.Equal(t, "", result)
}

func TestGetLoginUserIdWhenSet(t *testing.T) {
	loginUserId := "userId123"
	e := &EventBuilder{LoginUserId: loginUserId}
	result := e.getLoginUserId()
	assert.Equal(t, loginUserId, result)
}

func TestGetLoginUserIdWhenNoSet(t *testing.T) {
	e := &EventBuilder{}
	result := e.getLoginUserId()
	assert.Equal(t, "", result)
}

func TestGetLoginUserKeyWhenSet(t *testing.T) {
	loginUserKey := "userKey123"
	e := &EventBuilder{LoginUserKey: loginUserKey}
	result := e.getLoginUserKey()
	assert.Equal(t, loginUserKey, result)
}

func TestGetLoginUserKeyWhenNoSet(t *testing.T) {
	e := &EventBuilder{}
	result := e.getLoginUserKey()
	assert.Equal(t, "", result)
}

func TestGetItemIdWhenSet(t *testing.T) {
	itemId := "item123"
	e := &EventBuilder{ItemId: itemId}
	result := e.getItemId()
	assert.Equal(t, itemId, result)
}

func TestGetItemIdWhenNoSet(t *testing.T) {
	e := &EventBuilder{}
	result := e.getItemId()
	assert.Equal(t, "", result)
}

func TestGetItemKeyWhenSet(t *testing.T) {
	itemKey := "itemKey123"
	e := &EventBuilder{ItemKey: itemKey}
	result := e.getItemKey()
	assert.Equal(t, itemKey, result)
}

func TestGetItemKeyWhenNoSet(t *testing.T) {
	e := &EventBuilder{}
	result := e.getItemKey()
	assert.Equal(t, "", result)
}

func TestGetAttributes(t *testing.T) {
	tests := []struct {
		name       string
		attributes map[string]interface{}
		expected   map[string]string
	}{
		{
			name: "Single string value",
			attributes: map[string]interface{}{
				"key1": "value1",
			},
			expected: map[string]string{
				"key1": "value1",
			},
		},
		{
			name: "String slice",
			attributes: map[string]interface{}{
				"key1": []string{"value1", "value2"},
			},
			expected: map[string]string{
				"key1": "value1||value2",
			},
		},
		{
			name: "Bool slice",
			attributes: map[string]interface{}{
				"key1": []bool{true, false},
			},
			expected: map[string]string{
				"key1": "true||false",
			},
		},
		{
			name: "Int slice",
			attributes: map[string]interface{}{
				"key1": []int{1, 2, 3},
			},
			expected: map[string]string{
				"key1": "1||2||3",
			},
		},
		{
			name: "Int32 slice",
			attributes: map[string]interface{}{
				"key1": []int32{10, 20},
			},
			expected: map[string]string{
				"key1": "10||20",
			},
		},
		{
			name: "Float64 slice",
			attributes: map[string]interface{}{
				"key1": []float64{1.1, 2.2, 3.3},
			},
			expected: map[string]string{
				"key1": "1.1||2.2||3.3",
			},
		},
		{
			name: "Map of strings",
			attributes: map[string]interface{}{
				"key1": map[string]string{"subkey1": "subval1", "subkey2": "subval2"},
			},
			expected: map[string]string{
				"key1": `{"subkey1":"subval1","subkey2":"subval2"}`,
			},
		},
		{
			name: "Interface slice",
			attributes: map[string]interface{}{
				"key1": []interface{}{"val1", 123, true},
			},
			expected: map[string]string{
				"key1": "val1||123||true",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &EventBuilder{Attributes: tt.attributes}
			got := e.getAttributes()

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("getAttributes() = %v, expected %v", got, tt.expected)
			}
		})
	}
}
