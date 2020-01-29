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

// AttributeSummary Summary of an entity attribute.
type AttributeSummary struct {

	// Unique attribute key that is immutable.
	Key *string `mandatory:"true" json:"key"`

	// A user-friendly display name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Detailed description of the attribute.
	Description *string `mandatory:"false" json:"description"`

	// The unique key of the parent entity.
	EntityKey *string `mandatory:"false" json:"entityKey"`

	// Unique external key of this attribute in the external source system.
	ExternalKey *string `mandatory:"false" json:"externalKey"`

	// Max allowed length of the attribute value.
	Length *int64 `mandatory:"false" json:"length"`

	// Property that identifies if this attribute can be assigned null values.
	IsNullable *bool `mandatory:"false" json:"isNullable"`

	// URI to the attribute instance in the API.
	Uri *string `mandatory:"false" json:"uri"`

	// State of the attribute.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The date and time the attribute was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2019-03-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Data type of the attribute as defined in the external source system.
	ExternalDataType *string `mandatory:"false" json:"externalDataType"`
}

func (m AttributeSummary) String() string {
	return common.PointerString(m)
}
