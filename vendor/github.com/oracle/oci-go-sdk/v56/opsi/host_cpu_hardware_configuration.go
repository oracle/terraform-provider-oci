// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// HostCpuHardwareConfiguration CPU Hardware Configuration metric for the host
type HostCpuHardwareConfiguration struct {

	// Collection timestamp
	// Example: `"2020-05-06T00:00:00.000Z"`
	TimeCollected *common.SDKTime `mandatory:"true" json:"timeCollected"`

	// Total number of CPU Sockets
	TotalSockets *int `mandatory:"false" json:"totalSockets"`

	// Name of the CPU vendor
	VendorName *string `mandatory:"false" json:"vendorName"`

	// Clock frequency of the processor in megahertz
	FrequencyInMhz *float64 `mandatory:"false" json:"frequencyInMhz"`

	// Size of cache memory in megabytes
	CacheInMB *float64 `mandatory:"false" json:"cacheInMB"`

	// Model name of processor
	CpuImplementation *string `mandatory:"false" json:"cpuImplementation"`

	// CPU model
	Model *string `mandatory:"false" json:"model"`

	// Type of processor in the system
	CpuFamily *string `mandatory:"false" json:"cpuFamily"`

	// Number of cores per socket
	CoresPerSocket *int `mandatory:"false" json:"coresPerSocket"`

	// Number of threads per socket
	ThreadsPerSocket *int `mandatory:"false" json:"threadsPerSocket"`

	// Indicates if hyper-threading is enabled or not
	HyperThreadingEnabled *string `mandatory:"false" json:"hyperThreadingEnabled"`
}

//GetTimeCollected returns TimeCollected
func (m HostCpuHardwareConfiguration) GetTimeCollected() *common.SDKTime {
	return m.TimeCollected
}

func (m HostCpuHardwareConfiguration) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m HostCpuHardwareConfiguration) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeHostCpuHardwareConfiguration HostCpuHardwareConfiguration
	s := struct {
		DiscriminatorParam string `json:"metricName"`
		MarshalTypeHostCpuHardwareConfiguration
	}{
		"HOST_CPU_HARDWARE_CONFIGURATION",
		(MarshalTypeHostCpuHardwareConfiguration)(m),
	}

	return json.Marshal(&s)
}
