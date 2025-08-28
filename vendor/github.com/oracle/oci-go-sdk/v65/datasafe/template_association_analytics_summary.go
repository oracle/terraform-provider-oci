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

// TemplateAssociationAnalyticsSummary The summary of template association analytics data.
type TemplateAssociationAnalyticsSummary struct {

	// The name of the aggregation metric.
	MetricName TemplateAssociationAnalyticsSummaryMetricNameEnum `mandatory:"true" json:"metricName"`

	// The total count for the aggregation metric.
	Count *int64 `mandatory:"true" json:"count"`

	Dimensions *TemplateAssociationAnalyticsDimensions `mandatory:"false" json:"dimensions"`
}

func (m TemplateAssociationAnalyticsSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TemplateAssociationAnalyticsSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingTemplateAssociationAnalyticsSummaryMetricNameEnum(string(m.MetricName)); !ok && m.MetricName != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MetricName: %s. Supported values are: %s.", m.MetricName, strings.Join(GetTemplateAssociationAnalyticsSummaryMetricNameEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TemplateAssociationAnalyticsSummaryMetricNameEnum Enum with underlying type: string
type TemplateAssociationAnalyticsSummaryMetricNameEnum string

// Set of constants representing the allowable values for TemplateAssociationAnalyticsSummaryMetricNameEnum
const (
	TemplateAssociationAnalyticsSummaryMetricNameTemplateAssociationStats TemplateAssociationAnalyticsSummaryMetricNameEnum = "TEMPLATE_ASSOCIATION_STATS"
)

var mappingTemplateAssociationAnalyticsSummaryMetricNameEnum = map[string]TemplateAssociationAnalyticsSummaryMetricNameEnum{
	"TEMPLATE_ASSOCIATION_STATS": TemplateAssociationAnalyticsSummaryMetricNameTemplateAssociationStats,
}

var mappingTemplateAssociationAnalyticsSummaryMetricNameEnumLowerCase = map[string]TemplateAssociationAnalyticsSummaryMetricNameEnum{
	"template_association_stats": TemplateAssociationAnalyticsSummaryMetricNameTemplateAssociationStats,
}

// GetTemplateAssociationAnalyticsSummaryMetricNameEnumValues Enumerates the set of values for TemplateAssociationAnalyticsSummaryMetricNameEnum
func GetTemplateAssociationAnalyticsSummaryMetricNameEnumValues() []TemplateAssociationAnalyticsSummaryMetricNameEnum {
	values := make([]TemplateAssociationAnalyticsSummaryMetricNameEnum, 0)
	for _, v := range mappingTemplateAssociationAnalyticsSummaryMetricNameEnum {
		values = append(values, v)
	}
	return values
}

// GetTemplateAssociationAnalyticsSummaryMetricNameEnumStringValues Enumerates the set of values in String for TemplateAssociationAnalyticsSummaryMetricNameEnum
func GetTemplateAssociationAnalyticsSummaryMetricNameEnumStringValues() []string {
	return []string{
		"TEMPLATE_ASSOCIATION_STATS",
	}
}

// GetMappingTemplateAssociationAnalyticsSummaryMetricNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTemplateAssociationAnalyticsSummaryMetricNameEnum(val string) (TemplateAssociationAnalyticsSummaryMetricNameEnum, bool) {
	enum, ok := mappingTemplateAssociationAnalyticsSummaryMetricNameEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
