package baremetal

import "net/http"

// VolumeAttachment describes a cloud block storage attachment
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/VolumeAttachment/
type VolumeAttachment struct {
	OPCRequestIDUnmarshaller
	ETagUnmarshaller
	AttachmentType     string `json:"attachmentType"`
	AvailabilityDomain string `json:"availabilityDomain"`
	CompartmentID      string `json:"compartmentId"`
	DisplayName        string `json:"displayName"`
	ID                 string `json:"id"`
	InstanceID         string `json:"instanceId"`
	State              string `json:"lifecycleState"`
	TimeCreated        Time   `json:"timeCreated"`
	VolumeID           string `json:"volumeId"`
}

// ListVolumeAttachments contains a list of volume attachments
//
type ListVolumeAttachments struct {
	OPCRequestIDUnmarshaller
	NextPageUnmarshaller
	VolumeAttachments []VolumeAttachment
}

func (l *ListVolumeAttachments) GetList() interface{} {
	return &l.VolumeAttachments
}

//AttachVolume attaches a storage volume to the specified instance
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/VolumeAttachment/AttachVolume
func (c *Client) AttachVolume(attachmentType, instanceID, volumeID string, opts *CreateOptions) (res *VolumeAttachment, e error) {
	required := struct {
		AttachmentType string `header:"-" json:"attachmentType" url:"-"`
		InstanceID     string `header:"-" json:"instanceId" url:"-"`
		VolumeID       string `header:"-" json:"volumeId" url:"-"`
	}{
		AttachmentType: attachmentType,
		InstanceID:     instanceID,
		VolumeID:       volumeID,
	}

	details := &requestDetails{
		name:     resourceVolumeAttachments,
		optional: opts,
		required: required,
	}

	var resp *response
	if resp, e = c.coreApi.request(http.MethodPost, details); e != nil {
		return
	}

	res = &VolumeAttachment{}
	e = resp.unmarshal(res)
	return
}

// GetVolumeAttachment gets information about the specified volume attachment
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/VolumeAttachment/GetVolumeAttachment
func (c *Client) GetVolumeAttachment(id string) (res *VolumeAttachment, e error) {
	details := &requestDetails{
		ids:  urlParts{id},
		name: resourceVolumeAttachments,
	}

	var resp *response
	if resp, e = c.coreApi.getRequest(details); e != nil {
		return
	}

	res = &VolumeAttachment{}
	e = resp.unmarshal(res)
	return
}

// DetachVolume detaches a storage volume from the specified instance
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/Volume/DetachVolume
func (c *Client) DetachVolume(id string, opts *IfMatchOptions) (e error) {
	details := &requestDetails{
		ids:      urlParts{id},
		name:     resourceVolumeAttachments,
		optional: opts,
	}

	return c.coreApi.deleteRequest(details)
}

// ListVolumeAttachments gets a list of the volume attachments in the specified
// compartment
//
// See https://docs.us-az-phoenix-1.oracleiaas.com/api/#/en/core/20160918/VolumeAttachment/ListVolumeAttachments
func (c *Client) ListVolumeAttachments(compartmentID string, opts *ListVolumeAttachmentsOptions) (res *ListVolumeAttachments, e error) {
	details := &requestDetails{
		name:     resourceVolumeAttachments,
		optional: opts,
		required: ocidRequirement{compartmentID},
	}

	var resp *response
	if resp, e = c.coreApi.getRequest(details); e != nil {
		return
	}

	res = &ListVolumeAttachments{}
	e = resp.unmarshal(res)
	return
}
