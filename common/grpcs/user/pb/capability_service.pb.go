// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.3
// source: capability_service.proto

package pb

import (
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

type CreateCapReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Capability *Capability `protobuf:"bytes,1,opt,name=capability,proto3" json:"capability,omitempty"`
}

func (x *CreateCapReq) Reset() {
	*x = CreateCapReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_capability_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCapReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCapReq) ProtoMessage() {}

func (x *CreateCapReq) ProtoReflect() protoreflect.Message {
	mi := &file_capability_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCapReq.ProtoReflect.Descriptor instead.
func (*CreateCapReq) Descriptor() ([]byte, []int) {
	return file_capability_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCapReq) GetCapability() *Capability {
	if x != nil {
		return x.Capability
	}
	return nil
}

type CreateCapRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Capability *Capability `protobuf:"bytes,1,opt,name=capability,proto3" json:"capability,omitempty"`
}

func (x *CreateCapRes) Reset() {
	*x = CreateCapRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_capability_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCapRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCapRes) ProtoMessage() {}

func (x *CreateCapRes) ProtoReflect() protoreflect.Message {
	mi := &file_capability_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCapRes.ProtoReflect.Descriptor instead.
func (*CreateCapRes) Descriptor() ([]byte, []int) {
	return file_capability_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCapRes) GetCapability() *Capability {
	if x != nil {
		return x.Capability
	}
	return nil
}

type ReadCapReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ReadCapReq) Reset() {
	*x = ReadCapReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_capability_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadCapReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadCapReq) ProtoMessage() {}

func (x *ReadCapReq) ProtoReflect() protoreflect.Message {
	mi := &file_capability_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadCapReq.ProtoReflect.Descriptor instead.
func (*ReadCapReq) Descriptor() ([]byte, []int) {
	return file_capability_service_proto_rawDescGZIP(), []int{2}
}

func (x *ReadCapReq) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ReadCapRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Capability *Capability `protobuf:"bytes,1,opt,name=capability,proto3" json:"capability,omitempty"`
}

func (x *ReadCapRes) Reset() {
	*x = ReadCapRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_capability_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadCapRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadCapRes) ProtoMessage() {}

func (x *ReadCapRes) ProtoReflect() protoreflect.Message {
	mi := &file_capability_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadCapRes.ProtoReflect.Descriptor instead.
func (*ReadCapRes) Descriptor() ([]byte, []int) {
	return file_capability_service_proto_rawDescGZIP(), []int{3}
}

func (x *ReadCapRes) GetCapability() *Capability {
	if x != nil {
		return x.Capability
	}
	return nil
}

type UpdateCapReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Capability *Capability `protobuf:"bytes,1,opt,name=capability,proto3" json:"capability,omitempty"`
}

func (x *UpdateCapReq) Reset() {
	*x = UpdateCapReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_capability_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCapReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCapReq) ProtoMessage() {}

func (x *UpdateCapReq) ProtoReflect() protoreflect.Message {
	mi := &file_capability_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCapReq.ProtoReflect.Descriptor instead.
func (*UpdateCapReq) Descriptor() ([]byte, []int) {
	return file_capability_service_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateCapReq) GetCapability() *Capability {
	if x != nil {
		return x.Capability
	}
	return nil
}

type UpdateCapRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Capability *Capability `protobuf:"bytes,1,opt,name=capability,proto3" json:"capability,omitempty"`
}

func (x *UpdateCapRes) Reset() {
	*x = UpdateCapRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_capability_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCapRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCapRes) ProtoMessage() {}

func (x *UpdateCapRes) ProtoReflect() protoreflect.Message {
	mi := &file_capability_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCapRes.ProtoReflect.Descriptor instead.
func (*UpdateCapRes) Descriptor() ([]byte, []int) {
	return file_capability_service_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateCapRes) GetCapability() *Capability {
	if x != nil {
		return x.Capability
	}
	return nil
}

type DeleteCapReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteCapReq) Reset() {
	*x = DeleteCapReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_capability_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCapReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCapReq) ProtoMessage() {}

