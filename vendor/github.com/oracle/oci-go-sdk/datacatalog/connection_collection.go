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

// ConnectionCollection Results of a connections listing. Each member of the result is a summary representation of a connection to a data asset.
type ConnectionCollection struct {

	// Collection of connection summaries.
	Items []ConnectionSummary `mandatory:"true" json:"items"`

	// Total number of items returned.
	Count *int `mandatory:"false" json:"count"`
}

func (m ConnectionCollection) String() string {
	return common.PointerString(m)
}
