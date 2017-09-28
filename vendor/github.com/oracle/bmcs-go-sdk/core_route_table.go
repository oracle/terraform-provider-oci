// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

import "net/http"

type RouteRule struct {
	CidrBlock       string `json:"cidrBlock"`
	NetworkEntityID string `json:"networkEntityId"`
}

// RouteTable describes a route table
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/RouteTable/
type RouteTable struct {
	OPCRequestIDUnmarshaller
	ETagUnmarshaller
	CompartmentID string      `json:"compartmentId"`
	DisplayName   string      `json:"displayName"`
	ID            string      `json:"id"`
	TimeModified  Time        `json:"timeModified"`
	RouteRules    []RouteRule `json:"routeRules"`
	State         string      `json:"lifecycleState"`
	TimeCreated   Time        `json:"timeCreated"`
}

// ListRouteTables contains a list of route tables
type ListRouteTables struct {
	OPCRequestIDUnmarshaller
	NextPageUnmarshaller
	RouteTables []RouteTable
}

func (l *ListRouteTables) GetList() interface{} {
	return &l.RouteTables
}

// CreateRouteTable is used to create a route table
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/RouteTable/CreateRouteTable
func (c *Client) CreateRouteTable(compartmentID, vcnID string, routeRules []RouteRule, opts *CreateOptions) (res *RouteTable, e error) {
	required := struct {
		ocidRequirement
		RouteRules []RouteRule `header:"-" json:"routeRules" url:"-"`
		VcnID      string      `header:"-" json:"vcnId" url:"-"`
	}{
		RouteRules: routeRules,
		VcnID:      vcnID,
	}
	required.CompartmentID = compartmentID

	details := &requestDetails{
		name:     resourceRouteTables,
		optional: opts,
		required: required,
	}

	var resp *response
	if resp, e = c.coreApi.postRequest(details); e != nil {
		return
	}

	res = &RouteTable{}
	e = resp.unmarshal(res)
	return
}

// GetRouteTable is used to get information about a route table
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/RouteTable/GetRouteTable
func (c *Client) GetRouteTable(id string) (res *RouteTable, e error) {
	details := &requestDetails{
		name: resourceRouteTables,
		ids:  urlParts{id},
	}

	var resp *response
	if resp, e = c.coreApi.getRequest(details); e != nil {
		return
	}

	res = &RouteTable{}
	e = resp.unmarshal(res)
	return
}

// UpdateRouteTable is used to update a route table
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/RouteTable/UpdateRouteTable
func (c *Client) UpdateRouteTable(id string, opts *UpdateRouteTableOptions) (res *RouteTable, e error) {
	details := &requestDetails{
		ids:      urlParts{id},
		name:     resourceRouteTables,
		optional: opts,
	}

	var resp *response
	if resp, e = c.coreApi.request(http.MethodPut, details); e != nil {
		return
	}

	res = &RouteTable{}
	e = resp.unmarshal(res)
	return
}

// DeleteRouteTable is used to delete a route table
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/RouteTable/DeleteRouteTable
func (c *Client) DeleteRouteTable(id string, opts *IfMatchOptions) (e error) {
	details := &requestDetails{
		ids:      urlParts{id},
		name:     resourceRouteTables,
		optional: opts,
	}

	return c.coreApi.deleteRequest(details)
}

// ListRouteTables is used to list route tables in a given compartment and vcn
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/RouteTable/ListRouteTables
func (c *Client) ListRouteTables(compartmentID, vcnID string, opts *ListOptions) (res *ListRouteTables, e error) {
	required := struct {
		listOCIDRequirement
		VcnID string `header:"-" json:"-" url:"vcnId"`
	}{
		VcnID: vcnID,
	}
	required.CompartmentID = compartmentID

	details := &requestDetails{
		name:     resourceRouteTables,
		optional: opts,
		required: required,
	}

	var resp *response
	if resp, e = c.coreApi.getRequest(details); e != nil {
		return
	}

	res = &ListRouteTables{}
	e = resp.unmarshal(res)
	return
}
