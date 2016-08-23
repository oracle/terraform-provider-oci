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
	State             string `json:"state"`
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
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#IPSecConnectionDeviceConfig
type IPSecConnectionDeviceConfig struct {
	RequestableResource
	IPSecConnectionDevice
	Tunnels []TunnelConfig `json:"tunnels"`
}

func (l *IPSecConnectionDeviceConfig) GetList() interface{} {
	return &l.Tunnels
}

// IPSecConnectionDeviceConfig information on a IPSec tunnel status
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#IPSecConnectionDeviceStatus
type IPSecConnectionDeviceStatus struct {
	RequestableResource
	IPSecConnectionDevice
	Tunnels []TunnelStatus `json:"tunnels"`
}

func (l *IPSecConnectionDeviceStatus) GetList() interface{} {
	return &l.Tunnels
}

// IPSecConnection information about an IPSec connection
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#IPSecConnection
type IPSecConnection struct {
	ETaggedResource
	CompartmentID string   `json:"compartmentId"`
	CpeID         string   `json:"cpeId"`
	DisplayName   string   `json:"displayName"`
	DrgID         string   `json:"drgId"`
	ID            string   `json:"id"`
	State         string   `json:"state"`
	StaticRoutes  []string `json:"staticRoutes"`
	TimeCreated   Time     `json:"timeCreated"`
}

// ListIPSecConnections contains a list of IPSec connections as well as
// a NextPage tag that can be passed to a subsequent request for paging.
type ListIPSecConnections struct {
	ResourceContainer
	Connections []IPSecConnection
}

func (l *ListIPSecConnections) GetList() interface{} {
	return &l.Connections
}

// CreateIPSecConnection create an IPSec connection.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#createIPSecConnection
func (c *Client) CreateIPSecConnection(compartmentID, cpeID, drgID string, staticRoutes []string, opts ...Options) (conn *IPSecConnection, e error) {
	var displayName string
	if len(opts) > 0 {
		displayName = opts[0].DisplayName
	}

	body := struct {
		CompartmentID string   `json:"compartmentId"`
		CpeID         string   `json:"cpeId"`
		DisplayName   string   `json:"displayName"`
		DrgID         string   `json:"drgId"`
		StaticRoutes  []string `json:"staticRoutes"`
	}{
		CompartmentID: compartmentID,
		CpeID:         cpeID,
		DisplayName:   displayName,
		DrgID:         drgID,
		StaticRoutes:  staticRoutes,
	}

	req := &sdkRequestOptions{
		name:    resourceIPSecConnections,
		body:    body,
		options: opts,
	}

	var response *requestResponse
	if response, e = c.coreApi.request(http.MethodPost, req); e != nil {
		return
	}

	conn = &IPSecConnection{}
	e = response.unmarshal(conn)

	return
}

// ListIPSecConnections retrieves a list of IPSec connections in a compartment.
// Results can be further refined by optional parameters DrgID and/or CpeID. Paging
// is supported by providing optional Page and Limit parameters.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#listIPSecConnections
func (c *Client) ListIPSecConnections(compartmentID string, opts ...Options) (conns *ListIPSecConnections, e error) {
	req := &sdkRequestOptions{
		name:    resourceIPSecConnections,
		ocid:    compartmentID,
		options: opts,
	}

	var response *requestResponse
	if response, e = c.coreApi.getRequest(req); e != nil {
		return
	}

	conns = &ListIPSecConnections{}
	e = response.unmarshal(conns)
	return
}

// GetIPSecConnection retrieves the IPSec connection identified by id
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#getIPSecConnection
func (c *Client) GetIPSecConnection(id string) (conn *IPSecConnection, e error) {
	req := &sdkRequestOptions{
		name: resourceIPSecConnections,
		ids:  urlParts{id},
	}

	var response *requestResponse
	if response, e = c.coreApi.getRequest(req); e != nil {
		return
	}

	conn = &IPSecConnection{}
	e = response.unmarshal(conn)
	return
}

// DeleteIPSecConnection deletes an IPSec connection.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#deleteIPSecConnection
func (c *Client) DeleteIPSecConnection(id string, opts ...Options) (e error) {
	req := &sdkRequestOptions{
		name:    resourceIPSecConnections,
		ids:     urlParts{id},
		options: opts,
	}

	return c.coreApi.deleteRequest(req)
}

// GetIPSecConnectionDeviceConfig retrieves router configuration to set up
// IPSec tunnel on customer premise device.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#getIPSecConnectionDeviceConfig
func (c *Client) GetIPSecConnectionDeviceConfig(ipsecID string) (config *IPSecConnectionDeviceConfig, e error) {
	req := &sdkRequestOptions{
		name: resourceIPSecConnections,
		ids:  urlParts{ipsecID, deviceConfig},
	}

	var response *requestResponse
	if response, e = c.coreApi.getRequest(req); e != nil {
		return
	}

	config = &IPSecConnectionDeviceConfig{}
	e = response.unmarshal(config)
	return

}

// GetIPSecConnectionDeviceStatus get status on an IPSec tunnel.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#getIPSecConnectionDeviceStatus
func (c *Client) GetIPSecConnectionDeviceStatus(ipsecID string) (status *IPSecConnectionDeviceStatus, e error) {
	req := &sdkRequestOptions{
		name: resourceIPSecConnections,
		ids:  urlParts{ipsecID, deviceStatus},
	}

	var response *requestResponse
	if response, e = c.coreApi.getRequest(req); e != nil {
		return
	}

	status = &IPSecConnectionDeviceStatus{}
	e = response.unmarshal(status)
	return
}
