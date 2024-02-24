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
        // userRouter.DELETE("/:userId", middlewares.AuthMiddleware(), usercontroller.Delete)
        router.POST("/photos", productcontroller.Create)
        router.GET("/photos", productcontroller.Show)
        router.PUT("/photos/:photoId", productcontroller.Update)
        router.DELETE("/photos/:photoId", productcontroller.Delete)

    return router
}
