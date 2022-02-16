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
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// SummarizeExadataInsightResourceStatisticsAggregationCollection Returns list of the resources with resource statistics like usage,capacity,utilization and usage change percent.
type SummarizeExadataInsightResourceStatisticsAggregationCollection struct {

	// The start timestamp that was passed into the request.
	TimeIntervalStart *common.SDKTime `mandatory:"true" json:"timeIntervalStart"`

	// The end timestamp that was passed into the request.
	TimeIntervalEnd *common.SDKTime `mandatory:"true" json:"timeIntervalEnd"`

	// Collection of Resource Statistics items
	Items []ExadataInsightResourceStatisticsAggregation `mandatory:"true" json:"items"`

	// Displays usage unit ( CORES, GB)
	UsageUnit UsageUnitEnum `mandatory:"true" json:"usageUnit"`

	// Defines the type of exadata resource metric (example: CPU, STORAGE)
	ExadataResourceMetric SummarizeExadataInsightResourceStatisticsAggregationCollectionExadataResourceMetricEnum `mandatory:"true" json:"exadataResourceMetric"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata insight.
	ExadataInsightId *string `mandatory:"true" json:"exadataInsightId"`
}

func (m SummarizeExadataInsightResourceStatisticsAggregationCollection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SummarizeExadataInsightResourceStatisticsAggregationCollection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingUsageUnitEnum(string(m.UsageUnit)); !ok && m.UsageUnit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UsageUnit: %s. Supported values are: %s.", m.UsageUnit, strings.Join(GetUsageUnitEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeExadataInsightResourceStatisticsAggregationCollectionExadataResourceMetricEnum(string(m.ExadataResourceMetric)); !ok && m.ExadataResourceMetric != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ExadataResourceMetric: %s. Supported values are: %s.", m.ExadataResourceMetric, strings.Join(GetSummarizeExadataInsightResourceStatisticsAggregationCollectionExadataResourceMetricEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *SummarizeExadataInsightResourceStatisticsAggregationCollection) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		TimeIntervalStart     *common.SDKTime                                                                         `json:"timeIntervalStart"`
		TimeIntervalEnd       *common.SDKTime                                                                         `json:"timeIntervalEnd"`
		Items                 []exadatainsightresourcestatisticsaggregation                                           `json:"items"`
		UsageUnit             UsageUnitEnum                                                                           `json:"usageUnit"`
		ExadataResourceMetric SummarizeExadataInsightResourceStatisticsAggregationCollectionExadataResourceMetricEnum `json:"exadataResourceMetric"`
		ExadataInsightId      *string                                                                                 `json:"exadataInsightId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.TimeIntervalStart = model.TimeIntervalStart

	m.TimeIntervalEnd = model.TimeIntervalEnd

	m.Items = make([]ExadataInsightResourceStatisticsAggregation, len(model.Items))
	for i, n := range model.Items {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Items[i] = nn.(ExadataInsightResourceStatisticsAggregation)
		} else {
			m.Items[i] = nil
		}
	}

	m.UsageUnit = model.UsageUnit

	m.ExadataResourceMetric = model.ExadataResourceMetric

	m.ExadataInsightId = model.ExadataInsightId

	return
}

// SummarizeExadataInsightResourceStatisticsAggregationCollectionExadataResourceMetricEnum Enum with underlying type: string
type SummarizeExadataInsightResourceStatisticsAggregationCollectionExadataResourceMetricEnum string

