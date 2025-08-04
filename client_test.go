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
	"os"
	"path/filepath"
	"testing"
	"time"

	core "github.com/growingio/growingio-sdk-go/internal/core"
	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	// Reset global state
	core.InitializedSuccessfully = false

	// Test with valid config
	config := createValidConfig()
	client, err := NewClient(config)
	assert.NoError(t, err)
	assert.NotNil(t, client)
	assert.True(t, client.initialized)
	assert.Equal(t, config, client.config)

	// Test with invalid config
	invalidConfig := &Config{}
	client, err = NewClient(invalidConfig)
	assert.Error(t, err)
	assert.Nil(t, client)

	// Test with nil config
	client, err = NewClient(nil)
	assert.Error(t, err)
	assert.Nil(t, client)
}

func TestNewClientFromFile(t *testing.T) {
	// Create a temporary config file
	configContent := `
sdk:
  account_id: "file-test-account"
  data_source_id: "file-test-datasource"
  log_level: "debug"
http:
  server_host: "https://file-api.test.com"
  timeout: 45
batch:
  enable: false
  max_size: 200
  flush_after: 8
  routine_count: 6
  max_cache_size: 2000
`

	tmpDir := t.TempDir()
	configFile := filepath.Join(tmpDir, "test_config.yml")
	err := os.WriteFile(configFile, []byte(configContent), 0644)
	assert.NoError(t, err)

	// Test successful client creation from file
	client, err := NewClientFromFile(configFile)
	assert.NoError(t, err)
	assert.NotNil(t, client)
	assert.True(t, client.initialized)
	assert.Equal(t, "file-test-account", client.config.SdkConfig.AccountId)
	assert.Equal(t, "file-test-datasource", client.config.SdkConfig.DataSourceId)
	assert.Equal(t, DebugLogLevel, client.config.SdkConfig.LogLevel)

	// Test with non-existent file
	client, err = NewClientFromFile("non-existent.yml")
	assert.Error(t, err)
	assert.Nil(t, client)

	// Test with invalid YAML
	invalidConfigFile := filepath.Join(tmpDir, "invalid.yml")
	err = os.WriteFile(invalidConfigFile, []byte("invalid: yaml: [content"), 0644)
	assert.NoError(t, err)

	client, err = NewClientFromFile(invalidConfigFile)
	assert.Error(t, err)
	assert.Nil(t, client)
}

func TestClient_Initialize(t *testing.T) {
	// Reset global state
	core.InitializedSuccessfully = false

	config := createValidConfig()
	client := &Client{config: config}

	// Test successful initialization
	err := client.initialize()
	assert.NoError(t, err)
	assert.True(t, client.initialized)
	assert.True(t, core.InitializedSuccessfully)

	// Test duplicate initialization
	err = client.initialize()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "already initialized")
}

func TestClient_InitializeWithBatch(t *testing.T) {
	// Reset global state
	core.InitializedSuccessfully = false
	core.BatchEnable = false

	config := createValidConfig()
	config.BatchConfig.Enable = true
	config.BatchConfig.MaxSize = 10
	config.BatchConfig.FlushAfter = 2
	config.BatchConfig.RoutineCount = 2
	config.BatchConfig.MaxCacheSize = 100

	client := &Client{config: config}

	// Test initialization with batch enabled
	err := client.initialize()
	assert.NoError(t, err)
	assert.True(t, client.initialized)
	assert.True(t, core.BatchEnable)
	assert.Equal(t, 10, core.MaxSize)
	assert.Equal(t, 2, core.FlushAfter)
	assert.Equal(t, 2, core.RoutineCount)
	assert.Equal(t, 100, core.MaxCacheSize)
}

func TestClient_TrackCustomEvent(t *testing.T) {
	// Reset global state
	core.InitializedSuccessfully = false

	config := createValidConfig()
	client, err := NewClient(config)
	assert.NoError(t, err)

	// Test successful custom event tracking
	customEvent := NewCustomEvent("client-test-event").
		WithAnonymousId("client-anonymous-123").
		WithEventTime(time.Now().UnixMilli()).
		WithAttributes(map[string]interface{}{
			"test_key": "test_value",
			"number":   42,
		})

	err = client.TrackCustomEvent(customEvent)
	assert.NoError(t, err)

	// Test with nil builder
	err = client.TrackCustomEvent(nil)
	assert.Error(t, err)

	// Test with invalid builder (empty event name)
	invalidEvent := NewCustomEvent("")
	err = client.TrackCustomEvent(invalidEvent)
	assert.Error(t, err)

	// Test with uninitialized client
	uninitializedClient := &Client{config: config, initialized: false}
	err = uninitializedClient.TrackCustomEvent(customEvent)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not initialized")
}

