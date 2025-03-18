// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.29.3
// source: idl/auth.proto

package auth_service

import (
	context "context"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  bool   `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_idl_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_idl_auth_proto_rawDescGZIP(), []int{0}
}

func (x *Status) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *Status) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type CreateAuthReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid uint64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *CreateAuthReq) Reset() {
	*x = CreateAuthReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAuthReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAuthReq) ProtoMessage() {}

func (x *CreateAuthReq) ProtoReflect() protoreflect.Message {
	mi := &file_idl_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAuthReq.ProtoReflect.Descriptor instead.
func (*CreateAuthReq) Descriptor() ([]byte, []int) {
	return file_idl_auth_proto_rawDescGZIP(), []int{1}
}

func (x *CreateAuthReq) GetUid() uint64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type CreateAuthResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status       *Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	AccessToken  string  `protobuf:"bytes,2,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	RefreshToken string  `protobuf:"bytes,3,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
}

func (x *CreateAuthResp) Reset() {
	*x = CreateAuthResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAuthResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAuthResp) ProtoMessage() {}

func (x *CreateAuthResp) ProtoReflect() protoreflect.Message {
	mi := &file_idl_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAuthResp.ProtoReflect.Descriptor instead.
func (*CreateAuthResp) Descriptor() ([]byte, []int) {
	return file_idl_auth_proto_rawDescGZIP(), []int{2}
}

func (x *CreateAuthResp) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *CreateAuthResp) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *CreateAuthResp) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type UpdateAuthReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RefreshToken string `protobuf:"bytes,1,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
}

func (x *UpdateAuthReq) Reset() {
	*x = UpdateAuthReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAuthReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAuthReq) ProtoMessage() {}

func (x *UpdateAuthReq) ProtoReflect() protoreflect.Message {
	mi := &file_idl_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAuthReq.ProtoReflect.Descriptor instead.
func (*UpdateAuthReq) Descriptor() ([]byte, []int) {
	return file_idl_auth_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateAuthReq) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type UpdateAuthResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status        *Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	AccessToken   string  `protobuf:"bytes,2,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	AccessExpire  string  `protobuf:"bytes,3,opt,name=access_expire,json=accessExpire,proto3" json:"access_expire,omitempty"`
	RefreshToken  string  `protobuf:"bytes,4,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	RefreshExpire string  `protobuf:"bytes,5,opt,name=refresh_expire,json=refreshExpire,proto3" json:"refresh_expire,omitempty"`
}

func (x *UpdateAuthResp) Reset() {
	*x = UpdateAuthResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAuthResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAuthResp) ProtoMessage() {}

func (x *UpdateAuthResp) ProtoReflect() protoreflect.Message {
	mi := &file_idl_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAuthResp.ProtoReflect.Descriptor instead.
func (*UpdateAuthResp) Descriptor() ([]byte, []int) {
	return file_idl_auth_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateAuthResp) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *UpdateAuthResp) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *UpdateAuthResp) GetAccessExpire() string {
	if x != nil {
		return x.AccessExpire
	}
	return ""
}

func (x *UpdateAuthResp) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *UpdateAuthResp) GetRefreshExpire() string {
	if x != nil {
		return x.RefreshExpire
	}
	return ""
}

type EmailReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=Email,proto3" json:"Email,omitempty"`
}

func (x *EmailReq) Reset() {
	*x = EmailReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailReq) ProtoMessage() {}

