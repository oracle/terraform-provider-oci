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

// ManagedComputeClusterMetricExpressionRule The metric expression rule base.
type ManagedComputeClusterMetricExpressionRule interface {
}

type managedcomputeclustermetricexpressionrule struct {
	JsonData                 []byte
	MetricExpressionRuleType string `json:"metricExpressionRuleType"`
}

// UnmarshalJSON unmarshals json
func (m *managedcomputeclustermetricexpressionrule) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalermanagedcomputeclustermetricexpressionrule managedcomputeclustermetricexpressionrule
	s := struct {
		Model Unmarshalermanagedcomputeclustermetricexpressionrule
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.MetricExpressionRuleType = s.Model.MetricExpressionRuleType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *managedcomputeclustermetricexpressionrule) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.MetricExpressionRuleType {
	case "PREDEFINED_EXPRESSION":
		mm := ManagedComputeClusterPredefinedMetricExpressionRule{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CUSTOM_EXPRESSION":
		mm := ManagedComputeClusterCustomMetricExpressionRule{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for ManagedComputeClusterMetricExpressionRule: %s.", m.MetricExpressionRuleType)
		return *m, nil
	}
}

func (m managedcomputeclustermetricexpressionrule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m managedcomputeclustermetricexpressionrule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ManagedComputeClusterMetricExpressionRuleMetricExpressionRuleTypeEnum Enum with underlying type: string
type ManagedComputeClusterMetricExpressionRuleMetricExpressionRuleTypeEnum string

// Set of constants representing the allowable values for ManagedComputeClusterMetricExpressionRuleMetricExpressionRuleTypeEnum
const (
	ManagedComputeClusterMetricExpressionRuleMetricExpressionRuleTypePredefinedExpression ManagedComputeClusterMetricExpressionRuleMetricExpressionRuleTypeEnum = "PREDEFINED_EXPRESSION"
	ManagedComputeClusterMetricExpressionRuleMetricExpressionRuleTypeCustomExpression     ManagedComputeClusterMetricExpressionRuleMetricExpressionRuleTypeEnum = "CUSTOM_EXPRESSION"
)

var mappingManagedComputeClusterMetricExpressionRuleMetricExpressionRuleTypeEnum = map[string]ManagedComputeClusterMetricExpressionRuleMetricExpressionRuleTypeEnum{
	"PREDEFINED_EXPRESSION": ManagedComputeClusterMetricExpressionRuleMetricExpressionRuleTypePredefinedExpression,
	"CUSTOM_EXPRESSION":     ManagedComputeClusterMetricExpressionRuleMetricExpressionRuleTypeCustomExpression,
}

var mappingManagedComputeClusterMetricExpressionRuleMetricExpressionRuleTypeEnumLowerCase = map[string]ManagedComputeClusterMetricExpressionRuleMetricExpressionRuleTypeEnum{
	"predefined_expression": ManagedComputeClusterMetricExpressionRuleMetricExpressionRuleTypePredefinedExpression,
	"custom_expression":     ManagedComputeClusterMetricExpressionRuleMetricExpressionRuleTypeCustomExpression,
}

// GetManagedComputeClusterMetricExpressionRuleMetricExpressionRuleTypeEnumValues Enumerates the set of values for ManagedComputeClusterMetricExpressionRuleMetricExpressionRuleTypeEnum
func GetManagedComputeClusterMetricExpressionRuleMetricExpressionRuleTypeEnumValues() []ManagedComputeClusterMetricExpressionRuleMetricExpressionRuleTypeEnum {
	values := make([]ManagedComputeClusterMetricExpressionRuleMetricExpressionRuleTypeEnum, 0)
	for _, v := range mappingManagedComputeClusterMetricExpressionRuleMetricExpressionRuleTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetManagedComputeClusterMetricExpressionRuleMetricExpressionRuleTypeEnumStringValues Enumerates the set of values in String for ManagedComputeClusterMetricExpressionRuleMetricExpressionRuleTypeEnum
func GetManagedComputeClusterMetricExpressionRuleMetricExpressionRuleTypeEnumStringValues() []string {
	return []string{
		"PREDEFINED_EXPRESSION",
		"CUSTOM_EXPRESSION",
	}
}

// GetMappingManagedComputeClusterMetricExpressionRuleMetricExpressionRuleTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingManagedComputeClusterMetricExpressionRuleMetricExpressionRuleTypeEnum(val string) (ManagedComputeClusterMetricExpressionRuleMetricExpressionRuleTypeEnum, bool) {
	enum, ok := mappingManagedComputeClusterMetricExpressionRuleMetricExpressionRuleTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
