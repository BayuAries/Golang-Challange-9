package database

import (
	"fmt"
	"log"
	"sesi_12/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "botepang"
	dbPort   = "5432"
	dbName   = "simple_api"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, dbPort)
	dsn := config
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("error conecting to database :", err)
	}

	fmt.Println("sukses koneksi to database")
	db.Debug().AutoMigrate(models.User{}, models.Product{})

}

func GetDB() *gorm.DB {
	return db
}
