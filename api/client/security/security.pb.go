// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: security/security.proto

package security

import (
	proto "github.com/golang/protobuf/proto"
	authn "github.com/spinnaker/kleat/api/client/security/authn"
	authz "github.com/spinnaker/kleat/api/client/security/authz"
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

// Configuration for security settings.
type Security struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Configuration for the API server's addressable URL and CORS policies.
	ApiSecurity *ApiSecurity `protobuf:"bytes,1,opt,name=apiSecurity,proto3" json:"apiSecurity,omitempty"`
	// Configuration for the UI server's addressable URL.
	UiSecurity *UiSecurity `protobuf:"bytes,2,opt,name=uiSecurity,proto3" json:"uiSecurity,omitempty"`
	// Configuration of how users authenticate against Spinnaker.
	Authn *authn.Authentication `protobuf:"bytes,3,opt,name=authn,proto3" json:"authn,omitempty"`
	// Configuration for what resources users of Spinnaker can read and modify.
	Authz *authz.Authorization `protobuf:"bytes,4,opt,name=authz,proto3" json:"authz,omitempty"`
}

func (x *Security) Reset() {
	*x = Security{}
	if protoimpl.UnsafeEnabled {
		mi := &file_security_security_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Security) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Security) ProtoMessage() {}

func (x *Security) ProtoReflect() protoreflect.Message {
	mi := &file_security_security_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Security.ProtoReflect.Descriptor instead.
func (*Security) Descriptor() ([]byte, []int) {
	return file_security_security_proto_rawDescGZIP(), []int{0}
}

func (x *Security) GetApiSecurity() *ApiSecurity {
	if x != nil {
		return x.ApiSecurity
	}
	return nil
}

func (x *Security) GetUiSecurity() *UiSecurity {
	if x != nil {
		return x.UiSecurity
	}
	return nil
}

func (x *Security) GetAuthn() *authn.Authentication {
	if x != nil {
		return x.Authn
	}
	return nil
}

func (x *Security) GetAuthz() *authz.Authorization {
	if x != nil {
		return x.Authz
	}
	return nil
}

var File_security_security_proto protoreflect.FileDescriptor

var file_security_security_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72,
	0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x1a, 0x1a, 0x73, 0x65, 0x63, 0x75, 0x72,
	0x69, 0x74, 0x79, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x7a, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x12, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2f, 0x73, 0x73, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfc, 0x01, 0x0a, 0x08, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x12, 0x3d, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x41, 0x70, 0x69, 0x53, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x52, 0x0b, 0x61, 0x70, 0x69, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x12, 0x3a, 0x0a, 0x0a, 0x75, 0x69, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x55, 0x69, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x52, 0x0a, 0x75, 0x69, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x12, 0x3a, 0x0a,
	0x05, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x6e, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x05, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x12, 0x39, 0x0a, 0x05, 0x61, 0x75, 0x74,
	0x68, 0x7a, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2e,
	0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x61,
	0x75, 0x74, 0x68, 0x7a, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72, 0x2f, 0x6b, 0x6c, 0x65,
	0x61, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x73, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_security_security_proto_rawDescOnce sync.Once
	file_security_security_proto_rawDescData = file_security_security_proto_rawDesc
)

func file_security_security_proto_rawDescGZIP() []byte {
	file_security_security_proto_rawDescOnce.Do(func() {
		file_security_security_proto_rawDescData = protoimpl.X.CompressGZIP(file_security_security_proto_rawDescData)
	})
	return file_security_security_proto_rawDescData
}

var file_security_security_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_security_security_proto_goTypes = []interface{}{
	(*Security)(nil),             // 0: proto.security.Security
	(*ApiSecurity)(nil),          // 1: proto.security.ApiSecurity
	(*UiSecurity)(nil),           // 2: proto.security.UiSecurity
	(*authn.Authentication)(nil), // 3: proto.security.authn.Authentication
	(*authz.Authorization)(nil),  // 4: proto.security.authz.Authorization
}
var file_security_security_proto_depIdxs = []int32{
	1, // 0: proto.security.Security.apiSecurity:type_name -> proto.security.ApiSecurity
	2, // 1: proto.security.Security.uiSecurity:type_name -> proto.security.UiSecurity
	3, // 2: proto.security.Security.authn:type_name -> proto.security.authn.Authentication
	4, // 3: proto.security.Security.authz:type_name -> proto.security.authz.Authorization
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_security_security_proto_init() }
func file_security_security_proto_init() {
	if File_security_security_proto != nil {
		return
	}
	file_security_ssl_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_security_security_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Security); i {
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
			RawDescriptor: file_security_security_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_security_security_proto_goTypes,
		DependencyIndexes: file_security_security_proto_depIdxs,
		MessageInfos:      file_security_security_proto_msgTypes,
	}.Build()
	File_security_security_proto = out.File
	file_security_security_proto_rawDesc = nil
	file_security_security_proto_goTypes = nil
	file_security_security_proto_depIdxs = nil
}
