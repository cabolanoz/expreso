package models

import (
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `gorm:"unique"`
	Password string `json:"-"`
}
