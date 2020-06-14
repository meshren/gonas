package models

type UserFiles struct {
	BaseModel
	UserID       uint      `json:"user_id"`
	FileID       uint      `json:"file_id"`
	DirectoryID  uint      `json:"directory_id"`
	DirectoryPID uint      `gorm:"column:directory_pid",json:"directory_pid"`
	File         File      `json:"file"`
	Directory    Directory `json:"directory"`
}

func (UserFiles) TableName() string {
	return "user_files"
}

func (uf *UserFiles) Create() (ok bool, err error) {
	db, err := connection()
	if err != nil {
		return false, err
	}
	db.Debug().Create(uf)
	return true, nil
}
