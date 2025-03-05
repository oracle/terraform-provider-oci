// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// InstanceNetworkInterfaceAttachment Describes a network interface attachment.
type InstanceNetworkInterfaceAttachment struct {

	// The ID of the network interface attachment.
	AttachmentKey *string `mandatory:"false" json:"attachmentKey"`

	// The timestamp when the attachment initiated.
	TimeAttach *common.SDKTime `mandatory:"false" json:"timeAttach"`

	// Indicates whether the network interface is deleted when the instance is terminated.
	IsDeleteOnTermination *bool `mandatory:"false" json:"isDeleteOnTermination"`

	// The index of the device on the instance for the network interface attachment.
	DeviceIndex *int `mandatory:"false" json:"deviceIndex"`

	// The index of the network card.
	NetworkCardIndex *int `mandatory:"false" json:"networkCardIndex"`

	// The attachment state.
	Status *string `mandatory:"false" json:"status"`
}

func (m InstanceNetworkInterfaceAttachment) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InstanceNetworkInterfaceAttachment) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
