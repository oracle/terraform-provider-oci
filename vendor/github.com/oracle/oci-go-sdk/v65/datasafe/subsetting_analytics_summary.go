// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// SubsettingAnalyticsSummary Summary of subsetting analytics data
type SubsettingAnalyticsSummary struct {

	// The name of the aggregation metric
	MetricName SubsettingAnalyticsSummaryMetricNameEnum `mandatory:"true" json:"metricName"`

	// The total count for the aggregation metric
	Count *int64 `mandatory:"true" json:"count"`

	Dimensions *SubsettingAnalyticsDimensions `mandatory:"false" json:"dimensions"`

	// The date and time the target database was last subsetted using a subsetting policy, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339)
	TimeLastSubsetted *common.SDKTime `mandatory:"false" json:"timeLastSubsetted"`
}

func (m SubsettingAnalyticsSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SubsettingAnalyticsSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSubsettingAnalyticsSummaryMetricNameEnum(string(m.MetricName)); !ok && m.MetricName != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MetricName: %s. Supported values are: %s.", m.MetricName, strings.Join(GetSubsettingAnalyticsSummaryMetricNameEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SubsettingAnalyticsSummaryMetricNameEnum Enum with underlying type: string
type SubsettingAnalyticsSummaryMetricNameEnum string

// Set of constants representing the allowable values for SubsettingAnalyticsSummaryMetricNameEnum
const (
	SubsettingAnalyticsSummaryMetricNameSubsettingPolicy           SubsettingAnalyticsSummaryMetricNameEnum = "SUBSETTING_POLICY"
	SubsettingAnalyticsSummaryMetricNameSubsettingDatabase         SubsettingAnalyticsSummaryMetricNameEnum = "SUBSETTING_DATABASE"
	SubsettingAnalyticsSummaryMetricNameSubsettingWorkRequest      SubsettingAnalyticsSummaryMetricNameEnum = "SUBSETTING_WORK_REQUEST"
	SubsettingAnalyticsSummaryMetricNameSubsettedSchema            SubsettingAnalyticsSummaryMetricNameEnum = "SUBSETTED_SCHEMA"
	SubsettingAnalyticsSummaryMetricNameSubsettedTable             SubsettingAnalyticsSummaryMetricNameEnum = "SUBSETTED_TABLE"
	SubsettingAnalyticsSummaryMetricNameReducedRowCount            SubsettingAnalyticsSummaryMetricNameEnum = "REDUCED_ROW_COUNT"
	SubsettingAnalyticsSummaryMetricNameReducedSizeInBytes         SubsettingAnalyticsSummaryMetricNameEnum = "REDUCED_SIZE_IN_BYTES"
	SubsettingAnalyticsSummaryMetricNameReducedPercent             SubsettingAnalyticsSummaryMetricNameEnum = "REDUCED_PERCENT"
	SubsettingAnalyticsSummaryMetricNameOriginalSizeInBytes        SubsettingAnalyticsSummaryMetricNameEnum = "ORIGINAL_SIZE_IN_BYTES"
	SubsettingAnalyticsSummaryMetricNameSizeAfterSubsettingInBytes SubsettingAnalyticsSummaryMetricNameEnum = "SIZE_AFTER_SUBSETTING_IN_BYTES"
)

var mappingSubsettingAnalyticsSummaryMetricNameEnum = map[string]SubsettingAnalyticsSummaryMetricNameEnum{
	"SUBSETTING_POLICY":              SubsettingAnalyticsSummaryMetricNameSubsettingPolicy,
	"SUBSETTING_DATABASE":            SubsettingAnalyticsSummaryMetricNameSubsettingDatabase,
	"SUBSETTING_WORK_REQUEST":        SubsettingAnalyticsSummaryMetricNameSubsettingWorkRequest,
	"SUBSETTED_SCHEMA":               SubsettingAnalyticsSummaryMetricNameSubsettedSchema,
	"SUBSETTED_TABLE":                SubsettingAnalyticsSummaryMetricNameSubsettedTable,
	"REDUCED_ROW_COUNT":              SubsettingAnalyticsSummaryMetricNameReducedRowCount,
	"REDUCED_SIZE_IN_BYTES":          SubsettingAnalyticsSummaryMetricNameReducedSizeInBytes,
	"REDUCED_PERCENT":                SubsettingAnalyticsSummaryMetricNameReducedPercent,
	"ORIGINAL_SIZE_IN_BYTES":         SubsettingAnalyticsSummaryMetricNameOriginalSizeInBytes,
	"SIZE_AFTER_SUBSETTING_IN_BYTES": SubsettingAnalyticsSummaryMetricNameSizeAfterSubsettingInBytes,
}

var mappingSubsettingAnalyticsSummaryMetricNameEnumLowerCase = map[string]SubsettingAnalyticsSummaryMetricNameEnum{
	"subsetting_policy":              SubsettingAnalyticsSummaryMetricNameSubsettingPolicy,
	"subsetting_database":            SubsettingAnalyticsSummaryMetricNameSubsettingDatabase,
	"subsetting_work_request":        SubsettingAnalyticsSummaryMetricNameSubsettingWorkRequest,
	"subsetted_schema":               SubsettingAnalyticsSummaryMetricNameSubsettedSchema,
	"subsetted_table":                SubsettingAnalyticsSummaryMetricNameSubsettedTable,
	"reduced_row_count":              SubsettingAnalyticsSummaryMetricNameReducedRowCount,
	"reduced_size_in_bytes":          SubsettingAnalyticsSummaryMetricNameReducedSizeInBytes,
	"reduced_percent":                SubsettingAnalyticsSummaryMetricNameReducedPercent,
	"original_size_in_bytes":         SubsettingAnalyticsSummaryMetricNameOriginalSizeInBytes,
	"size_after_subsetting_in_bytes": SubsettingAnalyticsSummaryMetricNameSizeAfterSubsettingInBytes,
}

// GetSubsettingAnalyticsSummaryMetricNameEnumValues Enumerates the set of values for SubsettingAnalyticsSummaryMetricNameEnum
func GetSubsettingAnalyticsSummaryMetricNameEnumValues() []SubsettingAnalyticsSummaryMetricNameEnum {
	values := make([]SubsettingAnalyticsSummaryMetricNameEnum, 0)
	for _, v := range mappingSubsettingAnalyticsSummaryMetricNameEnum {
		values = append(values, v)
	}
	return values
}

// GetSubsettingAnalyticsSummaryMetricNameEnumStringValues Enumerates the set of values in String for SubsettingAnalyticsSummaryMetricNameEnum
func GetSubsettingAnalyticsSummaryMetricNameEnumStringValues() []string {
	return []string{
		"SUBSETTING_POLICY",
		"SUBSETTING_DATABASE",
		"SUBSETTING_WORK_REQUEST",
		"SUBSETTED_SCHEMA",
		"SUBSETTED_TABLE",
		"REDUCED_ROW_COUNT",
		"REDUCED_SIZE_IN_BYTES",
		"REDUCED_PERCENT",
		"ORIGINAL_SIZE_IN_BYTES",
		"SIZE_AFTER_SUBSETTING_IN_BYTES",
	}
}

// GetMappingSubsettingAnalyticsSummaryMetricNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSubsettingAnalyticsSummaryMetricNameEnum(val string) (SubsettingAnalyticsSummaryMetricNameEnum, bool) {
	enum, ok := mappingSubsettingAnalyticsSummaryMetricNameEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
