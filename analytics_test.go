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

func TestValidateConfig(t *testing.T) {
	tests := []struct {
		name    string
		config  *Config
		wantErr bool
	}{
		{
			name:    "nil config",
			config:  nil,
			wantErr: true,
		},
		{
			name: "empty AccountId",
			config: &Config{
				SdkConfig: SdkConfig{
					AccountId:    "",
					DataSourceId: "test-data-source",
				},
				HttpConfig: HttpConfig{
					ServerHost: "https://api.example.com",
				},
			},
			wantErr: true,
		},
		{
			name: "empty DataSourceId",
			config: &Config{
				SdkConfig: SdkConfig{
					AccountId:    "test-account",
					DataSourceId: "",
				},
				HttpConfig: HttpConfig{
					ServerHost: "https://api.example.com",
				},
			},
			wantErr: true,
		},
		{
			name: "empty ServerHost",
			config: &Config{
				SdkConfig: SdkConfig{
					AccountId:    "test-account",
					DataSourceId: "test-data-source",
				},
				HttpConfig: HttpConfig{
					ServerHost: "",
				},
			},
			wantErr: true,
		},
		{
			name: "negative RequestTimeout",
			config: &Config{
				SdkConfig: SdkConfig{
					AccountId:    "test-account",
					DataSourceId: "test-data-source",
				},
				HttpConfig: HttpConfig{
					ServerHost:     "https://api.example.com",
					RequestTimeout: -1,
				},
			},
			wantErr: true,
		},
		{
			name: "valid config",
			config: &Config{
				SdkConfig: SdkConfig{
					AccountId:    "test-account",
					DataSourceId: "test-data-source",
				},
				HttpConfig: HttpConfig{
					ServerHost:     "https://api.example.com",
					RequestTimeout: 30,
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateConfig(tt.config)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestValidateCustomEventBuilder(t *testing.T) {
	tests := []struct {
		name    string
		builder *CustomEventBuilder
		wantErr bool
	}{
		{
			name:    "nil builder",
			builder: nil,
			wantErr: true,
		},
		{
			name: "empty EventName",
			builder: &CustomEventBuilder{
				EventName: "",
			},
			wantErr: true,
		},
		{
			name: "no user identification",
			builder: &CustomEventBuilder{
				EventName: "test-event",
			},
			wantErr: true,
		},
		{
			name: "valid with AnonymousId",
			builder: &CustomEventBuilder{
				EventName: "test-event",
				CommonEventFields: CommonEventFields{
					AnonymousId: "test-anonymous-id",
				},
			},
			wantErr: false,
		},
		{
			name: "valid with LoginUserId",
			builder: &CustomEventBuilder{
				EventName: "test-event",
				CommonEventFields: CommonEventFields{
					LoginUserId: "test-user-id",
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateCustomEventBuilder(tt.builder)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestValidateUserEventBuilder(t *testing.T) {
	tests := []struct {
		name    string
		builder *UserEventBuilder
		wantErr bool
	}{
		{
			name:    "nil builder",
			builder: nil,
			wantErr: true,
		},
		{
			name: "empty LoginUserId",
			builder: &UserEventBuilder{
				CommonEventFields: CommonEventFields{
					LoginUserId: "",
				},
			},
			wantErr: true,
		},
		{
			name: "valid builder",
			builder: &UserEventBuilder{
				CommonEventFields: CommonEventFields{
					LoginUserId: "test-user-id",
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateUserEventBuilder(tt.builder)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestValidateItemBuilder(t *testing.T) {
	tests := []struct {
		name    string
		builder *ItemBuilder
		wantErr bool
	}{
		{
			name:    "nil builder",
			builder: nil,
			wantErr: true,
		},
		{
			name: "empty ItemId",
			builder: &ItemBuilder{
				ItemId:  "",
				ItemKey: "test-key",
			},
			wantErr: true,
		},
		{
			name: "empty ItemKey",
			builder: &ItemBuilder{
				ItemId:  "test-id",
				ItemKey: "",
			},
			wantErr: true,
		},
		{
			name: "valid builder",
			builder: &ItemBuilder{
				ItemId:  "test-id",
				ItemKey: "test-key",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateItemBuilder(tt.builder)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}