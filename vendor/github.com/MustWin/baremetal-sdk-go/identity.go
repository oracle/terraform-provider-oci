package baremetal

import (
	"fmt"
	"net/http"
	"time"
)

// Resource contains information representing Users, Groups,
// Policies and other elements
type IdentityResource struct {
	ETaggedResource
	// Unique identifier for a particular item such as a User or a Group
	ID string `json:"id"`
	// CompartmentID is the ID of the tenancy containing the compartment
	CompartmentID string    `json:"compartmentId"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	TimeCreated   time.Time `json:"timeCreated"`
	TimeModified  time.Time `json:"timeModified"`
	State         string    `json:"state"`
}

// ListResponse response for List commands.
type ListResourceResponse struct {
	ResourceContainer
	Items []IdentityResource
}

func (l *ListResourceResponse) GetList() interface{} {
	return &l.Items
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
	e = response.unmarshal(resource)
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

	resource = &IdentityResource{}
	e = response.unmarshal(resource)
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

	resp = &ListResourceResponse{}
	e = getResp.unmarshal(resp)
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
	e = response.unmarshal(resource)
	return
}
