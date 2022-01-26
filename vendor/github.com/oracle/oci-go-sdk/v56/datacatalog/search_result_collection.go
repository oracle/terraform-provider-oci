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

// SearchResultCollection The list of search result items matching the criteria returned from the search operation. Search errors and
// messages, if any , will be part of the standard error response.
type SearchResultCollection struct {

	// Total number of items returned.
	Count *int `mandatory:"false" json:"count"`

	// Search result set.
	Items []SearchResult `mandatory:"false" json:"items"`

	// String that data objects are to be searched with.
	Query *string `mandatory:"false" json:"query"`

	// Aggregations/facets on properties of data objects.
	FacetedSearchAggregation []FacetedSearchAggregation `mandatory:"false" json:"facetedSearchAggregation"`

	// A list of fields or properties used in the sorting of a search result.
	SortableFields []string `mandatory:"false" json:"sortableFields"`
}

func (m SearchResultCollection) String() string {
	return common.PointerString(m)
}
