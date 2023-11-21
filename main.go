package main

import (
	"generale-go/common"
	"generale-go/routes"
	"github.com/gin-gonic/gin"
	"log"
	//"strconv"
)

func main() {

	engine := gin.Default()

	_ = common.InitUserDB() //init

	routes.SetupRouter(engine)
	errEngineRun := engine.Run(":7070")
	if errEngineRun != nil {
		log.Println(errEngineRun)
		return
	}
}
