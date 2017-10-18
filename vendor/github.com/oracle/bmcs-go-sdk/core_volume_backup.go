// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

import "net/http"

// VolumeBackup describe a point-in-time copy of a volume
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/VolumeBackup/
type VolumeBackup struct {
	OPCRequestIDUnmarshaller
	ETagUnmarshaller
	CompartmentID       string `json:"compartmentId"`
	DisplayName         string `json:"displayName"`
	ID                  string `json:"id"`
	SizeInMBs           uint64 `json:"sizeInMBs"`
	SizeInGBs           uint64 `json:"sizeInGBs"`
	State               string `json:"lifecycleState"`
	TimeCreated         Time   `json:"timeCreated"`
	TimeRequestReceived Time   `json:"timeRequestReceived"`
	UniqueSizeInMBs     uint64 `json:"uniqueSizeInMBs"`
	UniqueSizeInGBs     uint64 `json:"uniqueSizeInGBs"`
	VolumeID            string `json:"volumeId"`
}

// ListVolumeBackups contains a list of volume backups
//
type ListVolumeBackups struct {
	OPCRequestIDUnmarshaller
	NextPageUnmarshaller
	VolumeBackups []VolumeBackup
}

func (l *ListVolumeBackups) GetList() interface{} {
	return &l.VolumeBackups
}

// CreateVolumeBackup Creates a new backup of the specified volume
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/VolumeBackup/CreateVolumeBackup
func (c *Client) CreateVolumeBackup(volumeID string, opts *CreateOptions) (vol *VolumeBackup, e error) {
	required := struct {
		VolumeID string `header:"-" json:"volumeId" url:"-"`
	}{
		VolumeID: volumeID,
	}

	details := &requestDetails{
		name:     resourceVolumeBackups,
		optional: opts,
		required: required,
	}

	var resp *response
	if resp, e = c.coreApi.postRequest(details); e != nil {
		return
	}

	vol = &VolumeBackup{}
	e = resp.unmarshal(vol)
	return
}

// GetVolumeBackup gets information for the specified volumeBackup
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/VolumeBackup/GetVolumeBackup
func (c *Client) GetVolumeBackup(id string) (vol *VolumeBackup, e error) {
	details := &requestDetails{
		ids:  urlParts{id},
		name: resourceVolumeBackups,
	}

	var resp *response
	if resp, e = c.coreApi.getRequest(details); e != nil {
		return
	}

	vol = &VolumeBackup{}
	e = resp.unmarshal(vol)
	return
}

// UpdateVolumeBackup updates a volume's display name
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/VolumeBackup/UpdateVolumeBackup
func (c *Client) UpdateVolumeBackup(id string, opts *IfMatchDisplayNameOptions) (vol *VolumeBackup, e error) {
	details := &requestDetails{
		ids:      urlParts{id},
		name:     resourceVolumeBackups,
		optional: opts,
	}

	var resp *response
	if resp, e = c.coreApi.request(http.MethodPut, details); e != nil {
		return
	}

	vol = &VolumeBackup{}
	e = resp.unmarshal(vol)
	return
}

// DeleteVolumeBackup deletes a volumeBackup
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/VolumeBackup/DeleteVolumeBackup
func (c *Client) DeleteVolumeBackup(id string, opts *IfMatchOptions) (e error) {
	details := &requestDetails{
		ids:      urlParts{id},
		name:     resourceVolumeBackups,
		optional: opts,
	}

	return c.coreApi.deleteRequest(details)
}

// ListVolumeBackups returns a list of volumes for a particular compartment
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/VolumeBackup/ListVolumeBackups
func (c *Client) ListVolumeBackups(compartmentID string, opts *ListBackupsOptions) (vols *ListVolumeBackups, e error) {
	details := &requestDetails{
		name:     resourceVolumeBackups,
		optional: opts,
		required: ocidRequirement{compartmentID},
	}

	var resp *response
	if resp, e = c.coreApi.getRequest(details); e != nil {
		return
	}

	vols = &ListVolumeBackups{}
	e = resp.unmarshal(vols)
	return
}
