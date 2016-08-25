package baremetal

import "net/http"

// Drg describes a dynamic routing gateway
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#Drg
type Drg struct {
	ETaggedResource
	CompartmentID string `json:"compartmentId"`
	DisplayName   string `json:"displayName"`
	ID            string `json:"id"`
	State         string `json:"state"`
	TimeCreated   Time   `json:"timeCreated"`
}

// DrgList contains a list of gateways
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
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#createDrg
func (c *Client) CreateDrg(compartmentID string, opts ...Options) (res *Drg, e error) {
	body := struct {
		CompartmentID string `json:"compartmentId"`
		DisplayName   string `json:"displayName,omitempty"`
	}{
		CompartmentID: compartmentID,
	}

	if len(opts) > 0 {
		body.DisplayName = opts[0].DisplayName
	}

	reqOpts := &sdkRequestOptions{
		body:    body,
		name:    resourceDrgs,
		options: opts,
	}

	var response *requestResponse
	if response, e = c.coreApi.request(http.MethodPost, reqOpts); e != nil {
		return
	}

	res = &Drg{}
	e = response.unmarshal(res)
	return
}

// GetDrg retrieves information about a gateway
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#getDrg
func (c *Client) GetDrg(id string, opts ...Options) (res *Drg, e error) {
	reqOpts := &sdkRequestOptions{
		name:    resourceDrgs,
		options: opts,
		ids:     urlParts{id},
	}
	var resp *requestResponse
	if resp, e = c.coreApi.getRequest(reqOpts); e != nil {
		return
	}

	res = &Drg{}
	e = resp.unmarshal(res)
	return
}

// DeleteDrg removes a gateway
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#deleteDrg
func (c *Client) DeleteDrg(id string, opts ...Options) (e error) {
	reqOpts := &sdkRequestOptions{
		name:    resourceDrgs,
		options: opts,
		ids:     urlParts{id},
	}
	return c.coreApi.deleteRequest(reqOpts)
}

// ListDrgs returns a list of gateways for a compartment
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#listDrgs
func (c *Client) ListDrgs(compartmentID string, opts ...Options) (res *ListDrgs, e error) {
	reqOpts := &sdkRequestOptions{
		name:    resourceDrgs,
		ocid:    compartmentID,
		options: opts,
	}

	var resp *requestResponse
	if resp, e = c.coreApi.getRequest(reqOpts); e != nil {
		return
	}

	res = &ListDrgs{}
	e = resp.unmarshal(res)
	return
}
