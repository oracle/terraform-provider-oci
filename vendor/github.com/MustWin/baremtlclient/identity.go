package baremtlsdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// APIKey is returned for operations that create or modify user API keys.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#ApiKey
type APIKey struct {
	KeyID        string    `json:"keyId"`
	KeyValue     string    `json:"keyValue"`
	Fingerprint  string    `json:"fingerprint"`
	UserID       string    `json:"userId"`
	TimeCreated  time.Time `json:"timeCreated"`
	TimeModified time.Time `json:"timeModified"`
	State        string    `json:"state"`
}

// ListAPIKeyResponse contains a list of API keys
type ListAPIKeyResponse struct {
	OPCRequestID string
	Keys         []APIKey
}

// UIPassword is returned for change or create password operations.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#UIPassword
type UIPassword struct {
	NewPassword  string    `json:"password"`
	UserID       string    `json:"userId"`
	TimeCreated  time.Time `json:"timeCreated"`
	TimeModified time.Time `json:"timeModified"`
	State        string    `json:"state"`
	ETag         string    `json:"etag,omitempty"`
	OPCRequestID string    `json:"opc-request-id,omitempty"`
}

// Options is used to pass optional values to to package methods. Note that zero-value
// fields will be ignored. Typically options are passed on a variadic paramter
// to SDK methods. Note that only the first option, if present, will be used.
type Options struct {
	// OPCIdempotencyToken (Optional) - A token you supply to uniquely identify the request and provide idempotency
	//if the request is retried. Idempotency tokens expire after 30 minutes.
	OPCIdempotencyToken string
	// IfMatch (Optional) is for optimistic concurrency control. In the PUT or DELETE call for
	// a resource, set the if-match parameter to the value of the etag from a previous GET or POST response
	// for that resource. The resource will be updated or deleted only if the etag you
	// provide matches the resource's current etag value.
	IfMatch string
}

// AvailablityDomain contains name and then tenancy ID that an
// availability domain belongs to.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#AvailabilityDomain
type AvailabilityDomain struct {
	Name          string `json:"name"`
	CompartmentID string `json:"compartmentId"`
}

