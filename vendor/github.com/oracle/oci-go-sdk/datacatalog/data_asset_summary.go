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

// DataAssetSummary Summary of a data asset. A physical store, or stream, of data known to the data catalog and containing one or
// many data entities, possibly in an organized structure of folders. A data asset is often synonymous with
// a 'System', such as a database, or may be a file container or a message stream.
type DataAssetSummary struct {

	// Unique data asset key that is immutable.
	Key *string `mandatory:"true" json:"key"`

	// A user-friendly display name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Detailed description of the data asset.
	Description *string `mandatory:"false" json:"description"`

	// The data catalog's OCID.
	CatalogId *string `mandatory:"false" json:"catalogId"`

	// External URI that can be used to reference the object. Format will differ based on the type of object.
	ExternalKey *string `mandatory:"false" json:"externalKey"`

	// URI to the data asset instance in the API.
	Uri *string `mandatory:"false" json:"uri"`

	// The date and time the data asset was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2019-03-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The key of the object type. Type keys's can be found via the '/types' endpoint.
	TypeKey *string `mandatory:"false" json:"typeKey"`

	// State of the data asset.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m DataAssetSummary) String() string {
	return common.PointerString(m)
}
