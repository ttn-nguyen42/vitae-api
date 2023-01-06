package utils

import (
	"Vitae/config"
	"os"
)

func IsProduction() bool {
	mode := os.Getenv(config.EnvGinMode)
	if mode == "" {
		return false
	}
	if mode != "production" {
		return false
	}
	return true
}