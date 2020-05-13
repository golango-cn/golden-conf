package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/golango.cn/golden-conf/conf-database/conf"
	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

func Init(c *conf.Config) error {

	var db *gorm.DB
	var err error

	if db, err = gorm.Open(c.DB.DriverName, c.DB.URL); err != nil {
		return err
	}
	db.SingularTable(true)       //表名采用单数形式
	db.DB().SetMaxOpenConns(100) //SetMaxOpenConns用于设置最大打开的连接数
	db.DB().SetMaxIdleConns(10)  //SetMaxIdleConns用于设置闲置的连接数
	db.LogMode(true)

	if err = db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
		&Environment{},
		&App{},
		&AppEnviroment{},
		&AppConfig{},
		&AppConfigHistory{},
		&Admin{},
	).Error; err != nil {
		_ = db.Close()
		return err
	}

	Db = db
	return nil
}
