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
	"errors"
	"sync"

	logger "github.com/growingio/growingio-sdk-go/internal/logger"
)

var initMutex sync.Mutex
var defaultClient *Client

func validateConfig(config *Config) error {
	if config == nil {
		return errors.New("config cannot be nil")
	}
	
	if config.SdkConfig.AccountId == "" {
		return errors.New("AccountId is required")
	}
	
	if config.SdkConfig.DataSourceId == "" {
		return errors.New("DataSourceId is required")
	}
	
	if config.HttpConfig.ServerHost == "" {
		return errors.New("ServerHost is required")
	}
	
	if config.HttpConfig.RequestTimeout < 0 {
		return errors.New("RequestTimeout cannot be negative")
	}
	
	if config.BatchConfig.Enable {
		if config.BatchConfig.MaxSize < 0 {
			return errors.New("BatchConfig.MaxSize cannot be negative")
		}
		if config.BatchConfig.FlushAfter < 0 {
			return errors.New("BatchConfig.FlushAfter cannot be negative")
		}
		if config.BatchConfig.RoutineCount < 0 {
			return errors.New("BatchConfig.RoutineCount cannot be negative")
		}
		if config.BatchConfig.MaxCacheSize < 0 {
			return errors.New("BatchConfig.MaxCacheSize cannot be negative")
		}
	}
	
	return nil
}

func validateCustomEventBuilder(b *CustomEventBuilder) error {
	if b == nil {
		return errors.New("CustomEventBuilder cannot be nil")
	}
	
	if len(b.EventName) == 0 {
		return errors.New("EventName is required")
	}
	
	if len(b.AnonymousId)+len(b.LoginUserId) == 0 {
		return errors.New("either AnonymousId or LoginUserId is required")
	}
	
	return nil
}

func validateUserEventBuilder(b *UserEventBuilder) error {
	if b == nil {
		return errors.New("UserEventBuilder cannot be nil")
	}
	
	if len(b.LoginUserId) == 0 {
		return errors.New("LoginUserId is required")
	}
	
	return nil
}

func validateItemBuilder(b *ItemBuilder) error {
	if b == nil {
		return errors.New("ItemBuilder cannot be nil")
	}
	
	if len(b.ItemId) == 0 {
		return errors.New("ItemId is required")
	}
	
	if len(b.ItemKey) == 0 {
		return errors.New("ItemKey is required")
	}
	
	return nil
}

func InitAnalytics(config *Config) error {
	initMutex.Lock()
	defer initMutex.Unlock()

	if defaultClient != nil {
		err := errors.New("initialization failed, already initialized")
		logger.Error(err, "Do not initialize GrowingAnalytics repeatedly")
		return err
	}

	client, err := NewClient(config)
	if err != nil {
		return err
	}

	defaultClient = client
	return nil
}

func InitAnalyticsByConfigFile(file string) error {
	initMutex.Lock()
	defer initMutex.Unlock()

	if defaultClient != nil {
		err := errors.New("initialization failed, already initialized")
		logger.Error(err, "Do not initialize GrowingAnalytics repeatedly")
		return err
	}

	client, err := NewClientFromFile(file)
	if err != nil {
		return err
	}

	defaultClient = client
	return nil
}

func TrackCustomEvent(b *CustomEventBuilder) error {
	if defaultClient == nil {
		err := errors.New("GrowingAnalytics has not been initialized")
		logger.Error(err, "Please init GrowingAnalytics first")
		return err
	}
	return defaultClient.TrackCustomEvent(b)
}

func TrackUser(b *UserEventBuilder) error {
	if defaultClient == nil {
		err := errors.New("GrowingAnalytics has not been initialized")
		logger.Error(err, "Please init GrowingAnalytics first")
		return err
	}
	return defaultClient.TrackUser(b)
}

func SubmitItem(b *ItemBuilder) error {
	if defaultClient == nil {
		err := errors.New("GrowingAnalytics has not been initialized")
		logger.Error(err, "Please init GrowingAnalytics first")
		return err
	}
	return defaultClient.SubmitItem(b)
}
