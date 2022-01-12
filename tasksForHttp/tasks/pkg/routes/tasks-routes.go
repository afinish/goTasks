package routes

import (
	"github.com/heavenmise/goTasks/tasksForHttp/tasks/pkg/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterNewRouter() *gin.Engine {
	r := gin.Default()
	v0 := r.Group("/")
	{
		v0.GET("tasks", controllers.GetTasks)
		v0.GET("task", controllers.GetTask)
		v0.POST("task", controllers.CreateTask)
		v0.PUT("task", controllers.UpdateTask)
		v0.DELETE("task", controllers.DeleteTask)
	}
	return r
}