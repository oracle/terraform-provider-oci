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

// ConnectionAliasSummary Summary representation of database aliases parsed from the file metadata.
type ConnectionAliasSummary struct {

	// A user-friendly display name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	AliasName *string `mandatory:"true" json:"aliasName"`

	// The description about the database alias parsed from the file metadata.
	AliasDetails *string `mandatory:"false" json:"aliasDetails"`
}

func (m ConnectionAliasSummary) String() string {
	return common.PointerString(m)
}
