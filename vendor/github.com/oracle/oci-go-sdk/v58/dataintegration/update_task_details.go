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

// UpdateTaskDetails Properties used in task create operations.
type UpdateTaskDetails interface {

	// Generated key that can be used in API calls to identify task. On scenarios where reference to the task is needed, a value can be passed in create.
	GetKey() *string

	// The version of the object that is used to track changes in the object instance.
	GetObjectVersion() *int

	// The object's model version.
	GetModelVersion() *string

	GetParentRef() *ParentReference

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	GetName() *string

	// Detailed description for the object.
	GetDescription() *string

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	GetObjectStatus() *int

	// Value can only contain upper case letters, underscore and numbers. It should begin with upper case letter or underscore. The value can be modified.
	GetIdentifier() *string

	// An array of input ports.
	GetInputPorts() []InputPort

	// An array of output ports.
	GetOutputPorts() []OutputPort

	// An array of parameters.
	GetParameters() []Parameter

	GetOpConfigValues() *ConfigValues

	GetConfigProviderDelegate() *ConfigProvider

	GetRegistryMetadata() *RegistryMetadata
}

type updatetaskdetails struct {
	JsonData               []byte
	Key                    *string           `mandatory:"true" json:"key"`
	ObjectVersion          *int              `mandatory:"true" json:"objectVersion"`
	ModelVersion           *string           `mandatory:"false" json:"modelVersion"`
	ParentRef              *ParentReference  `mandatory:"false" json:"parentRef"`
	Name                   *string           `mandatory:"false" json:"name"`
	Description            *string           `mandatory:"false" json:"description"`
	ObjectStatus           *int              `mandatory:"false" json:"objectStatus"`
	Identifier             *string           `mandatory:"false" json:"identifier"`
	InputPorts             []InputPort       `mandatory:"false" json:"inputPorts"`
	OutputPorts            []OutputPort      `mandatory:"false" json:"outputPorts"`
	Parameters             []Parameter       `mandatory:"false" json:"parameters"`
	OpConfigValues         *ConfigValues     `mandatory:"false" json:"opConfigValues"`
	ConfigProviderDelegate *ConfigProvider   `mandatory:"false" json:"configProviderDelegate"`
	RegistryMetadata       *RegistryMetadata `mandatory:"false" json:"registryMetadata"`
	ModelType              string            `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *updatetaskdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdatetaskdetails updatetaskdetails
	s := struct {
		Model Unmarshalerupdatetaskdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Key = s.Model.Key
	m.ObjectVersion = s.Model.ObjectVersion
	m.ModelVersion = s.Model.ModelVersion
	m.ParentRef = s.Model.ParentRef
	m.Name = s.Model.Name
	m.Description = s.Model.Description
	m.ObjectStatus = s.Model.ObjectStatus
	m.Identifier = s.Model.Identifier
	m.InputPorts = s.Model.InputPorts
	m.OutputPorts = s.Model.OutputPorts
	m.Parameters = s.Model.Parameters
	m.OpConfigValues = s.Model.OpConfigValues
	m.ConfigProviderDelegate = s.Model.ConfigProviderDelegate
	m.RegistryMetadata = s.Model.RegistryMetadata
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updatetaskdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "REST_TASK":
		mm := UpdateTaskFromRestTask{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PIPELINE_TASK":
		mm := UpdateTaskFromPipelineTask{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OCI_DATAFLOW_TASK":
		mm := UpdateTaskFromOciDataflowTask{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SQL_TASK":
		mm := UpdateTaskFromSqlTask{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "DATA_LOADER_TASK":
		mm := UpdateTaskFromDataLoaderTask{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "INTEGRATION_TASK":
		mm := UpdateTaskFromIntegrationTask{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetKey returns Key
func (m updatetaskdetails) GetKey() *string {
	return m.Key
}

//GetObjectVersion returns ObjectVersion
func (m updatetaskdetails) GetObjectVersion() *int {
	return m.ObjectVersion
}

//GetModelVersion returns ModelVersion
func (m updatetaskdetails) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m updatetaskdetails) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetName returns Name
func (m updatetaskdetails) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m updatetaskdetails) GetDescription() *string {
	return m.Description
}

//GetObjectStatus returns ObjectStatus
func (m updatetaskdetails) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetIdentifier returns Identifier
func (m updatetaskdetails) GetIdentifier() *string {
	return m.Identifier
}

//GetInputPorts returns InputPorts
func (m updatetaskdetails) GetInputPorts() []InputPort {
	return m.InputPorts
}

//GetOutputPorts returns OutputPorts
func (m updatetaskdetails) GetOutputPorts() []OutputPort {
	return m.OutputPorts
}

//GetParameters returns Parameters
func (m updatetaskdetails) GetParameters() []Parameter {
	return m.Parameters
}

//GetOpConfigValues returns OpConfigValues
func (m updatetaskdetails) GetOpConfigValues() *ConfigValues {
	return m.OpConfigValues
}

//GetConfigProviderDelegate returns ConfigProviderDelegate
func (m updatetaskdetails) GetConfigProviderDelegate() *ConfigProvider {
	return m.ConfigProviderDelegate
}

//GetRegistryMetadata returns RegistryMetadata
func (m updatetaskdetails) GetRegistryMetadata() *RegistryMetadata {
	return m.RegistryMetadata
}

func (m updatetaskdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updatetaskdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateTaskDetailsModelTypeEnum Enum with underlying type: string
type UpdateTaskDetailsModelTypeEnum string

// Set of constants representing the allowable values for UpdateTaskDetailsModelTypeEnum
const (
	UpdateTaskDetailsModelTypeIntegrationTask UpdateTaskDetailsModelTypeEnum = "INTEGRATION_TASK"
	UpdateTaskDetailsModelTypeDataLoaderTask  UpdateTaskDetailsModelTypeEnum = "DATA_LOADER_TASK"
	UpdateTaskDetailsModelTypePipelineTask    UpdateTaskDetailsModelTypeEnum = "PIPELINE_TASK"
	UpdateTaskDetailsModelTypeSqlTask         UpdateTaskDetailsModelTypeEnum = "SQL_TASK"
	UpdateTaskDetailsModelTypeOciDataflowTask UpdateTaskDetailsModelTypeEnum = "OCI_DATAFLOW_TASK"
	UpdateTaskDetailsModelTypeRestTask        UpdateTaskDetailsModelTypeEnum = "REST_TASK"
)

var mappingUpdateTaskDetailsModelTypeEnum = map[string]UpdateTaskDetailsModelTypeEnum{
	"INTEGRATION_TASK":  UpdateTaskDetailsModelTypeIntegrationTask,
	"DATA_LOADER_TASK":  UpdateTaskDetailsModelTypeDataLoaderTask,
	"PIPELINE_TASK":     UpdateTaskDetailsModelTypePipelineTask,
	"SQL_TASK":          UpdateTaskDetailsModelTypeSqlTask,
	"OCI_DATAFLOW_TASK": UpdateTaskDetailsModelTypeOciDataflowTask,
	"REST_TASK":         UpdateTaskDetailsModelTypeRestTask,
}

// GetUpdateTaskDetailsModelTypeEnumValues Enumerates the set of values for UpdateTaskDetailsModelTypeEnum
func GetUpdateTaskDetailsModelTypeEnumValues() []UpdateTaskDetailsModelTypeEnum {
	values := make([]UpdateTaskDetailsModelTypeEnum, 0)
	for _, v := range mappingUpdateTaskDetailsModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateTaskDetailsModelTypeEnumStringValues Enumerates the set of values in String for UpdateTaskDetailsModelTypeEnum
func GetUpdateTaskDetailsModelTypeEnumStringValues() []string {
	return []string{
		"INTEGRATION_TASK",
		"DATA_LOADER_TASK",
		"PIPELINE_TASK",
		"SQL_TASK",
		"OCI_DATAFLOW_TASK",
		"REST_TASK",
	}
}

// GetMappingUpdateTaskDetailsModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateTaskDetailsModelTypeEnum(val string) (UpdateTaskDetailsModelTypeEnum, bool) {
	mappingUpdateTaskDetailsModelTypeEnumIgnoreCase := make(map[string]UpdateTaskDetailsModelTypeEnum)
	for k, v := range mappingUpdateTaskDetailsModelTypeEnum {
		mappingUpdateTaskDetailsModelTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingUpdateTaskDetailsModelTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
