package models

import "github.com/jinzhu/gorm"

type GoProducts struct {
	gorm.Model
	Name string `json:"name"`
	Description string `json:"description"`
	Picture string `json:"picture"`
	Price float32 `json:"price"`
	Category string `json:"category"`
}

