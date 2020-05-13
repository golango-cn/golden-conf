package hander

import (
	"context"
	"github.com/golango.cn/golden-conf/conf-database/model"
	proto "github.com/golango.cn/golden-conf/conf-proto"
	"github.com/golango.cn/golden-conf/conf-proto/common"
)

type ConfigHandler struct {
}

// 获取单个配置记录
func (c ConfigHandler) GetConfig(ctx context.Context, requ *proto.GetConfigRequ, resp *proto.GetConfigResp) error {

	config := &model.AppConfig{
		AppConfigId: requ.AppConfigId,
	}

	if err := config.Find(); err != nil {
		return err
	}

	resp.BaseResp = &common.BaseResp{
		IsSuccess:    true,
		ErrorMessage: "",
	}
	resp.AppConfig = &proto.AppConfig{
		AppConfigId:     config.AppConfigId,
		AppId:           config.AppId,
		AppCode:         config.AppCode,
		EnvironmentName: config.EnvironmentName,
		Key:             config.Key,
		Value:           config.Value,
		Status:          config.Status,
		Memo:            config.Memo,
		Admin: &proto.Admin{
			AdminId:   config.Admin.AdminId,
			AdminName: config.Admin.AdminName,
		},
		CreateAt: config.CreatedAt.Unix(),
		UpdateAt: config.UpdatedAt.Unix(),
	}

	return nil
}

// 获取配置记录
func (c ConfigHandler) GetAppConfigs(ctx context.Context, requ *proto.GetAppConfigsRequ, resp *proto.GetAppConfigsResp) error {

	config := &model.AppConfig{}
	results, totalCount, err := config.Finds(
		requ.Key,
		requ.AppCode,
		requ.EnvironmentName,
		requ.BaseRequ.PageIndex,
		requ.BaseRequ.PageSize,
	)

	if err != nil {
		return err
	}

	resp.BaseResp = &common.BaseResp{
		IsSuccess:    true,
		ErrorMessage: "",
		TotalCount:   totalCount,
	}

	for _, v := range results {
		resp.AppConfigs = append(resp.AppConfigs, &proto.AppConfig{
			AppConfigId:     v.AppConfigId,
			AppId:           v.AppId,
			AppCode:         v.AppCode,
			EnvironmentName: v.EnvironmentName,
			Key:             v.Key,
			Value:           v.Value,
			Status:          v.Status,
			Memo:            v.Memo,
			Admin: &proto.Admin{
				AdminId:   v.Admin.AdminId,
				AdminName: v.Admin.AdminName,
			},
			CreateAt: v.CreatedAt.Unix(),
			UpdateAt: v.UpdatedAt.Unix(),
		})
	}

	return nil
}

// 创建配置
func (c ConfigHandler) CreateAppConfig(ctx context.Context, requ *proto.CreateAppConfigRequ, resp *proto.CreateAppConfigResp) error {

	config := model.AppConfig{
		AppId:           requ.AppConfig.AppId,
		AppCode:         requ.AppConfig.AppCode,
		EnvironmentName: requ.AppConfig.EnvironmentName,
		Key:             requ.AppConfig.Key,
		Value:           requ.AppConfig.Value,
		Status:          requ.AppConfig.Status,
		Memo:            requ.AppConfig.Memo,
		AdminId:         requ.AppConfig.Admin.AdminId,
	}
	err := config.Create()
	if err != nil {
		return err
	}

	resp.BaseResp = &common.BaseResp{
		IsSuccess:    true,
		ErrorMessage: "",
	}

	return nil
}

// 更新配置
func (c ConfigHandler) UpdateAppConfig(ctx context.Context, requ *proto.UpdateAppConfigRequ, resp *proto.UpdateAppConfigResp) error {

	config := model.AppConfig{
		Value:       requ.AppConfig.Value,
		Memo:        requ.AppConfig.Memo,
		AdminId:     requ.AppConfig.Admin.AdminId,
		AppConfigId: requ.AppConfig.AppConfigId,
	}
	err := config.Update()
	if err != nil {
		return err
	}

	resp.BaseResp = &common.BaseResp{
		IsSuccess:    true,
		ErrorMessage: "",
	}

	return nil
}

// 获取配置历史记录
func (c ConfigHandler) GetAppConfigHistorys(ctx context.Context, requ *proto.GetAppConfigHistorysRequ, resp *proto.GetAppConfigHistorysResp) error {


	history := &model.AppConfigHistory{}
	results, totalCount, err := history.Finds(requ.AppConfigId, requ.BaseRequ.PageIndex, requ.BaseRequ.PageSize)

	if err != nil {
		return err
	}

	resp.BaseResp = &common.BaseResp{
		IsSuccess:    true,
		ErrorMessage: "",
		TotalCount:   totalCount,
	}

	for _, v := range results {
		resp.AppConfigHistorys = append(resp.AppConfigHistorys, &proto.AppConfigHistory{
			AppConfigHistoryId: v.AppConfigHistoryId,
			AppConfigId:        v.AppConfigId,
			AppId:              v.AppId,
			AppCode:            v.AppCode,
			EnvironmentName:    v.EnvironmentName,
			Key:                v.Key,
			Value:              v.Value,
			Status:             v.Status,
			Memo:               v.Memo,
			Admin: &proto.Admin{
				AdminId:   v.Admin.AdminId,
				AdminName: v.Admin.AdminName,
			},
			CreateAt: v.CreatedAt.Unix(),
			UpdateAt: v.UpdatedAt.Unix(),
		})
	}

	return nil

}

// 回滚配置
func (c ConfigHandler) RollbackConfig(ctx context.Context, requ *proto.RollbackConfigRequ, resp *proto.RollbackConfigResp) error {

	history := &model.AppConfigHistory{
		AppConfigHistoryId: requ.AppConfigHistoryId,
	}
	err := history.Rollback(requ.AdminId)

	if err != nil {
		return err
	}

	resp.BaseResp = &common.BaseResp{
		IsSuccess:    true,
		ErrorMessage: "",
	}
	return nil
}

// 发布配置
func (c ConfigHandler) PublicConfig(ctx context.Context, requ *proto.PublicConfigRequ, resp *proto.PublicConfigResp) error {

	config := &model.AppConfig{
		AppConfigId: requ.AppConfigId,
	}
	err := config.Publish()

	if err != nil {
		return err
	}

	resp.BaseResp = &common.BaseResp{
		IsSuccess:    true,
		ErrorMessage: "",
	}
	return nil

}
