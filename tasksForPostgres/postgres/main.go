package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	host	= "localhost"
	port	= 5432
	user	= "postgres"
	password = "fastestmiseintheworld"
	dbname 	= "contactsdb"
)

type Contact struct {
	FirstName 	string
	LastName 	string
	Phone 		string
	Email 		string
	Position 	string
}

func create(firstName, lastName, phone, email, position string) {
	fmt.Println("Executing create function")
	insertStmt := `INSERT INTO contacts(first_name, last_name, phone, email, position) VALUES ($1, $2, $3, $4, $5);`
	_, e := db.Exec(insertStmt, firstName, lastName, phone, email, position)
	if e != nil {
		panic(e)
	}
}

func update(id, firstName, lastName, phone, email, position string) {
	fmt.Println("Executing update function")
	update := `UPDATE contacts SET first_name=$1, last_name=$2, phone=$3, email=$4, position=$5 WHERE id=$6`
	_, err := db.Exec(update, firstName, lastName, phone, email, position, id)
	if err != nil {
		panic(err)
	}
}

func get(id string) {
	fmt.Println("Executing get function")
	selectStmt := fmt.Sprintf("SELECT * FROM contacts WHERE id=%s", id)
	rows, err := db.Query(selectStmt)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var firstName, lastName, phone, email, position string
		err = rows.Scan(&id, &firstName, &lastName, &phone, &email, &position)
		if err != nil {
			panic(err)
		}
		fmt.Println(firstName, lastName, phone, email, position)
	}
	fmt.Println("==============================================================")
}

func getAll() {
	fmt.Println("Executing getAll function")
	rows, err := db.Query(`SELECT * FROM contacts`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var firstName, lastName, phone, email, position string
		err = rows.Scan(&id, &firstName, &lastName, &phone, &email, &position)
		if err != nil {
			panic(err)
		}
		fmt.Println(id, firstName, lastName, phone, email, position)
	}
	fmt.Println("==============================================================")
}

func deleteContact(id int) {
	fmt.Println("Executing deleteContact function")
	deleteStmt	:= `DELETE FROM contacts WHERE id=$1`
	_, err := db.Exec(deleteStmt, id)
	if err != nil {
		panic(err)
	}
}

func main() {
	dbURL := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	var err error
	db, err = sql.Open("postgres", dbURL)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	create(
		"Sundar",
		"Pichai",
		"+2110000011",
		"sundar@exmpl.com",
		"CEO",
	)
	getAll()
	get("2")
	update(
		"2",
		"Mike",
		"Michaels",
		"+9871111111",
		"mich@exmpl.com",
		"Team Lead",
	)
	get("2")
	getAll()
	deleteContact(3)
	getAll()
}
