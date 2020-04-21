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

// ConnectionSummary Summary representation of a connection to a data asset.
type ConnectionSummary struct {

	// Unique connection key that is immutable.
	Key *string `mandatory:"true" json:"key"`

	// A description of the connection.
	Description *string `mandatory:"false" json:"description"`

	// A user-friendly display name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The unique key of the parent data asset.
	DataAssetKey *string `mandatory:"false" json:"dataAssetKey"`

	// The key of the object type. Type key's can be found via the '/types' endpoint.
	TypeKey *string `mandatory:"false" json:"typeKey"`

	// URI to the connection instance in the API.
	Uri *string `mandatory:"false" json:"uri"`

	// Unique external key for this object as defined in the source systems.
	ExternalKey *string `mandatory:"false" json:"externalKey"`

	// The current state of the connection.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Indicates whether this connection is the default connection.
	IsDefault *bool `mandatory:"false" json:"isDefault"`

	// The date and time the connection was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2019-03-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m ConnectionSummary) String() string {
	return common.PointerString(m)
}
