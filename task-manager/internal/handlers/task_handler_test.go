package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"task-manager/internal/models"
	"task-manager/internal/storage"
	"testing"
)

// mocking task storage implements task-storage for testing
type MockTaskStorage struct {
	mock.Mock
}

func (m *MockTaskStorage) Create(task *models.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockTaskStorage) GetByID(id string) (task *models.Task, err error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Task), args.Error(1)
}

func (m *MockTaskStorage) GetAll() ([]*models.Task, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*models.Task), args.Error(1)
}

func (m *MockTaskStorage) Update(id string, update *models.TaskUpdate) (task *models.Task, err error) {
	args := m.Called(id, update)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Task), args.Error(1)
}

func (m *MockTaskStorage) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockTaskStorage) GetByStatus(status string) ([]*models.Task, error) {
	args := m.Called(status)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*models.Task), args.Error(1)
}

// test  handlers and helper functions
func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	return router
}
func setupTestHandler() (*TaskHandler, *MockTaskStorage) {
	mockStorage := new(MockTaskStorage)
	handler := NewTaskHandler(mockStorage)
	return handler, mockStorage
}

func performRequest(router http.Handler, method, path string, body interface{}) *httptest.ResponseRecorder {
	var reqBody []byte
	if body != nil {
		reqBody, _ = json.Marshal(body)
	}
	req, _ := http.NewRequest(method, path, bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}

func stringPtr(s string) *string {
	return &s
}
