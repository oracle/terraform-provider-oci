// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
// For more information, see Data Catalog (https://docs.oracle.com/iaas/data-catalog/home.htm).
//

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/v56/common"
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

// FacetedSearchSortRequestSortOrderEnum Enum with underlying type: string
type FacetedSearchSortRequestSortOrderEnum string

// Set of constants representing the allowable values for FacetedSearchSortRequestSortOrderEnum
const (
	FacetedSearchSortRequestSortOrderAsc  FacetedSearchSortRequestSortOrderEnum = "ASC"
	FacetedSearchSortRequestSortOrderDesc FacetedSearchSortRequestSortOrderEnum = "DESC"
)

var mappingFacetedSearchSortRequestSortOrder = map[string]FacetedSearchSortRequestSortOrderEnum{
	"ASC":  FacetedSearchSortRequestSortOrderAsc,
	"DESC": FacetedSearchSortRequestSortOrderDesc,
}

// GetFacetedSearchSortRequestSortOrderEnumValues Enumerates the set of values for FacetedSearchSortRequestSortOrderEnum
func GetFacetedSearchSortRequestSortOrderEnumValues() []FacetedSearchSortRequestSortOrderEnum {
	values := make([]FacetedSearchSortRequestSortOrderEnum, 0)
	for _, v := range mappingFacetedSearchSortRequestSortOrder {
		values = append(values, v)
	}
	return values
}
