package go_logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestSingleton(t *testing.T) {
	// using object logger without creating new instance
	logrus.Info("Hello World!")
	logrus.Warn("Hello World!")

	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.Info("Hello World!")
	logrus.Warn("Hello World!")
}
