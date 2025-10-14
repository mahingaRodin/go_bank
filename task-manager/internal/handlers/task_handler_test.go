package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"task-manager/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	//"task-manager/internal/storage"
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

// actual test functions
func TestTaskHandler_CreateTask(t *testing.T) {
	router := setupRouter()
	handler, mockStorage := setupTestHandler()

	router.POST("/tasks", handler.CreateTask)
	t.Run("Success", func(t *testing.T) {
		task := models.Task{
			Title:       "Test Task",
			Description: "Test Description",
		}
		mockStorage.On("Create", mock.AnythingOfType("*models.Task")).Return(nil)
		w := performRequest(router, "POST", "/tasks", task)
		assert.Equal(t, http.StatusCreated, w.Code)
		mockStorage.AssertCalled(t, "Create", mock.AnythingOfType("*models.Task"))
	})

	t.Run("Invalid JSON", func(t *testing.T) {
		w := performRequest(router, "POST", "/tasks", "invalid json")

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Storage Error", func(t *testing.T) {
		task := models.Task{
			Title:       "Test Task",
			Description: "Test Description",
		}

		mockStorage.On("Create", mock.AnythingOfType("*models.Task")).Return(assert.AnError)

		w := performRequest(router, "POST", "/tasks", task)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
