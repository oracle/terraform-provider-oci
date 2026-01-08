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

// ManagedComputeClusterPredefinedMetricExpressionRule An expression built using CPU or Memory metrics for triggering an autoscaling action on the managed compute cluster type compute target .
type ManagedComputeClusterPredefinedMetricExpressionRule struct {
	ScaleInConfiguration *ManagedComputeClusterPredefinedExpressionThresholdScalingConfiguration `mandatory:"true" json:"scaleInConfiguration"`

	ScaleOutConfiguration *ManagedComputeClusterPredefinedExpressionThresholdScalingConfiguration `mandatory:"true" json:"scaleOutConfiguration"`

	// Metric type
	MetricType ManagedComputeClusterPredefinedMetricExpressionRuleMetricTypeEnum `mandatory:"true" json:"metricType"`
}

func (m ManagedComputeClusterPredefinedMetricExpressionRule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ManagedComputeClusterPredefinedMetricExpressionRule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingManagedComputeClusterPredefinedMetricExpressionRuleMetricTypeEnum(string(m.MetricType)); !ok && m.MetricType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MetricType: %s. Supported values are: %s.", m.MetricType, strings.Join(GetManagedComputeClusterPredefinedMetricExpressionRuleMetricTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ManagedComputeClusterPredefinedMetricExpressionRule) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeManagedComputeClusterPredefinedMetricExpressionRule ManagedComputeClusterPredefinedMetricExpressionRule
	s := struct {
		DiscriminatorParam string `json:"metricExpressionRuleType"`
		MarshalTypeManagedComputeClusterPredefinedMetricExpressionRule
	}{
		"PREDEFINED_EXPRESSION",
		(MarshalTypeManagedComputeClusterPredefinedMetricExpressionRule)(m),
	}

	return json.Marshal(&s)
}

// ManagedComputeClusterPredefinedMetricExpressionRuleMetricTypeEnum Enum with underlying type: string
type ManagedComputeClusterPredefinedMetricExpressionRuleMetricTypeEnum string

// Set of constants representing the allowable values for ManagedComputeClusterPredefinedMetricExpressionRuleMetricTypeEnum
const (
	ManagedComputeClusterPredefinedMetricExpressionRuleMetricTypeCpuUtilization    ManagedComputeClusterPredefinedMetricExpressionRuleMetricTypeEnum = "CPU_UTILIZATION"
	ManagedComputeClusterPredefinedMetricExpressionRuleMetricTypeMemoryUtilization ManagedComputeClusterPredefinedMetricExpressionRuleMetricTypeEnum = "MEMORY_UTILIZATION"
)

var mappingManagedComputeClusterPredefinedMetricExpressionRuleMetricTypeEnum = map[string]ManagedComputeClusterPredefinedMetricExpressionRuleMetricTypeEnum{
	"CPU_UTILIZATION":    ManagedComputeClusterPredefinedMetricExpressionRuleMetricTypeCpuUtilization,
	"MEMORY_UTILIZATION": ManagedComputeClusterPredefinedMetricExpressionRuleMetricTypeMemoryUtilization,
}

var mappingManagedComputeClusterPredefinedMetricExpressionRuleMetricTypeEnumLowerCase = map[string]ManagedComputeClusterPredefinedMetricExpressionRuleMetricTypeEnum{
	"cpu_utilization":    ManagedComputeClusterPredefinedMetricExpressionRuleMetricTypeCpuUtilization,
	"memory_utilization": ManagedComputeClusterPredefinedMetricExpressionRuleMetricTypeMemoryUtilization,
}

// GetManagedComputeClusterPredefinedMetricExpressionRuleMetricTypeEnumValues Enumerates the set of values for ManagedComputeClusterPredefinedMetricExpressionRuleMetricTypeEnum
func GetManagedComputeClusterPredefinedMetricExpressionRuleMetricTypeEnumValues() []ManagedComputeClusterPredefinedMetricExpressionRuleMetricTypeEnum {
	values := make([]ManagedComputeClusterPredefinedMetricExpressionRuleMetricTypeEnum, 0)
	for _, v := range mappingManagedComputeClusterPredefinedMetricExpressionRuleMetricTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetManagedComputeClusterPredefinedMetricExpressionRuleMetricTypeEnumStringValues Enumerates the set of values in String for ManagedComputeClusterPredefinedMetricExpressionRuleMetricTypeEnum
func GetManagedComputeClusterPredefinedMetricExpressionRuleMetricTypeEnumStringValues() []string {
	return []string{
		"CPU_UTILIZATION",
		"MEMORY_UTILIZATION",
	}
}

// GetMappingManagedComputeClusterPredefinedMetricExpressionRuleMetricTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingManagedComputeClusterPredefinedMetricExpressionRuleMetricTypeEnum(val string) (ManagedComputeClusterPredefinedMetricExpressionRuleMetricTypeEnum, bool) {
	enum, ok := mappingManagedComputeClusterPredefinedMetricExpressionRuleMetricTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
