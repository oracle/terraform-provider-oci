// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

// UserGroupMembership returned by GetUserGroupMembership and related methods.
import (
	"time"
)

// UserGroupMembership represents the membership of a user in a group
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/identity/20160918/UserGroupMembership/
type UserGroupMembership struct {
	OPCRequestIDUnmarshaller
	ETagUnmarshaller
	CompartmentID  string    `json:"compartmentId"`
	GroupID        string    `json:"groupId"`
	ID             string    `json:"id"`
	InactiveStatus uint16    `json:"inactiveStatus"`
	State          string    `json:"lifecycleState"`
	TimeCreated    time.Time `json:"timeCreated"`
	UserID         string    `json:"userId"`
}

type ListUserGroupMemberships struct {
	OPCRequestIDUnmarshaller
	NextPageUnmarshaller
	Memberships []UserGroupMembership
}

func (l *ListUserGroupMemberships) GetList() interface{} {
	return &l.Memberships
}

// AddUserToGroup adds a user to a group.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/identity/20160918/UserGroupMembership/AddUserToGroup
func (c *Client) AddUserToGroup(userID, groupID string, opts *RetryTokenOptions) (res *UserGroupMembership, e error) {
	required := struct {
		GroupID string `header:"-" json:"groupId" url:"-"`
		UserID  string `header:"-" json:"userId" url:"-"`
	}{
		GroupID: groupID,
		UserID:  userID,
	}

	details := &requestDetails{
		name:     resourceUserGroupMemberships,
		optional: opts,
		required: required,
	}

	var resp *response
	if resp, e = c.identityApi.postRequest(details); e != nil {
		return
	}

	res = &UserGroupMembership{}
	e = resp.unmarshal(res)
	return
}

// GetUserGroupMembership returns a UserGroupMembership identified by id.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/identity/20160918/UserGroupMembership/GetUserGroupMembership
func (c *Client) GetUserGroupMembership(id string) (res *UserGroupMembership, e error) {
	details := &requestDetails{
		ids:  urlParts{id},
		name: resourceUserGroupMemberships,
	}

	var resp *response
	if resp, e = c.identityApi.getRequest(details); e != nil {
		return
	}

	res = &UserGroupMembership{}
	e = resp.unmarshal(res)
	return
}

// DeleteUserGroupMembership removes a user from a group.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/identity/20160918/UserGroupMembership/RemoveUserFromGroup
func (c *Client) DeleteUserGroupMembership(id string, opts *IfMatchOptions) (e error) {
	details := &requestDetails{
		ids:      urlParts{id},
		name:     resourceUserGroupMemberships,
		optional: opts,
	}

	return c.identityApi.deleteRequest(details)
}

// ListUserGroupMemberships lists the UserGroupMembership objects in a user's tenancy.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/identity/20160918/UserGroupMembership/ListUserGroupMemberships
func (c *Client) ListUserGroupMemberships(opts *ListMembershipsOptions) (resources *ListUserGroupMemberships, e error) {
	details := &requestDetails{
		name:     resourceUserGroupMemberships,
		optional: opts,
		required: ocidRequirement{c.authInfo.tenancyOCID},
	}

	var resp *response
	if resp, e = c.identityApi.getRequest(details); e != nil {
		return
	}

	resources = &ListUserGroupMemberships{}
	e = resp.unmarshal(resources)
	return
}
