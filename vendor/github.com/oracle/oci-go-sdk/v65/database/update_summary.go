// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateSummary A maintenance update for a cloud VM cluster. Applies to Exadata Cloud Service instances only.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access,
// see Getting Started with Policies (https://docs.cloud.oracle.com/Content/Identity/Concepts/policygetstarted.htm).
type UpdateSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the maintenance update.
	Id *string `mandatory:"true" json:"id"`

	// Details of the maintenance update package.
	Description *string `mandatory:"true" json:"description"`

	// The type of cloud VM cluster maintenance update.
	UpdateType UpdateSummaryUpdateTypeEnum `mandatory:"true" json:"updateType"`

	// The date and time the maintenance update was released.
	TimeReleased *common.SDKTime `mandatory:"true" json:"timeReleased"`

	// The version of the maintenance update package.
	Version *string `mandatory:"true" json:"version"`

	// The update action.
	LastAction UpdateSummaryLastActionEnum `mandatory:"false" json:"lastAction,omitempty"`

	// The possible actions performed by the update operation on the infrastructure components.
	AvailableActions []UpdateSummaryAvailableActionsEnum `mandatory:"false" json:"availableActions,omitempty"`

	// Descriptive text providing additional details about the lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The current state of the maintenance update. Dependent on value of `lastAction`.
	LifecycleState UpdateSummaryLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m UpdateSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingUpdateSummaryUpdateTypeEnum(string(m.UpdateType)); !ok && m.UpdateType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpdateType: %s. Supported values are: %s.", m.UpdateType, strings.Join(GetUpdateSummaryUpdateTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingUpdateSummaryLastActionEnum(string(m.LastAction)); !ok && m.LastAction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LastAction: %s. Supported values are: %s.", m.LastAction, strings.Join(GetUpdateSummaryLastActionEnumStringValues(), ",")))
	}
	for _, val := range m.AvailableActions {
		if _, ok := GetMappingUpdateSummaryAvailableActionsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AvailableActions: %s. Supported values are: %s.", val, strings.Join(GetUpdateSummaryAvailableActionsEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingUpdateSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetUpdateSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateSummaryLastActionEnum Enum with underlying type: string
type UpdateSummaryLastActionEnum string

// Set of constants representing the allowable values for UpdateSummaryLastActionEnum
const (
	UpdateSummaryLastActionRollingApply    UpdateSummaryLastActionEnum = "ROLLING_APPLY"
	UpdateSummaryLastActionNonRollingApply UpdateSummaryLastActionEnum = "NON_ROLLING_APPLY"
	UpdateSummaryLastActionPrecheck        UpdateSummaryLastActionEnum = "PRECHECK"
	UpdateSummaryLastActionRollback        UpdateSummaryLastActionEnum = "ROLLBACK"
)

var mappingUpdateSummaryLastActionEnum = map[string]UpdateSummaryLastActionEnum{
	"ROLLING_APPLY":     UpdateSummaryLastActionRollingApply,
	"NON_ROLLING_APPLY": UpdateSummaryLastActionNonRollingApply,
	"PRECHECK":          UpdateSummaryLastActionPrecheck,
	"ROLLBACK":          UpdateSummaryLastActionRollback,
}

var mappingUpdateSummaryLastActionEnumLowerCase = map[string]UpdateSummaryLastActionEnum{
	"rolling_apply":     UpdateSummaryLastActionRollingApply,
	"non_rolling_apply": UpdateSummaryLastActionNonRollingApply,
	"precheck":          UpdateSummaryLastActionPrecheck,
	"rollback":          UpdateSummaryLastActionRollback,
}

// GetUpdateSummaryLastActionEnumValues Enumerates the set of values for UpdateSummaryLastActionEnum
func GetUpdateSummaryLastActionEnumValues() []UpdateSummaryLastActionEnum {
	values := make([]UpdateSummaryLastActionEnum, 0)
	for _, v := range mappingUpdateSummaryLastActionEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateSummaryLastActionEnumStringValues Enumerates the set of values in String for UpdateSummaryLastActionEnum
func GetUpdateSummaryLastActionEnumStringValues() []string {
	return []string{
		"ROLLING_APPLY",
		"NON_ROLLING_APPLY",
		"PRECHECK",
		"ROLLBACK",
	}
}

// GetMappingUpdateSummaryLastActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateSummaryLastActionEnum(val string) (UpdateSummaryLastActionEnum, bool) {
	enum, ok := mappingUpdateSummaryLastActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// UpdateSummaryAvailableActionsEnum Enum with underlying type: string
type UpdateSummaryAvailableActionsEnum string

// Set of constants representing the allowable values for UpdateSummaryAvailableActionsEnum
const (
	UpdateSummaryAvailableActionsRollingApply    UpdateSummaryAvailableActionsEnum = "ROLLING_APPLY"
	UpdateSummaryAvailableActionsNonRollingApply UpdateSummaryAvailableActionsEnum = "NON_ROLLING_APPLY"
	UpdateSummaryAvailableActionsPrecheck        UpdateSummaryAvailableActionsEnum = "PRECHECK"
	UpdateSummaryAvailableActionsRollback        UpdateSummaryAvailableActionsEnum = "ROLLBACK"
)

var mappingUpdateSummaryAvailableActionsEnum = map[string]UpdateSummaryAvailableActionsEnum{
	"ROLLING_APPLY":     UpdateSummaryAvailableActionsRollingApply,
	"NON_ROLLING_APPLY": UpdateSummaryAvailableActionsNonRollingApply,
	"PRECHECK":          UpdateSummaryAvailableActionsPrecheck,
	"ROLLBACK":          UpdateSummaryAvailableActionsRollback,
}

var mappingUpdateSummaryAvailableActionsEnumLowerCase = map[string]UpdateSummaryAvailableActionsEnum{
	"rolling_apply":     UpdateSummaryAvailableActionsRollingApply,
	"non_rolling_apply": UpdateSummaryAvailableActionsNonRollingApply,
	"precheck":          UpdateSummaryAvailableActionsPrecheck,
	"rollback":          UpdateSummaryAvailableActionsRollback,
}

// GetUpdateSummaryAvailableActionsEnumValues Enumerates the set of values for UpdateSummaryAvailableActionsEnum
func GetUpdateSummaryAvailableActionsEnumValues() []UpdateSummaryAvailableActionsEnum {
	values := make([]UpdateSummaryAvailableActionsEnum, 0)
	for _, v := range mappingUpdateSummaryAvailableActionsEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateSummaryAvailableActionsEnumStringValues Enumerates the set of values in String for UpdateSummaryAvailableActionsEnum
func GetUpdateSummaryAvailableActionsEnumStringValues() []string {
	return []string{
		"ROLLING_APPLY",
		"NON_ROLLING_APPLY",
		"PRECHECK",
		"ROLLBACK",
	}
}

// GetMappingUpdateSummaryAvailableActionsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateSummaryAvailableActionsEnum(val string) (UpdateSummaryAvailableActionsEnum, bool) {
	enum, ok := mappingUpdateSummaryAvailableActionsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// UpdateSummaryUpdateTypeEnum Enum with underlying type: string
type UpdateSummaryUpdateTypeEnum string

// Set of constants representing the allowable values for UpdateSummaryUpdateTypeEnum
const (
	UpdateSummaryUpdateTypeGiUpgrade UpdateSummaryUpdateTypeEnum = "GI_UPGRADE"
	UpdateSummaryUpdateTypeGiPatch   UpdateSummaryUpdateTypeEnum = "GI_PATCH"
	UpdateSummaryUpdateTypeOsUpdate  UpdateSummaryUpdateTypeEnum = "OS_UPDATE"
)

var mappingUpdateSummaryUpdateTypeEnum = map[string]UpdateSummaryUpdateTypeEnum{
	"GI_UPGRADE": UpdateSummaryUpdateTypeGiUpgrade,
	"GI_PATCH":   UpdateSummaryUpdateTypeGiPatch,
	"OS_UPDATE":  UpdateSummaryUpdateTypeOsUpdate,
}

var mappingUpdateSummaryUpdateTypeEnumLowerCase = map[string]UpdateSummaryUpdateTypeEnum{
	"gi_upgrade": UpdateSummaryUpdateTypeGiUpgrade,
	"gi_patch":   UpdateSummaryUpdateTypeGiPatch,
	"os_update":  UpdateSummaryUpdateTypeOsUpdate,
}

// GetUpdateSummaryUpdateTypeEnumValues Enumerates the set of values for UpdateSummaryUpdateTypeEnum
func GetUpdateSummaryUpdateTypeEnumValues() []UpdateSummaryUpdateTypeEnum {
	values := make([]UpdateSummaryUpdateTypeEnum, 0)
	for _, v := range mappingUpdateSummaryUpdateTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateSummaryUpdateTypeEnumStringValues Enumerates the set of values in String for UpdateSummaryUpdateTypeEnum
func GetUpdateSummaryUpdateTypeEnumStringValues() []string {
	return []string{
		"GI_UPGRADE",
		"GI_PATCH",
		"OS_UPDATE",
	}
}

// GetMappingUpdateSummaryUpdateTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateSummaryUpdateTypeEnum(val string) (UpdateSummaryUpdateTypeEnum, bool) {
	enum, ok := mappingUpdateSummaryUpdateTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// UpdateSummaryLifecycleStateEnum Enum with underlying type: string
type UpdateSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for UpdateSummaryLifecycleStateEnum
const (
	UpdateSummaryLifecycleStateAvailable  UpdateSummaryLifecycleStateEnum = "AVAILABLE"
	UpdateSummaryLifecycleStateSuccess    UpdateSummaryLifecycleStateEnum = "SUCCESS"
	UpdateSummaryLifecycleStateInProgress UpdateSummaryLifecycleStateEnum = "IN_PROGRESS"
	UpdateSummaryLifecycleStateFailed     UpdateSummaryLifecycleStateEnum = "FAILED"
)

var mappingUpdateSummaryLifecycleStateEnum = map[string]UpdateSummaryLifecycleStateEnum{
	"AVAILABLE":   UpdateSummaryLifecycleStateAvailable,
	"SUCCESS":     UpdateSummaryLifecycleStateSuccess,
	"IN_PROGRESS": UpdateSummaryLifecycleStateInProgress,
	"FAILED":      UpdateSummaryLifecycleStateFailed,
}

var mappingUpdateSummaryLifecycleStateEnumLowerCase = map[string]UpdateSummaryLifecycleStateEnum{
	"available":   UpdateSummaryLifecycleStateAvailable,
	"success":     UpdateSummaryLifecycleStateSuccess,
	"in_progress": UpdateSummaryLifecycleStateInProgress,
	"failed":      UpdateSummaryLifecycleStateFailed,
}

// GetUpdateSummaryLifecycleStateEnumValues Enumerates the set of values for UpdateSummaryLifecycleStateEnum
func GetUpdateSummaryLifecycleStateEnumValues() []UpdateSummaryLifecycleStateEnum {
	values := make([]UpdateSummaryLifecycleStateEnum, 0)
	for _, v := range mappingUpdateSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for UpdateSummaryLifecycleStateEnum
func GetUpdateSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"AVAILABLE",
		"SUCCESS",
		"IN_PROGRESS",
		"FAILED",
	}
}

// GetMappingUpdateSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateSummaryLifecycleStateEnum(val string) (UpdateSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingUpdateSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
