package main

import (
	"github.com/althaafka/task-5-pbi-fullstack-developer-Althaaf-Khasyi-Atisomya.git/database"
	"github.com/althaafka/task-5-pbi-fullstack-developer-Althaaf-Khasyi-Atisomya.git/router"
)

func main() {
	database.Connect()

	r := router.SetupRouter()

	r.Run()
}
