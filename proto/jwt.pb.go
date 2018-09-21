// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/jwt.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type EJwtType int32

const (
	EJwtType_PROTO EJwtType = 0
	EJwtType_JSON  EJwtType = 1
)

var EJwtType_name = map[int32]string{
	0: "PROTO",
	1: "JSON",
}
var EJwtType_value = map[string]int32{
	"PROTO": 0,
	"JSON":  1,
}

func (x EJwtType) String() string {
	return proto.EnumName(EJwtType_name, int32(x))
}
func (EJwtType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_jwt_e81f20fa50718fd7, []int{0}
}

// jwt header信息
type PJwtHeader struct {
	Alg                  string   `protobuf:"bytes,1,opt,name=Alg,proto3" json:"Alg,omitempty"`
	Typ                  string   `protobuf:"bytes,2,opt,name=Typ,proto3" json:"Typ,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PJwtHeader) Reset()         { *m = PJwtHeader{} }
func (m *PJwtHeader) String() string { return proto.CompactTextString(m) }
func (*PJwtHeader) ProtoMessage()    {}
func (*PJwtHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_jwt_e81f20fa50718fd7, []int{0}
}
func (m *PJwtHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PJwtHeader.Unmarshal(m, b)
}
func (m *PJwtHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PJwtHeader.Marshal(b, m, deterministic)
}
func (dst *PJwtHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PJwtHeader.Merge(dst, src)
}
func (m *PJwtHeader) XXX_Size() int {
	return xxx_messageInfo_PJwtHeader.Size(m)
}
func (m *PJwtHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_PJwtHeader.DiscardUnknown(m)
}

var xxx_messageInfo_PJwtHeader proto.InternalMessageInfo

func (m *PJwtHeader) GetAlg() string {
	if m != nil {
		return m.Alg
	}
	return ""
}

func (m *PJwtHeader) GetTyp() string {
	if m != nil {
		return m.Typ
	}
	return ""
}

// jwt payload信息
type PJwtPayload struct {
	LoginType            ELoginType `protobuf:"varint,1,opt,name=login_type,json=loginType,proto3,enum=proto.ELoginType" json:"login_type,omitempty"`
	Uid                  int32      `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Time                 int64      `protobuf:"varint,3,opt,name=time,proto3" json:"time,omitempty"`
	Os                   Os         `protobuf:"varint,4,opt,name=os,proto3,enum=proto.Os" json:"os,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *PJwtPayload) Reset()         { *m = PJwtPayload{} }
func (m *PJwtPayload) String() string { return proto.CompactTextString(m) }
func (*PJwtPayload) ProtoMessage()    {}
func (*PJwtPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_jwt_e81f20fa50718fd7, []int{1}
}
func (m *PJwtPayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PJwtPayload.Unmarshal(m, b)
}
func (m *PJwtPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PJwtPayload.Marshal(b, m, deterministic)
}
func (dst *PJwtPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PJwtPayload.Merge(dst, src)
}
func (m *PJwtPayload) XXX_Size() int {
	return xxx_messageInfo_PJwtPayload.Size(m)
}
func (m *PJwtPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_PJwtPayload.DiscardUnknown(m)
}

var xxx_messageInfo_PJwtPayload proto.InternalMessageInfo

func (m *PJwtPayload) GetLoginType() ELoginType {
	if m != nil {
		return m.LoginType
	}
	return ELoginType_ACCOUNT
}

func (m *PJwtPayload) GetUid() int32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *PJwtPayload) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *PJwtPayload) GetOs() Os {
	if m != nil {
		return m.Os
	}
	return Os_ANDROID
}

// jwt token信息
type PJwtToken struct {
	Header               *PJwtHeader  `protobuf:"bytes,1,opt,name=Header,proto3" json:"Header,omitempty"`
	Payload              *PJwtPayload `protobuf:"bytes,2,opt,name=Payload,proto3" json:"Payload,omitempty"`
	Signature            string       `protobuf:"bytes,3,opt,name=Signature,proto3" json:"Signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *PJwtToken) Reset()         { *m = PJwtToken{} }
func (m *PJwtToken) String() string { return proto.CompactTextString(m) }
func (*PJwtToken) ProtoMessage()    {}
func (*PJwtToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_jwt_e81f20fa50718fd7, []int{2}
}
func (m *PJwtToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PJwtToken.Unmarshal(m, b)
}
func (m *PJwtToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PJwtToken.Marshal(b, m, deterministic)
}
func (dst *PJwtToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PJwtToken.Merge(dst, src)
}
func (m *PJwtToken) XXX_Size() int {
	return xxx_messageInfo_PJwtToken.Size(m)
}
func (m *PJwtToken) XXX_DiscardUnknown() {
	xxx_messageInfo_PJwtToken.DiscardUnknown(m)
}

var xxx_messageInfo_PJwtToken proto.InternalMessageInfo

func (m *PJwtToken) GetHeader() *PJwtHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *PJwtToken) GetPayload() *PJwtPayload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *PJwtToken) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func init() {
	proto.RegisterType((*PJwtHeader)(nil), "proto.PJwtHeader")
	proto.RegisterType((*PJwtPayload)(nil), "proto.PJwtPayload")
	proto.RegisterType((*PJwtToken)(nil), "proto.PJwtToken")
	proto.RegisterEnum("proto.EJwtType", EJwtType_name, EJwtType_value)
}

