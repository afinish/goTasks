package routes

import (
	"github.com/afinish/goTasks/tasksForKafka/contacts/consumer/pkg/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterNewRouter() *gin.Engine {
	r := gin.Default()
	v0 := r.Group("/")
	{
		v0.GET("contacts", controllers.GetContacts)
		v0.GET("contact", controllers.GetContact)
	}
	return r
}
