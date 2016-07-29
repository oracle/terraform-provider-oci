package baremtlclient

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// CreateResourceRequest contains information to create a new user.
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#CreateUserRequest
type CreateResourceRequest struct {
	CompartmentID string `json:"compartmentId"`
	Name          string `json:"name"`
	Description   string `json:"description"`
}

// AvailablityDomain contains name and then tenancy ID that an
// availability domain belongs to.
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

// ListRequest contains arguments for List requests. It contains optional
// fields to support pagination
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#listUsers
type ListResourceRequest struct {
	// Your tenancy's OCID
	CompartmentID string
	// (Optional) The value of OPCNextPage from ListUsersResponse used for
	// paging results.
	Page string
	// (Optional) The maximum number of results that ListUsers is to return.
	Limit uint64
}

// UpdateResourceRequest updates description for Compartment, Group,
// User
type UpdateResourceRequest struct {
	Description string `json:"description"`
}

func (l *ListResourceRequest) buildQuery(resource resourceName) (string, error) {

	query := url.Values{}
	if l.CompartmentID == "" {
		return "", errors.New("Missing Compartment ID")
	}

	query.Set("compartmentId", l.CompartmentID)

	if l.Limit > 0 {
		query.Set("limit", strconv.FormatUint(l.Limit, 10))
	}

	if l.Page != "" {
		query.Set("page", l.Page)
	}

	return buildIdentityURL(resource, &query), nil
}

// ListResponse response for List commands
type ListResourceResponse struct {
	// Page can be passed in the ListUsersRequest argument of the next
	// call to ListUsers in order to page results.
	Page  string
	Items []Resource
}

// Policy returned by GetPolicy and other policy related methods
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#Policy
type Policy struct {
	Resource
	Statements []string `json:"statements"`
}

// UserGroupMembership returned by GetUserGroupMembership and related methods
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#UserGroupMembership
type UserGroupMembership struct {
	Resource
	GroupID string `json:"groupId"`
	UserID  string `json:"userId"`
}

// Error is returned from unsuccessful API calls. The OPCRequestID if present
// is used to reference the failing requests for support.
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

// CreateCompartment create a new compartment.
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#createCompartment
func (c *Client) CreateCompartment(request CreateResourceRequest, headers http.Header) (compartment *Resource, e error) {
	return c.createResource(resourceCompartments, request, headers)
}

// CreateGroup create a new group.
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#createGroup
func (c *Client) CreateGroup(request CreateResourceRequest, headers http.Header) (response *Resource, e error) {
	return c.createResource(resourceGroups, request, headers)
}

func (c *Client) createResource(resourceType resourceName, request CreateResourceRequest, headers http.Header) (resource *Resource, e error) {
	urlStr := buildIdentityURL(resourceType, nil)
	var bodyBuffer []byte

	if bodyBuffer, e = c.identityAPI.request(http.MethodPost, urlStr, request, headers); e != nil {
		return
	}

	resource = &Resource{}
	e = json.Unmarshal(bodyBuffer, resource)
	return

}

// CreateUser is used to create a user.
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#createUser
func (c *Client) CreateUser(createRequest CreateResourceRequest, headers http.Header) (user *Resource, e error) {
	return c.createResource(resourceUsers, createRequest, headers)
}

func (c *Client) getIdentity(resource resourceName, ids ...string) (item *Resource, e error) {
	url := buildIdentityURL(resource, nil, ids...)
	var getResp *getResponse
	if getResp, e = c.identityAPI.getRequest(url, nil); e != nil {
		return
	}

	reader := bytes.NewBuffer(getResp.body)
	decoder := json.NewDecoder(reader)
	item = &Resource{}
	e = decoder.Decode(item)
	return

}

// GetCompartment returns the compartment identified by compartmentID
// See https://docs.us-az-phoenix-1.oracleiaas.com/#apiref.htm
func (c *Client) GetCompartment(compartmentID string) (compartment *Resource, e error) {
	compartment, e = c.getIdentity(resourceCompartments, compartmentID)
	return
}

// GetGroup returns a group identified by groupID
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#getGroup
func (c *Client) GetGroup(groupID string) (group *Resource, e error) {
	group, e = c.getIdentity(resourceGroups, groupID)
	return
}

