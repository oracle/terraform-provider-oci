package baremetal

import (
	"net/http"
	"time"
)

type Group struct {
	ETaggedResource
	CompartmentID  string    `json:"compartmentId"`
	Description    string    `json:"description"`
	ID             string    `json:"id"`
	InactiveStatus uint16    `json:"inactiveStatus"`
	Name           string    `json:"name"`
	State          string    `json:"lifecycleState"`
	TimeCreated    time.Time `json:"timeCreated"`
}

type ListGroups struct {
	ResourceContainer
	Groups []Group
}

func (l *ListGroups) GetList() interface{} {
	return &l.Groups
}

// CreateGroup create a new group. groupName MUST be supplied and MUST be
// unique. groupDescription is optional. You MAY supply one option with an
// idempotency token.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/identity/20160918/Group/CreateGroup
func (c *Client) CreateGroup(name, desc string, opts *RetryTokenOptions) (res *Group, e error) {
	required := identityCreationRequirement{
		CompartmentID: c.authInfo.tenancyOCID,
		Description:   desc,
		Name:          name,
	}

	details := &requestDetails{
		name:     resourceGroups,
		optional: opts,
		required: required,
	}

	var response *requestResponse
	if response, e = c.identityApi.request(http.MethodPost, details); e != nil {
		return
	}

	res = &Group{}
	e = response.unmarshal(res)
	return
}

// GetGroup returns a group identified by groupID.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/identity/20160918/Group/GetGroup
func (c *Client) GetGroup(id string) (res *Group, e error) {
	details := &requestDetails{
		ids:  urlParts{id},
		name: resourceGroups,
	}

	var response *requestResponse
	if response, e = c.identityApi.getRequest(details); e != nil {
		return
	}

	res = &Group{}
	e = response.unmarshal(res)
	return
}

// UpdateGroup updates the description of a group.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/identity/20160918/Group/UpdateGroup
func (c *Client) UpdateGroup(id string, opts *UpdateIdentityOptions) (res *Group, e error) {
	details := &requestDetails{
		ids:      urlParts{id},
		name:     resourceGroups,
		optional: opts,
	}

	var response *requestResponse
	if response, e = c.identityApi.request(http.MethodPut, details); e != nil {
		return
	}

	res = &Group{}
	e = response.unmarshal(res)
	return
}

// DeleteGroup removes a group identified by groupID. Optionally pass an
// etag for optmistic concurrency control.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/identity/20160918/Group/DeleteGroup
func (c *Client) DeleteGroup(id string, opts *IfMatchOptions) (e error) {
	details := &requestDetails{
		ids:      urlParts{id},
		name:     resourceGroups,
		optional: opts,
	}

	return c.identityApi.deleteRequest(details)
}

// ListGroups returns a list of Groups in a tenancy. The request MAY contain optional paging arguments.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/identity/20160918/Group/ListGroups
func (c *Client) ListGroups(opts *ListOptions) (resources *ListGroups, e error) {
	details := &requestDetails{
		name:     resourceGroups,
		optional: opts,
		required: ocidRequirement{c.authInfo.tenancyOCID},
	}

	var response *requestResponse
	if response, e = c.identityApi.getRequest(details); e != nil {
		return
	}

	resources = &ListGroups{}
	e = response.unmarshal(resources)
	return
}
