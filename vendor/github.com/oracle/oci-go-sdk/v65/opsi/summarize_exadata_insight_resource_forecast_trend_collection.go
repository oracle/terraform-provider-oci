// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.oracle.com/iaas/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SummarizeExadataInsightResourceForecastTrendCollection Usage and Forecast results with breakdown by databases, hosts or storage servers.
type SummarizeExadataInsightResourceForecastTrendCollection struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata insight.
	ExadataInsightId *string `mandatory:"true" json:"exadataInsightId"`

	// Defines the resource type for an exadata  (example: DATABASE, STORAGE_SERVER, HOST, DISKGROUP)
	ExadataResourceType SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceTypeEnum `mandatory:"true" json:"exadataResourceType"`

	// Defines the type of exadata resource metric (example: CPU, STORAGE)
	ExadataResourceMetric SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricEnum `mandatory:"true" json:"exadataResourceMetric"`

	// The start timestamp that was passed into the request.
	TimeIntervalStart *common.SDKTime `mandatory:"true" json:"timeIntervalStart"`

	// The end timestamp that was passed into the request.
	TimeIntervalEnd *common.SDKTime `mandatory:"true" json:"timeIntervalEnd"`

	// Displays usage unit ( CORES, GB , PERCENT, MBPS)
	UsageUnit SummarizeExadataInsightResourceForecastTrendCollectionUsageUnitEnum `mandatory:"true" json:"usageUnit"`

	// Collection of id, name , daysToReach Capacity, historical usage and projected usage forecast.
	Items []ExadataInsightResourceForecastTrendSummary `mandatory:"true" json:"items"`
}

