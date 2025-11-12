// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OrderClause report filters.
type OrderClause struct {

	// Sort By.
	SortBy *string `mandatory:"true" json:"sortBy"`

	// Sort Order.
	SortOrder OrderClauseSortOrderEnum `mandatory:"false" json:"sortOrder,omitempty"`
}

func (m OrderClause) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OrderClause) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOrderClauseSortOrderEnum(string(m.SortOrder)); !ok && m.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", m.SortOrder, strings.Join(GetOrderClauseSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OrderClauseSortOrderEnum Enum with underlying type: string
type OrderClauseSortOrderEnum string

// Set of constants representing the allowable values for OrderClauseSortOrderEnum
const (
	OrderClauseSortOrderAsc  OrderClauseSortOrderEnum = "ASC"
	OrderClauseSortOrderDesc OrderClauseSortOrderEnum = "DESC"
)

var mappingOrderClauseSortOrderEnum = map[string]OrderClauseSortOrderEnum{
	"ASC":  OrderClauseSortOrderAsc,
	"DESC": OrderClauseSortOrderDesc,
}

var mappingOrderClauseSortOrderEnumLowerCase = map[string]OrderClauseSortOrderEnum{
	"asc":  OrderClauseSortOrderAsc,
	"desc": OrderClauseSortOrderDesc,
}

// GetOrderClauseSortOrderEnumValues Enumerates the set of values for OrderClauseSortOrderEnum
func GetOrderClauseSortOrderEnumValues() []OrderClauseSortOrderEnum {
	values := make([]OrderClauseSortOrderEnum, 0)
	for _, v := range mappingOrderClauseSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetOrderClauseSortOrderEnumStringValues Enumerates the set of values in String for OrderClauseSortOrderEnum
func GetOrderClauseSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingOrderClauseSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOrderClauseSortOrderEnum(val string) (OrderClauseSortOrderEnum, bool) {
	enum, ok := mappingOrderClauseSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
