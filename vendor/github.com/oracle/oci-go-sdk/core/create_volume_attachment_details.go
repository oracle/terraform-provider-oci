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

// CreateVolumeAttachmentDetails The representation of CreateVolumeAttachmentDetails
type CreateVolumeAttachmentDetails interface {

	// The OCID of the volume.
	GetVolumeId() *string

	// A user-friendly name. Does not have to be unique, and it cannot be changed. Avoid entering confidential information.
	GetDisplayName() *string

	// Whether the attachment should be created in read-only mode.
	GetIsReadOnly() *bool

	// Whether the attachment should be created in shareable mode. If an attachment
	// is created in shareable mode, then other instances can attach the same volume, provided
	// that they also create their attachments in shareable mode. Only certain volume types can
	// be attached in shareable mode. Defaults to false if not specified.
	GetIsShareable() *bool
}

type createvolumeattachmentdetails struct {
	JsonData    []byte
	VolumeId    *string `mandatory:"true" json:"volumeId"`
	DisplayName *string `mandatory:"false" json:"displayName"`
	IsReadOnly  *bool   `mandatory:"false" json:"isReadOnly"`
	IsShareable *bool   `mandatory:"false" json:"isShareable"`
	Type        string  `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *createvolumeattachmentdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatevolumeattachmentdetails createvolumeattachmentdetails
	s := struct {
		Model Unmarshalercreatevolumeattachmentdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.VolumeId = s.Model.VolumeId
	m.DisplayName = s.Model.DisplayName
	m.IsReadOnly = s.Model.IsReadOnly
	m.IsShareable = s.Model.IsShareable
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createvolumeattachmentdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {
	var err error
	switch m.Type {
	case "iscsi":
		mm := CreateIscsiVolumeAttachmentDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return m, nil
	}
}

//GetVolumeId returns VolumeId
func (m createvolumeattachmentdetails) GetVolumeId() *string {
	return m.VolumeId
}

//GetDisplayName returns DisplayName
func (m createvolumeattachmentdetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetIsReadOnly returns IsReadOnly
func (m createvolumeattachmentdetails) GetIsReadOnly() *bool {
	return m.IsReadOnly
}

//GetIsShareable returns IsShareable
func (m createvolumeattachmentdetails) GetIsShareable() *bool {
	return m.IsShareable
}

func (m createvolumeattachmentdetails) String() string {
	return common.PointerString(m)
}
