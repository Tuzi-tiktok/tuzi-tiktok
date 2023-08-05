package logger

import (
	"testing"
)

func TestInit(t *testing.T) {
	logger.Debug("Debug")
	logger.Info("Info")
	logger.Warn("Warn")
	logger.Error("Error")
}
