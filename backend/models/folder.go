package models

type Folder struct {
	BaseModel
	UserID   int
	Display  string `json:"display"`
	ParentID int
}


