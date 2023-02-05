package config

import (
	"os"
	"testing"
)

func ConfigTest(t *testing.T) {
	os.Setenv("ENVIRONMENT", "test")

	// Init Config
	if err := Init(); err != nil {
		t.Error("Config init error")
	}
}
