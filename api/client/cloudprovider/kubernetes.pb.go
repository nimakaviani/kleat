// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.3
// source: cloudprovider/kubernetes.proto

package cloudprovider

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

// Configuration for the Kubernetes provider.
type Kubernetes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether the provider is enabled.
	Enabled *wrappers.BoolValue `protobuf:"bytes,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// The list of configured accounts.
	Accounts []*KubernetesAccount `protobuf:"bytes,2,rep,name=accounts,proto3" json:"accounts,omitempty"`
	// The name of the primary account.
	PrimaryAccount string `protobuf:"bytes,3,opt,name=primaryAccount,proto3" json:"primaryAccount,omitempty"`
}

func (x *Kubernetes) Reset() {
	*x = Kubernetes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_kubernetes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Kubernetes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Kubernetes) ProtoMessage() {}

func (x *Kubernetes) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_kubernetes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Kubernetes.ProtoReflect.Descriptor instead.
func (*Kubernetes) Descriptor() ([]byte, []int) {
	return file_cloudprovider_kubernetes_proto_rawDescGZIP(), []int{0}
}

func (x *Kubernetes) GetEnabled() *wrappers.BoolValue {
	if x != nil {
		return x.Enabled
	}
	return nil
}

func (x *Kubernetes) GetAccounts() []*KubernetesAccount {
	if x != nil {
		return x.Accounts
	}
	return nil
}

func (x *Kubernetes) GetPrimaryAccount() string {
	if x != nil {
		return x.PrimaryAccount
	}
	return ""
}

// Configuration for a Spinnaker Kubernetes account. An account maps to a
// credential that can authenticate against your Kubernetes cluster.
type KubernetesAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the account.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A list of resource kinds this Spinnaker account can deploy and
	// will cache. When no kinds are configured, this defaults to all kinds
	// described here:
	// https://spinnaker.io/reference/providers/kubernetes-v2/.
	// This can only be set when omitKinds is empty or not set.
	Kinds []string `protobuf:"bytes,3,rep,name=kinds,proto3" json:"kinds,omitempty"`
	// A list of resource kinds this Spinnaker account cannot deploy to or
	// cache. This can only be set when kinds is empty or not set.
	OmitKinds []string `protobuf:"bytes,4,rep,name=omitKinds,proto3" json:"omitKinds,omitempty"`
	// The kubernetes context to be managed by Spinnaker. See
	// http://kubernetes.io/docs/user-guide/kubeconfig-file/#context for
	// more information. When no context is configured for an account the
	// `current-context` in your kubeconfig is assumed.
	Context string `protobuf:"bytes,5,opt,name=context,proto3" json:"context,omitempty"`
	// Number of caching agents for this kubernetes account.
	// Each agent handles a subset of the namespaces available to this account.
	// By default, only 1 agent caches all kinds for all namespaces in the
	// account.
	CacheThreads int32 `protobuf:"varint,6,opt,name=cacheThreads,proto3" json:"cacheThreads,omitempty"`
	// A list of namespaces this Spinnaker account can deploy to and will
	// cache. When no namespaces are configured, this defaults to all
	// namespaces.
	Namespaces []string `protobuf:"bytes,7,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
	// A list of namespaces this Spinnaker account cannot deploy to or
	// cache. This can only be set when namespaces is empty or not set.
	OmitNamespaces []string `protobuf:"bytes,8,rep,name=omitNamespaces,proto3" json:"omitNamespaces,omitempty"`
	// The list of custom resources Clouddriver will manage and make
	// available for use in Patch and Delete (Manifest) stages.
	CustomResources []*KubernetesCustomResource `protobuf:"bytes,9,rep,name=customResources,proto3" json:"customResources,omitempty"`
	// The list of kind-specific caching policies.
	CachingPolicies []*KubernetesCachingPolicy `protobuf:"bytes,10,rep,name=cachingPolicies,proto3" json:"cachingPolicies,omitempty"`
	// The list of the Spinnaker docker registry account names this
	// Spinnaker account can use as image sources. These docker registry
	// accounts must be registered in your halconfig before you can add them
	// here.
	DockerRegistries []*KubernetesAccountDockerRegistry `protobuf:"bytes,11,rep,name=dockerRegistries,proto3" json:"dockerRegistries,omitempty"`
	// The list of OAuth scopes used by kubectl to fetch an OAuth token.
	OAuthScopes []string `protobuf:"bytes,12,rep,name=oAuthScopes,proto3" json:"oAuthScopes,omitempty"`
	// The path to your kubeconfig file. By default, it will be under the
	// Spinnaker user's home directory in the typical .kube/config location.
	KubeconfigFile string `protobuf:"bytes,13,opt,name=kubeconfigFile,proto3" json:"kubeconfigFile,omitempty"`
	// Fiat permissions configuration.
	Permissions *client.Permissions `protobuf:"bytes,14,opt,name=permissions,proto3" json:"permissions,omitempty"`
	// (Deprecated): List of required Fiat permission groups. Configure
	// `permissions` instead.
	RequiredGroupMembership []string `protobuf:"bytes,15,rep,name=requiredGroupMembership,proto3" json:"requiredGroupMembership,omitempty"`
	// When true, clouddriver will query manifest status during pipeline executions
	// using live data rather than the cache. This eliminates all time spent in the
	// "force cache refresh" task in pipelines, greatly reducing execution time.
	// Defaults to false.
	LiveManifestCalls *wrappers.BoolValue `protobuf:"bytes,16,opt,name=liveManifestCalls,proto3" json:"liveManifestCalls,omitempty"`
	// When true, Spinnaker attempt to authenticate against Kubernetes using a
	// Kubernetes service account. This only works when Halyard & Spinnaker are
	// deployed in Kubernetes. Read more about service accounts here:
	// https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/.
	// Defaults to false.
	ServiceAccount *wrappers.BoolValue `protobuf:"bytes,17,opt,name=serviceAccount,proto3" json:"serviceAccount,omitempty"`
	// The raw contents of your kubeconfig file. Ignored if kubeconfigFile is set.
	KubeconfigContents string `protobuf:"bytes,18,opt,name=kubeconfigContents,proto3" json:"kubeconfigContents,omitempty"`
	// The path to the kubectl executable. This should be omitted unless you
	// want to override the default kubectl exectuable.
	KubectlPath string `protobuf:"bytes,19,opt,name=kubectlPath,proto3" json:"kubectlPath,omitempty"`
	// If set, all calls to kubectl will time out after the specified number of
	// seconds.
	KubectlRequestTimeoutSeconds int32 `protobuf:"varint,22,opt,name=kubectlRequestTimeoutSeconds,proto3" json:"kubectlRequestTimeoutSeconds,omitempty"`
	// Whether to check whether the account has permission to read configured
	// kinds before caching them. Kinds that the account does not have permission
	// to read will be omitted from caching.
	// This field defaults to true, and it is recommended to leave it at the
	// default. If this field is set to false, any Kubernetes objects that
	// are unreadable by the account will break caching for all objects.
	CheckPermissionsOnStartup *wrappers.BoolValue `protobuf:"bytes,23,opt,name=checkPermissionsOnStartup,proto3" json:"checkPermissionsOnStartup,omitempty"`
	// When using OAuth to authenticate with your cluster, the name of the
	// service account to use.
	OAuthServiceAccount string `protobuf:"bytes,24,opt,name=oAuthServiceAccount,proto3" json:"oAuthServiceAccount,omitempty"`
	// If true, only cache Kubernetes objects that have been deployed by
	// Spinnaker, and ignore any other objects that exist in the cluster.
	// Defaults to false.
	OnlySpinnakerManaged *wrappers.BoolValue `protobuf:"bytes,27,opt,name=onlySpinnakerManaged,proto3" json:"onlySpinnakerManaged,omitempty"`
	// If true, enable detailed logging for all communications with the
	// Kubernetes cluster for this account. Defaults to false.
	Debug *wrappers.BoolValue `protobuf:"bytes,28,opt,name=debug,proto3" json:"debug,omitempty"`
}

