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

// InstanceConfigurationAttachVolumeDetails Volume attachmentDetails. Please see AttachVolumeDetails
type InstanceConfigurationAttachVolumeDetails interface {

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

type instanceconfigurationattachvolumedetails struct {
	JsonData    []byte
	DisplayName *string `mandatory:"false" json:"displayName"`
	IsReadOnly  *bool   `mandatory:"false" json:"isReadOnly"`
	IsShareable *bool   `mandatory:"false" json:"isShareable"`
	Type        string  `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *instanceconfigurationattachvolumedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerinstanceconfigurationattachvolumedetails instanceconfigurationattachvolumedetails
	s := struct {
		Model Unmarshalerinstanceconfigurationattachvolumedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DisplayName = s.Model.DisplayName
	m.IsReadOnly = s.Model.IsReadOnly
	m.IsShareable = s.Model.IsShareable
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *instanceconfigurationattachvolumedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {
	var err error
	switch m.Type {
	case "iscsi":
		mm := InstanceConfigurationIscsiAttachVolumeDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return m, nil
	}
}

//GetDisplayName returns DisplayName
func (m instanceconfigurationattachvolumedetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetIsReadOnly returns IsReadOnly
func (m instanceconfigurationattachvolumedetails) GetIsReadOnly() *bool {
	return m.IsReadOnly
}

//GetIsShareable returns IsShareable
func (m instanceconfigurationattachvolumedetails) GetIsShareable() *bool {
	return m.IsShareable
}

func (m instanceconfigurationattachvolumedetails) String() string {
	return common.PointerString(m)
}
