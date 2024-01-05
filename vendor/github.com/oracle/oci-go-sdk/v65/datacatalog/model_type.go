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

// ModelType Full data catalog type definition. Fully defines a type of the data catalog. All types are statically defined
// in the system and are immutable. It isn't possible to create new types or update existing types via the API.
type ModelType struct {

	// Unique type key that is immutable.
	Key *string `mandatory:"true" json:"key"`

	// The immutable name of the type.
	Name *string `mandatory:"false" json:"name"`

	// Detailed description of the type.
	Description *string `mandatory:"false" json:"description"`

	// The data catalog's OCID.
	CatalogId *string `mandatory:"false" json:"catalogId"`

	// A map of arrays which defines the type specific properties, both required and optional. The map keys are
	// category names and the values are arrays contiaing all property details. Every property is contained inside
	// of a category. Most types have required properties within the "default" category.
	// Example:
	// `{
	//    "properties": {
	//      "default": {
	//        "attributes:": [
	//          {
	//            "name": "host",
	//            "type": "string",
	//            "isRequired": true,
	//            "isUpdatable": false
	//          },
	//          ...
	//        ]
	//      }
	//    }
	//  }`
	Properties map[string][]PropertyDefinition `mandatory:"false" json:"properties"`

	// The current state of the type.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Indicates whether the type is internal, making it unavailable for use by metadata elements.
	IsInternal *bool `mandatory:"false" json:"isInternal"`

	// Indicates whether the type can be used for tagging metadata elements.
	IsTag *bool `mandatory:"false" json:"isTag"`

	// Indicates whether the type is approved for use as a classifying object.
	IsApproved *bool `mandatory:"false" json:"isApproved"`

	// Indicates the category this type belongs to. For instance, data assets, connections.
	TypeCategory *string `mandatory:"false" json:"typeCategory"`

	// Mapping type equivalence in the external system.
	ExternalTypeName *string `mandatory:"false" json:"externalTypeName"`

	// URI to the type instance in the API.
	Uri *string `mandatory:"false" json:"uri"`

	// Custom properties associated with this Type.
	CustomProperties []CustomPropertySummary `mandatory:"false" json:"customProperties"`

	// Unique key of the parent type.
	ParentTypeKey *string `mandatory:"false" json:"parentTypeKey"`

	// Name of the parent type.
	ParentTypeName *string `mandatory:"false" json:"parentTypeName"`
}

func (m ModelType) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ModelType) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
