// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
// For more information, see Data Catalog (https://docs.oracle.com/iaas/data-catalog/home.htm).
//

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// CustomProperty Custom Property Definition
type CustomProperty struct {

	// Unique data asset key that is immutable.
	Key *string `mandatory:"true" json:"key"`

	// Display name of the custom property
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Data type of the custom property
	DataType CustomPropertyDataTypeEnum `mandatory:"false" json:"dataType,omitempty"`

	// Description for the custom property
	Description *string `mandatory:"false" json:"description"`

	// Namespace name of the custom property
	NamespaceName *string `mandatory:"false" json:"namespaceName"`

	// Is this property allowed to have list of values
	IsListType *bool `mandatory:"false" json:"isListType"`

	// If this field allows to sort from UI
	IsSortable *bool `mandatory:"false" json:"isSortable"`

	// If this field allows to filter or create facets from UI
	IsFilterable *bool `mandatory:"false" json:"isFilterable"`

	// If this field allows multiple values to be set
	IsMultiValued *bool `mandatory:"false" json:"isMultiValued"`

	// If this field is a hidden field
	IsHidden *bool `mandatory:"false" json:"isHidden"`

	// If this field is a editable field
	IsEditable *bool `mandatory:"false" json:"isEditable"`

	// If this field is displayed in a list view of applicable objects.
	IsShownInList *bool `mandatory:"false" json:"isShownInList"`

	// If this field is defined by service or by a user
	IsServiceDefined *bool `mandatory:"false" json:"isServiceDefined"`

	// If this field is allowed to pop in search results
	IsHiddenInSearch *bool `mandatory:"false" json:"isHiddenInSearch"`

	// The current state of the custom property.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The date and time the custom property was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2019-03-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The last time that any change was made to the custom property. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// OCID of the user who created the custom property.
	CreatedById *string `mandatory:"false" json:"createdById"`

	// OCID of the user who last modified the custom property.
	UpdatedById *string `mandatory:"false" json:"updatedById"`

	// Total number of first class objects using this custom property
	UsageCount *int `mandatory:"false" json:"usageCount"`

	// If an OCI Event will be emitted when the custom property is modified.
	IsEventEnabled *bool `mandatory:"false" json:"isEventEnabled"`

	// The set of object types to which the custom property applies.
	Scope []CustomPropertyTypeUsage `mandatory:"false" json:"scope"`

	// Allowed values for the custom property if any
	AllowedValues []string `mandatory:"false" json:"allowedValues"`

	// Event configuration for this custom property, against the desired subset of object types to which the property applies.
	Events []EventConfig `mandatory:"false" json:"events"`

	// A map of maps that contains the properties which are specific to the asset type. Each data asset type
	// definition defines it's set of required and optional properties. The map keys are category names and the
	// values are maps of property name to property value. Every property is contained inside of a category. Most
	// data assets have required properties within the "default" category.
	// Example: `{"properties": { "default": { "host": "host1", "port": "1521", "database": "orcl"}}}`
	Properties map[string]map[string]string `mandatory:"false" json:"properties"`
}

func (m CustomProperty) String() string {
	return common.PointerString(m)
}
