package main

import (
	"github.com/say-bonjour/be-say-bonjour/model"

	"gorm.io/gorm"
)

func migrateDB(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
}
