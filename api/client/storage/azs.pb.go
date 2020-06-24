// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: storage/azs.proto

package storage

import (
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// Configuration for an Azure Storage persistent store.
type Azs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether this persistent store is enabled.
	Enabled *wrappers.BoolValue `protobuf:"bytes,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// The name of an Azure Storage Account.
	StorageAccountName string `protobuf:"bytes,2,opt,name=storageAccountName,proto3" json:"storageAccountName,omitempty"`
	// The key to access the Azure Storage Account.
	StorageAccountKey string `protobuf:"bytes,3,opt,name=storageAccountKey,proto3" json:"storageAccountKey,omitempty"`
	// The container name in the chosen storage account to place Spinnaker's
	// persistent data. Defaults to 'spinnaker' if unspecified.
	StorageContainerName string `protobuf:"bytes,4,opt,name=storageContainerName,proto3" json:"storageContainerName,omitempty"`
}

func (x *Azs) Reset() {
	*x = Azs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storage_azs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Azs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Azs) ProtoMessage() {}

func (x *Azs) ProtoReflect() protoreflect.Message {
	mi := &file_storage_azs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Azs.ProtoReflect.Descriptor instead.
func (*Azs) Descriptor() ([]byte, []int) {
	return file_storage_azs_proto_rawDescGZIP(), []int{0}
}

func (x *Azs) GetEnabled() *wrappers.BoolValue {
	if x != nil {
		return x.Enabled
	}
	return nil
}

func (x *Azs) GetStorageAccountName() string {
	if x != nil {
		return x.StorageAccountName
	}
	return ""
}

func (x *Azs) GetStorageAccountKey() string {
	if x != nil {
		return x.StorageAccountKey
	}
	return ""
}

func (x *Azs) GetStorageContainerName() string {
	if x != nil {
		return x.StorageContainerName
	}
	return ""
}

var File_storage_azs_proto protoreflect.FileDescriptor

var file_storage_azs_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x61, 0x7a, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xcd, 0x01, 0x0a, 0x03, 0x41, 0x7a, 0x73, 0x12, 0x34, 0x0a, 0x07, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f,
	0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x12, 0x2e, 0x0a, 0x12, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x2c, 0x0a, 0x11, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x4b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x32,
	0x0a, 0x14, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72, 0x2f, 0x6b, 0x6c, 0x65, 0x61, 0x74,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_storage_azs_proto_rawDescOnce sync.Once
	file_storage_azs_proto_rawDescData = file_storage_azs_proto_rawDesc
)

func file_storage_azs_proto_rawDescGZIP() []byte {
	file_storage_azs_proto_rawDescOnce.Do(func() {
		file_storage_azs_proto_rawDescData = protoimpl.X.CompressGZIP(file_storage_azs_proto_rawDescData)
	})
	return file_storage_azs_proto_rawDescData
}

var file_storage_azs_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_storage_azs_proto_goTypes = []interface{}{
	(*Azs)(nil),                // 0: proto.storage.Azs
	(*wrappers.BoolValue)(nil), // 1: google.protobuf.BoolValue
}
var file_storage_azs_proto_depIdxs = []int32{
	1, // 0: proto.storage.Azs.enabled:type_name -> google.protobuf.BoolValue
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_storage_azs_proto_init() }
func file_storage_azs_proto_init() {
	if File_storage_azs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_storage_azs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Azs); i {
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
			RawDescriptor: file_storage_azs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_storage_azs_proto_goTypes,
		DependencyIndexes: file_storage_azs_proto_depIdxs,
		MessageInfos:      file_storage_azs_proto_msgTypes,
	}.Build()
	File_storage_azs_proto = out.File
	file_storage_azs_proto_rawDesc = nil
	file_storage_azs_proto_goTypes = nil
	file_storage_azs_proto_depIdxs = nil
}
