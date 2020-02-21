package envyLogrus

import (
	"github.com/sirupsen/logrus"
	"os"
	"testing"
)

func TestGetLogLevel(t *testing.T) {
	if result := GetLogLevel("LOG_LEVEL", "TEST"); result != logrus.DebugLevel {
		t.Error("Expected debug (TEST), got ", result)
	}
	if result := GetLogLevel("LOG_LEVEL", "DEBUG"); result != logrus.DebugLevel {
		t.Error("Expected debug, got ", result)
	}
	if result := GetLogLevel("LOG_LEVEL", "INFO"); result != logrus.InfoLevel {
		t.Error("Expected INFO, got ", result)
	}
	if result := GetLogLevel("LOG_LEVEL", "WARN"); result != logrus.WarnLevel {
		t.Error("Expected WARN, got ", result)
	}
	if result := GetLogLevel("LOG_LEVEL", "ERROR"); result != logrus.ErrorLevel {
		t.Error("Expected ERROR, got ", result)
	}
	if result := GetLogLevel("LOG_LEVEL", "FATAL"); result != logrus.DebugLevel {
		t.Error("Expected DEBUG (from ENV), got", result)
	}
	os.Setenv("LOG_LEVEL", "ERROR")
	if result := GetLogLevel("LOG_LEVEL", "FATAL"); result != logrus.ErrorLevel {
		t.Error("Expected ERROR (from ENV), got", result)
	}
}

func TestGetCustomLogLevel(t *testing.T) {
	customLogLevels := map[string]logrus.Level{
		"FATAL": logrus.FatalLevel,
		"GARBAGE": logrus.TraceLevel,
	}
	if result := GetCustomLogLevel("LOG_LEVEL", "FATAL", customLogLevels); result != logrus.FatalLevel {
		t.Error("Custom Log Level: Expected FATAL, received", result)
	}
}
