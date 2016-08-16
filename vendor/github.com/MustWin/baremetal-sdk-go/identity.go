package baremetal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

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

type CreateIdentityResourceRequest struct {
	CompartmentID string `json:"compartmentId"`
	Name          string `json:"name"`
	Description   string `json:"description"`
}

type UpdateIdentityResourceRequest struct {
	Description string `json:"description"`
}

func (c *Client) createIdentityResource(name resourceName, body CreateIdentityResourceRequest, opts []Options) (resource *IdentityResource, e error) {
	reqOpts := &sdkRequestOptions{
		body:    body,
		name:    name,
		options: opts,
	}

	var response *requestResponse
	if response, e = c.identityApi.request(http.MethodPost, reqOpts); e != nil {
		return
	}

	resource = &IdentityResource{}
	e = json.Unmarshal(response.body, resource)

	if respHeader := response.header; respHeader != nil {
		resource.ETag = respHeader.Get(headerETag)
	}

	return
}

func (c *Client) getIdentityResource(name resourceName, ids ...interface{}) (resource *IdentityResource, e error) {
	reqOpts := &sdkRequestOptions{
		name: name,
		ids:  ids,
	}

	var response *requestResponse
	if response, e = c.identityApi.getRequest(reqOpts); e != nil {
		return
	}

	reader := bytes.NewBuffer(response.body)
	decoder := json.NewDecoder(reader)
	resource = &IdentityResource{}
	e = decoder.Decode(resource)

	if respHeader := response.header; respHeader != nil {
		resource.ETag = respHeader.Get(headerETag)
	}

	return
}

func (c *Client) listIdentityResources(name resourceName, options ...Options) (resp *ListResourceResponse, e error) {
	conf := &sdkRequestOptions{
		name:    name,
		ocid:    c.authInfo.tenancyOCID,
		options: options,
	}

	var getResp *requestResponse
	if getResp, e = c.identityApi.getRequest(conf); e != nil {
		return
	}

	reader := bytes.NewBuffer(getResp.body)
	decoder := json.NewDecoder(reader)
	var items []IdentityResource
	if e = decoder.Decode(&items); e != nil {
		return
	}

	resp = &ListResourceResponse{
		Page:  getResp.header.Get(headerOPCNextPage),
		Items: items,
	}

	return
}

func (c *Client) updateIdentityResource(name resourceName, resourceID string, body interface{}, opts []Options) (resource *IdentityResource, e error) {
	conf := &sdkRequestOptions{
		body:    body,
		name:    name,
		options: opts,
		ids:     []interface{}{resourceID},
	}

	var response *requestResponse
	if response, e = c.identityApi.request(http.MethodPut, conf); e != nil {
		return
	}

	resource = &IdentityResource{}
	e = json.Unmarshal(response.body, resource)

	if respHeader := response.header; respHeader != nil {
		resource.ETag = respHeader.Get(headerETag)
	}

	return
}
