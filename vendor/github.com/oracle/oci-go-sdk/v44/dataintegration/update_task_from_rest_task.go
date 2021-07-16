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
	"github.com/oracle/oci-go-sdk/v44/common"
)

// UpdateTaskFromRestTask The information about the Generic REST task.
type UpdateTaskFromRestTask struct {

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

	AuthDetails *AuthDetails `mandatory:"false" json:"authDetails"`

	Endpoint *Expression `mandatory:"false" json:"endpoint"`

	// The headers for the REST call.
	Headers *interface{} `mandatory:"false" json:"headers"`

	// Header value.
	AdditionalProperties *string `mandatory:"false" json:"additionalProperties"`

	// JSON data for payload body.
	JsonData *string `mandatory:"false" json:"jsonData"`

	CancelEndpoint *Expression `mandatory:"false" json:"cancelEndpoint"`

	// The REST method to use.
	MethodType UpdateTaskFromRestTaskMethodTypeEnum `mandatory:"false" json:"methodType,omitempty"`

	// The invocation type to be used for Generic REST invocation.
	ApiCallMode UpdateTaskFromRestTaskApiCallModeEnum `mandatory:"false" json:"apiCallMode,omitempty"`

	// The REST method to use for canceling the original request.
	CancelMethodType UpdateTaskFromRestTaskCancelMethodTypeEnum `mandatory:"false" json:"cancelMethodType,omitempty"`
}

//GetKey returns Key
func (m UpdateTaskFromRestTask) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m UpdateTaskFromRestTask) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m UpdateTaskFromRestTask) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetName returns Name
func (m UpdateTaskFromRestTask) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m UpdateTaskFromRestTask) GetDescription() *string {
	return m.Description
}

//GetObjectStatus returns ObjectStatus
func (m UpdateTaskFromRestTask) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetObjectVersion returns ObjectVersion
func (m UpdateTaskFromRestTask) GetObjectVersion() *int {
	return m.ObjectVersion
}

//GetIdentifier returns Identifier
func (m UpdateTaskFromRestTask) GetIdentifier() *string {
	return m.Identifier
}

//GetInputPorts returns InputPorts
func (m UpdateTaskFromRestTask) GetInputPorts() []InputPort {
	return m.InputPorts
}

//GetOutputPorts returns OutputPorts
func (m UpdateTaskFromRestTask) GetOutputPorts() []OutputPort {
	return m.OutputPorts
}

//GetParameters returns Parameters
func (m UpdateTaskFromRestTask) GetParameters() []Parameter {
	return m.Parameters
}

//GetOpConfigValues returns OpConfigValues
func (m UpdateTaskFromRestTask) GetOpConfigValues() *ConfigValues {
	return m.OpConfigValues
}

//GetConfigProviderDelegate returns ConfigProviderDelegate
func (m UpdateTaskFromRestTask) GetConfigProviderDelegate() *ConfigProvider {
	return m.ConfigProviderDelegate
}

//GetRegistryMetadata returns RegistryMetadata
func (m UpdateTaskFromRestTask) GetRegistryMetadata() *RegistryMetadata {
	return m.RegistryMetadata
}

func (m UpdateTaskFromRestTask) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m UpdateTaskFromRestTask) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateTaskFromRestTask UpdateTaskFromRestTask
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeUpdateTaskFromRestTask
	}{
		"REST_TASK",
		(MarshalTypeUpdateTaskFromRestTask)(m),
	}

	return json.Marshal(&s)
}

// UpdateTaskFromRestTaskMethodTypeEnum Enum with underlying type: string
type UpdateTaskFromRestTaskMethodTypeEnum string

// Set of constants representing the allowable values for UpdateTaskFromRestTaskMethodTypeEnum
const (
	UpdateTaskFromRestTaskMethodTypeGet    UpdateTaskFromRestTaskMethodTypeEnum = "GET"
	UpdateTaskFromRestTaskMethodTypePost   UpdateTaskFromRestTaskMethodTypeEnum = "POST"
	UpdateTaskFromRestTaskMethodTypePatch  UpdateTaskFromRestTaskMethodTypeEnum = "PATCH"
	UpdateTaskFromRestTaskMethodTypeDelete UpdateTaskFromRestTaskMethodTypeEnum = "DELETE"
	UpdateTaskFromRestTaskMethodTypePut    UpdateTaskFromRestTaskMethodTypeEnum = "PUT"
)

