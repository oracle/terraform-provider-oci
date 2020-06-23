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

// TaskFromDataLoaderTaskDetails The information about a data flow task.
type TaskFromDataLoaderTaskDetails struct {

	// Generated key that can be used in API calls to identify task. On scenarios where reference to the task is needed, a value can be passed in create.
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

	Metadata *ObjectMetadata `mandatory:"false" json:"metadata"`

	// A map, if provided key is replaced with generated key, this structure provides mapping between user provided key and generated key
	KeyMap map[string]string `mandatory:"false" json:"keyMap"`

	DataFlow *DataFlow `mandatory:"false" json:"dataFlow"`
}

//GetKey returns Key
func (m TaskFromDataLoaderTaskDetails) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m TaskFromDataLoaderTaskDetails) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m TaskFromDataLoaderTaskDetails) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetName returns Name
func (m TaskFromDataLoaderTaskDetails) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m TaskFromDataLoaderTaskDetails) GetDescription() *string {
	return m.Description
}

//GetObjectVersion returns ObjectVersion
func (m TaskFromDataLoaderTaskDetails) GetObjectVersion() *int {
	return m.ObjectVersion
}

//GetObjectStatus returns ObjectStatus
func (m TaskFromDataLoaderTaskDetails) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetIdentifier returns Identifier
func (m TaskFromDataLoaderTaskDetails) GetIdentifier() *string {
	return m.Identifier
}

//GetInputPorts returns InputPorts
func (m TaskFromDataLoaderTaskDetails) GetInputPorts() []InputPort {
	return m.InputPorts
}

//GetOutputPorts returns OutputPorts
func (m TaskFromDataLoaderTaskDetails) GetOutputPorts() []OutputPort {
	return m.OutputPorts
}

//GetParameters returns Parameters
func (m TaskFromDataLoaderTaskDetails) GetParameters() []Parameter {
	return m.Parameters
}

//GetOpConfigValues returns OpConfigValues
func (m TaskFromDataLoaderTaskDetails) GetOpConfigValues() *ConfigValues {
	return m.OpConfigValues
}

//GetConfigProviderDelegate returns ConfigProviderDelegate
func (m TaskFromDataLoaderTaskDetails) GetConfigProviderDelegate() *ConfigProvider {
	return m.ConfigProviderDelegate
}

//GetMetadata returns Metadata
func (m TaskFromDataLoaderTaskDetails) GetMetadata() *ObjectMetadata {
	return m.Metadata
}

//GetKeyMap returns KeyMap
func (m TaskFromDataLoaderTaskDetails) GetKeyMap() map[string]string {
	return m.KeyMap
}

func (m TaskFromDataLoaderTaskDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m TaskFromDataLoaderTaskDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeTaskFromDataLoaderTaskDetails TaskFromDataLoaderTaskDetails
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeTaskFromDataLoaderTaskDetails
	}{
		"DATA_LOADER_TASK",
		(MarshalTypeTaskFromDataLoaderTaskDetails)(m),
	}

	return json.Marshal(&s)
}
