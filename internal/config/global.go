package config

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"

	"screencapturer/internal/domain/model"
)

const WEB_SERVER_PORT = 8949
const CLIENT_WEB_PORT = 9999

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
