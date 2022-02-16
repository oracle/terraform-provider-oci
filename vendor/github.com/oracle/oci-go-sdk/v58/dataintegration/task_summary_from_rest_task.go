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

// TaskSummaryFromRestTask The information about the Generic REST task. The endpoint and cancelEndpoint  properties are deprecated, use the properties executeRestCallConfig, cancelRestCallConfig and pollRestCallConfig for execute, cancel and polling of the calls.
type TaskSummaryFromRestTask struct {

	// Generated key that can be used in API calls to identify task. On scenarios where reference to the task is needed, a value can be passed in create.
	Key *string `mandatory:"false" json:"key"`

	// The object's model version.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"false" json:"name"`

	// Detailed description for the object.
	Description *string `mandatory:"false" json:"description"`

	// The version of the object that is used to track changes in the object instance.
	ObjectVersion *int `mandatory:"false" json:"objectVersion"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	// Value can only contain upper case letters, underscore, and numbers. It should begin with upper case letter or underscore. The value can be modified.
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

	// A key map. If provided, key is replaced with generated key. This structure provides mapping between user provided key and generated key.
	KeyMap map[string]string `mandatory:"false" json:"keyMap"`

	AuthDetails *AuthDetails `mandatory:"false" json:"authDetails"`

	Endpoint *Expression `mandatory:"false" json:"endpoint"`

	Headers *interface{} `mandatory:"false" json:"headers"`

	// JSON data for payload body. This property is deprecated, use ExecuteRestCallConfig's payload config param instead.
	JsonData *string `mandatory:"false" json:"jsonData"`

	CancelEndpoint *Expression `mandatory:"false" json:"cancelEndpoint"`

	ExecuteRestCallConfig *ExecuteRestCallConfig `mandatory:"false" json:"executeRestCallConfig"`

	CancelRestCallConfig *CancelRestCallConfig `mandatory:"false" json:"cancelRestCallConfig"`

	// The REST method to use. This property is deprecated, use ExecuteRestCallConfig's methodType property instead.
	MethodType TaskSummaryFromRestTaskMethodTypeEnum `mandatory:"false" json:"methodType,omitempty"`

	// The REST invocation pattern to use. ASYNC_OCI_WORKREQUEST is being deprecated as well as cancelEndpoint/MethodType.
	ApiCallMode TaskSummaryFromRestTaskApiCallModeEnum `mandatory:"false" json:"apiCallMode,omitempty"`

	// The REST method to use for canceling the original request.
	CancelMethodType TaskSummaryFromRestTaskCancelMethodTypeEnum `mandatory:"false" json:"cancelMethodType,omitempty"`
}

//GetKey returns Key
func (m TaskSummaryFromRestTask) GetKey() *string {
	return m.Key
}

//GetModelVersion returns ModelVersion
func (m TaskSummaryFromRestTask) GetModelVersion() *string {
	return m.ModelVersion
}

//GetParentRef returns ParentRef
func (m TaskSummaryFromRestTask) GetParentRef() *ParentReference {
	return m.ParentRef
}

//GetName returns Name
func (m TaskSummaryFromRestTask) GetName() *string {
	return m.Name
}

//GetDescription returns Description
func (m TaskSummaryFromRestTask) GetDescription() *string {
	return m.Description
}

//GetObjectVersion returns ObjectVersion
func (m TaskSummaryFromRestTask) GetObjectVersion() *int {
	return m.ObjectVersion
}

//GetObjectStatus returns ObjectStatus
func (m TaskSummaryFromRestTask) GetObjectStatus() *int {
	return m.ObjectStatus
}

//GetIdentifier returns Identifier
func (m TaskSummaryFromRestTask) GetIdentifier() *string {
	return m.Identifier
}

//GetInputPorts returns InputPorts
func (m TaskSummaryFromRestTask) GetInputPorts() []InputPort {
	return m.InputPorts
}

//GetOutputPorts returns OutputPorts
func (m TaskSummaryFromRestTask) GetOutputPorts() []OutputPort {
	return m.OutputPorts
}

//GetParameters returns Parameters
func (m TaskSummaryFromRestTask) GetParameters() []Parameter {
	return m.Parameters
}

//GetOpConfigValues returns OpConfigValues
func (m TaskSummaryFromRestTask) GetOpConfigValues() *ConfigValues {
	return m.OpConfigValues
}

//GetConfigProviderDelegate returns ConfigProviderDelegate
func (m TaskSummaryFromRestTask) GetConfigProviderDelegate() *ConfigProvider {
	return m.ConfigProviderDelegate
}

//GetMetadata returns Metadata
func (m TaskSummaryFromRestTask) GetMetadata() *ObjectMetadata {
	return m.Metadata
}

//GetKeyMap returns KeyMap
func (m TaskSummaryFromRestTask) GetKeyMap() map[string]string {
	return m.KeyMap
}

func (m TaskSummaryFromRestTask) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TaskSummaryFromRestTask) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingTaskSummaryFromRestTaskMethodTypeEnum(string(m.MethodType)); !ok && m.MethodType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MethodType: %s. Supported values are: %s.", m.MethodType, strings.Join(GetTaskSummaryFromRestTaskMethodTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTaskSummaryFromRestTaskApiCallModeEnum(string(m.ApiCallMode)); !ok && m.ApiCallMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ApiCallMode: %s. Supported values are: %s.", m.ApiCallMode, strings.Join(GetTaskSummaryFromRestTaskApiCallModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTaskSummaryFromRestTaskCancelMethodTypeEnum(string(m.CancelMethodType)); !ok && m.CancelMethodType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CancelMethodType: %s. Supported values are: %s.", m.CancelMethodType, strings.Join(GetTaskSummaryFromRestTaskCancelMethodTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m TaskSummaryFromRestTask) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeTaskSummaryFromRestTask TaskSummaryFromRestTask
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeTaskSummaryFromRestTask
	}{
		"REST_TASK",
		(MarshalTypeTaskSummaryFromRestTask)(m),
	}

	return json.Marshal(&s)
}

// TaskSummaryFromRestTaskMethodTypeEnum Enum with underlying type: string
type TaskSummaryFromRestTaskMethodTypeEnum string

// Set of constants representing the allowable values for TaskSummaryFromRestTaskMethodTypeEnum
const (
	TaskSummaryFromRestTaskMethodTypeGet    TaskSummaryFromRestTaskMethodTypeEnum = "GET"
	TaskSummaryFromRestTaskMethodTypePost   TaskSummaryFromRestTaskMethodTypeEnum = "POST"
	TaskSummaryFromRestTaskMethodTypePatch  TaskSummaryFromRestTaskMethodTypeEnum = "PATCH"
	TaskSummaryFromRestTaskMethodTypeDelete TaskSummaryFromRestTaskMethodTypeEnum = "DELETE"
	TaskSummaryFromRestTaskMethodTypePut    TaskSummaryFromRestTaskMethodTypeEnum = "PUT"
)

var mappingTaskSummaryFromRestTaskMethodTypeEnum = map[string]TaskSummaryFromRestTaskMethodTypeEnum{
	"GET":    TaskSummaryFromRestTaskMethodTypeGet,
	"POST":   TaskSummaryFromRestTaskMethodTypePost,
	"PATCH":  TaskSummaryFromRestTaskMethodTypePatch,
	"DELETE": TaskSummaryFromRestTaskMethodTypeDelete,
	"PUT":    TaskSummaryFromRestTaskMethodTypePut,
}

// GetTaskSummaryFromRestTaskMethodTypeEnumValues Enumerates the set of values for TaskSummaryFromRestTaskMethodTypeEnum
func GetTaskSummaryFromRestTaskMethodTypeEnumValues() []TaskSummaryFromRestTaskMethodTypeEnum {
	values := make([]TaskSummaryFromRestTaskMethodTypeEnum, 0)
	for _, v := range mappingTaskSummaryFromRestTaskMethodTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTaskSummaryFromRestTaskMethodTypeEnumStringValues Enumerates the set of values in String for TaskSummaryFromRestTaskMethodTypeEnum
func GetTaskSummaryFromRestTaskMethodTypeEnumStringValues() []string {
	return []string{
		"GET",
		"POST",
		"PATCH",
		"DELETE",
		"PUT",
	}
}

// GetMappingTaskSummaryFromRestTaskMethodTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTaskSummaryFromRestTaskMethodTypeEnum(val string) (TaskSummaryFromRestTaskMethodTypeEnum, bool) {
	mappingTaskSummaryFromRestTaskMethodTypeEnumIgnoreCase := make(map[string]TaskSummaryFromRestTaskMethodTypeEnum)
	for k, v := range mappingTaskSummaryFromRestTaskMethodTypeEnum {
		mappingTaskSummaryFromRestTaskMethodTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingTaskSummaryFromRestTaskMethodTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// TaskSummaryFromRestTaskApiCallModeEnum Enum with underlying type: string
type TaskSummaryFromRestTaskApiCallModeEnum string

// Set of constants representing the allowable values for TaskSummaryFromRestTaskApiCallModeEnum
const (
	TaskSummaryFromRestTaskApiCallModeSynchronous         TaskSummaryFromRestTaskApiCallModeEnum = "SYNCHRONOUS"
	TaskSummaryFromRestTaskApiCallModeAsyncOciWorkrequest TaskSummaryFromRestTaskApiCallModeEnum = "ASYNC_OCI_WORKREQUEST"
	TaskSummaryFromRestTaskApiCallModeAsyncGeneric        TaskSummaryFromRestTaskApiCallModeEnum = "ASYNC_GENERIC"
)

var mappingTaskSummaryFromRestTaskApiCallModeEnum = map[string]TaskSummaryFromRestTaskApiCallModeEnum{
	"SYNCHRONOUS":           TaskSummaryFromRestTaskApiCallModeSynchronous,
	"ASYNC_OCI_WORKREQUEST": TaskSummaryFromRestTaskApiCallModeAsyncOciWorkrequest,
	"ASYNC_GENERIC":         TaskSummaryFromRestTaskApiCallModeAsyncGeneric,
}

// GetTaskSummaryFromRestTaskApiCallModeEnumValues Enumerates the set of values for TaskSummaryFromRestTaskApiCallModeEnum
func GetTaskSummaryFromRestTaskApiCallModeEnumValues() []TaskSummaryFromRestTaskApiCallModeEnum {
	values := make([]TaskSummaryFromRestTaskApiCallModeEnum, 0)
	for _, v := range mappingTaskSummaryFromRestTaskApiCallModeEnum {
		values = append(values, v)
	}
	return values
}

// GetTaskSummaryFromRestTaskApiCallModeEnumStringValues Enumerates the set of values in String for TaskSummaryFromRestTaskApiCallModeEnum
func GetTaskSummaryFromRestTaskApiCallModeEnumStringValues() []string {
	return []string{
		"SYNCHRONOUS",
		"ASYNC_OCI_WORKREQUEST",
		"ASYNC_GENERIC",
	}
}

// GetMappingTaskSummaryFromRestTaskApiCallModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTaskSummaryFromRestTaskApiCallModeEnum(val string) (TaskSummaryFromRestTaskApiCallModeEnum, bool) {
	mappingTaskSummaryFromRestTaskApiCallModeEnumIgnoreCase := make(map[string]TaskSummaryFromRestTaskApiCallModeEnum)
	for k, v := range mappingTaskSummaryFromRestTaskApiCallModeEnum {
		mappingTaskSummaryFromRestTaskApiCallModeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingTaskSummaryFromRestTaskApiCallModeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// TaskSummaryFromRestTaskCancelMethodTypeEnum Enum with underlying type: string
type TaskSummaryFromRestTaskCancelMethodTypeEnum string

// Set of constants representing the allowable values for TaskSummaryFromRestTaskCancelMethodTypeEnum
const (
	TaskSummaryFromRestTaskCancelMethodTypeGet    TaskSummaryFromRestTaskCancelMethodTypeEnum = "GET"
	TaskSummaryFromRestTaskCancelMethodTypePost   TaskSummaryFromRestTaskCancelMethodTypeEnum = "POST"
	TaskSummaryFromRestTaskCancelMethodTypePatch  TaskSummaryFromRestTaskCancelMethodTypeEnum = "PATCH"
	TaskSummaryFromRestTaskCancelMethodTypeDelete TaskSummaryFromRestTaskCancelMethodTypeEnum = "DELETE"
	TaskSummaryFromRestTaskCancelMethodTypePut    TaskSummaryFromRestTaskCancelMethodTypeEnum = "PUT"
)

var mappingTaskSummaryFromRestTaskCancelMethodTypeEnum = map[string]TaskSummaryFromRestTaskCancelMethodTypeEnum{
	"GET":    TaskSummaryFromRestTaskCancelMethodTypeGet,
	"POST":   TaskSummaryFromRestTaskCancelMethodTypePost,
	"PATCH":  TaskSummaryFromRestTaskCancelMethodTypePatch,
	"DELETE": TaskSummaryFromRestTaskCancelMethodTypeDelete,
	"PUT":    TaskSummaryFromRestTaskCancelMethodTypePut,
}

// GetTaskSummaryFromRestTaskCancelMethodTypeEnumValues Enumerates the set of values for TaskSummaryFromRestTaskCancelMethodTypeEnum
func GetTaskSummaryFromRestTaskCancelMethodTypeEnumValues() []TaskSummaryFromRestTaskCancelMethodTypeEnum {
	values := make([]TaskSummaryFromRestTaskCancelMethodTypeEnum, 0)
	for _, v := range mappingTaskSummaryFromRestTaskCancelMethodTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTaskSummaryFromRestTaskCancelMethodTypeEnumStringValues Enumerates the set of values in String for TaskSummaryFromRestTaskCancelMethodTypeEnum
func GetTaskSummaryFromRestTaskCancelMethodTypeEnumStringValues() []string {
	return []string{
		"GET",
		"POST",
		"PATCH",
		"DELETE",
		"PUT",
	}
}

// GetMappingTaskSummaryFromRestTaskCancelMethodTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTaskSummaryFromRestTaskCancelMethodTypeEnum(val string) (TaskSummaryFromRestTaskCancelMethodTypeEnum, bool) {
	mappingTaskSummaryFromRestTaskCancelMethodTypeEnumIgnoreCase := make(map[string]TaskSummaryFromRestTaskCancelMethodTypeEnum)
	for k, v := range mappingTaskSummaryFromRestTaskCancelMethodTypeEnum {
		mappingTaskSummaryFromRestTaskCancelMethodTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingTaskSummaryFromRestTaskCancelMethodTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
