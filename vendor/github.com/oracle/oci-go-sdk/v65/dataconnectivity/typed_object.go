// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Connectivity Management API
//
// Use the Data Connectivity Management Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataconnectivity

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TypedObject The `TypedObject` class is a base class for any model object that has a type.
type TypedObject interface {

	// The key of the object.
	GetKey() *string

	// The model version of an object.
	GetModelVersion() *string

	GetParentRef() *ParentReference

	GetConfigValues() *ConfigValues

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	GetObjectStatus() *int

	// Free form text without any restriction on the permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	GetName() *string

	// A detailed description of the object.
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
	case "SHAPE":
		mm := Shape{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SHAPE_FIELD":
		mm := ShapeField{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "INPUT_PORT":
		mm := InputPort{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PARAMETER":
		mm := Parameter{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "NATIVE_SHAPE_FIELD":
		mm := NativeShapeField{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for TypedObject: %s.", m.ModelType)
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m typedobject) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TypedObjectModelTypeEnum Enum with underlying type: string
type TypedObjectModelTypeEnum string

// Set of constants representing the allowable values for TypedObjectModelTypeEnum
const (
	TypedObjectModelTypeShape            TypedObjectModelTypeEnum = "SHAPE"
	TypedObjectModelTypeShapeField       TypedObjectModelTypeEnum = "SHAPE_FIELD"
	TypedObjectModelTypeNativeShapeField TypedObjectModelTypeEnum = "NATIVE_SHAPE_FIELD"
)

var mappingTypedObjectModelTypeEnum = map[string]TypedObjectModelTypeEnum{
	"SHAPE":              TypedObjectModelTypeShape,
	"SHAPE_FIELD":        TypedObjectModelTypeShapeField,
	"NATIVE_SHAPE_FIELD": TypedObjectModelTypeNativeShapeField,
}

var mappingTypedObjectModelTypeEnumLowerCase = map[string]TypedObjectModelTypeEnum{
	"shape":              TypedObjectModelTypeShape,
	"shape_field":        TypedObjectModelTypeShapeField,
	"native_shape_field": TypedObjectModelTypeNativeShapeField,
}

// GetTypedObjectModelTypeEnumValues Enumerates the set of values for TypedObjectModelTypeEnum
func GetTypedObjectModelTypeEnumValues() []TypedObjectModelTypeEnum {
	values := make([]TypedObjectModelTypeEnum, 0)
	for _, v := range mappingTypedObjectModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTypedObjectModelTypeEnumStringValues Enumerates the set of values in String for TypedObjectModelTypeEnum
func GetTypedObjectModelTypeEnumStringValues() []string {
	return []string{
		"SHAPE",
		"SHAPE_FIELD",
		"NATIVE_SHAPE_FIELD",
	}
}

// GetMappingTypedObjectModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTypedObjectModelTypeEnum(val string) (TypedObjectModelTypeEnum, bool) {
	enum, ok := mappingTypedObjectModelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
