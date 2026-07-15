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

// ComputeCapacityDetails Capacity details of the Database Infrastructure.
type ComputeCapacityDetails struct {

	// Total CPU cores count.
	TotalCores *int `mandatory:"false" json:"totalCores"`

	// Total CPU cores count allocated..
	AllocatedCores *int `mandatory:"false" json:"allocatedCores"`

	// Total Reserved CPU cores count.
	ReservedCores *int `mandatory:"false" json:"reservedCores"`

	// Total available CPU cores count.
	AvailableCores *int `mandatory:"false" json:"availableCores"`

	// Total memory allocated, in gigabytes (GB).
	TotalMemoryInGBs *int64 `mandatory:"false" json:"totalMemoryInGBs"`

	// Memory allocated to Oracle database virtual machine cluster or Instance, in gigabytes (GB).
	UsedMemoryInGBs *int64 `mandatory:"false" json:"usedMemoryInGBs"`

	// Reserved memory, in gigabytes (GB).
	ReservedMemoryInGBs *int64 `mandatory:"false" json:"reservedMemoryInGBs"`

	// Available memory, in gigabytes (GB).
	AvailableMemoryInGBs *int64 `mandatory:"false" json:"availableMemoryInGBs"`
}

func (m ComputeCapacityDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ComputeCapacityDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
