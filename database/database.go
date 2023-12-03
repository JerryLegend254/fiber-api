package database

import (
	"log"

	"github.com/JerryLegend254/fiber-api/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBInstance struct{
	Db *gorm.DB
}

var Database DBInstance

func ConnectDB() {
	dsn := "root:@tcp(127.0.0.1:3306)/fiber?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		log.Fatal("Error connecting to database\n", err.Error())
	}

	log.Println("Connected to the database successfully")
	db.Logger = logger.Default.LogMode(logger.Info)

	db.AutoMigrate(&models.User{})

	Database = DBInstance{Db: db}
	
	
}
