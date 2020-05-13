// Code generated by protoc-gen-go. DO NOT EDIT.
// source: environment.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	common "github.com/golango.cn/golden-conf/conf-proto/common"
	math "math"
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

type GetEnvironmentsRequ struct {
	BaseRequ             *common.BaseRequ `protobuf:"bytes,1,opt,name=baseRequ,proto3" json:"baseRequ,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GetEnvironmentsRequ) Reset()         { *m = GetEnvironmentsRequ{} }
func (m *GetEnvironmentsRequ) String() string { return proto.CompactTextString(m) }
func (*GetEnvironmentsRequ) ProtoMessage()    {}
func (*GetEnvironmentsRequ) Descriptor() ([]byte, []int) {
	return fileDescriptor_64e647b85623514a, []int{0}
}

func (m *GetEnvironmentsRequ) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEnvironmentsRequ.Unmarshal(m, b)
}
func (m *GetEnvironmentsRequ) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEnvironmentsRequ.Marshal(b, m, deterministic)
}
func (m *GetEnvironmentsRequ) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEnvironmentsRequ.Merge(m, src)
}
func (m *GetEnvironmentsRequ) XXX_Size() int {
	return xxx_messageInfo_GetEnvironmentsRequ.Size(m)
}
func (m *GetEnvironmentsRequ) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEnvironmentsRequ.DiscardUnknown(m)
}

var xxx_messageInfo_GetEnvironmentsRequ proto.InternalMessageInfo

func (m *GetEnvironmentsRequ) GetBaseRequ() *common.BaseRequ {
	if m != nil {
		return m.BaseRequ
	}
	return nil
}

type GetEnvironmentsResp struct {
	BaseResp             *common.BaseResp `protobuf:"bytes,1,opt,name=baseResp,proto3" json:"baseResp,omitempty"`
	Environments         []*Environment   `protobuf:"bytes,3,rep,name=environments,proto3" json:"environments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GetEnvironmentsResp) Reset()         { *m = GetEnvironmentsResp{} }
func (m *GetEnvironmentsResp) String() string { return proto.CompactTextString(m) }
func (*GetEnvironmentsResp) ProtoMessage()    {}
func (*GetEnvironmentsResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_64e647b85623514a, []int{1}
}

func (m *GetEnvironmentsResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEnvironmentsResp.Unmarshal(m, b)
}
func (m *GetEnvironmentsResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEnvironmentsResp.Marshal(b, m, deterministic)
}
func (m *GetEnvironmentsResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEnvironmentsResp.Merge(m, src)
}
func (m *GetEnvironmentsResp) XXX_Size() int {
	return xxx_messageInfo_GetEnvironmentsResp.Size(m)
}
func (m *GetEnvironmentsResp) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEnvironmentsResp.DiscardUnknown(m)
}

var xxx_messageInfo_GetEnvironmentsResp proto.InternalMessageInfo

func (m *GetEnvironmentsResp) GetBaseResp() *common.BaseResp {
	if m != nil {
		return m.BaseResp
	}
	return nil
}

func (m *GetEnvironmentsResp) GetEnvironments() []*Environment {
	if m != nil {
		return m.Environments
	}
	return nil
}

