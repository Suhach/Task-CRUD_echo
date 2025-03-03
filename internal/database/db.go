package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ---------переменная----для---использования---DB----\\
var DB *gorm.DB

// ----------------------func--------------------------\\
func InitDB() {
	dsn := "host=localhost user=postgres password=1234567 dbname=postgres port=5432 sslmode=disable" // dsn
	var err error                                                                                    // error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})                                          // init DB
	if err != nil {
		log.Fatalf("Failed to connect to DataBase: %v", err)
	}
}
