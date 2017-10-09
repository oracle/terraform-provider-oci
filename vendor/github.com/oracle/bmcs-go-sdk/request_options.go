// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

// Options
// To get the body, optional and required are marshalled and merged.
// To get the query string, optional and required are merged.
// To get the header, optional and required are merged.
// Required options get built inline within exported functions, based on
// function parameters.
// A single options struct gets passed to exported functions, representing
// optional params.
// Both required and optional fields need to explicitly tag header, json and
// url, excluding appropriately.

type IfMatchOptions struct {
	IfMatch string `header:"If-Match,omitempty" json:"-" url:"-"`
}

type IfNoneMatchOptions struct {
	IfNoneMatch string `header:"If-None-Match,omitempty" json:"-" url:"-"`
}

type RetryTokenOptions struct {
	RetryToken string `header:"opc-retry-token,omitempty" json:"-" url:"-"`
}

type HeaderOptions struct {
	IfMatchOptions
	RetryTokenOptions
}

type ClientRequestOptions struct {
	OPCClientRequestID string `header:"opc-client-request-id,omitempty" json:"-" url:"-"`
}

type DisplayNameOptions struct {
	DisplayName string `header:"-" json:"displayName,omitempty" url:"-"`
}

type VersionDateOptions struct {
	VersionDate string `header:"-" json:"versionDate,omitempty" url:"-"`
}

// Creation Options

type CreateOptions struct {
	RetryTokenOptions
	DisplayNameOptions
}

type CreateBucketOptions struct {
	Metadata   map[string]string `header:"-" json:"metadata,omitempty" url:"-"`
	AccessType BucketAccessType  `header:"-" json:"publicAccessType,omitempty" url:"-"`
}

type CreatePreauthenticatedRequestDetails struct {
	ClientRequestOptions
	Name        string        `header:"-" json:"name" url:"-"`
	ObjectName  string        `header:"-" json:"objectName,omitempty" url:"-"`
	AccessType  PARAccessType `header:"-" json:"accessType" url:"-"`
	TimeExpires Time          `header:"-" json:"timeExpires" url:"-"`
}

type CreatePrivateIPOptions struct {
	CreateOptions
	HostnameLabel string `header:"-" json:"hostnameLabel,omitempty" url:"-"`
	IPAddress     string `header:"-" json:"ipAddress,omitempty" url:"-"`
}

type CreateVcnOptions struct {
	CreateOptions
	DnsLabel string `header:"-" json:"dnsLabel,omitempty" url:"-"`
}

type CreateSubnetOptions struct {
	CreateOptions
	DHCPOptionsID          string   `header:"-" json:"dhcpOptionsId,omitempty" url:"-"`
	DNSLabel               string   `header:"-" json:"dnsLabel,omitempty" url:"-"`
	ProhibitPublicIpOnVnic bool     `header:"-" json:"prohibitPublicIpOnVnic,omitempty" url:"-"`
	RouteTableID           string   `header:"-" json:"routeTableId,omitempty" url:"-"`
	SecurityListIDs        []string `header:"-" json:"securityListIds,omitempty" url:"-"`
}

type AttachVnicOptions struct {
	CreateOptions
}

type LoadBalancerOptions struct {
	ClientRequestOptions
	RetryTokenOptions
}

type CreateLoadBalancerOptions struct {
	LoadBalancerOptions
	DisplayNameOptions
	IsPrivate bool `header:"-" json:"isPrivate,omitempty" url:"-"`
}

type CreateLoadBalancerBackendOptions struct {
	LoadBalancerOptions
	Backup  bool `header:"-" json:"backup,omitempty" url:"-"`
	Drain   bool `header:"-" json:"drain,omitempty" url:"-"`
	Offline bool `header:"-" json:"offline,omitempty" url:"-"`
	Weight  int  `header:"-" json:"weight,omitempty" url:"-"`
}

