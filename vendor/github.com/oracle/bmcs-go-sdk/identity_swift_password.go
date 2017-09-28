// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

import (
	"net/http"
	"time"
)

type SwiftPassword struct {
	ETagUnmarshaller
	OPCRequestIDUnmarshaller
	Password       string    `json:"password"`
	ID             string    `json:"id"`
	UserID         string    `json:"userId"`
	Description    string    `json:"description"`
	State          string    `json:"lifecycleState"`
	InactiveStatus uint64    `json:"inactiveStatus"`
	ExpiresOn      time.Time `json:"expiresOn"`
	TimeCreated    time.Time `json:"timeCreated"`
}

type ListSwiftPasswords struct {
	OPCRequestIDUnmarshaller
	NextPageUnmarshaller
	SwiftPasswords []SwiftPassword
}

func (l *ListSwiftPasswords) GetList() interface{} {
	return &l.SwiftPasswords
}

// CreateSwiftPassword creates a new SwiftPassword for userID.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/identity/20160918/SwiftPassword/CreateSwiftPassword
func (c *Client) CreateSwiftPassword(userID, desc string, opts *RetryTokenOptions) (res *SwiftPassword, e error) {
	required := struct {
		Description string `header:"-" json:"description" url:"-"`
	}{
		Description: desc,
	}

	details := &requestDetails{
		ids:      urlParts{userID, resourceSwiftPasswords},
		name:     resourceUsers,
		optional: opts,
		required: required,
	}

	var resp *response
	if resp, e = c.identityApi.postRequest(details); e != nil {
		return
	}

	res = &SwiftPassword{}
	e = resp.unmarshal(res)
	return
}

// UpdateSwiftPassword updates a SwiftPassword.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/identity/20160918/SwiftPassword/UpdateSwiftPassword
func (c *Client) UpdateSwiftPassword(id, userID string, opts *UpdateIdentityOptions) (res *SwiftPassword, e error) {
	details := &requestDetails{
		ids:      urlParts{userID, resourceSwiftPasswords, id},
		name:     resourceUsers,
		optional: opts,
	}

	var resp *response
	if resp, e = c.identityApi.request(http.MethodPut, details); e != nil {
		return
	}

	res = &SwiftPassword{}
	e = resp.unmarshal(res)
	return
}

// DeleteSwiftPassword deletes a SwiftPassword id for userID.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/identity/20160918/SwiftPassword/DeleteSwiftPassword
func (c *Client) DeleteSwiftPassword(id, userID string, opts *IfMatchOptions) (e error) {
	details := &requestDetails{
		ids:      urlParts{userID, resourceSwiftPasswords, id},
		name:     resourceUsers,
		optional: opts,
	}

	return c.identityApi.deleteRequest(details)
}

// ListSwiftPasswords gets all SwiftPasswords for userID.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/identity/20160918/SwiftPassword/ListSwiftPasswords
func (c *Client) ListSwiftPasswords(userID string) (resources *ListSwiftPasswords, e error) {
	details := &requestDetails{
		ids:  urlParts{userID, resourceSwiftPasswords},
		name: resourceUsers,
	}

	var resp *response
	if resp, e = c.identityApi.getRequest(details); e != nil {
		return
	}

	resources = &ListSwiftPasswords{}
	e = resp.unmarshal(resources)
	return
}