func init() { proto.RegisterFile("proto/jwt.proto", fileDescriptor_jwt_e81f20fa50718fd7) }

var fileDescriptor_jwt_e81f20fa50718fd7 = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0x4d, 0x4b, 0xf3, 0x40,
	0x14, 0x85, 0xdf, 0xc9, 0x47, 0xdf, 0xce, 0x0d, 0x68, 0xbc, 0x6e, 0xa2, 0x08, 0x96, 0xac, 0xaa,
	0x48, 0x2d, 0xf5, 0x17, 0xb8, 0x28, 0x48, 0x11, 0x13, 0xa6, 0xd9, 0x4b, 0x6c, 0x86, 0x10, 0x8d,
	0x99, 0x90, 0x4c, 0x28, 0xd9, 0x89, 0xbf, 0x5c, 0xe6, 0xa3, 0xd6, 0xd5, 0x1c, 0xce, 0xbd, 0xf3,
	0x9c, 0x73, 0xe1, 0xb4, 0xed, 0x84, 0x14, 0xf7, 0xef, 0x7b, 0xb9, 0xd0, 0x0a, 0x7d, 0xfd, 0x5c,
	0x9e, 0x1b, 0x3f, 0xdf, 0xed, 0xc4, 0xd0, 0xd8, 0x59, 0xbc, 0x04, 0x48, 0x37, 0x7b, 0xf9, 0xc4,
	0xf3, 0x82, 0x77, 0x18, 0x82, 0xfb, 0x58, 0x97, 0x11, 0x99, 0x91, 0x39, 0x65, 0x4a, 0x2a, 0x27,
	0x1b, 0xdb, 0xc8, 0x31, 0x4e, 0x36, 0xb6, 0xf1, 0x17, 0x81, 0x40, 0x7d, 0x49, 0xf3, 0xb1, 0x16,
	0x79, 0x81, 0x4b, 0x80, 0x5a, 0x94, 0x55, 0xf3, 0x2a, 0xc7, 0x96, 0xeb, 0xaf, 0x27, 0xab, 0x33,
	0x43, 0x5f, 0xac, 0x9f, 0xd5, 0x24, 0x1b, 0x5b, 0xce, 0x68, 0x7d, 0x90, 0x8a, 0x39, 0x54, 0x85,
	0x66, 0xfa, 0x4c, 0x49, 0x44, 0xf0, 0x64, 0xf5, 0xc9, 0x23, 0x77, 0x46, 0xe6, 0x2e, 0xd3, 0x1a,
	0x2f, 0xc0, 0x11, 0x7d, 0xe4, 0x69, 0x1e, 0xb5, 0xbc, 0xa4, 0x67, 0x8e, 0xe8, 0xe3, 0x6f, 0x02,
	0x54, 0x55, 0xc8, 0xc4, 0x07, 0x6f, 0xf0, 0x06, 0x26, 0xa6, 0xbe, 0x0e, 0x0f, 0x7e, 0xc3, 0x8f,
	0x77, 0x31, 0xbb, 0x80, 0x77, 0xf0, 0xdf, 0xd6, 0xd6, 0xe9, 0xc1, 0x0a, 0xff, 0xec, 0xda, 0x09,
	0x3b, 0xac, 0xe0, 0x15, 0xd0, 0x6d, 0x55, 0x36, 0xb9, 0x1c, 0x3a, 0x53, 0x8d, 0xb2, 0xa3, 0x71,
	0x7b, 0x0d, 0xd3, 0xb5, 0xea, 0xa0, 0x2e, 0xa2, 0xe0, 0xa7, 0x2c, 0xc9, 0x92, 0xf0, 0x1f, 0x4e,
	0xc1, 0xdb, 0x6c, 0x93, 0x97, 0x90, 0xbc, 0x4d, 0x34, 0xfa, 0xe1, 0x27, 0x00, 0x00, 0xff, 0xff,
	0x62, 0xe6, 0x93, 0x34, 0x90, 0x01, 0x00, 0x00,
}