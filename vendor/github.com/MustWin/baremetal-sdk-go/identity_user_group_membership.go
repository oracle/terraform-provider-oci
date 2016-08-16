package baremetal

// UserGroupMembership returned by GetUserGroupMembership and related methods.
import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#UserGroupMembership
type UserGroupMembership struct {
	// Unique identifier for a particular item such as a User or a Group
	ID            string    `json:"id"`
	CompartmentID string    `json:"compartmentId"`
	GroupID       string    `json:"groupId"`
	UserID        string    `json:"userId"`
	TimeCreated   time.Time `json:"timeCreated"`
	TimeModified  time.Time `json:"timeModified"`
	State         string    `json:"state"`
	ETag          string    `json:"etag,omitempty"`
	OPCRequestID  string    `json:"opc-request-id,omitempty"`
}

type AddUserToGroupRequest struct {
	GroupID string `json:"groupId"`
	UserID  string `json:"userId"`
}

// UserGroupMembershipResponse contains matches from a List request and optionally a
// Page field that can be used in subsequent List requests in conjunction with
// the Limit field to support pagination of results.
type UserGroupMembershipResponse struct {
	Page        string
	Memberships []UserGroupMembership
}

// AddUserToGroup adds a user to a group.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#addUserToGroup
func (c *Client) AddUserToGroup(userID, groupID string, opts ...Options) (resource *UserGroupMembership, e error) {
	body := AddUserToGroupRequest{
		UserID:  userID,
		GroupID: groupID,
	}

	reqOpts := &sdkRequestOptions{
		body:    body,
		name:    resourceUserGroupMemberships,
		options: opts,
	}

	response, e := c.identityApi.request(http.MethodPost, reqOpts)
	if e != nil {
		return
	}

	resource = &UserGroupMembership{}
	e = json.Unmarshal(response.body, resource)

	if respHeader := response.header; respHeader != nil {
		resource.ETag = respHeader.Get(headerETag)
	}

	return
}

// DeleteUserGroupMembership removes a user from a group.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#removeUserFromGroup
func (c *Client) DeleteUserGroupMembership(userGroupMembershipID string, opts ...Options) (e error) {
	reqOpts := &sdkRequestOptions{
		name:    resourceUserGroupMemberships,
		options: opts,
		ids:     urlParts{userGroupMembershipID},
	}
	return c.identityApi.deleteRequest(reqOpts)
}

// GetUserGroupMembership returns a UserGroupMembership identified by id.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#getUserGroupMembership
func (c *Client) GetUserGroupMembership(id string) (resource *UserGroupMembership, e error) {
	reqOpts := &sdkRequestOptions{
		name: resourceUserGroupMemberships,
		ids:  urlParts{id},
	}

	var response *requestResponse
	if response, e = c.identityApi.getRequest(reqOpts); e != nil {
		return
	}

	reader := bytes.NewBuffer(response.body)
	decoder := json.NewDecoder(reader)
	resource = &UserGroupMembership{}
	e = decoder.Decode(resource)

	if respHeader := response.header; respHeader != nil {
		resource.ETag = respHeader.Get(headerETag)
	}

	return
}

func (c *Client) ListUserGroupMemberships(options ...Options) (response *UserGroupMembershipResponse, e error) {
	reqOpts := &sdkRequestOptions{
		name:    resourceUserGroupMemberships,
		options: options,
		ocid:    c.authInfo.tenancyOCID,
	}

	var getResp *requestResponse
	if getResp, e = c.identityApi.getRequest(reqOpts); e != nil {
		return
	}

	reader := bytes.NewBuffer(getResp.body)
	decoder := json.NewDecoder(reader)
	var items []UserGroupMembership

	if e = decoder.Decode(&items); e != nil {
		return
	}

	response = &UserGroupMembershipResponse{
		Page:        getResp.header.Get(headerOPCNextPage),
		Memberships: items,
	}

	return
}
