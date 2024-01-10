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

// DecisionOutputPort The conditional output port details, used in operators such as decision operator.
type DecisionOutputPort struct {

	// The key of the object.
	Key *string `mandatory:"false" json:"key"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	ConfigValues *ConfigValues `mandatory:"false" json:"configValues"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"false" json:"name"`

	// Detailed description for the object.
	Description *string `mandatory:"false" json:"description"`

	// An array of fields.
	Fields []TypedObject `mandatory:"false" json:"fields"`

	// The port details for the data asset.Type.
	PortType DecisionOutputPortPortTypeEnum `mandatory:"true" json:"portType"`

	// The port based on what decision expression evaluates to.
	DecisionOutputPortType DecisionOutputPortDecisionOutputPortTypeEnum `mandatory:"true" json:"decisionOutputPortType"`
}

// GetKey returns Key
func (m DecisionOutputPort) GetKey() *string {
	return m.Key
}

// GetModelVersion returns ModelVersion
func (m DecisionOutputPort) GetModelVersion() *string {
	return m.ModelVersion
}

// GetParentRef returns ParentRef
func (m DecisionOutputPort) GetParentRef() *ParentReference {
	return m.ParentRef
}

// GetConfigValues returns ConfigValues
func (m DecisionOutputPort) GetConfigValues() *ConfigValues {
	return m.ConfigValues
}

// GetObjectStatus returns ObjectStatus
func (m DecisionOutputPort) GetObjectStatus() *int {
	return m.ObjectStatus
}

// GetName returns Name
func (m DecisionOutputPort) GetName() *string {
	return m.Name
}

// GetDescription returns Description
func (m DecisionOutputPort) GetDescription() *string {
	return m.Description
}

func (m DecisionOutputPort) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DecisionOutputPort) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDecisionOutputPortPortTypeEnum(string(m.PortType)); !ok && m.PortType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PortType: %s. Supported values are: %s.", m.PortType, strings.Join(GetDecisionOutputPortPortTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDecisionOutputPortDecisionOutputPortTypeEnum(string(m.DecisionOutputPortType)); !ok && m.DecisionOutputPortType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DecisionOutputPortType: %s. Supported values are: %s.", m.DecisionOutputPortType, strings.Join(GetDecisionOutputPortDecisionOutputPortTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DecisionOutputPort) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDecisionOutputPort DecisionOutputPort
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeDecisionOutputPort
	}{
		"DECISION_OUTPUT_PORT",
		(MarshalTypeDecisionOutputPort)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *DecisionOutputPort) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Key                    *string                                      `json:"key"`
		ModelVersion           *string                                      `json:"modelVersion"`
		ParentRef              *ParentReference                             `json:"parentRef"`
		ConfigValues           *ConfigValues                                `json:"configValues"`
		ObjectStatus           *int                                         `json:"objectStatus"`
		Name                   *string                                      `json:"name"`
		Description            *string                                      `json:"description"`
		Fields                 []typedobject                                `json:"fields"`
		PortType               DecisionOutputPortPortTypeEnum               `json:"portType"`
		DecisionOutputPortType DecisionOutputPortDecisionOutputPortTypeEnum `json:"decisionOutputPortType"`
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

	m.Fields = make([]TypedObject, len(model.Fields))
	for i, n := range model.Fields {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Fields[i] = nn.(TypedObject)
		} else {
			m.Fields[i] = nil
		}
	}
	m.PortType = model.PortType

	m.DecisionOutputPortType = model.DecisionOutputPortType

	return
}

// DecisionOutputPortPortTypeEnum Enum with underlying type: string
type DecisionOutputPortPortTypeEnum string

// Set of constants representing the allowable values for DecisionOutputPortPortTypeEnum
const (
	DecisionOutputPortPortTypeData    DecisionOutputPortPortTypeEnum = "DATA"
	DecisionOutputPortPortTypeControl DecisionOutputPortPortTypeEnum = "CONTROL"
	DecisionOutputPortPortTypeModel   DecisionOutputPortPortTypeEnum = "MODEL"
)

var mappingDecisionOutputPortPortTypeEnum = map[string]DecisionOutputPortPortTypeEnum{
	"DATA":    DecisionOutputPortPortTypeData,
	"CONTROL": DecisionOutputPortPortTypeControl,
	"MODEL":   DecisionOutputPortPortTypeModel,
}

var mappingDecisionOutputPortPortTypeEnumLowerCase = map[string]DecisionOutputPortPortTypeEnum{
	"data":    DecisionOutputPortPortTypeData,
	"control": DecisionOutputPortPortTypeControl,
	"model":   DecisionOutputPortPortTypeModel,
}

// GetDecisionOutputPortPortTypeEnumValues Enumerates the set of values for DecisionOutputPortPortTypeEnum
func GetDecisionOutputPortPortTypeEnumValues() []DecisionOutputPortPortTypeEnum {
	values := make([]DecisionOutputPortPortTypeEnum, 0)
	for _, v := range mappingDecisionOutputPortPortTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDecisionOutputPortPortTypeEnumStringValues Enumerates the set of values in String for DecisionOutputPortPortTypeEnum
func GetDecisionOutputPortPortTypeEnumStringValues() []string {
	return []string{
		"DATA",
		"CONTROL",
		"MODEL",
	}
}

// GetMappingDecisionOutputPortPortTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDecisionOutputPortPortTypeEnum(val string) (DecisionOutputPortPortTypeEnum, bool) {
	enum, ok := mappingDecisionOutputPortPortTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DecisionOutputPortDecisionOutputPortTypeEnum Enum with underlying type: string
type DecisionOutputPortDecisionOutputPortTypeEnum string

// Set of constants representing the allowable values for DecisionOutputPortDecisionOutputPortTypeEnum
const (
	DecisionOutputPortDecisionOutputPortTypeError DecisionOutputPortDecisionOutputPortTypeEnum = "EVAL_ERROR"
	DecisionOutputPortDecisionOutputPortTypeTrue  DecisionOutputPortDecisionOutputPortTypeEnum = "EVAL_TRUE"
	DecisionOutputPortDecisionOutputPortTypeFalse DecisionOutputPortDecisionOutputPortTypeEnum = "EVAL_FALSE"
)

var mappingDecisionOutputPortDecisionOutputPortTypeEnum = map[string]DecisionOutputPortDecisionOutputPortTypeEnum{
	"EVAL_ERROR": DecisionOutputPortDecisionOutputPortTypeError,
	"EVAL_TRUE":  DecisionOutputPortDecisionOutputPortTypeTrue,
	"EVAL_FALSE": DecisionOutputPortDecisionOutputPortTypeFalse,
}

var mappingDecisionOutputPortDecisionOutputPortTypeEnumLowerCase = map[string]DecisionOutputPortDecisionOutputPortTypeEnum{
	"eval_error": DecisionOutputPortDecisionOutputPortTypeError,
	"eval_true":  DecisionOutputPortDecisionOutputPortTypeTrue,
	"eval_false": DecisionOutputPortDecisionOutputPortTypeFalse,
}

// GetDecisionOutputPortDecisionOutputPortTypeEnumValues Enumerates the set of values for DecisionOutputPortDecisionOutputPortTypeEnum
func GetDecisionOutputPortDecisionOutputPortTypeEnumValues() []DecisionOutputPortDecisionOutputPortTypeEnum {
	values := make([]DecisionOutputPortDecisionOutputPortTypeEnum, 0)
	for _, v := range mappingDecisionOutputPortDecisionOutputPortTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDecisionOutputPortDecisionOutputPortTypeEnumStringValues Enumerates the set of values in String for DecisionOutputPortDecisionOutputPortTypeEnum
func GetDecisionOutputPortDecisionOutputPortTypeEnumStringValues() []string {
	return []string{
		"EVAL_ERROR",
		"EVAL_TRUE",
		"EVAL_FALSE",
	}
}

// GetMappingDecisionOutputPortDecisionOutputPortTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDecisionOutputPortDecisionOutputPortTypeEnum(val string) (DecisionOutputPortDecisionOutputPortTypeEnum, bool) {
	enum, ok := mappingDecisionOutputPortDecisionOutputPortTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
