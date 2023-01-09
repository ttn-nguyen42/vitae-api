package database

import (
	"Vitae/config"
	"fmt"
	"os"
	"testing"
)

func TestGetDriver_ReturnsCorrectDriverWhenAllEnvSupplied(t *testing.T) {
	os.Clearenv()
	mockClusterId := "Cluster0"
	os.Setenv(config.EnvMongoClusterID, mockClusterId)
	mockUsername := "Username"
	os.Setenv(config.EnvMongoUsername, mockUsername)
	mockPassword := "Password"
	os.Setenv(config.EnvMongoPassword, mockPassword)

	gotDriver, gotDatabaseId, err := getDriver()

	if err != nil {
		t.Errorf("Expected no error, got %v", err.Error())
	}
	expectedDriver := fmt.Sprintf("mongodb+srv://%v:%v@%v.mongodb.net/%v", mockUsername, mockPassword, mockClusterId, "?retryWrites=true&w=majority")
	if gotDriver != expectedDriver {
		t.Errorf("Expected driver equals '%v', got '%v'", expectedDriver, gotDriver)
	}
	if gotDatabaseId != mockClusterId {
		t.Errorf("Expected database ID equals %v, got %v", mockClusterId, gotDatabaseId)
	}
}

func TestGetDriver_ReturnsErrorWhenMissingUsername(t *testing.T) {
	os.Clearenv()
	mockClusterId := "Cluster0"
	os.Setenv(config.EnvMongoClusterID, mockClusterId)
	mockPassword := "Password"
	os.Setenv(config.EnvMongoPassword, mockPassword)

	gotDriver, gotDatabaseId, err := getDriver()

	if err == nil {
		t.Errorf("Expected error of %v", err.Error())
	}

	if gotDriver != "" {
		t.Errorf("Expected no driver returned, got %v", gotDriver)
	}
	if gotDatabaseId != "" {
		t.Errorf("Expected no database ID returned, got %v", gotDatabaseId)
	}
}

func TestGetDriver_ReturnsErrorWhenMissingPassword(t *testing.T) {
	os.Clearenv()
	mockClusterId := "Cluster0"
	os.Setenv(config.EnvMongoClusterID, mockClusterId)
	mockUsername := "Username"
	os.Setenv(config.EnvMongoUsername, mockUsername)

	gotDriver, gotDatabaseId, err := getDriver()

	if err == nil {
		t.Errorf("Expected error of %v", err.Error())
	}

	if gotDriver != "" {
		t.Errorf("Expected no driver returned, got %v", gotDriver)
	}
	if gotDatabaseId != "" {
		t.Errorf("Expected no database ID returned, got %v", gotDatabaseId)
	}
}

func TestGetDriver_ReturnsErrorWhenMissingClusterId(t *testing.T) {
	os.Clearenv()
	mockUsername := "Username"
	os.Setenv(config.EnvMongoUsername, mockUsername)
	mockPassword := "Password"
	os.Setenv(config.EnvMongoPassword, mockPassword)

	gotDriver, gotDatabaseId, err := getDriver()

	if err == nil {
		t.Errorf("Expected error of %v", err.Error())
	}

	if gotDriver != "" {
		t.Errorf("Expected no driver returned, got %v", gotDriver)
	}
	if gotDatabaseId != "" {
		t.Errorf("Expected no database ID returned, got %v", gotDatabaseId)
	}
}
