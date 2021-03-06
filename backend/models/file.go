package models

import (
	"errors"
	"gonas/utils"
	"log"
)
// 是否要把文件夹和文件放入同一个表中？
// file type 1. directory 2. document(txt,doc,pdf,xlsx...) 3. image 4. video 5. audio 6. binary
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
	Users        []*User    `json:"-",gorm:"many2many:user_files"`
}

func (f *File) Create() (fileID uint, err error) {
	db, err := connection()
	defer db.Close()
	if err != nil {
		return 0, err
	}
	log.Println("file: ", f)
	ok := db.Create(f)
	log.Println("ok: ", ok)
	return f.ID, nil
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
