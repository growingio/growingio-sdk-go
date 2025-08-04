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

package analytics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCustomEvent(t *testing.T) {
	eventName := "test-event"
	builder := NewCustomEvent(eventName)
	
	assert.NotNil(t, builder)
	assert.Equal(t, eventName, builder.EventName)
}

func TestNewUser(t *testing.T) {
	loginUserId := "test-user-id"
	builder := NewUser(loginUserId)
	
	assert.NotNil(t, builder)
	assert.Equal(t, loginUserId, builder.LoginUserId)
}

func TestNewItem(t *testing.T) {
	itemId := "test-item-id"
	itemKey := "test-item-key"
	builder := NewItem(itemId, itemKey)
	
	assert.NotNil(t, builder)
	assert.Equal(t, itemId, builder.ItemId)
	assert.Equal(t, itemKey, builder.ItemKey)
}

func TestCustomEventBuilder_WithMethods(t *testing.T) {
	builder := NewCustomEvent("test-event")
	
	// Test WithEventTime
	eventTime := int64(1234567890)
	result := builder.WithEventTime(eventTime)
	assert.Equal(t, builder, result) // Should return same builder
	assert.Equal(t, eventTime, builder.EventTime)
	
	// Test WithAnonymousId
	anonymousId := "test-anonymous-id"
	result = builder.WithAnonymousId(anonymousId)
	assert.Equal(t, builder, result)
	assert.Equal(t, anonymousId, builder.AnonymousId)
	
	// Test WithLoginUserId
	loginUserId := "test-login-user-id"
	result = builder.WithLoginUserId(loginUserId)
	assert.Equal(t, builder, result)
	assert.Equal(t, loginUserId, builder.LoginUserId)
	
	// Test WithLoginUserKey
	loginUserKey := "test-login-user-key"
	result = builder.WithLoginUserKey(loginUserKey)
	assert.Equal(t, builder, result)
	assert.Equal(t, loginUserKey, builder.LoginUserKey)
	
	// Test WithAttributes
	attributes := map[string]interface{}{
		"key1": "value1",
		"key2": 123,
	}
	result = builder.WithAttributes(attributes)
	assert.Equal(t, builder, result)
	assert.Equal(t, attributes, builder.Attributes)
}

func TestUserEventBuilder_WithMethods(t *testing.T) {
	builder := NewUser("test-user-id")
	
	// Test WithEventTime
	eventTime := int64(1234567890)
	result := builder.WithEventTime(eventTime)
	assert.Equal(t, builder, result)
	assert.Equal(t, eventTime, builder.EventTime)
	
	// Test WithAnonymousId
	anonymousId := "test-anonymous-id"
	result = builder.WithAnonymousId(anonymousId)
	assert.Equal(t, builder, result)
	assert.Equal(t, anonymousId, builder.AnonymousId)
	
	// Test WithLoginUserKey
	loginUserKey := "test-login-user-key"
	result = builder.WithLoginUserKey(loginUserKey)
	assert.Equal(t, builder, result)
	assert.Equal(t, loginUserKey, builder.LoginUserKey)
	
	// Test WithAttributes
	attributes := map[string]interface{}{
		"key1": "value1",
		"key2": 123,
	}
	result = builder.WithAttributes(attributes)
	assert.Equal(t, builder, result)
	assert.Equal(t, attributes, builder.Attributes)
}

func TestItemBuilder_WithMethods(t *testing.T) {
	builder := NewItem("test-item-id", "test-item-key")
	
	// Test WithAttributes
	attributes := map[string]interface{}{
		"key1": "value1",
		"key2": 123,
	}
	result := builder.WithAttributes(attributes)
	assert.Equal(t, builder, result)
	assert.Equal(t, attributes, builder.Attributes)
}

func TestCommonEventFields_HelperMethods(t *testing.T) {
	fields := &CommonEventFields{}
	
	// Test setEventTime
	eventTime := int64(1234567890)
	fields.setEventTime(eventTime)
	assert.Equal(t, eventTime, fields.EventTime)
	
	// Test setAnonymousId
	anonymousId := "test-anonymous-id"
	fields.setAnonymousId(anonymousId)
	assert.Equal(t, anonymousId, fields.AnonymousId)
	
	// Test setLoginUserId
	loginUserId := "test-login-user-id"
	fields.setLoginUserId(loginUserId)
	assert.Equal(t, loginUserId, fields.LoginUserId)
	
	// Test setLoginUserKey
	loginUserKey := "test-login-user-key"
	fields.setLoginUserKey(loginUserKey)
	assert.Equal(t, loginUserKey, fields.LoginUserKey)
	
	// Test setAttributes
	attributes := map[string]interface{}{
		"key1": "value1",
		"key2": 123,
	}
	fields.setAttributes(attributes)
	assert.Equal(t, attributes, fields.Attributes)
}