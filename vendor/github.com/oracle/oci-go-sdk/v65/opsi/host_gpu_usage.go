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

// HostGpuUsage GPU performance metrics
type HostGpuUsage struct {

	// Collection timestamp
	// Example: `"2020-05-06T00:00:00.000Z"`
	TimeCollected *common.SDKTime `mandatory:"true" json:"timeCollected"`

	// GPU Identifier
	GpuId *int `mandatory:"false" json:"gpuId"`

	// GPU Utilization Percent
	Utilization *float64 `mandatory:"false" json:"utilization"`

	// GPU Memory Utilization Percent
	MemoryUtilization *float64 `mandatory:"false" json:"memoryUtilization"`

	// GPU Power Draw in Watts
	PowerDraw *float64 `mandatory:"false" json:"powerDraw"`

	// GPU Temperature in Celsius
	Temperature *float64 `mandatory:"false" json:"temperature"`

	// GPU Fan Utilization
	FanUtilization *float64 `mandatory:"false" json:"fanUtilization"`

	// GPU Graphics (Shader) Clock Speed
	ClockSpeedGraphics *float64 `mandatory:"false" json:"clockSpeedGraphics"`

	// GPU SM (Streaming Multiprocessor) Clock Speed
	ClockSpeedSm *float64 `mandatory:"false" json:"clockSpeedSm"`

	// GPU Video Clock Speed
	ClockSpeedVideo *float64 `mandatory:"false" json:"clockSpeedVideo"`

	// GPU Memory Clock Speed
	ClockSpeedMemory *float64 `mandatory:"false" json:"clockSpeedMemory"`

	// GPU Performance State
	PerformanceState *float64 `mandatory:"false" json:"performanceState"`

	// GPU ECC Single Bit Errors
	EccSingleBitErrors *int `mandatory:"false" json:"eccSingleBitErrors"`

	// GPU ECC Double Bit Errors
	EccDoubleBitErrors *int `mandatory:"false" json:"eccDoubleBitErrors"`

	// Nothing running on CPU, clocks are idle
	ClockEventIdle *int `mandatory:"false" json:"clockEventIdle"`

	// HW Thermal Slowdown (reducing the core clocks by a factor of 2 or more) is engaged. Temp too high
	ClockEventHwThermalSlowDown *int `mandatory:"false" json:"clockEventHwThermalSlowDown"`

	// SW Power Scaling algorithm is reducing the clocks below requested clocks because the GPU is consuming too much power
	ClockEventSwPowerCap *int `mandatory:"false" json:"clockEventSwPowerCap"`

	// GPU clocks are limited by applications clocks setting
	ClockEventAppClockSetting *int `mandatory:"false" json:"clockEventAppClockSetting"`

	// HW Power Brake Slowdown (reducing the core clocks by a factor of 2 or more) is engaged
	ClockEventHwPowerBreak *int `mandatory:"false" json:"clockEventHwPowerBreak"`

	// SW Thermal capping algorithm is reducing clocks below requested clocks because GPU temperature is higher than Max Operating Temp
	ClockEventSwThermalSlowdown *int `mandatory:"false" json:"clockEventSwThermalSlowdown"`

	// HW Power Brake Slowdown (reducing the core clocks by a factor of 2 or more) is engaged
	ClockEventSyncBoost *int `mandatory:"false" json:"clockEventSyncBoost"`
}

// GetTimeCollected returns TimeCollected
func (m HostGpuUsage) GetTimeCollected() *common.SDKTime {
	return m.TimeCollected
}

func (m HostGpuUsage) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HostGpuUsage) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m HostGpuUsage) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeHostGpuUsage HostGpuUsage
	s := struct {
		DiscriminatorParam string `json:"metricName"`
		MarshalTypeHostGpuUsage
	}{
		"HOST_GPU_USAGE",
		(MarshalTypeHostGpuUsage)(m),
	}

	return json.Marshal(&s)
}
