// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

// CreatePolicy creates a new policy.
import (
	"net/http"
	"time"
)

// Policy returned by GetPolicy and other policy related methods.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/identity.html#Policy
type Policy struct {
	OPCRequestIDUnmarshaller
	ETagUnmarshaller
	CompartmentID  string    `json:"compartmentId"`
	Description    string    `json:"description"`
	ID             string    `json:"id"`
	InactiveStatus uint16    `json:"inactiveStatus"`
	Name           string    `json:"name"`
	State          string    `json:"lifecycleState"`
	Statements     []string  `json:"statements"`
	TimeCreated    time.Time `json:"timeCreated"`
	VersionDate    string    `json:"versionDate"`
}

type ListPolicies struct {
	OPCRequestIDUnmarshaller
	NextPageUnmarshaller
	Policies []Policy
}

func (l *ListPolicies) GetList() interface{} {
	return &l.Policies
}

// CreatePolicy create a policy in a compartment.
//
// For information on defining policies and statements.
// See https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/overview.htm#Example
//
// For information on the CreatePolicy API,
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/identity/20160918/Policy/CreatePolicy
func (c *Client) CreatePolicy(name, desc, compartmentID string, statements []string, opts *CreatePolicyOptions) (res *Policy, e error) {
	required := struct {
		identityCreationRequirement
		Statements []string `header:"-" json:"statements" url:"-"`
	}{
		Statements: statements,
	}
	required.CompartmentID = compartmentID
	required.Description = desc
	required.Name = name

	details := &requestDetails{
		name:     resourcePolicies,
		optional: opts,
		required: required,
	}

	var resp *response
	if resp, e = c.identityApi.postRequest(details); e != nil {
		return
	}

	res = &Policy{}
	e = resp.unmarshal(res)
	return
}

// GetPolicy returns a policy identified by a policyID.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/identity/20160918/Policy/GetPolicy
func (c *Client) GetPolicy(id string) (res *Policy, e error) {
	details := &requestDetails{
		ids:  urlParts{id},
		name: resourcePolicies,
	}

	var resp *response
	if resp, e = c.identityApi.getRequest(details); e != nil {
		return
	}

	res = &Policy{}
	e = resp.unmarshal(res)
	return
}

// UpdatePolicy can be called to modify the description and statements of an existing policy.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/identity/20160918/Policy/UpdatePolicy
func (c *Client) UpdatePolicy(id string, opts *UpdatePolicyOptions) (res *Policy, e error) {
	details := &requestDetails{
		ids:      urlParts{id},
		name:     resourcePolicies,
		optional: opts,
	}

	var resp *response
	if resp, e = c.identityApi.request(http.MethodPut, details); e != nil {
		return
	}

	res = &Policy{}
	e = resp.unmarshal(res)
	return
}

// DeletePolicy removes a policy identified by policyID. Optionally pass an
// etag for optmistic concurrency control.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/identity/20160918/Policy/DeletePolicy
func (c *Client) DeletePolicy(id string, opts *IfMatchOptions) (e error) {
	details := &requestDetails{
		ids:      urlParts{id},
		name:     resourcePolicies,
		optional: opts,
	}

	return c.identityApi.deleteRequest(details)
}

// ListPolicies lists policies by compartmentId
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/identity/20160918/Policy/ListPolicies
func (c *Client) ListPolicies(compartmentID string, opts *ListOptions) (resources *ListPolicies, e error) {
	details := &requestDetails{
		name:     resourcePolicies,
		optional: opts,
		required: listOCIDRequirement{CompartmentID: compartmentID},
	}

	var resp *response
	if resp, e = c.identityApi.getRequest(details); e != nil {
		return
	}

	resources = &ListPolicies{}
	e = resp.unmarshal(resources)
	return
}
