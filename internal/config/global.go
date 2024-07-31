package config

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"

	"screencapturer/internal/domain/model"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&model.Computer{}); err != nil {
		return nil, err
	}

	DB = db
	return db, nil
}
