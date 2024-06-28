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

// SummarizeExadataInsightResourceUsageCollection Resource usage , allocation, utilization and usage ChangePercent for the current time period
type SummarizeExadataInsightResourceUsageCollection struct {

	// The start timestamp that was passed into the request.
	TimeIntervalStart *common.SDKTime `mandatory:"true" json:"timeIntervalStart"`

	// The end timestamp that was passed into the request.
	TimeIntervalEnd *common.SDKTime `mandatory:"true" json:"timeIntervalEnd"`

	// Defines the type of exadata resource metric (example: CPU, STORAGE)
	ExadataResourceMetric SummarizeExadataInsightResourceUsageCollectionExadataResourceMetricEnum `mandatory:"true" json:"exadataResourceMetric"`

	// Defines the resource type for an exadata  (example: DATABASE, STORAGE_SERVER, HOST, DISKGROUP)
	ExadataResourceType SummarizeExadataInsightResourceUsageCollectionExadataResourceTypeEnum `mandatory:"true" json:"exadataResourceType"`

	// Displays usage unit ( CORES, GB , PERCENT, MBPS)
	UsageUnit SummarizeExadataInsightResourceUsageCollectionUsageUnitEnum `mandatory:"true" json:"usageUnit"`

	// Collection of Resource Usage Summary items
	Items []ResourceUsageSummary `mandatory:"true" json:"items"`
}

