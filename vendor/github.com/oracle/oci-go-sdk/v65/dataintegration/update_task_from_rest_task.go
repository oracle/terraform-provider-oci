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

// UpdateTaskFromRestTask The information about the Generic REST task. The endpoint and cancelEndpoint  properties are deprecated, use the properties executeRestCallConfig, cancelRestCallConfig and pollRestCallConfig for execute, cancel and polling of the calls.
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

	// Whether the same task can be executed concurrently.
	IsConcurrentAllowed *bool `mandatory:"false" json:"isConcurrentAllowed"`

	RegistryMetadata *RegistryMetadata `mandatory:"false" json:"registryMetadata"`

	AuthDetails *AuthDetails `mandatory:"false" json:"authDetails"`

	AuthConfig AuthConfig `mandatory:"false" json:"authConfig"`

	Endpoint *Expression `mandatory:"false" json:"endpoint"`

	// Headers data for the request.
	Headers *interface{} `mandatory:"false" json:"headers"`

	// Header value.
	AdditionalProperties *string `mandatory:"false" json:"additionalProperties"`

	// JSON data for payload body. This property is deprecated, use ExecuteRestCallConfig's payload config param instead.
	JsonData *string `mandatory:"false" json:"jsonData"`

	CancelEndpoint *Expression `mandatory:"false" json:"cancelEndpoint"`

	ExecuteRestCallConfig *ExecuteRestCallConfig `mandatory:"false" json:"executeRestCallConfig"`

	CancelRestCallConfig *CancelRestCallConfig `mandatory:"false" json:"cancelRestCallConfig"`

	PollRestCallConfig *PollRestCallConfig `mandatory:"false" json:"pollRestCallConfig"`

	// List of typed expressions.
	TypedExpressions []TypedExpression `mandatory:"false" json:"typedExpressions"`

	// The REST method to use. This property is deprecated, use ExecuteRestCallConfig's methodType property instead.
	MethodType UpdateTaskFromRestTaskMethodTypeEnum `mandatory:"false" json:"methodType,omitempty"`

	// The REST invocation pattern to use. ASYNC_OCI_WORKREQUEST is being deprecated as well as cancelEndpoint/MethodType.
	ApiCallMode UpdateTaskFromRestTaskApiCallModeEnum `mandatory:"false" json:"apiCallMode,omitempty"`

	// The REST method to use for canceling the original request.
	CancelMethodType UpdateTaskFromRestTaskCancelMethodTypeEnum `mandatory:"false" json:"cancelMethodType,omitempty"`
}

// GetKey returns Key
func (m UpdateTaskFromRestTask) GetKey() *string {
	return m.Key
}

// GetModelVersion returns ModelVersion
func (m UpdateTaskFromRestTask) GetModelVersion() *string {
	return m.ModelVersion
}

// GetParentRef returns ParentRef
func (m UpdateTaskFromRestTask) GetParentRef() *ParentReference {
	return m.ParentRef
}

// GetName returns Name
func (m UpdateTaskFromRestTask) GetName() *string {
	return m.Name
}

// GetDescription returns Description
func (m UpdateTaskFromRestTask) GetDescription() *string {
	return m.Description
}

// GetObjectStatus returns ObjectStatus
func (m UpdateTaskFromRestTask) GetObjectStatus() *int {
	return m.ObjectStatus
}

// GetObjectVersion returns ObjectVersion
func (m UpdateTaskFromRestTask) GetObjectVersion() *int {
	return m.ObjectVersion
}

// GetIdentifier returns Identifier
func (m UpdateTaskFromRestTask) GetIdentifier() *string {
	return m.Identifier
}

// GetInputPorts returns InputPorts
func (m UpdateTaskFromRestTask) GetInputPorts() []InputPort {
	return m.InputPorts
}

// GetOutputPorts returns OutputPorts
func (m UpdateTaskFromRestTask) GetOutputPorts() []OutputPort {
	return m.OutputPorts
}

// GetParameters returns Parameters
func (m UpdateTaskFromRestTask) GetParameters() []Parameter {
	return m.Parameters
}

