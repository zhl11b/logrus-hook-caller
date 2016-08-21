package hook

import (
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

	logrus.Fatal("cool!")
	logrus.WithField("class", 5).Warn("mybe cool")
}
