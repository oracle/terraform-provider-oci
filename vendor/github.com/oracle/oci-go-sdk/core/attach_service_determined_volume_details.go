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

// AttachServiceDeterminedVolumeDetails The representation of AttachServiceDeterminedVolumeDetails
type AttachServiceDeterminedVolumeDetails struct {

	// The OCID of the instance.
	InstanceId *string `mandatory:"true" json:"instanceId"`

	// The OCID of the volume.
	VolumeId *string `mandatory:"true" json:"volumeId"`

	// A user-friendly name. Does not have to be unique, and it cannot be changed. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Whether the attachment was created in read-only mode.
	IsReadOnly *bool `mandatory:"false" json:"isReadOnly"`

	// Whether the attachment should be created in shareable mode. If an attachment
	// is created in shareable mode, then other instances can attach the same volume, provided
	// that they also create their attachments in shareable mode. Only certain volume types can
	// be attached in shareable mode. Defaults to false if not specified.
	IsShareable *bool `mandatory:"false" json:"isShareable"`
}

//GetDisplayName returns DisplayName
func (m AttachServiceDeterminedVolumeDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetInstanceId returns InstanceId
func (m AttachServiceDeterminedVolumeDetails) GetInstanceId() *string {
	return m.InstanceId
}

//GetIsReadOnly returns IsReadOnly
func (m AttachServiceDeterminedVolumeDetails) GetIsReadOnly() *bool {
	return m.IsReadOnly
}

//GetIsShareable returns IsShareable
func (m AttachServiceDeterminedVolumeDetails) GetIsShareable() *bool {
	return m.IsShareable
}

//GetVolumeId returns VolumeId
func (m AttachServiceDeterminedVolumeDetails) GetVolumeId() *string {
	return m.VolumeId
}

func (m AttachServiceDeterminedVolumeDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m AttachServiceDeterminedVolumeDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAttachServiceDeterminedVolumeDetails AttachServiceDeterminedVolumeDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeAttachServiceDeterminedVolumeDetails
	}{
		"service_determined",
		(MarshalTypeAttachServiceDeterminedVolumeDetails)(m),
	}

	return json.Marshal(&s)
}
