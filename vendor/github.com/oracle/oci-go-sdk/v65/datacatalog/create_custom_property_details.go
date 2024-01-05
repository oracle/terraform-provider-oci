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

// CreateCustomPropertyDetails Properties used in custom property create operations.
type CreateCustomPropertyDetails struct {

	// A user-friendly display name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Detailed description of the custom property.
	Description *string `mandatory:"false" json:"description"`

	// The data type of the custom property
	DataType CustomPropertyDataTypeEnum `mandatory:"false" json:"dataType,omitempty"`

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

	// If this field is allowed to pop in search results
	IsHiddenInSearch *bool `mandatory:"false" json:"isHiddenInSearch"`

	// If an OCI Event will be emitted when the custom property is modified.
	IsEventEnabled *bool `mandatory:"false" json:"isEventEnabled"`

	// Allowed values for the custom property if any
	AllowedValues []string `mandatory:"false" json:"allowedValues"`

	// A map of maps that contains the properties which are specific to the data asset type. Each data asset type
	// definition defines it's set of required and optional properties. The map keys are category names and the
	// values are maps of property name to property value. Every property is contained inside of a category. Most
	// data assets have required properties within the "default" category. To determine the set of optional and
	// required properties for a data asset type, a query can be done on '/types?type=dataAsset' that returns a
	// collection of all data asset types. The appropriate data asset type, which includes definitions of all of
	// it's properties, can be identified from this collection.
	// Example: `{"properties": { "default": { "host": "host1", "port": "1521", "database": "orcl"}}}`
	Properties map[string]map[string]string `mandatory:"false" json:"properties"`
}

func (m CreateCustomPropertyDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateCustomPropertyDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCustomPropertyDataTypeEnum(string(m.DataType)); !ok && m.DataType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DataType: %s. Supported values are: %s.", m.DataType, strings.Join(GetCustomPropertyDataTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
