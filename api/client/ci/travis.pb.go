// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: ci/travis.proto

package ci

import (
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	client "github.com/spinnaker/kleat/api/client"
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

// Configuration for Travis.
type Travis struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether Travis is enabled.
	Enabled *wrappers.BoolValue `protobuf:"bytes,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// The list of configured Travis accounts.
	Masters []*TravisAccount `protobuf:"bytes,2,rep,name=masters,proto3" json:"masters,omitempty"`
}

func (x *Travis) Reset() {
	*x = Travis{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ci_travis_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Travis) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Travis) ProtoMessage() {}

func (x *Travis) ProtoReflect() protoreflect.Message {
	mi := &file_ci_travis_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Travis.ProtoReflect.Descriptor instead.
func (*Travis) Descriptor() ([]byte, []int) {
	return file_ci_travis_proto_rawDescGZIP(), []int{0}
}

func (x *Travis) GetEnabled() *wrappers.BoolValue {
	if x != nil {
		return x.Enabled
	}
	return nil
}

func (x *Travis) GetMasters() []*TravisAccount {
	if x != nil {
		return x.Masters
	}
	return nil
}

// Configuration for a Travis account.
type TravisAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the account.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// (Required) The address of the Travis API (https://api.travis-ci.org).
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// (Required) The base URL to the Travis UI (https://travis-ci.org).
	BaseUrl string `protobuf:"bytes,3,opt,name=baseUrl,proto3" json:"baseUrl,omitempty"`
	// The GitHub token with which to authenticate against Travis.
	GithubToken string `protobuf:"bytes,4,opt,name=githubToken,proto3" json:"githubToken,omitempty"`
	// Fiat permissions configuration. A user must have at least one of the READ
	// roles in order to view this build account or use it as a trigger source.
	// A user must have at least one of the WRITE roles in order to run jobs on
	// this build account.
	Permissions *client.Permissions `protobuf:"bytes,6,opt,name=permissions,proto3" json:"permissions,omitempty"`
	// Defines how many jobs the Travis integration should retrieve per polling
	// cycle. Defaults to 100.
	NumberOfJobs string `protobuf:"bytes,7,opt,name=numberOfJobs,proto3" json:"numberOfJobs,omitempty"`
	// Defines how many builds Igor should return when querying for builds for a
	// specific repo. This affects for instance how many builds will be displayed
	// in the drop down when starting a manual execution of a pipeline. If set too
	// high, the Travis API might return an error for jobs that writes a lot of logs,
	// which is why the default setting is a bit conservative. Defaults to 10.
	BuildResultLimit string `protobuf:"bytes,8,opt,name=buildResultLimit,proto3" json:"buildResultLimit,omitempty"`
}

func (x *TravisAccount) Reset() {
	*x = TravisAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ci_travis_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TravisAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TravisAccount) ProtoMessage() {}

func (x *TravisAccount) ProtoReflect() protoreflect.Message {
	mi := &file_ci_travis_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TravisAccount.ProtoReflect.Descriptor instead.
func (*TravisAccount) Descriptor() ([]byte, []int) {
	return file_ci_travis_proto_rawDescGZIP(), []int{1}
}

func (x *TravisAccount) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TravisAccount) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *TravisAccount) GetBaseUrl() string {
	if x != nil {
		return x.BaseUrl
	}
	return ""
}

func (x *TravisAccount) GetGithubToken() string {
	if x != nil {
		return x.GithubToken
	}
	return ""
}

func (x *TravisAccount) GetPermissions() *client.Permissions {
	if x != nil {
		return x.Permissions
	}
	return nil
}

func (x *TravisAccount) GetNumberOfJobs() string {
	if x != nil {
		return x.NumberOfJobs
	}
	return ""
}

func (x *TravisAccount) GetBuildResultLimit() string {
	if x != nil {
		return x.BuildResultLimit
	}
	return ""
}

var File_ci_travis_proto protoreflect.FileDescriptor

var file_ci_travis_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x69, 0x2f, 0x74, 0x72, 0x61, 0x76, 0x69, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x69, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x71,
	0x0a, 0x06, 0x54, 0x72, 0x61, 0x76, 0x69, 0x73, 0x12, 0x34, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x31,
	0x0a, 0x07, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x69, 0x2e, 0x54, 0x72, 0x61, 0x76, 0x69,
	0x73, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x07, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x73, 0x22, 0xff, 0x01, 0x0a, 0x0d, 0x54, 0x72, 0x61, 0x76, 0x69, 0x73, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x73, 0x65, 0x55, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x62, 0x61, 0x73, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x34, 0x0a,
	0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x4a,
	0x6f, 0x62, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x4f, 0x66, 0x4a, 0x6f, 0x62, 0x73, 0x12, 0x2a, 0x0a, 0x10, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x10, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72, 0x2f, 0x6b, 0x6c, 0x65, 0x61,
	0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x69, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ci_travis_proto_rawDescOnce sync.Once
	file_ci_travis_proto_rawDescData = file_ci_travis_proto_rawDesc
)

func file_ci_travis_proto_rawDescGZIP() []byte {
	file_ci_travis_proto_rawDescOnce.Do(func() {
		file_ci_travis_proto_rawDescData = protoimpl.X.CompressGZIP(file_ci_travis_proto_rawDescData)
	})
	return file_ci_travis_proto_rawDescData
}

var file_ci_travis_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ci_travis_proto_goTypes = []interface{}{
	(*Travis)(nil),             // 0: proto.ci.Travis
	(*TravisAccount)(nil),      // 1: proto.ci.TravisAccount
	(*wrappers.BoolValue)(nil), // 2: google.protobuf.BoolValue
	(*client.Permissions)(nil), // 3: proto.Permissions
}
var file_ci_travis_proto_depIdxs = []int32{
	2, // 0: proto.ci.Travis.enabled:type_name -> google.protobuf.BoolValue
	1, // 1: proto.ci.Travis.masters:type_name -> proto.ci.TravisAccount
	3, // 2: proto.ci.TravisAccount.permissions:type_name -> proto.Permissions
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_ci_travis_proto_init() }
func file_ci_travis_proto_init() {
	if File_ci_travis_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ci_travis_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Travis); i {
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
		file_ci_travis_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TravisAccount); i {
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
			RawDescriptor: file_ci_travis_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ci_travis_proto_goTypes,
		DependencyIndexes: file_ci_travis_proto_depIdxs,
		MessageInfos:      file_ci_travis_proto_msgTypes,
	}.Build()
	File_ci_travis_proto = out.File
	file_ci_travis_proto_rawDesc = nil
	file_ci_travis_proto_goTypes = nil
	file_ci_travis_proto_depIdxs = nil
}
