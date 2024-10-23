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

import (
	"fmt"
	stdlog "log"
	"os"
	"runtime"
	"sync"

	"github.com/go-logr/logr"
	"github.com/go-logr/stdr"
)

func SetLogLevel(logLevel int) {
	stdr.SetVerbosity(logLevel)
}

func Debug(msg string, keysAndValues ...any) {
	combined := append([]any{"goroutineId", getGoroutineID()}, keysAndValues...)
	sharedInstance().V(debugLogLevel).Info(msg, combined...)
}

func Info(msg string, keysAndValues ...any) {
	combined := append([]any{"goroutineId", getGoroutineID()}, keysAndValues...)
	sharedInstance().V(infoLogLevel).Info(msg, combined...)
}

func Warn(msg string, keysAndValues ...any) {
	combined := append([]any{"goroutineId", getGoroutineID()}, keysAndValues...)
	sharedInstance().V(warnLogLevel).Info(msg, combined...)
}

func Error(err error, msg string, keysAndValues ...any) {
	combined := append([]any{"goroutineId", getGoroutineID()}, keysAndValues...)
	sharedInstance().Error(err, msg, combined...)
}

func getGoroutineID() uint64 {
	var buf [64]byte
	runtime.Stack(buf[:], false)
	var id uint64
	fmt.Sscanf(string(buf[:]), "goroutine %d", &id)
	return id
}

const (
	debugLogLevel int = 8
	infoLogLevel  int = 4
	warnLogLevel  int = 1
	errorLogLevel int = 0
)

var (
	logger logr.Logger
	once   sync.Once
)

func initLogger() {
	logger = stdr.New(stdlog.New(os.Stderr, "[GrowingAnalytics] ", stdlog.LstdFlags|stdlog.Lmsgprefix))
}

func sharedInstance() logr.Logger {
	once.Do(initLogger)
	return logger
}