func (x *KubernetesAccount) Reset() {
	*x = KubernetesAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_kubernetes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubernetesAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubernetesAccount) ProtoMessage() {}

func (x *KubernetesAccount) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_kubernetes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubernetesAccount.ProtoReflect.Descriptor instead.
func (*KubernetesAccount) Descriptor() ([]byte, []int) {
	return file_cloudprovider_kubernetes_proto_rawDescGZIP(), []int{1}
}

func (x *KubernetesAccount) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *KubernetesAccount) GetKinds() []string {
	if x != nil {
		return x.Kinds
	}
	return nil
}

func (x *KubernetesAccount) GetOmitKinds() []string {
	if x != nil {
		return x.OmitKinds
	}
	return nil
}

func (x *KubernetesAccount) GetContext() string {
	if x != nil {
		return x.Context
	}
	return ""
}

func (x *KubernetesAccount) GetCacheThreads() int32 {
	if x != nil {
		return x.CacheThreads
	}
	return 0
}

func (x *KubernetesAccount) GetNamespaces() []string {
	if x != nil {
		return x.Namespaces
	}
	return nil
}

func (x *KubernetesAccount) GetOmitNamespaces() []string {
	if x != nil {
		return x.OmitNamespaces
	}
	return nil
}

func (x *KubernetesAccount) GetCustomResources() []*KubernetesCustomResource {
	if x != nil {
		return x.CustomResources
	}
	return nil
}