func (x *DeleteCapReq) ProtoReflect() protoreflect.Message {
	mi := &file_capability_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCapReq.ProtoReflect.Descriptor instead.
func (*DeleteCapReq) Descriptor() ([]byte, []int) {
	return file_capability_service_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteCapReq) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteCapRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *DeleteCapRes) Reset() {
	*x = DeleteCapRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_capability_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCapRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCapRes) ProtoMessage() {}

func (x *DeleteCapRes) ProtoReflect() protoreflect.Message {
	mi := &file_capability_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCapRes.ProtoReflect.Descriptor instead.
func (*DeleteCapRes) Descriptor() ([]byte, []int) {
	return file_capability_service_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteCapRes) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type ListCapReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListCapReq) Reset() {
	*x = ListCapReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_capability_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCapReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCapReq) ProtoMessage() {}

func (x *ListCapReq) ProtoReflect() protoreflect.Message {
	mi := &file_capability_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCapReq.ProtoReflect.Descriptor instead.
func (*ListCapReq) Descriptor() ([]byte, []int) {
	return file_capability_service_proto_rawDescGZIP(), []int{8}
}

type ListCapRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Capability *Capability `protobuf:"bytes,1,opt,name=capability,proto3" json:"capability,omitempty"`
}

func (x *ListCapRes) Reset() {
	*x = ListCapRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_capability_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCapRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCapRes) ProtoMessage() {}

func (x *ListCapRes) ProtoReflect() protoreflect.Message {
	mi := &file_capability_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCapRes.ProtoReflect.Descriptor instead.
func (*ListCapRes) Descriptor() ([]byte, []int) {
	return file_capability_service_proto_rawDescGZIP(), []int{9}
}

func (x *ListCapRes) GetCapability() *Capability {
	if x != nil {
		return x.Capability
	}
	return nil
}

type AddRoleCapabilitiesReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoleCapability *RoleCapability `protobuf:"bytes,1,opt,name=roleCapability,proto3" json:"roleCapability,omitempty"`
}

func (x *AddRoleCapabilitiesReq) Reset() {
	*x = AddRoleCapabilitiesReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_capability_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRoleCapabilitiesReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRoleCapabilitiesReq) ProtoMessage() {}

func (x *AddRoleCapabilitiesReq) ProtoReflect() protoreflect.Message {
	mi := &file_capability_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRoleCapabilitiesReq.ProtoReflect.Descriptor instead.
func (*AddRoleCapabilitiesReq) Descriptor() ([]byte, []int) {
	return file_capability_service_proto_rawDescGZIP(), []int{10}
}

func (x *AddRoleCapabilitiesReq) GetRoleCapability() *RoleCapability {
	if x != nil {
		return x.RoleCapability
	}
	return nil
}

type AddRoleCapabilitiesRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *AddRoleCapabilitiesRes) Reset() {
	*x = AddRoleCapabilitiesRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_capability_service_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRoleCapabilitiesRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRoleCapabilitiesRes) ProtoMessage() {}

