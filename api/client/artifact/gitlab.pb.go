// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: artifact/gitlab.proto

package artifact

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

// Configuration for the GitLab artifact provider.
type GitLab struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether the GitLab artifact provider is enabled.
	Enabled *wrappers.BoolValue `protobuf:"bytes,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// The list of configured GitLab accounts.
	Accounts []*GitLabAccount `protobuf:"bytes,2,rep,name=accounts,proto3" json:"accounts,omitempty"`
}

func (x *GitLab) Reset() {
	*x = GitLab{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artifact_gitlab_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GitLab) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GitLab) ProtoMessage() {}

func (x *GitLab) ProtoReflect() protoreflect.Message {
	mi := &file_artifact_gitlab_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GitLab.ProtoReflect.Descriptor instead.
func (*GitLab) Descriptor() ([]byte, []int) {
	return file_artifact_gitlab_proto_rawDescGZIP(), []int{0}
}

func (x *GitLab) GetEnabled() *wrappers.BoolValue {
	if x != nil {
		return x.Enabled
	}
	return nil
}

func (x *GitLab) GetAccounts() []*GitLabAccount {
	if x != nil {
		return x.Accounts
	}
	return nil
}

// Configuration for a GitLab artifact account. Either `token` or
// `tokenFile` should be specified as means of authentication.
type GitLabAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the account.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The GitLab access token.
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	// The path to a file containing the GitLab access token.
	TokenFile string `protobuf:"bytes,3,opt,name=tokenFile,proto3" json:"tokenFile,omitempty"`
}

func (x *GitLabAccount) Reset() {
	*x = GitLabAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artifact_gitlab_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GitLabAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GitLabAccount) ProtoMessage() {}

func (x *GitLabAccount) ProtoReflect() protoreflect.Message {
	mi := &file_artifact_gitlab_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GitLabAccount.ProtoReflect.Descriptor instead.
func (*GitLabAccount) Descriptor() ([]byte, []int) {
	return file_artifact_gitlab_proto_rawDescGZIP(), []int{1}
}

func (x *GitLabAccount) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GitLabAccount) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *GitLabAccount) GetTokenFile() string {
	if x != nil {
		return x.TokenFile
	}
	return ""
}

var File_artifact_gitlab_proto protoreflect.FileDescriptor

var file_artifact_gitlab_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2f, 0x67, 0x69, 0x74, 0x6c, 0x61,
	0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61,
	0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x79, 0x0a, 0x06, 0x47, 0x69, 0x74, 0x4c, 0x61,
	0x62, 0x12, 0x34, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x39, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e, 0x47, 0x69, 0x74, 0x4c, 0x61,
	0x62, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x22, 0x57, 0x0a, 0x0d, 0x47, 0x69, 0x74, 0x4c, 0x61, 0x62, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c, 0x0a,
	0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x46, 0x69, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x46, 0x69, 0x6c, 0x65, 0x42, 0x30, 0x5a, 0x2e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x69, 0x6e, 0x6e, 0x61,
	0x6b, 0x65, 0x72, 0x2f, 0x6b, 0x6c, 0x65, 0x61, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_artifact_gitlab_proto_rawDescOnce sync.Once
	file_artifact_gitlab_proto_rawDescData = file_artifact_gitlab_proto_rawDesc
)

func file_artifact_gitlab_proto_rawDescGZIP() []byte {
	file_artifact_gitlab_proto_rawDescOnce.Do(func() {
		file_artifact_gitlab_proto_rawDescData = protoimpl.X.CompressGZIP(file_artifact_gitlab_proto_rawDescData)
	})
	return file_artifact_gitlab_proto_rawDescData
}

var file_artifact_gitlab_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_artifact_gitlab_proto_goTypes = []interface{}{
	(*GitLab)(nil),             // 0: proto.artifact.GitLab
	(*GitLabAccount)(nil),      // 1: proto.artifact.GitLabAccount
	(*wrappers.BoolValue)(nil), // 2: google.protobuf.BoolValue
}
var file_artifact_gitlab_proto_depIdxs = []int32{
	2, // 0: proto.artifact.GitLab.enabled:type_name -> google.protobuf.BoolValue
	1, // 1: proto.artifact.GitLab.accounts:type_name -> proto.artifact.GitLabAccount
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_artifact_gitlab_proto_init() }
func file_artifact_gitlab_proto_init() {
	if File_artifact_gitlab_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_artifact_gitlab_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GitLab); i {
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
		file_artifact_gitlab_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GitLabAccount); i {
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
			RawDescriptor: file_artifact_gitlab_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_artifact_gitlab_proto_goTypes,
		DependencyIndexes: file_artifact_gitlab_proto_depIdxs,
		MessageInfos:      file_artifact_gitlab_proto_msgTypes,
	}.Build()
	File_artifact_gitlab_proto = out.File
	file_artifact_gitlab_proto_rawDesc = nil
	file_artifact_gitlab_proto_goTypes = nil
	file_artifact_gitlab_proto_depIdxs = nil
}
