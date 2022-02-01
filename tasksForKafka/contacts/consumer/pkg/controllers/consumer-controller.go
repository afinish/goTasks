package controllers

import (
	"net/http"

	"github.com/afinish/goTasks/tasksForKafka/contacts/consumer/pkg/models"
	"github.com/gin-gonic/gin"
)

func GetContacts(c *gin.Context) {
	contacts := models.GetContacts()
	c.JSON(http.StatusOK, gin.H{
		"data": contacts,
	})
}

func GetContact(c *gin.Context) {
	var id string
	query := c.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value)-1]
		if key == "id" {
			id = queryValue
		}
	}
	contact, err := models.GetContact(id)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": contact,
	})
}
