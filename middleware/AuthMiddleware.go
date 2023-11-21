package middleware

import (
	"generale-go/common"
	"generale-go/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenString, errTokenString := c.Cookie("jwt")
		println(tokenString)
		if errTokenString != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    40101,
				"message": "Insufficient permissions1",
				"error":   errTokenString,
			})
			return
		}

		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    40101,
				"message": "Insufficient permissions2",
			})
			return
		}

		// get userId
		userUid := claims.UserUid
		DB := common.GetUserDB()
		var user model.User
		DB.First(&user, userUid) //find user from database

		if len(user.Uid) == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    40101,
				"message": "Insufficient permissions3",
			})
			return
		}

		c.Set("user", user) //c.Get(key) to get key
		c.Next()

	}
}