// Set of constants representing the allowable values for SummarizeExadataInsightResourceStatisticsAggregationCollectionExadataResourceMetricEnum
const (
	SummarizeExadataInsightResourceStatisticsAggregationCollectionExadataResourceMetricCpu        SummarizeExadataInsightResourceStatisticsAggregationCollectionExadataResourceMetricEnum = "CPU"
	SummarizeExadataInsightResourceStatisticsAggregationCollectionExadataResourceMetricStorage    SummarizeExadataInsightResourceStatisticsAggregationCollectionExadataResourceMetricEnum = "STORAGE"
	SummarizeExadataInsightResourceStatisticsAggregationCollectionExadataResourceMetricIo         SummarizeExadataInsightResourceStatisticsAggregationCollectionExadataResourceMetricEnum = "IO"
	SummarizeExadataInsightResourceStatisticsAggregationCollectionExadataResourceMetricMemory     SummarizeExadataInsightResourceStatisticsAggregationCollectionExadataResourceMetricEnum = "MEMORY"
	SummarizeExadataInsightResourceStatisticsAggregationCollectionExadataResourceMetricIops       SummarizeExadataInsightResourceStatisticsAggregationCollectionExadataResourceMetricEnum = "IOPS"
	SummarizeExadataInsightResourceStatisticsAggregationCollectionExadataResourceMetricThroughput SummarizeExadataInsightResourceStatisticsAggregationCollectionExadataResourceMetricEnum = "THROUGHPUT"
)

var mappingSummarizeExadataInsightResourceStatisticsAggregationCollectionExadataResourceMetricEnum = map[string]SummarizeExadataInsightResourceStatisticsAggregationCollectionExadataResourceMetricEnum{
	"CPU":        SummarizeExadataInsightResourceStatisticsAggregationCollectionExadataResourceMetricCpu,
	"STORAGE":    SummarizeExadataInsightResourceStatisticsAggregationCollectionExadataResourceMetricStorage,
	"IO":         SummarizeExadataInsightResourceStatisticsAggregationCollectionExadataResourceMetricIo,
	"MEMORY":     SummarizeExadataInsightResourceStatisticsAggregationCollectionExadataResourceMetricMemory,
	"IOPS":       SummarizeExadataInsightResourceStatisticsAggregationCollectionExadataResourceMetricIops,
	"THROUGHPUT": SummarizeExadataInsightResourceStatisticsAggregationCollectionExadataResourceMetricThroughput,
}

// GetSummarizeExadataInsightResourceStatisticsAggregationCollectionExadataResourceMetricEnumValues Enumerates the set of values for SummarizeExadataInsightResourceStatisticsAggregationCollectionExadataResourceMetricEnum
func GetSummarizeExadataInsightResourceStatisticsAggregationCollectionExadataResourceMetricEnumValues() []SummarizeExadataInsightResourceStatisticsAggregationCollectionExadataResourceMetricEnum {
	values := make([]SummarizeExadataInsightResourceStatisticsAggregationCollectionExadataResourceMetricEnum, 0)
	for _, v := range mappingSummarizeExadataInsightResourceStatisticsAggregationCollectionExadataResourceMetricEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeExadataInsightResourceStatisticsAggregationCollectionExadataResourceMetricEnumStringValues Enumerates the set of values in String for SummarizeExadataInsightResourceStatisticsAggregationCollectionExadataResourceMetricEnum
func GetSummarizeExadataInsightResourceStatisticsAggregationCollectionExadataResourceMetricEnumStringValues() []string {
	return []string{
		"CPU",
		"STORAGE",
		"IO",
		"MEMORY",
		"IOPS",
		"THROUGHPUT",
	}
}

// GetMappingSummarizeExadataInsightResourceStatisticsAggregationCollectionExadataResourceMetricEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeExadataInsightResourceStatisticsAggregationCollectionExadataResourceMetricEnum(val string) (SummarizeExadataInsightResourceStatisticsAggregationCollectionExadataResourceMetricEnum, bool) {
	mappingSummarizeExadataInsightResourceStatisticsAggregationCollectionExadataResourceMetricEnumIgnoreCase := make(map[string]SummarizeExadataInsightResourceStatisticsAggregationCollectionExadataResourceMetricEnum)
	for k, v := range mappingSummarizeExadataInsightResourceStatisticsAggregationCollectionExadataResourceMetricEnum {
		mappingSummarizeExadataInsightResourceStatisticsAggregationCollectionExadataResourceMetricEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingSummarizeExadataInsightResourceStatisticsAggregationCollectionExadataResourceMetricEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
