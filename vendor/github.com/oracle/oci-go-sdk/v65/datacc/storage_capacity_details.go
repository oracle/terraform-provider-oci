// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Infrastructure Cloud@Customer Service API
//
// API for Database Infrastructure Cloud@Customer Service. Use this API to manage Database Infrastructure VM clusters, Application VMs, and related resources.
//

package datacc

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// StorageCapacityDetails Capacity details of the Storage disk group.
type StorageCapacityDetails struct {

	// The Disk redundancy for Database Infrastructure storage.
	StorageDiskRedundancy StorageDiskRedundancyEnum `mandatory:"false" json:"storageDiskRedundancy,omitempty"`

	// Disk group name.
	DiskGroup *string `mandatory:"false" json:"diskGroup"`

	// The total amount of physical disk space available in a disk group, in gigabytes (GB).
	PhyTotalSpaceInGBs *int64 `mandatory:"false" json:"phyTotalSpaceInGBs"`

	// The total amount of physical disk space that is reserved for system use in a disk group, in gigabytes (GB).
	PhyReservedSpaceInGBs *int64 `mandatory:"false" json:"phyReservedSpaceInGBs"`

	// The total amount of physical disk space that is currently available for use in a disk group, in gigabytes (GB).
	PhyFreeSpaceInGBs *int64 `mandatory:"false" json:"phyFreeSpaceInGBs"`

	// The total amount of logical disk space that is currently available for use in a disk group, in gigabytes (GB).
	LogicalFreeSpaceInGBs *int64 `mandatory:"false" json:"logicalFreeSpaceInGBs"`
}

func (m StorageCapacityDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m StorageCapacityDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingStorageDiskRedundancyEnum(string(m.StorageDiskRedundancy)); !ok && m.StorageDiskRedundancy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for StorageDiskRedundancy: %s. Supported values are: %s.", m.StorageDiskRedundancy, strings.Join(GetStorageDiskRedundancyEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
