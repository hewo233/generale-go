package routes

import "github.com/gin-gonic/gin"

func SetupRouter(engin *gin.Engine) {
	engine.POST("/user/register", controller.Register())
}
