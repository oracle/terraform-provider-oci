package baremetal

import (
	"net/http"
)

// Requirements

type identityCreationRequirement struct {
	CompartmentID string `json:"compartmentId" url:"-"`
	Description   string `json:"description" url:"-"`
	Name          string `json:"name" url:"-"`
}

type ocidRequirement struct {
	CompartmentID string `json:"compartmentId" url:"-"`
}

type listOCIDRequirement struct {
	CompartmentID string `json:"-" url:"compartmentId"`
}

// Options

// To get the body, optional and required are marshalled and merged.
// To get the query string, optional and required are merged.
// To get the header,
//   optional is type asserted as a HeaderGenerator
//   required is type asserted as a HeaderGenerator
//   both results are combined
// Required options get built inline within funcs based on args.
// A single options struct gets passed in as optional.
// Both need to explicitly handle json and url tags, excluding appropriately.

type requestOptions interface {
	url(urlBuilderFn) (val string, e error)
	header() http.Header
	getBody() ([]byte, error)
}

type HeaderGenerator interface {
	Header() http.Header
}

type IfMatchOptions struct {
	IfMatch string `json:"-" url:"-"`
}

func (opt *IfMatchOptions) Header() http.Header {
	header := http.Header{}
	if opt != nil && opt.IfMatch != "" {
		header.Set(headerIfMatch, opt.IfMatch)
	}
	return header
}

type IfNoneMatchOptions struct {
	IfNoneMatch string `json:"-" url:"-"`
}

func (opt *IfNoneMatchOptions) Header() http.Header {
	header := http.Header{}
	if opt != nil && opt.IfNoneMatch != "" {
		header.Set(headerIfNoneMatch, opt.IfNoneMatch)
	}
	return header
}

type RetryTokenOptions struct {
	RetryToken string `json:"-" url:"-"`
}

func (opt *RetryTokenOptions) Header() http.Header {
	header := http.Header{}
	if opt != nil && opt.RetryToken != "" {
		header.Set(headerRetryToken, opt.RetryToken)
	}
	return header
}

// Both header options cannot be embedded into the same struct separately.
// Duplicating here keeps the API clean, allowing this option to be embedded in
// non-header structs with easy-to-set struct values.
type HeaderOptions struct {
	IfMatch    string `json:"-" url:"-"`
	RetryToken string `json:"-" url:"-"`
}

func (opt *HeaderOptions) Header() http.Header {
	header := http.Header{}
	if opt != nil {
		if opt.IfMatch != "" {
			header.Set(headerIfMatch, opt.IfMatch)
		}
		if opt.RetryToken != "" {
			header.Set(headerRetryToken, opt.RetryToken)
		}
	}
	return header
}

type ClientRequestOptions struct {
	OPCClientRequestID string
}

func (ref *ClientRequestOptions) Header() http.Header {
	h := http.Header{}
	if ref.OPCClientRequestID != "" {
		h.Add(headerOPCClientRequestID, ref.OPCClientRequestID)
	}
	return h
}

type DisplayNameOptions struct {
	DisplayName string `json:"displayName,omitempty" url:"-"`
}

type VersionDateOptions struct {
	VersionDate string `json:"versionDate,omitempty" url:"-"`
}

// Creation Options

type CreateOptions struct {
	RetryTokenOptions
	DisplayNameOptions
}

type CreateBucketOptions struct {
	Metadata map[string]string `json:"metadata,omitempty" url:"-"`
}

type CreateSubnetOptions struct {
	CreateOptions
	DHCPOptionsID   string   `json:"dhcpOptionsId,omitempty" url:"-"`
	RouteTableID    string   `json:"routeTableId,omitempty" url:"-"`
	SecurityListIDs []string `json:"securityListIds,omitempty" url:"-"`
}

type CreateVolumeOptions struct {
	CreateOptions
	VolumeBackupID string `json:"volumeBackupId,omitempty" url:"-"`
}

type CreatePolicyOptions struct {
	RetryTokenOptions
	VersionDateOptions
}

type LaunchInstanceOptions struct {
	CreateOptions
	Metadata map[string]string `json:"metadata,omitempty" url:"-"`
}

type LaunchDBSystemOptions struct {
	CreateOptions
	DatabaseEdition DatabaseEdition     `json:"databaseEdition,omitempty" url:"-"`
	DBHome          createDBHomeDetails `json:"dbHome,omitempty" url:"-"`
	DiskRedundancy  DiskRedundancy      `json:"diskRedundancy,omitempty" url:"-"`
	Domain          string              `json:"domain,omitempty" url:"-"`
	Hostname        string              `json:"hostname,omitempty" url:"-"`
}

// Update Options

type UpdateOptions struct {
	HeaderOptions
	DisplayNameOptions
}

type UpdateBackupOptions struct {
	IfMatchOptions
	DisplayNameOptions
}

type UpdateBucketOptions struct {
	Name      string            `json:"name,omitempty" url:"-"`
	Namespace Namespace         `json:"namespace,omitempty" url:"-"`
	Metadata  map[string]string `json:"metadata,omitempty" url:"-"`
}

type UpdateIdentityOptions struct {
	IfMatchOptions
	Description string `json:"description,omitempty" url:"-"`
}

