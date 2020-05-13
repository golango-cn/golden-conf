package model

import (
	"strings"
	"time"
)

type Environment struct {
	EnvironmentID   int64 `gorm:"primary_key;AUTO_INCREMENT"`
	EnvironmentName string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time `sql:"index"`
}

func (Environment) TableName() string {
	return "environment"
}

// 初始化环境数据
func (e *Environment) Create(envs []string) error {
	if len(envs) == 0 {
		return nil
	}
	for _, v := range envs {
		if len(strings.TrimSpace(v)) == 0 {
			continue
		}
		var cnt int32
		if err := Db.Table(e.TableName()).
			Where("environment_name = ?", v).
			Count(&cnt).Error; err != nil {
			return err
		}
		if cnt == 0 {
			if err := Db.Create(&Environment{
				EnvironmentName: v,
				CreatedAt:       time.Now(),
				UpdatedAt:       time.Now(),
			}).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

// 获取所有记录
func (e *Environment) Finds() ([]*Environment, error) {

	var envs []*Environment
	err := Db.Table(e.TableName()).Find(&envs).Error

	if err != nil {
		return nil, err
	}
	return envs, nil
}
