package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string `json:"name"`
	Email        string `gorm:"unique" json:"email"`
	PhoneNumber  string `json:"-"`
	Password     string `json:"-"`
	ProfilePhoto string `json:"profile_photo"`
	Role         string `json:"role"`
	Cars         []Car  `json:"cars" gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
