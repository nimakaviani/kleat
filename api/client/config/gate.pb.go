// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: config/gate.proto

package config

import (
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	security "github.com/spinnaker/kleat/api/client/security"
	authn "github.com/spinnaker/kleat/api/client/security/authn"
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

// Configuration for the gate microservice.
type Gate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Web server configuration.
	Server *ServerConfig `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	// Configuration for cross-origin resource sharing.
	Cors *Cors `protobuf:"bytes,2,opt,name=cors,proto3" json:"cors,omitempty"`
	// Wrapper for Spring configuration properties (including OAuth2 authentication).
	Security *SpringSecurity `protobuf:"bytes,3,opt,name=security,proto3" json:"security,omitempty"`
	// Configuration for SAML authentication.
	Saml *authn.Saml `protobuf:"bytes,4,opt,name=saml,proto3" json:"saml,omitempty"`
	// Configuration for LDAP authentication.
	Ldap *authn.Ldap `protobuf:"bytes,5,opt,name=ldap,proto3" json:"ldap,omitempty"`
	// Configuration for X509 authentication.
	X509 *authn.X509 `protobuf:"bytes,6,opt,name=x509,proto3" json:"x509,omitempty"`
	// Wrapper for Google-specific authentication (ex: IAP).
	Google *Gate_GoogleConfig `protobuf:"bytes,7,opt,name=google,proto3" json:"google,omitempty"`
	// Configuration for the status of non-core services.
	Services     *Gate_Services     `protobuf:"bytes,8,opt,name=services,proto3" json:"services,omitempty"`
	Integrations *Gate_Integrations `protobuf:"bytes,9,opt,name=integrations,proto3" json:"integrations,omitempty"`
}

func (x *Gate) Reset() {
	*x = Gate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_gate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Gate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Gate) ProtoMessage() {}

func (x *Gate) ProtoReflect() protoreflect.Message {
	mi := &file_config_gate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Gate.ProtoReflect.Descriptor instead.
func (*Gate) Descriptor() ([]byte, []int) {
	return file_config_gate_proto_rawDescGZIP(), []int{0}
}

func (x *Gate) GetServer() *ServerConfig {
	if x != nil {
		return x.Server
	}
	return nil
}

func (x *Gate) GetCors() *Cors {
	if x != nil {
		return x.Cors
	}
	return nil
}

func (x *Gate) GetSecurity() *SpringSecurity {
	if x != nil {
		return x.Security
	}
	return nil
}

func (x *Gate) GetSaml() *authn.Saml {
	if x != nil {
		return x.Saml
	}
	return nil
}

func (x *Gate) GetLdap() *authn.Ldap {
	if x != nil {
		return x.Ldap
	}
	return nil
}

func (x *Gate) GetX509() *authn.X509 {
	if x != nil {
		return x.X509
	}
	return nil
}

func (x *Gate) GetGoogle() *Gate_GoogleConfig {
	if x != nil {
		return x.Google
	}
	return nil
}

func (x *Gate) GetServices() *Gate_Services {
	if x != nil {
		return x.Services
	}
	return nil
}

func (x *Gate) GetIntegrations() *Gate_Integrations {
	if x != nil {
		return x.Integrations
	}
	return nil
}

// Wrapper for Spring security configuration properties.
type SpringSecurity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Configuration for OAuth2 authentication.
	Oauth2 *authn.OAuth2 `protobuf:"bytes,1,opt,name=oauth2,proto3" json:"oauth2,omitempty"`
	// Configuration for basic authentication.
	Basic *authn.Basic `protobuf:"bytes,2,opt,name=basic,proto3" json:"basic,omitempty"`
}

func (x *SpringSecurity) Reset() {
	*x = SpringSecurity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_gate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpringSecurity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpringSecurity) ProtoMessage() {}

func (x *SpringSecurity) ProtoReflect() protoreflect.Message {
	mi := &file_config_gate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpringSecurity.ProtoReflect.Descriptor instead.
func (*SpringSecurity) Descriptor() ([]byte, []int) {
	return file_config_gate_proto_rawDescGZIP(), []int{1}
}

func (x *SpringSecurity) GetOauth2() *authn.OAuth2 {
	if x != nil {
		return x.Oauth2
	}
	return nil
}

func (x *SpringSecurity) GetBasic() *authn.Basic {
	if x != nil {
		return x.Basic
	}
	return nil
}

// Web server configuration.
type ServerConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// SSL configuration.
	Ssl *security.ApiSsl `protobuf:"bytes,1,opt,name=ssl,proto3" json:"ssl,omitempty"`
}

func (x *ServerConfig) Reset() {
	*x = ServerConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_gate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerConfig) ProtoMessage() {}

func (x *ServerConfig) ProtoReflect() protoreflect.Message {
	mi := &file_config_gate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerConfig.ProtoReflect.Descriptor instead.
func (*ServerConfig) Descriptor() ([]byte, []int) {
	return file_config_gate_proto_rawDescGZIP(), []int{2}
}

func (x *ServerConfig) GetSsl() *security.ApiSsl {
	if x != nil {
		return x.Ssl
	}
	return nil
}

// Configuration for cross-origin resource sharing.
type Cors struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A regex matching all URLs authentication redirects may come from.
	AllowedOriginsPattern string `protobuf:"bytes,1,opt,name=allowedOriginsPattern,proto3" json:"allowedOriginsPattern,omitempty"`
}

func (x *Cors) Reset() {
	*x = Cors{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_gate_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cors) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cors) ProtoMessage() {}

func (x *Cors) ProtoReflect() protoreflect.Message {
	mi := &file_config_gate_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cors.ProtoReflect.Descriptor instead.
func (*Cors) Descriptor() ([]byte, []int) {
	return file_config_gate_proto_rawDescGZIP(), []int{3}
}

func (x *Cors) GetAllowedOriginsPattern() string {
	if x != nil {
		return x.AllowedOriginsPattern
	}
	return ""
}

type Gate_Services struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kayenta *ServiceSettings `protobuf:"bytes,1,opt,name=kayenta,proto3" json:"kayenta,omitempty"`
	Deck    *ServiceSettings `protobuf:"bytes,2,opt,name=deck,proto3" json:"deck,omitempty"`
}

func (x *Gate_Services) Reset() {
	*x = Gate_Services{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_gate_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Gate_Services) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Gate_Services) ProtoMessage() {}

func (x *Gate_Services) ProtoReflect() protoreflect.Message {
	mi := &file_config_gate_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Gate_Services.ProtoReflect.Descriptor instead.
func (*Gate_Services) Descriptor() ([]byte, []int) {
	return file_config_gate_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Gate_Services) GetKayenta() *ServiceSettings {
	if x != nil {
		return x.Kayenta
	}
	return nil
}

func (x *Gate_Services) GetDeck() *ServiceSettings {
	if x != nil {
		return x.Deck
	}
	return nil
}

// Wrapper for Google-specific authentication.
type Gate_GoogleConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Configuration for Identity-Aware Proxy authentication.
	Iap *authn.Iap `protobuf:"bytes,1,opt,name=iap,proto3" json:"iap,omitempty"`
}

func (x *Gate_GoogleConfig) Reset() {
	*x = Gate_GoogleConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_gate_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Gate_GoogleConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Gate_GoogleConfig) ProtoMessage() {}

func (x *Gate_GoogleConfig) ProtoReflect() protoreflect.Message {
	mi := &file_config_gate_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Gate_GoogleConfig.ProtoReflect.Descriptor instead.
func (*Gate_GoogleConfig) Descriptor() ([]byte, []int) {
	return file_config_gate_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Gate_GoogleConfig) GetIap() *authn.Iap {
	if x != nil {
		return x.Iap
	}
	return nil
}

// Wrapper for Gate integrations.
type Gate_Integrations struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Gremlin *Gate_Integrations_Gremlin `protobuf:"bytes,1,opt,name=gremlin,proto3" json:"gremlin,omitempty"`
}

func (x *Gate_Integrations) Reset() {
	*x = Gate_Integrations{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_gate_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Gate_Integrations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Gate_Integrations) ProtoMessage() {}

func (x *Gate_Integrations) ProtoReflect() protoreflect.Message {
	mi := &file_config_gate_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Gate_Integrations.ProtoReflect.Descriptor instead.
func (*Gate_Integrations) Descriptor() ([]byte, []int) {
	return file_config_gate_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Gate_Integrations) GetGremlin() *Gate_Integrations_Gremlin {
	if x != nil {
		return x.Gremlin
	}
	return nil
}

// Configuration for Gremlin fault-injection support.
type Gate_Integrations_Gremlin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether Gremlin is enabled.
	Enabled *wrappers.BoolValue `protobuf:"bytes,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
}