func (x *KubernetesAccount) GetCachingPolicies() []*KubernetesCachingPolicy {
	if x != nil {
		return x.CachingPolicies
	}
	return nil
}

func (x *KubernetesAccount) GetDockerRegistries() []*KubernetesAccountDockerRegistry {
	if x != nil {
		return x.DockerRegistries
	}
	return nil
}

func (x *KubernetesAccount) GetOAuthScopes() []string {
	if x != nil {
		return x.OAuthScopes
	}
	return nil
}

func (x *KubernetesAccount) GetKubeconfigFile() string {
	if x != nil {
		return x.KubeconfigFile
	}
	return ""
}

func (x *KubernetesAccount) GetPermissions() *client.Permissions {
	if x != nil {
		return x.Permissions
	}
	return nil
}

func (x *KubernetesAccount) GetRequiredGroupMembership() []string {
	if x != nil {
		return x.RequiredGroupMembership
	}
	return nil
}

func (x *KubernetesAccount) GetLiveManifestCalls() *wrappers.BoolValue {
	if x != nil {
		return x.LiveManifestCalls
	}
	return nil
}

func (x *KubernetesAccount) GetServiceAccount() *wrappers.BoolValue {
	if x != nil {
		return x.ServiceAccount
	}
	return nil
}

func (x *KubernetesAccount) GetKubeconfigContents() string {
	if x != nil {
		return x.KubeconfigContents
	}
	return ""
}

func (x *KubernetesAccount) GetKubectlPath() string {
	if x != nil {
		return x.KubectlPath
	}
	return ""
}

func (x *KubernetesAccount) GetKubectlRequestTimeoutSeconds() int32 {
	if x != nil {
		return x.KubectlRequestTimeoutSeconds
	}
	return 0
}

func (x *KubernetesAccount) GetCheckPermissionsOnStartup() *wrappers.BoolValue {
	if x != nil {
		return x.CheckPermissionsOnStartup
	}
	return nil
}

func (x *KubernetesAccount) GetOAuthServiceAccount() string {
	if x != nil {
		return x.OAuthServiceAccount
	}
	return ""
}

func (x *KubernetesAccount) GetOnlySpinnakerManaged() *wrappers.BoolValue {
	if x != nil {
		return x.OnlySpinnakerManaged
	}
	return nil
}

func (x *KubernetesAccount) GetDebug() *wrappers.BoolValue {
	if x != nil {
		return x.Debug
	}
	return nil
}

// Configuration for a CRD to be managed by Spinnaker. If Spinnaker does not
// have permission to list CRDs but you need Spinnaker to manage CRDs, you
// need to explicitly register each CRD.
type KubernetesCustomResource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Kubernetes kind of the custom resource.
	KubernetesKind string `protobuf:"bytes,1,opt,name=kubernetesKind,proto3" json:"kubernetesKind,omitempty"`
	// The Spinnaker kind to which you would like the custom resource to
	// map.
	SpinnakerKind string `protobuf:"bytes,2,opt,name=spinnakerKind,proto3" json:"spinnakerKind,omitempty"`
	// An integer representing the deployment priority of this resource.
	// Resources with lower values are deployed before resources with higher
	// values.
	DeployPriority string `protobuf:"bytes,3,opt,name=deployPriority,proto3" json:"deployPriority,omitempty"`
	// Whether Spinnaker should manage versioning this resource.
	Versioned *wrappers.BoolValue `protobuf:"bytes,4,opt,name=versioned,proto3" json:"versioned,omitempty"`
	// Whether the resource is namespaced. Defaults to true.
	Namespaced *wrappers.BoolValue `protobuf:"bytes,5,opt,name=namespaced,proto3" json:"namespaced,omitempty"`
}