func (m SummarizeExadataInsightResourceUsageCollection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SummarizeExadataInsightResourceUsageCollection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSummarizeExadataInsightResourceUsageCollectionExadataResourceMetricEnum(string(m.ExadataResourceMetric)); !ok && m.ExadataResourceMetric != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ExadataResourceMetric: %s. Supported values are: %s.", m.ExadataResourceMetric, strings.Join(GetSummarizeExadataInsightResourceUsageCollectionExadataResourceMetricEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeExadataInsightResourceUsageCollectionExadataResourceTypeEnum(string(m.ExadataResourceType)); !ok && m.ExadataResourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ExadataResourceType: %s. Supported values are: %s.", m.ExadataResourceType, strings.Join(GetSummarizeExadataInsightResourceUsageCollectionExadataResourceTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeExadataInsightResourceUsageCollectionUsageUnitEnum(string(m.UsageUnit)); !ok && m.UsageUnit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UsageUnit: %s. Supported values are: %s.", m.UsageUnit, strings.Join(GetSummarizeExadataInsightResourceUsageCollectionUsageUnitEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeExadataInsightResourceUsageCollectionExadataResourceMetricEnum Enum with underlying type: string
type SummarizeExadataInsightResourceUsageCollectionExadataResourceMetricEnum string

// Set of constants representing the allowable values for SummarizeExadataInsightResourceUsageCollectionExadataResourceMetricEnum
const (
	SummarizeExadataInsightResourceUsageCollectionExadataResourceMetricCpu        SummarizeExadataInsightResourceUsageCollectionExadataResourceMetricEnum = "CPU"
	SummarizeExadataInsightResourceUsageCollectionExadataResourceMetricStorage    SummarizeExadataInsightResourceUsageCollectionExadataResourceMetricEnum = "STORAGE"
	SummarizeExadataInsightResourceUsageCollectionExadataResourceMetricIo         SummarizeExadataInsightResourceUsageCollectionExadataResourceMetricEnum = "IO"
	SummarizeExadataInsightResourceUsageCollectionExadataResourceMetricMemory     SummarizeExadataInsightResourceUsageCollectionExadataResourceMetricEnum = "MEMORY"
	SummarizeExadataInsightResourceUsageCollectionExadataResourceMetricIops       SummarizeExadataInsightResourceUsageCollectionExadataResourceMetricEnum = "IOPS"
	SummarizeExadataInsightResourceUsageCollectionExadataResourceMetricThroughput SummarizeExadataInsightResourceUsageCollectionExadataResourceMetricEnum = "THROUGHPUT"
)

var mappingSummarizeExadataInsightResourceUsageCollectionExadataResourceMetricEnum = map[string]SummarizeExadataInsightResourceUsageCollectionExadataResourceMetricEnum{
	"CPU":        SummarizeExadataInsightResourceUsageCollectionExadataResourceMetricCpu,
	"STORAGE":    SummarizeExadataInsightResourceUsageCollectionExadataResourceMetricStorage,
	"IO":         SummarizeExadataInsightResourceUsageCollectionExadataResourceMetricIo,
	"MEMORY":     SummarizeExadataInsightResourceUsageCollectionExadataResourceMetricMemory,
	"IOPS":       SummarizeExadataInsightResourceUsageCollectionExadataResourceMetricIops,
	"THROUGHPUT": SummarizeExadataInsightResourceUsageCollectionExadataResourceMetricThroughput,
}

var mappingSummarizeExadataInsightResourceUsageCollectionExadataResourceMetricEnumLowerCase = map[string]SummarizeExadataInsightResourceUsageCollectionExadataResourceMetricEnum{
	"cpu":        SummarizeExadataInsightResourceUsageCollectionExadataResourceMetricCpu,
	"storage":    SummarizeExadataInsightResourceUsageCollectionExadataResourceMetricStorage,
	"io":         SummarizeExadataInsightResourceUsageCollectionExadataResourceMetricIo,
	"memory":     SummarizeExadataInsightResourceUsageCollectionExadataResourceMetricMemory,
	"iops":       SummarizeExadataInsightResourceUsageCollectionExadataResourceMetricIops,
	"throughput": SummarizeExadataInsightResourceUsageCollectionExadataResourceMetricThroughput,
}

// GetSummarizeExadataInsightResourceUsageCollectionExadataResourceMetricEnumValues Enumerates the set of values for SummarizeExadataInsightResourceUsageCollectionExadataResourceMetricEnum
func GetSummarizeExadataInsightResourceUsageCollectionExadataResourceMetricEnumValues() []SummarizeExadataInsightResourceUsageCollectionExadataResourceMetricEnum {
	values := make([]SummarizeExadataInsightResourceUsageCollectionExadataResourceMetricEnum, 0)
	for _, v := range mappingSummarizeExadataInsightResourceUsageCollectionExadataResourceMetricEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeExadataInsightResourceUsageCollectionExadataResourceMetricEnumStringValues Enumerates the set of values in String for SummarizeExadataInsightResourceUsageCollectionExadataResourceMetricEnum
func GetSummarizeExadataInsightResourceUsageCollectionExadataResourceMetricEnumStringValues() []string {
	return []string{
		"CPU",
		"STORAGE",
		"IO",
		"MEMORY",
		"IOPS",
		"THROUGHPUT",
	}
}

// GetMappingSummarizeExadataInsightResourceUsageCollectionExadataResourceMetricEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeExadataInsightResourceUsageCollectionExadataResourceMetricEnum(val string) (SummarizeExadataInsightResourceUsageCollectionExadataResourceMetricEnum, bool) {
	enum, ok := mappingSummarizeExadataInsightResourceUsageCollectionExadataResourceMetricEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeExadataInsightResourceUsageCollectionExadataResourceTypeEnum Enum with underlying type: string
type SummarizeExadataInsightResourceUsageCollectionExadataResourceTypeEnum string

// Set of constants representing the allowable values for SummarizeExadataInsightResourceUsageCollectionExadataResourceTypeEnum
const (
	SummarizeExadataInsightResourceUsageCollectionExadataResourceTypeDatabase      SummarizeExadataInsightResourceUsageCollectionExadataResourceTypeEnum = "DATABASE"
	SummarizeExadataInsightResourceUsageCollectionExadataResourceTypeHost          SummarizeExadataInsightResourceUsageCollectionExadataResourceTypeEnum = "HOST"
	SummarizeExadataInsightResourceUsageCollectionExadataResourceTypeStorageServer SummarizeExadataInsightResourceUsageCollectionExadataResourceTypeEnum = "STORAGE_SERVER"
	SummarizeExadataInsightResourceUsageCollectionExadataResourceTypeDiskgroup     SummarizeExadataInsightResourceUsageCollectionExadataResourceTypeEnum = "DISKGROUP"
)

var mappingSummarizeExadataInsightResourceUsageCollectionExadataResourceTypeEnum = map[string]SummarizeExadataInsightResourceUsageCollectionExadataResourceTypeEnum{
	"DATABASE":       SummarizeExadataInsightResourceUsageCollectionExadataResourceTypeDatabase,
	"HOST":           SummarizeExadataInsightResourceUsageCollectionExadataResourceTypeHost,
	"STORAGE_SERVER": SummarizeExadataInsightResourceUsageCollectionExadataResourceTypeStorageServer,
	"DISKGROUP":      SummarizeExadataInsightResourceUsageCollectionExadataResourceTypeDiskgroup,
}

var mappingSummarizeExadataInsightResourceUsageCollectionExadataResourceTypeEnumLowerCase = map[string]SummarizeExadataInsightResourceUsageCollectionExadataResourceTypeEnum{
	"database":       SummarizeExadataInsightResourceUsageCollectionExadataResourceTypeDatabase,
	"host":           SummarizeExadataInsightResourceUsageCollectionExadataResourceTypeHost,
	"storage_server": SummarizeExadataInsightResourceUsageCollectionExadataResourceTypeStorageServer,
	"diskgroup":      SummarizeExadataInsightResourceUsageCollectionExadataResourceTypeDiskgroup,
}

// GetSummarizeExadataInsightResourceUsageCollectionExadataResourceTypeEnumValues Enumerates the set of values for SummarizeExadataInsightResourceUsageCollectionExadataResourceTypeEnum
func GetSummarizeExadataInsightResourceUsageCollectionExadataResourceTypeEnumValues() []SummarizeExadataInsightResourceUsageCollectionExadataResourceTypeEnum {
	values := make([]SummarizeExadataInsightResourceUsageCollectionExadataResourceTypeEnum, 0)
	for _, v := range mappingSummarizeExadataInsightResourceUsageCollectionExadataResourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeExadataInsightResourceUsageCollectionExadataResourceTypeEnumStringValues Enumerates the set of values in String for SummarizeExadataInsightResourceUsageCollectionExadataResourceTypeEnum
func GetSummarizeExadataInsightResourceUsageCollectionExadataResourceTypeEnumStringValues() []string {
	return []string{
		"DATABASE",
		"HOST",
		"STORAGE_SERVER",
		"DISKGROUP",
	}
}

// GetMappingSummarizeExadataInsightResourceUsageCollectionExadataResourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeExadataInsightResourceUsageCollectionExadataResourceTypeEnum(val string) (SummarizeExadataInsightResourceUsageCollectionExadataResourceTypeEnum, bool) {
	enum, ok := mappingSummarizeExadataInsightResourceUsageCollectionExadataResourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeExadataInsightResourceUsageCollectionUsageUnitEnum Enum with underlying type: string
type SummarizeExadataInsightResourceUsageCollectionUsageUnitEnum string

// Set of constants representing the allowable values for SummarizeExadataInsightResourceUsageCollectionUsageUnitEnum
const (
	SummarizeExadataInsightResourceUsageCollectionUsageUnitCores   SummarizeExadataInsightResourceUsageCollectionUsageUnitEnum = "CORES"
	SummarizeExadataInsightResourceUsageCollectionUsageUnitGb      SummarizeExadataInsightResourceUsageCollectionUsageUnitEnum = "GB"
	SummarizeExadataInsightResourceUsageCollectionUsageUnitMbps    SummarizeExadataInsightResourceUsageCollectionUsageUnitEnum = "MBPS"
	SummarizeExadataInsightResourceUsageCollectionUsageUnitIops    SummarizeExadataInsightResourceUsageCollectionUsageUnitEnum = "IOPS"
	SummarizeExadataInsightResourceUsageCollectionUsageUnitPercent SummarizeExadataInsightResourceUsageCollectionUsageUnitEnum = "PERCENT"
)

var mappingSummarizeExadataInsightResourceUsageCollectionUsageUnitEnum = map[string]SummarizeExadataInsightResourceUsageCollectionUsageUnitEnum{
	"CORES":   SummarizeExadataInsightResourceUsageCollectionUsageUnitCores,
	"GB":      SummarizeExadataInsightResourceUsageCollectionUsageUnitGb,
	"MBPS":    SummarizeExadataInsightResourceUsageCollectionUsageUnitMbps,
	"IOPS":    SummarizeExadataInsightResourceUsageCollectionUsageUnitIops,
	"PERCENT": SummarizeExadataInsightResourceUsageCollectionUsageUnitPercent,
}

var mappingSummarizeExadataInsightResourceUsageCollectionUsageUnitEnumLowerCase = map[string]SummarizeExadataInsightResourceUsageCollectionUsageUnitEnum{
	"cores":   SummarizeExadataInsightResourceUsageCollectionUsageUnitCores,
	"gb":      SummarizeExadataInsightResourceUsageCollectionUsageUnitGb,
	"mbps":    SummarizeExadataInsightResourceUsageCollectionUsageUnitMbps,
	"iops":    SummarizeExadataInsightResourceUsageCollectionUsageUnitIops,
	"percent": SummarizeExadataInsightResourceUsageCollectionUsageUnitPercent,
}

// GetSummarizeExadataInsightResourceUsageCollectionUsageUnitEnumValues Enumerates the set of values for SummarizeExadataInsightResourceUsageCollectionUsageUnitEnum
func GetSummarizeExadataInsightResourceUsageCollectionUsageUnitEnumValues() []SummarizeExadataInsightResourceUsageCollectionUsageUnitEnum {
	values := make([]SummarizeExadataInsightResourceUsageCollectionUsageUnitEnum, 0)
	for _, v := range mappingSummarizeExadataInsightResourceUsageCollectionUsageUnitEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeExadataInsightResourceUsageCollectionUsageUnitEnumStringValues Enumerates the set of values in String for SummarizeExadataInsightResourceUsageCollectionUsageUnitEnum
func GetSummarizeExadataInsightResourceUsageCollectionUsageUnitEnumStringValues() []string {
	return []string{
		"CORES",
		"GB",
		"MBPS",
		"IOPS",
		"PERCENT",
	}
}

// GetMappingSummarizeExadataInsightResourceUsageCollectionUsageUnitEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeExadataInsightResourceUsageCollectionUsageUnitEnum(val string) (SummarizeExadataInsightResourceUsageCollectionUsageUnitEnum, bool) {
	enum, ok := mappingSummarizeExadataInsightResourceUsageCollectionUsageUnitEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