type UpdateLoadBalancerBackendOptions struct {
	LoadBalancerOptions
	Backup  bool `header:"-" json:"backup" url:"-"`
	Drain   bool `header:"-" json:"drain" url:"-"`
	Offline bool `header:"-" json:"offline" url:"-"`
	Weight  int  `header:"-" json:"weight" url:"-"`
}

type UpdateLoadBalancerBackendSetOptions struct {
	LoadBalancerOptions
	RetryTokenOptions
	Backends      []Backend         `header:"-" json:"backends" url:"-"`
	HealthChecker *HealthChecker    `header:"-" json:"healthChecker,omitempty" url:"-"`
	Policy        string            `header:"-" json:"policy,omitempty" url:"-"`
	SSLConfig     *SSLConfiguration `header:"-" json:"sslConfiguration,omitempty" url:"-"`
}

type UpdateLoadBalancerListenerOptions struct {
	LoadBalancerOptions
	DefaultBackendSetName string            `header:"-" json:"defaultBackendSetName" url:"-"`
	Port                  int               `header:"-" json:"port" url:"-"`
	Protocol              string            `header:"-" json:"protocol" url:"-"`
	SSLConfig             *SSLConfiguration `header:"-" json:"sslConfiguration,omitempty" url:"-"`
}

type ListLoadBalancerPolicyOptions struct {
	ClientRequestOptions
	ListOptions
}

type ListLoadBalancerOptions struct {
	ClientRequestOptions
	ListOptions
	CompartmentID string `header:"-" json:"-" url:"compartmentId,omitempty"`
	Detail        string `header:"-" json:"-" url:"detail,omitempty"`
}

type UpdateLoadBalancerOptions struct {
	LoadBalancerOptions
	DisplayNameOptions
}

type UpdatePrivateIPOptions struct {
	UpdateOptions
	HostnameLabel string `header:"-" json:"hostnameLabel,omitempty" url:"-"`
	VnicID        string `header:"-" json:"vnicId,omitempty" url:"-"`
}

type UpdateVnicOptions struct {
	UpdateOptions
	HostnameLabel       string `header:"-" json:"hostnameLabel,omitempty" url:"-"`
	SkipSourceDestCheck *bool  `header:"-" json:"skipSourceDestCheck,omitempty" url:"-"`
}

type CreateVolumeOptions struct {
	CreateOptions
	SizeInMBs      int    `header:"-" json:"sizeInMBs,omitempty" url:"-"`
	SizeInGBs      int    `header:"-" json:"sizeInGBs,omitempty" url:"-"`
	VolumeBackupID string `header:"-" json:"volumeBackupId,omitempty" url:"-"`
}

type CreatePolicyOptions struct {
	RetryTokenOptions
	VersionDateOptions
}

type CreateVnicOptions struct {
	AssignPublicIp      *bool  `header:"-" json:"assignPublicIp,omitempty" url:"-"`
	DisplayName         string `header:"-" json:"displayName,omitempty" url:"-"`
	HostnameLabel       string `header:"-" json:"hostnameLabel,omitempty" url:"-"`
	PrivateIp           string `header:"-" json:"privateIp,omitempty" url:"-"`
	SkipSourceDestCheck *bool  `header:"-" json:"skipSourceDestCheck,omitempty" url:"-"`
	SubnetID            string `header:"-" json:"subnetId,omitempty" url:"-"`
}

type LaunchInstanceOptions struct {
	CreateOptions
	CreateVnicOptions *CreateVnicOptions     `header:"-" json:"createVnicDetails,omitempty" url:"-"`
	HostnameLabel     string                 `header:"-" json:"hostnameLabel,omitempty" url:"-"`
	IpxeScript        string                 `header:"-" json:"ipxeScript,omitempty" url:"-"`
	Metadata          map[string]string      `header:"-" json:"metadata,omitempty" url:"-"`
	ExtendedMetadata  map[string]interface{} `header:"-" json:"extendedMetadata,omitempty" url:"-"`
}

