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

// DataAssetCollection Results of a data assets listing. A data asset is often synonymous with a 'System', such as a database, or may be a file container or a message stream.
type DataAssetCollection struct {

	// Collection of data asset summaries.
	Items []DataAssetSummary `mandatory:"true" json:"items"`
}

func (m DataAssetCollection) String() string {
	return common.PointerString(m)
}
