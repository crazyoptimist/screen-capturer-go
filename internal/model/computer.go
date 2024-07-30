package model

import "gorm.io/gorm"

type Computer struct {
	gorm.Model
	Name   string `json:"name"`
	Status bool   `json:"status"`
}
