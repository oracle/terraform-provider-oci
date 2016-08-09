package baremtlsdk

// UserGroupMembership returned by GetUserGroupMembership and related methods.
import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
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
func (c *Client) AddUserToGroup(userID, groupID string, opts ...Options) (membership *UserGroupMembership, e error) {
	var headers http.Header
	if len(opts) > 0 {
		if opts[0].OPCIdempotencyToken != "" {
			headers = http.Header{}
			headers.Set(headerOPCIdempotencyToken, opts[0].OPCIdempotencyToken)
		}
	}

	request := AddUserToGroupRequest{
		UserID:  userID,
		GroupID: groupID,
	}

	strURL := buildIdentityURL(resourceUserGroupMemberships, nil)

	resp, e := c.api.request(http.MethodPost, strURL, request, headers)

	if e != nil {
		return
	}

	membership = &UserGroupMembership{}
	e = json.Unmarshal(resp.body, membership)
	return
}

// DeleteUserGroupMembership removes a user from a group.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#removeUserFromGroup
func (c *Client) DeleteUserGroupMembership(userGroupMembershipID string, opts ...Options) (e error) {
	var headers http.Header
	if len(opts) > 0 {
		if opts[0].IfMatch != "" {
			headers = http.Header{}
			headers.Set(headerIfMatch, opts[0].IfMatch)
		}
	}

	url := buildIdentityURL(resourceUserGroupMemberships, nil, userGroupMembershipID)

	return c.api.deleteRequest(url, headers)

}

// GetUserGroupMembership returns a UserGroupMembership identified by id.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#getUserGroupMembership
func (c *Client) GetUserGroupMembership(id string) (userGroupMembership *UserGroupMembership, e error) {

	url := buildIdentityURL(resourceUserGroupMemberships, nil, id)

	var getResp *requestResponse
	if getResp, e = c.api.getRequest(url, nil); e != nil {
		return
	}

	reader := bytes.NewBuffer(getResp.body)
	decoder := json.NewDecoder(reader)
	userGroupMembership = &UserGroupMembership{}
	e = decoder.Decode(userGroupMembership)
	return

}

func (c *Client) ListUserGroupMemberships(options ...ListOptions) (response *UserGroupMembershipResponse, e error) {
	q := url.Values{}
	q.Set(queryCompartmentID, c.authInfo.tenancyOCID)
	if len(options) > 0 {
		opt := options[0]

		if opt.UserID != "" {
			q.Set(queryUserID, opt.UserID)
		}

		if opt.GroupID != "" {
			q.Set(queryGroupID, opt.GroupID)
		}

		if opt.Page != "" {
			q.Set(queryPage, opt.Page)
		}

		if opt.Limit > 0 {
			q.Set(queryLimit, strconv.FormatUint(opt.Limit, 10))
		}

		resourceURL := buildIdentityURL(resourceUserGroupMemberships, q)

		var getResp *requestResponse
		if getResp, e = c.api.getRequest(resourceURL, nil); e != nil {
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

	return
}
