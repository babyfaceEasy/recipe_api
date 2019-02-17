package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// https://medium.com/aubergine-solutions/how-i-handled-null-possible-values-from-database-rows-in-golang-521fb0ee267

// Recipe the model Recipe
type Recipe struct {
	gorm.Model
	Name       string `gorm:"size:100 NOT NULL" json:"name"`
	PrepTime   string `gorm:"size:10 NOT NULL" json:"prepTime"`
	Difficulty int    `gorm:"size:3 NOT NULL" json:"difficulty"`
	Vegetarian bool   `json:"vegetarian"`
}
