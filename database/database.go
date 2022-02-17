package database

import (
	"log"

	"github.com/thanh/go-book1/model"
	"gorm.io/gorm/logger"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("book.db"), &gorm.Config{})
	if err != nil {
		panic("Failed connection to database")
	}
	log.Println("Connected Successfully")
	db.Logger = logger.Default.LogMode(logger.Info) //tao thoi gian moi
	log.Println("Running Migrations")

	db.AutoMigrate(&model.Book{})
	Database = DbInstance{
		Db: db,
	}
}
