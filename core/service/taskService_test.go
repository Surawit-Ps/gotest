package service

import (
	"golangTest/core/entity"
	"golangTest/pkg/errs"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock TaskRepository สำหรับ testing
type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) AddTask(task entity.Task) error {
	args := m.Called(task)
	return args.Error(0)
}

func (m *MockTaskRepository) GetTasks(assignName, status string, page, limit int) ([]entity.Task, error) {
	args := m.Called(assignName, status, page, limit)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entity.Task), args.Error(1)
}

func (m *MockTaskRepository) GetATask(id string) (*entity.Task, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Task), args.Error(1)
}

func (m *MockTaskRepository) EditTask(id string, task entity.Task) error {
	args := m.Called(id, task)
	return args.Error(0)
}

func (m *MockTaskRepository) EditTaskStatus(id, status string) error {
	args := m.Called(id, status)
	return args.Error(0)
}

// TestCreateTask - ทดสอบการสร้าง task ทั้ง success และ error cases
func TestCreateTask(t *testing.T) {
	tests := []struct {
		name      string
		task      entity.Task
		mockError error
		wantErr   bool
	}{
		{
			name: "Success",
			task: entity.Task{
				Title:       "Test Task",
				Description: "Test Description",
				AssignName:  "John Doe",
			},
			mockError: nil,
			wantErr:   false,
		},
		{
			name: "DatabaseError",
			task: entity.Task{
				Title:       "Test Task",
				Description: "Test Description",
				AssignName:  "John Doe",
			},
			mockError: errs.ErrDatabase,
			wantErr:   true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(MockTaskRepository)
			mockRepo.On("AddTask", tc.task).Return(tc.mockError)

			service := NewTaskService(mockRepo)
			err := service.CreateTask(tc.task)

			if tc.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tc.mockError, err)
			} else {
				assert.NoError(t, err)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}

// TestGetTasks - ทดสอบการดึง tasks ทั้ง success และ error cases
func TestGetTasks(t *testing.T) {
	tests := []struct {
		name       string
		assignName string
		status     string
		page       int
		limit      int
		mockTasks  []entity.Task
		mockError  error
		wantErr    bool
	}{
		{
			name:       "Success_AllTasks",
			assignName: "",
			status:     "",
			page:       1,
			limit:      5,
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
			mockError: nil,
			wantErr:   false,
		},
		{
			name:       "Success_WithFilter",
			assignName: "John",
			status:     "todo",
			page:       1,
			limit:      5,
			mockTasks: []entity.Task{
				{
					Id:          "1",
					Title:       "Task 1",
					Description: "Desc 1",
					Status:      "todo",
					AssignName:  "John",
				},
			},
			mockError: nil,
			wantErr:   false,
		},
		{
			name:       "DatabaseError",
			assignName: "",
			status:     "",
			page:       1,
			limit:      5,
			mockTasks:  nil,
			mockError:  errs.ErrDatabase,
			wantErr:    true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(MockTaskRepository)
			mockRepo.On("GetTasks", tc.assignName, tc.status, tc.page, tc.limit).Return(tc.mockTasks, tc.mockError)

			service := NewTaskService(mockRepo)
			tasks, err := service.GetTasks(tc.assignName, tc.status, tc.page, tc.limit)

			if tc.wantErr {
				assert.Error(t, err)
				assert.Nil(t, tasks)
				assert.Equal(t, tc.mockError, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.mockTasks, tasks)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}

// TestGetATask - ทดสอบการดึง task เดี่ยว ทั้ง success และ error cases
func TestGetATask(t *testing.T) {
	tests := []struct {
		name      string
		taskID    string
		mockTask  *entity.Task
		mockError error
		wantErr   bool
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
			mockError: nil,
			wantErr:   false,
		},
		{
			name:      "NotFound",
			taskID:    "999",
			mockTask:  nil,
			mockError: errs.ErrTaskNotFound,
			wantErr:   true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockRepo := new(MockTaskRepository)
			mockRepo.On("GetATask", tc.taskID).Return(tc.mockTask, tc.mockError)

			service := NewTaskService(mockRepo)
			task, err := service.GetATask(tc.taskID)

			if tc.wantErr {
				assert.Error(t, err)
				assert.Nil(t, task)
				assert.Equal(t, tc.mockError, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.mockTask, task)
			}
			mockRepo.AssertExpectations(t)
		})
	}
}

// TestUpdateTask - ทดสอบการอัปเดต task ทั้ง success และ error cases
