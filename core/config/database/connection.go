package database

import (
	"fmt"
	"learn-fiber/core/config"
	"log"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	port, err := strconv.Atoi(config.GetEnv("DB_PORT"))
	if err != nil {
		panic("Failed to parse database port")
	}

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.GetEnv("DB_HOST"),
		port,
		config.GetEnv("DB_USER"),
		config.GetEnv("DB_PASSWORD"),
		config.GetEnv("DB_DATABASE"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return
	}

	DB = db
	fmt.Println("DB Connection opened")

	// _, err := db.DB()
	// if err != nil {
	// 	log.Fatalf("Failed to get raw DB from GORM: %v", err)
	// }

	// return db, sqlDb
}
