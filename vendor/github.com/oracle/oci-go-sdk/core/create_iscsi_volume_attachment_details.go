// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// CreateIscsiVolumeAttachmentDetails The representation of CreateIscsiVolumeAttachmentDetails
type CreateIscsiVolumeAttachmentDetails struct {

	// The OCID of the volume.
	VolumeId *string `mandatory:"true" json:"volumeId"`

	// A user-friendly name. Does not have to be unique, and it cannot be changed. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Whether the attachment should be created in read-only mode.
	IsReadOnly *bool `mandatory:"false" json:"isReadOnly"`

	// Whether the attachment should be created in shareable mode. If an attachment
	// is created in shareable mode, then other instances can attach the same volume, provided
	// that they also create their attachments in shareable mode. Only certain volume types can
	// be attached in shareable mode. Defaults to false if not specified.
	IsShareable *bool `mandatory:"false" json:"isShareable"`

	// Whether to use CHAP authentication for the volume attachment. Defaults to false.
	UseChap *bool `mandatory:"false" json:"useChap"`
}

//GetDisplayName returns DisplayName
func (m CreateIscsiVolumeAttachmentDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetIsReadOnly returns IsReadOnly
func (m CreateIscsiVolumeAttachmentDetails) GetIsReadOnly() *bool {
	return m.IsReadOnly
}

//GetIsShareable returns IsShareable
func (m CreateIscsiVolumeAttachmentDetails) GetIsShareable() *bool {
	return m.IsShareable
}

//GetVolumeId returns VolumeId
func (m CreateIscsiVolumeAttachmentDetails) GetVolumeId() *string {
	return m.VolumeId
}

func (m CreateIscsiVolumeAttachmentDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateIscsiVolumeAttachmentDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateIscsiVolumeAttachmentDetails CreateIscsiVolumeAttachmentDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeCreateIscsiVolumeAttachmentDetails
	}{
		"iscsi",
		(MarshalTypeCreateIscsiVolumeAttachmentDetails)(m),
	}

	return json.Marshal(&s)
}
