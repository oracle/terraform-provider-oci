// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// MaskingAnalyticsSummary Summary of masking analytics data.
type MaskingAnalyticsSummary struct {

	// The name of the aggregation metric.
	MetricName MaskingAnalyticsSummaryMetricNameEnum `mandatory:"true" json:"metricName"`

	// The total count for the aggregation metric.
	Count *int64 `mandatory:"true" json:"count"`

	Dimensions *MaskingAnalyticsDimensions `mandatory:"false" json:"dimensions"`
}

func (m MaskingAnalyticsSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MaskingAnalyticsSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMaskingAnalyticsSummaryMetricNameEnum(string(m.MetricName)); !ok && m.MetricName != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MetricName: %s. Supported values are: %s.", m.MetricName, strings.Join(GetMaskingAnalyticsSummaryMetricNameEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MaskingAnalyticsSummaryMetricNameEnum Enum with underlying type: string
type MaskingAnalyticsSummaryMetricNameEnum string

// Set of constants representing the allowable values for MaskingAnalyticsSummaryMetricNameEnum
const (
	MaskingAnalyticsSummaryMetricNameMaskingPolicy       MaskingAnalyticsSummaryMetricNameEnum = "MASKING_POLICY"
	MaskingAnalyticsSummaryMetricNameMaskingDatabase     MaskingAnalyticsSummaryMetricNameEnum = "MASKING_DATABASE"
	MaskingAnalyticsSummaryMetricNameMaskingWorkRequest  MaskingAnalyticsSummaryMetricNameEnum = "MASKING_WORK_REQUEST"
	MaskingAnalyticsSummaryMetricNameMaskedSensitiveType MaskingAnalyticsSummaryMetricNameEnum = "MASKED_SENSITIVE_TYPE"
	MaskingAnalyticsSummaryMetricNameMaskedSchema        MaskingAnalyticsSummaryMetricNameEnum = "MASKED_SCHEMA"
	MaskingAnalyticsSummaryMetricNameMaskedTable         MaskingAnalyticsSummaryMetricNameEnum = "MASKED_TABLE"
	MaskingAnalyticsSummaryMetricNameMaskedColumn        MaskingAnalyticsSummaryMetricNameEnum = "MASKED_COLUMN"
	MaskingAnalyticsSummaryMetricNameMaskedDataValue     MaskingAnalyticsSummaryMetricNameEnum = "MASKED_DATA_VALUE"
)

var mappingMaskingAnalyticsSummaryMetricNameEnum = map[string]MaskingAnalyticsSummaryMetricNameEnum{
	"MASKING_POLICY":        MaskingAnalyticsSummaryMetricNameMaskingPolicy,
	"MASKING_DATABASE":      MaskingAnalyticsSummaryMetricNameMaskingDatabase,
	"MASKING_WORK_REQUEST":  MaskingAnalyticsSummaryMetricNameMaskingWorkRequest,
	"MASKED_SENSITIVE_TYPE": MaskingAnalyticsSummaryMetricNameMaskedSensitiveType,
	"MASKED_SCHEMA":         MaskingAnalyticsSummaryMetricNameMaskedSchema,
	"MASKED_TABLE":          MaskingAnalyticsSummaryMetricNameMaskedTable,
	"MASKED_COLUMN":         MaskingAnalyticsSummaryMetricNameMaskedColumn,
	"MASKED_DATA_VALUE":     MaskingAnalyticsSummaryMetricNameMaskedDataValue,
}

// GetMaskingAnalyticsSummaryMetricNameEnumValues Enumerates the set of values for MaskingAnalyticsSummaryMetricNameEnum
func GetMaskingAnalyticsSummaryMetricNameEnumValues() []MaskingAnalyticsSummaryMetricNameEnum {
	values := make([]MaskingAnalyticsSummaryMetricNameEnum, 0)
	for _, v := range mappingMaskingAnalyticsSummaryMetricNameEnum {
		values = append(values, v)
	}
	return values
}

// GetMaskingAnalyticsSummaryMetricNameEnumStringValues Enumerates the set of values in String for MaskingAnalyticsSummaryMetricNameEnum
func GetMaskingAnalyticsSummaryMetricNameEnumStringValues() []string {
	return []string{
		"MASKING_POLICY",
		"MASKING_DATABASE",
		"MASKING_WORK_REQUEST",
		"MASKED_SENSITIVE_TYPE",
		"MASKED_SCHEMA",
		"MASKED_TABLE",
		"MASKED_COLUMN",
		"MASKED_DATA_VALUE",
	}
}

// GetMappingMaskingAnalyticsSummaryMetricNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMaskingAnalyticsSummaryMetricNameEnum(val string) (MaskingAnalyticsSummaryMetricNameEnum, bool) {
	mappingMaskingAnalyticsSummaryMetricNameEnumIgnoreCase := make(map[string]MaskingAnalyticsSummaryMetricNameEnum)
	for k, v := range mappingMaskingAnalyticsSummaryMetricNameEnum {
		mappingMaskingAnalyticsSummaryMetricNameEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingMaskingAnalyticsSummaryMetricNameEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
