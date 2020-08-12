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

// Attribute Details of an entity attribute. An attribute of a data entity describing an item of data,
// with a name and data type. Synonymous with 'column' in a database.
type Attribute struct {

	// Unique attribute key that is immutable.
	Key *string `mandatory:"true" json:"key"`

	// A user-friendly display name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Detailed description of the attribute.
	Description *string `mandatory:"false" json:"description"`

	// The unique key of the parent entity.
	EntityKey *string `mandatory:"false" json:"entityKey"`

	// State of the attribute.
	LifecycleState LifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The date and time the attribute was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2019-03-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The last time that any change was made to the attribute. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// OCID of the user who created this attribute in the data catalog.
	CreatedById *string `mandatory:"false" json:"createdById"`

	// OCID of the user who modified this attribute in the data catalog.
	UpdatedById *string `mandatory:"false" json:"updatedById"`

	// Data type of the attribute as defined in the external system. Type mapping across systems can be achieved
	// through term associations across domains in the ontology. The attribute can also be tagged to the datatype in
	// the domain ontology to resolve any ambiguity arising from type name similarity that can occur with user
	// defined types.
	ExternalDataType *string `mandatory:"false" json:"externalDataType"`

	// Unique external key of this attribute in the external source system.
	ExternalKey *string `mandatory:"false" json:"externalKey"`

	// Property that identifies if this attribute can be used as a watermark to extract incremental data.
	IsIncrementalData *bool `mandatory:"false" json:"isIncrementalData"`

	// Property that identifies if this attribute can be assigned null values.
	IsNullable *bool `mandatory:"false" json:"isNullable"`

	// The minimum count for the number of instances of a given type stored in this collection type attribute,applicable if this attribute is a complex type.
	MinCollectionCount *int `mandatory:"false" json:"minCollectionCount"`

	// The maximum count for the number of instances of a given type stored in this collection type attribute,applicable if this attribute is a complex type.
	// For type specifications in systems that specify only "capacity" without upper or lower bound , this property can also be used to just mean "capacity".
	// Some examples are Varray size in Oracle , Occurs Clause in Cobol , capacity in XmlSchemaObjectCollection , maxOccurs in  Xml , maxItems in Json
	MaxCollectionCount *int `mandatory:"false" json:"maxCollectionCount"`

	// Entity key that represents the datatype of this attribute , applicable if this attribute is a complex type.
	DatatypeEntityKey *string `mandatory:"false" json:"datatypeEntityKey"`

	// External entity key that represents the datatype of this attribute , applicable if this attribute is a complex type.
	ExternalDatatypeEntityKey *string `mandatory:"false" json:"externalDatatypeEntityKey"`

	// Attribute key that represents the parent attribute of this attribute , applicable if the parent attribute is of complex datatype.
	ParentAttributeKey *string `mandatory:"false" json:"parentAttributeKey"`

	// External attribute key that represents the parent attribute  of this attribute , applicable if the parent attribute is of complex type.
	ExternalParentAttributeKey *string `mandatory:"false" json:"externalParentAttributeKey"`

	// Max allowed length of the attribute value.
	Length *int64 `mandatory:"false" json:"length"`

	// Position of the attribute in the record definition.
	Position *int `mandatory:"false" json:"position"`

	// Precision of the attribute value usually applies to float data type.
	Precision *int `mandatory:"false" json:"precision"`

	// Scale of the attribute value usually applies to float data type.
	Scale *int `mandatory:"false" json:"scale"`

	// Last modified timestamp of this object in the external system.
	TimeExternal *common.SDKTime `mandatory:"false" json:"timeExternal"`

	// URI to the attribute instance in the API.
	Uri *string `mandatory:"false" json:"uri"`

	// Full path of the attribute.
	Path *string `mandatory:"false" json:"path"`

	// A map of maps that contains the properties which are specific to the attribute type. Each attribute type
	// definition defines it's set of required and optional properties. The map keys are category names and the
	// values are maps of property name to property value. Every property is contained inside of a category. Most
	// attributes have required properties within the "default" category.
	// Example: `{"properties": { "default": { "key1": "value1"}}}`
	Properties map[string]map[string]string `mandatory:"false" json:"properties"`
}

func (m Attribute) String() string {
	return common.PointerString(m)
}
