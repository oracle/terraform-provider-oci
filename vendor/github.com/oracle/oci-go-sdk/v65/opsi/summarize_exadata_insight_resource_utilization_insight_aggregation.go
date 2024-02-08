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

// SummarizeExadataInsightResourceUtilizationInsightAggregation Insights response containing utilization values for exadata systems.
type SummarizeExadataInsightResourceUtilizationInsightAggregation struct {

	// The start timestamp that was passed into the request.
	TimeIntervalStart *common.SDKTime `mandatory:"true" json:"timeIntervalStart"`

	// The end timestamp that was passed into the request.
	TimeIntervalEnd *common.SDKTime `mandatory:"true" json:"timeIntervalEnd"`

	// Defines the type of exadata resource metric (example: CPU, STORAGE)
	ExadataResourceMetric SummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceMetricEnum `mandatory:"true" json:"exadataResourceMetric"`

	// Defines the resource type for an exadata  (example: DATABASE, STORAGE_SERVER, HOST, DISKGROUP)
	ExadataResourceType SummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceTypeEnum `mandatory:"true" json:"exadataResourceType"`

	// Collection of Exadata system utilization
	Utilization []ExadataInsightResourceInsightUtilizationItem `mandatory:"true" json:"utilization"`
}

func (m SummarizeExadataInsightResourceUtilizationInsightAggregation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SummarizeExadataInsightResourceUtilizationInsightAggregation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceMetricEnum(string(m.ExadataResourceMetric)); !ok && m.ExadataResourceMetric != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ExadataResourceMetric: %s. Supported values are: %s.", m.ExadataResourceMetric, strings.Join(GetSummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceMetricEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceTypeEnum(string(m.ExadataResourceType)); !ok && m.ExadataResourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ExadataResourceType: %s. Supported values are: %s.", m.ExadataResourceType, strings.Join(GetSummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceMetricEnum Enum with underlying type: string
type SummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceMetricEnum string

// Set of constants representing the allowable values for SummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceMetricEnum
const (
	SummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceMetricCpu        SummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceMetricEnum = "CPU"
	SummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceMetricStorage    SummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceMetricEnum = "STORAGE"
	SummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceMetricIo         SummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceMetricEnum = "IO"
	SummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceMetricMemory     SummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceMetricEnum = "MEMORY"
	SummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceMetricIops       SummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceMetricEnum = "IOPS"
	SummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceMetricThroughput SummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceMetricEnum = "THROUGHPUT"
)

var mappingSummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceMetricEnum = map[string]SummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceMetricEnum{
	"CPU":        SummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceMetricCpu,
	"STORAGE":    SummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceMetricStorage,
	"IO":         SummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceMetricIo,
	"MEMORY":     SummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceMetricMemory,
	"IOPS":       SummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceMetricIops,
	"THROUGHPUT": SummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceMetricThroughput,
}

var mappingSummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceMetricEnumLowerCase = map[string]SummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceMetricEnum{
	"cpu":        SummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceMetricCpu,
	"storage":    SummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceMetricStorage,
	"io":         SummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceMetricIo,
	"memory":     SummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceMetricMemory,
	"iops":       SummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceMetricIops,
	"throughput": SummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceMetricThroughput,
}

// GetSummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceMetricEnumValues Enumerates the set of values for SummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceMetricEnum
func GetSummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceMetricEnumValues() []SummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceMetricEnum {
	values := make([]SummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceMetricEnum, 0)
	for _, v := range mappingSummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceMetricEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceMetricEnumStringValues Enumerates the set of values in String for SummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceMetricEnum
func GetSummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceMetricEnumStringValues() []string {
	return []string{
		"CPU",
		"STORAGE",
		"IO",
		"MEMORY",
		"IOPS",
		"THROUGHPUT",
	}
}

// GetMappingSummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceMetricEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceMetricEnum(val string) (SummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceMetricEnum, bool) {
	enum, ok := mappingSummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceMetricEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceTypeEnum Enum with underlying type: string
type SummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceTypeEnum string

// Set of constants representing the allowable values for SummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceTypeEnum
const (
	SummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceTypeDatabase      SummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceTypeEnum = "DATABASE"
	SummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceTypeHost          SummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceTypeEnum = "HOST"
	SummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceTypeStorageServer SummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceTypeEnum = "STORAGE_SERVER"
	SummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceTypeDiskgroup     SummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceTypeEnum = "DISKGROUP"
)

var mappingSummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceTypeEnum = map[string]SummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceTypeEnum{
	"DATABASE":       SummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceTypeDatabase,
	"HOST":           SummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceTypeHost,
	"STORAGE_SERVER": SummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceTypeStorageServer,
	"DISKGROUP":      SummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceTypeDiskgroup,
}

var mappingSummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceTypeEnumLowerCase = map[string]SummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceTypeEnum{
	"database":       SummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceTypeDatabase,
	"host":           SummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceTypeHost,
	"storage_server": SummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceTypeStorageServer,
	"diskgroup":      SummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceTypeDiskgroup,
}

// GetSummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceTypeEnumValues Enumerates the set of values for SummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceTypeEnum
func GetSummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceTypeEnumValues() []SummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceTypeEnum {
	values := make([]SummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceTypeEnum, 0)
	for _, v := range mappingSummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceTypeEnumStringValues Enumerates the set of values in String for SummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceTypeEnum
func GetSummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceTypeEnumStringValues() []string {
	return []string{
		"DATABASE",
		"HOST",
		"STORAGE_SERVER",
		"DISKGROUP",
	}
}

// GetMappingSummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceTypeEnum(val string) (SummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceTypeEnum, bool) {
	enum, ok := mappingSummarizeExadataInsightResourceUtilizationInsightAggregationExadataResourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
