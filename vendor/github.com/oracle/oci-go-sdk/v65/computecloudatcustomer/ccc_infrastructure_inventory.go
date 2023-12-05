// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Compute Cloud@Customer API
//
// Use the Compute Cloud@Customer API to manage Compute Cloud@Customer infrastructures and upgrade schedules.
// For more information see Compute Cloud@Customer documentation (https://docs.cloud.oracle.com/iaas/compute-cloud-at-customer/home.htm).
//

package computecloudatcustomer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CccInfrastructureInventory Inventory for a Compute Cloud@Customer infrastructure. This information
// cannot be updated and is from the infrastructure. The information will only be available
// after the connectionState is transitioned to CONNECTED.
type CccInfrastructureInventory struct {

	// The serial number of the Compute Cloud@Customer infrastructure rack.
	SerialNumber *string `mandatory:"false" json:"serialNumber"`

	// The number of management nodes that are available and in active use
	// on the Compute Cloud@Customer infrastructure rack.
	ManagementNodeCount *int `mandatory:"false" json:"managementNodeCount"`

	// The number of compute nodes that are available and usable
	// on the Compute Cloud@Customer infrastructure rack. There is no distinction
	// of compute node type in this information.
	ComputeNodeCount *int `mandatory:"false" json:"computeNodeCount"`

	// The number of storage trays in the Compute Cloud@Customer infrastructure rack that are designated for capacity storage.
	CapacityStorageTrayCount *int `mandatory:"false" json:"capacityStorageTrayCount"`

	// The number of storage trays in the Compute Cloud@Customer infrastructure rack that are designated for performance storage.
	PerformanceStorageTrayCount *int `mandatory:"false" json:"performanceStorageTrayCount"`
}

func (m CccInfrastructureInventory) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CccInfrastructureInventory) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
