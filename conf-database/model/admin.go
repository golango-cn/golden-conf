package model

import "time"

type Admin struct {
	AdminId   int64 `gorm:"primary_key;AUTO_INCREMENT"`
	AdminName string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
