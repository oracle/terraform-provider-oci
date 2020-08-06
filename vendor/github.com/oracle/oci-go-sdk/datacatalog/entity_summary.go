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

// EntitySummary Summary of an data entity. A representation of data with a set of attributes, normally representing a single
// business entity. Synonymous with 'table' or 'view' in a database, or a single logical file structure
// that one or many files may match.
type EntitySummary struct {

	// Unique data entity key that is immutable.
	Key *string `mandatory:"true" json:"key"`

	// A user-friendly display name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Detailed description of a data entity.
	Description *string `mandatory:"false" json:"description"`

	// Unique key of the parent data asset.
	DataAssetKey *string `mandatory:"false" json:"dataAssetKey"`

	// Key of the associated folder.
	FolderKey *string `mandatory:"false" json:"folderKey"`

	// Name of the associated folder. This name is harvested from the source data asset when the parent folder for the entiy is harvested.
	FolderName *string `mandatory:"false" json:"folderName"`

	// Unique external key of this object in the source system.
	ExternalKey *string `mandatory:"false" json:"externalKey"`

	// Full path of the data entity.
	Path *string `mandatory:"false" json:"path"`

	// The date and time the data entity was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2019-03-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The last time that any change was made to the data entity. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// OCID of the user who updated this object in the data catalog.
	UpdatedById *string `mandatory:"false" json:"updatedById"`

	// URI to the data entity instance in the API.
	Uri *string `mandatory:"false" json:"uri"`

	// State of the data entity.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m EntitySummary) String() string {
	return common.PointerString(m)
}
