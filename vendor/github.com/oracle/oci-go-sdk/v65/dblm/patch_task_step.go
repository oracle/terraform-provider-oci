// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// PatchTaskStep A step in task of a patch operation used to patch one resource.
type PatchTaskStep struct {

	// Unique identifier that is system generated
	Key *int64 `mandatory:"true" json:"key"`

	// PatchTask Identifier which contains this patch step
	TaskKey *int64 `mandatory:"false" json:"taskKey"`

	// PatchOperation Identifier which contains this step
	OperationId *string `mandatory:"false" json:"operationId"`

	// The type of the PatchTaskStep
	Type PatchTaskStepTypeEnum `mandatory:"false" json:"type,omitempty"`

	// Step Name based on the step type
	Name *string `mandatory:"false" json:"name"`

	// Detailed description of step
	Description *string `mandatory:"false" json:"description"`

	// The time the PatchTaskStep was actually started. An RFC3339 formatted datetime string
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The time the PatchTaskStep was actually completed. An RFC3339 formatted datetime string
	TimeCompleted *common.SDKTime `mandatory:"false" json:"timeCompleted"`

	// The elapsed time for the PatchTaskStep in seconds.
	TimeElapsedInSeconds *int64 `mandatory:"false" json:"timeElapsedInSeconds"`

	// The work state of the PatchTaskStep.
	WorkState PatchTaskStepWorkStateEnum `mandatory:"false" json:"workState,omitempty"`

	// The status of the PatchTaskStep.
	Status PatchTaskStepStatusEnum `mandatory:"false" json:"status,omitempty"`

	Properties *PatchTaskStepPropertiesList `mandatory:"false" json:"properties"`

	// The type of the PatchTaskStep
	OutputType PatchTaskStepOutputTypeEnum `mandatory:"false" json:"outputType,omitempty"`

	Output *PatchTaskStepOutputList `mandatory:"false" json:"output"`

	ValidationOutput *PatchTaskStepValidationMessages `mandatory:"false" json:"validationOutput"`

	// Location of log file for the step
	LogFileLocation *string `mandatory:"false" json:"logFileLocation"`

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

func (m PatchTaskStep) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PatchTaskStep) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingPatchTaskStepTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetPatchTaskStepTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPatchTaskStepWorkStateEnum(string(m.WorkState)); !ok && m.WorkState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for WorkState: %s. Supported values are: %s.", m.WorkState, strings.Join(GetPatchTaskStepWorkStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPatchTaskStepStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetPatchTaskStepStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPatchTaskStepOutputTypeEnum(string(m.OutputType)); !ok && m.OutputType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OutputType: %s. Supported values are: %s.", m.OutputType, strings.Join(GetPatchTaskStepOutputTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PatchTaskStepTypeEnum Enum with underlying type: string
type PatchTaskStepTypeEnum string

// Set of constants representing the allowable values for PatchTaskStepTypeEnum
const (
	PatchTaskStepTypeInitializeDeploy             PatchTaskStepTypeEnum = "INITIALIZE_DEPLOY"
	PatchTaskStepTypeSubmitDeploy                 PatchTaskStepTypeEnum = "SUBMIT_DEPLOY"
	PatchTaskStepTypeSubmitRootDeploy             PatchTaskStepTypeEnum = "SUBMIT_ROOT_DEPLOY"
	PatchTaskStepTypeInitializeUpdate             PatchTaskStepTypeEnum = "INITIALIZE_UPDATE"
	PatchTaskStepTypeSubmitModifyClusterResources PatchTaskStepTypeEnum = "SUBMIT_MODIFY_CLUSTER_RESOURCES"
	PatchTaskStepTypeSubmitUpdate                 PatchTaskStepTypeEnum = "SUBMIT_UPDATE"
	PatchTaskStepTypeSubmitCleanup                PatchTaskStepTypeEnum = "SUBMIT_CLEANUP"
	PatchTaskStepTypeSubmitApplyPatches           PatchTaskStepTypeEnum = "SUBMIT_APPLY_PATCHES"
	PatchTaskStepTypeInitializeMigrateListener    PatchTaskStepTypeEnum = "INITIALIZE_MIGRATE_LISTENER"
	PatchTaskStepTypeSubmitMigrateListener        PatchTaskStepTypeEnum = "SUBMIT_MIGRATE_LISTENER"
	PatchTaskStepTypeMetadataValidations          PatchTaskStepTypeEnum = "METADATA_VALIDATIONS"
	PatchTaskStepTypeSubmitHostValidations        PatchTaskStepTypeEnum = "SUBMIT_HOST_VALIDATIONS"
	PatchTaskStepTypePollInitializeDeploy         PatchTaskStepTypeEnum = "POLL_INITIALIZE_DEPLOY"
	PatchTaskStepTypePollJobDeploy                PatchTaskStepTypeEnum = "POLL_JOB_DEPLOY"
	PatchTaskStepTypePollJobRootDeploy            PatchTaskStepTypeEnum = "POLL_JOB_ROOT_DEPLOY"
	PatchTaskStepTypePollInitializeUpdate         PatchTaskStepTypeEnum = "POLL_INITIALIZE_UPDATE"
	PatchTaskStepTypePollModifyClusterResources   PatchTaskStepTypeEnum = "POLL_MODIFY_CLUSTER_RESOURCES"
	PatchTaskStepTypePollJobUpdate                PatchTaskStepTypeEnum = "POLL_JOB_UPDATE"
	PatchTaskStepTypePollJobCleanup               PatchTaskStepTypeEnum = "POLL_JOB_CLEANUP"
	PatchTaskStepTypePollApplyPatches             PatchTaskStepTypeEnum = "POLL_APPLY_PATCHES"
	PatchTaskStepTypePollJobMigrateListener       PatchTaskStepTypeEnum = "POLL_JOB_MIGRATE_LISTENER"
	PatchTaskStepTypePollHostValidations          PatchTaskStepTypeEnum = "POLL_HOST_VALIDATIONS"
	PatchTaskStepTypeFetchDbHome                  PatchTaskStepTypeEnum = "FETCH_DB_HOME"
	PatchTaskStepTypePreCheckDatabase             PatchTaskStepTypeEnum = "PRE_CHECK_DATABASE"
	PatchTaskStepTypePollPreCheckDatabase         PatchTaskStepTypeEnum = "POLL_PRE_CHECK_DATABASE"
	PatchTaskStepTypeApplyUpdateDatabase          PatchTaskStepTypeEnum = "APPLY_UPDATE_DATABASE"
	PatchTaskStepTypePollApplyUpdateDatabase      PatchTaskStepTypeEnum = "POLL_APPLY_UPDATE_DATABASE"
)

var mappingPatchTaskStepTypeEnum = map[string]PatchTaskStepTypeEnum{
	"INITIALIZE_DEPLOY":               PatchTaskStepTypeInitializeDeploy,
	"SUBMIT_DEPLOY":                   PatchTaskStepTypeSubmitDeploy,
	"SUBMIT_ROOT_DEPLOY":              PatchTaskStepTypeSubmitRootDeploy,
	"INITIALIZE_UPDATE":               PatchTaskStepTypeInitializeUpdate,
	"SUBMIT_MODIFY_CLUSTER_RESOURCES": PatchTaskStepTypeSubmitModifyClusterResources,
	"SUBMIT_UPDATE":                   PatchTaskStepTypeSubmitUpdate,
	"SUBMIT_CLEANUP":                  PatchTaskStepTypeSubmitCleanup,
	"SUBMIT_APPLY_PATCHES":            PatchTaskStepTypeSubmitApplyPatches,
	"INITIALIZE_MIGRATE_LISTENER":     PatchTaskStepTypeInitializeMigrateListener,
	"SUBMIT_MIGRATE_LISTENER":         PatchTaskStepTypeSubmitMigrateListener,
	"METADATA_VALIDATIONS":            PatchTaskStepTypeMetadataValidations,
	"SUBMIT_HOST_VALIDATIONS":         PatchTaskStepTypeSubmitHostValidations,
	"POLL_INITIALIZE_DEPLOY":          PatchTaskStepTypePollInitializeDeploy,
	"POLL_JOB_DEPLOY":                 PatchTaskStepTypePollJobDeploy,
	"POLL_JOB_ROOT_DEPLOY":            PatchTaskStepTypePollJobRootDeploy,
	"POLL_INITIALIZE_UPDATE":          PatchTaskStepTypePollInitializeUpdate,
	"POLL_MODIFY_CLUSTER_RESOURCES":   PatchTaskStepTypePollModifyClusterResources,
	"POLL_JOB_UPDATE":                 PatchTaskStepTypePollJobUpdate,
	"POLL_JOB_CLEANUP":                PatchTaskStepTypePollJobCleanup,
	"POLL_APPLY_PATCHES":              PatchTaskStepTypePollApplyPatches,
	"POLL_JOB_MIGRATE_LISTENER":       PatchTaskStepTypePollJobMigrateListener,
	"POLL_HOST_VALIDATIONS":           PatchTaskStepTypePollHostValidations,
	"FETCH_DB_HOME":                   PatchTaskStepTypeFetchDbHome,
	"PRE_CHECK_DATABASE":              PatchTaskStepTypePreCheckDatabase,
	"POLL_PRE_CHECK_DATABASE":         PatchTaskStepTypePollPreCheckDatabase,
	"APPLY_UPDATE_DATABASE":           PatchTaskStepTypeApplyUpdateDatabase,
	"POLL_APPLY_UPDATE_DATABASE":      PatchTaskStepTypePollApplyUpdateDatabase,
}

var mappingPatchTaskStepTypeEnumLowerCase = map[string]PatchTaskStepTypeEnum{
	"initialize_deploy":               PatchTaskStepTypeInitializeDeploy,
	"submit_deploy":                   PatchTaskStepTypeSubmitDeploy,
	"submit_root_deploy":              PatchTaskStepTypeSubmitRootDeploy,
	"initialize_update":               PatchTaskStepTypeInitializeUpdate,
	"submit_modify_cluster_resources": PatchTaskStepTypeSubmitModifyClusterResources,
	"submit_update":                   PatchTaskStepTypeSubmitUpdate,
	"submit_cleanup":                  PatchTaskStepTypeSubmitCleanup,
	"submit_apply_patches":            PatchTaskStepTypeSubmitApplyPatches,
	"initialize_migrate_listener":     PatchTaskStepTypeInitializeMigrateListener,
	"submit_migrate_listener":         PatchTaskStepTypeSubmitMigrateListener,
	"metadata_validations":            PatchTaskStepTypeMetadataValidations,
	"submit_host_validations":         PatchTaskStepTypeSubmitHostValidations,
	"poll_initialize_deploy":          PatchTaskStepTypePollInitializeDeploy,
	"poll_job_deploy":                 PatchTaskStepTypePollJobDeploy,
	"poll_job_root_deploy":            PatchTaskStepTypePollJobRootDeploy,
	"poll_initialize_update":          PatchTaskStepTypePollInitializeUpdate,
	"poll_modify_cluster_resources":   PatchTaskStepTypePollModifyClusterResources,
	"poll_job_update":                 PatchTaskStepTypePollJobUpdate,
	"poll_job_cleanup":                PatchTaskStepTypePollJobCleanup,
	"poll_apply_patches":              PatchTaskStepTypePollApplyPatches,
	"poll_job_migrate_listener":       PatchTaskStepTypePollJobMigrateListener,
	"poll_host_validations":           PatchTaskStepTypePollHostValidations,
	"fetch_db_home":                   PatchTaskStepTypeFetchDbHome,
	"pre_check_database":              PatchTaskStepTypePreCheckDatabase,
	"poll_pre_check_database":         PatchTaskStepTypePollPreCheckDatabase,
	"apply_update_database":           PatchTaskStepTypeApplyUpdateDatabase,
	"poll_apply_update_database":      PatchTaskStepTypePollApplyUpdateDatabase,
}

// GetPatchTaskStepTypeEnumValues Enumerates the set of values for PatchTaskStepTypeEnum
func GetPatchTaskStepTypeEnumValues() []PatchTaskStepTypeEnum {
	values := make([]PatchTaskStepTypeEnum, 0)
	for _, v := range mappingPatchTaskStepTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetPatchTaskStepTypeEnumStringValues Enumerates the set of values in String for PatchTaskStepTypeEnum
func GetPatchTaskStepTypeEnumStringValues() []string {
	return []string{
		"INITIALIZE_DEPLOY",
		"SUBMIT_DEPLOY",
		"SUBMIT_ROOT_DEPLOY",
		"INITIALIZE_UPDATE",
		"SUBMIT_MODIFY_CLUSTER_RESOURCES",
		"SUBMIT_UPDATE",
		"SUBMIT_CLEANUP",
		"SUBMIT_APPLY_PATCHES",
		"INITIALIZE_MIGRATE_LISTENER",
		"SUBMIT_MIGRATE_LISTENER",
		"METADATA_VALIDATIONS",
		"SUBMIT_HOST_VALIDATIONS",
		"POLL_INITIALIZE_DEPLOY",
		"POLL_JOB_DEPLOY",
		"POLL_JOB_ROOT_DEPLOY",
		"POLL_INITIALIZE_UPDATE",
		"POLL_MODIFY_CLUSTER_RESOURCES",
		"POLL_JOB_UPDATE",
		"POLL_JOB_CLEANUP",
		"POLL_APPLY_PATCHES",
		"POLL_JOB_MIGRATE_LISTENER",
		"POLL_HOST_VALIDATIONS",
		"FETCH_DB_HOME",
		"PRE_CHECK_DATABASE",
		"POLL_PRE_CHECK_DATABASE",
		"APPLY_UPDATE_DATABASE",
		"POLL_APPLY_UPDATE_DATABASE",
	}
}

// GetMappingPatchTaskStepTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPatchTaskStepTypeEnum(val string) (PatchTaskStepTypeEnum, bool) {
	enum, ok := mappingPatchTaskStepTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// PatchTaskStepWorkStateEnum Enum with underlying type: string
type PatchTaskStepWorkStateEnum string

// Set of constants representing the allowable values for PatchTaskStepWorkStateEnum
const (
	PatchTaskStepWorkStateScheduled PatchTaskStepWorkStateEnum = "SCHEDULED"
	PatchTaskStepWorkStateRunning   PatchTaskStepWorkStateEnum = "RUNNING"
	PatchTaskStepWorkStateCompleted PatchTaskStepWorkStateEnum = "COMPLETED"
	PatchTaskStepWorkStateFailed    PatchTaskStepWorkStateEnum = "FAILED"
)

var mappingPatchTaskStepWorkStateEnum = map[string]PatchTaskStepWorkStateEnum{
	"SCHEDULED": PatchTaskStepWorkStateScheduled,
	"RUNNING":   PatchTaskStepWorkStateRunning,
	"COMPLETED": PatchTaskStepWorkStateCompleted,
	"FAILED":    PatchTaskStepWorkStateFailed,
}

var mappingPatchTaskStepWorkStateEnumLowerCase = map[string]PatchTaskStepWorkStateEnum{
	"scheduled": PatchTaskStepWorkStateScheduled,
	"running":   PatchTaskStepWorkStateRunning,
	"completed": PatchTaskStepWorkStateCompleted,
	"failed":    PatchTaskStepWorkStateFailed,
}

// GetPatchTaskStepWorkStateEnumValues Enumerates the set of values for PatchTaskStepWorkStateEnum
func GetPatchTaskStepWorkStateEnumValues() []PatchTaskStepWorkStateEnum {
	values := make([]PatchTaskStepWorkStateEnum, 0)
	for _, v := range mappingPatchTaskStepWorkStateEnum {
		values = append(values, v)
	}
	return values
}

// GetPatchTaskStepWorkStateEnumStringValues Enumerates the set of values in String for PatchTaskStepWorkStateEnum
func GetPatchTaskStepWorkStateEnumStringValues() []string {
	return []string{
		"SCHEDULED",
		"RUNNING",
		"COMPLETED",
		"FAILED",
	}
}

// GetMappingPatchTaskStepWorkStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPatchTaskStepWorkStateEnum(val string) (PatchTaskStepWorkStateEnum, bool) {
	enum, ok := mappingPatchTaskStepWorkStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// PatchTaskStepStatusEnum Enum with underlying type: string
type PatchTaskStepStatusEnum string

// Set of constants representing the allowable values for PatchTaskStepStatusEnum
const (
	PatchTaskStepStatusScheduled  PatchTaskStepStatusEnum = "SCHEDULED"
	PatchTaskStepStatusRunning    PatchTaskStepStatusEnum = "RUNNING"
	PatchTaskStepStatusSuccessful PatchTaskStepStatusEnum = "SUCCESSFUL"
	PatchTaskStepStatusWarnings   PatchTaskStepStatusEnum = "WARNINGS"
	PatchTaskStepStatusFailed     PatchTaskStepStatusEnum = "FAILED"
	PatchTaskStepStatusSkipped    PatchTaskStepStatusEnum = "SKIPPED"
)

var mappingPatchTaskStepStatusEnum = map[string]PatchTaskStepStatusEnum{
	"SCHEDULED":  PatchTaskStepStatusScheduled,
	"RUNNING":    PatchTaskStepStatusRunning,
	"SUCCESSFUL": PatchTaskStepStatusSuccessful,
	"WARNINGS":   PatchTaskStepStatusWarnings,
	"FAILED":     PatchTaskStepStatusFailed,
	"SKIPPED":    PatchTaskStepStatusSkipped,
}

var mappingPatchTaskStepStatusEnumLowerCase = map[string]PatchTaskStepStatusEnum{
	"scheduled":  PatchTaskStepStatusScheduled,
	"running":    PatchTaskStepStatusRunning,
	"successful": PatchTaskStepStatusSuccessful,
	"warnings":   PatchTaskStepStatusWarnings,
	"failed":     PatchTaskStepStatusFailed,
	"skipped":    PatchTaskStepStatusSkipped,
}

// GetPatchTaskStepStatusEnumValues Enumerates the set of values for PatchTaskStepStatusEnum
func GetPatchTaskStepStatusEnumValues() []PatchTaskStepStatusEnum {
	values := make([]PatchTaskStepStatusEnum, 0)
	for _, v := range mappingPatchTaskStepStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetPatchTaskStepStatusEnumStringValues Enumerates the set of values in String for PatchTaskStepStatusEnum
func GetPatchTaskStepStatusEnumStringValues() []string {
	return []string{
		"SCHEDULED",
		"RUNNING",
		"SUCCESSFUL",
		"WARNINGS",
		"FAILED",
		"SKIPPED",
	}
}

// GetMappingPatchTaskStepStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPatchTaskStepStatusEnum(val string) (PatchTaskStepStatusEnum, bool) {
	enum, ok := mappingPatchTaskStepStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// PatchTaskStepOutputTypeEnum Enum with underlying type: string
type PatchTaskStepOutputTypeEnum string

// Set of constants representing the allowable values for PatchTaskStepOutputTypeEnum
const (
	PatchTaskStepOutputTypeValidation PatchTaskStepOutputTypeEnum = "VALIDATION"
	PatchTaskStepOutputTypeMessages   PatchTaskStepOutputTypeEnum = "MESSAGES"
)

var mappingPatchTaskStepOutputTypeEnum = map[string]PatchTaskStepOutputTypeEnum{
	"VALIDATION": PatchTaskStepOutputTypeValidation,
	"MESSAGES":   PatchTaskStepOutputTypeMessages,
}

var mappingPatchTaskStepOutputTypeEnumLowerCase = map[string]PatchTaskStepOutputTypeEnum{
	"validation": PatchTaskStepOutputTypeValidation,
	"messages":   PatchTaskStepOutputTypeMessages,
}

// GetPatchTaskStepOutputTypeEnumValues Enumerates the set of values for PatchTaskStepOutputTypeEnum
func GetPatchTaskStepOutputTypeEnumValues() []PatchTaskStepOutputTypeEnum {
	values := make([]PatchTaskStepOutputTypeEnum, 0)
	for _, v := range mappingPatchTaskStepOutputTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetPatchTaskStepOutputTypeEnumStringValues Enumerates the set of values in String for PatchTaskStepOutputTypeEnum
func GetPatchTaskStepOutputTypeEnumStringValues() []string {
	return []string{
		"VALIDATION",
		"MESSAGES",
	}
}

// GetMappingPatchTaskStepOutputTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPatchTaskStepOutputTypeEnum(val string) (PatchTaskStepOutputTypeEnum, bool) {
	enum, ok := mappingPatchTaskStepOutputTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