func (x *KubernetesCustomResource) Reset() {
	*x = KubernetesCustomResource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_kubernetes_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubernetesCustomResource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubernetesCustomResource) ProtoMessage() {}

func (x *KubernetesCustomResource) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_kubernetes_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubernetesCustomResource.ProtoReflect.Descriptor instead.
func (*KubernetesCustomResource) Descriptor() ([]byte, []int) {
	return file_cloudprovider_kubernetes_proto_rawDescGZIP(), []int{2}
}

func (x *KubernetesCustomResource) GetKubernetesKind() string {
	if x != nil {
		return x.KubernetesKind
	}
	return ""
}

func (x *KubernetesCustomResource) GetSpinnakerKind() string {
	if x != nil {
		return x.SpinnakerKind
	}
	return ""
}

func (x *KubernetesCustomResource) GetDeployPriority() string {
	if x != nil {
		return x.DeployPriority
	}
	return ""
}

func (x *KubernetesCustomResource) GetVersioned() *wrappers.BoolValue {
	if x != nil {
		return x.Versioned
	}
	return nil
}

func (x *KubernetesCustomResource) GetNamespaced() *wrappers.BoolValue {
	if x != nil {
		return x.Namespaced
	}
	return nil
}

// Configuration for a kind-specific caching policy.
type KubernetesCachingPolicy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Kubernetes kind to which the policy applies.
	KubernetesKind string `protobuf:"bytes,1,opt,name=kubernetesKind,proto3" json:"kubernetesKind,omitempty"`
	// The maximum number of resources an agent will cache of the specified
	// Kubernetes kind.
	MaxEntriesPerAgent int32 `protobuf:"varint,2,opt,name=maxEntriesPerAgent,proto3" json:"maxEntriesPerAgent,omitempty"`
}

func (x *KubernetesCachingPolicy) Reset() {
	*x = KubernetesCachingPolicy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_kubernetes_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubernetesCachingPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubernetesCachingPolicy) ProtoMessage() {}

func (x *KubernetesCachingPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_kubernetes_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubernetesCachingPolicy.ProtoReflect.Descriptor instead.
func (*KubernetesCachingPolicy) Descriptor() ([]byte, []int) {
	return file_cloudprovider_kubernetes_proto_rawDescGZIP(), []int{3}
}

func (x *KubernetesCachingPolicy) GetKubernetesKind() string {
	if x != nil {
		return x.KubernetesKind
	}
	return ""
}

func (x *KubernetesCachingPolicy) GetMaxEntriesPerAgent() int32 {
	if x != nil {
		return x.MaxEntriesPerAgent
	}
	return 0
}

// Configuration for a Docker registry.
type KubernetesAccountDockerRegistry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The configured name of the Docker registry.
	AccountName string `protobuf:"bytes,1,opt,name=accountName,proto3" json:"accountName,omitempty"`
	// The list of Docker registry namespaces usable as image sources.
	Namespaces []string `protobuf:"bytes,2,rep,name=namespaces,proto3" json:"namespaces,omitempty"`
}

func (x *KubernetesAccountDockerRegistry) Reset() {
	*x = KubernetesAccountDockerRegistry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_kubernetes_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KubernetesAccountDockerRegistry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubernetesAccountDockerRegistry) ProtoMessage() {}

func (x *KubernetesAccountDockerRegistry) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_kubernetes_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubernetesAccountDockerRegistry.ProtoReflect.Descriptor instead.
func (*KubernetesAccountDockerRegistry) Descriptor() ([]byte, []int) {
	return file_cloudprovider_kubernetes_proto_rawDescGZIP(), []int{4}
}

func (x *KubernetesAccountDockerRegistry) GetAccountName() string {
	if x != nil {
		return x.AccountName
	}
	return ""
}

func (x *KubernetesAccountDockerRegistry) GetNamespaces() []string {
	if x != nil {
		return x.Namespaces
	}
	return nil
}

var File_cloudprovider_kubernetes_proto protoreflect.FileDescriptor

