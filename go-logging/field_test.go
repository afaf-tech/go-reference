package go_logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

type MyType struct {
	Username string
	Email    string
}

func TestField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("user", MyType{Username: "@local", Email: "local@gmail.com"}).Info("Hello World!")
	logger.WithFields(logrus.Fields{
		"username": "fatan",
		"email":    "fatan@gmail.com",
	}).Info("Hello World!")
}
