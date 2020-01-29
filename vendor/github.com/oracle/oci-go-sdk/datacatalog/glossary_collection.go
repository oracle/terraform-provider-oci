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

// GlossaryCollection Results of a glossaries listing.  Glossary is an organizing concept for business terms to provide a unified semantic model across disparate data assets.
type GlossaryCollection struct {

	// Collection of glossaries.
	Items []GlossarySummary `mandatory:"true" json:"items"`
}

func (m GlossaryCollection) String() string {
	return common.PointerString(m)
}