type LaunchDBSystemOptions struct {
	CreateOptions
	BackupSubnetId        string         `header:"-" json:"backupSubnetId,omitempty" url:"-"`
	ClusterName           string         `header:"-" json:"clusterName,omitempty" url:"-"`
	DataStoragePercentage int            `header:"-" json:"dataStoragePercentage,omitempty" url:"-"`
	DiskRedundancy        DiskRedundancy `header:"-" json:"diskRedundancy,omitempty" url:"-"`
	Domain                string         `header:"-" json:"domain,omitempty" url:"-"`
}

type CreateDBHomeOptions struct {
	DisplayNameOptions
}

type CreateDatabaseOptions struct {
	CharacterSet  string
	NCharacterSet string
	DBWorkload    string
	PDBName       string
}

// Read Options

type GetObjectOptions struct {
	IfMatchOptions
	IfNoneMatchOptions
	ClientRequestOptions
	Range string `header:"Range,omitempty" json:"-" url:"-"`
}

// Update Options

type UpdateOptions struct {
	HeaderOptions
	DisplayNameOptions
}

type IfMatchDisplayNameOptions struct {
	IfMatchOptions
	DisplayNameOptions
}

type UpdateBucketOptions struct {
	IfMatchOptions
	Name       string            `header:"-" json:"name,omitempty" url:"-"`
	Namespace  Namespace         `header:"-" json:"namespace,omitempty" url:"-"`
	AccessType BucketAccessType  `header:"-" json:"publicAccessType,omitempty" url:"-"`
	Metadata   map[string]string `header:"-" json:"metadata,omitempty" url:"-"`
}

type UpdateIdentityOptions struct {
	IfMatchOptions
	Description string `header:"-" json:"description,omitempty" url:"-"`
}

type UpdateCompartmentOptions struct {
	UpdateIdentityOptions
	Name string `header:"-" json:"name,omitempty" url:"-"`
}

type UpdateUserStateOptions struct {
	IfMatchOptions
	Blocked *bool `header:"-" json:"blocked,omitempty" url:"-"`
}

type UpdatePolicyOptions struct {
	UpdateIdentityOptions
	VersionDateOptions
	Statements []string `header:"-" json:"statements,omitempty" url:"-"`
}

type UpdateDHCPDNSOptions struct {
	CreateOptions
	Options []DHCPDNSOption `header:"-" json:"options,omitempty" url:"-"`
}

type UpdateGatewayOptions struct {
	IfMatchOptions
	DisplayNameOptions
	IsEnabled *bool `header:"-" json:"isEnabled,omitempty" url:"-"`
}

type UpdateRouteTableOptions struct {
	CreateOptions
	RouteRules []RouteRule `header:"-" json:"routeRules,omitempty" url:"-"`
}

type UpdateSecurityListOptions struct {
	IfMatchDisplayNameOptions
	EgressRules  []EgressSecurityRule  `header:"-" json:"egressSecurityRules" url:"-"`
	IngressRules []IngressSecurityRule `header:"-" json:"ingressSecurityRules" url:"-"`
}

type PutObjectOptions struct {
	IfMatchOptions
	IfNoneMatchOptions
	ClientRequestOptions
	MetadataUnmarshaller
	Expect          string `header:"Expect,omitempty" json:"-" url:"-"`
	ContentMD5      string `header:"Content-MD5,omitempty" json:"-" url:"-"`
	ContentType     string `header:"Content-Type,omitempty" json:"-" url:"-"`
	ContentLanguage string `header:"Content-Language,omitempty" json:"-" url:"-"`
	ContentEncoding string `header:"Content-Encoding,omitempty" json:"-" url:"-"`
}

// Delete Options

type DeleteObjectOptions struct {
	IfMatchOptions
	ClientRequestOptions
}

// List Options

type PageListOptions struct {
	Page string `header:"-" json:"-" url:"page,omitempty"`
}

type LimitListOptions struct {
	Limit uint64 `header:"-" json:"-" url:"limit,omitempty"`
}

type ListOptions struct {
	LimitListOptions
	PageListOptions
}

