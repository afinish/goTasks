package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/heavenmise/goTasks/tasksForHttp/tasks/pkg/config"
	"github.com/heavenmise/goTasks/tasksForHttp/tasks/pkg/models"
	"github.com/heavenmise/goTasks/tasksForHttp/tasks/pkg/utils"
)

func GetTasks(c *gin.Context) {
	fmt.Println("Executing GetTasks function")
	var Tasks []models.Task
	fmt.Println("Panic1")
	rows, err := config.DB.Query(`SELECT * FROM tasks`)
	fmt.Println("Panic2")
	if err != nil {
		fmt.Println("Panic3")
		panic(err)
	}
	
	defer rows.Close()

	for rows.Next() {
		fmt.Println("in rows loop")
		var id int
		var task models.Task
		err = rows.Scan(&id, &task.Name, &task.Status, &task.Priority, &task.CreatedAt, &task.CreatedBy, &task.DueDate)
		if err != nil {
			panic(err)
		}
		fmt.Println(task)
		Tasks = append(Tasks, task)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": Tasks,
	})
}

func GetTask(c *gin.Context) {
	var id string
	var Tasks []models.Task
	query := c.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value) - 1]
		if key == "id" {
			id = queryValue
		}
	}
	sqlQuery := fmt.Sprintf(`SELECT * FROM tasks WHERE id=%s`, id)

	rows, err := config.DB.Query(sqlQuery)
	if err != nil {
		panic(err)
	}
	fmt.Println("Panic1")
	defer rows.Close()
	for rows.Next() {
		var id int
		var name, status, priority, createdBy, createdAt, dueDate string
		err = rows.Scan(&id, &name, &status, &priority, &createdAt, &createdBy, &dueDate)
		if err != nil {
			panic(err)
		}
		task := models.Task{
			Name: name, 
			Status: status,
			Priority: priority,
			CreatedAt: createdAt,
			CreatedBy: createdBy,
			DueDate: dueDate,
		}
		Tasks = append(Tasks, task)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": Tasks,
	})
}

func CreateTask(c *gin.Context) {
	NewTask := &models.Task{}
	utils.ParseBody(c, NewTask)
	Task := NewTask.CreateTask()
	c.JSON(http.StatusOK, gin.H{
		"data": Task,
	})
}

func UpdateTask(c *gin.Context) {
	var task models.Task
	var updateTask = &models.Task{}
	utils.ParseBody(c, updateTask)

	var id string
	query := c.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value) - 1]
		if key == "id" {
			id = queryValue
		}
	}
	sqlQuery := fmt.Sprintf(`SELECT * FROM tasks WHERE id=%s`, id)

	taskRows, err := config.DB.Query(sqlQuery)
	if err != nil {
		panic(err)
	}
	fmt.Println("Panic1")
	defer taskRows.Close()
	for taskRows.Next() {
		var id int
		var name, status, priority, createdBy, createdAt, dueDate string
		err = taskRows.Scan(&id, &name, &status, &priority, &createdAt, &createdBy, &dueDate)
		if err != nil {
			panic(err)
		}
		task = models.Task{
			Name: name, 
			Status: status,
			Priority: priority,
		}
	}
	if updateTask.Name != "" {
		task.Name = updateTask.Name
	}
	if updateTask.Status != "" {
		task.Status = updateTask.Status
	}
	if updateTask.Priority != "" {
		task.Priority = updateTask.Priority
	}
	update := `UPDATE tasks SET name=$1, status=$2, priority=$3 WHERE id=$4`
	_, e := config.DB.Exec(update, task.Name, task.Status, task.Priority, id)
	if e != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": task,
	})
}

func DeleteTask(c *gin.Context) {
	var id string
	query := c.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value) - 1]
		if key == "id" {
			id = queryValue
		}
	}
	fmt.Println("ID: " + id)
	statement := `DELETE FROM tasks WHERE id=$1`
	_, err := config.DB.Exec(statement, id)
	if err != nil {
		panic(err)
	}
}