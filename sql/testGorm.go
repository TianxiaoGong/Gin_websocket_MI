package main

import (
	"Gin_WebSocket_IM/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:19990609gtx@tcp(127.0.0.1:3306)/ginchat?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	db.AutoMigrate(&models.Community{})

	//// Create
	//user := &models.UserBasic{}
	//user.Name = "gtx"
	//db.Create(user)
	//
	//// Read
	//fmt.Println(db.First(user, 1))
	//
	//// Update
	//db.Model(user).Update("PassWord", 1234)

}
