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
	"github.com/oracle/oci-go-sdk/v56/common"
)

// SummarizeHostInsightResourceCapacityTrendAggregationCollection Top level response object.
type SummarizeHostInsightResourceCapacityTrendAggregationCollection struct {

	// The start timestamp that was passed into the request.
	TimeIntervalStart *common.SDKTime `mandatory:"true" json:"timeIntervalStart"`

	// The end timestamp that was passed into the request.
	TimeIntervalEnd *common.SDKTime `mandatory:"true" json:"timeIntervalEnd"`

	// Defines the type of resource metric (CPU, Physical Memory, Logical Memory)
	ResourceMetric SummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum `mandatory:"true" json:"resourceMetric"`

	// Displays usage unit (CORES, GB)
	UsageUnit UsageUnitEnum `mandatory:"true" json:"usageUnit"`

	// Time duration in milliseconds between data points (one hour or one day).
	ItemDurationInMs *int64 `mandatory:"true" json:"itemDurationInMs"`

	// Capacity Data with timestamp.
	CapacityData []HostResourceCapacityTrendAggregation `mandatory:"true" json:"capacityData"`
}

func (m SummarizeHostInsightResourceCapacityTrendAggregationCollection) String() string {
	return common.PointerString(m)
}

// SummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum Enum with underlying type: string
type SummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum string

// Set of constants representing the allowable values for SummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum
const (
	SummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricCpu           SummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum = "CPU"
	SummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricMemory        SummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum = "MEMORY"
	SummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricLogicalMemory SummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum = "LOGICAL_MEMORY"
)

var mappingSummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetric = map[string]SummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum{
	"CPU":            SummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricCpu,
	"MEMORY":         SummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricMemory,
	"LOGICAL_MEMORY": SummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricLogicalMemory,
}

// GetSummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricEnumValues Enumerates the set of values for SummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum
func GetSummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricEnumValues() []SummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum {
	values := make([]SummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetricEnum, 0)
	for _, v := range mappingSummarizeHostInsightResourceCapacityTrendAggregationCollectionResourceMetric {
		values = append(values, v)
	}
	return values
}
