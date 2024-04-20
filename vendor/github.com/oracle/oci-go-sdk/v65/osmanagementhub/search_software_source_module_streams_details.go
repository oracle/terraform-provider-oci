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

// SearchSoftwareSourceModuleStreamsDetails Provides the information used to search for a set of module streams from a list software sources.
type SearchSoftwareSourceModuleStreamsDetails struct {

	// List of software source OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	SoftwareSourceIds []string `mandatory:"true" json:"softwareSourceIds"`

	// The sort order.
	SortOrder SearchSoftwareSourceModuleStreamsDetailsSortOrderEnum `mandatory:"false" json:"sortOrder,omitempty"`

	// The name of a module.
	ModuleName *string `mandatory:"false" json:"moduleName"`

	// The field to sort by.
	SortBy SearchSoftwareSourceModuleStreamsDetailsSortByEnum `mandatory:"false" json:"sortBy,omitempty"`
}

func (m SearchSoftwareSourceModuleStreamsDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SearchSoftwareSourceModuleStreamsDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSearchSoftwareSourceModuleStreamsDetailsSortOrderEnum(string(m.SortOrder)); !ok && m.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", m.SortOrder, strings.Join(GetSearchSoftwareSourceModuleStreamsDetailsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSearchSoftwareSourceModuleStreamsDetailsSortByEnum(string(m.SortBy)); !ok && m.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", m.SortBy, strings.Join(GetSearchSoftwareSourceModuleStreamsDetailsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SearchSoftwareSourceModuleStreamsDetailsSortOrderEnum Enum with underlying type: string
type SearchSoftwareSourceModuleStreamsDetailsSortOrderEnum string

// Set of constants representing the allowable values for SearchSoftwareSourceModuleStreamsDetailsSortOrderEnum
const (
	SearchSoftwareSourceModuleStreamsDetailsSortOrderAsc  SearchSoftwareSourceModuleStreamsDetailsSortOrderEnum = "ASC"
	SearchSoftwareSourceModuleStreamsDetailsSortOrderDesc SearchSoftwareSourceModuleStreamsDetailsSortOrderEnum = "DESC"
)

var mappingSearchSoftwareSourceModuleStreamsDetailsSortOrderEnum = map[string]SearchSoftwareSourceModuleStreamsDetailsSortOrderEnum{
	"ASC":  SearchSoftwareSourceModuleStreamsDetailsSortOrderAsc,
	"DESC": SearchSoftwareSourceModuleStreamsDetailsSortOrderDesc,
}

var mappingSearchSoftwareSourceModuleStreamsDetailsSortOrderEnumLowerCase = map[string]SearchSoftwareSourceModuleStreamsDetailsSortOrderEnum{
	"asc":  SearchSoftwareSourceModuleStreamsDetailsSortOrderAsc,
	"desc": SearchSoftwareSourceModuleStreamsDetailsSortOrderDesc,
}

// GetSearchSoftwareSourceModuleStreamsDetailsSortOrderEnumValues Enumerates the set of values for SearchSoftwareSourceModuleStreamsDetailsSortOrderEnum
func GetSearchSoftwareSourceModuleStreamsDetailsSortOrderEnumValues() []SearchSoftwareSourceModuleStreamsDetailsSortOrderEnum {
	values := make([]SearchSoftwareSourceModuleStreamsDetailsSortOrderEnum, 0)
	for _, v := range mappingSearchSoftwareSourceModuleStreamsDetailsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetSearchSoftwareSourceModuleStreamsDetailsSortOrderEnumStringValues Enumerates the set of values in String for SearchSoftwareSourceModuleStreamsDetailsSortOrderEnum
func GetSearchSoftwareSourceModuleStreamsDetailsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingSearchSoftwareSourceModuleStreamsDetailsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSearchSoftwareSourceModuleStreamsDetailsSortOrderEnum(val string) (SearchSoftwareSourceModuleStreamsDetailsSortOrderEnum, bool) {
	enum, ok := mappingSearchSoftwareSourceModuleStreamsDetailsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SearchSoftwareSourceModuleStreamsDetailsSortByEnum Enum with underlying type: string
type SearchSoftwareSourceModuleStreamsDetailsSortByEnum string

// Set of constants representing the allowable values for SearchSoftwareSourceModuleStreamsDetailsSortByEnum
const (
	SearchSoftwareSourceModuleStreamsDetailsSortByModulename SearchSoftwareSourceModuleStreamsDetailsSortByEnum = "MODULENAME"
)

var mappingSearchSoftwareSourceModuleStreamsDetailsSortByEnum = map[string]SearchSoftwareSourceModuleStreamsDetailsSortByEnum{
	"MODULENAME": SearchSoftwareSourceModuleStreamsDetailsSortByModulename,
}

var mappingSearchSoftwareSourceModuleStreamsDetailsSortByEnumLowerCase = map[string]SearchSoftwareSourceModuleStreamsDetailsSortByEnum{
	"modulename": SearchSoftwareSourceModuleStreamsDetailsSortByModulename,
}

// GetSearchSoftwareSourceModuleStreamsDetailsSortByEnumValues Enumerates the set of values for SearchSoftwareSourceModuleStreamsDetailsSortByEnum
func GetSearchSoftwareSourceModuleStreamsDetailsSortByEnumValues() []SearchSoftwareSourceModuleStreamsDetailsSortByEnum {
	values := make([]SearchSoftwareSourceModuleStreamsDetailsSortByEnum, 0)
	for _, v := range mappingSearchSoftwareSourceModuleStreamsDetailsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetSearchSoftwareSourceModuleStreamsDetailsSortByEnumStringValues Enumerates the set of values in String for SearchSoftwareSourceModuleStreamsDetailsSortByEnum
func GetSearchSoftwareSourceModuleStreamsDetailsSortByEnumStringValues() []string {
	return []string{
		"MODULENAME",
	}
}

// GetMappingSearchSoftwareSourceModuleStreamsDetailsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSearchSoftwareSourceModuleStreamsDetailsSortByEnum(val string) (SearchSoftwareSourceModuleStreamsDetailsSortByEnum, bool) {
	enum, ok := mappingSearchSoftwareSourceModuleStreamsDetailsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