func TestClient_TrackUser(t *testing.T) {
	// Reset global state
	core.InitializedSuccessfully = false

	config := createValidConfig()
	client, err := NewClient(config)
	assert.NoError(t, err)

	// Test successful user tracking
	userEvent := NewUser("client-user-123").
		WithEventTime(time.Now().UnixMilli()).
		WithAnonymousId("client-anonymous-456").
		WithLoginUserKey("client-user-key-789").
		WithAttributes(map[string]interface{}{
			"name":  "Client Test User",
			"email": "client@test.com",
		})

	err = client.TrackUser(userEvent)
	assert.NoError(t, err)

	// Test with nil builder
	err = client.TrackUser(nil)
	assert.Error(t, err)

	// Test with invalid builder (empty login user id)
	invalidUser := NewUser("")
	err = client.TrackUser(invalidUser)
	assert.Error(t, err)

	// Test with uninitialized client
	uninitializedClient := &Client{config: config, initialized: false}
	err = uninitializedClient.TrackUser(userEvent)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not initialized")
}

func TestClient_SubmitItem(t *testing.T) {
	// Reset global state
	core.InitializedSuccessfully = false

	config := createValidConfig()
	client, err := NewClient(config)
	assert.NoError(t, err)

	// Test successful item submission
	item := NewItem("client-item-123", "client-product").
		WithAttributes(map[string]interface{}{
			"name":        "Client Test Item",
			"price":       29.99,
			"category":    "test",
			"in_stock":    true,
			"tags":        []string{"client", "test"},
		})

	err = client.SubmitItem(item)
	assert.NoError(t, err)

	// Test with nil builder
	err = client.SubmitItem(nil)
	assert.Error(t, err)

	// Test with invalid builder (empty item id)
	invalidItem := NewItem("", "test-key")
	err = client.SubmitItem(invalidItem)
	assert.Error(t, err)

	// Test with invalid builder (empty item key)
	invalidItem2 := NewItem("test-id", "")
	err = client.SubmitItem(invalidItem2)
	assert.Error(t, err)

	// Test with uninitialized client
	uninitializedClient := &Client{config: config, initialized: false}
	err = uninitializedClient.SubmitItem(item)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not initialized")
}

func TestClient_ConcurrentAccess(t *testing.T) {
	// Reset global state
	core.InitializedSuccessfully = false

	config := createValidConfig()
	client, err := NewClient(config)
	assert.NoError(t, err)

	// Test concurrent access to client methods
	done := make(chan bool, 3)

	// Goroutine 1: Track custom events
	go func() {
		defer func() { done <- true }()
		for i := 0; i < 10; i++ {
			event := NewCustomEvent("concurrent-event").
				WithAnonymousId("anonymous-concurrent")
			client.TrackCustomEvent(event)
		}
	}()

	// Goroutine 2: Track users
	go func() {
		defer func() { done <- true }()
		for i := 0; i < 10; i++ {
			user := NewUser("concurrent-user")
			client.TrackUser(user)
		}
	}()

	// Goroutine 3: Submit items
	go func() {
		defer func() { done <- true }()
		for i := 0; i < 10; i++ {
			item := NewItem("concurrent-item", "concurrent-key")
			client.SubmitItem(item)
		}
	}()

	// Wait for all goroutines to complete
	for i := 0; i < 3; i++ {
		<-done
	}

	// Verify client is still in good state
	assert.True(t, client.initialized)
}

func TestClient_MultipleInstances(t *testing.T) {
	// Reset global state
	core.InitializedSuccessfully = false

	// Create multiple client instances
	config1 := createValidConfig()
	config1.SdkConfig.AccountId = "account-1"

	config2 := createValidConfig()
	config2.SdkConfig.AccountId = "account-2"

	client1, err := NewClient(config1)
	assert.NoError(t, err)
	assert.Equal(t, "account-1", client1.config.SdkConfig.AccountId)

	// Note: Due to global state in core package, second client will override first
	// This is a limitation of the current architecture
	client2, err := NewClient(config2)
	assert.NoError(t, err)
	assert.Equal(t, "account-2", client2.config.SdkConfig.AccountId)

	// Both clients should be initialized
	assert.True(t, client1.initialized)
	assert.True(t, client2.initialized)
}