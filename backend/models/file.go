package models

import (
	"errors"
	"gonas/utils"
	"log"
)

// file type 1. folder 2. document(txt,doc,pdf,xlsx...) 3. image 4. video 5. audio 6. binary
const (
	FileTypeText = iota
)

type File struct {
	BaseModel
	Display      string     `json:"display"`
	Hash         string     `json:"hash"`
	Type         int8       `json:"type"`
	Size         int64      `json:"size"`
	OriginalName string     `json:"original_name"`
	DeviceID     uint       `json:"device_id"`
	Users        []*User    `gorm:"many2many:user_files",json:"users"`
}

func (f *File) Create() error {
	db, err := connection()
	defer db.Close()
	if err != nil {
		return err
	}
	log.Println("file: ", f)
	ok := db.Create(f)
	log.Println("ok: ", ok)
	return nil
}

type FileI interface {
	GetAll() (files []File, err error)
	Create() error
}

func (f *File) GetAll(order string, limit uint, offset uint) (files []File, err error) {
	db, err := connection()
	defer db.Close()
	if err != nil {
		utils.ErrDetail(err)
		return
	}
	db.Debug().Where(f).Order(order).Limit(limit).Offset(offset).Find(&files)
	return
}

func (f *File) CreateFolder() error {
	if f.Display == "" {
		return errors.New("目录名称不能为空！")
	}
	db, err := connection()
	defer db.Close()
	if err != nil {
		return err
	}
	db.Create(f)
	return nil
}
