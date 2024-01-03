// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// FindingAnalyticsSummary The summary of information about the analytics data of findings or top findings.
// It includes details such as metric name, findinKey,
// title (topFindingCategory for top finding), severity (topFindingStatus for top finding) and targetId.
type FindingAnalyticsSummary struct {

	// The name of the aggregation metric.
	MetricName FindingAnalyticsSummaryMetricNameEnum `mandatory:"true" json:"metricName"`

	// The total count for the aggregation metric.
	Count *int64 `mandatory:"true" json:"count"`

	Dimensions *FindingAnalyticsDimensions `mandatory:"false" json:"dimensions"`
}

func (m FindingAnalyticsSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FindingAnalyticsSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingFindingAnalyticsSummaryMetricNameEnum(string(m.MetricName)); !ok && m.MetricName != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MetricName: %s. Supported values are: %s.", m.MetricName, strings.Join(GetFindingAnalyticsSummaryMetricNameEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// FindingAnalyticsSummaryMetricNameEnum Enum with underlying type: string
type FindingAnalyticsSummaryMetricNameEnum string

// Set of constants representing the allowable values for FindingAnalyticsSummaryMetricNameEnum
const (
	FindingAnalyticsSummaryMetricNameTopFindingStats FindingAnalyticsSummaryMetricNameEnum = "TOP_FINDING_STATS"
	FindingAnalyticsSummaryMetricNameFindingStats    FindingAnalyticsSummaryMetricNameEnum = "FINDING_STATS"
)

var mappingFindingAnalyticsSummaryMetricNameEnum = map[string]FindingAnalyticsSummaryMetricNameEnum{
	"TOP_FINDING_STATS": FindingAnalyticsSummaryMetricNameTopFindingStats,
	"FINDING_STATS":     FindingAnalyticsSummaryMetricNameFindingStats,
}

var mappingFindingAnalyticsSummaryMetricNameEnumLowerCase = map[string]FindingAnalyticsSummaryMetricNameEnum{
	"top_finding_stats": FindingAnalyticsSummaryMetricNameTopFindingStats,
	"finding_stats":     FindingAnalyticsSummaryMetricNameFindingStats,
}

// GetFindingAnalyticsSummaryMetricNameEnumValues Enumerates the set of values for FindingAnalyticsSummaryMetricNameEnum
func GetFindingAnalyticsSummaryMetricNameEnumValues() []FindingAnalyticsSummaryMetricNameEnum {
	values := make([]FindingAnalyticsSummaryMetricNameEnum, 0)
	for _, v := range mappingFindingAnalyticsSummaryMetricNameEnum {
		values = append(values, v)
	}
	return values
}

// GetFindingAnalyticsSummaryMetricNameEnumStringValues Enumerates the set of values in String for FindingAnalyticsSummaryMetricNameEnum
func GetFindingAnalyticsSummaryMetricNameEnumStringValues() []string {
	return []string{
		"TOP_FINDING_STATS",
		"FINDING_STATS",
	}
}

// GetMappingFindingAnalyticsSummaryMetricNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFindingAnalyticsSummaryMetricNameEnum(val string) (FindingAnalyticsSummaryMetricNameEnum, bool) {
	enum, ok := mappingFindingAnalyticsSummaryMetricNameEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
