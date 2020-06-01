package models

type UserFiles struct {
	BaseModel
	UserID      uint      `json:"user_id"`
	FileID      uint      `json:"file_id"`
	DirectoryID uint      `json:"directory_id"`
	File        File      `json:"file"`
	Directory   Directory `json:"directory"`
}

func (UserFiles) TableName() string {
	return "user_files"
}

