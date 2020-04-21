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

// TypeSummary Summary data catalog type information. All types are statically defined in the system and are immutable.
// It isn't possible to create new types or update existing types via the API.
type TypeSummary struct {

	// Unique type key that is immutable.
	Key *string `mandatory:"true" json:"key"`

	// The immutable name of the type.
	Name *string `mandatory:"false" json:"name"`

	// Detailed description of the type.
	Description *string `mandatory:"false" json:"description"`

	// The data catalog's OCID.
	CatalogId *string `mandatory:"false" json:"catalogId"`

	// Indicates the category this type belongs to. For instance, data assets, connections.
	TypeCategory *string `mandatory:"false" json:"typeCategory"`

	// URI to the type instance in the API.
	Uri *string `mandatory:"false" json:"uri"`

	// State of the folder.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m TypeSummary) String() string {
	return common.PointerString(m)
}
