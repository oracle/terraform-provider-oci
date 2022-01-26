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

// SearchCriteria Search Query object that allows complex search predicates that cannot be expressed through simple query params.
type SearchCriteria struct {

	// Search query dsl that defines the query components including fields and predicates.
	Query *string `mandatory:"false" json:"query"`

	// Query string that a dataObject is to be searched with. Used in the faceted query request
	FacetedQuery *string `mandatory:"false" json:"facetedQuery"`

	// List of properties of dataObjects that needs to aggregated on for facets.
	Dimensions []string `mandatory:"false" json:"dimensions"`

	// Array of objects having details about sort field and order.
	Sort []FacetedSearchSortRequest `mandatory:"false" json:"sort"`

	Filters *FacetedSearchFilterRequest `mandatory:"false" json:"filters"`
}

func (m SearchCriteria) String() string {
	return common.PointerString(m)
}
