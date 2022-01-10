package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	host	= "localhost"
	port	= 5432
	user	= "postgres"
	password = "fastestmiseintheworld"
	dbname 	= "tasksdb"
)

type Task struct {
	Name 		string
	Status 		string
	Priority 	string
	UpdatedAt	time.Time
	CreatedBy 	string
	DueDate 	time.Time
}

func create(Name string, Status string, Priority string, UpdatedAt time.Time, CreatedBy string, DueDate time.Time) {
	fmt.Println("Executing create function")
	insertStmt := `INSERT INTO tasks(name, status, priority, updated_at, created_by, due_date) VALUES ($1, $2, $3, $4, $5, $6);`
	_, e := db.Exec(insertStmt, Name, Status, Priority, UpdatedAt.String(), CreatedBy, DueDate.String())
	if e != nil {
		panic(e)
	}
}

func update(id, Name string, Status string, Priority string, UpdatedAt time.Time, CreatedBy string, DueDate time.Time) {
	fmt.Println("Executing update function")
	update := `UPDATE tasks SET name=$1, status=$2, priority=$3, updated_at=$4, created_by=$5, due_date=$6 WHERE id=$7`
	_, err := db.Exec(update, Name, Status, Priority, UpdatedAt.String(), CreatedBy, DueDate.String(), id)
	if err != nil {
		panic(err)
	}
}

func get(id string) {
	fmt.Println("Executing get function")
	selectStmt := fmt.Sprintf("SELECT * FROM tasks WHERE id=%s", id)
	rows, err := db.Query(selectStmt)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var Name, Status, Priority, UpdatedAt, CreatedBy, DueDate string
		err = rows.Scan(&id, &Name, &Status, &Priority, &UpdatedAt, &CreatedBy, &DueDate)
		if err != nil {
			panic(err)
		}
		fmt.Println(Name, Status, Priority, UpdatedAt, CreatedBy, DueDate)
	}
	fmt.Println("==============================================================")
}

func getAll() {
	fmt.Println("Executing getAll function")
	rows, err := db.Query(`SELECT * FROM tasks`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var Name, Status, Priority, UpdatedAt, CreatedBy, DueDate string
		err = rows.Scan(&id, &Name, &Status, &Priority, &UpdatedAt, &CreatedBy, &DueDate)
		if err != nil {
			panic(err)
		}
		fmt.Println(Name, Status, Priority, UpdatedAt, CreatedBy, DueDate)
	}
	fmt.Println("==============================================================")
}

func deleteTask(id int) {
	fmt.Println("Executing deleteTask function")
	deleteStmt	:= `DELETE FROM tasks WHERE id=$1`
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
		"Groceries",
		"not-done",
		"high",
		time.Now(),
		"Josh",
		time.Date(2022, time.January, 15, 12, 0, 0, 0, time.UTC),
	)
	create(
		"Change car oil",
		"not-done",
		"medium",
		time.Now(),
		"Mike",
		time.Date(2022, time.January, 13, 12, 0, 0, 0, time.UTC),
	)
	getAll()
	update(
		"2",
		"Change car oil",
		"done",
		"medium",
		time.Now(),
		"Mike",
		time.Date(2022, time.January, 13, 12, 0, 0, 0, time.UTC),
	)
	get("2")
	deleteTask(2)
	getAll()
}