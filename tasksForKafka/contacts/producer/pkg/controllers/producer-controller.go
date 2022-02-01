package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/afinish/goTasks/tasksForKafka/contacts/producer/pkg/models"
	"github.com/afinish/goTasks/tasksForKafka/contacts/producer/pkg/utils"
	"github.com/gin-gonic/gin"
)

func GetContacts(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "http://localhost:8081/contacts")
}

func GetContact(c *gin.Context) {
	id := utils.GetContactID(c)
	location := fmt.Sprintf("http://localhost:8081/contact?id=%s", id)
	c.Redirect(http.StatusMovedPermanently, location)
}

func CreateContact(c *gin.Context) {
	newContact := &models.Contact{}
	utils.ParseBody(c, newContact)
	message, err := json.Marshal(newContact)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": newContact,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": newContact,
		})
		utils.PostMessage("POST", message)
	}
}

func UpdateContact(c *gin.Context) {
	var updateContact = &models.Contact{}
	utils.ParseBody(c, updateContact)
	message, err := json.Marshal(updateContact)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": updateContact,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": updateContact,
		})
		utils.PostMessage("PUT", message)
	}
}

func DeleteContact(c *gin.Context) {
	contactID := utils.GetContactID(c)
	message, err := json.Marshal(contactID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"data": contactID,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"data": contactID,
		})
		utils.PostMessage("DELETE", message)
	}
}