func (x *AddRoleCapabilitiesRes) ProtoReflect() protoreflect.Message {
	mi := &file_capability_service_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRoleCapabilitiesRes.ProtoReflect.Descriptor instead.
func (*AddRoleCapabilitiesRes) Descriptor() ([]byte, []int) {
	return file_capability_service_proto_rawDescGZIP(), []int{11}
}

func (x *AddRoleCapabilitiesRes) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_capability_service_proto protoreflect.FileDescriptor

var file_capability_service_proto_rawDesc = []byte{
	0x0a, 0x18, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x63, 0x61, 0x70, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3b, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61,
	0x70, 0x52, 0x65, 0x71, 0x12, 0x2b, 0x0a, 0x0a, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x43, 0x61, 0x70, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x0a, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x22, 0x3b, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x70, 0x52, 0x65,
	0x73, 0x12, 0x2b, 0x0a, 0x0a, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x52, 0x0a, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x22, 0x1c,
	0x0a, 0x0a, 0x52, 0x65, 0x61, 0x64, 0x43, 0x61, 0x70, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x39, 0x0a, 0x0a,
	0x52, 0x65, 0x61, 0x64, 0x43, 0x61, 0x70, 0x52, 0x65, 0x73, 0x12, 0x2b, 0x0a, 0x0a, 0x63, 0x61,
	0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x0a, 0x63, 0x61, 0x70,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x22, 0x3b, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x43, 0x61, 0x70, 0x52, 0x65, 0x71, 0x12, 0x2b, 0x0a, 0x0a, 0x63, 0x61, 0x70, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x43, 0x61,
	0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x0a, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x22, 0x3b, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61,
	0x70, 0x52, 0x65, 0x73, 0x12, 0x2b, 0x0a, 0x0a, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x43, 0x61, 0x70, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x0a, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x22, 0x1e, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x70, 0x52, 0x65,
	0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x28, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x70, 0x52, 0x65,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x0c, 0x0a, 0x0a, 0x4c,
	0x69, 0x73, 0x74, 0x43, 0x61, 0x70, 0x52, 0x65, 0x71, 0x22, 0x39, 0x0a, 0x0a, 0x4c, 0x69, 0x73,
	0x74, 0x43, 0x61, 0x70, 0x52, 0x65, 0x73, 0x12, 0x2b, 0x0a, 0x0a, 0x63, 0x61, 0x70, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x43, 0x61,
	0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x0a, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x22, 0x51, 0x0a, 0x16, 0x41, 0x64, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x43,
	0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x12, 0x37,
	0x0a, 0x0e, 0x72, 0x6f, 0x6c, 0x65, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x43, 0x61, 0x70,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x0e, 0x72, 0x6f, 0x6c, 0x65, 0x43, 0x61, 0x70,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x22, 0x32, 0x0a, 0x16, 0x41, 0x64, 0x64, 0x52, 0x6f,
	0x6c, 0x65, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0xb3, 0x02, 0x0a, 0x11,
	0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x2b, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x70, 0x12, 0x0d,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x70, 0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x70, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x25,
	0x0a, 0x07, 0x52, 0x65, 0x61, 0x64, 0x43, 0x61, 0x70, 0x12, 0x0b, 0x2e, 0x52, 0x65, 0x61, 0x64,
	0x43, 0x61, 0x70, 0x52, 0x65, 0x71, 0x1a, 0x0b, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x43, 0x61, 0x70,
	0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x61, 0x70, 0x12, 0x0d, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x70, 0x52, 0x65,
	0x71, 0x1a, 0x0d, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x70, 0x52, 0x65, 0x73,
	0x22, 0x00, 0x12, 0x2b, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x70, 0x12,
	0x0d, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x70, 0x52, 0x65, 0x71, 0x1a, 0x0d,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x70, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12,
	0x27, 0x0a, 0x07, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x70, 0x12, 0x0b, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x43, 0x61, 0x70, 0x52, 0x65, 0x71, 0x1a, 0x0b, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61,
	0x70, 0x52, 0x65, 0x73, 0x22, 0x00, 0x30, 0x01, 0x12, 0x47, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x52,
	0x6f, 0x6c, 0x65, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x17, 0x2e,
	0x41, 0x64, 0x64, 0x52, 0x6f, 0x6c, 0x65, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x6f, 0x6c, 0x65,
	0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x22,
	0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_capability_service_proto_rawDescOnce sync.Once
	file_capability_service_proto_rawDescData = file_capability_service_proto_rawDesc
)

func file_capability_service_proto_rawDescGZIP() []byte {
	file_capability_service_proto_rawDescOnce.Do(func() {
		file_capability_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_capability_service_proto_rawDescData)
	})
	return file_capability_service_proto_rawDescData
}

