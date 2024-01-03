// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// TaskSummaryFromSqlTask The information about the SQL task.
type TaskSummaryFromSqlTask struct {

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

	Script *Script `mandatory:"false" json:"script"`

	// Describes the shape of the execution result
	Operation *interface{} `mandatory:"false" json:"operation"`

	// Indicates whether the task is invoking a custom SQL script or stored procedure.
	SqlScriptType TaskSummaryFromSqlTaskSqlScriptTypeEnum `mandatory:"false" json:"sqlScriptType,omitempty"`
}

// GetKey returns Key
func (m TaskSummaryFromSqlTask) GetKey() *string {
	return m.Key
}

// GetModelVersion returns ModelVersion
func (m TaskSummaryFromSqlTask) GetModelVersion() *string {
	return m.ModelVersion
}

// GetParentRef returns ParentRef
func (m TaskSummaryFromSqlTask) GetParentRef() *ParentReference {
	return m.ParentRef
}

// GetName returns Name
func (m TaskSummaryFromSqlTask) GetName() *string {
	return m.Name
}

// GetDescription returns Description
func (m TaskSummaryFromSqlTask) GetDescription() *string {
	return m.Description
}

// GetObjectVersion returns ObjectVersion
func (m TaskSummaryFromSqlTask) GetObjectVersion() *int {
	return m.ObjectVersion
}

// GetObjectStatus returns ObjectStatus
func (m TaskSummaryFromSqlTask) GetObjectStatus() *int {
	return m.ObjectStatus
}

// GetIdentifier returns Identifier
func (m TaskSummaryFromSqlTask) GetIdentifier() *string {
	return m.Identifier
}

// GetInputPorts returns InputPorts
func (m TaskSummaryFromSqlTask) GetInputPorts() []InputPort {
	return m.InputPorts
}

// GetOutputPorts returns OutputPorts
func (m TaskSummaryFromSqlTask) GetOutputPorts() []OutputPort {
	return m.OutputPorts
}

// GetParameters returns Parameters
func (m TaskSummaryFromSqlTask) GetParameters() []Parameter {
	return m.Parameters
}

// GetOpConfigValues returns OpConfigValues
func (m TaskSummaryFromSqlTask) GetOpConfigValues() *ConfigValues {
	return m.OpConfigValues
}

// GetConfigProviderDelegate returns ConfigProviderDelegate
func (m TaskSummaryFromSqlTask) GetConfigProviderDelegate() *ConfigProvider {
	return m.ConfigProviderDelegate
}

// GetIsConcurrentAllowed returns IsConcurrentAllowed
func (m TaskSummaryFromSqlTask) GetIsConcurrentAllowed() *bool {
	return m.IsConcurrentAllowed
}

// GetMetadata returns Metadata
func (m TaskSummaryFromSqlTask) GetMetadata() *ObjectMetadata {
	return m.Metadata
}

// GetKeyMap returns KeyMap
func (m TaskSummaryFromSqlTask) GetKeyMap() map[string]string {
	return m.KeyMap
}

func (m TaskSummaryFromSqlTask) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TaskSummaryFromSqlTask) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingTaskSummaryFromSqlTaskSqlScriptTypeEnum(string(m.SqlScriptType)); !ok && m.SqlScriptType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SqlScriptType: %s. Supported values are: %s.", m.SqlScriptType, strings.Join(GetTaskSummaryFromSqlTaskSqlScriptTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m TaskSummaryFromSqlTask) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeTaskSummaryFromSqlTask TaskSummaryFromSqlTask
	s := struct {
		DiscriminatorParam string `json:"modelType"`
		MarshalTypeTaskSummaryFromSqlTask
	}{
		"SQL_TASK",
		(MarshalTypeTaskSummaryFromSqlTask)(m),
	}

	return json.Marshal(&s)
}

// TaskSummaryFromSqlTaskSqlScriptTypeEnum Enum with underlying type: string
type TaskSummaryFromSqlTaskSqlScriptTypeEnum string

// Set of constants representing the allowable values for TaskSummaryFromSqlTaskSqlScriptTypeEnum
const (
	TaskSummaryFromSqlTaskSqlScriptTypeStoredProcedure TaskSummaryFromSqlTaskSqlScriptTypeEnum = "STORED_PROCEDURE"
	TaskSummaryFromSqlTaskSqlScriptTypeSqlCode         TaskSummaryFromSqlTaskSqlScriptTypeEnum = "SQL_CODE"
)

var mappingTaskSummaryFromSqlTaskSqlScriptTypeEnum = map[string]TaskSummaryFromSqlTaskSqlScriptTypeEnum{
	"STORED_PROCEDURE": TaskSummaryFromSqlTaskSqlScriptTypeStoredProcedure,
	"SQL_CODE":         TaskSummaryFromSqlTaskSqlScriptTypeSqlCode,
}

var mappingTaskSummaryFromSqlTaskSqlScriptTypeEnumLowerCase = map[string]TaskSummaryFromSqlTaskSqlScriptTypeEnum{
	"stored_procedure": TaskSummaryFromSqlTaskSqlScriptTypeStoredProcedure,
	"sql_code":         TaskSummaryFromSqlTaskSqlScriptTypeSqlCode,
}

// GetTaskSummaryFromSqlTaskSqlScriptTypeEnumValues Enumerates the set of values for TaskSummaryFromSqlTaskSqlScriptTypeEnum
func GetTaskSummaryFromSqlTaskSqlScriptTypeEnumValues() []TaskSummaryFromSqlTaskSqlScriptTypeEnum {
	values := make([]TaskSummaryFromSqlTaskSqlScriptTypeEnum, 0)
	for _, v := range mappingTaskSummaryFromSqlTaskSqlScriptTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTaskSummaryFromSqlTaskSqlScriptTypeEnumStringValues Enumerates the set of values in String for TaskSummaryFromSqlTaskSqlScriptTypeEnum
func GetTaskSummaryFromSqlTaskSqlScriptTypeEnumStringValues() []string {
	return []string{
		"STORED_PROCEDURE",
		"SQL_CODE",
	}
}

// GetMappingTaskSummaryFromSqlTaskSqlScriptTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTaskSummaryFromSqlTaskSqlScriptTypeEnum(val string) (TaskSummaryFromSqlTaskSqlScriptTypeEnum, bool) {
	enum, ok := mappingTaskSummaryFromSqlTaskSqlScriptTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
