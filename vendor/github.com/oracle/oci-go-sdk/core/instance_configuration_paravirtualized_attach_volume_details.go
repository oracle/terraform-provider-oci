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

// InstanceConfigurationParavirtualizedAttachVolumeDetails The representation of InstanceConfigurationParavirtualizedAttachVolumeDetails
type InstanceConfigurationParavirtualizedAttachVolumeDetails struct {

	// A user-friendly name. Does not have to be unique, and it cannot be changed. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Whether the attachment should be created in read-only mode.
	IsReadOnly *bool `mandatory:"false" json:"isReadOnly"`
}

//GetDisplayName returns DisplayName
func (m InstanceConfigurationParavirtualizedAttachVolumeDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetIsReadOnly returns IsReadOnly
func (m InstanceConfigurationParavirtualizedAttachVolumeDetails) GetIsReadOnly() *bool {
	return m.IsReadOnly
}

func (m InstanceConfigurationParavirtualizedAttachVolumeDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m InstanceConfigurationParavirtualizedAttachVolumeDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeInstanceConfigurationParavirtualizedAttachVolumeDetails InstanceConfigurationParavirtualizedAttachVolumeDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeInstanceConfigurationParavirtualizedAttachVolumeDetails
	}{
		"paravirtualized",
		(MarshalTypeInstanceConfigurationParavirtualizedAttachVolumeDetails)(m),
	}

	return json.Marshal(&s)
}
