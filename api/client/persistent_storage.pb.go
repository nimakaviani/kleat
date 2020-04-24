// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.2
// source: persistent_storage.proto

package client

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

// Configuration for a Google Cloud Storage persistent store
type GcsPersistentStore struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether this persistent store is enabled.
	Enabled string `protobuf:"bytes,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// A path to a JSON service account with permission to read and write to the bucket to be used as a backing store.
	JsonPath string `protobuf:"bytes,2,opt,name=jsonPath,proto3" json:"jsonPath,omitempty"`
	// The Google Cloud Platform project you are using to host the GCS bucket as a backing store.
	Project string `protobuf:"bytes,3,opt,name=project,proto3" json:"project,omitempty"`
	// The name of a storage bucket that your specified account has access to. If not specified, a
	// random name will be chosen. If you specify a globally unique bucket name that does not exist
	// yet, Halyard will create that bucket for you.
	Bucket string `protobuf:"bytes,4,opt,name=bucket,proto3" json:"bucket,omitempty"`
	// The root folder in the chosen bucket to place all of Spinnaker's persistent data in.
	RootFolder string `protobuf:"bytes,5,opt,name=rootFolder,proto3" json:"rootFolder,omitempty"`
	// This is only required if the bucket you specify does not exist yet.
	BucketLocation string `protobuf:"bytes,6,opt,name=bucketLocation,proto3" json:"bucketLocation,omitempty"`
}

func (x *GcsPersistentStore) Reset() {
	*x = GcsPersistentStore{}
	if protoimpl.UnsafeEnabled {
		mi := &file_persistent_storage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GcsPersistentStore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GcsPersistentStore) ProtoMessage() {}

func (x *GcsPersistentStore) ProtoReflect() protoreflect.Message {
	mi := &file_persistent_storage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GcsPersistentStore.ProtoReflect.Descriptor instead.
func (*GcsPersistentStore) Descriptor() ([]byte, []int) {
	return file_persistent_storage_proto_rawDescGZIP(), []int{0}
}

func (x *GcsPersistentStore) GetEnabled() string {
	if x != nil {
		return x.Enabled
	}
	return ""
}

func (x *GcsPersistentStore) GetJsonPath() string {
	if x != nil {
		return x.JsonPath
	}
	return ""
}

func (x *GcsPersistentStore) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *GcsPersistentStore) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *GcsPersistentStore) GetRootFolder() string {
	if x != nil {
		return x.RootFolder
	}
	return ""
}

func (x *GcsPersistentStore) GetBucketLocation() string {
	if x != nil {
		return x.BucketLocation
	}
	return ""
}

// Configuration for an Azure Storage persistent store.
type AzsPersistentStore struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether this persistent store is enabled.
	Enabled string `protobuf:"bytes,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// The name of an Azure Storage Account.
	StorageAccountName string `protobuf:"bytes,2,opt,name=storageAccountName,proto3" json:"storageAccountName,omitempty"`
	// The key to access the Azure Storage Account.
	StorageAccountKey string `protobuf:"bytes,3,opt,name=storageAccountKey,proto3" json:"storageAccountKey,omitempty"`
	// The container name in the chosen storage account to place Spinnaker’s
	// persistent data. Defaults to 'spinnaker' if unspecified.
	StorageContainerName string `protobuf:"bytes,4,opt,name=storageContainerName,proto3" json:"storageContainerName,omitempty"`
}

