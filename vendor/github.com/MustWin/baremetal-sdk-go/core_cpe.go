package baremetal

import "net/http"

// Cpe describes customer premise equipment
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#Cpe
type Cpe struct {
	ETaggedResource
	ID            string `json:"id"`
	CompartmentID string `json:"compartmentId"`
	DisplayName   string `json:"displayName"`
	IPAddress     string `json:"ipAddress"`
	TimeCreated   Time   `json:"timeCreated"`
}

// CpeList contains a list of customer premise equipment
//
type CpeList struct {
	ResourceContainer
	Cpes []Cpe
}

func (l *CpeList) GetList() interface{} {
	return &l.Cpes
}

type CreateCpeRequest struct {
	CompartmentID string `json:"compartmentId"`
	DisplayName   string `json:"displayName,omitempty"`
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
	e = resp.unmarshal(cpes)
	return
}

// CreateCpe is used to define customer premise equipment such as routers
// in the Bare Metal cloud
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#createCpe
func (c *Client) CreateCpe(compartmentID, IPAddress string, opts ...Options) (cpe *Cpe, e error) {
	createRequest := CreateCpeRequest{
		CompartmentID: compartmentID,
		IPAddress:     IPAddress,
	}

	if len(opts) > 0 {
		createRequest.DisplayName = opts[0].DisplayName
	}

	reqOpts := &sdkRequestOptions{
		body:    createRequest,
		name:    resourceCustomerPremiseEquipment,
		options: opts,
	}

	var resp *requestResponse
	if resp, e = c.coreApi.request(http.MethodPost, reqOpts); e != nil {
		return
	}

	cpe = &Cpe{}
	e = resp.unmarshal(cpe)
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
	e = resp.unmarshal(cpe)
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
