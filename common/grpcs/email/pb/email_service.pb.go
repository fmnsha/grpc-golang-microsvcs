// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.3
// source: email_service.proto

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

type SendEmailReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email *EmailMessage `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *SendEmailReq) Reset() {
	*x = SendEmailReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_email_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendEmailReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendEmailReq) ProtoMessage() {}

func (x *SendEmailReq) ProtoReflect() protoreflect.Message {
	mi := &file_email_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendEmailReq.ProtoReflect.Descriptor instead.
func (*SendEmailReq) Descriptor() ([]byte, []int) {
	return file_email_service_proto_rawDescGZIP(), []int{0}
}

func (x *SendEmailReq) GetEmail() *EmailMessage {
	if x != nil {
		return x.Email
	}
	return nil
}

type SendEmailRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *SendEmailRes) Reset() {
	*x = SendEmailRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_email_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendEmailRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendEmailRes) ProtoMessage() {}

func (x *SendEmailRes) ProtoReflect() protoreflect.Message {
	mi := &file_email_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendEmailRes.ProtoReflect.Descriptor instead.
func (*SendEmailRes) Descriptor() ([]byte, []int) {
	return file_email_service_proto_rawDescGZIP(), []int{1}
}

func (x *SendEmailRes) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_email_service_proto protoreflect.FileDescriptor

var file_email_service_proto_rawDesc = []byte{
	0x0a, 0x13, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x33, 0x0a, 0x0c, 0x53, 0x65,
	0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22,
	0x28, 0x0a, 0x0c, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0x3b, 0x0a, 0x0c, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2b, 0x0a, 0x09, 0x53, 0x65, 0x6e,
	0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x0d, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_email_service_proto_rawDescOnce sync.Once
	file_email_service_proto_rawDescData = file_email_service_proto_rawDesc
)

func file_email_service_proto_rawDescGZIP() []byte {
	file_email_service_proto_rawDescOnce.Do(func() {
		file_email_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_email_service_proto_rawDescData)
	})
	return file_email_service_proto_rawDescData
}

var file_email_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_email_service_proto_goTypes = []interface{}{
	(*SendEmailReq)(nil), // 0: SendEmailReq
	(*SendEmailRes)(nil), // 1: SendEmailRes
	(*EmailMessage)(nil), // 2: EmailMessage
}
var file_email_service_proto_depIdxs = []int32{
	2, // 0: SendEmailReq.email:type_name -> EmailMessage
	0, // 1: EmailService.SendEmail:input_type -> SendEmailReq
	1, // 2: EmailService.SendEmail:output_type -> SendEmailRes
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_email_service_proto_init() }
func file_email_service_proto_init() {
	if File_email_service_proto != nil {
		return
	}
	file_email_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_email_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendEmailReq); i {
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
		file_email_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendEmailRes); i {
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
			RawDescriptor: file_email_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_email_service_proto_goTypes,
		DependencyIndexes: file_email_service_proto_depIdxs,
		MessageInfos:      file_email_service_proto_msgTypes,
	}.Build()
	File_email_service_proto = out.File
	file_email_service_proto_rawDesc = nil
	file_email_service_proto_goTypes = nil
	file_email_service_proto_depIdxs = nil
}
