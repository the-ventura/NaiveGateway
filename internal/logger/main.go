package logger

import (
	"os"

	"naivegateway/internal/config"

	"github.com/sirupsen/logrus"
)

// Log is a logger instance to be used in other packages
var Log = logrus.New()
var cfg = config.GetConfig()

func init() {
	Log.Out = os.Stdout
	SetLogLevel(Log, cfg.LogLevel, logrus.InfoLevel)
}

// SetLogLevel sets the log level
func SetLogLevel(logger *logrus.Logger, l string, def logrus.Level) {
	level, err := logrus.ParseLevel(l)
	if err != nil {
		logger.Warnf("Could not parse log level, setting to %v", def)
		level = def
	}
	// Enables more verbose output if the log level is above info
	if level > 4 {
		logger.SetReportCaller(true)
	}
	logger.SetLevel(level)
	logger.Infof("Log level set to %s", logger.GetLevel().String())
}
