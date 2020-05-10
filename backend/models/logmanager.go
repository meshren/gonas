package models

import "time"

type LogManager struct {
	ID         uint
	UserID     uint
	ResourceID uint
	Type       uint8
	Extra      string
	CreatedAt  time.Time
}

const (
	LogManagerTypeOther = 0
	LogManagerTypeLogin = 1
)
