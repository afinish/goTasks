package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/afinish/goTasks/tasksForKafka/contacts/producer/pkg/models"
	"github.com/afinish/goTasks/tasksForKafka/contacts/producer/pkg/utils"
	"github.com/gin-gonic/gin"
)

// func GetContacts(c *gin.Context) {
// 	resp, err := http.Get("localhost:8081/contacts")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"data": resp.Body,
// 	})
// }

// func GetContact(c *gin.Context) {
// 	contactID := utils.GetContactID(c)
// 	resp, err := http.Get(fmt.Sprintf("localhost:8081/contact?%s", contactID))
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"data": resp.Body,
// 	})
// }

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
		utils.PostMessage("a", message)
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
		utils.PostMessage("b", message)
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
		utils.PostMessage("c", message)
	}
}
