package about

import (
	"Vitae/repositories"
	v1 "Vitae/routes/v1"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type aboutReaderServiceMock struct {
	IReader

	FakeGetOne func(dto *GetResponse, id string) error
	FakeGetAll func(amount int) ([]GetResponse, error)
}

type aboutWriterServiceMock struct {
	IWriter

	FakeAddOne func(dto PostRequest) (string, error)
}

func (r *aboutReaderServiceMock) GetOne(dto *GetResponse, id string) error {
	return r.FakeGetOne(dto, id)
}

func (r *aboutReaderServiceMock) GetAll(amount int) ([]GetResponse, error) {
	return r.FakeGetAll(amount)
}

func (r *aboutWriterServiceMock) AddOne(dto PostRequest) (string, error) {
	return r.FakeAddOne(dto)
}

func TestPost_Returns201AndIdWhenSuccessful(t *testing.T) {
	server := gin.Default()
	mockId := primitive.NewObjectID().Hex()
	mockService := aboutWriterServiceMock{
		FakeAddOne: func(dto PostRequest) (string, error) {
			return mockId, nil
		},
	}
	mockDto := PostRequest{
		FirstName:  "Nguyen",
		LastName:   "Tran",
		Email:      "ttn.nguyen42@gmail.com",
		City:       "Hanoi",
		Country:    "Vietnam",
		Occupation: "Software Engineer Intern",
	}

	server.POST("/api/v1/test/about", Post(&mockService))
	bodyBytes, err := json.Marshal(&mockDto)
	if err != nil {
		t.Errorf("Test is broken, error is %v", err.Error())
	}
	bodyReader := bytes.NewReader(bodyBytes)
	req, _ := http.NewRequest("POST", "/api/v1/test/about", bodyReader)
	recorder := httptest.NewRecorder()
	server.ServeHTTP(recorder, req)

	responseBytes, _ := ioutil.ReadAll(recorder.Body)
	if recorder.Result().StatusCode != http.StatusCreated {
		t.Errorf("Expected status code of %v, got %v", http.StatusCreated, recorder.Result().StatusCode)
	}

	var gotResponse v1.IdResponse
	err = json.Unmarshal(responseBytes, &gotResponse)
	if err != nil {
		t.Errorf("Expected IdResponse but got \n%v", string(responseBytes))
	}
	if gotResponse.Id != mockId {
		t.Errorf("Expected ID %v, got %v", mockId, gotResponse.Id)
	}
}

func TestPost_Returns400WhenRequestIsInvalid(t *testing.T) {
	server := gin.Default()
	mockId := primitive.NewObjectID().Hex()
	mockService := aboutWriterServiceMock{
		FakeAddOne: func(dto PostRequest) (string, error) {
			return mockId, nil
		},
	}
	mockDto := v1.IdResponse{
		Id: mockId,
	}

	server.POST("/api/v1/test/about", Post(&mockService))
	bodyBytes, err := json.Marshal(&mockDto)
	if err != nil {
		t.Errorf("Test is broken, error is %v", err.Error())
	}
	bodyReader := bytes.NewReader(bodyBytes)
	req, _ := http.NewRequest("POST", "/api/v1/test/about", bodyReader)
	recorder := httptest.NewRecorder()
	server.ServeHTTP(recorder, req)

	responseBytes, _ := ioutil.ReadAll(recorder.Body)
	if recorder.Result().StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status code of %v, got %v", http.StatusInternalServerError, recorder.Result().StatusCode)
	}

	var gotResponse v1.MessageResponse
	err = json.Unmarshal(responseBytes, &gotResponse)
	if err != nil {
		t.Errorf("Expected MessageResponse but got \n%v", string(responseBytes))
	}
}

