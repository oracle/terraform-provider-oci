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

// InstanceConfigurationIscsiAttachVolumeDetails The representation of InstanceConfigurationIscsiAttachVolumeDetails
type InstanceConfigurationIscsiAttachVolumeDetails struct {

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
func (m InstanceConfigurationIscsiAttachVolumeDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetIsReadOnly returns IsReadOnly
func (m InstanceConfigurationIscsiAttachVolumeDetails) GetIsReadOnly() *bool {
	return m.IsReadOnly
}

//GetIsShareable returns IsShareable
func (m InstanceConfigurationIscsiAttachVolumeDetails) GetIsShareable() *bool {
	return m.IsShareable
}

func (m InstanceConfigurationIscsiAttachVolumeDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m InstanceConfigurationIscsiAttachVolumeDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeInstanceConfigurationIscsiAttachVolumeDetails InstanceConfigurationIscsiAttachVolumeDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeInstanceConfigurationIscsiAttachVolumeDetails
	}{
		"iscsi",
		(MarshalTypeInstanceConfigurationIscsiAttachVolumeDetails)(m),
	}

	return json.Marshal(&s)
}
