package repository

import (
	"testing"
	"time"

	"golangTest/core/entity"
	"golangTest/pkg/errs"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// setupTestDB - สร้าง in-memory database สำหรับ testing
func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	// Migrate the schema
	err = db.AutoMigrate(&Task{})
	assert.NoError(t, err)

	return db
}

// TestAddTask - ทดสอบการเพิ่ม task ทั้ง success และ default values
func TestAddTask(t *testing.T) {
	tests := []struct {
		name    string
		task    entity.Task
		checkFn func(*testing.T, *gorm.DB, entity.Task)
	}{
		{
			name: "Success",
			task: entity.Task{
				Title:       "Test Task",
				Description: "Test Description",
				AssignName:  "John Doe",
			},
			checkFn: func(t *testing.T, db *gorm.DB, inputTask entity.Task) {
				var savedTask Task
				result := db.First(&savedTask, "title = ?", "Test Task")
				assert.NoError(t, result.Error)
				assert.Equal(t, "Test Task", savedTask.Title)
				assert.Equal(t, "Test Description", savedTask.Description)
				assert.Equal(t, "John Doe", savedTask.AssignName)
				assert.NotEmpty(t, savedTask.Id)
			},
		},
		{
			name: "SetDefaultValues",
			task: entity.Task{
				Title:       "Test Task 2",
				Description: "Test Description",
				AssignName:  "Jane Doe",
			},
			checkFn: func(t *testing.T, db *gorm.DB, inputTask entity.Task) {
				var savedTask Task
				db.First(&savedTask, "title = ?", "Test Task 2")
				assert.Equal(t, "todo", savedTask.Status)
				assert.NotZero(t, savedTask.CreateAt)
				assert.NotZero(t, savedTask.UpdateAt)
				assert.NotEmpty(t, savedTask.Id)
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			db := setupTestDB(t)
			repo := NewTaskRepositoryDB(db)

			err := repo.AddTask(tc.task)
			assert.NoError(t, err)

			tc.checkFn(t, db, tc.task)
		})
	}
}

// TestGetTasks - ทดสอบการดึง tasks ทั้ง all, filter, pagination และ empty cases
func TestGetTasks(t *testing.T) {
	tests := []struct {
		name       string
		assignName string
		status     string
		page       int
		limit      int
		setupFn    func(*gorm.DB)
		expected   int
	}{
		{
			name:       "AllTasks",
			assignName: "",
			status:     "",
			page:       1,
			limit:      10,
			setupFn: func(db *gorm.DB) {
				tasks := []Task{
					{Id: "1", Title: "Task 1", AssignName: "John", Status: "todo", CreateAt: time.Now(), UpdateAt: time.Now()},
					{Id: "2", Title: "Task 2", AssignName: "Jane", Status: "todo", CreateAt: time.Now(), UpdateAt: time.Now()},
					{Id: "3", Title: "Task 3", AssignName: "Bob", Status: "done", CreateAt: time.Now(), UpdateAt: time.Now()},
				}
				for _, t := range tasks {
					db.Create(&t)
				}
			},
			expected: 3,
		},
		{
			name:       "FilterByAssignName",
			assignName: "John",
			status:     "",
			page:       1,
			limit:      10,
			setupFn: func(db *gorm.DB) {
				tasks := []Task{
					{Id: "1", Title: "Task 1", AssignName: "John", Status: "todo", CreateAt: time.Now(), UpdateAt: time.Now()},
					{Id: "2", Title: "Task 2", AssignName: "Jane", Status: "todo", CreateAt: time.Now(), UpdateAt: time.Now()},
					{Id: "3", Title: "Task 3", AssignName: "John", Status: "done", CreateAt: time.Now(), UpdateAt: time.Now()},
				}
				for _, t := range tasks {
					db.Create(&t)
				}
			},
			expected: 2,
		},
		{
			name:       "FilterByStatus",
			assignName: "",
			status:     "todo",
			page:       1,
			limit:      10,
			setupFn: func(db *gorm.DB) {
				tasks := []Task{
					{Id: "1", Title: "Task 1", AssignName: "John", Status: "todo", CreateAt: time.Now(), UpdateAt: time.Now()},
					{Id: "2", Title: "Task 2", AssignName: "Jane", Status: "done", CreateAt: time.Now(), UpdateAt: time.Now()},
				}
				for _, t := range tasks {
					db.Create(&t)
				}
			},
			expected: 1,
		},
		{
			name:       "FilterByBoth",
			assignName: "John",
			status:     "todo",
			page:       1,
			limit:      10,
			setupFn: func(db *gorm.DB) {
				tasks := []Task{
					{Id: "1", Title: "Task 1", AssignName: "John", Status: "todo", CreateAt: time.Now(), UpdateAt: time.Now()},
					{Id: "2", Title: "Task 2", AssignName: "John", Status: "done", CreateAt: time.Now(), UpdateAt: time.Now()},
					{Id: "3", Title: "Task 3", AssignName: "Jane", Status: "todo", CreateAt: time.Now(), UpdateAt: time.Now()},
				}
				for _, t := range tasks {
					db.Create(&t)
				}
			},
			expected: 1,
		},
		{
			name:       "Pagination_Page1",
			assignName: "",
			status:     "",
			page:       1,
			limit:      5,
			setupFn: func(db *gorm.DB) {
				for i := 1; i <= 10; i++ {
					db.Create(&Task{
						Id:         string(rune(i)),
						Title:      "Task " + string(rune(i+'0')),
						AssignName: "John",
						Status:     "todo",
						CreateAt:   time.Now(),
						UpdateAt:   time.Now(),
					})
				}
			},
			expected: 5,
		},
		{
			name:       "NoResults",
			assignName: "",
			status:     "",
			page:       1,
			limit:      10,
			setupFn:    func(db *gorm.DB) {},
			expected:   0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			db := setupTestDB(t)
			tc.setupFn(db)

			repo := NewTaskRepositoryDB(db)
			result, err := repo.GetTasks(tc.assignName, tc.status, tc.page, tc.limit)

			assert.NoError(t, err)
			assert.Equal(t, tc.expected, len(result))
		})
	}
}

// TestGetATask - ทดสอบการดึง task เดี่ยว ทั้ง success และ not found
func TestGetATask(t *testing.T) {
	tests := []struct {
		name      string
		taskID    string
		setupFn   func(*gorm.DB)
		wantErr   bool
		errorType error
	}{
		{
			name:   "Success",
			taskID: "test-1",
			setupFn: func(db *gorm.DB) {
				createdTime := time.Now()
				db.Create(&Task{
					Id:          "test-1",
					Title:       "Test Task",
					Description: "Test Description",
					AssignName:  "John",
					Status:      "todo",
					CreateAt:    createdTime,
					UpdateAt:    createdTime,
				})
			},
			wantErr: false,
		},
		{
			name:      "NotFound",
			taskID:    "non-existent",
			setupFn:   func(db *gorm.DB) {},
			wantErr:   true,
			errorType: errs.ErrTaskNotFound,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			db := setupTestDB(t)
			tc.setupFn(db)

			repo := NewTaskRepositoryDB(db)
			result, err := repo.GetATask(tc.taskID)

			if tc.wantErr {
				assert.Error(t, err)
				assert.Equal(t, tc.errorType, err)
				assert.Nil(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
				assert.Equal(t, tc.taskID, result.Id)
				assert.Equal(t, "Test Task", result.Title)
			}
		})
	}
}


