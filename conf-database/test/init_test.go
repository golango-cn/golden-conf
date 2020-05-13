package test

import (
	"fmt"
	"github.com/golango.cn/golden-conf/conf-database/conf"
	"github.com/golango.cn/golden-conf/conf-database/model"
	"testing"
)

func init() {
	err := model.Init(&conf.Config{
		DB: conf.DataBase{
			DriverName: "mysql",
			URL:        "root:system@(127.0.0.1:3306)/nconf?charset=utf8&parseTime=true&loc=Local",
		},
		BroadcastType: "",
	})
	fmt.Println(err)
}

func TestInit(t *testing.T) {

	env := &model.Environment{}
	err := env.FirstOrInit([]string{"pro", "dev"})
	fmt.Println(err)
}
