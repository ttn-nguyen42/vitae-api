package utils

import (
	"Vitae/config"
	"fmt"
	"os"
)

func IsProduction() bool {
	mode := os.Getenv(config.EnvGinMode)
	if mode == "" {
		return false
	}
	if mode != "release" {
		return false
	}
	return true
}

func GetDatabaseName(baseName string) string {
    mode := os.Getenv(config.EnvGinMode)
    if mode == "test" {
        return fmt.Sprintf("%s-dev", baseName)
    }
    if mode == "staging" {
        return fmt.Sprintf("%s-staging", baseName)
    }
    if mode == "release" {
        return fmt.Sprintf("%s-prod", baseName)
    }
    panic("Unknown database name, due to unknown GIN_MODE")
}
