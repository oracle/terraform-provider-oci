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

// CreateTaskDetails Properties used in task create operations.
type CreateTaskDetails interface {

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	GetName() *string

	// Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
	GetIdentifier() *string

	GetRegistryMetadata() *RegistryMetadata

	// Generated key that can be used in API calls to identify task. On scenarios where reference to the task is needed, a value can be passed in create.
	GetKey() *string

	// The object's model version.
	GetModelVersion() *string

	GetParentRef() *ParentReference

	// Detailed description for the object.
	GetDescription() *string

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	GetObjectStatus() *int

	// An array of input ports.
	GetInputPorts() []InputPort

	// An array of output ports.
	GetOutputPorts() []OutputPort

	// An array of parameters.
	GetParameters() []Parameter

	GetOpConfigValues() *ConfigValues

	GetConfigProviderDelegate() *CreateConfigProvider
}

type createtaskdetails struct {
	JsonData               []byte
	Name                   *string               `mandatory:"true" json:"name"`
	Identifier             *string               `mandatory:"true" json:"identifier"`
	RegistryMetadata       *RegistryMetadata     `mandatory:"true" json:"registryMetadata"`
	Key                    *string               `mandatory:"false" json:"key"`
	ModelVersion           *string               `mandatory:"false" json:"modelVersion"`
	ParentRef              *ParentReference      `mandatory:"false" json:"parentRef"`
	Description            *string               `mandatory:"false" json:"description"`
	ObjectStatus           *int                  `mandatory:"false" json:"objectStatus"`
	InputPorts             []InputPort           `mandatory:"false" json:"inputPorts"`
	OutputPorts            []OutputPort          `mandatory:"false" json:"outputPorts"`
	Parameters             []Parameter           `mandatory:"false" json:"parameters"`
	OpConfigValues         *ConfigValues         `mandatory:"false" json:"opConfigValues"`
	ConfigProviderDelegate *CreateConfigProvider `mandatory:"false" json:"configProviderDelegate"`
	ModelType              string                `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *createtaskdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatetaskdetails createtaskdetails
	s := struct {
		Model Unmarshalercreatetaskdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Name = s.Model.Name
	m.Identifier = s.Model.Identifier
	m.RegistryMetadata = s.Model.RegistryMetadata
	m.Key = s.Model.Key
	m.ModelVersion = s.Model.ModelVersion
	m.ParentRef = s.Model.ParentRef
	m.Description = s.Model.Description
	m.ObjectStatus = s.Model.ObjectStatus
	m.InputPorts = s.Model.InputPorts
	m.OutputPorts = s.Model.OutputPorts
	m.Parameters = s.Model.Parameters
	m.OpConfigValues = s.Model.OpConfigValues
	m.ConfigProviderDelegate = s.Model.ConfigProviderDelegate
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createtaskdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "INTEGRATION_TASK":
		mm := CreateTaskFromIntegrationTask{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DATA_LOADER_TASK":
		mm := CreateTaskFromDataLoaderTask{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PIPELINE_TASK":
		mm := CreateTaskFromPipelineTask{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_DATAFLOW_TASK":
		mm := CreateTaskFromOciDataflowTask{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SQL_TASK":
		mm := CreateTaskFromSqlTask{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "REST_TASK":
		mm := CreateTaskFromRestTask{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetName returns Name
func (m createtaskdetails) GetName() *string {
	return m.Name
}

//GetIdentifier returns Identifier
func (m createtaskdetails) GetIdentifier() *string {
	return m.Identifier
}

//GetRegistryMetadata returns RegistryMetadata
func (m createtaskdetails) GetRegistryMetadata() *RegistryMetadata {
	return m.RegistryMetadata
}

//GetKey returns Key
func (m createtaskdetails) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m createtaskdetails) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m createtaskdetails) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetDescription returns Description
func (m createtaskdetails) GetDescription() *string {
	return m.Description
}

//GetObjectStatus returns ObjectStatus
func (m createtaskdetails) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetInputPorts returns InputPorts
func (m createtaskdetails) GetInputPorts() []InputPort {
	return m.InputPorts
}

//GetOutputPorts returns OutputPorts
func (m createtaskdetails) GetOutputPorts() []OutputPort {
	return m.OutputPorts
}

//GetParameters returns Parameters
func (m createtaskdetails) GetParameters() []Parameter {
	return m.Parameters
}

//GetOpConfigValues returns OpConfigValues
func (m createtaskdetails) GetOpConfigValues() *ConfigValues {
	return m.OpConfigValues
}

//GetConfigProviderDelegate returns ConfigProviderDelegate
func (m createtaskdetails) GetConfigProviderDelegate() *CreateConfigProvider {
	return m.ConfigProviderDelegate
}

func (m createtaskdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createtaskdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateTaskDetailsModelTypeEnum Enum with underlying type: string
type CreateTaskDetailsModelTypeEnum string

// Set of constants representing the allowable values for CreateTaskDetailsModelTypeEnum
const (
	CreateTaskDetailsModelTypeIntegrationTask CreateTaskDetailsModelTypeEnum = "INTEGRATION_TASK"
	CreateTaskDetailsModelTypeDataLoaderTask  CreateTaskDetailsModelTypeEnum = "DATA_LOADER_TASK"
	CreateTaskDetailsModelTypePipelineTask    CreateTaskDetailsModelTypeEnum = "PIPELINE_TASK"
	CreateTaskDetailsModelTypeSqlTask         CreateTaskDetailsModelTypeEnum = "SQL_TASK"
	CreateTaskDetailsModelTypeOciDataflowTask CreateTaskDetailsModelTypeEnum = "OCI_DATAFLOW_TASK"
	CreateTaskDetailsModelTypeRestTask        CreateTaskDetailsModelTypeEnum = "REST_TASK"
)

var mappingCreateTaskDetailsModelTypeEnum = map[string]CreateTaskDetailsModelTypeEnum{
	"INTEGRATION_TASK":  CreateTaskDetailsModelTypeIntegrationTask,
	"DATA_LOADER_TASK":  CreateTaskDetailsModelTypeDataLoaderTask,
	"PIPELINE_TASK":     CreateTaskDetailsModelTypePipelineTask,
	"SQL_TASK":          CreateTaskDetailsModelTypeSqlTask,
	"OCI_DATAFLOW_TASK": CreateTaskDetailsModelTypeOciDataflowTask,
	"REST_TASK":         CreateTaskDetailsModelTypeRestTask,
}

// GetCreateTaskDetailsModelTypeEnumValues Enumerates the set of values for CreateTaskDetailsModelTypeEnum
func GetCreateTaskDetailsModelTypeEnumValues() []CreateTaskDetailsModelTypeEnum {
	values := make([]CreateTaskDetailsModelTypeEnum, 0)
	for _, v := range mappingCreateTaskDetailsModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateTaskDetailsModelTypeEnumStringValues Enumerates the set of values in String for CreateTaskDetailsModelTypeEnum
func GetCreateTaskDetailsModelTypeEnumStringValues() []string {
	return []string{
		"INTEGRATION_TASK",
		"DATA_LOADER_TASK",
		"PIPELINE_TASK",
		"SQL_TASK",
		"OCI_DATAFLOW_TASK",
		"REST_TASK",
	}
}

// GetMappingCreateTaskDetailsModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateTaskDetailsModelTypeEnum(val string) (CreateTaskDetailsModelTypeEnum, bool) {
	mappingCreateTaskDetailsModelTypeEnumIgnoreCase := make(map[string]CreateTaskDetailsModelTypeEnum)
	for k, v := range mappingCreateTaskDetailsModelTypeEnum {
		mappingCreateTaskDetailsModelTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingCreateTaskDetailsModelTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
