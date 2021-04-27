package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Log is a logger instance to be used in other packages
var Log = logrus.New()

func init() {
	Log.Out = os.Stdout
	SetLogLevel(Log, "debug", logrus.InfoLevel)
}

// SetLogLevel sets the log level
func SetLogLevel(logger *logrus.Logger, l string, def logrus.Level) {
	level, err := logrus.ParseLevel(l)
	if err != nil {
		logger.Warnf("Could not parse log level, setting to %v", def)
		level = def
	}
	logger.SetLevel(level)
	logger.Infof("Log level set to %s", logger.GetLevel().String())
}
