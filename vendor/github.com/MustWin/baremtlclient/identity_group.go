package baremtlsdk

// CreateGroup create a new group. groupName MUST be supplied and MUST be
import (
	"encoding/json"
	"net/http"
)

// unique. groupDescription is optional. You MAY supply one option with an
// idempotency token.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#createGroup
func (c *Client) CreateGroup(groupName, groupDescription string, options ...Options) (response *IdentityResource, e error) {
	createRequest := CreateIdentityResourceRequest{
		CompartmentID: c.authInfo.tenancyOCID,
		Name:          groupName,
		Description:   groupDescription,
	}
	var headers http.Header
	if len(options) > 0 {
		if options[0].OPCIdempotencyToken != "" {
			headers = http.Header{}
			headers.Set(headerOPCIdempotencyToken, options[0].OPCIdempotencyToken)
		}
	}

	return c.createIdentityResource(resourceGroups, createRequest, headers)
}

// DeleteGroup removes a group identified by groupID. Optionally pass an
// etag for optmistic concurrency control.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#deleteGroup
func (c *Client) DeleteGroup(groupID string, opts ...Options) (e error) {
	var headers http.Header
	if len(opts) > 0 {
		if opts[0].IfMatch != "" {
			headers = http.Header{}
			headers.Set(headerIfMatch, opts[0].IfMatch)
		}
	}

	url := buildIdentityURL(resourceGroups, nil, groupID)

	return c.api.deleteRequest(url, headers)
}

// GetGroup returns a group identified by groupID.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#getGroup
func (c *Client) GetGroup(groupID string) (group *IdentityResource, e error) {
	group, e = c.getIdentityResource(resourceGroups, groupID)
	return
}

// ListGroups returns a list of Groups in a tenancy. The request MAY contain optional paging arguments.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#listGroups
func (c *Client) ListGroups(options ...ListOptions) (response *ListResourceResponse, e error) {
	return c.listIdentityResources(resourceGroups, options...)
}

// UpdateGroup updates the description of a group.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#updateGroup
func (c *Client) UpdateGroup(groupID, description string, options ...Options) (group *IdentityResource, e error) {

	headers := getUpdateHeaders(options...)

	request := UpdateIdentityResourceRequest{
		Description: description,
	}

	var resp []byte
	if resp, e = c.updateIdentityResource(resourceGroups, groupID, request, headers); e != nil {
		return
	}

	group = &IdentityResource{}
	e = json.Unmarshal(resp, group)
	return
}
