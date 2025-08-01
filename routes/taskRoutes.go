package routes

import (
	"github.com/gin-gonic/gin"
	"task-api/controllers"
)

func SetupTaskRoutes(r *gin.Engine) {
	tasks := r.Group("/tasks")
	{
		tasks.POST("/", controllers.CreateTask)
		tasks.GET("/", controllers.GetTasks)
		tasks.GET("/:id", controllers.GetTaskById)
		tasks.PUT("/:id", controllers.UpdateTask)
		tasks.DELETE("/:id", controllers.DeleteTask)
	}
}
