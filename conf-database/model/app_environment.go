package model

import "time"

type AppEnviroment struct {
	AppEnviromentId int64 `gorm:"primary_key;AUTO_INCREMENT"`
	AppID           int64
	AppCode         string
	EnvironmentName string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time `sql:"index"`
}

func (AppEnviroment) TableName() string {
	return "app_enviroment"
}
