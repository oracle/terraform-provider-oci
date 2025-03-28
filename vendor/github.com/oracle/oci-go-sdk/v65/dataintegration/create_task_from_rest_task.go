// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// CreateTaskFromRestTask The information about the Generic REST task. The endpoint and cancelEndpoint  properties are deprecated, use the properties executeRestCallConfig, cancelRestCallConfig and pollRestCallConfig for execute, cancel and polling of the calls.
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

	// Whether the same task can be executed concurrently.
	IsConcurrentAllowed *bool `mandatory:"false" json:"isConcurrentAllowed"`

	AuthDetails *AuthDetails `mandatory:"false" json:"authDetails"`

	AuthConfig AuthConfig `mandatory:"false" json:"authConfig"`

	Endpoint *Expression `mandatory:"false" json:"endpoint"`

	// Headers data for the request.
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
	MethodType CreateTaskFromRestTaskMethodTypeEnum `mandatory:"false" json:"methodType,omitempty"`

	// The REST invocation pattern to use. ASYNC_OCI_WORKREQUEST is being deprecated as well as cancelEndpoint/MethodType.
	ApiCallMode CreateTaskFromRestTaskApiCallModeEnum `mandatory:"false" json:"apiCallMode,omitempty"`

	// The REST method to use for canceling the original request.
	CancelMethodType CreateTaskFromRestTaskCancelMethodTypeEnum `mandatory:"false" json:"cancelMethodType,omitempty"`
}

// GetKey returns Key
func (m CreateTaskFromRestTask) GetKey() *string {
	return m.Key
}

// GetModelVersion returns ModelVersion
func (m CreateTaskFromRestTask) GetModelVersion() *string {
	return m.ModelVersion
}

// GetParentRef returns ParentRef
func (m CreateTaskFromRestTask) GetParentRef() *ParentReference {
	return m.ParentRef
}

// GetName returns Name
func (m CreateTaskFromRestTask) GetName() *string {
	return m.Name
}

// GetDescription returns Description
func (m CreateTaskFromRestTask) GetDescription() *string {
	return m.Description
}

// GetObjectStatus returns ObjectStatus
func (m CreateTaskFromRestTask) GetObjectStatus() *int {
	return m.ObjectStatus
}

// GetIdentifier returns Identifier
func (m CreateTaskFromRestTask) GetIdentifier() *string {
	return m.Identifier
}

// GetInputPorts returns InputPorts
func (m CreateTaskFromRestTask) GetInputPorts() []InputPort {
	return m.InputPorts
}

// GetOutputPorts returns OutputPorts
func (m CreateTaskFromRestTask) GetOutputPorts() []OutputPort {
	return m.OutputPorts
}

// GetParameters returns Parameters
func (m CreateTaskFromRestTask) GetParameters() []Parameter {
	return m.Parameters
}

// GetOpConfigValues returns OpConfigValues
func (m CreateTaskFromRestTask) GetOpConfigValues() *ConfigValues {
	return m.OpConfigValues
}

// GetConfigProviderDelegate returns ConfigProviderDelegate
func (m CreateTaskFromRestTask) GetConfigProviderDelegate() *CreateConfigProvider {
	return m.ConfigProviderDelegate
}

// GetIsConcurrentAllowed returns IsConcurrentAllowed
func (m CreateTaskFromRestTask) GetIsConcurrentAllowed() *bool {
	return m.IsConcurrentAllowed
}

// GetRegistryMetadata returns RegistryMetadata
func (m CreateTaskFromRestTask) GetRegistryMetadata() *RegistryMetadata {
	return m.RegistryMetadata
}

