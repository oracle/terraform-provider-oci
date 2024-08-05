// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SummarizeRepositoryAnalyticsDetails Details of the user configured settings for viewing the metrics at repository level.
type SummarizeRepositoryAnalyticsDetails struct {

	// The name of the metric to be filtered.
	RepositoryMetrics []MetricNameEnum `mandatory:"true" json:"repositoryMetrics"`

	// The beginning of the metric data query time range.
	StartTime *common.SDKTime `mandatory:"true" json:"startTime"`

	// Email address of the author.
	AuthorEmail *string `mandatory:"false" json:"authorEmail"`

	// Metrics aggregated for the defined period.
	AggregationDuration SummarizeRepositoryAnalyticsDetailsAggregationDurationEnum `mandatory:"false" json:"aggregationDuration,omitempty"`

	// The end of the metric data query time range.
	EndTime *common.SDKTime `mandatory:"false" json:"endTime"`

	// Attribute by which metric data has to be grouped
	GroupBy SummarizeRepositoryAnalyticsDetailsGroupByEnum `mandatory:"false" json:"groupBy,omitempty"`
}

func (m SummarizeRepositoryAnalyticsDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SummarizeRepositoryAnalyticsDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range m.RepositoryMetrics {
		if _, ok := GetMappingMetricNameEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RepositoryMetrics: %s. Supported values are: %s.", val, strings.Join(GetMetricNameEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingSummarizeRepositoryAnalyticsDetailsAggregationDurationEnum(string(m.AggregationDuration)); !ok && m.AggregationDuration != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AggregationDuration: %s. Supported values are: %s.", m.AggregationDuration, strings.Join(GetSummarizeRepositoryAnalyticsDetailsAggregationDurationEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeRepositoryAnalyticsDetailsGroupByEnum(string(m.GroupBy)); !ok && m.GroupBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GroupBy: %s. Supported values are: %s.", m.GroupBy, strings.Join(GetSummarizeRepositoryAnalyticsDetailsGroupByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeRepositoryAnalyticsDetailsAggregationDurationEnum Enum with underlying type: string
type SummarizeRepositoryAnalyticsDetailsAggregationDurationEnum string

// Set of constants representing the allowable values for SummarizeRepositoryAnalyticsDetailsAggregationDurationEnum
const (
	SummarizeRepositoryAnalyticsDetailsAggregationDurationDaily   SummarizeRepositoryAnalyticsDetailsAggregationDurationEnum = "DAILY"
	SummarizeRepositoryAnalyticsDetailsAggregationDurationWeekly  SummarizeRepositoryAnalyticsDetailsAggregationDurationEnum = "WEEKLY"
	SummarizeRepositoryAnalyticsDetailsAggregationDurationMonthly SummarizeRepositoryAnalyticsDetailsAggregationDurationEnum = "MONTHLY"
	SummarizeRepositoryAnalyticsDetailsAggregationDurationYearly  SummarizeRepositoryAnalyticsDetailsAggregationDurationEnum = "YEARLY"
)

var mappingSummarizeRepositoryAnalyticsDetailsAggregationDurationEnum = map[string]SummarizeRepositoryAnalyticsDetailsAggregationDurationEnum{
	"DAILY":   SummarizeRepositoryAnalyticsDetailsAggregationDurationDaily,
	"WEEKLY":  SummarizeRepositoryAnalyticsDetailsAggregationDurationWeekly,
	"MONTHLY": SummarizeRepositoryAnalyticsDetailsAggregationDurationMonthly,
	"YEARLY":  SummarizeRepositoryAnalyticsDetailsAggregationDurationYearly,
}

var mappingSummarizeRepositoryAnalyticsDetailsAggregationDurationEnumLowerCase = map[string]SummarizeRepositoryAnalyticsDetailsAggregationDurationEnum{
	"daily":   SummarizeRepositoryAnalyticsDetailsAggregationDurationDaily,
	"weekly":  SummarizeRepositoryAnalyticsDetailsAggregationDurationWeekly,
	"monthly": SummarizeRepositoryAnalyticsDetailsAggregationDurationMonthly,
	"yearly":  SummarizeRepositoryAnalyticsDetailsAggregationDurationYearly,
}

// GetSummarizeRepositoryAnalyticsDetailsAggregationDurationEnumValues Enumerates the set of values for SummarizeRepositoryAnalyticsDetailsAggregationDurationEnum
func GetSummarizeRepositoryAnalyticsDetailsAggregationDurationEnumValues() []SummarizeRepositoryAnalyticsDetailsAggregationDurationEnum {
	values := make([]SummarizeRepositoryAnalyticsDetailsAggregationDurationEnum, 0)
	for _, v := range mappingSummarizeRepositoryAnalyticsDetailsAggregationDurationEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeRepositoryAnalyticsDetailsAggregationDurationEnumStringValues Enumerates the set of values in String for SummarizeRepositoryAnalyticsDetailsAggregationDurationEnum
func GetSummarizeRepositoryAnalyticsDetailsAggregationDurationEnumStringValues() []string {
	return []string{
		"DAILY",
		"WEEKLY",
		"MONTHLY",
		"YEARLY",
	}
}

// GetMappingSummarizeRepositoryAnalyticsDetailsAggregationDurationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeRepositoryAnalyticsDetailsAggregationDurationEnum(val string) (SummarizeRepositoryAnalyticsDetailsAggregationDurationEnum, bool) {
	enum, ok := mappingSummarizeRepositoryAnalyticsDetailsAggregationDurationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeRepositoryAnalyticsDetailsGroupByEnum Enum with underlying type: string
type SummarizeRepositoryAnalyticsDetailsGroupByEnum string

// Set of constants representing the allowable values for SummarizeRepositoryAnalyticsDetailsGroupByEnum
const (
	SummarizeRepositoryAnalyticsDetailsGroupByAuthor SummarizeRepositoryAnalyticsDetailsGroupByEnum = "AUTHOR"
)

var mappingSummarizeRepositoryAnalyticsDetailsGroupByEnum = map[string]SummarizeRepositoryAnalyticsDetailsGroupByEnum{
	"AUTHOR": SummarizeRepositoryAnalyticsDetailsGroupByAuthor,
}

var mappingSummarizeRepositoryAnalyticsDetailsGroupByEnumLowerCase = map[string]SummarizeRepositoryAnalyticsDetailsGroupByEnum{
	"author": SummarizeRepositoryAnalyticsDetailsGroupByAuthor,
}

// GetSummarizeRepositoryAnalyticsDetailsGroupByEnumValues Enumerates the set of values for SummarizeRepositoryAnalyticsDetailsGroupByEnum
func GetSummarizeRepositoryAnalyticsDetailsGroupByEnumValues() []SummarizeRepositoryAnalyticsDetailsGroupByEnum {
	values := make([]SummarizeRepositoryAnalyticsDetailsGroupByEnum, 0)
	for _, v := range mappingSummarizeRepositoryAnalyticsDetailsGroupByEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeRepositoryAnalyticsDetailsGroupByEnumStringValues Enumerates the set of values in String for SummarizeRepositoryAnalyticsDetailsGroupByEnum
func GetSummarizeRepositoryAnalyticsDetailsGroupByEnumStringValues() []string {
	return []string{
		"AUTHOR",
	}
}

// GetMappingSummarizeRepositoryAnalyticsDetailsGroupByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeRepositoryAnalyticsDetailsGroupByEnum(val string) (SummarizeRepositoryAnalyticsDetailsGroupByEnum, bool) {
	enum, ok := mappingSummarizeRepositoryAnalyticsDetailsGroupByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
