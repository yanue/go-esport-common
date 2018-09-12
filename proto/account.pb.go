// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/account.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type User struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Gender               int32    `protobuf:"varint,3,opt,name=gender,proto3" json:"gender,omitempty"`
	SchoolId             int32    `protobuf:"varint,4,opt,name=school_id,json=schoolId,proto3" json:"school_id,omitempty"`
	ClassId              int32    `protobuf:"varint,5,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	AreaId               int32    `protobuf:"varint,6,opt,name=area_id,json=areaId,proto3" json:"area_id,omitempty"`
	IdentityStatus       int32    `protobuf:"varint,7,opt,name=identity_status,json=identityStatus,proto3" json:"identity_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_5d2fb36eecf4fb7f, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetGender() int32 {
	if m != nil {
		return m.Gender
	}
	return 0
}

func (m *User) GetSchoolId() int32 {
	if m != nil {
		return m.SchoolId
	}
	return 0
}

func (m *User) GetClassId() int32 {
	if m != nil {
		return m.ClassId
	}
	return 0
}

func (m *User) GetAreaId() int32 {
	if m != nil {
		return m.AreaId
	}
	return 0
}

func (m *User) GetIdentityStatus() int32 {
	if m != nil {
		return m.IdentityStatus
	}
	return 0
}

type RegData struct {
	Phone                string   `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	Code                 string   `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegData) Reset()         { *m = RegData{} }
func (m *RegData) String() string { return proto.CompactTextString(m) }
func (*RegData) ProtoMessage()    {}
func (*RegData) Descriptor() ([]byte, []int) {
	return fileDescriptor_account_5d2fb36eecf4fb7f, []int{1}
}
func (m *RegData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegData.Unmarshal(m, b)
}
func (m *RegData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegData.Marshal(b, m, deterministic)
}
func (dst *RegData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegData.Merge(dst, src)
}
func (m *RegData) XXX_Size() int {
	return xxx_messageInfo_RegData.Size(m)
}
func (m *RegData) XXX_DiscardUnknown() {
	xxx_messageInfo_RegData.DiscardUnknown(m)
}

var xxx_messageInfo_RegData proto.InternalMessageInfo

func (m *RegData) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *RegData) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "proto.User")
	proto.RegisterType((*RegData)(nil), "proto.regData")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Account service

type AccountClient interface {
	Login(ctx context.Context, in *PSingleString, opts ...client.CallOption) (*User, error)
	Reg(ctx context.Context, in *PSingleString, opts ...client.CallOption) (*User, error)
	GetVerifyCode(ctx context.Context, in *PSingleString, opts ...client.CallOption) (*PBool, error)
	GetUserInfo(ctx context.Context, in *PInt32, opts ...client.CallOption) (*User, error)
}

type accountClient struct {
	c           client.Client
	serviceName string
}

func NewAccountClient(serviceName string, c client.Client) AccountClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "proto"
	}
	return &accountClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *accountClient) Login(ctx context.Context, in *PSingleString, opts ...client.CallOption) (*User, error) {
	req := c.c.NewRequest(c.serviceName, "Account.Login", in)
	out := new(User)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) Reg(ctx context.Context, in *PSingleString, opts ...client.CallOption) (*User, error) {
	req := c.c.NewRequest(c.serviceName, "Account.Reg", in)
	out := new(User)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) GetVerifyCode(ctx context.Context, in *PSingleString, opts ...client.CallOption) (*PBool, error) {
	req := c.c.NewRequest(c.serviceName, "Account.GetVerifyCode", in)
	out := new(PBool)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) GetUserInfo(ctx context.Context, in *PInt32, opts ...client.CallOption) (*User, error) {
	req := c.c.NewRequest(c.serviceName, "Account.GetUserInfo", in)
	out := new(User)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Account service

type AccountHandler interface {
	Login(context.Context, *PSingleString, *User) error
	Reg(context.Context, *PSingleString, *User) error
	GetVerifyCode(context.Context, *PSingleString, *PBool) error
	GetUserInfo(context.Context, *PInt32, *User) error
}

func RegisterAccountHandler(s server.Server, hdlr AccountHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Account{hdlr}, opts...))
}

type Account struct {
	AccountHandler
}

func (h *Account) Login(ctx context.Context, in *PSingleString, out *User) error {
	return h.AccountHandler.Login(ctx, in, out)
}

func (h *Account) Reg(ctx context.Context, in *PSingleString, out *User) error {
	return h.AccountHandler.Reg(ctx, in, out)
}

func (h *Account) GetVerifyCode(ctx context.Context, in *PSingleString, out *PBool) error {
	return h.AccountHandler.GetVerifyCode(ctx, in, out)
}

