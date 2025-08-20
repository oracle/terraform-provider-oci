// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Dblm API
//
// A description of the Dblm API
//

package dblm

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PatchTask A task of a patch operation used to patch one resource.
type PatchTask struct {

	// Unique identifier that is system generated
	Key *int64 `mandatory:"true" json:"key"`

	// PatchOperation Identifier which contains this patch task
	OperationId *string `mandatory:"false" json:"operationId"`

	// The type of the PatchTask
	Type PatchTaskTypeEnum `mandatory:"false" json:"type,omitempty"`

	// The unique identifier of resource being patched
	ResourceId *string `mandatory:"false" json:"resourceId"`

	// Name of the resource being patched
	ResourceName *string `mandatory:"false" json:"resourceName"`

	// If true then patch task is skipped
	IsSkipped *bool `mandatory:"false" json:"isSkipped"`

	// Reason for the patch task operation skip
	ReasonSkipped *string `mandatory:"false" json:"reasonSkipped"`

	// The time the PatchTask was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the PatchTask was actually started. An RFC3339 formatted datetime string
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The time the PatchTask was actually completed. An RFC3339 formatted datetime string
	TimeCompleted *common.SDKTime `mandatory:"false" json:"timeCompleted"`

	// The elapsed time for the PatchTask in seconds.
	TimeElapsedInSeconds *int64 `mandatory:"false" json:"timeElapsedInSeconds"`

	// The work state of the PatchTask.
	WorkState PatchTaskWorkStateEnum `mandatory:"false" json:"workState,omitempty"`

	// The status of the PatchTask.
	Status PatchTaskStatusEnum `mandatory:"false" json:"status,omitempty"`

	// The total number of retries done for the PatchTask
	RetryCount *int `mandatory:"false" json:"retryCount"`

	DeployDetails *DeployDetails `mandatory:"false" json:"deployDetails"`

	UpdateDetails *UpdateDetails `mandatory:"false" json:"updateDetails"`

	MigrateListenerDetails *MigrateListenerDetails `mandatory:"false" json:"migrateListenerDetails"`

	CleanupDetails *CleanupDetails `mandatory:"false" json:"cleanupDetails"`

	RollbackListenerDetails *RollbackListenerDetails `mandatory:"false" json:"rollbackListenerDetails"`

	RollbackDetails *RollbackDetails `mandatory:"false" json:"rollbackDetails"`

	Messages *PatchTaskMessages `mandatory:"false" json:"messages"`

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

func (m PatchTask) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PatchTask) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingPatchTaskTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetPatchTaskTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPatchTaskWorkStateEnum(string(m.WorkState)); !ok && m.WorkState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for WorkState: %s. Supported values are: %s.", m.WorkState, strings.Join(GetPatchTaskWorkStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPatchTaskStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetPatchTaskStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PatchTaskTypeEnum Enum with underlying type: string
type PatchTaskTypeEnum string

// Set of constants representing the allowable values for PatchTaskTypeEnum
const (
	PatchTaskTypeDeploy           PatchTaskTypeEnum = "DEPLOY"
	PatchTaskTypeMigrateListener  PatchTaskTypeEnum = "MIGRATE_LISTENER"
	PatchTaskTypeUpdate           PatchTaskTypeEnum = "UPDATE"
	PatchTaskTypeCleanup          PatchTaskTypeEnum = "CLEANUP"
	PatchTaskTypeRollback         PatchTaskTypeEnum = "ROLLBACK"
	PatchTaskTypeRollbackListener PatchTaskTypeEnum = "ROLLBACK_LISTENER"
)

var mappingPatchTaskTypeEnum = map[string]PatchTaskTypeEnum{
	"DEPLOY":            PatchTaskTypeDeploy,
	"MIGRATE_LISTENER":  PatchTaskTypeMigrateListener,
	"UPDATE":            PatchTaskTypeUpdate,
	"CLEANUP":           PatchTaskTypeCleanup,
	"ROLLBACK":          PatchTaskTypeRollback,
	"ROLLBACK_LISTENER": PatchTaskTypeRollbackListener,
}

var mappingPatchTaskTypeEnumLowerCase = map[string]PatchTaskTypeEnum{
	"deploy":            PatchTaskTypeDeploy,
	"migrate_listener":  PatchTaskTypeMigrateListener,
	"update":            PatchTaskTypeUpdate,
	"cleanup":           PatchTaskTypeCleanup,
	"rollback":          PatchTaskTypeRollback,
	"rollback_listener": PatchTaskTypeRollbackListener,
}

