// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SingleValueFilter Content Source details.
type SingleValueFilter struct {

	// Name of report.
	FieldName *string `mandatory:"true" json:"fieldName"`

	// Value.
	Value *string `mandatory:"true" json:"value"`

	// Expression for Filter
	Expression SingleValueFilterExpressionEnum `mandatory:"true" json:"expression"`
}

// GetFieldName returns FieldName
func (m SingleValueFilter) GetFieldName() *string {
	return m.FieldName
}

func (m SingleValueFilter) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SingleValueFilter) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSingleValueFilterExpressionEnum(string(m.Expression)); !ok && m.Expression != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Expression: %s. Supported values are: %s.", m.Expression, strings.Join(GetSingleValueFilterExpressionEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m SingleValueFilter) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeSingleValueFilter SingleValueFilter
	s := struct {
		DiscriminatorParam string `json:"filterType"`
		MarshalTypeSingleValueFilter
	}{
		"SINGLE_VALUE",
		(MarshalTypeSingleValueFilter)(m),
	}

	return json.Marshal(&s)
}

// SingleValueFilterExpressionEnum Enum with underlying type: string
type SingleValueFilterExpressionEnum string

// Set of constants representing the allowable values for SingleValueFilterExpressionEnum
const (
	SingleValueFilterExpressionEqual                SingleValueFilterExpressionEnum = "EQUAL"
	SingleValueFilterExpressionNotEqual             SingleValueFilterExpressionEnum = "NOT_EQUAL"
	SingleValueFilterExpressionLessThan             SingleValueFilterExpressionEnum = "LESS_THAN"
	SingleValueFilterExpressionGreaterThan          SingleValueFilterExpressionEnum = "GREATER_THAN"
	SingleValueFilterExpressionLessThanOrEqualTo    SingleValueFilterExpressionEnum = "LESS_THAN_OR_EQUAL_TO"
	SingleValueFilterExpressionGreaterThanOrEqualTo SingleValueFilterExpressionEnum = "GREATER_THAN_OR_EQUAL_TO"
)

var mappingSingleValueFilterExpressionEnum = map[string]SingleValueFilterExpressionEnum{
	"EQUAL":                    SingleValueFilterExpressionEqual,
	"NOT_EQUAL":                SingleValueFilterExpressionNotEqual,
	"LESS_THAN":                SingleValueFilterExpressionLessThan,
	"GREATER_THAN":             SingleValueFilterExpressionGreaterThan,
	"LESS_THAN_OR_EQUAL_TO":    SingleValueFilterExpressionLessThanOrEqualTo,
	"GREATER_THAN_OR_EQUAL_TO": SingleValueFilterExpressionGreaterThanOrEqualTo,
}

var mappingSingleValueFilterExpressionEnumLowerCase = map[string]SingleValueFilterExpressionEnum{
	"equal":                    SingleValueFilterExpressionEqual,
	"not_equal":                SingleValueFilterExpressionNotEqual,
	"less_than":                SingleValueFilterExpressionLessThan,
	"greater_than":             SingleValueFilterExpressionGreaterThan,
	"less_than_or_equal_to":    SingleValueFilterExpressionLessThanOrEqualTo,
	"greater_than_or_equal_to": SingleValueFilterExpressionGreaterThanOrEqualTo,
}

// GetSingleValueFilterExpressionEnumValues Enumerates the set of values for SingleValueFilterExpressionEnum
func GetSingleValueFilterExpressionEnumValues() []SingleValueFilterExpressionEnum {
	values := make([]SingleValueFilterExpressionEnum, 0)
	for _, v := range mappingSingleValueFilterExpressionEnum {
		values = append(values, v)
	}
	return values
}

// GetSingleValueFilterExpressionEnumStringValues Enumerates the set of values in String for SingleValueFilterExpressionEnum
func GetSingleValueFilterExpressionEnumStringValues() []string {
	return []string{
		"EQUAL",
		"NOT_EQUAL",
		"LESS_THAN",
		"GREATER_THAN",
		"LESS_THAN_OR_EQUAL_TO",
		"GREATER_THAN_OR_EQUAL_TO",
	}
}

// GetMappingSingleValueFilterExpressionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSingleValueFilterExpressionEnum(val string) (SingleValueFilterExpressionEnum, bool) {
	enum, ok := mappingSingleValueFilterExpressionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
