package baremetal

import (
	"net/http"
	"net/url"
)

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
	Code string `json:"code"`
	Type string `json:"type"`
}

// https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/IngressSecurityRule/
type IngressSecurityRule struct {
	ICMPOptions ICMPOptions `json:"icmpOptions"`
	Protocol    string      `json:"protocol"`
	Source      string      `json:"source"`
	TCPOptions  TCPOptions  `json:"tcpOptions"`
	UDPOptions  UDPOptions  `json:"udpOptions"`
}

// https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/EgressSecurityRule/
type EgressSecurityRule struct {
	Destination string      `json:"destination"`
	ICMPOptions ICMPOptions `json:"icmpOptions"`
	Protocol    string      `json:"protocol"`
	TCPOptions  TCPOptions  `json:"tcpOptions"`
	UDPOptions  UDPOptions  `json:"udpOptions"`
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
	opts ...Options,
) (res *SecurityList, e error) {
	body := struct {
		CompartmentID string                `json:"compartmentId"`
		DisplayName   string                `json:"displayName,omitempty"`
		EgressRules   []EgressSecurityRule  `json:"egressSecurityRules"`
		IngressRules  []IngressSecurityRule `json:"ingressSecurityRules"`
		VcnID         string                `json:"vcnId"`
	}{
		CompartmentID: compartmentID,
		EgressRules:   egressRules,
		IngressRules:  ingressRules,
		VcnID:         vcnID,
	}
	if len(opts) > 0 {
		body.DisplayName = opts[0].DisplayName
	}

	reqOpts := &sdkRequestOptions{
		body:    body,
		name:    resourceSecurityLists,
		options: opts,
	}

	var response *requestResponse
	if response, e = c.coreApi.request(http.MethodPost, reqOpts); e != nil {
		return
	}

	res = &SecurityList{}
	e = response.unmarshal(res)
	return
}

// GetSecurityList gets the specified security list's information
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/SecurityList/GetSecurityList
func (c *Client) GetSecurityList(id string, opts ...Options) (res *SecurityList, e error) {
	reqOpts := &sdkRequestOptions{
		name:    resourceSecurityLists,
		options: opts,
		ids:     urlParts{id},
	}
	var resp *requestResponse
	if resp, e = c.coreApi.getRequest(reqOpts); e != nil {
		return
	}

	res = &SecurityList{}
	e = resp.unmarshal(res)
	return
}

// UpdateSecurityList updates the specified security list's rules
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/SecurityList/UpdateSecurityList
func (c *Client) UpdateSecurityList(
	id string,
	egressRules []EgressSecurityRule,
	ingressRules []IngressSecurityRule,
	opts ...Options,
) (res *SecurityList, e error) {
	body := struct {
		DisplayName  string                `json:"displayName,omitempty"`
		EgressRules  []EgressSecurityRule  `json:"egressSecurityRules"`
		IngressRules []IngressSecurityRule `json:"ingressSecurityRules"`
	}{
		EgressRules:  egressRules,
		IngressRules: ingressRules,
	}
	if len(opts) > 0 {
		body.DisplayName = opts[0].DisplayName
	}

	reqOpts := &sdkRequestOptions{
		body:    body,
		name:    resourceSecurityLists,
		options: opts,
		ids:     urlParts{id},
	}

	var response *requestResponse
	if response, e = c.coreApi.request(http.MethodPut, reqOpts); e != nil {
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
func (c *Client) DeleteSecurityList(id string, opts ...Options) (e error) {
	reqOpts := &sdkRequestOptions{
		name:    resourceSecurityLists,
		options: opts,
		ids:     urlParts{id},
	}
	return c.coreApi.deleteRequest(reqOpts)
}

// ListSecurityLists gets a list of the security lists in the specified VCN and
// compartment
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/SecurityList/ListSecurityLists
func (c *Client) ListSecurityLists(compartmentID, vcnID string, opts ...Options) (res *ListSecurityLists, e error) {
	query := url.Values{}
	query.Set(queryVcnID, vcnID)

	reqOpts := &sdkRequestOptions{
		name:    resourceSecurityLists,
		ocid:    compartmentID,
		options: opts,
		query:   query,
	}

	var resp *requestResponse
	if resp, e = c.coreApi.getRequest(reqOpts); e != nil {
		return
	}

	res = &ListSecurityLists{}
	e = resp.unmarshal(res)
	return
}
