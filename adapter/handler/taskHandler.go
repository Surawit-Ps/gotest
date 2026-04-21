package handler

import (

	"golangTest/core/entity"
	"golangTest/core/port"
	"strconv"
	e "golangTest/pkg/errs"
	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	taskService port.TaskServicePort
}

func NewTaskHandler(taskService port.TaskServicePort) TaskHandler {
	return TaskHandler{taskService: taskService}
}


func (h *TaskHandler) CreateTask(c *gin.Context) {
	var task entity.TaskRequest
	if err := c.ShouldBindJSON(&task); err != nil {
		handleError(c,e.ErrInvalidInput)
		return
	}
	taskEntity := entity.Task{
		Title: task.Title,
		Description: task.Description,
		AssignName: task.AssignName,
	}
	err := h.taskService.CreateTask(taskEntity)
	if err != nil {
		handleError(c, err)
		return 
	}
	ResponseCreateSuccess(c,"task created successfully", task)
}

func (h *TaskHandler) GetTasks(c *gin.Context) {
	status := c.Query("status")
	assign_name := c.Query("assign_name")
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "5")
	pageInt, _ := strconv.Atoi(page)
	if pageInt <= 0 {
		handleError(c, e.ErrInvalidInput)
		return
	}
	limitInt, _ := strconv.Atoi(limit)
	if limitInt <= 0 {
		handleError(c, e.ErrInvalidInput)
		return
	}
	tasks, err := h.taskService.GetTasks(status, assign_name, pageInt, limitInt)	
	if err != nil {
		handleError(c, err)
		return 
	}
	ResponseSuccess(c, tasks)
}

func (h *TaskHandler) GetATask(c *gin.Context) {
	id := c.Param("id")
	task, err := h.taskService.GetATask(id)
	if err != nil {
		handleError(c, err)
		return
	}
	ResponseSuccess(c, task)
}

func (h *TaskHandler) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var task entity.TaskRequest
	if err := c.ShouldBindJSON(&task); err != nil {
		handleError(c, err)
		return
	}

	TaskEntity := entity.Task{
		Title: task.Title,
		Description: task.Description,
		AssignName: task.AssignName,
	}
	err := h.taskService.UpdateTask(id, TaskEntity)
	if err != nil {
		handleError(c, err)
		return
	}
	ResponseSuccess(c, "Task updated successfully")
}

func (h *TaskHandler) UpdateTaskStatus(c *gin.Context) {
	id := c.Param("id")
	status := c.Param("status")
	err := h.taskService.UpdateTaskStatus(id, status)
	if err != nil {
		handleError(c, err)
		return
	}
	ResponseSuccess(c, "Task status updated successfully")
}