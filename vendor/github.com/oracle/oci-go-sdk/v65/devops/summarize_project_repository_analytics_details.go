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

// SummarizeProjectRepositoryAnalyticsDetails Details of the user configured settings for viewing the metrics at project level.
type SummarizeProjectRepositoryAnalyticsDetails struct {

	// The name of the metric to be filtered.
	RepositoryMetrics []MetricNameEnum `mandatory:"true" json:"repositoryMetrics"`

	// The beginning of the metric data query time range.
	StartTime *common.SDKTime `mandatory:"true" json:"startTime"`

	// Email address of the author.
	AuthorEmail *string `mandatory:"false" json:"authorEmail"`

	// Metrics aggregated for the defined period.
	AggregationDuration SummarizeProjectRepositoryAnalyticsDetailsAggregationDurationEnum `mandatory:"false" json:"aggregationDuration,omitempty"`

	// The end of the metric data query time range.
	EndTime *common.SDKTime `mandatory:"false" json:"endTime"`

	// Attribute by which metric data has to be grouped
	GroupBy SummarizeProjectRepositoryAnalyticsDetailsGroupByEnum `mandatory:"false" json:"groupBy,omitempty"`
}

func (m SummarizeProjectRepositoryAnalyticsDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SummarizeProjectRepositoryAnalyticsDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range m.RepositoryMetrics {
		if _, ok := GetMappingMetricNameEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RepositoryMetrics: %s. Supported values are: %s.", val, strings.Join(GetMetricNameEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingSummarizeProjectRepositoryAnalyticsDetailsAggregationDurationEnum(string(m.AggregationDuration)); !ok && m.AggregationDuration != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AggregationDuration: %s. Supported values are: %s.", m.AggregationDuration, strings.Join(GetSummarizeProjectRepositoryAnalyticsDetailsAggregationDurationEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeProjectRepositoryAnalyticsDetailsGroupByEnum(string(m.GroupBy)); !ok && m.GroupBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GroupBy: %s. Supported values are: %s.", m.GroupBy, strings.Join(GetSummarizeProjectRepositoryAnalyticsDetailsGroupByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeProjectRepositoryAnalyticsDetailsAggregationDurationEnum Enum with underlying type: string
type SummarizeProjectRepositoryAnalyticsDetailsAggregationDurationEnum string

// Set of constants representing the allowable values for SummarizeProjectRepositoryAnalyticsDetailsAggregationDurationEnum
const (
	SummarizeProjectRepositoryAnalyticsDetailsAggregationDurationDaily   SummarizeProjectRepositoryAnalyticsDetailsAggregationDurationEnum = "DAILY"
	SummarizeProjectRepositoryAnalyticsDetailsAggregationDurationWeekly  SummarizeProjectRepositoryAnalyticsDetailsAggregationDurationEnum = "WEEKLY"
	SummarizeProjectRepositoryAnalyticsDetailsAggregationDurationMonthly SummarizeProjectRepositoryAnalyticsDetailsAggregationDurationEnum = "MONTHLY"
	SummarizeProjectRepositoryAnalyticsDetailsAggregationDurationYearly  SummarizeProjectRepositoryAnalyticsDetailsAggregationDurationEnum = "YEARLY"
)

var mappingSummarizeProjectRepositoryAnalyticsDetailsAggregationDurationEnum = map[string]SummarizeProjectRepositoryAnalyticsDetailsAggregationDurationEnum{
	"DAILY":   SummarizeProjectRepositoryAnalyticsDetailsAggregationDurationDaily,
	"WEEKLY":  SummarizeProjectRepositoryAnalyticsDetailsAggregationDurationWeekly,
	"MONTHLY": SummarizeProjectRepositoryAnalyticsDetailsAggregationDurationMonthly,
	"YEARLY":  SummarizeProjectRepositoryAnalyticsDetailsAggregationDurationYearly,
}

var mappingSummarizeProjectRepositoryAnalyticsDetailsAggregationDurationEnumLowerCase = map[string]SummarizeProjectRepositoryAnalyticsDetailsAggregationDurationEnum{
	"daily":   SummarizeProjectRepositoryAnalyticsDetailsAggregationDurationDaily,
	"weekly":  SummarizeProjectRepositoryAnalyticsDetailsAggregationDurationWeekly,
	"monthly": SummarizeProjectRepositoryAnalyticsDetailsAggregationDurationMonthly,
	"yearly":  SummarizeProjectRepositoryAnalyticsDetailsAggregationDurationYearly,
}

// GetSummarizeProjectRepositoryAnalyticsDetailsAggregationDurationEnumValues Enumerates the set of values for SummarizeProjectRepositoryAnalyticsDetailsAggregationDurationEnum
func GetSummarizeProjectRepositoryAnalyticsDetailsAggregationDurationEnumValues() []SummarizeProjectRepositoryAnalyticsDetailsAggregationDurationEnum {
	values := make([]SummarizeProjectRepositoryAnalyticsDetailsAggregationDurationEnum, 0)
	for _, v := range mappingSummarizeProjectRepositoryAnalyticsDetailsAggregationDurationEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeProjectRepositoryAnalyticsDetailsAggregationDurationEnumStringValues Enumerates the set of values in String for SummarizeProjectRepositoryAnalyticsDetailsAggregationDurationEnum
func GetSummarizeProjectRepositoryAnalyticsDetailsAggregationDurationEnumStringValues() []string {
	return []string{
		"DAILY",
		"WEEKLY",
		"MONTHLY",
		"YEARLY",
	}
}

// GetMappingSummarizeProjectRepositoryAnalyticsDetailsAggregationDurationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeProjectRepositoryAnalyticsDetailsAggregationDurationEnum(val string) (SummarizeProjectRepositoryAnalyticsDetailsAggregationDurationEnum, bool) {
	enum, ok := mappingSummarizeProjectRepositoryAnalyticsDetailsAggregationDurationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeProjectRepositoryAnalyticsDetailsGroupByEnum Enum with underlying type: string
type SummarizeProjectRepositoryAnalyticsDetailsGroupByEnum string

// Set of constants representing the allowable values for SummarizeProjectRepositoryAnalyticsDetailsGroupByEnum
const (
	SummarizeProjectRepositoryAnalyticsDetailsGroupByAuthor SummarizeProjectRepositoryAnalyticsDetailsGroupByEnum = "AUTHOR"
)

var mappingSummarizeProjectRepositoryAnalyticsDetailsGroupByEnum = map[string]SummarizeProjectRepositoryAnalyticsDetailsGroupByEnum{
	"AUTHOR": SummarizeProjectRepositoryAnalyticsDetailsGroupByAuthor,
}

var mappingSummarizeProjectRepositoryAnalyticsDetailsGroupByEnumLowerCase = map[string]SummarizeProjectRepositoryAnalyticsDetailsGroupByEnum{
	"author": SummarizeProjectRepositoryAnalyticsDetailsGroupByAuthor,
}

// GetSummarizeProjectRepositoryAnalyticsDetailsGroupByEnumValues Enumerates the set of values for SummarizeProjectRepositoryAnalyticsDetailsGroupByEnum
func GetSummarizeProjectRepositoryAnalyticsDetailsGroupByEnumValues() []SummarizeProjectRepositoryAnalyticsDetailsGroupByEnum {
	values := make([]SummarizeProjectRepositoryAnalyticsDetailsGroupByEnum, 0)
	for _, v := range mappingSummarizeProjectRepositoryAnalyticsDetailsGroupByEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeProjectRepositoryAnalyticsDetailsGroupByEnumStringValues Enumerates the set of values in String for SummarizeProjectRepositoryAnalyticsDetailsGroupByEnum
func GetSummarizeProjectRepositoryAnalyticsDetailsGroupByEnumStringValues() []string {
	return []string{
		"AUTHOR",
	}
}

// GetMappingSummarizeProjectRepositoryAnalyticsDetailsGroupByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeProjectRepositoryAnalyticsDetailsGroupByEnum(val string) (SummarizeProjectRepositoryAnalyticsDetailsGroupByEnum, bool) {
	enum, ok := mappingSummarizeProjectRepositoryAnalyticsDetailsGroupByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
