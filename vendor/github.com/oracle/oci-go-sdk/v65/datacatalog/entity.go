// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// Entity Data entity details. A representation of data with a set of attributes, normally representing a single
// business entity. Synonymous with 'table' or 'view' in a database, or a single logical file structure
// that one or many files may match.
type Entity struct {

	// Unique data entity key that is immutable.
	Key *string `mandatory:"true" json:"key"`

	// A user-friendly display name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Optional user friendly business name of the data entity. If set, this supplements the harvested display name of the object.
	BusinessName *string `mandatory:"false" json:"businessName"`

	// Detailed description of a data entity.
	Description *string `mandatory:"false" json:"description"`

	// The date and time the data entity was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2019-03-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The last time that any change was made to the data entity. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// OCID of the user who created this object in the data catalog.
	CreatedById *string `mandatory:"false" json:"createdById"`

	// OCID of the user who updated this object in the data catalog.
	UpdatedById *string `mandatory:"false" json:"updatedById"`

	// The current state of the data entity.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail. An object not in ACTIVE state may have functional limitations,
	// see service documentation for details.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Unique external key of this object in the source system.
	ExternalKey *string `mandatory:"false" json:"externalKey"`

	// Key of the associated pattern if this is a logical entity.
	PatternKey *string `mandatory:"false" json:"patternKey"`

	// The expression realized after resolving qualifiers . Used in deriving this logical entity
	RealizedExpression *string `mandatory:"false" json:"realizedExpression"`

	// Last modified timestamp of this object in the external system.
	TimeExternal *common.SDKTime `mandatory:"false" json:"timeExternal"`

	// The date and time the entity was harvested, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeHarvested *common.SDKTime `mandatory:"false" json:"timeHarvested"`

	// List of objects and their relationships to this entity.
	ObjectRelationships []ObjectRelationship `mandatory:"false" json:"objectRelationships"`

	// Time that the data entities status was last updated. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeStatusUpdated *common.SDKTime `mandatory:"false" json:"timeStatusUpdated"`

	// Property that identifies if the object is a physical object (materialized) or virtual/logical object
	// defined on other objects.
	IsLogical *bool `mandatory:"false" json:"isLogical"`

	// Property that identifies if an object is a sub object of a physical or materialized parent object.
	IsPartition *bool `mandatory:"false" json:"isPartition"`

	// Unique key of the parent data asset.
	DataAssetKey *string `mandatory:"false" json:"dataAssetKey"`

	// Key of the associated folder.
	FolderKey *string `mandatory:"false" json:"folderKey"`

	// Name of the associated folder. This name is harvested from the source data asset when the parent folder for the entiy is harvested.
	FolderName *string `mandatory:"false" json:"folderName"`

	// Full path of the data entity.
	Path *string `mandatory:"false" json:"path"`

	// Status of the object as updated by the harvest process.
	HarvestStatus HarvestStatusEnum `mandatory:"false" json:"harvestStatus,omitempty"`

	// Key of the last harvest process to update this object.
	LastJobKey *string `mandatory:"false" json:"lastJobKey"`

	// The type of data entity object. Type key's can be found via the '/types' endpoint.
	TypeKey *string `mandatory:"false" json:"typeKey"`

	// URI to the data entity instance in the API.
	Uri *string `mandatory:"false" json:"uri"`

	// URL of the data entity in the object store.
	ObjectStorageUrl *string `mandatory:"false" json:"objectStorageUrl"`

	// The list of customized properties along with the values for this object
	CustomPropertyMembers []CustomPropertyGetUsage `mandatory:"false" json:"customPropertyMembers"`

	// A map of maps that contains the properties which are specific to the entity type. Each entity type
	// definition defines it's set of required and optional properties. The map keys are category names and the
	// values are maps of property name to property value. Every property is contained inside of a category. Most
	// data entities have required properties within the "default" category.
	// Example: `{"properties": { "default": { "key1": "value1"}}}`
	Properties map[string]map[string]string `mandatory:"false" json:"properties"`
}

func (m Entity) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Entity) ValidateEnumValue() (bool, error) {
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
