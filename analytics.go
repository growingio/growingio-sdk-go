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

	core "github.com/growingio/growingio-sdk-go/internal/core"
	logger "github.com/growingio/growingio-sdk-go/internal/logger"
	"gopkg.in/yaml.v3"
)

var lock = make(chan struct{}, 1)

func InitAnalytics(config *Config) error {
	lock <- struct{}{}
	defer func() { <-lock }()
	// create the logger before any calls are made
	logger.NewLogger()

	if core.InitializedSuccessfully {
		err := errors.New("initialization failed, already initialized")
		logger.Error(err, "Do not initialize GrowingAnalytics repeatedly")
		return err
	}

	logger.SetLogLevel(int(config.SdkConfig.LogLevel))

	if config.SdkConfig.AccountId == "" {
		err := errors.New("initialization failed, AccountId is empty")
		logger.Error(err, "Please enter your AccountId")
		return err
	}
	if config.SdkConfig.DataSourceId == "" {
		err := errors.New("initialization failed, DataSourceId is empty")
		logger.Error(err, "Please enter your DataSourceId")
		return err
	}
	core.AccountId = config.SdkConfig.AccountId
	core.DataSourceId = config.SdkConfig.DataSourceId

	if config.HttpConfig.ServerHost == "" {
		err := errors.New("initialization failed, ServerHost is empty")
		logger.Error(err, "Please enter your ServerHost")
		return err
	}
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
		err := errors.New("TrackCustomEvent failed, GrowingAnalytics has not been initialized")
		logger.Error(err, "Please init GrowingAnalytics first")
		return err
	}

	if len(b.EventName) == 0 {
		err := errors.New("TrackCustomEvent failed, EventName is empty")
		logger.Error(err, "Please enter eventName for customEvent")
		return err
	}

	if len(b.AnonymousId)+len(b.LoginUserId) == 0 {
		err := errors.New("TrackCustomEvent failed, AnonymousId and LoginUserId are empty")
		logger.Error(err, "Both AnonymousId and LoginUserId are empty. Please enter at least one of them")
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
		err := errors.New("TrackUser failed, GrowingAnalytics has not been initialized")
		logger.Error(err, "Please init GrowingAnalytics first")
		return err
	}

	if len(b.LoginUserId) == 0 {
		err := errors.New("TrackUser failed, LoginUserId is empty")
		logger.Error(err, "Please enter loginUserId for user")
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
		err := errors.New("SubmitItem failed, GrowingAnalytics has not been initialized")
		logger.Error(err, "Please init GrowingAnalytics first")
		return err
	}

	if len(b.ItemId) == 0 {
		err := errors.New("SubmitItem failed, ItemId is empty")
		logger.Error(err, "Please enter itemId for item")
		return err
	}

	if len(b.ItemKey) == 0 {
		err := errors.New("SubmitItem failed, ItemKey is empty")
		logger.Error(err, "Please enter itemKey for item")
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
