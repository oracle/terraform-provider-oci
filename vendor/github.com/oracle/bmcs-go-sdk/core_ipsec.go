// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

import "net/http"

// TunnelConfig represents the coniguration for one end of an IPSec tunnel.
type TunnelConfig struct {
	IPAddress    string `json:"ipAddress"`
	SharedSecret string `json:"sharedSecret"`
	TimeCreated  Time   `json:"timeCreated"`
}

// TunnelStatus represents the status of one end of an IPSec tunnel.
type TunnelStatus struct {
	IPAddress         string `json:"ipAddress"`
	State             string `json:"lifecycleState"`
	TimeCreated       Time   `json:"timeCreated"`
	TimeStateModified Time   `json:"timeStateModified"`
}

type IPSecConnectionDevice struct {
	CompartmentID string `json:"compartmentId"`
	ID            string `json:"id"`
	TimeCreated   Time   `json:"timeCreated"`
}

// IPSecConnectionDeviceConfig information to set up an IPSec tunnel.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/IPSecConnectionDeviceConfig/
type IPSecConnectionDeviceConfig struct {
	OPCRequestIDUnmarshaller
	IPSecConnectionDevice
	Tunnels []TunnelConfig `json:"tunnels"`
}

// IPSecConnectionDeviceStatus information on a IPSec tunnel status
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/IPSecConnectionDeviceStatus/
type IPSecConnectionDeviceStatus struct {
	OPCRequestIDUnmarshaller
	IPSecConnectionDevice
	Tunnels []TunnelStatus `json:"tunnels"`
}

// IPSecConnection information about an IPSec connection
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/IPSecConnection/
type IPSecConnection struct {
	OPCRequestIDUnmarshaller
	ETagUnmarshaller
	CompartmentID string   `json:"compartmentId"`
	CpeID         string   `json:"cpeId"`
	DisplayName   string   `json:"displayName"`
	DrgID         string   `json:"drgId"`
	ID            string   `json:"id"`
	State         string   `json:"lifecycleState"`
	StaticRoutes  []string `json:"staticRoutes"`
	TimeCreated   Time     `json:"timeCreated"`
}

// ListIPSecConnections contains a list of IPSec connections as well as
// a NextPage tag that can be passed to a subsequent request for paging.
type ListIPSecConnections struct {
	OPCRequestIDUnmarshaller
	NextPageUnmarshaller
	Connections []IPSecConnection
}

func (l *ListIPSecConnections) GetList() interface{} {
	return &l.Connections
}

// CreateIPSecConnection create an IPSec connection.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/IPSecConnection/CreateIPSecConnection
func (c *Client) CreateIPSecConnection(compartmentID, cpeID, drgID string, staticRoutes []string, opts *CreateOptions) (conn *IPSecConnection, e error) {
	required := struct {
		ocidRequirement
		CpeID        string   `header:"-" json:"cpeId" url:"-"`
		DrgID        string   `header:"-" json:"drgId" url:"-"`
		StaticRoutes []string `header:"-" json:"staticRoutes" url:"-"`
	}{
		CpeID:        cpeID,
		DrgID:        drgID,
		StaticRoutes: staticRoutes,
	}
	required.CompartmentID = compartmentID

	details := &requestDetails{
		name:     resourceIPSecConnections,
		optional: opts,
		required: required,
	}

	var resp *response
	if resp, e = c.coreApi.postRequest(details); e != nil {
		return
	}

	conn = &IPSecConnection{}
	e = resp.unmarshal(conn)
	return
}

// ListIPSecConnections retrieves a list of IPSec connections in a compartment.
// Results can be further refined by optional parameters DrgID and/or CpeID. Paging
// is supported by providing optional Page and Limit parameters.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/IPSecConnection/ListIPSecConnections
func (c *Client) ListIPSecConnections(compartmentID string, opts *ListIPSecConnsOptions) (conns *ListIPSecConnections, e error) {
	details := &requestDetails{
		name:     resourceIPSecConnections,
		optional: opts,
		required: listOCIDRequirement{CompartmentID: compartmentID},
	}

	var resp *response
	if resp, e = c.coreApi.getRequest(details); e != nil {
		return
	}

	conns = &ListIPSecConnections{}
	e = resp.unmarshal(conns)
	return
}

// GetIPSecConnection retrieves the IPSec connection identified by id
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/IPSecConnection/GetIPSecConnection
func (c *Client) GetIPSecConnection(id string) (conn *IPSecConnection, e error) {
	details := &requestDetails{
		name: resourceIPSecConnections,
		ids:  urlParts{id},
	}

	var resp *response
	if resp, e = c.coreApi.getRequest(details); e != nil {
		return
	}

	conn = &IPSecConnection{}
	e = resp.unmarshal(conn)
	return
}

// UpdateIPSecConnection updates the display name for the specified IPSec connection.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/IPSecConnection/UpdateIPSecConnection
func (c *Client) UpdateIPSecConnection(id string, opts *IfMatchDisplayNameOptions) (conn *IPSecConnection, e error) {
	details := &requestDetails{
		name:     resourceIPSecConnections,
		ids:      urlParts{id},
		optional: opts,
	}
	var resp *response
	if resp, e = c.coreApi.request(http.MethodPut, details); e != nil {
		return
	}

	conn = &IPSecConnection{}
	e = resp.unmarshal(conn)

	return
}

// DeleteIPSecConnection deletes an IPSec connection.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/IPSecConnection/DeleteIPSecConnection
func (c *Client) DeleteIPSecConnection(id string, opts *IfMatchOptions) (e error) {
	details := &requestDetails{
		ids:      urlParts{id},
		name:     resourceIPSecConnections,
		optional: opts,
	}

	return c.coreApi.deleteRequest(details)
}

// GetIPSecConnectionDeviceConfig retrieves router configuration to set up
// IPSec tunnel on customer premise device.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/IPSecConnectionDeviceConfig/GetIPSecConnectionDeviceConfig
func (c *Client) GetIPSecConnectionDeviceConfig(id string) (config *IPSecConnectionDeviceConfig, e error) {
	details := &requestDetails{
		name: resourceIPSecConnections,
		ids:  urlParts{id, deviceConfig},
	}

	var resp *response
	if resp, e = c.coreApi.getRequest(details); e != nil {
		return
	}

	config = &IPSecConnectionDeviceConfig{}
	e = resp.unmarshal(config)
	return
}

// GetIPSecConnectionDeviceStatus get status on an IPSec tunnel.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/IPSecConnectionDeviceStatus/GetIPSecConnectionDeviceStatus
func (c *Client) GetIPSecConnectionDeviceStatus(id string) (status *IPSecConnectionDeviceStatus, e error) {
	details := &requestDetails{
		name: resourceIPSecConnections,
		ids:  urlParts{id, deviceStatus},
	}

	var resp *response
	if resp, e = c.coreApi.getRequest(details); e != nil {
		return
	}

	status = &IPSecConnectionDeviceStatus{}
	e = resp.unmarshal(status)
	return
}
