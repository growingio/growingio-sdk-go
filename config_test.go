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
	"gopkg.in/yaml.v3"
)

func TestConfig_String(t *testing.T) {
	config := &Config{
		SdkConfig: SdkConfig{
			AccountId:    "test-account",
			DataSourceId: "test-data-source",
			LogLevel:     InfoLogLevel,
		},
		HttpConfig: HttpConfig{
			ServerHost:     "https://api.example.com",
			RequestTimeout: 30,
		},
		BatchConfig: BatchConfig{
			Enable:       true,
			MaxSize:      100,
			FlushAfter:   5,
			RoutineCount: 8,
			MaxCacheSize: 1000,
		},
	}

	jsonStr := config.String()
	assert.NotEmpty(t, jsonStr)
	assert.Contains(t, jsonStr, "test-account")
	assert.Contains(t, jsonStr, "test-data-source")
}

func TestLogLevel_UnmarshalYAML(t *testing.T) {
	tests := []struct {
		name     string
		yamlStr  string
		expected LogLevel
		wantErr  bool
	}{
		{
			name:     "debug level",
			yamlStr:  "debug",
			expected: DebugLogLevel,
			wantErr:  false,
		},
		{
			name:     "info level",
			yamlStr:  "info",
			expected: InfoLogLevel,
			wantErr:  false,
		},
		{
			name:     "warn level",
			yamlStr:  "warn",
			expected: WarnLogLevel,
			wantErr:  false,
		},
		{
			name:     "error level",
			yamlStr:  "error",
			expected: ErrorLogLevel,
			wantErr:  false,
		},
		{
			name:    "invalid level",
			yamlStr: "invalid",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var level LogLevel
			err := yaml.Unmarshal([]byte(tt.yamlStr), &level)
			
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, level)
			}
		})
	}
}

func TestLogLevel_Constants(t *testing.T) {
	assert.Equal(t, LogLevel(8), DebugLogLevel)
	assert.Equal(t, LogLevel(4), InfoLogLevel)
	assert.Equal(t, LogLevel(1), WarnLogLevel)
	assert.Equal(t, LogLevel(0), ErrorLogLevel)
}