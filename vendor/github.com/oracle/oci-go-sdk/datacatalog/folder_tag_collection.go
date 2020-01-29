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

// FolderTagCollection Results of a folders tag listing. Folder tags allow association of folder objects to business terms.
type FolderTagCollection struct {

	// Collection of folder tags.
	Items []FolderTagSummary `mandatory:"true" json:"items"`
}

func (m FolderTagCollection) String() string {
	return common.PointerString(m)
}
