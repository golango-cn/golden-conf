package test

import (
	"fmt"
	"github.com/golango.cn/golden-conf/conf-database/model"
	"testing"
)

func TestApp(t *testing.T) {

	app := &model.App{
		AppCode: "member-core",
		AppName: "会员组核心项目",
		AppDesc: "会员组核心项目",
		AdminId: 3,
	}

	err := app.Create([]string{"env", "dev"})
	fmt.Println(err)

}

func TestFind(t *testing.T) {

	app := &model.App{}
	finds, err := app.Finds()

	fmt.Println(finds, err)

	for _, v := range finds {
		fmt.Printf("%#v\n", v)
		fmt.Println(v.CreatedAt.Unix())
	}
}
