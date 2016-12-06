package baremetal

import "net/http"

type RouteRule struct {
	CidrBlock         string            `json:"cidrBlock"`
	DisplayName       string            `json:"displayName,omitempty"`
	NetworkEntityID   string            `json:"networkEntityId"`
	NetworkEntityType NetworkEntityType `json:"networkEntityType"`
	TimeCreated       Time              `json:"timeCreated,omitempty"`
}

// RouteTable describes a route table
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/RouteTable/
type RouteTable struct {
	ETaggedResource
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
	ResourceContainer
	RouteTables []RouteTable
}

func (l *ListRouteTables) GetList() interface{} {
	return &l.RouteTables
}

// CreateRouteTable is used to create a route table
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/RouteTable/CreateRouteTable
func (c *Client) CreateRouteTable(compartmentID, vcnID string, routeRules []RouteRule, opts *CreateOptions) (res *RouteTable, e error) {
	required := struct {
		ocidRequirement
		RouteRules []RouteRule `json:"routeRules" url:"-"`
		VcnID      string      `json:"vcnId" url:"-"`
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

	var response *requestResponse
	if response, e = c.coreApi.request(http.MethodPost, details); e != nil {
		return
	}

	res = &RouteTable{}
	e = response.unmarshal(res)
	return
}

// GetRouteTable is used to get information about a route table
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/RouteTable/GetRouteTable
func (c *Client) GetRouteTable(id string) (res *RouteTable, e error) {
	details := &requestDetails{
		name: resourceRouteTables,
		ids:  urlParts{id},
	}

	var response *requestResponse
	if response, e = c.coreApi.getRequest(details); e != nil {
		return
	}

	res = &RouteTable{}
	e = response.unmarshal(res)
	return
}

// UpdateRouteTable is used to update a route table
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/RouteTable/UpdateRouteTable
func (c *Client) UpdateRouteTable(id string, opts *UpdateRouteTableOptions) (res *RouteTable, e error) {
	details := &requestDetails{
		ids:      urlParts{id},
		name:     resourceRouteTables,
		optional: opts,
	}

	var response *requestResponse
	if response, e = c.coreApi.request(http.MethodPut, details); e != nil {
		return
	}

	res = &RouteTable{}
	e = response.unmarshal(res)
	return
}

// DeleteRouteTable is used to delete a route table
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/RouteTable/DeleteRouteTable
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
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/RouteTable/ListRouteTables
func (c *Client) ListRouteTables(compartmentID, vcnID string, opts *ListOptions) (res *ListRouteTables, e error) {
	required := struct {
		listOCIDRequirement
		VcnID string `json:"-" url:"vcnId"`
	}{
		VcnID: vcnID,
	}
	required.CompartmentID = compartmentID

	details := &requestDetails{
		name:     resourceRouteTables,
		optional: opts,
		required: required,
	}

	var response *requestResponse
	if response, e = c.coreApi.getRequest(details); e != nil {
		return
	}

	res = &ListRouteTables{}
	e = response.unmarshal(res)
	return
}
