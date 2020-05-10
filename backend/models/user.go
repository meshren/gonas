package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	Username string
	Password string
	DeviceID uint
	LastLoginIp uint32
	LastLoginAt time.Time
	LoginCount uint64
}
