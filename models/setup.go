package models

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB


func ConnectDatabase() {
    os.Setenv("GIN_MODE", "release")

    dsn := "root:Abc12345@tcp(localhost:3306)/go_restapi_gin"
    database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic(err)
    }

    database.AutoMigrate(&Product{})

    DB = database
}