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

// TaskFromRestTaskDetails The information about the Generic REST task. The endpoint and cancelEndpoint  properties are deprecated, use the properties executeRestCallConfig, cancelRestCallConfig and pollRestCallConfig for execute, cancel and polling of the calls.
type TaskFromRestTaskDetails struct {

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

	// Whether the same task can be executed concurrently.
	IsConcurrentAllowed *bool `mandatory:"false" json:"isConcurrentAllowed"`

	Metadata *ObjectMetadata `mandatory:"false" json:"metadata"`

	// A key map. If provided, key is replaced with generated key. This structure provides mapping between user provided key and generated key.
	KeyMap map[string]string `mandatory:"false" json:"keyMap"`

	RegistryMetadata *RegistryMetadata `mandatory:"false" json:"registryMetadata"`

	AuthDetails *AuthDetails `mandatory:"false" json:"authDetails"`

	AuthConfig AuthConfig `mandatory:"false" json:"authConfig"`

	Endpoint *Expression `mandatory:"false" json:"endpoint"`

	// The headers for the REST call. This property is deprecated, use ExecuteRestCallConfig's headers property instead.
	Headers *interface{} `mandatory:"false" json:"headers"`

	// JSON data for payload body. This property is deprecated, use ExecuteRestCallConfig's payload config param instead.
	JsonData *string `mandatory:"false" json:"jsonData"`

	CancelEndpoint *Expression `mandatory:"false" json:"cancelEndpoint"`

	ExecuteRestCallConfig *ExecuteRestCallConfig `mandatory:"false" json:"executeRestCallConfig"`

	CancelRestCallConfig *CancelRestCallConfig `mandatory:"false" json:"cancelRestCallConfig"`

	PollRestCallConfig *PollRestCallConfig `mandatory:"false" json:"pollRestCallConfig"`

	// List of typed expressions.
	TypedExpressions []TypedExpression `mandatory:"false" json:"typedExpressions"`

	// The REST method to use. This property is deprecated, use ExecuteRestCallConfig's methodType property instead.
	MethodType TaskFromRestTaskDetailsMethodTypeEnum `mandatory:"false" json:"methodType,omitempty"`

	// The REST invocation pattern to use. ASYNC_OCI_WORKREQUEST is being deprecated as well as cancelEndpoint/MethodType.
	ApiCallMode TaskFromRestTaskDetailsApiCallModeEnum `mandatory:"false" json:"apiCallMode,omitempty"`

	// The REST method to use for canceling the original request.
	CancelMethodType TaskFromRestTaskDetailsCancelMethodTypeEnum `mandatory:"false" json:"cancelMethodType,omitempty"`
}

// GetKey returns Key
func (m TaskFromRestTaskDetails) GetKey() *string {
	return m.Key
}

// GetModelVersion returns ModelVersion
func (m TaskFromRestTaskDetails) GetModelVersion() *string {
	return m.ModelVersion
}

// GetParentRef returns ParentRef
func (m TaskFromRestTaskDetails) GetParentRef() *ParentReference {
	return m.ParentRef
}

// GetName returns Name
func (m TaskFromRestTaskDetails) GetName() *string {
	return m.Name
}

// GetDescription returns Description
func (m TaskFromRestTaskDetails) GetDescription() *string {
	return m.Description
}

// GetObjectVersion returns ObjectVersion
func (m TaskFromRestTaskDetails) GetObjectVersion() *int {
	return m.ObjectVersion
}

// GetObjectStatus returns ObjectStatus
func (m TaskFromRestTaskDetails) GetObjectStatus() *int {
	return m.ObjectStatus
}

// GetIdentifier returns Identifier
func (m TaskFromRestTaskDetails) GetIdentifier() *string {
	return m.Identifier
}

// GetInputPorts returns InputPorts
func (m TaskFromRestTaskDetails) GetInputPorts() []InputPort {
	return m.InputPorts
}

// GetOutputPorts returns OutputPorts
func (m TaskFromRestTaskDetails) GetOutputPorts() []OutputPort {
	return m.OutputPorts
}

