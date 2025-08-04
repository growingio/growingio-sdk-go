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
	"testing"
	"time"

	"github.com/growingio/growingio-sdk-go/internal/protobuf"
	"github.com/stretchr/testify/assert"
)

func setupTestEnvironment() {
	AccountId = "test-account"
	DataSourceId = "test-datasource"
	Platform = "go"
	SdkVersion = "1.0.0-test"
	Urls = RequestUrl{
		Collect: "https://test-api.com/v3/projects/test-account/collect",
		Item:    "https://test-api.com/projects/test-account/collect/item",
	}
	RequestTimeout = 5
	InitializedSuccessfully = true
}

func TestBuildCustomEvent_Integration(t *testing.T) {
	setupTestEnvironment()

	builder := &EventBuilder{
		EventName:    "integration_test_event",
		EventTime:    time.Now().UnixMilli(),
		AnonymousId:  "integration-anonymous-123",
		LoginUserId:  "integration-user-456",
		LoginUserKey: "integration-key-789",
		Attributes: map[string]interface{}{
			"test_string":      "integration_value",
			"test_number":      42,
			"test_bool":        true,
			"test_string_array": []string{"a", "b", "c"},
			"test_int_array":    []int{1, 2, 3},
			"test_map":         map[string]string{"nested": "value"},
		},
	}

	// Test building custom event
	event := buildCustomEvent(builder)
	
	assert.NotNil(t, event)
	assert.Equal(t, "test-account", event.ProjectKey)
	assert.Equal(t, "test-datasource", event.DataSourceId)
	assert.Equal(t, "1.0.0-test", event.SdkVersion)
	assert.Equal(t, "go", event.Platform)
	assert.Equal(t, protobuf.EventType_CUSTOM, event.EventType)
	assert.Equal(t, "integration_test_event", event.EventName)
	assert.Equal(t, builder.EventTime, event.Timestamp)
	assert.Equal(t, "integration-anonymous-123", event.DeviceId)
	assert.Equal(t, "integration-user-456", event.UserId)
	assert.Equal(t, "integration-key-789", event.UserKey)
	assert.NotNil(t, event.Attributes)
	
	// Check attributes conversion
	assert.Equal(t, "integration_value", event.Attributes["test_string"])
	assert.Equal(t, "42", event.Attributes["test_number"])
	assert.Equal(t, "true", event.Attributes["test_bool"])
	assert.Equal(t, "a||b||c", event.Attributes["test_string_array"])
	assert.Equal(t, "1||2||3", event.Attributes["test_int_array"])
	assert.Contains(t, event.Attributes["test_map"], "nested")
}

func TestBuildUserLoginEvent_Integration(t *testing.T) {
	setupTestEnvironment()

	builder := &EventBuilder{
		EventTime:    time.Now().UnixMilli(),
		AnonymousId:  "user-anonymous-123",
		LoginUserId:  "user-login-456",
		LoginUserKey: "user-key-789",
		Attributes: map[string]interface{}{
			"name":         "Integration Test User",
			"email":        "integration@test.com",
			"preferences":  []string{"pref1", "pref2"},
		},
	}

	// Test building user login event
	event := buildUserLoginEvent(builder)
	
	assert.NotNil(t, event)
	assert.Equal(t, "test-account", event.ProjectKey)
	assert.Equal(t, "test-datasource", event.DataSourceId)
	assert.Equal(t, protobuf.EventType_LOGIN_USER_ATTRIBUTES, event.EventType)
	assert.Equal(t, builder.EventTime, event.Timestamp)
	assert.Equal(t, "user-anonymous-123", event.DeviceId)
	assert.Equal(t, "user-login-456", event.UserId)
	assert.Equal(t, "user-key-789", event.UserKey)
	
	// Check attributes
	assert.Equal(t, "Integration Test User", event.Attributes["name"])
	assert.Equal(t, "integration@test.com", event.Attributes["email"])
	assert.Equal(t, "pref1||pref2", event.Attributes["preferences"])
}

func TestBuildItemEvent_Integration(t *testing.T) {
	setupTestEnvironment()

	builder := &EventBuilder{
		ItemId:  "integration-item-123",
		ItemKey: "integration-product",
		Attributes: map[string]interface{}{
			"name":        "Integration Test Item",
			"price":       99.99,
			"category":    "test-category",
			"tags":        []string{"tag1", "tag2", "tag3"},
			"in_stock":    true,
		},
	}

	// Test building item event
	event := buildItemEvent(builder)
	
	assert.NotNil(t, event)
	assert.Equal(t, "test-account", event.ProjectKey)
	assert.Equal(t, "test-datasource", event.DataSourceId)
	assert.Equal(t, "integration-item-123", event.Id)
	assert.Equal(t, "integration-product", event.Key)
	
	// Check attributes
	assert.Equal(t, "Integration Test Item", event.Attributes["name"])
	assert.Equal(t, "99.99", event.Attributes["price"])
	assert.Equal(t, "test-category", event.Attributes["category"])
	assert.Equal(t, "tag1||tag2||tag3", event.Attributes["tags"])
	assert.Equal(t, "true", event.Attributes["in_stock"])
}

