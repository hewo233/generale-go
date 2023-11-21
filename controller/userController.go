package controller

import (
	"generale-go/common"
	"generale-go/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func Register(c *gin.Context) {
	dbUser := common.GetUserDB()
	name := c.PostForm("name")
	password := c.PostForm("password")
	uid := uuid.New().String()

	if len(name) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    40001, //40001 means Register Wrong
			"message": "name should not be null",
		})
	}

	if len(password) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    40001,
			"message": "Password needs to be more than 6 characters",
		})
	}

	var user model.User
	dbUser.Where("name = ?", name).First(&user)
	if len(user.Uid) != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    40001,
			"message": "this user name has been used.",
		})
	}

	//creat new user
	HashPassword, errPassHash := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if errPassHash != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    50001,
			"message": "Password encryption error",
		})
		return
	}

	newUser := model.User{
		Name:     name,
		Password: string(HashPassword),
		Uid:      uid,
	}
	errUserCreat := dbUser.Create(&newUser).Error
	if errUserCreat != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"code":    50002,
			"message": errUserCreat,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    20000,
		"message": "Register successfully",
	})
}

func Login(c *gin.Context) {
	dbUser := common.GetUserDB()
	var requestUser model.User

	errBindUser := c.Bind(&requestUser)
	if errBindUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    40002,
			"message": "Bind wrong",
		})
	}

	name := requestUser.Name
	password := requestUser.Password

	//name := c.PostForm("name")

	if len(name) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    40002,
			"message": "Wrong Name",
		})
		return
	}

	if len(password) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    40002,
			"message": "Password needs to be more than 6 characters",
		})
		return
	}

	var user model.User
	dbUser.Where("name = ?", name).First(&user)
	if len(user.Uid) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    40002,
			"message": "User not found",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    40002,
			"message": "Wrong Password",
		})
		return
	}

	token, err := common.ReleaseToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    50000,
			"message": "System error",
		})
		log.Printf("Token generate error: %v", err) //打印错误日志
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 20000,
		"data": gin.H{
			"token": token,
		},
		"message": "登录成功",
	})

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "jwt",
		Value:    token,
		HttpOnly: true,
		Path:     "/",
		MaxAge:   3600, // 3600s (1 h)
	})

}

func Info(c *gin.Context) {
	user, _ := c.Get("user") //get user from key "user"
	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"user": user,
		},
	})
}
