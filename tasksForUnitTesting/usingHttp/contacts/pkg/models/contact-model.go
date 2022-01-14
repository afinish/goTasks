package models

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/heavenmise/goTasks/tasksForHttp/contacts/pkg/config"
)

type Contact struct {
	ID			int64	`json:"id"`
	FirstName	string	`json:"first_name"`
	LastName	string	`json:"last_name"`
	Phone		string	`json:"phone"`
	Email		string	`json:"email"`
	Position	string	`json:"position"`
}

func GetContacts() *[]Contact {
	var contacts []Contact
	rows, err := config.DB.Query(`SELECT * FROM contacts`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var contact Contact
		err := rows.Scan(&contact.ID, &contact.FirstName, &contact.LastName, &contact.Phone, &contact.Email, &contact.Position)
		if err != nil {
			panic(err)
		}
		contacts = append(contacts, contact)
	}
	return &contacts
}
func GetContact(id string) (*Contact, error) {
	var contact Contact
	row := config.DB.QueryRow(`SELECT * FROM contacts WHERE id=$1`, id)
	err := row.Scan(&contact.ID, &contact.FirstName, &contact.LastName, &contact.Phone, &contact.Email, &contact.Position)
	switch err {
    case sql.ErrNoRows:
        fmt.Println("No rows were returned!")
        return &contact, nil
    case nil:
        return &contact, nil
    default:
        log.Fatalf("Unable to scan the row. %v", err)
    }
	return &contact, nil
}

func (contact *Contact) CreateContact() int64 {
	var id int64
	insertStmt := `INSERT INTO contacts(first_name, last_name, phone, email, position) VALUES ($1, $2, $3, $4, $5) RETURNING id;`
	err := config.DB.QueryRow(insertStmt, contact.FirstName, contact.LastName, contact.Phone, contact.Email, contact.Position).Scan(&id)
	if err != nil {
		panic(err)
	}
	return id
}

func DeleteContact(id string) {
	_, err := config.DB.Query(`DELETE FROM contacts WHERE id=$1`, id)
	if err != nil {
		panic(err)
	}
}

func UpdateContact(id string, updateContact *Contact) {
	contact, err := GetContact(id)
	if err != nil {
		panic(err)
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
}