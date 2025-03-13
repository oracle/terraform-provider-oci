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

// PatchOperation An operation to patch one or more resources.
type PatchOperation struct {

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// PatchOperation Identifier, can be renamed
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment Identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The time the PatchOperation was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the PatchOperation.
	LifecycleState PatchOperationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// If true, only validates that prerequisites are met
	IsPrerequisitesOnly *bool `mandatory:"false" json:"isPrerequisitesOnly"`

	// If true, only quick validations are run. This is applicable only when isPrerequisitesOnly=true.
	IsQuickPrerequisitesCheck *bool `mandatory:"false" json:"isQuickPrerequisitesCheck"`

	// Working directory for staging binaries and temporary files
	WorkingDirectory *string `mandatory:"false" json:"workingDirectory"`

	DeployHomesSpecification *DeployHomesSpecification `mandatory:"false" json:"deployHomesSpecification"`

	MigrateListenerSpecification *MigrateListenerSpecification `mandatory:"false" json:"migrateListenerSpecification"`

	UpdateSpecification *UpdateSpecification `mandatory:"false" json:"updateSpecification"`

	CleanupHomesSpecification *CleanupHomesSpecification `mandatory:"false" json:"cleanupHomesSpecification"`

	RollbackListenerSpecification *RollbackListenerSpecification `mandatory:"false" json:"rollbackListenerSpecification"`

	RollbackSpecification *RollbackSpecification `mandatory:"false" json:"rollbackSpecification"`

	// The unique identifier of the user who invoked this operation
	InvokedByUserId *string `mandatory:"false" json:"invokedByUserId"`

	// The time the PatchOperation was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The actual start time of the PatchOperation. An RFC3339 formatted datetime string
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The actual completion time of the PatchOperation. An RFC3339 formatted datetime string
	TimeCompleted *common.SDKTime `mandatory:"false" json:"timeCompleted"`

	// The elapsed time for the PatchOperation in seconds.
	TimeElapsedInSeconds *int64 `mandatory:"false" json:"timeElapsedInSeconds"`

	// The work state of the PatchOperation.
	WorkState PatchOperationWorkStateEnum `mandatory:"false" json:"workState,omitempty"`

	// The status of the PatchOperation.
	Status PatchOperationStatusEnum `mandatory:"false" json:"status,omitempty"`

	// The total number of warnings found during the PatchOperation execution
	WarningsCount *int `mandatory:"false" json:"warningsCount"`

	// The total number of retries done for the PatchOperation
	RetryCount *int `mandatory:"false" json:"retryCount"`

	// The number of resources patched by the PatchOperation
	ResourceCount *int `mandatory:"false" json:"resourceCount"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	TaskSummary *PatchOperationTaskSummary `mandatory:"false" json:"taskSummary"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m PatchOperation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PatchOperation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPatchOperationLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetPatchOperationLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingPatchOperationWorkStateEnum(string(m.WorkState)); !ok && m.WorkState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for WorkState: %s. Supported values are: %s.", m.WorkState, strings.Join(GetPatchOperationWorkStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPatchOperationStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetPatchOperationStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PatchOperationWorkStateEnum Enum with underlying type: string
type PatchOperationWorkStateEnum string

// Set of constants representing the allowable values for PatchOperationWorkStateEnum
const (
	PatchOperationWorkStateScheduled PatchOperationWorkStateEnum = "SCHEDULED"
	PatchOperationWorkStateRunning   PatchOperationWorkStateEnum = "RUNNING"
	PatchOperationWorkStateCompleted PatchOperationWorkStateEnum = "COMPLETED"
	PatchOperationWorkStateFailed    PatchOperationWorkStateEnum = "FAILED"
)

var mappingPatchOperationWorkStateEnum = map[string]PatchOperationWorkStateEnum{
	"SCHEDULED": PatchOperationWorkStateScheduled,
	"RUNNING":   PatchOperationWorkStateRunning,
	"COMPLETED": PatchOperationWorkStateCompleted,
	"FAILED":    PatchOperationWorkStateFailed,
}

var mappingPatchOperationWorkStateEnumLowerCase = map[string]PatchOperationWorkStateEnum{
	"scheduled": PatchOperationWorkStateScheduled,
	"running":   PatchOperationWorkStateRunning,
	"completed": PatchOperationWorkStateCompleted,
	"failed":    PatchOperationWorkStateFailed,
}

// GetPatchOperationWorkStateEnumValues Enumerates the set of values for PatchOperationWorkStateEnum
func GetPatchOperationWorkStateEnumValues() []PatchOperationWorkStateEnum {
	values := make([]PatchOperationWorkStateEnum, 0)
	for _, v := range mappingPatchOperationWorkStateEnum {
		values = append(values, v)
	}
	return values
}

// GetPatchOperationWorkStateEnumStringValues Enumerates the set of values in String for PatchOperationWorkStateEnum
func GetPatchOperationWorkStateEnumStringValues() []string {
	return []string{
		"SCHEDULED",
		"RUNNING",
		"COMPLETED",
		"FAILED",
	}
}

// GetMappingPatchOperationWorkStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPatchOperationWorkStateEnum(val string) (PatchOperationWorkStateEnum, bool) {
	enum, ok := mappingPatchOperationWorkStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// PatchOperationStatusEnum Enum with underlying type: string
type PatchOperationStatusEnum string

// Set of constants representing the allowable values for PatchOperationStatusEnum
const (
	PatchOperationStatusScheduled  PatchOperationStatusEnum = "SCHEDULED"
	PatchOperationStatusRunning    PatchOperationStatusEnum = "RUNNING"
	PatchOperationStatusSuccessful PatchOperationStatusEnum = "SUCCESSFUL"
	PatchOperationStatusWarnings   PatchOperationStatusEnum = "WARNINGS"
	PatchOperationStatusFailed     PatchOperationStatusEnum = "FAILED"
)

var mappingPatchOperationStatusEnum = map[string]PatchOperationStatusEnum{
	"SCHEDULED":  PatchOperationStatusScheduled,
	"RUNNING":    PatchOperationStatusRunning,
	"SUCCESSFUL": PatchOperationStatusSuccessful,
	"WARNINGS":   PatchOperationStatusWarnings,
	"FAILED":     PatchOperationStatusFailed,
}

var mappingPatchOperationStatusEnumLowerCase = map[string]PatchOperationStatusEnum{
	"scheduled":  PatchOperationStatusScheduled,
	"running":    PatchOperationStatusRunning,
	"successful": PatchOperationStatusSuccessful,
	"warnings":   PatchOperationStatusWarnings,
	"failed":     PatchOperationStatusFailed,
}

// GetPatchOperationStatusEnumValues Enumerates the set of values for PatchOperationStatusEnum
func GetPatchOperationStatusEnumValues() []PatchOperationStatusEnum {
	values := make([]PatchOperationStatusEnum, 0)
	for _, v := range mappingPatchOperationStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetPatchOperationStatusEnumStringValues Enumerates the set of values in String for PatchOperationStatusEnum
func GetPatchOperationStatusEnumStringValues() []string {
	return []string{
		"SCHEDULED",
		"RUNNING",
		"SUCCESSFUL",
		"WARNINGS",
		"FAILED",
	}
}

// GetMappingPatchOperationStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPatchOperationStatusEnum(val string) (PatchOperationStatusEnum, bool) {
	enum, ok := mappingPatchOperationStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// PatchOperationLifecycleStateEnum Enum with underlying type: string
type PatchOperationLifecycleStateEnum string

// Set of constants representing the allowable values for PatchOperationLifecycleStateEnum
const (
	PatchOperationLifecycleStateCreating PatchOperationLifecycleStateEnum = "CREATING"
	PatchOperationLifecycleStateUpdating PatchOperationLifecycleStateEnum = "UPDATING"
	PatchOperationLifecycleStateActive   PatchOperationLifecycleStateEnum = "ACTIVE"
	PatchOperationLifecycleStateDeleting PatchOperationLifecycleStateEnum = "DELETING"
	PatchOperationLifecycleStateDeleted  PatchOperationLifecycleStateEnum = "DELETED"
	PatchOperationLifecycleStateFailed   PatchOperationLifecycleStateEnum = "FAILED"
)

var mappingPatchOperationLifecycleStateEnum = map[string]PatchOperationLifecycleStateEnum{
	"CREATING": PatchOperationLifecycleStateCreating,
	"UPDATING": PatchOperationLifecycleStateUpdating,
	"ACTIVE":   PatchOperationLifecycleStateActive,
	"DELETING": PatchOperationLifecycleStateDeleting,
	"DELETED":  PatchOperationLifecycleStateDeleted,
	"FAILED":   PatchOperationLifecycleStateFailed,
}

var mappingPatchOperationLifecycleStateEnumLowerCase = map[string]PatchOperationLifecycleStateEnum{
	"creating": PatchOperationLifecycleStateCreating,
	"updating": PatchOperationLifecycleStateUpdating,
	"active":   PatchOperationLifecycleStateActive,
	"deleting": PatchOperationLifecycleStateDeleting,
	"deleted":  PatchOperationLifecycleStateDeleted,
	"failed":   PatchOperationLifecycleStateFailed,
}

// GetPatchOperationLifecycleStateEnumValues Enumerates the set of values for PatchOperationLifecycleStateEnum
func GetPatchOperationLifecycleStateEnumValues() []PatchOperationLifecycleStateEnum {
	values := make([]PatchOperationLifecycleStateEnum, 0)
	for _, v := range mappingPatchOperationLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetPatchOperationLifecycleStateEnumStringValues Enumerates the set of values in String for PatchOperationLifecycleStateEnum
func GetPatchOperationLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingPatchOperationLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPatchOperationLifecycleStateEnum(val string) (PatchOperationLifecycleStateEnum, bool) {
	enum, ok := mappingPatchOperationLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
