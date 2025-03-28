// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// OutputPort The output port details.
type OutputPort struct {

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
	PortType OutputPortPortTypeEnum `mandatory:"false" json:"portType,omitempty"`
}

// GetKey returns Key
func (m OutputPort) GetKey() *string {
	return m.Key
}

// GetModelVersion returns ModelVersion
func (m OutputPort) GetModelVersion() *string {
	return m.ModelVersion
}

// GetParentRef returns ParentRef
func (m OutputPort) GetParentRef() *ParentReference {
	return m.ParentRef
}

// GetConfigValues returns ConfigValues
func (m OutputPort) GetConfigValues() *ConfigValues {
	return m.ConfigValues
}

// GetObjectStatus returns ObjectStatus
func (m OutputPort) GetObjectStatus() *int {
	return m.ObjectStatus
}

// GetName returns Name
func (m OutputPort) GetName() *string {
	return m.Name
}

// GetDescription returns Description
func (m OutputPort) GetDescription() *string {
	return m.Description
}

func (m OutputPort) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OutputPort) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOutputPortPortTypeEnum(string(m.PortType)); !ok && m.PortType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PortType: %s. Supported values are: %s.", m.PortType, strings.Join(GetOutputPortPortTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m OutputPort) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOutputPort OutputPort
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeOutputPort
	}{
		"OUTPUT_PORT",
		(MarshalTypeOutputPort)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *OutputPort) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Key          *string                `json:"key"`
		ModelVersion *string                `json:"modelVersion"`
		ParentRef    *ParentReference       `json:"parentRef"`
		ConfigValues *ConfigValues          `json:"configValues"`
		ObjectStatus *int                   `json:"objectStatus"`
		Name         *string                `json:"name"`
		Description  *string                `json:"description"`
		PortType     OutputPortPortTypeEnum `json:"portType"`
		Fields       []typedobject          `json:"fields"`
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

	m.PortType = model.PortType

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
	return
}

// OutputPortPortTypeEnum Enum with underlying type: string
type OutputPortPortTypeEnum string

// Set of constants representing the allowable values for OutputPortPortTypeEnum
const (
	OutputPortPortTypeData    OutputPortPortTypeEnum = "DATA"
	OutputPortPortTypeControl OutputPortPortTypeEnum = "CONTROL"
	OutputPortPortTypeModel   OutputPortPortTypeEnum = "MODEL"
)

var mappingOutputPortPortTypeEnum = map[string]OutputPortPortTypeEnum{
	"DATA":    OutputPortPortTypeData,
	"CONTROL": OutputPortPortTypeControl,
	"MODEL":   OutputPortPortTypeModel,
}

var mappingOutputPortPortTypeEnumLowerCase = map[string]OutputPortPortTypeEnum{
	"data":    OutputPortPortTypeData,
	"control": OutputPortPortTypeControl,
	"model":   OutputPortPortTypeModel,
}

// GetOutputPortPortTypeEnumValues Enumerates the set of values for OutputPortPortTypeEnum
func GetOutputPortPortTypeEnumValues() []OutputPortPortTypeEnum {
	values := make([]OutputPortPortTypeEnum, 0)
	for _, v := range mappingOutputPortPortTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOutputPortPortTypeEnumStringValues Enumerates the set of values in String for OutputPortPortTypeEnum
func GetOutputPortPortTypeEnumStringValues() []string {
	return []string{
		"DATA",
		"CONTROL",
		"MODEL",
	}
}

// GetMappingOutputPortPortTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOutputPortPortTypeEnum(val string) (OutputPortPortTypeEnum, bool) {
	enum, ok := mappingOutputPortPortTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
