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

// DiscoveryAnalyticsSummary Summary of discovery analytics data.
type DiscoveryAnalyticsSummary struct {

	// The name of the aggregation metric.
	MetricName DiscoveryAnalyticsSummaryMetricNameEnum `mandatory:"true" json:"metricName"`

	// The total count for the aggregation metric.
	Count *int64 `mandatory:"true" json:"count"`

	Dimensions *Dimensions `mandatory:"false" json:"dimensions"`
}

func (m DiscoveryAnalyticsSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DiscoveryAnalyticsSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDiscoveryAnalyticsSummaryMetricNameEnum(string(m.MetricName)); !ok && m.MetricName != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MetricName: %s. Supported values are: %s.", m.MetricName, strings.Join(GetDiscoveryAnalyticsSummaryMetricNameEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DiscoveryAnalyticsSummaryMetricNameEnum Enum with underlying type: string
type DiscoveryAnalyticsSummaryMetricNameEnum string

// Set of constants representing the allowable values for DiscoveryAnalyticsSummaryMetricNameEnum
const (
	DiscoveryAnalyticsSummaryMetricNameDataModel DiscoveryAnalyticsSummaryMetricNameEnum = "SENSITIVE_DATA_MODEL"
	DiscoveryAnalyticsSummaryMetricNameType      DiscoveryAnalyticsSummaryMetricNameEnum = "SENSITIVE_TYPE"
	DiscoveryAnalyticsSummaryMetricNameSchema    DiscoveryAnalyticsSummaryMetricNameEnum = "SENSITIVE_SCHEMA"
	DiscoveryAnalyticsSummaryMetricNameTable     DiscoveryAnalyticsSummaryMetricNameEnum = "SENSITIVE_TABLE"
	DiscoveryAnalyticsSummaryMetricNameColumn    DiscoveryAnalyticsSummaryMetricNameEnum = "SENSITIVE_COLUMN"
	DiscoveryAnalyticsSummaryMetricNameDataValue DiscoveryAnalyticsSummaryMetricNameEnum = "SENSITIVE_DATA_VALUE"
)

var mappingDiscoveryAnalyticsSummaryMetricNameEnum = map[string]DiscoveryAnalyticsSummaryMetricNameEnum{
	"SENSITIVE_DATA_MODEL": DiscoveryAnalyticsSummaryMetricNameDataModel,
	"SENSITIVE_TYPE":       DiscoveryAnalyticsSummaryMetricNameType,
	"SENSITIVE_SCHEMA":     DiscoveryAnalyticsSummaryMetricNameSchema,
	"SENSITIVE_TABLE":      DiscoveryAnalyticsSummaryMetricNameTable,
	"SENSITIVE_COLUMN":     DiscoveryAnalyticsSummaryMetricNameColumn,
	"SENSITIVE_DATA_VALUE": DiscoveryAnalyticsSummaryMetricNameDataValue,
}

// GetDiscoveryAnalyticsSummaryMetricNameEnumValues Enumerates the set of values for DiscoveryAnalyticsSummaryMetricNameEnum
func GetDiscoveryAnalyticsSummaryMetricNameEnumValues() []DiscoveryAnalyticsSummaryMetricNameEnum {
	values := make([]DiscoveryAnalyticsSummaryMetricNameEnum, 0)
	for _, v := range mappingDiscoveryAnalyticsSummaryMetricNameEnum {
		values = append(values, v)
	}
	return values
}

// GetDiscoveryAnalyticsSummaryMetricNameEnumStringValues Enumerates the set of values in String for DiscoveryAnalyticsSummaryMetricNameEnum
func GetDiscoveryAnalyticsSummaryMetricNameEnumStringValues() []string {
	return []string{
		"SENSITIVE_DATA_MODEL",
		"SENSITIVE_TYPE",
		"SENSITIVE_SCHEMA",
		"SENSITIVE_TABLE",
		"SENSITIVE_COLUMN",
		"SENSITIVE_DATA_VALUE",
	}
}

// GetMappingDiscoveryAnalyticsSummaryMetricNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDiscoveryAnalyticsSummaryMetricNameEnum(val string) (DiscoveryAnalyticsSummaryMetricNameEnum, bool) {
	mappingDiscoveryAnalyticsSummaryMetricNameEnumIgnoreCase := make(map[string]DiscoveryAnalyticsSummaryMetricNameEnum)
	for k, v := range mappingDiscoveryAnalyticsSummaryMetricNameEnum {
		mappingDiscoveryAnalyticsSummaryMetricNameEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingDiscoveryAnalyticsSummaryMetricNameEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
