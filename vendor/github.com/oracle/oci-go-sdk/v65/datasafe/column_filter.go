// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ColumnFilter Filters that are applied to the data at the column level.
type ColumnFilter struct {

	// Name of the column on which the filter must be applied.
	FieldName *string `mandatory:"true" json:"fieldName"`

	// Specifies the type of operator that must be applied for example in, eq etc.
	Operator ColumnFilterOperatorEnum `mandatory:"true" json:"operator"`

	// An array of expressions based on the operator type. A filter may have one or more expressions.
	Expressions []string `mandatory:"true" json:"expressions"`

	// Indicates whether the filter is enabled. Values can either be 'true' or 'false'.
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// Indicates whether the filter is hidden. Values can either be 'true' or 'false'.
	IsHidden *bool `mandatory:"true" json:"isHidden"`
}

func (m ColumnFilter) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ColumnFilter) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingColumnFilterOperatorEnum(string(m.Operator)); !ok && m.Operator != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Operator: %s. Supported values are: %s.", m.Operator, strings.Join(GetColumnFilterOperatorEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ColumnFilterOperatorEnum Enum with underlying type: string
type ColumnFilterOperatorEnum string

// Set of constants representing the allowable values for ColumnFilterOperatorEnum
const (
	ColumnFilterOperatorIn       ColumnFilterOperatorEnum = "IN"
	ColumnFilterOperatorEq       ColumnFilterOperatorEnum = "EQ"
	ColumnFilterOperatorEqCs     ColumnFilterOperatorEnum = "EQ_CS"
	ColumnFilterOperatorGt       ColumnFilterOperatorEnum = "GT"
	ColumnFilterOperatorGe       ColumnFilterOperatorEnum = "GE"
	ColumnFilterOperatorLt       ColumnFilterOperatorEnum = "LT"
	ColumnFilterOperatorLe       ColumnFilterOperatorEnum = "LE"
	ColumnFilterOperatorAnd      ColumnFilterOperatorEnum = "AND"
	ColumnFilterOperatorOr       ColumnFilterOperatorEnum = "OR"
	ColumnFilterOperatorNe       ColumnFilterOperatorEnum = "NE"
	ColumnFilterOperatorCo       ColumnFilterOperatorEnum = "CO"
	ColumnFilterOperatorCoCs     ColumnFilterOperatorEnum = "CO_CS"
	ColumnFilterOperatorNot      ColumnFilterOperatorEnum = "NOT"
	ColumnFilterOperatorNotIn    ColumnFilterOperatorEnum = "NOT_IN"
	ColumnFilterOperatorPr       ColumnFilterOperatorEnum = "PR"
	ColumnFilterOperatorInSet    ColumnFilterOperatorEnum = "IN_SET"
	ColumnFilterOperatorNotInSet ColumnFilterOperatorEnum = "NOT_IN_SET"
)

var mappingColumnFilterOperatorEnum = map[string]ColumnFilterOperatorEnum{
	"IN":         ColumnFilterOperatorIn,
	"EQ":         ColumnFilterOperatorEq,
	"EQ_CS":      ColumnFilterOperatorEqCs,
	"GT":         ColumnFilterOperatorGt,
	"GE":         ColumnFilterOperatorGe,
	"LT":         ColumnFilterOperatorLt,
	"LE":         ColumnFilterOperatorLe,
	"AND":        ColumnFilterOperatorAnd,
	"OR":         ColumnFilterOperatorOr,
	"NE":         ColumnFilterOperatorNe,
	"CO":         ColumnFilterOperatorCo,
	"CO_CS":      ColumnFilterOperatorCoCs,
	"NOT":        ColumnFilterOperatorNot,
	"NOT_IN":     ColumnFilterOperatorNotIn,
	"PR":         ColumnFilterOperatorPr,
	"IN_SET":     ColumnFilterOperatorInSet,
	"NOT_IN_SET": ColumnFilterOperatorNotInSet,
}

var mappingColumnFilterOperatorEnumLowerCase = map[string]ColumnFilterOperatorEnum{
	"in":         ColumnFilterOperatorIn,
	"eq":         ColumnFilterOperatorEq,
	"eq_cs":      ColumnFilterOperatorEqCs,
	"gt":         ColumnFilterOperatorGt,
	"ge":         ColumnFilterOperatorGe,
	"lt":         ColumnFilterOperatorLt,
	"le":         ColumnFilterOperatorLe,
	"and":        ColumnFilterOperatorAnd,
	"or":         ColumnFilterOperatorOr,
	"ne":         ColumnFilterOperatorNe,
	"co":         ColumnFilterOperatorCo,
	"co_cs":      ColumnFilterOperatorCoCs,
	"not":        ColumnFilterOperatorNot,
	"not_in":     ColumnFilterOperatorNotIn,
	"pr":         ColumnFilterOperatorPr,
	"in_set":     ColumnFilterOperatorInSet,
	"not_in_set": ColumnFilterOperatorNotInSet,
}

// GetColumnFilterOperatorEnumValues Enumerates the set of values for ColumnFilterOperatorEnum
func GetColumnFilterOperatorEnumValues() []ColumnFilterOperatorEnum {
	values := make([]ColumnFilterOperatorEnum, 0)
	for _, v := range mappingColumnFilterOperatorEnum {
		values = append(values, v)
	}
	return values
}

// GetColumnFilterOperatorEnumStringValues Enumerates the set of values in String for ColumnFilterOperatorEnum
func GetColumnFilterOperatorEnumStringValues() []string {
	return []string{
		"IN",
		"EQ",
		"EQ_CS",
		"GT",
		"GE",
		"LT",
		"LE",
		"AND",
		"OR",
		"NE",
		"CO",
		"CO_CS",
		"NOT",
		"NOT_IN",
		"PR",
		"IN_SET",
		"NOT_IN_SET",
	}
}

// GetMappingColumnFilterOperatorEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingColumnFilterOperatorEnum(val string) (ColumnFilterOperatorEnum, bool) {
	enum, ok := mappingColumnFilterOperatorEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
