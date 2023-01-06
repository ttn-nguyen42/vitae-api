package utils

import (
	"Vitae/config"
	"os"
	"testing"
)

func TestIsProduction_ReturnsFalseWhenInDebugMode(t *testing.T) {
	os.Setenv(config.EnvGinMode, "debug")
	got := IsProduction()
	if got != false {
		t.Errorf("Expected %v, got %v", false, got)
	}
}

func TestIsProduction_ReturnsFalseWhenNotInReleaseMode(t *testing.T) {
	os.Setenv(config.EnvGinMode, "not release")
	got := IsProduction()
	if got != false {
		t.Errorf("Expected %v, got %v", false, got)
	}
}

func TestIsProduction_ReturnsTrueWhenInReleaseMode(t *testing.T) {
	os.Setenv(config.EnvGinMode, "release")
	got := IsProduction()
	if got != true {
		t.Errorf("Expected %v, got %v", true, got)
	}
}

func TestIsProduction_ReturnsFalseWhenModeIsNotSet(t *testing.T) {
	os.Clearenv()
	got := IsProduction()
	if got != false {
		t.Errorf("Expected %v, got %v", false, got)
	}
}