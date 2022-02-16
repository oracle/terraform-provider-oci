// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

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

	if _, ok := GetMappingSortClauseOrderEnum(string(m.Order)); !ok && m.Order != "" {
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

// GetMappingSortClauseOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSortClauseOrderEnum(val string) (SortClauseOrderEnum, bool) {
	mappingSortClauseOrderEnumIgnoreCase := make(map[string]SortClauseOrderEnum)
	for k, v := range mappingSortClauseOrderEnum {
		mappingSortClauseOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingSortClauseOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
