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
	"os"
	"strings"
	"sync"

	core "github.com/growingio/growingio-sdk-go/internal/core"
	logger "github.com/growingio/growingio-sdk-go/internal/logger"
	"gopkg.in/yaml.v3"
)

var initMutex sync.Mutex

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
	// create the logger before any calls are made
	logger.NewLogger()

	if core.InitializedSuccessfully {
		err := errors.New("initialization failed, already initialized")
		logger.Error(err, "Do not initialize GrowingAnalytics repeatedly")
		return err
	}

	logger.SetLogLevel(int(config.SdkConfig.LogLevel))

	if err := validateConfig(config); err != nil {
		logger.Error(err, "Config validation failed")
		return err
	}
	core.AccountId = config.SdkConfig.AccountId
	core.DataSourceId = config.SdkConfig.DataSourceId

	serverHost := strings.TrimSuffix(config.HttpConfig.ServerHost, "/")

	core.Urls = core.RequestUrl{
		Collect: serverHost + "/v3/projects/" + config.SdkConfig.AccountId + "/collect",
		Item:    serverHost + "/projects/" + config.SdkConfig.AccountId + "/collect/item",
	}
	if config.HttpConfig.RequestTimeout > 0 {
		core.RequestTimeout = config.HttpConfig.RequestTimeout
	}

	if config.BatchConfig.Enable {
		core.BatchEnable = config.BatchConfig.Enable
		core.MaxSize = config.BatchConfig.MaxSize
		core.FlushAfter = config.BatchConfig.FlushAfter
		core.RoutineCount = config.BatchConfig.RoutineCount
		core.MaxCacheSize = config.BatchConfig.MaxCacheSize
		core.InitBatch()
		core.RunBatch()
	}

	core.InitializedSuccessfully = true
	logger.Info("Thank you very much for using GrowingIO. We will do our best to provide you with the best service.", "sdkVersion", core.SdkVersion)
	logger.Debug(config.String())
	return nil
}

func InitAnalyticsByConfigFile(file string) error {
	yamlFile, err := os.ReadFile(file)
	if err != nil {
		logger.Error(err, "Read File failed", "filePath", file)
		return err
	}
	config := &Config{}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		logger.Error(err, "yaml unmarshal failed", "filePath", file)
		return err
	}
	return InitAnalytics(config)
}

func TrackCustomEvent(b *CustomEventBuilder) error {
	if !core.InitializedSuccessfully {
		err := errors.New("GrowingAnalytics has not been initialized")
		logger.Error(err, "Please init GrowingAnalytics first")
		return err
	}

	if err := validateCustomEventBuilder(b); err != nil {
		logger.Error(err, "CustomEventBuilder validation failed")
		return err
	}

	builder := &core.EventBuilder{
		EventName:    b.EventName,
		EventTime:    b.EventTime,
		AnonymousId:  b.AnonymousId,
		LoginUserId:  b.LoginUserId,
		LoginUserKey: b.LoginUserKey,
		Attributes:   b.Attributes,
	}
	core.BuildCustomEvent(builder)
	return nil
}

func TrackUser(b *UserEventBuilder) error {
	if !core.InitializedSuccessfully {
		err := errors.New("GrowingAnalytics has not been initialized")
		logger.Error(err, "Please init GrowingAnalytics first")
		return err
	}

	if err := validateUserEventBuilder(b); err != nil {
		logger.Error(err, "UserEventBuilder validation failed")
		return err
	}

	builder := &core.EventBuilder{
		EventTime:    b.EventTime,
		AnonymousId:  b.AnonymousId,
		LoginUserId:  b.LoginUserId,
		LoginUserKey: b.LoginUserKey,
		Attributes:   b.Attributes,
	}
	core.BuildUserLoginEvent(builder)
	return nil
}

func SubmitItem(b *ItemBuilder) error {
	if !core.InitializedSuccessfully {
		err := errors.New("GrowingAnalytics has not been initialized")
		logger.Error(err, "Please init GrowingAnalytics first")
		return err
	}

	if err := validateItemBuilder(b); err != nil {
		logger.Error(err, "ItemBuilder validation failed")
		return err
	}

	builder := &core.EventBuilder{
		ItemId:     b.ItemId,
		ItemKey:    b.ItemKey,
		Attributes: b.Attributes,
	}
	core.BuildItemEvent(builder)
	return nil
}
