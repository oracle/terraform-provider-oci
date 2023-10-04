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

// RuntimeOperatorSummary The information about RuntimeOperator.
type RuntimeOperatorSummary struct {

	// The RuntimeOperator key.
	Key *string `mandatory:"false" json:"key"`

	// The TaskRun key.
	TaskRunKey *string `mandatory:"false" json:"taskRunKey"`

	// The runtime operator start time.
	StartTimeInMillis *int64 `mandatory:"false" json:"startTimeInMillis"`

	// The runtime operator end time.
	EndTimeInMillis *int64 `mandatory:"false" json:"endTimeInMillis"`

	// Status of RuntimeOperator. This field is deprecated, use RuntimeOperator's executionState field instead.
	Status RuntimeOperatorSummaryStatusEnum `mandatory:"false" json:"status,omitempty"`

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
	ExecutionState RuntimeOperatorSummaryExecutionStateEnum `mandatory:"false" json:"executionState,omitempty"`

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
	TaskType RuntimeOperatorSummaryTaskTypeEnum `mandatory:"false" json:"taskType,omitempty"`

	ConfigProvider *ConfigProvider `mandatory:"false" json:"configProvider"`

	// The type of Runtime Operator
	OperatorType RuntimeOperatorSummaryOperatorTypeEnum `mandatory:"false" json:"operatorType,omitempty"`

	// A map metrics for the task run.
	Metrics map[string]float32 `mandatory:"false" json:"metrics"`
}

