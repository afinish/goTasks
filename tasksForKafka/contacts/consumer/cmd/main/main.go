package main

import (
	"database/sql"

	"github.com/afinish/goTasks/tasksForKafka/contacts/consumer/pkg/config"
	"github.com/afinish/goTasks/tasksForKafka/contacts/consumer/pkg/routes"
	"github.com/afinish/goTasks/tasksForKafka/contacts/consumer/pkg/subscribe"
	_ "github.com/lib/pq"
)

func main() {
	var err error
	config.DB, err = sql.Open("postgres", config.DbURL())
	if err != nil {
		panic(err)
	}
	defer config.DB.Close()

	go subscribe.Post()
	go subscribe.Put()
	go subscribe.Delete()

	r := routes.RegisterNewRouter()
	r.Run(":8081")
}
