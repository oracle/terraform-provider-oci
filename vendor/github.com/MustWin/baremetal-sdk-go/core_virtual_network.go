package baremetal

import "net/http"

// VirtualNetwork describes virtual cloud network
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/Vcn/
type VirtualNetwork struct {
	ETaggedResource
	CidrBlock             string `json:"cidrBlock"`
	CompartmentID         string `json:"compartmentId"`
	DefaultRoutingTableID string `json:"defaultRouteTableId"`
	DefaultSecurityListID string `json:"defaultSecurityListId"`
	DisplayName           string `json:"displayName"`
	ID                    string `json:"id"`
	State                 string `json:"lifecycleState"`
	TimeCreated           Time   `json:"timeCreated"`
}

// ListVirtualNetworks contains a list of virtual networks
//
type ListVirtualNetworks struct {
	ResourceContainer
	VirtualNetworks []VirtualNetwork
}

func (l *ListVirtualNetworks) GetList() interface{} {
	return &l.VirtualNetworks
}

// CreateVirtualNeworkRequest describes the body of a virtual network create
// request
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/Vcn/CreateVcn
// TODO: This has been changed to CreateVirtualNetworkDetails
type CreateVirtualNetworkRequest struct {
	CidrBlock     string `json:"cidrBlock"`
	CompartmentID string `json:"compartmentId"`
	DisplayName   string `json:"displayName,omitempty"`
}

// CreateVirtualNetwork is used to create a virtual network
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/Vcn/CreateVcn
func (c *Client) CreateVirtualNetwork(cidrBlock, compartmentID string, opts ...Options) (vcn *VirtualNetwork, e error) {
	createRequest := CreateVirtualNetworkRequest{
		CidrBlock:     cidrBlock,
		CompartmentID: compartmentID,
	}

	if len(opts) > 0 {
		createRequest.DisplayName = opts[0].DisplayName
	}

	reqOpts := &sdkRequestOptions{
		body:    createRequest,
		name:    resourceVirtualNetworks,
		options: opts,
	}

	var response *requestResponse
	if response, e = c.coreApi.request(http.MethodPost, reqOpts); e != nil {
		return
	}

	vcn = &VirtualNetwork{}
	e = response.unmarshal(vcn)
	return
}

// GetVirtualNetwork retrieves information about a virtual network
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/Vcn/GetVcn
func (c *Client) GetVirtualNetwork(id string, opts ...Options) (vcn *VirtualNetwork, e error) {
	reqOpts := &sdkRequestOptions{
		name:    resourceVirtualNetworks,
		options: opts,
		ids:     urlParts{id},
	}
	var resp *requestResponse
	if resp, e = c.coreApi.getRequest(reqOpts); e != nil {
		return
	}

	vcn = &VirtualNetwork{}
	e = resp.unmarshal(vcn)
	return
}

// DeleteVirtualNetwork removes a virtual network
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/Vcn/DeleteVcn
func (c *Client) DeleteVirtualNetwork(id string, opts ...Options) (e error) {
	reqOpts := &sdkRequestOptions{
		name:    resourceVirtualNetworks,
		options: opts,
		ids:     urlParts{id},
	}
	return c.coreApi.deleteRequest(reqOpts)
}

// ListVirtualNetworks returns a list of virtual networks for a particular
// compartment
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/Vcn/ListVcns
func (c *Client) ListVirtualNetworks(compartmentID string, opts ...Options) (vcns *ListVirtualNetworks, e error) {
	reqOpts := &sdkRequestOptions{
		name:    resourceVirtualNetworks,
		ocid:    compartmentID,
		options: opts,
	}

	var resp *requestResponse
	if resp, e = c.coreApi.getRequest(reqOpts); e != nil {
		return
	}

	vcns = &ListVirtualNetworks{}
	e = resp.unmarshal(vcns)
	return
}