type UpdatePolicyOptions struct {
	UpdateIdentityOptions
	VersionDateOptions
	Statements []string `json:"statements,omitempty" url:"-"`
}

type UpdateDHCPDNSOptions struct {
	CreateOptions
	Options []DHCPDNSOption `json:"options,omitempty" url:"-"`
}

type UpdateGatewayOptions struct {
	IfMatchOptions
	DisplayNameOptions
	IsEnabled bool `json:"isEnabled,omitempty" url:"-"`
}

type UpdateRouteTableOptions struct {
	CreateOptions
	RouteRules []RouteRule `json:"routeRules,omitempty" url:"-"`
}

type UpdateSecurityListOptions struct {
	CreateOptions
	EgressRules  []EgressSecurityRule  `json:"egressSecurityRules,omitempty" url:"-"`
	IngressRules []IngressSecurityRule `json:"ingressSecurityRules,omitempty" url:"-"`
}

// List Options

type PageListOptions struct {
	Page string `json:"-" url:"page,omitempty"`
}

type LimitListOptions struct {
	Limit uint64 `json:"-" url:"limit,omitempty"`
}

type ListOptions struct {
	LimitListOptions
	PageListOptions
}

type DisplayNameListOptions struct {
	DisplayName string `json:"-" url:"displayName,omitempty"`
}

type AvailabilityDomainListOptions struct {
	AvailabilityDomain string `json:"-" url:"availabilityDomain,omitempty"`
}

type DrgIDListOptions struct {
	DrgID string `json:"-" url:"drgId,omitempty"`
}

type InstanceIDListOptions struct {
	InstanceID string `json:"-" url:"instanceId,omitempty"`
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
	VcnID string `json:"-" url:"vcnId,omitempty"`
}

type ListImagesOptions struct {
	DisplayNameListOptions
	ListOptions
	OperatingSystem        string `json:"-" url:"operatingSystem,omitempty"`
	OperatingSystemVersion string `json:"-" url:"operatingSystemVersion,omitempty"`
}

type ListIPSecConnsOptions struct {
	DrgIDListOptions
	ListOptions
	CpeID string `json:"-" url:"cpeId,omitempty"`
}

type ListShapesOptions struct {
	AvailabilityDomainListOptions
	ListOptions
	ImageID string `json:"-" url:"imageId,omitempty"`
}

type ListVnicAttachmentsOptions struct {
	AvailabilityDomainListOptions
	InstanceIDListOptions
	ListOptions
	VnicID string `json:"-" url:"vnicId,omitempty"`
}

type ListVolumesOptions struct {
	AvailabilityDomainListOptions
	ListOptions
}

type ListVolumeAttachmentsOptions struct {
	AvailabilityDomainListOptions
	InstanceIDListOptions
	ListOptions
	VolumeID string `json:"-" url:"volumeId,omitempty"`
}

type ListBackupsOptions struct {
	ListOptions
	VolumeID string `json:"-" url:"volumeId,omitempty"`
}

type ListMembershipsOptions struct {
	ListOptions
	GroupID string `json:"-" url:"groupId,omitempty"`
	UserID  string `json:"-" url:"userId,omitempty"`
}

// Misc Options

type ConsoleHistoryDataOptions struct {
	Length uint64 `json:"-" url:"length,omitempty"`
	Offset uint64 `json:"-" url:"offset,omitempty"`
}

// -------- Objects API ------
type ListBucketsOptions struct {
	ListOptions
	ClientRequestOptions
}

type ListObjectsOptions struct {
	ClientRequestOptions
	Prefix    string `json:"-" url:"prefix"`
	Start     string `json:"-" url:"start"`
	End       string `json:"-" url:"end"`
	Limit     string `json:"-" url:"limit"`
	Delimiter string `json:"-" url:"delimiter"`
	Fields    string `json:"-" url:"fields"`
}

type DeleteObjectOptions struct {
	IfMatchOptions
	ClientRequestableResource
}

type PutObjectOptions struct {
	IfMatchOptions
	IfNoneMatchOptions
	ClientRequestableResource
	ContentResource
	Expect string
	MetadataResource
}

func (opt *PutObjectOptions) Header() http.Header {
	header := http.Header{}
	if opt != nil {
		if opt.IfMatch != "" {
			header.Set(headerIfMatch, opt.IfMatch)
		}
		if opt.IfNoneMatch != "" {
			header.Set(headerIfNoneMatch, opt.IfMatch)
		}
		if opt.ClientRequestID != "" {
			header.Set(headerOPCClientRequestID, opt.ClientRequestID)
		}
		if opt.Expect != "" {
			header.Set(headerExpect, opt.Expect)
		}
		if opt.ContentLength != 0 {
			header.Set(headerContentLength, string(opt.ContentLength))
		}
		if opt.MD5 != "" {
			header.Set(headerContentMD5, opt.MD5)
		}
		if opt.ContentType != "" {
			header.Set(headerContentType, opt.ContentType)
		}
		if opt.ContentLanguage != "" {
			header.Set(headerContentLanguage, opt.ContentLanguage)
		}
		if opt.ContentEncoding != "" {
			header.Set(headerContentEncoding, opt.ContentEncoding)
		}
	}
	return header
}