func TestPost_Returns500WhenRandomErrorOccurred(t *testing.T) {
	server := gin.Default()
	mockService := aboutWriterServiceMock{
		FakeAddOne: func(dto PostRequest) (string, error) {
			return "", repositories.NewInternalError("Random error")
		},
	}
	mockDto := PostRequest{
		FirstName:  "Nguyen",
		LastName:   "Tran",
		Email:      "ttn.nguyen42@gmail.com",
		City:       "Hanoi",
		Country:    "Vietnam",
		Occupation: "Software Engineer Intern",
	}

	server.POST("/api/v1/test/about", Post(&mockService))
	bodyBytes, err := json.Marshal(&mockDto)
	if err != nil {
		t.Errorf("Test is broken, error is %v", err.Error())
	}
	bodyReader := bytes.NewReader(bodyBytes)
	req, _ := http.NewRequest("POST", "/api/v1/test/about", bodyReader)
	recorder := httptest.NewRecorder()
	server.ServeHTTP(recorder, req)

	responseBytes, _ := ioutil.ReadAll(recorder.Body)
	if recorder.Result().StatusCode != http.StatusInternalServerError {
		t.Errorf("Expected status code of %v, got %v", http.StatusInternalServerError, recorder.Result().StatusCode)
	}

	var gotResponse v1.MessageResponse
	err = json.Unmarshal(responseBytes, &gotResponse)
	if err != nil {
		t.Errorf("Expected MessageResponse but got \n%v", string(responseBytes))
	}
}

func TestGetAll_Returns500WhenUnknownError(t *testing.T) {
	server := gin.Default()
	mockErrorMessage := v1.MessageResponse{
		Message: http.StatusText(http.StatusInternalServerError),
	}
	mockService := aboutReaderServiceMock{
		FakeGetAll: func(amount int) ([]GetResponse, error) {
			return []GetResponse{}, repositories.NewInternalError("Random error occured")
		},
	}
	server.GET("/api/v1/test/about", GetAll(&mockService))
	req, _ := http.NewRequest("GET", "/api/v1/test/about", nil)
	recorder := httptest.NewRecorder()
	server.ServeHTTP(recorder, req)

	responseBytes, _ := ioutil.ReadAll(recorder.Body)
	gotResponse := string(responseBytes)
	if recorder.Result().StatusCode != http.StatusInternalServerError {
		t.Errorf("Expected status code of %v, got %v", http.StatusInternalServerError, recorder.Result().StatusCode)
	}
	mockResponseString, err := json.Marshal(mockErrorMessage)
	if err != nil {
		t.Errorf("Test is broken, error is %v", err)
	}
	if gotResponse != string(mockResponseString) {
		t.Errorf("Expected \n %v, got \n %v", string(mockResponseString), gotResponse)
	}
}

func TestGetAll_Returns200With10DataEntryWhenNotSupplyAmount(t *testing.T) {
	server := gin.Default()
	mockId := primitive.NewObjectID().Hex()
	mockResult := []GetResponse{}
	for i := 0; i < 20; i += 1 {
		mockResult = append(mockResult, GetResponse{mockId, "Nguyen", "Tran", "ttn.nguyen42@gmail.com", "Hanoi", "Vietnam", "Software Engineer Intern"})
	}
	mockService := aboutReaderServiceMock{
		FakeGetAll: func(amount int) ([]GetResponse, error) {
			return mockResult[0:10], nil
		},
	}
	server.GET("/api/v1/test/about", GetAll(&mockService))
	req, _ := http.NewRequest("GET", "/api/v1/test/about", nil)
	recorder := httptest.NewRecorder()
	server.ServeHTTP(recorder, req)

	responseBytes, _ := ioutil.ReadAll(recorder.Body)
	if recorder.Result().StatusCode != http.StatusOK {
		t.Errorf("Expected status code of %v, got %v", http.StatusOK, recorder.Result().StatusCode)
	}
	var gotResponse []GetResponse
	err := json.Unmarshal(responseBytes, &gotResponse)
	if err != nil {
		t.Errorf("Test is broken, error is %v", err)
	}
	if len(gotResponse) != 10 {
		t.Errorf("Expected \n %v, got \n %v", 10, len(gotResponse))
	}
}

