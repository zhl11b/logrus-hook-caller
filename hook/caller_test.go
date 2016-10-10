package hook

import (
	"os"
	"testing"

	"github.com/Sirupsen/logrus"
)

func TestCaller(t *testing.T) {
	// init
	logrus.AddHook(&CallerHook{})
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{})

	// display log
	logrus.Debug("debug info")
	logrus.WithFields(logrus.Fields{
		"name":   "john smith",
		"age":    23,
		"ismale": false,
	}).Info("debug info")

	// logrus.Fatal("cool!") // 注释这一句，会日志的fatal会直接退出程序
	logrus.Warn("cool")
	logrus.WithField("class", 5).Warn("mybe cool")
}

func BenchmarkCaller(b *testing.B) {
	// init
	filepath := "test_caller.log"
	logrus.AddHook(&CallerHook{})
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{})
	fd, err := os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		panic("open logfile failed!")
	}
	defer os.Remove(filepath)
	defer fd.Close()

	logrus.SetOutput(fd)

	// test
	for i := 0; i < b.N; i++ {
		logrus.Debug("nice boy")
	}
}

func BenchmarkWithField(b *testing.B) {
	// init
	filepath := "test_caller.log"
	logrus.AddHook(&CallerHook{})
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{})
	fd, err := os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		panic("open logfile failed!")
	}
	defer os.Remove(filepath)
	defer fd.Close()
	logrus.SetOutput(fd)

	// test
	for i := 0; i < b.N; i++ {
		logrus.WithField("name", "john").Warn("is married")
	}
}

func BenchmarkWithFields(b *testing.B) {
	// init
	filepath := "test_caller.log"
	logrus.AddHook(&CallerHook{})
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{})
	fd, err := os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		panic("open logfile failed!")
	}
	defer os.Remove(filepath)
	defer fd.Close()
	logrus.SetOutput(fd)

	// test
	for i := 0; i < b.N; i++ {
		logrus.WithFields(logrus.Fields{
			"name":   "john smith",
			"age":    32,
			"ismale": false,
			"class":  3,
		}).Info("record his info.")
	}
}
