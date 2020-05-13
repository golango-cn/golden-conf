package model

import (
	"fmt"
	"time"
)

type App struct {
	AppId     int64 `gorm:"primary_key;AUTO_INCREMENT"`
	AppCode   string
	AppName   string
	AppDesc   string
	AdminId   int64
	Admin     Admin `gorm:"foreignkey:AdminId;association_foreignkey:AdminId"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (App) TableName() string {
	return "app"
}

// 初始化
func (a *App) Create(envs []string) error {

	if len(envs) == 0 {
		return fmt.Errorf("[CreateApp]The environment configuration cannot be empty")
	}

	var err error
	tx := Db.Begin()
	defer func() {
		if err != nil {
			err = fmt.Errorf("[CreateApp]Create app:%s-%s error: %s", a.AppName, a.AppCode, err.Error())
			tx.Rollback()
		}
	}()

	if err = tx.Create(a).Error; err != nil {
		return err
	}

	for _, v := range envs {
		ae := &AppEnviroment{
			AppID:           a.AppId,
			AppCode:         a.AppCode,
			EnvironmentName: v,
		}
		if err = tx.Create(ae).Error; err != nil {
			return err
		}
	}

	tx.Commit()
	return nil
}

// 获取所有记录
func (e *App) Finds() ([]*App, error) {

	var apps []*App
	err := Db.Preload("Admin").Find(&apps).Error

	if err != nil {
		return nil, err
	}
	return apps, nil
}

// 查找一条记录
func (e *App) Find() error {

	err := Db.Preload("Admin").First(&e, e.AppId).Error
	if err != nil {
		return err
	}

	return nil
}
