package about

import (
	"Vitae/models"
	"Vitae/repositories"
	"Vitae/repositories/about"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type aboutRepositoryMock struct {
	/* Magic trick
	 * Because interfaces are embedded into each structs as methods of nil pointers
	 * It satisfies the IRepository interface since all methods are implemented, as nil pointers
	 * We just need to implement the required methods
	 */
	about.IRepository

	// Fake methods
	FakeAddOne func(document models.About) (string, error)
	FakeGetOne func(result *models.About, id string) error
}

func (m *aboutRepositoryMock) AddOne(document models.About) (string, error) {
	return m.FakeAddOne(document)
}

func (m *aboutRepositoryMock) GetOne(dto *models.About, id string) error {
	return m.FakeGetOne(dto, id)
}

func TestGetOne_WillGetOneWhenGivenCorrectId(t *testing.T) {
	testId := primitive.NewObjectIDFromTimestamp(time.Now())
	test := models.About{
		Id: testId,
		FirstName: "Nguyen",
		LastName: "Tran",
		Email: "ttn.nguyen42@gmail.com",
		Birthday: time.Now(),
		City: "Hanoi",
		Country: "Vietnam",
	}
	mockRepository := aboutRepositoryMock{
		FakeGetOne: func(result *models.About, id string) error {
			*result = test
			return nil
		},
	}
	service := NewService(&mockRepository)

	var gotData GetResponse
	gotErr := service.GetOne(&gotData, testId.Hex())

	if gotErr != nil {
		t.Errorf("Expected no error, got %v", gotErr.Error())
	}
	if gotData.FirstName != test.FirstName {
		t.Errorf("Expected FirstName to be %v, got %v", test.FirstName, gotData.FirstName)
	}
}

func TestGetOne_BreaksWhenGivenNotExistingId(t *testing.T) {
	testId := primitive.NewObjectIDFromTimestamp(time.Now())
	testMessage := "No existing ID found"
	mockRepository := aboutRepositoryMock{
		FakeGetOne: func(result *models.About, id string) error {
			return repositories.NewNotFoundError(testMessage)
		},
	}
	service := NewService(&mockRepository)
	var gotData GetResponse
	gotErr := service.GetOne(&gotData, testId.Hex())

	if gotErr == nil {
		t.Errorf("Expected error %v, got none", gotErr.Error())
	}
	if gotErr.Error() != testMessage {
		t.Errorf("Expected error %v, got %v", testMessage, gotErr.Error())
	}
	if gotData.FirstName != "" {
		t.Errorf("Expected empty FirstName field, got %v", gotData.FirstName)
	}
}

func TestGetOne_BreaksWhenDatabaseIsBroken(t *testing.T) {
	testId := primitive.NewObjectIDFromTimestamp(time.Now())
	testMessage := "Broken database"
	mockRepository := aboutRepositoryMock{
		FakeGetOne: func(result *models.About, id string) error {
			return repositories.NewInternalError(testMessage)
		},
	}
	service := NewService(&mockRepository)
	var gotData GetResponse
	gotErr := service.GetOne(&gotData, testId.Hex())

	if gotErr == nil {
		t.Errorf("Expected error %v, got none", testMessage)
	}
	if gotErr.Error() != testMessage {
		t.Errorf("Expected error %v, got %v", testMessage, gotErr.Error())
	}
	if gotData.FirstName != "" {
		t.Errorf("Expected empty FirstName field, got %v", gotData.FirstName)
	}
}

func TestGetOne_BreaksWhenInvalidIdIsSupplied(t *testing.T) {
	testId := "Definitely an invalid ID"
	testMessage := "Invalid ID"
	mockRepository := aboutRepositoryMock{
		FakeGetOne: func(result *models.About, id string) error {
			return repositories.NewInvalidIdError(testMessage)
		},
	}

	service := NewService(&mockRepository)
	var gotData GetResponse
	gotErr := service.GetOne(&gotData, testId)
	
	if gotErr == nil {
		t.Errorf("Expected error %v, got none", testMessage)
	}
	if gotErr.Error() != testMessage {
		t.Errorf("Expected error %v, got %v", testMessage, gotErr.Error())
	}
	if gotData.FirstName != "" {
		t.Errorf("Expected empty FirstName field, got %v", gotData.FirstName)
	}
}

func TestAddOne_WillAddOneWhenGivenCorrectDto(t *testing.T) {
	testId := primitive.NewObjectIDFromTimestamp(time.Now()).Hex()
	mockRepository := aboutRepositoryMock{
		FakeAddOne: func(document models.About) (string, error) {
			return testId, nil
		},
	}
	service := NewService(&mockRepository)
	test := PostRequest{
		FirstName: "Nguyen",
		LastName: "Tran",
		Email: "ttn.nguyen42@gmail.com",
		Birthday: time.Now(),
		City: "Hanoi",
		Country: "Vietnam",
	}
	gotId, gotErr := service.AddOne(test)
	if gotErr != nil {
		t.Errorf("Expected no error but got: %v", gotErr.Error())
	}
	if len(gotId) == 0 {
		t.Error("Expected to receive 0 length ID but got one")
	}
	if gotId != testId {
		t.Errorf("Expected ID of %v, got %v", testId, gotId)
	}
}

func TestAddOne_BreaksWhenDatabaseIsBroken(t *testing.T) {
	testErrorMessage := "Broken database"
	mockRepository := aboutRepositoryMock{
		FakeAddOne: func(document models.About) (string, error) {
			return "", repositories.NewInternalError(testErrorMessage)
		},
	}
	test := PostRequest{
		FirstName: "Nguyen",
		LastName: "Tran",
		Email: "ttn.nguyen42@gmail.com",
		Birthday: time.Now(),
		City: "Hanoi",
		Country: "Vietnam",
	}
	service := NewService(&mockRepository)
	gotId, gotErr := service.AddOne(test)
	if gotErr == nil {
		t.Errorf("Expected error '%v', got none", &testErrorMessage)
	}
	if len(gotErr.Error()) == 0 {
		t.Errorf("Expected error message to be '%v', got none", testErrorMessage)
	}
	if gotErr.Error() != testErrorMessage {
		t.Errorf("Expected error message to be '%v', got '%v'", testErrorMessage, gotErr.Error())
	}
	if gotId != "" {
		t.Errorf("Expected ID to be empty, got %v", gotId)
	}
}