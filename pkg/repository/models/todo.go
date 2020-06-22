package models

import "github.com/jinzhu/gorm"

// Todo struct contains info about todo.
type Todo struct {
	gorm.Model
	Title      string
	IsFinished bool
}
