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

// ManagedComputeClusterWorkloadMetricExpressionRule The metric expression rule base for workload scaling.
type ManagedComputeClusterWorkloadMetricExpressionRule interface {
}

type managedcomputeclusterworkloadmetricexpressionrule struct {
	JsonData                 []byte
	MetricExpressionRuleType string `json:"metricExpressionRuleType"`
}

// UnmarshalJSON unmarshals json
func (m *managedcomputeclusterworkloadmetricexpressionrule) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalermanagedcomputeclusterworkloadmetricexpressionrule managedcomputeclusterworkloadmetricexpressionrule
	s := struct {
		Model Unmarshalermanagedcomputeclusterworkloadmetricexpressionrule
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.MetricExpressionRuleType = s.Model.MetricExpressionRuleType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *managedcomputeclusterworkloadmetricexpressionrule) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.MetricExpressionRuleType {
	case "TARGET_CUSTOM_EXPRESSION":
		mm := TargetCustomMetricExpressionRule{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "TARGET_PREDEFINED_EXPRESSION":
		mm := TargetPredefinedMetricExpressionRule{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for ManagedComputeClusterWorkloadMetricExpressionRule: %s.", m.MetricExpressionRuleType)
		return *m, nil
	}
}

func (m managedcomputeclusterworkloadmetricexpressionrule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m managedcomputeclusterworkloadmetricexpressionrule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ManagedComputeClusterWorkloadMetricExpressionRuleMetricExpressionRuleTypeEnum Enum with underlying type: string
type ManagedComputeClusterWorkloadMetricExpressionRuleMetricExpressionRuleTypeEnum string

// Set of constants representing the allowable values for ManagedComputeClusterWorkloadMetricExpressionRuleMetricExpressionRuleTypeEnum
const (
	ManagedComputeClusterWorkloadMetricExpressionRuleMetricExpressionRuleTypePredefinedExpression ManagedComputeClusterWorkloadMetricExpressionRuleMetricExpressionRuleTypeEnum = "TARGET_PREDEFINED_EXPRESSION"
	ManagedComputeClusterWorkloadMetricExpressionRuleMetricExpressionRuleTypeCustomExpression     ManagedComputeClusterWorkloadMetricExpressionRuleMetricExpressionRuleTypeEnum = "TARGET_CUSTOM_EXPRESSION"
)

var mappingManagedComputeClusterWorkloadMetricExpressionRuleMetricExpressionRuleTypeEnum = map[string]ManagedComputeClusterWorkloadMetricExpressionRuleMetricExpressionRuleTypeEnum{
	"TARGET_PREDEFINED_EXPRESSION": ManagedComputeClusterWorkloadMetricExpressionRuleMetricExpressionRuleTypePredefinedExpression,
	"TARGET_CUSTOM_EXPRESSION":     ManagedComputeClusterWorkloadMetricExpressionRuleMetricExpressionRuleTypeCustomExpression,
}

var mappingManagedComputeClusterWorkloadMetricExpressionRuleMetricExpressionRuleTypeEnumLowerCase = map[string]ManagedComputeClusterWorkloadMetricExpressionRuleMetricExpressionRuleTypeEnum{
	"target_predefined_expression": ManagedComputeClusterWorkloadMetricExpressionRuleMetricExpressionRuleTypePredefinedExpression,
	"target_custom_expression":     ManagedComputeClusterWorkloadMetricExpressionRuleMetricExpressionRuleTypeCustomExpression,
}

// GetManagedComputeClusterWorkloadMetricExpressionRuleMetricExpressionRuleTypeEnumValues Enumerates the set of values for ManagedComputeClusterWorkloadMetricExpressionRuleMetricExpressionRuleTypeEnum
func GetManagedComputeClusterWorkloadMetricExpressionRuleMetricExpressionRuleTypeEnumValues() []ManagedComputeClusterWorkloadMetricExpressionRuleMetricExpressionRuleTypeEnum {
	values := make([]ManagedComputeClusterWorkloadMetricExpressionRuleMetricExpressionRuleTypeEnum, 0)
	for _, v := range mappingManagedComputeClusterWorkloadMetricExpressionRuleMetricExpressionRuleTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetManagedComputeClusterWorkloadMetricExpressionRuleMetricExpressionRuleTypeEnumStringValues Enumerates the set of values in String for ManagedComputeClusterWorkloadMetricExpressionRuleMetricExpressionRuleTypeEnum
func GetManagedComputeClusterWorkloadMetricExpressionRuleMetricExpressionRuleTypeEnumStringValues() []string {
	return []string{
		"TARGET_PREDEFINED_EXPRESSION",
		"TARGET_CUSTOM_EXPRESSION",
	}
}

// GetMappingManagedComputeClusterWorkloadMetricExpressionRuleMetricExpressionRuleTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingManagedComputeClusterWorkloadMetricExpressionRuleMetricExpressionRuleTypeEnum(val string) (ManagedComputeClusterWorkloadMetricExpressionRuleMetricExpressionRuleTypeEnum, bool) {
	enum, ok := mappingManagedComputeClusterWorkloadMetricExpressionRuleMetricExpressionRuleTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
