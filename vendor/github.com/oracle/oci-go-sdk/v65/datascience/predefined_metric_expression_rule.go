// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// PredefinedMetricExpressionRule An expression built using CPU or Memory metrics for triggering an autoscaling action on the model deployment.
type PredefinedMetricExpressionRule struct {
	ScaleInConfiguration *PredefinedExpressionThresholdScalingConfiguration `mandatory:"true" json:"scaleInConfiguration"`

	ScaleOutConfiguration *PredefinedExpressionThresholdScalingConfiguration `mandatory:"true" json:"scaleOutConfiguration"`

	// Metric type
	MetricType PredefinedMetricExpressionRuleMetricTypeEnum `mandatory:"true" json:"metricType"`
}

func (m PredefinedMetricExpressionRule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PredefinedMetricExpressionRule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPredefinedMetricExpressionRuleMetricTypeEnum(string(m.MetricType)); !ok && m.MetricType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MetricType: %s. Supported values are: %s.", m.MetricType, strings.Join(GetPredefinedMetricExpressionRuleMetricTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m PredefinedMetricExpressionRule) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePredefinedMetricExpressionRule PredefinedMetricExpressionRule
	s := struct {
		DiscriminatorParam string `json:"metricExpressionRuleType"`
		MarshalTypePredefinedMetricExpressionRule
	}{
		"PREDEFINED_EXPRESSION",
		(MarshalTypePredefinedMetricExpressionRule)(m),
	}

	return json.Marshal(&s)
}

// PredefinedMetricExpressionRuleMetricTypeEnum Enum with underlying type: string
type PredefinedMetricExpressionRuleMetricTypeEnum string

// Set of constants representing the allowable values for PredefinedMetricExpressionRuleMetricTypeEnum
const (
	PredefinedMetricExpressionRuleMetricTypeCpuUtilization    PredefinedMetricExpressionRuleMetricTypeEnum = "CPU_UTILIZATION"
	PredefinedMetricExpressionRuleMetricTypeMemoryUtilization PredefinedMetricExpressionRuleMetricTypeEnum = "MEMORY_UTILIZATION"
)

var mappingPredefinedMetricExpressionRuleMetricTypeEnum = map[string]PredefinedMetricExpressionRuleMetricTypeEnum{
	"CPU_UTILIZATION":    PredefinedMetricExpressionRuleMetricTypeCpuUtilization,
	"MEMORY_UTILIZATION": PredefinedMetricExpressionRuleMetricTypeMemoryUtilization,
}

var mappingPredefinedMetricExpressionRuleMetricTypeEnumLowerCase = map[string]PredefinedMetricExpressionRuleMetricTypeEnum{
	"cpu_utilization":    PredefinedMetricExpressionRuleMetricTypeCpuUtilization,
	"memory_utilization": PredefinedMetricExpressionRuleMetricTypeMemoryUtilization,
}

// GetPredefinedMetricExpressionRuleMetricTypeEnumValues Enumerates the set of values for PredefinedMetricExpressionRuleMetricTypeEnum
func GetPredefinedMetricExpressionRuleMetricTypeEnumValues() []PredefinedMetricExpressionRuleMetricTypeEnum {
	values := make([]PredefinedMetricExpressionRuleMetricTypeEnum, 0)
	for _, v := range mappingPredefinedMetricExpressionRuleMetricTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetPredefinedMetricExpressionRuleMetricTypeEnumStringValues Enumerates the set of values in String for PredefinedMetricExpressionRuleMetricTypeEnum
func GetPredefinedMetricExpressionRuleMetricTypeEnumStringValues() []string {
	return []string{
		"CPU_UTILIZATION",
		"MEMORY_UTILIZATION",
	}
}

// GetMappingPredefinedMetricExpressionRuleMetricTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPredefinedMetricExpressionRuleMetricTypeEnum(val string) (PredefinedMetricExpressionRuleMetricTypeEnum, bool) {
	enum, ok := mappingPredefinedMetricExpressionRuleMetricTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