func (m CreateTaskFromRestTask) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateTaskFromRestTask) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateTaskFromRestTaskMethodTypeEnum(string(m.MethodType)); !ok && m.MethodType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MethodType: %s. Supported values are: %s.", m.MethodType, strings.Join(GetCreateTaskFromRestTaskMethodTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCreateTaskFromRestTaskApiCallModeEnum(string(m.ApiCallMode)); !ok && m.ApiCallMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ApiCallMode: %s. Supported values are: %s.", m.ApiCallMode, strings.Join(GetCreateTaskFromRestTaskApiCallModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCreateTaskFromRestTaskCancelMethodTypeEnum(string(m.CancelMethodType)); !ok && m.CancelMethodType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CancelMethodType: %s. Supported values are: %s.", m.CancelMethodType, strings.Join(GetCreateTaskFromRestTaskCancelMethodTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

// UnmarshalJSON unmarshals from json
func (m *CreateTaskFromRestTask) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Key                    *string                                    `json:"key"`
		ModelVersion           *string                                    `json:"modelVersion"`
		ParentRef              *ParentReference                           `json:"parentRef"`
		Description            *string                                    `json:"description"`
		ObjectStatus           *int                                       `json:"objectStatus"`
		InputPorts             []InputPort                                `json:"inputPorts"`
		OutputPorts            []OutputPort                               `json:"outputPorts"`
		Parameters             []Parameter                                `json:"parameters"`
		OpConfigValues         *ConfigValues                              `json:"opConfigValues"`
		ConfigProviderDelegate *CreateConfigProvider                      `json:"configProviderDelegate"`
		IsConcurrentAllowed    *bool                                      `json:"isConcurrentAllowed"`
		AuthDetails            *AuthDetails                               `json:"authDetails"`
		AuthConfig             authconfig                                 `json:"authConfig"`
		Endpoint               *Expression                                `json:"endpoint"`
		MethodType             CreateTaskFromRestTaskMethodTypeEnum       `json:"methodType"`
		Headers                *interface{}                               `json:"headers"`
		JsonData               *string                                    `json:"jsonData"`
		ApiCallMode            CreateTaskFromRestTaskApiCallModeEnum      `json:"apiCallMode"`
		CancelEndpoint         *Expression                                `json:"cancelEndpoint"`
		CancelMethodType       CreateTaskFromRestTaskCancelMethodTypeEnum `json:"cancelMethodType"`
		ExecuteRestCallConfig  *ExecuteRestCallConfig                     `json:"executeRestCallConfig"`
		CancelRestCallConfig   *CancelRestCallConfig                      `json:"cancelRestCallConfig"`
		PollRestCallConfig     *PollRestCallConfig                        `json:"pollRestCallConfig"`
		TypedExpressions       []TypedExpression                          `json:"typedExpressions"`
		Name                   *string                                    `json:"name"`
		Identifier             *string                                    `json:"identifier"`
		RegistryMetadata       *RegistryMetadata                          `json:"registryMetadata"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Key = model.Key

	m.ModelVersion = model.ModelVersion

	m.ParentRef = model.ParentRef

	m.Description = model.Description

	m.ObjectStatus = model.ObjectStatus

	m.InputPorts = make([]InputPort, len(model.InputPorts))
	copy(m.InputPorts, model.InputPorts)
	m.OutputPorts = make([]OutputPort, len(model.OutputPorts))
	copy(m.OutputPorts, model.OutputPorts)
	m.Parameters = make([]Parameter, len(model.Parameters))
	copy(m.Parameters, model.Parameters)
	m.OpConfigValues = model.OpConfigValues

	m.ConfigProviderDelegate = model.ConfigProviderDelegate

	m.IsConcurrentAllowed = model.IsConcurrentAllowed

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
	m.Name = model.Name

	m.Identifier = model.Identifier

	m.RegistryMetadata = model.RegistryMetadata

	return
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

var mappingCreateTaskFromRestTaskMethodTypeEnum = map[string]CreateTaskFromRestTaskMethodTypeEnum{
	"GET":    CreateTaskFromRestTaskMethodTypeGet,
	"POST":   CreateTaskFromRestTaskMethodTypePost,
	"PATCH":  CreateTaskFromRestTaskMethodTypePatch,
	"DELETE": CreateTaskFromRestTaskMethodTypeDelete,
	"PUT":    CreateTaskFromRestTaskMethodTypePut,
}

var mappingCreateTaskFromRestTaskMethodTypeEnumLowerCase = map[string]CreateTaskFromRestTaskMethodTypeEnum{
	"get":    CreateTaskFromRestTaskMethodTypeGet,
	"post":   CreateTaskFromRestTaskMethodTypePost,
	"patch":  CreateTaskFromRestTaskMethodTypePatch,
	"delete": CreateTaskFromRestTaskMethodTypeDelete,
	"put":    CreateTaskFromRestTaskMethodTypePut,
}

// GetCreateTaskFromRestTaskMethodTypeEnumValues Enumerates the set of values for CreateTaskFromRestTaskMethodTypeEnum
func GetCreateTaskFromRestTaskMethodTypeEnumValues() []CreateTaskFromRestTaskMethodTypeEnum {
	values := make([]CreateTaskFromRestTaskMethodTypeEnum, 0)
	for _, v := range mappingCreateTaskFromRestTaskMethodTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateTaskFromRestTaskMethodTypeEnumStringValues Enumerates the set of values in String for CreateTaskFromRestTaskMethodTypeEnum
func GetCreateTaskFromRestTaskMethodTypeEnumStringValues() []string {
	return []string{
		"GET",
		"POST",
		"PATCH",
		"DELETE",
		"PUT",
	}
}

// GetMappingCreateTaskFromRestTaskMethodTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateTaskFromRestTaskMethodTypeEnum(val string) (CreateTaskFromRestTaskMethodTypeEnum, bool) {
	enum, ok := mappingCreateTaskFromRestTaskMethodTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreateTaskFromRestTaskApiCallModeEnum Enum with underlying type: string
type CreateTaskFromRestTaskApiCallModeEnum string

// Set of constants representing the allowable values for CreateTaskFromRestTaskApiCallModeEnum
const (
	CreateTaskFromRestTaskApiCallModeSynchronous         CreateTaskFromRestTaskApiCallModeEnum = "SYNCHRONOUS"
	CreateTaskFromRestTaskApiCallModeAsyncOciWorkrequest CreateTaskFromRestTaskApiCallModeEnum = "ASYNC_OCI_WORKREQUEST"
	CreateTaskFromRestTaskApiCallModeAsyncGeneric        CreateTaskFromRestTaskApiCallModeEnum = "ASYNC_GENERIC"
)

var mappingCreateTaskFromRestTaskApiCallModeEnum = map[string]CreateTaskFromRestTaskApiCallModeEnum{
	"SYNCHRONOUS":           CreateTaskFromRestTaskApiCallModeSynchronous,
	"ASYNC_OCI_WORKREQUEST": CreateTaskFromRestTaskApiCallModeAsyncOciWorkrequest,
	"ASYNC_GENERIC":         CreateTaskFromRestTaskApiCallModeAsyncGeneric,
}

var mappingCreateTaskFromRestTaskApiCallModeEnumLowerCase = map[string]CreateTaskFromRestTaskApiCallModeEnum{
	"synchronous":           CreateTaskFromRestTaskApiCallModeSynchronous,
	"async_oci_workrequest": CreateTaskFromRestTaskApiCallModeAsyncOciWorkrequest,
	"async_generic":         CreateTaskFromRestTaskApiCallModeAsyncGeneric,
}

// GetCreateTaskFromRestTaskApiCallModeEnumValues Enumerates the set of values for CreateTaskFromRestTaskApiCallModeEnum
func GetCreateTaskFromRestTaskApiCallModeEnumValues() []CreateTaskFromRestTaskApiCallModeEnum {
	values := make([]CreateTaskFromRestTaskApiCallModeEnum, 0)
	for _, v := range mappingCreateTaskFromRestTaskApiCallModeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateTaskFromRestTaskApiCallModeEnumStringValues Enumerates the set of values in String for CreateTaskFromRestTaskApiCallModeEnum
func GetCreateTaskFromRestTaskApiCallModeEnumStringValues() []string {
	return []string{
		"SYNCHRONOUS",
		"ASYNC_OCI_WORKREQUEST",
		"ASYNC_GENERIC",
	}
}

// GetMappingCreateTaskFromRestTaskApiCallModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateTaskFromRestTaskApiCallModeEnum(val string) (CreateTaskFromRestTaskApiCallModeEnum, bool) {
	enum, ok := mappingCreateTaskFromRestTaskApiCallModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
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

var mappingCreateTaskFromRestTaskCancelMethodTypeEnum = map[string]CreateTaskFromRestTaskCancelMethodTypeEnum{
	"GET":    CreateTaskFromRestTaskCancelMethodTypeGet,
	"POST":   CreateTaskFromRestTaskCancelMethodTypePost,
	"PATCH":  CreateTaskFromRestTaskCancelMethodTypePatch,
	"DELETE": CreateTaskFromRestTaskCancelMethodTypeDelete,
	"PUT":    CreateTaskFromRestTaskCancelMethodTypePut,
}

var mappingCreateTaskFromRestTaskCancelMethodTypeEnumLowerCase = map[string]CreateTaskFromRestTaskCancelMethodTypeEnum{
	"get":    CreateTaskFromRestTaskCancelMethodTypeGet,
	"post":   CreateTaskFromRestTaskCancelMethodTypePost,
	"patch":  CreateTaskFromRestTaskCancelMethodTypePatch,
	"delete": CreateTaskFromRestTaskCancelMethodTypeDelete,
	"put":    CreateTaskFromRestTaskCancelMethodTypePut,
}

// GetCreateTaskFromRestTaskCancelMethodTypeEnumValues Enumerates the set of values for CreateTaskFromRestTaskCancelMethodTypeEnum
func GetCreateTaskFromRestTaskCancelMethodTypeEnumValues() []CreateTaskFromRestTaskCancelMethodTypeEnum {
	values := make([]CreateTaskFromRestTaskCancelMethodTypeEnum, 0)
	for _, v := range mappingCreateTaskFromRestTaskCancelMethodTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateTaskFromRestTaskCancelMethodTypeEnumStringValues Enumerates the set of values in String for CreateTaskFromRestTaskCancelMethodTypeEnum
func GetCreateTaskFromRestTaskCancelMethodTypeEnumStringValues() []string {
	return []string{
		"GET",
		"POST",
		"PATCH",
		"DELETE",
		"PUT",
	}
}

// GetMappingCreateTaskFromRestTaskCancelMethodTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateTaskFromRestTaskCancelMethodTypeEnum(val string) (CreateTaskFromRestTaskCancelMethodTypeEnum, bool) {
	enum, ok := mappingCreateTaskFromRestTaskCancelMethodTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
