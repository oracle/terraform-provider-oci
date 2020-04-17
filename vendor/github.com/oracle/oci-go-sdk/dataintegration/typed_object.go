// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataintegration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// TypedObject The TypedObject class is a base class for any model object that has a type.
type TypedObject interface {

	// The key of the object.
	GetKey() *string

	// The model version of an object.
	GetModelVersion() *string

	GetParentRef() *ParentReference

	GetConfigValues() *ConfigValues

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	GetObjectStatus() *int

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value can be edited by the user and it is restricted to 1000 characters
	GetName() *string

	// Detailed description for the object.
	GetDescription() *string
}

type typedobject struct {
	JsonData     []byte
	Key          *string          `mandatory:"false" json:"key"`
	ModelVersion *string          `mandatory:"false" json:"modelVersion"`
	ParentRef    *ParentReference `mandatory:"false" json:"parentRef"`
	ConfigValues *ConfigValues    `mandatory:"false" json:"configValues"`
	ObjectStatus *int             `mandatory:"false" json:"objectStatus"`
	Name         *string          `mandatory:"false" json:"name"`
	Description  *string          `mandatory:"false" json:"description"`
	ModelType    string           `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *typedobject) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalertypedobject typedobject
	s := struct {
		Model Unmarshalertypedobject
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Key = s.Model.Key
	m.ModelVersion = s.Model.ModelVersion
	m.ParentRef = s.Model.ParentRef
	m.ConfigValues = s.Model.ConfigValues
	m.ObjectStatus = s.Model.ObjectStatus
	m.Name = s.Model.Name
	m.Description = s.Model.Description
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *typedobject) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "OUTPUT_PORT":
		mm := OutputPort{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DYNAMIC_INPUT_FIELD":
		mm := DynamicInputField{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FIELD":
		mm := AbstractField{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "INPUT_FIELD":
		mm := InputField{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SHAPE":
		mm := Shape{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "INPUT_PORT":
		mm := InputPort{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PROXY_FIELD":
		mm := ProxyField{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DYNAMIC_PROXY_FIELD":
		mm := DynamicProxyField{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SHAPE_FIELD":
		mm := ShapeField{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PARAMETER":
		mm := Parameter{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OUTPUT_FIELD":
		mm := OutputField{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DERIVED_FIELD":
		mm := DerivedField{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FLOW_PORT":
		mm := FlowPort{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetKey returns Key
func (m typedobject) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m typedobject) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m typedobject) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetConfigValues returns ConfigValues
func (m typedobject) GetConfigValues() *ConfigValues {
	return m.ConfigValues
}

//GetObjectStatus returns ObjectStatus
func (m typedobject) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetName returns Name
func (m typedobject) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m typedobject) GetDescription() *string {
	return m.Description
}

func (m typedobject) String() string {
	return common.PointerString(m)
}

// TypedObjectModelTypeEnum Enum with underlying type: string
type TypedObjectModelTypeEnum string

// Set of constants representing the allowable values for TypedObjectModelTypeEnum
const (
	TypedObjectModelTypeShape             TypedObjectModelTypeEnum = "SHAPE"
	TypedObjectModelTypeInputPort         TypedObjectModelTypeEnum = "INPUT_PORT"
	TypedObjectModelTypeShapeField        TypedObjectModelTypeEnum = "SHAPE_FIELD"
	TypedObjectModelTypeInputField        TypedObjectModelTypeEnum = "INPUT_FIELD"
	TypedObjectModelTypeDerivedField      TypedObjectModelTypeEnum = "DERIVED_FIELD"
	TypedObjectModelTypeOutputField       TypedObjectModelTypeEnum = "OUTPUT_FIELD"
	TypedObjectModelTypeDynamicProxyField TypedObjectModelTypeEnum = "DYNAMIC_PROXY_FIELD"
	TypedObjectModelTypeOutputPort        TypedObjectModelTypeEnum = "OUTPUT_PORT"
	TypedObjectModelTypeDynamicInputField TypedObjectModelTypeEnum = "DYNAMIC_INPUT_FIELD"
	TypedObjectModelTypeProxyField        TypedObjectModelTypeEnum = "PROXY_FIELD"
	TypedObjectModelTypeParameter         TypedObjectModelTypeEnum = "PARAMETER"
)

var mappingTypedObjectModelType = map[string]TypedObjectModelTypeEnum{
	"SHAPE":               TypedObjectModelTypeShape,
	"INPUT_PORT":          TypedObjectModelTypeInputPort,
	"SHAPE_FIELD":         TypedObjectModelTypeShapeField,
	"INPUT_FIELD":         TypedObjectModelTypeInputField,
	"DERIVED_FIELD":       TypedObjectModelTypeDerivedField,
	"OUTPUT_FIELD":        TypedObjectModelTypeOutputField,
	"DYNAMIC_PROXY_FIELD": TypedObjectModelTypeDynamicProxyField,
	"OUTPUT_PORT":         TypedObjectModelTypeOutputPort,
	"DYNAMIC_INPUT_FIELD": TypedObjectModelTypeDynamicInputField,
	"PROXY_FIELD":         TypedObjectModelTypeProxyField,
	"PARAMETER":           TypedObjectModelTypeParameter,
}

// GetTypedObjectModelTypeEnumValues Enumerates the set of values for TypedObjectModelTypeEnum
func GetTypedObjectModelTypeEnumValues() []TypedObjectModelTypeEnum {
	values := make([]TypedObjectModelTypeEnum, 0)
	for _, v := range mappingTypedObjectModelType {
		values = append(values, v)
	}
	return values
}
