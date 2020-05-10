package models

import (
	"github.com/jinzhu/gorm"
	"gonas/utils"
)

// file type 0 unknown 1 document(txt,doc,pdf,xlsx...) 2 image 3 video 4 audio 5 binary

type File struct {
	gorm.Model
	Display      string
	Hash         string
	Type         int8
	Size         int64
	DirectoryID  int
	Directory    string
	OriginalName string
	DeviceID     uint
}

func (f *File) GetAll(order string, limit uint, offset uint) (files []File, err error) {
	db, err := connection()
	defer db.Close()
	if err != nil {
		utils.ErrDetail(err)
		return
	}
	if order != "" {
		db.Where(f).Order(order)
	}
	if limit > 0 {
		db.Limit(limit)
	}
	if offset >= 0 {
		db.Offset(offset)
	}
	db.Find(&files)
	return
}
