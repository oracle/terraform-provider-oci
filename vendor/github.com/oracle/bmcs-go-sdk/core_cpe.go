// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

import "net/http"

// Cpe describes customer premise equipment
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Cpe/
type Cpe struct {
	OPCRequestIDUnmarshaller
	ETagUnmarshaller
	ID            string `json:"id"`
	CompartmentID string `json:"compartmentId"`
	DisplayName   string `json:"displayName"`
	IPAddress     string `json:"ipAddress"`
	TimeCreated   Time   `json:"timeCreated"`
}

// ListCpes contains a list of customer premise equipment
//
type ListCpes struct {
	OPCRequestIDUnmarshaller
	NextPageUnmarshaller
	Cpes []Cpe
}

func (l *ListCpes) GetList() interface{} {
	return &l.Cpes
}

// ListCpes returns a list of customer premise equipment for a particular compartment
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Cpe/ListCpes
func (c *Client) ListCpes(compartmentID string, opts *ListOptions) (cpes *ListCpes, e error) {
	details := &requestDetails{
		name:     resourceCustomerPremiseEquipment,
		required: listOCIDRequirement{CompartmentID: compartmentID},
		optional: opts,
	}

	var resp *response
	if resp, e = c.coreApi.getRequest(details); e != nil {
		return
	}

	cpes = &ListCpes{}
	e = resp.unmarshal(cpes)
	return
}

// CreateCpe is used to define customer premise equipment such as routers
// in the Bare Metal cloud
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Cpe/CreateCpe
func (c *Client) CreateCpe(compartmentID, ipAddress string, opts *CreateOptions) (cpe *Cpe, e error) {
	required := struct {
		ocidRequirement
		IPAddress string `header:"-" json:"ipAddress" url:"-"`
	}{
		IPAddress: ipAddress,
	}
	required.CompartmentID = compartmentID

	details := &requestDetails{
		name:     resourceCustomerPremiseEquipment,
		optional: opts,
		required: required,
	}

	var resp *response
	if resp, e = c.coreApi.postRequest(details); e != nil {
		return
	}

	cpe = &Cpe{}
	e = resp.unmarshal(cpe)
	return
}

// GetCpe retrieves information on a customer premise equipment resource.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Cpe/GetCpe
func (c *Client) GetCpe(id string) (cpe *Cpe, e error) {
	details := &requestDetails{
		name: resourceCustomerPremiseEquipment,
		ids:  urlParts{id},
	}
	var resp *response
	if resp, e = c.coreApi.getRequest(details); e != nil {
		return
	}

	cpe = &Cpe{}
	e = resp.unmarshal(cpe)
	return
}

// UpdateCpe updates the specified CPE's display name.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Cpe/UpdateCpe
func (c *Client) UpdateCpe(id string, opts *IfMatchDisplayNameOptions) (cpe *Cpe, e error) {
	details := &requestDetails{
		name:     resourceCustomerPremiseEquipment,
		ids:      urlParts{id},
		optional: opts,
	}
	var resp *response
	if resp, e = c.coreApi.request(http.MethodPut, details); e != nil {
		return
	}

	cpe = &Cpe{}
	e = resp.unmarshal(cpe)

	return
}

// DeleteCpe removes customer premise equipment resource
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Cpe/DeleteCpe
func (c *Client) DeleteCpe(id string, opts *IfMatchOptions) (e error) {
	details := &requestDetails{
		name:     resourceCustomerPremiseEquipment,
		ids:      urlParts{id},
		optional: opts,
	}
	return c.coreApi.deleteRequest(details)
}
