// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// TemplateAnalyticSummary The summary of template analytics data.
type TemplateAnalyticSummary struct {

	// The name of the aggregation metric.
	MetricName TemplateAnalyticSummaryMetricNameEnum `mandatory:"true" json:"metricName"`

	// The total count for the aggregation metric.
	Count *int64 `mandatory:"true" json:"count"`

	Dimensions *TemplateAnalyticDimensions `mandatory:"false" json:"dimensions"`
}

func (m TemplateAnalyticSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TemplateAnalyticSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingTemplateAnalyticSummaryMetricNameEnum(string(m.MetricName)); !ok && m.MetricName != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MetricName: %s. Supported values are: %s.", m.MetricName, strings.Join(GetTemplateAnalyticSummaryMetricNameEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TemplateAnalyticSummaryMetricNameEnum Enum with underlying type: string
type TemplateAnalyticSummaryMetricNameEnum string

// Set of constants representing the allowable values for TemplateAnalyticSummaryMetricNameEnum
const (
	TemplateAnalyticSummaryMetricNameTemplateStats TemplateAnalyticSummaryMetricNameEnum = "TEMPLATE_STATS"
)

var mappingTemplateAnalyticSummaryMetricNameEnum = map[string]TemplateAnalyticSummaryMetricNameEnum{
	"TEMPLATE_STATS": TemplateAnalyticSummaryMetricNameTemplateStats,
}

var mappingTemplateAnalyticSummaryMetricNameEnumLowerCase = map[string]TemplateAnalyticSummaryMetricNameEnum{
	"template_stats": TemplateAnalyticSummaryMetricNameTemplateStats,
}

// GetTemplateAnalyticSummaryMetricNameEnumValues Enumerates the set of values for TemplateAnalyticSummaryMetricNameEnum
func GetTemplateAnalyticSummaryMetricNameEnumValues() []TemplateAnalyticSummaryMetricNameEnum {
	values := make([]TemplateAnalyticSummaryMetricNameEnum, 0)
	for _, v := range mappingTemplateAnalyticSummaryMetricNameEnum {
		values = append(values, v)
	}
	return values
}

// GetTemplateAnalyticSummaryMetricNameEnumStringValues Enumerates the set of values in String for TemplateAnalyticSummaryMetricNameEnum
func GetTemplateAnalyticSummaryMetricNameEnumStringValues() []string {
	return []string{
		"TEMPLATE_STATS",
	}
}

// GetMappingTemplateAnalyticSummaryMetricNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTemplateAnalyticSummaryMetricNameEnum(val string) (TemplateAnalyticSummaryMetricNameEnum, bool) {
	enum, ok := mappingTemplateAnalyticSummaryMetricNameEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
