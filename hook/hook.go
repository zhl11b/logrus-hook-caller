package hook

import (
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/Sirupsen/logrus"
)

// CallerHook 需要的钩子
type CallerHook struct {
}

// Fire 实现
func (hook *CallerHook) Fire(entry *logrus.Entry) error {
	if len(entry.Data) == 0 {
		entry.Data["caller"] = hook.caller2()
	} else {
		entry.Data["caller"] = hook.caller()
	}

	return nil
}

// Levels 实现的级别
func (hook *CallerHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
	}
}

// caller 带withfields的函数调用
func (hook *CallerHook) caller() string {
	if _, file, line, ok := runtime.Caller(5); ok {
		return strings.Join([]string{filepath.Base(file), strconv.Itoa(line)}, ":")
	}
	// not sure what the convention should be here
	return "???"
}

// caller2 不带withfields的函数调用
func (hook *CallerHook) caller2() string {
	if _, file, line, ok := runtime.Caller(7); ok {
		return strings.Join([]string{filepath.Base(file), strconv.Itoa(line)}, ":")
	}
	// not sure what the convention should be here
	return "???"
}
