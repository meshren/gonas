package models

import "time"

type LogUser struct {
	ID         uint
	ManagerID  uint
	ResourceID uint
	Type       uint8
	Extra      string
	CreatedAt  time.Time
}

func (l *LogUser) Create() (bool, error) {
}

const (
	LogUserTypeOther = 0
	LogUserTypeLogin = 1
)

type logUserI interface {
	Create() (bool, error)
}
