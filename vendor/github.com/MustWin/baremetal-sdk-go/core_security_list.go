package baremetal

import "net/http"

// https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/PortRange/
type PortRange struct {
	Max uint64 `json:"max"`
	Min uint64 `json:"min"`
}

// https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/UdpOptions/
type UDPOptions struct {
	DestinationPortRange PortRange `json:"destinationPortRange"`
}

// https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/TcpOptions/
type TCPOptions struct {
	DestinationPortRange PortRange `json:"destinationPortRange"`
}

// https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/IcmpOptions/
type ICMPOptions struct {
	Code uint64 `json:"code,omitempty"`
	Type uint64 `json:"type"`
}

// https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/IngressSecurityRule/
type IngressSecurityRule struct {
	ICMPOptions *ICMPOptions `json:"icmpOptions,omitempty"`
	Protocol    string       `json:"protocol"`
	Source      string       `json:"source"`
	TCPOptions  *TCPOptions  `json:"tcpOptions,omitempty"`
	UDPOptions  *UDPOptions  `json:"udpOptions,omitempty"`
}

// https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/EgressSecurityRule/
type EgressSecurityRule struct {
	Destination string       `json:"destination"`
	ICMPOptions *ICMPOptions `json:"icmpOptions,omitempty"`
	Protocol    string       `json:"protocol"`
	TCPOptions  *TCPOptions  `json:"tcpOptions,omitempty"`
	UDPOptions  *UDPOptions  `json:"udpOptions,omitempty"`
}

// SecurityList describes a set of virtual, stateful firewall rules for your VCN
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/SecurityList/
type SecurityList struct {
	ETaggedResource
	CompartmentID        string                `json:"compartmentId"`
	DisplayName          string                `json:"displayName"`
	EgressSecurityRules  []EgressSecurityRule  `json:"egressSecurityRules"`
	ID                   string                `json:"id"`
	IngressSecurityRules []IngressSecurityRule `json:"ingressSecurityRules"`
	State                string                `json:"lifecycleState"`
	TimeCreated          Time                  `json:"timeCreated"`
	VcnID                string                `json:"vcnId"`
}

// ListSecurityLists contains a list of images
//
type ListSecurityLists struct {
	ResourceContainer
	SecurityLists []SecurityList
}

func (l *ListSecurityLists) GetList() interface{} {
	return &l.SecurityLists
}

// CreateSecurityList creates a new security list for the specified VCN
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/SecurityList/CreateSecurityList
func (c *Client) CreateSecurityList(
	compartmentID, vcnID string,
	egressRules []EgressSecurityRule,
	ingressRules []IngressSecurityRule,
	opts *CreateOptions,
) (res *SecurityList, e error) {
	required := struct {
		ocidRequirement
		EgressRules  []EgressSecurityRule  `json:"egressSecurityRules" url:"-"`
		IngressRules []IngressSecurityRule `json:"ingressSecurityRules" url:"-"`
		VcnID        string                `json:"vcnId" url:"-"`
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

	var response *requestResponse
	if response, e = c.coreApi.request(http.MethodPost, details); e != nil {
		return
	}

	res = &SecurityList{}
	e = response.unmarshal(res)
	return
}

// GetSecurityList gets the specified security list's information
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/SecurityList/GetSecurityList
func (c *Client) GetSecurityList(id string) (res *SecurityList, e error) {
	details := &requestDetails{
		name: resourceSecurityLists,
		ids:  urlParts{id},
	}

	var resp *requestResponse
	if resp, e = c.coreApi.getRequest(details); e != nil {
		return
	}

	res = &SecurityList{}
	e = resp.unmarshal(res)
	return
}

// UpdateSecurityList updates the specified security list's rules
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/SecurityList/UpdateSecurityList
func (c *Client) UpdateSecurityList(id string, opts *UpdateSecurityListOptions) (res *SecurityList, e error) {
	details := &requestDetails{
		ids:      urlParts{id},
		name:     resourceSecurityLists,
		optional: opts,
	}

	var response *requestResponse
	if response, e = c.coreApi.request(http.MethodPut, details); e != nil {
		return
	}

	res = &SecurityList{}
	e = response.unmarshal(res)
	return
}

// DeleteSecurityList deletes the specified security list, but only if it's not
// in use
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/SecurityList/DeleteSecurityList
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
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/SecurityList/ListSecurityLists
func (c *Client) ListSecurityLists(compartmentID, vcnID string, opts *ListOptions) (res *ListSecurityLists, e error) {
	required := struct {
		listOCIDRequirement
		VcnID string `json:"-" url:"vcnId"`
	}{
		VcnID: vcnID,
	}
	required.CompartmentID = compartmentID

	details := &requestDetails{
		name:     resourceSecurityLists,
		optional: opts,
		required: required,
	}

	var resp *requestResponse
	if resp, e = c.coreApi.getRequest(details); e != nil {
		return
	}

	res = &ListSecurityLists{}
	e = resp.unmarshal(res)
	return
}
