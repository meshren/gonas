package models

type Category struct {
	BaseModel
	UserID   uint
	Name     string
	Describe string
}
