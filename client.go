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

// Client represents a GrowingIO Analytics client instance
type Client struct {
	config      *Config
	initialized bool
	mutex       sync.RWMutex
}

// NewClient creates a new Analytics client with the provided configuration
func NewClient(config *Config) (*Client, error) {
	if err := validateConfig(config); err != nil {
		return nil, err
	}

	client := &Client{
		config: config,
	}

	if err := client.initialize(); err != nil {
		return nil, err
	}

	return client, nil
}

// NewClientFromFile creates a new Analytics client from a YAML configuration file
func NewClientFromFile(configFile string) (*Client, error) {
	yamlFile, err := os.ReadFile(configFile)
	if err != nil {
		logger.Error(err, "Read File failed", "filePath", configFile)
		return nil, err
	}
	
	config := &Config{}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		logger.Error(err, "yaml unmarshal failed", "filePath", configFile)
		return nil, err
	}
	
	return NewClient(config)
}

func (c *Client) initialize() error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.initialized {
		return errors.New("client already initialized")
	}

	// create the logger before any calls are made
	logger.NewLogger()
	logger.SetLogLevel(int(c.config.SdkConfig.LogLevel))

	core.AccountId = c.config.SdkConfig.AccountId
	core.DataSourceId = c.config.SdkConfig.DataSourceId

	serverHost := strings.TrimSuffix(c.config.HttpConfig.ServerHost, "/")

	core.Urls = core.RequestUrl{
		Collect: serverHost + "/v3/projects/" + c.config.SdkConfig.AccountId + "/collect",
		Item:    serverHost + "/projects/" + c.config.SdkConfig.AccountId + "/collect/item",
	}
	
	if c.config.HttpConfig.RequestTimeout > 0 {
		core.RequestTimeout = c.config.HttpConfig.RequestTimeout
	}

	if c.config.BatchConfig.Enable {
		core.BatchEnable = c.config.BatchConfig.Enable
		core.MaxSize = c.config.BatchConfig.MaxSize
		core.FlushAfter = c.config.BatchConfig.FlushAfter
		core.RoutineCount = c.config.BatchConfig.RoutineCount
		core.MaxCacheSize = c.config.BatchConfig.MaxCacheSize
		core.InitBatch()
		core.RunBatch()
	}

	core.InitializedSuccessfully = true
	c.initialized = true
	
	logger.Info("Thank you very much for using GrowingIO. We will do our best to provide you with the best service.", "sdkVersion", core.SdkVersion)
	logger.Debug(c.config.String())
	
	return nil
}

// TrackCustomEvent tracks a custom event using this client instance
func (c *Client) TrackCustomEvent(b *CustomEventBuilder) error {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	if !c.initialized {
		return errors.New("client not initialized")
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

// TrackUser tracks a user event using this client instance
func (c *Client) TrackUser(b *UserEventBuilder) error {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	if !c.initialized {
		return errors.New("client not initialized")
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

// SubmitItem submits an item using this client instance
func (c *Client) SubmitItem(b *ItemBuilder) error {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	if !c.initialized {
		return errors.New("client not initialized")
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