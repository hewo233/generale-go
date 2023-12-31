package routes

import (
	"generale-go/controller"
	"generale-go/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(engine *gin.Engine) {

	//User
	engine.POST("/user/register", controller.Register)
	engine.POST("/user/login", controller.Login)
	engine.GET("/user/userinfo", middleware.AuthMiddleware(), controller.Info)

}
