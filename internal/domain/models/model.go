package models

import "time"

type User struct {
	ID            int64
	ApplicationID uint
	Email         string
	Password      string
	CreateAt      time.Time
}
