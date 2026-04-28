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

// DiskGroupCapacityDetails Capacity details of the storage disk group.
type DiskGroupCapacityDetails struct {

	// The storage type for the Cloud Database Infrastructure.
	StorageType *string `mandatory:"false" json:"storageType"`

	// The total amount of logical disk space available, in gigabytes (GB).
	TotalSpaceInGbs *float64 `mandatory:"false" json:"totalSpaceInGbs"`

	// The total amount of logical disk space that is reserved for system use, in gigabytes (GB).
	ReservedSpaceInGbs *float64 `mandatory:"false" json:"reservedSpaceInGbs"`

	// The total amount of logical disk space that is currently available for use, in gigabytes (GB).
	FreeSpaceInGbs *float64 `mandatory:"false" json:"freeSpaceInGbs"`
}

func (m DiskGroupCapacityDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DiskGroupCapacityDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
