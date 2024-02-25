package router

import (
	"github.com/gin-gonic/gin"
	"github.com/althaafka/task-5-pbi-fullstack-developer-Althaaf-Khasyi-Atisomya.git/controllers/usercontroller"
	"github.com/althaafka/task-5-pbi-fullstack-developer-Althaaf-Khasyi-Atisomya.git/controllers/productcontroller"
	"github.com/althaafka/task-5-pbi-fullstack-developer-Althaaf-Khasyi-Atisomya.git/middlewares"
)

func SetupRouter() *gin.Engine {
    router := gin.Default()

        router.POST("/users/register", usercontroller.Register)
        router.POST("/users/login", usercontroller.Login)
        router.PUT("/users/:userId", middlewares.AuthMiddleware(), usercontroller.Update)
        router.DELETE("/:userId", middlewares.AuthMiddleware(), usercontroller.Delete)

        router.POST("/photos", middlewares.AuthMiddleware(), productcontroller.Create)
        router.GET("/photos", middlewares.AuthMiddleware(), productcontroller.Show)
        router.PUT("/photos/:photoId", middlewares.AuthMiddleware(), productcontroller.Update)
        router.DELETE("/photos/:photoId",middlewares.AuthMiddleware(), productcontroller.Delete)

    return router
}
