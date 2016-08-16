package baremetal

import (
	"encoding/json"
	"net/http"
)

// Cpe describes customer premise equipment
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#Cpe
type Cpe struct {
	ID            string `json:"id"`
	CompartmentID string `json:"compartmentId"`
	DisplayName   string `json:"displayName"`
	IPAddress     string `json:"ipAddress"`
	TimeCreated   Time   `json:"timeCreated"`
	ETag          string `json:"etag,omitempty"`
	OPCRequestID  string `json:"opc-request-id,omitempty"`
}

// CpeList contains a list of customer premise equipment
//
type CpeList struct {
	OPCNextPage  string
	OPCRequestID string
	Cpes         []Cpe
}

type CreateCpeRequest struct {
	CompartmentID string `json:"compartmentId"`
	DisplayName   string `json:"displayName"`
	IPAddress     string `json:"ipAddress"`
}

// ListCpes returns a list of customer premise equipment for a particular compartment
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#listCpes
func (c *Client) ListCpes(compartmentID string, opts ...Options) (cpes *CpeList, e error) {
	reqOpts := &sdkRequestOptions{
		name:    resourceCustomerPremiseEquipment,
		ocid:    compartmentID,
		options: opts,
	}

	var resp *requestResponse
	if resp, e = c.coreApi.getRequest(reqOpts); e != nil {
		return
	}

	cpes = &CpeList{}
	if e = json.Unmarshal(resp.body, &cpes.Cpes); e != nil {
		return
	}

	cpes.OPCNextPage = resp.header.Get(headerOPCNextPage)
	cpes.OPCRequestID = resp.header.Get(headerOPCRequestID)

	return

}

// CreateCpe is used to define customer premise equipment such as routers
// in the Bare Metal cloud
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#createCpe
func (c *Client) CreateCpe(compartmentID, displayName, IPAddress string, opts ...Options) (cpe *Cpe, e error) {
	createRequest := CreateCpeRequest{
		CompartmentID: compartmentID,
		DisplayName:   displayName,
		IPAddress:     IPAddress,
	}

	reqOpts := &sdkRequestOptions{
		body:    createRequest,
		name:    resourceCustomerPremiseEquipment,
		options: opts,
	}

	var response *requestResponse
	if response, e = c.coreApi.request(http.MethodPost, reqOpts); e != nil {
		return
	}

	cpe = &Cpe{}

	if e = json.Unmarshal(response.body, cpe); e != nil {
		return
	}

	cpe.ETag = response.header.Get(headerETag)
	cpe.OPCRequestID = response.header.Get(headerOPCRequestID)

	return
}

// GetCpe retrieves information on a customer premise equipment resource.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#getCpe
func (c *Client) GetCpe(id string, opts ...Options) (cpe *Cpe, e error) {
	reqOpts := &sdkRequestOptions{
		name:    resourceCustomerPremiseEquipment,
		options: opts,
		ids:     urlParts{id},
	}
	var resp *requestResponse
	if resp, e = c.coreApi.getRequest(reqOpts); e != nil {
		return
	}

	cpe = &Cpe{}

	if e = json.Unmarshal(resp.body, cpe); e != nil {
		return
	}

	cpe.ETag = resp.header.Get(headerETag)
	cpe.OPCRequestID = resp.header.Get(headerOPCRequestID)

	return

}

// DeleteCpe removes customer premise equipment resource
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#deleteCpe
func (c *Client) DeleteCpe(id string, opts ...Options) (e error) {
	reqOpts := &sdkRequestOptions{
		name:    resourceCustomerPremiseEquipment,
		options: opts,
		ids:     urlParts{id},
	}
	return c.coreApi.deleteRequest(reqOpts)
}
