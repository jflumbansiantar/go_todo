package models

import "github.com/jinzhu/gorm"

type TodoItem struct {
	gorm.Model
	Description string
	IsCompleted bool
}