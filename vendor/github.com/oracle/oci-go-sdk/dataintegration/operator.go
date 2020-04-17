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

// Operator An operator defines some data integration semantics in a data flow. It may be reading/writing data or transforming the data.
type Operator interface {

	// The key of the object.
	GetKey() *string

	// The model version of an object.
	GetModelVersion() *string

	GetParentRef() *ParentReference

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value can be edited by the user and it is restricted to 1000 characters
	GetName() *string

	// Detailed description for the object.
	GetDescription() *string

	// The version of the object that is used to track changes in the object instance.
	GetObjectVersion() *int

	// An array of input ports.
	GetInputPorts() []InputPort

	// An array of output ports.
	GetOutputPorts() []OutputPort

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	GetObjectStatus() *int

	// Value can only contain upper case letters, underscore and numbers. It should begin with upper case letter or underscore. The value can be edited by the user.
	GetIdentifier() *string

	// An array of parameters.
	GetParameters() []Parameter

	GetOpConfigValues() *ConfigValues
}

type operator struct {
	JsonData       []byte
	Key            *string          `mandatory:"false" json:"key"`
	ModelVersion   *string          `mandatory:"false" json:"modelVersion"`
	ParentRef      *ParentReference `mandatory:"false" json:"parentRef"`
	Name           *string          `mandatory:"false" json:"name"`
	Description    *string          `mandatory:"false" json:"description"`
	ObjectVersion  *int             `mandatory:"false" json:"objectVersion"`
	InputPorts     []InputPort      `mandatory:"false" json:"inputPorts"`
	OutputPorts    []OutputPort     `mandatory:"false" json:"outputPorts"`
	ObjectStatus   *int             `mandatory:"false" json:"objectStatus"`
	Identifier     *string          `mandatory:"false" json:"identifier"`
	Parameters     []Parameter      `mandatory:"false" json:"parameters"`
	OpConfigValues *ConfigValues    `mandatory:"false" json:"opConfigValues"`
	ModelType      string           `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *operator) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshaleroperator operator
	s := struct {
		Model Unmarshaleroperator
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Key = s.Model.Key
	m.ModelVersion = s.Model.ModelVersion
	m.ParentRef = s.Model.ParentRef
	m.Name = s.Model.Name
	m.Description = s.Model.Description
	m.ObjectVersion = s.Model.ObjectVersion
	m.InputPorts = s.Model.InputPorts
	m.OutputPorts = s.Model.OutputPorts
	m.ObjectStatus = s.Model.ObjectStatus
	m.Identifier = s.Model.Identifier
	m.Parameters = s.Model.Parameters
	m.OpConfigValues = s.Model.OpConfigValues
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *operator) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "TARGET_OPERATOR":
		mm := Target{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "JOINER_OPERATOR":
		mm := Joiner{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "FILTER_OPERATOR":
		mm := Filter{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AGGREGATOR_OPERATOR":
		mm := Aggregator{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PROJECTION_OPERATOR":
		mm := Projection{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SOURCE_OPERATOR":
		mm := Source{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetKey returns Key
func (m operator) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m operator) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m operator) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetName returns Name
func (m operator) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m operator) GetDescription() *string {
	return m.Description
}

//GetObjectVersion returns ObjectVersion
func (m operator) GetObjectVersion() *int {
	return m.ObjectVersion
}

//GetInputPorts returns InputPorts
func (m operator) GetInputPorts() []InputPort {
	return m.InputPorts
}

//GetOutputPorts returns OutputPorts
func (m operator) GetOutputPorts() []OutputPort {
	return m.OutputPorts
}

//GetObjectStatus returns ObjectStatus
func (m operator) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetIdentifier returns Identifier
func (m operator) GetIdentifier() *string {
	return m.Identifier
}

//GetParameters returns Parameters
func (m operator) GetParameters() []Parameter {
	return m.Parameters
}

//GetOpConfigValues returns OpConfigValues
func (m operator) GetOpConfigValues() *ConfigValues {
	return m.OpConfigValues
}

func (m operator) String() string {
	return common.PointerString(m)
}

// OperatorModelTypeEnum Enum with underlying type: string
type OperatorModelTypeEnum string

// Set of constants representing the allowable values for OperatorModelTypeEnum
const (
	OperatorModelTypeSourceOperator     OperatorModelTypeEnum = "SOURCE_OPERATOR"
	OperatorModelTypeFilterOperator     OperatorModelTypeEnum = "FILTER_OPERATOR"
	OperatorModelTypeJoinerOperator     OperatorModelTypeEnum = "JOINER_OPERATOR"
	OperatorModelTypeAggregatorOperator OperatorModelTypeEnum = "AGGREGATOR_OPERATOR"
	OperatorModelTypeProjectionOperator OperatorModelTypeEnum = "PROJECTION_OPERATOR"
	OperatorModelTypeTargetOperator     OperatorModelTypeEnum = "TARGET_OPERATOR"
)

var mappingOperatorModelType = map[string]OperatorModelTypeEnum{
	"SOURCE_OPERATOR":     OperatorModelTypeSourceOperator,
	"FILTER_OPERATOR":     OperatorModelTypeFilterOperator,
	"JOINER_OPERATOR":     OperatorModelTypeJoinerOperator,
	"AGGREGATOR_OPERATOR": OperatorModelTypeAggregatorOperator,
	"PROJECTION_OPERATOR": OperatorModelTypeProjectionOperator,
	"TARGET_OPERATOR":     OperatorModelTypeTargetOperator,
}

// GetOperatorModelTypeEnumValues Enumerates the set of values for OperatorModelTypeEnum
func GetOperatorModelTypeEnumValues() []OperatorModelTypeEnum {
	values := make([]OperatorModelTypeEnum, 0)
	for _, v := range mappingOperatorModelType {
		values = append(values, v)
	}
	return values
}
