package baremetal

// UserGroupMembership returned by GetUserGroupMembership and related methods.
import (
	"net/http"
	"time"
)

//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#UserGroupMembership
type UserGroupMembership struct {
	ETaggedResource
	// Unique identifier for a particular item such as a User or a Group
	ID            string    `json:"id"`
	CompartmentID string    `json:"compartmentId"`
	GroupID       string    `json:"groupId"`
	UserID        string    `json:"userId"`
	TimeCreated   time.Time `json:"timeCreated"`
	TimeModified  time.Time `json:"timeModified"`
	State         string    `json:"state"`
}

type AddUserToGroupRequest struct {
	GroupID string `json:"groupId"`
	UserID  string `json:"userId"`
}

// UserGroupMembershipResponse contains matches from a List request and optionally a
// Page field that can be used in subsequent List requests in conjunction with
// the Limit field to support pagination of results.
type UserGroupMembershipResponse struct {
	ResourceContainer
	Memberships []UserGroupMembership
}

func (l *UserGroupMembershipResponse) GetList() interface{} {
	return &l.Memberships
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
	e = response.unmarshal(resource)
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

	resource = &UserGroupMembership{}
	e = response.unmarshal(resource)
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

	response = &UserGroupMembershipResponse{}
	e = getResp.unmarshal(response)
	return
}
