package config

import (
	"os"
	"testing"
)

func ServerTest(t *testing.T) {
	os.Setenv("ENVIRONMENT", "test")

	// Init Config
	if err := Init(); err != nil {
		t.Error("Config init error")
	}
}
