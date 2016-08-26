package baremetal

import (
	"net/http"
	"net/url"
)

type RouteRule struct {
	CidrBlock         string            `json:"cidrBlock"`
	DisplayName       string            `json:"displayName"`
	NetworkEntityID   string            `json:"networkEntityId"`
	NetworkEntityType networkEntityType `json:"networkEntityType"`
	TimeCreated       Time              `json:"timeCreated"`
}

// RouteTable describes a route table
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#RouteTable
type RouteTable struct {
	ETaggedResource
	CompartmentID string      `json:"compartmentId"`
	DisplayName   string      `json:"displayName"`
	ID            string      `json:"id"`
	TimeModified  Time        `json:"timeModified"`
	RouteRules    []RouteRule `json:"routeRules"`
	State         string      `json:"state"`
	TimeCreated   Time        `json:"timeCreated"`
}

// ListRouteTables contains a list of route tables
type ListRouteTables struct {
	ResourceContainer
	RouteTables []RouteTable
}

func (l *ListRouteTables) GetList() interface{} {
	return &l.RouteTables
}

// CreateRouteTable is used to create a route table
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#createRouteTable
func (c *Client) CreateRouteTable(compartmentID, vcnID string, routeRules []RouteRule, opts ...Options) (res *RouteTable, e error) {
	var displayName string
	if len(opts) > 0 {
		displayName = opts[0].DisplayName
	}

	body := struct {
		CompartmentID string      `json:"compartmentId"`
		DisplayName   string      `json:"displayName"`
		RouteRules    []RouteRule `json:"routeRules"`
		VcnID         string      `json:"vcnId"`
	}{
		CompartmentID: compartmentID,
		DisplayName:   displayName,
		RouteRules:    routeRules,
		VcnID:         vcnID,
	}

	req := &sdkRequestOptions{
		name:    resourceRouteTables,
		body:    body,
		options: opts,
	}

	var response *requestResponse
	if response, e = c.coreApi.request(http.MethodPost, req); e != nil {
		return
	}

	res = &RouteTable{}
	e = response.unmarshal(res)

	return
}

// GetRouteTable is used to get information about a route table
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#getRouteTable
func (c *Client) GetRouteTable(id string) (res *RouteTable, e error) {
	req := &sdkRequestOptions{
		name: resourceRouteTables,
		ids:  urlParts{id},
	}

	var response *requestResponse
	if response, e = c.coreApi.getRequest(req); e != nil {
		return
	}

	res = &RouteTable{}
	e = response.unmarshal(res)
	return
}

// UpdateRouteTable is used to update a route table
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#updateRouteTable
func (c *Client) UpdateRouteTable(id string, routeRules []RouteRule, opts ...Options) (res *RouteTable, e error) {
	body := struct {
		RouteRules []RouteRule `json:"routeRules"`
	}{
		RouteRules: routeRules,
	}

	reqOpts := &sdkRequestOptions{
		body:    body,
		name:    resourceRouteTables,
		options: opts,
		ids:     urlParts{id},
	}

	var response *requestResponse
	if response, e = c.coreApi.request(http.MethodPut, reqOpts); e != nil {
		return
	}

	res = &RouteTable{}
	e = response.unmarshal(res)
	return
}

// DeleteRouteTable is used to delete a route table
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#deleteRouteTable
func (c *Client) DeleteRouteTable(id string, opts ...Options) (e error) {
	req := &sdkRequestOptions{
		name:    resourceRouteTables,
		ids:     urlParts{id},
		options: opts,
	}

	return c.coreApi.deleteRequest(req)
}

// ListRouteTables is used to list route tables in a given compartment and vcn
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#listRouteTables
func (c *Client) ListRouteTables(compartmentID, vcnID string, opts ...Options) (res *ListRouteTables, e error) {
	query := url.Values{}
	query.Set(queryVcnID, vcnID)

	req := &sdkRequestOptions{
		name:    resourceRouteTables,
		ocid:    compartmentID,
		options: opts,
		query:   query,
	}

	var response *requestResponse
	if response, e = c.coreApi.getRequest(req); e != nil {
		return
	}

	res = &ListRouteTables{}
	e = response.unmarshal(res)
	return
}
