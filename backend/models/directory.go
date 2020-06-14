package models

type Directory struct {
	BaseModel
	UserID   uint   `json:"user_id"`
	Display  string `json:"display"`
	ParentID uint   `json:"parent_id"`
}

func (d *Directory) Create() (ok bool, err error) {
	db, err := connection()
	if err != nil {
		return false, err
	}
	db.Debug().Create(d)
	return true, nil
}

func (d *Directory) Update() bool {
	panic("implement me")
}

func (d *Directory) Delete() bool {
	panic("implement me")
}

type DirectoryI interface {
	Create() (ok bool, err error)
	Update() bool
	Delete() bool
}
