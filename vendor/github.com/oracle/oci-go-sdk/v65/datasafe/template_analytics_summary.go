// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// TemplateAnalyticsSummary The summary of template analytics data.
type TemplateAnalyticsSummary struct {

	// The name of the aggregation metric.
	MetricName TemplateAnalyticsSummaryMetricNameEnum `mandatory:"true" json:"metricName"`

	// The total count for the aggregation metric.
	Count *int64 `mandatory:"true" json:"count"`

	Dimensions *TemplateAnalyticsDimensions `mandatory:"false" json:"dimensions"`
}

func (m TemplateAnalyticsSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TemplateAnalyticsSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingTemplateAnalyticsSummaryMetricNameEnum(string(m.MetricName)); !ok && m.MetricName != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MetricName: %s. Supported values are: %s.", m.MetricName, strings.Join(GetTemplateAnalyticsSummaryMetricNameEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TemplateAnalyticsSummaryMetricNameEnum Enum with underlying type: string
type TemplateAnalyticsSummaryMetricNameEnum string

// Set of constants representing the allowable values for TemplateAnalyticsSummaryMetricNameEnum
const (
	TemplateAnalyticsSummaryMetricNameTemplateStats TemplateAnalyticsSummaryMetricNameEnum = "TEMPLATE_STATS"
)

var mappingTemplateAnalyticsSummaryMetricNameEnum = map[string]TemplateAnalyticsSummaryMetricNameEnum{
	"TEMPLATE_STATS": TemplateAnalyticsSummaryMetricNameTemplateStats,
}

var mappingTemplateAnalyticsSummaryMetricNameEnumLowerCase = map[string]TemplateAnalyticsSummaryMetricNameEnum{
	"template_stats": TemplateAnalyticsSummaryMetricNameTemplateStats,
}

// GetTemplateAnalyticsSummaryMetricNameEnumValues Enumerates the set of values for TemplateAnalyticsSummaryMetricNameEnum
func GetTemplateAnalyticsSummaryMetricNameEnumValues() []TemplateAnalyticsSummaryMetricNameEnum {
	values := make([]TemplateAnalyticsSummaryMetricNameEnum, 0)
	for _, v := range mappingTemplateAnalyticsSummaryMetricNameEnum {
		values = append(values, v)
	}
	return values
}

// GetTemplateAnalyticsSummaryMetricNameEnumStringValues Enumerates the set of values in String for TemplateAnalyticsSummaryMetricNameEnum
func GetTemplateAnalyticsSummaryMetricNameEnumStringValues() []string {
	return []string{
		"TEMPLATE_STATS",
	}
}

// GetMappingTemplateAnalyticsSummaryMetricNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTemplateAnalyticsSummaryMetricNameEnum(val string) (TemplateAnalyticsSummaryMetricNameEnum, bool) {
	enum, ok := mappingTemplateAnalyticsSummaryMetricNameEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
