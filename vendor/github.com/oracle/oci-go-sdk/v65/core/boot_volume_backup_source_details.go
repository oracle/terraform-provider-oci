// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
//

package core

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BootVolumeBackupSourceDetails Define volume parameters that will be utilized to define a BootVolume and a BootVolumeAttachment on the
// creation of an Instance.
type BootVolumeBackupSourceDetails struct {

	// The ocid of the VolumeBackup that will be utilized as the source for creating the Volume to attach to the
	// Instance.
	// Warning: This VolumeBackup will not longer be usable or accessible once this operation is completed. It will
	// become a part of the Image.
	VolumeBackupId *string `mandatory:"true" json:"volumeBackupId"`
}

//GetVolumeBackupId returns VolumeBackupId
func (m BootVolumeBackupSourceDetails) GetVolumeBackupId() *string {
	return m.VolumeBackupId
}

func (m BootVolumeBackupSourceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BootVolumeBackupSourceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m BootVolumeBackupSourceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeBootVolumeBackupSourceDetails BootVolumeBackupSourceDetails
	s := struct {
		DiscriminatorParam string `json:"volumeType"`
		MarshalTypeBootVolumeBackupSourceDetails
	}{
		"BOOT_VOLUME",
		(MarshalTypeBootVolumeBackupSourceDetails)(m),
	}

	return json.Marshal(&s)
}
