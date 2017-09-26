// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package baremetal

// VolumeAttachment describes a cloud block storage attachment
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/VolumeAttachment/
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
	// IScsiVolumeAttachment extended attributes
	// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/IScsiVolumeAttachment/
	CHAPSecret   string `json:"chapSecret,omitempty"`
	CHAPUsername string `json:"chapUsername,omitempty"`
	IPv4         string `json:"ipv4,omitempty"`
	IQN          string `json:"iqn,omitempty"` // iSCSI Qualified Name per RFC 3720
	Port         int    `json:"port,omitempty"`
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
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/VolumeAttachment/AttachVolume
func (c *Client) AttachVolume(attachmentType, instanceID, volumeID string, opts *CreateOptions) (res *VolumeAttachment, e error) {
	required := struct {
		AttachmentType string `header:"-" json:"type" url:"-"`
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
	if resp, e = c.coreApi.postRequest(details); e != nil {
		return
	}

	res = &VolumeAttachment{}
	e = resp.unmarshal(res)
	return
}

// GetVolumeAttachment gets information about the specified volume attachment
//
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/VolumeAttachment/GetVolumeAttachment
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
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/Volume/DetachVolume
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
// See https://docs.us-phoenix-1.oraclecloud.com/api/#/en/iaas/20160918/VolumeAttachment/ListVolumeAttachments
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
