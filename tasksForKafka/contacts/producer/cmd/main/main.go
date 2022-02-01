package main

import "github.com/afinish/goTasks/tasksForKafka/contacts/producer/pkg/routes"

func main() {
	r := routes.RegisterNewRouter()
	r.Run(":8080")
}
