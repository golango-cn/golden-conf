package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type AppConfigHistory struct {
	AppConfigHistoryId int64 `gorm:"primary_key;AUTO_INCREMENT"` // 主键
	AppConfigId        int64
	AppId              int64  // 应用编号
	AppCode            string // 应用标识
	EnvironmentName    string // 环境
	Key                string // 键
	Value              string // 值
	Status             int64  // 状态 0未发布 1已发布
	Memo               string // 备注
	AdminId            int64  // 操作人
	Admin              Admin  `gorm:"foreignkey:AdminId;association_foreignkey:AdminId"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          *time.Time `sql:"index"`
}

// 表名
func (AppConfigHistory) TableName() string {
	return "app_config_history"
}

// 创建配置
func CloneAppConfig2History(tx *gorm.DB, appConfig *AppConfig) error {

	history := &AppConfigHistory{
		AppConfigId:     appConfig.AppConfigId,
		AppId:           appConfig.AppId,
		AppCode:         appConfig.AppCode,
		EnvironmentName: appConfig.EnvironmentName,
		Key:             appConfig.Key,
		Value:           appConfig.Value,
		Status:          appConfig.Status,
		Memo:            appConfig.Memo,
		AdminId:         appConfig.AdminId,
	}

	if err := tx.Create(history).Error; err != nil {
		return err
	}
	return nil

}

// 查找历史记录
func (a *AppConfigHistory) Finds(appConfigId int64, pageIndex, pageSize int64) ([]*AppConfigHistory, int64, error) {

	db := Db
	db = db.Where("app_config_id = ?", appConfigId)

	cnt := 0
	if err := db.Table(a.TableName()).Count(&cnt).Error; err != nil {
		return nil, 0, err
	}

	var historys []*AppConfigHistory
	err := db.Preload("Admin").Limit(pageSize).Offset((pageIndex - 1) * pageSize).
		Find(&historys).Error

	if err != nil {
		return nil, 0, err
	}
	return historys, int64(cnt), nil

}

// 回滚
func (a *AppConfigHistory) Rollback(adminId int64) error {

	if err := Db.Table(a.TableName()).First(&a, a.AppConfigHistoryId).Error; err != nil {
		return err
	}

	config := &AppConfig{
		AppConfigId: a.AppConfigId,
		Value:       a.Value,
		AdminId:     adminId,
	}
	err := config.Update()
	if err != nil {
		return err
	}

	return nil

}
