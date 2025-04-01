package models

import (
	"gorm.io/gorm"
)

type Car struct {
	gorm.Model     `swaggerignore:"true"`
	UserId  uint   `json:"user_id"`
	Make    string `json:"make"`
	Pattern string `json:"pattern"` // e.g. "Toyota" - This is replacement for "model"
	Year    int    `json:"year"`
	Plate   string `json:"plate"`
	Color   string `json:"color"`
	Seats   int    `json:"seats"`
}
