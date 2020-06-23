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

// Filter The information about the filter object.
type Filter struct {

	// The key of the object.
	Key *string `mandatory:"false" json:"key"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value can be edited by the user and it is restricted to 1000 characters
	Name *string `mandatory:"false" json:"name"`

	// Detailed description for the object.
	Description *string `mandatory:"false" json:"description"`

	// The version of the object that is used to track changes in the object instance.
	ObjectVersion *int `mandatory:"false" json:"objectVersion"`

	// An array of input ports.
	InputPorts []InputPort `mandatory:"false" json:"inputPorts"`

	// An array of output ports.
	OutputPorts []OutputPort `mandatory:"false" json:"outputPorts"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Value can only contain upper case letters, underscore and numbers. It should begin with upper case letter or underscore. The value can be edited by the user.
	Identifier *string `mandatory:"false" json:"identifier"`

	// An array of parameters.
	Parameters []Parameter `mandatory:"false" json:"parameters"`

	OpConfigValues *ConfigValues `mandatory:"false" json:"opConfigValues"`

	FilterCondition *Expression `mandatory:"false" json:"filterCondition"`
}

//GetKey returns Key
func (m Filter) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m Filter) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m Filter) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetName returns Name
func (m Filter) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m Filter) GetDescription() *string {
	return m.Description
}

//GetObjectVersion returns ObjectVersion
func (m Filter) GetObjectVersion() *int {
	return m.ObjectVersion
}

//GetInputPorts returns InputPorts
func (m Filter) GetInputPorts() []InputPort {
	return m.InputPorts
}

//GetOutputPorts returns OutputPorts
func (m Filter) GetOutputPorts() []OutputPort {
	return m.OutputPorts
}

//GetObjectStatus returns ObjectStatus
func (m Filter) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetIdentifier returns Identifier
func (m Filter) GetIdentifier() *string {
	return m.Identifier
}

//GetParameters returns Parameters
func (m Filter) GetParameters() []Parameter {
	return m.Parameters
}

//GetOpConfigValues returns OpConfigValues
func (m Filter) GetOpConfigValues() *ConfigValues {
	return m.OpConfigValues
}

func (m Filter) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m Filter) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeFilter Filter
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeFilter
	}{
		"FILTER_OPERATOR",
		(MarshalTypeFilter)(m),
	}

	return json.Marshal(&s)
}
