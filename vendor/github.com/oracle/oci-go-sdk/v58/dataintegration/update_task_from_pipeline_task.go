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
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// UpdateTaskFromPipelineTask The information about the pipeline task.
type UpdateTaskFromPipelineTask struct {

	// Generated key that can be used in API calls to identify task. On scenarios where reference to the task is needed, a value can be passed in create.
	Key *string `mandatory:"true" json:"key"`

	// The version of the object that is used to track changes in the object instance.
	ObjectVersion *int `mandatory:"true" json:"objectVersion"`

	// The object's model version.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"false" json:"name"`

	// Detailed description for the object.
	Description *string `mandatory:"false" json:"description"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Value can only contain upper case letters, underscore and numbers. It should begin with upper case letter or underscore. The value can be modified.
	Identifier *string `mandatory:"false" json:"identifier"`

	// An array of input ports.
	InputPorts []InputPort `mandatory:"false" json:"inputPorts"`

	// An array of output ports.
	OutputPorts []OutputPort `mandatory:"false" json:"outputPorts"`

	// An array of parameters.
	Parameters []Parameter `mandatory:"false" json:"parameters"`

	OpConfigValues *ConfigValues `mandatory:"false" json:"opConfigValues"`

	ConfigProviderDelegate *ConfigProvider `mandatory:"false" json:"configProviderDelegate"`

	RegistryMetadata *RegistryMetadata `mandatory:"false" json:"registryMetadata"`

	Pipeline *Pipeline `mandatory:"false" json:"pipeline"`
}

//GetKey returns Key
func (m UpdateTaskFromPipelineTask) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m UpdateTaskFromPipelineTask) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m UpdateTaskFromPipelineTask) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetName returns Name
func (m UpdateTaskFromPipelineTask) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m UpdateTaskFromPipelineTask) GetDescription() *string {
	return m.Description
}

//GetObjectStatus returns ObjectStatus
func (m UpdateTaskFromPipelineTask) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetObjectVersion returns ObjectVersion
func (m UpdateTaskFromPipelineTask) GetObjectVersion() *int {
	return m.ObjectVersion
}

//GetIdentifier returns Identifier
func (m UpdateTaskFromPipelineTask) GetIdentifier() *string {
	return m.Identifier
}

//GetInputPorts returns InputPorts
func (m UpdateTaskFromPipelineTask) GetInputPorts() []InputPort {
	return m.InputPorts
}

//GetOutputPorts returns OutputPorts
func (m UpdateTaskFromPipelineTask) GetOutputPorts() []OutputPort {
	return m.OutputPorts
}

//GetParameters returns Parameters
func (m UpdateTaskFromPipelineTask) GetParameters() []Parameter {
	return m.Parameters
}

//GetOpConfigValues returns OpConfigValues
func (m UpdateTaskFromPipelineTask) GetOpConfigValues() *ConfigValues {
	return m.OpConfigValues
}

//GetConfigProviderDelegate returns ConfigProviderDelegate
func (m UpdateTaskFromPipelineTask) GetConfigProviderDelegate() *ConfigProvider {
	return m.ConfigProviderDelegate
}

//GetRegistryMetadata returns RegistryMetadata
func (m UpdateTaskFromPipelineTask) GetRegistryMetadata() *RegistryMetadata {
	return m.RegistryMetadata
}

func (m UpdateTaskFromPipelineTask) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateTaskFromPipelineTask) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateTaskFromPipelineTask) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateTaskFromPipelineTask UpdateTaskFromPipelineTask
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeUpdateTaskFromPipelineTask
	}{
		"PIPELINE_TASK",
		(MarshalTypeUpdateTaskFromPipelineTask)(m),
	}

	return json.Marshal(&s)
}
