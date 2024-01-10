// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ConditionBlock A condition block. This could represent a single condition, or have nested condition blocks under it.
// To form a single condition, specify the fieldName, labelConditionOperator and labelConditionValue(s).
// To form nested conditions, specify the conditions in conditionBlocks, and how to join them in conditionBlocksOperator.
type ConditionBlock struct {

	// Operator using which the conditionBlocks should be joined. Specify this for nested conditions.
	ConditionBlocksOperator ConditionBlockConditionBlocksOperatorEnum `mandatory:"false" json:"conditionBlocksOperator,omitempty"`

	// The name of the field the condition is based on. Specify this if this condition block represents a single condition.
	FieldName *string `mandatory:"false" json:"fieldName"`

	// The condition operator. Specify this if this condition block represents a single condition.
	LabelConditionOperator *string `mandatory:"false" json:"labelConditionOperator"`

	// The condition value. Specify this if this condition block represents a single condition.
	LabelConditionValue *string `mandatory:"false" json:"labelConditionValue"`

	// A list of condition values. Specify this if this condition block represents a single condition.
	LabelConditionValues []string `mandatory:"false" json:"labelConditionValues"`

	// Condition blocks to evaluate within this condition block. Specify this for nested conditions.
	ConditionBlocks []ConditionBlock `mandatory:"false" json:"conditionBlocks"`
}

func (m ConditionBlock) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ConditionBlock) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingConditionBlockConditionBlocksOperatorEnum(string(m.ConditionBlocksOperator)); !ok && m.ConditionBlocksOperator != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ConditionBlocksOperator: %s. Supported values are: %s.", m.ConditionBlocksOperator, strings.Join(GetConditionBlockConditionBlocksOperatorEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ConditionBlockConditionBlocksOperatorEnum Enum with underlying type: string
type ConditionBlockConditionBlocksOperatorEnum string

// Set of constants representing the allowable values for ConditionBlockConditionBlocksOperatorEnum
const (
	ConditionBlockConditionBlocksOperatorAnd    ConditionBlockConditionBlocksOperatorEnum = "AND"
	ConditionBlockConditionBlocksOperatorOr     ConditionBlockConditionBlocksOperatorEnum = "OR"
	ConditionBlockConditionBlocksOperatorNotAnd ConditionBlockConditionBlocksOperatorEnum = "NOT_AND"
	ConditionBlockConditionBlocksOperatorNotOr  ConditionBlockConditionBlocksOperatorEnum = "NOT_OR"
)

var mappingConditionBlockConditionBlocksOperatorEnum = map[string]ConditionBlockConditionBlocksOperatorEnum{
	"AND":     ConditionBlockConditionBlocksOperatorAnd,
	"OR":      ConditionBlockConditionBlocksOperatorOr,
	"NOT_AND": ConditionBlockConditionBlocksOperatorNotAnd,
	"NOT_OR":  ConditionBlockConditionBlocksOperatorNotOr,
}

var mappingConditionBlockConditionBlocksOperatorEnumLowerCase = map[string]ConditionBlockConditionBlocksOperatorEnum{
	"and":     ConditionBlockConditionBlocksOperatorAnd,
	"or":      ConditionBlockConditionBlocksOperatorOr,
	"not_and": ConditionBlockConditionBlocksOperatorNotAnd,
	"not_or":  ConditionBlockConditionBlocksOperatorNotOr,
}

// GetConditionBlockConditionBlocksOperatorEnumValues Enumerates the set of values for ConditionBlockConditionBlocksOperatorEnum
func GetConditionBlockConditionBlocksOperatorEnumValues() []ConditionBlockConditionBlocksOperatorEnum {
	values := make([]ConditionBlockConditionBlocksOperatorEnum, 0)
	for _, v := range mappingConditionBlockConditionBlocksOperatorEnum {
		values = append(values, v)
	}
	return values
}

// GetConditionBlockConditionBlocksOperatorEnumStringValues Enumerates the set of values in String for ConditionBlockConditionBlocksOperatorEnum
func GetConditionBlockConditionBlocksOperatorEnumStringValues() []string {
	return []string{
		"AND",
		"OR",
		"NOT_AND",
		"NOT_OR",
	}
}

// GetMappingConditionBlockConditionBlocksOperatorEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConditionBlockConditionBlocksOperatorEnum(val string) (ConditionBlockConditionBlocksOperatorEnum, bool) {
	enum, ok := mappingConditionBlockConditionBlocksOperatorEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