func TestGetAll_Returns200WithSpecifiedAmount(t *testing.T) {
	server := gin.Default()
	mockId := primitive.NewObjectID().Hex()
	length := 12
	mockResult := []GetResponse{}
	for i := 0; i < 20; i += 1 {
		mockResult = append(mockResult, GetResponse{mockId, "Nguyen", "Tran", "ttn.nguyen42@gmail.com", "Hanoi", "Vietnam", "Software Engineer Intern"})
	}
	mockService := aboutReaderServiceMock{
		FakeGetAll: func(amount int) ([]GetResponse, error) {
			return mockResult[0:length], nil
		},
	}
	server.GET("/api/v1/test/about", GetAll(&mockService))
	req, _ := http.NewRequest("GET", fmt.Sprintf("/api/v1/test/about?amount=%v", length), nil)
	recorder := httptest.NewRecorder()
	server.ServeHTTP(recorder, req)

	responseBytes, _ := ioutil.ReadAll(recorder.Body)
	if recorder.Result().StatusCode != http.StatusOK {
		t.Errorf("Expected status code of %v, got %v", http.StatusOK, recorder.Result().StatusCode)
	}
	var gotResponse []GetResponse
	err := json.Unmarshal(responseBytes, &gotResponse)
	if err != nil {
		t.Errorf("Test is broken, error is %v", err)
	}
	if len(gotResponse) != length {
		t.Errorf("Expected \n %v, got \n %v", length, len(gotResponse))
	}
}

func TestGetAll_Returns400WhenSupplyIncorrectAmount(t *testing.T) {
	server := gin.Default()
	amount := "not-a-correct-query"
	mockService := aboutReaderServiceMock{
		FakeGetAll: func(amount int) ([]GetResponse, error) {
			return []GetResponse{}, repositories.NewInternalError("Random error occured")
		},
	}
	server.GET("/api/v1/test/about", GetAll(&mockService))
	req, _ := http.NewRequest("GET", fmt.Sprintf("/api/v1/test/about?amount=%v", amount), nil)
	recorder := httptest.NewRecorder()
	server.ServeHTTP(recorder, req)

	if recorder.Result().StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status code of %v, got %v", http.StatusBadRequest, recorder.Result().StatusCode)
	}
}

func TestGetOne_Returns200WhenSuccessfullyFoundData(t *testing.T) {
	server := gin.Default()
	mockId := primitive.NewObjectID().Hex()
	mockDto := GetResponse{
		FirstName:  "Nguyen",
		LastName:   "Tran",
		Email:      "ttn.nguyen42@gmail.com",
		City:       "Hanoi",
		Country:    "Vietnam",
		Occupation: "Software Engineer Intern",
	}
	mockService := aboutReaderServiceMock{
		FakeGetOne: func(dto *GetResponse, id string) error {
			*dto = mockDto
			return nil
		},
	}
	server.GET("/api/v1/test/about/:id", GetOne(&mockService))
	req, _ := http.NewRequest("GET", fmt.Sprintf("/api/v1/test/about/%v", mockId), nil)
	recorder := httptest.NewRecorder()
	server.ServeHTTP(recorder, req)

	responseBytes, _ := ioutil.ReadAll(recorder.Body)
	gotResponse := string(responseBytes)
	if recorder.Result().StatusCode != http.StatusOK {
		t.Errorf("Expected status code of %v, got %v", http.StatusOK, recorder.Result().StatusCode)
	}
	mockStringResponse, err := json.Marshal(&mockDto)
	if err != nil {
		t.Errorf("Test is broken, error %v", mockStringResponse)
	}
	if string(mockStringResponse) != gotResponse {
		t.Errorf("Expected \n %v, got \n %v", string(mockStringResponse), gotResponse)
	}
}

