// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package user

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type UpdateUserRequest struct {
	NickName             string   `protobuf:"bytes,1,opt,name=nickName,proto3" json:"nickName,omitempty"`
	AvatarUrl            string   `protobuf:"bytes,2,opt,name=avatarUrl,proto3" json:"avatarUrl,omitempty"`
	Gender               int32    `protobuf:"varint,3,opt,name=gender,proto3" json:"gender,omitempty"`
	Country              string   `protobuf:"bytes,4,opt,name=country,proto3" json:"country,omitempty"`
	Province             string   `protobuf:"bytes,5,opt,name=province,proto3" json:"province,omitempty"`
	City                 string   `protobuf:"bytes,6,opt,name=city,proto3" json:"city,omitempty"`
	Language             string   `protobuf:"bytes,7,opt,name=language,proto3" json:"language,omitempty"`
	Uid                  string   `protobuf:"bytes,8,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserRequest) Reset()         { *m = UpdateUserRequest{} }
func (m *UpdateUserRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateUserRequest) ProtoMessage()    {}
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *UpdateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserRequest.Unmarshal(m, b)
}
func (m *UpdateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserRequest.Marshal(b, m, deterministic)
}
func (m *UpdateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserRequest.Merge(m, src)
}
func (m *UpdateUserRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateUserRequest.Size(m)
}
func (m *UpdateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserRequest proto.InternalMessageInfo

func (m *UpdateUserRequest) GetNickName() string {
	if m != nil {
		return m.NickName
	}
	return ""
}

func (m *UpdateUserRequest) GetAvatarUrl() string {
	if m != nil {
		return m.AvatarUrl
	}
	return ""
}

func (m *UpdateUserRequest) GetGender() int32 {
	if m != nil {
		return m.Gender
	}
	return 0
}

func (m *UpdateUserRequest) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *UpdateUserRequest) GetProvince() string {
	if m != nil {
		return m.Province
	}
	return ""
}

func (m *UpdateUserRequest) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *UpdateUserRequest) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

func (m *UpdateUserRequest) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

type UpdateUserResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserResponse) Reset()         { *m = UpdateUserResponse{} }
func (m *UpdateUserResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateUserResponse) ProtoMessage()    {}
func (*UpdateUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *UpdateUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserResponse.Unmarshal(m, b)
}
func (m *UpdateUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserResponse.Marshal(b, m, deterministic)
}
func (m *UpdateUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserResponse.Merge(m, src)
}
func (m *UpdateUserResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateUserResponse.Size(m)
}
func (m *UpdateUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserResponse proto.InternalMessageInfo

type VerificationLoginRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Uid                  string   `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerificationLoginRequest) Reset()         { *m = VerificationLoginRequest{} }
func (m *VerificationLoginRequest) String() string { return proto.CompactTextString(m) }
func (*VerificationLoginRequest) ProtoMessage()    {}
func (*VerificationLoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *VerificationLoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerificationLoginRequest.Unmarshal(m, b)
}
func (m *VerificationLoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerificationLoginRequest.Marshal(b, m, deterministic)
}
func (m *VerificationLoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerificationLoginRequest.Merge(m, src)
}
func (m *VerificationLoginRequest) XXX_Size() int {
	return xxx_messageInfo_VerificationLoginRequest.Size(m)
}
func (m *VerificationLoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VerificationLoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VerificationLoginRequest proto.InternalMessageInfo

func (m *VerificationLoginRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *VerificationLoginRequest) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

type VerificationLoginResponse struct {
	Result               bool     `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerificationLoginResponse) Reset()         { *m = VerificationLoginResponse{} }
func (m *VerificationLoginResponse) String() string { return proto.CompactTextString(m) }
func (*VerificationLoginResponse) ProtoMessage()    {}
func (*VerificationLoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *VerificationLoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerificationLoginResponse.Unmarshal(m, b)
}
func (m *VerificationLoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerificationLoginResponse.Marshal(b, m, deterministic)
}
func (m *VerificationLoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerificationLoginResponse.Merge(m, src)
}
func (m *VerificationLoginResponse) XXX_Size() int {
	return xxx_messageInfo_VerificationLoginResponse.Size(m)
}
func (m *VerificationLoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VerificationLoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VerificationLoginResponse proto.InternalMessageInfo

func (m *VerificationLoginResponse) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

type UserLoginRequest struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserLoginRequest) Reset()         { *m = UserLoginRequest{} }
func (m *UserLoginRequest) String() string { return proto.CompactTextString(m) }
func (*UserLoginRequest) ProtoMessage()    {}
func (*UserLoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{4}
}

func (m *UserLoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserLoginRequest.Unmarshal(m, b)
}
func (m *UserLoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserLoginRequest.Marshal(b, m, deterministic)
}
func (m *UserLoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserLoginRequest.Merge(m, src)
}
func (m *UserLoginRequest) XXX_Size() int {
	return xxx_messageInfo_UserLoginRequest.Size(m)
}
func (m *UserLoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserLoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserLoginRequest proto.InternalMessageInfo

func (m *UserLoginRequest) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

type UserLoginResponse struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserLoginResponse) Reset()         { *m = UserLoginResponse{} }
func (m *UserLoginResponse) String() string { return proto.CompactTextString(m) }
func (*UserLoginResponse) ProtoMessage()    {}
func (*UserLoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{5}
}

func (m *UserLoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserLoginResponse.Unmarshal(m, b)
}
func (m *UserLoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserLoginResponse.Marshal(b, m, deterministic)
}
func (m *UserLoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserLoginResponse.Merge(m, src)
}
func (m *UserLoginResponse) XXX_Size() int {
	return xxx_messageInfo_UserLoginResponse.Size(m)
}
func (m *UserLoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UserLoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UserLoginResponse proto.InternalMessageInfo

func (m *UserLoginResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// 请求参数-根据自己的需求定义
type GetUserRequest struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserRequest) Reset()         { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()    {}
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{6}
}

func (m *GetUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserRequest.Unmarshal(m, b)
}
func (m *GetUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserRequest.Marshal(b, m, deterministic)
}
func (m *GetUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserRequest.Merge(m, src)
}
func (m *GetUserRequest) XXX_Size() int {
	return xxx_messageInfo_GetUserRequest.Size(m)
}
func (m *GetUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserRequest proto.InternalMessageInfo

func (m *GetUserRequest) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

// 返回参数-根据自己的需求定义
type GetUserResponse struct {
	NickName             string   `protobuf:"bytes,1,opt,name=nickName,proto3" json:"nickName,omitempty"`
	AvatarUrl            string   `protobuf:"bytes,2,opt,name=avatarUrl,proto3" json:"avatarUrl,omitempty"`
	Gender               int32    `protobuf:"varint,3,opt,name=gender,proto3" json:"gender,omitempty"`
	Country              string   `protobuf:"bytes,4,opt,name=country,proto3" json:"country,omitempty"`
	Province             string   `protobuf:"bytes,5,opt,name=province,proto3" json:"province,omitempty"`
	City                 string   `protobuf:"bytes,6,opt,name=city,proto3" json:"city,omitempty"`
	Language             string   `protobuf:"bytes,7,opt,name=language,proto3" json:"language,omitempty"`
	Init                 int32    `protobuf:"varint,8,opt,name=init,proto3" json:"init,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUserResponse) Reset()         { *m = GetUserResponse{} }
func (m *GetUserResponse) String() string { return proto.CompactTextString(m) }
func (*GetUserResponse) ProtoMessage()    {}
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{7}
}

func (m *GetUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUserResponse.Unmarshal(m, b)
}
func (m *GetUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUserResponse.Marshal(b, m, deterministic)
}
func (m *GetUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUserResponse.Merge(m, src)
}
func (m *GetUserResponse) XXX_Size() int {
	return xxx_messageInfo_GetUserResponse.Size(m)
}
func (m *GetUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetUserResponse proto.InternalMessageInfo

func (m *GetUserResponse) GetNickName() string {
	if m != nil {
		return m.NickName
	}
	return ""
}

func (m *GetUserResponse) GetAvatarUrl() string {
	if m != nil {
		return m.AvatarUrl
	}
	return ""
}

func (m *GetUserResponse) GetGender() int32 {
	if m != nil {
		return m.Gender
	}
	return 0
}

func (m *GetUserResponse) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *GetUserResponse) GetProvince() string {
	if m != nil {
		return m.Province
	}
	return ""
}

func (m *GetUserResponse) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *GetUserResponse) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

