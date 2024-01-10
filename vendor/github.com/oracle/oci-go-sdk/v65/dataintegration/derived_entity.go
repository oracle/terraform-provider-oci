// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DerivedEntity The Derive entity object
type DerivedEntity struct {
	Metadata *ObjectMetadata `mandatory:"false" json:"metadata"`

	// The object key.
	Key *string `mandatory:"false" json:"key"`

	// The object's model version.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"false" json:"name"`

	// The version of the object that is used to track changes in the object instance.
	ObjectVersion *int `mandatory:"false" json:"objectVersion"`

	Shape *Shape `mandatory:"false" json:"shape"`

	// The shape ID.
	ShapeId *string `mandatory:"false" json:"shapeId"`

	// The resource name.
	ResourceName *string `mandatory:"false" json:"resourceName"`

	// The status of an object that can be set to value 1 for shallow reference across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
	Identifier *string `mandatory:"false" json:"identifier"`

	RefDataObject ReferencedDataObject `mandatory:"false" json:"refDataObject"`

	// Property-bag (key-value pairs where key is Shape Field resource name and value is object)
	DerivedProperties map[string]interface{} `mandatory:"false" json:"derivedProperties"`

	// Determines whether entity is treated as source or target
	Mode DerivedEntityModeEnum `mandatory:"false" json:"mode,omitempty"`
}

// GetMetadata returns Metadata
func (m DerivedEntity) GetMetadata() *ObjectMetadata {
	return m.Metadata
}

func (m DerivedEntity) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DerivedEntity) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDerivedEntityModeEnum(string(m.Mode)); !ok && m.Mode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Mode: %s. Supported values are: %s.", m.Mode, strings.Join(GetDerivedEntityModeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DerivedEntity) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDerivedEntity DerivedEntity
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeDerivedEntity
	}{
		"DERIVED_ENTITY",
		(MarshalTypeDerivedEntity)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *DerivedEntity) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Metadata          *ObjectMetadata        `json:"metadata"`
		Key               *string                `json:"key"`
		ModelVersion      *string                `json:"modelVersion"`
		ParentRef         *ParentReference       `json:"parentRef"`
		Name              *string                `json:"name"`
		ObjectVersion     *int                   `json:"objectVersion"`
		Shape             *Shape                 `json:"shape"`
		ShapeId           *string                `json:"shapeId"`
		ResourceName      *string                `json:"resourceName"`
		ObjectStatus      *int                   `json:"objectStatus"`
		Identifier        *string                `json:"identifier"`
		RefDataObject     referenceddataobject   `json:"refDataObject"`
		Mode              DerivedEntityModeEnum  `json:"mode"`
		DerivedProperties map[string]interface{} `json:"derivedProperties"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Metadata = model.Metadata

	m.Key = model.Key

	m.ModelVersion = model.ModelVersion

	m.ParentRef = model.ParentRef

	m.Name = model.Name

	m.ObjectVersion = model.ObjectVersion

	m.Shape = model.Shape

	m.ShapeId = model.ShapeId

	m.ResourceName = model.ResourceName

	m.ObjectStatus = model.ObjectStatus

	m.Identifier = model.Identifier

	nn, e = model.RefDataObject.UnmarshalPolymorphicJSON(model.RefDataObject.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.RefDataObject = nn.(ReferencedDataObject)
	} else {
		m.RefDataObject = nil
	}

	m.Mode = model.Mode

	m.DerivedProperties = model.DerivedProperties

	return
}

// DerivedEntityModeEnum Enum with underlying type: string
type DerivedEntityModeEnum string

// Set of constants representing the allowable values for DerivedEntityModeEnum
const (
	DerivedEntityModeIn  DerivedEntityModeEnum = "IN"
	DerivedEntityModeOut DerivedEntityModeEnum = "OUT"
)

var mappingDerivedEntityModeEnum = map[string]DerivedEntityModeEnum{
	"IN":  DerivedEntityModeIn,
	"OUT": DerivedEntityModeOut,
}

var mappingDerivedEntityModeEnumLowerCase = map[string]DerivedEntityModeEnum{
	"in":  DerivedEntityModeIn,
	"out": DerivedEntityModeOut,
}

// GetDerivedEntityModeEnumValues Enumerates the set of values for DerivedEntityModeEnum
func GetDerivedEntityModeEnumValues() []DerivedEntityModeEnum {
	values := make([]DerivedEntityModeEnum, 0)
	for _, v := range mappingDerivedEntityModeEnum {
		values = append(values, v)
	}
	return values
}

// GetDerivedEntityModeEnumStringValues Enumerates the set of values in String for DerivedEntityModeEnum
func GetDerivedEntityModeEnumStringValues() []string {
	return []string{
		"IN",
		"OUT",
	}
}

// GetMappingDerivedEntityModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDerivedEntityModeEnum(val string) (DerivedEntityModeEnum, bool) {
	enum, ok := mappingDerivedEntityModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