// GetOpConfigValues returns OpConfigValues
func (m UpdateTaskFromRestTask) GetOpConfigValues() *ConfigValues {
	return m.OpConfigValues
}

// GetConfigProviderDelegate returns ConfigProviderDelegate
func (m UpdateTaskFromRestTask) GetConfigProviderDelegate() *ConfigProvider {
	return m.ConfigProviderDelegate
}

// GetIsConcurrentAllowed returns IsConcurrentAllowed
func (m UpdateTaskFromRestTask) GetIsConcurrentAllowed() *bool {
	return m.IsConcurrentAllowed
}

// GetRegistryMetadata returns RegistryMetadata
func (m UpdateTaskFromRestTask) GetRegistryMetadata() *RegistryMetadata {
	return m.RegistryMetadata
}

func (m UpdateTaskFromRestTask) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateTaskFromRestTask) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingUpdateTaskFromRestTaskMethodTypeEnum(string(m.MethodType)); !ok && m.MethodType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MethodType: %s. Supported values are: %s.", m.MethodType, strings.Join(GetUpdateTaskFromRestTaskMethodTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUpdateTaskFromRestTaskApiCallModeEnum(string(m.ApiCallMode)); !ok && m.ApiCallMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ApiCallMode: %s. Supported values are: %s.", m.ApiCallMode, strings.Join(GetUpdateTaskFromRestTaskApiCallModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUpdateTaskFromRestTaskCancelMethodTypeEnum(string(m.CancelMethodType)); !ok && m.CancelMethodType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CancelMethodType: %s. Supported values are: %s.", m.CancelMethodType, strings.Join(GetUpdateTaskFromRestTaskCancelMethodTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

// UnmarshalJSON unmarshals from json
func (m *UpdateTaskFromRestTask) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ModelVersion           *string                                    `json:"modelVersion"`
		ParentRef              *ParentReference                           `json:"parentRef"`
		Name                   *string                                    `json:"name"`
		Description            *string                                    `json:"description"`
		ObjectStatus           *int                                       `json:"objectStatus"`
		Identifier             *string                                    `json:"identifier"`
		InputPorts             []InputPort                                `json:"inputPorts"`
		OutputPorts            []OutputPort                               `json:"outputPorts"`
		Parameters             []Parameter                                `json:"parameters"`
		OpConfigValues         *ConfigValues                              `json:"opConfigValues"`
		ConfigProviderDelegate *ConfigProvider                            `json:"configProviderDelegate"`
		IsConcurrentAllowed    *bool                                      `json:"isConcurrentAllowed"`
		RegistryMetadata       *RegistryMetadata                          `json:"registryMetadata"`
		AuthDetails            *AuthDetails                               `json:"authDetails"`
		AuthConfig             authconfig                                 `json:"authConfig"`
		Endpoint               *Expression                                `json:"endpoint"`
		MethodType             UpdateTaskFromRestTaskMethodTypeEnum       `json:"methodType"`
		Headers                *interface{}                               `json:"headers"`
		AdditionalProperties   *string                                    `json:"additionalProperties"`
		JsonData               *string                                    `json:"jsonData"`
		ApiCallMode            UpdateTaskFromRestTaskApiCallModeEnum      `json:"apiCallMode"`
		CancelEndpoint         *Expression                                `json:"cancelEndpoint"`
		CancelMethodType       UpdateTaskFromRestTaskCancelMethodTypeEnum `json:"cancelMethodType"`
		ExecuteRestCallConfig  *ExecuteRestCallConfig                     `json:"executeRestCallConfig"`
		CancelRestCallConfig   *CancelRestCallConfig                      `json:"cancelRestCallConfig"`
		PollRestCallConfig     *PollRestCallConfig                        `json:"pollRestCallConfig"`
		TypedExpressions       []TypedExpression                          `json:"typedExpressions"`
		Key                    *string                                    `json:"key"`
		ObjectVersion          *int                                       `json:"objectVersion"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.ModelVersion = model.ModelVersion

	m.ParentRef = model.ParentRef

	m.Name = model.Name

	m.Description = model.Description

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

	m.AdditionalProperties = model.AdditionalProperties

	m.JsonData = model.JsonData

	m.ApiCallMode = model.ApiCallMode

	m.CancelEndpoint = model.CancelEndpoint

	m.CancelMethodType = model.CancelMethodType

	m.ExecuteRestCallConfig = model.ExecuteRestCallConfig

	m.CancelRestCallConfig = model.CancelRestCallConfig

	m.PollRestCallConfig = model.PollRestCallConfig

	m.TypedExpressions = make([]TypedExpression, len(model.TypedExpressions))
	copy(m.TypedExpressions, model.TypedExpressions)
	m.Key = model.Key

	m.ObjectVersion = model.ObjectVersion

	return
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

var mappingUpdateTaskFromRestTaskMethodTypeEnum = map[string]UpdateTaskFromRestTaskMethodTypeEnum{
	"GET":    UpdateTaskFromRestTaskMethodTypeGet,
	"POST":   UpdateTaskFromRestTaskMethodTypePost,
	"PATCH":  UpdateTaskFromRestTaskMethodTypePatch,
	"DELETE": UpdateTaskFromRestTaskMethodTypeDelete,
	"PUT":    UpdateTaskFromRestTaskMethodTypePut,
}

var mappingUpdateTaskFromRestTaskMethodTypeEnumLowerCase = map[string]UpdateTaskFromRestTaskMethodTypeEnum{
	"get":    UpdateTaskFromRestTaskMethodTypeGet,
	"post":   UpdateTaskFromRestTaskMethodTypePost,
	"patch":  UpdateTaskFromRestTaskMethodTypePatch,
	"delete": UpdateTaskFromRestTaskMethodTypeDelete,
	"put":    UpdateTaskFromRestTaskMethodTypePut,
}

// GetUpdateTaskFromRestTaskMethodTypeEnumValues Enumerates the set of values for UpdateTaskFromRestTaskMethodTypeEnum
func GetUpdateTaskFromRestTaskMethodTypeEnumValues() []UpdateTaskFromRestTaskMethodTypeEnum {
	values := make([]UpdateTaskFromRestTaskMethodTypeEnum, 0)
	for _, v := range mappingUpdateTaskFromRestTaskMethodTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateTaskFromRestTaskMethodTypeEnumStringValues Enumerates the set of values in String for UpdateTaskFromRestTaskMethodTypeEnum
func GetUpdateTaskFromRestTaskMethodTypeEnumStringValues() []string {
	return []string{
		"GET",
		"POST",
		"PATCH",
		"DELETE",
		"PUT",
	}
}

// GetMappingUpdateTaskFromRestTaskMethodTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateTaskFromRestTaskMethodTypeEnum(val string) (UpdateTaskFromRestTaskMethodTypeEnum, bool) {
	enum, ok := mappingUpdateTaskFromRestTaskMethodTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// UpdateTaskFromRestTaskApiCallModeEnum Enum with underlying type: string
type UpdateTaskFromRestTaskApiCallModeEnum string

// Set of constants representing the allowable values for UpdateTaskFromRestTaskApiCallModeEnum
const (
	UpdateTaskFromRestTaskApiCallModeSynchronous         UpdateTaskFromRestTaskApiCallModeEnum = "SYNCHRONOUS"
	UpdateTaskFromRestTaskApiCallModeAsyncOciWorkrequest UpdateTaskFromRestTaskApiCallModeEnum = "ASYNC_OCI_WORKREQUEST"
	UpdateTaskFromRestTaskApiCallModeAsyncGeneric        UpdateTaskFromRestTaskApiCallModeEnum = "ASYNC_GENERIC"
)

var mappingUpdateTaskFromRestTaskApiCallModeEnum = map[string]UpdateTaskFromRestTaskApiCallModeEnum{
	"SYNCHRONOUS":           UpdateTaskFromRestTaskApiCallModeSynchronous,
	"ASYNC_OCI_WORKREQUEST": UpdateTaskFromRestTaskApiCallModeAsyncOciWorkrequest,
	"ASYNC_GENERIC":         UpdateTaskFromRestTaskApiCallModeAsyncGeneric,
}

var mappingUpdateTaskFromRestTaskApiCallModeEnumLowerCase = map[string]UpdateTaskFromRestTaskApiCallModeEnum{
	"synchronous":           UpdateTaskFromRestTaskApiCallModeSynchronous,
	"async_oci_workrequest": UpdateTaskFromRestTaskApiCallModeAsyncOciWorkrequest,
	"async_generic":         UpdateTaskFromRestTaskApiCallModeAsyncGeneric,
}

// GetUpdateTaskFromRestTaskApiCallModeEnumValues Enumerates the set of values for UpdateTaskFromRestTaskApiCallModeEnum
func GetUpdateTaskFromRestTaskApiCallModeEnumValues() []UpdateTaskFromRestTaskApiCallModeEnum {
	values := make([]UpdateTaskFromRestTaskApiCallModeEnum, 0)
	for _, v := range mappingUpdateTaskFromRestTaskApiCallModeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateTaskFromRestTaskApiCallModeEnumStringValues Enumerates the set of values in String for UpdateTaskFromRestTaskApiCallModeEnum
func GetUpdateTaskFromRestTaskApiCallModeEnumStringValues() []string {
	return []string{
		"SYNCHRONOUS",
		"ASYNC_OCI_WORKREQUEST",
		"ASYNC_GENERIC",
	}
}

// GetMappingUpdateTaskFromRestTaskApiCallModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateTaskFromRestTaskApiCallModeEnum(val string) (UpdateTaskFromRestTaskApiCallModeEnum, bool) {
	enum, ok := mappingUpdateTaskFromRestTaskApiCallModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
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

var mappingUpdateTaskFromRestTaskCancelMethodTypeEnum = map[string]UpdateTaskFromRestTaskCancelMethodTypeEnum{
	"GET":    UpdateTaskFromRestTaskCancelMethodTypeGet,
	"POST":   UpdateTaskFromRestTaskCancelMethodTypePost,
	"PATCH":  UpdateTaskFromRestTaskCancelMethodTypePatch,
	"DELETE": UpdateTaskFromRestTaskCancelMethodTypeDelete,
	"PUT":    UpdateTaskFromRestTaskCancelMethodTypePut,
}

var mappingUpdateTaskFromRestTaskCancelMethodTypeEnumLowerCase = map[string]UpdateTaskFromRestTaskCancelMethodTypeEnum{
	"get":    UpdateTaskFromRestTaskCancelMethodTypeGet,
	"post":   UpdateTaskFromRestTaskCancelMethodTypePost,
	"patch":  UpdateTaskFromRestTaskCancelMethodTypePatch,
	"delete": UpdateTaskFromRestTaskCancelMethodTypeDelete,
	"put":    UpdateTaskFromRestTaskCancelMethodTypePut,
}

// GetUpdateTaskFromRestTaskCancelMethodTypeEnumValues Enumerates the set of values for UpdateTaskFromRestTaskCancelMethodTypeEnum
func GetUpdateTaskFromRestTaskCancelMethodTypeEnumValues() []UpdateTaskFromRestTaskCancelMethodTypeEnum {
	values := make([]UpdateTaskFromRestTaskCancelMethodTypeEnum, 0)
	for _, v := range mappingUpdateTaskFromRestTaskCancelMethodTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateTaskFromRestTaskCancelMethodTypeEnumStringValues Enumerates the set of values in String for UpdateTaskFromRestTaskCancelMethodTypeEnum
func GetUpdateTaskFromRestTaskCancelMethodTypeEnumStringValues() []string {
	return []string{
		"GET",
		"POST",
		"PATCH",
		"DELETE",
		"PUT",
	}
}

// GetMappingUpdateTaskFromRestTaskCancelMethodTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateTaskFromRestTaskCancelMethodTypeEnum(val string) (UpdateTaskFromRestTaskCancelMethodTypeEnum, bool) {
	enum, ok := mappingUpdateTaskFromRestTaskCancelMethodTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
