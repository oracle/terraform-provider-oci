// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// Filter Query builder filter action to apply edit to queryString.
type Filter struct {

	// Operator to apply when editing the query string.
	Operator FilterOperatorEnum `mandatory:"true" json:"operator"`

	// Field filter references when inserting filter into the query string. Field must be a valid logging analytics out-of-the-box field, virtual field calculated in the query or a user defined field.
	FieldName *string `mandatory:"false" json:"fieldName"`

	// Field values that will be inserted into the query string for the specified fieldName. Please note all values should reflect the fields data type otherwise the insert is subject to fail.
	Values []interface{} `mandatory:"false" json:"values"`
}

func (m Filter) String() string {
	return common.PointerString(m)
}

// FilterOperatorEnum Enum with underlying type: string
type FilterOperatorEnum string

// Set of constants representing the allowable values for FilterOperatorEnum
const (
	FilterOperatorClear                  FilterOperatorEnum = "CLEAR"
	FilterOperatorReplace                FilterOperatorEnum = "REPLACE"
	FilterOperatorEquals                 FilterOperatorEnum = "EQUALS"
	FilterOperatorNotEquals              FilterOperatorEnum = "NOT_EQUALS"
	FilterOperatorStartsWith             FilterOperatorEnum = "STARTS_WITH"
	FilterOperatorDoesNotStartWith       FilterOperatorEnum = "DOES_NOT_START_WITH"
	FilterOperatorEndsWith               FilterOperatorEnum = "ENDS_WITH"
	FilterOperatorDoesNotEndWith         FilterOperatorEnum = "DOES_NOT_END_WITH"
	FilterOperatorContains               FilterOperatorEnum = "CONTAINS"
	FilterOperatorDoesNotContain         FilterOperatorEnum = "DOES_NOT_CONTAIN"
	FilterOperatorIsLessThan             FilterOperatorEnum = "IS_LESS_THAN"
	FilterOperatorIsLessThanOrEqualTo    FilterOperatorEnum = "IS_LESS_THAN_OR_EQUAL_TO"
	FilterOperatorIsGreaterThan          FilterOperatorEnum = "IS_GREATER_THAN"
	FilterOperatorIsGreaterThanOrEqualTo FilterOperatorEnum = "IS_GREATER_THAN_OR_EQUAL_TO"
	FilterOperatorIsBetween              FilterOperatorEnum = "IS_BETWEEN"
	FilterOperatorIsNotBetween           FilterOperatorEnum = "IS_NOT_BETWEEN"
	FilterOperatorAddSubquery            FilterOperatorEnum = "ADD_SUBQUERY"
	FilterOperatorClearSubquery          FilterOperatorEnum = "CLEAR_SUBQUERY"
)

var mappingFilterOperator = map[string]FilterOperatorEnum{
	"CLEAR":                       FilterOperatorClear,
	"REPLACE":                     FilterOperatorReplace,
	"EQUALS":                      FilterOperatorEquals,
	"NOT_EQUALS":                  FilterOperatorNotEquals,
	"STARTS_WITH":                 FilterOperatorStartsWith,
	"DOES_NOT_START_WITH":         FilterOperatorDoesNotStartWith,
	"ENDS_WITH":                   FilterOperatorEndsWith,
	"DOES_NOT_END_WITH":           FilterOperatorDoesNotEndWith,
	"CONTAINS":                    FilterOperatorContains,
	"DOES_NOT_CONTAIN":            FilterOperatorDoesNotContain,
	"IS_LESS_THAN":                FilterOperatorIsLessThan,
	"IS_LESS_THAN_OR_EQUAL_TO":    FilterOperatorIsLessThanOrEqualTo,
	"IS_GREATER_THAN":             FilterOperatorIsGreaterThan,
	"IS_GREATER_THAN_OR_EQUAL_TO": FilterOperatorIsGreaterThanOrEqualTo,
	"IS_BETWEEN":                  FilterOperatorIsBetween,
	"IS_NOT_BETWEEN":              FilterOperatorIsNotBetween,
	"ADD_SUBQUERY":                FilterOperatorAddSubquery,
	"CLEAR_SUBQUERY":              FilterOperatorClearSubquery,
}

// GetFilterOperatorEnumValues Enumerates the set of values for FilterOperatorEnum
func GetFilterOperatorEnumValues() []FilterOperatorEnum {
	values := make([]FilterOperatorEnum, 0)
	for _, v := range mappingFilterOperator {
		values = append(values, v)
	}
	return values
}
