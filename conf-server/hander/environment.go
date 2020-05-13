package hander

import (
	"context"
	"github.com/golango.cn/golden-conf/conf-database/model"
	proto "github.com/golango.cn/golden-conf/conf-proto"
	"github.com/golango.cn/golden-conf/conf-proto/common"
)

type EnvironmentHandler struct {
}

// 初始化环境
func (e EnvironmentHandler) CreateEnvironments(ctx context.Context, requ *proto.CreateEnvironmentsRequ, resp *proto.CreateEnvironmentsResp) error {

	environment := model.Environment{}
	err := environment.Create(requ.Environments)

	if err != nil {
		return err
	}

	resp.BaseResp = &common.BaseResp{
		IsSuccess:    true,
		ErrorMessage: "",
	}

	return nil

}

// 环境列表
func (e EnvironmentHandler) GetEnvironments(ctx context.Context, requ *proto.GetEnvironmentsRequ, resp *proto.GetEnvironmentsResp) error {

	environment := model.Environment{}
	results, err := environment.Finds()

	if err != nil {
		return err
	}

	for _, v := range results {
		resp.Environments = append(resp.Environments, &proto.Environment{
			EnvironmentID:   v.EnvironmentID,
			EnvironmentName: v.EnvironmentName,
		})
	}

	resp.BaseResp = &common.BaseResp{
		IsSuccess:    true,
		ErrorMessage: "",
	}

	return nil

}
