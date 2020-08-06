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

// CreateAttributeDetails Properties used in attribute create operations.
type CreateAttributeDetails struct {

	// A user-friendly display name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Data type of the attribute as defined in the external system.
	ExternalDataType *string `mandatory:"true" json:"externalDataType"`

	// Last modified timestamp of this object in the external system.
	TimeExternal *common.SDKTime `mandatory:"true" json:"timeExternal"`

	// Detailed description of the attribute.
	Description *string `mandatory:"false" json:"description"`

	// Property that identifies if this attribute can be used as a watermark to extract incremental data.
	IsIncrementalData *bool `mandatory:"false" json:"isIncrementalData"`

	// Property that identifies if this attribute can be assigned null values.
	IsNullable *bool `mandatory:"false" json:"isNullable"`

	// Max allowed length of the attribute value.
	Length *int64 `mandatory:"false" json:"length"`

	// Position of the attribute in the record definition.
	Position *int `mandatory:"false" json:"position"`

	// Precision of the attribute value usually applies to float data type.
	Precision *int `mandatory:"false" json:"precision"`

	// Scale of the attribute value usually applies to float data type.
	Scale *int `mandatory:"false" json:"scale"`

	// The minimum count for the number of instances of a given type stored in this collection type attribute,applicable if this attribute is a complex type.
	MinCollectionCount *int `mandatory:"false" json:"minCollectionCount"`

	// The maximum count for the number of instances of a given type stored in this collection type attribute,applicable if this attribute is a complex type.
	// For type specifications in systems that specify only "capacity" without upper or lower bound , this property can also be used to just mean "capacity".
	// Some examples are Varray size in Oracle , Occurs Clause in Cobol , capacity in XmlSchemaObjectCollection , maxOccurs in  Xml , maxItems in Json
	MaxCollectionCount *int `mandatory:"false" json:"maxCollectionCount"`

	// External entity key that represents the datatype of this attribute , applicable if this attribute is a complex type.
	ExternalDatatypeEntityKey *string `mandatory:"false" json:"externalDatatypeEntityKey"`

	// External attribute key that represents the parent attribute  of this attribute , applicable if the parent attribute is of complex type.
	ExternalParentAttributeKey *string `mandatory:"false" json:"externalParentAttributeKey"`

	// A map of maps that contains the properties which are specific to the attribute type. Each attribute type
	// definition defines it's set of required and optional properties. The map keys are category names and the
	// values are maps of property name to property value. Every property is contained inside of a category. Most
	// attributes have required properties within the "default" category. To determine the set of required and
	// optional properties for an attribute type, a query can be done on '/types?type=attribute' that returns a
	// collection of all attribute types. The appropriate attribute type, which will include definitions of all
	// of it's properties, can be identified from this collection.
	// Example: `{"properties": { "default": { "key1": "value1"}}}`
	Properties map[string]map[string]string `mandatory:"false" json:"properties"`
}

func (m CreateAttributeDetails) String() string {
	return common.PointerString(m)
}
