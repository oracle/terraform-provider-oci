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

// FolderCollection Results of a folders listing. Folders are external organization concept that groups data entities.
type FolderCollection struct {

	// Collection of folders.
	Items []FolderSummary `mandatory:"true" json:"items"`
}

func (m FolderCollection) String() string {
	return common.PointerString(m)
}
