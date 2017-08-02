// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

import "net/http"

// Subnet represents a network subnet
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Subnet/
type Subnet struct {
	OPCRequestIDUnmarshaller
	ETagUnmarshaller
	AvailabilityDomain     string   `json:"availabilityDomain"`
	CIDRBlock              string   `json:"cidrBlock"`
	CompartmentID          string   `json:"compartmentId"`
	DisplayName            string   `json:"displayName"`
	DHCPOptionsID          string   `json:"dhcpOptionsId"`
	DNSLabel               string   `json:"dnsLabel"`
	ID                     string   `json:"id"`
	RouteTableID           string   `json:"routeTableId"`
	SecurityListIDs        []string `json:"securityListIds"`
	State                  string   `json:"lifecycleState"`
	SubnetDomainName       string   `json:"subnetDomainName"`
	TimeCreated            Time     `json:"timeCreated"`
	VcnID                  string   `json:"vcnId"`
	ProhibitPublicIpOnVnic bool     `json:"prohibitPublicIpOnVnic"`
	VirtualRouterIP        string   `json:"virtualRouterIp"`
	VirtualRouterMac       string   `json:"virtualRouterMac"`
}

// ListSubnets contains a list of Subnet
type ListSubnets struct {
	OPCRequestIDUnmarshaller
	NextPageUnmarshaller
	Subnets []Subnet
}

func (l *ListSubnets) GetList() interface{} {
	return &l.Subnets
}

// CreateSubnet will create a new subnet.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Subnet/CreateSubnet
func (c *Client) CreateSubnet(
	availabilityDomain,
	cidrBlock,
	compartmentID,
	vcnID string,
	opts *CreateSubnetOptions,
) (sn *Subnet, e error) {

	required := struct {
		ocidRequirement
		AvailabilityDomain string `header:"-" json:"availabilityDomain" url:"-"`
		CIDRBlock          string `header:"-" json:"cidrBlock" url:"-"`
		VcnID              string `header:"-" json:"vcnId" url:"-"`
	}{
		AvailabilityDomain: availabilityDomain,
		CIDRBlock:          cidrBlock,
		VcnID:              vcnID,
	}
	required.CompartmentID = compartmentID

	details := &requestDetails{
		name:     resourceSubnets,
		optional: opts,
		required: required,
	}

	var resp *response
	if resp, e = c.coreApi.postRequest(details); e != nil {
		return
	}

	sn = &Subnet{}
	e = resp.unmarshal(sn)
	return
}

// GetSubnet will retrieve Subnet for subnetID
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Subnet/GetSubnet
func (c *Client) GetSubnet(id string) (subnet *Subnet, e error) {
	details := &requestDetails{
		ids:  urlParts{id},
		name: resourceSubnets,
	}

	var resp *response
	if resp, e = c.coreApi.getRequest(details); e != nil {
		return
	}

	subnet = &Subnet{}
	e = resp.unmarshal(subnet)
	return
}

// UpdateSubnet updates the display name for the specified Subnet
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Subnet/UpdateSubnet
func (c *Client) UpdateSubnet(id string, opts *IfMatchDisplayNameOptions) (subnet *Subnet, e error) {
	details := &requestDetails{
		name:     resourceSubnets,
		ids:      urlParts{id},
		optional: opts,
	}
	var resp *response
	if resp, e = c.coreApi.request(http.MethodPut, details); e != nil {
		return
	}

	subnet = &Subnet{}
	e = resp.unmarshal(subnet)

	return
}

// DeleteSubnet will delete a subnet with subnetID
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Subnet/DeleteSubnet
func (c *Client) DeleteSubnet(id string, opts *IfMatchOptions) error {
	details := &requestDetails{
		ids:      urlParts{id},
		name:     resourceSubnets,
		optional: opts,
	}

	return c.coreApi.deleteRequest(details)
}

// ListSubnets returns a list of subnets in compartment for a virtual cloud network.
// The size of results may be limited by assigning values to the Limit field of
// Options.  Results may be paged by assigning the NewPage from the last
// response to the Page member of Options.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Subnet/ListSubnets
func (c *Client) ListSubnets(compartmentID, vcnID string, opts *ListOptions) (subnets *ListSubnets, e error) {
	required := struct {
		listOCIDRequirement
		VcnID string `header:"-" json:"-" url:"vcn"`
	}{
		VcnID: vcnID,
	}
	required.CompartmentID = compartmentID

	details := &requestDetails{
		name:     resourceSubnets,
		optional: opts,
		required: required,
	}

	var resp *response
	if resp, e = c.coreApi.getRequest(details); e != nil {
		return
	}

	subnets = &ListSubnets{}
	e = resp.unmarshal(subnets)
	return
}
