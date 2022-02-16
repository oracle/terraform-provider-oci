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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SummarizeDatabaseInsightResourceUtilizationInsightAggregation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetricEnum(string(m.ResourceMetric)); !ok && m.ResourceMetric != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResourceMetric: %s. Supported values are: %s.", m.ResourceMetric, strings.Join(GetSummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetricEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingSummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetricEnum = map[string]SummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetricEnum{
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
	for _, v := range mappingSummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetricEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetricEnumStringValues Enumerates the set of values in String for SummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetricEnum
func GetSummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetricEnumStringValues() []string {
	return []string{
		"CPU",
		"STORAGE",
		"IO",
		"MEMORY",
		"MEMORY_PGA",
		"MEMORY_SGA",
	}
}

// GetMappingSummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetricEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetricEnum(val string) (SummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetricEnum, bool) {
	mappingSummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetricEnumIgnoreCase := make(map[string]SummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetricEnum)
	for k, v := range mappingSummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetricEnum {
		mappingSummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetricEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingSummarizeDatabaseInsightResourceUtilizationInsightAggregationResourceMetricEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
