package main

import (
	"database/sql"

	"github.com/heavenmise/goTasks/tasksForHttp/contacts/pkg/config"
	"github.com/heavenmise/goTasks/tasksForHttp/contacts/pkg/routes"
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