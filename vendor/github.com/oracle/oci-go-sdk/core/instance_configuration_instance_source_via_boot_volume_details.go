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

// InstanceConfigurationInstanceSourceViaBootVolumeDetails The representation of InstanceConfigurationInstanceSourceViaBootVolumeDetails
type InstanceConfigurationInstanceSourceViaBootVolumeDetails struct {

	// The OCID of the boot volume used to boot the instance.
	BootVolumeId *string `mandatory:"false" json:"bootVolumeId"`
}

func (m InstanceConfigurationInstanceSourceViaBootVolumeDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m InstanceConfigurationInstanceSourceViaBootVolumeDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeInstanceConfigurationInstanceSourceViaBootVolumeDetails InstanceConfigurationInstanceSourceViaBootVolumeDetails
	s := struct {
		DiscriminatorParam string `json:"sourceType"`
		MarshalTypeInstanceConfigurationInstanceSourceViaBootVolumeDetails
	}{
		"bootVolume",
		(MarshalTypeInstanceConfigurationInstanceSourceViaBootVolumeDetails)(m),
	}

	return json.Marshal(&s)
}
