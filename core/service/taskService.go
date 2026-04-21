package service

import (
	"golangTest/core/entity"
	"golangTest/core/port"
	e "golangTest/pkg/errs"
)

type taskService struct {
	taskRepo port.TaskRepositoryPort
}

func NewTaskService(taskRepo port.TaskRepositoryPort) *taskService {
	return &taskService{taskRepo: taskRepo}
}

func (s *taskService) CreateTask(task entity.Task) error {
	return s.taskRepo.AddTask(task)
}

func (s *taskService) GetTasks(assign_name string, status string, page int, limit int) ([]entity.Task, error) {
	return s.taskRepo.GetTasks(assign_name, status, page, limit)
}

func (s *taskService) GetATask(id string) (*entity.Task, error) {
	return s.taskRepo.GetATask(id)
}

func (s *taskService) UpdateTask(id string, task entity.Task) error {
	return s.taskRepo.EditTask(id, task)
}

func (s *taskService) UpdateTaskStatus(id string, status string) error {
	oldTask, err := s.taskRepo.GetATask(id)
	if err != nil {
		return err
	}
	if oldTask.Status == "done" {
		return e.ErrStatusUnchanged
	} else if oldTask.Status == status {
		return e.ErrStatusUnchanged
	}
	return s.taskRepo.EditTaskStatus(id, status)
}

func (s *taskService) DeleteTask(id string) error {
	return s.taskRepo.RemoveTask(id)
}
