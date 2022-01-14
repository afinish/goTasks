package routes

import (
	"github.com/heavenmise/goTasks/tasksForHttp/contacts/pkg/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterNewRouter() *gin.Engine {
	r := gin.Default()
	v0 := r.Group("/")
	{
		v0.GET("contacts", controllers.GetContacts)
		v0.GET("contact", controllers.GetContact)
		v0.POST("contact", controllers.CreateContact)
		v0.PUT("contact", controllers.UpdateContact)
		v0.DELETE("contact", controllers.DeleteContact)
	}
	return r
}