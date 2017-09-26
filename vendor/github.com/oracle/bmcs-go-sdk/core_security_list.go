// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

import "net/http"

// PortRange specifies a set of ports for UDPOptions or TCPOptions
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/PortRange/
type PortRange struct {
	Max uint64 `header:"-" json:"max" url:"-"`
	Min uint64 `header:"-" json:"min" url:"-"`
}

// UDPOptions specifies ports for a UDP rule
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/UdpOptions/
type UDPOptions struct {
	DestinationPortRange PortRange `header:"-" json:"destinationPortRange" url:"-"`
}

// TCPOptions specifies ports for a TCP rule
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/TcpOptions/
type TCPOptions struct {
	DestinationPortRange PortRange `header:"-" json:"destinationPortRange" url:"-"`
}

// ICMPOptions specifies a particular ICMP type and code
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/IcmpOptions/
type ICMPOptions struct {
	Code uint64 `header:"-" json:"code,omitempty" url:"-"`
	Type uint64 `header:"-" json:"type" url:"-"`
}

// IngressSecurityRule is a rule for allowing inbound IP packets.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/IngressSecurityRule/
type IngressSecurityRule struct {
	ICMPOptions *ICMPOptions `header:"-" json:"icmpOptions,omitempty" url:"-"`
	Protocol    string       `header:"-" json:"protocol" url:"-"`
	Source      string       `header:"-" json:"source" url:"-"`
	TCPOptions  *TCPOptions  `header:"-" json:"tcpOptions,omitempty" url:"-"`
	UDPOptions  *UDPOptions  `header:"-" json:"udpOptions,omitempty" url:"-"`
	IsStateless bool         `header:"-" json:"isStateless" url:"-"`
}

// EgressSecurityRule is a rule for allowing outbound IP packets.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/EgressSecurityRule/
type EgressSecurityRule struct {
	Destination string       `header:"-" json:"destination" url:"-"`
	ICMPOptions *ICMPOptions `header:"-" json:"icmpOptions,omitempty" url:"-"`
	Protocol    string       `header:"-" json:"protocol" url:"-"`
	TCPOptions  *TCPOptions  `header:"-" json:"tcpOptions,omitempty" url:"-"`
	UDPOptions  *UDPOptions  `header:"-" json:"udpOptions,omitempty" url:"-"`
	IsStateless bool         `header:"-" json:"isStateless" url:"-"`
}

// SecurityList describes a set of virtual, stateful firewall rules for your VCN
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/SecurityList/
type SecurityList struct {
	ETagUnmarshaller
	OPCRequestIDUnmarshaller
	CompartmentID        string                `json:"compartmentId"`
	DisplayName          string                `json:"displayName"`
	EgressSecurityRules  []EgressSecurityRule  `json:"egressSecurityRules"`
	ID                   string                `json:"id"`
	IngressSecurityRules []IngressSecurityRule `json:"ingressSecurityRules"`
	State                string                `json:"lifecycleState"`
	TimeCreated          Time                  `json:"timeCreated"`
	VcnID                string                `json:"vcnId"`
}

// ListSecurityLists is the response from a ListSecurityLists() request
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/SecurityList/ListSecurityLists
type ListSecurityLists struct {
	NextPageUnmarshaller
	OPCRequestIDUnmarshaller
	SecurityLists []SecurityList
}

func (l *ListSecurityLists) GetList() interface{} {
	return &l.SecurityLists
}

// CreateSecurityList creates a new security list for the specified VCN
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/SecurityList/CreateSecurityList
func (c *Client) CreateSecurityList(
	compartmentID, vcnID string,
	egressRules []EgressSecurityRule,
	ingressRules []IngressSecurityRule,
	opts *CreateOptions,
) (res *SecurityList, e error) {
	required := struct {
		ocidRequirement
		EgressRules  []EgressSecurityRule  `header:"-" json:"egressSecurityRules" url:"-"`
		IngressRules []IngressSecurityRule `header:"-" json:"ingressSecurityRules" url:"-"`
		VcnID        string                `header:"-" json:"vcnId" url:"-"`
	}{
		EgressRules:  egressRules,
		IngressRules: ingressRules,
		VcnID:        vcnID,
	}
	required.CompartmentID = compartmentID

	details := &requestDetails{
		name:     resourceSecurityLists,
		optional: opts,
		required: required,
	}

	var resp *response
	if resp, e = c.coreApi.postRequest(details); e != nil {
		return
	}

	res = &SecurityList{}
	e = resp.unmarshal(res)
	return
}

// GetSecurityList gets the specified security list's information
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/SecurityList/GetSecurityList
func (c *Client) GetSecurityList(id string) (res *SecurityList, e error) {
	details := &requestDetails{
		name: resourceSecurityLists,
		ids:  urlParts{id},
	}

	var resp *response
	if resp, e = c.coreApi.getRequest(details); e != nil {
		return
	}

	res = &SecurityList{}
	e = resp.unmarshal(res)
	return
}

// UpdateSecurityList updates the specified security list's rules
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/SecurityList/UpdateSecurityList
func (c *Client) UpdateSecurityList(
	id string,
	opts *UpdateSecurityListOptions,
) (res *SecurityList, e error) {
	details := &requestDetails{
		ids:      urlParts{id},
		name:     resourceSecurityLists,
		optional: opts,
	}

	var resp *response
	if resp, e = c.coreApi.request(http.MethodPut, details); e != nil {
		return
	}

	res = &SecurityList{}
	e = resp.unmarshal(res)
	return
}

// DeleteSecurityList deletes the specified security list, but only if it's not
// in use
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/SecurityList/DeleteSecurityList
func (c *Client) DeleteSecurityList(id string, opts *IfMatchOptions) (e error) {
	details := &requestDetails{
		ids:      urlParts{id},
		name:     resourceSecurityLists,
		optional: opts,
	}
	return c.coreApi.deleteRequest(details)
}

// ListSecurityLists gets a list of the security lists in the specified VCN and
// compartment
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/SecurityList/ListSecurityLists
func (c *Client) ListSecurityLists(compartmentID, vcnID string, opts *ListOptions) (res *ListSecurityLists, e error) {
	required := struct {
		listOCIDRequirement
		VcnID string `header:"-" json:"-" url:"vcnId"`
	}{
		VcnID: vcnID,
	}
	required.CompartmentID = compartmentID

	details := &requestDetails{
		name:     resourceSecurityLists,
		optional: opts,
		required: required,
	}

	var resp *response
	if resp, e = c.coreApi.getRequest(details); e != nil {
		return
	}

	res = &ListSecurityLists{}
	e = resp.unmarshal(res)
	return
}
