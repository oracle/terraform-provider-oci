// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration Service APIs to perform common extract, load, and transform (ETL) tasks.
//

package dataintegration

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v45/common"
)

// CreateTaskFromRestTask The information about the Generic REST task.
type CreateTaskFromRestTask struct {

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

	AuthDetails *AuthDetails `mandatory:"false" json:"authDetails"`

	Endpoint *Expression `mandatory:"false" json:"endpoint"`

	// The headers for the REST call.
	Headers *interface{} `mandatory:"false" json:"headers"`

	// JSON data for payload body.
	JsonData *string `mandatory:"false" json:"jsonData"`

	CancelEndpoint *Expression `mandatory:"false" json:"cancelEndpoint"`

	// The REST method to use.
	MethodType CreateTaskFromRestTaskMethodTypeEnum `mandatory:"false" json:"methodType,omitempty"`

	// The invocation type to be used for Generic REST invocation.
	ApiCallMode CreateTaskFromRestTaskApiCallModeEnum `mandatory:"false" json:"apiCallMode,omitempty"`

	// The REST method to use for canceling the original request.
	CancelMethodType CreateTaskFromRestTaskCancelMethodTypeEnum `mandatory:"false" json:"cancelMethodType,omitempty"`
}

//GetKey returns Key
func (m CreateTaskFromRestTask) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m CreateTaskFromRestTask) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m CreateTaskFromRestTask) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetName returns Name
func (m CreateTaskFromRestTask) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m CreateTaskFromRestTask) GetDescription() *string {
	return m.Description
}

//GetObjectStatus returns ObjectStatus
func (m CreateTaskFromRestTask) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetIdentifier returns Identifier
func (m CreateTaskFromRestTask) GetIdentifier() *string {
	return m.Identifier
}

//GetInputPorts returns InputPorts
func (m CreateTaskFromRestTask) GetInputPorts() []InputPort {
	return m.InputPorts
}

//GetOutputPorts returns OutputPorts
func (m CreateTaskFromRestTask) GetOutputPorts() []OutputPort {
	return m.OutputPorts
}

//GetParameters returns Parameters
func (m CreateTaskFromRestTask) GetParameters() []Parameter {
	return m.Parameters
}

//GetOpConfigValues returns OpConfigValues
func (m CreateTaskFromRestTask) GetOpConfigValues() *ConfigValues {
	return m.OpConfigValues
}

//GetConfigProviderDelegate returns ConfigProviderDelegate
func (m CreateTaskFromRestTask) GetConfigProviderDelegate() *CreateConfigProvider {
	return m.ConfigProviderDelegate
}

//GetRegistryMetadata returns RegistryMetadata
func (m CreateTaskFromRestTask) GetRegistryMetadata() *RegistryMetadata {
	return m.RegistryMetadata
}

func (m CreateTaskFromRestTask) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m CreateTaskFromRestTask) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateTaskFromRestTask CreateTaskFromRestTask
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeCreateTaskFromRestTask
	}{
		"REST_TASK",
		(MarshalTypeCreateTaskFromRestTask)(m),
	}

	return json.Marshal(&s)
}

// CreateTaskFromRestTaskMethodTypeEnum Enum with underlying type: string
type CreateTaskFromRestTaskMethodTypeEnum string

// Set of constants representing the allowable values for CreateTaskFromRestTaskMethodTypeEnum
const (
	CreateTaskFromRestTaskMethodTypeGet    CreateTaskFromRestTaskMethodTypeEnum = "GET"
	CreateTaskFromRestTaskMethodTypePost   CreateTaskFromRestTaskMethodTypeEnum = "POST"
	CreateTaskFromRestTaskMethodTypePatch  CreateTaskFromRestTaskMethodTypeEnum = "PATCH"
	CreateTaskFromRestTaskMethodTypeDelete CreateTaskFromRestTaskMethodTypeEnum = "DELETE"
	CreateTaskFromRestTaskMethodTypePut    CreateTaskFromRestTaskMethodTypeEnum = "PUT"
)

var mappingCreateTaskFromRestTaskMethodType = map[string]CreateTaskFromRestTaskMethodTypeEnum{
	"GET":    CreateTaskFromRestTaskMethodTypeGet,
	"POST":   CreateTaskFromRestTaskMethodTypePost,
	"PATCH":  CreateTaskFromRestTaskMethodTypePatch,
	"DELETE": CreateTaskFromRestTaskMethodTypeDelete,
	"PUT":    CreateTaskFromRestTaskMethodTypePut,
}

