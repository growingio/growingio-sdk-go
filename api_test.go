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

func resetGlobalState() {
	defaultClient = nil
	core.InitializedSuccessfully = false
	core.BatchEnable = false
}

func createValidConfig() *Config {
	return &Config{
		SdkConfig: SdkConfig{
			AccountId:    "test-account-id",
			DataSourceId: "test-data-source-id",
			LogLevel:     InfoLogLevel,
		},
		HttpConfig: HttpConfig{
			ServerHost:     "https://api.test.com",
			RequestTimeout: 30,
		},
		BatchConfig: BatchConfig{
			Enable:       false, // Disable batch to avoid goroutines in tests
			MaxSize:      100,
			FlushAfter:   5,
			RoutineCount: 4,
			MaxCacheSize: 1000,
		},
	}
}

func TestInitAnalytics(t *testing.T) {
	resetGlobalState()
	defer resetGlobalState()

	config := createValidConfig()
	
	// Test successful initialization
	err := InitAnalytics(config)
	assert.NoError(t, err)
	assert.NotNil(t, defaultClient)
	assert.True(t, core.InitializedSuccessfully)

	// Test duplicate initialization prevention
	err = InitAnalytics(config)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "already initialized")
}

func TestInitAnalytics_InvalidConfig(t *testing.T) {
	resetGlobalState()
	defer resetGlobalState()

	tests := []struct {
		name   string
		config *Config
	}{
		{
			name:   "nil config",
			config: nil,
		},
		{
			name: "empty AccountId",
			config: &Config{
				SdkConfig: SdkConfig{
					AccountId:    "",
					DataSourceId: "test-data-source-id",
				},
				HttpConfig: HttpConfig{
					ServerHost: "https://api.test.com",
				},
			},
		},
		{
			name: "empty DataSourceId",
			config: &Config{
				SdkConfig: SdkConfig{
					AccountId:    "test-account-id",
					DataSourceId: "",
				},
				HttpConfig: HttpConfig{
					ServerHost: "https://api.test.com",
				},
			},
		},
		{
			name: "empty ServerHost",
			config: &Config{
				SdkConfig: SdkConfig{
					AccountId:    "test-account-id",
					DataSourceId: "test-data-source-id",
				},
				HttpConfig: HttpConfig{
					ServerHost: "",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resetGlobalState()
			err := InitAnalytics(tt.config)
			assert.Error(t, err)
			assert.Nil(t, defaultClient)
		})
	}
}

func TestInitAnalyticsByConfigFile(t *testing.T) {
	resetGlobalState()
	defer resetGlobalState()

	// Create a temporary config file
	configContent := `
sdk:
  account_id: "test-account-id"
  data_source_id: "test-data-source-id"
  log_level: "info"
http:
  server_host: "https://api.test.com"
  timeout: 30
batch:
  enable: false
  max_size: 100
  flush_after: 5
  routine_count: 4
  max_cache_size: 1000
`

	tmpDir := t.TempDir()
	configFile := filepath.Join(tmpDir, "config.yml")
	err := os.WriteFile(configFile, []byte(configContent), 0644)
	assert.NoError(t, err)

	// Test successful initialization from file
	err = InitAnalyticsByConfigFile(configFile)
	assert.NoError(t, err)
	assert.NotNil(t, defaultClient)

	// Test duplicate initialization prevention
	err = InitAnalyticsByConfigFile(configFile)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "already initialized")
}

func TestInitAnalyticsByConfigFile_InvalidFile(t *testing.T) {
	resetGlobalState()
	defer resetGlobalState()

	// Test non-existent file
	err := InitAnalyticsByConfigFile("non-existent-file.yml")
	assert.Error(t, err)
	assert.Nil(t, defaultClient)

	// Test invalid YAML
	tmpDir := t.TempDir()
	invalidConfigFile := filepath.Join(tmpDir, "invalid.yml")
	err = os.WriteFile(invalidConfigFile, []byte("invalid: yaml: content: ["), 0644)
	assert.NoError(t, err)

	err = InitAnalyticsByConfigFile(invalidConfigFile)
	assert.Error(t, err)
	assert.Nil(t, defaultClient)
}

