package baremetal

import (
	"net/http"
	"time"
)

type User struct {
	ETaggedResource
	CompartmentID  string    `json:"compartmentId"`
	Description    string    `json:"description"`
	ID             string    `json:"id"`
	InactiveStatus uint16    `json:"inactiveStatus"`
	Name           string    `json:"name"`
	State          string    `json:"lifecycleState"`
	TimeCreated    time.Time `json:"timeCreated"`
}

type ListUsers struct {
	ResourceContainer
	Users []User
}

func (l *ListUsers) GetList() interface{} {
	return &l.Users
}

// CreateUser is used to create a user. userName MUST be unique. description
// contains a comment about the user. The caller can supply 0 or 1 options. Options
// MAY contain an idempotency token.
// The caller specifies this token so that subsequent calls to create user will
// be idempotent. The token expires after 30 minutes.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/identity/20160918/User/CreateUser
func (c *Client) CreateUser(name, desc string, opts *RetryTokenOptions) (res *User, e error) {
	required := identityCreationRequirement{
		CompartmentID: c.authInfo.tenancyOCID,
		Description:   desc,
		Name:          name,
	}

	details := &requestDetails{
		name:     resourceUsers,
		optional: opts,
		required: required,
	}

	var response *requestResponse
	if response, e = c.identityApi.request(http.MethodPost, details); e != nil {
		return
	}

	res = &User{}
	e = response.unmarshal(res)
	return
}

// GetUser returns a user identified by userID.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/identity/20160918/User/GetUser
func (c *Client) GetUser(id string) (res *User, e error) {
	details := &requestDetails{
		ids:  urlParts{id},
		name: resourceUsers,
	}

	var response *requestResponse
	if response, e = c.identityApi.getRequest(details); e != nil {
		return
	}

	res = &User{}
	e = response.unmarshal(res)
	return
}

// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/identity/20160918/User/UpdateUser
func (c *Client) UpdateUser(id string, opts *UpdateIdentityOptions) (res *User, e error) {
	details := &requestDetails{
		ids:      urlParts{id},
		name:     resourceUsers,
		optional: opts,
	}

	var response *requestResponse
	if response, e = c.identityApi.request(http.MethodPut, details); e != nil {
		return
	}

	res = &User{}
	e = response.unmarshal(res)
	return
}

// TODO: UpdateUserState
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/identity/20160918/User/UpdateUserState

// DeleteUser deletes a user.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/identity/20160918/User/DeleteUser
func (c *Client) DeleteUser(id string, opts *IfMatchOptions) (e error) {
	details := &requestDetails{
		ids:      urlParts{id},
		name:     resourceUsers,
		optional: opts,
	}

	return c.identityApi.deleteRequest(details)
}

// ListUsers returns an array of users for the current tenancy.  The requestor
// MAY supply paging options.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/identity/20160918/User/ListUsers
func (c *Client) ListUsers(opts *ListOptions) (resources *ListUsers, e error) {
	details := &requestDetails{
		name:     resourceUsers,
		optional: opts,
		required: ocidRequirement{c.authInfo.tenancyOCID},
	}

	var response *requestResponse
	if response, e = c.identityApi.getRequest(details); e != nil {
		return
	}

	resources = &ListUsers{}
	e = response.unmarshal(resources)
	return
}
