package models

import "github.com/jinzhu/gorm"

// Todo struct keeps info about todo.
type Todo struct {
	gorm.Model
	Title      string
	IsFinished bool
}
