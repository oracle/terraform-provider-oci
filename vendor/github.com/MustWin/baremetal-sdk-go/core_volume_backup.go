package baremetal

import "net/http"

// VolumeBackup describe a point-in-time copy of a volume
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/VolumeBackup/
type VolumeBackup struct {
	ETaggedResource
	CompartmentID       string `json:"compartmentId"`
	DisplayName         string `json:"displayName"`
	ID                  string `json:"id"`
	SizeInMBs           uint64 `json:"sizeInMBs"`
	State               string `json:"lifecycleState"`
	TimeCreated         Time   `json:"timeCreated"`
	TimeRequestReceived Time   `json:"timeRequestReceived"`
	UniqueSizeInMBs     uint64 `json:"uniqueSizeInMBs"`
	VolumeID            string `json:"volumeId"`
}

// ListVolumeBackups contains a list of volume backups
//
type ListVolumeBackups struct {
	ResourceContainer
	VolumeBackups []VolumeBackup
}

func (l *ListVolumeBackups) GetList() interface{} {
	return &l.VolumeBackups
}

// CreateVolumeBackup Creates a new backup of the specified volume
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/VolumeBackup/CreateVolumeBackup
func (c *Client) CreateVolumeBackup(volumeID string, opts ...Options) (vol *VolumeBackup, e error) {
	body := struct {
		DisplayName string `json:"displayName,omitempty"`
		VolumeID    string `json:"volumeId"`
	}{VolumeID: volumeID}
	if len(opts) > 0 {
		body.DisplayName = opts[0].DisplayName
	}

	reqOpts := &sdkRequestOptions{
		body:    body,
		name:    resourceVolumeBackups,
		options: opts,
	}

	var response *requestResponse
	if response, e = c.coreApi.request(http.MethodPost, reqOpts); e != nil {
		return
	}

	vol = &VolumeBackup{}
	e = response.unmarshal(vol)
	return
}

// GetVolumeBackup gets information for the specified volumeBackup
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/VolumeBackup/GetVolumeBackup
func (c *Client) GetVolumeBackup(id string, opts ...Options) (vol *VolumeBackup, e error) {
	reqOpts := &sdkRequestOptions{
		name:    resourceVolumeBackups,
		options: opts,
		ids:     urlParts{id},
	}
	var resp *requestResponse
	if resp, e = c.coreApi.getRequest(reqOpts); e != nil {
		return
	}

	vol = &VolumeBackup{}
	e = resp.unmarshal(vol)
	return
}

// UpdateVolumeBackup updates a volume's display name
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/VolumeBackup/UpdateVolumeBackup
func (c *Client) UpdateVolumeBackup(id string, opts ...Options) (vol *VolumeBackup, e error) {
	body := struct {
		DisplayName string `json:"displayName,omitempty"`
	}{}
	if len(opts) > 0 {
		body.DisplayName = opts[0].DisplayName
	}

	reqOpts := &sdkRequestOptions{
		body:    body,
		name:    resourceVolumeBackups,
		options: opts,
		ids:     urlParts{id},
	}

	var response *requestResponse
	if response, e = c.coreApi.request(http.MethodPut, reqOpts); e != nil {
		return
	}

	vol = &VolumeBackup{}
	e = response.unmarshal(vol)
	return
}

// DeleteVolumeBackup deletes a volumeBackup
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/VolumeBackup/DeleteVolumeBackup
func (c *Client) DeleteVolumeBackup(id string, opts ...Options) (e error) {
	reqOpts := &sdkRequestOptions{
		name:    resourceVolumeBackups,
		options: opts,
		ids:     urlParts{id},
	}
	return c.coreApi.deleteRequest(reqOpts)
}

// ListVolumeBackups returns a list of volumes for a particular compartment
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/VolumeBackup/ListVolumeBackups
func (c *Client) ListVolumeBackups(compartmentID string, opts ...Options) (vols *ListVolumeBackups, e error) {
	reqOpts := &sdkRequestOptions{
		name:    resourceVolumeBackups,
		ocid:    compartmentID,
		options: opts,
	}

	var resp *requestResponse
	if resp, e = c.coreApi.getRequest(reqOpts); e != nil {
		return
	}

	vols = &ListVolumeBackups{}
	e = resp.unmarshal(vols)
	return
}