type Environment struct {
	EnvironmentID        int64    `protobuf:"varint,1,opt,name=environmentID,proto3" json:"environmentID,omitempty"`
	EnvironmentName      string   `protobuf:"bytes,2,opt,name=environmentName,proto3" json:"environmentName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Environment) Reset()         { *m = Environment{} }
func (m *Environment) String() string { return proto.CompactTextString(m) }
func (*Environment) ProtoMessage()    {}
func (*Environment) Descriptor() ([]byte, []int) {
	return fileDescriptor_64e647b85623514a, []int{2}
}

func (m *Environment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Environment.Unmarshal(m, b)
}
func (m *Environment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Environment.Marshal(b, m, deterministic)
}
func (m *Environment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Environment.Merge(m, src)
}
func (m *Environment) XXX_Size() int {
	return xxx_messageInfo_Environment.Size(m)
}
func (m *Environment) XXX_DiscardUnknown() {
	xxx_messageInfo_Environment.DiscardUnknown(m)
}

var xxx_messageInfo_Environment proto.InternalMessageInfo

func (m *Environment) GetEnvironmentID() int64 {
	if m != nil {
		return m.EnvironmentID
	}
	return 0
}

func (m *Environment) GetEnvironmentName() string {
	if m != nil {
		return m.EnvironmentName
	}
	return ""
}

type CreateEnvironmentsRequ struct {
	BaseRequ             *common.BaseRequ `protobuf:"bytes,1,opt,name=baseRequ,proto3" json:"baseRequ,omitempty"`
	Environments         []string         `protobuf:"bytes,2,rep,name=environments,proto3" json:"environments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *CreateEnvironmentsRequ) Reset()         { *m = CreateEnvironmentsRequ{} }
func (m *CreateEnvironmentsRequ) String() string { return proto.CompactTextString(m) }
func (*CreateEnvironmentsRequ) ProtoMessage()    {}
func (*CreateEnvironmentsRequ) Descriptor() ([]byte, []int) {
	return fileDescriptor_64e647b85623514a, []int{3}
}

func (m *CreateEnvironmentsRequ) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateEnvironmentsRequ.Unmarshal(m, b)
}
func (m *CreateEnvironmentsRequ) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateEnvironmentsRequ.Marshal(b, m, deterministic)
}
func (m *CreateEnvironmentsRequ) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateEnvironmentsRequ.Merge(m, src)
}
func (m *CreateEnvironmentsRequ) XXX_Size() int {
	return xxx_messageInfo_CreateEnvironmentsRequ.Size(m)
}
func (m *CreateEnvironmentsRequ) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateEnvironmentsRequ.DiscardUnknown(m)
}

var xxx_messageInfo_CreateEnvironmentsRequ proto.InternalMessageInfo

func (m *CreateEnvironmentsRequ) GetBaseRequ() *common.BaseRequ {
	if m != nil {
		return m.BaseRequ
	}
	return nil
}

func (m *CreateEnvironmentsRequ) GetEnvironments() []string {
	if m != nil {
		return m.Environments
	}
	return nil
}

