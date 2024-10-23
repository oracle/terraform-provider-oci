// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TaskRecord Details of a task.
type TaskRecord struct {

	// The OCID of the resource.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	// Example: `My new resource`
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Task type.
	Type TaskRecordTypeEnum `mandatory:"true" json:"type"`

	// The current state of the TaskRecord.
	LifecycleState TaskRecordLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The time this resource was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time this resource was last updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	Details *Details `mandatory:"true" json:"details"`

	// OCID of the compartment to which the resource belongs to.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly description. To provide some insight about the resource.
	// Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The version of the task
	Version *string `mandatory:"false" json:"version"`

	// Associated region
	ResourceRegion *string `mandatory:"false" json:"resourceRegion"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m TaskRecord) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TaskRecord) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingTaskRecordTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetTaskRecordTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTaskRecordLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetTaskRecordLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TaskRecordTypeEnum Enum with underlying type: string
type TaskRecordTypeEnum string

// Set of constants representing the allowable values for TaskRecordTypeEnum
const (
	TaskRecordTypeUserDefined   TaskRecordTypeEnum = "USER_DEFINED"
	TaskRecordTypeOracleDefined TaskRecordTypeEnum = "ORACLE_DEFINED"
	TaskRecordTypeSystemDefined TaskRecordTypeEnum = "SYSTEM_DEFINED"
)

var mappingTaskRecordTypeEnum = map[string]TaskRecordTypeEnum{
	"USER_DEFINED":   TaskRecordTypeUserDefined,
	"ORACLE_DEFINED": TaskRecordTypeOracleDefined,
	"SYSTEM_DEFINED": TaskRecordTypeSystemDefined,
}

var mappingTaskRecordTypeEnumLowerCase = map[string]TaskRecordTypeEnum{
	"user_defined":   TaskRecordTypeUserDefined,
	"oracle_defined": TaskRecordTypeOracleDefined,
	"system_defined": TaskRecordTypeSystemDefined,
}

// GetTaskRecordTypeEnumValues Enumerates the set of values for TaskRecordTypeEnum
func GetTaskRecordTypeEnumValues() []TaskRecordTypeEnum {
	values := make([]TaskRecordTypeEnum, 0)
	for _, v := range mappingTaskRecordTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetTaskRecordTypeEnumStringValues Enumerates the set of values in String for TaskRecordTypeEnum
func GetTaskRecordTypeEnumStringValues() []string {
	return []string{
		"USER_DEFINED",
		"ORACLE_DEFINED",
		"SYSTEM_DEFINED",
	}
}

// GetMappingTaskRecordTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTaskRecordTypeEnum(val string) (TaskRecordTypeEnum, bool) {
	enum, ok := mappingTaskRecordTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// TaskRecordLifecycleStateEnum Enum with underlying type: string
type TaskRecordLifecycleStateEnum string

// Set of constants representing the allowable values for TaskRecordLifecycleStateEnum
const (
	TaskRecordLifecycleStateActive   TaskRecordLifecycleStateEnum = "ACTIVE"
	TaskRecordLifecycleStateInactive TaskRecordLifecycleStateEnum = "INACTIVE"
	TaskRecordLifecycleStateDeleted  TaskRecordLifecycleStateEnum = "DELETED"
	TaskRecordLifecycleStateDeleting TaskRecordLifecycleStateEnum = "DELETING"
	TaskRecordLifecycleStateFailed   TaskRecordLifecycleStateEnum = "FAILED"
	TaskRecordLifecycleStateUpdating TaskRecordLifecycleStateEnum = "UPDATING"
)

var mappingTaskRecordLifecycleStateEnum = map[string]TaskRecordLifecycleStateEnum{
	"ACTIVE":   TaskRecordLifecycleStateActive,
	"INACTIVE": TaskRecordLifecycleStateInactive,
	"DELETED":  TaskRecordLifecycleStateDeleted,
	"DELETING": TaskRecordLifecycleStateDeleting,
	"FAILED":   TaskRecordLifecycleStateFailed,
	"UPDATING": TaskRecordLifecycleStateUpdating,
}

var mappingTaskRecordLifecycleStateEnumLowerCase = map[string]TaskRecordLifecycleStateEnum{
	"active":   TaskRecordLifecycleStateActive,
	"inactive": TaskRecordLifecycleStateInactive,
	"deleted":  TaskRecordLifecycleStateDeleted,
	"deleting": TaskRecordLifecycleStateDeleting,
	"failed":   TaskRecordLifecycleStateFailed,
	"updating": TaskRecordLifecycleStateUpdating,
}

// GetTaskRecordLifecycleStateEnumValues Enumerates the set of values for TaskRecordLifecycleStateEnum
func GetTaskRecordLifecycleStateEnumValues() []TaskRecordLifecycleStateEnum {
	values := make([]TaskRecordLifecycleStateEnum, 0)
	for _, v := range mappingTaskRecordLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetTaskRecordLifecycleStateEnumStringValues Enumerates the set of values in String for TaskRecordLifecycleStateEnum
func GetTaskRecordLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
		"DELETED",
		"DELETING",
		"FAILED",
		"UPDATING",
	}
}

// GetMappingTaskRecordLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTaskRecordLifecycleStateEnum(val string) (TaskRecordLifecycleStateEnum, bool) {
	enum, ok := mappingTaskRecordLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
