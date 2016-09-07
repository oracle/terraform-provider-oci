package baremetal

// Options is used to pass optional values to to package methods. Note that zero-value
// fields will be ignored. Typically options are passed on a variadic paramter
// to SDK methods. Note that only the first option, if present, will be used.
type Options struct {
	AvailabilityDomain     string
	DisplayName            string
	ImageID                string
	InstanceID             string
	VolumeID               string
	VnicID                 string
	VcnID                  string
	OperatingSystem        string
	OperatingSystemVersion string
	// IfMatch (Optional) is for optimistic concurrency control. In the PUT or DELETE call for
	// a resource, set the if-match parameter to the value of the etag from a previous GET or POST response
	// for that resource. The resource will be updated or deleted only if the etag you
	// provide matches the resource's current etag value.
	IfMatch string
	// OPCIdempotencyToken (Optional) - A token you supply to uniquely identify the request and provide idempotency
	//if the request is retried. Idempotency tokens expire after 30 minutes.
	OPCIdempotencyToken string

	// ListOptions contains arguments to support pagination for List requests. ListOptions
	// is typically a variadic paramter to List SDK functions.  Only the first ListOption will
	// be used.  If multiple ListOptions are passed the subsequent values after the first
	// will be discarded.
	UserID  string
	GroupID string
	// Page the value of OPCNextPage from ListUsersResponse used for
	// paging results.
	Page string
	// Limit he maximum number of results that ListUsers is to return.
	Limit uint64
	// DrgID (Optional) the ID of the VPN headend
	DrgID string
	// CpeID (Optional) the ID of the customer premise equipment
	CpeID string
	// Length (Optional) max number of bytes to return in console history
	Length int
	// Offset (Optional) offset of console history to return.
	Offset int
}
