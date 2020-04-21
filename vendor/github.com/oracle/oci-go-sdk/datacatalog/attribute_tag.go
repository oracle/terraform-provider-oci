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

// AttributeTag Represents an association of an entity attribute to a term.
type AttributeTag struct {

	// Unique tag key that is immutable.
	Key *string `mandatory:"true" json:"key"`

	// Name of the tag which matches the term name.
	Name *string `mandatory:"false" json:"name"`

	// Unique key of the related term.
	TermKey *string `mandatory:"false" json:"termKey"`

	// Path of the related term.
	TermPath *string `mandatory:"false" json:"termPath"`

	// Description of the related term.
	TermDescription *string `mandatory:"false" json:"termDescription"`

	// The current state of the tag.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The date and time the tag was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2019-03-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// OCID of the user who created the tag.
	CreatedById *string `mandatory:"false" json:"createdById"`

	// URI to the tag instance in the API.
	Uri *string `mandatory:"false" json:"uri"`

	// The unique key of the parent attribute.
	AttributeKey *string `mandatory:"false" json:"attributeKey"`
}

func (m AttributeTag) String() string {
	return common.PointerString(m)
}
