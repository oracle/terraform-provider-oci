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

// Pivot Pivot operator has one input and one output. Pivot operator takes group by columns, a pivot key with values and aggregations. Output is the pivoted table.
type Pivot struct {

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

	GroupByColumns *DynamicProxyField `mandatory:"false" json:"groupByColumns"`

	MaterializedGroupByColumns *MaterializedDynamicField `mandatory:"false" json:"materializedGroupByColumns"`

	PivotKeys *PivotKeys `mandatory:"false" json:"pivotKeys"`
}

// GetKey returns Key
func (m Pivot) GetKey() *string {
	return m.Key
}

// GetModelVersion returns ModelVersion
func (m Pivot) GetModelVersion() *string {
	return m.ModelVersion
}

// GetParentRef returns ParentRef
func (m Pivot) GetParentRef() *ParentReference {
	return m.ParentRef
}

// GetName returns Name
func (m Pivot) GetName() *string {
	return m.Name
}

// GetDescription returns Description
func (m Pivot) GetDescription() *string {
	return m.Description
}

// GetObjectVersion returns ObjectVersion
func (m Pivot) GetObjectVersion() *int {
	return m.ObjectVersion
}

// GetInputPorts returns InputPorts
func (m Pivot) GetInputPorts() []InputPort {
	return m.InputPorts
}

// GetOutputPorts returns OutputPorts
func (m Pivot) GetOutputPorts() []TypedObject {
	return m.OutputPorts
}

// GetObjectStatus returns ObjectStatus
func (m Pivot) GetObjectStatus() *int {
	return m.ObjectStatus
}

// GetIdentifier returns Identifier
func (m Pivot) GetIdentifier() *string {
	return m.Identifier
}

// GetParameters returns Parameters
func (m Pivot) GetParameters() []Parameter {
	return m.Parameters
}

// GetOpConfigValues returns OpConfigValues
func (m Pivot) GetOpConfigValues() *ConfigValues {
	return m.OpConfigValues
}

func (m Pivot) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Pivot) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m Pivot) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePivot Pivot
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypePivot
	}{
		"PIVOT_OPERATOR",
		(MarshalTypePivot)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *Pivot) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Key                        *string                   `json:"key"`
		ModelVersion               *string                   `json:"modelVersion"`
		ParentRef                  *ParentReference          `json:"parentRef"`
		Name                       *string                   `json:"name"`
		Description                *string                   `json:"description"`
		ObjectVersion              *int                      `json:"objectVersion"`
		InputPorts                 []InputPort               `json:"inputPorts"`
		OutputPorts                []typedobject             `json:"outputPorts"`
		ObjectStatus               *int                      `json:"objectStatus"`
		Identifier                 *string                   `json:"identifier"`
		Parameters                 []Parameter               `json:"parameters"`
		OpConfigValues             *ConfigValues             `json:"opConfigValues"`
		GroupByColumns             *DynamicProxyField        `json:"groupByColumns"`
		MaterializedGroupByColumns *MaterializedDynamicField `json:"materializedGroupByColumns"`
		PivotKeys                  *PivotKeys                `json:"pivotKeys"`
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

	m.GroupByColumns = model.GroupByColumns

	m.MaterializedGroupByColumns = model.MaterializedGroupByColumns

	m.PivotKeys = model.PivotKeys

	return
}
