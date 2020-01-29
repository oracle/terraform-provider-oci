// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
//

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
)

// SearchResultCollection The list of search result items matching the criteria returned from the search operation. Search errors and
// messages, if any , will be part of the standard error response.
type SearchResultCollection struct {

	// Total number of items returned.
	Count *int `mandatory:"false" json:"count"`

	// Search result set.
	Items []SearchResult `mandatory:"false" json:"items"`
}

func (m SearchResultCollection) String() string {
	return common.PointerString(m)
}
