// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TargetPredefinedMetricExpressionRule An expression built using CPU or Memory metrics for triggering an autoscaling action for workload.
type TargetPredefinedMetricExpressionRule struct {
	ScaleConfiguration *TargetPredefinedExpressionThresholdScalingConfiguration `mandatory:"true" json:"scaleConfiguration"`

	// Metric type
	MetricType TargetPredefinedMetricExpressionRuleMetricTypeEnum `mandatory:"true" json:"metricType"`
}

func (m TargetPredefinedMetricExpressionRule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TargetPredefinedMetricExpressionRule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingTargetPredefinedMetricExpressionRuleMetricTypeEnum(string(m.MetricType)); !ok && m.MetricType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MetricType: %s. Supported values are: %s.", m.MetricType, strings.Join(GetTargetPredefinedMetricExpressionRuleMetricTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m TargetPredefinedMetricExpressionRule) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeTargetPredefinedMetricExpressionRule TargetPredefinedMetricExpressionRule
	s := struct {
		DiscriminatorParam string `json:"metricExpressionRuleType"`
		MarshalTypeTargetPredefinedMetricExpressionRule
	}{
		"TARGET_PREDEFINED_EXPRESSION",
		(MarshalTypeTargetPredefinedMetricExpressionRule)(m),
	}

	return json.Marshal(&s)
}

// TargetPredefinedMetricExpressionRuleMetricTypeEnum Enum with underlying type: string
type TargetPredefinedMetricExpressionRuleMetricTypeEnum string

// Set of constants representing the allowable values for TargetPredefinedMetricExpressionRuleMetricTypeEnum
const (
	TargetPredefinedMetricExpressionRuleMetricTypeCpuUtilization    TargetPredefinedMetricExpressionRuleMetricTypeEnum = "CPU_UTILIZATION"
	TargetPredefinedMetricExpressionRuleMetricTypeMemoryUtilization TargetPredefinedMetricExpressionRuleMetricTypeEnum = "MEMORY_UTILIZATION"
)

var mappingTargetPredefinedMetricExpressionRuleMetricTypeEnum = map[string]TargetPredefinedMetricExpressionRuleMetricTypeEnum{
	"CPU_UTILIZATION":    TargetPredefinedMetricExpressionRuleMetricTypeCpuUtilization,
	"MEMORY_UTILIZATION": TargetPredefinedMetricExpressionRuleMetricTypeMemoryUtilization,
}

var mappingTargetPredefinedMetricExpressionRuleMetricTypeEnumLowerCase = map[string]TargetPredefinedMetricExpressionRuleMetricTypeEnum{
	"cpu_utilization":    TargetPredefinedMetricExpressionRuleMetricTypeCpuUtilization,
	"memory_utilization": TargetPredefinedMetricExpressionRuleMetricTypeMemoryUtilization,
}

// GetTargetPredefinedMetricExpressionRuleMetricTypeEnumValues Enumerates the set of values for TargetPredefinedMetricExpressionRuleMetricTypeEnum
func GetTargetPredefinedMetricExpressionRuleMetricTypeEnumValues() []TargetPredefinedMetricExpressionRuleMetricTypeEnum {
	values := make([]TargetPredefinedMetricExpressionRuleMetricTypeEnum, 0)
	for _, v := range mappingTargetPredefinedMetricExpressionRuleMetricTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTargetPredefinedMetricExpressionRuleMetricTypeEnumStringValues Enumerates the set of values in String for TargetPredefinedMetricExpressionRuleMetricTypeEnum
func GetTargetPredefinedMetricExpressionRuleMetricTypeEnumStringValues() []string {
	return []string{
		"CPU_UTILIZATION",
		"MEMORY_UTILIZATION",
	}
}

// GetMappingTargetPredefinedMetricExpressionRuleMetricTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTargetPredefinedMetricExpressionRuleMetricTypeEnum(val string) (TargetPredefinedMetricExpressionRuleMetricTypeEnum, bool) {
	enum, ok := mappingTargetPredefinedMetricExpressionRuleMetricTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
