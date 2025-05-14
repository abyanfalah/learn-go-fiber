package db

import (
	"context"
	"fmt"
	"learn-fiber/core/config"
	"log"
	"strconv"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db          *gorm.DB
	redisClient *redis.Client
	redisCtx    = context.Background()
)

func init() {
	dbConnect()
	redisConnect()
	// InitMigration(db)
}

func Use() *gorm.DB {
	return db
}

func dbConnect() {
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

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return
	}

	log.Println("DB Connection opened")
}

func Close() {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
		return
	}

	err = sqlDB.Close()
	if err != nil {
		log.Fatalf("Failed to close database connection: %v", err)
		return
	}

	fmt.Println("DB Connection closed")
}

func redisConnect() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     config.GetEnv("REDIS_HOST") + ":" + config.GetEnv("REDIS_PORT"),
		Password: config.GetEnv("REDIS_PASSWORD"),
		DB:       0,
	})

	_, err := redisClient.Ping(redisCtx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis : %v", err)
	}
	log.Println("Connected to Redis")
}

func Redis() *redis.Client {
	return redisClient
}

func RedisCtx() context.Context {
	return redisCtx
}