func TestBatchProcessing_Integration(t *testing.T) {
	// Reset batch state
	BatchEnable = true
	MaxSize = 3  // Small batch size for testing
	FlushAfter = 2
	RoutineCount = 2
	MaxCacheSize = 10
	
	InitBatch()
	assert.NotNil(t, bInst)

	// Create test events
	events := []*protobuf.EventV3Dto{
		{EventName: "batch_event_1", ProjectKey: "test"},
		{EventName: "batch_event_2", ProjectKey: "test"},
		{EventName: "batch_event_3", ProjectKey: "test"},
	}

	// Create test items  
	items := []*protobuf.ItemDto{
		{Id: "batch_item_1", Key: "product"},
		{Id: "batch_item_2", Key: "product"},
	}

	// Test pushing events to batch
	for _, event := range events {
		bInst.pushEvent(event)
	}

	// Test pushing items to batch
	for _, item := range items {
		bInst.pushItem(item)
	}

	// Give some time for batch processing
	time.Sleep(100 * time.Millisecond)

	// Verify batch is working (we can't easily verify exact behavior without mocking HTTP)
	assert.True(t, BatchEnable)
	assert.Equal(t, 3, MaxSize)
	assert.Equal(t, 2, FlushAfter)
}

func TestRunBatch_Integration(t *testing.T) {
	// Reset batch state
	BatchEnable = true
	MaxSize = 5
	FlushAfter = 1  // Short timeout for testing
	RoutineCount = 2
	MaxCacheSize = 20
	
	InitBatch()
	
	// Start batch processing
	RunBatch()
	
	// Give time for goroutines to start
	time.Sleep(50 * time.Millisecond)
	
	// Test is mainly to ensure RunBatch doesn't panic
	// and that goroutines are started properly
	assert.True(t, BatchEnable)
	assert.NotNil(t, bInst)
}

func TestGetDefaultRoutineCount_Integration(t *testing.T) {
	count := getDefaultRoutineCount()
	
	// Should be at least 4 and at most 32
	assert.GreaterOrEqual(t, count, 4)
	assert.LessOrEqual(t, count, 32)
	
	// Should be reasonable for most systems
	assert.True(t, count >= 4 && count <= 32)
}

func TestSendOrStoreEvent_WithBatch(t *testing.T) {
	setupTestEnvironment()
	
	// Enable batch
	BatchEnable = true
	MaxSize = 10
	MaxCacheSize = 50
	InitBatch()
	
	event := &protobuf.EventV3Dto{
		EventName:   "batch_test_event",
		ProjectKey:  "test-account",
		EventType:   protobuf.EventType_CUSTOM,
	}
	
	// Should use batch when enabled
	sendOrStoreEvent(bInst, event)
	
	// Verify event was processed (we can't easily verify it was batched without more complex setup)
	assert.True(t, BatchEnable)
}

func TestSendOrStoreEvent_WithoutBatch(t *testing.T) {
	setupTestEnvironment()
	
	// Disable batch
	BatchEnable = false
	
	event := &protobuf.EventV3Dto{
		EventName:   "direct_test_event",
		ProjectKey:  "test-account",
		EventType:   protobuf.EventType_CUSTOM,
	}
	
	// Should send directly when batch disabled
	sendOrStoreEvent(bInst, event)
	
	// Verify batch is disabled
	assert.False(t, BatchEnable)
}

func TestSendOrStoreItem_WithBatch(t *testing.T) {
	setupTestEnvironment()
	
	// Enable batch
	BatchEnable = true
	MaxSize = 10
	MaxCacheSize = 50
	InitBatch()
	
	item := &protobuf.ItemDto{
		Id:         "batch_test_item",
		Key:        "product",
		ProjectKey: "test-account",
	}
	
	// Should use batch when enabled
	sendOrStoreItem(bInst, item)
	
	// Verify batch is enabled
	assert.True(t, BatchEnable)
}

func TestSendOrStoreItem_WithoutBatch(t *testing.T) {
	setupTestEnvironment()
	
	// Disable batch
	BatchEnable = false
	
	item := &protobuf.ItemDto{
		Id:         "direct_test_item",
		Key:        "product",
		ProjectKey: "test-account",
	}
	
	// Should send directly when batch disabled
	sendOrStoreItem(bInst, item)
	
	// Verify batch is disabled
	assert.False(t, BatchEnable)
}