func (x *AzsPersistentStore) Reset() {
	*x = AzsPersistentStore{}
	if protoimpl.UnsafeEnabled {
		mi := &file_persistent_storage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AzsPersistentStore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AzsPersistentStore) ProtoMessage() {}

func (x *AzsPersistentStore) ProtoReflect() protoreflect.Message {
	mi := &file_persistent_storage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AzsPersistentStore.ProtoReflect.Descriptor instead.
func (*AzsPersistentStore) Descriptor() ([]byte, []int) {
	return file_persistent_storage_proto_rawDescGZIP(), []int{1}
}

func (x *AzsPersistentStore) GetEnabled() string {
	if x != nil {
		return x.Enabled
	}
	return ""
}

func (x *AzsPersistentStore) GetStorageAccountName() string {
	if x != nil {
		return x.StorageAccountName
	}
	return ""
}

func (x *AzsPersistentStore) GetStorageAccountKey() string {
	if x != nil {
		return x.StorageAccountKey
	}
	return ""
}

func (x *AzsPersistentStore) GetStorageContainerName() string {
	if x != nil {
		return x.StorageContainerName
	}
	return ""
}

// Configuration for an Oracle persistent store.
type OraclePersistentStore struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether this persistent store is enabled.
	Enabled string `protobuf:"bytes,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// The bucket name to store persistent state object in.
	BucketName string `protobuf:"bytes,2,opt,name=bucketName,proto3" json:"bucketName,omitempty"`
	// The namespace the bucket and objects should be created in.
	Namespace string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// An Oracle region (e.g., us-phoenix-1).
	Region string `protobuf:"bytes,4,opt,name=region,proto3" json:"region,omitempty"`
	// The OCID of the Oracle User you’re authenticating as.
	UserId string `protobuf:"bytes,5,opt,name=userId,proto3" json:"userId,omitempty"`
	// Fingerprint of the public key.
	Fingerprint string `protobuf:"bytes,6,opt,name=fingerprint,proto3" json:"fingerprint,omitempty"`
	// Path to the private key in PEM format.
	SshPrivateKeyFilePath string `protobuf:"bytes,7,opt,name=sshPrivateKeyFilePath,proto3" json:"sshPrivateKeyFilePath,omitempty"`
	// Passphrase used for the private key, if it is encrypted.
	PrivateKeyPassphrase string `protobuf:"bytes,8,opt,name=privateKeyPassphrase,proto3" json:"privateKeyPassphrase,omitempty"`
	// The OCID of the Oracle Tenancy to use.
	TenancyId string `protobuf:"bytes,9,opt,name=tenancyId,proto3" json:"tenancyId,omitempty"`
}

func (x *OraclePersistentStore) Reset() {
	*x = OraclePersistentStore{}
	if protoimpl.UnsafeEnabled {
		mi := &file_persistent_storage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OraclePersistentStore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OraclePersistentStore) ProtoMessage() {}

func (x *OraclePersistentStore) ProtoReflect() protoreflect.Message {
	mi := &file_persistent_storage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OraclePersistentStore.ProtoReflect.Descriptor instead.
func (*OraclePersistentStore) Descriptor() ([]byte, []int) {
	return file_persistent_storage_proto_rawDescGZIP(), []int{2}
}

func (x *OraclePersistentStore) GetEnabled() string {
	if x != nil {
		return x.Enabled
	}
	return ""
}

func (x *OraclePersistentStore) GetBucketName() string {
	if x != nil {
		return x.BucketName
	}
	return ""
}

func (x *OraclePersistentStore) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *OraclePersistentStore) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *OraclePersistentStore) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *OraclePersistentStore) GetFingerprint() string {
	if x != nil {
		return x.Fingerprint
	}
	return ""
}

func (x *OraclePersistentStore) GetSshPrivateKeyFilePath() string {
	if x != nil {
		return x.SshPrivateKeyFilePath
	}
	return ""
}

func (x *OraclePersistentStore) GetPrivateKeyPassphrase() string {
	if x != nil {
		return x.PrivateKeyPassphrase
	}
	return ""
}

func (x *OraclePersistentStore) GetTenancyId() string {
	if x != nil {
		return x.TenancyId
	}
	return ""
}

var File_persistent_storage_proto protoreflect.FileDescriptor

var file_persistent_storage_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xc4, 0x01, 0x0a, 0x12, 0x47, 0x63, 0x73, 0x50, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74,
	0x65, 0x6e, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6a, 0x73, 0x6f, 0x6e, 0x50, 0x61, 0x74, 0x68, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6a, 0x73, 0x6f, 0x6e, 0x50, 0x61, 0x74, 0x68, 0x12, 0x18,
	0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x6f, 0x6f, 0x74, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x6f, 0x6f, 0x74, 0x46, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x12, 0x26, 0x0a, 0x0e, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xc0, 0x01, 0x0a, 0x12, 0x41, 0x7a, 0x73,
	0x50, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x2e, 0x0a, 0x12, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x32, 0x0a, 0x14, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xc9, 0x02, 0x0a, 0x15,
	0x4f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x50, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x74,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12,
	0x1e, 0x0a, 0x0a, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a,
	0x0b, 0x66, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x66, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x12,
	0x34, 0x0a, 0x15, 0x73, 0x73, 0x68, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79,
	0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15,
	0x73, 0x73, 0x68, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x46, 0x69, 0x6c,
	0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x32, 0x0a, 0x14, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x4b, 0x65, 0x79, 0x50, 0x61, 0x73, 0x73, 0x70, 0x68, 0x72, 0x61, 0x73, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x14, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x50,
	0x61, 0x73, 0x73, 0x70, 0x68, 0x72, 0x61, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x63, 0x79, 0x49, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x65,
	0x6e, 0x61, 0x6e, 0x63, 0x79, 0x49, 0x64, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72, 0x2f,
	0x6b, 0x6c, 0x65, 0x61, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_persistent_storage_proto_rawDescOnce sync.Once
	file_persistent_storage_proto_rawDescData = file_persistent_storage_proto_rawDesc
)

func file_persistent_storage_proto_rawDescGZIP() []byte {
	file_persistent_storage_proto_rawDescOnce.Do(func() {
		file_persistent_storage_proto_rawDescData = protoimpl.X.CompressGZIP(file_persistent_storage_proto_rawDescData)
	})
	return file_persistent_storage_proto_rawDescData
}

var file_persistent_storage_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_persistent_storage_proto_goTypes = []interface{}{
	(*GcsPersistentStore)(nil),    // 0: proto.GcsPersistentStore
	(*AzsPersistentStore)(nil),    // 1: proto.AzsPersistentStore
	(*OraclePersistentStore)(nil), // 2: proto.OraclePersistentStore
}
var file_persistent_storage_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_persistent_storage_proto_init() }
func file_persistent_storage_proto_init() {
	if File_persistent_storage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_persistent_storage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GcsPersistentStore); i {
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
		file_persistent_storage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AzsPersistentStore); i {
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
		file_persistent_storage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OraclePersistentStore); i {
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
			RawDescriptor: file_persistent_storage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_persistent_storage_proto_goTypes,
		DependencyIndexes: file_persistent_storage_proto_depIdxs,
		MessageInfos:      file_persistent_storage_proto_msgTypes,
	}.Build()
	File_persistent_storage_proto = out.File
	file_persistent_storage_proto_rawDesc = nil
	file_persistent_storage_proto_goTypes = nil
	file_persistent_storage_proto_depIdxs = nil
}
