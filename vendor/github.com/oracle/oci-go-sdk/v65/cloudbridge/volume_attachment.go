// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Bridge API
//
// API for Oracle Cloud Bridge service.
//

package cloudbridge

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// VolumeAttachment Describes volume attachment details.
type VolumeAttachment struct {

	// Indicates whether the EBS volume is deleted on instance termination.
	IsDeleteOnTermination *bool `mandatory:"false" json:"isDeleteOnTermination"`

	// The device name.
	Device *string `mandatory:"false" json:"device"`

	// The ID of the instance.
	InstanceKey *string `mandatory:"false" json:"instanceKey"`

	// The attachment state of the volume.
	Status *string `mandatory:"false" json:"status"`

	// The ID of the volume.
	VolumeKey *string `mandatory:"false" json:"volumeKey"`
}

func (m VolumeAttachment) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VolumeAttachment) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
