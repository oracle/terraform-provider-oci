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

// CreateTaskFromSqlTask The information about the SQL task.
type CreateTaskFromSqlTask struct {

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

	Script *Script `mandatory:"false" json:"script"`

	// Describes the shape of the execution result
	Operation *interface{} `mandatory:"false" json:"operation"`

	// Indicates whether the task is invoking a custom SQL script or stored procedure.
	SqlScriptType CreateTaskFromSqlTaskSqlScriptTypeEnum `mandatory:"false" json:"sqlScriptType,omitempty"`
}

//GetKey returns Key
func (m CreateTaskFromSqlTask) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m CreateTaskFromSqlTask) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m CreateTaskFromSqlTask) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetName returns Name
func (m CreateTaskFromSqlTask) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m CreateTaskFromSqlTask) GetDescription() *string {
	return m.Description
}

//GetObjectStatus returns ObjectStatus
func (m CreateTaskFromSqlTask) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetIdentifier returns Identifier
func (m CreateTaskFromSqlTask) GetIdentifier() *string {
	return m.Identifier
}

//GetInputPorts returns InputPorts
func (m CreateTaskFromSqlTask) GetInputPorts() []InputPort {
	return m.InputPorts
}

//GetOutputPorts returns OutputPorts
func (m CreateTaskFromSqlTask) GetOutputPorts() []OutputPort {
	return m.OutputPorts
}

//GetParameters returns Parameters
func (m CreateTaskFromSqlTask) GetParameters() []Parameter {
	return m.Parameters
}

//GetOpConfigValues returns OpConfigValues
func (m CreateTaskFromSqlTask) GetOpConfigValues() *ConfigValues {
	return m.OpConfigValues
}

//GetConfigProviderDelegate returns ConfigProviderDelegate
func (m CreateTaskFromSqlTask) GetConfigProviderDelegate() *CreateConfigProvider {
	return m.ConfigProviderDelegate
}

//GetRegistryMetadata returns RegistryMetadata
func (m CreateTaskFromSqlTask) GetRegistryMetadata() *RegistryMetadata {
	return m.RegistryMetadata
}

func (m CreateTaskFromSqlTask) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateTaskFromSqlTask) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateTaskFromSqlTask CreateTaskFromSqlTask
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeCreateTaskFromSqlTask
	}{
		"SQL_TASK",
		(MarshalTypeCreateTaskFromSqlTask)(m),
	}

	return json.Marshal(&s)
}

// CreateTaskFromSqlTaskSqlScriptTypeEnum Enum with underlying type: string
type CreateTaskFromSqlTaskSqlScriptTypeEnum string

// Set of constants representing the allowable values for CreateTaskFromSqlTaskSqlScriptTypeEnum
const (
	CreateTaskFromSqlTaskSqlScriptTypeStoredProcedure CreateTaskFromSqlTaskSqlScriptTypeEnum = "STORED_PROCEDURE"
	CreateTaskFromSqlTaskSqlScriptTypeSqlCode         CreateTaskFromSqlTaskSqlScriptTypeEnum = "SQL_CODE"
)

var mappingCreateTaskFromSqlTaskSqlScriptType = map[string]CreateTaskFromSqlTaskSqlScriptTypeEnum{
	"STORED_PROCEDURE": CreateTaskFromSqlTaskSqlScriptTypeStoredProcedure,
	"SQL_CODE":         CreateTaskFromSqlTaskSqlScriptTypeSqlCode,
}

// GetCreateTaskFromSqlTaskSqlScriptTypeEnumValues Enumerates the set of values for CreateTaskFromSqlTaskSqlScriptTypeEnum
func GetCreateTaskFromSqlTaskSqlScriptTypeEnumValues() []CreateTaskFromSqlTaskSqlScriptTypeEnum {
	values := make([]CreateTaskFromSqlTaskSqlScriptTypeEnum, 0)
	for _, v := range mappingCreateTaskFromSqlTaskSqlScriptType {
		values = append(values, v)
	}
	return values
}