var file_capability_service_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_capability_service_proto_goTypes = []interface{}{
	(*CreateCapReq)(nil),           // 0: CreateCapReq
	(*CreateCapRes)(nil),           // 1: CreateCapRes
	(*ReadCapReq)(nil),             // 2: ReadCapReq
	(*ReadCapRes)(nil),             // 3: ReadCapRes
	(*UpdateCapReq)(nil),           // 4: UpdateCapReq
	(*UpdateCapRes)(nil),           // 5: UpdateCapRes
	(*DeleteCapReq)(nil),           // 6: DeleteCapReq
	(*DeleteCapRes)(nil),           // 7: DeleteCapRes
	(*ListCapReq)(nil),             // 8: ListCapReq
	(*ListCapRes)(nil),             // 9: ListCapRes
	(*AddRoleCapabilitiesReq)(nil), // 10: AddRoleCapabilitiesReq
	(*AddRoleCapabilitiesRes)(nil), // 11: AddRoleCapabilitiesRes
	(*Capability)(nil),             // 12: Capability
	(*RoleCapability)(nil),         // 13: RoleCapability
}
var file_capability_service_proto_depIdxs = []int32{
	12, // 0: CreateCapReq.capability:type_name -> Capability
	12, // 1: CreateCapRes.capability:type_name -> Capability
	12, // 2: ReadCapRes.capability:type_name -> Capability
	12, // 3: UpdateCapReq.capability:type_name -> Capability
	12, // 4: UpdateCapRes.capability:type_name -> Capability
	12, // 5: ListCapRes.capability:type_name -> Capability
	13, // 6: AddRoleCapabilitiesReq.roleCapability:type_name -> RoleCapability
	0,  // 7: CapabilityService.CreateCap:input_type -> CreateCapReq
	2,  // 8: CapabilityService.ReadCap:input_type -> ReadCapReq
	4,  // 9: CapabilityService.UpdateCap:input_type -> UpdateCapReq
	6,  // 10: CapabilityService.DeleteCap:input_type -> DeleteCapReq
	8,  // 11: CapabilityService.ListCap:input_type -> ListCapReq
	10, // 12: CapabilityService.AddRoleCapability:input_type -> AddRoleCapabilitiesReq
	1,  // 13: CapabilityService.CreateCap:output_type -> CreateCapRes
	3,  // 14: CapabilityService.ReadCap:output_type -> ReadCapRes
	5,  // 15: CapabilityService.UpdateCap:output_type -> UpdateCapRes
	7,  // 16: CapabilityService.DeleteCap:output_type -> DeleteCapRes
	9,  // 17: CapabilityService.ListCap:output_type -> ListCapRes
	11, // 18: CapabilityService.AddRoleCapability:output_type -> AddRoleCapabilitiesRes
	13, // [13:19] is the sub-list for method output_type
	7,  // [7:13] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_capability_service_proto_init() }
func file_capability_service_proto_init() {
	if File_capability_service_proto != nil {
		return
	}
	file_capability_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_capability_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCapReq); i {
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
		file_capability_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCapRes); i {
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
		file_capability_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadCapReq); i {
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
		file_capability_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadCapRes); i {
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
		file_capability_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCapReq); i {
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
		file_capability_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCapRes); i {
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
		file_capability_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCapReq); i {
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
		file_capability_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCapRes); i {
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
		file_capability_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCapReq); i {
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
		file_capability_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCapRes); i {
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
		file_capability_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRoleCapabilitiesReq); i {
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
		file_capability_service_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRoleCapabilitiesRes); i {
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
			RawDescriptor: file_capability_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_capability_service_proto_goTypes,
		DependencyIndexes: file_capability_service_proto_depIdxs,
		MessageInfos:      file_capability_service_proto_msgTypes,
	}.Build()
	File_capability_service_proto = out.File
	file_capability_service_proto_rawDesc = nil
	file_capability_service_proto_goTypes = nil
	file_capability_service_proto_depIdxs = nil
}