func (m SummarizeExadataInsightResourceForecastTrendCollection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SummarizeExadataInsightResourceForecastTrendCollection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSummarizeExadataInsightResourceForecastTrendCollectionExadataResourceTypeEnum(string(m.ExadataResourceType)); !ok && m.ExadataResourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ExadataResourceType: %s. Supported values are: %s.", m.ExadataResourceType, strings.Join(GetSummarizeExadataInsightResourceForecastTrendCollectionExadataResourceTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricEnum(string(m.ExadataResourceMetric)); !ok && m.ExadataResourceMetric != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ExadataResourceMetric: %s. Supported values are: %s.", m.ExadataResourceMetric, strings.Join(GetSummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeExadataInsightResourceForecastTrendCollectionUsageUnitEnum(string(m.UsageUnit)); !ok && m.UsageUnit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UsageUnit: %s. Supported values are: %s.", m.UsageUnit, strings.Join(GetSummarizeExadataInsightResourceForecastTrendCollectionUsageUnitEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceTypeEnum Enum with underlying type: string
type SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceTypeEnum string

// Set of constants representing the allowable values for SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceTypeEnum
const (
	SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceTypeDatabase      SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceTypeEnum = "DATABASE"
	SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceTypeHost          SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceTypeEnum = "HOST"
	SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceTypeStorageServer SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceTypeEnum = "STORAGE_SERVER"
	SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceTypeDiskgroup     SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceTypeEnum = "DISKGROUP"
)

var mappingSummarizeExadataInsightResourceForecastTrendCollectionExadataResourceTypeEnum = map[string]SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceTypeEnum{
	"DATABASE":       SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceTypeDatabase,
	"HOST":           SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceTypeHost,
	"STORAGE_SERVER": SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceTypeStorageServer,
	"DISKGROUP":      SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceTypeDiskgroup,
}

var mappingSummarizeExadataInsightResourceForecastTrendCollectionExadataResourceTypeEnumLowerCase = map[string]SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceTypeEnum{
	"database":       SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceTypeDatabase,
	"host":           SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceTypeHost,
	"storage_server": SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceTypeStorageServer,
	"diskgroup":      SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceTypeDiskgroup,
}

// GetSummarizeExadataInsightResourceForecastTrendCollectionExadataResourceTypeEnumValues Enumerates the set of values for SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceTypeEnum
func GetSummarizeExadataInsightResourceForecastTrendCollectionExadataResourceTypeEnumValues() []SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceTypeEnum {
	values := make([]SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceTypeEnum, 0)
	for _, v := range mappingSummarizeExadataInsightResourceForecastTrendCollectionExadataResourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeExadataInsightResourceForecastTrendCollectionExadataResourceTypeEnumStringValues Enumerates the set of values in String for SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceTypeEnum
func GetSummarizeExadataInsightResourceForecastTrendCollectionExadataResourceTypeEnumStringValues() []string {
	return []string{
		"DATABASE",
		"HOST",
		"STORAGE_SERVER",
		"DISKGROUP",
	}
}

// GetMappingSummarizeExadataInsightResourceForecastTrendCollectionExadataResourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeExadataInsightResourceForecastTrendCollectionExadataResourceTypeEnum(val string) (SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceTypeEnum, bool) {
	enum, ok := mappingSummarizeExadataInsightResourceForecastTrendCollectionExadataResourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricEnum Enum with underlying type: string
type SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricEnum string

// Set of constants representing the allowable values for SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricEnum
const (
	SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricCpu        SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricEnum = "CPU"
	SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricStorage    SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricEnum = "STORAGE"
	SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricIo         SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricEnum = "IO"
	SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricMemory     SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricEnum = "MEMORY"
	SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricIops       SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricEnum = "IOPS"
	SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricThroughput SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricEnum = "THROUGHPUT"
)

var mappingSummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricEnum = map[string]SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricEnum{
	"CPU":        SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricCpu,
	"STORAGE":    SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricStorage,
	"IO":         SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricIo,
	"MEMORY":     SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricMemory,
	"IOPS":       SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricIops,
	"THROUGHPUT": SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricThroughput,
}

var mappingSummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricEnumLowerCase = map[string]SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricEnum{
	"cpu":        SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricCpu,
	"storage":    SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricStorage,
	"io":         SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricIo,
	"memory":     SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricMemory,
	"iops":       SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricIops,
	"throughput": SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricThroughput,
}

// GetSummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricEnumValues Enumerates the set of values for SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricEnum
func GetSummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricEnumValues() []SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricEnum {
	values := make([]SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricEnum, 0)
	for _, v := range mappingSummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricEnumStringValues Enumerates the set of values in String for SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricEnum
func GetSummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricEnumStringValues() []string {
	return []string{
		"CPU",
		"STORAGE",
		"IO",
		"MEMORY",
		"IOPS",
		"THROUGHPUT",
	}
}

// GetMappingSummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricEnum(val string) (SummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricEnum, bool) {
	enum, ok := mappingSummarizeExadataInsightResourceForecastTrendCollectionExadataResourceMetricEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeExadataInsightResourceForecastTrendCollectionUsageUnitEnum Enum with underlying type: string
type SummarizeExadataInsightResourceForecastTrendCollectionUsageUnitEnum string

// Set of constants representing the allowable values for SummarizeExadataInsightResourceForecastTrendCollectionUsageUnitEnum
const (
	SummarizeExadataInsightResourceForecastTrendCollectionUsageUnitCores   SummarizeExadataInsightResourceForecastTrendCollectionUsageUnitEnum = "CORES"
	SummarizeExadataInsightResourceForecastTrendCollectionUsageUnitGb      SummarizeExadataInsightResourceForecastTrendCollectionUsageUnitEnum = "GB"
	SummarizeExadataInsightResourceForecastTrendCollectionUsageUnitMbps    SummarizeExadataInsightResourceForecastTrendCollectionUsageUnitEnum = "MBPS"
	SummarizeExadataInsightResourceForecastTrendCollectionUsageUnitIops    SummarizeExadataInsightResourceForecastTrendCollectionUsageUnitEnum = "IOPS"
	SummarizeExadataInsightResourceForecastTrendCollectionUsageUnitPercent SummarizeExadataInsightResourceForecastTrendCollectionUsageUnitEnum = "PERCENT"
)

var mappingSummarizeExadataInsightResourceForecastTrendCollectionUsageUnitEnum = map[string]SummarizeExadataInsightResourceForecastTrendCollectionUsageUnitEnum{
	"CORES":   SummarizeExadataInsightResourceForecastTrendCollectionUsageUnitCores,
	"GB":      SummarizeExadataInsightResourceForecastTrendCollectionUsageUnitGb,
	"MBPS":    SummarizeExadataInsightResourceForecastTrendCollectionUsageUnitMbps,
	"IOPS":    SummarizeExadataInsightResourceForecastTrendCollectionUsageUnitIops,
	"PERCENT": SummarizeExadataInsightResourceForecastTrendCollectionUsageUnitPercent,
}

var mappingSummarizeExadataInsightResourceForecastTrendCollectionUsageUnitEnumLowerCase = map[string]SummarizeExadataInsightResourceForecastTrendCollectionUsageUnitEnum{
	"cores":   SummarizeExadataInsightResourceForecastTrendCollectionUsageUnitCores,
	"gb":      SummarizeExadataInsightResourceForecastTrendCollectionUsageUnitGb,
	"mbps":    SummarizeExadataInsightResourceForecastTrendCollectionUsageUnitMbps,
	"iops":    SummarizeExadataInsightResourceForecastTrendCollectionUsageUnitIops,
	"percent": SummarizeExadataInsightResourceForecastTrendCollectionUsageUnitPercent,
}

// GetSummarizeExadataInsightResourceForecastTrendCollectionUsageUnitEnumValues Enumerates the set of values for SummarizeExadataInsightResourceForecastTrendCollectionUsageUnitEnum
func GetSummarizeExadataInsightResourceForecastTrendCollectionUsageUnitEnumValues() []SummarizeExadataInsightResourceForecastTrendCollectionUsageUnitEnum {
	values := make([]SummarizeExadataInsightResourceForecastTrendCollectionUsageUnitEnum, 0)
	for _, v := range mappingSummarizeExadataInsightResourceForecastTrendCollectionUsageUnitEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeExadataInsightResourceForecastTrendCollectionUsageUnitEnumStringValues Enumerates the set of values in String for SummarizeExadataInsightResourceForecastTrendCollectionUsageUnitEnum
func GetSummarizeExadataInsightResourceForecastTrendCollectionUsageUnitEnumStringValues() []string {
	return []string{
		"CORES",
		"GB",
		"MBPS",
		"IOPS",
		"PERCENT",
	}
}

// GetMappingSummarizeExadataInsightResourceForecastTrendCollectionUsageUnitEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeExadataInsightResourceForecastTrendCollectionUsageUnitEnum(val string) (SummarizeExadataInsightResourceForecastTrendCollectionUsageUnitEnum, bool) {
	enum, ok := mappingSummarizeExadataInsightResourceForecastTrendCollectionUsageUnitEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
