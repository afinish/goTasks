package main

import (
	// "database/sql"

	"database/sql"

	"github.com/heavenmise/goTasks/tasksForHttp/tasks/pkg/config"
	"github.com/heavenmise/goTasks/tasksForHttp/tasks/pkg/routes"
	_ "github.com/lib/pq"
)

func main() {
	var err error
	config.DB, err = sql.Open("postgres", config.DbURL())
	if err != nil {
		panic(err)
	}
	defer config.DB.Close()

	r := routes.RegisterNewRouter()
	r.Run()
}