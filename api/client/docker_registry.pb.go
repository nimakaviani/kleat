// Code generated by protoc-gen-go. DO NOT EDIT.
// source: docker_registry.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Configuration for the Docker Registry provider.
type DockerRegistryProvider struct {
	// Whether the provider is enabled.
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// The list of configured accounts.
	Accounts []*DockerRegistryAccount `protobuf:"bytes,2,rep,name=accounts,proto3" json:"accounts,omitempty"`
	// The name of the primary account.
	PrimaryAccount       string   `protobuf:"bytes,3,opt,name=primaryAccount,proto3" json:"primaryAccount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DockerRegistryProvider) Reset()         { *m = DockerRegistryProvider{} }
func (m *DockerRegistryProvider) String() string { return proto.CompactTextString(m) }
func (*DockerRegistryProvider) ProtoMessage()    {}
func (*DockerRegistryProvider) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef83379b340f7e71, []int{0}
}

func (m *DockerRegistryProvider) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DockerRegistryProvider.Unmarshal(m, b)
}
func (m *DockerRegistryProvider) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DockerRegistryProvider.Marshal(b, m, deterministic)
}
func (m *DockerRegistryProvider) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DockerRegistryProvider.Merge(m, src)
}
func (m *DockerRegistryProvider) XXX_Size() int {
	return xxx_messageInfo_DockerRegistryProvider.Size(m)
}
func (m *DockerRegistryProvider) XXX_DiscardUnknown() {
	xxx_messageInfo_DockerRegistryProvider.DiscardUnknown(m)
}

var xxx_messageInfo_DockerRegistryProvider proto.InternalMessageInfo

func (m *DockerRegistryProvider) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *DockerRegistryProvider) GetAccounts() []*DockerRegistryAccount {
	if m != nil {
		return m.Accounts
	}
	return nil
}

func (m *DockerRegistryProvider) GetPrimaryAccount() string {
	if m != nil {
		return m.PrimaryAccount
	}
	return ""
}

// A credential able to authenticate against a set of Docker repositories.
type DockerRegistryAccount struct {
	// The name of the account.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// (Required) The registry address from which to pull and deploy images
	// (e.g., `https://index.docker.io`).
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// The number of seconds between polling the Docker registry. Certain
	// registries are sensitive to over-polling, and larger intervals
	// (e.g., 10 minutes = 600 seconds) are desirable if you experience rate
	// limiting. Defaults to `30`.
	CacheIntervalSeconds int32 `protobuf:"varint,3,opt,name=cacheIntervalSeconds,proto3" json:"cacheIntervalSeconds,omitempty"`
	// The number of threads on which to cache all provided repositories.
	// Really only useful if you have a ton of repos. Defaults to 1.
	CacheThreads int32 `protobuf:"varint,4,opt,name=cacheThreads,proto3" json:"cacheThreads,omitempty"`
	// Timeout in milliseconds for provided repositories. Defaults to
	// `60,000`.
	ClientTimeoutMillis int32 `protobuf:"varint,5,opt,name=clientTimeoutMillis,proto3" json:"clientTimeoutMillis,omitempty"`
	// The email associated with your Docker registry. Often this only needs to
	// be well-formed, rather than be a real address.
	Email string `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
	// The environment name for the account. Many accounts can share the
	// same environment (e.g., dev, test, prod).
	Environment string `protobuf:"bytes,7,opt,name=environment,proto3" json:"environment,omitempty"`
	// If `true`, Spinnaker will treat the Docker registry as insecure and not
	// validate the SSL certificate. Defaults to `false`.
	InsecureRegistry bool `protobuf:"varint,8,opt,name=insecureRegistry,proto3" json:"insecureRegistry,omitempty"`
	// Pagination size for the Docker `repository _catalog` endpoint. Defaults
	// to `100`.
	PaginateSize int32 `protobuf:"varint,9,opt,name=paginateSize,proto3" json:"paginateSize,omitempty"`
	// The Docker registry password. Only one of `password`, `passwordCommand`,
	// and `passwordFile` should be specified.
	Password string `protobuf:"bytes,10,opt,name=password,proto3" json:"password,omitempty"`
	// Command to retrieve Docker token/password. The command must be available
	// in the environment. Only one of `password`, `passwordCommand`, and
	// `passwordFile` should be specified.
	PasswordCommand string `protobuf:"bytes,11,opt,name=passwordCommand,proto3" json:"passwordCommand,omitempty"`
	// The path to a file containing your Docker password in plaintext (not a
	// Docker `config.json` file). Only one of `password`, `passwordCommand`,
	// and `passwordFile` should be specified.
	PasswordFile string `protobuf:"bytes,12,opt,name=passwordFile,proto3" json:"passwordFile,omitempty"`
	// Fiat permissions configuration.
	Permissions *Permissions `protobuf:"bytes,13,opt,name=permissions,proto3" json:"permissions,omitempty"`
	// (Deprecated) List of required Fiat permission groups. Configure
	// `permissions` instead.
	RequiredGroupMemberships []string `protobuf:"bytes,14,rep,name=requiredGroupMemberships,proto3" json:"requiredGroupMemberships,omitempty"`
	// An optional list of repositories from which to cache images. If not
	// provided, Spinnaker will attempt to read accessible repositories from the
	// `registries _catalog` endpoint.
	Repositories []string `protobuf:"bytes,15,rep,name=repositories,proto3" json:"repositories,omitempty"`
	// If `true`, Spinnaker will sort tags by creation date. Defaults to
	// `false`.
	SortTagsByDate bool `protobuf:"varint,16,opt,name=sortTagsByDate,proto3" json:"sortTagsByDate,omitempty"`
	// If `true`, Spinnaker will track digest changes. This is not recommended
	// because it greatly increases queries to the registry, and most
	// registries are flaky. Defaults to `false`.
	TrackDigests bool `protobuf:"varint,17,opt,name=trackDigests,proto3" json:"trackDigests,omitempty"`
	// The username associated with this Docker registry.
	Username             string   `protobuf:"bytes,18,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DockerRegistryAccount) Reset()         { *m = DockerRegistryAccount{} }
func (m *DockerRegistryAccount) String() string { return proto.CompactTextString(m) }
func (*DockerRegistryAccount) ProtoMessage()    {}
func (*DockerRegistryAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef83379b340f7e71, []int{1}
}

func (m *DockerRegistryAccount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DockerRegistryAccount.Unmarshal(m, b)
}
func (m *DockerRegistryAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DockerRegistryAccount.Marshal(b, m, deterministic)
}
func (m *DockerRegistryAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DockerRegistryAccount.Merge(m, src)
}
func (m *DockerRegistryAccount) XXX_Size() int {
	return xxx_messageInfo_DockerRegistryAccount.Size(m)
}
func (m *DockerRegistryAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_DockerRegistryAccount.DiscardUnknown(m)
}

var xxx_messageInfo_DockerRegistryAccount proto.InternalMessageInfo

func (m *DockerRegistryAccount) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DockerRegistryAccount) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *DockerRegistryAccount) GetCacheIntervalSeconds() int32 {
	if m != nil {
		return m.CacheIntervalSeconds
	}
	return 0
}

