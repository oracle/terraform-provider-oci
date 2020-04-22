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

// Folder A generic term used in the data catalog for an external organization concept used for a collection of data entities
// or processes within a data asset. This term is an internal term which models multiple external types of folder,
// such as file directories, database schemas, and so on. Some data assets, such as Object Store containers, may contain
// many levels of folders.
type Folder struct {

	// Unique folder key that is immutable.
	Key *string `mandatory:"true" json:"key"`

	// A user-friendly display name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Detailed description of a folder.
	Description *string `mandatory:"false" json:"description"`

	// The unique key of the containing folder or null if there is no parent folder.
	ParentFolderKey *string `mandatory:"false" json:"parentFolderKey"`

	// Full path of the folder.
	Path *string `mandatory:"false" json:"path"`

	// The key of the associated data asset.
	DataAssetKey *string `mandatory:"false" json:"dataAssetKey"`

	// A map of maps that contains the properties which are specific to the folder type. Each folder type
	// definition defines it's set of required and optional properties. The map keys are category names and the
	// values are maps of property name to property value. Every property is contained inside of a category. Most
	// folders have required properties within the "default" category.
	// Example: `{"properties": { "default": { "key1": "value1"}}}`
	Properties map[string]map[string]string `mandatory:"false" json:"properties"`

	// Unique external key of this object in the source system.
	ExternalKey *string `mandatory:"false" json:"externalKey"`

	// The date and time the folder was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2019-03-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The last time that any change was made to the folder. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// OCID of the user who created the folder.
	CreatedById *string `mandatory:"false" json:"createdById"`

	// OCID of the user who modified the folder.
	UpdatedById *string `mandatory:"false" json:"updatedById"`

	// Last modified timestamp of this object in the external system.
	TimeExternal *common.SDKTime `mandatory:"false" json:"timeExternal"`

	// The current state of the folder.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Status of the object as updated by the harvest process.
	HarvestStatus HarvestStatusEnum `mandatory:"false" json:"harvestStatus,omitempty"`

	// The key of the last harvest process to update the metadata of this object.
	LastJobKey *string `mandatory:"false" json:"lastJobKey"`

	// URI to the folder instance in the API.
	Uri *string `mandatory:"false" json:"uri"`
}

func (m Folder) String() string {
	return common.PointerString(m)
}