// GetCreateTaskFromRestTaskMethodTypeEnumValues Enumerates the set of values for CreateTaskFromRestTaskMethodTypeEnum
func GetCreateTaskFromRestTaskMethodTypeEnumValues() []CreateTaskFromRestTaskMethodTypeEnum {
	values := make([]CreateTaskFromRestTaskMethodTypeEnum, 0)
	for _, v := range mappingCreateTaskFromRestTaskMethodType {
		values = append(values, v)
	}
	return values
}

// CreateTaskFromRestTaskApiCallModeEnum Enum with underlying type: string
type CreateTaskFromRestTaskApiCallModeEnum string

// Set of constants representing the allowable values for CreateTaskFromRestTaskApiCallModeEnum
const (
	CreateTaskFromRestTaskApiCallModeSynchronous         CreateTaskFromRestTaskApiCallModeEnum = "SYNCHRONOUS"
	CreateTaskFromRestTaskApiCallModeAsyncOciWorkrequest CreateTaskFromRestTaskApiCallModeEnum = "ASYNC_OCI_WORKREQUEST"
)

var mappingCreateTaskFromRestTaskApiCallMode = map[string]CreateTaskFromRestTaskApiCallModeEnum{
	"SYNCHRONOUS":           CreateTaskFromRestTaskApiCallModeSynchronous,
	"ASYNC_OCI_WORKREQUEST": CreateTaskFromRestTaskApiCallModeAsyncOciWorkrequest,
}

// GetCreateTaskFromRestTaskApiCallModeEnumValues Enumerates the set of values for CreateTaskFromRestTaskApiCallModeEnum
func GetCreateTaskFromRestTaskApiCallModeEnumValues() []CreateTaskFromRestTaskApiCallModeEnum {
	values := make([]CreateTaskFromRestTaskApiCallModeEnum, 0)
	for _, v := range mappingCreateTaskFromRestTaskApiCallMode {
		values = append(values, v)
	}
	return values
}

// CreateTaskFromRestTaskCancelMethodTypeEnum Enum with underlying type: string
type CreateTaskFromRestTaskCancelMethodTypeEnum string

// Set of constants representing the allowable values for CreateTaskFromRestTaskCancelMethodTypeEnum
const (
	CreateTaskFromRestTaskCancelMethodTypeGet    CreateTaskFromRestTaskCancelMethodTypeEnum = "GET"
	CreateTaskFromRestTaskCancelMethodTypePost   CreateTaskFromRestTaskCancelMethodTypeEnum = "POST"
	CreateTaskFromRestTaskCancelMethodTypePatch  CreateTaskFromRestTaskCancelMethodTypeEnum = "PATCH"
	CreateTaskFromRestTaskCancelMethodTypeDelete CreateTaskFromRestTaskCancelMethodTypeEnum = "DELETE"
	CreateTaskFromRestTaskCancelMethodTypePut    CreateTaskFromRestTaskCancelMethodTypeEnum = "PUT"
)

var mappingCreateTaskFromRestTaskCancelMethodType = map[string]CreateTaskFromRestTaskCancelMethodTypeEnum{
	"GET":    CreateTaskFromRestTaskCancelMethodTypeGet,
	"POST":   CreateTaskFromRestTaskCancelMethodTypePost,
	"PATCH":  CreateTaskFromRestTaskCancelMethodTypePatch,
	"DELETE": CreateTaskFromRestTaskCancelMethodTypeDelete,
	"PUT":    CreateTaskFromRestTaskCancelMethodTypePut,
}

// GetCreateTaskFromRestTaskCancelMethodTypeEnumValues Enumerates the set of values for CreateTaskFromRestTaskCancelMethodTypeEnum
func GetCreateTaskFromRestTaskCancelMethodTypeEnumValues() []CreateTaskFromRestTaskCancelMethodTypeEnum {
	values := make([]CreateTaskFromRestTaskCancelMethodTypeEnum, 0)
	for _, v := range mappingCreateTaskFromRestTaskCancelMethodType {
		values = append(values, v)
	}
	return values
}
