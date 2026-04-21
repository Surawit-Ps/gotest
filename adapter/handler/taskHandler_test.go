package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"golangTest/core/entity"
	"golangTest/pkg/errs"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock TaskService สำหรับ testing
type MockTaskService struct {
	mock.Mock
}

func (m *MockTaskService) CreateTask(task entity.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockTaskService) GetTasks(assignName, status string, page, limit int) ([]entity.Task, error) {
	args := m.Called(assignName, status, page, limit)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entity.Task), args.Error(1)
}

func (m *MockTaskService) GetATask(id string) (*entity.Task, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Task), args.Error(1)
}

func (m *MockTaskService) UpdateTask(id string, task entity.Task) error {
	args := m.Called(id, task)
	return args.Error(0)
}

func (m *MockTaskService) UpdateTaskStatus(id, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

// TestCreateTask - ทดสอบการสร้าง task ทั้ง success และ error cases
func TestCreateTask(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		body           interface{}
		mockError      error
		expectedStatus int
		shouldCallMock bool
	}{
		{
			name: "Success",
			body: entity.TaskRequest{
				Title:       "Test Task",
				Description: "Test Description",
				AssignName:  "John Doe",
			},
			mockError:      nil,
			expectedStatus: http.StatusCreated,
			shouldCallMock: true,
		},
		{
			name:           "InvalidInput",
			body:           "invalid json",
			mockError:      nil,
			expectedStatus: http.StatusBadRequest,
			shouldCallMock: false,
		},
		{
			name: "DatabaseError",
			body: entity.TaskRequest{
				Title:       "Test Task",
				Description: "Test Description",
				AssignName:  "John Doe",
			},
			mockError:      errs.ErrDatabase,
			expectedStatus: http.StatusInternalServerError,
			shouldCallMock: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockService := new(MockTaskService)

			if tc.shouldCallMock {
				mockService.On("CreateTask", mock.AnythingOfType("entity.Task")).Return(tc.mockError)
			}

			handler := NewTaskHandler(mockService)
			router := gin.New()
			router.POST("/tasks", handler.CreateTask)

			body, _ := json.Marshal(tc.body)
			req := httptest.NewRequest("POST", "/tasks", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedStatus, w.Code)
		})
	}
}

// TestGetTasks - ทดสอบการดึง tasks ทั้ง success และ error cases
func TestGetTasks(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		queryParams    string
		mockTasks      []entity.Task
		mockError      error
		expectedStatus int
	}{
		{
			name:        "Success_AllTasks",
			queryParams: "?page=1&limit=5",
			mockTasks: []entity.Task{
				{
					Id:          "1",
					Title:       "Task 1",
					Description: "Desc 1",
					Status:      "todo",
					AssignName:  "John",
				},
				{
					Id:          "2",
					Title:       "Task 2",
					Description: "Desc 2",
					Status:      "in-progress",
					AssignName:  "Jane",
				},
			},
			mockError:      nil,
			expectedStatus: http.StatusOK,
		},
		{
			name:        "Success_WithFilter",
			queryParams: "?status=todo&assign_name=John&page=1&limit=5",
			mockTasks: []entity.Task{
				{
					Id:          "1",
					Title:       "Task 1",
					Description: "Desc 1",
					Status:      "todo",
					AssignName:  "John",
				},
			},
			mockError:      nil,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "DatabaseError",
			queryParams:    "?page=1&limit=5",
			mockTasks:      nil,
			mockError:      errs.ErrDatabase,
			expectedStatus: http.StatusInternalServerError,
		},
		{
			name:           "InvalidPageLimit",
			queryParams:    "?page=abc&limit=xyz",
			mockTasks:      nil,
			mockError:      nil,
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockService := new(MockTaskService)
			mockService.On("GetTasks", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(tc.mockTasks, tc.mockError)

			handler := NewTaskHandler(mockService)
			router := gin.New()
			router.GET("/tasks", handler.GetTasks)

			req := httptest.NewRequest("GET", "/tasks"+tc.queryParams, nil)

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedStatus, w.Code)
		})
	}
}

// TestGetATask - ทดสอบการดึง task เดี่ยว
func TestGetATask(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		taskID         string
		mockTask       *entity.Task
		mockError      error
		expectedStatus int
	}{
		{
			name:   "Success",
			taskID: "1",
			mockTask: &entity.Task{
				Id:          "1",
				Title:       "Task 1",
				Description: "Desc 1",
				Status:      "todo",
				AssignName:  "John",
			},
			mockError:      nil,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "NotFound",
			taskID:         "999",
			mockTask:       nil,
			mockError:      errs.ErrTaskNotFound,
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "DatabaseError",
			taskID:         "1",
			mockTask:       nil,
			mockError:      errs.ErrDatabase,
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockService := new(MockTaskService)
			mockService.On("GetATask", tc.taskID).Return(tc.mockTask, tc.mockError)

			handler := NewTaskHandler(mockService)
			router := gin.New()
			router.GET("/tasks/:id", handler.GetATask)

			req := httptest.NewRequest("GET", "/tasks/"+tc.taskID, nil)

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tc.expectedStatus, w.Code)
		})
	}
}


