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

// SummarizeHostInsightResourceStatisticsAggregationCollection Returns list of hosts with resource statistics like usage, capacity, utilization, usage change percent and load.
type SummarizeHostInsightResourceStatisticsAggregationCollection struct {

	// The start timestamp that was passed into the request.
	TimeIntervalStart *common.SDKTime `mandatory:"true" json:"timeIntervalStart"`

	// The end timestamp that was passed into the request.
	TimeIntervalEnd *common.SDKTime `mandatory:"true" json:"timeIntervalEnd"`

	// Defines the type of resource metric (CPU, Physical Memory, Logical Memory)
	ResourceMetric SummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricEnum `mandatory:"true" json:"resourceMetric"`

	// Displays usage unit.
	UsageUnit UsageUnitEnum `mandatory:"true" json:"usageUnit"`

	// Collection of Resource Statistics items
	Items []HostInsightResourceStatisticsAggregation `mandatory:"true" json:"items"`
}

func (m SummarizeHostInsightResourceStatisticsAggregationCollection) String() string {
	return common.PointerString(m)
}

// SummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricEnum Enum with underlying type: string
type SummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricEnum string

// Set of constants representing the allowable values for SummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricEnum
const (
	SummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricCpu           SummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricEnum = "CPU"
	SummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricMemory        SummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricEnum = "MEMORY"
	SummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricLogicalMemory SummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricEnum = "LOGICAL_MEMORY"
)

var mappingSummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetric = map[string]SummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricEnum{
	"CPU":            SummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricCpu,
	"MEMORY":         SummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricMemory,
	"LOGICAL_MEMORY": SummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricLogicalMemory,
}

// GetSummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricEnumValues Enumerates the set of values for SummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricEnum
func GetSummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricEnumValues() []SummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricEnum {
	values := make([]SummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetricEnum, 0)
	for _, v := range mappingSummarizeHostInsightResourceStatisticsAggregationCollectionResourceMetric {
		values = append(values, v)
	}
	return values
}
