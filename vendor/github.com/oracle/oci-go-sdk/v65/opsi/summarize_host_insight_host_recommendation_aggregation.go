// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SummarizeHostInsightHostRecommendationAggregation Returns list of hosts with resource statistics like usage, capacity, utilization, usage change percent and load.
type SummarizeHostInsightHostRecommendationAggregation struct {

	// Defines the type of resource metric (CPU, Physical Memory, Logical Memory)
	ResourceMetric SummarizeHostInsightHostRecommendationAggregationResourceMetricEnum `mandatory:"true" json:"resourceMetric"`

	// Displays usage unit ( CORES, GB , PERCENT, MBPS)
	UsageUnit SummarizeHostInsightHostRecommendationAggregationUsageUnitEnum `mandatory:"true" json:"usageUnit"`

	// Time duration in milliseconds between data points (one hour or one day).
	ItemDurationInMs *int64 `mandatory:"true" json:"itemDurationInMs"`

	// The start timestamp that was passed into the request.
	TimeIntervalStart *common.SDKTime `mandatory:"true" json:"timeIntervalStart"`

	// The end timestamp that was passed into the request.
	TimeIntervalEnd *common.SDKTime `mandatory:"true" json:"timeIntervalEnd"`

	Details HostInsightHostRecommendations `mandatory:"false" json:"details"`
}

