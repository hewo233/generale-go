package middleware

import (
	"generale-go/common"
	"generale-go/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenString := c.GetHeader("Authorization") //from Header get token

		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    40101,
				"message": "Insufficient permissions",
			})
			return
		}

		tokenString = tokenString[7:] // remove Bearer

		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    40101,
				"message": "Insufficient permissions",
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
				"message": "Insufficient permissions",
			})
			return
		}

		c.Set("user", user) //c.Get(key) to get key
		c.Next()

	}
}
