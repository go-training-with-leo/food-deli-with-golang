package database

import (
	"fmt"
	"log"

	"api-gateway/src/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateDbInstance() *gorm.DB {
	username := config.App.DatabaseUsername
	password := config.App.DatabasePassword
	host := config.App.DatabaseHost
	port := config.App.DatabasePort
	name := config.App.DatabaseName
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, username, password, name, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Cannot connect to DB: ", err)
	}

	fmt.Println("Connected to PostgresSQL!")

	return db
}
