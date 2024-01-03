// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
// For more information, see Data Catalog (https://docs.oracle.com/iaas/data-catalog/home.htm).
//

package datacatalog

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
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

	// Optional user friendly business name of the folder. If set, this supplements the harvested display name of the object.
	BusinessName *string `mandatory:"false" json:"businessName"`

	// Detailed description of a folder.
	Description *string `mandatory:"false" json:"description"`

	// The unique key of the containing folder or null if there is no parent folder.
	ParentFolderKey *string `mandatory:"false" json:"parentFolderKey"`

	// The type of folder object. Type keys can be found via the '/types' endpoint.
	TypeKey *string `mandatory:"false" json:"typeKey"`

	// The date and time the folder was harvested, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeHarvested *common.SDKTime `mandatory:"false" json:"timeHarvested"`

	// List of objects and their relationships to this folder.
	ObjectRelationships []ObjectRelationship `mandatory:"false" json:"objectRelationships"`

	// Full path of the folder.
	Path *string `mandatory:"false" json:"path"`

	// The key of the associated data asset.
	DataAssetKey *string `mandatory:"false" json:"dataAssetKey"`

	// The list of customized properties along with the values for this object
	CustomPropertyMembers []CustomPropertyGetUsage `mandatory:"false" json:"customPropertyMembers"`

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

	// A message describing the current state in more detail. An object not in ACTIVE state may have functional limitations,
	// see service documentation for details.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Status of the object as updated by the harvest process.
	HarvestStatus HarvestStatusEnum `mandatory:"false" json:"harvestStatus,omitempty"`

	// The key of the last harvest process to update the metadata of this object.
	LastJobKey *string `mandatory:"false" json:"lastJobKey"`

	// URI to the folder instance in the API.
	Uri *string `mandatory:"false" json:"uri"`

	// URL of the folder in the object store.
	ObjectStorageUrl *string `mandatory:"false" json:"objectStorageUrl"`
}

func (m Folder) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Folder) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingHarvestStatusEnum(string(m.HarvestStatus)); !ok && m.HarvestStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for HarvestStatus: %s. Supported values are: %s.", m.HarvestStatus, strings.Join(GetHarvestStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
