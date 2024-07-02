// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SummarizeHostInsightStorageUsageTrendAggregationCollection Top level response object.
type SummarizeHostInsightStorageUsageTrendAggregationCollection struct {

	// The start timestamp that was passed into the request.
	TimeIntervalStart *common.SDKTime `mandatory:"true" json:"timeIntervalStart"`

	// The end timestamp that was passed into the request.
	TimeIntervalEnd *common.SDKTime `mandatory:"true" json:"timeIntervalEnd"`

	// Displays usage unit ( CORES, GB , PERCENT, MBPS)
	UsageUnit SummarizeHostInsightStorageUsageTrendAggregationCollectionUsageUnitEnum `mandatory:"true" json:"usageUnit"`

	// Time duration in milliseconds between data points (one hour or one day).
	ItemDurationInMs *int64 `mandatory:"true" json:"itemDurationInMs"`

	// Collection of Usage Data with time stamps for all filesystems.
	Items []StorageUsageTrendAggregation `mandatory:"true" json:"items"`
}

func (m SummarizeHostInsightStorageUsageTrendAggregationCollection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SummarizeHostInsightStorageUsageTrendAggregationCollection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSummarizeHostInsightStorageUsageTrendAggregationCollectionUsageUnitEnum(string(m.UsageUnit)); !ok && m.UsageUnit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UsageUnit: %s. Supported values are: %s.", m.UsageUnit, strings.Join(GetSummarizeHostInsightStorageUsageTrendAggregationCollectionUsageUnitEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeHostInsightStorageUsageTrendAggregationCollectionUsageUnitEnum Enum with underlying type: string
type SummarizeHostInsightStorageUsageTrendAggregationCollectionUsageUnitEnum string

// Set of constants representing the allowable values for SummarizeHostInsightStorageUsageTrendAggregationCollectionUsageUnitEnum
const (
	SummarizeHostInsightStorageUsageTrendAggregationCollectionUsageUnitCores   SummarizeHostInsightStorageUsageTrendAggregationCollectionUsageUnitEnum = "CORES"
	SummarizeHostInsightStorageUsageTrendAggregationCollectionUsageUnitGb      SummarizeHostInsightStorageUsageTrendAggregationCollectionUsageUnitEnum = "GB"
	SummarizeHostInsightStorageUsageTrendAggregationCollectionUsageUnitMbps    SummarizeHostInsightStorageUsageTrendAggregationCollectionUsageUnitEnum = "MBPS"
	SummarizeHostInsightStorageUsageTrendAggregationCollectionUsageUnitIops    SummarizeHostInsightStorageUsageTrendAggregationCollectionUsageUnitEnum = "IOPS"
	SummarizeHostInsightStorageUsageTrendAggregationCollectionUsageUnitPercent SummarizeHostInsightStorageUsageTrendAggregationCollectionUsageUnitEnum = "PERCENT"
)

var mappingSummarizeHostInsightStorageUsageTrendAggregationCollectionUsageUnitEnum = map[string]SummarizeHostInsightStorageUsageTrendAggregationCollectionUsageUnitEnum{
	"CORES":   SummarizeHostInsightStorageUsageTrendAggregationCollectionUsageUnitCores,
	"GB":      SummarizeHostInsightStorageUsageTrendAggregationCollectionUsageUnitGb,
	"MBPS":    SummarizeHostInsightStorageUsageTrendAggregationCollectionUsageUnitMbps,
	"IOPS":    SummarizeHostInsightStorageUsageTrendAggregationCollectionUsageUnitIops,
	"PERCENT": SummarizeHostInsightStorageUsageTrendAggregationCollectionUsageUnitPercent,
}

var mappingSummarizeHostInsightStorageUsageTrendAggregationCollectionUsageUnitEnumLowerCase = map[string]SummarizeHostInsightStorageUsageTrendAggregationCollectionUsageUnitEnum{
	"cores":   SummarizeHostInsightStorageUsageTrendAggregationCollectionUsageUnitCores,
	"gb":      SummarizeHostInsightStorageUsageTrendAggregationCollectionUsageUnitGb,
	"mbps":    SummarizeHostInsightStorageUsageTrendAggregationCollectionUsageUnitMbps,
	"iops":    SummarizeHostInsightStorageUsageTrendAggregationCollectionUsageUnitIops,
	"percent": SummarizeHostInsightStorageUsageTrendAggregationCollectionUsageUnitPercent,
}

// GetSummarizeHostInsightStorageUsageTrendAggregationCollectionUsageUnitEnumValues Enumerates the set of values for SummarizeHostInsightStorageUsageTrendAggregationCollectionUsageUnitEnum
func GetSummarizeHostInsightStorageUsageTrendAggregationCollectionUsageUnitEnumValues() []SummarizeHostInsightStorageUsageTrendAggregationCollectionUsageUnitEnum {
	values := make([]SummarizeHostInsightStorageUsageTrendAggregationCollectionUsageUnitEnum, 0)
	for _, v := range mappingSummarizeHostInsightStorageUsageTrendAggregationCollectionUsageUnitEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeHostInsightStorageUsageTrendAggregationCollectionUsageUnitEnumStringValues Enumerates the set of values in String for SummarizeHostInsightStorageUsageTrendAggregationCollectionUsageUnitEnum
func GetSummarizeHostInsightStorageUsageTrendAggregationCollectionUsageUnitEnumStringValues() []string {
	return []string{
		"CORES",
		"GB",
		"MBPS",
		"IOPS",
		"PERCENT",
	}
}

// GetMappingSummarizeHostInsightStorageUsageTrendAggregationCollectionUsageUnitEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeHostInsightStorageUsageTrendAggregationCollectionUsageUnitEnum(val string) (SummarizeHostInsightStorageUsageTrendAggregationCollectionUsageUnitEnum, bool) {
	enum, ok := mappingSummarizeHostInsightStorageUsageTrendAggregationCollectionUsageUnitEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
