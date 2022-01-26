// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// StructuredType A `StructuredType` object represents a data type that exists in a physical data asset object such as a table column, but is more complex. For example, an Oracle database `OBJECT` type. It can be composed of multiple `DataType` objects.
type StructuredType struct {

	// The property which disciminates the subtypes.
	ModelType StructuredTypeModelTypeEnum `mandatory:"true" json:"modelType"`

	// The key of the object.
	Key *string `mandatory:"false" json:"key"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"false" json:"name"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// A user defined description for the object.
	Description *string `mandatory:"false" json:"description"`

	// The data type.
	DtType StructuredTypeDtTypeEnum `mandatory:"false" json:"dtType,omitempty"`

	// The data type system name.
	TypeSystemName *string `mandatory:"false" json:"typeSystemName"`

	ConfigDefinition *ConfigDefinition `mandatory:"false" json:"configDefinition"`

	Schema BaseType `mandatory:"false" json:"schema"`
}

func (m StructuredType) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *StructuredType) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Key              *string                     `json:"key"`
		ModelVersion     *string                     `json:"modelVersion"`
		ParentRef        *ParentReference            `json:"parentRef"`
		Name             *string                     `json:"name"`
		ObjectStatus     *int                        `json:"objectStatus"`
		Description      *string                     `json:"description"`
		DtType           StructuredTypeDtTypeEnum    `json:"dtType"`
		TypeSystemName   *string                     `json:"typeSystemName"`
		ConfigDefinition *ConfigDefinition           `json:"configDefinition"`
		Schema           basetype                    `json:"schema"`
		ModelType        StructuredTypeModelTypeEnum `json:"modelType"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Key = model.Key

	m.ModelVersion = model.ModelVersion

	m.ParentRef = model.ParentRef

	m.Name = model.Name

	m.ObjectStatus = model.ObjectStatus

	m.Description = model.Description

	m.DtType = model.DtType

	m.TypeSystemName = model.TypeSystemName

	m.ConfigDefinition = model.ConfigDefinition

	nn, e = model.Schema.UnmarshalPolymorphicJSON(model.Schema.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Schema = nn.(BaseType)
	} else {
		m.Schema = nil
	}

	m.ModelType = model.ModelType

	return
}

// StructuredTypeModelTypeEnum Enum with underlying type: string
type StructuredTypeModelTypeEnum string

// Set of constants representing the allowable values for StructuredTypeModelTypeEnum
const (
	StructuredTypeModelTypeDynamicType    StructuredTypeModelTypeEnum = "DYNAMIC_TYPE"
	StructuredTypeModelTypeStructuredType StructuredTypeModelTypeEnum = "STRUCTURED_TYPE"
	StructuredTypeModelTypeDataType       StructuredTypeModelTypeEnum = "DATA_TYPE"
	StructuredTypeModelTypeJavaType       StructuredTypeModelTypeEnum = "JAVA_TYPE"
	StructuredTypeModelTypeConfiguredType StructuredTypeModelTypeEnum = "CONFIGURED_TYPE"
	StructuredTypeModelTypeCompositeType  StructuredTypeModelTypeEnum = "COMPOSITE_TYPE"
	StructuredTypeModelTypeDerivedType    StructuredTypeModelTypeEnum = "DERIVED_TYPE"
)

var mappingStructuredTypeModelType = map[string]StructuredTypeModelTypeEnum{
	"DYNAMIC_TYPE":    StructuredTypeModelTypeDynamicType,
	"STRUCTURED_TYPE": StructuredTypeModelTypeStructuredType,
	"DATA_TYPE":       StructuredTypeModelTypeDataType,
	"JAVA_TYPE":       StructuredTypeModelTypeJavaType,
	"CONFIGURED_TYPE": StructuredTypeModelTypeConfiguredType,
	"COMPOSITE_TYPE":  StructuredTypeModelTypeCompositeType,
	"DERIVED_TYPE":    StructuredTypeModelTypeDerivedType,
}

// GetStructuredTypeModelTypeEnumValues Enumerates the set of values for StructuredTypeModelTypeEnum
func GetStructuredTypeModelTypeEnumValues() []StructuredTypeModelTypeEnum {
	values := make([]StructuredTypeModelTypeEnum, 0)
	for _, v := range mappingStructuredTypeModelType {
		values = append(values, v)
	}
	return values
}

// StructuredTypeDtTypeEnum Enum with underlying type: string
type StructuredTypeDtTypeEnum string

// Set of constants representing the allowable values for StructuredTypeDtTypeEnum
const (
	StructuredTypeDtTypePrimitive  StructuredTypeDtTypeEnum = "PRIMITIVE"
	StructuredTypeDtTypeStructured StructuredTypeDtTypeEnum = "STRUCTURED"
)

var mappingStructuredTypeDtType = map[string]StructuredTypeDtTypeEnum{
	"PRIMITIVE":  StructuredTypeDtTypePrimitive,
	"STRUCTURED": StructuredTypeDtTypeStructured,
}

// GetStructuredTypeDtTypeEnumValues Enumerates the set of values for StructuredTypeDtTypeEnum
func GetStructuredTypeDtTypeEnumValues() []StructuredTypeDtTypeEnum {
	values := make([]StructuredTypeDtTypeEnum, 0)
	for _, v := range mappingStructuredTypeDtType {
		values = append(values, v)
	}
	return values
}
