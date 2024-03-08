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

// MetricExpressionRule The metric expression rule base.
type MetricExpressionRule interface {
}

type metricexpressionrule struct {
	JsonData                 []byte
	MetricExpressionRuleType string `json:"metricExpressionRuleType"`
}

// UnmarshalJSON unmarshals json
func (m *metricexpressionrule) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalermetricexpressionrule metricexpressionrule
	s := struct {
		Model Unmarshalermetricexpressionrule
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.MetricExpressionRuleType = s.Model.MetricExpressionRuleType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *metricexpressionrule) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.MetricExpressionRuleType {
	case "PREDEFINED_EXPRESSION":
		mm := PredefinedMetricExpressionRule{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "CUSTOM_EXPRESSION":
		mm := CustomMetricExpressionRule{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for MetricExpressionRule: %s.", m.MetricExpressionRuleType)
		return *m, nil
	}
}

func (m metricexpressionrule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m metricexpressionrule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MetricExpressionRuleMetricExpressionRuleTypeEnum Enum with underlying type: string
type MetricExpressionRuleMetricExpressionRuleTypeEnum string

// Set of constants representing the allowable values for MetricExpressionRuleMetricExpressionRuleTypeEnum
const (
	MetricExpressionRuleMetricExpressionRuleTypePredefinedExpression MetricExpressionRuleMetricExpressionRuleTypeEnum = "PREDEFINED_EXPRESSION"
	MetricExpressionRuleMetricExpressionRuleTypeCustomExpression     MetricExpressionRuleMetricExpressionRuleTypeEnum = "CUSTOM_EXPRESSION"
)

var mappingMetricExpressionRuleMetricExpressionRuleTypeEnum = map[string]MetricExpressionRuleMetricExpressionRuleTypeEnum{
	"PREDEFINED_EXPRESSION": MetricExpressionRuleMetricExpressionRuleTypePredefinedExpression,
	"CUSTOM_EXPRESSION":     MetricExpressionRuleMetricExpressionRuleTypeCustomExpression,
}

var mappingMetricExpressionRuleMetricExpressionRuleTypeEnumLowerCase = map[string]MetricExpressionRuleMetricExpressionRuleTypeEnum{
	"predefined_expression": MetricExpressionRuleMetricExpressionRuleTypePredefinedExpression,
	"custom_expression":     MetricExpressionRuleMetricExpressionRuleTypeCustomExpression,
}

// GetMetricExpressionRuleMetricExpressionRuleTypeEnumValues Enumerates the set of values for MetricExpressionRuleMetricExpressionRuleTypeEnum
func GetMetricExpressionRuleMetricExpressionRuleTypeEnumValues() []MetricExpressionRuleMetricExpressionRuleTypeEnum {
	values := make([]MetricExpressionRuleMetricExpressionRuleTypeEnum, 0)
	for _, v := range mappingMetricExpressionRuleMetricExpressionRuleTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetMetricExpressionRuleMetricExpressionRuleTypeEnumStringValues Enumerates the set of values in String for MetricExpressionRuleMetricExpressionRuleTypeEnum
func GetMetricExpressionRuleMetricExpressionRuleTypeEnumStringValues() []string {
	return []string{
		"PREDEFINED_EXPRESSION",
		"CUSTOM_EXPRESSION",
	}
}

// GetMappingMetricExpressionRuleMetricExpressionRuleTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMetricExpressionRuleMetricExpressionRuleTypeEnum(val string) (MetricExpressionRuleMetricExpressionRuleTypeEnum, bool) {
	enum, ok := mappingMetricExpressionRuleMetricExpressionRuleTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