func (m RuntimeOperatorSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RuntimeOperatorSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingRuntimeOperatorSummaryStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetRuntimeOperatorSummaryStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRuntimeOperatorSummaryExecutionStateEnum(string(m.ExecutionState)); !ok && m.ExecutionState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ExecutionState: %s. Supported values are: %s.", m.ExecutionState, strings.Join(GetRuntimeOperatorSummaryExecutionStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRuntimeOperatorSummaryTaskTypeEnum(string(m.TaskType)); !ok && m.TaskType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TaskType: %s. Supported values are: %s.", m.TaskType, strings.Join(GetRuntimeOperatorSummaryTaskTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRuntimeOperatorSummaryOperatorTypeEnum(string(m.OperatorType)); !ok && m.OperatorType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperatorType: %s. Supported values are: %s.", m.OperatorType, strings.Join(GetRuntimeOperatorSummaryOperatorTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *RuntimeOperatorSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Key               *string                                  `json:"key"`
		TaskRunKey        *string                                  `json:"taskRunKey"`
		StartTimeInMillis *int64                                   `json:"startTimeInMillis"`
		EndTimeInMillis   *int64                                   `json:"endTimeInMillis"`
		Status            RuntimeOperatorSummaryStatusEnum         `json:"status"`
		ModelType         *string                                  `json:"modelType"`
		ModelVersion      *string                                  `json:"modelVersion"`
		ParentRef         *ParentReference                         `json:"parentRef"`
		Name              *string                                  `json:"name"`
		ObjectVersion     *int                                     `json:"objectVersion"`
		Identifier        *string                                  `json:"identifier"`
		ExecutionState    RuntimeOperatorSummaryExecutionStateEnum `json:"executionState"`
		Parameters        []Parameter                              `json:"parameters"`
		ObjectStatus      *int                                     `json:"objectStatus"`
		Metadata          *ObjectMetadata                          `json:"metadata"`
		Operator          operator                                 `json:"operator"`
		Inputs            map[string]ParameterValue                `json:"inputs"`
		Outputs           map[string]ParameterValue                `json:"outputs"`
		TaskType          RuntimeOperatorSummaryTaskTypeEnum       `json:"taskType"`
		ConfigProvider    *ConfigProvider                          `json:"configProvider"`
		OperatorType      RuntimeOperatorSummaryOperatorTypeEnum   `json:"operatorType"`
		Metrics           map[string]float32                       `json:"metrics"`
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

// RuntimeOperatorSummaryStatusEnum Enum with underlying type: string
type RuntimeOperatorSummaryStatusEnum string

// Set of constants representing the allowable values for RuntimeOperatorSummaryStatusEnum
const (
	RuntimeOperatorSummaryStatusNotStarted  RuntimeOperatorSummaryStatusEnum = "NOT_STARTED"
	RuntimeOperatorSummaryStatusQueued      RuntimeOperatorSummaryStatusEnum = "QUEUED"
	RuntimeOperatorSummaryStatusRunning     RuntimeOperatorSummaryStatusEnum = "RUNNING"
	RuntimeOperatorSummaryStatusTerminating RuntimeOperatorSummaryStatusEnum = "TERMINATING"
	RuntimeOperatorSummaryStatusTerminated  RuntimeOperatorSummaryStatusEnum = "TERMINATED"
	RuntimeOperatorSummaryStatusSuccess     RuntimeOperatorSummaryStatusEnum = "SUCCESS"
	RuntimeOperatorSummaryStatusError       RuntimeOperatorSummaryStatusEnum = "ERROR"
)

var mappingRuntimeOperatorSummaryStatusEnum = map[string]RuntimeOperatorSummaryStatusEnum{
	"NOT_STARTED": RuntimeOperatorSummaryStatusNotStarted,
	"QUEUED":      RuntimeOperatorSummaryStatusQueued,
	"RUNNING":     RuntimeOperatorSummaryStatusRunning,
	"TERMINATING": RuntimeOperatorSummaryStatusTerminating,
	"TERMINATED":  RuntimeOperatorSummaryStatusTerminated,
	"SUCCESS":     RuntimeOperatorSummaryStatusSuccess,
	"ERROR":       RuntimeOperatorSummaryStatusError,
}

var mappingRuntimeOperatorSummaryStatusEnumLowerCase = map[string]RuntimeOperatorSummaryStatusEnum{
	"not_started": RuntimeOperatorSummaryStatusNotStarted,
	"queued":      RuntimeOperatorSummaryStatusQueued,
	"running":     RuntimeOperatorSummaryStatusRunning,
	"terminating": RuntimeOperatorSummaryStatusTerminating,
	"terminated":  RuntimeOperatorSummaryStatusTerminated,
	"success":     RuntimeOperatorSummaryStatusSuccess,
	"error":       RuntimeOperatorSummaryStatusError,
}

// GetRuntimeOperatorSummaryStatusEnumValues Enumerates the set of values for RuntimeOperatorSummaryStatusEnum
func GetRuntimeOperatorSummaryStatusEnumValues() []RuntimeOperatorSummaryStatusEnum {
	values := make([]RuntimeOperatorSummaryStatusEnum, 0)
	for _, v := range mappingRuntimeOperatorSummaryStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetRuntimeOperatorSummaryStatusEnumStringValues Enumerates the set of values in String for RuntimeOperatorSummaryStatusEnum
func GetRuntimeOperatorSummaryStatusEnumStringValues() []string {
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

// GetMappingRuntimeOperatorSummaryStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRuntimeOperatorSummaryStatusEnum(val string) (RuntimeOperatorSummaryStatusEnum, bool) {
	enum, ok := mappingRuntimeOperatorSummaryStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// RuntimeOperatorSummaryExecutionStateEnum Enum with underlying type: string
type RuntimeOperatorSummaryExecutionStateEnum string

// Set of constants representing the allowable values for RuntimeOperatorSummaryExecutionStateEnum
const (
	RuntimeOperatorSummaryExecutionStateNotStarted RuntimeOperatorSummaryExecutionStateEnum = "NOT_STARTED"
	RuntimeOperatorSummaryExecutionStateRunning    RuntimeOperatorSummaryExecutionStateEnum = "RUNNING"
	RuntimeOperatorSummaryExecutionStateTerminated RuntimeOperatorSummaryExecutionStateEnum = "TERMINATED"
	RuntimeOperatorSummaryExecutionStateSuccess    RuntimeOperatorSummaryExecutionStateEnum = "SUCCESS"
	RuntimeOperatorSummaryExecutionStateError      RuntimeOperatorSummaryExecutionStateEnum = "ERROR"
	RuntimeOperatorSummaryExecutionStateSkipped    RuntimeOperatorSummaryExecutionStateEnum = "SKIPPED"
	RuntimeOperatorSummaryExecutionStateUnknown    RuntimeOperatorSummaryExecutionStateEnum = "UNKNOWN"
	RuntimeOperatorSummaryExecutionStateIgnored    RuntimeOperatorSummaryExecutionStateEnum = "IGNORED"
)

var mappingRuntimeOperatorSummaryExecutionStateEnum = map[string]RuntimeOperatorSummaryExecutionStateEnum{
	"NOT_STARTED": RuntimeOperatorSummaryExecutionStateNotStarted,
	"RUNNING":     RuntimeOperatorSummaryExecutionStateRunning,
	"TERMINATED":  RuntimeOperatorSummaryExecutionStateTerminated,
	"SUCCESS":     RuntimeOperatorSummaryExecutionStateSuccess,
	"ERROR":       RuntimeOperatorSummaryExecutionStateError,
	"SKIPPED":     RuntimeOperatorSummaryExecutionStateSkipped,
	"UNKNOWN":     RuntimeOperatorSummaryExecutionStateUnknown,
	"IGNORED":     RuntimeOperatorSummaryExecutionStateIgnored,
}

var mappingRuntimeOperatorSummaryExecutionStateEnumLowerCase = map[string]RuntimeOperatorSummaryExecutionStateEnum{
	"not_started": RuntimeOperatorSummaryExecutionStateNotStarted,
	"running":     RuntimeOperatorSummaryExecutionStateRunning,
	"terminated":  RuntimeOperatorSummaryExecutionStateTerminated,
	"success":     RuntimeOperatorSummaryExecutionStateSuccess,
	"error":       RuntimeOperatorSummaryExecutionStateError,
	"skipped":     RuntimeOperatorSummaryExecutionStateSkipped,
	"unknown":     RuntimeOperatorSummaryExecutionStateUnknown,
	"ignored":     RuntimeOperatorSummaryExecutionStateIgnored,
}

// GetRuntimeOperatorSummaryExecutionStateEnumValues Enumerates the set of values for RuntimeOperatorSummaryExecutionStateEnum
func GetRuntimeOperatorSummaryExecutionStateEnumValues() []RuntimeOperatorSummaryExecutionStateEnum {
	values := make([]RuntimeOperatorSummaryExecutionStateEnum, 0)
	for _, v := range mappingRuntimeOperatorSummaryExecutionStateEnum {
		values = append(values, v)
	}
	return values
}

// GetRuntimeOperatorSummaryExecutionStateEnumStringValues Enumerates the set of values in String for RuntimeOperatorSummaryExecutionStateEnum
func GetRuntimeOperatorSummaryExecutionStateEnumStringValues() []string {
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

// GetMappingRuntimeOperatorSummaryExecutionStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRuntimeOperatorSummaryExecutionStateEnum(val string) (RuntimeOperatorSummaryExecutionStateEnum, bool) {
	enum, ok := mappingRuntimeOperatorSummaryExecutionStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// RuntimeOperatorSummaryTaskTypeEnum Enum with underlying type: string
type RuntimeOperatorSummaryTaskTypeEnum string

// Set of constants representing the allowable values for RuntimeOperatorSummaryTaskTypeEnum
const (
	RuntimeOperatorSummaryTaskTypeIntegrationTask RuntimeOperatorSummaryTaskTypeEnum = "INTEGRATION_TASK"
	RuntimeOperatorSummaryTaskTypeDataLoaderTask  RuntimeOperatorSummaryTaskTypeEnum = "DATA_LOADER_TASK"
	RuntimeOperatorSummaryTaskTypePipelineTask    RuntimeOperatorSummaryTaskTypeEnum = "PIPELINE_TASK"
	RuntimeOperatorSummaryTaskTypeSqlTask         RuntimeOperatorSummaryTaskTypeEnum = "SQL_TASK"
	RuntimeOperatorSummaryTaskTypeOciDataflowTask RuntimeOperatorSummaryTaskTypeEnum = "OCI_DATAFLOW_TASK"
	RuntimeOperatorSummaryTaskTypeRestTask        RuntimeOperatorSummaryTaskTypeEnum = "REST_TASK"
)

var mappingRuntimeOperatorSummaryTaskTypeEnum = map[string]RuntimeOperatorSummaryTaskTypeEnum{
	"INTEGRATION_TASK":  RuntimeOperatorSummaryTaskTypeIntegrationTask,
	"DATA_LOADER_TASK":  RuntimeOperatorSummaryTaskTypeDataLoaderTask,
	"PIPELINE_TASK":     RuntimeOperatorSummaryTaskTypePipelineTask,
	"SQL_TASK":          RuntimeOperatorSummaryTaskTypeSqlTask,
	"OCI_DATAFLOW_TASK": RuntimeOperatorSummaryTaskTypeOciDataflowTask,
	"REST_TASK":         RuntimeOperatorSummaryTaskTypeRestTask,
}

var mappingRuntimeOperatorSummaryTaskTypeEnumLowerCase = map[string]RuntimeOperatorSummaryTaskTypeEnum{
	"integration_task":  RuntimeOperatorSummaryTaskTypeIntegrationTask,
	"data_loader_task":  RuntimeOperatorSummaryTaskTypeDataLoaderTask,
	"pipeline_task":     RuntimeOperatorSummaryTaskTypePipelineTask,
	"sql_task":          RuntimeOperatorSummaryTaskTypeSqlTask,
	"oci_dataflow_task": RuntimeOperatorSummaryTaskTypeOciDataflowTask,
	"rest_task":         RuntimeOperatorSummaryTaskTypeRestTask,
}

// GetRuntimeOperatorSummaryTaskTypeEnumValues Enumerates the set of values for RuntimeOperatorSummaryTaskTypeEnum
func GetRuntimeOperatorSummaryTaskTypeEnumValues() []RuntimeOperatorSummaryTaskTypeEnum {
	values := make([]RuntimeOperatorSummaryTaskTypeEnum, 0)
	for _, v := range mappingRuntimeOperatorSummaryTaskTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRuntimeOperatorSummaryTaskTypeEnumStringValues Enumerates the set of values in String for RuntimeOperatorSummaryTaskTypeEnum
func GetRuntimeOperatorSummaryTaskTypeEnumStringValues() []string {
	return []string{
		"INTEGRATION_TASK",
		"DATA_LOADER_TASK",
		"PIPELINE_TASK",
		"SQL_TASK",
		"OCI_DATAFLOW_TASK",
		"REST_TASK",
	}
}

// GetMappingRuntimeOperatorSummaryTaskTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRuntimeOperatorSummaryTaskTypeEnum(val string) (RuntimeOperatorSummaryTaskTypeEnum, bool) {
	enum, ok := mappingRuntimeOperatorSummaryTaskTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// RuntimeOperatorSummaryOperatorTypeEnum Enum with underlying type: string
type RuntimeOperatorSummaryOperatorTypeEnum string

// Set of constants representing the allowable values for RuntimeOperatorSummaryOperatorTypeEnum
const (
	RuntimeOperatorSummaryOperatorTypeBashOperator       RuntimeOperatorSummaryOperatorTypeEnum = "BASH_OPERATOR"
	RuntimeOperatorSummaryOperatorTypeTaskOperator       RuntimeOperatorSummaryOperatorTypeEnum = "TASK_OPERATOR"
	RuntimeOperatorSummaryOperatorTypeRestOperator       RuntimeOperatorSummaryOperatorTypeEnum = "REST_OPERATOR"
	RuntimeOperatorSummaryOperatorTypeStartOperator      RuntimeOperatorSummaryOperatorTypeEnum = "START_OPERATOR"
	RuntimeOperatorSummaryOperatorTypeEndOperator        RuntimeOperatorSummaryOperatorTypeEnum = "END_OPERATOR"
	RuntimeOperatorSummaryOperatorTypeExpressionOperator RuntimeOperatorSummaryOperatorTypeEnum = "EXPRESSION_OPERATOR"
	RuntimeOperatorSummaryOperatorTypeMergeOperator      RuntimeOperatorSummaryOperatorTypeEnum = "MERGE_OPERATOR"
	RuntimeOperatorSummaryOperatorTypeDecisionOperator   RuntimeOperatorSummaryOperatorTypeEnum = "DECISION_OPERATOR"
	RuntimeOperatorSummaryOperatorTypeLoopOperator       RuntimeOperatorSummaryOperatorTypeEnum = "LOOP_OPERATOR"
	RuntimeOperatorSummaryOperatorTypeActualEndOperator  RuntimeOperatorSummaryOperatorTypeEnum = "ACTUAL_END_OPERATOR"
)

var mappingRuntimeOperatorSummaryOperatorTypeEnum = map[string]RuntimeOperatorSummaryOperatorTypeEnum{
	"BASH_OPERATOR":       RuntimeOperatorSummaryOperatorTypeBashOperator,
	"TASK_OPERATOR":       RuntimeOperatorSummaryOperatorTypeTaskOperator,
	"REST_OPERATOR":       RuntimeOperatorSummaryOperatorTypeRestOperator,
	"START_OPERATOR":      RuntimeOperatorSummaryOperatorTypeStartOperator,
	"END_OPERATOR":        RuntimeOperatorSummaryOperatorTypeEndOperator,
	"EXPRESSION_OPERATOR": RuntimeOperatorSummaryOperatorTypeExpressionOperator,
	"MERGE_OPERATOR":      RuntimeOperatorSummaryOperatorTypeMergeOperator,
	"DECISION_OPERATOR":   RuntimeOperatorSummaryOperatorTypeDecisionOperator,
	"LOOP_OPERATOR":       RuntimeOperatorSummaryOperatorTypeLoopOperator,
	"ACTUAL_END_OPERATOR": RuntimeOperatorSummaryOperatorTypeActualEndOperator,
}

var mappingRuntimeOperatorSummaryOperatorTypeEnumLowerCase = map[string]RuntimeOperatorSummaryOperatorTypeEnum{
	"bash_operator":       RuntimeOperatorSummaryOperatorTypeBashOperator,
	"task_operator":       RuntimeOperatorSummaryOperatorTypeTaskOperator,
	"rest_operator":       RuntimeOperatorSummaryOperatorTypeRestOperator,
	"start_operator":      RuntimeOperatorSummaryOperatorTypeStartOperator,
	"end_operator":        RuntimeOperatorSummaryOperatorTypeEndOperator,
	"expression_operator": RuntimeOperatorSummaryOperatorTypeExpressionOperator,
	"merge_operator":      RuntimeOperatorSummaryOperatorTypeMergeOperator,
	"decision_operator":   RuntimeOperatorSummaryOperatorTypeDecisionOperator,
	"loop_operator":       RuntimeOperatorSummaryOperatorTypeLoopOperator,
	"actual_end_operator": RuntimeOperatorSummaryOperatorTypeActualEndOperator,
}

// GetRuntimeOperatorSummaryOperatorTypeEnumValues Enumerates the set of values for RuntimeOperatorSummaryOperatorTypeEnum
func GetRuntimeOperatorSummaryOperatorTypeEnumValues() []RuntimeOperatorSummaryOperatorTypeEnum {
	values := make([]RuntimeOperatorSummaryOperatorTypeEnum, 0)
	for _, v := range mappingRuntimeOperatorSummaryOperatorTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRuntimeOperatorSummaryOperatorTypeEnumStringValues Enumerates the set of values in String for RuntimeOperatorSummaryOperatorTypeEnum
func GetRuntimeOperatorSummaryOperatorTypeEnumStringValues() []string {
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

// GetMappingRuntimeOperatorSummaryOperatorTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRuntimeOperatorSummaryOperatorTypeEnum(val string) (RuntimeOperatorSummaryOperatorTypeEnum, bool) {
	enum, ok := mappingRuntimeOperatorSummaryOperatorTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
