package database

import (
	"learn-fiber/model"

	"gorm.io/gorm"
)

func InitMigration(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
}
