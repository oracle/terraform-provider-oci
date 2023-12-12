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

// RuntimeOperator Runtime operator model which holds the runtime metadata of the task operator executed.
type RuntimeOperator struct {

	// The RuntimeOperator key.
	Key *string `mandatory:"false" json:"key"`

	// The TaskRun key.
	TaskRunKey *string `mandatory:"false" json:"taskRunKey"`

	// The runtime operator start time.
	StartTimeInMillis *int64 `mandatory:"false" json:"startTimeInMillis"`

	// The runtime operator end time.
	EndTimeInMillis *int64 `mandatory:"false" json:"endTimeInMillis"`

	// Status of RuntimeOperator. This field is deprecated, use RuntimeOperator's executionState field instead.
	Status RuntimeOperatorStatusEnum `mandatory:"false" json:"status,omitempty"`

	// The type of the object.
	ModelType *string `mandatory:"false" json:"modelType"`

	// The model version of an object.
	ModelVersion *string `mandatory:"false" json:"modelVersion"`

	ParentRef *ParentReference `mandatory:"false" json:"parentRef"`

	// Free form text without any restriction on permitted characters. Name can have letters, numbers, and special characters. The value is editable and is restricted to 1000 characters.
	Name *string `mandatory:"false" json:"name"`

	// The version of the object that is used to track changes in the object instance.
	ObjectVersion *int `mandatory:"false" json:"objectVersion"`

	// Value can only contain upper case letters, underscore and numbers. It should begin with upper case letter or underscore. The value can be modified.
	Identifier *string `mandatory:"false" json:"identifier"`

	// status
	ExecutionState RuntimeOperatorExecutionStateEnum `mandatory:"false" json:"executionState,omitempty"`

	// A list of parameters for the pipeline, this allows certain aspects of the pipeline to be configured when the pipeline is executed.
	Parameters []Parameter `mandatory:"false" json:"parameters"`

	// The status of an object that can be set to value 1 for shallow references across objects, other values reserved.
	ObjectStatus *int `mandatory:"false" json:"objectStatus"`

	Metadata *ObjectMetadata `mandatory:"false" json:"metadata"`

	Operator Operator `mandatory:"false" json:"operator"`

	// The configuration provider bindings.
	Inputs map[string]ParameterValue `mandatory:"false" json:"inputs"`

	// The configuration provider bindings.
	Outputs map[string]ParameterValue `mandatory:"false" json:"outputs"`

	// The type of task run.
	TaskType RuntimeOperatorTaskTypeEnum `mandatory:"false" json:"taskType,omitempty"`

	ConfigProvider *ConfigProvider `mandatory:"false" json:"configProvider"`

	// The type of Runtime Operator
	OperatorType RuntimeOperatorOperatorTypeEnum `mandatory:"false" json:"operatorType,omitempty"`

	// A map metrics for the task run.
	Metrics map[string]float32 `mandatory:"false" json:"metrics"`
}

