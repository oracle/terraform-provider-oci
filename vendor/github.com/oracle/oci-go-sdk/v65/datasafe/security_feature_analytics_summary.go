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

// SecurityFeatureAnalyticsSummary The summary of database security feature analytics data.
type SecurityFeatureAnalyticsSummary struct {

	// The name of the aggregation metric.
	MetricName SecurityFeatureAnalyticsSummaryMetricNameEnum `mandatory:"true" json:"metricName"`

	// The total count for the aggregation metric.
	Count *int64 `mandatory:"true" json:"count"`

	Dimensions *SecurityFeatureAnalyticsDimensions `mandatory:"false" json:"dimensions"`
}

func (m SecurityFeatureAnalyticsSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SecurityFeatureAnalyticsSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSecurityFeatureAnalyticsSummaryMetricNameEnum(string(m.MetricName)); !ok && m.MetricName != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MetricName: %s. Supported values are: %s.", m.MetricName, strings.Join(GetSecurityFeatureAnalyticsSummaryMetricNameEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SecurityFeatureAnalyticsSummaryMetricNameEnum Enum with underlying type: string
type SecurityFeatureAnalyticsSummaryMetricNameEnum string

// Set of constants representing the allowable values for SecurityFeatureAnalyticsSummaryMetricNameEnum
const (
	SecurityFeatureAnalyticsSummaryMetricNameSecurityFeatureStats SecurityFeatureAnalyticsSummaryMetricNameEnum = "SECURITY_FEATURE_STATS"
)

var mappingSecurityFeatureAnalyticsSummaryMetricNameEnum = map[string]SecurityFeatureAnalyticsSummaryMetricNameEnum{
	"SECURITY_FEATURE_STATS": SecurityFeatureAnalyticsSummaryMetricNameSecurityFeatureStats,
}

var mappingSecurityFeatureAnalyticsSummaryMetricNameEnumLowerCase = map[string]SecurityFeatureAnalyticsSummaryMetricNameEnum{
	"security_feature_stats": SecurityFeatureAnalyticsSummaryMetricNameSecurityFeatureStats,
}

// GetSecurityFeatureAnalyticsSummaryMetricNameEnumValues Enumerates the set of values for SecurityFeatureAnalyticsSummaryMetricNameEnum
func GetSecurityFeatureAnalyticsSummaryMetricNameEnumValues() []SecurityFeatureAnalyticsSummaryMetricNameEnum {
	values := make([]SecurityFeatureAnalyticsSummaryMetricNameEnum, 0)
	for _, v := range mappingSecurityFeatureAnalyticsSummaryMetricNameEnum {
		values = append(values, v)
	}
	return values
}

// GetSecurityFeatureAnalyticsSummaryMetricNameEnumStringValues Enumerates the set of values in String for SecurityFeatureAnalyticsSummaryMetricNameEnum
func GetSecurityFeatureAnalyticsSummaryMetricNameEnumStringValues() []string {
	return []string{
		"SECURITY_FEATURE_STATS",
	}
}

// GetMappingSecurityFeatureAnalyticsSummaryMetricNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSecurityFeatureAnalyticsSummaryMetricNameEnum(val string) (SecurityFeatureAnalyticsSummaryMetricNameEnum, bool) {
	enum, ok := mappingSecurityFeatureAnalyticsSummaryMetricNameEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
