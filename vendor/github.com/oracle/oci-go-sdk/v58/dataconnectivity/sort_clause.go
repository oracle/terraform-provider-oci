// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Connectivity Management API
//
// Use the DCMS APIs to perform Metadata/Data operations.
//

package dataconnectivity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// SortClause The information about the sort object.
type SortClause struct {
	Field *ShapeField `mandatory:"false" json:"field"`

	// The sort order.
	Order SortClauseOrderEnum `mandatory:"false" json:"order,omitempty"`
}

func (m SortClause) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SortClause) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := mappingSortClauseOrderEnum[string(m.Order)]; !ok && m.Order != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Order: %s. Supported values are: %s.", m.Order, strings.Join(GetSortClauseOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SortClauseOrderEnum Enum with underlying type: string
type SortClauseOrderEnum string

// Set of constants representing the allowable values for SortClauseOrderEnum
const (
	SortClauseOrderAsc  SortClauseOrderEnum = "ASC"
	SortClauseOrderDesc SortClauseOrderEnum = "DESC"
)

var mappingSortClauseOrderEnum = map[string]SortClauseOrderEnum{
	"ASC":  SortClauseOrderAsc,
	"DESC": SortClauseOrderDesc,
}

// GetSortClauseOrderEnumValues Enumerates the set of values for SortClauseOrderEnum
func GetSortClauseOrderEnumValues() []SortClauseOrderEnum {
	values := make([]SortClauseOrderEnum, 0)
	for _, v := range mappingSortClauseOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetSortClauseOrderEnumStringValues Enumerates the set of values in String for SortClauseOrderEnum
func GetSortClauseOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}
