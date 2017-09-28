// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

import "net/http"

// DHCPDNSOption specifies how DNS (host name resolution) is handled in the VCN
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/DhcpDnsOption/
type DHCPDNSOption struct {
	Type              string   `json:"type"`
	CustomDNSServers  []string `json:"customDnsServers,omitempty"`
	ServerType        string   `json:"serverType,omitempty"`
	SearchDomainNames []string `json:"searchDomainNames,omitempty"`
}

// DHCPOptions contains a set of dhcp options
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/DhcpOptions/
type DHCPOptions struct {
	OPCRequestIDUnmarshaller
	ETagUnmarshaller
	CompartmentID string          `json:"compartmentId"`
	DisplayName   string          `json:"displayName"`
	ID            string          `json:"id"`
	Options       []DHCPDNSOption `json:"options"`
	State         string          `json:"lifecycleState"`
	TimeCreated   Time            `json:"timeCreated"`
}

// ListDHCPOptions contains a list of dhcp options
//
type ListDHCPOptions struct {
	OPCRequestIDUnmarshaller
	NextPageUnmarshaller
	DHCPOptions []DHCPOptions
}

func (l *ListDHCPOptions) GetList() interface{} {
	return &l.DHCPOptions
}

// CreateDHCPOptions creates a new set of DHCP options for the specified VCN
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/DhcpOptions/CreateDhcpOptions
func (c *Client) CreateDHCPOptions(compartmentID, vcnID string, dhcpOptions []DHCPDNSOption, opts *CreateOptions) (res *DHCPOptions, e error) {
	required := struct {
		ocidRequirement
		Options []DHCPDNSOption `header:"-" json:"options" url:"-"`
		VcnID   string          `header:"-" json:"vcnId" url:"-"`
	}{
		Options: dhcpOptions,
		VcnID:   vcnID,
	}
	required.CompartmentID = compartmentID

	details := &requestDetails{
		name:     resourceDHCPOptions,
		optional: opts,
		required: required,
	}

	var resp *response
	if resp, e = c.coreApi.postRequest(details); e != nil {
		return
	}

	res = &DHCPOptions{}
	e = resp.unmarshal(res)
	return
}

// GetDHCPOptions gets the specified set of DHCP options
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/DhcpOptions/GetDhcpOptions
func (c *Client) GetDHCPOptions(id string) (res *DHCPOptions, e error) {
	details := &requestDetails{
		name: resourceDHCPOptions,
		ids:  urlParts{id},
	}

	var resp *response
	if resp, e = c.coreApi.getRequest(details); e != nil {
		return
	}

	res = &DHCPOptions{}
	e = resp.unmarshal(res)
	return
}

// UpdateDHCPOptions updates the specified set of DHCP options
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/DhcpOptions/UpdateDhcpOptions
func (c *Client) UpdateDHCPOptions(id string, opts *UpdateDHCPDNSOptions) (res *DHCPOptions, e error) {
	details := &requestDetails{
		name:     resourceDHCPOptions,
		ids:      urlParts{id},
		optional: opts,
	}

	var resp *response
	if resp, e = c.coreApi.request(http.MethodPut, details); e != nil {
		return
	}

	res = &DHCPOptions{}
	e = resp.unmarshal(res)
	return
}

// DeleteDHCPOptions deletes the specified set of DHCP options, but only if it's
// not in use by a subnet
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/DhcpOptions/DeleteDhcpOptions
func (c *Client) DeleteDHCPOptions(id string, opts *IfMatchOptions) (e error) {
	details := &requestDetails{
		name:     resourceDHCPOptions,
		ids:      urlParts{id},
		optional: opts,
	}
	return c.coreApi.deleteRequest(details)
}

// ListDHCPOptions gets a list of the sets of DHCP options in the specified VCN
// and specified compartment
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/DhcpOptions/ListDhcpOptions
func (c *Client) ListDHCPOptions(compartmentID, vcnID string, opts *ListOptions) (res *ListDHCPOptions, e error) {
	required := struct {
		listOCIDRequirement
		VcnID string `header:"-" json:"-" url:"vcnId"`
	}{
		VcnID: vcnID,
	}
	required.CompartmentID = compartmentID

	details := &requestDetails{
		name:     resourceDHCPOptions,
		required: required,
		optional: opts,
	}

	var resp *response
	if resp, e = c.coreApi.getRequest(details); e != nil {
		return
	}

	res = &ListDHCPOptions{}
	e = resp.unmarshal(res)
	return
}
