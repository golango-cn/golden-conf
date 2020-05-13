package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	proto "github.com/golango.cn/golden-conf/conf-proto"
	"github.com/golango.cn/golden-conf/conf-proto/common"
	"net/http"
	"strconv"
)

type ConfigController struct {
}

// 获取列表
func (ConfigController) Gets(c *gin.Context) {

	pageIndex, _ := strconv.ParseInt(c.DefaultQuery("pageIndex", "1"), 10, 64)
	pageSize, _ := strconv.ParseInt(c.DefaultQuery("pageSize", "20"), 10, 64)

	key := c.DefaultQuery("key", "")
	appCode := c.DefaultQuery("appCode", "")
	environmentName := c.DefaultQuery("environmentName", "")

	resp, err := configServer.GetAppConfigs(context.Background(), &proto.GetAppConfigsRequ{
		BaseRequ: &common.BaseRequ{
			PageIndex: pageIndex,
			PageSize:  pageSize,
		},
		Key:             key,
		AppCode:         appCode,
		EnvironmentName: environmentName,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)

}

// 获取单条记录
func (ConfigController) Get(c *gin.Context) {

	appConfigId, _ := strconv.ParseInt(c.Param("appConfigId"), 10, 64)
	resp, err := configServer.GetConfig(context.Background(), &proto.GetConfigRequ{
		AppConfigId: appConfigId,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)

}

// 发布配置
func (ConfigController) Publish(c *gin.Context) {

	var requ = struct {
		AppConfigId int64 `form:"appConfigId"        binding:"required"`
	}{}
	if err := c.Bind(&requ); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	resp, err := configServer.PublicConfig(context.Background(), &proto.PublicConfigRequ{
		AppConfigId: requ.AppConfigId,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// 更新配置
func (ConfigController) Update(c *gin.Context) {

	appConfigId, _ := strconv.ParseInt(c.Param("appConfigId"), 10, 64)
	var requ = struct {
		Value   string `form:"value"        binding:"required"`   // 值
		Memo    string `form:"memo"        binding:"required"`    // 备注
		AdminId int64  `form:"adminId"        binding:"required"` // 操作人
	}{}
	if err := c.Bind(&requ); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	resp, err := configServer.UpdateAppConfig(context.Background(), &proto.UpdateAppConfigRequ{
		AppConfig: &proto.AppConfig{
			AppConfigId: appConfigId,
			Value:       requ.Value,
			Memo:        requ.Memo,
			Admin: &proto.Admin{
				AdminId: requ.AdminId,
			},
		},
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// 新建配置
func (ConfigController) Post(c *gin.Context) {

	var requ = struct {
		AppId           int64  `form:"appId"        binding:"required"`           // 应用编号
		AppCode         string `form:"appCode"        binding:"required"`         // 应用标识
		EnvironmentName string `form:"environmentName"        binding:"required"` // 环境
		Key             string `form:"key"        binding:"required"`             // 键
		Value           string `form:"value"        binding:"required"`           // 值
		Memo            string `form:"memo"        binding:"required"`            // 备注
		AdminId         int64  `form:"adminId"        binding:"required"`         // 操作人
	}{}

	if err := c.Bind(&requ); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	resp, err := configServer.CreateAppConfig(context.Background(), &proto.CreateAppConfigRequ{
		AppConfig: &proto.AppConfig{
			AppId:           requ.AppId,
			AppCode:         requ.AppCode,
			EnvironmentName: requ.EnvironmentName,
			Key:             requ.Key,
			Value:           requ.Value,
			Memo:            requ.Memo,
			Admin: &proto.Admin{
				AdminId: requ.AdminId,
			},
		},
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)

}

// 获取历史记录
func (ConfigController) Historys(c *gin.Context) {

	appConfigId, _ := strconv.ParseInt(c.Param("appConfigId"), 10, 64)
	pageIndex, _ := strconv.ParseInt(c.DefaultQuery("pageIndex", "1"), 10, 64)
	pageSize, _ := strconv.ParseInt(c.DefaultQuery("pageSize", "20"), 10, 64)

	resp, err := configServer.GetAppConfigHistorys(context.Background(), &proto.GetAppConfigHistorysRequ{
		AppConfigId: appConfigId,
		BaseRequ: &common.BaseRequ{
			PageIndex: pageIndex,
			PageSize:  pageSize,
		},
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)

}

// 回滚
func (ConfigController) Rollback(c *gin.Context) {

	var requ = struct {
		AppConfigHistoryId int64 `form:"appConfigHistoryId"        binding:"required"` // 应用编号
		AdminId            int64 `form:"adminId"        binding:"required"`            // 操作人
	}{}

	if err := c.Bind(&requ); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	resp, err := configServer.RollbackConfig(context.Background(), &proto.RollbackConfigRequ{
		AppConfigHistoryId: requ.AppConfigHistoryId,
		AdminId:            requ.AdminId,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)

}