// GetPolicy returns a policy identified by a policyID
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#getPolicy
func (c *Client) GetPolicy(policyID string) (policy *Policy, e error) {

	url := buildIdentityURL(resourcePolicies, nil, policyID)

	var getResp *getResponse
	if getResp, e = c.identityAPI.getRequest(url, nil); e != nil {
		return
	}

	reader := bytes.NewBuffer(getResp.body)
	decoder := json.NewDecoder(reader)
	policy = &Policy{}
	e = decoder.Decode(policy)
	return

}

// GetUser returns a user identified by userID
// See https://docs.us-az-phoenix-1.oracleiaas.com/#apiref.htm
func (c *Client) GetUser(userID string) (user *Resource, e error) {
	user, e = c.getIdentity(resourceUsers, userID)
	return
}

// GetUserGroupMembership returns a UserGroupMembership identified by id
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#getUserGroupMembership
func (c *Client) GetUserGroupMembership(id string) (userGroupMembership *UserGroupMembership, e error) {

	url := buildIdentityURL(resourceUserGroupMemberships, nil, id)

	var getResp *getResponse
	if getResp, e = c.identityAPI.getRequest(url, nil); e != nil {
		return
	}

	reader := bytes.NewBuffer(getResp.body)
	decoder := json.NewDecoder(reader)
	userGroupMembership = &UserGroupMembership{}
	e = decoder.Decode(userGroupMembership)
	return

}

func (c *Client) ListAvailablityDomains(compartmentID string) (ads []AvailabilityDomain, e error) {
	url := buildIdentityURL(resourceAvailabilityDomains, &url.Values{
		"compartmentId": []string{compartmentID},
	})
	var getResp *getResponse

	if getResp, e = c.identityAPI.getRequest(url, nil); e != nil {
		return
	}

	reader := bytes.NewBuffer(getResp.body)
	decoder := json.NewDecoder(reader)
	e = decoder.Decode(&ads)
	return
}

// ListCompartments returns a list of compartments. The request argument MUST
// supply a compartment ID (tenancy) and MAY contain optional paging arguments
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#listCompartments
func (c *Client) ListCompartments(request *ListResourceRequest) (response *ListResourceResponse, e error) {
	return c.listItems(resourceCompartments, request)
}

// ListGroups returns a list of Groups in a tenancy. The request argument MUST
// supply a compartment ID (tenancy) and MAY contain optional paging arguments
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#listGroups
func (c *Client) ListGroups(request *ListResourceRequest) (response *ListResourceResponse, e error) {
	return c.listItems(resourceGroups, request)
}

func (c *Client) listItems(resource resourceName, request *ListResourceRequest) (resp *ListResourceResponse, e error) {
	var url string
	if url, e = request.buildQuery(resource); e != nil {
		return
	}

	var getResp *getResponse
	if getResp, e = c.identityAPI.getRequest(url, nil); e != nil {
		return
	}

	reader := bytes.NewBuffer(getResp.body)
	decoder := json.NewDecoder(reader)
	var items []Resource
	if e = decoder.Decode(&items); e != nil {
		return
	}

	resp = &ListResourceResponse{
		Page:  getResp.header.Get("opc-next-page"),
		Items: items,
	}

	return
}

// ListUsers returns an array of users for a particular tenancy.  The requestor
// MUST supply a tenancyOCID.  See ListUsersRequest
// for possible query paramters. The response contains an array of users and
// possible a page ID that can be used in subsequent requests.
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/identity.html#listUsers
func (c *Client) ListUsers(request *ListResourceRequest) (response *ListResourceResponse, e error) {
	return c.listItems(resourceUsers, request)
}

func (c *Client) UpdateCompartment(compartmentID string, request UpdateResourceRequest, headers http.Header) (compartment *Resource, e error) {
	var resp []byte
	if resp, e = c.updateResource(resourceCompartments, compartmentID, request, headers); e != nil {
		return
	}

	compartment = &Resource{}
	e = json.Unmarshal(resp, compartment)
	return

}

func (c *Client) updateResource(resource resourceName, resourceID string, request interface{}, headers http.Header) (resp []byte, e error) {
	url := buildIdentityURL(resource, nil, resourceID)
	resp, e = c.identityAPI.request(http.MethodPut, url, request, headers)
	return
}
