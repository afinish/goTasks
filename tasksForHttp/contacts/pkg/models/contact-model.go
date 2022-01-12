package models

import "github.com/heavenmise/goTasks/tasksForHttp/contacts/pkg/config"

type Contact struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Position  string `json:"position"`
}

func (contact *Contact) CreateContact() *Contact {
	insertStmt := `INSERT INTO contacts(first_name, last_name, phone, email, position) VALUES ($1, $2, $3, $4, $5);`
	config.DB.Exec(insertStmt, contact.FirstName, contact.LastName, contact.Phone, contact.Email, contact.Position)
	return contact
}