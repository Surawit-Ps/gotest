package main

import (
	"golangTest/adapter/handler"
	"golangTest/adapter/repository"
	"golangTest/core/entity"
	"golangTest/core/port"
	"golangTest/core/service"
	"golangTest/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&repository.Task{})

	taskRepo := repository.NewTaskRepositoryDB(db)
	taskService := service.NewTaskService(taskRepo)
	tasks, err := taskService.GetTasks("", "", 1, 1)
	if err == nil && len(tasks) == 0 {
		addTestData(taskService)
	}
	taskHandler := handler.NewTaskHandler(taskService)

	// Set up Gin router and routes
	router := gin.Default()
	routes.SetupRouter(router, taskHandler)
	// Start the server
	router.Run(":8080")
}

func addTestData(taskService port.TaskServicePort) {
	task := []entity.Task{
		{Title: "Task 1", Description: "Description for Task 1", Status: "pending", AssignName: "Alice"},
		{Title: "Task 2", Description: "Description for Task 2", Status: "in_progress", AssignName: "Bob"},
		{Title: "Task 3", Description: "Description for Task 3", Status: "done", AssignName: "Charlie"},
		{Title: "Task 4", Description: "Description for Task 4", Status: "pending", AssignName: "David"},
		{Title: "Task 5", Description: "Description for Task 5", Status: "in_progress", AssignName: "Eve"},
		{Title: "Task 6", Description: "Description for Task 6", Status: "done", AssignName: "Frank"},
		{Title: "Task 7", Description: "Description for Task 7", Status: "pending", AssignName: "Grace"},
		{Title: "Task 8", Description: "Description for Task 8", Status: "in_progress", AssignName: "Heidi"},
		{Title: "Task 9", Description: "Description for Task 9", Status: "done", AssignName: "Ivan"},
		{Title: "Task 10", Description: "Description for Task 10", Status: "pending", AssignName: "Judy"},
		{Title: "Task 11", Description: "Description for Task 11", Status: "in_progress", AssignName: "Karl"},
		{Title: "Task 12", Description: "Description for Task 12", Status: "done", AssignName: "Leo"},
		{Title: "Task 13", Description: "Description for Task 13", Status: "pending", AssignName: "Mallory"},
		{Title: "Task 14", Description: "Description for Task 14", Status: "in_progress", AssignName: "Nina"},
		{Title: "Task 15", Description: "Description for Task 15", Status: "done", AssignName: "Oscar"},
		{Title: "Task 16", Description: "Description for Task 16", Status: "pending", AssignName: "Peggy"},
		{Title: "Task 17", Description: "Description for Task 17", Status: "in_progress", AssignName: "Quentin"},
		{Title: "Task 18", Description: "Description for Task 18", Status: "done", AssignName: "Ruth"},
		{Title: "Task 19", Description: "Description for Task 19", Status: "pending", AssignName: "Sybil"},
		{Title: "Task 20", Description: "Description for Task 20", Status: "in_progress", AssignName: "Trent"},
		{Title: "Task 21", Description: "Description for Task 21", Status: "done", AssignName: "Uma"},
		{Title: "Task 22", Description: "Description for Task 22", Status: "pending", AssignName: "Victor"},
		{Title: "Task 23", Description: "Description for Task 23", Status: "in_progress", AssignName: "Wendy"},
		{Title: "Task 24", Description: "Description for Task 24", Status: "done", AssignName: "Xavier"},
		{Title: "Task 25", Description: "Description for Task 25", Status: "pending", AssignName: "Yvonne"},
		{Title: "Task 26", Description: "Description for Task 26", Status: "in_progress", AssignName: "Zack"},
		{Title: "Task 27", Description: "Description for Task 27", Status: "done", AssignName: "Alice"},
		{Title: "Task 28", Description: "Description for Task 28", Status: "pending", AssignName: "Bob"},
		{Title: "Task 29", Description: "Description for Task 29", Status: "in_progress", AssignName: "Charlie"},
		{Title: "Task 30", Description: "Description for Task 30", Status: "done", AssignName: "David"},
		{Title: "Task 31", Description: "Description for Task 31", Status: "pending", AssignName: "Eve"},
		{Title: "Task 32", Description: "Description for Task 32", Status: "in_progress", AssignName: "Frank"},
		{Title: "Task 33", Description: "Description for Task 33", Status: "done", AssignName: "Grace"},
		{Title: "Task 34", Description: "Description for Task 34", Status: "pending", AssignName: "Heidi"},
		{Title: "Task 35", Description: "Description for Task 35", Status: "in_progress", AssignName: "Ivan"},
		{Title: "Task 36", Description: "Description for Task 36", Status: "done", AssignName: "Judy"},
		{Title: "Task 37", Description: "Description for Task 37", Status: "pending", AssignName: "Karl"},
		{Title: "Task 38", Description: "Description for Task 38", Status: "in_progress", AssignName: "Leo"},
		{Title: "Task 39", Description: "Description for Task 39", Status: "done", AssignName: "Mallory"},
	}
	for _, t := range task {
		err := taskService.CreateTask(t)
		if err != nil {
			panic("failed to create task")
		}
	}
}
