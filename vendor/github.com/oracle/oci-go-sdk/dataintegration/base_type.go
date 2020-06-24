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

// BaseType Base type for the type system
type BaseType interface {

	// The key of the object.
	GetKey() *string

	// The model version of an object.
	GetModelVersion() *string

	GetParentRef() *ParentReference

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value can be edited by the user and it is restricted to 1000 characters
	GetName() *string

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	GetObjectStatus() *int

	// Detailed description for the object.
	GetDescription() *string
}

type basetype struct {
	JsonData     []byte
	Key          *string          `mandatory:"false" json:"key"`
	ModelVersion *string          `mandatory:"false" json:"modelVersion"`
	ParentRef    *ParentReference `mandatory:"false" json:"parentRef"`
	Name         *string          `mandatory:"false" json:"name"`
	ObjectStatus *int             `mandatory:"false" json:"objectStatus"`
	Description  *string          `mandatory:"false" json:"description"`
	ModelType    string           `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *basetype) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerbasetype basetype
	s := struct {
		Model Unmarshalerbasetype
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Key = s.Model.Key
	m.ModelVersion = s.Model.ModelVersion
	m.ParentRef = s.Model.ParentRef
	m.Name = s.Model.Name
	m.ObjectStatus = s.Model.ObjectStatus
	m.Description = s.Model.Description
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *basetype) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "CONFIGURED_TYPE":
		mm := ConfiguredType{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "JAVA_TYPE":
		mm := JavaType{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DYNAMIC_TYPE":
		mm := DynamicType{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DERIVED_TYPE":
		mm := DerivedType{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DATA_TYPE":
		mm := DataType{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "COMPOSITE_TYPE":
		mm := CompositeType{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetKey returns Key
func (m basetype) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m basetype) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m basetype) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetName returns Name
func (m basetype) GetName() *string {
	return m.Name
}

//GetObjectStatus returns ObjectStatus
func (m basetype) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetDescription returns Description
func (m basetype) GetDescription() *string {
	return m.Description
}

func (m basetype) String() string {
	return common.PointerString(m)
}

// BaseTypeModelTypeEnum Enum with underlying type: string
type BaseTypeModelTypeEnum string

// Set of constants representing the allowable values for BaseTypeModelTypeEnum
const (
	BaseTypeModelTypeDynamicType    BaseTypeModelTypeEnum = "DYNAMIC_TYPE"
	BaseTypeModelTypeStructuredType BaseTypeModelTypeEnum = "STRUCTURED_TYPE"
	BaseTypeModelTypeDataType       BaseTypeModelTypeEnum = "DATA_TYPE"
	BaseTypeModelTypeJavaType       BaseTypeModelTypeEnum = "JAVA_TYPE"
	BaseTypeModelTypeConfiguredType BaseTypeModelTypeEnum = "CONFIGURED_TYPE"
	BaseTypeModelTypeCompositeType  BaseTypeModelTypeEnum = "COMPOSITE_TYPE"
)

var mappingBaseTypeModelType = map[string]BaseTypeModelTypeEnum{
	"DYNAMIC_TYPE":    BaseTypeModelTypeDynamicType,
	"STRUCTURED_TYPE": BaseTypeModelTypeStructuredType,
	"DATA_TYPE":       BaseTypeModelTypeDataType,
	"JAVA_TYPE":       BaseTypeModelTypeJavaType,
	"CONFIGURED_TYPE": BaseTypeModelTypeConfiguredType,
	"COMPOSITE_TYPE":  BaseTypeModelTypeCompositeType,
}

// GetBaseTypeModelTypeEnumValues Enumerates the set of values for BaseTypeModelTypeEnum
func GetBaseTypeModelTypeEnumValues() []BaseTypeModelTypeEnum {
	values := make([]BaseTypeModelTypeEnum, 0)
	for _, v := range mappingBaseTypeModelType {
		values = append(values, v)
	}
	return values
}
