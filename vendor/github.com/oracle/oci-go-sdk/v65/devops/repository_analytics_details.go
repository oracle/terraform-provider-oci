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

// RepositoryAnalyticsDetails Details of the user configured settings for viewing the metrics.
type RepositoryAnalyticsDetails struct {

	// The name of the metric to be filtered.
	RepositoryMetrics []MetricNameEnum `mandatory:"true" json:"repositoryMetrics"`

	// The beginning of the metric data query time range.
	StartTime *common.SDKTime `mandatory:"true" json:"startTime"`

	// Email address of the author.
	AuthorEmail *string `mandatory:"false" json:"authorEmail"`

	// Metrics aggregated for the defined period.
	AggregationDuration RepositoryAnalyticsDetailsAggregationDurationEnum `mandatory:"false" json:"aggregationDuration,omitempty"`

	// The end of the metric data query time range.
	EndTime *common.SDKTime `mandatory:"false" json:"endTime"`

	// Attribute by which metric data has to be grouped
	GroupBy RepositoryAnalyticsDetailsGroupByEnum `mandatory:"false" json:"groupBy,omitempty"`
}

func (m RepositoryAnalyticsDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RepositoryAnalyticsDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range m.RepositoryMetrics {
		if _, ok := GetMappingMetricNameEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RepositoryMetrics: %s. Supported values are: %s.", val, strings.Join(GetMetricNameEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingRepositoryAnalyticsDetailsAggregationDurationEnum(string(m.AggregationDuration)); !ok && m.AggregationDuration != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AggregationDuration: %s. Supported values are: %s.", m.AggregationDuration, strings.Join(GetRepositoryAnalyticsDetailsAggregationDurationEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRepositoryAnalyticsDetailsGroupByEnum(string(m.GroupBy)); !ok && m.GroupBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GroupBy: %s. Supported values are: %s.", m.GroupBy, strings.Join(GetRepositoryAnalyticsDetailsGroupByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RepositoryAnalyticsDetailsAggregationDurationEnum Enum with underlying type: string
type RepositoryAnalyticsDetailsAggregationDurationEnum string

// Set of constants representing the allowable values for RepositoryAnalyticsDetailsAggregationDurationEnum
const (
	RepositoryAnalyticsDetailsAggregationDurationDaily   RepositoryAnalyticsDetailsAggregationDurationEnum = "DAILY"
	RepositoryAnalyticsDetailsAggregationDurationWeekly  RepositoryAnalyticsDetailsAggregationDurationEnum = "WEEKLY"
	RepositoryAnalyticsDetailsAggregationDurationMonthly RepositoryAnalyticsDetailsAggregationDurationEnum = "MONTHLY"
	RepositoryAnalyticsDetailsAggregationDurationYearly  RepositoryAnalyticsDetailsAggregationDurationEnum = "YEARLY"
)

var mappingRepositoryAnalyticsDetailsAggregationDurationEnum = map[string]RepositoryAnalyticsDetailsAggregationDurationEnum{
	"DAILY":   RepositoryAnalyticsDetailsAggregationDurationDaily,
	"WEEKLY":  RepositoryAnalyticsDetailsAggregationDurationWeekly,
	"MONTHLY": RepositoryAnalyticsDetailsAggregationDurationMonthly,
	"YEARLY":  RepositoryAnalyticsDetailsAggregationDurationYearly,
}

var mappingRepositoryAnalyticsDetailsAggregationDurationEnumLowerCase = map[string]RepositoryAnalyticsDetailsAggregationDurationEnum{
	"daily":   RepositoryAnalyticsDetailsAggregationDurationDaily,
	"weekly":  RepositoryAnalyticsDetailsAggregationDurationWeekly,
	"monthly": RepositoryAnalyticsDetailsAggregationDurationMonthly,
	"yearly":  RepositoryAnalyticsDetailsAggregationDurationYearly,
}

// GetRepositoryAnalyticsDetailsAggregationDurationEnumValues Enumerates the set of values for RepositoryAnalyticsDetailsAggregationDurationEnum
func GetRepositoryAnalyticsDetailsAggregationDurationEnumValues() []RepositoryAnalyticsDetailsAggregationDurationEnum {
	values := make([]RepositoryAnalyticsDetailsAggregationDurationEnum, 0)
	for _, v := range mappingRepositoryAnalyticsDetailsAggregationDurationEnum {
		values = append(values, v)
	}
	return values
}

// GetRepositoryAnalyticsDetailsAggregationDurationEnumStringValues Enumerates the set of values in String for RepositoryAnalyticsDetailsAggregationDurationEnum
func GetRepositoryAnalyticsDetailsAggregationDurationEnumStringValues() []string {
	return []string{
		"DAILY",
		"WEEKLY",
		"MONTHLY",
		"YEARLY",
	}
}

// GetMappingRepositoryAnalyticsDetailsAggregationDurationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRepositoryAnalyticsDetailsAggregationDurationEnum(val string) (RepositoryAnalyticsDetailsAggregationDurationEnum, bool) {
	enum, ok := mappingRepositoryAnalyticsDetailsAggregationDurationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// RepositoryAnalyticsDetailsGroupByEnum Enum with underlying type: string
type RepositoryAnalyticsDetailsGroupByEnum string

// Set of constants representing the allowable values for RepositoryAnalyticsDetailsGroupByEnum
const (
	RepositoryAnalyticsDetailsGroupByAuthor RepositoryAnalyticsDetailsGroupByEnum = "AUTHOR"
)

var mappingRepositoryAnalyticsDetailsGroupByEnum = map[string]RepositoryAnalyticsDetailsGroupByEnum{
	"AUTHOR": RepositoryAnalyticsDetailsGroupByAuthor,
}

var mappingRepositoryAnalyticsDetailsGroupByEnumLowerCase = map[string]RepositoryAnalyticsDetailsGroupByEnum{
	"author": RepositoryAnalyticsDetailsGroupByAuthor,
}

// GetRepositoryAnalyticsDetailsGroupByEnumValues Enumerates the set of values for RepositoryAnalyticsDetailsGroupByEnum
func GetRepositoryAnalyticsDetailsGroupByEnumValues() []RepositoryAnalyticsDetailsGroupByEnum {
	values := make([]RepositoryAnalyticsDetailsGroupByEnum, 0)
	for _, v := range mappingRepositoryAnalyticsDetailsGroupByEnum {
		values = append(values, v)
	}
	return values
}

// GetRepositoryAnalyticsDetailsGroupByEnumStringValues Enumerates the set of values in String for RepositoryAnalyticsDetailsGroupByEnum
func GetRepositoryAnalyticsDetailsGroupByEnumStringValues() []string {
	return []string{
		"AUTHOR",
	}
}

// GetMappingRepositoryAnalyticsDetailsGroupByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRepositoryAnalyticsDetailsGroupByEnum(val string) (RepositoryAnalyticsDetailsGroupByEnum, bool) {
	enum, ok := mappingRepositoryAnalyticsDetailsGroupByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
