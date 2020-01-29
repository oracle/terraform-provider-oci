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

// AttributeCollection Results of an attributes listing. Attributes describe an item of data with name and datatype.
type AttributeCollection struct {

	// Collection of attributes.
	Items []AttributeSummary `mandatory:"true" json:"items"`
}

func (m AttributeCollection) String() string {
	return common.PointerString(m)
}
