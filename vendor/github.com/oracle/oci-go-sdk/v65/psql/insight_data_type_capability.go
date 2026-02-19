// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// PGSQL Control Plane API
//
// Use the OCI Database with PostgreSQL API to manage resources such as database systems, database nodes, backups, and configurations.
// For information, see the user guide documentation for the service (https://docs.oracle.com/iaas/Content/postgresql/home.htm).
//

package psql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InsightDataTypeCapability Capability metadata for a specific insight data type.
type InsightDataTypeCapability struct {

	// Insight data type identifier (for example, AAS_TIME_SERIES).
	InsightDataType InsightDataTypeCapabilityInsightDataTypeEnum `mandatory:"true" json:"insightDataType"`

	DataContract *InsightDataContract `mandatory:"true" json:"dataContract"`

	// Human-readable description of the insight data type.
	Description *string `mandatory:"false" json:"description"`

	DateTimeRangeSupport *DateTimeRangeCapability `mandatory:"false" json:"dateTimeRangeSupport"`

	Granularity *GranularityCapability `mandatory:"false" json:"granularity"`

	// Supported filters for this insight data type.
	Filters []InsightFilterCapability `mandatory:"false" json:"filters"`

	Sorting *SortingCapability `mandatory:"false" json:"sorting"`

	Pagination *PaginationCapability `mandatory:"false" json:"pagination"`

	Limits *InsightLimits `mandatory:"false" json:"limits"`
}

func (m InsightDataTypeCapability) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InsightDataTypeCapability) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingInsightDataTypeCapabilityInsightDataTypeEnum(string(m.InsightDataType)); !ok && m.InsightDataType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for InsightDataType: %s. Supported values are: %s.", m.InsightDataType, strings.Join(GetInsightDataTypeCapabilityInsightDataTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// InsightDataTypeCapabilityInsightDataTypeEnum Enum with underlying type: string
type InsightDataTypeCapabilityInsightDataTypeEnum string

// Set of constants representing the allowable values for InsightDataTypeCapabilityInsightDataTypeEnum
const (
	InsightDataTypeCapabilityInsightDataTypeAasTimeSeries InsightDataTypeCapabilityInsightDataTypeEnum = "AAS_TIME_SERIES"
	InsightDataTypeCapabilityInsightDataTypeTopQueries    InsightDataTypeCapabilityInsightDataTypeEnum = "TOP_QUERIES"
)

var mappingInsightDataTypeCapabilityInsightDataTypeEnum = map[string]InsightDataTypeCapabilityInsightDataTypeEnum{
	"AAS_TIME_SERIES": InsightDataTypeCapabilityInsightDataTypeAasTimeSeries,
	"TOP_QUERIES":     InsightDataTypeCapabilityInsightDataTypeTopQueries,
}

var mappingInsightDataTypeCapabilityInsightDataTypeEnumLowerCase = map[string]InsightDataTypeCapabilityInsightDataTypeEnum{
	"aas_time_series": InsightDataTypeCapabilityInsightDataTypeAasTimeSeries,
	"top_queries":     InsightDataTypeCapabilityInsightDataTypeTopQueries,
}

// GetInsightDataTypeCapabilityInsightDataTypeEnumValues Enumerates the set of values for InsightDataTypeCapabilityInsightDataTypeEnum
func GetInsightDataTypeCapabilityInsightDataTypeEnumValues() []InsightDataTypeCapabilityInsightDataTypeEnum {
	values := make([]InsightDataTypeCapabilityInsightDataTypeEnum, 0)
	for _, v := range mappingInsightDataTypeCapabilityInsightDataTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetInsightDataTypeCapabilityInsightDataTypeEnumStringValues Enumerates the set of values in String for InsightDataTypeCapabilityInsightDataTypeEnum
func GetInsightDataTypeCapabilityInsightDataTypeEnumStringValues() []string {
	return []string{
		"AAS_TIME_SERIES",
		"TOP_QUERIES",
	}
}

// GetMappingInsightDataTypeCapabilityInsightDataTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInsightDataTypeCapabilityInsightDataTypeEnum(val string) (InsightDataTypeCapabilityInsightDataTypeEnum, bool) {
	enum, ok := mappingInsightDataTypeCapabilityInsightDataTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
