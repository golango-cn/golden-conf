// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: config.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golango.cn/golden-conf/conf-proto/common"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for ConfigServer service

func NewConfigServerEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for ConfigServer service

type ConfigServerService interface {
	GetConfig(ctx context.Context, in *GetConfigRequ, opts ...client.CallOption) (*GetConfigResp, error)
	GetAppConfigs(ctx context.Context, in *GetAppConfigsRequ, opts ...client.CallOption) (*GetAppConfigsResp, error)
	CreateAppConfig(ctx context.Context, in *CreateAppConfigRequ, opts ...client.CallOption) (*CreateAppConfigResp, error)
	UpdateAppConfig(ctx context.Context, in *UpdateAppConfigRequ, opts ...client.CallOption) (*UpdateAppConfigResp, error)
	GetAppConfigHistorys(ctx context.Context, in *GetAppConfigHistorysRequ, opts ...client.CallOption) (*GetAppConfigHistorysResp, error)
	RollbackConfig(ctx context.Context, in *RollbackConfigRequ, opts ...client.CallOption) (*RollbackConfigResp, error)
	PublicConfig(ctx context.Context, in *PublicConfigRequ, opts ...client.CallOption) (*PublicConfigResp, error)
}

type configServerService struct {
	c    client.Client
	name string
}

func NewConfigServerService(name string, c client.Client) ConfigServerService {
	return &configServerService{
		c:    c,
		name: name,
	}
}

func (c *configServerService) GetConfig(ctx context.Context, in *GetConfigRequ, opts ...client.CallOption) (*GetConfigResp, error) {
	req := c.c.NewRequest(c.name, "ConfigServer.GetConfig", in)
	out := new(GetConfigResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServerService) GetAppConfigs(ctx context.Context, in *GetAppConfigsRequ, opts ...client.CallOption) (*GetAppConfigsResp, error) {
	req := c.c.NewRequest(c.name, "ConfigServer.GetAppConfigs", in)
	out := new(GetAppConfigsResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServerService) CreateAppConfig(ctx context.Context, in *CreateAppConfigRequ, opts ...client.CallOption) (*CreateAppConfigResp, error) {
	req := c.c.NewRequest(c.name, "ConfigServer.CreateAppConfig", in)
	out := new(CreateAppConfigResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServerService) UpdateAppConfig(ctx context.Context, in *UpdateAppConfigRequ, opts ...client.CallOption) (*UpdateAppConfigResp, error) {
	req := c.c.NewRequest(c.name, "ConfigServer.UpdateAppConfig", in)
	out := new(UpdateAppConfigResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServerService) GetAppConfigHistorys(ctx context.Context, in *GetAppConfigHistorysRequ, opts ...client.CallOption) (*GetAppConfigHistorysResp, error) {
	req := c.c.NewRequest(c.name, "ConfigServer.GetAppConfigHistorys", in)
	out := new(GetAppConfigHistorysResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServerService) RollbackConfig(ctx context.Context, in *RollbackConfigRequ, opts ...client.CallOption) (*RollbackConfigResp, error) {
	req := c.c.NewRequest(c.name, "ConfigServer.RollbackConfig", in)
	out := new(RollbackConfigResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServerService) PublicConfig(ctx context.Context, in *PublicConfigRequ, opts ...client.CallOption) (*PublicConfigResp, error) {
	req := c.c.NewRequest(c.name, "ConfigServer.PublicConfig", in)
	out := new(PublicConfigResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ConfigServer service

type ConfigServerHandler interface {
	GetConfig(context.Context, *GetConfigRequ, *GetConfigResp) error
	GetAppConfigs(context.Context, *GetAppConfigsRequ, *GetAppConfigsResp) error
	CreateAppConfig(context.Context, *CreateAppConfigRequ, *CreateAppConfigResp) error
	UpdateAppConfig(context.Context, *UpdateAppConfigRequ, *UpdateAppConfigResp) error
	GetAppConfigHistorys(context.Context, *GetAppConfigHistorysRequ, *GetAppConfigHistorysResp) error
	RollbackConfig(context.Context, *RollbackConfigRequ, *RollbackConfigResp) error
	PublicConfig(context.Context, *PublicConfigRequ, *PublicConfigResp) error
}

func RegisterConfigServerHandler(s server.Server, hdlr ConfigServerHandler, opts ...server.HandlerOption) error {
	type configServer interface {
		GetConfig(ctx context.Context, in *GetConfigRequ, out *GetConfigResp) error
		GetAppConfigs(ctx context.Context, in *GetAppConfigsRequ, out *GetAppConfigsResp) error
		CreateAppConfig(ctx context.Context, in *CreateAppConfigRequ, out *CreateAppConfigResp) error
		UpdateAppConfig(ctx context.Context, in *UpdateAppConfigRequ, out *UpdateAppConfigResp) error
		GetAppConfigHistorys(ctx context.Context, in *GetAppConfigHistorysRequ, out *GetAppConfigHistorysResp) error
		RollbackConfig(ctx context.Context, in *RollbackConfigRequ, out *RollbackConfigResp) error
		PublicConfig(ctx context.Context, in *PublicConfigRequ, out *PublicConfigResp) error
	}
	type ConfigServer struct {
		configServer
	}
	h := &configServerHandler{hdlr}
	return s.Handle(s.NewHandler(&ConfigServer{h}, opts...))
}

type configServerHandler struct {
	ConfigServerHandler
}

func (h *configServerHandler) GetConfig(ctx context.Context, in *GetConfigRequ, out *GetConfigResp) error {
	return h.ConfigServerHandler.GetConfig(ctx, in, out)
}

func (h *configServerHandler) GetAppConfigs(ctx context.Context, in *GetAppConfigsRequ, out *GetAppConfigsResp) error {
	return h.ConfigServerHandler.GetAppConfigs(ctx, in, out)
}

func (h *configServerHandler) CreateAppConfig(ctx context.Context, in *CreateAppConfigRequ, out *CreateAppConfigResp) error {
	return h.ConfigServerHandler.CreateAppConfig(ctx, in, out)
}

func (h *configServerHandler) UpdateAppConfig(ctx context.Context, in *UpdateAppConfigRequ, out *UpdateAppConfigResp) error {
	return h.ConfigServerHandler.UpdateAppConfig(ctx, in, out)
}

func (h *configServerHandler) GetAppConfigHistorys(ctx context.Context, in *GetAppConfigHistorysRequ, out *GetAppConfigHistorysResp) error {
	return h.ConfigServerHandler.GetAppConfigHistorys(ctx, in, out)
}

func (h *configServerHandler) RollbackConfig(ctx context.Context, in *RollbackConfigRequ, out *RollbackConfigResp) error {
	return h.ConfigServerHandler.RollbackConfig(ctx, in, out)
}

func (h *configServerHandler) PublicConfig(ctx context.Context, in *PublicConfigRequ, out *PublicConfigResp) error {
	return h.ConfigServerHandler.PublicConfig(ctx, in, out)
}
