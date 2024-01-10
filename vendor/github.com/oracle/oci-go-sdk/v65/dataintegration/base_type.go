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

// BaseType Base type for the type system.
type BaseType interface {

	// The key of the object.
	GetKey() *string

	// The model version of an object.
	GetModelVersion() *string

	GetParentRef() *ParentReference

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	GetName() *string

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	GetObjectStatus() *int

	// A user defined description for the object.
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
	case "ARRAY_TYPE":
		mm := ArrayType{}
		err = json.Unmarshal(data, &mm)
		return mm, err
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
	case "MATERIALIZED_COMPOSITE_TYPE":
		mm := MaterializedCompositeType{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MAP_TYPE":
		mm := MapType{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "COMPOSITE_TYPE":
		mm := CompositeType{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for BaseType: %s.", m.ModelType)
		return *m, nil
	}
}

// GetKey returns Key
func (m basetype) GetKey() *string {
	return m.Key
}

// GetModelVersion returns ModelVersion
func (m basetype) GetModelVersion() *string {
	return m.ModelVersion
}

// GetParentRef returns ParentRef
func (m basetype) GetParentRef() *ParentReference {
	return m.ParentRef
}

// GetName returns Name
func (m basetype) GetName() *string {
	return m.Name
}

// GetObjectStatus returns ObjectStatus
func (m basetype) GetObjectStatus() *int {
	return m.ObjectStatus
}

// GetDescription returns Description
func (m basetype) GetDescription() *string {
	return m.Description
}

func (m basetype) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m basetype) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BaseTypeModelTypeEnum Enum with underlying type: string
type BaseTypeModelTypeEnum string

// Set of constants representing the allowable values for BaseTypeModelTypeEnum
const (
	BaseTypeModelTypeDynamicType               BaseTypeModelTypeEnum = "DYNAMIC_TYPE"
	BaseTypeModelTypeStructuredType            BaseTypeModelTypeEnum = "STRUCTURED_TYPE"
	BaseTypeModelTypeDataType                  BaseTypeModelTypeEnum = "DATA_TYPE"
	BaseTypeModelTypeJavaType                  BaseTypeModelTypeEnum = "JAVA_TYPE"
	BaseTypeModelTypeConfiguredType            BaseTypeModelTypeEnum = "CONFIGURED_TYPE"
	BaseTypeModelTypeCompositeType             BaseTypeModelTypeEnum = "COMPOSITE_TYPE"
	BaseTypeModelTypeDerivedType               BaseTypeModelTypeEnum = "DERIVED_TYPE"
	BaseTypeModelTypeArrayType                 BaseTypeModelTypeEnum = "ARRAY_TYPE"
	BaseTypeModelTypeMapType                   BaseTypeModelTypeEnum = "MAP_TYPE"
	BaseTypeModelTypeMaterializedCompositeType BaseTypeModelTypeEnum = "MATERIALIZED_COMPOSITE_TYPE"
)

var mappingBaseTypeModelTypeEnum = map[string]BaseTypeModelTypeEnum{
	"DYNAMIC_TYPE":                BaseTypeModelTypeDynamicType,
	"STRUCTURED_TYPE":             BaseTypeModelTypeStructuredType,
	"DATA_TYPE":                   BaseTypeModelTypeDataType,
	"JAVA_TYPE":                   BaseTypeModelTypeJavaType,
	"CONFIGURED_TYPE":             BaseTypeModelTypeConfiguredType,
	"COMPOSITE_TYPE":              BaseTypeModelTypeCompositeType,
	"DERIVED_TYPE":                BaseTypeModelTypeDerivedType,
	"ARRAY_TYPE":                  BaseTypeModelTypeArrayType,
	"MAP_TYPE":                    BaseTypeModelTypeMapType,
	"MATERIALIZED_COMPOSITE_TYPE": BaseTypeModelTypeMaterializedCompositeType,
}

var mappingBaseTypeModelTypeEnumLowerCase = map[string]BaseTypeModelTypeEnum{
	"dynamic_type":                BaseTypeModelTypeDynamicType,
	"structured_type":             BaseTypeModelTypeStructuredType,
	"data_type":                   BaseTypeModelTypeDataType,
	"java_type":                   BaseTypeModelTypeJavaType,
	"configured_type":             BaseTypeModelTypeConfiguredType,
	"composite_type":              BaseTypeModelTypeCompositeType,
	"derived_type":                BaseTypeModelTypeDerivedType,
	"array_type":                  BaseTypeModelTypeArrayType,
	"map_type":                    BaseTypeModelTypeMapType,
	"materialized_composite_type": BaseTypeModelTypeMaterializedCompositeType,
}

// GetBaseTypeModelTypeEnumValues Enumerates the set of values for BaseTypeModelTypeEnum
func GetBaseTypeModelTypeEnumValues() []BaseTypeModelTypeEnum {
	values := make([]BaseTypeModelTypeEnum, 0)
	for _, v := range mappingBaseTypeModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetBaseTypeModelTypeEnumStringValues Enumerates the set of values in String for BaseTypeModelTypeEnum
func GetBaseTypeModelTypeEnumStringValues() []string {
	return []string{
		"DYNAMIC_TYPE",
		"STRUCTURED_TYPE",
		"DATA_TYPE",
		"JAVA_TYPE",
		"CONFIGURED_TYPE",
		"COMPOSITE_TYPE",
		"DERIVED_TYPE",
		"ARRAY_TYPE",
		"MAP_TYPE",
		"MATERIALIZED_COMPOSITE_TYPE",
	}
}

// GetMappingBaseTypeModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBaseTypeModelTypeEnum(val string) (BaseTypeModelTypeEnum, bool) {
	enum, ok := mappingBaseTypeModelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
