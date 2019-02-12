package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Recipe the model Recipe
type Recipe struct {
	gorm.Model
	Name       string `gorm:"size:100"`
	PrepTime   string `gorm:"size:10"`
	Difficulty int16  `gorm:"size:3"`
	Vegetarian bool
}
