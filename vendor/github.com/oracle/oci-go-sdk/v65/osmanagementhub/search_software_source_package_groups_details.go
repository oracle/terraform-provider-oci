// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for the operating system environments in your private data centers through a single management console. For more information, see Overview of OS Management Hub (https://docs.cloud.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SearchSoftwareSourcePackageGroupsDetails Contains a list of software sources to get the list of associated package groups.
type SearchSoftwareSourcePackageGroupsDetails struct {

	// List of software source OCIDs.
	SoftwareSourceIds []string `mandatory:"true" json:"softwareSourceIds"`

	// The sort order.
	SortOrder SearchSoftwareSourcePackageGroupsDetailsSortOrderEnum `mandatory:"false" json:"sortOrder,omitempty"`

	// The field to sort by.
	SortBy SearchSoftwareSourcePackageGroupsDetailsSortByEnum `mandatory:"false" json:"sortBy,omitempty"`

	// filters results, allowing only those with a Name which contains the string.
	NameContains *string `mandatory:"false" json:"nameContains"`

	// Indicates if this is a group, category or environment.
	GroupType PackageGroupGroupTypeEnum `mandatory:"false" json:"groupType,omitempty"`
}

func (m SearchSoftwareSourcePackageGroupsDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SearchSoftwareSourcePackageGroupsDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSearchSoftwareSourcePackageGroupsDetailsSortOrderEnum(string(m.SortOrder)); !ok && m.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", m.SortOrder, strings.Join(GetSearchSoftwareSourcePackageGroupsDetailsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSearchSoftwareSourcePackageGroupsDetailsSortByEnum(string(m.SortBy)); !ok && m.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", m.SortBy, strings.Join(GetSearchSoftwareSourcePackageGroupsDetailsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPackageGroupGroupTypeEnum(string(m.GroupType)); !ok && m.GroupType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GroupType: %s. Supported values are: %s.", m.GroupType, strings.Join(GetPackageGroupGroupTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SearchSoftwareSourcePackageGroupsDetailsSortOrderEnum Enum with underlying type: string
type SearchSoftwareSourcePackageGroupsDetailsSortOrderEnum string

// Set of constants representing the allowable values for SearchSoftwareSourcePackageGroupsDetailsSortOrderEnum
const (
	SearchSoftwareSourcePackageGroupsDetailsSortOrderAsc  SearchSoftwareSourcePackageGroupsDetailsSortOrderEnum = "ASC"
	SearchSoftwareSourcePackageGroupsDetailsSortOrderDesc SearchSoftwareSourcePackageGroupsDetailsSortOrderEnum = "DESC"
)

var mappingSearchSoftwareSourcePackageGroupsDetailsSortOrderEnum = map[string]SearchSoftwareSourcePackageGroupsDetailsSortOrderEnum{
	"ASC":  SearchSoftwareSourcePackageGroupsDetailsSortOrderAsc,
	"DESC": SearchSoftwareSourcePackageGroupsDetailsSortOrderDesc,
}

var mappingSearchSoftwareSourcePackageGroupsDetailsSortOrderEnumLowerCase = map[string]SearchSoftwareSourcePackageGroupsDetailsSortOrderEnum{
	"asc":  SearchSoftwareSourcePackageGroupsDetailsSortOrderAsc,
	"desc": SearchSoftwareSourcePackageGroupsDetailsSortOrderDesc,
}

// GetSearchSoftwareSourcePackageGroupsDetailsSortOrderEnumValues Enumerates the set of values for SearchSoftwareSourcePackageGroupsDetailsSortOrderEnum
func GetSearchSoftwareSourcePackageGroupsDetailsSortOrderEnumValues() []SearchSoftwareSourcePackageGroupsDetailsSortOrderEnum {
	values := make([]SearchSoftwareSourcePackageGroupsDetailsSortOrderEnum, 0)
	for _, v := range mappingSearchSoftwareSourcePackageGroupsDetailsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetSearchSoftwareSourcePackageGroupsDetailsSortOrderEnumStringValues Enumerates the set of values in String for SearchSoftwareSourcePackageGroupsDetailsSortOrderEnum
func GetSearchSoftwareSourcePackageGroupsDetailsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingSearchSoftwareSourcePackageGroupsDetailsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSearchSoftwareSourcePackageGroupsDetailsSortOrderEnum(val string) (SearchSoftwareSourcePackageGroupsDetailsSortOrderEnum, bool) {
	enum, ok := mappingSearchSoftwareSourcePackageGroupsDetailsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SearchSoftwareSourcePackageGroupsDetailsSortByEnum Enum with underlying type: string
type SearchSoftwareSourcePackageGroupsDetailsSortByEnum string

// Set of constants representing the allowable values for SearchSoftwareSourcePackageGroupsDetailsSortByEnum
const (
	SearchSoftwareSourcePackageGroupsDetailsSortByName SearchSoftwareSourcePackageGroupsDetailsSortByEnum = "NAME"
)

var mappingSearchSoftwareSourcePackageGroupsDetailsSortByEnum = map[string]SearchSoftwareSourcePackageGroupsDetailsSortByEnum{
	"NAME": SearchSoftwareSourcePackageGroupsDetailsSortByName,
}

var mappingSearchSoftwareSourcePackageGroupsDetailsSortByEnumLowerCase = map[string]SearchSoftwareSourcePackageGroupsDetailsSortByEnum{
	"name": SearchSoftwareSourcePackageGroupsDetailsSortByName,
}

// GetSearchSoftwareSourcePackageGroupsDetailsSortByEnumValues Enumerates the set of values for SearchSoftwareSourcePackageGroupsDetailsSortByEnum
func GetSearchSoftwareSourcePackageGroupsDetailsSortByEnumValues() []SearchSoftwareSourcePackageGroupsDetailsSortByEnum {
	values := make([]SearchSoftwareSourcePackageGroupsDetailsSortByEnum, 0)
	for _, v := range mappingSearchSoftwareSourcePackageGroupsDetailsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetSearchSoftwareSourcePackageGroupsDetailsSortByEnumStringValues Enumerates the set of values in String for SearchSoftwareSourcePackageGroupsDetailsSortByEnum
func GetSearchSoftwareSourcePackageGroupsDetailsSortByEnumStringValues() []string {
	return []string{
		"NAME",
	}
}

// GetMappingSearchSoftwareSourcePackageGroupsDetailsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSearchSoftwareSourcePackageGroupsDetailsSortByEnum(val string) (SearchSoftwareSourcePackageGroupsDetailsSortByEnum, bool) {
	enum, ok := mappingSearchSoftwareSourcePackageGroupsDetailsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
