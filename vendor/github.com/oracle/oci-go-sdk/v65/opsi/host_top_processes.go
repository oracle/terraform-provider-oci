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

// HostTopProcesses Top Processes metric for the host
type HostTopProcesses struct {

	// Collection timestamp
	// Example: `"2020-05-06T00:00:00.000Z"`
	TimeCollected *common.SDKTime `mandatory:"true" json:"timeCollected"`

	// process id
	Pid *float32 `mandatory:"false" json:"pid"`

	// User that started the process
	UserName *string `mandatory:"false" json:"userName"`

	// Memory utilization percentage
	MemoryUtilizationPercent *float64 `mandatory:"false" json:"memoryUtilizationPercent"`

	// CPU utilization percentage
	CpuUtilizationPercent *float64 `mandatory:"false" json:"cpuUtilizationPercent"`

	// CPU usage in seconds
	CpuUsageInSeconds *float64 `mandatory:"false" json:"cpuUsageInSeconds"`

	// Command line executed for the process
	Command *string `mandatory:"false" json:"command"`

	// Virtual memory in megabytes
	VirtualMemoryInMBs *float64 `mandatory:"false" json:"virtualMemoryInMBs"`

	// Physical memory in megabytes
	PhysicalMemoryInMBs *float64 `mandatory:"false" json:"physicalMemoryInMBs"`

	// Process Start Time
	// Example: `"2020-03-31T00:00:00.000Z"`
	StartTime *common.SDKTime `mandatory:"false" json:"startTime"`

	// Number of processes running at the time of collection
	TotalProcesses *float32 `mandatory:"false" json:"totalProcesses"`

	// Container id if this process corresponds to a running container in the host
	ContainerId *string `mandatory:"false" json:"containerId"`

	// Bytes Read
	DiskBytesRead *float64 `mandatory:"false" json:"diskBytesRead"`

	// Bytes Written
	DiskBytesWritten *float64 `mandatory:"false" json:"diskBytesWritten"`

	// Read transactions per second
	DiskIopsRead *float64 `mandatory:"false" json:"diskIopsRead"`

	// Write transactions per second
	DiskIopsWritten *float64 `mandatory:"false" json:"diskIopsWritten"`

	// IO Transactions per second
	DiskIops *float64 `mandatory:"false" json:"diskIops"`
}

// GetTimeCollected returns TimeCollected
func (m HostTopProcesses) GetTimeCollected() *common.SDKTime {
	return m.TimeCollected
}

func (m HostTopProcesses) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HostTopProcesses) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m HostTopProcesses) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeHostTopProcesses HostTopProcesses
	s := struct {
		DiscriminatorParam string `json:"metricName"`
		MarshalTypeHostTopProcesses
	}{
		"HOST_TOP_PROCESSES",
		(MarshalTypeHostTopProcesses)(m),
	}

	return json.Marshal(&s)
}
