// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v37/common"
)

// LogAnalyticsMetric LogAnalyticsMetric
type LogAnalyticsMetric struct {

	// aggregation field
	AggregationField *string `mandatory:"false" json:"aggregationField"`

	// bucket metadata
	BucketMetadata *string `mandatory:"false" json:"bucketMetadata"`

	// clock period
	ClockPeriod *string `mandatory:"false" json:"clockPeriod"`

	// description
	Description *string `mandatory:"false" json:"description"`

	// edit version
	EditVersion *int64 `mandatory:"false" json:"editVersion"`

	// field name
	FieldName *string `mandatory:"false" json:"fieldName"`

	// field value array
	FieldValues []string `mandatory:"false" json:"fieldValues"`

	// grouping fields
	GroupingField *string `mandatory:"false" json:"groupingField"`

	// is enabled flag
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// is system flag
	IsSystem *bool `mandatory:"false" json:"isSystem"`

	// metric display name
	DisplayName *string `mandatory:"false" json:"displayName"`

	// metric Id
	MetricReference *int64 `mandatory:"false" json:"metricReference"`

	// name
	Name *string `mandatory:"false" json:"name"`

	// metric type
	MetricType LogAnalyticsMetricMetricTypeEnum `mandatory:"false" json:"metricType,omitempty"`

	// is metric source map enabled flag
	IsMetricSourceEnabled *bool `mandatory:"false" json:"isMetricSourceEnabled"`

	// operator
	Operator LogAnalyticsMetricOperatorEnum `mandatory:"false" json:"operator,omitempty"`

	// sources
	Sources []LogAnalyticsSource `mandatory:"false" json:"sources"`

	// entity type
	EntityType *string `mandatory:"false" json:"entityType"`

	// last updated date
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// unit type
	UnitType *string `mandatory:"false" json:"unitType"`

	// user customized
	IsUserCustomized *bool `mandatory:"false" json:"isUserCustomized"`
}

func (m LogAnalyticsMetric) String() string {
	return common.PointerString(m)
}

// LogAnalyticsMetricMetricTypeEnum Enum with underlying type: string
type LogAnalyticsMetricMetricTypeEnum string

// Set of constants representing the allowable values for LogAnalyticsMetricMetricTypeEnum
const (
	LogAnalyticsMetricMetricTypeCount               LogAnalyticsMetricMetricTypeEnum = "COUNT"
	LogAnalyticsMetricMetricTypeSum                 LogAnalyticsMetricMetricTypeEnum = "SUM"
	LogAnalyticsMetricMetricTypeAverage             LogAnalyticsMetricMetricTypeEnum = "AVERAGE"
	LogAnalyticsMetricMetricTypeCountDistribution   LogAnalyticsMetricMetricTypeEnum = "COUNT_DISTRIBUTION"
	LogAnalyticsMetricMetricTypeSumDistribution     LogAnalyticsMetricMetricTypeEnum = "SUM_DISTRIBUTION"
	LogAnalyticsMetricMetricTypeAverageDistribution LogAnalyticsMetricMetricTypeEnum = "AVERAGE_DISTRIBUTION"
)

var mappingLogAnalyticsMetricMetricType = map[string]LogAnalyticsMetricMetricTypeEnum{
	"COUNT":                LogAnalyticsMetricMetricTypeCount,
	"SUM":                  LogAnalyticsMetricMetricTypeSum,
	"AVERAGE":              LogAnalyticsMetricMetricTypeAverage,
	"COUNT_DISTRIBUTION":   LogAnalyticsMetricMetricTypeCountDistribution,
	"SUM_DISTRIBUTION":     LogAnalyticsMetricMetricTypeSumDistribution,
	"AVERAGE_DISTRIBUTION": LogAnalyticsMetricMetricTypeAverageDistribution,
}

// GetLogAnalyticsMetricMetricTypeEnumValues Enumerates the set of values for LogAnalyticsMetricMetricTypeEnum
func GetLogAnalyticsMetricMetricTypeEnumValues() []LogAnalyticsMetricMetricTypeEnum {
	values := make([]LogAnalyticsMetricMetricTypeEnum, 0)
	for _, v := range mappingLogAnalyticsMetricMetricType {
		values = append(values, v)
	}
	return values
}

// LogAnalyticsMetricOperatorEnum Enum with underlying type: string
type LogAnalyticsMetricOperatorEnum string

// Set of constants representing the allowable values for LogAnalyticsMetricOperatorEnum
const (
	LogAnalyticsMetricOperatorContainsIgnoreCase LogAnalyticsMetricOperatorEnum = "CONTAINS_IGNORE_CASE"
	LogAnalyticsMetricOperatorInIgnoreCase       LogAnalyticsMetricOperatorEnum = "IN_IGNORE_CASE"
	LogAnalyticsMetricOperatorEqualIgnoreCase    LogAnalyticsMetricOperatorEnum = "EQUAL_IGNORE_CASE"
	LogAnalyticsMetricOperatorNotNull            LogAnalyticsMetricOperatorEnum = "NOT_NULL"
)

var mappingLogAnalyticsMetricOperator = map[string]LogAnalyticsMetricOperatorEnum{
	"CONTAINS_IGNORE_CASE": LogAnalyticsMetricOperatorContainsIgnoreCase,
	"IN_IGNORE_CASE":       LogAnalyticsMetricOperatorInIgnoreCase,
	"EQUAL_IGNORE_CASE":    LogAnalyticsMetricOperatorEqualIgnoreCase,
	"NOT_NULL":             LogAnalyticsMetricOperatorNotNull,
}

// GetLogAnalyticsMetricOperatorEnumValues Enumerates the set of values for LogAnalyticsMetricOperatorEnum
func GetLogAnalyticsMetricOperatorEnumValues() []LogAnalyticsMetricOperatorEnum {
	values := make([]LogAnalyticsMetricOperatorEnum, 0)
	for _, v := range mappingLogAnalyticsMetricOperator {
		values = append(values, v)
	}
	return values
}
