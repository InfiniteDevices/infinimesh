// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/node/nodepb/node.proto

package nodepb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Action int32

const (
	Action_NONE  Action = 0
	Action_READ  Action = 1
	Action_WRITE Action = 2
)

var Action_name = map[int32]string{
	0: "NONE",
	1: "READ",
	2: "WRITE",
}

var Action_value = map[string]int32{
	"NONE":  0,
	"READ":  1,
	"WRITE": 2,
}

func (x Action) String() string {
	return proto.EnumName(Action_name, int32(x))
}

func (Action) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_af4ee7a0c648c985, []int{0}
}

type IsAuthorizedRequest struct {
	Node                 string   `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	Account              string   `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	Action               Action   `protobuf:"varint,3,opt,name=action,proto3,enum=infinimesh.node.Action" json:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IsAuthorizedRequest) Reset()         { *m = IsAuthorizedRequest{} }
func (m *IsAuthorizedRequest) String() string { return proto.CompactTextString(m) }
func (*IsAuthorizedRequest) ProtoMessage()    {}
func (*IsAuthorizedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_af4ee7a0c648c985, []int{0}
}

func (m *IsAuthorizedRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsAuthorizedRequest.Unmarshal(m, b)
}
func (m *IsAuthorizedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsAuthorizedRequest.Marshal(b, m, deterministic)
}
func (m *IsAuthorizedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsAuthorizedRequest.Merge(m, src)
}
func (m *IsAuthorizedRequest) XXX_Size() int {
	return xxx_messageInfo_IsAuthorizedRequest.Size(m)
}
func (m *IsAuthorizedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IsAuthorizedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IsAuthorizedRequest proto.InternalMessageInfo

func (m *IsAuthorizedRequest) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

func (m *IsAuthorizedRequest) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *IsAuthorizedRequest) GetAction() Action {
	if m != nil {
		return m.Action
	}
	return Action_NONE
}

type IsAuthorizedResponse struct {
	Decision             *wrappers.BoolValue `protobuf:"bytes,1,opt,name=decision,proto3" json:"decision,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *IsAuthorizedResponse) Reset()         { *m = IsAuthorizedResponse{} }
func (m *IsAuthorizedResponse) String() string { return proto.CompactTextString(m) }
func (*IsAuthorizedResponse) ProtoMessage()    {}
func (*IsAuthorizedResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_af4ee7a0c648c985, []int{1}
}

func (m *IsAuthorizedResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsAuthorizedResponse.Unmarshal(m, b)
}
func (m *IsAuthorizedResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsAuthorizedResponse.Marshal(b, m, deterministic)
}
func (m *IsAuthorizedResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsAuthorizedResponse.Merge(m, src)
}
func (m *IsAuthorizedResponse) XXX_Size() int {
	return xxx_messageInfo_IsAuthorizedResponse.Size(m)
}
func (m *IsAuthorizedResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IsAuthorizedResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IsAuthorizedResponse proto.InternalMessageInfo

func (m *IsAuthorizedResponse) GetDecision() *wrappers.BoolValue {
	if m != nil {
		return m.Decision
	}
	return nil
}

type LoginRequest struct {
	ClientID             string   `protobuf:"bytes,1,opt,name=clientID,proto3" json:"clientID,omitempty"`
	ClientSecret         string   `protobuf:"bytes,2,opt,name=clientSecret,proto3" json:"clientSecret,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_af4ee7a0c648c985, []int{2}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

func (m *LoginRequest) GetClientSecret() string {
	if m != nil {
		return m.ClientSecret
	}
	return ""
}

type LoginResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_af4ee7a0c648c985, []int{3}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

type CreateAccountRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateAccountRequest) Reset()         { *m = CreateAccountRequest{} }
func (m *CreateAccountRequest) String() string { return proto.CompactTextString(m) }
func (*CreateAccountRequest) ProtoMessage()    {}
func (*CreateAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_af4ee7a0c648c985, []int{4}
}

func (m *CreateAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAccountRequest.Unmarshal(m, b)
}
func (m *CreateAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAccountRequest.Marshal(b, m, deterministic)
}
func (m *CreateAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAccountRequest.Merge(m, src)
}
func (m *CreateAccountRequest) XXX_Size() int {
	return xxx_messageInfo_CreateAccountRequest.Size(m)
}
func (m *CreateAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAccountRequest proto.InternalMessageInfo

func (m *CreateAccountRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CreateAccountResponse struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateAccountResponse) Reset()         { *m = CreateAccountResponse{} }
func (m *CreateAccountResponse) String() string { return proto.CompactTextString(m) }
func (*CreateAccountResponse) ProtoMessage()    {}
func (*CreateAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_af4ee7a0c648c985, []int{5}
}

func (m *CreateAccountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateAccountResponse.Unmarshal(m, b)
}
func (m *CreateAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateAccountResponse.Marshal(b, m, deterministic)
}
func (m *CreateAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateAccountResponse.Merge(m, src)
}
func (m *CreateAccountResponse) XXX_Size() int {
	return xxx_messageInfo_CreateAccountResponse.Size(m)
}
func (m *CreateAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateAccountResponse proto.InternalMessageInfo

func (m *CreateAccountResponse) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

type AuthorizeRequest struct {
	Account              string   `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Node                 string   `protobuf:"bytes,2,opt,name=node,proto3" json:"node,omitempty"`
	Action               string   `protobuf:"bytes,3,opt,name=action,proto3" json:"action,omitempty"`
	Inherit              bool     `protobuf:"varint,4,opt,name=inherit,proto3" json:"inherit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthorizeRequest) Reset()         { *m = AuthorizeRequest{} }
func (m *AuthorizeRequest) String() string { return proto.CompactTextString(m) }
func (*AuthorizeRequest) ProtoMessage()    {}
func (*AuthorizeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_af4ee7a0c648c985, []int{6}
}

func (m *AuthorizeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthorizeRequest.Unmarshal(m, b)
}
func (m *AuthorizeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthorizeRequest.Marshal(b, m, deterministic)
}
func (m *AuthorizeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthorizeRequest.Merge(m, src)
}
func (m *AuthorizeRequest) XXX_Size() int {
	return xxx_messageInfo_AuthorizeRequest.Size(m)
}
func (m *AuthorizeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthorizeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthorizeRequest proto.InternalMessageInfo

func (m *AuthorizeRequest) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *AuthorizeRequest) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

func (m *AuthorizeRequest) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *AuthorizeRequest) GetInherit() bool {
	if m != nil {
		return m.Inherit
	}
	return false
}

type AuthorizeResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthorizeResponse) Reset()         { *m = AuthorizeResponse{} }
func (m *AuthorizeResponse) String() string { return proto.CompactTextString(m) }
func (*AuthorizeResponse) ProtoMessage()    {}
func (*AuthorizeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_af4ee7a0c648c985, []int{7}
}

func (m *AuthorizeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthorizeResponse.Unmarshal(m, b)
}
func (m *AuthorizeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthorizeResponse.Marshal(b, m, deterministic)
}
func (m *AuthorizeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthorizeResponse.Merge(m, src)
}
func (m *AuthorizeResponse) XXX_Size() int {
	return xxx_messageInfo_AuthorizeResponse.Size(m)
}
func (m *AuthorizeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthorizeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthorizeResponse proto.InternalMessageInfo

type Context struct {
	Account              string   `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Root                 bool     `protobuf:"varint,2,opt,name=root,proto3" json:"root,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Context) Reset()         { *m = Context{} }
func (m *Context) String() string { return proto.CompactTextString(m) }
func (*Context) ProtoMessage()    {}
func (*Context) Descriptor() ([]byte, []int) {
	return fileDescriptor_af4ee7a0c648c985, []int{8}
}

func (m *Context) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Context.Unmarshal(m, b)
}
func (m *Context) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Context.Marshal(b, m, deterministic)
}
func (m *Context) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Context.Merge(m, src)
}
func (m *Context) XXX_Size() int {
	return xxx_messageInfo_Context.Size(m)
}
func (m *Context) XXX_DiscardUnknown() {
	xxx_messageInfo_Context.DiscardUnknown(m)
}

var xxx_messageInfo_Context proto.InternalMessageInfo

func (m *Context) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *Context) GetRoot() bool {
	if m != nil {
		return m.Root
	}
	return false
}

// TODO maybe path-based; path as string instead of parent id? id = uuid, i.e.
// caller can give it? currently DB assigns it
type CreateObjectRequest struct {
	Context              *Context `protobuf:"bytes,1,opt,name=context,proto3" json:"context,omitempty"`
	Parent               string   `protobuf:"bytes,2,opt,name=parent,proto3" json:"parent,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateObjectRequest) Reset()         { *m = CreateObjectRequest{} }
func (m *CreateObjectRequest) String() string { return proto.CompactTextString(m) }
func (*CreateObjectRequest) ProtoMessage()    {}
func (*CreateObjectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_af4ee7a0c648c985, []int{9}
}

func (m *CreateObjectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateObjectRequest.Unmarshal(m, b)
}
func (m *CreateObjectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateObjectRequest.Marshal(b, m, deterministic)
}
func (m *CreateObjectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateObjectRequest.Merge(m, src)
}
func (m *CreateObjectRequest) XXX_Size() int {
	return xxx_messageInfo_CreateObjectRequest.Size(m)
}
func (m *CreateObjectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateObjectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateObjectRequest proto.InternalMessageInfo

func (m *CreateObjectRequest) GetContext() *Context {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *CreateObjectRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *CreateObjectRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CreateObjectResponse struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateObjectResponse) Reset()         { *m = CreateObjectResponse{} }
func (m *CreateObjectResponse) String() string { return proto.CompactTextString(m) }
func (*CreateObjectResponse) ProtoMessage()    {}
func (*CreateObjectResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_af4ee7a0c648c985, []int{10}
}

func (m *CreateObjectResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateObjectResponse.Unmarshal(m, b)
}
func (m *CreateObjectResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateObjectResponse.Marshal(b, m, deterministic)
}
func (m *CreateObjectResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateObjectResponse.Merge(m, src)
}
func (m *CreateObjectResponse) XXX_Size() int {
	return xxx_messageInfo_CreateObjectResponse.Size(m)
}
func (m *CreateObjectResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateObjectResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateObjectResponse proto.InternalMessageInfo

func (m *CreateObjectResponse) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

type Object struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Object) Reset()         { *m = Object{} }
func (m *Object) String() string { return proto.CompactTextString(m) }
func (*Object) ProtoMessage()    {}
func (*Object) Descriptor() ([]byte, []int) {
	return fileDescriptor_af4ee7a0c648c985, []int{11}
}

func (m *Object) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Object.Unmarshal(m, b)
}
func (m *Object) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Object.Marshal(b, m, deterministic)
}
func (m *Object) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Object.Merge(m, src)
}
func (m *Object) XXX_Size() int {
	return xxx_messageInfo_Object.Size(m)
}
func (m *Object) XXX_DiscardUnknown() {
	xxx_messageInfo_Object.DiscardUnknown(m)
}

var xxx_messageInfo_Object proto.InternalMessageInfo

func (m *Object) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *Object) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Device struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Device) Reset()         { *m = Device{} }
func (m *Device) String() string { return proto.CompactTextString(m) }
func (*Device) ProtoMessage()    {}
func (*Device) Descriptor() ([]byte, []int) {
	return fileDescriptor_af4ee7a0c648c985, []int{12}
}

