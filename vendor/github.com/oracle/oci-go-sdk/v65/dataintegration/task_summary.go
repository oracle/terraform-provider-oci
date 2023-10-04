// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// TaskSummary The task summary object type contains the audit summary information and the definition of the task summary object.
type TaskSummary interface {

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
}

type tasksummary struct {
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
	ModelType              string            `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *tasksummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalertasksummary tasksummary
	s := struct {
		Model Unmarshalertasksummary
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
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *tasksummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "INTEGRATION_TASK":
		mm := TaskSummaryFromIntegrationTask{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SQL_TASK":
		mm := TaskSummaryFromSqlTask{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_DATAFLOW_TASK":
		mm := TaskSummaryFromOciDataflowTask{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "REST_TASK":
		mm := TaskSummaryFromRestTask{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PIPELINE_TASK":
		mm := TaskSummaryFromPipelineTask{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DATA_LOADER_TASK":
		mm := TaskSummaryFromDataLoaderTask{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for TaskSummary: %s.", m.ModelType)
		return *m, nil
	}
}

// GetKey returns Key
func (m tasksummary) GetKey() *string {
	return m.Key
}

// GetModelVersion returns ModelVersion
func (m tasksummary) GetModelVersion() *string {
	return m.ModelVersion
}

// GetParentRef returns ParentRef
func (m tasksummary) GetParentRef() *ParentReference {
	return m.ParentRef
}

// GetName returns Name
func (m tasksummary) GetName() *string {
	return m.Name
}

// GetDescription returns Description
func (m tasksummary) GetDescription() *string {
	return m.Description
}

// GetObjectVersion returns ObjectVersion
func (m tasksummary) GetObjectVersion() *int {
	return m.ObjectVersion
}

// GetObjectStatus returns ObjectStatus
func (m tasksummary) GetObjectStatus() *int {
	return m.ObjectStatus
}

// GetIdentifier returns Identifier
func (m tasksummary) GetIdentifier() *string {
	return m.Identifier
}

// GetInputPorts returns InputPorts
func (m tasksummary) GetInputPorts() []InputPort {
	return m.InputPorts
}

// GetOutputPorts returns OutputPorts
func (m tasksummary) GetOutputPorts() []OutputPort {
	return m.OutputPorts
}

// GetParameters returns Parameters
func (m tasksummary) GetParameters() []Parameter {
	return m.Parameters
}

// GetOpConfigValues returns OpConfigValues
func (m tasksummary) GetOpConfigValues() *ConfigValues {
	return m.OpConfigValues
}

// GetConfigProviderDelegate returns ConfigProviderDelegate
func (m tasksummary) GetConfigProviderDelegate() *ConfigProvider {
	return m.ConfigProviderDelegate
}

// GetIsConcurrentAllowed returns IsConcurrentAllowed
func (m tasksummary) GetIsConcurrentAllowed() *bool {
	return m.IsConcurrentAllowed
}

// GetMetadata returns Metadata
func (m tasksummary) GetMetadata() *ObjectMetadata {
	return m.Metadata
}

// GetKeyMap returns KeyMap
func (m tasksummary) GetKeyMap() map[string]string {
	return m.KeyMap
}

func (m tasksummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m tasksummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TaskSummaryModelTypeEnum Enum with underlying type: string
type TaskSummaryModelTypeEnum string

// Set of constants representing the allowable values for TaskSummaryModelTypeEnum
const (
	TaskSummaryModelTypeIntegrationTask TaskSummaryModelTypeEnum = "INTEGRATION_TASK"
	TaskSummaryModelTypeDataLoaderTask  TaskSummaryModelTypeEnum = "DATA_LOADER_TASK"
	TaskSummaryModelTypePipelineTask    TaskSummaryModelTypeEnum = "PIPELINE_TASK"
	TaskSummaryModelTypeSqlTask         TaskSummaryModelTypeEnum = "SQL_TASK"
	TaskSummaryModelTypeOciDataflowTask TaskSummaryModelTypeEnum = "OCI_DATAFLOW_TASK"
	TaskSummaryModelTypeRestTask        TaskSummaryModelTypeEnum = "REST_TASK"
)

var mappingTaskSummaryModelTypeEnum = map[string]TaskSummaryModelTypeEnum{
	"INTEGRATION_TASK":  TaskSummaryModelTypeIntegrationTask,
	"DATA_LOADER_TASK":  TaskSummaryModelTypeDataLoaderTask,
	"PIPELINE_TASK":     TaskSummaryModelTypePipelineTask,
	"SQL_TASK":          TaskSummaryModelTypeSqlTask,
	"OCI_DATAFLOW_TASK": TaskSummaryModelTypeOciDataflowTask,
	"REST_TASK":         TaskSummaryModelTypeRestTask,
}

var mappingTaskSummaryModelTypeEnumLowerCase = map[string]TaskSummaryModelTypeEnum{
	"integration_task":  TaskSummaryModelTypeIntegrationTask,
	"data_loader_task":  TaskSummaryModelTypeDataLoaderTask,
	"pipeline_task":     TaskSummaryModelTypePipelineTask,
	"sql_task":          TaskSummaryModelTypeSqlTask,
	"oci_dataflow_task": TaskSummaryModelTypeOciDataflowTask,
	"rest_task":         TaskSummaryModelTypeRestTask,
}

// GetTaskSummaryModelTypeEnumValues Enumerates the set of values for TaskSummaryModelTypeEnum
func GetTaskSummaryModelTypeEnumValues() []TaskSummaryModelTypeEnum {
	values := make([]TaskSummaryModelTypeEnum, 0)
	for _, v := range mappingTaskSummaryModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTaskSummaryModelTypeEnumStringValues Enumerates the set of values in String for TaskSummaryModelTypeEnum
func GetTaskSummaryModelTypeEnumStringValues() []string {
	return []string{
		"INTEGRATION_TASK",
		"DATA_LOADER_TASK",
		"PIPELINE_TASK",
		"SQL_TASK",
		"OCI_DATAFLOW_TASK",
		"REST_TASK",
	}
}

// GetMappingTaskSummaryModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTaskSummaryModelTypeEnum(val string) (TaskSummaryModelTypeEnum, bool) {
	enum, ok := mappingTaskSummaryModelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
