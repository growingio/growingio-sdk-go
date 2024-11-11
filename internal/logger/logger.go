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
	stdlog "log"
	"os"

	"github.com/go-logr/logr"
	"github.com/go-logr/stdr"
)

type logger struct {
	vendor logr.Logger
}

func newLogger() Logger {
	vendor := stdr.New(stdlog.New(os.Stderr, "[GrowingAnalytics] ", stdlog.LstdFlags|stdlog.Lmsgprefix))
	return &logger{
		vendor: vendor,
	}
}

func (l *logger) setLogLevel(logLevel int) {
	stdr.SetVerbosity(logLevel)
}

func (l *logger) debug(msg string, keysAndValues ...any) {
	l.vendor.V(debugLogLevel).Info(msg, keysAndValues...)
}

func (l *logger) info(msg string, keysAndValues ...any) {
	l.vendor.V(infoLogLevel).Info(msg, keysAndValues...)
}

func (l *logger) warn(msg string, keysAndValues ...any) {
	l.vendor.V(warnLogLevel).Info(msg, keysAndValues...)
}

func (l *logger) error_log(err error, msg string, keysAndValues ...any) {
	l.vendor.Error(err, msg, keysAndValues...)
}
