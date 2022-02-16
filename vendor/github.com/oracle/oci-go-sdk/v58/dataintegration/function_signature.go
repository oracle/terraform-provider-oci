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
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// FunctionSignature The function signature can specify function paramaters and/or function return type.
type FunctionSignature struct {

	// The key of the object.
	Key *string `mandatory:"false" json:"key"`

	// The type of the object.
	ModelType FunctionSignatureModelTypeEnum `mandatory:"false" json:"modelType,omitempty"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"false" json:"name"`

	RetType *ConfiguredType `mandatory:"false" json:"retType"`

	// An array of function arguments.
	Arguments []TypedObject `mandatory:"false" json:"arguments"`

	// Detailed description for the object.
	Description *string `mandatory:"false" json:"description"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`
}

func (m FunctionSignature) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FunctionSignature) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingFunctionSignatureModelTypeEnum(string(m.ModelType)); !ok && m.ModelType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ModelType: %s. Supported values are: %s.", m.ModelType, strings.Join(GetFunctionSignatureModelTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *FunctionSignature) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Key          *string                        `json:"key"`
		ModelType    FunctionSignatureModelTypeEnum `json:"modelType"`
		ModelVersion *string                        `json:"modelVersion"`
		ParentRef    *ParentReference               `json:"parentRef"`
		Name         *string                        `json:"name"`
		RetType      *ConfiguredType                `json:"retType"`
		Arguments    []typedobject                  `json:"arguments"`
		Description  *string                        `json:"description"`
		ObjectStatus *int                           `json:"objectStatus"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Key = model.Key

	m.ModelType = model.ModelType

	m.ModelVersion = model.ModelVersion

	m.ParentRef = model.ParentRef

	m.Name = model.Name

	m.RetType = model.RetType

	m.Arguments = make([]TypedObject, len(model.Arguments))
	for i, n := range model.Arguments {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Arguments[i] = nn.(TypedObject)
		} else {
			m.Arguments[i] = nil
		}
	}

	m.Description = model.Description

	m.ObjectStatus = model.ObjectStatus

	return
}

// FunctionSignatureModelTypeEnum Enum with underlying type: string
type FunctionSignatureModelTypeEnum string

// Set of constants representing the allowable values for FunctionSignatureModelTypeEnum
const (
	FunctionSignatureModelTypeDisFunctionSignature FunctionSignatureModelTypeEnum = "DIS_FUNCTION_SIGNATURE"
)

var mappingFunctionSignatureModelTypeEnum = map[string]FunctionSignatureModelTypeEnum{
	"DIS_FUNCTION_SIGNATURE": FunctionSignatureModelTypeDisFunctionSignature,
}

// GetFunctionSignatureModelTypeEnumValues Enumerates the set of values for FunctionSignatureModelTypeEnum
func GetFunctionSignatureModelTypeEnumValues() []FunctionSignatureModelTypeEnum {
	values := make([]FunctionSignatureModelTypeEnum, 0)
	for _, v := range mappingFunctionSignatureModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetFunctionSignatureModelTypeEnumStringValues Enumerates the set of values in String for FunctionSignatureModelTypeEnum
func GetFunctionSignatureModelTypeEnumStringValues() []string {
	return []string{
		"DIS_FUNCTION_SIGNATURE",
	}
}

// GetMappingFunctionSignatureModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFunctionSignatureModelTypeEnum(val string) (FunctionSignatureModelTypeEnum, bool) {
	mappingFunctionSignatureModelTypeEnumIgnoreCase := make(map[string]FunctionSignatureModelTypeEnum)
	for k, v := range mappingFunctionSignatureModelTypeEnum {
		mappingFunctionSignatureModelTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingFunctionSignatureModelTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