var mappingUpdateTaskFromRestTaskMethodType = map[string]UpdateTaskFromRestTaskMethodTypeEnum{
	"GET":    UpdateTaskFromRestTaskMethodTypeGet,
	"POST":   UpdateTaskFromRestTaskMethodTypePost,
	"PATCH":  UpdateTaskFromRestTaskMethodTypePatch,
	"DELETE": UpdateTaskFromRestTaskMethodTypeDelete,
	"PUT":    UpdateTaskFromRestTaskMethodTypePut,
}

// GetUpdateTaskFromRestTaskMethodTypeEnumValues Enumerates the set of values for UpdateTaskFromRestTaskMethodTypeEnum
func GetUpdateTaskFromRestTaskMethodTypeEnumValues() []UpdateTaskFromRestTaskMethodTypeEnum {
	values := make([]UpdateTaskFromRestTaskMethodTypeEnum, 0)
	for _, v := range mappingUpdateTaskFromRestTaskMethodType {
		values = append(values, v)
	}
	return values
}

// UpdateTaskFromRestTaskApiCallModeEnum Enum with underlying type: string
type UpdateTaskFromRestTaskApiCallModeEnum string

// Set of constants representing the allowable values for UpdateTaskFromRestTaskApiCallModeEnum
const (
	UpdateTaskFromRestTaskApiCallModeSynchronous         UpdateTaskFromRestTaskApiCallModeEnum = "SYNCHRONOUS"
	UpdateTaskFromRestTaskApiCallModeAsyncOciWorkrequest UpdateTaskFromRestTaskApiCallModeEnum = "ASYNC_OCI_WORKREQUEST"
)

var mappingUpdateTaskFromRestTaskApiCallMode = map[string]UpdateTaskFromRestTaskApiCallModeEnum{
	"SYNCHRONOUS":           UpdateTaskFromRestTaskApiCallModeSynchronous,
	"ASYNC_OCI_WORKREQUEST": UpdateTaskFromRestTaskApiCallModeAsyncOciWorkrequest,
}

// GetUpdateTaskFromRestTaskApiCallModeEnumValues Enumerates the set of values for UpdateTaskFromRestTaskApiCallModeEnum
func GetUpdateTaskFromRestTaskApiCallModeEnumValues() []UpdateTaskFromRestTaskApiCallModeEnum {
	values := make([]UpdateTaskFromRestTaskApiCallModeEnum, 0)
	for _, v := range mappingUpdateTaskFromRestTaskApiCallMode {
		values = append(values, v)
	}
	return values
}

// UpdateTaskFromRestTaskCancelMethodTypeEnum Enum with underlying type: string
type UpdateTaskFromRestTaskCancelMethodTypeEnum string

// Set of constants representing the allowable values for UpdateTaskFromRestTaskCancelMethodTypeEnum
const (
	UpdateTaskFromRestTaskCancelMethodTypeGet    UpdateTaskFromRestTaskCancelMethodTypeEnum = "GET"
	UpdateTaskFromRestTaskCancelMethodTypePost   UpdateTaskFromRestTaskCancelMethodTypeEnum = "POST"
	UpdateTaskFromRestTaskCancelMethodTypePatch  UpdateTaskFromRestTaskCancelMethodTypeEnum = "PATCH"
	UpdateTaskFromRestTaskCancelMethodTypeDelete UpdateTaskFromRestTaskCancelMethodTypeEnum = "DELETE"
	UpdateTaskFromRestTaskCancelMethodTypePut    UpdateTaskFromRestTaskCancelMethodTypeEnum = "PUT"
)

var mappingUpdateTaskFromRestTaskCancelMethodType = map[string]UpdateTaskFromRestTaskCancelMethodTypeEnum{
	"GET":    UpdateTaskFromRestTaskCancelMethodTypeGet,
	"POST":   UpdateTaskFromRestTaskCancelMethodTypePost,
	"PATCH":  UpdateTaskFromRestTaskCancelMethodTypePatch,
	"DELETE": UpdateTaskFromRestTaskCancelMethodTypeDelete,
	"PUT":    UpdateTaskFromRestTaskCancelMethodTypePut,
}

// GetUpdateTaskFromRestTaskCancelMethodTypeEnumValues Enumerates the set of values for UpdateTaskFromRestTaskCancelMethodTypeEnum
func GetUpdateTaskFromRestTaskCancelMethodTypeEnumValues() []UpdateTaskFromRestTaskCancelMethodTypeEnum {
	values := make([]UpdateTaskFromRestTaskCancelMethodTypeEnum, 0)
	for _, v := range mappingUpdateTaskFromRestTaskCancelMethodType {
		values = append(values, v)
	}
	return values
}
