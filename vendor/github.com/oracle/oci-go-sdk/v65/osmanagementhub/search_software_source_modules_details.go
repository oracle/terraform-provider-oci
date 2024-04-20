// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for instances in OCI, your private data center, or 3rd-party clouds.
// For more information, see Overview of OS Management Hub (https://docs.cloud.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SearchSoftwareSourceModulesDetails Provides the information used to search for a set of modules from a list software sources.
type SearchSoftwareSourceModulesDetails struct {

	// List of sofware source OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	SoftwareSourceIds []string `mandatory:"true" json:"softwareSourceIds"`

	// The sort order.
	SortOrder SearchSoftwareSourceModulesDetailsSortOrderEnum `mandatory:"false" json:"sortOrder,omitempty"`

	// The name of a module.
	Name *string `mandatory:"false" json:"name"`

	// A filter to return modules with a name that contains the given string.
	NameContains *string `mandatory:"false" json:"nameContains"`

	// The field to sort by.
	SortBy SearchSoftwareSourceModulesDetailsSortByEnum `mandatory:"false" json:"sortBy,omitempty"`
}

func (m SearchSoftwareSourceModulesDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SearchSoftwareSourceModulesDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSearchSoftwareSourceModulesDetailsSortOrderEnum(string(m.SortOrder)); !ok && m.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", m.SortOrder, strings.Join(GetSearchSoftwareSourceModulesDetailsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSearchSoftwareSourceModulesDetailsSortByEnum(string(m.SortBy)); !ok && m.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", m.SortBy, strings.Join(GetSearchSoftwareSourceModulesDetailsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SearchSoftwareSourceModulesDetailsSortOrderEnum Enum with underlying type: string
type SearchSoftwareSourceModulesDetailsSortOrderEnum string

// Set of constants representing the allowable values for SearchSoftwareSourceModulesDetailsSortOrderEnum
const (
	SearchSoftwareSourceModulesDetailsSortOrderAsc  SearchSoftwareSourceModulesDetailsSortOrderEnum = "ASC"
	SearchSoftwareSourceModulesDetailsSortOrderDesc SearchSoftwareSourceModulesDetailsSortOrderEnum = "DESC"
)

var mappingSearchSoftwareSourceModulesDetailsSortOrderEnum = map[string]SearchSoftwareSourceModulesDetailsSortOrderEnum{
	"ASC":  SearchSoftwareSourceModulesDetailsSortOrderAsc,
	"DESC": SearchSoftwareSourceModulesDetailsSortOrderDesc,
}

var mappingSearchSoftwareSourceModulesDetailsSortOrderEnumLowerCase = map[string]SearchSoftwareSourceModulesDetailsSortOrderEnum{
	"asc":  SearchSoftwareSourceModulesDetailsSortOrderAsc,
	"desc": SearchSoftwareSourceModulesDetailsSortOrderDesc,
}

// GetSearchSoftwareSourceModulesDetailsSortOrderEnumValues Enumerates the set of values for SearchSoftwareSourceModulesDetailsSortOrderEnum
func GetSearchSoftwareSourceModulesDetailsSortOrderEnumValues() []SearchSoftwareSourceModulesDetailsSortOrderEnum {
	values := make([]SearchSoftwareSourceModulesDetailsSortOrderEnum, 0)
	for _, v := range mappingSearchSoftwareSourceModulesDetailsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetSearchSoftwareSourceModulesDetailsSortOrderEnumStringValues Enumerates the set of values in String for SearchSoftwareSourceModulesDetailsSortOrderEnum
func GetSearchSoftwareSourceModulesDetailsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingSearchSoftwareSourceModulesDetailsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSearchSoftwareSourceModulesDetailsSortOrderEnum(val string) (SearchSoftwareSourceModulesDetailsSortOrderEnum, bool) {
	enum, ok := mappingSearchSoftwareSourceModulesDetailsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SearchSoftwareSourceModulesDetailsSortByEnum Enum with underlying type: string
type SearchSoftwareSourceModulesDetailsSortByEnum string

// Set of constants representing the allowable values for SearchSoftwareSourceModulesDetailsSortByEnum
const (
	SearchSoftwareSourceModulesDetailsSortByName SearchSoftwareSourceModulesDetailsSortByEnum = "NAME"
)

var mappingSearchSoftwareSourceModulesDetailsSortByEnum = map[string]SearchSoftwareSourceModulesDetailsSortByEnum{
	"NAME": SearchSoftwareSourceModulesDetailsSortByName,
}

var mappingSearchSoftwareSourceModulesDetailsSortByEnumLowerCase = map[string]SearchSoftwareSourceModulesDetailsSortByEnum{
	"name": SearchSoftwareSourceModulesDetailsSortByName,
}

// GetSearchSoftwareSourceModulesDetailsSortByEnumValues Enumerates the set of values for SearchSoftwareSourceModulesDetailsSortByEnum
func GetSearchSoftwareSourceModulesDetailsSortByEnumValues() []SearchSoftwareSourceModulesDetailsSortByEnum {
	values := make([]SearchSoftwareSourceModulesDetailsSortByEnum, 0)
	for _, v := range mappingSearchSoftwareSourceModulesDetailsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetSearchSoftwareSourceModulesDetailsSortByEnumStringValues Enumerates the set of values in String for SearchSoftwareSourceModulesDetailsSortByEnum
func GetSearchSoftwareSourceModulesDetailsSortByEnumStringValues() []string {
	return []string{
		"NAME",
	}
}

// GetMappingSearchSoftwareSourceModulesDetailsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSearchSoftwareSourceModulesDetailsSortByEnum(val string) (SearchSoftwareSourceModulesDetailsSortByEnum, bool) {
	enum, ok := mappingSearchSoftwareSourceModulesDetailsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
