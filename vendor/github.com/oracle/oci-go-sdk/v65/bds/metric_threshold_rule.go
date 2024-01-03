// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Big Data Service API
//
// REST API for Oracle Big Data Service. Use this API to build, deploy, and manage fully elastic Big Data Service clusters. Build on Hadoop, Spark and Data Science distributions, which can be fully integrated with existing enterprise data in Oracle Database and Oracle applications.
//

package bds

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MetricThresholdRule An autoscale action is triggered when a performance metric exceeds a threshold.
type MetricThresholdRule struct {

	// This value is the minimum period of time the metric value exceeds the threshold value before the action is triggered. The value is in minutes.
	DurationInMinutes *int `mandatory:"true" json:"durationInMinutes"`

	// The comparison operator to use. Options are greater than (GT) or less than (LT).
	Operator MetricThresholdRuleOperatorEnum `mandatory:"true" json:"operator"`

	// Integer non-negative value. 0 < value < 100
	Value *int `mandatory:"true" json:"value"`
}

func (m MetricThresholdRule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MetricThresholdRule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMetricThresholdRuleOperatorEnum(string(m.Operator)); !ok && m.Operator != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Operator: %s. Supported values are: %s.", m.Operator, strings.Join(GetMetricThresholdRuleOperatorEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MetricThresholdRuleOperatorEnum Enum with underlying type: string
type MetricThresholdRuleOperatorEnum string

// Set of constants representing the allowable values for MetricThresholdRuleOperatorEnum
const (
	MetricThresholdRuleOperatorGt MetricThresholdRuleOperatorEnum = "GT"
	MetricThresholdRuleOperatorLt MetricThresholdRuleOperatorEnum = "LT"
)

var mappingMetricThresholdRuleOperatorEnum = map[string]MetricThresholdRuleOperatorEnum{
	"GT": MetricThresholdRuleOperatorGt,
	"LT": MetricThresholdRuleOperatorLt,
}

var mappingMetricThresholdRuleOperatorEnumLowerCase = map[string]MetricThresholdRuleOperatorEnum{
	"gt": MetricThresholdRuleOperatorGt,
	"lt": MetricThresholdRuleOperatorLt,
}

// GetMetricThresholdRuleOperatorEnumValues Enumerates the set of values for MetricThresholdRuleOperatorEnum
func GetMetricThresholdRuleOperatorEnumValues() []MetricThresholdRuleOperatorEnum {
	values := make([]MetricThresholdRuleOperatorEnum, 0)
	for _, v := range mappingMetricThresholdRuleOperatorEnum {
		values = append(values, v)
	}
	return values
}

// GetMetricThresholdRuleOperatorEnumStringValues Enumerates the set of values in String for MetricThresholdRuleOperatorEnum
func GetMetricThresholdRuleOperatorEnumStringValues() []string {
	return []string{
		"GT",
		"LT",
	}
}

// GetMappingMetricThresholdRuleOperatorEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMetricThresholdRuleOperatorEnum(val string) (MetricThresholdRuleOperatorEnum, bool) {
	enum, ok := mappingMetricThresholdRuleOperatorEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