// GetParameters returns Parameters
func (m TaskFromRestTaskDetails) GetParameters() []Parameter {
	return m.Parameters
}

// GetOpConfigValues returns OpConfigValues
func (m TaskFromRestTaskDetails) GetOpConfigValues() *ConfigValues {
	return m.OpConfigValues
}

// GetConfigProviderDelegate returns ConfigProviderDelegate
func (m TaskFromRestTaskDetails) GetConfigProviderDelegate() *ConfigProvider {
	return m.ConfigProviderDelegate
}

// GetIsConcurrentAllowed returns IsConcurrentAllowed
func (m TaskFromRestTaskDetails) GetIsConcurrentAllowed() *bool {
	return m.IsConcurrentAllowed
}

// GetMetadata returns Metadata
func (m TaskFromRestTaskDetails) GetMetadata() *ObjectMetadata {
	return m.Metadata
}

// GetKeyMap returns KeyMap
func (m TaskFromRestTaskDetails) GetKeyMap() map[string]string {
	return m.KeyMap
}

// GetRegistryMetadata returns RegistryMetadata
func (m TaskFromRestTaskDetails) GetRegistryMetadata() *RegistryMetadata {
	return m.RegistryMetadata
}

func (m TaskFromRestTaskDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TaskFromRestTaskDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingTaskFromRestTaskDetailsMethodTypeEnum(string(m.MethodType)); !ok && m.MethodType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MethodType: %s. Supported values are: %s.", m.MethodType, strings.Join(GetTaskFromRestTaskDetailsMethodTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTaskFromRestTaskDetailsApiCallModeEnum(string(m.ApiCallMode)); !ok && m.ApiCallMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ApiCallMode: %s. Supported values are: %s.", m.ApiCallMode, strings.Join(GetTaskFromRestTaskDetailsApiCallModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTaskFromRestTaskDetailsCancelMethodTypeEnum(string(m.CancelMethodType)); !ok && m.CancelMethodType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CancelMethodType: %s. Supported values are: %s.", m.CancelMethodType, strings.Join(GetTaskFromRestTaskDetailsCancelMethodTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m TaskFromRestTaskDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeTaskFromRestTaskDetails TaskFromRestTaskDetails
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeTaskFromRestTaskDetails
	}{
		"REST_TASK",
		(MarshalTypeTaskFromRestTaskDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *TaskFromRestTaskDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Key                    *string                                     `json:"key"`
		ModelVersion           *string                                     `json:"modelVersion"`
		ParentRef              *ParentReference                            `json:"parentRef"`
		Name                   *string                                     `json:"name"`
		Description            *string                                     `json:"description"`
		ObjectVersion          *int                                        `json:"objectVersion"`
		ObjectStatus           *int                                        `json:"objectStatus"`
		Identifier             *string                                     `json:"identifier"`
		InputPorts             []InputPort                                 `json:"inputPorts"`
		OutputPorts            []OutputPort                                `json:"outputPorts"`
		Parameters             []Parameter                                 `json:"parameters"`
		OpConfigValues         *ConfigValues                               `json:"opConfigValues"`
		ConfigProviderDelegate *ConfigProvider                             `json:"configProviderDelegate"`
		IsConcurrentAllowed    *bool                                       `json:"isConcurrentAllowed"`
		Metadata               *ObjectMetadata                             `json:"metadata"`
		KeyMap                 map[string]string                           `json:"keyMap"`
		RegistryMetadata       *RegistryMetadata                           `json:"registryMetadata"`
		AuthDetails            *AuthDetails                                `json:"authDetails"`
		AuthConfig             authconfig                                  `json:"authConfig"`
		Endpoint               *Expression                                 `json:"endpoint"`
		MethodType             TaskFromRestTaskDetailsMethodTypeEnum       `json:"methodType"`
		Headers                *interface{}                                `json:"headers"`
		JsonData               *string                                     `json:"jsonData"`
		ApiCallMode            TaskFromRestTaskDetailsApiCallModeEnum      `json:"apiCallMode"`
		CancelEndpoint         *Expression                                 `json:"cancelEndpoint"`
		CancelMethodType       TaskFromRestTaskDetailsCancelMethodTypeEnum `json:"cancelMethodType"`
		ExecuteRestCallConfig  *ExecuteRestCallConfig                      `json:"executeRestCallConfig"`
		CancelRestCallConfig   *CancelRestCallConfig                       `json:"cancelRestCallConfig"`
		PollRestCallConfig     *PollRestCallConfig                         `json:"pollRestCallConfig"`
		TypedExpressions       []TypedExpression                           `json:"typedExpressions"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Key = model.Key

	m.ModelVersion = model.ModelVersion

	m.ParentRef = model.ParentRef

	m.Name = model.Name

	m.Description = model.Description

	m.ObjectVersion = model.ObjectVersion

	m.ObjectStatus = model.ObjectStatus

	m.Identifier = model.Identifier

	m.InputPorts = make([]InputPort, len(model.InputPorts))
	copy(m.InputPorts, model.InputPorts)
	m.OutputPorts = make([]OutputPort, len(model.OutputPorts))
	copy(m.OutputPorts, model.OutputPorts)
	m.Parameters = make([]Parameter, len(model.Parameters))
	copy(m.Parameters, model.Parameters)
	m.OpConfigValues = model.OpConfigValues

	m.ConfigProviderDelegate = model.ConfigProviderDelegate

	m.IsConcurrentAllowed = model.IsConcurrentAllowed

	m.Metadata = model.Metadata

	m.KeyMap = model.KeyMap

	m.RegistryMetadata = model.RegistryMetadata

	m.AuthDetails = model.AuthDetails

	nn, e = model.AuthConfig.UnmarshalPolymorphicJSON(model.AuthConfig.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.AuthConfig = nn.(AuthConfig)
	} else {
		m.AuthConfig = nil
	}

	m.Endpoint = model.Endpoint

	m.MethodType = model.MethodType

	m.Headers = model.Headers

	m.JsonData = model.JsonData

	m.ApiCallMode = model.ApiCallMode

	m.CancelEndpoint = model.CancelEndpoint

	m.CancelMethodType = model.CancelMethodType

	m.ExecuteRestCallConfig = model.ExecuteRestCallConfig

	m.CancelRestCallConfig = model.CancelRestCallConfig

	m.PollRestCallConfig = model.PollRestCallConfig

	m.TypedExpressions = make([]TypedExpression, len(model.TypedExpressions))
	copy(m.TypedExpressions, model.TypedExpressions)
	return
}

// TaskFromRestTaskDetailsMethodTypeEnum Enum with underlying type: string
type TaskFromRestTaskDetailsMethodTypeEnum string

// Set of constants representing the allowable values for TaskFromRestTaskDetailsMethodTypeEnum
const (
	TaskFromRestTaskDetailsMethodTypeGet    TaskFromRestTaskDetailsMethodTypeEnum = "GET"
	TaskFromRestTaskDetailsMethodTypePost   TaskFromRestTaskDetailsMethodTypeEnum = "POST"
	TaskFromRestTaskDetailsMethodTypePatch  TaskFromRestTaskDetailsMethodTypeEnum = "PATCH"
	TaskFromRestTaskDetailsMethodTypeDelete TaskFromRestTaskDetailsMethodTypeEnum = "DELETE"
	TaskFromRestTaskDetailsMethodTypePut    TaskFromRestTaskDetailsMethodTypeEnum = "PUT"
)

var mappingTaskFromRestTaskDetailsMethodTypeEnum = map[string]TaskFromRestTaskDetailsMethodTypeEnum{
	"GET":    TaskFromRestTaskDetailsMethodTypeGet,
	"POST":   TaskFromRestTaskDetailsMethodTypePost,
	"PATCH":  TaskFromRestTaskDetailsMethodTypePatch,
	"DELETE": TaskFromRestTaskDetailsMethodTypeDelete,
	"PUT":    TaskFromRestTaskDetailsMethodTypePut,
}

var mappingTaskFromRestTaskDetailsMethodTypeEnumLowerCase = map[string]TaskFromRestTaskDetailsMethodTypeEnum{
	"get":    TaskFromRestTaskDetailsMethodTypeGet,
	"post":   TaskFromRestTaskDetailsMethodTypePost,
	"patch":  TaskFromRestTaskDetailsMethodTypePatch,
	"delete": TaskFromRestTaskDetailsMethodTypeDelete,
	"put":    TaskFromRestTaskDetailsMethodTypePut,
}

// GetTaskFromRestTaskDetailsMethodTypeEnumValues Enumerates the set of values for TaskFromRestTaskDetailsMethodTypeEnum
func GetTaskFromRestTaskDetailsMethodTypeEnumValues() []TaskFromRestTaskDetailsMethodTypeEnum {
	values := make([]TaskFromRestTaskDetailsMethodTypeEnum, 0)
	for _, v := range mappingTaskFromRestTaskDetailsMethodTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTaskFromRestTaskDetailsMethodTypeEnumStringValues Enumerates the set of values in String for TaskFromRestTaskDetailsMethodTypeEnum
func GetTaskFromRestTaskDetailsMethodTypeEnumStringValues() []string {
	return []string{
		"GET",
		"POST",
		"PATCH",
		"DELETE",
		"PUT",
	}
}

// GetMappingTaskFromRestTaskDetailsMethodTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTaskFromRestTaskDetailsMethodTypeEnum(val string) (TaskFromRestTaskDetailsMethodTypeEnum, bool) {
	enum, ok := mappingTaskFromRestTaskDetailsMethodTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// TaskFromRestTaskDetailsApiCallModeEnum Enum with underlying type: string
type TaskFromRestTaskDetailsApiCallModeEnum string

// Set of constants representing the allowable values for TaskFromRestTaskDetailsApiCallModeEnum
const (
	TaskFromRestTaskDetailsApiCallModeSynchronous         TaskFromRestTaskDetailsApiCallModeEnum = "SYNCHRONOUS"
	TaskFromRestTaskDetailsApiCallModeAsyncOciWorkrequest TaskFromRestTaskDetailsApiCallModeEnum = "ASYNC_OCI_WORKREQUEST"
	TaskFromRestTaskDetailsApiCallModeAsyncGeneric        TaskFromRestTaskDetailsApiCallModeEnum = "ASYNC_GENERIC"
)

var mappingTaskFromRestTaskDetailsApiCallModeEnum = map[string]TaskFromRestTaskDetailsApiCallModeEnum{
	"SYNCHRONOUS":           TaskFromRestTaskDetailsApiCallModeSynchronous,
	"ASYNC_OCI_WORKREQUEST": TaskFromRestTaskDetailsApiCallModeAsyncOciWorkrequest,
	"ASYNC_GENERIC":         TaskFromRestTaskDetailsApiCallModeAsyncGeneric,
}

var mappingTaskFromRestTaskDetailsApiCallModeEnumLowerCase = map[string]TaskFromRestTaskDetailsApiCallModeEnum{
	"synchronous":           TaskFromRestTaskDetailsApiCallModeSynchronous,
	"async_oci_workrequest": TaskFromRestTaskDetailsApiCallModeAsyncOciWorkrequest,
	"async_generic":         TaskFromRestTaskDetailsApiCallModeAsyncGeneric,
}

// GetTaskFromRestTaskDetailsApiCallModeEnumValues Enumerates the set of values for TaskFromRestTaskDetailsApiCallModeEnum
func GetTaskFromRestTaskDetailsApiCallModeEnumValues() []TaskFromRestTaskDetailsApiCallModeEnum {
	values := make([]TaskFromRestTaskDetailsApiCallModeEnum, 0)
	for _, v := range mappingTaskFromRestTaskDetailsApiCallModeEnum {
		values = append(values, v)
	}
	return values
}

// GetTaskFromRestTaskDetailsApiCallModeEnumStringValues Enumerates the set of values in String for TaskFromRestTaskDetailsApiCallModeEnum
func GetTaskFromRestTaskDetailsApiCallModeEnumStringValues() []string {
	return []string{
		"SYNCHRONOUS",
		"ASYNC_OCI_WORKREQUEST",
		"ASYNC_GENERIC",
	}
}

// GetMappingTaskFromRestTaskDetailsApiCallModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTaskFromRestTaskDetailsApiCallModeEnum(val string) (TaskFromRestTaskDetailsApiCallModeEnum, bool) {
	enum, ok := mappingTaskFromRestTaskDetailsApiCallModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// TaskFromRestTaskDetailsCancelMethodTypeEnum Enum with underlying type: string
type TaskFromRestTaskDetailsCancelMethodTypeEnum string

// Set of constants representing the allowable values for TaskFromRestTaskDetailsCancelMethodTypeEnum
const (
	TaskFromRestTaskDetailsCancelMethodTypeGet    TaskFromRestTaskDetailsCancelMethodTypeEnum = "GET"
	TaskFromRestTaskDetailsCancelMethodTypePost   TaskFromRestTaskDetailsCancelMethodTypeEnum = "POST"
	TaskFromRestTaskDetailsCancelMethodTypePatch  TaskFromRestTaskDetailsCancelMethodTypeEnum = "PATCH"
	TaskFromRestTaskDetailsCancelMethodTypeDelete TaskFromRestTaskDetailsCancelMethodTypeEnum = "DELETE"
	TaskFromRestTaskDetailsCancelMethodTypePut    TaskFromRestTaskDetailsCancelMethodTypeEnum = "PUT"
)

var mappingTaskFromRestTaskDetailsCancelMethodTypeEnum = map[string]TaskFromRestTaskDetailsCancelMethodTypeEnum{
	"GET":    TaskFromRestTaskDetailsCancelMethodTypeGet,
	"POST":   TaskFromRestTaskDetailsCancelMethodTypePost,
	"PATCH":  TaskFromRestTaskDetailsCancelMethodTypePatch,
	"DELETE": TaskFromRestTaskDetailsCancelMethodTypeDelete,
	"PUT":    TaskFromRestTaskDetailsCancelMethodTypePut,
}

var mappingTaskFromRestTaskDetailsCancelMethodTypeEnumLowerCase = map[string]TaskFromRestTaskDetailsCancelMethodTypeEnum{
	"get":    TaskFromRestTaskDetailsCancelMethodTypeGet,
	"post":   TaskFromRestTaskDetailsCancelMethodTypePost,
	"patch":  TaskFromRestTaskDetailsCancelMethodTypePatch,
	"delete": TaskFromRestTaskDetailsCancelMethodTypeDelete,
	"put":    TaskFromRestTaskDetailsCancelMethodTypePut,
}

// GetTaskFromRestTaskDetailsCancelMethodTypeEnumValues Enumerates the set of values for TaskFromRestTaskDetailsCancelMethodTypeEnum
func GetTaskFromRestTaskDetailsCancelMethodTypeEnumValues() []TaskFromRestTaskDetailsCancelMethodTypeEnum {
	values := make([]TaskFromRestTaskDetailsCancelMethodTypeEnum, 0)
	for _, v := range mappingTaskFromRestTaskDetailsCancelMethodTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTaskFromRestTaskDetailsCancelMethodTypeEnumStringValues Enumerates the set of values in String for TaskFromRestTaskDetailsCancelMethodTypeEnum
func GetTaskFromRestTaskDetailsCancelMethodTypeEnumStringValues() []string {
	return []string{
		"GET",
		"POST",
		"PATCH",
		"DELETE",
		"PUT",
	}
}

// GetMappingTaskFromRestTaskDetailsCancelMethodTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTaskFromRestTaskDetailsCancelMethodTypeEnum(val string) (TaskFromRestTaskDetailsCancelMethodTypeEnum, bool) {
	enum, ok := mappingTaskFromRestTaskDetailsCancelMethodTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
