package main

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ------var----for----use-with---DB---\\
var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost user=postgres password=1234567 dbname=postgres port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to DataBase: ", err)
	}
}
