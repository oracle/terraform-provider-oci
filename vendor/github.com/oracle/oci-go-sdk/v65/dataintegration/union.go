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

// Union The information about a union object.
type Union struct {

	// The key of the object.
	Key *string `mandatory:"false" json:"key"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"false" json:"name"`

	// Details about the operator.
	Description *string `mandatory:"false" json:"description"`

	// The version of the object that is used to track changes in the object instance.
	ObjectVersion *int `mandatory:"false" json:"objectVersion"`

	// An array of input ports.
	InputPorts []InputPort `mandatory:"false" json:"inputPorts"`

	// An array of output ports.
	OutputPorts []TypedObject `mandatory:"false" json:"outputPorts"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
	Identifier *string `mandatory:"false" json:"identifier"`

	// An array of parameters used in the data flow.
	Parameters []Parameter `mandatory:"false" json:"parameters"`

	OpConfigValues *ConfigValues `mandatory:"false" json:"opConfigValues"`

	// The information about the union all.
	IsAll *bool `mandatory:"false" json:"isAll"`

	// unionType
	UnionType UnionUnionTypeEnum `mandatory:"false" json:"unionType,omitempty"`
}

// GetKey returns Key
func (m Union) GetKey() *string {
	return m.Key
}

// GetModelVersion returns ModelVersion
func (m Union) GetModelVersion() *string {
	return m.ModelVersion
}

// GetParentRef returns ParentRef
func (m Union) GetParentRef() *ParentReference {
	return m.ParentRef
}

// GetName returns Name
func (m Union) GetName() *string {
	return m.Name
}

// GetDescription returns Description
func (m Union) GetDescription() *string {
	return m.Description
}

// GetObjectVersion returns ObjectVersion
func (m Union) GetObjectVersion() *int {
	return m.ObjectVersion
}

// GetInputPorts returns InputPorts
func (m Union) GetInputPorts() []InputPort {
	return m.InputPorts
}

// GetOutputPorts returns OutputPorts
func (m Union) GetOutputPorts() []TypedObject {
	return m.OutputPorts
}

// GetObjectStatus returns ObjectStatus
func (m Union) GetObjectStatus() *int {
	return m.ObjectStatus
}

// GetIdentifier returns Identifier
func (m Union) GetIdentifier() *string {
	return m.Identifier
}

// GetParameters returns Parameters
func (m Union) GetParameters() []Parameter {
	return m.Parameters
}

// GetOpConfigValues returns OpConfigValues
func (m Union) GetOpConfigValues() *ConfigValues {
	return m.OpConfigValues
}

func (m Union) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Union) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingUnionUnionTypeEnum(string(m.UnionType)); !ok && m.UnionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UnionType: %s. Supported values are: %s.", m.UnionType, strings.Join(GetUnionUnionTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m Union) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUnion Union
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeUnion
	}{
		"UNION_OPERATOR",
		(MarshalTypeUnion)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *Union) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Key            *string            `json:"key"`
		ModelVersion   *string            `json:"modelVersion"`
		ParentRef      *ParentReference   `json:"parentRef"`
		Name           *string            `json:"name"`
		Description    *string            `json:"description"`
		ObjectVersion  *int               `json:"objectVersion"`
		InputPorts     []InputPort        `json:"inputPorts"`
		OutputPorts    []typedobject      `json:"outputPorts"`
		ObjectStatus   *int               `json:"objectStatus"`
		Identifier     *string            `json:"identifier"`
		Parameters     []Parameter        `json:"parameters"`
		OpConfigValues *ConfigValues      `json:"opConfigValues"`
		UnionType      UnionUnionTypeEnum `json:"unionType"`
		IsAll          *bool              `json:"isAll"`
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

	m.Description = model.Description

	m.ObjectVersion = model.ObjectVersion

	m.InputPorts = make([]InputPort, len(model.InputPorts))
	copy(m.InputPorts, model.InputPorts)
	m.OutputPorts = make([]TypedObject, len(model.OutputPorts))
	for i, n := range model.OutputPorts {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.OutputPorts[i] = nn.(TypedObject)
		} else {
			m.OutputPorts[i] = nil
		}
	}
	m.ObjectStatus = model.ObjectStatus

	m.Identifier = model.Identifier

	m.Parameters = make([]Parameter, len(model.Parameters))
	copy(m.Parameters, model.Parameters)
	m.OpConfigValues = model.OpConfigValues

	m.UnionType = model.UnionType

	m.IsAll = model.IsAll

	return
}

// UnionUnionTypeEnum Enum with underlying type: string
type UnionUnionTypeEnum string

// Set of constants representing the allowable values for UnionUnionTypeEnum
const (
	UnionUnionTypeName     UnionUnionTypeEnum = "NAME"
	UnionUnionTypePosition UnionUnionTypeEnum = "POSITION"
)

var mappingUnionUnionTypeEnum = map[string]UnionUnionTypeEnum{
	"NAME":     UnionUnionTypeName,
	"POSITION": UnionUnionTypePosition,
}

var mappingUnionUnionTypeEnumLowerCase = map[string]UnionUnionTypeEnum{
	"name":     UnionUnionTypeName,
	"position": UnionUnionTypePosition,
}

// GetUnionUnionTypeEnumValues Enumerates the set of values for UnionUnionTypeEnum
func GetUnionUnionTypeEnumValues() []UnionUnionTypeEnum {
	values := make([]UnionUnionTypeEnum, 0)
	for _, v := range mappingUnionUnionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUnionUnionTypeEnumStringValues Enumerates the set of values in String for UnionUnionTypeEnum
func GetUnionUnionTypeEnumStringValues() []string {
	return []string{
		"NAME",
		"POSITION",
	}
}

// GetMappingUnionUnionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUnionUnionTypeEnum(val string) (UnionUnionTypeEnum, bool) {
	enum, ok := mappingUnionUnionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