func (m RuntimeOperator) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RuntimeOperator) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingRuntimeOperatorStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetRuntimeOperatorStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRuntimeOperatorExecutionStateEnum(string(m.ExecutionState)); !ok && m.ExecutionState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ExecutionState: %s. Supported values are: %s.", m.ExecutionState, strings.Join(GetRuntimeOperatorExecutionStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRuntimeOperatorTaskTypeEnum(string(m.TaskType)); !ok && m.TaskType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TaskType: %s. Supported values are: %s.", m.TaskType, strings.Join(GetRuntimeOperatorTaskTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRuntimeOperatorOperatorTypeEnum(string(m.OperatorType)); !ok && m.OperatorType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperatorType: %s. Supported values are: %s.", m.OperatorType, strings.Join(GetRuntimeOperatorOperatorTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *RuntimeOperator) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Key               *string                           `json:"key"`
		TaskRunKey        *string                           `json:"taskRunKey"`
		StartTimeInMillis *int64                            `json:"startTimeInMillis"`
		EndTimeInMillis   *int64                            `json:"endTimeInMillis"`
		Status            RuntimeOperatorStatusEnum         `json:"status"`
		ModelType         *string                           `json:"modelType"`
		ModelVersion      *string                           `json:"modelVersion"`
		ParentRef         *ParentReference                  `json:"parentRef"`
		Name              *string                           `json:"name"`
		ObjectVersion     *int                              `json:"objectVersion"`
		Identifier        *string                           `json:"identifier"`
		ExecutionState    RuntimeOperatorExecutionStateEnum `json:"executionState"`
		Parameters        []Parameter                       `json:"parameters"`
		ObjectStatus      *int                              `json:"objectStatus"`
		Metadata          *ObjectMetadata                   `json:"metadata"`
		Operator          operator                          `json:"operator"`
		Inputs            map[string]ParameterValue         `json:"inputs"`
		Outputs           map[string]ParameterValue         `json:"outputs"`
		TaskType          RuntimeOperatorTaskTypeEnum       `json:"taskType"`
		ConfigProvider    *ConfigProvider                   `json:"configProvider"`
		OperatorType      RuntimeOperatorOperatorTypeEnum   `json:"operatorType"`
		Metrics           map[string]float32                `json:"metrics"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Key = model.Key

	m.TaskRunKey = model.TaskRunKey

	m.StartTimeInMillis = model.StartTimeInMillis

	m.EndTimeInMillis = model.EndTimeInMillis

	m.Status = model.Status

	m.ModelType = model.ModelType

	m.ModelVersion = model.ModelVersion

	m.ParentRef = model.ParentRef

	m.Name = model.Name

	m.ObjectVersion = model.ObjectVersion

	m.Identifier = model.Identifier

	m.ExecutionState = model.ExecutionState

	m.Parameters = make([]Parameter, len(model.Parameters))
	copy(m.Parameters, model.Parameters)
	m.ObjectStatus = model.ObjectStatus

	m.Metadata = model.Metadata

	nn, e = model.Operator.UnmarshalPolymorphicJSON(model.Operator.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Operator = nn.(Operator)
	} else {
		m.Operator = nil
	}

	m.Inputs = model.Inputs

	m.Outputs = model.Outputs

	m.TaskType = model.TaskType

	m.ConfigProvider = model.ConfigProvider

	m.OperatorType = model.OperatorType

	m.Metrics = model.Metrics

	return
}

// RuntimeOperatorStatusEnum Enum with underlying type: string
type RuntimeOperatorStatusEnum string

// Set of constants representing the allowable values for RuntimeOperatorStatusEnum
const (
	RuntimeOperatorStatusNotStarted  RuntimeOperatorStatusEnum = "NOT_STARTED"
	RuntimeOperatorStatusQueued      RuntimeOperatorStatusEnum = "QUEUED"
	RuntimeOperatorStatusRunning     RuntimeOperatorStatusEnum = "RUNNING"
	RuntimeOperatorStatusTerminating RuntimeOperatorStatusEnum = "TERMINATING"
	RuntimeOperatorStatusTerminated  RuntimeOperatorStatusEnum = "TERMINATED"
	RuntimeOperatorStatusSuccess     RuntimeOperatorStatusEnum = "SUCCESS"
	RuntimeOperatorStatusError       RuntimeOperatorStatusEnum = "ERROR"
)

var mappingRuntimeOperatorStatusEnum = map[string]RuntimeOperatorStatusEnum{
	"NOT_STARTED": RuntimeOperatorStatusNotStarted,
	"QUEUED":      RuntimeOperatorStatusQueued,
	"RUNNING":     RuntimeOperatorStatusRunning,
	"TERMINATING": RuntimeOperatorStatusTerminating,
	"TERMINATED":  RuntimeOperatorStatusTerminated,
	"SUCCESS":     RuntimeOperatorStatusSuccess,
	"ERROR":       RuntimeOperatorStatusError,
}

var mappingRuntimeOperatorStatusEnumLowerCase = map[string]RuntimeOperatorStatusEnum{
	"not_started": RuntimeOperatorStatusNotStarted,
	"queued":      RuntimeOperatorStatusQueued,
	"running":     RuntimeOperatorStatusRunning,
	"terminating": RuntimeOperatorStatusTerminating,
	"terminated":  RuntimeOperatorStatusTerminated,
	"success":     RuntimeOperatorStatusSuccess,
	"error":       RuntimeOperatorStatusError,
}

// GetRuntimeOperatorStatusEnumValues Enumerates the set of values for RuntimeOperatorStatusEnum
func GetRuntimeOperatorStatusEnumValues() []RuntimeOperatorStatusEnum {
	values := make([]RuntimeOperatorStatusEnum, 0)
	for _, v := range mappingRuntimeOperatorStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetRuntimeOperatorStatusEnumStringValues Enumerates the set of values in String for RuntimeOperatorStatusEnum
func GetRuntimeOperatorStatusEnumStringValues() []string {
	return []string{
		"NOT_STARTED",
		"QUEUED",
		"RUNNING",
		"TERMINATING",
		"TERMINATED",
		"SUCCESS",
		"ERROR",
	}
}

// GetMappingRuntimeOperatorStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRuntimeOperatorStatusEnum(val string) (RuntimeOperatorStatusEnum, bool) {
	enum, ok := mappingRuntimeOperatorStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// RuntimeOperatorExecutionStateEnum Enum with underlying type: string
type RuntimeOperatorExecutionStateEnum string

// Set of constants representing the allowable values for RuntimeOperatorExecutionStateEnum
const (
	RuntimeOperatorExecutionStateNotStarted RuntimeOperatorExecutionStateEnum = "NOT_STARTED"
	RuntimeOperatorExecutionStateRunning    RuntimeOperatorExecutionStateEnum = "RUNNING"
	RuntimeOperatorExecutionStateTerminated RuntimeOperatorExecutionStateEnum = "TERMINATED"
	RuntimeOperatorExecutionStateSuccess    RuntimeOperatorExecutionStateEnum = "SUCCESS"
	RuntimeOperatorExecutionStateError      RuntimeOperatorExecutionStateEnum = "ERROR"
	RuntimeOperatorExecutionStateSkipped    RuntimeOperatorExecutionStateEnum = "SKIPPED"
	RuntimeOperatorExecutionStateUnknown    RuntimeOperatorExecutionStateEnum = "UNKNOWN"
	RuntimeOperatorExecutionStateIgnored    RuntimeOperatorExecutionStateEnum = "IGNORED"
)

var mappingRuntimeOperatorExecutionStateEnum = map[string]RuntimeOperatorExecutionStateEnum{
	"NOT_STARTED": RuntimeOperatorExecutionStateNotStarted,
	"RUNNING":     RuntimeOperatorExecutionStateRunning,
	"TERMINATED":  RuntimeOperatorExecutionStateTerminated,
	"SUCCESS":     RuntimeOperatorExecutionStateSuccess,
	"ERROR":       RuntimeOperatorExecutionStateError,
	"SKIPPED":     RuntimeOperatorExecutionStateSkipped,
	"UNKNOWN":     RuntimeOperatorExecutionStateUnknown,
	"IGNORED":     RuntimeOperatorExecutionStateIgnored,
}

var mappingRuntimeOperatorExecutionStateEnumLowerCase = map[string]RuntimeOperatorExecutionStateEnum{
	"not_started": RuntimeOperatorExecutionStateNotStarted,
	"running":     RuntimeOperatorExecutionStateRunning,
	"terminated":  RuntimeOperatorExecutionStateTerminated,
	"success":     RuntimeOperatorExecutionStateSuccess,
	"error":       RuntimeOperatorExecutionStateError,
	"skipped":     RuntimeOperatorExecutionStateSkipped,
	"unknown":     RuntimeOperatorExecutionStateUnknown,
	"ignored":     RuntimeOperatorExecutionStateIgnored,
}

// GetRuntimeOperatorExecutionStateEnumValues Enumerates the set of values for RuntimeOperatorExecutionStateEnum
func GetRuntimeOperatorExecutionStateEnumValues() []RuntimeOperatorExecutionStateEnum {
	values := make([]RuntimeOperatorExecutionStateEnum, 0)
	for _, v := range mappingRuntimeOperatorExecutionStateEnum {
		values = append(values, v)
	}
	return values
}

// GetRuntimeOperatorExecutionStateEnumStringValues Enumerates the set of values in String for RuntimeOperatorExecutionStateEnum
func GetRuntimeOperatorExecutionStateEnumStringValues() []string {
	return []string{
		"NOT_STARTED",
		"RUNNING",
		"TERMINATED",
		"SUCCESS",
		"ERROR",
		"SKIPPED",
		"UNKNOWN",
		"IGNORED",
	}
}

// GetMappingRuntimeOperatorExecutionStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRuntimeOperatorExecutionStateEnum(val string) (RuntimeOperatorExecutionStateEnum, bool) {
	enum, ok := mappingRuntimeOperatorExecutionStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// RuntimeOperatorTaskTypeEnum Enum with underlying type: string
type RuntimeOperatorTaskTypeEnum string

// Set of constants representing the allowable values for RuntimeOperatorTaskTypeEnum
const (
	RuntimeOperatorTaskTypeIntegrationTask RuntimeOperatorTaskTypeEnum = "INTEGRATION_TASK"
	RuntimeOperatorTaskTypeDataLoaderTask  RuntimeOperatorTaskTypeEnum = "DATA_LOADER_TASK"
	RuntimeOperatorTaskTypePipelineTask    RuntimeOperatorTaskTypeEnum = "PIPELINE_TASK"
	RuntimeOperatorTaskTypeSqlTask         RuntimeOperatorTaskTypeEnum = "SQL_TASK"
	RuntimeOperatorTaskTypeOciDataflowTask RuntimeOperatorTaskTypeEnum = "OCI_DATAFLOW_TASK"
	RuntimeOperatorTaskTypeRestTask        RuntimeOperatorTaskTypeEnum = "REST_TASK"
)

var mappingRuntimeOperatorTaskTypeEnum = map[string]RuntimeOperatorTaskTypeEnum{
	"INTEGRATION_TASK":  RuntimeOperatorTaskTypeIntegrationTask,
	"DATA_LOADER_TASK":  RuntimeOperatorTaskTypeDataLoaderTask,
	"PIPELINE_TASK":     RuntimeOperatorTaskTypePipelineTask,
	"SQL_TASK":          RuntimeOperatorTaskTypeSqlTask,
	"OCI_DATAFLOW_TASK": RuntimeOperatorTaskTypeOciDataflowTask,
	"REST_TASK":         RuntimeOperatorTaskTypeRestTask,
}

var mappingRuntimeOperatorTaskTypeEnumLowerCase = map[string]RuntimeOperatorTaskTypeEnum{
	"integration_task":  RuntimeOperatorTaskTypeIntegrationTask,
	"data_loader_task":  RuntimeOperatorTaskTypeDataLoaderTask,
	"pipeline_task":     RuntimeOperatorTaskTypePipelineTask,
	"sql_task":          RuntimeOperatorTaskTypeSqlTask,
	"oci_dataflow_task": RuntimeOperatorTaskTypeOciDataflowTask,
	"rest_task":         RuntimeOperatorTaskTypeRestTask,
}

// GetRuntimeOperatorTaskTypeEnumValues Enumerates the set of values for RuntimeOperatorTaskTypeEnum
func GetRuntimeOperatorTaskTypeEnumValues() []RuntimeOperatorTaskTypeEnum {
	values := make([]RuntimeOperatorTaskTypeEnum, 0)
	for _, v := range mappingRuntimeOperatorTaskTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRuntimeOperatorTaskTypeEnumStringValues Enumerates the set of values in String for RuntimeOperatorTaskTypeEnum
func GetRuntimeOperatorTaskTypeEnumStringValues() []string {
	return []string{
		"INTEGRATION_TASK",
		"DATA_LOADER_TASK",
		"PIPELINE_TASK",
		"SQL_TASK",
		"OCI_DATAFLOW_TASK",
		"REST_TASK",
	}
}

// GetMappingRuntimeOperatorTaskTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRuntimeOperatorTaskTypeEnum(val string) (RuntimeOperatorTaskTypeEnum, bool) {
	enum, ok := mappingRuntimeOperatorTaskTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// RuntimeOperatorOperatorTypeEnum Enum with underlying type: string
type RuntimeOperatorOperatorTypeEnum string

// Set of constants representing the allowable values for RuntimeOperatorOperatorTypeEnum
const (
	RuntimeOperatorOperatorTypeBashOperator       RuntimeOperatorOperatorTypeEnum = "BASH_OPERATOR"
	RuntimeOperatorOperatorTypeTaskOperator       RuntimeOperatorOperatorTypeEnum = "TASK_OPERATOR"
	RuntimeOperatorOperatorTypeRestOperator       RuntimeOperatorOperatorTypeEnum = "REST_OPERATOR"
	RuntimeOperatorOperatorTypeStartOperator      RuntimeOperatorOperatorTypeEnum = "START_OPERATOR"
	RuntimeOperatorOperatorTypeEndOperator        RuntimeOperatorOperatorTypeEnum = "END_OPERATOR"
	RuntimeOperatorOperatorTypeExpressionOperator RuntimeOperatorOperatorTypeEnum = "EXPRESSION_OPERATOR"
	RuntimeOperatorOperatorTypeMergeOperator      RuntimeOperatorOperatorTypeEnum = "MERGE_OPERATOR"
	RuntimeOperatorOperatorTypeDecisionOperator   RuntimeOperatorOperatorTypeEnum = "DECISION_OPERATOR"
	RuntimeOperatorOperatorTypeLoopOperator       RuntimeOperatorOperatorTypeEnum = "LOOP_OPERATOR"
	RuntimeOperatorOperatorTypeActualEndOperator  RuntimeOperatorOperatorTypeEnum = "ACTUAL_END_OPERATOR"
)

var mappingRuntimeOperatorOperatorTypeEnum = map[string]RuntimeOperatorOperatorTypeEnum{
	"BASH_OPERATOR":       RuntimeOperatorOperatorTypeBashOperator,
	"TASK_OPERATOR":       RuntimeOperatorOperatorTypeTaskOperator,
	"REST_OPERATOR":       RuntimeOperatorOperatorTypeRestOperator,
	"START_OPERATOR":      RuntimeOperatorOperatorTypeStartOperator,
	"END_OPERATOR":        RuntimeOperatorOperatorTypeEndOperator,
	"EXPRESSION_OPERATOR": RuntimeOperatorOperatorTypeExpressionOperator,
	"MERGE_OPERATOR":      RuntimeOperatorOperatorTypeMergeOperator,
	"DECISION_OPERATOR":   RuntimeOperatorOperatorTypeDecisionOperator,
	"LOOP_OPERATOR":       RuntimeOperatorOperatorTypeLoopOperator,
	"ACTUAL_END_OPERATOR": RuntimeOperatorOperatorTypeActualEndOperator,
}

var mappingRuntimeOperatorOperatorTypeEnumLowerCase = map[string]RuntimeOperatorOperatorTypeEnum{
	"bash_operator":       RuntimeOperatorOperatorTypeBashOperator,
	"task_operator":       RuntimeOperatorOperatorTypeTaskOperator,
	"rest_operator":       RuntimeOperatorOperatorTypeRestOperator,
	"start_operator":      RuntimeOperatorOperatorTypeStartOperator,
	"end_operator":        RuntimeOperatorOperatorTypeEndOperator,
	"expression_operator": RuntimeOperatorOperatorTypeExpressionOperator,
	"merge_operator":      RuntimeOperatorOperatorTypeMergeOperator,
	"decision_operator":   RuntimeOperatorOperatorTypeDecisionOperator,
	"loop_operator":       RuntimeOperatorOperatorTypeLoopOperator,
	"actual_end_operator": RuntimeOperatorOperatorTypeActualEndOperator,
}

// GetRuntimeOperatorOperatorTypeEnumValues Enumerates the set of values for RuntimeOperatorOperatorTypeEnum
func GetRuntimeOperatorOperatorTypeEnumValues() []RuntimeOperatorOperatorTypeEnum {
	values := make([]RuntimeOperatorOperatorTypeEnum, 0)
	for _, v := range mappingRuntimeOperatorOperatorTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRuntimeOperatorOperatorTypeEnumStringValues Enumerates the set of values in String for RuntimeOperatorOperatorTypeEnum
func GetRuntimeOperatorOperatorTypeEnumStringValues() []string {
	return []string{
		"BASH_OPERATOR",
		"TASK_OPERATOR",
		"REST_OPERATOR",
		"START_OPERATOR",
		"END_OPERATOR",
		"EXPRESSION_OPERATOR",
		"MERGE_OPERATOR",
		"DECISION_OPERATOR",
		"LOOP_OPERATOR",
		"ACTUAL_END_OPERATOR",
	}
}

// GetMappingRuntimeOperatorOperatorTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRuntimeOperatorOperatorTypeEnum(val string) (RuntimeOperatorOperatorTypeEnum, bool) {
	enum, ok := mappingRuntimeOperatorOperatorTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
