package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/heavenmise/goTasks/tasksForHttp/contacts/pkg/config"
	"github.com/heavenmise/goTasks/tasksForHttp/contacts/pkg/models"
	"github.com/heavenmise/goTasks/tasksForHttp/contacts/pkg/utils"
)

func GetContacts(c *gin.Context) {
	fmt.Println("Executing GetContacts function")
	var Contacts []models.Contact
	fmt.Println("Panic1")
	contactRows, err := config.DB.Query(`SELECT * FROM contacts`)
	fmt.Println("Panic2")
	if err != nil {
		fmt.Println("Panic3")
		panic(err)
	}
	
	defer contactRows.Close()

	for contactRows.Next() {
		fmt.Println("in contactRows loop")
		var id int
		var contact models.Contact
		err = contactRows.Scan(&id, &contact.FirstName, &contact.LastName, &contact.Phone, &contact.Email, &contact.Position)
		if err != nil {
			panic(err)
		}
		fmt.Println(contact)
		Contacts = append(Contacts, contact)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": Contacts,
	})
}

func GetContact(c *gin.Context) {
	var id string
	var Contacts []models.Contact
	query := c.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value) - 1]
		if key == "id" {
			id = queryValue
		}
	}
	sqlQuery := fmt.Sprintf(`SELECT * FROM contacts WHERE id=%s`, id)

	contactRows, err := config.DB.Query(sqlQuery)
	if err != nil {
		panic(err)
	}
	fmt.Println("Panic1")
	defer contactRows.Close()
	for contactRows.Next() {
		var id int
		var firstName, lastName, phone, email, position string
		err = contactRows.Scan(&id, &firstName, &lastName, &phone, &email, &position)
		if err != nil {
			panic(err)
		}
		contact := models.Contact{
			FirstName: firstName, 
			LastName: lastName, 
			Phone: phone, 
			Email: email, 
			Position: position,
		}
		Contacts = append(Contacts, contact)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": Contacts,
	})
}

func CreateContact(c *gin.Context) {
	NewContact := &models.Contact{}
	utils.ParseBody(c, NewContact)
	Contact := NewContact.CreateContact()
	c.JSON(http.StatusOK, gin.H{
		"data": Contact,
	})
}

func UpdateContact(c *gin.Context) {
	var contact models.Contact
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
	sqlQuery := fmt.Sprintf(`SELECT * FROM contacts WHERE id=%s`, id)

	contactRows, err := config.DB.Query(sqlQuery)
	if err != nil {
		panic(err)
	}
	fmt.Println("Panic1")
	defer contactRows.Close()
	for contactRows.Next() {
		var id int
		var firstName, lastName, phone, email, position string
		err = contactRows.Scan(&id, &firstName, &lastName, &phone, &email, &position)
		if err != nil {
			panic(err)
		}
		contact = models.Contact{
			FirstName: firstName, 
			LastName: lastName, 
			Phone: phone, 
			Email: email, 
			Position: position,
		}
	}
	if updateContact.FirstName != "" {
		contact.FirstName = updateContact.FirstName
	}
	if updateContact.LastName != "" {
		contact.LastName = updateContact.LastName
	}
	if updateContact.Phone != "" {
		contact.Phone = updateContact.Phone
	}
	if updateContact.Email != "" {
		contact.Email = updateContact.Email
	}
	if updateContact.Position != "" {
		contact.Position = updateContact.Position
	}
	update := `UPDATE contacts SET first_name=$1, last_name=$2, phone=$3, email=$4, position=$5 WHERE id=$6`
	_, e := config.DB.Exec(update, contact.FirstName, contact.LastName, contact.Phone, contact.Email, contact.Position, id)
	if e != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": contact,
	})
}

func DeleteContact(c *gin.Context) {
	var name string
	query := c.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value) - 1]
		if key == "first_name" {
			name = queryValue
		}
	}
	deleteStmt := `DELETE FROM contacts WHERE first_name=$1`
	_, err := config.DB.Exec(deleteStmt, name)
	if err != nil {
		panic(err)
	}
}