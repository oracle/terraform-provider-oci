// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// HostMemoryUsage Memory usage metric for the host
type HostMemoryUsage struct {

	// Collection timestamp
	// Example: `"2020-05-06T00:00:00.000Z"`
	TimeCollected *common.SDKTime `mandatory:"true" json:"timeCollected"`

	// Amount of physical memory used in gigabytes
	MemoryUsedInGB *float64 `mandatory:"false" json:"memoryUsedInGB"`

	// Amount of physical memory used in percentage
	MemoryUtilizationInPercent *float32 `mandatory:"false" json:"memoryUtilizationInPercent"`

	// Load on memory in gigabytes
	MemoryLoadInGB *float64 `mandatory:"false" json:"memoryLoadInGB"`

	// Amount of usable physical memory in kilobytes
	RealMemoryInKB *float64 `mandatory:"false" json:"realMemoryInKB"`

	// Amount of available physical memory in kilobytes
	FreeMemoryInKB *float64 `mandatory:"false" json:"freeMemoryInKB"`

	// Memory used excluding buffers and cache in gigabytes
	LogicalMemoryUsedInGB *float64 `mandatory:"false" json:"logicalMemoryUsedInGB"`

	// Amount of logical memory used in percentage
	LogicalMemoryUtilizationInPercent *float32 `mandatory:"false" json:"logicalMemoryUtilizationInPercent"`

	// Amount of avaiable virtual memory in kilobytes
	FreeLogicalMemoryInKB *float64 `mandatory:"false" json:"freeLogicalMemoryInKB"`

	// Number of major page faults
	MajorPageFaults *int `mandatory:"false" json:"majorPageFaults"`

	// Amount of available swap space in kilobytes
	SwapFreeInKB *float64 `mandatory:"false" json:"swapFreeInKB"`

	// Amount of memory used for anon huge pages in kilobytes
	AnonHugePagesInKB *float64 `mandatory:"false" json:"anonHugePagesInKB"`

	// Number of available huge pages
	HugePagesFree *int `mandatory:"false" json:"hugePagesFree"`

	// Number of reserved huge pages
	HugePagesReserved *int `mandatory:"false" json:"hugePagesReserved"`

	// Number of surplus huge pages
	HugePagesSurplus *int `mandatory:"false" json:"hugePagesSurplus"`
}

// GetTimeCollected returns TimeCollected
func (m HostMemoryUsage) GetTimeCollected() *common.SDKTime {
	return m.TimeCollected
}

func (m HostMemoryUsage) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HostMemoryUsage) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m HostMemoryUsage) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeHostMemoryUsage HostMemoryUsage
	s := struct {
		DiscriminatorParam string `json:"metricName"`
		MarshalTypeHostMemoryUsage
	}{
		"HOST_MEMORY_USAGE",
		(MarshalTypeHostMemoryUsage)(m),
	}

	return json.Marshal(&s)
}
