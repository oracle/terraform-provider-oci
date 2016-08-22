package baremetal

import "net/http"

// VolumeAttachment describes a cloud block storage attachment
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#VolumeAttachment
type VolumeAttachment struct {
	ETaggedResource
	AttachmentType     string `json:"attachmentType"`
	AvailabilityDomain string `json:"availabilityDomain"`
	CompartmentID      string `json:"compartmentId"`
	DisplayName        string `json:"displayName"`
	ID                 string `json:"id"`
	InstanceID         string `json:"instanceId"`
	State              string `json:"state"`
	TimeCreated        Time   `json:"timeCreated"`
	VolumeID           string `json:"volumeId"`
}

// VolumeAttachmentList contains a list of volume attachments
//
type VolumeAttachmentList struct {
	ResourceContainer
	VolumeAttachments []VolumeAttachment
}

func (l *VolumeAttachmentList) GetList() interface{} {
	return &l.VolumeAttachments
}

// AttachVolumeRequest describes the body of a volume attachment creation request
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#AttachVolumeRequest
type AttachVolumeRequest struct {
	CompartmentID  string `json:"compartmentId"`
	InstanceID     string `json:"instanceId"`
	AttachmentType string `json:"type"`
	VolumeID       string `json:"volumeId"`
}

//AttachVolume attaches a storage volume to the specified instance
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#attachVolume
func (c *Client) AttachVolume(compartmentID, instanceID, attachmentType, volumeID string, opts ...Options) (res *VolumeAttachment, e error) {
	createRequest := AttachVolumeRequest{
		CompartmentID:  compartmentID,
		InstanceID:     instanceID,
		AttachmentType: attachmentType,
		VolumeID:       volumeID,
	}

	reqOpts := &sdkRequestOptions{
		body:    createRequest,
		name:    resourceVolumeAttachments,
		options: opts,
	}

	var response *requestResponse
	if response, e = c.coreApi.request(http.MethodPost, reqOpts); e != nil {
		return
	}

	res = &VolumeAttachment{}
	e = response.unmarshal(res)
	return
}

// GetVolumeAttachment gets information about the specified volume attachment
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#getVolumeAttachment
func (c *Client) GetVolumeAttachment(id string, opts ...Options) (res *VolumeAttachment, e error) {
	reqOpts := &sdkRequestOptions{
		name:    resourceVolumeAttachments,
		options: opts,
		ids:     urlParts{id},
	}
	var resp *requestResponse
	if resp, e = c.coreApi.getRequest(reqOpts); e != nil {
		return
	}

	res = &VolumeAttachment{}
	e = resp.unmarshal(res)
	return
}

// DetachVolume detaches a storage volume from the specified instance
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#detachVolume
func (c *Client) DetachVolume(id string, opts ...Options) (e error) {
	reqOpts := &sdkRequestOptions{
		name:    resourceVolumeAttachments,
		options: opts,
		ids:     urlParts{id},
	}
	return c.coreApi.deleteRequest(reqOpts)
}

// ListVolumeAttachments gets a list of the volume attachments in the specified
// compartment
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/core.html#listVolumeAttachments
func (c *Client) ListVolumeAttachments(compartmentID string, opts ...Options) (res *VolumeAttachmentList, e error) {
	reqOpts := &sdkRequestOptions{
		name:    resourceVolumeAttachments,
		ocid:    compartmentID,
		options: opts,
	}

	var resp *requestResponse
	if resp, e = c.coreApi.getRequest(reqOpts); e != nil {
		return
	}

	res = &VolumeAttachmentList{}
	e = resp.unmarshal(res)
	return
}
