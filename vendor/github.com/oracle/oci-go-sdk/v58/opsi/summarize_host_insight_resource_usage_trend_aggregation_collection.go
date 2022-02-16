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

// SummarizeHostInsightResourceUsageTrendAggregationCollection Top level response object.
type SummarizeHostInsightResourceUsageTrendAggregationCollection struct {

	// The start timestamp that was passed into the request.
	TimeIntervalStart *common.SDKTime `mandatory:"true" json:"timeIntervalStart"`

	// The end timestamp that was passed into the request.
	TimeIntervalEnd *common.SDKTime `mandatory:"true" json:"timeIntervalEnd"`

	// Defines the type of resource metric (CPU, Physical Memory, Logical Memory)
	ResourceMetric SummarizeHostInsightResourceUsageTrendAggregationCollectionResourceMetricEnum `mandatory:"true" json:"resourceMetric"`

	// Displays usage unit (CORES, GB)
	UsageUnit UsageUnitEnum `mandatory:"true" json:"usageUnit"`

	// Time duration in milliseconds between data points (one hour or one day).
	ItemDurationInMs *int64 `mandatory:"true" json:"itemDurationInMs"`

	// Usage Data with timestamp.
	UsageData []ResourceUsageTrendAggregation `mandatory:"true" json:"usageData"`
}

func (m SummarizeHostInsightResourceUsageTrendAggregationCollection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SummarizeHostInsightResourceUsageTrendAggregationCollection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSummarizeHostInsightResourceUsageTrendAggregationCollectionResourceMetricEnum(string(m.ResourceMetric)); !ok && m.ResourceMetric != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResourceMetric: %s. Supported values are: %s.", m.ResourceMetric, strings.Join(GetSummarizeHostInsightResourceUsageTrendAggregationCollectionResourceMetricEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUsageUnitEnum(string(m.UsageUnit)); !ok && m.UsageUnit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UsageUnit: %s. Supported values are: %s.", m.UsageUnit, strings.Join(GetUsageUnitEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeHostInsightResourceUsageTrendAggregationCollectionResourceMetricEnum Enum with underlying type: string
type SummarizeHostInsightResourceUsageTrendAggregationCollectionResourceMetricEnum string

// Set of constants representing the allowable values for SummarizeHostInsightResourceUsageTrendAggregationCollectionResourceMetricEnum
const (
	SummarizeHostInsightResourceUsageTrendAggregationCollectionResourceMetricCpu           SummarizeHostInsightResourceUsageTrendAggregationCollectionResourceMetricEnum = "CPU"
	SummarizeHostInsightResourceUsageTrendAggregationCollectionResourceMetricMemory        SummarizeHostInsightResourceUsageTrendAggregationCollectionResourceMetricEnum = "MEMORY"
	SummarizeHostInsightResourceUsageTrendAggregationCollectionResourceMetricLogicalMemory SummarizeHostInsightResourceUsageTrendAggregationCollectionResourceMetricEnum = "LOGICAL_MEMORY"
)

var mappingSummarizeHostInsightResourceUsageTrendAggregationCollectionResourceMetricEnum = map[string]SummarizeHostInsightResourceUsageTrendAggregationCollectionResourceMetricEnum{
	"CPU":            SummarizeHostInsightResourceUsageTrendAggregationCollectionResourceMetricCpu,
	"MEMORY":         SummarizeHostInsightResourceUsageTrendAggregationCollectionResourceMetricMemory,
	"LOGICAL_MEMORY": SummarizeHostInsightResourceUsageTrendAggregationCollectionResourceMetricLogicalMemory,
}

// GetSummarizeHostInsightResourceUsageTrendAggregationCollectionResourceMetricEnumValues Enumerates the set of values for SummarizeHostInsightResourceUsageTrendAggregationCollectionResourceMetricEnum
func GetSummarizeHostInsightResourceUsageTrendAggregationCollectionResourceMetricEnumValues() []SummarizeHostInsightResourceUsageTrendAggregationCollectionResourceMetricEnum {
	values := make([]SummarizeHostInsightResourceUsageTrendAggregationCollectionResourceMetricEnum, 0)
	for _, v := range mappingSummarizeHostInsightResourceUsageTrendAggregationCollectionResourceMetricEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeHostInsightResourceUsageTrendAggregationCollectionResourceMetricEnumStringValues Enumerates the set of values in String for SummarizeHostInsightResourceUsageTrendAggregationCollectionResourceMetricEnum
func GetSummarizeHostInsightResourceUsageTrendAggregationCollectionResourceMetricEnumStringValues() []string {
	return []string{
		"CPU",
		"MEMORY",
		"LOGICAL_MEMORY",
	}
}

// GetMappingSummarizeHostInsightResourceUsageTrendAggregationCollectionResourceMetricEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeHostInsightResourceUsageTrendAggregationCollectionResourceMetricEnum(val string) (SummarizeHostInsightResourceUsageTrendAggregationCollectionResourceMetricEnum, bool) {
	mappingSummarizeHostInsightResourceUsageTrendAggregationCollectionResourceMetricEnumIgnoreCase := make(map[string]SummarizeHostInsightResourceUsageTrendAggregationCollectionResourceMetricEnum)
	for k, v := range mappingSummarizeHostInsightResourceUsageTrendAggregationCollectionResourceMetricEnum {
		mappingSummarizeHostInsightResourceUsageTrendAggregationCollectionResourceMetricEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingSummarizeHostInsightResourceUsageTrendAggregationCollectionResourceMetricEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
