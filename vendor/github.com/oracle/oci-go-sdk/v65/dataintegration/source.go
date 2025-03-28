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

// Source The information about the source object.
type Source struct {

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

	Entity DataEntity `mandatory:"false" json:"entity"`

	// Specifies the read access.
	IsReadAccess *bool `mandatory:"false" json:"isReadAccess"`

	// Specifies the copy fields.
	IsCopyFields *bool `mandatory:"false" json:"isCopyFields"`

	// Specifies if this uses a predefined shape.
	IsPredefinedShape *bool `mandatory:"false" json:"isPredefinedShape"`

	SchemaDriftConfig *SchemaDriftConfig `mandatory:"false" json:"schemaDriftConfig"`

	FixedDataShape *Shape `mandatory:"false" json:"fixedDataShape"`

	ReadOperationConfig *ReadOperationConfig `mandatory:"false" json:"readOperationConfig"`
}

// GetKey returns Key
func (m Source) GetKey() *string {
	return m.Key
}

// GetModelVersion returns ModelVersion
func (m Source) GetModelVersion() *string {
	return m.ModelVersion
}

// GetParentRef returns ParentRef
func (m Source) GetParentRef() *ParentReference {
	return m.ParentRef
}

// GetName returns Name
func (m Source) GetName() *string {
	return m.Name
}

// GetDescription returns Description
func (m Source) GetDescription() *string {
	return m.Description
}

// GetObjectVersion returns ObjectVersion
func (m Source) GetObjectVersion() *int {
	return m.ObjectVersion
}

// GetInputPorts returns InputPorts
func (m Source) GetInputPorts() []InputPort {
	return m.InputPorts
}

// GetOutputPorts returns OutputPorts
func (m Source) GetOutputPorts() []TypedObject {
	return m.OutputPorts
}

// GetObjectStatus returns ObjectStatus
func (m Source) GetObjectStatus() *int {
	return m.ObjectStatus
}

// GetIdentifier returns Identifier
func (m Source) GetIdentifier() *string {
	return m.Identifier
}

// GetParameters returns Parameters
func (m Source) GetParameters() []Parameter {
	return m.Parameters
}

// GetOpConfigValues returns OpConfigValues
func (m Source) GetOpConfigValues() *ConfigValues {
	return m.OpConfigValues
}

func (m Source) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Source) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m Source) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeSource Source
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeSource
	}{
		"SOURCE_OPERATOR",
		(MarshalTypeSource)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *Source) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Key                 *string              `json:"key"`
		ModelVersion        *string              `json:"modelVersion"`
		ParentRef           *ParentReference     `json:"parentRef"`
		Name                *string              `json:"name"`
		Description         *string              `json:"description"`
		ObjectVersion       *int                 `json:"objectVersion"`
		InputPorts          []InputPort          `json:"inputPorts"`
		OutputPorts         []typedobject        `json:"outputPorts"`
		ObjectStatus        *int                 `json:"objectStatus"`
		Identifier          *string              `json:"identifier"`
		Parameters          []Parameter          `json:"parameters"`
		OpConfigValues      *ConfigValues        `json:"opConfigValues"`
		Entity              dataentity           `json:"entity"`
		IsReadAccess        *bool                `json:"isReadAccess"`
		IsCopyFields        *bool                `json:"isCopyFields"`
		IsPredefinedShape   *bool                `json:"isPredefinedShape"`
		SchemaDriftConfig   *SchemaDriftConfig   `json:"schemaDriftConfig"`
		FixedDataShape      *Shape               `json:"fixedDataShape"`
		ReadOperationConfig *ReadOperationConfig `json:"readOperationConfig"`
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

	nn, e = model.Entity.UnmarshalPolymorphicJSON(model.Entity.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Entity = nn.(DataEntity)
	} else {
		m.Entity = nil
	}

	m.IsReadAccess = model.IsReadAccess

	m.IsCopyFields = model.IsCopyFields

	m.IsPredefinedShape = model.IsPredefinedShape

	m.SchemaDriftConfig = model.SchemaDriftConfig

	m.FixedDataShape = model.FixedDataShape

	m.ReadOperationConfig = model.ReadOperationConfig

	return
}
