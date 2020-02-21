package envyLogrus

import (
	"github.com/hmschreck/envy"
	"github.com/sirupsen/logrus"
)

var DefaultLogLevels = map[string]logrus.Level{
	"DEBUG": logrus.DebugLevel,
	"INFO": logrus.InfoLevel,
	"WARN": logrus.WarnLevel,
	"ERROR": logrus.ErrorLevel,
}

func GetLogLevel(env_key, fallback string) (logLevel logrus.Level) {
	return GetCustomLogLevel(env_key, fallback, DefaultLogLevels)
}

func GetCustomLogLevel(env_key, fallback string, logLevels map[string]logrus.Level) (output logrus.Level){
	logFromEnv := envy.GetEnv(env_key, fallback)
	if val, ok := logLevels[logFromEnv]; ok {
		return logLevels[logFromEnv]
	} else if val, ok := logLevels[fallback]; ok {
		return logLevels[fallback]
	} else {
		return logrus.DebugLevel
	}
}
