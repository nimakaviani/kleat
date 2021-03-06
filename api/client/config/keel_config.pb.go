// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: config/keel_config.proto

package config

import (
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	plugins "github.com/spinnaker/kleat/api/client/plugins"
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

type KeelConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Plugins         *plugins.Plugins            `protobuf:"bytes,1,opt,name=plugins,proto3" json:"plugins,omitempty"`
	ResourceCheck   *KeelConfig_ResourceCheck   `protobuf:"bytes,2,opt,name=resourceCheck,proto3" json:"resourceCheck,omitempty"`
	ArtifactRefresh *KeelConfig_ArtifactRefresh `protobuf:"bytes,3,opt,name=artifactRefresh,json=artifact-refresh,proto3" json:"artifactRefresh,omitempty"`
	Constraints     *KeelConfig_Constraints     `protobuf:"bytes,4,opt,name=constraints,proto3" json:"constraints,omitempty"`
}

func (x *KeelConfig) Reset() {
	*x = KeelConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_keel_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeelConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeelConfig) ProtoMessage() {}

func (x *KeelConfig) ProtoReflect() protoreflect.Message {
	mi := &file_config_keel_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeelConfig.ProtoReflect.Descriptor instead.
func (*KeelConfig) Descriptor() ([]byte, []int) {
	return file_config_keel_config_proto_rawDescGZIP(), []int{0}
}

func (x *KeelConfig) GetPlugins() *plugins.Plugins {
	if x != nil {
		return x.Plugins
	}
	return nil
}

func (x *KeelConfig) GetResourceCheck() *KeelConfig_ResourceCheck {
	if x != nil {
		return x.ResourceCheck
	}
	return nil
}

func (x *KeelConfig) GetArtifactRefresh() *KeelConfig_ArtifactRefresh {
	if x != nil {
		return x.ArtifactRefresh
	}
	return nil
}

func (x *KeelConfig) GetConstraints() *KeelConfig_Constraints {
	if x != nil {
		return x.Constraints
	}
	return nil
}

type KeelConfig_ResourceCheck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MinAgeDuration *duration.Duration `protobuf:"bytes,1,opt,name=minAgeDuration,proto3" json:"minAgeDuration,omitempty"`
}

