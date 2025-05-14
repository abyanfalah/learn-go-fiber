package db

import (
	"fmt"

	"gorm.io/gorm"
)

func InitMigration(db *gorm.DB) {
	// db.AutoMigrate(&model.User{})
	fmt.Println()
}
