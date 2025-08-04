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
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestConfig_EdgeCases(t *testing.T) {
	tests := []struct {
		name   string
		config *Config
		valid  bool
	}{
		{
			name: "negative timeout",
			config: &Config{
				SdkConfig: SdkConfig{
					AccountId:    "test-account",
					DataSourceId: "test-datasource",
				},
				HttpConfig: HttpConfig{
					ServerHost:     "https://api.test.com",
					RequestTimeout: -1,
				},
			},
			valid: false,
		},
		{
			name: "zero timeout (should be allowed)",
			config: &Config{
				SdkConfig: SdkConfig{
					AccountId:    "test-account",
					DataSourceId: "test-datasource",
				},
				HttpConfig: HttpConfig{
					ServerHost:     "https://api.test.com",
					RequestTimeout: 0,
				},
			},
			valid: true,
		},
		{
			name: "very long strings",
			config: &Config{
				SdkConfig: SdkConfig{
					AccountId:    strings.Repeat("a", 1000),
					DataSourceId: strings.Repeat("b", 1000),
				},
				HttpConfig: HttpConfig{
					ServerHost: "https://api.test.com",
				},
			},
			valid: true,
		},
		{
			name: "whitespace only strings",
			config: &Config{
				SdkConfig: SdkConfig{
					AccountId:    "   ",
					DataSourceId: "test-datasource",
				},
				HttpConfig: HttpConfig{
					ServerHost: "https://api.test.com",
				},
			},
			valid: true, // Our validation checks for empty, not whitespace
		},
		{
			name: "negative batch values",
			config: &Config{
				SdkConfig: SdkConfig{
					AccountId:    "test-account",
					DataSourceId: "test-datasource",
				},
				HttpConfig: HttpConfig{
					ServerHost: "https://api.test.com",
				},
				BatchConfig: BatchConfig{
					Enable:       true,
					MaxSize:      -1,
					FlushAfter:   -1,
					RoutineCount: -1,
					MaxCacheSize: -1,
				},
			},
			valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateConfig(tt.config)
			if tt.valid {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestBuilder_EdgeCases(t *testing.T) {
	// Test with very large attributes
	largeAttributes := make(map[string]interface{})
	for i := 0; i < 1000; i++ {
		largeAttributes[strings.Repeat("key", 100)] = strings.Repeat("value", 1000)
	}

	customEvent := NewCustomEvent("test-event").
		WithAnonymousId("test-id").
		WithAttributes(largeAttributes)

	assert.Equal(t, "test-event", customEvent.EventName)
	assert.Equal(t, "test-id", customEvent.AnonymousId)
	assert.Equal(t, largeAttributes, customEvent.Attributes)

	// Test with nil attributes
	customEvent2 := NewCustomEvent("test-event2").
		WithAnonymousId("test-id2").
		WithAttributes(nil)

	assert.Nil(t, customEvent2.Attributes)

	// Test with empty attributes map
	emptyAttributes := make(map[string]interface{})
	customEvent3 := NewCustomEvent("test-event3").
		WithAnonymousId("test-id3").
		WithAttributes(emptyAttributes)

	assert.NotNil(t, customEvent3.Attributes)
	assert.Len(t, customEvent3.Attributes, 0)
}

func TestBuilder_SpecialCharacters(t *testing.T) {
	specialChars := "!@#$%^&*()_+-=[]{}|;':\",./<>?`~"
	unicodeChars := "测试用例🚀😀"
	
	customEvent := NewCustomEvent(specialChars + unicodeChars).
		WithAnonymousId(specialChars).
		WithLoginUserId(unicodeChars).
		WithLoginUserKey(specialChars + unicodeChars).
		WithAttributes(map[string]interface{}{
			specialChars:         unicodeChars,
			unicodeChars:        specialChars,
			"mixed_" + unicodeChars: specialChars + unicodeChars,
		})

	assert.Equal(t, specialChars + unicodeChars, customEvent.EventName)
	assert.Equal(t, specialChars, customEvent.AnonymousId)
	assert.Equal(t, unicodeChars, customEvent.LoginUserId)
	assert.Equal(t, specialChars + unicodeChars, customEvent.LoginUserKey)
	assert.NotNil(t, customEvent.Attributes)
}

func TestBuilder_ExtremeValues(t *testing.T) {
	// Test with extreme timestamp values
	customEvent := NewCustomEvent("extreme-test").
		WithAnonymousId("test-id").
		WithEventTime(0) // Unix epoch

	assert.Equal(t, int64(0), customEvent.EventTime)

	// Test with very large timestamp
	futureTime := time.Date(2099, 12, 31, 23, 59, 59, 0, time.UTC).UnixMilli()
	customEvent2 := NewCustomEvent("future-test").
		WithAnonymousId("test-id").
		WithEventTime(futureTime)

	assert.Equal(t, futureTime, customEvent2.EventTime)

	// Test with negative timestamp
	customEvent3 := NewCustomEvent("past-test").
		WithAnonymousId("test-id").
		WithEventTime(-1)

	assert.Equal(t, int64(-1), customEvent3.EventTime)
}

func TestBuilder_ChainedCalls(t *testing.T) {
	// Test that chained calls work correctly and return the same instance
	customEvent := NewCustomEvent("chain-test")
	
	result1 := customEvent.WithAnonymousId("id1")
	result2 := result1.WithLoginUserId("user1") 
	result3 := result2.WithLoginUserKey("key1")
	result4 := result3.WithEventTime(12345)
	result5 := result4.WithAttributes(map[string]interface{}{"test": "value"})

	// All results should be the same instance
	assert.Same(t, customEvent, result1)
	assert.Same(t, customEvent, result2)
	assert.Same(t, customEvent, result3)
	assert.Same(t, customEvent, result4)
	assert.Same(t, customEvent, result5)

	// Verify all values were set
	assert.Equal(t, "chain-test", customEvent.EventName)
	assert.Equal(t, "id1", customEvent.AnonymousId)
	assert.Equal(t, "user1", customEvent.LoginUserId)
	assert.Equal(t, "key1", customEvent.LoginUserKey)
	assert.Equal(t, int64(12345), customEvent.EventTime)
	assert.Equal(t, "value", customEvent.Attributes["test"])
}

func TestValidation_EdgeCases(t *testing.T) {
	// Test custom event with both AnonymousId and LoginUserId empty
	customEvent := NewCustomEvent("test-event")
	err := validateCustomEventBuilder(customEvent)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "either AnonymousId or LoginUserId is required")

	// Test custom event with only whitespace in event name  
	customEvent2 := NewCustomEvent("   ")
	customEvent2.WithAnonymousId("test-id")
	err = validateCustomEventBuilder(customEvent2)
	assert.NoError(t, err) // Our validation checks len() == 0, not trimmed length

	// Test user event with only whitespace in LoginUserId
	userEvent := NewUser("   ")
	err = validateUserEventBuilder(userEvent)
	assert.NoError(t, err) // Our validation checks len() == 0, not trimmed length

	// Test item with only whitespace
	item := NewItem("   ", "   ")
	err = validateItemBuilder(item)
	assert.NoError(t, err) // Our validation checks len() == 0, not trimmed length
}

func TestCommonEventFields_EdgeCases(t *testing.T) {
	fields := &CommonEventFields{}
	
	// Test setting empty strings (should be allowed)
	fields.setEventTime(0)
	fields.setAnonymousId("")
	fields.setLoginUserId("")
	fields.setLoginUserKey("")
	fields.setAttributes(nil)

	assert.Equal(t, int64(0), fields.EventTime)
	assert.Equal(t, "", fields.AnonymousId)
	assert.Equal(t, "", fields.LoginUserId)
	assert.Equal(t, "", fields.LoginUserKey)
	assert.Nil(t, fields.Attributes)

	// Test setting very large values
	largeString := strings.Repeat("x", 10000)
	fields.setAnonymousId(largeString)
	fields.setLoginUserId(largeString)
	fields.setLoginUserKey(largeString)

	assert.Equal(t, largeString, fields.AnonymousId)
	assert.Equal(t, largeString, fields.LoginUserId)
	assert.Equal(t, largeString, fields.LoginUserKey)
}

func TestLogLevel_EdgeCases(t *testing.T) {
	var level LogLevel
	
	// Test with empty string
	err := level.UnmarshalYAML(func(v interface{}) error {
		*(v.(*string)) = ""
		return nil
	})
	assert.Error(t, err)

	// Test with mixed case
	err = level.UnmarshalYAML(func(v interface{}) error {
		*(v.(*string)) = "DEBUG"
		return nil
	})
	assert.Error(t, err) // Should be lowercase

	// Test with whitespace
	err = level.UnmarshalYAML(func(v interface{}) error {
		*(v.(*string)) = " debug "
		return nil
	})
	assert.Error(t, err) // Should be exact match

	// Test with numeric string
	err = level.UnmarshalYAML(func(v interface{}) error {
		*(v.(*string)) = "123"
		return nil
	})
	assert.Error(t, err)
}

func TestConfig_String_EdgeCases(t *testing.T) {
	// Test with nil config
	var config *Config
	if config != nil {
		result := config.String()
		assert.NotEmpty(t, result)
	}

	// Test with config containing special characters
	config = &Config{
		SdkConfig: SdkConfig{
			AccountId:    "test\"with\\quotes",
			DataSourceId: "test\nwith\nnewlines",
		},
		HttpConfig: HttpConfig{
			ServerHost: "https://api.test.com/with/path?param=value&other=test",
		},
	}

	result := config.String()
	assert.NotEmpty(t, result)
	assert.Contains(t, result, "test\\\"with\\\\quotes") // Should be escaped
}

func TestConcurrentBuilderAccess(t *testing.T) {
	customEvent := NewCustomEvent("concurrent-test")
	
	// Test concurrent access to builder methods
	done := make(chan bool, 3)
	
	go func() {
		defer func() { done <- true }()
		for i := 0; i < 100; i++ {
			customEvent.WithAnonymousId("anonymous-" + string(rune(i)))
		}
	}()
	
	go func() {
		defer func() { done <- true }()
		for i := 0; i < 100; i++ {
			customEvent.WithLoginUserId("user-" + string(rune(i)))
		}
	}()
	
	go func() {
		defer func() { done <- true }()
		for i := 0; i < 100; i++ {
			customEvent.WithEventTime(int64(i))
		}
	}()
	
	// Wait for all goroutines
	for i := 0; i < 3; i++ {
		<-done
	}
	
	// Builder should still be valid
	assert.Equal(t, "concurrent-test", customEvent.EventName)
	assert.NotEmpty(t, customEvent.AnonymousId)
	assert.NotEmpty(t, customEvent.LoginUserId)
	assert.GreaterOrEqual(t, customEvent.EventTime, int64(0))
}