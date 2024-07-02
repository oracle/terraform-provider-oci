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

// HostCpuStatistics Contains CPU statistics.
type HostCpuStatistics struct {

	// Total amount used of the resource metric type (CPU, STORAGE).
	Usage *float64 `mandatory:"true" json:"usage"`

	// The maximum allocated amount of the resource metric type  (CPU, STORAGE) for a set of databases.
	Capacity *float64 `mandatory:"true" json:"capacity"`

	// Resource utilization in percentage.
	UtilizationPercent *float64 `mandatory:"true" json:"utilizationPercent"`

	// Change in resource utilization in percentage
	UsageChangePercent *float64 `mandatory:"true" json:"usageChangePercent"`

	// The baseline utilization is a fraction of each CPU core expressed in percentages, either 12.5% or 50%. The baseline provides the minimum CPUs that can be used constantly.
	CpuBaseline *float64 `mandatory:"false" json:"cpuBaseline"`

	Load *SummaryStatistics `mandatory:"false" json:"load"`
}

// GetUsage returns Usage
func (m HostCpuStatistics) GetUsage() *float64 {
	return m.Usage
}

// GetCapacity returns Capacity
func (m HostCpuStatistics) GetCapacity() *float64 {
	return m.Capacity
}

// GetUtilizationPercent returns UtilizationPercent
func (m HostCpuStatistics) GetUtilizationPercent() *float64 {
	return m.UtilizationPercent
}

// GetUsageChangePercent returns UsageChangePercent
func (m HostCpuStatistics) GetUsageChangePercent() *float64 {
	return m.UsageChangePercent
}

func (m HostCpuStatistics) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HostCpuStatistics) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m HostCpuStatistics) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeHostCpuStatistics HostCpuStatistics
	s := struct {
		DiscriminatorParam string `json:"resourceName"`
		MarshalTypeHostCpuStatistics
	}{
		"HOST_CPU_STATISTICS",
		(MarshalTypeHostCpuStatistics)(m),
	}

	return json.Marshal(&s)
}
