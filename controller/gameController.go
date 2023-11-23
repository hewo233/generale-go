package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type gameBeginData struct {
	Usernum int    `json:"usernum"`
	Seed    string `json:"seed"`
}

func mapInit(c *gin.Context) {

}

func gameBegin(c *gin.Context) {
	var data gameBeginData
	errBindData := c.BindJSON(&data)
	if errBindData != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    40003,
			"message": "Bind data error",
		})

		return
	}
}
