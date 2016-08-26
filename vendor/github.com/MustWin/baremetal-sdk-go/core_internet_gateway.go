package baremetal

import (
	"net/http"
	"net/url"
)

// InternetGateway information on an internet gateway hosted in a
// virtual cloud network
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#InternetGateway
type InternetGateway struct {
	ETaggedResource
	CompartmentID string `json:"compartmentId"`
	DisplayName   string `json:"displayName,omitempty"`
	ID            string `json:"id"`
	IsEnabled     bool   `json:"isEnabled"`
	ModifiedTime  Time   `json:"modifiedTime"`
	State         string `json:"state"`
	TimeCreated   Time   `json:"timeCreated"`
}

// ListInternetGateways contains a set of internet gateways
type ListInternetGateways struct {
	ResourceContainer
	Gateways []InternetGateway
}

func (ig *ListInternetGateways) GetList() interface{} {
	return &ig.Gateways
}

// CreateInternetGateway creates an internet gateway. compartmentID is the compartment
// hosting the gateway, vcnID is the ID of the virtual cloud network, isEnabled
// determines if the gateway is enabled on creation. An optional display name may
// be provided in opts.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#createInternetGateway
func (c *Client) CreateInternetGateway(compartmentID, vcnID string, isEnabled bool, opts ...Options) (gw *InternetGateway, e error) {
	var displayName string
	if len(opts) > 0 {
		displayName = opts[0].DisplayName
	}

	createRequest := struct {
		CompartmentID string `json:"compartmentId"`
		DisplayName   string `json:"displayName,omitempty"`
		IsEnabled     bool   `json:"isEnabled"`
		VcnID         string `json:"vcnId"`
	}{
		CompartmentID: compartmentID,
		DisplayName:   displayName,
		IsEnabled:     isEnabled,
		VcnID:         vcnID,
	}

	requestOptions := &sdkRequestOptions{
		body:    createRequest,
		name:    resourceInternetGateways,
		options: opts,
	}

	var response *requestResponse
	if response, e = c.coreApi.request(http.MethodPost, requestOptions); e != nil {
		return
	}

	gw = &InternetGateway{}
	e = response.unmarshal(gw)
	return
}

// GetInternetGateway retrieves information for the Internet Gateway identified
// by id.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#getInternetGateway
func (c *Client) GetInternetGateway(id string) (gw *InternetGateway, e error) {
	reqOpts := &sdkRequestOptions{
		name: resourceInternetGateways,
		ids:  urlParts{id},
	}

	var response *requestResponse
	if response, e = c.coreApi.getRequest(reqOpts); e != nil {
		return
	}

	gw = &InternetGateway{}
	e = response.unmarshal(gw)
	return

}

// UpdateInternetGateway enables or disables internet gateway
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#UpdateInternetGatewayRequest
func (c *Client) UpdateInternetGateway(id string, isEnabled bool, opts ...Options) (gw *InternetGateway, e error) {

	body := struct {
		IsEnabled bool `json:"isEnabled"`
	}{
		IsEnabled: isEnabled,
	}

	reqOpts := &sdkRequestOptions{
		name:    resourceInternetGateways,
		ids:     urlParts{id},
		body:    body,
		options: opts,
	}

	var response *requestResponse
	if response, e = c.coreApi.request(http.MethodPut, reqOpts); e != nil {
		return
	}

	gw = &InternetGateway{}
	e = response.unmarshal(gw)
	return

}

// DeleteInternetGateway removes internet gateway
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#deleteInternetGateway
func (c *Client) DeleteInternetGateway(id string, opts ...Options) (e error) {
	request := &sdkRequestOptions{
		name:    resourceInternetGateways,
		ids:     urlParts{id},
		options: opts,
	}
	return c.coreApi.deleteRequest(request)
}

// ListInternetGateways is used to fetch a list of internet gateways.
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#listInternetGateways
func (c *Client) ListInternetGateways(compartmentID, vcnID string, opts ...Options) (list *ListInternetGateways, e error) {
	query := url.Values{}
	query.Set(queryVcnID, vcnID)

	request := &sdkRequestOptions{
		ocid:    compartmentID,
		options: opts,
		query:   query,
		name:    resourceInternetGateways,
	}

	var response *requestResponse
	if response, e = c.coreApi.getRequest(request); e != nil {
		return
	}

	list = &ListInternetGateways{}
	e = response.unmarshal(list)
	return
}
