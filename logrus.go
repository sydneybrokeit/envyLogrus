package envyLogrus

import (
	"github.com/sydneybrokeit/envy"
	"github.com/sirupsen/logrus"
)

var DefaultLogLevels = map[string]logrus.Level{
	"DEBUG": logrus.DebugLevel,
	"INFO": logrus.InfoLevel,
	"WARN": logrus.WarnLevel,
	"ERROR": logrus.ErrorLevel,
}

// Takes an environment variable name (env_key), and a fallback value (fallback)
// and returns the 'default' mapping of logrus LogLevels
func GetLogLevel(env_key, fallback string) (logLevel logrus.Level) {
	return GetCustomLogLevel(env_key, fallback, DefaultLogLevels)
}

func GetCustomLogLevel(env_key, fallback string, logLevels map[string]logrus.Level) (output logrus.Level){
	logFromEnv := envy.GetEnv(env_key, fallback)
	if val, ok := logLevels[logFromEnv]; ok {
		return val
	} else if val, ok := logLevels[fallback]; ok {
		return val
	} else {
		return logrus.DebugLevel
	}
}

