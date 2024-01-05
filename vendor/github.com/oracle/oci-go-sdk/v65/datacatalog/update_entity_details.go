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

// UpdateEntityDetails Properties used in entity update operations.
type UpdateEntityDetails struct {

	// A user-friendly display name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Optional user friendly business name of the data entity. If set, this supplements the harvested display name of the object.
	BusinessName *string `mandatory:"false" json:"businessName"`

	// Detailed description of a data entity.
	Description *string `mandatory:"false" json:"description"`

	// Last modified timestamp of the object in the external system.
	TimeExternal *common.SDKTime `mandatory:"false" json:"timeExternal"`

	// Property to indicate if the object is a physical materialized object or virtual. For example, View.
	IsLogical *bool `mandatory:"false" json:"isLogical"`

	// Property to indicate if the object is a sub object of a parent physical object.
	IsPartition *bool `mandatory:"false" json:"isPartition"`

	// Key of the associated folder.
	FolderKey *string `mandatory:"false" json:"folderKey"`

	// Key of the associated pattern if this is a logical entity.
	PatternKey *string `mandatory:"false" json:"patternKey"`

	// The expression realized after resolving qualifiers . Used in deriving this logical entity
	RealizedExpression *string `mandatory:"false" json:"realizedExpression"`

	// Status of the object as updated by the harvest process. When an entity object is created, it's harvest status
	// will indicate if the entity's metadata has been fully harvested or not. The harvest process can perform
	// shallow harvesting to allow users to browse the metadata and can on-demand deep harvest on any object
	// This requires a harvest status indicator for catalog objects.
	HarvestStatus HarvestStatusEnum `mandatory:"false" json:"harvestStatus,omitempty"`

	// Key of the last harvest process to update this object.
	LastJobKey *string `mandatory:"false" json:"lastJobKey"`

	// The list of customized properties along with the values for this object
	CustomPropertyMembers []CustomPropertySetUsage `mandatory:"false" json:"customPropertyMembers"`

	// A map of maps that contains the properties which are specific to the entity type. Each entity type
	// definition defines it's set of required and optional properties. The map keys are category names and the
	// values are maps of property name to property value. Every property is contained inside of a category. Most
	// entities have required properties within the "default" category. To determine the set of required and
	// optional properties for an entity type, a query can be done on '/types?type=dataEntity' that returns a
	// collection of all entity types. The appropriate entity type, which includes definitions of all of
	// it's properties, can be identified from this collection.
	// Example: `{"properties": { "default": { "key1": "value1"}}}`
	Properties map[string]map[string]string `mandatory:"false" json:"properties"`
}

func (m UpdateEntityDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateEntityDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingHarvestStatusEnum(string(m.HarvestStatus)); !ok && m.HarvestStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for HarvestStatus: %s. Supported values are: %s.", m.HarvestStatus, strings.Join(GetHarvestStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
