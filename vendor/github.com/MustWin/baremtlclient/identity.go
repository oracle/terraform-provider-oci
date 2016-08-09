package baremtlsdk

import (
	"fmt"
	"time"
)

// Options is used to pass optional values to to package methods. Note that zero-value
// fields will be ignored. Typically options are passed on a variadic paramter
// to SDK methods. Note that only the first option, if present, will be used.
type Options struct {
	// OPCIdempotencyToken (Optional) - A token you supply to uniquely identify the request and provide idempotency
	//if the request is retried. Idempotency tokens expire after 30 minutes.
	OPCIdempotencyToken string
	// IfMatch (Optional) is for optimistic concurrency control. In the PUT or DELETE call for
	// a resource, set the if-match parameter to the value of the etag from a previous GET or POST response
	// for that resource. The resource will be updated or deleted only if the etag you
	// provide matches the resource's current etag value.
	IfMatch string
}

// Resource contains information representing Users, Groups,
// Policies and other elements
type IdentityResource struct {
	// Unique identifier for a particular item such as a User or a Group
	ID string `json:"id"`
	// CompartmentID is the ID of the tenancy containing the compartment
	CompartmentID string    `json:"compartmentId"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	TimeCreated   time.Time `json:"timeCreated"`
	TimeModified  time.Time `json:"timeModified"`
	State         string    `json:"state"`
	ETag          string    `json:"etag,omitempty"`
	OPCRequestID  string    `json:"opc-request-id,omitempty"`
}

// ListOptions contains arguments to support pagination for List requests. ListOptions
// is typically a variadic paramter to List SDK functions.  Only the first ListOption will
// be used.  If multiple ListOptions are passed the subsequent values after the first
// will be discarded.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#listUsers
type ListOptions struct {
	// Page the value of OPCNextPage from ListUsersResponse used for
	// paging results.
	Page string
	// Limit he maximum number of results that ListUsers is to return.
	Limit   uint64
	UserID  string
	GroupID string
}

// ListResponse response for List commands.
type ListResourceResponse struct {
	// Page can be passed in the ListUsersRequest argument of the next
	// call to ListUsers in order to page results.
	Page  string
	Items []IdentityResource
}

// Error is returned from unsuccessful API calls. The OPCRequestID if present
// is used to reference the failing requests for support.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#Error
type Error struct {
	Code         string `json:"code"`
	Message      string `json:"message"`
	OPCRequestID string `json:"opc-request-id,omitempty"`
}

// Error returns a formatted description of an API error.
func (e *Error) Error() string {
	return fmt.Sprintf("Code: %s; OPC Request ID: %s; Message: %s",
		e.Code,
		e.OPCRequestID,
		e.Message,
	)
}
