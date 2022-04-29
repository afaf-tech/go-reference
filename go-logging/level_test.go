package go_logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLevel(t *testing.T) {
	logger := logrus.New()

	logger.Trace("Something very low level.")
	logger.Debug("Useful debugging information.")
	logger.Info("Something noteworthy happened!")
	logger.Warn("You should probably take a look at this.")
	logger.Error("Something failed but I'm not quitting.")
	// Calls os.Exit(1) after logging
	logger.Fatal("Bye.")
	// Calls panic() after logging
	logger.Panic("I'm bailing.")
}
func TestLoggingLevel(t *testing.T) {
	logger := logrus.New()
	// when we set level to global logger settings, it will behave like the level specified
	logger.SetLevel(logrus.TraceLevel)

	logger.Trace("This is Trace")
	logger.Debug("This is Debug")
	logger.Info("This is Info")
	logger.Warn("This is Warn")
	logger.Error("This is Error")
}
