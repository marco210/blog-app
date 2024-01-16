package database

import (
	"log"
	"os"

	"github.com/marco210/blog-app/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)


var DBConn *gorm.DB

func ConnectDB()  {
	user := os.Getenv("db_user")
	password := os.Getenv("db_password")
	dbname := os.Getenv("db_name")

	dsn := "root:password123@tcp(127.0.0.1:3306)/blog-db?charset=utf8mb4&parseTime=True&loc=Local"
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})

	if err != nil {
		panic("Database connect failed")
	}

	log.Println("Connection successful")

	db.AutoMigrate(new(model.Blog))

	DBConn = db

}