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

// HostGpuConfiguration GPU configuration metrics
type HostGpuConfiguration struct {

	// Collection timestamp
	// Example: `"2020-05-06T00:00:00.000Z"`
	TimeCollected *common.SDKTime `mandatory:"true" json:"timeCollected"`

	// GPU Identifier
	GpuId *int `mandatory:"true" json:"gpuId"`

	// GPU Product Name
	ProductName *string `mandatory:"true" json:"productName"`

	// GPU Vendor
	Vendor *string `mandatory:"true" json:"vendor"`

	// Bus Identifier
	BusId *string `mandatory:"true" json:"busId"`

	// Bus Width
	BusWidth *int `mandatory:"true" json:"busWidth"`

	// Power Capacity
	TotalPower *float64 `mandatory:"true" json:"totalPower"`

	// Total Memory Allocated to GPU
	TotalMemory *float64 `mandatory:"true" json:"totalMemory"`

	// Max Video Clock Speed
	TotalVideoClockSpeed *float64 `mandatory:"true" json:"totalVideoClockSpeed"`

	// Max SM (Streaming Multiprocessor) Clock Speed
	TotalSmClockSpeed *float64 `mandatory:"true" json:"totalSmClockSpeed"`

	// Max Graphics Clock Speed
	TotalGraphicsClockSpeed *float64 `mandatory:"true" json:"totalGraphicsClockSpeed"`

	// Max Memory Clock Speed
	TotalMemoryClockSpeed *float64 `mandatory:"true" json:"totalMemoryClockSpeed"`

	// CUDA library version
	CudaVersion *string `mandatory:"true" json:"cudaVersion"`

	// GPU Driver version
	DriverVersion *string `mandatory:"true" json:"driverVersion"`

	// GPU Capabilities
	GpuCapabilities *string `mandatory:"false" json:"gpuCapabilities"`
}

// GetTimeCollected returns TimeCollected
func (m HostGpuConfiguration) GetTimeCollected() *common.SDKTime {
	return m.TimeCollected
}

func (m HostGpuConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HostGpuConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m HostGpuConfiguration) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeHostGpuConfiguration HostGpuConfiguration
	s := struct {
		DiscriminatorParam string `json:"metricName"`
		MarshalTypeHostGpuConfiguration
	}{
		"HOST_GPU_CONFIGURATION",
		(MarshalTypeHostGpuConfiguration)(m),
	}

	return json.Marshal(&s)
}
