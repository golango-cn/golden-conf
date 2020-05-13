package hander

import (
	"context"
	"github.com/golango.cn/golden-conf/conf-database/model"
	proto "github.com/golango.cn/golden-conf/conf-proto"
	"github.com/golango.cn/golden-conf/conf-proto/common"
)

type AppHandler struct {
}

// 获取单条记录
func (a AppHandler) GetApp(ctx context.Context, requ *proto.GetAppRequ, resp *proto.GetAppResp) error {

	app := &model.App{AppId: requ.AppId}
	err := app.Find()

	if err != nil {
		return err
	}

	resp.BaseResp = &common.BaseResp{
		IsSuccess:    true,
		ErrorMessage: "",
	}
	resp.App = &proto.App{
		AppId:   app.AppId,
		AppCode: app.AppCode,
		AppDesc: app.AppDesc,
		Admin: &proto.Admin{
			AdminId:   app.Admin.AdminId,
			AdminName: app.Admin.AdminName,
		},
		CreateAt: app.CreatedAt.Unix(),
		UpdateAt: app.UpdatedAt.Unix(),
	}
	return nil
}

// 获取列表
func (a AppHandler) GetApps(ctx context.Context, requ *proto.GetAppsRequ, resp *proto.GetAppsResp) error {

	app := &model.App{}
	results, err := app.Finds()

	if err != nil {
		return err
	}

	resp.BaseResp = &common.BaseResp{
		IsSuccess:    true,
		ErrorMessage: "",
	}

	for _, v := range results {
		resp.Apps = append(resp.Apps, &proto.App{
			AppId:   v.AppId,
			AppCode: v.AppCode,
			AppDesc: v.AppDesc,
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

// 创建App项目
func (a AppHandler) CreateApp(ctx context.Context, requ *proto.CreateAppRequ, resp *proto.CreateAppResp) error {

	app := &model.App{
		AppCode: requ.App.AppCode,
		AppDesc: requ.App.AppDesc,
		AdminId: requ.App.Admin.AdminId,
		AppName: requ.App.AppName,
	}
	err := app.Create(requ.Environments)

	if err != nil {
		return err
	}

	resp.BaseResp = &common.BaseResp{
		IsSuccess:    true,
		ErrorMessage: "",
	}
	resp.AppId = app.AppId

	return nil
}
