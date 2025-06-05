package router

import (
	"Kuraz-Tech-Challenge/handlers"

	"github.com/gin-gonic/gin"
)



func SetUpRouter(rg *gin.RouterGroup, TaskHandler *handlers.TaskHandler) {
	taskrouter := rg.Group("/tasks") 
	{
		taskrouter.POST("/", TaskHandler.CreateTask)
		taskrouter.GET("/", TaskHandler.GetTasks)
		taskrouter.PUT("/", TaskHandler.UpdateTask)
		taskrouter.DELETE("/:id", TaskHandler.DeleteTask)

	}
}