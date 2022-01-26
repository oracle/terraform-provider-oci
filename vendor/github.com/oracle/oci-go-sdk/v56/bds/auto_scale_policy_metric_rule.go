// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Big Data Service API
//
// REST API for Oracle Big Data Service. Use this API to build, deploy, and manage fully elastic Big Data Service clusters. Build on Hadoop, Spark and Data Science distributions, which can be fully integrated with existing enterprise data in Oracle Database and Oracle applications.
//

package bds

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// AutoScalePolicyMetricRule Metric and threshold details for triggering an autoscale action.
type AutoScalePolicyMetricRule struct {

	// Allowed value is CPU_UTILIZATION.
	MetricType AutoScalePolicyMetricRuleMetricTypeEnum `mandatory:"true" json:"metricType"`

	Threshold *MetricThresholdRule `mandatory:"true" json:"threshold"`
}

func (m AutoScalePolicyMetricRule) String() string {
	return common.PointerString(m)
}

// AutoScalePolicyMetricRuleMetricTypeEnum Enum with underlying type: string
type AutoScalePolicyMetricRuleMetricTypeEnum string

// Set of constants representing the allowable values for AutoScalePolicyMetricRuleMetricTypeEnum
const (
	AutoScalePolicyMetricRuleMetricTypeCpuUtilization AutoScalePolicyMetricRuleMetricTypeEnum = "CPU_UTILIZATION"
)

var mappingAutoScalePolicyMetricRuleMetricType = map[string]AutoScalePolicyMetricRuleMetricTypeEnum{
	"CPU_UTILIZATION": AutoScalePolicyMetricRuleMetricTypeCpuUtilization,
}

// GetAutoScalePolicyMetricRuleMetricTypeEnumValues Enumerates the set of values for AutoScalePolicyMetricRuleMetricTypeEnum
func GetAutoScalePolicyMetricRuleMetricTypeEnumValues() []AutoScalePolicyMetricRuleMetricTypeEnum {
	values := make([]AutoScalePolicyMetricRuleMetricTypeEnum, 0)
	for _, v := range mappingAutoScalePolicyMetricRuleMetricType {
		values = append(values, v)
	}
	return values
}
