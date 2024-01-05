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

// LineageObject Object describing an individual element of object lineage.
type LineageObject struct {

	// Key of the object, such as an entity, about which this lineage applies.
	ObjectKey *string `mandatory:"false" json:"objectKey"`

	// Display name of the object.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Detailed description of the object.
	Description *string `mandatory:"false" json:"description"`

	// Indicates if intra-lineage is available for this given object. If yes, drill-down can be requested for
	// this object.
	IsIntraLineageAvailable *bool `mandatory:"false" json:"isIntraLineageAvailable"`

	// Key of the parent object for this object.
	ParentKey *string `mandatory:"false" json:"parentKey"`

	// Full path of the parent object.
	ParentPath *string `mandatory:"false" json:"parentPath"`

	// The time that this object was created. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time that this object was updated. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Type name of the object. Type keys can be found via the '/types' endpoint.
	TypeName *string `mandatory:"false" json:"typeName"`

	// Type key of the object. Type keys can be found via the '/types' endpoint.
	TypeKey *string `mandatory:"false" json:"typeKey"`

	// A map of maps that contains the properties which are specific to the entity type. Each entity type
	// definition defines it's set of required and optional properties. The map keys are category names and the
	// values are maps of property name to property value. Every property is contained inside of a category. Most
	// data entities have required properties within the "default" category.
	// Example: `{"properties": { "default": { "key1": "value1"}}}`
	Properties map[string]map[string]string `mandatory:"false" json:"properties"`
}

func (m LineageObject) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LineageObject) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
