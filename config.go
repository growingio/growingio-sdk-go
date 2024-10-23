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

package growingio

import (
	"encoding/json"
	"errors"
)

type Config struct {
	SdkConfig   SdkConfig   `yaml:"sdk"`
	HttpConfig  HttpConfig  `yaml:"http"`
	BatchConfig BatchConfig `yaml:"batch"`
}

func (c *Config) String() string {
	jsonString, err := json.Marshal(c)
	if err != nil {
		return ""
	}
	return string(jsonString)
}

type LogLevel int

const (
	DebugLogLevel LogLevel = 8
	InfoLogLevel  LogLevel = 4
	WarnLogLevel  LogLevel = 1
	ErrorLogLevel LogLevel = 0
)

func (l *LogLevel) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var level string
	if err := unmarshal(&level); err != nil {
		return err
	}

	switch level {
	case "debug":
		*l = DebugLogLevel
	case "info":
		*l = InfoLogLevel
	case "warn":
		*l = WarnLogLevel
	case "error":
		*l = ErrorLogLevel
	default:
		err := errors.New("unmarshal log_level failed, please check it")
		return err
	}
	return nil
}

type SdkConfig struct {
	AccountId    string   `yaml:"account_id"`
	DataSourceId string   `yaml:"data_source_id"`
	LogLevel     LogLevel `yaml:"log_level"`
}

type HttpConfig struct {
	ServerHost     string `yaml:"server_host"`
	RequestTimeout int    `yaml:"timeout"`
}

type BatchConfig struct {
	Enable       bool `yaml:"enable"`
	MaxSize      int  `yaml:"max_size"`
	FlushAfter   int  `yaml:"flush_after"`
	RoutineCount int  `yaml:"routine_count"`
	MaxCacheSize int  `yaml:"max_cache_size"`
}
