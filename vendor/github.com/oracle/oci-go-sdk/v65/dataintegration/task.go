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

// Task The task type contains the audit summary information and the definition of the task.
type Task interface {

	// Generated key that can be used in API calls to identify task. On scenarios where reference to the task is needed, a value can be passed in create.
	GetKey() *string

	// The object's model version.
	GetModelVersion() *string

	GetParentRef() *ParentReference

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	GetName() *string

	// Detailed description for the object.
	GetDescription() *string

	// The version of the object that is used to track changes in the object instance.
	GetObjectVersion() *int

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	GetObjectStatus() *int

	// Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
	GetIdentifier() *string

	// An array of input ports.
	GetInputPorts() []InputPort

	// An array of output ports.
	GetOutputPorts() []OutputPort

	// An array of parameters.
	GetParameters() []Parameter

	GetOpConfigValues() *ConfigValues

	GetConfigProviderDelegate() *ConfigProvider

	// Whether the same task can be executed concurrently.
	GetIsConcurrentAllowed() *bool

	GetMetadata() *ObjectMetadata

	// A key map. If provided, key is replaced with generated key. This structure provides mapping between user provided key and generated key.
	GetKeyMap() map[string]string

	GetRegistryMetadata() *RegistryMetadata
}

type task struct {
	JsonData               []byte
	Key                    *string           `mandatory:"false" json:"key"`
	ModelVersion           *string           `mandatory:"false" json:"modelVersion"`
	ParentRef              *ParentReference  `mandatory:"false" json:"parentRef"`
	Name                   *string           `mandatory:"false" json:"name"`
	Description            *string           `mandatory:"false" json:"description"`
	ObjectVersion          *int              `mandatory:"false" json:"objectVersion"`
	ObjectStatus           *int              `mandatory:"false" json:"objectStatus"`
	Identifier             *string           `mandatory:"false" json:"identifier"`
	InputPorts             []InputPort       `mandatory:"false" json:"inputPorts"`
	OutputPorts            []OutputPort      `mandatory:"false" json:"outputPorts"`
	Parameters             []Parameter       `mandatory:"false" json:"parameters"`
	OpConfigValues         *ConfigValues     `mandatory:"false" json:"opConfigValues"`
	ConfigProviderDelegate *ConfigProvider   `mandatory:"false" json:"configProviderDelegate"`
	IsConcurrentAllowed    *bool             `mandatory:"false" json:"isConcurrentAllowed"`
	Metadata               *ObjectMetadata   `mandatory:"false" json:"metadata"`
	KeyMap                 map[string]string `mandatory:"false" json:"keyMap"`
	RegistryMetadata       *RegistryMetadata `mandatory:"false" json:"registryMetadata"`
	ModelType              string            `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *task) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalertask task
	s := struct {
		Model Unmarshalertask
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
	m.ObjectStatus = s.Model.ObjectStatus
	m.Identifier = s.Model.Identifier
	m.InputPorts = s.Model.InputPorts
	m.OutputPorts = s.Model.OutputPorts
	m.Parameters = s.Model.Parameters
	m.OpConfigValues = s.Model.OpConfigValues
	m.ConfigProviderDelegate = s.Model.ConfigProviderDelegate
	m.IsConcurrentAllowed = s.Model.IsConcurrentAllowed
	m.Metadata = s.Model.Metadata
	m.KeyMap = s.Model.KeyMap
	m.RegistryMetadata = s.Model.RegistryMetadata
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *task) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "PIPELINE_TASK":
		mm := TaskFromPipelineTaskDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "INTEGRATION_TASK":
		mm := TaskFromIntegrationTaskDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SQL_TASK":
		mm := TaskFromSqlTaskDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "REST_TASK":
		mm := TaskFromRestTaskDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_DATAFLOW_TASK":
		mm := TaskFromOciDataflowTaskDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DATA_LOADER_TASK":
		mm := TaskFromDataLoaderTaskDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for Task: %s.", m.ModelType)
		return *m, nil
	}
}

// GetKey returns Key
func (m task) GetKey() *string {
	return m.Key
}

// GetModelVersion returns ModelVersion
func (m task) GetModelVersion() *string {
	return m.ModelVersion
}

// GetParentRef returns ParentRef
func (m task) GetParentRef() *ParentReference {
	return m.ParentRef
}

// GetName returns Name
func (m task) GetName() *string {
	return m.Name
}

// GetDescription returns Description
func (m task) GetDescription() *string {
	return m.Description
}

// GetObjectVersion returns ObjectVersion
func (m task) GetObjectVersion() *int {
	return m.ObjectVersion
}

// GetObjectStatus returns ObjectStatus
func (m task) GetObjectStatus() *int {
	return m.ObjectStatus
}

// GetIdentifier returns Identifier
func (m task) GetIdentifier() *string {
	return m.Identifier
}

// GetInputPorts returns InputPorts
func (m task) GetInputPorts() []InputPort {
	return m.InputPorts
}

// GetOutputPorts returns OutputPorts
func (m task) GetOutputPorts() []OutputPort {
	return m.OutputPorts
}

// GetParameters returns Parameters
func (m task) GetParameters() []Parameter {
	return m.Parameters
}

// GetOpConfigValues returns OpConfigValues
func (m task) GetOpConfigValues() *ConfigValues {
	return m.OpConfigValues
}

// GetConfigProviderDelegate returns ConfigProviderDelegate
func (m task) GetConfigProviderDelegate() *ConfigProvider {
	return m.ConfigProviderDelegate
}

// GetIsConcurrentAllowed returns IsConcurrentAllowed
func (m task) GetIsConcurrentAllowed() *bool {
	return m.IsConcurrentAllowed
}

// GetMetadata returns Metadata
func (m task) GetMetadata() *ObjectMetadata {
	return m.Metadata
}

// GetKeyMap returns KeyMap
func (m task) GetKeyMap() map[string]string {
	return m.KeyMap
}

// GetRegistryMetadata returns RegistryMetadata
func (m task) GetRegistryMetadata() *RegistryMetadata {
	return m.RegistryMetadata
}

func (m task) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m task) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TaskModelTypeEnum Enum with underlying type: string
type TaskModelTypeEnum string

// Set of constants representing the allowable values for TaskModelTypeEnum
const (
	TaskModelTypeIntegrationTask TaskModelTypeEnum = "INTEGRATION_TASK"
	TaskModelTypeDataLoaderTask  TaskModelTypeEnum = "DATA_LOADER_TASK"
	TaskModelTypePipelineTask    TaskModelTypeEnum = "PIPELINE_TASK"
	TaskModelTypeSqlTask         TaskModelTypeEnum = "SQL_TASK"
	TaskModelTypeOciDataflowTask TaskModelTypeEnum = "OCI_DATAFLOW_TASK"
	TaskModelTypeRestTask        TaskModelTypeEnum = "REST_TASK"
)

var mappingTaskModelTypeEnum = map[string]TaskModelTypeEnum{
	"INTEGRATION_TASK":  TaskModelTypeIntegrationTask,
	"DATA_LOADER_TASK":  TaskModelTypeDataLoaderTask,
	"PIPELINE_TASK":     TaskModelTypePipelineTask,
	"SQL_TASK":          TaskModelTypeSqlTask,
	"OCI_DATAFLOW_TASK": TaskModelTypeOciDataflowTask,
	"REST_TASK":         TaskModelTypeRestTask,
}

var mappingTaskModelTypeEnumLowerCase = map[string]TaskModelTypeEnum{
	"integration_task":  TaskModelTypeIntegrationTask,
	"data_loader_task":  TaskModelTypeDataLoaderTask,
	"pipeline_task":     TaskModelTypePipelineTask,
	"sql_task":          TaskModelTypeSqlTask,
	"oci_dataflow_task": TaskModelTypeOciDataflowTask,
	"rest_task":         TaskModelTypeRestTask,
}

// GetTaskModelTypeEnumValues Enumerates the set of values for TaskModelTypeEnum
func GetTaskModelTypeEnumValues() []TaskModelTypeEnum {
	values := make([]TaskModelTypeEnum, 0)
	for _, v := range mappingTaskModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTaskModelTypeEnumStringValues Enumerates the set of values in String for TaskModelTypeEnum
func GetTaskModelTypeEnumStringValues() []string {
	return []string{
		"INTEGRATION_TASK",
		"DATA_LOADER_TASK",
		"PIPELINE_TASK",
		"SQL_TASK",
		"OCI_DATAFLOW_TASK",
		"REST_TASK",
	}
}

// GetMappingTaskModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTaskModelTypeEnum(val string) (TaskModelTypeEnum, bool) {
	enum, ok := mappingTaskModelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
