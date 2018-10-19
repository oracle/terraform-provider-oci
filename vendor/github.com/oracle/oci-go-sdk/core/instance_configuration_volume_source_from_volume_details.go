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

// InstanceConfigurationVolumeSourceFromVolumeDetails Specifies the source volume.
type InstanceConfigurationVolumeSourceFromVolumeDetails struct {

	// The OCID of the volume.
	Id *string `mandatory:"false" json:"id"`
}

func (m InstanceConfigurationVolumeSourceFromVolumeDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m InstanceConfigurationVolumeSourceFromVolumeDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeInstanceConfigurationVolumeSourceFromVolumeDetails InstanceConfigurationVolumeSourceFromVolumeDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeInstanceConfigurationVolumeSourceFromVolumeDetails
	}{
		"volume",
		(MarshalTypeInstanceConfigurationVolumeSourceFromVolumeDetails)(m),
	}

	return json.Marshal(&s)
}
