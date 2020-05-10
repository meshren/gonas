package models

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	UserID uint
	Name string
	Describe string
}
