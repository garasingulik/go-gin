package lib

import (
	"testing"
)

func LoggerTest(t *testing.T) {
	// Create logger
	log := CreateLogger("test")

	if log == nil {
		t.Error("Logger failed to create")
	}
}
