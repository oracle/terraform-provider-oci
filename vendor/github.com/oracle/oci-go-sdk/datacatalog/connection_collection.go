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

// ConnectionCollection Results of a connections listing. Each member of the result is a summary representation of a connection to a data asset.
type ConnectionCollection struct {

	// Collection of connection summaries.
	Items []ConnectionSummary `mandatory:"true" json:"items"`
}

func (m ConnectionCollection) String() string {
	return common.PointerString(m)
}
