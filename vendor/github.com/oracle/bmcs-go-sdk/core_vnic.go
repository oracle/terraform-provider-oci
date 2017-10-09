// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

import "net/http"

// Vnic describes a virtual network interface.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Vnic/
type Vnic struct {
	OPCRequestIDUnmarshaller
	ETagUnmarshaller
	AvailabilityDomain  string `json:"availabilityDomain"`
	CompartmentID       string `json:"compartmentId"`
	DisplayName         string `json:"displayName"`
	HostnameLabel       string `json:"hostnameLabel"`
	ID                  string `json:"id"`
	IsPrimary           bool   `json:"isPrimary"`
	MacAddress          string `json:"macAddress"`
	State               string `json:"lifecycleState"`
	PrivateIPAddress    string `json:"privateIp"`
	PublicIPAddress     string `json:"publicIp"`
	SkipSourceDestCheck bool   `json:"skipSourceDestCheck"`
	SubnetID            string `json:"subnetId"`
	TimeCreated         Time   `json:"timeCreated"`
}

// GetVnic retrieves information about a virtual network interface identified
// by vnicID. ListVnicAttachments can be used to retrieve Vnic IDs.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Vnic/GetVnic
func (c *Client) GetVnic(id string) (vnic *Vnic, e error) {
	details := &requestDetails{
		name: resourceVnics,
		ids:  urlParts{id},
	}

	var resp *response
	if resp, e = c.coreApi.getRequest(details); e != nil {
		return
	}

	vnic = &Vnic{}
	e = resp.unmarshal(vnic)
	return
}

// UpdateVnic updates the specified VNIC.
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Vnic/UpdateVnic
func (c *Client) UpdateVnic(id string, opts *UpdateVnicOptions) (res *Vnic, e error) {
	details := &requestDetails{
		ids:      urlParts{id},
		name:     resourceVnics,
		optional: opts,
	}

	var resp *response
	if resp, e = c.coreApi.request(http.MethodPut, details); e != nil {
		return
	}

	res = &Vnic{}
	e = resp.unmarshal(res)
	return
}
