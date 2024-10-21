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
