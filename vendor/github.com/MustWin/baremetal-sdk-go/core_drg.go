package baremetal

import "net/http"

// Drg describes a dynamic routing gateway
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/Drg/
type Drg struct {
	ETaggedResource
	CompartmentID string `json:"compartmentId"`
	DisplayName   string `json:"displayName"`
	ID            string `json:"id"`
	State         string `json:"lifecycleState"`
	TimeCreated   Time   `json:"timeCreated"`
}

// ListDrgs contains a list of gateways
//
type ListDrgs struct {
	ResourceContainer
	Drgs []Drg
}

func (l *ListDrgs) GetList() interface{} {
	return &l.Drgs
}

// CreateDrg is used to create a gateway
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/Drg/CreateDrg
func (c *Client) CreateDrg(compartmentID string, opts *CreateOptions) (res *Drg, e error) {
	details := &requestDetails{
		name:     resourceDrgs,
		optional: opts,
		required: ocidRequirement{compartmentID},
	}

	var response *requestResponse
	if response, e = c.coreApi.request(http.MethodPost, details); e != nil {
		return
	}

	res = &Drg{}
	e = response.unmarshal(res)
	return
}

// GetDrg retrieves information about a gateway
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/Drg/GetDrg
func (c *Client) GetDrg(id string) (res *Drg, e error) {
	details := &requestDetails{
		name: resourceDrgs,
		ids:  urlParts{id},
	}

	var resp *requestResponse
	if resp, e = c.coreApi.getRequest(details); e != nil {
		return
	}

	res = &Drg{}
	e = resp.unmarshal(res)
	return
}

// DeleteDrg removes a gateway
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/Drg/DeleteDrg
func (c *Client) DeleteDrg(id string, opts *IfMatchOptions) (e error) {
	details := &requestDetails{
		name:     resourceDrgs,
		ids:      urlParts{id},
		optional: opts,
	}
	return c.coreApi.deleteRequest(details)
}

// ListDrgs returns a list of gateways for a compartment
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/Drg/ListDrgs
func (c *Client) ListDrgs(compartmentID string, opts *ListOptions) (res *ListDrgs, e error) {
	details := &requestDetails{
		name:     resourceDHCPOptions,
		required: listOCIDRequirement{compartmentID},
		optional: opts,
	}

	var resp *requestResponse
	if resp, e = c.coreApi.getRequest(details); e != nil {
		return
	}

	res = &ListDrgs{}
	e = resp.unmarshal(res)
	return
}
