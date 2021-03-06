// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: plugins/plugins.proto

package plugins

import (
	proto "github.com/golang/protobuf/proto"
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

// Configuration of Spinnaker's plugins
type Plugins struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bakery         *Bakery         `protobuf:"bytes,1,opt,name=bakery,proto3" json:"bakery,omitempty"`
	DeliveryConfig *DeliveryConfig `protobuf:"bytes,2,opt,name=deliveryConfig,proto3" json:"deliveryConfig,omitempty"`
	Ec2            *EC2            `protobuf:"bytes,3,opt,name=ec2,proto3" json:"ec2,omitempty"`
	K8S            *Kubernetes     `protobuf:"bytes,4,opt,name=k8s,proto3" json:"k8s,omitempty"`
	Titus          *Titus          `protobuf:"bytes,5,opt,name=titus,proto3" json:"titus,omitempty"`
}

func (x *Plugins) Reset() {
	*x = Plugins{}
	if protoimpl.UnsafeEnabled {
		mi := &file_plugins_plugins_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Plugins) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Plugins) ProtoMessage() {}

func (x *Plugins) ProtoReflect() protoreflect.Message {
	mi := &file_plugins_plugins_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Plugins.ProtoReflect.Descriptor instead.
func (*Plugins) Descriptor() ([]byte, []int) {
	return file_plugins_plugins_proto_rawDescGZIP(), []int{0}
}

func (x *Plugins) GetBakery() *Bakery {
	if x != nil {
		return x.Bakery
	}
	return nil
}

func (x *Plugins) GetDeliveryConfig() *DeliveryConfig {
	if x != nil {
		return x.DeliveryConfig
	}
	return nil
}

func (x *Plugins) GetEc2() *EC2 {
	if x != nil {
		return x.Ec2
	}
	return nil
}

func (x *Plugins) GetK8S() *Kubernetes {
	if x != nil {
		return x.K8S
	}
	return nil
}

func (x *Plugins) GetTitus() *Titus {
	if x != nil {
		return x.Titus
	}
	return nil
}

var File_plugins_plugins_proto protoreflect.FileDescriptor

var file_plugins_plugins_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70,
	0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x1a, 0x14, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2f,
	0x62, 0x61, 0x6b, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x73, 0x2f, 0x65, 0x63, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x6b, 0x38, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x13, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x74, 0x69, 0x74, 0x75, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfe, 0x01, 0x0a, 0x07, 0x50, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x73, 0x12, 0x2d, 0x0a, 0x06, 0x62, 0x61, 0x6b, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x73, 0x2e, 0x42, 0x61, 0x6b, 0x65, 0x72, 0x79, 0x52, 0x06, 0x62, 0x61, 0x6b, 0x65, 0x72,
	0x79, 0x12, 0x45, 0x0a, 0x0e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65,
	0x72, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65,
	0x72, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x24, 0x0a, 0x03, 0x65, 0x63, 0x32, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x45, 0x43, 0x32, 0x52, 0x03, 0x65, 0x63, 0x32, 0x12, 0x2b,
	0x0a, 0x03, 0x6b, 0x38, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x4b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x52, 0x03, 0x6b, 0x38, 0x73, 0x12, 0x2a, 0x0a, 0x05, 0x74,
	0x69, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x54, 0x69, 0x74, 0x75, 0x73,
	0x52, 0x05, 0x74, 0x69, 0x74, 0x75, 0x73, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72, 0x2f,
	0x6b, 0x6c, 0x65, 0x61, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_plugins_plugins_proto_rawDescOnce sync.Once
	file_plugins_plugins_proto_rawDescData = file_plugins_plugins_proto_rawDesc
)

func file_plugins_plugins_proto_rawDescGZIP() []byte {
	file_plugins_plugins_proto_rawDescOnce.Do(func() {
		file_plugins_plugins_proto_rawDescData = protoimpl.X.CompressGZIP(file_plugins_plugins_proto_rawDescData)
	})
	return file_plugins_plugins_proto_rawDescData
}

var file_plugins_plugins_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_plugins_plugins_proto_goTypes = []interface{}{
	(*Plugins)(nil),        // 0: proto.plugins.Plugins
	(*Bakery)(nil),         // 1: proto.plugins.Bakery
	(*DeliveryConfig)(nil), // 2: proto.plugins.DeliveryConfig
	(*EC2)(nil),            // 3: proto.plugins.EC2
	(*Kubernetes)(nil),     // 4: proto.plugins.Kubernetes
	(*Titus)(nil),          // 5: proto.plugins.Titus
}
var file_plugins_plugins_proto_depIdxs = []int32{
	1, // 0: proto.plugins.Plugins.bakery:type_name -> proto.plugins.Bakery
	2, // 1: proto.plugins.Plugins.deliveryConfig:type_name -> proto.plugins.DeliveryConfig
	3, // 2: proto.plugins.Plugins.ec2:type_name -> proto.plugins.EC2
	4, // 3: proto.plugins.Plugins.k8s:type_name -> proto.plugins.Kubernetes
	5, // 4: proto.plugins.Plugins.titus:type_name -> proto.plugins.Titus
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_plugins_plugins_proto_init() }
func file_plugins_plugins_proto_init() {
	if File_plugins_plugins_proto != nil {
		return
	}
	file_plugins_bakery_proto_init()
	file_plugins_delivery_config_proto_init()
	file_plugins_ec2_proto_init()
	file_plugins_k8s_proto_init()
	file_plugins_titus_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_plugins_plugins_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Plugins); i {
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
			RawDescriptor: file_plugins_plugins_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_plugins_plugins_proto_goTypes,
		DependencyIndexes: file_plugins_plugins_proto_depIdxs,
		MessageInfos:      file_plugins_plugins_proto_msgTypes,
	}.Build()
	File_plugins_plugins_proto = out.File
	file_plugins_plugins_proto_rawDesc = nil
	file_plugins_plugins_proto_goTypes = nil
	file_plugins_plugins_proto_depIdxs = nil
}
