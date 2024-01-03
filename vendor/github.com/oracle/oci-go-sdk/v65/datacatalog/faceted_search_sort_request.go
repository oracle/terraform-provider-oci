// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
// For more information, see Data Catalog (https://docs.oracle.com/iaas/data-catalog/home.htm).
//

package datacatalog

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// FacetedSearchSortRequest Object with sort criteria details
type FacetedSearchSortRequest struct {

	// Filed name that needs to be sorted by.
	SortBy *string `mandatory:"false" json:"sortBy"`

	// Sort order for search results.
	SortOrder FacetedSearchSortRequestSortOrderEnum `mandatory:"false" json:"sortOrder,omitempty"`
}

func (m FacetedSearchSortRequest) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FacetedSearchSortRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingFacetedSearchSortRequestSortOrderEnum(string(m.SortOrder)); !ok && m.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", m.SortOrder, strings.Join(GetFacetedSearchSortRequestSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// FacetedSearchSortRequestSortOrderEnum Enum with underlying type: string
type FacetedSearchSortRequestSortOrderEnum string

// Set of constants representing the allowable values for FacetedSearchSortRequestSortOrderEnum
const (
	FacetedSearchSortRequestSortOrderAsc  FacetedSearchSortRequestSortOrderEnum = "ASC"
	FacetedSearchSortRequestSortOrderDesc FacetedSearchSortRequestSortOrderEnum = "DESC"
)

var mappingFacetedSearchSortRequestSortOrderEnum = map[string]FacetedSearchSortRequestSortOrderEnum{
	"ASC":  FacetedSearchSortRequestSortOrderAsc,
	"DESC": FacetedSearchSortRequestSortOrderDesc,
}

var mappingFacetedSearchSortRequestSortOrderEnumLowerCase = map[string]FacetedSearchSortRequestSortOrderEnum{
	"asc":  FacetedSearchSortRequestSortOrderAsc,
	"desc": FacetedSearchSortRequestSortOrderDesc,
}

// GetFacetedSearchSortRequestSortOrderEnumValues Enumerates the set of values for FacetedSearchSortRequestSortOrderEnum
func GetFacetedSearchSortRequestSortOrderEnumValues() []FacetedSearchSortRequestSortOrderEnum {
	values := make([]FacetedSearchSortRequestSortOrderEnum, 0)
	for _, v := range mappingFacetedSearchSortRequestSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetFacetedSearchSortRequestSortOrderEnumStringValues Enumerates the set of values in String for FacetedSearchSortRequestSortOrderEnum
func GetFacetedSearchSortRequestSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingFacetedSearchSortRequestSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFacetedSearchSortRequestSortOrderEnum(val string) (FacetedSearchSortRequestSortOrderEnum, bool) {
	enum, ok := mappingFacetedSearchSortRequestSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