type DisplayNameListOptions struct {
	DisplayName string `header:"-" json:"-" url:"displayName,omitempty"`
}

type AvailabilityDomainListOptions struct {
	AvailabilityDomain string `header:"-" json:"-" url:"availabilityDomain,omitempty"`
}

type DrgIDListOptions struct {
	DrgID string `header:"-" json:"-" url:"drgId,omitempty"`
}

type InstanceIDListOptions struct {
	InstanceID string `header:"-" json:"-" url:"instanceId,omitempty"`
}

type ListInstancesOptions struct {
	AvailabilityDomainListOptions
	DisplayNameListOptions
	ListOptions
}

type ListConsoleHistoriesOptions struct {
	AvailabilityDomainListOptions
	InstanceIDListOptions
	ListOptions
}

type ListDrgAttachmentsOptions struct {
	DrgIDListOptions
	ListOptions
	VcnID string `header:"-" json:"-" url:"vcnId,omitempty"`
}

type ListImagesOptions struct {
	DisplayNameListOptions
	ListOptions
	OperatingSystem        string `header:"-" json:"-" url:"operatingSystem,omitempty"`
	OperatingSystemVersion string `header:"-" json:"-" url:"operatingSystemVersion,omitempty"`
}

type ListIPSecConnsOptions struct {
	DrgIDListOptions
	ListOptions
	CpeID string `header:"-" json:"-" url:"cpeId,omitempty"`
}

type ListPrivateIPsOptions struct {
	ListOptions
	IPAddress string `header:"-" json:"-" url:"ipAddress,omitempty"`
	SubnetID  string `header:"-" json:"-" url:"subnetId,omitempty"`
	VnicID    string `header:"-" json:"-" url:"vnicId,omitempty"`
}

type ListShapesOptions struct {
	AvailabilityDomainListOptions
	ListOptions
	ImageID string `header:"-" json:"-" url:"imageId,omitempty"`
}

type ListVnicAttachmentsOptions struct {
	AvailabilityDomainListOptions
	InstanceIDListOptions
	ListOptions
	VnicID string `header:"-" json:"-" url:"vnicId,omitempty"`
}

type ListVolumesOptions struct {
	AvailabilityDomainListOptions
	ListOptions
}

type ListVolumeAttachmentsOptions struct {
	AvailabilityDomainListOptions
	InstanceIDListOptions
	ListOptions
	VolumeID string `header:"-" json:"-" url:"volumeId,omitempty"`
}

type ListBackupsOptions struct {
	ListOptions
	VolumeID string `header:"-" json:"-" url:"volumeId,omitempty"`
}

type ListMembershipsOptions struct {
	ListOptions
	GroupID string `header:"-" json:"-" url:"groupId,omitempty"`
	UserID  string `header:"-" json:"-" url:"userId,omitempty"`
}

type ListBucketsOptions struct {
	ListOptions
	ClientRequestOptions
}

type ListPreauthenticatedRequestOptions struct {
	ListOptions
	ClientRequestOptions
	ObjectNamePrefix string `header:"-" json:"-" url:"objectNamePrefix,omitempty"`
}

type ListObjectsOptions struct {
	ClientRequestOptions
	LimitListOptions
	Prefix    string `header:"-" json:"-" url:"prefix,omitempty"`
	Start     string `header:"-" json:"-" url:"start,omitempty"`
	End       string `header:"-" json:"-" url:"end,omitempty"`
	Delimiter string `header:"-" json:"-" url:"delimiter,omitempty"`
	Fields    string `header:"-" json:"-" url:"fields,omitempty"`
}

// Misc Options

type HeadObjectOptions struct {
	IfMatchOptions
	IfNoneMatchOptions
	ClientRequestOptions
}

type ConsoleHistoryDataOptions struct {
	Length uint64 `header:"-" json:"-" url:"length,omitempty"`
	Offset uint64 `header:"-" json:"-" url:"offset,omitempty"`
}
