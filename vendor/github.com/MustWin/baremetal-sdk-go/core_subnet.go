package baremetal

import (
	"net/http"
	"net/url"
)

// Subnet represents a network subnet
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#Subnet
type Subnet struct {
	ETaggedResource
	AvailabilityDomain string   `json:"availabilityDomain"`
	CIDRBlock          string   `json:"cidrBlock"`
	CompartmentID      string   `json:"compartmentId"`
	DisplayName        string   `json:"displayName"`
	ID                 string   `json:"id"`
	RouteTableID       string   `json:"routeTableId"`
	SecurityListIDs    []string `json:"securityListIds"`
	State              string   `json:"state"`
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
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#createSubnet
func (c *Client) CreateSubnet(
	availabilityDomain,
	cidrBlock,
	compartmentID,
	routeTableID,
	vcnID string,
	securityListIDs []string,
	opts ...Options,
) (sn *Subnet, e error) {

	var displayName string
	if len(opts) > 0 {
		displayName = opts[0].DisplayName
	}

	requestBody := struct {
		AvailabilityDomain string   `json:"availabilityDomain"`
		CIDRBlock          string   `json:"cidrBlock"`
		CompartmentID      string   `json:"compartmentId"`
		DisplayName        string   `json:"displayName,omitempty"`
		RouteTableID       string   `json:"routeTableId"`
		SecurityListIDs    []string `json:"securityListIds"`
		VcnID              string   `json:"vcnId"`
	}{
		AvailabilityDomain: availabilityDomain,
		CIDRBlock:          cidrBlock,
		CompartmentID:      compartmentID,
		DisplayName:        displayName,
		RouteTableID:       routeTableID,
		SecurityListIDs:    securityListIDs,
		VcnID:              vcnID,
	}

	req := &sdkRequestOptions{
		name:    resourceSubnets,
		body:    requestBody,
		options: opts,
	}

	var response *requestResponse
	if response, e = c.coreApi.request(http.MethodPost, req); e != nil {
		return
	}

	sn = &Subnet{}
	e = response.unmarshal(sn)
	return
}

// ListSubnets returns a list of subnets in compartment for a virtual cloud network.
// The size of results may be limited by assigning values to the Limit field of
// Options.  Results may be paged by assigning the NewPage from the last
// response to the Page member of Options.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#listSubnets
func (c *Client) ListSubnets(compartmentID, vcnID string, opts ...Options) (subnets *ListSubnets, e error) {
	query := url.Values{}
	query.Set(queryVcnID, vcnID)

	req := &sdkRequestOptions{
		name:    resourceSubnets,
		ocid:    compartmentID,
		query:   query,
		options: opts,
	}

	var response *requestResponse
	if response, e = c.coreApi.getRequest(req); e != nil {
		return
	}

	subnets = &ListSubnets{}
	e = response.unmarshal(subnets)
	return
}

// GetSubnet will retrieve Subnet for subnetID.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#getSubnet
func (c *Client) GetSubnet(subnetID string) (subnet *Subnet, e error) {
	req := &sdkRequestOptions{
		name: resourceSubnets,
		ids:  urlParts{subnetID},
	}

	var response *requestResponse
	if response, e = c.coreApi.getRequest(req); e != nil {
		return
	}

	subnet = &Subnet{}
	e = response.unmarshal(subnet)
	return
}

// DeleteSubnet will delete a subnet with subnetID.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#deleteSubnet
func (c *Client) DeleteSubnet(subnetID string, opts ...Options) error {
	req := &sdkRequestOptions{
		name:    resourceSubnets,
		ids:     urlParts{subnetID},
		options: opts,
	}

	return c.coreApi.deleteRequest(req)
}
