package test

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"testing"
)

// 用户
type Users struct {
	UserID int64  `gorm:"primary_key;"`
	Name   string `gorm:"not null"`
}

// 文章
type Topics struct {
	TopicId    int        `gorm:"primary_key"`
	Title      string     `gorm:"not null"`
	CategoryId int        `gorm:"not null"`
	Category   Categories `gorm:"foreignkey:CategoryId"` //文章所属分类外键
	UserId     int64
	User       Users `gorm:"foreignkey:UserId;association_foreignkey:UserId"` //文章所属用户外键
}

// 分类
type Categories struct {
	CategoryId int    `gorm:"primary_key"`
	Name       string `gorm:"not null"`
}

func GetDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:system@(127.0.0.1:3306)/nconf?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		fmt.Println("db error:", err)
	} else {
		fmt.Println("database connection success")
	}
	//defer db.Close()
	return db
}

func Test_Main(t *testing.T) {


	db := GetDB()
	db.LogMode(true)

	models := []interface{}{
		&Topics{},
		&Users{},
		&Categories{},
	}
	//1.执行建表语句
	err := db.Debug().AutoMigrate(models...).Error
	if err != nil {
		fmt.Println("db error:", err)
	}
	//2.执行sql
	//INSERT INTO topics("id", "title", "user_id", "category_id") VALUES (1, '测试', 1, 1);
	//INSERT INTO categories("id", "name") VALUES (1, '测试分类');
	//INSERT INTO users("id", "name") VALUES (1, '测试用户');

	//3.执行预加载
	topics, err := GetTopicsById(db, 1)
	if err != nil {
		fmt.Println("get topics error:", err)
	}
	fmt.Println(topics)

	////
	//type tts struct {
	//	TopicId int64
	//	UserId int64
	//	Name string
	//}
	//var topics []tts
	//
	//db = db.Table("topics a").
	//	Select("a.topic_id, b.user_id, b.name").
	//	Joins("inner join users b on a.user_id = b.user_id")
	//
	//var cnt int64
	//db.Count(&cnt)
	//
	//err  = db.Scan(&topics).Limit(10).Offset(1).Error
	//
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Println(topics)


	//db = db.Raw(`select a.topic_id, b.user_id, b.name
	//			from topics a inner join users b on a.user_id = b.user_id`)
	//
	//var cnt int64
	//db.Count(&cnt)
	//
	//err  = db.Scan(&topics).Limit(10).Offset(1).Error
	//
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Println(topics)
}

func GetTopicsById(db *gorm.DB, id int) (*Topics, error) {
	var topic Topics
	//查询方法1
	//err := db.Model(&topic).Where("id=?", id).First(&topic).
	//	Related(&topic.Category, "CategoryId").
	//	Related(&topic.User, "UserId").Error

	//查询方法2
	err := db.Where("topic_id=?", id).
		Preload("Category").
		Preload("User").
		First(&topic).Error
	if err != nil {
		return nil, err
	}
	return &topic, nil
}
