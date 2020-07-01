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

// Parameter Parameters are created and assigned values that can be deferred to execution/runtime.
type Parameter struct {

	// The key of the object.
	Key *string `mandatory:"false" json:"key"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	ConfigValues *ConfigValues `mandatory:"false" json:"configValues"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value can be edited by the user and it is restricted to 1000 characters
	Name *string `mandatory:"false" json:"name"`

	// Detailed description for the object.
	Description *string `mandatory:"false" json:"description"`

	Type BaseType `mandatory:"false" json:"type"`

	// The default value of the parameter.
	DefaultValue *interface{} `mandatory:"false" json:"defaultValue"`

	// The default value of the parameter which can be an object in DIS, such as a data entity.
	RootObjectDefaultValue *interface{} `mandatory:"false" json:"rootObjectDefaultValue"`

	// Whether the parameter is input value.
	IsInput *bool `mandatory:"false" json:"isInput"`

	// Whether the parameter is output value.
	IsOutput *bool `mandatory:"false" json:"isOutput"`

	// The name of the object type.
	TypeName *string `mandatory:"false" json:"typeName"`

	// The output aggregation type
	OutputAggregationType ParameterOutputAggregationTypeEnum `mandatory:"false" json:"outputAggregationType,omitempty"`
}

//GetKey returns Key
func (m Parameter) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m Parameter) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m Parameter) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetConfigValues returns ConfigValues
func (m Parameter) GetConfigValues() *ConfigValues {
	return m.ConfigValues
}

//GetObjectStatus returns ObjectStatus
func (m Parameter) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetName returns Name
func (m Parameter) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m Parameter) GetDescription() *string {
	return m.Description
}

func (m Parameter) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m Parameter) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeParameter Parameter
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeParameter
	}{
		"PARAMETER",
		(MarshalTypeParameter)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *Parameter) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Key                    *string                            `json:"key"`
		ModelVersion           *string                            `json:"modelVersion"`
		ParentRef              *ParentReference                   `json:"parentRef"`
		ConfigValues           *ConfigValues                      `json:"configValues"`
		ObjectStatus           *int                               `json:"objectStatus"`
		Name                   *string                            `json:"name"`
		Description            *string                            `json:"description"`
		Type                   basetype                           `json:"type"`
		DefaultValue           *interface{}                       `json:"defaultValue"`
		RootObjectDefaultValue *interface{}                       `json:"rootObjectDefaultValue"`
		IsInput                *bool                              `json:"isInput"`
		IsOutput               *bool                              `json:"isOutput"`
		OutputAggregationType  ParameterOutputAggregationTypeEnum `json:"outputAggregationType"`
		TypeName               *string                            `json:"typeName"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Key = model.Key

	m.ModelVersion = model.ModelVersion

	m.ParentRef = model.ParentRef

	m.ConfigValues = model.ConfigValues

	m.ObjectStatus = model.ObjectStatus

	m.Name = model.Name

	m.Description = model.Description

	nn, e = model.Type.UnmarshalPolymorphicJSON(model.Type.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Type = nn.(BaseType)
	} else {
		m.Type = nil
	}

	m.DefaultValue = model.DefaultValue

	m.RootObjectDefaultValue = model.RootObjectDefaultValue

	m.IsInput = model.IsInput

	m.IsOutput = model.IsOutput

	m.OutputAggregationType = model.OutputAggregationType

	m.TypeName = model.TypeName

	return
}

// ParameterOutputAggregationTypeEnum Enum with underlying type: string
type ParameterOutputAggregationTypeEnum string

// Set of constants representing the allowable values for ParameterOutputAggregationTypeEnum
const (
	ParameterOutputAggregationTypeMin   ParameterOutputAggregationTypeEnum = "MIN"
	ParameterOutputAggregationTypeMax   ParameterOutputAggregationTypeEnum = "MAX"
	ParameterOutputAggregationTypeCount ParameterOutputAggregationTypeEnum = "COUNT"
	ParameterOutputAggregationTypeSum   ParameterOutputAggregationTypeEnum = "SUM"
)

var mappingParameterOutputAggregationType = map[string]ParameterOutputAggregationTypeEnum{
	"MIN":   ParameterOutputAggregationTypeMin,
	"MAX":   ParameterOutputAggregationTypeMax,
	"COUNT": ParameterOutputAggregationTypeCount,
	"SUM":   ParameterOutputAggregationTypeSum,
}

// GetParameterOutputAggregationTypeEnumValues Enumerates the set of values for ParameterOutputAggregationTypeEnum
func GetParameterOutputAggregationTypeEnumValues() []ParameterOutputAggregationTypeEnum {
	values := make([]ParameterOutputAggregationTypeEnum, 0)
	for _, v := range mappingParameterOutputAggregationType {
		values = append(values, v)
	}
	return values
}
