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

// EntityCollection Results of a data entities listing. Data entities are representation of a dataset with a set of attributes.
type EntityCollection struct {

	// Collection of data entities.
	Items []EntitySummary `mandatory:"true" json:"items"`
}

func (m EntityCollection) String() string {
	return common.PointerString(m)
}
