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

// HostCpuUsage CPU Usage metric for the host
type HostCpuUsage struct {

	// Collection timestamp
	// Example: `"2020-05-06T00:00:00.000Z"`
	TimeCollected *common.SDKTime `mandatory:"true" json:"timeCollected"`

	// Percentage of CPU time spent in user mode
	CpuUserModeInPercent *float32 `mandatory:"false" json:"cpuUserModeInPercent"`

	// Percentage of CPU time spent in system mode
	CpuSystemModeInPercent *float32 `mandatory:"false" json:"cpuSystemModeInPercent"`

	// Amount of CPU Time spent in seconds
	CpuUsageInSec *float64 `mandatory:"false" json:"cpuUsageInSec"`

	// Amount of CPU Time spent in percentage
	CpuUtilizationInPercent *float32 `mandatory:"false" json:"cpuUtilizationInPercent"`

	// Amount of CPU time stolen in percentage
	CpuStolenInPercent *float32 `mandatory:"false" json:"cpuStolenInPercent"`

	// Amount of CPU idle time in percentage
	CpuIdleInPercent *float32 `mandatory:"false" json:"cpuIdleInPercent"`

	// Load average in the last 1 minute
	CpuLoad1min *float32 `mandatory:"false" json:"cpuLoad1min"`

	// Load average in the last 5 minutes
	CpuLoad5min *float32 `mandatory:"false" json:"cpuLoad5min"`

	// Load average in the last 15 minutes
	CpuLoad15min *float32 `mandatory:"false" json:"cpuLoad15min"`
}

//GetTimeCollected returns TimeCollected
func (m HostCpuUsage) GetTimeCollected() *common.SDKTime {
	return m.TimeCollected
}

func (m HostCpuUsage) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m HostCpuUsage) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeHostCpuUsage HostCpuUsage
	s := struct {
		DiscriminatorParam string `json:"metricName"`
		MarshalTypeHostCpuUsage
	}{
		"HOST_CPU_USAGE",
		(MarshalTypeHostCpuUsage)(m),
	}

	return json.Marshal(&s)
}
