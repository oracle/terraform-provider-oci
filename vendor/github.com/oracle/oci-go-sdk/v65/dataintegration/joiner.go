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

// Joiner The information about a joiner object.
type Joiner struct {

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

	JoinCondition *Expression `mandatory:"false" json:"joinCondition"`

	// joinType
	JoinType JoinerJoinTypeEnum `mandatory:"false" json:"joinType,omitempty"`
}

// GetKey returns Key
func (m Joiner) GetKey() *string {
	return m.Key
}

// GetModelVersion returns ModelVersion
func (m Joiner) GetModelVersion() *string {
	return m.ModelVersion
}

// GetParentRef returns ParentRef
func (m Joiner) GetParentRef() *ParentReference {
	return m.ParentRef
}

// GetName returns Name
func (m Joiner) GetName() *string {
	return m.Name
}

// GetDescription returns Description
func (m Joiner) GetDescription() *string {
	return m.Description
}

// GetObjectVersion returns ObjectVersion
func (m Joiner) GetObjectVersion() *int {
	return m.ObjectVersion
}

// GetInputPorts returns InputPorts
func (m Joiner) GetInputPorts() []InputPort {
	return m.InputPorts
}

// GetOutputPorts returns OutputPorts
func (m Joiner) GetOutputPorts() []TypedObject {
	return m.OutputPorts
}

// GetObjectStatus returns ObjectStatus
func (m Joiner) GetObjectStatus() *int {
	return m.ObjectStatus
}

// GetIdentifier returns Identifier
func (m Joiner) GetIdentifier() *string {
	return m.Identifier
}

// GetParameters returns Parameters
func (m Joiner) GetParameters() []Parameter {
	return m.Parameters
}

// GetOpConfigValues returns OpConfigValues
func (m Joiner) GetOpConfigValues() *ConfigValues {
	return m.OpConfigValues
}

func (m Joiner) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Joiner) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingJoinerJoinTypeEnum(string(m.JoinType)); !ok && m.JoinType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for JoinType: %s. Supported values are: %s.", m.JoinType, strings.Join(GetJoinerJoinTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m Joiner) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeJoiner Joiner
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeJoiner
	}{
		"JOINER_OPERATOR",
		(MarshalTypeJoiner)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *Joiner) UnmarshalJSON(data []byte) (e error) {
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
		JoinType       JoinerJoinTypeEnum `json:"joinType"`
		JoinCondition  *Expression        `json:"joinCondition"`
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

	m.JoinType = model.JoinType

	m.JoinCondition = model.JoinCondition

	return
}

// JoinerJoinTypeEnum Enum with underlying type: string
type JoinerJoinTypeEnum string

// Set of constants representing the allowable values for JoinerJoinTypeEnum
const (
	JoinerJoinTypeInner JoinerJoinTypeEnum = "INNER"
	JoinerJoinTypeFull  JoinerJoinTypeEnum = "FULL"
	JoinerJoinTypeLeft  JoinerJoinTypeEnum = "LEFT"
	JoinerJoinTypeRight JoinerJoinTypeEnum = "RIGHT"
)

var mappingJoinerJoinTypeEnum = map[string]JoinerJoinTypeEnum{
	"INNER": JoinerJoinTypeInner,
	"FULL":  JoinerJoinTypeFull,
	"LEFT":  JoinerJoinTypeLeft,
	"RIGHT": JoinerJoinTypeRight,
}

var mappingJoinerJoinTypeEnumLowerCase = map[string]JoinerJoinTypeEnum{
	"inner": JoinerJoinTypeInner,
	"full":  JoinerJoinTypeFull,
	"left":  JoinerJoinTypeLeft,
	"right": JoinerJoinTypeRight,
}

// GetJoinerJoinTypeEnumValues Enumerates the set of values for JoinerJoinTypeEnum
func GetJoinerJoinTypeEnumValues() []JoinerJoinTypeEnum {
	values := make([]JoinerJoinTypeEnum, 0)
	for _, v := range mappingJoinerJoinTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetJoinerJoinTypeEnumStringValues Enumerates the set of values in String for JoinerJoinTypeEnum
func GetJoinerJoinTypeEnumStringValues() []string {
	return []string{
		"INNER",
		"FULL",
		"LEFT",
		"RIGHT",
	}
}

// GetMappingJoinerJoinTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingJoinerJoinTypeEnum(val string) (JoinerJoinTypeEnum, bool) {
	enum, ok := mappingJoinerJoinTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
