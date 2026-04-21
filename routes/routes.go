package routes

import (
	"golangTest/adapter/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter(g *gin.Engine, taskHandler handler.TaskHandler) {
	taskRoutes := g.Group("/tasks")
	{
		taskRoutes.POST("/", taskHandler.CreateTask)
		taskRoutes.GET("/", taskHandler.GetTasks)
		taskRoutes.GET("/:id", taskHandler.GetATask)
		taskRoutes.PUT("/:id", taskHandler.UpdateTask)
		taskRoutes.PATCH("/:id/:status", taskHandler.UpdateTaskStatus)
	}
}