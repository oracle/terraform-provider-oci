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

// SummarizeDatabaseInsightResourceUtilizationInsightAggregation Insights response containing current/projected groups for storage or CPU.
type SummarizeDatabaseInsightResourceUtilizationInsightAggregation struct {

	// The start timestamp that was passed into the request.
	TimeIntervalStart *common.SDKTime `mandatory:"true" json:"timeIntervalStart"`

	// The end timestamp that was passed into the request.
	TimeIntervalEnd *common.SDKTime `mandatory:"true" json:"timeIntervalEnd"`

	// Defines the type of resource metric (example: CPU, STORAGE)
	ResourceMetric SummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetricEnum `mandatory:"true" json:"resourceMetric"`

	ProjectedUtilization *ResourceInsightProjectedUtilization `mandatory:"true" json:"projectedUtilization"`

	CurrentUtilization *ResourceInsightCurrentUtilization `mandatory:"true" json:"currentUtilization"`
}

func (m SummarizeDatabaseInsightResourceUtilizationInsightAggregation) String() string {
	return common.PointerString(m)
}

// SummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetricEnum Enum with underlying type: string
type SummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetricEnum string

// Set of constants representing the allowable values for SummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetricEnum
const (
	SummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetricCpu       SummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetricEnum = "CPU"
	SummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetricStorage   SummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetricEnum = "STORAGE"
	SummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetricIo        SummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetricEnum = "IO"
	SummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetricMemory    SummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetricEnum = "MEMORY"
	SummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetricMemoryPga SummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetricEnum = "MEMORY_PGA"
	SummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetricMemorySga SummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetricEnum = "MEMORY_SGA"
)

var mappingSummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetric = map[string]SummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetricEnum{
	"CPU":        SummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetricCpu,
	"STORAGE":    SummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetricStorage,
	"IO":         SummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetricIo,
	"MEMORY":     SummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetricMemory,
	"MEMORY_PGA": SummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetricMemoryPga,
	"MEMORY_SGA": SummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetricMemorySga,
}

// GetSummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetricEnumValues Enumerates the set of values for SummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetricEnum
func GetSummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetricEnumValues() []SummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetricEnum {
	values := make([]SummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetricEnum, 0)
	for _, v := range mappingSummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetric {
		values = append(values, v)
	}
	return values
}
