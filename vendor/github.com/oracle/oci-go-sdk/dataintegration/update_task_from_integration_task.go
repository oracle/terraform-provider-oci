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

// UpdateTaskFromIntegrationTask The information about the integration task.
type UpdateTaskFromIntegrationTask struct {

	// Generated key that can be used in API calls to identify task. On scenarios where reference to the task is needed, a value can be passed in create.
	Key *string `mandatory:"true" json:"key"`

	// The version of the object that is used to track changes in the object instance.
	ObjectVersion *int `mandatory:"true" json:"objectVersion"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value can be edited by the user and it is restricted to 1000 characters
	Name *string `mandatory:"false" json:"name"`

	// Detailed description for the object.
	Description *string `mandatory:"false" json:"description"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Value can only contain upper case letters, underscore and numbers. It should begin with upper case letter or underscore. The value can be edited by the user.
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

	DataFlow *DataFlow `mandatory:"false" json:"dataFlow"`
}

//GetKey returns Key
func (m UpdateTaskFromIntegrationTask) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m UpdateTaskFromIntegrationTask) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m UpdateTaskFromIntegrationTask) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetName returns Name
func (m UpdateTaskFromIntegrationTask) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m UpdateTaskFromIntegrationTask) GetDescription() *string {
	return m.Description
}

//GetObjectStatus returns ObjectStatus
func (m UpdateTaskFromIntegrationTask) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetObjectVersion returns ObjectVersion
func (m UpdateTaskFromIntegrationTask) GetObjectVersion() *int {
	return m.ObjectVersion
}

//GetIdentifier returns Identifier
func (m UpdateTaskFromIntegrationTask) GetIdentifier() *string {
	return m.Identifier
}

//GetInputPorts returns InputPorts
func (m UpdateTaskFromIntegrationTask) GetInputPorts() []InputPort {
	return m.InputPorts
}

//GetOutputPorts returns OutputPorts
func (m UpdateTaskFromIntegrationTask) GetOutputPorts() []OutputPort {
	return m.OutputPorts
}

//GetParameters returns Parameters
func (m UpdateTaskFromIntegrationTask) GetParameters() []Parameter {
	return m.Parameters
}

//GetOpConfigValues returns OpConfigValues
func (m UpdateTaskFromIntegrationTask) GetOpConfigValues() *ConfigValues {
	return m.OpConfigValues
}

//GetConfigProviderDelegate returns ConfigProviderDelegate
func (m UpdateTaskFromIntegrationTask) GetConfigProviderDelegate() *ConfigProvider {
	return m.ConfigProviderDelegate
}

//GetRegistryMetadata returns RegistryMetadata
func (m UpdateTaskFromIntegrationTask) GetRegistryMetadata() *RegistryMetadata {
	return m.RegistryMetadata
}

func (m UpdateTaskFromIntegrationTask) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m UpdateTaskFromIntegrationTask) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateTaskFromIntegrationTask UpdateTaskFromIntegrationTask
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeUpdateTaskFromIntegrationTask
	}{
		"INTEGRATION_TASK",
		(MarshalTypeUpdateTaskFromIntegrationTask)(m),
	}

	return json.Marshal(&s)
}
