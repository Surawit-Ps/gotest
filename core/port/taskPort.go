package port

import (
	"golangTest/core/entity"
)

type TaskRepositoryPort interface {
	AddTask(entity.Task) error
	GetTasks(string,string,int,int) ([]entity.Task, error)
	GetATask(string) (*entity.Task, error)
	EditTask(string, entity.Task) error
	EditTaskStatus(string, string) error
	RemoveTask(string) error
}

type TaskServicePort interface {
	CreateTask(entity.Task) error
	GetTasks(string, string, int, int) ([]entity.Task, error)
	GetATask(string) (*entity.Task, error)
	UpdateTask(string, entity.Task) error
	UpdateTaskStatus(string, string) error
	DeleteTask(string) error
}
