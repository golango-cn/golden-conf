package model

import (
	"fmt"
	"strings"
	"time"
)

type AppConfig struct {
	AppConfigId     int64  `gorm:"primary_key;AUTO_INCREMENT"` // 主键
	AppId           int64  // 应用编号
	AppCode         string // 应用标识
	EnvironmentName string // 环境
	Key             string // 键
	Value           string // 值
	Status          int64  // 状态 0未发布 1已发布
	Memo            string // 备注
	AdminId         int64  // 操作人
	Admin           Admin  `gorm:"foreignkey:AdminId;association_foreignkey:AdminId"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time `sql:"index"`
}

func (AppConfig) TableName() string {
	return "app_config"
}

// 创建配置
func (a *AppConfig) Create() error {

	var cnt int32
	if err := Db.Table(a.TableName()).
		Where("app_id = ? and environment_name = ? and `key` = ?", a.AppId, a.EnvironmentName, a.Key).
		Count(&cnt).Error; err != nil {
		return err
	}

	if cnt > 0 {
		return fmt.Errorf("[CreateAppConfig]The same configuration exists for the current environment")
	}

	var err error
	tx := Db.Begin()
	defer func() {
		if err != nil {
			err = fmt.Errorf("[CreateAppConfig]Create app:%s-%s-%s error: %s", a.Key, a.EnvironmentName, a.AppCode, err.Error())
			tx.Rollback()
		}
	}()
	if err = tx.Create(a).Error; err != nil {
		return err
	}
	if err = CloneAppConfig2History(tx, a); err != nil {
		return err
	}

	tx.Commit()
	return nil
}

// 发布配置
func (a *AppConfig) Publish() error {
	return Db.Model(&a).Update("status", 1).Error
}

// 更新配置
func (a *AppConfig) Update() error {

	var err error
	tx := Db.Begin()
	defer func() {
		if err != nil {
			err = fmt.Errorf("[UpdateAppConfig]Create app:%s-%s-%s error: %s", a.Key, a.EnvironmentName, a.AppCode, err.Error())
			tx.Rollback()
		}
	}()

	if err = tx.Model(&a).Updates(map[string]interface{}{
		"status":   0,
		"value":    a.Value,
		"memo":     a.Memo,
		"admin_id": a.AdminId,
	}).Error; err != nil {
		return err
	}
	if err = tx.First(&a, a.AppConfigId).Error; err != nil {
		return err
	}
	if err = CloneAppConfig2History(tx, a); err != nil {
		return err
	}

	tx.Commit()
	return err
}

// 查询记录
func (a *AppConfig) Finds(key, appCode, environmentName string, pageIndex, pageSize int64) ([]*AppConfig, int64, error) {

	db := Db
	if len(strings.TrimSpace(key)) > 0 {
		db = db.Where("`key` like ?", fmt.Sprintf("%%%s%%", key))
	}
	if len(appCode) > 0 {
		db = db.Where("app_code = ?", appCode)
	}
	if len(environmentName) > 0 {
		db = db.Where("environment_name = ?", environmentName)
	}

	var cnt int64
	if err := db.Table(a.TableName()).Count(&cnt).Error; err != nil {
		return nil, 0, err
	}

	var configs []*AppConfig
	if err := db.Table(a.TableName()).Preload("Admin").Limit(pageSize).Offset((pageIndex - 1) * pageSize).Find(&configs).Error; err != nil {
		return nil, 0, err
	}
	return configs, cnt, nil

}

// 查找记录
func (a *AppConfig) Find() error {
	err := Db.Table(a.TableName()).
		Preload("Admin").
		First(&a, a.AppConfigId).Error
	if err != nil {
		return err
	}
	return nil
}
