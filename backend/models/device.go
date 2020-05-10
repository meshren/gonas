package models

import "github.com/jinzhu/gorm"

type Device struct {
	gorm.Model
	Display string
	Memo    string
	Mac     string
	LastIp  uint32
	Type    uint8
}

func (d *Device) New() (bool, error) {
	db, err := connection()
	defer db.Close()
	if err != nil {
		return false, err
	}
	db.Create(d)
	return true, nil
}

func (d *Device) Update() (bool, error) {
	db, err := connection()
	defer db.Close()
	if err != nil {
		return false, err
	}
	db.First(d)
	db.Save(d)
	return true, nil
}

func (d *Device) DeleteByID() (bool, error) {
	db, err := connection()
	defer db.Close()
	if err != nil {
		return false, err
	}
	if d.ID != 0 {
		db.Delete(d)
	}
	return true, nil

}

type DeviceI interface {
	New() (bool, error)
	Update() (bool, error)
	DeleteByID() (bool, error)
}
