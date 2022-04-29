package go_logging

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestOutput(t *testing.T) {
	logger := logrus.New()

	file, _ := os.OpenFile("application.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	logger.SetOutput(file)
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.Info("Hello Logging")
	logger.Warn("Hello Logging")
	logger.Error("Hello Logging")
}
