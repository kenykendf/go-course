package database

import (
	"fmt"
	"log"

	"final-project/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "gocourse"
	password = "hacktiv8"
	port     = 5432
	dbname   = "final-project"
	DB       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	DB, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database :", err)
	}

	DB.Debug().AutoMigrate(models.User{}, models.Photo{}, models.Comment{}, models.Social{})

	fmt.Println("Success To Connect Docker Postgresql")
}

func GetDB() *gorm.DB {
	return DB
}
