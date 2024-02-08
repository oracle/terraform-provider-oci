// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TopProcessesUsage Aggregated data for top processes on a specific date.
type TopProcessesUsage struct {

	// Command line and arguments used to launch process.
	Command *string `mandatory:"true" json:"command"`

	// Unique identifier for a process.
	ProcessHash *string `mandatory:"true" json:"processHash"`

	// Process CPU usage.
	CpuUsage *float64 `mandatory:"true" json:"cpuUsage"`

	// Process CPU utilization percentage.
	CpuUtilization *float64 `mandatory:"true" json:"cpuUtilization"`

	// Process memory utilization percentage.
	MemoryUtilization *float64 `mandatory:"true" json:"memoryUtilization"`

	// Process virtual memory in Megabytes.
	VirtualMemoryInMBs *float64 `mandatory:"true" json:"virtualMemoryInMBs"`

	// Procress physical memory in Megabytes.
	PhysicalMemoryInMBs *float64 `mandatory:"true" json:"physicalMemoryInMBs"`

	// Maximum number of processes running at time of collection.
	MaxProcessCount *int `mandatory:"true" json:"maxProcessCount"`
}

func (m TopProcessesUsage) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TopProcessesUsage) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