// Resource contains information representing Users, Groups,
// Policies and other elements
type Resource struct {
	// Unique identifier for a particular item such as a User or a Group
	ID string `json:"id"`
	// CompartmentID is the ID of the tenancy containing the compartment
	CompartmentID string    `json:"compartmentId"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	TimeCreated   time.Time `json:"timeCreated"`
	TimeModified  time.Time `json:"timeModified"`
	State         string    `json:"state"`
	ETag          string    `json:"etag,omitempty"`
	OPCRequestID  string    `json:"opc-request-id,omitempty"`
}

// ListOptions contains arguments to support pagination for List requests. ListOptions
// is typically a variadic paramter to List SDK functions.  Only the first ListOption will
// be used.  If multiple ListOptions are passed the subsequent values after the first
// will be discarded.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#listUsers
type ListOptions struct {
	// Page the value of OPCNextPage from ListUsersResponse used for
	// paging results.
	Page string
	// Limit he maximum number of results that ListUsers is to return.
	Limit   uint64
	UserID  string
	GroupID string
}

// ListResponse response for List commands.
type ListResourceResponse struct {
	// Page can be passed in the ListUsersRequest argument of the next
	// call to ListUsers in order to page results.
	Page  string
	Items []Resource
}

// Policy returned by GetPolicy and other policy related methods.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#Policy
type Policy struct {
	Resource
	Statements []string `json:"statements"`
}

// UserGroupMembership returned by GetUserGroupMembership and related methods.
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

// CreateCompartment create a new compartment.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#createCompartment
func (c *Client) CreateCompartment(compartmentName, compartmentDescription string, options ...Options) (compartment *Resource, e error) {
	createRequest := CreateResourceRequest{
		CompartmentID: c.authInfo.tenancyOCID,
		Name:          compartmentName,
		Description:   compartmentDescription,
	}
	var headers http.Header
	if len(options) > 0 {
		if options[0].OPCIdempotencyToken != "" {
			headers = http.Header{}
			headers.Set(headerOPCIdempotencyToken, options[0].OPCIdempotencyToken)
		}
	}

	return c.createResource(resourceCompartments, createRequest, headers)
}

// CreateGroup create a new group. groupName MUST be supplied and MUST be
// unique. groupDescription is optional. You MAY supply one option with an
// idempotency token.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#createGroup
func (c *Client) CreateGroup(groupName, groupDescription string, options ...Options) (response *Resource, e error) {
	createRequest := CreateResourceRequest{
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

	return c.createResource(resourceGroups, createRequest, headers)
}

// CreateOrResetUIPassword - creates or resets password for user identified by
// userID. You MAY supply an idempotency token.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#createOrResetUIPassword
func (c *Client) CreateOrResetUIPassword(password, userID string, opts ...Options) (newpassword *UIPassword, e error) {
	var headers http.Header
	if len(opts) > 0 && opts[0].OPCIdempotencyToken != "" {
		headers = http.Header{}
		headers.Set(headerOPCIdempotencyToken, opts[0].OPCIdempotencyToken)
	}

	url := buildIdentityURL(resourceUsers, nil, userID, uiPassword)
	request := UpdateUIPasswordRequest{
		Password: password,
	}

	var response *requestResponse
	if response, e = c.api.request(http.MethodPost, url, request, headers); e != nil {
		return
	}

	newpassword = &UIPassword{}
	if e = json.Unmarshal(response.body, newpassword); e != nil {
		return
	}

	if response.header != nil {
		newpassword.ETag = response.header.Get(headerIfMatch)
		newpassword.OPCRequestID = response.header.Get(headerOPCRequestID)
	}

	return
}

// CreatePolicy creates a new policy.
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#createPolicy
func (c *Client) CreatePolicy(policyName, policyDescription string, statements []string, options ...Options) (policy *Policy, e error) {
	urlStr := buildIdentityURL(resourcePolicies, nil)
	var headers http.Header

	if len(options) > 0 {
		if options[0].OPCIdempotencyToken != "" {
			headers = http.Header{}
			headers.Set(headerOPCIdempotencyToken, options[0].OPCIdempotencyToken)
		}
	}

	var request CreatePolicyRequest
	request.CompartmentID = c.authInfo.tenancyOCID
	request.Name = policyName
	request.Description = policyDescription
	request.Statements = statements

	var resp *requestResponse
	if resp, e = c.api.request(http.MethodPost, urlStr, request, headers); e != nil {
		return
	}

	policy = &Policy{}
	e = json.Unmarshal(resp.body, policy)
	return

}

// CreateUser is used to create a user. userName MUST be unique. description
// contains a comment about the user. The caller can supply 0 or 1 options. Options
// MAY contain an idempotency token.
// The caller specifies this token so that subsequent calls to create user will
// be idempotent. The token expires after 30 minutes.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#createUser
func (c *Client) CreateUser(userName, userDescription string, options ...Options) (user *Resource, e error) {
	createRequest := CreateResourceRequest{
		CompartmentID: c.authInfo.tenancyOCID,
		Name:          userName,
		Description:   userDescription,
	}
	var headers http.Header
	if len(options) > 0 {
		if options[0].OPCIdempotencyToken != "" {
			headers = http.Header{}
			headers.Set(headerOPCIdempotencyToken, options[0].OPCIdempotencyToken)
		}
	}
	return c.createResource(resourceUsers, createRequest, headers)
}

// Deletes an API key belonging to a user.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#deleteApiKey
func (c *Client) DeleteAPIKey(userID, fingerprint string, opts ...Options) (e error) {
	var headers http.Header
	if len(opts) > 0 {
		if opts[0].IfMatch != "" {
			headers = http.Header{}
			headers.Set(headerIfMatch, opts[0].IfMatch)
		}
	}

	url := buildIdentityURL(resourceUsers, nil, userID, apiKeys, fingerprint)
	return c.api.deleteRequest(url, headers)

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

// DeletePolicy removes a policy identified by policyID. Optionally pass an
// etag for optmistic concurrency control.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#DeletePolicy
func (c *Client) DeletePolicy(policyID string, opts ...Options) (e error) {
	var headers http.Header
	if len(opts) > 0 {
		if opts[0].IfMatch != "" {
			headers = http.Header{}
			headers.Set(headerIfMatch, opts[0].IfMatch)
		}
	}

	url := buildIdentityURL(resourcePolicies, nil, policyID)

	return c.api.deleteRequest(url, headers)

}

// DeleteUser deletes a user.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#deleteUser
func (c *Client) DeleteUser(userID string, opts ...Options) (e error) {
	var headers http.Header
	if len(opts) > 0 {
		if opts[0].IfMatch != "" {
			headers = http.Header{}
			headers.Set(headerIfMatch, opts[0].IfMatch)
		}
	}

	url := buildIdentityURL(resourceUsers, nil, userID)

	return c.api.deleteRequest(url, headers)

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

// GetCompartment returns the compartment identified by compartmentID.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/#apiref.htm
func (c *Client) GetCompartment(compartmentID string) (compartment *Resource, e error) {
	compartment, e = c.getIdentity(resourceCompartments, compartmentID)
	return
}

// GetGroup returns a group identified by groupID.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#getGroup
func (c *Client) GetGroup(groupID string) (group *Resource, e error) {
	group, e = c.getIdentity(resourceGroups, groupID)
	return
}

// GetPolicy returns a policy identified by a policyID.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#getPolicy
func (c *Client) GetPolicy(policyID string) (policy *Policy, e error) {

	url := buildIdentityURL(resourcePolicies, nil, policyID)

	var getResp *requestResponse
	if getResp, e = c.api.getRequest(url, nil); e != nil {
		return
	}

	reader := bytes.NewBuffer(getResp.body)
	decoder := json.NewDecoder(reader)
	policy = &Policy{}
	e = decoder.Decode(policy)
	return

}

// GetUser returns a user identified by userID.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/#apiref.htm
func (c *Client) GetUser(userID string) (user *Resource, e error) {
	user, e = c.getIdentity(resourceUsers, userID)
	return
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

// ListAPIKeys returns information about a user's API keys.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#listApiKeys
func (c *Client) ListAPIKeys(userID string) (response *ListAPIKeyResponse, e error) {
	url := buildIdentityURL(resourceUsers, nil, userID, apiKeys, "/")
	var getResp *requestResponse
	if getResp, e = c.api.getRequest(url, nil); e != nil {
		return
	}
	reader := bytes.NewBuffer(getResp.body)
	decoder := json.NewDecoder(reader)
	var keys []APIKey

	if e = decoder.Decode(&keys); e != nil {
		return
	}

	response = &ListAPIKeyResponse{
		Keys:         keys,
		OPCRequestID: getResp.header.Get(headerOPCRequestID),
	}

	return

}

// ListAvailablityDomains lists availability domains in a user's root tenancy.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#listAvailabilityDomains
func (c *Client) ListAvailablityDomains() (ads []AvailabilityDomain, e error) {
	url := buildIdentityURL(resourceAvailabilityDomains, url.Values{
		queryCompartmentID: []string{c.authInfo.tenancyOCID},
	})
	var getResp *requestResponse

	if getResp, e = c.api.getRequest(url, nil); e != nil {
		return
	}

	reader := bytes.NewBuffer(getResp.body)
	decoder := json.NewDecoder(reader)
	e = decoder.Decode(&ads)
	return
}

// ListCompartments returns a list of compartments. The request MAY contain optional paging arguments.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#listCompartments
func (c *Client) ListCompartments(options ...ListOptions) (response *ListResourceResponse, e error) {
	return c.listItems(resourceCompartments, options...)
}

// ListGroups returns a list of Groups in a tenancy. The request MAY contain optional paging arguments.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#listGroups
func (c *Client) ListGroups(options ...ListOptions) (response *ListResourceResponse, e error) {
	return c.listItems(resourceGroups, options...)
}

// ListUsers returns an array of users for the current tenancy.  The requestor
// MAY supply paging options.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#listUsers
func (c *Client) ListUsers(options ...ListOptions) (response *ListResourceResponse, e error) {
	return c.listItems(resourceUsers, options...)
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

// UpdateCompartment updates the description of a compartment.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#updateCompartment
func (c *Client) UpdateCompartment(compartmentID, description string, options ...Options) (compartment *Resource, e error) {

	headers := getUpdateHeaders(options...)

	request := UpdateResourceRequest{
		Description: description,
	}

	var resp []byte
	if resp, e = c.updateResource(resourceCompartments, compartmentID, request, headers); e != nil {
		return
	}

	compartment = &Resource{}
	e = json.Unmarshal(resp, compartment)
	return

}

// UpdateGroup updates the description of a group.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#updateGroup
func (c *Client) UpdateGroup(groupID, description string, options ...Options) (group *Resource, e error) {

	headers := getUpdateHeaders(options...)

	request := UpdateResourceRequest{
		Description: description,
	}

	var resp []byte
	if resp, e = c.updateResource(resourceGroups, groupID, request, headers); e != nil {
		return
	}

	group = &Resource{}
	e = json.Unmarshal(resp, group)
	return
}

// UpdatePolicy can be called to modify the description and statements of an existing policy.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#updatePolicy
func (c *Client) UpdatePolicy(policyID, policyDescription string, policyStatements []string, opts ...Options) (policy *Policy, e error) {
	var headers http.Header
	if len(opts) > 0 {
		if opts[0].IfMatch != "" {
			headers = http.Header{}
			headers.Set(headerIfMatch, opts[0].IfMatch)
		}
	}

	var request UpdatePolicyRequest
	request.Description = policyDescription
	request.Statements = policyStatements

	var resp []byte
	if resp, e = c.updateResource(resourcePolicies, policyID, request, headers); e != nil {
		return
	}

	policy = &Policy{}
	e = json.Unmarshal(resp, policy)
	return
}

func (c *Client) UpdateUser(userID, userDescription string, opts ...Options) (user *Resource, e error) {
	headers := getUpdateHeaders(opts...)
	request := UpdateResourceRequest{
		Description: userDescription,
	}

	var resp []byte
	if resp, e = c.updateResource(resourceUsers, userID, request, headers); e != nil {
		return
	}

	user = &Resource{}
	e = json.Unmarshal(resp, user)
	return

}

// UpdateUserUIPassword - Changes the password of a user identified by userID. An
// ETAG MAY be passed as an option for optimistic concurrency control.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#createOrResetUIPassword
func (c *Client) UpdateUserUIPassword(newPassword, userID string, opts ...Options) (uipwd *UIPassword, e error) {
	var headers http.Header
	if len(opts) > 0 && opts[0].IfMatch != "" {
		headers = http.Header{}
		headers.Set(headerIfMatch, opts[0].IfMatch)
	}

	request := UpdateUIPasswordRequest{
		Password: newPassword,
	}

	url := buildIdentityURL(resourceUsers, nil, userID, uiPassword)

	var response *requestResponse
	if response, e = c.api.request(http.MethodPut, url, request, headers); e != nil {
		return
	}

	uipwd = &UIPassword{}
	if e = json.Unmarshal(response.body, uipwd); e != nil {
		return
	}

	if response.header != nil {
		uipwd.ETag = response.header.Get(headerIfMatch)
		uipwd.OPCRequestID = response.header.Get(headerOPCRequestID)
	}

	return
}

// UploadAPIKey - add an API signing key for user. The key must be an RSA public
// key in pem format.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#uploadApiKey
func (c *Client) UploadAPIKey(userID, key string, opts ...Options) (apiKey *APIKey, e error) {
	url := buildIdentityURL(resourceUsers, nil, userID, apiKeys, "/")
	request := CreateAPIKeyRequest{
		Key: key,
	}

	var headers http.Header

	if len(opts) > 0 {
		if opts[0].OPCIdempotencyToken != "" {
			headers = http.Header{}
			headers.Set(headerOPCIdempotencyToken, opts[0].OPCIdempotencyToken)
		}
	}

	var resp *requestResponse
	if resp, e = c.api.request(http.MethodPost, url, request, headers); e != nil {
		return
	}

	apiKey = &APIKey{}
	e = json.Unmarshal(resp.body, apiKey)
	return

}