type CreateEnvironmentsResp struct {
	BaseResp             *common.BaseResp `protobuf:"bytes,1,opt,name=baseResp,proto3" json:"baseResp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *CreateEnvironmentsResp) Reset()         { *m = CreateEnvironmentsResp{} }
func (m *CreateEnvironmentsResp) String() string { return proto.CompactTextString(m) }
func (*CreateEnvironmentsResp) ProtoMessage()    {}
func (*CreateEnvironmentsResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_64e647b85623514a, []int{4}
}

func (m *CreateEnvironmentsResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateEnvironmentsResp.Unmarshal(m, b)
}
func (m *CreateEnvironmentsResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateEnvironmentsResp.Marshal(b, m, deterministic)
}
func (m *CreateEnvironmentsResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateEnvironmentsResp.Merge(m, src)
}
func (m *CreateEnvironmentsResp) XXX_Size() int {
	return xxx_messageInfo_CreateEnvironmentsResp.Size(m)
}
func (m *CreateEnvironmentsResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateEnvironmentsResp.DiscardUnknown(m)
}

var xxx_messageInfo_CreateEnvironmentsResp proto.InternalMessageInfo

func (m *CreateEnvironmentsResp) GetBaseResp() *common.BaseResp {
	if m != nil {
		return m.BaseResp
	}
	return nil
}

func init() {
	proto.RegisterType((*GetEnvironmentsRequ)(nil), "proto.GetEnvironmentsRequ")
	proto.RegisterType((*GetEnvironmentsResp)(nil), "proto.GetEnvironmentsResp")
	proto.RegisterType((*Environment)(nil), "proto.Environment")
	proto.RegisterType((*CreateEnvironmentsRequ)(nil), "proto.CreateEnvironmentsRequ")
	proto.RegisterType((*CreateEnvironmentsResp)(nil), "proto.CreateEnvironmentsResp")
}

func init() {
	proto.RegisterFile("environment.proto", fileDescriptor_64e647b85623514a)
}

var fileDescriptor_64e647b85623514a = []byte{
	// 307 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x4d, 0x83, 0x62, 0xa7, 0x4a, 0xed, 0x08, 0x12, 0x02, 0x42, 0x58, 0x3c, 0xe4, 0x60,
	0x37, 0x10, 0xc1, 0xa3, 0x97, 0xfa, 0x07, 0x11, 0x3c, 0xac, 0x5e, 0x3d, 0x24, 0x71, 0x8c, 0x15,
	0xb3, 0x1b, 0xb3, 0x69, 0x2f, 0x7e, 0x2f, 0x3f, 0x9f, 0x24, 0xa9, 0xed, 0x36, 0x46, 0x41, 0xbc,
	0x24, 0xcb, 0x7b, 0x6f, 0x7f, 0xbb, 0xf3, 0x12, 0x18, 0x91, 0x9c, 0x4f, 0x0b, 0x25, 0x33, 0x92,
	0x25, 0xcf, 0x0b, 0x55, 0x2a, 0xdc, 0xac, 0x5f, 0xee, 0x59, 0x3a, 0x2d, 0x9f, 0x67, 0x31, 0x4f,
	0x54, 0x16, 0xa4, 0xea, 0x35, 0x92, 0xa9, 0xe2, 0x89, 0xac, 0x96, 0x8f, 0x24, 0xc7, 0x89, 0x92,
	0x4f, 0x41, 0xf5, 0x18, 0xd7, 0xe9, 0x20, 0x51, 0x59, 0xa6, 0x64, 0x10, 0x47, 0x9a, 0x1a, 0x0c,
	0x9b, 0xc0, 0xfe, 0x15, 0x95, 0x17, 0x2b, 0xbc, 0x16, 0xf4, 0x36, 0xc3, 0x63, 0xd8, 0xae, 0x42,
	0xd5, 0xda, 0xb1, 0x3c, 0xcb, 0x1f, 0x84, 0x7b, 0xbc, 0xd9, 0xcc, 0xbf, 0x74, 0xb1, 0x4c, 0xb0,
	0xf7, 0x0e, 0x88, 0xce, 0x57, 0x10, 0x9d, 0x77, 0x43, 0x74, 0x2e, 0x96, 0x09, 0x3c, 0x85, 0x1d,
	0x63, 0x4a, 0xed, 0xd8, 0x9e, 0xed, 0x0f, 0x42, 0x6c, 0xee, 0xc9, 0x0d, 0xb8, 0x58, 0xcb, 0xb1,
	0x07, 0x18, 0x18, 0x26, 0x1e, 0xc1, 0xae, 0x61, 0x5f, 0x9f, 0xd7, 0x27, 0xdb, 0x62, 0x5d, 0x44,
	0x1f, 0x86, 0x86, 0x70, 0x1b, 0x65, 0xe4, 0xf4, 0x3c, 0xcb, 0xef, 0x8b, 0xb6, 0xcc, 0x5e, 0xe0,
	0x60, 0x52, 0x50, 0x54, 0xd2, 0xff, 0x3a, 0x42, 0xd6, 0x1a, 0xaf, 0xe7, 0xd9, 0x7e, 0xbf, 0x35,
	0xca, 0x65, 0xf7, 0x59, 0x7f, 0xad, 0x32, 0xfc, 0xb0, 0x60, 0x64, 0x20, 0xee, 0xa8, 0x98, 0x53,
	0x81, 0x37, 0x30, 0x6c, 0x7d, 0x25, 0x74, 0x17, 0xed, 0x76, 0xfc, 0x02, 0xee, 0x8f, 0x9e, 0xce,
	0xd9, 0x06, 0xde, 0x03, 0x7e, 0xbf, 0x2a, 0x1e, 0x2e, 0xf6, 0x74, 0x37, 0xe6, 0xfe, 0x66, 0x57,
	0xd4, 0x78, 0xab, 0xf6, 0x4f, 0x3e, 0x03, 0x00, 0x00, 0xff, 0xff, 0x88, 0xca, 0xc7, 0x0c, 0xf0,
	0x02, 0x00, 0x00,
}
