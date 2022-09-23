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
// The required permissions are documented in the
// Details for the Core Services (https://docs.cloud.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InstanceFirmwareCollection The firmware inventory for a baremetal instance. This inventory
// includes information about the devices attached to the host of the
// instance.
type InstanceFirmwareCollection struct {

	// Instance id
	Id *string `mandatory:"true" json:"id"`

	// Shape of the instance
	Shape *string `mandatory:"true" json:"shape"`

	// List of Devices and installed firmware versions information
	Items []InstanceFirmwareSummary `mandatory:"true" json:"items"`

	// Host serial for the bare metal host of this instance
	HostSerial *string `mandatory:"false" json:"hostSerial"`

	// The time the inventory was detected for this host. An RFC3339 formatted datetime string
	TimeLastDetected *common.SDKTime `mandatory:"false" json:"timeLastDetected"`
}

func (m InstanceFirmwareCollection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InstanceFirmwareCollection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