// GetPatchTaskTypeEnumValues Enumerates the set of values for PatchTaskTypeEnum
func GetPatchTaskTypeEnumValues() []PatchTaskTypeEnum {
	values := make([]PatchTaskTypeEnum, 0)
	for _, v := range mappingPatchTaskTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetPatchTaskTypeEnumStringValues Enumerates the set of values in String for PatchTaskTypeEnum
func GetPatchTaskTypeEnumStringValues() []string {
	return []string{
		"DEPLOY",
		"MIGRATE_LISTENER",
		"UPDATE",
		"CLEANUP",
		"ROLLBACK",
		"ROLLBACK_LISTENER",
	}
}

// GetMappingPatchTaskTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPatchTaskTypeEnum(val string) (PatchTaskTypeEnum, bool) {
	enum, ok := mappingPatchTaskTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// PatchTaskWorkStateEnum Enum with underlying type: string
type PatchTaskWorkStateEnum string

// Set of constants representing the allowable values for PatchTaskWorkStateEnum
const (
	PatchTaskWorkStateScheduled PatchTaskWorkStateEnum = "SCHEDULED"
	PatchTaskWorkStateRunning   PatchTaskWorkStateEnum = "RUNNING"
	PatchTaskWorkStateCompleted PatchTaskWorkStateEnum = "COMPLETED"
	PatchTaskWorkStateFailed    PatchTaskWorkStateEnum = "FAILED"
)

var mappingPatchTaskWorkStateEnum = map[string]PatchTaskWorkStateEnum{
	"SCHEDULED": PatchTaskWorkStateScheduled,
	"RUNNING":   PatchTaskWorkStateRunning,
	"COMPLETED": PatchTaskWorkStateCompleted,
	"FAILED":    PatchTaskWorkStateFailed,
}

var mappingPatchTaskWorkStateEnumLowerCase = map[string]PatchTaskWorkStateEnum{
	"scheduled": PatchTaskWorkStateScheduled,
	"running":   PatchTaskWorkStateRunning,
	"completed": PatchTaskWorkStateCompleted,
	"failed":    PatchTaskWorkStateFailed,
}

// GetPatchTaskWorkStateEnumValues Enumerates the set of values for PatchTaskWorkStateEnum
func GetPatchTaskWorkStateEnumValues() []PatchTaskWorkStateEnum {
	values := make([]PatchTaskWorkStateEnum, 0)
	for _, v := range mappingPatchTaskWorkStateEnum {
		values = append(values, v)
	}
	return values
}

// GetPatchTaskWorkStateEnumStringValues Enumerates the set of values in String for PatchTaskWorkStateEnum
func GetPatchTaskWorkStateEnumStringValues() []string {
	return []string{
		"SCHEDULED",
		"RUNNING",
		"COMPLETED",
		"FAILED",
	}
}

// GetMappingPatchTaskWorkStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPatchTaskWorkStateEnum(val string) (PatchTaskWorkStateEnum, bool) {
	enum, ok := mappingPatchTaskWorkStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// PatchTaskStatusEnum Enum with underlying type: string
type PatchTaskStatusEnum string

// Set of constants representing the allowable values for PatchTaskStatusEnum
const (
	PatchTaskStatusScheduled  PatchTaskStatusEnum = "SCHEDULED"
	PatchTaskStatusRunning    PatchTaskStatusEnum = "RUNNING"
	PatchTaskStatusSuccessful PatchTaskStatusEnum = "SUCCESSFUL"
	PatchTaskStatusWarnings   PatchTaskStatusEnum = "WARNINGS"
	PatchTaskStatusFailed     PatchTaskStatusEnum = "FAILED"
)

var mappingPatchTaskStatusEnum = map[string]PatchTaskStatusEnum{
	"SCHEDULED":  PatchTaskStatusScheduled,
	"RUNNING":    PatchTaskStatusRunning,
	"SUCCESSFUL": PatchTaskStatusSuccessful,
	"WARNINGS":   PatchTaskStatusWarnings,
	"FAILED":     PatchTaskStatusFailed,
}

var mappingPatchTaskStatusEnumLowerCase = map[string]PatchTaskStatusEnum{
	"scheduled":  PatchTaskStatusScheduled,
	"running":    PatchTaskStatusRunning,
	"successful": PatchTaskStatusSuccessful,
	"warnings":   PatchTaskStatusWarnings,
	"failed":     PatchTaskStatusFailed,
}

// GetPatchTaskStatusEnumValues Enumerates the set of values for PatchTaskStatusEnum
func GetPatchTaskStatusEnumValues() []PatchTaskStatusEnum {
	values := make([]PatchTaskStatusEnum, 0)
	for _, v := range mappingPatchTaskStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetPatchTaskStatusEnumStringValues Enumerates the set of values in String for PatchTaskStatusEnum
func GetPatchTaskStatusEnumStringValues() []string {
	return []string{
		"SCHEDULED",
		"RUNNING",
		"SUCCESSFUL",
		"WARNINGS",
		"FAILED",
	}
}

// GetMappingPatchTaskStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPatchTaskStatusEnum(val string) (PatchTaskStatusEnum, bool) {
	enum, ok := mappingPatchTaskStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
