// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

import (
	"net/http"
)

// VirtualNetwork describes virtual cloud network
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Vcn/
type VirtualNetwork struct {
	OPCRequestIDUnmarshaller
	ETagUnmarshaller
	CidrBlock             string `json:"cidrBlock"`
	CompartmentID         string `json:"compartmentId"`
	DefaultRouteTableID   string `json:"defaultRouteTableId"`
	DefaultSecurityListID string `json:"defaultSecurityListId"`
	DefaultDHCPOptionsID  string `json:"defaultDhcpOptionsId"`
	DisplayName           string `json:"displayName"`
	DnsLabel              string `json:"dnsLabel"`
	ID                    string `json:"id"`
	State                 string `json:"lifecycleState"`
	TimeCreated           Time   `json:"timeCreated"`
	VcnDomainName         string `json:"vcnDomainName"`
}

// ListVirtualNetworks contains a list of virtual networks
//
type ListVirtualNetworks struct {
	OPCRequestIDUnmarshaller
	NextPageUnmarshaller
	VirtualNetworks []VirtualNetwork
}

func (l *ListVirtualNetworks) GetList() interface{} {
	return &l.VirtualNetworks
}

// CreateVirtualNetwork is used to create a virtual network
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Vcn/CreateVcn
func (c *Client) CreateVirtualNetwork(cidrBlock, compartmentID string, opts *CreateVcnOptions) (vcn *VirtualNetwork, e error) {
	required := struct {
		ocidRequirement
		CidrBlock string `header:"-" json:"cidrBlock,omitempty" url:"-"`
	}{
		CidrBlock: cidrBlock,
	}
	required.CompartmentID = compartmentID

	details := &requestDetails{
		name:     resourceVirtualNetworks,
		optional: opts,
		required: required,
	}

	var resp *response
	if resp, e = c.coreApi.postRequest(details); e != nil {
		return
	}

	vcn = &VirtualNetwork{}
	e = resp.unmarshal(vcn)
	return
}

// GetVirtualNetwork retrieves information about a virtual network
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Vcn/GetVcn
func (c *Client) GetVirtualNetwork(id string) (vcn *VirtualNetwork, e error) {
	details := &requestDetails{
		ids:  urlParts{id},
		name: resourceVirtualNetworks,
	}

	var resp *response
	if resp, e = c.coreApi.getRequest(details); e != nil {
		return
	}

	vcn = &VirtualNetwork{}
	e = resp.unmarshal(vcn)
	return
}

// UpdateVirtualNetwork updates information about a virtual network
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Vcn/UpdateVcn
func (c *Client) UpdateVirtualNetwork(id string, opts *IfMatchDisplayNameOptions) (vcn *VirtualNetwork, e error) {
	details := &requestDetails{
		ids:      urlParts{id},
		name:     resourceVirtualNetworks,
		optional: opts,
	}

	var resp *response
	if resp, e = c.coreApi.request(http.MethodPut, details); e != nil {
		return
	}

	vcn = &VirtualNetwork{}
	e = resp.unmarshal(vcn)
	return
}

// DeleteVirtualNetwork removes a virtual network
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Vcn/DeleteVcn
func (c *Client) DeleteVirtualNetwork(id string, opts *IfMatchOptions) (e error) {
	details := &requestDetails{
		ids:      urlParts{id},
		name:     resourceVirtualNetworks,
		optional: opts,
	}
	return c.coreApi.deleteRequest(details)
}

// ListVirtualNetworks returns a list of virtual networks for a particular
// compartment
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Vcn/ListVcns
func (c *Client) ListVirtualNetworks(compartmentID string, opts *ListOptions) (vcns *ListVirtualNetworks, e error) {
	details := &requestDetails{
		name:     resourceVirtualNetworks,
		optional: opts,
		required: listOCIDRequirement{CompartmentID: compartmentID},
	}

	var resp *response
	if resp, e = c.coreApi.getRequest(details); e != nil {
		return
	}

	vcns = &ListVirtualNetworks{}
	e = resp.unmarshal(vcns)
	return
}
