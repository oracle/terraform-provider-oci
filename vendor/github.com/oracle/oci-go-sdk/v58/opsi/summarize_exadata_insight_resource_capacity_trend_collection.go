// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// SummarizeExadataInsightResourceCapacityTrendCollection capacity results with breakdown by databases, hosts, storage servers or diskgroup.
type SummarizeExadataInsightResourceCapacityTrendCollection struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata insight.
	ExadataInsightId *string `mandatory:"true" json:"exadataInsightId"`

	// Defines the resource type for an exadata  (example: DATABASE, STORAGE_SERVER, HOST, DISKGROUP)
	ExadataResourceType SummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceTypeEnum `mandatory:"true" json:"exadataResourceType"`

	// Defines the type of exadata resource metric (example: CPU, STORAGE)
	ExadataResourceMetric SummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceMetricEnum `mandatory:"true" json:"exadataResourceMetric"`

	// The start timestamp that was passed into the request.
	TimeIntervalStart *common.SDKTime `mandatory:"true" json:"timeIntervalStart"`

	// The end timestamp that was passed into the request.
	TimeIntervalEnd *common.SDKTime `mandatory:"true" json:"timeIntervalEnd"`

	// Displays usage unit ( CORES, GB)
	UsageUnit UsageUnitEnum `mandatory:"true" json:"usageUnit"`

	// Capacity Data with time interval
	Items []ExadataInsightResourceCapacityTrendSummary `mandatory:"true" json:"items"`
}

func (m SummarizeExadataInsightResourceCapacityTrendCollection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SummarizeExadataInsightResourceCapacityTrendCollection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceTypeEnum(string(m.ExadataResourceType)); !ok && m.ExadataResourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ExadataResourceType: %s. Supported values are: %s.", m.ExadataResourceType, strings.Join(GetSummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceMetricEnum(string(m.ExadataResourceMetric)); !ok && m.ExadataResourceMetric != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ExadataResourceMetric: %s. Supported values are: %s.", m.ExadataResourceMetric, strings.Join(GetSummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceMetricEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUsageUnitEnum(string(m.UsageUnit)); !ok && m.UsageUnit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UsageUnit: %s. Supported values are: %s.", m.UsageUnit, strings.Join(GetUsageUnitEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceTypeEnum Enum with underlying type: string
type SummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceTypeEnum string

// Set of constants representing the allowable values for SummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceTypeEnum
const (
	SummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceTypeDatabase      SummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceTypeEnum = "DATABASE"
	SummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceTypeHost          SummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceTypeEnum = "HOST"
	SummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceTypeStorageServer SummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceTypeEnum = "STORAGE_SERVER"
	SummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceTypeDiskgroup     SummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceTypeEnum = "DISKGROUP"
)

var mappingSummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceTypeEnum = map[string]SummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceTypeEnum{
	"DATABASE":       SummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceTypeDatabase,
	"HOST":           SummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceTypeHost,
	"STORAGE_SERVER": SummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceTypeStorageServer,
	"DISKGROUP":      SummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceTypeDiskgroup,
}

// GetSummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceTypeEnumValues Enumerates the set of values for SummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceTypeEnum
func GetSummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceTypeEnumValues() []SummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceTypeEnum {
	values := make([]SummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceTypeEnum, 0)
	for _, v := range mappingSummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceTypeEnumStringValues Enumerates the set of values in String for SummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceTypeEnum
func GetSummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceTypeEnumStringValues() []string {
	return []string{
		"DATABASE",
		"HOST",
		"STORAGE_SERVER",
		"DISKGROUP",
	}
}

// GetMappingSummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceTypeEnum(val string) (SummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceTypeEnum, bool) {
	mappingSummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceTypeEnumIgnoreCase := make(map[string]SummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceTypeEnum)
	for k, v := range mappingSummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceTypeEnum {
		mappingSummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingSummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceMetricEnum Enum with underlying type: string
type SummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceMetricEnum string

// Set of constants representing the allowable values for SummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceMetricEnum
const (
	SummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceMetricCpu        SummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceMetricEnum = "CPU"
	SummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceMetricStorage    SummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceMetricEnum = "STORAGE"
	SummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceMetricIo         SummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceMetricEnum = "IO"
	SummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceMetricMemory     SummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceMetricEnum = "MEMORY"
	SummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceMetricIops       SummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceMetricEnum = "IOPS"
	SummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceMetricThroughput SummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceMetricEnum = "THROUGHPUT"
)

var mappingSummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceMetricEnum = map[string]SummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceMetricEnum{
	"CPU":        SummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceMetricCpu,
	"STORAGE":    SummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceMetricStorage,
	"IO":         SummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceMetricIo,
	"MEMORY":     SummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceMetricMemory,
	"IOPS":       SummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceMetricIops,
	"THROUGHPUT": SummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceMetricThroughput,
}

// GetSummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceMetricEnumValues Enumerates the set of values for SummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceMetricEnum
func GetSummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceMetricEnumValues() []SummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceMetricEnum {
	values := make([]SummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceMetricEnum, 0)
	for _, v := range mappingSummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceMetricEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceMetricEnumStringValues Enumerates the set of values in String for SummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceMetricEnum
func GetSummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceMetricEnumStringValues() []string {
	return []string{
		"CPU",
		"STORAGE",
		"IO",
		"MEMORY",
		"IOPS",
		"THROUGHPUT",
	}
}

// GetMappingSummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceMetricEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceMetricEnum(val string) (SummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceMetricEnum, bool) {
	mappingSummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceMetricEnumIgnoreCase := make(map[string]SummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceMetricEnum)
	for k, v := range mappingSummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceMetricEnum {
		mappingSummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceMetricEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingSummarizeExadataInsightResourceCapacityTrendCollectionExadataResourceMetricEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
