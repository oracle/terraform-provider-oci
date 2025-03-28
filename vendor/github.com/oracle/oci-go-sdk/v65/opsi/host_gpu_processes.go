// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.oracle.com/iaas/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// HostGpuProcesses GPU processes metrics, processes using GPUs.
type HostGpuProcesses struct {

	// Collection timestamp
	// Example: `"2020-05-06T00:00:00.000Z"`
	TimeCollected *common.SDKTime `mandatory:"true" json:"timeCollected"`

	// GPU Identifier
	GpuId *int `mandatory:"false" json:"gpuId"`

	// Process Identifier
	Pid *int `mandatory:"false" json:"pid"`

	// Process Name (process using GPU)
	ProcessName *string `mandatory:"false" json:"processName"`

	// Process elapsed time
	ElapsedTime *float64 `mandatory:"false" json:"elapsedTime"`

	// Memory Used by Process in MBs
	GpuMemoryUsage *float64 `mandatory:"false" json:"gpuMemoryUsage"`
}

// GetTimeCollected returns TimeCollected
func (m HostGpuProcesses) GetTimeCollected() *common.SDKTime {
	return m.TimeCollected
}

func (m HostGpuProcesses) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HostGpuProcesses) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m HostGpuProcesses) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeHostGpuProcesses HostGpuProcesses
	s := struct {
		DiscriminatorParam string `json:"metricName"`
		MarshalTypeHostGpuProcesses
	}{
		"HOST_GPU_PROCESSES",
		(MarshalTypeHostGpuProcesses)(m),
	}

	return json.Marshal(&s)
}
