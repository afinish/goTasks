package models

import (
	"time"

	"github.com/heavenmise/goTasks/tasksForHttp/tasks/pkg/config"
)

type Task struct {
	Name      string	`json:"name"`
	Status    string	`json:"status"`
	Priority  string	`json:"priority"`
	CreatedAt string	`json:"created_at"`
	CreatedBy string	`json:"created_by"`
	DueDate   string	`json:"due_date"`
}

func (task *Task) CreateTask() *Task {
	statement := `INSERT INTO tasks(name, status, priority, created_at, created_by, due_date) VALUES ($1, $2, $3, $4, $5, $6);`
	_, err := config.DB.Exec(statement, task.Name, task.Status, task.Priority, time.Now().String(), task.CreatedBy, task.DueDate)
	if err != nil {
		panic(err)
	}
	return task
}