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
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// SummarizeHostInsightResourceUsageAggregation Resource usage summation for the current time period.
type SummarizeHostInsightResourceUsageAggregation struct {

	// The start timestamp that was passed into the request.
	TimeIntervalStart *common.SDKTime `mandatory:"true" json:"timeIntervalStart"`

	// The end timestamp that was passed into the request.
	TimeIntervalEnd *common.SDKTime `mandatory:"true" json:"timeIntervalEnd"`

	// Defines the type of resource metric (CPU, Physical Memory, Logical Memory)
	ResourceMetric SummarizeHostInsightResourceUsageAggregationResourceMetricEnum `mandatory:"true" json:"resourceMetric"`

	// Displays usage unit (CORES, GB)
	UsageUnit UsageUnitEnum `mandatory:"true" json:"usageUnit"`

	// Total amount used of the resource metric type (CPU, STORAGE).
	Usage *float64 `mandatory:"true" json:"usage"`

	// The maximum allocated amount of the resource metric type  (CPU, STORAGE).
	Capacity *float64 `mandatory:"true" json:"capacity"`

	// Percentage change in resource usage during the current period calculated using linear regression functions
	UsageChangePercent *float64 `mandatory:"true" json:"usageChangePercent"`
}

func (m SummarizeHostInsightResourceUsageAggregation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SummarizeHostInsightResourceUsageAggregation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSummarizeHostInsightResourceUsageAggregationResourceMetricEnum(string(m.ResourceMetric)); !ok && m.ResourceMetric != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResourceMetric: %s. Supported values are: %s.", m.ResourceMetric, strings.Join(GetSummarizeHostInsightResourceUsageAggregationResourceMetricEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUsageUnitEnum(string(m.UsageUnit)); !ok && m.UsageUnit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UsageUnit: %s. Supported values are: %s.", m.UsageUnit, strings.Join(GetUsageUnitEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeHostInsightResourceUsageAggregationResourceMetricEnum Enum with underlying type: string
type SummarizeHostInsightResourceUsageAggregationResourceMetricEnum string

// Set of constants representing the allowable values for SummarizeHostInsightResourceUsageAggregationResourceMetricEnum
const (
	SummarizeHostInsightResourceUsageAggregationResourceMetricCpu           SummarizeHostInsightResourceUsageAggregationResourceMetricEnum = "CPU"
	SummarizeHostInsightResourceUsageAggregationResourceMetricMemory        SummarizeHostInsightResourceUsageAggregationResourceMetricEnum = "MEMORY"
	SummarizeHostInsightResourceUsageAggregationResourceMetricLogicalMemory SummarizeHostInsightResourceUsageAggregationResourceMetricEnum = "LOGICAL_MEMORY"
)

var mappingSummarizeHostInsightResourceUsageAggregationResourceMetricEnum = map[string]SummarizeHostInsightResourceUsageAggregationResourceMetricEnum{
	"CPU":            SummarizeHostInsightResourceUsageAggregationResourceMetricCpu,
	"MEMORY":         SummarizeHostInsightResourceUsageAggregationResourceMetricMemory,
	"LOGICAL_MEMORY": SummarizeHostInsightResourceUsageAggregationResourceMetricLogicalMemory,
}

// GetSummarizeHostInsightResourceUsageAggregationResourceMetricEnumValues Enumerates the set of values for SummarizeHostInsightResourceUsageAggregationResourceMetricEnum
func GetSummarizeHostInsightResourceUsageAggregationResourceMetricEnumValues() []SummarizeHostInsightResourceUsageAggregationResourceMetricEnum {
	values := make([]SummarizeHostInsightResourceUsageAggregationResourceMetricEnum, 0)
	for _, v := range mappingSummarizeHostInsightResourceUsageAggregationResourceMetricEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeHostInsightResourceUsageAggregationResourceMetricEnumStringValues Enumerates the set of values in String for SummarizeHostInsightResourceUsageAggregationResourceMetricEnum
func GetSummarizeHostInsightResourceUsageAggregationResourceMetricEnumStringValues() []string {
	return []string{
		"CPU",
		"MEMORY",
		"LOGICAL_MEMORY",
	}
}

// GetMappingSummarizeHostInsightResourceUsageAggregationResourceMetricEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeHostInsightResourceUsageAggregationResourceMetricEnum(val string) (SummarizeHostInsightResourceUsageAggregationResourceMetricEnum, bool) {
	mappingSummarizeHostInsightResourceUsageAggregationResourceMetricEnumIgnoreCase := make(map[string]SummarizeHostInsightResourceUsageAggregationResourceMetricEnum)
	for k, v := range mappingSummarizeHostInsightResourceUsageAggregationResourceMetricEnum {
		mappingSummarizeHostInsightResourceUsageAggregationResourceMetricEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingSummarizeHostInsightResourceUsageAggregationResourceMetricEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
