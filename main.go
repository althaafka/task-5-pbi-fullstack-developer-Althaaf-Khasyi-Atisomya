package main

import (
	"github.com/althaafka/task-5-pbi-fullstack-developer-Althaaf-Khasyi-Atisomya.git/database"
	"github.com/althaafka/task-5-pbi-fullstack-developer-Althaaf-Khasyi-Atisomya.git/controllers/usercontroller"
	"github.com/althaafka/task-5-pbi-fullstack-developer-Althaaf-Khasyi-Atisomya.git/controllers/productcontroller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database.Connect()

	r.POST("/users/register", usercontroller.Register)
	r.POST("/users/login", usercontroller.Login)
	r.PUT("/users/:userId", usercontroller.Update)

	r.POST("/photos", productcontroller.Create)
	r.GET("/photos",productcontroller.Show)

	r.Run()
}