func TestTrackCustomEvent(t *testing.T) {
	resetGlobalState()
	defer resetGlobalState()

	// Test without initialization
	customEvent := NewCustomEvent("test-event").
		WithAnonymousId("test-anonymous-id")
	
	err := TrackCustomEvent(customEvent)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not been initialized")

	// Initialize SDK
	config := createValidConfig()
	err = InitAnalytics(config)
	assert.NoError(t, err)

	// Test successful event tracking
	err = TrackCustomEvent(customEvent)
	assert.NoError(t, err)

	// Test with nil builder
	err = TrackCustomEvent(nil)
	assert.Error(t, err)

	// Test with invalid builder
	invalidEvent := NewCustomEvent("")
	err = TrackCustomEvent(invalidEvent)
	assert.Error(t, err)
}

func TestTrackUser(t *testing.T) {
	resetGlobalState()
	defer resetGlobalState()

	// Test without initialization
	userEvent := NewUser("test-user-id")
	
	err := TrackUser(userEvent)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not been initialized")

	// Initialize SDK
	config := createValidConfig()
	err = InitAnalytics(config)
	assert.NoError(t, err)

	// Test successful user tracking
	err = TrackUser(userEvent)
	assert.NoError(t, err)

	// Test with nil builder
	err = TrackUser(nil)
	assert.Error(t, err)

	// Test with invalid builder
	invalidUser := NewUser("")
	err = TrackUser(invalidUser)
	assert.Error(t, err)
}

func TestSubmitItem(t *testing.T) {
	resetGlobalState()
	defer resetGlobalState()

	// Test without initialization
	item := NewItem("test-item-id", "test-item-key")
	
	err := SubmitItem(item)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "not been initialized")

	// Initialize SDK
	config := createValidConfig()
	err = InitAnalytics(config)
	assert.NoError(t, err)

	// Test successful item submission
	err = SubmitItem(item)
	assert.NoError(t, err)

	// Test with nil builder
	err = SubmitItem(nil)
	assert.Error(t, err)

	// Test with invalid builder
	invalidItem := NewItem("", "test-key")
	err = SubmitItem(invalidItem)
	assert.Error(t, err)
}

func TestComplexEventTracking(t *testing.T) {
	resetGlobalState()
	defer resetGlobalState()

	// Initialize SDK
	config := createValidConfig()
	err := InitAnalytics(config)
	assert.NoError(t, err)

	// Test complex custom event with all fields
	customEvent := NewCustomEvent("purchase").
		WithEventTime(time.Now().UnixMilli()).
		WithAnonymousId("anonymous-123").
		WithLoginUserId("user-456").
		WithLoginUserKey("user-key-789").
		WithAttributes(map[string]interface{}{
			"product_id":    "prod-123",
			"price":         99.99,
			"category":      "electronics",
			"tags":          []string{"new", "popular", "sale"},
			"is_premium":    true,
			"metadata":      map[string]string{"source": "mobile", "campaign": "holiday"},
		})

	err = TrackCustomEvent(customEvent)
	assert.NoError(t, err)

	// Test complex user event
	userEvent := NewUser("user-789").
		WithEventTime(time.Now().UnixMilli()).
		WithAnonymousId("anonymous-456").
		WithLoginUserKey("user-key-123").
		WithAttributes(map[string]interface{}{
			"name":         "John Doe",
			"email":        "john@example.com",
			"age":          30,
			"preferences":  []string{"tech", "gaming"},
			"is_verified":  true,
		})

	err = TrackUser(userEvent)
	assert.NoError(t, err)

	// Test complex item
	item := NewItem("item-123", "product").
		WithAttributes(map[string]interface{}{
			"name":        "Smartphone",
			"brand":       "TechCorp",
			"price":       599.99,
			"in_stock":    true,
			"categories":  []string{"electronics", "mobile", "devices"},
			"specs":       map[string]string{"memory": "8GB", "storage": "256GB"},
		})

	err = SubmitItem(item)
	assert.NoError(t, err)
}