func (x *Gate_Integrations_Gremlin) Reset() {
	*x = Gate_Integrations_Gremlin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_gate_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Gate_Integrations_Gremlin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Gate_Integrations_Gremlin) ProtoMessage() {}

func (x *Gate_Integrations_Gremlin) ProtoReflect() protoreflect.Message {
	mi := &file_config_gate_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Gate_Integrations_Gremlin.ProtoReflect.Descriptor instead.
func (*Gate_Integrations_Gremlin) Descriptor() ([]byte, []int) {
	return file_config_gate_proto_rawDescGZIP(), []int{0, 2, 0}
}

func (x *Gate_Integrations_Gremlin) GetEnabled() *wrappers.BoolValue {
	if x != nil {
		return x.Enabled
	}
	return nil
}

var File_config_gate_proto protoreflect.FileDescriptor

var file_config_gate_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1a, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x6e, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2f, 0x73, 0x73, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xad, 0x06, 0x0a, 0x04, 0x47, 0x61, 0x74, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x04,
	0x63, 0x6f, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x72, 0x73, 0x52, 0x04,
	0x63, 0x6f, 0x72, 0x73, 0x12, 0x38, 0x0a, 0x08, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x53, 0x70, 0x72, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x52, 0x08, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x12, 0x2e,
	0x0a, 0x04, 0x73, 0x61, 0x6d, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x6e, 0x2e, 0x53, 0x61, 0x6d, 0x6c, 0x52, 0x04, 0x73, 0x61, 0x6d, 0x6c, 0x12, 0x2e,
	0x0a, 0x04, 0x6c, 0x64, 0x61, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x6e, 0x2e, 0x4c, 0x64, 0x61, 0x70, 0x52, 0x04, 0x6c, 0x64, 0x61, 0x70, 0x12, 0x2e,
	0x0a, 0x04, 0x78, 0x35, 0x30, 0x39, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x6e, 0x2e, 0x58, 0x35, 0x30, 0x39, 0x52, 0x04, 0x78, 0x35, 0x30, 0x39, 0x12, 0x37,
	0x0a, 0x06, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x47, 0x61,
	0x74, 0x65, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x06, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x12, 0x37, 0x0a, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x12, 0x43, 0x0a, 0x0c, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0c, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x76, 0x0a, 0x08, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x12, 0x37, 0x0a, 0x07, 0x6b, 0x61, 0x79, 0x65, 0x6e, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x52, 0x07, 0x6b, 0x61, 0x79, 0x65, 0x6e, 0x74, 0x61, 0x12, 0x31, 0x0a, 0x04, 0x64, 0x65,
	0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x04, 0x64, 0x65, 0x63, 0x6b, 0x1a, 0x3b, 0x0a,
	0x0c, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2b, 0x0a,
	0x03, 0x69, 0x61, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x6e, 0x2e, 0x49, 0x61, 0x70, 0x52, 0x03, 0x69, 0x61, 0x70, 0x1a, 0x92, 0x01, 0x0a, 0x0c, 0x49,
	0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x41, 0x0a, 0x07, 0x67,
	0x72, 0x65, 0x6d, 0x6c, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x47, 0x61, 0x74, 0x65,
	0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x47, 0x72,
	0x65, 0x6d, 0x6c, 0x69, 0x6e, 0x52, 0x07, 0x67, 0x72, 0x65, 0x6d, 0x6c, 0x69, 0x6e, 0x1a, 0x3f,
	0x0a, 0x07, 0x47, 0x72, 0x65, 0x6d, 0x6c, 0x69, 0x6e, 0x12, 0x34, 0x0a, 0x07, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f,
	0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x22,
	0x79, 0x0a, 0x0e, 0x53, 0x70, 0x72, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x12, 0x34, 0x0a, 0x06, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x2e, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x32, 0x52,
	0x06, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x32, 0x12, 0x31, 0x0a, 0x05, 0x62, 0x61, 0x73, 0x69, 0x63,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6e, 0x2e, 0x42, 0x61,
	0x73, 0x69, 0x63, 0x52, 0x05, 0x62, 0x61, 0x73, 0x69, 0x63, 0x22, 0x38, 0x0a, 0x0c, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x28, 0x0a, 0x03, 0x73, 0x73,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x41, 0x70, 0x69, 0x53, 0x73, 0x6c, 0x52,
	0x03, 0x73, 0x73, 0x6c, 0x22, 0x3c, 0x0a, 0x04, 0x43, 0x6f, 0x72, 0x73, 0x12, 0x34, 0x0a, 0x15,
	0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x64, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x73, 0x50, 0x61,
	0x74, 0x74, 0x65, 0x72, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x61, 0x6c, 0x6c,
	0x6f, 0x77, 0x65, 0x64, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x73, 0x50, 0x61, 0x74, 0x74, 0x65,
	0x72, 0x6e, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72, 0x2f, 0x6b, 0x6c, 0x65, 0x61, 0x74,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_gate_proto_rawDescOnce sync.Once
	file_config_gate_proto_rawDescData = file_config_gate_proto_rawDesc
)

