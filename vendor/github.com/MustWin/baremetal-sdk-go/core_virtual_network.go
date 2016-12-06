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

// CreateVirtualNetwork is used to create a virtual network
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/Vcn/CreateVcn
func (c *Client) CreateVirtualNetwork(cidrBlock, compartmentID string, opts *CreateOptions) (vcn *VirtualNetwork, e error) {
	required := struct {
		ocidRequirement
		CidrBlock string
	}{
		CidrBlock: cidrBlock,
	}
	required.CompartmentID = compartmentID

	details := &requestDetails{
		name:     resourceVirtualNetworks,
		optional: opts,
		required: required,
	}

	var response *requestResponse
	if response, e = c.coreApi.request(http.MethodPost, details); e != nil {
		return
	}

	vcn = &VirtualNetwork{}
	e = response.unmarshal(vcn)
	return
}

// GetVirtualNetwork retrieves information about a virtual network
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/Vcn/GetVcn
func (c *Client) GetVirtualNetwork(id string) (vcn *VirtualNetwork, e error) {
	details := &requestDetails{
		ids:  urlParts{id},
		name: resourceVirtualNetworks,
	}

	var resp *requestResponse
	if resp, e = c.coreApi.getRequest(details); e != nil {
		return
	}

	vcn = &VirtualNetwork{}
	e = resp.unmarshal(vcn)
	return
}

// DeleteVirtualNetwork removes a virtual network
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/Vcn/DeleteVcn
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
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/Vcn/ListVcns
func (c *Client) ListVirtualNetworks(compartmentID string, opts *ListOptions) (vcns *ListVirtualNetworks, e error) {
	details := &requestDetails{
		name:     resourceVirtualNetworks,
		optional: opts,
		required: listOCIDRequirement{CompartmentID: compartmentID},
	}

	var resp *requestResponse
	if resp, e = c.coreApi.getRequest(details); e != nil {
		return
	}

	vcns = &ListVirtualNetworks{}
	e = resp.unmarshal(vcns)
	return
}
