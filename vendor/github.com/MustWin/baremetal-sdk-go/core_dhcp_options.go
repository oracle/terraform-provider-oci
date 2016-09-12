package baremetal

import (
	"net/http"
	"net/url"
)

// DHCPDNSOption specifies how DNS (host name resolution) is handled in the VCN
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/DHCPDNSOption/
type DHCPDNSOption struct {
	Type             string   `json:"type"`
	CustomDNSServers []string `json:"customDnsServers"`
	ServerType       string   `json:"serverType"`
}

// DHCPOptions contains a set of dhcp options
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/DHCPOptions/
type DHCPOptions struct {
	RequestableResource
	ETaggedResource
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
	ResourceContainer
	DHCPOptions []DHCPOptions
}

func (l *ListDHCPOptions) GetList() interface{} {
	return &l.DHCPOptions
}

// CreateDHCPOptions creates a new set of DHCP options for the specified VCN
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/DHCPOptions/CreateDHCPOptions
func (c *Client) CreateDHCPOptions(compartmentID, vcnID string, dhcpOptions []DHCPDNSOption, opts ...Options) (res *DHCPOptions, e error) {
	body := struct {
		CompartmentID string          `json:"compartmentId"`
		DisplayName   string          `json:"displayName,omitempty"`
		Options       []DHCPDNSOption `json:"options"`
		VcnID         string          `json:"vcnId"`
	}{
		CompartmentID: compartmentID,
		Options:       dhcpOptions,
		VcnID:         vcnID,
	}
	if len(opts) > 0 {
		body.DisplayName = opts[0].DisplayName
	}

	reqOpts := &sdkRequestOptions{
		body:    body,
		name:    resourceDHCPOptions,
		options: opts,
	}

	var response *requestResponse
	if response, e = c.coreApi.request(http.MethodPost, reqOpts); e != nil {
		return
	}

	res = &DHCPOptions{}
	e = response.unmarshal(res)
	return
}

// GetDHCPOptions gets the specified set of DHCP options
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/DHCPOptions/GetDHCPOptions
func (c *Client) GetDHCPOptions(id string) (res *DHCPOptions, e error) {
	reqOpts := &sdkRequestOptions{
		name: resourceDHCPOptions,
		ids:  urlParts{id},
	}

	var resp *requestResponse
	if resp, e = c.coreApi.getRequest(reqOpts); e != nil {
		return
	}

	res = &DHCPOptions{}
	e = resp.unmarshal(res)
	return
}

// UpdateDHCPOptions updates the specified set of DHCP options
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/DHCPOptions/UpdateDHCPOptions
func (c *Client) UpdateDHCPOptions(id string, opts ...Options) (res *DHCPOptions, e error) {
	body := struct {
		Options []DHCPDNSOption `json:"options,omitempty"`
	}{}
	if len(opts) > 0 {
		body.Options = opts[0].DHCPOptions
	}

	reqOpts := &sdkRequestOptions{
		body:    body,
		name:    resourceDHCPOptions,
		options: opts,
		ids:     urlParts{id},
	}

	var response *requestResponse
	if response, e = c.coreApi.request(http.MethodPut, reqOpts); e != nil {
		return
	}

	res = &DHCPOptions{}
	e = response.unmarshal(res)
	return
}

// DeleteDHCPOptions deletes the specified set of DHCP options, but only if it's
// not in use by a subnet
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/DHCPOptions/DeleteDHCPOptions
func (c *Client) DeleteDHCPOptions(id string, opts ...Options) (e error) {
	reqOpts := &sdkRequestOptions{
		ids:     urlParts{id},
		name:    resourceDHCPOptions,
		options: opts,
	}
	return c.coreApi.deleteRequest(reqOpts)
}

// ListDHCPOptions gets a list of the sets of DHCP options in the specified VCN
// and specified compartment
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/DHCPOptions/ListDHCPOptions
func (c *Client) ListDHCPOptions(compartmentID, vcnID string, opts ...Options) (res *ListDHCPOptions, e error) {
	query := url.Values{}
	query.Set(queryVcnID, vcnID)

	reqOpts := &sdkRequestOptions{
		name:    resourceDHCPOptions,
		ocid:    compartmentID,
		options: opts,
		query:   query,
	}

	var resp *requestResponse
	if resp, e = c.coreApi.getRequest(reqOpts); e != nil {
		return
	}

	res = &ListDHCPOptions{}
	e = resp.unmarshal(res)
	return
}
