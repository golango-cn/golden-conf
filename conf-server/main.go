package main

import (
	"github.com/golango.cn/golden-conf/conf-database/conf"
	"github.com/golango.cn/golden-conf/conf-database/model"
	proto "github.com/golango.cn/golden-conf/conf-proto"
	"github.com/golango.cn/golden-conf/conf-server/hander"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
)

func main() {

	service := micro.NewService(
		micro.Name("go.micro.conf.server"),
		micro.Flags(
			&cli.StringFlag{
				Name:    "database_driver",
				EnvVars: []string{"DATABASE_DRIVER"},
				Usage:   "database_driver",
				Value:   "mysql",
			},
			&cli.StringFlag{
				Name:    "database_url",
				EnvVars: []string{"DATABASE_URL"},
				Usage:   "database_url",
				Value:   "127.0.0.1",
			}),
	)
	service.Init(
		micro.Action(func(c *cli.Context) error {
			return nil
		}),
		micro.BeforeStart(func() error {

			err := model.Init(&conf.Config{
				DB: conf.DataBase{
					DriverName: "mysql",
					URL:        "root:system@(127.0.0.1:3306)/nconf?charset=utf8&parseTime=true&loc=Local",
				},
				BroadcastType: "",
			})
			if err != nil {
				return err
			}
			if err := model.Db.DB().Ping(); err != nil {
				return err
			}
			return nil
		}),
		micro.BeforeStop(func() error {
			model.Db.Close()
			return nil
		}),
	)

	if err := proto.RegisterAppServerHandler(service.Server(), new(hander.AppHandler)); err != nil {
		panic(err)
	}
	if err := proto.RegisterConfigServerHandler(service.Server(), new(hander.ConfigHandler)); err != nil {
		panic(err)
	}
	if err := proto.RegisterEnvironmentServerHandler(service.Server(), new(hander.EnvironmentHandler)); err != nil {
		panic(err)
	}


	if err := service.Run(); err != nil {
		panic(err)
	}
}