func (m *Device) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Device.Unmarshal(m, b)
}
func (m *Device) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Device.Marshal(b, m, deterministic)
}
func (m *Device) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Device.Merge(m, src)
}
func (m *Device) XXX_Size() int {
	return xxx_messageInfo_Device.Size(m)
}
func (m *Device) XXX_DiscardUnknown() {
	xxx_messageInfo_Device.DiscardUnknown(m)
}

var xxx_messageInfo_Device proto.InternalMessageInfo

func (m *Device) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *Device) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ListObjectsRequest struct {
	Account              string   `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListObjectsRequest) Reset()         { *m = ListObjectsRequest{} }
func (m *ListObjectsRequest) String() string { return proto.CompactTextString(m) }
func (*ListObjectsRequest) ProtoMessage()    {}
func (*ListObjectsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_af4ee7a0c648c985, []int{13}
}

func (m *ListObjectsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListObjectsRequest.Unmarshal(m, b)
}
func (m *ListObjectsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListObjectsRequest.Marshal(b, m, deterministic)
}
func (m *ListObjectsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListObjectsRequest.Merge(m, src)
}
func (m *ListObjectsRequest) XXX_Size() int {
	return xxx_messageInfo_ListObjectsRequest.Size(m)
}
func (m *ListObjectsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListObjectsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListObjectsRequest proto.InternalMessageInfo

func (m *ListObjectsRequest) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

type ListObjectsResponse struct {
	Objects              []*ObjectList `protobuf:"bytes,1,rep,name=objects,proto3" json:"objects,omitempty"`
	Devices              []*Device     `protobuf:"bytes,2,rep,name=devices,proto3" json:"devices,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ListObjectsResponse) Reset()         { *m = ListObjectsResponse{} }
func (m *ListObjectsResponse) String() string { return proto.CompactTextString(m) }
func (*ListObjectsResponse) ProtoMessage()    {}
func (*ListObjectsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_af4ee7a0c648c985, []int{14}
}

func (m *ListObjectsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListObjectsResponse.Unmarshal(m, b)
}
func (m *ListObjectsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListObjectsResponse.Marshal(b, m, deterministic)
}
func (m *ListObjectsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListObjectsResponse.Merge(m, src)
}
func (m *ListObjectsResponse) XXX_Size() int {
	return xxx_messageInfo_ListObjectsResponse.Size(m)
}
func (m *ListObjectsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListObjectsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListObjectsResponse proto.InternalMessageInfo

func (m *ListObjectsResponse) GetObjects() []*ObjectList {
	if m != nil {
		return m.Objects
	}
	return nil
}

func (m *ListObjectsResponse) GetDevices() []*Device {
	if m != nil {
		return m.Devices
	}
	return nil
}

type ObjectList struct {
	Uid                  string        `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Objects              []*ObjectList `protobuf:"bytes,3,rep,name=objects,proto3" json:"objects,omitempty"`
	Devices              []*Device     `protobuf:"bytes,4,rep,name=devices,proto3" json:"devices,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ObjectList) Reset()         { *m = ObjectList{} }
func (m *ObjectList) String() string { return proto.CompactTextString(m) }
func (*ObjectList) ProtoMessage()    {}
func (*ObjectList) Descriptor() ([]byte, []int) {
	return fileDescriptor_af4ee7a0c648c985, []int{15}
}

func (m *ObjectList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectList.Unmarshal(m, b)
}
func (m *ObjectList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectList.Marshal(b, m, deterministic)
}
func (m *ObjectList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectList.Merge(m, src)
}
func (m *ObjectList) XXX_Size() int {
	return xxx_messageInfo_ObjectList.Size(m)
}
func (m *ObjectList) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectList.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectList proto.InternalMessageInfo

func (m *ObjectList) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ObjectList) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ObjectList) GetObjects() []*ObjectList {
	if m != nil {
		return m.Objects
	}
	return nil
}

func (m *ObjectList) GetDevices() []*Device {
	if m != nil {
		return m.Devices
	}
	return nil
}

func init() {
	proto.RegisterEnum("infinimesh.node.Action", Action_name, Action_value)
	proto.RegisterType((*IsAuthorizedRequest)(nil), "infinimesh.node.IsAuthorizedRequest")
	proto.RegisterType((*IsAuthorizedResponse)(nil), "infinimesh.node.IsAuthorizedResponse")
	proto.RegisterType((*LoginRequest)(nil), "infinimesh.node.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "infinimesh.node.LoginResponse")
	proto.RegisterType((*CreateAccountRequest)(nil), "infinimesh.node.CreateAccountRequest")
	proto.RegisterType((*CreateAccountResponse)(nil), "infinimesh.node.CreateAccountResponse")
	proto.RegisterType((*AuthorizeRequest)(nil), "infinimesh.node.AuthorizeRequest")
	proto.RegisterType((*AuthorizeResponse)(nil), "infinimesh.node.AuthorizeResponse")
	proto.RegisterType((*Context)(nil), "infinimesh.node.Context")
	proto.RegisterType((*CreateObjectRequest)(nil), "infinimesh.node.CreateObjectRequest")
	proto.RegisterType((*CreateObjectResponse)(nil), "infinimesh.node.CreateObjectResponse")
	proto.RegisterType((*Object)(nil), "infinimesh.node.Object")
	proto.RegisterType((*Device)(nil), "infinimesh.node.Device")
	proto.RegisterType((*ListObjectsRequest)(nil), "infinimesh.node.ListObjectsRequest")
	proto.RegisterType((*ListObjectsResponse)(nil), "infinimesh.node.ListObjectsResponse")
	proto.RegisterType((*ObjectList)(nil), "infinimesh.node.ObjectList")
}

func init() { proto.RegisterFile("pkg/node/nodepb/node.proto", fileDescriptor_af4ee7a0c648c985) }

var fileDescriptor_af4ee7a0c648c985 = []byte{
	// 653 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xcb, 0x6e, 0xd3, 0x4c,
	0x14, 0xfe, 0x9d, 0xe4, 0xb7, 0x9d, 0xd3, 0x94, 0x86, 0x49, 0xa1, 0x96, 0x91, 0x50, 0x18, 0x5a,
	0x08, 0x5d, 0x38, 0xc2, 0x08, 0x58, 0xa7, 0x97, 0x45, 0xa5, 0x92, 0x4a, 0x53, 0x04, 0x08, 0x24,
	0x24, 0xc7, 0x99, 0xa6, 0x03, 0xa9, 0xc7, 0xf8, 0x02, 0x88, 0x0d, 0xef, 0x81, 0x78, 0x10, 0x1e,
	0x0f, 0xd9, 0x33, 0xe3, 0xda, 0x75, 0x9a, 0x54, 0x6c, 0x92, 0xb9, 0x7c, 0xe7, 0xf2, 0x7d, 0x67,
	0xce, 0x31, 0xd8, 0xe1, 0xe7, 0xd9, 0x30, 0xe0, 0x53, 0x9a, 0xff, 0x84, 0x93, 0xfc, 0xcf, 0x09,
	0x23, 0x9e, 0x70, 0xb4, 0xc1, 0x82, 0x33, 0x16, 0xb0, 0x0b, 0x1a, 0x9f, 0x3b, 0xd9, 0xb1, 0x7d,
	0x7f, 0xc6, 0xf9, 0x6c, 0x4e, 0x87, 0xf9, 0xf5, 0x24, 0x3d, 0x1b, 0x7e, 0x8b, 0xbc, 0x30, 0xa4,
	0x51, 0x2c, 0x0c, 0x70, 0x02, 0xbd, 0xa3, 0x78, 0x94, 0x26, 0xe7, 0x3c, 0x62, 0x3f, 0xe8, 0x94,
	0xd0, 0x2f, 0x29, 0x8d, 0x13, 0x84, 0xa0, 0x95, 0x99, 0x5b, 0x5a, 0x5f, 0x1b, 0xb4, 0x49, 0xbe,
	0x46, 0x16, 0x18, 0x9e, 0xef, 0xf3, 0x34, 0x48, 0xac, 0x46, 0x7e, 0xac, 0xb6, 0x68, 0x08, 0xba,
	0xe7, 0x27, 0x8c, 0x07, 0x56, 0xb3, 0xaf, 0x0d, 0x6e, 0xb9, 0x5b, 0xce, 0x95, 0x34, 0x9c, 0x51,
	0x7e, 0x4d, 0x24, 0x0c, 0x8f, 0x61, 0xb3, 0x1a, 0x35, 0x0e, 0x79, 0x10, 0x53, 0xf4, 0x02, 0xcc,
	0x29, 0xf5, 0x59, 0x9c, 0xb9, 0xca, 0x42, 0xaf, 0xb9, 0xb6, 0x23, 0x08, 0x38, 0x8a, 0x80, 0xb3,
	0xc7, 0xf9, 0xfc, 0x8d, 0x37, 0x4f, 0x29, 0x29, 0xb0, 0x78, 0x0c, 0x9d, 0x63, 0x3e, 0x63, 0x81,
	0x4a, 0xdf, 0x06, 0xd3, 0x9f, 0x33, 0x1a, 0x24, 0x47, 0x07, 0x92, 0x42, 0xb1, 0x47, 0x18, 0x3a,
	0x62, 0x7d, 0x4a, 0xfd, 0x88, 0x2a, 0x2e, 0x95, 0x33, 0xbc, 0x01, 0xeb, 0xd2, 0x9f, 0x48, 0x0c,
	0xef, 0xc2, 0xe6, 0x7e, 0x44, 0xbd, 0x84, 0x8e, 0x04, 0xe5, 0xb2, 0x4e, 0xde, 0xc5, 0xa5, 0x4e,
	0xde, 0x05, 0xc5, 0x4f, 0xe0, 0xce, 0x15, 0xac, 0x64, 0xd7, 0x85, 0x66, 0xca, 0xa6, 0x12, 0x9b,
	0x2d, 0x71, 0x04, 0xdd, 0x42, 0x05, 0xe5, 0xb2, 0x24, 0xb3, 0x56, 0x95, 0x59, 0x15, 0xa5, 0x51,
	0x2a, 0xca, 0xdd, 0x8a, 0xf4, 0x6d, 0xa5, 0x70, 0xe6, 0x85, 0x05, 0xe7, 0x34, 0x62, 0x89, 0xd5,
	0xea, 0x6b, 0x03, 0x93, 0xa8, 0x2d, 0xee, 0xc1, 0xed, 0x52, 0x4c, 0xc9, 0xef, 0x25, 0x18, 0xfb,
	0x3c, 0x48, 0xe8, 0xf7, 0x15, 0xf1, 0x23, 0xce, 0x85, 0x62, 0x26, 0xc9, 0xd7, 0x38, 0x85, 0x9e,
	0x20, 0x7b, 0x32, 0xf9, 0x44, 0xfd, 0x42, 0x17, 0x17, 0x0c, 0x5f, 0xf8, 0x93, 0x75, 0xb4, 0x6a,
	0x4f, 0x42, 0xc6, 0x23, 0x0a, 0x98, 0x51, 0x09, 0xbd, 0x88, 0x16, 0xcf, 0x4b, 0xee, 0x0a, 0x8d,
	0x9b, 0x25, 0x8d, 0x07, 0xaa, 0x1e, 0x2a, 0xec, 0xb5, 0x12, 0x3b, 0xa0, 0x0b, 0x4c, 0xfd, 0x6e,
	0xa1, 0x67, 0x07, 0xf4, 0x03, 0xfa, 0x95, 0xf9, 0x74, 0x09, 0xbe, 0x51, 0xc1, 0xa3, 0x63, 0x16,
	0x27, 0x22, 0x46, 0xbc, 0xb2, 0x88, 0xf8, 0x27, 0xf4, 0x2a, 0x78, 0x99, 0xf8, 0x73, 0x30, 0xb8,
	0x38, 0xb2, 0xb4, 0x7e, 0x73, 0xb0, 0xe6, 0xde, 0xab, 0x09, 0x26, 0x4c, 0x32, 0x63, 0xa2, 0xb0,
	0xe8, 0x29, 0x18, 0xd3, 0x3c, 0xdb, 0xd8, 0x6a, 0xe4, 0x66, 0xf5, 0xd6, 0x13, 0x6c, 0x88, 0xc2,
	0xe1, 0xdf, 0x1a, 0xc0, 0xa5, 0xab, 0x9b, 0xb1, 0x2c, 0xa7, 0xd7, 0xfc, 0xb7, 0xf4, 0x5a, 0x37,
	0x4b, 0x6f, 0xf7, 0x31, 0xe8, 0x62, 0x58, 0x20, 0x13, 0x5a, 0xe3, 0x93, 0xf1, 0x61, 0xf7, 0xbf,
	0x6c, 0x45, 0x0e, 0x47, 0x07, 0x5d, 0x0d, 0xb5, 0xe1, 0xff, 0xb7, 0xe4, 0xe8, 0xf5, 0x61, 0xb7,
	0xe1, 0xfe, 0x6a, 0x80, 0x21, 0x3b, 0x0c, 0x7d, 0x84, 0xf5, 0x4a, 0xcb, 0xa1, 0x9d, 0xfa, 0x73,
	0x5b, 0xd0, 0xbe, 0xf6, 0xa3, 0x55, 0x30, 0x59, 0x9d, 0x0f, 0xd0, 0x29, 0xcf, 0x2b, 0xb4, 0x5d,
	0xb3, 0x5b, 0x30, 0x44, 0xed, 0x9d, 0x15, 0x28, 0xe9, 0x9c, 0x40, 0xbb, 0x38, 0x45, 0x0f, 0xea,
	0xa3, 0xf3, 0xca, 0x80, 0xb0, 0xf1, 0x32, 0x88, 0xf0, 0xe9, 0xfe, 0xd1, 0x60, 0x5d, 0x14, 0xe4,
	0x94, 0x46, 0xf9, 0x6b, 0x7e, 0x05, 0x9d, 0x72, 0xc7, 0x2c, 0xa0, 0xb0, 0xa0, 0x8f, 0xed, 0xad,
	0x6b, 0xca, 0x8c, 0xde, 0xc1, 0x5a, 0xe9, 0x19, 0xa3, 0x87, 0x35, 0x5c, 0xbd, 0x29, 0xec, 0xed,
	0xe5, 0x20, 0x91, 0xfa, 0x9e, 0xf9, 0x5e, 0x17, 0xdf, 0xb5, 0x89, 0x9e, 0xcf, 0xfc, 0x67, 0x7f,
	0x03, 0x00, 0x00, 0xff, 0xff, 0x7b, 0xe0, 0xc9, 0x16, 0xf1, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AccountClient is the client API for Account service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountClient interface {
	CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error)
	IsAuthorized(ctx context.Context, in *IsAuthorizedRequest, opts ...grpc.CallOption) (*IsAuthorizedResponse, error)
	Authorize(ctx context.Context, in *AuthorizeRequest, opts ...grpc.CallOption) (*AuthorizeResponse, error)
}

type accountClient struct {
	cc *grpc.ClientConn
}

func NewAccountClient(cc *grpc.ClientConn) AccountClient {
	return &accountClient{cc}
}

func (c *accountClient) CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*CreateAccountResponse, error) {
	out := new(CreateAccountResponse)
	err := c.cc.Invoke(ctx, "/infinimesh.node.Account/CreateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) IsAuthorized(ctx context.Context, in *IsAuthorizedRequest, opts ...grpc.CallOption) (*IsAuthorizedResponse, error) {
	out := new(IsAuthorizedResponse)
	err := c.cc.Invoke(ctx, "/infinimesh.node.Account/IsAuthorized", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) Authorize(ctx context.Context, in *AuthorizeRequest, opts ...grpc.CallOption) (*AuthorizeResponse, error) {
	out := new(AuthorizeResponse)
	err := c.cc.Invoke(ctx, "/infinimesh.node.Account/Authorize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServer is the server API for Account service.
type AccountServer interface {
	CreateAccount(context.Context, *CreateAccountRequest) (*CreateAccountResponse, error)
	IsAuthorized(context.Context, *IsAuthorizedRequest) (*IsAuthorizedResponse, error)
	Authorize(context.Context, *AuthorizeRequest) (*AuthorizeResponse, error)
}

func RegisterAccountServer(s *grpc.Server, srv AccountServer) {
	s.RegisterService(&_Account_serviceDesc, srv)
}

func _Account_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infinimesh.node.Account/CreateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).CreateAccount(ctx, req.(*CreateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_IsAuthorized_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsAuthorizedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).IsAuthorized(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infinimesh.node.Account/IsAuthorized",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).IsAuthorized(ctx, req.(*IsAuthorizedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_Authorize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).Authorize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infinimesh.node.Account/Authorize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).Authorize(ctx, req.(*AuthorizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Account_serviceDesc = grpc.ServiceDesc{
	ServiceName: "infinimesh.node.Account",
	HandlerType: (*AccountServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAccount",
			Handler:    _Account_CreateAccount_Handler,
		},
		{
			MethodName: "IsAuthorized",
			Handler:    _Account_IsAuthorized_Handler,
		},
		{
			MethodName: "Authorize",
			Handler:    _Account_Authorize_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/node/nodepb/node.proto",
}

// ObjectServiceClient is the client API for ObjectService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ObjectServiceClient interface {
	CreateObject(ctx context.Context, in *CreateObjectRequest, opts ...grpc.CallOption) (*Object, error)
	ListObjects(ctx context.Context, in *ListObjectsRequest, opts ...grpc.CallOption) (*ListObjectsResponse, error)
}

type objectServiceClient struct {
	cc *grpc.ClientConn
}

func NewObjectServiceClient(cc *grpc.ClientConn) ObjectServiceClient {
	return &objectServiceClient{cc}
}

func (c *objectServiceClient) CreateObject(ctx context.Context, in *CreateObjectRequest, opts ...grpc.CallOption) (*Object, error) {
	out := new(Object)
	err := c.cc.Invoke(ctx, "/infinimesh.node.ObjectService/CreateObject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectServiceClient) ListObjects(ctx context.Context, in *ListObjectsRequest, opts ...grpc.CallOption) (*ListObjectsResponse, error) {
	out := new(ListObjectsResponse)
	err := c.cc.Invoke(ctx, "/infinimesh.node.ObjectService/ListObjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ObjectServiceServer is the server API for ObjectService service.
type ObjectServiceServer interface {
	CreateObject(context.Context, *CreateObjectRequest) (*Object, error)
	ListObjects(context.Context, *ListObjectsRequest) (*ListObjectsResponse, error)
}

func RegisterObjectServiceServer(s *grpc.Server, srv ObjectServiceServer) {
	s.RegisterService(&_ObjectService_serviceDesc, srv)
}

func _ObjectService_CreateObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectServiceServer).CreateObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infinimesh.node.ObjectService/CreateObject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectServiceServer).CreateObject(ctx, req.(*CreateObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectService_ListObjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListObjectsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectServiceServer).ListObjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/infinimesh.node.ObjectService/ListObjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectServiceServer).ListObjects(ctx, req.(*ListObjectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ObjectService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "infinimesh.node.ObjectService",
	HandlerType: (*ObjectServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateObject",
			Handler:    _ObjectService_CreateObject_Handler,
		},
		{
			MethodName: "ListObjects",
			Handler:    _ObjectService_ListObjects_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/node/nodepb/node.proto",
}