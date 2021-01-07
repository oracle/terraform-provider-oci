// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Big Data Service API
//
// API for the Big Data Service. Use this API to build, deploy, and manage fully elastic Big Data Service
// build on Hadoop, Spark and Data Science distribution, which can be fully integrated with existing enterprise
// data in Oracle Database and Oracle Applications..
//

package bds

import (
	"github.com/oracle/oci-go-sdk/v31/common"
)

// MetricThresholdRule An autoscaling action is triggered when a performance metric meets or exceeds a threshold
type MetricThresholdRule struct {

	// This value is the minimum period of time metric value meets or exceeds threshold value before action is trigger. The value is in minutes.
	DurationInMinutes *int `mandatory:"true" json:"durationInMinutes"`

	// The comparison operator to use. Options are greater than (GT), less than (LT).
	Operator MetricThresholdRuleOperatorEnum `mandatory:"true" json:"operator"`

	// integer non negative value. 0 < value < 100
	Value *int `mandatory:"true" json:"value"`
}

func (m MetricThresholdRule) String() string {
	return common.PointerString(m)
}

// MetricThresholdRuleOperatorEnum Enum with underlying type: string
type MetricThresholdRuleOperatorEnum string

// Set of constants representing the allowable values for MetricThresholdRuleOperatorEnum
const (
	MetricThresholdRuleOperatorGt MetricThresholdRuleOperatorEnum = "GT"
	MetricThresholdRuleOperatorLt MetricThresholdRuleOperatorEnum = "LT"
)

var mappingMetricThresholdRuleOperator = map[string]MetricThresholdRuleOperatorEnum{
	"GT": MetricThresholdRuleOperatorGt,
	"LT": MetricThresholdRuleOperatorLt,
}

// GetMetricThresholdRuleOperatorEnumValues Enumerates the set of values for MetricThresholdRuleOperatorEnum
func GetMetricThresholdRuleOperatorEnumValues() []MetricThresholdRuleOperatorEnum {
	values := make([]MetricThresholdRuleOperatorEnum, 0)
	for _, v := range mappingMetricThresholdRuleOperator {
		values = append(values, v)
	}
	return values
}
