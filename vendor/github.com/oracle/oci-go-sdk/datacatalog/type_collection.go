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

// TypeCollection Results of a types listing. Types define the basic type of catalog objects and are immutable.
type TypeCollection struct {

	// Collection of types.
	Items []TypeSummary `mandatory:"true" json:"items"`
}

func (m TypeCollection) String() string {
	return common.PointerString(m)
}
