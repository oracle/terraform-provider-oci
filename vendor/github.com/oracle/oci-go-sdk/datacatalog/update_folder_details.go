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

// UpdateFolderDetails Properties used in folder update operations.
type UpdateFolderDetails struct {

	// A user-friendly display name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Detailed description of a folder.
	Description *string `mandatory:"false" json:"description"`

	// The key of the containing folder.
	ParentFolderKey *string `mandatory:"false" json:"parentFolderKey"`

	// A map of maps that contains the properties which are specific to the folder type. Each folder type
	// definition defines it's set of required and optional properties. The map keys are category names and the
	// values are maps of property name to property value. Every property is contained inside of a category. Most
	// folders have required properties within the "default" category. To determine the set of optional and
	// required properties for a folder type, a query can be done on '/types?type=folder' that returns a
	// collection of all folder types. The appropriate folder type, which includes definitions of all of
	// it's properties, can be identified from this collection.
	// Example: `{"properties": { "default": { "key1": "value1"}}}`
	Properties map[string]map[string]string `mandatory:"false" json:"properties"`

	// Last modified timestamp of this object in the external system.
	TimeExternal *common.SDKTime `mandatory:"false" json:"timeExternal"`

	// Harvest status of the folder.
	HarvestStatus HarvestStatusEnum `mandatory:"false" json:"harvestStatus,omitempty"`

	// The key of the last harvest process to update the metadata of this object.
	LastJobKey *string `mandatory:"false" json:"lastJobKey"`
}

func (m UpdateFolderDetails) String() string {
	return common.PointerString(m)
}
