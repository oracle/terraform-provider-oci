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

// SearchTermSummary Summary of a term associated with an object. This is a brief summary returned as part of the search result.
type SearchTermSummary struct {

	// A user-friendly display name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Unique term key that is immutable.
	Key *string `mandatory:"false" json:"key"`

	// Unique id of the parent glossary.
	GlossaryKey *string `mandatory:"false" json:"glossaryKey"`

	// Name of the parent glossary.
	GlossaryName *string `mandatory:"false" json:"glossaryName"`

	// This terms parent term key. Will be null if the term has no parent term.
	ParentTermKey *string `mandatory:"false" json:"parentTermKey"`

	// Name of the parent term key. Will be null if the term has no parent term.
	ParentTermName *string `mandatory:"false" json:"parentTermName"`
}

func (m SearchTermSummary) String() string {
	return common.PointerString(m)
}
