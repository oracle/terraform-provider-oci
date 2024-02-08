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

// SummarizeHostInsightNetworkUsageTrendAggregationCollection Top level response object.
type SummarizeHostInsightNetworkUsageTrendAggregationCollection struct {

	// The start timestamp that was passed into the request.
	TimeIntervalStart *common.SDKTime `mandatory:"true" json:"timeIntervalStart"`

	// The end timestamp that was passed into the request.
	TimeIntervalEnd *common.SDKTime `mandatory:"true" json:"timeIntervalEnd"`

	// Displays usage unit ( CORES, GB , PERCENT, MBPS)
	UsageUnit SummarizeHostInsightNetworkUsageTrendAggregationCollectionUsageUnitEnum `mandatory:"true" json:"usageUnit"`

	// Time duration in milliseconds between data points (one hour or one day).
	ItemDurationInMs *int64 `mandatory:"true" json:"itemDurationInMs"`

	// Collection of Usage Data with time stamps for all network interfaces.
	Items []NetworkUsageTrendAggregation `mandatory:"true" json:"items"`
}

func (m SummarizeHostInsightNetworkUsageTrendAggregationCollection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SummarizeHostInsightNetworkUsageTrendAggregationCollection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSummarizeHostInsightNetworkUsageTrendAggregationCollectionUsageUnitEnum(string(m.UsageUnit)); !ok && m.UsageUnit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UsageUnit: %s. Supported values are: %s.", m.UsageUnit, strings.Join(GetSummarizeHostInsightNetworkUsageTrendAggregationCollectionUsageUnitEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeHostInsightNetworkUsageTrendAggregationCollectionUsageUnitEnum Enum with underlying type: string
type SummarizeHostInsightNetworkUsageTrendAggregationCollectionUsageUnitEnum string

// Set of constants representing the allowable values for SummarizeHostInsightNetworkUsageTrendAggregationCollectionUsageUnitEnum
const (
	SummarizeHostInsightNetworkUsageTrendAggregationCollectionUsageUnitCores   SummarizeHostInsightNetworkUsageTrendAggregationCollectionUsageUnitEnum = "CORES"
	SummarizeHostInsightNetworkUsageTrendAggregationCollectionUsageUnitGb      SummarizeHostInsightNetworkUsageTrendAggregationCollectionUsageUnitEnum = "GB"
	SummarizeHostInsightNetworkUsageTrendAggregationCollectionUsageUnitMbps    SummarizeHostInsightNetworkUsageTrendAggregationCollectionUsageUnitEnum = "MBPS"
	SummarizeHostInsightNetworkUsageTrendAggregationCollectionUsageUnitIops    SummarizeHostInsightNetworkUsageTrendAggregationCollectionUsageUnitEnum = "IOPS"
	SummarizeHostInsightNetworkUsageTrendAggregationCollectionUsageUnitPercent SummarizeHostInsightNetworkUsageTrendAggregationCollectionUsageUnitEnum = "PERCENT"
)

var mappingSummarizeHostInsightNetworkUsageTrendAggregationCollectionUsageUnitEnum = map[string]SummarizeHostInsightNetworkUsageTrendAggregationCollectionUsageUnitEnum{
	"CORES":   SummarizeHostInsightNetworkUsageTrendAggregationCollectionUsageUnitCores,
	"GB":      SummarizeHostInsightNetworkUsageTrendAggregationCollectionUsageUnitGb,
	"MBPS":    SummarizeHostInsightNetworkUsageTrendAggregationCollectionUsageUnitMbps,
	"IOPS":    SummarizeHostInsightNetworkUsageTrendAggregationCollectionUsageUnitIops,
	"PERCENT": SummarizeHostInsightNetworkUsageTrendAggregationCollectionUsageUnitPercent,
}

var mappingSummarizeHostInsightNetworkUsageTrendAggregationCollectionUsageUnitEnumLowerCase = map[string]SummarizeHostInsightNetworkUsageTrendAggregationCollectionUsageUnitEnum{
	"cores":   SummarizeHostInsightNetworkUsageTrendAggregationCollectionUsageUnitCores,
	"gb":      SummarizeHostInsightNetworkUsageTrendAggregationCollectionUsageUnitGb,
	"mbps":    SummarizeHostInsightNetworkUsageTrendAggregationCollectionUsageUnitMbps,
	"iops":    SummarizeHostInsightNetworkUsageTrendAggregationCollectionUsageUnitIops,
	"percent": SummarizeHostInsightNetworkUsageTrendAggregationCollectionUsageUnitPercent,
}

// GetSummarizeHostInsightNetworkUsageTrendAggregationCollectionUsageUnitEnumValues Enumerates the set of values for SummarizeHostInsightNetworkUsageTrendAggregationCollectionUsageUnitEnum
func GetSummarizeHostInsightNetworkUsageTrendAggregationCollectionUsageUnitEnumValues() []SummarizeHostInsightNetworkUsageTrendAggregationCollectionUsageUnitEnum {
	values := make([]SummarizeHostInsightNetworkUsageTrendAggregationCollectionUsageUnitEnum, 0)
	for _, v := range mappingSummarizeHostInsightNetworkUsageTrendAggregationCollectionUsageUnitEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeHostInsightNetworkUsageTrendAggregationCollectionUsageUnitEnumStringValues Enumerates the set of values in String for SummarizeHostInsightNetworkUsageTrendAggregationCollectionUsageUnitEnum
func GetSummarizeHostInsightNetworkUsageTrendAggregationCollectionUsageUnitEnumStringValues() []string {
	return []string{
		"CORES",
		"GB",
		"MBPS",
		"IOPS",
		"PERCENT",
	}
}

// GetMappingSummarizeHostInsightNetworkUsageTrendAggregationCollectionUsageUnitEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeHostInsightNetworkUsageTrendAggregationCollectionUsageUnitEnum(val string) (SummarizeHostInsightNetworkUsageTrendAggregationCollectionUsageUnitEnum, bool) {
	enum, ok := mappingSummarizeHostInsightNetworkUsageTrendAggregationCollectionUsageUnitEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
