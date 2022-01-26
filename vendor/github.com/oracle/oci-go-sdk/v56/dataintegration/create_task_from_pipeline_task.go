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
	"github.com/oracle/oci-go-sdk/v56/common"
)

// CreateTaskFromPipelineTask The information about the pipeline task.
type CreateTaskFromPipelineTask struct {

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"true" json:"name"`

	// Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
	Identifier *string `mandatory:"true" json:"identifier"`

	RegistryMetadata *RegistryMetadata `mandatory:"true" json:"registryMetadata"`

	// Generated key that can be used in API calls to identify task. On scenarios where reference to the task is needed, a value can be passed in create.
	Key *string `mandatory:"false" json:"key"`

	// The object's model version.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Detailed description for the object.
	Description *string `mandatory:"false" json:"description"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// An array of input ports.
	InputPorts []InputPort `mandatory:"false" json:"inputPorts"`

	// An array of output ports.
	OutputPorts []OutputPort `mandatory:"false" json:"outputPorts"`

	// An array of parameters.
	Parameters []Parameter `mandatory:"false" json:"parameters"`

	OpConfigValues *ConfigValues `mandatory:"false" json:"opConfigValues"`

	ConfigProviderDelegate *CreateConfigProvider `mandatory:"false" json:"configProviderDelegate"`

	Pipeline *Pipeline `mandatory:"false" json:"pipeline"`
}

//GetKey returns Key
func (m CreateTaskFromPipelineTask) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m CreateTaskFromPipelineTask) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m CreateTaskFromPipelineTask) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetName returns Name
func (m CreateTaskFromPipelineTask) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m CreateTaskFromPipelineTask) GetDescription() *string {
	return m.Description
}

//GetObjectStatus returns ObjectStatus
func (m CreateTaskFromPipelineTask) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetIdentifier returns Identifier
func (m CreateTaskFromPipelineTask) GetIdentifier() *string {
	return m.Identifier
}

//GetInputPorts returns InputPorts
func (m CreateTaskFromPipelineTask) GetInputPorts() []InputPort {
	return m.InputPorts
}

//GetOutputPorts returns OutputPorts
func (m CreateTaskFromPipelineTask) GetOutputPorts() []OutputPort {
	return m.OutputPorts
}

//GetParameters returns Parameters
func (m CreateTaskFromPipelineTask) GetParameters() []Parameter {
	return m.Parameters
}

//GetOpConfigValues returns OpConfigValues
func (m CreateTaskFromPipelineTask) GetOpConfigValues() *ConfigValues {
	return m.OpConfigValues
}

//GetConfigProviderDelegate returns ConfigProviderDelegate
func (m CreateTaskFromPipelineTask) GetConfigProviderDelegate() *CreateConfigProvider {
	return m.ConfigProviderDelegate
}

//GetRegistryMetadata returns RegistryMetadata
func (m CreateTaskFromPipelineTask) GetRegistryMetadata() *RegistryMetadata {
	return m.RegistryMetadata
}

func (m CreateTaskFromPipelineTask) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateTaskFromPipelineTask) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateTaskFromPipelineTask CreateTaskFromPipelineTask
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeCreateTaskFromPipelineTask
	}{
		"PIPELINE_TASK",
		(MarshalTypeCreateTaskFromPipelineTask)(m),
	}

	return json.Marshal(&s)
}
