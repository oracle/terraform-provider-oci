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
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SummarizeDatabaseInsightTablespaceUsageTrendAggregationCollection Top level response object.
type SummarizeDatabaseInsightTablespaceUsageTrendAggregationCollection struct {

	// The start timestamp that was passed into the request.
	TimeIntervalStart *common.SDKTime `mandatory:"true" json:"timeIntervalStart"`

	// The end timestamp that was passed into the request.
	TimeIntervalEnd *common.SDKTime `mandatory:"true" json:"timeIntervalEnd"`

	// Displays usage unit ( CORES, GB , PERCENT, MBPS)
	UsageUnit SummarizeDatabaseInsightTablespaceUsageTrendAggregationCollectionUsageUnitEnum `mandatory:"true" json:"usageUnit"`

	// Time duration in milliseconds between data points (one hour or one day).
	ItemDurationInMs *int64 `mandatory:"true" json:"itemDurationInMs"`

	// Collection of Usage Data with time stamps for top five tablespace
	Items []TablespaceUsageTrendAggregation `mandatory:"true" json:"items"`
}

func (m SummarizeDatabaseInsightTablespaceUsageTrendAggregationCollection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SummarizeDatabaseInsightTablespaceUsageTrendAggregationCollection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSummarizeDatabaseInsightTablespaceUsageTrendAggregationCollectionUsageUnitEnum(string(m.UsageUnit)); !ok && m.UsageUnit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UsageUnit: %s. Supported values are: %s.", m.UsageUnit, strings.Join(GetSummarizeDatabaseInsightTablespaceUsageTrendAggregationCollectionUsageUnitEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeDatabaseInsightTablespaceUsageTrendAggregationCollectionUsageUnitEnum Enum with underlying type: string
type SummarizeDatabaseInsightTablespaceUsageTrendAggregationCollectionUsageUnitEnum string

// Set of constants representing the allowable values for SummarizeDatabaseInsightTablespaceUsageTrendAggregationCollectionUsageUnitEnum
const (
	SummarizeDatabaseInsightTablespaceUsageTrendAggregationCollectionUsageUnitCores   SummarizeDatabaseInsightTablespaceUsageTrendAggregationCollectionUsageUnitEnum = "CORES"
	SummarizeDatabaseInsightTablespaceUsageTrendAggregationCollectionUsageUnitGb      SummarizeDatabaseInsightTablespaceUsageTrendAggregationCollectionUsageUnitEnum = "GB"
	SummarizeDatabaseInsightTablespaceUsageTrendAggregationCollectionUsageUnitMbps    SummarizeDatabaseInsightTablespaceUsageTrendAggregationCollectionUsageUnitEnum = "MBPS"
	SummarizeDatabaseInsightTablespaceUsageTrendAggregationCollectionUsageUnitIops    SummarizeDatabaseInsightTablespaceUsageTrendAggregationCollectionUsageUnitEnum = "IOPS"
	SummarizeDatabaseInsightTablespaceUsageTrendAggregationCollectionUsageUnitPercent SummarizeDatabaseInsightTablespaceUsageTrendAggregationCollectionUsageUnitEnum = "PERCENT"
)

var mappingSummarizeDatabaseInsightTablespaceUsageTrendAggregationCollectionUsageUnitEnum = map[string]SummarizeDatabaseInsightTablespaceUsageTrendAggregationCollectionUsageUnitEnum{
	"CORES":   SummarizeDatabaseInsightTablespaceUsageTrendAggregationCollectionUsageUnitCores,
	"GB":      SummarizeDatabaseInsightTablespaceUsageTrendAggregationCollectionUsageUnitGb,
	"MBPS":    SummarizeDatabaseInsightTablespaceUsageTrendAggregationCollectionUsageUnitMbps,
	"IOPS":    SummarizeDatabaseInsightTablespaceUsageTrendAggregationCollectionUsageUnitIops,
	"PERCENT": SummarizeDatabaseInsightTablespaceUsageTrendAggregationCollectionUsageUnitPercent,
}

var mappingSummarizeDatabaseInsightTablespaceUsageTrendAggregationCollectionUsageUnitEnumLowerCase = map[string]SummarizeDatabaseInsightTablespaceUsageTrendAggregationCollectionUsageUnitEnum{
	"cores":   SummarizeDatabaseInsightTablespaceUsageTrendAggregationCollectionUsageUnitCores,
	"gb":      SummarizeDatabaseInsightTablespaceUsageTrendAggregationCollectionUsageUnitGb,
	"mbps":    SummarizeDatabaseInsightTablespaceUsageTrendAggregationCollectionUsageUnitMbps,
	"iops":    SummarizeDatabaseInsightTablespaceUsageTrendAggregationCollectionUsageUnitIops,
	"percent": SummarizeDatabaseInsightTablespaceUsageTrendAggregationCollectionUsageUnitPercent,
}

// GetSummarizeDatabaseInsightTablespaceUsageTrendAggregationCollectionUsageUnitEnumValues Enumerates the set of values for SummarizeDatabaseInsightTablespaceUsageTrendAggregationCollectionUsageUnitEnum
func GetSummarizeDatabaseInsightTablespaceUsageTrendAggregationCollectionUsageUnitEnumValues() []SummarizeDatabaseInsightTablespaceUsageTrendAggregationCollectionUsageUnitEnum {
	values := make([]SummarizeDatabaseInsightTablespaceUsageTrendAggregationCollectionUsageUnitEnum, 0)
	for _, v := range mappingSummarizeDatabaseInsightTablespaceUsageTrendAggregationCollectionUsageUnitEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeDatabaseInsightTablespaceUsageTrendAggregationCollectionUsageUnitEnumStringValues Enumerates the set of values in String for SummarizeDatabaseInsightTablespaceUsageTrendAggregationCollectionUsageUnitEnum
func GetSummarizeDatabaseInsightTablespaceUsageTrendAggregationCollectionUsageUnitEnumStringValues() []string {
	return []string{
		"CORES",
		"GB",
		"MBPS",
		"IOPS",
		"PERCENT",
	}
}

// GetMappingSummarizeDatabaseInsightTablespaceUsageTrendAggregationCollectionUsageUnitEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeDatabaseInsightTablespaceUsageTrendAggregationCollectionUsageUnitEnum(val string) (SummarizeDatabaseInsightTablespaceUsageTrendAggregationCollectionUsageUnitEnum, bool) {
	enum, ok := mappingSummarizeDatabaseInsightTablespaceUsageTrendAggregationCollectionUsageUnitEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
