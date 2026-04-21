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
		panic(err)
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
		{Title: "Task 1", Description: "Description for Task 1", AssignName: "Alice"},
		{Title: "Task 2", Description: "Description for Task 2", AssignName: "Bob"},
		{Title: "Task 3", Description: "Description for Task 3", AssignName: "Charlie"},
		{Title: "Task 4", Description: "Description for Task 4", AssignName: "David"},
		{Title: "Task 5", Description: "Description for Task 5", AssignName: "Eve"},
		{Title: "Task 6", Description: "Description for Task 6", AssignName: "Frank"},
		{Title: "Task 7", Description: "Description for Task 7", AssignName: "Grace"},
		{Title: "Task 8", Description: "Description for Task 8", AssignName: "Heidi"},
		{Title: "Task 9", Description: "Description for Task 9", AssignName: "Ivan"},
		{Title: "Task 10", Description: "Description for Task 10", AssignName: "Judy"},
		{Title: "Task 11", Description: "Description for Task 11", AssignName: "Karl"},
		{Title: "Task 12", Description: "Description for Task 12", AssignName: "Leo"},
		{Title: "Task 13", Description: "Description for Task 13", AssignName: "Mallory"},
		{Title: "Task 14", Description: "Description for Task 14", AssignName: "Nina"},
		{Title: "Task 15", Description: "Description for Task 15", AssignName: "Oscar"},
		{Title: "Task 16", Description: "Description for Task 16", AssignName: "Peggy"},
		{Title: "Task 17", Description: "Description for Task 17", AssignName: "Quentin"},
		{Title: "Task 18", Description: "Description for Task 18", AssignName: "Ruth"},
		{Title: "Task 19", Description: "Description for Task 19", AssignName: "Sybil"},
		{Title: "Task 20", Description: "Description for Task 20", AssignName: "Trent"},
		{Title: "Task 21", Description: "Description for Task 21", AssignName: "Uma"},
		{Title: "Task 22", Description: "Description for Task 22", AssignName: "Victor"},
		{Title: "Task 23", Description: "Description for Task 23", AssignName: "Wendy"},
		{Title: "Task 24", Description: "Description for Task 24", AssignName: "Xavier"},
		{Title: "Task 25", Description: "Description for Task 25", AssignName: "Yvonne"},
		{Title: "Task 26", Description: "Description for Task 26", AssignName: "Zack"},
		{Title: "Task 27", Description: "Description for Task 27", AssignName: "Alice"},
		{Title: "Task 28", Description: "Description for Task 28", AssignName: "Bob"},
		{Title: "Task 29", Description: "Description for Task 29", AssignName: "Charlie"},
		{Title: "Task 30", Description: "Description for Task 30", AssignName: "David"},
		{Title: "Task 31", Description: "Description for Task 31", AssignName: "Eve"},
		{Title: "Task 32", Description: "Description for Task 32", AssignName: "Frank"},
		{Title: "Task 33", Description: "Description for Task 33", AssignName: "Grace"},
		{Title: "Task 34", Description: "Description for Task 34", AssignName: "Heidi"},
		{Title: "Task 35", Description: "Description for Task 35", AssignName: "Ivan"},
		{Title: "Task 36", Description: "Description for Task 36", AssignName: "Judy"},
		{Title: "Task 37", Description: "Description for Task 37", AssignName: "Karl"},
		{Title: "Task 38", Description: "Description for Task 38", AssignName: "Leo"},
		{Title: "Task 39", Description: "Description for Task 39", AssignName: "Mallory"},
	}
	for _, t := range task {
		err := taskService.CreateTask(t)
		if err != nil {
			panic("failed to create task")
		}
	}
}
