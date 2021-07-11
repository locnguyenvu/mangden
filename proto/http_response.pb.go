// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.14.0
// source: http_response.proto

package proto

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

type Success struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  string `protobuf:"bytes,1,opt,name=Status,proto3" json:"Status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
}

func (x *Success) Reset() {
	*x = Success{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Success) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Success) ProtoMessage() {}

func (x *Success) ProtoReflect() protoreflect.Message {
	mi := &file_http_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Success.ProtoReflect.Descriptor instead.
func (*Success) Descriptor() ([]byte, []int) {
	return file_http_response_proto_rawDescGZIP(), []int{0}
}

func (x *Success) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Success) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ConfigDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Value     string `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
	CreatedAt string `protobuf:"bytes,3,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt string `protobuf:"bytes,4,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
}

func (x *ConfigDetail) Reset() {
	*x = ConfigDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_http_response_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigDetail) ProtoMessage() {}

func (x *ConfigDetail) ProtoReflect() protoreflect.Message {
	mi := &file_http_response_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigDetail.ProtoReflect.Descriptor instead.
func (*ConfigDetail) Descriptor() ([]byte, []int) {
	return file_http_response_proto_rawDescGZIP(), []int{1}
}

func (x *ConfigDetail) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ConfigDetail) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *ConfigDetail) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *ConfigDetail) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

var File_http_response_proto protoreflect.FileDescriptor

var file_http_response_proto_rawDesc = []byte{
	0x0a, 0x13, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3b, 0x0a, 0x07,
	0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x74, 0x0a, 0x0c, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42,
	0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x6f,
	0x63, 0x6e, 0x67, 0x75, 0x79, 0x65, 0x6e, 0x76, 0x75, 0x2f, 0x6d, 0x61, 0x6e, 0x67, 0x64, 0x65,
	0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_http_response_proto_rawDescOnce sync.Once
	file_http_response_proto_rawDescData = file_http_response_proto_rawDesc
)

func file_http_response_proto_rawDescGZIP() []byte {
	file_http_response_proto_rawDescOnce.Do(func() {
		file_http_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_http_response_proto_rawDescData)
	})
	return file_http_response_proto_rawDescData
}

var file_http_response_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_http_response_proto_goTypes = []interface{}{
	(*Success)(nil),      // 0: proto.Success
	(*ConfigDetail)(nil), // 1: proto.ConfigDetail
}
var file_http_response_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_http_response_proto_init() }
func file_http_response_proto_init() {
	if File_http_response_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_http_response_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Success); i {
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
		file_http_response_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigDetail); i {
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
			RawDescriptor: file_http_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_http_response_proto_goTypes,
		DependencyIndexes: file_http_response_proto_depIdxs,
		MessageInfos:      file_http_response_proto_msgTypes,
	}.Build()
	File_http_response_proto = out.File
	file_http_response_proto_rawDesc = nil
	file_http_response_proto_goTypes = nil
	file_http_response_proto_depIdxs = nil
}