func file_config_gate_proto_rawDescGZIP() []byte {
	file_config_gate_proto_rawDescOnce.Do(func() {
		file_config_gate_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_gate_proto_rawDescData)
	})
	return file_config_gate_proto_rawDescData
}

var file_config_gate_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_config_gate_proto_goTypes = []interface{}{
	(*Gate)(nil),                      // 0: proto.config.Gate
	(*SpringSecurity)(nil),            // 1: proto.config.SpringSecurity
	(*ServerConfig)(nil),              // 2: proto.config.ServerConfig
	(*Cors)(nil),                      // 3: proto.config.Cors
	(*Gate_Services)(nil),             // 4: proto.config.Gate.Services
	(*Gate_GoogleConfig)(nil),         // 5: proto.config.Gate.GoogleConfig
	(*Gate_Integrations)(nil),         // 6: proto.config.Gate.Integrations
	(*Gate_Integrations_Gremlin)(nil), // 7: proto.config.Gate.Integrations.Gremlin
	(*authn.Saml)(nil),                // 8: proto.security.authn.Saml
	(*authn.Ldap)(nil),                // 9: proto.security.authn.Ldap
	(*authn.X509)(nil),                // 10: proto.security.authn.X509
	(*authn.OAuth2)(nil),              // 11: proto.security.authn.OAuth2
	(*authn.Basic)(nil),               // 12: proto.security.authn.Basic
	(*security.ApiSsl)(nil),           // 13: proto.security.ApiSsl
	(*ServiceSettings)(nil),           // 14: proto.config.ServiceSettings
	(*authn.Iap)(nil),                 // 15: proto.security.authn.Iap
	(*wrappers.BoolValue)(nil),        // 16: google.protobuf.BoolValue
}
var file_config_gate_proto_depIdxs = []int32{
	2,  // 0: proto.config.Gate.server:type_name -> proto.config.ServerConfig
	3,  // 1: proto.config.Gate.cors:type_name -> proto.config.Cors
	1,  // 2: proto.config.Gate.security:type_name -> proto.config.SpringSecurity
	8,  // 3: proto.config.Gate.saml:type_name -> proto.security.authn.Saml
	9,  // 4: proto.config.Gate.ldap:type_name -> proto.security.authn.Ldap
	10, // 5: proto.config.Gate.x509:type_name -> proto.security.authn.X509
	5,  // 6: proto.config.Gate.google:type_name -> proto.config.Gate.GoogleConfig
	4,  // 7: proto.config.Gate.services:type_name -> proto.config.Gate.Services
	6,  // 8: proto.config.Gate.integrations:type_name -> proto.config.Gate.Integrations
	11, // 9: proto.config.SpringSecurity.oauth2:type_name -> proto.security.authn.OAuth2
	12, // 10: proto.config.SpringSecurity.basic:type_name -> proto.security.authn.Basic
	13, // 11: proto.config.ServerConfig.ssl:type_name -> proto.security.ApiSsl
	14, // 12: proto.config.Gate.Services.kayenta:type_name -> proto.config.ServiceSettings
	14, // 13: proto.config.Gate.Services.deck:type_name -> proto.config.ServiceSettings
	15, // 14: proto.config.Gate.GoogleConfig.iap:type_name -> proto.security.authn.Iap
	7,  // 15: proto.config.Gate.Integrations.gremlin:type_name -> proto.config.Gate.Integrations.Gremlin
	16, // 16: proto.config.Gate.Integrations.Gremlin.enabled:type_name -> google.protobuf.BoolValue
	17, // [17:17] is the sub-list for method output_type
	17, // [17:17] is the sub-list for method input_type
	17, // [17:17] is the sub-list for extension type_name
	17, // [17:17] is the sub-list for extension extendee
	0,  // [0:17] is the sub-list for field type_name
}

func init() { file_config_gate_proto_init() }
func file_config_gate_proto_init() {
	if File_config_gate_proto != nil {
		return
	}
	file_config_service_enabled_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_config_gate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Gate); i {
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
		file_config_gate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpringSecurity); i {
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
		file_config_gate_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerConfig); i {
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
		file_config_gate_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cors); i {
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
		file_config_gate_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Gate_Services); i {
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
		file_config_gate_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Gate_GoogleConfig); i {
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
		file_config_gate_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Gate_Integrations); i {
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
		file_config_gate_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Gate_Integrations_Gremlin); i {
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
			RawDescriptor: file_config_gate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_gate_proto_goTypes,
		DependencyIndexes: file_config_gate_proto_depIdxs,
		MessageInfos:      file_config_gate_proto_msgTypes,
	}.Build()
	File_config_gate_proto = out.File
	file_config_gate_proto_rawDesc = nil
	file_config_gate_proto_goTypes = nil
	file_config_gate_proto_depIdxs = nil
}
