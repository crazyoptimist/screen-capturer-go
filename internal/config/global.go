package config

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

const WEB_SERVER_PORT = 8949

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	DB = db
	return db, nil
}