func (h *Account) GetUserInfo(ctx context.Context, in *PInt32, out *User) error {
	return h.AccountHandler.GetUserInfo(ctx, in, out)
}

func init() { proto.RegisterFile("proto/account.proto", fileDescriptor_account_5d2fb36eecf4fb7f) }

var fileDescriptor_account_5d2fb36eecf4fb7f = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x4f, 0x6b, 0xf2, 0x40,
	0x10, 0x87, 0x8d, 0x1a, 0xa3, 0xe3, 0xab, 0xef, 0xcb, 0xbc, 0xd2, 0xa6, 0xf6, 0x22, 0xb9, 0x54,
	0xda, 0x62, 0x41, 0xe9, 0x07, 0xe8, 0x1f, 0x90, 0x40, 0x0f, 0x25, 0xd2, 0x5e, 0x65, 0xcd, 0x8e,
	0x71, 0x21, 0xdd, 0x95, 0xec, 0x7a, 0xf0, 0xbb, 0xf5, 0xde, 0xaf, 0x55, 0x76, 0x57, 0x7b, 0x28,
	0x14, 0x7a, 0xca, 0x3c, 0xbf, 0x67, 0x32, 0xcb, 0x0c, 0xfc, 0xdf, 0x56, 0xca, 0xa8, 0x1b, 0x96,
	0xe7, 0x6a, 0x27, 0xcd, 0xc4, 0x11, 0x86, 0xee, 0x33, 0xfc, 0xe7, 0xdd, 0x8a, 0x69, 0xf2, 0x22,
	0x79, 0x0f, 0xa0, 0xf9, 0xa2, 0xa9, 0xc2, 0x3e, 0xd4, 0x05, 0x8f, 0x83, 0x51, 0x30, 0x0e, 0xb3,
	0xba, 0xe0, 0x88, 0xd0, 0x94, 0xec, 0x8d, 0xe2, 0xfa, 0x28, 0x18, 0x77, 0x32, 0x57, 0xe3, 0x09,
	0xb4, 0x0a, 0x92, 0x9c, 0xaa, 0xb8, 0xe1, 0xfa, 0x0e, 0x84, 0xe7, 0xd0, 0xd1, 0xf9, 0x46, 0xa9,
	0x72, 0x29, 0x78, 0xdc, 0x74, 0xaa, 0xed, 0x83, 0x94, 0xe3, 0x19, 0xb4, 0xf3, 0x92, 0x69, 0x6d,
	0x5d, 0xe8, 0x5c, 0xe4, 0x38, 0xe5, 0x78, 0x0a, 0x11, 0xab, 0x88, 0x59, 0xd3, 0xf2, 0x03, 0x2d,
	0xa6, 0x1c, 0x2f, 0xe0, 0xaf, 0xe0, 0x24, 0x8d, 0x30, 0xfb, 0xa5, 0x36, 0xcc, 0xec, 0x74, 0x1c,
	0xb9, 0x86, 0xfe, 0x31, 0x5e, 0xb8, 0x34, 0x99, 0x41, 0x54, 0x51, 0xf1, 0xc8, 0x0c, 0xc3, 0x01,
	0x84, 0xdb, 0x8d, 0x92, 0xe4, 0x76, 0xe8, 0x64, 0x1e, 0xec, 0x1a, 0xb9, 0xe2, 0x5f, 0x6b, 0xd8,
	0x7a, 0xfa, 0x11, 0x40, 0x74, 0xe7, 0xcf, 0x83, 0xd7, 0x10, 0x3e, 0xa9, 0x42, 0x48, 0x1c, 0xf8,
	0x83, 0x4c, 0x9e, 0x17, 0x42, 0x16, 0x25, 0x2d, 0x4c, 0x25, 0x64, 0x31, 0xec, 0x1e, 0x52, 0x7b,
	0xa2, 0xa4, 0x86, 0x97, 0xd0, 0xc8, 0xa8, 0xf8, 0x5d, 0xef, 0x2d, 0xf4, 0xe6, 0x64, 0x5e, 0xa9,
	0x12, 0xeb, 0xfd, 0x83, 0xe2, 0xf4, 0xc3, 0x5f, 0x7f, 0x8e, 0xe9, 0xbd, 0x52, 0x65, 0x52, 0xc3,
	0x2b, 0xe8, 0xce, 0xc9, 0xd8, 0x19, 0xa9, 0x5c, 0x2b, 0xec, 0x1d, 0x75, 0x2a, 0xcd, 0x6c, 0xfa,
	0xed, 0x8d, 0x55, 0xcb, 0xd1, 0xec, 0x33, 0x00, 0x00, 0xff, 0xff, 0x4a, 0x2d, 0x6a, 0xa5, 0xf4,
	0x01, 0x00, 0x00,
}