func (m SummarizeHostInsightHostRecommendationAggregation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SummarizeHostInsightHostRecommendationAggregation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSummarizeHostInsightHostRecommendationAggregationResourceMetricEnum(string(m.ResourceMetric)); !ok && m.ResourceMetric != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResourceMetric: %s. Supported values are: %s.", m.ResourceMetric, strings.Join(GetSummarizeHostInsightHostRecommendationAggregationResourceMetricEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeHostInsightHostRecommendationAggregationUsageUnitEnum(string(m.UsageUnit)); !ok && m.UsageUnit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UsageUnit: %s. Supported values are: %s.", m.UsageUnit, strings.Join(GetSummarizeHostInsightHostRecommendationAggregationUsageUnitEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *SummarizeHostInsightHostRecommendationAggregation) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Details           hostinsighthostrecommendations                                      `json:"details"`
		ResourceMetric    SummarizeHostInsightHostRecommendationAggregationResourceMetricEnum `json:"resourceMetric"`
		UsageUnit         SummarizeHostInsightHostRecommendationAggregationUsageUnitEnum      `json:"usageUnit"`
		ItemDurationInMs  *int64                                                              `json:"itemDurationInMs"`
		TimeIntervalStart *common.SDKTime                                                     `json:"timeIntervalStart"`
		TimeIntervalEnd   *common.SDKTime                                                     `json:"timeIntervalEnd"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.Details.UnmarshalPolymorphicJSON(model.Details.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Details = nn.(HostInsightHostRecommendations)
	} else {
		m.Details = nil
	}

	m.ResourceMetric = model.ResourceMetric

	m.UsageUnit = model.UsageUnit

	m.ItemDurationInMs = model.ItemDurationInMs

	m.TimeIntervalStart = model.TimeIntervalStart

	m.TimeIntervalEnd = model.TimeIntervalEnd

	return
}

// SummarizeHostInsightHostRecommendationAggregationResourceMetricEnum Enum with underlying type: string
type SummarizeHostInsightHostRecommendationAggregationResourceMetricEnum string

// Set of constants representing the allowable values for SummarizeHostInsightHostRecommendationAggregationResourceMetricEnum
const (
	SummarizeHostInsightHostRecommendationAggregationResourceMetricCpu           SummarizeHostInsightHostRecommendationAggregationResourceMetricEnum = "CPU"
	SummarizeHostInsightHostRecommendationAggregationResourceMetricMemory        SummarizeHostInsightHostRecommendationAggregationResourceMetricEnum = "MEMORY"
	SummarizeHostInsightHostRecommendationAggregationResourceMetricLogicalMemory SummarizeHostInsightHostRecommendationAggregationResourceMetricEnum = "LOGICAL_MEMORY"
	SummarizeHostInsightHostRecommendationAggregationResourceMetricStorage       SummarizeHostInsightHostRecommendationAggregationResourceMetricEnum = "STORAGE"
	SummarizeHostInsightHostRecommendationAggregationResourceMetricNetwork       SummarizeHostInsightHostRecommendationAggregationResourceMetricEnum = "NETWORK"
)

var mappingSummarizeHostInsightHostRecommendationAggregationResourceMetricEnum = map[string]SummarizeHostInsightHostRecommendationAggregationResourceMetricEnum{
	"CPU":            SummarizeHostInsightHostRecommendationAggregationResourceMetricCpu,
	"MEMORY":         SummarizeHostInsightHostRecommendationAggregationResourceMetricMemory,
	"LOGICAL_MEMORY": SummarizeHostInsightHostRecommendationAggregationResourceMetricLogicalMemory,
	"STORAGE":        SummarizeHostInsightHostRecommendationAggregationResourceMetricStorage,
	"NETWORK":        SummarizeHostInsightHostRecommendationAggregationResourceMetricNetwork,
}

var mappingSummarizeHostInsightHostRecommendationAggregationResourceMetricEnumLowerCase = map[string]SummarizeHostInsightHostRecommendationAggregationResourceMetricEnum{
	"cpu":            SummarizeHostInsightHostRecommendationAggregationResourceMetricCpu,
	"memory":         SummarizeHostInsightHostRecommendationAggregationResourceMetricMemory,
	"logical_memory": SummarizeHostInsightHostRecommendationAggregationResourceMetricLogicalMemory,
	"storage":        SummarizeHostInsightHostRecommendationAggregationResourceMetricStorage,
	"network":        SummarizeHostInsightHostRecommendationAggregationResourceMetricNetwork,
}

// GetSummarizeHostInsightHostRecommendationAggregationResourceMetricEnumValues Enumerates the set of values for SummarizeHostInsightHostRecommendationAggregationResourceMetricEnum
func GetSummarizeHostInsightHostRecommendationAggregationResourceMetricEnumValues() []SummarizeHostInsightHostRecommendationAggregationResourceMetricEnum {
	values := make([]SummarizeHostInsightHostRecommendationAggregationResourceMetricEnum, 0)
	for _, v := range mappingSummarizeHostInsightHostRecommendationAggregationResourceMetricEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeHostInsightHostRecommendationAggregationResourceMetricEnumStringValues Enumerates the set of values in String for SummarizeHostInsightHostRecommendationAggregationResourceMetricEnum
func GetSummarizeHostInsightHostRecommendationAggregationResourceMetricEnumStringValues() []string {
	return []string{
		"CPU",
		"MEMORY",
		"LOGICAL_MEMORY",
		"STORAGE",
		"NETWORK",
	}
}

// GetMappingSummarizeHostInsightHostRecommendationAggregationResourceMetricEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeHostInsightHostRecommendationAggregationResourceMetricEnum(val string) (SummarizeHostInsightHostRecommendationAggregationResourceMetricEnum, bool) {
	enum, ok := mappingSummarizeHostInsightHostRecommendationAggregationResourceMetricEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeHostInsightHostRecommendationAggregationUsageUnitEnum Enum with underlying type: string
type SummarizeHostInsightHostRecommendationAggregationUsageUnitEnum string

// Set of constants representing the allowable values for SummarizeHostInsightHostRecommendationAggregationUsageUnitEnum
const (
	SummarizeHostInsightHostRecommendationAggregationUsageUnitCores   SummarizeHostInsightHostRecommendationAggregationUsageUnitEnum = "CORES"
	SummarizeHostInsightHostRecommendationAggregationUsageUnitGb      SummarizeHostInsightHostRecommendationAggregationUsageUnitEnum = "GB"
	SummarizeHostInsightHostRecommendationAggregationUsageUnitMbps    SummarizeHostInsightHostRecommendationAggregationUsageUnitEnum = "MBPS"
	SummarizeHostInsightHostRecommendationAggregationUsageUnitIops    SummarizeHostInsightHostRecommendationAggregationUsageUnitEnum = "IOPS"
	SummarizeHostInsightHostRecommendationAggregationUsageUnitPercent SummarizeHostInsightHostRecommendationAggregationUsageUnitEnum = "PERCENT"
)

var mappingSummarizeHostInsightHostRecommendationAggregationUsageUnitEnum = map[string]SummarizeHostInsightHostRecommendationAggregationUsageUnitEnum{
	"CORES":   SummarizeHostInsightHostRecommendationAggregationUsageUnitCores,
	"GB":      SummarizeHostInsightHostRecommendationAggregationUsageUnitGb,
	"MBPS":    SummarizeHostInsightHostRecommendationAggregationUsageUnitMbps,
	"IOPS":    SummarizeHostInsightHostRecommendationAggregationUsageUnitIops,
	"PERCENT": SummarizeHostInsightHostRecommendationAggregationUsageUnitPercent,
}

var mappingSummarizeHostInsightHostRecommendationAggregationUsageUnitEnumLowerCase = map[string]SummarizeHostInsightHostRecommendationAggregationUsageUnitEnum{
	"cores":   SummarizeHostInsightHostRecommendationAggregationUsageUnitCores,
	"gb":      SummarizeHostInsightHostRecommendationAggregationUsageUnitGb,
	"mbps":    SummarizeHostInsightHostRecommendationAggregationUsageUnitMbps,
	"iops":    SummarizeHostInsightHostRecommendationAggregationUsageUnitIops,
	"percent": SummarizeHostInsightHostRecommendationAggregationUsageUnitPercent,
}

// GetSummarizeHostInsightHostRecommendationAggregationUsageUnitEnumValues Enumerates the set of values for SummarizeHostInsightHostRecommendationAggregationUsageUnitEnum
func GetSummarizeHostInsightHostRecommendationAggregationUsageUnitEnumValues() []SummarizeHostInsightHostRecommendationAggregationUsageUnitEnum {
	values := make([]SummarizeHostInsightHostRecommendationAggregationUsageUnitEnum, 0)
	for _, v := range mappingSummarizeHostInsightHostRecommendationAggregationUsageUnitEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeHostInsightHostRecommendationAggregationUsageUnitEnumStringValues Enumerates the set of values in String for SummarizeHostInsightHostRecommendationAggregationUsageUnitEnum
func GetSummarizeHostInsightHostRecommendationAggregationUsageUnitEnumStringValues() []string {
	return []string{
		"CORES",
		"GB",
		"MBPS",
		"IOPS",
		"PERCENT",
	}
}

// GetMappingSummarizeHostInsightHostRecommendationAggregationUsageUnitEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeHostInsightHostRecommendationAggregationUsageUnitEnum(val string) (SummarizeHostInsightHostRecommendationAggregationUsageUnitEnum, bool) {
	enum, ok := mappingSummarizeHostInsightHostRecommendationAggregationUsageUnitEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
