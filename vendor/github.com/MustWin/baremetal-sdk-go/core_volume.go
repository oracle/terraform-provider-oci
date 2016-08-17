package baremetal

import (
	"encoding/json"
	"net/http"
)

// Volume describes cloud block storage
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#Volume
type Volume struct {
	AvailabilityDomain string `json:"availabilityDomain"`
	CompartmentID      string `json:"compartmentId"`
	DisplayName        string `json:"displayName"`
	ID                 string `json:"id"`
	SizeInMBs          string `json:"sizeInMBs"`
	State              string `json:"state"`
	TimeCreated        Time   `json:"timeCreated"`
	ETag               string `json:"etag,omitempty"`
	OPCRequestID       string `json:"opc-request-id,omitempty"`
}

// VolumeList contains a list of block volumes
//
type VolumeList struct {
	OPCNextPage  string
	OPCRequestID string
	Volumes      []Volume
}

// CreateVolumeRequest describes the body of a volume creation request
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#CreateVolumeRequest
type CreateVolumeRequest struct {
	AvailabilityDomain string  `json:"availabilityDomain"`
	CompartmentID      string  `json:"compartmentId"`
	DisplayName        *string `json:"displayName"`
}

// UpdateVolumeRequest describes the body of a volume name update request
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#UpdateVolumeRequest
type UpdateVolumeRequest struct {
	DisplayName *string `json:"displayName"`
}

// CreateVolume is used to create a cloud block storage device
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#createVolume
func (c *Client) CreateVolume(availabilityDomain, compartmentID string, opts ...Options) (vol *Volume, e error) {
	createRequest := CreateVolumeRequest{
		AvailabilityDomain: availabilityDomain,
		CompartmentID:      compartmentID,
	}

	if len(opts) > 0 {
		createRequest.DisplayName = &opts[0].DisplayName
	}

	reqOpts := &sdkRequestOptions{
		body:    createRequest,
		name:    resourceVolumes,
		options: opts,
	}

	var response *requestResponse
	if response, e = c.coreApi.request(http.MethodPost, reqOpts); e != nil {
		return
	}

	vol = &Volume{}

	if e = json.Unmarshal(response.body, vol); e != nil {
		return
	}

	vol.ETag = response.header.Get(headerETag)
	vol.OPCRequestID = response.header.Get(headerOPCRequestID)

	return
}

// GetVolume retrieves information about a block volume
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#getVolume
func (c *Client) GetVolume(id string, opts ...Options) (vol *Volume, e error) {
	reqOpts := &sdkRequestOptions{
		name:    resourceVolumes,
		options: opts,
		ids:     urlParts{id},
	}
	var resp *requestResponse
	if resp, e = c.coreApi.getRequest(reqOpts); e != nil {
		return
	}

	vol = &Volume{}

	if e = json.Unmarshal(resp.body, vol); e != nil {
		return
	}

	vol.ETag = resp.header.Get(headerETag)
	vol.OPCRequestID = resp.header.Get(headerOPCRequestID)

	return
}

// UpdateVolume updates a volume's display name
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#updateVolume
func (c *Client) UpdateVolume(id string, opts ...Options) (vol *Volume, e error) {
	body := UpdateVolumeRequest{}

	if len(opts) > 0 {
		body.DisplayName = &opts[0].DisplayName
	}

	reqOpts := &sdkRequestOptions{
		body:    body,
		name:    resourceVolumes,
		options: opts,
		ids:     urlParts{id},
	}

	var response *requestResponse
	if response, e = c.coreApi.request(http.MethodPut, reqOpts); e != nil {
		return
	}

	vol = &Volume{}
	e = json.Unmarshal(response.body, vol)

	if respHeader := response.header; respHeader != nil {
		vol.ETag = respHeader.Get(headerETag)
	}

	return
}

// DeleteVolume removes a cloud block storage volume
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#deleteVolume
func (c *Client) DeleteVolume(id string, opts ...Options) (e error) {
	reqOpts := &sdkRequestOptions{
		name:    resourceVolumes,
		options: opts,
		ids:     urlParts{id},
	}
	return c.coreApi.deleteRequest(reqOpts)
}

// ListVolumes returns a list of volumes for a particular compartment
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#listVolumes
func (c *Client) ListVolumes(compartmentID string, opts ...Options) (vols *VolumeList, e error) {
	reqOpts := &sdkRequestOptions{
		name:    resourceVolumes,
		ocid:    compartmentID,
		options: opts,
	}

	var resp *requestResponse
	if resp, e = c.coreApi.getRequest(reqOpts); e != nil {
		return
	}

	vols = &VolumeList{}
	if e = json.Unmarshal(resp.body, &vols.Volumes); e != nil {
		return
	}

	vols.OPCNextPage = resp.header.Get(headerOPCNextPage)
	vols.OPCRequestID = resp.header.Get(headerOPCRequestID)

	return
}
