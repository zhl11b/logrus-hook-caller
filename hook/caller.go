package hook

import (
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

// CallerHook 需要的钩子
type CallerHook struct {
}

// Fire 实现
func (hook *CallerHook) Fire(entry *logrus.Entry) error {
	var delta int
	if _, ok := entry.Data["caller"]; ok {
		delta = -1
	}
	if len(entry.Data)+delta == 0 { // if don't use withfields
		entry.Data["caller"] = hook.caller(9)
	} else { // if use withfields
		entry.Data["caller"] = hook.caller(7)
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
func (hook *CallerHook) caller(skip int) string {
	if _, file, line, ok := runtime.Caller(skip); ok {
		return strings.Join([]string{filepath.Base(file), strconv.Itoa(line)}, ":")
	}
	// not sure what the convention should be here
	return "???"
}
