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

// TermRelationshipCollection Results of a terms relationship listing. Term relationships are associations between two terms in business glossary.
type TermRelationshipCollection struct {

	// Collection of term relationships.
	Items []TermRelationshipSummary `mandatory:"true" json:"items"`
}

func (m TermRelationshipCollection) String() string {
	return common.PointerString(m)
}
