// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

import (
	"net/http"
	"time"
)

type User struct {
	ETagUnmarshaller
	OPCRequestIDUnmarshaller
	CompartmentID  string    `json:"compartmentId"`
	Description    string    `json:"description"`
	ID             string    `json:"id"`
	InactiveStatus uint16    `json:"inactiveStatus"`
	Name           string    `json:"name"`
	State          string    `json:"lifecycleState"`
	TimeCreated    time.Time `json:"timeCreated"`
}

type ListUsers struct {
	OPCRequestIDUnmarshaller
	NextPageUnmarshaller
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
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/identity/20160918/User/CreateUser
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

	var resp *response
	if resp, e = c.identityApi.postRequest(details); e != nil {
		return
	}

	res = &User{}
	e = resp.unmarshal(res)
	return
}

// GetUser returns a user identified by userID.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/identity/20160918/User/GetUser
func (c *Client) GetUser(id string) (res *User, e error) {
	details := &requestDetails{
		ids:  urlParts{id},
		name: resourceUsers,
	}

	var resp *response
	if resp, e = c.identityApi.getRequest(details); e != nil {
		return
	}

	res = &User{}
	e = resp.unmarshal(res)
	return
}

// UpdateUser updates the description of the specified user
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/identity/20160918/User/UpdateUser
func (c *Client) UpdateUser(id string, opts *UpdateIdentityOptions) (res *User, e error) {
	details := &requestDetails{
		ids:      urlParts{id},
		name:     resourceUsers,
		optional: opts,
	}

	var resp *response
	if resp, e = c.identityApi.request(http.MethodPut, details); e != nil {
		return
	}

	res = &User{}
	e = resp.unmarshal(res)
	return
}

// UpdateUserState updates the state of the specified user
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/identity/20160918/User/UpdateUserState
func (c *Client) UpdateUserState(id string, opts *UpdateUserStateOptions) (res *User, e error) {
	details := &requestDetails{
		ids:      urlParts{id, "state"},
		name:     resourceUsers,
		optional: opts,
	}

	var resp *response
	if resp, e = c.identityApi.request(http.MethodPut, details); e != nil {
		return
	}

	res = &User{}
	e = resp.unmarshal(res)
	return
}

// DeleteUser deletes a user
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/identity/20160918/User/DeleteUser
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
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/identity/20160918/User/ListUsers
func (c *Client) ListUsers(opts *ListOptions) (resources *ListUsers, e error) {
	details := &requestDetails{
		name:     resourceUsers,
		optional: opts,
		required: listOCIDRequirement{c.authInfo.tenancyOCID},
	}
	var resp *response
	if resp, e = c.identityApi.getRequest(details); e != nil {
		return
	}

	resources = &ListUsers{}
	e = resp.unmarshal(resources)
	return
}
