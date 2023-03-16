package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateDbInstance() *gorm.DB {
	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Cannot connect to DB: ", err)
	}

	fmt.Println("Connected to PostgresSQL!")

	return db
}
