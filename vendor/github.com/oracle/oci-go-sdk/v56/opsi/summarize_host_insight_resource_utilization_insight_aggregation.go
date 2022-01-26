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

// SummarizeHostInsightResourceUtilizationInsightAggregation Insights response containing current/projected groups for CPU or memory.
type SummarizeHostInsightResourceUtilizationInsightAggregation struct {

	// The start timestamp that was passed into the request.
	TimeIntervalStart *common.SDKTime `mandatory:"true" json:"timeIntervalStart"`

	// The end timestamp that was passed into the request.
	TimeIntervalEnd *common.SDKTime `mandatory:"true" json:"timeIntervalEnd"`

	// Defines the type of resource metric (CPU, Physical Memory, Logical Memory)
	ResourceMetric SummarizeHostInsightResourceUtilizationInsightAggregationResourceMetricEnum `mandatory:"true" json:"resourceMetric"`

	ProjectedUtilization *ResourceInsightProjectedUtilization `mandatory:"true" json:"projectedUtilization"`

	CurrentUtilization *ResourceInsightCurrentUtilization `mandatory:"true" json:"currentUtilization"`
}

func (m SummarizeHostInsightResourceUtilizationInsightAggregation) String() string {
	return common.PointerString(m)
}

// SummarizeHostInsightResourceUtilizationInsightAggregationResourceMetricEnum Enum with underlying type: string
type SummarizeHostInsightResourceUtilizationInsightAggregationResourceMetricEnum string

// Set of constants representing the allowable values for SummarizeHostInsightResourceUtilizationInsightAggregationResourceMetricEnum
const (
	SummarizeHostInsightResourceUtilizationInsightAggregationResourceMetricCpu           SummarizeHostInsightResourceUtilizationInsightAggregationResourceMetricEnum = "CPU"
	SummarizeHostInsightResourceUtilizationInsightAggregationResourceMetricMemory        SummarizeHostInsightResourceUtilizationInsightAggregationResourceMetricEnum = "MEMORY"
	SummarizeHostInsightResourceUtilizationInsightAggregationResourceMetricLogicalMemory SummarizeHostInsightResourceUtilizationInsightAggregationResourceMetricEnum = "LOGICAL_MEMORY"
)

var mappingSummarizeHostInsightResourceUtilizationInsightAggregationResourceMetric = map[string]SummarizeHostInsightResourceUtilizationInsightAggregationResourceMetricEnum{
	"CPU":            SummarizeHostInsightResourceUtilizationInsightAggregationResourceMetricCpu,
	"MEMORY":         SummarizeHostInsightResourceUtilizationInsightAggregationResourceMetricMemory,
	"LOGICAL_MEMORY": SummarizeHostInsightResourceUtilizationInsightAggregationResourceMetricLogicalMemory,
}

// GetSummarizeHostInsightResourceUtilizationInsightAggregationResourceMetricEnumValues Enumerates the set of values for SummarizeHostInsightResourceUtilizationInsightAggregationResourceMetricEnum
func GetSummarizeHostInsightResourceUtilizationInsightAggregationResourceMetricEnumValues() []SummarizeHostInsightResourceUtilizationInsightAggregationResourceMetricEnum {
	values := make([]SummarizeHostInsightResourceUtilizationInsightAggregationResourceMetricEnum, 0)
	for _, v := range mappingSummarizeHostInsightResourceUtilizationInsightAggregationResourceMetric {
		values = append(values, v)
	}
	return values
}
