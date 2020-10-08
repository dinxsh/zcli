// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.7.1
// source: zeropsDaemonProtocol.proto

package zeropsDaemonProtocol

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ErrorCode int32

const (
	ErrorCode_NO_ERROR              ErrorCode = 0
	ErrorCode_INTERNAL_SERVER_ERROR ErrorCode = 1
)

// Enum value maps for ErrorCode.
var (
	ErrorCode_name = map[int32]string{
		0: "NO_ERROR",
		1: "INTERNAL_SERVER_ERROR",
	}
	ErrorCode_value = map[string]int32{
		"NO_ERROR":              0,
		"INTERNAL_SERVER_ERROR": 1,
	}
)

func (x ErrorCode) Enum() *ErrorCode {
	p := new(ErrorCode)
	*p = x
	return p
}

func (x ErrorCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorCode) Descriptor() protoreflect.EnumDescriptor {
	return file_zeropsDaemonProtocol_proto_enumTypes[0].Descriptor()
}

func (ErrorCode) Type() protoreflect.EnumType {
	return &file_zeropsDaemonProtocol_proto_enumTypes[0]
}

func (x ErrorCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorCode.Descriptor instead.
func (ErrorCode) EnumDescriptor() ([]byte, []int) {
	return file_zeropsDaemonProtocol_proto_rawDescGZIP(), []int{0}
}

type VpnStatus int32

const (
	VpnStatus_INACTIVE VpnStatus = 0
	VpnStatus_ACTIVE   VpnStatus = 1
)

// Enum value maps for VpnStatus.
var (
	VpnStatus_name = map[int32]string{
		0: "INACTIVE",
		1: "ACTIVE",
	}
	VpnStatus_value = map[string]int32{
		"INACTIVE": 0,
		"ACTIVE":   1,
	}
)

func (x VpnStatus) Enum() *VpnStatus {
	p := new(VpnStatus)
	*p = x
	return p
}

func (x VpnStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VpnStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_zeropsDaemonProtocol_proto_enumTypes[1].Descriptor()
}

func (VpnStatus) Type() protoreflect.EnumType {
	return &file_zeropsDaemonProtocol_proto_enumTypes[1]
}

func (x VpnStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VpnStatus.Descriptor instead.
func (VpnStatus) EnumDescriptor() ([]byte, []int) {
	return file_zeropsDaemonProtocol_proto_rawDescGZIP(), []int{1}
}

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    ErrorCode `protobuf:"varint,1,opt,name=code,proto3,enum=zeropsDaemonProtocol.ErrorCode" json:"code,omitempty"`
	Message string    `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zeropsDaemonProtocol_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_zeropsDaemonProtocol_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_zeropsDaemonProtocol_proto_rawDescGZIP(), []int{0}
}

func (x *Error) GetCode() ErrorCode {
	if x != nil {
		return x.Code
	}
	return ErrorCode_NO_ERROR
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type StartVpnRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiAddress string `protobuf:"bytes,1,opt,name=apiAddress,proto3" json:"apiAddress,omitempty"`
	VpnAddress string `protobuf:"bytes,2,opt,name=vpnAddress,proto3" json:"vpnAddress,omitempty"`
	ProjectId  string `protobuf:"bytes,3,opt,name=projectId,proto3" json:"projectId,omitempty"`
	Token      string `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *StartVpnRequest) Reset() {
	*x = StartVpnRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zeropsDaemonProtocol_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartVpnRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartVpnRequest) ProtoMessage() {}

func (x *StartVpnRequest) ProtoReflect() protoreflect.Message {
	mi := &file_zeropsDaemonProtocol_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartVpnRequest.ProtoReflect.Descriptor instead.
func (*StartVpnRequest) Descriptor() ([]byte, []int) {
	return file_zeropsDaemonProtocol_proto_rawDescGZIP(), []int{1}
}

func (x *StartVpnRequest) GetApiAddress() string {
	if x != nil {
		return x.ApiAddress
	}
	return ""
}

func (x *StartVpnRequest) GetVpnAddress() string {
	if x != nil {
		return x.VpnAddress
	}
	return ""
}

func (x *StartVpnRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *StartVpnRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type StartVpnResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error *Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *StartVpnResponse) Reset() {
	*x = StartVpnResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zeropsDaemonProtocol_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartVpnResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartVpnResponse) ProtoMessage() {}

func (x *StartVpnResponse) ProtoReflect() protoreflect.Message {
	mi := &file_zeropsDaemonProtocol_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartVpnResponse.ProtoReflect.Descriptor instead.
func (*StartVpnResponse) Descriptor() ([]byte, []int) {
	return file_zeropsDaemonProtocol_proto_rawDescGZIP(), []int{2}
}

func (x *StartVpnResponse) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type StopVpnRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StopVpnRequest) Reset() {
	*x = StopVpnRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zeropsDaemonProtocol_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopVpnRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopVpnRequest) ProtoMessage() {}

func (x *StopVpnRequest) ProtoReflect() protoreflect.Message {
	mi := &file_zeropsDaemonProtocol_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopVpnRequest.ProtoReflect.Descriptor instead.
func (*StopVpnRequest) Descriptor() ([]byte, []int) {
	return file_zeropsDaemonProtocol_proto_rawDescGZIP(), []int{3}
}

type StopVpnResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error *Error `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *StopVpnResponse) Reset() {
	*x = StopVpnResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zeropsDaemonProtocol_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopVpnResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopVpnResponse) ProtoMessage() {}

func (x *StopVpnResponse) ProtoReflect() protoreflect.Message {
	mi := &file_zeropsDaemonProtocol_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopVpnResponse.ProtoReflect.Descriptor instead.
func (*StopVpnResponse) Descriptor() ([]byte, []int) {
	return file_zeropsDaemonProtocol_proto_rawDescGZIP(), []int{4}
}

func (x *StopVpnResponse) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

type StatusVpnRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StatusVpnRequest) Reset() {
	*x = StatusVpnRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zeropsDaemonProtocol_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusVpnRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusVpnRequest) ProtoMessage() {}

func (x *StatusVpnRequest) ProtoReflect() protoreflect.Message {
	mi := &file_zeropsDaemonProtocol_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusVpnRequest.ProtoReflect.Descriptor instead.
func (*StatusVpnRequest) Descriptor() ([]byte, []int) {
	return file_zeropsDaemonProtocol_proto_rawDescGZIP(), []int{5}
}

type StatusVpnResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status VpnStatus `protobuf:"varint,1,opt,name=status,proto3,enum=zeropsDaemonProtocol.VpnStatus" json:"status,omitempty"`
	Error  *Error    `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *StatusVpnResponse) Reset() {
	*x = StatusVpnResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zeropsDaemonProtocol_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusVpnResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusVpnResponse) ProtoMessage() {}

func (x *StatusVpnResponse) ProtoReflect() protoreflect.Message {
	mi := &file_zeropsDaemonProtocol_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusVpnResponse.ProtoReflect.Descriptor instead.
func (*StatusVpnResponse) Descriptor() ([]byte, []int) {
	return file_zeropsDaemonProtocol_proto_rawDescGZIP(), []int{6}
}

func (x *StatusVpnResponse) GetStatus() VpnStatus {
	if x != nil {
		return x.Status
	}
	return VpnStatus_INACTIVE
}

func (x *StatusVpnResponse) GetError() *Error {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_zeropsDaemonProtocol_proto protoreflect.FileDescriptor

var file_zeropsDaemonProtocol_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x7a, 0x65, 0x72, 0x6f, 0x70, 0x73, 0x44, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x7a, 0x65,
	0x72, 0x6f, 0x70, 0x73, 0x44, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x22, 0x56, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x33, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x7a, 0x65, 0x72, 0x6f,
	0x70, 0x73, 0x44, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x85, 0x01, 0x0a, 0x0f, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x56, 0x70, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e,
	0x0a, 0x0a, 0x61, 0x70, 0x69, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1e,
	0x0a, 0x0a, 0x76, 0x70, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x76, 0x70, 0x6e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1c,
	0x0a, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x45, 0x0a, 0x10, 0x73, 0x74, 0x61, 0x72, 0x74, 0x56, 0x70, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x7a, 0x65, 0x72, 0x6f, 0x70, 0x73, 0x44, 0x61,
	0x65, 0x6d, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x10, 0x0a, 0x0e, 0x73, 0x74, 0x6f,
	0x70, 0x56, 0x70, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x44, 0x0a, 0x0f, 0x73,
	0x74, 0x6f, 0x70, 0x56, 0x70, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x7a, 0x65, 0x72, 0x6f, 0x70, 0x73, 0x44, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x22, 0x12, 0x0a, 0x10, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x56, 0x70, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x7f, 0x0a, 0x11, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x56,
	0x70, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x7a, 0x65, 0x72,
	0x6f, 0x70, 0x73, 0x44, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2e, 0x76, 0x70, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x31, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x7a, 0x65, 0x72, 0x6f, 0x70, 0x73, 0x44, 0x61, 0x65, 0x6d, 0x6f,
	0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x52,
	0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2a, 0x34, 0x0a, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x0c, 0x0a, 0x08, 0x4e, 0x4f, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10,
	0x00, 0x12, 0x19, 0x0a, 0x15, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x53, 0x45,
	0x52, 0x56, 0x45, 0x52, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x01, 0x2a, 0x25, 0x0a, 0x09,
	0x76, 0x70, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x4e, 0x41,
	0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43, 0x54, 0x49, 0x56,
	0x45, 0x10, 0x01, 0x32, 0xa7, 0x02, 0x0a, 0x14, 0x5a, 0x65, 0x72, 0x6f, 0x70, 0x73, 0x44, 0x61,
	0x65, 0x6d, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x59, 0x0a, 0x08,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x56, 0x70, 0x6e, 0x12, 0x25, 0x2e, 0x7a, 0x65, 0x72, 0x6f, 0x70,
	0x73, 0x44, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x56, 0x70, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x26, 0x2e, 0x7a, 0x65, 0x72, 0x6f, 0x70, 0x73, 0x44, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x73, 0x74, 0x61, 0x72, 0x74, 0x56, 0x70, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x70, 0x56,
	0x70, 0x6e, 0x12, 0x24, 0x2e, 0x7a, 0x65, 0x72, 0x6f, 0x70, 0x73, 0x44, 0x61, 0x65, 0x6d, 0x6f,
	0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x73, 0x74, 0x6f, 0x70, 0x56, 0x70,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x7a, 0x65, 0x72, 0x6f, 0x70,
	0x73, 0x44, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e,
	0x73, 0x74, 0x6f, 0x70, 0x56, 0x70, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x5c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x56, 0x70, 0x6e, 0x12, 0x26, 0x2e, 0x7a,
	0x65, 0x72, 0x6f, 0x70, 0x73, 0x44, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x56, 0x70, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x7a, 0x65, 0x72, 0x6f, 0x70, 0x73, 0x44, 0x61, 0x65,
	0x6d, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x56, 0x70, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x34, 0x5a,
	0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x65, 0x72, 0x6f,
	0x70, 0x73, 0x2d, 0x69, 0x6f, 0x2f, 0x7a, 0x63, 0x6c, 0x69, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x7a,
	0x65, 0x72, 0x6f, 0x70, 0x73, 0x44, 0x61, 0x65, 0x6d, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_zeropsDaemonProtocol_proto_rawDescOnce sync.Once
	file_zeropsDaemonProtocol_proto_rawDescData = file_zeropsDaemonProtocol_proto_rawDesc
)

func file_zeropsDaemonProtocol_proto_rawDescGZIP() []byte {
	file_zeropsDaemonProtocol_proto_rawDescOnce.Do(func() {
		file_zeropsDaemonProtocol_proto_rawDescData = protoimpl.X.CompressGZIP(file_zeropsDaemonProtocol_proto_rawDescData)
	})
	return file_zeropsDaemonProtocol_proto_rawDescData
}

var file_zeropsDaemonProtocol_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_zeropsDaemonProtocol_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_zeropsDaemonProtocol_proto_goTypes = []interface{}{
	(ErrorCode)(0),            // 0: zeropsDaemonProtocol.errorCode
	(VpnStatus)(0),            // 1: zeropsDaemonProtocol.vpnStatus
	(*Error)(nil),             // 2: zeropsDaemonProtocol.error
	(*StartVpnRequest)(nil),   // 3: zeropsDaemonProtocol.startVpnRequest
	(*StartVpnResponse)(nil),  // 4: zeropsDaemonProtocol.startVpnResponse
	(*StopVpnRequest)(nil),    // 5: zeropsDaemonProtocol.stopVpnRequest
	(*StopVpnResponse)(nil),   // 6: zeropsDaemonProtocol.stopVpnResponse
	(*StatusVpnRequest)(nil),  // 7: zeropsDaemonProtocol.statusVpnRequest
	(*StatusVpnResponse)(nil), // 8: zeropsDaemonProtocol.statusVpnResponse
}
var file_zeropsDaemonProtocol_proto_depIdxs = []int32{
	0, // 0: zeropsDaemonProtocol.error.code:type_name -> zeropsDaemonProtocol.errorCode
	2, // 1: zeropsDaemonProtocol.startVpnResponse.error:type_name -> zeropsDaemonProtocol.error
	2, // 2: zeropsDaemonProtocol.stopVpnResponse.error:type_name -> zeropsDaemonProtocol.error
	1, // 3: zeropsDaemonProtocol.statusVpnResponse.status:type_name -> zeropsDaemonProtocol.vpnStatus
	2, // 4: zeropsDaemonProtocol.statusVpnResponse.error:type_name -> zeropsDaemonProtocol.error
	3, // 5: zeropsDaemonProtocol.ZeropsDaemonProtocol.startVpn:input_type -> zeropsDaemonProtocol.startVpnRequest
	5, // 6: zeropsDaemonProtocol.ZeropsDaemonProtocol.stopVpn:input_type -> zeropsDaemonProtocol.stopVpnRequest
	7, // 7: zeropsDaemonProtocol.ZeropsDaemonProtocol.statusVpn:input_type -> zeropsDaemonProtocol.statusVpnRequest
	4, // 8: zeropsDaemonProtocol.ZeropsDaemonProtocol.startVpn:output_type -> zeropsDaemonProtocol.startVpnResponse
	6, // 9: zeropsDaemonProtocol.ZeropsDaemonProtocol.stopVpn:output_type -> zeropsDaemonProtocol.stopVpnResponse
	8, // 10: zeropsDaemonProtocol.ZeropsDaemonProtocol.statusVpn:output_type -> zeropsDaemonProtocol.statusVpnResponse
	8, // [8:11] is the sub-list for method output_type
	5, // [5:8] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_zeropsDaemonProtocol_proto_init() }
func file_zeropsDaemonProtocol_proto_init() {
	if File_zeropsDaemonProtocol_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_zeropsDaemonProtocol_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
		file_zeropsDaemonProtocol_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartVpnRequest); i {
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
		file_zeropsDaemonProtocol_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartVpnResponse); i {
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
		file_zeropsDaemonProtocol_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopVpnRequest); i {
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
		file_zeropsDaemonProtocol_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopVpnResponse); i {
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
		file_zeropsDaemonProtocol_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusVpnRequest); i {
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
		file_zeropsDaemonProtocol_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusVpnResponse); i {
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
			RawDescriptor: file_zeropsDaemonProtocol_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_zeropsDaemonProtocol_proto_goTypes,
		DependencyIndexes: file_zeropsDaemonProtocol_proto_depIdxs,
		EnumInfos:         file_zeropsDaemonProtocol_proto_enumTypes,
		MessageInfos:      file_zeropsDaemonProtocol_proto_msgTypes,
	}.Build()
	File_zeropsDaemonProtocol_proto = out.File
	file_zeropsDaemonProtocol_proto_rawDesc = nil
	file_zeropsDaemonProtocol_proto_goTypes = nil
	file_zeropsDaemonProtocol_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ZeropsDaemonProtocolClient is the client API for ZeropsDaemonProtocol service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ZeropsDaemonProtocolClient interface {
	StartVpn(ctx context.Context, in *StartVpnRequest, opts ...grpc.CallOption) (*StartVpnResponse, error)
	StopVpn(ctx context.Context, in *StopVpnRequest, opts ...grpc.CallOption) (*StopVpnResponse, error)
	StatusVpn(ctx context.Context, in *StatusVpnRequest, opts ...grpc.CallOption) (*StatusVpnResponse, error)
}

type zeropsDaemonProtocolClient struct {
	cc grpc.ClientConnInterface
}

func NewZeropsDaemonProtocolClient(cc grpc.ClientConnInterface) ZeropsDaemonProtocolClient {
	return &zeropsDaemonProtocolClient{cc}
}

func (c *zeropsDaemonProtocolClient) StartVpn(ctx context.Context, in *StartVpnRequest, opts ...grpc.CallOption) (*StartVpnResponse, error) {
	out := new(StartVpnResponse)
	err := c.cc.Invoke(ctx, "/zeropsDaemonProtocol.ZeropsDaemonProtocol/startVpn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zeropsDaemonProtocolClient) StopVpn(ctx context.Context, in *StopVpnRequest, opts ...grpc.CallOption) (*StopVpnResponse, error) {
	out := new(StopVpnResponse)
	err := c.cc.Invoke(ctx, "/zeropsDaemonProtocol.ZeropsDaemonProtocol/stopVpn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zeropsDaemonProtocolClient) StatusVpn(ctx context.Context, in *StatusVpnRequest, opts ...grpc.CallOption) (*StatusVpnResponse, error) {
	out := new(StatusVpnResponse)
	err := c.cc.Invoke(ctx, "/zeropsDaemonProtocol.ZeropsDaemonProtocol/statusVpn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ZeropsDaemonProtocolServer is the server API for ZeropsDaemonProtocol service.
type ZeropsDaemonProtocolServer interface {
	StartVpn(context.Context, *StartVpnRequest) (*StartVpnResponse, error)
	StopVpn(context.Context, *StopVpnRequest) (*StopVpnResponse, error)
	StatusVpn(context.Context, *StatusVpnRequest) (*StatusVpnResponse, error)
}

// UnimplementedZeropsDaemonProtocolServer can be embedded to have forward compatible implementations.
type UnimplementedZeropsDaemonProtocolServer struct {
}

func (*UnimplementedZeropsDaemonProtocolServer) StartVpn(context.Context, *StartVpnRequest) (*StartVpnResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartVpn not implemented")
}
func (*UnimplementedZeropsDaemonProtocolServer) StopVpn(context.Context, *StopVpnRequest) (*StopVpnResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopVpn not implemented")
}
func (*UnimplementedZeropsDaemonProtocolServer) StatusVpn(context.Context, *StatusVpnRequest) (*StatusVpnResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StatusVpn not implemented")
}

func RegisterZeropsDaemonProtocolServer(s *grpc.Server, srv ZeropsDaemonProtocolServer) {
	s.RegisterService(&_ZeropsDaemonProtocol_serviceDesc, srv)
}

func _ZeropsDaemonProtocol_StartVpn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartVpnRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZeropsDaemonProtocolServer).StartVpn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zeropsDaemonProtocol.ZeropsDaemonProtocol/StartVpn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZeropsDaemonProtocolServer).StartVpn(ctx, req.(*StartVpnRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZeropsDaemonProtocol_StopVpn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopVpnRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZeropsDaemonProtocolServer).StopVpn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zeropsDaemonProtocol.ZeropsDaemonProtocol/StopVpn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZeropsDaemonProtocolServer).StopVpn(ctx, req.(*StopVpnRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZeropsDaemonProtocol_StatusVpn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusVpnRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZeropsDaemonProtocolServer).StatusVpn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zeropsDaemonProtocol.ZeropsDaemonProtocol/StatusVpn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZeropsDaemonProtocolServer).StatusVpn(ctx, req.(*StatusVpnRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ZeropsDaemonProtocol_serviceDesc = grpc.ServiceDesc{
	ServiceName: "zeropsDaemonProtocol.ZeropsDaemonProtocol",
	HandlerType: (*ZeropsDaemonProtocolServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "startVpn",
			Handler:    _ZeropsDaemonProtocol_StartVpn_Handler,
		},
		{
			MethodName: "stopVpn",
			Handler:    _ZeropsDaemonProtocol_StopVpn_Handler,
		},
		{
			MethodName: "statusVpn",
			Handler:    _ZeropsDaemonProtocol_StatusVpn_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "zeropsDaemonProtocol.proto",
}
