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

// Connection Detailed representation of a connection to a data asset, minus any sensitive properties.
type Connection struct {

	// Unique connection key that is immutable.
	Key *string `mandatory:"true" json:"key"`

	// A description of the connection.
	Description *string `mandatory:"false" json:"description"`

	// A user-friendly display name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The date and time the connection was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2019-03-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The last time that any change was made to the connection. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// OCID of the user who created the connection.
	CreatedById *string `mandatory:"false" json:"createdById"`

	// OCID of the user who modified the connection.
	UpdatedById *string `mandatory:"false" json:"updatedById"`

	// A map of maps that contains the properties which are specific to the connection type. Each connection type
	// definition defines it's set of required and optional properties. The map keys are category names and the
	// values are maps of property name to property value. Every property is contained inside of a category. Most
	// connections have required properties within the "default" category.
	// Example: `{"properties": { "default": { "username": "user1"}}}`
	Properties map[string]map[string]string `mandatory:"false" json:"properties"`

	// Unique external key of this object from the source system.
	ExternalKey *string `mandatory:"false" json:"externalKey"`

	// Time that the connections status was last updated. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeStatusUpdated *common.SDKTime `mandatory:"false" json:"timeStatusUpdated"`

	// The current state of the connection.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Indicates whether this connection is the default connection.
	IsDefault *bool `mandatory:"false" json:"isDefault"`

	// Unique key of the parent data asset.
	DataAssetKey *string `mandatory:"false" json:"dataAssetKey"`

	// The key of the object type. Type key's can be found via the '/types' endpoint.
	TypeKey *string `mandatory:"false" json:"typeKey"`

	// URI to the connection instance in the API.
	Uri *string `mandatory:"false" json:"uri"`
}

func (m Connection) String() string {
	return common.PointerString(m)
}