func (m *DockerRegistryAccount) GetCacheThreads() int32 {
	if m != nil {
		return m.CacheThreads
	}
	return 0
}

func (m *DockerRegistryAccount) GetClientTimeoutMillis() int32 {
	if m != nil {
		return m.ClientTimeoutMillis
	}
	return 0
}

func (m *DockerRegistryAccount) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *DockerRegistryAccount) GetEnvironment() string {
	if m != nil {
		return m.Environment
	}
	return ""
}

func (m *DockerRegistryAccount) GetInsecureRegistry() bool {
	if m != nil {
		return m.InsecureRegistry
	}
	return false
}

func (m *DockerRegistryAccount) GetPaginateSize() int32 {
	if m != nil {
		return m.PaginateSize
	}
	return 0
}

func (m *DockerRegistryAccount) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *DockerRegistryAccount) GetPasswordCommand() string {
	if m != nil {
		return m.PasswordCommand
	}
	return ""
}

func (m *DockerRegistryAccount) GetPasswordFile() string {
	if m != nil {
		return m.PasswordFile
	}
	return ""
}

func (m *DockerRegistryAccount) GetPermissions() *Permissions {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func (m *DockerRegistryAccount) GetRequiredGroupMemberships() []string {
	if m != nil {
		return m.RequiredGroupMemberships
	}
	return nil
}

func (m *DockerRegistryAccount) GetRepositories() []string {
	if m != nil {
		return m.Repositories
	}
	return nil
}

func (m *DockerRegistryAccount) GetSortTagsByDate() bool {
	if m != nil {
		return m.SortTagsByDate
	}
	return false
}

func (m *DockerRegistryAccount) GetTrackDigests() bool {
	if m != nil {
		return m.TrackDigests
	}
	return false
}

func (m *DockerRegistryAccount) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func init() {
	proto.RegisterType((*DockerRegistryProvider)(nil), "proto.DockerRegistryProvider")
	proto.RegisterType((*DockerRegistryAccount)(nil), "proto.DockerRegistryAccount")
}

func init() { proto.RegisterFile("docker_registry.proto", fileDescriptor_ef83379b340f7e71) }

var fileDescriptor_ef83379b340f7e71 = []byte{
	// 474 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x15, 0x52, 0xb7, 0xc9, 0xba, 0xf4, 0xcf, 0xd2, 0xa2, 0x55, 0xc5, 0xc1, 0xca, 0x01,
	0x59, 0x1c, 0x2a, 0x14, 0x38, 0x20, 0x6e, 0x40, 0x04, 0xe2, 0x50, 0xa9, 0x72, 0x73, 0x47, 0x1b,
	0xef, 0x28, 0x19, 0xd5, 0xde, 0x35, 0x33, 0xeb, 0xa0, 0xf2, 0x1c, 0xbc, 0x0f, 0xaf, 0x86, 0xbc,
	0xae, 0x43, 0x1c, 0xc2, 0x29, 0x3b, 0xbf, 0xef, 0xdb, 0xd9, 0x2f, 0xe3, 0x11, 0x97, 0xc6, 0xe5,
	0xf7, 0x40, 0xdf, 0x08, 0x96, 0xc8, 0x9e, 0x1e, 0xae, 0x2b, 0x72, 0xde, 0xc9, 0x28, 0xfc, 0x5c,
	0x9d, 0x57, 0x40, 0x25, 0x32, 0xa3, 0xb3, 0xdc, 0x2a, 0x93, 0x5f, 0x03, 0xf1, 0x7c, 0x16, 0xee,
	0x64, 0x8f, 0x57, 0x6e, 0xc9, 0xad, 0xd1, 0x00, 0x49, 0x25, 0x8e, 0xc0, 0xea, 0x45, 0x01, 0x46,
	0x0d, 0x92, 0x41, 0x3a, 0xca, 0xba, 0x52, 0xbe, 0x13, 0x23, 0x9d, 0xe7, 0xae, 0xb6, 0x9e, 0xd5,
	0x93, 0x64, 0x98, 0xc6, 0xd3, 0x17, 0x6d, 0xbb, 0xeb, 0x7e, 0xab, 0x0f, 0xad, 0x29, 0xdb, 0xb8,
	0xe5, 0x4b, 0x71, 0x52, 0x11, 0x96, 0x7a, 0xa3, 0xa9, 0x61, 0x32, 0x48, 0xc7, 0xd9, 0x0e, 0x9d,
	0xfc, 0x8e, 0xc4, 0xe5, 0xde, 0x5e, 0x52, 0x8a, 0x03, 0xab, 0x4b, 0x08, 0x91, 0xc6, 0x59, 0x38,
	0x37, 0x49, 0xb5, 0x31, 0x04, 0xdc, 0xc4, 0x69, 0x70, 0x57, 0xca, 0xa9, 0xb8, 0xc8, 0x75, 0xbe,
	0x82, 0xaf, 0xd6, 0x03, 0xad, 0x75, 0x71, 0x07, 0xb9, 0xb3, 0x86, 0xc3, 0xab, 0x51, 0xb6, 0x57,
	0x93, 0x13, 0x71, 0x1c, 0xf8, 0x7c, 0x45, 0xa0, 0x0d, 0xab, 0x83, 0xe0, 0xed, 0x31, 0xf9, 0x5a,
	0x3c, 0xcb, 0x0b, 0x04, 0xeb, 0xe7, 0x58, 0x82, 0xab, 0xfd, 0x0d, 0x16, 0x05, 0xb2, 0x8a, 0x82,
	0x75, 0x9f, 0x24, 0x2f, 0x44, 0x04, 0xa5, 0xc6, 0x42, 0x1d, 0x86, 0x84, 0x6d, 0x21, 0x13, 0x11,
	0x83, 0x5d, 0x23, 0x39, 0x5b, 0x82, 0xf5, 0xea, 0x28, 0x68, 0xdb, 0x48, 0xbe, 0x12, 0x67, 0x68,
	0x19, 0xf2, 0x9a, 0xa0, 0x1b, 0x85, 0x1a, 0x85, 0xcf, 0xf1, 0x0f, 0x6f, 0x92, 0x57, 0x7a, 0x89,
	0x56, 0x7b, 0xb8, 0xc3, 0x9f, 0xa0, 0xc6, 0x6d, 0xf2, 0x6d, 0x26, 0xaf, 0xc4, 0xa8, 0xd2, 0xcc,
	0x3f, 0x1c, 0x19, 0x25, 0xc2, 0x73, 0x9b, 0x5a, 0xa6, 0xe2, 0xb4, 0x3b, 0x7f, 0x72, 0x65, 0xa9,
	0xad, 0x51, 0x71, 0xb0, 0xec, 0xe2, 0xf6, 0xa5, 0x16, 0x7d, 0xc6, 0x02, 0xd4, 0x71, 0xb0, 0xf5,
	0x98, 0x7c, 0x2b, 0xe2, 0xad, 0x7d, 0x53, 0x4f, 0x93, 0x41, 0x1a, 0x4f, 0xe5, 0xe3, 0xa2, 0xdc,
	0xfe, 0x55, 0xb2, 0x6d, 0x9b, 0x7c, 0x2f, 0x14, 0xc1, 0xf7, 0x1a, 0x09, 0xcc, 0x17, 0x72, 0x75,
	0x75, 0x03, 0xe5, 0x02, 0x88, 0x57, 0x58, 0xb1, 0x3a, 0x49, 0x86, 0xe9, 0x38, 0xfb, 0xaf, 0xde,
	0xa4, 0x22, 0xa8, 0x1c, 0xa3, 0x77, 0x84, 0xc0, 0xea, 0x34, 0xf8, 0x7b, 0xac, 0xd9, 0x40, 0x76,
	0xe4, 0xe7, 0x7a, 0xc9, 0x1f, 0x1f, 0x66, 0xda, 0x83, 0x3a, 0x0b, 0xd3, 0xdc, 0xa1, 0x4d, 0x2f,
	0x4f, 0x3a, 0xbf, 0x9f, 0xe1, 0x12, 0xd8, 0xb3, 0x3a, 0x0f, 0xae, 0x1e, 0x6b, 0x66, 0x59, 0x33,
	0x50, 0xd8, 0x47, 0xd9, 0xce, 0xb2, 0xab, 0x17, 0x87, 0xe1, 0x7f, 0xbe, 0xf9, 0x13, 0x00, 0x00,
	0xff, 0xff, 0xfa, 0xf9, 0xed, 0xd2, 0x92, 0x03, 0x00, 0x00,
}