// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Usage API
//
// Use the Usage API to view your Oracle Cloud usage and costs. The API allows you to request data that meets the specified filter criteria, and to group that data by the dimension of your choosing. The Usage API is used by the Cost Analysis tool in the Console. Also see Using the Usage API (https://docs.cloud.oracle.com/Content/Billing/Concepts/costanalysisoverview.htm#cost_analysis_using_the_api) for more information.
//

package usageapi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Filter The filter object for query usage.
type Filter struct {

	// The filter operator. Example: 'AND', 'OR', 'NOT'.
	Operator FilterOperatorEnum `mandatory:"false" json:"operator,omitempty"`

	// The dimensions to filter on.
	Dimensions []Dimension `mandatory:"false" json:"dimensions"`

	// The tags to filter on.
	Tags []Tag `mandatory:"false" json:"tags"`

	// The nested filter object.
	Filters []Filter `mandatory:"false" json:"filters"`
}

func (m Filter) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Filter) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingFilterOperatorEnum(string(m.Operator)); !ok && m.Operator != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Operator: %s. Supported values are: %s.", m.Operator, strings.Join(GetFilterOperatorEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// FilterOperatorEnum Enum with underlying type: string
type FilterOperatorEnum string

// Set of constants representing the allowable values for FilterOperatorEnum
const (
	FilterOperatorAnd FilterOperatorEnum = "AND"
	FilterOperatorNot FilterOperatorEnum = "NOT"
	FilterOperatorOr  FilterOperatorEnum = "OR"
)

var mappingFilterOperatorEnum = map[string]FilterOperatorEnum{
	"AND": FilterOperatorAnd,
	"NOT": FilterOperatorNot,
	"OR":  FilterOperatorOr,
}

var mappingFilterOperatorEnumLowerCase = map[string]FilterOperatorEnum{
	"and": FilterOperatorAnd,
	"not": FilterOperatorNot,
	"or":  FilterOperatorOr,
}

// GetFilterOperatorEnumValues Enumerates the set of values for FilterOperatorEnum
func GetFilterOperatorEnumValues() []FilterOperatorEnum {
	values := make([]FilterOperatorEnum, 0)
	for _, v := range mappingFilterOperatorEnum {
		values = append(values, v)
	}
	return values
}

// GetFilterOperatorEnumStringValues Enumerates the set of values in String for FilterOperatorEnum
func GetFilterOperatorEnumStringValues() []string {
	return []string{
		"AND",
		"NOT",
		"OR",
	}
}

// GetMappingFilterOperatorEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFilterOperatorEnum(val string) (FilterOperatorEnum, bool) {
	enum, ok := mappingFilterOperatorEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
