// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

import "net/http"

// InternetGateway information on an internet gateway hosted in a
// virtual cloud network
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/InternetGateway/
type InternetGateway struct {
	OPCRequestIDUnmarshaller
	ETagUnmarshaller
	CompartmentID string `json:"compartmentId"`
	DisplayName   string `json:"displayName,omitempty"`
	ID            string `json:"id"`
	IsEnabled     bool   `json:"isEnabled"`
	ModifiedTime  Time   `json:"modifiedTime"`
	State         string `json:"lifecycleState"`
	TimeCreated   Time   `json:"timeCreated"`
}

// ListInternetGateways contains a set of internet gateways
type ListInternetGateways struct {
	OPCRequestIDUnmarshaller
	NextPageUnmarshaller
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
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/InternetGateway/CreateInternetGateway
func (c *Client) CreateInternetGateway(compartmentID, vcnID string, isEnabled bool, opts *CreateOptions) (gw *InternetGateway, e error) {
	required := struct {
		ocidRequirement
		IsEnabled bool   `header:"-" json:"isEnabled" url:"-"`
		VcnID     string `header:"-" json:"vcnId" url:"-"`
	}{
		IsEnabled: isEnabled,
		VcnID:     vcnID,
	}
	required.CompartmentID = compartmentID

	details := &requestDetails{
		name:     resourceInternetGateways,
		optional: opts,
		required: required,
	}

	var resp *response
	if resp, e = c.coreApi.postRequest(details); e != nil {
		return
	}

	gw = &InternetGateway{}
	e = resp.unmarshal(gw)
	return
}

// GetInternetGateway retrieves information for the Internet Gateway identified
// by id.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/InternetGateway/GetInternetGateway
func (c *Client) GetInternetGateway(id string) (gw *InternetGateway, e error) {
	details := &requestDetails{
		name: resourceInternetGateways,
		ids:  urlParts{id},
	}

	var resp *response
	if resp, e = c.coreApi.getRequest(details); e != nil {
		return
	}

	gw = &InternetGateway{}
	e = resp.unmarshal(gw)
	return

}

// UpdateInternetGateway enables or disables internet gateway
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/InternetGateway/UpdateInternetGateway
func (c *Client) UpdateInternetGateway(id string, opts *UpdateGatewayOptions) (gw *InternetGateway, e error) {
	details := &requestDetails{
		ids:      urlParts{id},
		name:     resourceInternetGateways,
		optional: opts,
	}

	var resp *response
	if resp, e = c.coreApi.request(http.MethodPut, details); e != nil {
		return
	}

	gw = &InternetGateway{}
	e = resp.unmarshal(gw)
	return
}

// DeleteInternetGateway removes internet gateway
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/InternetGateway/DeleteInternetGateway
func (c *Client) DeleteInternetGateway(id string, opts *IfMatchOptions) (e error) {
	details := &requestDetails{
		name:     resourceInternetGateways,
		ids:      urlParts{id},
		optional: opts,
	}
	return c.coreApi.deleteRequest(details)
}

// ListInternetGateways is used to fetch a list of internet gateways.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/InternetGateway/ListInternetGateways
func (c *Client) ListInternetGateways(compartmentID, vcnID string, opts *ListOptions) (list *ListInternetGateways, e error) {
	required := struct {
		listOCIDRequirement
		VcnID string `header:"-" json:"-" url:"vcnId"`
	}{
		VcnID: vcnID,
	}
	required.CompartmentID = compartmentID

	details := &requestDetails{
		name:     resourceInternetGateways,
		optional: opts,
		required: required,
	}

	var resp *response
	if resp, e = c.coreApi.getRequest(details); e != nil {
		return
	}

	list = &ListInternetGateways{}
	e = resp.unmarshal(list)
	return
}
