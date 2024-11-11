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

package logger

import "sync"

const (
	debugLogLevel int = 8
	infoLogLevel  int = 4
	warnLogLevel  int = 1
	errorLogLevel int = 0
)

type Logger interface {
	setLogLevel(logLevel int)
	debug(msg string, keysAndValues ...any)
	info(msg string, keysAndValues ...any)
	warn(msg string, keysAndValues ...any)
	error_log(err error, msg string, keysAndValues ...any)
}

var (
	glog     Logger
	logLevel int
	once     sync.Once
)

func NewLogger() {
	once.Do(func() {
		glog = newLogger()
	})
}

func SetLogLevel(level int) {
	logLevel = level
	glog.setLogLevel(level)
}

func Debug(msg string, keysAndValues ...any) {
	if glog == nil {
		return
	}
	if logLevel == debugLogLevel {
		keysAndValues = append([]any{"goroutineId", curGoroutineID()}, keysAndValues...)
	}
	glog.debug(msg, keysAndValues...)
}

func Info(msg string, keysAndValues ...any) {
	if glog == nil {
		return
	}
	if logLevel == debugLogLevel {
		keysAndValues = append([]any{"goroutineId", curGoroutineID()}, keysAndValues...)
	}
	glog.info(msg, keysAndValues...)
}

func Warn(msg string, keysAndValues ...any) {
	if glog == nil {
		return
	}
	if logLevel == debugLogLevel {
		keysAndValues = append([]any{"goroutineId", curGoroutineID()}, keysAndValues...)
	}
	glog.warn(msg, keysAndValues...)
}

func Error(err error, msg string, keysAndValues ...any) {
	if glog == nil {
		return
	}
	keysAndValues = append([]any{"goroutineId", curGoroutineID()}, keysAndValues...)
	glog.error_log(err, msg, keysAndValues...)
}
