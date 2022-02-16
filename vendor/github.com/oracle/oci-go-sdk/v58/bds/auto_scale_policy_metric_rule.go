// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Big Data Service API
//
// REST API for Oracle Big Data Service. Use this API to build, deploy, and manage fully elastic Big Data Service clusters. Build on Hadoop, Spark and Data Science distributions, which can be fully integrated with existing enterprise data in Oracle Database and Oracle applications.
//

package bds

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AutoScalePolicyMetricRule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAutoScalePolicyMetricRuleMetricTypeEnum(string(m.MetricType)); !ok && m.MetricType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MetricType: %s. Supported values are: %s.", m.MetricType, strings.Join(GetAutoScalePolicyMetricRuleMetricTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AutoScalePolicyMetricRuleMetricTypeEnum Enum with underlying type: string
type AutoScalePolicyMetricRuleMetricTypeEnum string

// Set of constants representing the allowable values for AutoScalePolicyMetricRuleMetricTypeEnum
const (
	AutoScalePolicyMetricRuleMetricTypeCpuUtilization AutoScalePolicyMetricRuleMetricTypeEnum = "CPU_UTILIZATION"
)

var mappingAutoScalePolicyMetricRuleMetricTypeEnum = map[string]AutoScalePolicyMetricRuleMetricTypeEnum{
	"CPU_UTILIZATION": AutoScalePolicyMetricRuleMetricTypeCpuUtilization,
}

// GetAutoScalePolicyMetricRuleMetricTypeEnumValues Enumerates the set of values for AutoScalePolicyMetricRuleMetricTypeEnum
func GetAutoScalePolicyMetricRuleMetricTypeEnumValues() []AutoScalePolicyMetricRuleMetricTypeEnum {
	values := make([]AutoScalePolicyMetricRuleMetricTypeEnum, 0)
	for _, v := range mappingAutoScalePolicyMetricRuleMetricTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAutoScalePolicyMetricRuleMetricTypeEnumStringValues Enumerates the set of values in String for AutoScalePolicyMetricRuleMetricTypeEnum
func GetAutoScalePolicyMetricRuleMetricTypeEnumStringValues() []string {
	return []string{
		"CPU_UTILIZATION",
	}
}

// GetMappingAutoScalePolicyMetricRuleMetricTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutoScalePolicyMetricRuleMetricTypeEnum(val string) (AutoScalePolicyMetricRuleMetricTypeEnum, bool) {
	mappingAutoScalePolicyMetricRuleMetricTypeEnumIgnoreCase := make(map[string]AutoScalePolicyMetricRuleMetricTypeEnum)
	for k, v := range mappingAutoScalePolicyMetricRuleMetricTypeEnum {
		mappingAutoScalePolicyMetricRuleMetricTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingAutoScalePolicyMetricRuleMetricTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
