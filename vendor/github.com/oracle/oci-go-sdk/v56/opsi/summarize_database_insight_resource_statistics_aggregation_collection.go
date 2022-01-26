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

// SummarizeDatabaseInsightResourceStatisticsAggregationCollection Returns list of the Databases with resource statistics like usage,capacity,utilization and usage change percent.
type SummarizeDatabaseInsightResourceStatisticsAggregationCollection struct {

	// The start timestamp that was passed into the request.
	TimeIntervalStart *common.SDKTime `mandatory:"true" json:"timeIntervalStart"`

	// The end timestamp that was passed into the request.
	TimeIntervalEnd *common.SDKTime `mandatory:"true" json:"timeIntervalEnd"`

	// Defines the type of resource metric (example: CPU, STORAGE)
	ResourceMetric SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricEnum `mandatory:"true" json:"resourceMetric"`

	// Displays usage unit ( CORES, GB)
	UsageUnit UsageUnitEnum `mandatory:"true" json:"usageUnit"`

	// Collection of Resource Statistics items
	Items []ResourceStatisticsAggregation `mandatory:"true" json:"items"`
}

func (m SummarizeDatabaseInsightResourceStatisticsAggregationCollection) String() string {
	return common.PointerString(m)
}

// SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricEnum Enum with underlying type: string
type SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricEnum string

// Set of constants representing the allowable values for SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricEnum
const (
	SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricCpu       SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricEnum = "CPU"
	SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricStorage   SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricEnum = "STORAGE"
	SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricIo        SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricEnum = "IO"
	SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricMemory    SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricEnum = "MEMORY"
	SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricMemoryPga SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricEnum = "MEMORY_PGA"
	SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricMemorySga SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricEnum = "MEMORY_SGA"
)

var mappingSummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetric = map[string]SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricEnum{
	"CPU":        SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricCpu,
	"STORAGE":    SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricStorage,
	"IO":         SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricIo,
	"MEMORY":     SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricMemory,
	"MEMORY_PGA": SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricMemoryPga,
	"MEMORY_SGA": SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricMemorySga,
}

// GetSummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricEnumValues Enumerates the set of values for SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricEnum
func GetSummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricEnumValues() []SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricEnum {
	values := make([]SummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetricEnum, 0)
	for _, v := range mappingSummarizeDatabaseInsightResourceStatisticsAggregationCollectionResourceMetric {
		values = append(values, v)
	}
	return values
}
