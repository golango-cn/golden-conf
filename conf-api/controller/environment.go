package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	proto "github.com/golango.cn/golden-conf/conf-proto"
	"net/http"
)

type EnvironmentController struct {
}

// 获取所有环境
func (EnvironmentController) Gets(c *gin.Context) {

	resp, err := enviromentServer.GetEnvironments(context.Background(), &proto.GetEnvironmentsRequ{
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

// 创建环境
func (EnvironmentController) Post(c *gin.Context) {

	var requ = struct {
		Environments []string `form:"environments"        binding:"required"`
	}{}

	if err := c.Bind(&requ); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	resp, err := enviromentServer.CreateEnvironments(context.Background(), &proto.CreateEnvironmentsRequ{
		Environments: requ.Environments,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)

}
