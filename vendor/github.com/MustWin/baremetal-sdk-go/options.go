package baremetal

import (
	"net/http"
	"net/url"
	"strconv"
)

type urlParts []interface{}

// Options is used to pass optional values to to package methods. Note that zero-value
// fields will be ignored. Typically options are passed on a variadic paramter
// to SDK methods. Note that only the first option, if present, will be used.
// IfMatch (Optional) is for optimistic concurrency control. In the PUT or DELETE call for
// a resource, set the if-match parameter to the value of the etag from a previous GET or POST response
// for that resource. The resource will be updated or deleted only if the etag you
// provide matches the resource's current etag value.
// OPCIdempotencyToken (Optional) - A token you supply to uniquely identify the request and provide idempotency
// if the request is retried. Idempotency tokens expire after 30 minutes.
// ListOptions contains arguments to support pagination for List requests. ListOptions
// is typically a variadic paramter to List SDK functions.  Only the first ListOption will
// be used.  If multiple ListOptions are passed the subsequent values after the first
// will be discarded.
// Page the value of OPCNextPage from ListUsersResponse used for
// paging results.
// Limit he maximum number of results that ListUsers is to return.
// DrgID (Optional) the ID of the VPN headend
// CpeID (Optional) the ID of the customer premise equipment
// Length (Optional) max number of bytes to return in console history
// Offset (Optional) offset of console history to return.
type Options struct {
	AvailabilityDomain     string
	CpeID                  string
	DHCPOptions            []DHCPDNSOption
	DisplayName            string
	DrgID                  string
	GroupID                string
	IfMatch                string
	ImageID                string
	InstanceID             string
	Length                 int
	Limit                  uint64
	Offset                 int
	OPCIdempotencyToken    string
	OperatingSystem        string
	OperatingSystemVersion string
	Page                   string
	UserID                 string
	VcnID                  string
	VnicID                 string
	VolumeID               string
}

type requestOptions interface {
	header() http.Header
	getBody() interface{}
	url(urlBuilderFn) string
}

type sdkRequestOptions struct {
	body       interface{}
	name       resourceName
	options    []Options
	ocid       string
	ids        urlParts
	query      url.Values
	httpHeader http.Header
}

func (r *sdkRequestOptions) url(b urlBuilderFn) string {
	r.parseOptions()
	return b(r.name, r.query, r.ids...)
}

func (r *sdkRequestOptions) getBody() interface{} {
	r.parseOptions()
	return r.body
}

func (r *sdkRequestOptions) header() http.Header {
	r.parseOptions()
	return r.httpHeader
}

func (r *sdkRequestOptions) parseOptions() {
	if r.query == nil {
		r.query = url.Values{}
	}

	r.httpHeader = http.Header{}

	// Parse query options.
	if r.ocid != "" {
		r.query.Set(queryCompartmentID, r.ocid)
	}

	if len(r.options) > 0 {
		option := r.options[0]

		if option.AvailabilityDomain != "" {
			r.query.Set(queryAvailabilityDomain, option.AvailabilityDomain)
		}
		if option.CpeID != "" {
			r.query.Set(queryCpeID, option.CpeID)
		}
		if option.DrgID != "" {
			r.query.Set(queryDrgID, option.DrgID)
		}
		if option.GroupID != "" {
			r.query.Set(queryGroupID, option.GroupID)
		}
		if option.IfMatch != "" {
			r.httpHeader.Set(headerIfMatch, option.IfMatch)
		}
		if option.ImageID != "" {
			r.query.Set(queryImageID, option.ImageID)
		}
		if option.InstanceID != "" {
			r.query.Set(queryInstanceID, option.InstanceID)
		}
		if option.Length > 0 {
			r.query.Set(queryLength, strconv.Itoa(option.Length))
		}
		if option.Limit > 0 {
			r.query.Set(queryLimit, strconv.FormatUint(option.Limit, 10))
		}
		if option.Offset > 0 {
			r.query.Set(queryOffset, strconv.Itoa(option.Offset))
		}
		if option.OPCIdempotencyToken != "" {
			r.httpHeader.Set(headerRetryToken, option.OPCIdempotencyToken)
		}
		if option.OperatingSystem != "" {
			r.query.Set(queryOperatingSystem, option.OperatingSystem)
		}
		if option.OperatingSystemVersion != "" {
			r.query.Set(queryOperatingSystemVersion, option.OperatingSystemVersion)
		}
		if option.Page != "" {
			r.query.Set(queryPage, option.Page)
		}
		if option.UserID != "" {
			r.query.Set(queryUserID, option.UserID)
		}
		if option.VcnID != "" {
			r.query.Set(queryVcnID, option.VcnID)
		}
		if option.VnicID != "" {
			r.query.Set(queryVnicID, option.VnicID)
		}
		if option.VolumeID != "" {
			r.query.Set(queryVolumeID, option.VolumeID)
		}
	}
}
