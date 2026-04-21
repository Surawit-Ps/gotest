package repository

import (
	"fmt"
	"golangTest/core/entity"
	e "golangTest/pkg/errs"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type taskRepositoryDB struct {
	db *gorm.DB
}

func NewTaskRepositoryDB(db *gorm.DB) *taskRepositoryDB {
	return &taskRepositoryDB{db: db}
}

type Task struct {
	Id          string `gorm:"primaryKey"`
	Title       string
	Description string
	Status      string
	AssignName  string
	CreateAt	 time.Time
	UpdateAt	 time.Time
}

func EntityToModel(task entity.Task) Task {
	return Task{
		Id:          task.Id,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		AssignName:  task.AssignName,
		CreateAt:	 task.CreatedAt  ,
		UpdateAt:	 task.UpdateAt,
	}
}

func ModelToEntity(task Task) entity.Task {
	return entity.Task{
		Id:          task.Id,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		AssignName:  task.AssignName,
	}
}

func (r taskRepositoryDB) AddTask(task entity.Task) error {
	task.Id = uuid.New().String()
	task.CreatedAt = time.Now()
	task.UpdateAt = time.Now()
	task.Status = "todo"
	taskModel := EntityToModel(task)
	err := r.db.Create(&taskModel).Error
	if err != nil {
		return e.ErrDatabase
	}
	return nil
}

func (r taskRepositoryDB) GetTasks(assign_name string,status string, page int, limit int) ([]entity.Task, error) {
	var tasks []Task
	query := r.db.Model(&Task{})
	fmt.Print(query)
	if assign_name != "" {
		query = query.Where("assign_name = ?", assign_name)
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}
	err := query.Offset((page - 1) * limit).Limit(limit).Find(&tasks).Error
	if err != nil {
		return nil, e.ErrDatabase
	}
	var taskEntities []entity.Task
	for _, task := range tasks {
		taskEntities = append(taskEntities, ModelToEntity(task))
	}
	return taskEntities, nil

}

func (r taskRepositoryDB) GetATask(id string) (*entity.Task, error) {
	var task Task
	err := r.db.First(&task, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, e.ErrTaskNotFound
		}
		return nil, e.ErrDatabase
	}
	taskEntity := ModelToEntity(task)
	return &taskEntity, nil
}

func (r taskRepositoryDB) EditTask(id string, task entity.Task) error {
	var editTask Task
	err := r.db.First(&editTask, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return e.ErrTaskNotFound
		}
		return e.ErrDatabase
	}
	editTask.Title = task.Title
	editTask.Description = task.Description
	editTask.AssignName = task.AssignName
	editTask.UpdateAt = time.Now()
	err = r.db.Save(&editTask).Error
	if err != nil {
		return e.ErrDatabase
	}
	return nil
}

func (r taskRepositoryDB) EditTaskStatus(id string, status string) error {
	var task Task
	err := r.db.First(&task, "id = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return e.ErrTaskNotFound
		}
		return e.ErrDatabase
	}
	task.Status = status
	task.UpdateAt = time.Now()
	err = r.db.Save(&task).Error
	if err != nil {
		return e.ErrDatabase
	}
	return nil
}