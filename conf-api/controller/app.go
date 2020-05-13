package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	proto "github.com/golango.cn/golden-conf/conf-proto"
	"net/http"
	"strconv"
)

type ErrorResponse struct {
	Error string
}

type AppController struct {
}

// 获取所有记录
func (AppController) Gets(c *gin.Context) {
	resp, err := appServer.GetApps(context.Background(), &proto.GetAppsRequ{})
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// 获取单条记录
func (AppController) Get(c *gin.Context) {

	appId, _ := strconv.ParseInt(c.Param("appId"), 10, 64)
	resp, err := appServer.GetApp(context.Background(), &proto.GetAppRequ{
		AppId: appId,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)

}

// 保存记录
func (AppController) Post(c *gin.Context) {

	var requ = struct {
		AppCode string `form:"appCode"        binding:"required"`
		AppDesc string `form:"appDesc"        binding:"required"`
		AppName string `form:"appName"        binding:"required"`
		Admin   struct {
			AdminId int64 `form:"adminId"        binding:"required"`
		} `form:"admin"        binding:"required"`
		Environments []string `from:"environments" binding:"required"`
	}{}
	if err := c.Bind(&requ); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}
	resp, err := appServer.CreateApp(context.Background(), &proto.CreateAppRequ{
		App: &proto.App{
			AppCode: requ.AppCode,
			AppDesc: requ.AppDesc,
			Admin: &proto.Admin{
				AdminId:
				requ.Admin.AdminId,
			},
			AppName: requ.AppName,
		},
		Environments: requ.Environments,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)

}
