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

// SummarizeHostInsightsDiskStatisticsCollection Top level response object.
type SummarizeHostInsightsDiskStatisticsCollection struct {

	// The start timestamp that was passed into the request.
	TimeIntervalStart *common.SDKTime `mandatory:"true" json:"timeIntervalStart"`

	// The end timestamp that was passed into the request.
	TimeIntervalEnd *common.SDKTime `mandatory:"true" json:"timeIntervalEnd"`

	// Displays usage unit ( CORES, GB , PERCENT, MBPS)
	UsageUnit SummarizeHostInsightsDiskStatisticsCollectionUsageUnitEnum `mandatory:"true" json:"usageUnit"`

	// Time duration in milliseconds between data points (one hour or one day).
	ItemDurationInMs *int64 `mandatory:"true" json:"itemDurationInMs"`

	// Collection of Data for all disks in a host.
	Items []DiskStatistics `mandatory:"true" json:"items"`
}

func (m SummarizeHostInsightsDiskStatisticsCollection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SummarizeHostInsightsDiskStatisticsCollection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSummarizeHostInsightsDiskStatisticsCollectionUsageUnitEnum(string(m.UsageUnit)); !ok && m.UsageUnit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UsageUnit: %s. Supported values are: %s.", m.UsageUnit, strings.Join(GetSummarizeHostInsightsDiskStatisticsCollectionUsageUnitEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeHostInsightsDiskStatisticsCollectionUsageUnitEnum Enum with underlying type: string
type SummarizeHostInsightsDiskStatisticsCollectionUsageUnitEnum string

// Set of constants representing the allowable values for SummarizeHostInsightsDiskStatisticsCollectionUsageUnitEnum
const (
	SummarizeHostInsightsDiskStatisticsCollectionUsageUnitCores   SummarizeHostInsightsDiskStatisticsCollectionUsageUnitEnum = "CORES"
	SummarizeHostInsightsDiskStatisticsCollectionUsageUnitGb      SummarizeHostInsightsDiskStatisticsCollectionUsageUnitEnum = "GB"
	SummarizeHostInsightsDiskStatisticsCollectionUsageUnitMbps    SummarizeHostInsightsDiskStatisticsCollectionUsageUnitEnum = "MBPS"
	SummarizeHostInsightsDiskStatisticsCollectionUsageUnitIops    SummarizeHostInsightsDiskStatisticsCollectionUsageUnitEnum = "IOPS"
	SummarizeHostInsightsDiskStatisticsCollectionUsageUnitPercent SummarizeHostInsightsDiskStatisticsCollectionUsageUnitEnum = "PERCENT"
)

var mappingSummarizeHostInsightsDiskStatisticsCollectionUsageUnitEnum = map[string]SummarizeHostInsightsDiskStatisticsCollectionUsageUnitEnum{
	"CORES":   SummarizeHostInsightsDiskStatisticsCollectionUsageUnitCores,
	"GB":      SummarizeHostInsightsDiskStatisticsCollectionUsageUnitGb,
	"MBPS":    SummarizeHostInsightsDiskStatisticsCollectionUsageUnitMbps,
	"IOPS":    SummarizeHostInsightsDiskStatisticsCollectionUsageUnitIops,
	"PERCENT": SummarizeHostInsightsDiskStatisticsCollectionUsageUnitPercent,
}

var mappingSummarizeHostInsightsDiskStatisticsCollectionUsageUnitEnumLowerCase = map[string]SummarizeHostInsightsDiskStatisticsCollectionUsageUnitEnum{
	"cores":   SummarizeHostInsightsDiskStatisticsCollectionUsageUnitCores,
	"gb":      SummarizeHostInsightsDiskStatisticsCollectionUsageUnitGb,
	"mbps":    SummarizeHostInsightsDiskStatisticsCollectionUsageUnitMbps,
	"iops":    SummarizeHostInsightsDiskStatisticsCollectionUsageUnitIops,
	"percent": SummarizeHostInsightsDiskStatisticsCollectionUsageUnitPercent,
}

// GetSummarizeHostInsightsDiskStatisticsCollectionUsageUnitEnumValues Enumerates the set of values for SummarizeHostInsightsDiskStatisticsCollectionUsageUnitEnum
func GetSummarizeHostInsightsDiskStatisticsCollectionUsageUnitEnumValues() []SummarizeHostInsightsDiskStatisticsCollectionUsageUnitEnum {
	values := make([]SummarizeHostInsightsDiskStatisticsCollectionUsageUnitEnum, 0)
	for _, v := range mappingSummarizeHostInsightsDiskStatisticsCollectionUsageUnitEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeHostInsightsDiskStatisticsCollectionUsageUnitEnumStringValues Enumerates the set of values in String for SummarizeHostInsightsDiskStatisticsCollectionUsageUnitEnum
func GetSummarizeHostInsightsDiskStatisticsCollectionUsageUnitEnumStringValues() []string {
	return []string{
		"CORES",
		"GB",
		"MBPS",
		"IOPS",
		"PERCENT",
	}
}

// GetMappingSummarizeHostInsightsDiskStatisticsCollectionUsageUnitEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeHostInsightsDiskStatisticsCollectionUsageUnitEnum(val string) (SummarizeHostInsightsDiskStatisticsCollectionUsageUnitEnum, bool) {
	enum, ok := mappingSummarizeHostInsightsDiskStatisticsCollectionUsageUnitEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
