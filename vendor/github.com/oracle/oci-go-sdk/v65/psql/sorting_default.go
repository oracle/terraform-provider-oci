// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// PGSQL Control Plane API
//
// Use the OCI Database with PostgreSQL API to manage resources such as database systems, database nodes, backups, and configurations.
// For information, see the user guide documentation for the service (https://docs.oracle.com/iaas/Content/postgresql/home.htm).
//

package psql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SortingDefault Default sorting behavior for an insight type.
type SortingDefault struct {

	// Default field used for sorting.
	Field *string `mandatory:"false" json:"field"`

	// Default sort order.
	Order SortingDefaultOrderEnum `mandatory:"false" json:"order,omitempty"`
}

func (m SortingDefault) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SortingDefault) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSortingDefaultOrderEnum(string(m.Order)); !ok && m.Order != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Order: %s. Supported values are: %s.", m.Order, strings.Join(GetSortingDefaultOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SortingDefaultOrderEnum Enum with underlying type: string
type SortingDefaultOrderEnum string

// Set of constants representing the allowable values for SortingDefaultOrderEnum
const (
	SortingDefaultOrderAsc  SortingDefaultOrderEnum = "ASC"
	SortingDefaultOrderDesc SortingDefaultOrderEnum = "DESC"
)

var mappingSortingDefaultOrderEnum = map[string]SortingDefaultOrderEnum{
	"ASC":  SortingDefaultOrderAsc,
	"DESC": SortingDefaultOrderDesc,
}

var mappingSortingDefaultOrderEnumLowerCase = map[string]SortingDefaultOrderEnum{
	"asc":  SortingDefaultOrderAsc,
	"desc": SortingDefaultOrderDesc,
}

// GetSortingDefaultOrderEnumValues Enumerates the set of values for SortingDefaultOrderEnum
func GetSortingDefaultOrderEnumValues() []SortingDefaultOrderEnum {
	values := make([]SortingDefaultOrderEnum, 0)
	for _, v := range mappingSortingDefaultOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetSortingDefaultOrderEnumStringValues Enumerates the set of values in String for SortingDefaultOrderEnum
func GetSortingDefaultOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingSortingDefaultOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSortingDefaultOrderEnum(val string) (SortingDefaultOrderEnum, bool) {
	enum, ok := mappingSortingDefaultOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
