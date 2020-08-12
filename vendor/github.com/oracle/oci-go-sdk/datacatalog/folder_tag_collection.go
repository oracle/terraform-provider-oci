// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
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

	// Total number of items returned.
	Count *int `mandatory:"false" json:"count"`
}

func (m FolderTagCollection) String() string {
	return common.PointerString(m)
}