func (x *KeelConfig_ResourceCheck) Reset() {
	*x = KeelConfig_ResourceCheck{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_keel_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeelConfig_ResourceCheck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeelConfig_ResourceCheck) ProtoMessage() {}

func (x *KeelConfig_ResourceCheck) ProtoReflect() protoreflect.Message {
	mi := &file_config_keel_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeelConfig_ResourceCheck.ProtoReflect.Descriptor instead.
func (*KeelConfig_ResourceCheck) Descriptor() ([]byte, []int) {
	return file_config_keel_config_proto_rawDescGZIP(), []int{0, 0}
}

func (x *KeelConfig_ResourceCheck) GetMinAgeDuration() *duration.Duration {
	if x != nil {
		return x.MinAgeDuration
	}
	return nil
}

type KeelConfig_Constraints struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ManualJudgement *KeelConfig_ManualJudgement `protobuf:"bytes,1,opt,name=manualJudgement,json=manual-judgement,proto3" json:"manualJudgement,omitempty"`
}

func (x *KeelConfig_Constraints) Reset() {
	*x = KeelConfig_Constraints{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_keel_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeelConfig_Constraints) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeelConfig_Constraints) ProtoMessage() {}

func (x *KeelConfig_Constraints) ProtoReflect() protoreflect.Message {
	mi := &file_config_keel_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeelConfig_Constraints.ProtoReflect.Descriptor instead.
func (*KeelConfig_Constraints) Descriptor() ([]byte, []int) {
	return file_config_keel_config_proto_rawDescGZIP(), []int{0, 1}
}

func (x *KeelConfig_Constraints) GetManualJudgement() *KeelConfig_ManualJudgement {
	if x != nil {
		return x.ManualJudgement
	}
	return nil
}

type KeelConfig_ManualJudgement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InteractiveNotifactions *KeelConfig_InteractiveNotifications `protobuf:"bytes,1,opt,name=interactiveNotifactions,json=interactive-notifications,proto3" json:"interactiveNotifactions,omitempty"`
}

func (x *KeelConfig_ManualJudgement) Reset() {
	*x = KeelConfig_ManualJudgement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_keel_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeelConfig_ManualJudgement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeelConfig_ManualJudgement) ProtoMessage() {}

func (x *KeelConfig_ManualJudgement) ProtoReflect() protoreflect.Message {
	mi := &file_config_keel_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeelConfig_ManualJudgement.ProtoReflect.Descriptor instead.
func (*KeelConfig_ManualJudgement) Descriptor() ([]byte, []int) {
	return file_config_keel_config_proto_rawDescGZIP(), []int{0, 2}
}

func (x *KeelConfig_ManualJudgement) GetInteractiveNotifactions() *KeelConfig_InteractiveNotifications {
	if x != nil {
		return x.InteractiveNotifactions
	}
	return nil
}

type KeelConfig_InteractiveNotifications struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled *wrappers.BoolValue `protobuf:"bytes,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

func (x *KeelConfig_InteractiveNotifications) Reset() {
	*x = KeelConfig_InteractiveNotifications{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_keel_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeelConfig_InteractiveNotifications) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeelConfig_InteractiveNotifications) ProtoMessage() {}

func (x *KeelConfig_InteractiveNotifications) ProtoReflect() protoreflect.Message {
	mi := &file_config_keel_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeelConfig_InteractiveNotifications.ProtoReflect.Descriptor instead.
func (*KeelConfig_InteractiveNotifications) Descriptor() ([]byte, []int) {
	return file_config_keel_config_proto_rawDescGZIP(), []int{0, 3}
}

func (x *KeelConfig_InteractiveNotifications) GetEnabled() *wrappers.BoolValue {
	if x != nil {
		return x.Enabled
	}
	return nil
}

type KeelConfig_ArtifactRefresh struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Frequency string `protobuf:"bytes,1,opt,name=frequency,proto3" json:"frequency,omitempty"`
}

func (x *KeelConfig_ArtifactRefresh) Reset() {
	*x = KeelConfig_ArtifactRefresh{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_keel_config_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeelConfig_ArtifactRefresh) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeelConfig_ArtifactRefresh) ProtoMessage() {}

func (x *KeelConfig_ArtifactRefresh) ProtoReflect() protoreflect.Message {
	mi := &file_config_keel_config_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeelConfig_ArtifactRefresh.ProtoReflect.Descriptor instead.
func (*KeelConfig_ArtifactRefresh) Descriptor() ([]byte, []int) {
	return file_config_keel_config_proto_rawDescGZIP(), []int{0, 4}
}

func (x *KeelConfig_ArtifactRefresh) GetFrequency() string {
	if x != nil {
		return x.Frequency
	}
	return ""
}

var File_config_keel_config_proto protoreflect.FileDescriptor

var file_config_keel_config_proto_rawDesc = []byte{
	0x0a, 0x18, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6b, 0x65, 0x65, 0x6c, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x15, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x73, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xe7, 0x05, 0x0a, 0x0a, 0x4b, 0x65, 0x65, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x30,
	0x0a, 0x07, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x2e,
	0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73, 0x52, 0x07, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x73,
	0x12, 0x4c, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4b, 0x65, 0x65, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52,
	0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x53,
	0x0a, 0x0f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4b, 0x65, 0x65, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x52, 0x10, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2d, 0x72, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x12, 0x46, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e,
	0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4b, 0x65, 0x65, 0x6c, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x52, 0x0b,
	0x63, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x1a, 0x52, 0x0a, 0x0d, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x41, 0x0a, 0x0e,
	0x6d, 0x69, 0x6e, 0x41, 0x67, 0x65, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0e, 0x6d, 0x69, 0x6e, 0x41, 0x67, 0x65, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a,
	0x62, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x74, 0x73, 0x12, 0x53,
	0x0a, 0x0f, 0x6d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4b, 0x65, 0x65, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x4d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x4a, 0x75, 0x64, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x10, 0x6d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x2d, 0x6a, 0x75, 0x64, 0x67, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x1a, 0x80, 0x01, 0x0a, 0x0f, 0x4d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x4a, 0x75,
	0x64, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x6d, 0x0a, 0x17, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x4b, 0x65, 0x65, 0x6c, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x4e, 0x6f,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x19, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x2d, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x50, 0x0a, 0x18, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x34, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x1a, 0x2f, 0x0a, 0x0f, 0x41, 0x72, 0x74, 0x69,
	0x66, 0x61, 0x63, 0x74, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x66,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x66, 0x72, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65,
	0x72, 0x2f, 0x6b, 0x6c, 0x65, 0x61, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_config_keel_config_proto_rawDescOnce sync.Once
	file_config_keel_config_proto_rawDescData = file_config_keel_config_proto_rawDesc
)

func file_config_keel_config_proto_rawDescGZIP() []byte {
	file_config_keel_config_proto_rawDescOnce.Do(func() {
		file_config_keel_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_keel_config_proto_rawDescData)
	})
	return file_config_keel_config_proto_rawDescData
}

var file_config_keel_config_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_config_keel_config_proto_goTypes = []interface{}{
	(*KeelConfig)(nil),                          // 0: proto.config.KeelConfig
	(*KeelConfig_ResourceCheck)(nil),            // 1: proto.config.KeelConfig.ResourceCheck
	(*KeelConfig_Constraints)(nil),              // 2: proto.config.KeelConfig.Constraints
	(*KeelConfig_ManualJudgement)(nil),          // 3: proto.config.KeelConfig.ManualJudgement
	(*KeelConfig_InteractiveNotifications)(nil), // 4: proto.config.KeelConfig.InteractiveNotifications
	(*KeelConfig_ArtifactRefresh)(nil),          // 5: proto.config.KeelConfig.ArtifactRefresh
	(*plugins.Plugins)(nil),                     // 6: proto.plugins.Plugins
	(*duration.Duration)(nil),                   // 7: google.protobuf.Duration
	(*wrappers.BoolValue)(nil),                  // 8: google.protobuf.BoolValue
}
var file_config_keel_config_proto_depIdxs = []int32{
	6, // 0: proto.config.KeelConfig.plugins:type_name -> proto.plugins.Plugins
	1, // 1: proto.config.KeelConfig.resourceCheck:type_name -> proto.config.KeelConfig.ResourceCheck
	5, // 2: proto.config.KeelConfig.artifactRefresh:type_name -> proto.config.KeelConfig.ArtifactRefresh
	2, // 3: proto.config.KeelConfig.constraints:type_name -> proto.config.KeelConfig.Constraints
	7, // 4: proto.config.KeelConfig.ResourceCheck.minAgeDuration:type_name -> google.protobuf.Duration
	3, // 5: proto.config.KeelConfig.Constraints.manualJudgement:type_name -> proto.config.KeelConfig.ManualJudgement
	4, // 6: proto.config.KeelConfig.ManualJudgement.interactiveNotifactions:type_name -> proto.config.KeelConfig.InteractiveNotifications
	8, // 7: proto.config.KeelConfig.InteractiveNotifications.enabled:type_name -> google.protobuf.BoolValue
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_config_keel_config_proto_init() }
func file_config_keel_config_proto_init() {
	if File_config_keel_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_keel_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeelConfig); i {
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
		file_config_keel_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeelConfig_ResourceCheck); i {
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
		file_config_keel_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeelConfig_Constraints); i {
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
		file_config_keel_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeelConfig_ManualJudgement); i {
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
		file_config_keel_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeelConfig_InteractiveNotifications); i {
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
		file_config_keel_config_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeelConfig_ArtifactRefresh); i {
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
			RawDescriptor: file_config_keel_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_keel_config_proto_goTypes,
		DependencyIndexes: file_config_keel_config_proto_depIdxs,
		MessageInfos:      file_config_keel_config_proto_msgTypes,
	}.Build()
	File_config_keel_config_proto = out.File
	file_config_keel_config_proto_rawDesc = nil
	file_config_keel_config_proto_goTypes = nil
	file_config_keel_config_proto_depIdxs = nil
}
