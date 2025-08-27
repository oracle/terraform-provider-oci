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

// Column The description of the column.
type Column struct {

	// Name of the column displayed on UI.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Specifies the corresponding field name in the data source.
	FieldName *string `mandatory:"true" json:"fieldName"`

	// Indicates if the column is hidden. Values can either be 'true' or 'false'.
	IsHidden *bool `mandatory:"true" json:"isHidden"`

	// Specifies the display order of the column.
	DisplayOrder *int `mandatory:"true" json:"displayOrder"`

	// Specifies the data type of the column.
	DataType *string `mandatory:"false" json:"dataType"`

	// Specifies if column is virtual and can only be used as column filter.
	IsVirtual *bool `mandatory:"false" json:"isVirtual"`

	// An array of operators that can be supported by column fieldName.
	ApplicableOperators []ColumnApplicableOperatorsEnum `mandatory:"false" json:"applicableOperators,omitempty"`
}

func (m Column) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Column) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	for _, val := range m.ApplicableOperators {
		if _, ok := GetMappingColumnApplicableOperatorsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ApplicableOperators: %s. Supported values are: %s.", val, strings.Join(GetColumnApplicableOperatorsEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ColumnApplicableOperatorsEnum Enum with underlying type: string
type ColumnApplicableOperatorsEnum string

// Set of constants representing the allowable values for ColumnApplicableOperatorsEnum
const (
	ColumnApplicableOperatorsIn       ColumnApplicableOperatorsEnum = "IN"
	ColumnApplicableOperatorsEq       ColumnApplicableOperatorsEnum = "EQ"
	ColumnApplicableOperatorsEqCs     ColumnApplicableOperatorsEnum = "EQ_CS"
	ColumnApplicableOperatorsGt       ColumnApplicableOperatorsEnum = "GT"
	ColumnApplicableOperatorsGe       ColumnApplicableOperatorsEnum = "GE"
	ColumnApplicableOperatorsLt       ColumnApplicableOperatorsEnum = "LT"
	ColumnApplicableOperatorsLe       ColumnApplicableOperatorsEnum = "LE"
	ColumnApplicableOperatorsAnd      ColumnApplicableOperatorsEnum = "AND"
	ColumnApplicableOperatorsOr       ColumnApplicableOperatorsEnum = "OR"
	ColumnApplicableOperatorsNe       ColumnApplicableOperatorsEnum = "NE"
	ColumnApplicableOperatorsCo       ColumnApplicableOperatorsEnum = "CO"
	ColumnApplicableOperatorsCoCs     ColumnApplicableOperatorsEnum = "CO_CS"
	ColumnApplicableOperatorsNot      ColumnApplicableOperatorsEnum = "NOT"
	ColumnApplicableOperatorsNotIn    ColumnApplicableOperatorsEnum = "NOT_IN"
	ColumnApplicableOperatorsInSet    ColumnApplicableOperatorsEnum = "IN_SET"
	ColumnApplicableOperatorsNotInSet ColumnApplicableOperatorsEnum = "NOT_IN_SET"
)

var mappingColumnApplicableOperatorsEnum = map[string]ColumnApplicableOperatorsEnum{
	"IN":         ColumnApplicableOperatorsIn,
	"EQ":         ColumnApplicableOperatorsEq,
	"EQ_CS":      ColumnApplicableOperatorsEqCs,
	"GT":         ColumnApplicableOperatorsGt,
	"GE":         ColumnApplicableOperatorsGe,
	"LT":         ColumnApplicableOperatorsLt,
	"LE":         ColumnApplicableOperatorsLe,
	"AND":        ColumnApplicableOperatorsAnd,
	"OR":         ColumnApplicableOperatorsOr,
	"NE":         ColumnApplicableOperatorsNe,
	"CO":         ColumnApplicableOperatorsCo,
	"CO_CS":      ColumnApplicableOperatorsCoCs,
	"NOT":        ColumnApplicableOperatorsNot,
	"NOT_IN":     ColumnApplicableOperatorsNotIn,
	"IN_SET":     ColumnApplicableOperatorsInSet,
	"NOT_IN_SET": ColumnApplicableOperatorsNotInSet,
}

var mappingColumnApplicableOperatorsEnumLowerCase = map[string]ColumnApplicableOperatorsEnum{
	"in":         ColumnApplicableOperatorsIn,
	"eq":         ColumnApplicableOperatorsEq,
	"eq_cs":      ColumnApplicableOperatorsEqCs,
	"gt":         ColumnApplicableOperatorsGt,
	"ge":         ColumnApplicableOperatorsGe,
	"lt":         ColumnApplicableOperatorsLt,
	"le":         ColumnApplicableOperatorsLe,
	"and":        ColumnApplicableOperatorsAnd,
	"or":         ColumnApplicableOperatorsOr,
	"ne":         ColumnApplicableOperatorsNe,
	"co":         ColumnApplicableOperatorsCo,
	"co_cs":      ColumnApplicableOperatorsCoCs,
	"not":        ColumnApplicableOperatorsNot,
	"not_in":     ColumnApplicableOperatorsNotIn,
	"in_set":     ColumnApplicableOperatorsInSet,
	"not_in_set": ColumnApplicableOperatorsNotInSet,
}

// GetColumnApplicableOperatorsEnumValues Enumerates the set of values for ColumnApplicableOperatorsEnum
func GetColumnApplicableOperatorsEnumValues() []ColumnApplicableOperatorsEnum {
	values := make([]ColumnApplicableOperatorsEnum, 0)
	for _, v := range mappingColumnApplicableOperatorsEnum {
		values = append(values, v)
	}
	return values
}

// GetColumnApplicableOperatorsEnumStringValues Enumerates the set of values in String for ColumnApplicableOperatorsEnum
func GetColumnApplicableOperatorsEnumStringValues() []string {
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
		"IN_SET",
		"NOT_IN_SET",
	}
}

// GetMappingColumnApplicableOperatorsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingColumnApplicableOperatorsEnum(val string) (ColumnApplicableOperatorsEnum, bool) {
	enum, ok := mappingColumnApplicableOperatorsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
