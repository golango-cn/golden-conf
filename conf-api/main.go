package main

import (
	"github.com/gin-gonic/gin"
	"github.com/golango.cn/golden-conf/conf-api/common"
	"github.com/golango.cn/golden-conf/conf-api/conf"
	"github.com/golango.cn/golden-conf/conf-api/controller"
	proto "github.com/golango.cn/golden-conf/conf-proto"
	"github.com/micro/go-micro/v2/web"
	"log"
)

func main() {

	// 创建新服务
	service := web.NewService(
		web.Name("go.micro.conf.web.api"),
		web.Version("latest"),
		//web.Registry(micReg),
		web.Address(":8091"),
	)

	// 初始化服务
	if err := service.Init(); err != nil {
		panic(err)
	}

	config := conf.New()
	// 注册客户端
	if err := config.Register("go.micro.conf.server", service,
		func(name string, service web.Service, c *conf.Config) error {

			appServer := proto.NewAppServerService(name, service.Options().Service.Client())
			c.Services[common.APP_SERVER] = appServer

			envServer := proto.NewEnvironmentServerService(name, service.Options().Service.Client())
			c.Services[common.ENVIRONMENT_SERVER] = envServer

			cfgServer := proto.NewConfigServerService(name, service.Options().Service.Client())
			c.Services[common.CONFIG_SERVER] = cfgServer

			return nil
		}); err != nil {
		panic(err)

	}

	controller.Init()

	r := Router()
	service.Handle("/", r)

	// 运行服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func Router() *gin.Engine {

	router := gin.Default()

	// 注意，这里的/api一定要与服务go.micro.conf.web.api中的api对应，否则网关无法路由
	r := router.Group("/api")

	appController := controller.AppController{}

	r.GET("/app", appController.Gets)
	r.GET("/app/:appId", appController.Get)
	r.POST("/app", appController.Post)

	configController := controller.ConfigController{}

	r.GET("/config", configController.Gets)
	r.POST("/config", configController.Post)
	r.POST("/config/publish", configController.Publish)
	r.GET("/config/:appConfigId", configController.Get)
	r.GET("/config/:appConfigId/historys", configController.Historys)
	r.POST("/config/rollback", configController.Rollback)
	r.PUT("/config/:appConfigId", configController.Update)

	environmentController := controller.EnvironmentController{}

	r.GET("/environment", environmentController.Gets)
	r.POST("/environment", environmentController.Post)

	return router
}
