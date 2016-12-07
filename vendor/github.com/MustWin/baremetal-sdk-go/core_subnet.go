package baremetal

import "net/http"

// Subnet represents a network subnet
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/Subnet/
type Subnet struct {
	ETaggedResource
	AvailabilityDomain string   `json:"availabilityDomain"`
	CIDRBlock          string   `json:"cidrBlock"`
	CompartmentID      string   `json:"compartmentId"`
	DisplayName        string   `json:"displayName"`
	ID                 string   `json:"id"`
	RouteTableID       string   `json:"routeTableId"`
	SecurityListIDs    []string `json:"securityListIds"`
	State              string   `json:"lifecycleState"`
	TimeCreated        Time     `json:"timeCreated"`
	VcnID              string   `json:"vcnId"`
	VirtualRouterID    string   `json:"virtualRouterId"`
	VirtualRouterMac   string   `json:"virtualRouterMac"`
}

// ListSubnets contains a list of Subnet
type ListSubnets struct {
	ResourceContainer
	Subnets []Subnet
}

func (l *ListSubnets) GetList() interface{} {
	return &l.Subnets
}

// CreateSubnet will create a new subnet.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/Subnet/CreateSubnet
func (c *Client) CreateSubnet(
	availabilityDomain,
	cidrBlock,
	compartmentID,
	vcnID string,
	opts *CreateSubnetOptions,
) (sn *Subnet, e error) {

	required := struct {
		ocidRequirement
		AvailabilityDomain string `json:"availabilityDomain" url:"-"`
		CIDRBlock          string `json:"cidrBlock" url:"-"`
		VcnID              string `json:"vcnId" url:"-"`
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

	var response *requestResponse
	if response, e = c.coreApi.request(http.MethodPost, details); e != nil {
		return
	}

	sn = &Subnet{}
	e = response.unmarshal(sn)
	return
}

// GetSubnet will retrieve Subnet for subnetID.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/Subnet/GetSubnet
func (c *Client) GetSubnet(id string) (subnet *Subnet, e error) {
	details := &requestDetails{
		ids:  urlParts{id},
		name: resourceSubnets,
	}

	var response *requestResponse
	if response, e = c.coreApi.getRequest(details); e != nil {
		return
	}

	subnet = &Subnet{}
	e = response.unmarshal(subnet)
	return
}

// DeleteSubnet will delete a subnet with subnetID.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/Subnet/DeleteSubnet
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
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/Subnet/ListSubnets
func (c *Client) ListSubnets(compartmentID, vcnID string, opts *ListOptions) (subnets *ListSubnets, e error) {
	required := struct {
		listOCIDRequirement
		VcnID string `json:"-" url:"vcn"`
	}{
		VcnID: vcnID,
	}
	required.CompartmentID = compartmentID

	details := &requestDetails{
		name:     resourceSubnets,
		optional: opts,
		required: required,
	}

	var response *requestResponse
	if response, e = c.coreApi.getRequest(details); e != nil {
		return
	}

	subnets = &ListSubnets{}
	e = response.unmarshal(subnets)
	return
}