func TestGetOne_Returns404WhenIdNotFound(t *testing.T) {
	server := gin.Default()
	mockId := "not-a-correct-ID"
	mockError := v1.MessageResponse{
		Message: http.StatusText(http.StatusNotFound),
	}
	mockService := aboutReaderServiceMock{
		FakeGetOne: func(dto *GetResponse, id string) error {
			return repositories.NewNotFoundError("Not found")
		},
	}
	server.GET("/api/v1/test/about/:id", GetOne(&mockService))
	req, _ := http.NewRequest("GET", fmt.Sprintf("/api/v1/test/about/%v", mockId), nil)
	recorder := httptest.NewRecorder()
	server.ServeHTTP(recorder, req)

	responseBytes, _ := ioutil.ReadAll(recorder.Body)
	gotResponse := string(responseBytes)

	if recorder.Result().StatusCode != http.StatusNotFound {
		t.Errorf("Expected status code of %v, got %v", http.StatusNotFound, recorder.Result().StatusCode)
	}
	mockStringResponse, err := json.Marshal(mockError)
	if err != nil {
		t.Errorf("Test is broken, error is %v", err.Error())
	}
	if string(mockStringResponse) != gotResponse {
		t.Errorf("Expected \n %v, got \n %v", string(mockStringResponse), gotResponse)
	}
}

func TestGetOne_Returns400WhenIdIsNotInCorrectFormat(t *testing.T) {
	server := gin.Default()
	mockId := "not-a-correct-ID"
	mockError := v1.MessageResponse{
		Message: "Invalid ID format",
	}
	mockService := aboutReaderServiceMock{
		FakeGetOne: func(dto *GetResponse, id string) error {
			return repositories.NewInvalidIdError("ID has incorrect format")
		},
	}
	server.GET("/api/v1/test/about/:id", GetOne(&mockService))
	req, _ := http.NewRequest("GET", fmt.Sprintf("/api/v1/test/about/%v", mockId), nil)
	recorder := httptest.NewRecorder()
	server.ServeHTTP(recorder, req)
	responseBytes, _ := ioutil.ReadAll(recorder.Body)
	gotResponse := string(responseBytes)

	if recorder.Result().StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status code of %v, got %v", http.StatusBadRequest, recorder.Result().StatusCode)
	}
	mockStringResponse, err := json.Marshal(mockError)
	if err != nil {
		t.Errorf("Test is broken, error is %v", err.Error())
	}
	if string(mockStringResponse) != gotResponse {
		t.Errorf("Expected \n %v, got \n %v", string(mockStringResponse), gotResponse)
	}
}

func TestGetOne_Returns500OnUnknownError(t *testing.T) {
	server := gin.Default()
	mockId := "not-a-correct-ID"
	mockError := v1.MessageResponse{
		Message: http.StatusText(http.StatusInternalServerError),
	}
	mockService := aboutReaderServiceMock{
		FakeGetOne: func(dto *GetResponse, id string) error {
			return errors.New("Unknown error")
		},
	}
	server.GET("/api/v1/test/about/:id", GetOne(&mockService))
	req, _ := http.NewRequest("GET", fmt.Sprintf("/api/v1/test/about/%v", mockId), nil)
	recorder := httptest.NewRecorder()
	server.ServeHTTP(recorder, req)

	responseBytes, _ := ioutil.ReadAll(recorder.Body)
	gotResponse := string(responseBytes)

	if recorder.Result().StatusCode != http.StatusInternalServerError {
		t.Errorf("Expected status code of %v, got %v", http.StatusInternalServerError, recorder.Result().StatusCode)
	}
	mockStringResponse, err := json.Marshal(mockError)
	if err != nil {
		t.Errorf("Test is broken, error is %v", err.Error())
	}
	if string(mockStringResponse) != gotResponse {
		t.Errorf("Expected \n %v, got \n %v", string(mockStringResponse), gotResponse)
	}
}
