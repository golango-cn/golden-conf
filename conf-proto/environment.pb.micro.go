// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: environment.proto

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

// Api Endpoints for EnvironmentServer service

func NewEnvironmentServerEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for EnvironmentServer service

type EnvironmentServerService interface {
	GetEnvironments(ctx context.Context, in *GetEnvironmentsRequ, opts ...client.CallOption) (*GetEnvironmentsResp, error)
	CreateEnvironments(ctx context.Context, in *CreateEnvironmentsRequ, opts ...client.CallOption) (*CreateEnvironmentsResp, error)
}

type environmentServerService struct {
	c    client.Client
	name string
}

func NewEnvironmentServerService(name string, c client.Client) EnvironmentServerService {
	return &environmentServerService{
		c:    c,
		name: name,
	}
}

func (c *environmentServerService) GetEnvironments(ctx context.Context, in *GetEnvironmentsRequ, opts ...client.CallOption) (*GetEnvironmentsResp, error) {
	req := c.c.NewRequest(c.name, "EnvironmentServer.GetEnvironments", in)
	out := new(GetEnvironmentsResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *environmentServerService) CreateEnvironments(ctx context.Context, in *CreateEnvironmentsRequ, opts ...client.CallOption) (*CreateEnvironmentsResp, error) {
	req := c.c.NewRequest(c.name, "EnvironmentServer.CreateEnvironments", in)
	out := new(CreateEnvironmentsResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for EnvironmentServer service

type EnvironmentServerHandler interface {
	GetEnvironments(context.Context, *GetEnvironmentsRequ, *GetEnvironmentsResp) error
	CreateEnvironments(context.Context, *CreateEnvironmentsRequ, *CreateEnvironmentsResp) error
}

func RegisterEnvironmentServerHandler(s server.Server, hdlr EnvironmentServerHandler, opts ...server.HandlerOption) error {
	type environmentServer interface {
		GetEnvironments(ctx context.Context, in *GetEnvironmentsRequ, out *GetEnvironmentsResp) error
		CreateEnvironments(ctx context.Context, in *CreateEnvironmentsRequ, out *CreateEnvironmentsResp) error
	}
	type EnvironmentServer struct {
		environmentServer
	}
	h := &environmentServerHandler{hdlr}
	return s.Handle(s.NewHandler(&EnvironmentServer{h}, opts...))
}

type environmentServerHandler struct {
	EnvironmentServerHandler
}

func (h *environmentServerHandler) GetEnvironments(ctx context.Context, in *GetEnvironmentsRequ, out *GetEnvironmentsResp) error {
	return h.EnvironmentServerHandler.GetEnvironments(ctx, in, out)
}

func (h *environmentServerHandler) CreateEnvironments(ctx context.Context, in *CreateEnvironmentsRequ, out *CreateEnvironmentsResp) error {
	return h.EnvironmentServerHandler.CreateEnvironments(ctx, in, out)
}
