// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Usage API
//
// A description of the UsageApi API.
//

package usageapi

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Filter The filter object for query usage.
type Filter struct {

	// The operator of the filter. Example: 'AND', 'OR', 'NOT'.
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

// FilterOperatorEnum Enum with underlying type: string
type FilterOperatorEnum string

// Set of constants representing the allowable values for FilterOperatorEnum
const (
	FilterOperatorAnd FilterOperatorEnum = "AND"
	FilterOperatorNot FilterOperatorEnum = "NOT"
	FilterOperatorOr  FilterOperatorEnum = "OR"
)

var mappingFilterOperator = map[string]FilterOperatorEnum{
	"AND": FilterOperatorAnd,
	"NOT": FilterOperatorNot,
	"OR":  FilterOperatorOr,
}

// GetFilterOperatorEnumValues Enumerates the set of values for FilterOperatorEnum
func GetFilterOperatorEnumValues() []FilterOperatorEnum {
	values := make([]FilterOperatorEnum, 0)
	for _, v := range mappingFilterOperator {
		values = append(values, v)
	}
	return values
}
