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

// InstanceConfigurationVolumeSourceFromVolumeBackupDetails Specifies the volume backup.
type InstanceConfigurationVolumeSourceFromVolumeBackupDetails struct {

	// The OCID of the volume backup.
	Id *string `mandatory:"false" json:"id"`
}

func (m InstanceConfigurationVolumeSourceFromVolumeBackupDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m InstanceConfigurationVolumeSourceFromVolumeBackupDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeInstanceConfigurationVolumeSourceFromVolumeBackupDetails InstanceConfigurationVolumeSourceFromVolumeBackupDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeInstanceConfigurationVolumeSourceFromVolumeBackupDetails
	}{
		"volumeBackup",
		(MarshalTypeInstanceConfigurationVolumeSourceFromVolumeBackupDetails)(m),
	}

	return json.Marshal(&s)
}
