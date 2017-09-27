// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

import "net/http"

// Volume describes cloud block storage
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Volume/
type Volume struct {
	OPCRequestIDUnmarshaller
	ETagUnmarshaller
	AvailabilityDomain string `json:"availabilityDomain"`
	CompartmentID      string `json:"compartmentId"`
	DisplayName        string `json:"displayName"`
	ID                 string `json:"id"`
	SizeInMBs          int    `json:"sizeInMBs"`
	SizeInGBs          int    `json:"sizeInGBs"`
	State              string `json:"lifecycleState"`
	TimeCreated        Time   `json:"timeCreated"`
}

// ListVolumes contains a list of block volumes
//
type ListVolumes struct {
	OPCRequestIDUnmarshaller
	NextPageUnmarshaller
	Volumes []Volume
}

func (l *ListVolumes) GetList() interface{} {
	return &l.Volumes
}

// CreateVolume is used to create a cloud block storage device
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Volume/CreateVolume
func (c *Client) CreateVolume(availabilityDomain, compartmentID string, opts *CreateVolumeOptions) (res *Volume, e error) {
	required := struct {
		ocidRequirement
		AvailabilityDomain string `header:"-" json:"availabilityDomain" url:"-"`
	}{
		AvailabilityDomain: availabilityDomain,
	}
	required.CompartmentID = compartmentID

	details := &requestDetails{
		name:     resourceVolumes,
		optional: opts,
		required: required,
	}

	var resp *response
	if resp, e = c.coreApi.postRequest(details); e != nil {
		return
	}

	res = &Volume{}
	e = resp.unmarshal(res)
	return
}

// GetVolume retrieves information about a block volume
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Volume/GetVolume
func (c *Client) GetVolume(id string) (res *Volume, e error) {
	details := &requestDetails{
		name: resourceVolumes,
		ids:  urlParts{id},
	}

	var resp *response
	if resp, e = c.coreApi.getRequest(details); e != nil {
		return
	}

	res = &Volume{}
	e = resp.unmarshal(res)
	return
}

// UpdateVolume updates a volume's display name
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Volume/UpdateVolume
func (c *Client) UpdateVolume(id string, opts *UpdateOptions) (res *Volume, e error) {
	details := &requestDetails{
		ids:      urlParts{id},
		name:     resourceVolumes,
		optional: opts,
	}

	var resp *response
	if resp, e = c.coreApi.request(http.MethodPut, details); e != nil {
		return
	}

	res = &Volume{}
	e = resp.unmarshal(res)
	return
}

// DeleteVolume removes a cloud block storage volume
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Volume/DeleteVolume
func (c *Client) DeleteVolume(id string, opts *IfMatchOptions) (e error) {
	details := &requestDetails{
		ids:      urlParts{id},
		name:     resourceVolumes,
		optional: opts,
	}

	return c.coreApi.deleteRequest(details)
}

// ListVolumes returns a list of volumes for a particular compartment
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Volume/ListVolumes
func (c *Client) ListVolumes(compartmentID string, opts *ListVolumesOptions) (res *ListVolumes, e error) {
	details := &requestDetails{
		name:     resourceVolumes,
		optional: opts,
		required: ocidRequirement{compartmentID},
	}

	var resp *response
	if resp, e = c.coreApi.getRequest(details); e != nil {
		return
	}

	res = &ListVolumes{}
	e = resp.unmarshal(res)
	return
}
