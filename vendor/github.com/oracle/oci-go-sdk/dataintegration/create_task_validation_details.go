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

// CreateTaskValidationDetails The task type contains the audit summary information and the definition of the task.
type CreateTaskValidationDetails interface {

	// Generated key that can be used in API calls to identify task. On scenarios where reference to the task is needed, a value can be passed in create.
	GetKey() *string

	// The model version of an object.
	GetModelVersion() *string

	GetParentRef() *ParentReference

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value can be edited by the user and it is restricted to 1000 characters
	GetName() *string

	// Detailed description for the object.
	GetDescription() *string

	// The version of the object that is used to track changes in the object instance.
	GetObjectVersion() *int

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	GetObjectStatus() *int

	// Value can only contain upper case letters, underscore and numbers. It should begin with upper case letter or underscore. The value can be edited by the user.
	GetIdentifier() *string

	// An array of input ports.
	GetInputPorts() []InputPort

	// An array of output ports.
	GetOutputPorts() []OutputPort

	// An array of parameters.
	GetParameters() []Parameter

	GetOpConfigValues() *ConfigValues

	GetConfigProviderDelegate() *ConfigProvider

	GetMetadata() *ObjectMetadata
}

type createtaskvalidationdetails struct {
	JsonData               []byte
	Key                    *string          `mandatory:"false" json:"key"`
	ModelVersion           *string          `mandatory:"false" json:"modelVersion"`
	ParentRef              *ParentReference `mandatory:"false" json:"parentRef"`
	Name                   *string          `mandatory:"false" json:"name"`
	Description            *string          `mandatory:"false" json:"description"`
	ObjectVersion          *int             `mandatory:"false" json:"objectVersion"`
	ObjectStatus           *int             `mandatory:"false" json:"objectStatus"`
	Identifier             *string          `mandatory:"false" json:"identifier"`
	InputPorts             []InputPort      `mandatory:"false" json:"inputPorts"`
	OutputPorts            []OutputPort     `mandatory:"false" json:"outputPorts"`
	Parameters             []Parameter      `mandatory:"false" json:"parameters"`
	OpConfigValues         *ConfigValues    `mandatory:"false" json:"opConfigValues"`
	ConfigProviderDelegate *ConfigProvider  `mandatory:"false" json:"configProviderDelegate"`
	Metadata               *ObjectMetadata  `mandatory:"false" json:"metadata"`
	ModelType              string           `json:"modelType"`
}

// UnmarshalJSON unmarshals json
func (m *createtaskvalidationdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatetaskvalidationdetails createtaskvalidationdetails
	s := struct {
		Model Unmarshalercreatetaskvalidationdetails
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
	m.Metadata = s.Model.Metadata
	m.ModelType = s.Model.ModelType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createtaskvalidationdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ModelType {
	case "DATA_LOADER_TASK":
		mm := CreateTaskValidationFromDataLoaderTask{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "INTEGRATION_TASK":
		mm := CreateTaskValidationFromIntegrationTask{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetKey returns Key
func (m createtaskvalidationdetails) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m createtaskvalidationdetails) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m createtaskvalidationdetails) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetName returns Name
func (m createtaskvalidationdetails) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m createtaskvalidationdetails) GetDescription() *string {
	return m.Description
}

//GetObjectVersion returns ObjectVersion
func (m createtaskvalidationdetails) GetObjectVersion() *int {
	return m.ObjectVersion
}

//GetObjectStatus returns ObjectStatus
func (m createtaskvalidationdetails) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetIdentifier returns Identifier
func (m createtaskvalidationdetails) GetIdentifier() *string {
	return m.Identifier
}

//GetInputPorts returns InputPorts
func (m createtaskvalidationdetails) GetInputPorts() []InputPort {
	return m.InputPorts
}

//GetOutputPorts returns OutputPorts
func (m createtaskvalidationdetails) GetOutputPorts() []OutputPort {
	return m.OutputPorts
}

//GetParameters returns Parameters
func (m createtaskvalidationdetails) GetParameters() []Parameter {
	return m.Parameters
}

//GetOpConfigValues returns OpConfigValues
func (m createtaskvalidationdetails) GetOpConfigValues() *ConfigValues {
	return m.OpConfigValues
}

//GetConfigProviderDelegate returns ConfigProviderDelegate
func (m createtaskvalidationdetails) GetConfigProviderDelegate() *ConfigProvider {
	return m.ConfigProviderDelegate
}

//GetMetadata returns Metadata
func (m createtaskvalidationdetails) GetMetadata() *ObjectMetadata {
	return m.Metadata
}

func (m createtaskvalidationdetails) String() string {
	return common.PointerString(m)
}

// CreateTaskValidationDetailsModelTypeEnum Enum with underlying type: string
type CreateTaskValidationDetailsModelTypeEnum string

// Set of constants representing the allowable values for CreateTaskValidationDetailsModelTypeEnum
const (
	CreateTaskValidationDetailsModelTypeIntegrationTask CreateTaskValidationDetailsModelTypeEnum = "INTEGRATION_TASK"
	CreateTaskValidationDetailsModelTypeDataLoaderTask  CreateTaskValidationDetailsModelTypeEnum = "DATA_LOADER_TASK"
)

var mappingCreateTaskValidationDetailsModelType = map[string]CreateTaskValidationDetailsModelTypeEnum{
	"INTEGRATION_TASK": CreateTaskValidationDetailsModelTypeIntegrationTask,
	"DATA_LOADER_TASK": CreateTaskValidationDetailsModelTypeDataLoaderTask,
}

// GetCreateTaskValidationDetailsModelTypeEnumValues Enumerates the set of values for CreateTaskValidationDetailsModelTypeEnum
func GetCreateTaskValidationDetailsModelTypeEnumValues() []CreateTaskValidationDetailsModelTypeEnum {
	values := make([]CreateTaskValidationDetailsModelTypeEnum, 0)
	for _, v := range mappingCreateTaskValidationDetailsModelType {
		values = append(values, v)
	}
	return values
}
