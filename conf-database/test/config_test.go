package test

import (
	"fmt"
	"github.com/golango.cn/golden-conf/conf-database/model"
	"testing"
)

func TestConfig(t *testing.T) {

	config := &model.AppConfig{
		AppId:           1,
		AppCode:         "member-core",
		EnvironmentName: "dev",
		Key:             "sql_driver1",
		Value:           "mysql",
		Memo:            "mysql驱动",
		AdminId:         3,
	}
	err := config.Create()
	fmt.Println(err)

}

func TestPublish(t *testing.T) {
	config := &model.AppConfig{
		AppConfigId: 3,
	}

	err := config.Publish()
	fmt.Println(err)
}

func TestUpdate(t *testing.T) {
	config := &model.AppConfig{
		AppConfigId: 3,
		Value:       "mysql",
		AdminId:     2,
	}
	err := config.Update()
	fmt.Println(err)
}

func TestFinds2(t *testing.T) {

	config := new(model.AppConfig)
	results, i, err := config.Finds("d", "", "", 1, 50)

	fmt.Println(results, i, err)

}

func TestRollback(t *testing.T) {
	history := new(model.AppConfigHistory)
	err := history.Rollback(1, 1)
	fmt.Println(err)
}
