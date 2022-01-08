package tasks

import (
	"fmt"
	"strconv"
)

var tCounter int = 0
var tasks map[string]Task

type Task struct {
	Name 		string
	Status 		string
	Priority 	string
	CreatedAt 	string
	CreateBy 	string
	DueDate 	string
}

func create(name, status, priority, createdAt, createBy, dueDate string) {
	tCounter++
	id := strconv.Itoa(tCounter)
	tasks[id] = Task{
		name,
		status,
		priority,
		createdAt,
		createBy,
		dueDate,
	}
}

func update(id, name, status, priority, createdAt, createBy, dueDate string) {
	tasks[id] = Task{
		name,
		status,
		priority,
		createdAt,
		createBy,
		dueDate,
	}
}

func get(id string) Task {
	return tasks[id]
}

func getAll() map[string]Task {
	return tasks
}

func deleteTask(id string) {
	delete(tasks, id)
}

func init() {
	tasks = make(map[string]Task)
	create(
		"Groceries",
		"Not-done",
		"Important",
		"2021-10-11",
		"John",
		"2021-10-13",
	)

	create(
		"Homework",
		"Not-done",
		"Important",
		"2021-10-9",
		"Mark",
		"2021-10-10",
	)

	create(
		"Eat food",
		"done",
		"Neutral",
		"2021-10-8",
		"Cille",
		"2021-10-8",
	)

	fmt.Println(get("2"))
	fmt.Println(getAll())

	update(
		"1",
		"Groceries",
		"done",
		"Important",
		"2021-10-11",
		"John",
		"2021-10-13",
	)

	fmt.Println(get("1"))
	deleteTask("3")
	fmt.Println(getAll())
}