func (m *GetUserResponse) GetInit() int32 {
	if m != nil {
		return m.Init
	}
	return 0
}

type Config struct {
	Port                 string   `protobuf:"bytes,1,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{8}
}

func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetPort() string {
	if m != nil {
		return m.Port
	}
	return ""
}

func init() {
	proto.RegisterType((*UpdateUserRequest)(nil), "protobuf.UpdateUserRequest")
	proto.RegisterType((*UpdateUserResponse)(nil), "protobuf.UpdateUserResponse")
	proto.RegisterType((*VerificationLoginRequest)(nil), "protobuf.VerificationLoginRequest")
	proto.RegisterType((*VerificationLoginResponse)(nil), "protobuf.VerificationLoginResponse")
	proto.RegisterType((*UserLoginRequest)(nil), "protobuf.UserLoginRequest")
	proto.RegisterType((*UserLoginResponse)(nil), "protobuf.UserLoginResponse")
	proto.RegisterType((*GetUserRequest)(nil), "protobuf.GetUserRequest")
	proto.RegisterType((*GetUserResponse)(nil), "protobuf.GetUserResponse")
	proto.RegisterType((*Config)(nil), "protobuf.Config")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 423 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x53, 0xcd, 0x8e, 0xd3, 0x30,
	0x10, 0x26, 0x69, 0x92, 0xa6, 0x73, 0x80, 0xed, 0x68, 0x85, 0xbc, 0xa1, 0x87, 0x95, 0xe1, 0xb0,
	0x5c, 0x7a, 0x60, 0x5f, 0x00, 0x2d, 0x12, 0x3f, 0xd2, 0x8a, 0x43, 0xa4, 0xc2, 0x0d, 0xc9, 0x9b,
	0xba, 0x91, 0xb5, 0xc5, 0x0e, 0x8e, 0x53, 0xa9, 0xaf, 0xc3, 0x6b, 0xf0, 0x2c, 0x7d, 0x17, 0x14,
	0xc7, 0xf9, 0x69, 0x1b, 0xb8, 0xef, 0x29, 0xf3, 0xcd, 0x7c, 0xf9, 0xe6, 0xd7, 0x00, 0x55, 0xc9,
	0xf5, 0xb2, 0xd0, 0xca, 0x28, 0x8c, 0xed, 0xe7, 0xa1, 0xda, 0xd0, 0x83, 0x07, 0xf3, 0x55, 0xb1,
	0x66, 0x86, 0xaf, 0x4a, 0xae, 0x53, 0xfe, 0xab, 0xe2, 0xa5, 0xc1, 0x04, 0x62, 0x29, 0xb2, 0xc7,
	0xaf, 0xec, 0x27, 0x27, 0xde, 0xb5, 0x77, 0x33, 0x4b, 0x3b, 0x8c, 0x0b, 0x98, 0xb1, 0x1d, 0x33,
	0x4c, 0xaf, 0xf4, 0x96, 0xf8, 0x36, 0xd8, 0x3b, 0xf0, 0x25, 0x44, 0x39, 0x97, 0x6b, 0xae, 0xc9,
	0xe4, 0xda, 0xbb, 0x09, 0x53, 0x87, 0x90, 0xc0, 0x34, 0x53, 0x95, 0x34, 0x7a, 0x4f, 0x02, 0xfb,
	0x4f, 0x0b, 0xeb, 0x5c, 0x85, 0x56, 0x3b, 0x21, 0x33, 0x4e, 0xc2, 0x26, 0x57, 0x8b, 0x11, 0x21,
	0xc8, 0x84, 0xd9, 0x93, 0xc8, 0xfa, 0xad, 0x5d, 0xf3, 0xb7, 0x4c, 0xe6, 0x15, 0xcb, 0x39, 0x99,
	0x36, 0xfc, 0x16, 0xe3, 0x05, 0x4c, 0x2a, 0xb1, 0x26, 0xb1, 0x75, 0xd7, 0x26, 0xbd, 0x04, 0x1c,
	0xb6, 0x57, 0x16, 0x4a, 0x96, 0x9c, 0xde, 0x01, 0xf9, 0xc6, 0xb5, 0xd8, 0x88, 0x8c, 0x19, 0xa1,
	0xe4, 0xbd, 0xca, 0x85, 0x6c, 0x7b, 0xbf, 0x84, 0xd0, 0xa8, 0x47, 0x2e, 0x5d, 0xe3, 0x0d, 0x68,
	0x95, 0xfd, 0x5e, 0xf9, 0x16, 0xae, 0x46, 0x34, 0x9a, 0x04, 0xf5, 0x18, 0x34, 0x2f, 0xab, 0xad,
	0xb1, 0x2a, 0x71, 0xea, 0x10, 0x7d, 0x03, 0x17, 0x75, 0x21, 0x47, 0x09, 0x9d, 0xb4, 0xd7, 0x4b,
	0xbf, 0x85, 0xf9, 0x80, 0xe5, 0x24, 0x47, 0xeb, 0xa2, 0x14, 0x9e, 0x7f, 0xe2, 0x66, 0xb8, 0xbb,
	0x73, 0xb9, 0x83, 0x07, 0x2f, 0x3a, 0x92, 0x53, 0x7b, 0xaa, 0x1b, 0x46, 0x08, 0x84, 0x14, 0xc6,
	0xae, 0x38, 0x4c, 0xad, 0x4d, 0x17, 0x10, 0x7d, 0x50, 0x72, 0x23, 0xf2, 0x3a, 0x5a, 0x28, 0x6d,
	0x5c, 0x47, 0xd6, 0x7e, 0xf7, 0xc7, 0x87, 0xa0, 0x6e, 0x1d, 0xdf, 0xc3, 0xd4, 0x4d, 0x01, 0xc9,
	0xb2, 0x7d, 0x00, 0xcb, 0xe3, 0xe9, 0x25, 0x57, 0x23, 0x11, 0x77, 0x34, 0xcf, 0xf0, 0x23, 0xcc,
	0xba, 0xbd, 0x60, 0xd2, 0x33, 0x4f, 0x57, 0x9a, 0xbc, 0x1a, 0x8d, 0x75, 0x3a, 0x3f, 0x60, 0x7e,
	0x76, 0x3a, 0x48, 0xfb, 0x7f, 0xfe, 0x75, 0x9b, 0xc9, 0xeb, 0xff, 0x72, 0x3a, 0xfd, 0x2f, 0x00,
	0xfd, 0xd1, 0xe3, 0xb0, 0x98, 0xd3, 0x97, 0x9e, 0x2c, 0xc6, 0x83, 0xad, 0xd4, 0x5d, 0xf4, 0xdb,
	0x9f, 0x7c, 0xbe, 0xff, 0xfe, 0x10, 0x59, 0xda, 0xed, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x53,
	0xda, 0x9c, 0x87, 0x46, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	// Sends a greeting
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	UserLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginResponse, error)
	VerificationLogin(ctx context.Context, in *VerificationLoginRequest, opts ...grpc.CallOption) (*VerificationLoginResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error)
}

type userClient struct {
	cc *grpc.ClientConn
}

func NewUserClient(cc *grpc.ClientConn) UserClient {
	return &userClient{cc}
}

func (c *userClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, "/protobuf.User/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginResponse, error) {
	out := new(UserLoginResponse)
	err := c.cc.Invoke(ctx, "/protobuf.User/UserLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) VerificationLogin(ctx context.Context, in *VerificationLoginRequest, opts ...grpc.CallOption) (*VerificationLoginResponse, error) {
	out := new(VerificationLoginResponse)
	err := c.cc.Invoke(ctx, "/protobuf.User/VerificationLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error) {
	out := new(UpdateUserResponse)
	err := c.cc.Invoke(ctx, "/protobuf.User/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	// Sends a greeting
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	UserLogin(context.Context, *UserLoginRequest) (*UserLoginResponse, error)
	VerificationLogin(context.Context, *VerificationLoginRequest) (*VerificationLoginResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
}

// UnimplementedUserServer can be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (*UnimplementedUserServer) GetUser(ctx context.Context, req *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (*UnimplementedUserServer) UserLogin(ctx context.Context, req *UserLoginRequest) (*UserLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogin not implemented")
}
func (*UnimplementedUserServer) VerificationLogin(ctx context.Context, req *VerificationLoginRequest) (*VerificationLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerificationLogin not implemented")
}
func (*UnimplementedUserServer) UpdateUser(ctx context.Context, req *UpdateUserRequest) (*UpdateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.User/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.User/UserLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserLogin(ctx, req.(*UserLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_VerificationLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerificationLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).VerificationLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.User/VerificationLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).VerificationLogin(ctx, req.(*VerificationLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.User/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _User_GetUser_Handler,
		},
		{
			MethodName: "UserLogin",
			Handler:    _User_UserLogin_Handler,
		},
		{
			MethodName: "VerificationLogin",
			Handler:    _User_VerificationLogin_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _User_UpdateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
