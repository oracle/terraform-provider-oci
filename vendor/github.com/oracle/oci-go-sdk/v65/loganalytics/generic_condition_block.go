// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// GenericConditionBlock A condition block. This could represent a single condition, or have nested condition blocks under it.
// To form a single condition, specify the name, operator and value(s).
// To form nested conditions, specify the conditions in conditionBlocks, and how to join them in conditionBlocksOperator.
type GenericConditionBlock struct {

	// Operator using which the conditionBlocks should be joined. Specify this for nested conditions.
	ConditionBlocksOperator GenericConditionBlockConditionBlocksOperatorEnum `mandatory:"false" json:"conditionBlocksOperator,omitempty"`

	// The name of the field the condition is based on. Specify this if this condition block represents a single condition.
	Name *string `mandatory:"false" json:"name"`

	// The condition operator. Specify this if this condition block represents a single condition.
	Operator *string `mandatory:"false" json:"operator"`

	// The condition value. Specify this if this condition block represents a single condition.
	Value *string `mandatory:"false" json:"value"`

	// A list of condition values. Specify this if this condition block represents a single condition.
	Values []string `mandatory:"false" json:"values"`

	// Condition blocks to evaluate within this condition block. Specify this for nested conditions.
	GenericConditionBlocks []GenericConditionBlock `mandatory:"false" json:"genericConditionBlocks"`
}

func (m GenericConditionBlock) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GenericConditionBlock) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingGenericConditionBlockConditionBlocksOperatorEnum(string(m.ConditionBlocksOperator)); !ok && m.ConditionBlocksOperator != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ConditionBlocksOperator: %s. Supported values are: %s.", m.ConditionBlocksOperator, strings.Join(GetGenericConditionBlockConditionBlocksOperatorEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GenericConditionBlockConditionBlocksOperatorEnum Enum with underlying type: string
type GenericConditionBlockConditionBlocksOperatorEnum string

// Set of constants representing the allowable values for GenericConditionBlockConditionBlocksOperatorEnum
const (
	GenericConditionBlockConditionBlocksOperatorAnd    GenericConditionBlockConditionBlocksOperatorEnum = "AND"
	GenericConditionBlockConditionBlocksOperatorOr     GenericConditionBlockConditionBlocksOperatorEnum = "OR"
	GenericConditionBlockConditionBlocksOperatorNotAnd GenericConditionBlockConditionBlocksOperatorEnum = "NOT_AND"
	GenericConditionBlockConditionBlocksOperatorNotOr  GenericConditionBlockConditionBlocksOperatorEnum = "NOT_OR"
)

var mappingGenericConditionBlockConditionBlocksOperatorEnum = map[string]GenericConditionBlockConditionBlocksOperatorEnum{
	"AND":     GenericConditionBlockConditionBlocksOperatorAnd,
	"OR":      GenericConditionBlockConditionBlocksOperatorOr,
	"NOT_AND": GenericConditionBlockConditionBlocksOperatorNotAnd,
	"NOT_OR":  GenericConditionBlockConditionBlocksOperatorNotOr,
}

var mappingGenericConditionBlockConditionBlocksOperatorEnumLowerCase = map[string]GenericConditionBlockConditionBlocksOperatorEnum{
	"and":     GenericConditionBlockConditionBlocksOperatorAnd,
	"or":      GenericConditionBlockConditionBlocksOperatorOr,
	"not_and": GenericConditionBlockConditionBlocksOperatorNotAnd,
	"not_or":  GenericConditionBlockConditionBlocksOperatorNotOr,
}

// GetGenericConditionBlockConditionBlocksOperatorEnumValues Enumerates the set of values for GenericConditionBlockConditionBlocksOperatorEnum
func GetGenericConditionBlockConditionBlocksOperatorEnumValues() []GenericConditionBlockConditionBlocksOperatorEnum {
	values := make([]GenericConditionBlockConditionBlocksOperatorEnum, 0)
	for _, v := range mappingGenericConditionBlockConditionBlocksOperatorEnum {
		values = append(values, v)
	}
	return values
}

// GetGenericConditionBlockConditionBlocksOperatorEnumStringValues Enumerates the set of values in String for GenericConditionBlockConditionBlocksOperatorEnum
func GetGenericConditionBlockConditionBlocksOperatorEnumStringValues() []string {
	return []string{
		"AND",
		"OR",
		"NOT_AND",
		"NOT_OR",
	}
}

// GetMappingGenericConditionBlockConditionBlocksOperatorEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGenericConditionBlockConditionBlocksOperatorEnum(val string) (GenericConditionBlockConditionBlocksOperatorEnum, bool) {
	enum, ok := mappingGenericConditionBlockConditionBlocksOperatorEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
