package common

import (
	"generale-go/model"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var UserDB *gorm.DB

func InitUserDB() *gorm.DB {
	dbUser, errUser := gorm.Open(sqlite.Open("db/user.db"), &gorm.Config{})
	if errUser != nil {
		panic("failed to connect UserDatabase")
	}
	errAuto := dbUser.AutoMigrate(&model.User{})
	if errAuto != nil {
		log.Println("dbUser error.(AutoMigrate),: ", errAuto)
	}

	UserDB = dbUser

	return dbUser

}

func GetUserDB() *gorm.DB {
	return UserDB
}