func (x *EmailReq) ProtoReflect() protoreflect.Message {
	mi := &file_idl_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailReq.ProtoReflect.Descriptor instead.
func (*EmailReq) Descriptor() ([]byte, []int) {
	return file_idl_auth_proto_rawDescGZIP(), []int{5}
}

func (x *EmailReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type EmailResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *EmailResp) Reset() {
	*x = EmailResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_auth_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailResp) ProtoMessage() {}

func (x *EmailResp) ProtoReflect() protoreflect.Message {
	mi := &file_idl_auth_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailResp.ProtoReflect.Descriptor instead.
func (*EmailResp) Descriptor() ([]byte, []int) {
	return file_idl_auth_proto_rawDescGZIP(), []int{6}
}

func (x *EmailResp) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

var File_idl_auth_proto protoreflect.FileDescriptor

var file_idl_auth_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x69, 0x64, 0x6c, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x61, 0x75, 0x74, 0x68, 0x22, 0x3a, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x21, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68,
	0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x7e, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41,
	0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x12, 0x24, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x21, 0x0a,
	0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x34, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41,
	0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xca, 0x01, 0x0a, 0x0e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x12, 0x24,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x12, 0x23, 0x0a, 0x0d,
	0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x22, 0x20, 0x0a, 0x08, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x31, 0x0a, 0x09, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x12, 0x24, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xaf, 0x01,
	0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a,
	0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x12, 0x13, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71,
	0x1a, 0x14, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75,
	0x74, 0x68, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x12, 0x13, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x00, 0x12, 0x2a, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x0e, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x61,
	0x75, 0x74, 0x68, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42,
	0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x49, 0x54,
	0x75, 0x2d, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x57, 0x65, 0x47, 0x6f, 0x2f, 0x69, 0x74, 0x75, 0x5f,
	0x72, 0x70, 0x63, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x6b, 0x69, 0x74,
	0x65, 0x78, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_idl_auth_proto_rawDescOnce sync.Once
	file_idl_auth_proto_rawDescData = file_idl_auth_proto_rawDesc
)

func file_idl_auth_proto_rawDescGZIP() []byte {
	file_idl_auth_proto_rawDescOnce.Do(func() {
		file_idl_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_idl_auth_proto_rawDescData)
	})
	return file_idl_auth_proto_rawDescData
}

var file_idl_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_idl_auth_proto_goTypes = []interface{}{
	(*Status)(nil),         // 0: auth.Status
	(*CreateAuthReq)(nil),  // 1: auth.CreateAuthReq
	(*CreateAuthResp)(nil), // 2: auth.CreateAuthResp
	(*UpdateAuthReq)(nil),  // 3: auth.UpdateAuthReq
	(*UpdateAuthResp)(nil), // 4: auth.UpdateAuthResp
	(*EmailReq)(nil),       // 5: auth.EmailReq
	(*EmailResp)(nil),      // 6: auth.EmailResp
}
var file_idl_auth_proto_depIdxs = []int32{
	0, // 0: auth.CreateAuthResp.status:type_name -> auth.Status
	0, // 1: auth.UpdateAuthResp.status:type_name -> auth.Status
	0, // 2: auth.EmailResp.status:type_name -> auth.Status
	1, // 3: auth.AuthService.CreateAuth:input_type -> auth.CreateAuthReq
	3, // 4: auth.AuthService.UpdateAuth:input_type -> auth.UpdateAuthReq
	5, // 5: auth.AuthService.Email:input_type -> auth.EmailReq
	2, // 6: auth.AuthService.CreateAuth:output_type -> auth.CreateAuthResp
	4, // 7: auth.AuthService.UpdateAuth:output_type -> auth.UpdateAuthResp
	6, // 8: auth.AuthService.Email:output_type -> auth.EmailResp
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_idl_auth_proto_init() }
func file_idl_auth_proto_init() {
	if File_idl_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_idl_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_idl_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAuthReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_idl_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAuthResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_idl_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAuthReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_idl_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAuthResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_idl_auth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_idl_auth_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_idl_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_idl_auth_proto_goTypes,
		DependencyIndexes: file_idl_auth_proto_depIdxs,
		MessageInfos:      file_idl_auth_proto_msgTypes,
	}.Build()
	File_idl_auth_proto = out.File
	file_idl_auth_proto_rawDesc = nil
	file_idl_auth_proto_goTypes = nil
	file_idl_auth_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.12.3. DO NOT EDIT.

type AuthService interface {
	CreateAuth(ctx context.Context, req *CreateAuthReq) (res *CreateAuthResp, err error)
	UpdateAuth(ctx context.Context, req *UpdateAuthReq) (res *UpdateAuthResp, err error)
	Email(ctx context.Context, req *EmailReq) (res *EmailResp, err error)
}
