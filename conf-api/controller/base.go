package controller

import (
	"github.com/golango.cn/golden-conf/conf-api/common"
	"github.com/golango.cn/golden-conf/conf-api/conf"
	proto "github.com/golango.cn/golden-conf/conf-proto"
)

var config *conf.Config

var appServer proto.AppServerService
var configServer proto.ConfigServerService
var enviromentServer proto.EnvironmentServerService

func Init() {

	config = conf.GetConfig()

	appServer = config.Services[common.APP_SERVER].(proto.AppServerService)
	configServer = config.Services[common.CONFIG_SERVER].(proto.ConfigServerService)
	enviromentServer = config.Services[common.ENVIRONMENT_SERVER].(proto.EnvironmentServerService)

}
