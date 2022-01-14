package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/heavenmise/goTasks/tasksForHttp/contacts/pkg/models"
	"github.com/heavenmise/goTasks/tasksForHttp/contacts/pkg/utils"
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
		queryValue := value[len(value) - 1]
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

func CreateContact(c *gin.Context) {
	newContact := &models.Contact{}
	utils.ParseBody(c, newContact)
	contactId := newContact.CreateContact()

	c.JSON(http.StatusOK, gin.H{
		"id": contactId,
	})
}

func UpdateContact(c *gin.Context) {
	var updateContact = &models.Contact{}
	utils.ParseBody(c, updateContact)

	var id string
	query := c.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value) - 1]
		if key == "id" {
			id = queryValue
		}
	}
	
	models.UpdateContact(id, updateContact)
	updatedContact, err := models.GetContact(id)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": updatedContact,
	})
}

func DeleteContact(c *gin.Context) {
	var id string
	query := c.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value) - 1]
		if key == "id" {
			id = queryValue
		}
	}
	models.DeleteContact(id)

	contact, err := models.GetContact(id)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": contact,
	})
}