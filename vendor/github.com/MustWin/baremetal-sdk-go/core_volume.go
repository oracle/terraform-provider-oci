package baremetal

import "net/http"

// Volume describes cloud block storage
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/Volume/
type Volume struct {
	ETaggedResource
	AvailabilityDomain string `json:"availabilityDomain"`
	CompartmentID      string `json:"compartmentId"`
	DisplayName        string `json:"displayName"`
	ID                 string `json:"id"`
	SizeInMBs          string `json:"sizeInMBs"`
	State              string `json:"lifecycleState"`
	TimeCreated        Time   `json:"timeCreated"`
}

// ListVolumes contains a list of block volumes
//
type ListVolumes struct {
	ResourceContainer
	Volumes []Volume
}

func (l *ListVolumes) GetList() interface{} {
	return &l.Volumes
}

// CreateVolumeRequest describes the body of a volume creation request
//
// TODO: changed to CreateVolumeDetails
type CreateVolumeRequest struct {
	AvailabilityDomain string `json:"availabilityDomain"`
	CompartmentID      string `json:"compartmentId"`
	DisplayName        string `json:"displayName,omitempty"`
}

// UpdateVolumeRequest describes the body of a volume name update request
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#UpdateVolumeRequest
type UpdateVolumeRequest struct {
	DisplayName string `json:"displayName,omitempty"`
}

// CreateVolume is used to create a cloud block storage device
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/Volume/CreateVolume
func (c *Client) CreateVolume(availabilityDomain, compartmentID string, opts ...Options) (vol *Volume, e error) {
	createRequest := CreateVolumeRequest{
		AvailabilityDomain: availabilityDomain,
		CompartmentID:      compartmentID,
	}

	if len(opts) > 0 {
		createRequest.DisplayName = opts[0].DisplayName
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
	e = response.unmarshal(vol)
	return
}

// GetVolume retrieves information about a block volume
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/Volume/GetVolume
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
	e = resp.unmarshal(vol)
	return
}

// UpdateVolume updates a volume's display name
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/Volume/UpdateVolume
func (c *Client) UpdateVolume(id string, opts ...Options) (vol *Volume, e error) {
	body := UpdateVolumeRequest{}

	if len(opts) > 0 {
		body.DisplayName = opts[0].DisplayName
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
	e = response.unmarshal(vol)
	return
}

// DeleteVolume removes a cloud block storage volume
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/Volume/DeleteVolume
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
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/Volume/ListVolumes
func (c *Client) ListVolumes(compartmentID string, opts ...Options) (vols *ListVolumes, e error) {
	reqOpts := &sdkRequestOptions{
		name:    resourceVolumes,
		ocid:    compartmentID,
		options: opts,
	}

	var resp *requestResponse
	if resp, e = c.coreApi.getRequest(reqOpts); e != nil {
		return
	}

	vols = &ListVolumes{}
	e = resp.unmarshal(vols)
	return
}
