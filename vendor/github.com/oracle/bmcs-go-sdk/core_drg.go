// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

import "net/http"

// Drg describes a dynamic routing gateway
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Drg/
type Drg struct {
	OPCRequestIDUnmarshaller
	ETagUnmarshaller
	CompartmentID string `json:"compartmentId"`
	DisplayName   string `json:"displayName"`
	ID            string `json:"id"`
	State         string `json:"lifecycleState"`
	TimeCreated   Time   `json:"timeCreated"`
}

// ListDrgs contains a list of gateways
//
type ListDrgs struct {
	OPCRequestIDUnmarshaller
	NextPageUnmarshaller
	Drgs []Drg
}

func (l *ListDrgs) GetList() interface{} {
	return &l.Drgs
}

// CreateDrg is used to create a gateway
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Drg/CreateDrg
func (c *Client) CreateDrg(compartmentID string, opts *CreateOptions) (res *Drg, e error) {
	details := &requestDetails{
		name:     resourceDrgs,
		optional: opts,
		required: ocidRequirement{compartmentID},
	}

	var resp *response
	if resp, e = c.coreApi.postRequest(details); e != nil {
		return
	}

	res = &Drg{}
	e = resp.unmarshal(res)
	return
}

// GetDrg retrieves information about a gateway
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Drg/GetDrg
func (c *Client) GetDrg(id string) (res *Drg, e error) {
	details := &requestDetails{
		name: resourceDrgs,
		ids:  urlParts{id},
	}

	var resp *response
	if resp, e = c.coreApi.getRequest(details); e != nil {
		return
	}

	res = &Drg{}
	e = resp.unmarshal(res)
	return
}

// UpdateDrg updates the specified DRG's display name.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Drg/UpdateDrg
func (c *Client) UpdateDrg(id string, opts *IfMatchDisplayNameOptions) (drg *Drg, e error) {
	details := &requestDetails{
		name:     resourceDrgs,
		ids:      urlParts{id},
		optional: opts,
	}
	var resp *response
	if resp, e = c.coreApi.request(http.MethodPut, details); e != nil {
		return
	}

	drg = &Drg{}
	e = resp.unmarshal(drg)

	return
}

// DeleteDrg removes a gateway
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Drg/DeleteDrg
func (c *Client) DeleteDrg(id string, opts *IfMatchOptions) (e error) {
	details := &requestDetails{
		name:     resourceDrgs,
		ids:      urlParts{id},
		optional: opts,
	}
	return c.coreApi.deleteRequest(details)
}

// ListDrgs returns a list of gateways for a compartment
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Drg/ListDrgs
func (c *Client) ListDrgs(compartmentID string, opts *ListOptions) (res *ListDrgs, e error) {
	details := &requestDetails{
		name:     resourceDrgs,
		required: listOCIDRequirement{compartmentID},
		optional: opts,
	}

	var resp *response
	if resp, e = c.coreApi.getRequest(details); e != nil {
		return
	}

	res = &ListDrgs{}
	e = resp.unmarshal(res)
	return
}
