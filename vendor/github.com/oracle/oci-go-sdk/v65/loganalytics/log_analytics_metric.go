// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LogAnalyticsMetric LogAnalyticsMetric
type LogAnalyticsMetric struct {

	// The aggregation field.
	AggregationField *string `mandatory:"false" json:"aggregationField"`

	// The bucket metadata.
	BucketMetadata *string `mandatory:"false" json:"bucketMetadata"`

	// The clock period.
	ClockPeriod *string `mandatory:"false" json:"clockPeriod"`

	// The metric description.
	Description *string `mandatory:"false" json:"description"`

	// The metric edit version.
	EditVersion *int64 `mandatory:"false" json:"editVersion"`

	// The field name.
	FieldName *string `mandatory:"false" json:"fieldName"`

	// The field values.
	FieldValues []string `mandatory:"false" json:"fieldValues"`

	// The grouping fields.
	GroupingField *string `mandatory:"false" json:"groupingField"`

	// A flag inidcating whether or not the metric is enabled.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// The system flag.  A value of false denotes a custom, or user
	// defined object.  A value of true denotes a built in object.
	IsSystem *bool `mandatory:"false" json:"isSystem"`

	// The metric display name.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The metric unique identifier.
	MetricReference *int64 `mandatory:"false" json:"metricReference"`

	// The metric name.
	Name *string `mandatory:"false" json:"name"`

	// The metric type, specifying the type of aggreation to perform.  Default value
	// is COUNT.
	MetricType LogAnalyticsMetricMetricTypeEnum `mandatory:"false" json:"metricType,omitempty"`

	// A flag specifying whether or not the metric source is enabled.
	IsMetricSourceEnabled *bool `mandatory:"false" json:"isMetricSourceEnabled"`

	// The metric operator.
	Operator LogAnalyticsMetricOperatorEnum `mandatory:"false" json:"operator,omitempty"`

	// The metric sources.
	Sources []LogAnalyticsSource `mandatory:"false" json:"sources"`

	// The entity type.
	EntityType *string `mandatory:"false" json:"entityType"`

	// The last updated date.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The unit type.
	UnitType *string `mandatory:"false" json:"unitType"`

	// A flag specifying whether or not this is a custom (user defined) metric.
	IsUserCustomized *bool `mandatory:"false" json:"isUserCustomized"`
}

func (m LogAnalyticsMetric) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LogAnalyticsMetric) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingLogAnalyticsMetricMetricTypeEnum(string(m.MetricType)); !ok && m.MetricType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MetricType: %s. Supported values are: %s.", m.MetricType, strings.Join(GetLogAnalyticsMetricMetricTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLogAnalyticsMetricOperatorEnum(string(m.Operator)); !ok && m.Operator != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Operator: %s. Supported values are: %s.", m.Operator, strings.Join(GetLogAnalyticsMetricOperatorEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingLogAnalyticsMetricMetricTypeEnum = map[string]LogAnalyticsMetricMetricTypeEnum{
	"COUNT":                LogAnalyticsMetricMetricTypeCount,
	"SUM":                  LogAnalyticsMetricMetricTypeSum,
	"AVERAGE":              LogAnalyticsMetricMetricTypeAverage,
	"COUNT_DISTRIBUTION":   LogAnalyticsMetricMetricTypeCountDistribution,
	"SUM_DISTRIBUTION":     LogAnalyticsMetricMetricTypeSumDistribution,
	"AVERAGE_DISTRIBUTION": LogAnalyticsMetricMetricTypeAverageDistribution,
}

var mappingLogAnalyticsMetricMetricTypeEnumLowerCase = map[string]LogAnalyticsMetricMetricTypeEnum{
	"count":                LogAnalyticsMetricMetricTypeCount,
	"sum":                  LogAnalyticsMetricMetricTypeSum,
	"average":              LogAnalyticsMetricMetricTypeAverage,
	"count_distribution":   LogAnalyticsMetricMetricTypeCountDistribution,
	"sum_distribution":     LogAnalyticsMetricMetricTypeSumDistribution,
	"average_distribution": LogAnalyticsMetricMetricTypeAverageDistribution,
}

// GetLogAnalyticsMetricMetricTypeEnumValues Enumerates the set of values for LogAnalyticsMetricMetricTypeEnum
func GetLogAnalyticsMetricMetricTypeEnumValues() []LogAnalyticsMetricMetricTypeEnum {
	values := make([]LogAnalyticsMetricMetricTypeEnum, 0)
	for _, v := range mappingLogAnalyticsMetricMetricTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetLogAnalyticsMetricMetricTypeEnumStringValues Enumerates the set of values in String for LogAnalyticsMetricMetricTypeEnum
func GetLogAnalyticsMetricMetricTypeEnumStringValues() []string {
	return []string{
		"COUNT",
		"SUM",
		"AVERAGE",
		"COUNT_DISTRIBUTION",
		"SUM_DISTRIBUTION",
		"AVERAGE_DISTRIBUTION",
	}
}

// GetMappingLogAnalyticsMetricMetricTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLogAnalyticsMetricMetricTypeEnum(val string) (LogAnalyticsMetricMetricTypeEnum, bool) {
	enum, ok := mappingLogAnalyticsMetricMetricTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
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

var mappingLogAnalyticsMetricOperatorEnum = map[string]LogAnalyticsMetricOperatorEnum{
	"CONTAINS_IGNORE_CASE": LogAnalyticsMetricOperatorContainsIgnoreCase,
	"IN_IGNORE_CASE":       LogAnalyticsMetricOperatorInIgnoreCase,
	"EQUAL_IGNORE_CASE":    LogAnalyticsMetricOperatorEqualIgnoreCase,
	"NOT_NULL":             LogAnalyticsMetricOperatorNotNull,
}

var mappingLogAnalyticsMetricOperatorEnumLowerCase = map[string]LogAnalyticsMetricOperatorEnum{
	"contains_ignore_case": LogAnalyticsMetricOperatorContainsIgnoreCase,
	"in_ignore_case":       LogAnalyticsMetricOperatorInIgnoreCase,
	"equal_ignore_case":    LogAnalyticsMetricOperatorEqualIgnoreCase,
	"not_null":             LogAnalyticsMetricOperatorNotNull,
}

// GetLogAnalyticsMetricOperatorEnumValues Enumerates the set of values for LogAnalyticsMetricOperatorEnum
func GetLogAnalyticsMetricOperatorEnumValues() []LogAnalyticsMetricOperatorEnum {
	values := make([]LogAnalyticsMetricOperatorEnum, 0)
	for _, v := range mappingLogAnalyticsMetricOperatorEnum {
		values = append(values, v)
	}
	return values
}

// GetLogAnalyticsMetricOperatorEnumStringValues Enumerates the set of values in String for LogAnalyticsMetricOperatorEnum
func GetLogAnalyticsMetricOperatorEnumStringValues() []string {
	return []string{
		"CONTAINS_IGNORE_CASE",
		"IN_IGNORE_CASE",
		"EQUAL_IGNORE_CASE",
		"NOT_NULL",
	}
}

// GetMappingLogAnalyticsMetricOperatorEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLogAnalyticsMetricOperatorEnum(val string) (LogAnalyticsMetricOperatorEnum, bool) {
	enum, ok := mappingLogAnalyticsMetricOperatorEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