var file_cloudprovider_kubernetes_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f,
	0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xae, 0x01, 0x0a, 0x0a, 0x4b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x12, 0x34, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x42, 0x0a,
	0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x26, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x73, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x72, 0x69, 0x6d, 0x61,
	0x72, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xe0, 0x09, 0x0a, 0x11, 0x4b, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6b, 0x69, 0x6e, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x05, 0x6b, 0x69, 0x6e, 0x64, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x6f, 0x6d, 0x69,
	0x74, 0x4b, 0x69, 0x6e, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x6d,
	0x69, 0x74, 0x4b, 0x69, 0x6e, 0x64, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78,
	0x74, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x61, 0x63, 0x68, 0x65, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x63, 0x61, 0x63, 0x68, 0x65, 0x54, 0x68,
	0x72, 0x65, 0x61, 0x64, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x6f, 0x6d, 0x69, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x6f,
	0x6d, 0x69, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x12, 0x57, 0x0a,
	0x0f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x4b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x0f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x56, 0x0a, 0x0f, 0x63, 0x61, 0x63, 0x68, 0x69, 0x6e,
	0x67, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x43, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x0f, 0x63,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x12, 0x60,
	0x0a, 0x10, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x4b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x44, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x52, 0x10,
	0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x12, 0x20, 0x0a, 0x0b, 0x6f, 0x41, 0x75, 0x74, 0x68, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x73, 0x18,
	0x0c, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x41, 0x75, 0x74, 0x68, 0x53, 0x63, 0x6f, 0x70,
	0x65, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x6b, 0x75, 0x62, 0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x46, 0x69, 0x6c, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6b, 0x75, 0x62, 0x65,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x34, 0x0a, 0x0b, 0x70, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x38, 0x0a, 0x17, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x18, 0x0f, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x17, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x12, 0x48, 0x0a, 0x11, 0x6c, 0x69,
	0x76, 0x65, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x43, 0x61, 0x6c, 0x6c, 0x73, 0x18,
	0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x11, 0x6c, 0x69, 0x76, 0x65, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x43,
	0x61, 0x6c, 0x6c, 0x73, 0x12, 0x42, 0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42,
	0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x12, 0x6b, 0x75, 0x62, 0x65,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x12,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x6b, 0x75, 0x62, 0x65, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x6b, 0x75, 0x62, 0x65,
	0x63, 0x74, 0x6c, 0x50, 0x61, 0x74, 0x68, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6b,
	0x75, 0x62, 0x65, 0x63, 0x74, 0x6c, 0x50, 0x61, 0x74, 0x68, 0x12, 0x42, 0x0a, 0x1c, 0x6b, 0x75,
	0x62, 0x65, 0x63, 0x74, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x16, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x1c, 0x6b, 0x75, 0x62, 0x65, 0x63, 0x74, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x58,
	0x0a, 0x19, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x4f, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x75, 0x70, 0x18, 0x17, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x19, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x4f,
	0x6e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x75, 0x70, 0x12, 0x30, 0x0a, 0x13, 0x6f, 0x41, 0x75, 0x74,
	0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x6f, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x4e, 0x0a, 0x14, 0x6f, 0x6e,
	0x6c, 0x79, 0x53, 0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x64, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x14, 0x6f, 0x6e, 0x6c, 0x79, 0x53, 0x70, 0x69, 0x6e, 0x6e, 0x61,
	0x6b, 0x65, 0x72, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x12, 0x30, 0x0a, 0x05, 0x64, 0x65,
	0x62, 0x75, 0x67, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x64, 0x65, 0x62, 0x75, 0x67, 0x22, 0x86, 0x02, 0x0a,
	0x18, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x6b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x4b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x4b, 0x69, 0x6e,
	0x64, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72, 0x4b, 0x69,
	0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x70, 0x69, 0x6e, 0x6e, 0x61,
	0x6b, 0x65, 0x72, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x50, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x12,
	0x38, 0x0a, 0x09, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x64, 0x12, 0x3a, 0x0a, 0x0a, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x64, 0x22, 0x71, 0x0a, 0x17, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x43, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x67, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x12, 0x26, 0x0a, 0x0e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x4b, 0x69,
	0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x4b, 0x69, 0x6e, 0x64, 0x12, 0x2e, 0x0a, 0x12, 0x6d, 0x61, 0x78, 0x45,
	0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x50, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x6d, 0x61, 0x78, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73,
	0x50, 0x65, 0x72, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x22, 0x63, 0x0a, 0x1f, 0x4b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x6f, 0x63,
	0x6b, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0a, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x73, 0x42, 0x35, 0x5a,
	0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x69, 0x6e,
	0x6e, 0x61, 0x6b, 0x65, 0x72, 0x2f, 0x6b, 0x6c, 0x65, 0x61, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloudprovider_kubernetes_proto_rawDescOnce sync.Once
	file_cloudprovider_kubernetes_proto_rawDescData = file_cloudprovider_kubernetes_proto_rawDesc
)

