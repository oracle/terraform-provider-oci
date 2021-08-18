// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
//

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/v46/common"
)

// ObjectRelationship Details regarding a specific object and its relationship to the referencing object.
type ObjectRelationship struct {

	// Type of relationship with the referencing object.
	RelationshipType *string `mandatory:"false" json:"relationshipType"`

	// Unique id of the object.
	Key *string `mandatory:"false" json:"key"`

	// Name of the object.
	Name *string `mandatory:"false" json:"name"`

	// Type name of the object. Type names can be found via the '/types' endpoint.
	TypeName *string `mandatory:"false" json:"typeName"`

	// Type key of the object. Type keys can be found via the '/types' endpoint.
	TypeKey *string `mandatory:"false" json:"typeKey"`

	// The date and time the relationship was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2019-03-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The last time a change was made to this reference. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m ObjectRelationship) String() string {
	return common.PointerString(m)
}
