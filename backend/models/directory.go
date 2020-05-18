package models

type Directory struct {
	BaseModel
	UserID   uint   `json:"user_id"`
	Display  string `json:"display"`
	ParentID uint   `json:"parent_id"`
}
