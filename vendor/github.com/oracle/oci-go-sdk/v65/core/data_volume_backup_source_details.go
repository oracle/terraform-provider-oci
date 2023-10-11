// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.cloud.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DataVolumeBackupSourceDetails Define volume parameters that will be utilized to define a Volume and VolumeAttachment on the
// creation of an Instance.
type DataVolumeBackupSourceDetails struct {

	// The ocid of the VolumeBackup that will be utilized as the source for creating the Volume to attach to the
	// Instance.
	// Warning: This VolumeBackup will not longer be usable or accessible once this operation is completed. It will
	// become a part of the Image.
	VolumeBackupId *string `mandatory:"true" json:"volumeBackupId"`

	// A user-friendly name for the Volume. It does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`
}

// GetVolumeBackupId returns VolumeBackupId
func (m DataVolumeBackupSourceDetails) GetVolumeBackupId() *string {
	return m.VolumeBackupId
}

func (m DataVolumeBackupSourceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DataVolumeBackupSourceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DataVolumeBackupSourceDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDataVolumeBackupSourceDetails DataVolumeBackupSourceDetails
	s := struct {
		DiscriminatorParam string `json:"volumeType"`
		MarshalTypeDataVolumeBackupSourceDetails
	}{
		"DATA_VOLUME",
		(MarshalTypeDataVolumeBackupSourceDetails)(m),
	}

	return json.Marshal(&s)
}
