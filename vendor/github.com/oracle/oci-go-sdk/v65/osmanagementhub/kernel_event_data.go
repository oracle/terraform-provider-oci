// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for the operating system environments in your private data centers through a single management console. For more information, see Overview of OS Management Hub (https://docs.cloud.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// KernelEventData Kernel event data
type KernelEventData struct {
	Content *KernelEventContent `mandatory:"true" json:"content"`

	// Event count
	Count *int `mandatory:"true" json:"count"`

	// Event fingerprint
	EventFingerprint *string `mandatory:"true" json:"eventFingerprint"`

	// Event reason
	Reason *string `mandatory:"true" json:"reason"`

	// Event first occurred time
	TimeFirstOccurred *common.SDKTime `mandatory:"true" json:"timeFirstOccurred"`

	AdditionalDetails *KernelEventAdditionalDetails `mandatory:"false" json:"additionalDetails"`
}

func (m KernelEventData) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m KernelEventData) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