func file_cloudprovider_kubernetes_proto_rawDescGZIP() []byte {
	file_cloudprovider_kubernetes_proto_rawDescOnce.Do(func() {
		file_cloudprovider_kubernetes_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloudprovider_kubernetes_proto_rawDescData)
	})
	return file_cloudprovider_kubernetes_proto_rawDescData
}

var file_cloudprovider_kubernetes_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_cloudprovider_kubernetes_proto_goTypes = []interface{}{
	(*Kubernetes)(nil),                      // 0: proto.cloudprovider.Kubernetes
	(*KubernetesAccount)(nil),               // 1: proto.cloudprovider.KubernetesAccount
	(*KubernetesCustomResource)(nil),        // 2: proto.cloudprovider.KubernetesCustomResource
	(*KubernetesCachingPolicy)(nil),         // 3: proto.cloudprovider.KubernetesCachingPolicy
	(*KubernetesAccountDockerRegistry)(nil), // 4: proto.cloudprovider.KubernetesAccountDockerRegistry
	(*wrappers.BoolValue)(nil),              // 5: google.protobuf.BoolValue
	(*client.Permissions)(nil),              // 6: proto.Permissions
}
var file_cloudprovider_kubernetes_proto_depIdxs = []int32{
	5,  // 0: proto.cloudprovider.Kubernetes.enabled:type_name -> google.protobuf.BoolValue
	1,  // 1: proto.cloudprovider.Kubernetes.accounts:type_name -> proto.cloudprovider.KubernetesAccount
	2,  // 2: proto.cloudprovider.KubernetesAccount.customResources:type_name -> proto.cloudprovider.KubernetesCustomResource
	3,  // 3: proto.cloudprovider.KubernetesAccount.cachingPolicies:type_name -> proto.cloudprovider.KubernetesCachingPolicy
	4,  // 4: proto.cloudprovider.KubernetesAccount.dockerRegistries:type_name -> proto.cloudprovider.KubernetesAccountDockerRegistry
	6,  // 5: proto.cloudprovider.KubernetesAccount.permissions:type_name -> proto.Permissions
	5,  // 6: proto.cloudprovider.KubernetesAccount.liveManifestCalls:type_name -> google.protobuf.BoolValue
	5,  // 7: proto.cloudprovider.KubernetesAccount.serviceAccount:type_name -> google.protobuf.BoolValue
	5,  // 8: proto.cloudprovider.KubernetesAccount.checkPermissionsOnStartup:type_name -> google.protobuf.BoolValue
	5,  // 9: proto.cloudprovider.KubernetesAccount.onlySpinnakerManaged:type_name -> google.protobuf.BoolValue
	5,  // 10: proto.cloudprovider.KubernetesAccount.debug:type_name -> google.protobuf.BoolValue
	5,  // 11: proto.cloudprovider.KubernetesCustomResource.versioned:type_name -> google.protobuf.BoolValue
	5,  // 12: proto.cloudprovider.KubernetesCustomResource.namespaced:type_name -> google.protobuf.BoolValue
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_cloudprovider_kubernetes_proto_init() }
func file_cloudprovider_kubernetes_proto_init() {
	if File_cloudprovider_kubernetes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloudprovider_kubernetes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Kubernetes); i {
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
		file_cloudprovider_kubernetes_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubernetesAccount); i {
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
		file_cloudprovider_kubernetes_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubernetesCustomResource); i {
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
		file_cloudprovider_kubernetes_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubernetesCachingPolicy); i {
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
		file_cloudprovider_kubernetes_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KubernetesAccountDockerRegistry); i {
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
			RawDescriptor: file_cloudprovider_kubernetes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloudprovider_kubernetes_proto_goTypes,
		DependencyIndexes: file_cloudprovider_kubernetes_proto_depIdxs,
		MessageInfos:      file_cloudprovider_kubernetes_proto_msgTypes,
	}.Build()
	File_cloudprovider_kubernetes_proto = out.File
	file_cloudprovider_kubernetes_proto_rawDesc = nil
	file_cloudprovider_kubernetes_proto_goTypes = nil
	file_cloudprovider_kubernetes_proto_depIdxs = nil
}
