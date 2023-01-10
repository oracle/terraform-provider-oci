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

// Update The representation of Update
type Update struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the maintenance update.
	Id *string `mandatory:"true" json:"id"`

	// Details of the maintenance update package.
	Description *string `mandatory:"true" json:"description"`

	// The type of cloud VM cluster maintenance update.
	UpdateType UpdateUpdateTypeEnum `mandatory:"true" json:"updateType"`

	// The date and time the maintenance update was released.
	TimeReleased *common.SDKTime `mandatory:"true" json:"timeReleased"`

	// The version of the maintenance update package.
	Version *string `mandatory:"true" json:"version"`

	// The update action.
	LastAction UpdateLastActionEnum `mandatory:"false" json:"lastAction,omitempty"`

	// The possible actions performed by the update operation on the infrastructure components.
	AvailableActions []UpdateAvailableActionsEnum `mandatory:"false" json:"availableActions,omitempty"`

	// Descriptive text providing additional details about the lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The current state of the maintenance update. Dependent on value of `lastAction`.
	LifecycleState UpdateLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m Update) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Update) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingUpdateUpdateTypeEnum(string(m.UpdateType)); !ok && m.UpdateType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpdateType: %s. Supported values are: %s.", m.UpdateType, strings.Join(GetUpdateUpdateTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingUpdateLastActionEnum(string(m.LastAction)); !ok && m.LastAction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LastAction: %s. Supported values are: %s.", m.LastAction, strings.Join(GetUpdateLastActionEnumStringValues(), ",")))
	}
	for _, val := range m.AvailableActions {
		if _, ok := GetMappingUpdateAvailableActionsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AvailableActions: %s. Supported values are: %s.", val, strings.Join(GetUpdateAvailableActionsEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingUpdateLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetUpdateLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateLastActionEnum Enum with underlying type: string
type UpdateLastActionEnum string

// Set of constants representing the allowable values for UpdateLastActionEnum
const (
	UpdateLastActionRollingApply    UpdateLastActionEnum = "ROLLING_APPLY"
	UpdateLastActionNonRollingApply UpdateLastActionEnum = "NON_ROLLING_APPLY"
	UpdateLastActionPrecheck        UpdateLastActionEnum = "PRECHECK"
	UpdateLastActionRollback        UpdateLastActionEnum = "ROLLBACK"
)

var mappingUpdateLastActionEnum = map[string]UpdateLastActionEnum{
	"ROLLING_APPLY":     UpdateLastActionRollingApply,
	"NON_ROLLING_APPLY": UpdateLastActionNonRollingApply,
	"PRECHECK":          UpdateLastActionPrecheck,
	"ROLLBACK":          UpdateLastActionRollback,
}

var mappingUpdateLastActionEnumLowerCase = map[string]UpdateLastActionEnum{
	"rolling_apply":     UpdateLastActionRollingApply,
	"non_rolling_apply": UpdateLastActionNonRollingApply,
	"precheck":          UpdateLastActionPrecheck,
	"rollback":          UpdateLastActionRollback,
}

// GetUpdateLastActionEnumValues Enumerates the set of values for UpdateLastActionEnum
func GetUpdateLastActionEnumValues() []UpdateLastActionEnum {
	values := make([]UpdateLastActionEnum, 0)
	for _, v := range mappingUpdateLastActionEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateLastActionEnumStringValues Enumerates the set of values in String for UpdateLastActionEnum
func GetUpdateLastActionEnumStringValues() []string {
	return []string{
		"ROLLING_APPLY",
		"NON_ROLLING_APPLY",
		"PRECHECK",
		"ROLLBACK",
	}
}

// GetMappingUpdateLastActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateLastActionEnum(val string) (UpdateLastActionEnum, bool) {
	enum, ok := mappingUpdateLastActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// UpdateAvailableActionsEnum Enum with underlying type: string
type UpdateAvailableActionsEnum string

// Set of constants representing the allowable values for UpdateAvailableActionsEnum
const (
	UpdateAvailableActionsRollingApply    UpdateAvailableActionsEnum = "ROLLING_APPLY"
	UpdateAvailableActionsNonRollingApply UpdateAvailableActionsEnum = "NON_ROLLING_APPLY"
	UpdateAvailableActionsPrecheck        UpdateAvailableActionsEnum = "PRECHECK"
	UpdateAvailableActionsRollback        UpdateAvailableActionsEnum = "ROLLBACK"
)

var mappingUpdateAvailableActionsEnum = map[string]UpdateAvailableActionsEnum{
	"ROLLING_APPLY":     UpdateAvailableActionsRollingApply,
	"NON_ROLLING_APPLY": UpdateAvailableActionsNonRollingApply,
	"PRECHECK":          UpdateAvailableActionsPrecheck,
	"ROLLBACK":          UpdateAvailableActionsRollback,
}

var mappingUpdateAvailableActionsEnumLowerCase = map[string]UpdateAvailableActionsEnum{
	"rolling_apply":     UpdateAvailableActionsRollingApply,
	"non_rolling_apply": UpdateAvailableActionsNonRollingApply,
	"precheck":          UpdateAvailableActionsPrecheck,
	"rollback":          UpdateAvailableActionsRollback,
}

// GetUpdateAvailableActionsEnumValues Enumerates the set of values for UpdateAvailableActionsEnum
func GetUpdateAvailableActionsEnumValues() []UpdateAvailableActionsEnum {
	values := make([]UpdateAvailableActionsEnum, 0)
	for _, v := range mappingUpdateAvailableActionsEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateAvailableActionsEnumStringValues Enumerates the set of values in String for UpdateAvailableActionsEnum
func GetUpdateAvailableActionsEnumStringValues() []string {
	return []string{
		"ROLLING_APPLY",
		"NON_ROLLING_APPLY",
		"PRECHECK",
		"ROLLBACK",
	}
}

// GetMappingUpdateAvailableActionsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateAvailableActionsEnum(val string) (UpdateAvailableActionsEnum, bool) {
	enum, ok := mappingUpdateAvailableActionsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// UpdateUpdateTypeEnum Enum with underlying type: string
type UpdateUpdateTypeEnum string

// Set of constants representing the allowable values for UpdateUpdateTypeEnum
const (
	UpdateUpdateTypeGiUpgrade UpdateUpdateTypeEnum = "GI_UPGRADE"
	UpdateUpdateTypeGiPatch   UpdateUpdateTypeEnum = "GI_PATCH"
	UpdateUpdateTypeOsUpdate  UpdateUpdateTypeEnum = "OS_UPDATE"
)

var mappingUpdateUpdateTypeEnum = map[string]UpdateUpdateTypeEnum{
	"GI_UPGRADE": UpdateUpdateTypeGiUpgrade,
	"GI_PATCH":   UpdateUpdateTypeGiPatch,
	"OS_UPDATE":  UpdateUpdateTypeOsUpdate,
}

var mappingUpdateUpdateTypeEnumLowerCase = map[string]UpdateUpdateTypeEnum{
	"gi_upgrade": UpdateUpdateTypeGiUpgrade,
	"gi_patch":   UpdateUpdateTypeGiPatch,
	"os_update":  UpdateUpdateTypeOsUpdate,
}

// GetUpdateUpdateTypeEnumValues Enumerates the set of values for UpdateUpdateTypeEnum
func GetUpdateUpdateTypeEnumValues() []UpdateUpdateTypeEnum {
	values := make([]UpdateUpdateTypeEnum, 0)
	for _, v := range mappingUpdateUpdateTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateUpdateTypeEnumStringValues Enumerates the set of values in String for UpdateUpdateTypeEnum
func GetUpdateUpdateTypeEnumStringValues() []string {
	return []string{
		"GI_UPGRADE",
		"GI_PATCH",
		"OS_UPDATE",
	}
}

// GetMappingUpdateUpdateTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateUpdateTypeEnum(val string) (UpdateUpdateTypeEnum, bool) {
	enum, ok := mappingUpdateUpdateTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// UpdateLifecycleStateEnum Enum with underlying type: string
type UpdateLifecycleStateEnum string

// Set of constants representing the allowable values for UpdateLifecycleStateEnum
const (
	UpdateLifecycleStateAvailable  UpdateLifecycleStateEnum = "AVAILABLE"
	UpdateLifecycleStateSuccess    UpdateLifecycleStateEnum = "SUCCESS"
	UpdateLifecycleStateInProgress UpdateLifecycleStateEnum = "IN_PROGRESS"
	UpdateLifecycleStateFailed     UpdateLifecycleStateEnum = "FAILED"
)

var mappingUpdateLifecycleStateEnum = map[string]UpdateLifecycleStateEnum{
	"AVAILABLE":   UpdateLifecycleStateAvailable,
	"SUCCESS":     UpdateLifecycleStateSuccess,
	"IN_PROGRESS": UpdateLifecycleStateInProgress,
	"FAILED":      UpdateLifecycleStateFailed,
}

var mappingUpdateLifecycleStateEnumLowerCase = map[string]UpdateLifecycleStateEnum{
	"available":   UpdateLifecycleStateAvailable,
	"success":     UpdateLifecycleStateSuccess,
	"in_progress": UpdateLifecycleStateInProgress,
	"failed":      UpdateLifecycleStateFailed,
}

// GetUpdateLifecycleStateEnumValues Enumerates the set of values for UpdateLifecycleStateEnum
func GetUpdateLifecycleStateEnumValues() []UpdateLifecycleStateEnum {
	values := make([]UpdateLifecycleStateEnum, 0)
	for _, v := range mappingUpdateLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateLifecycleStateEnumStringValues Enumerates the set of values in String for UpdateLifecycleStateEnum
func GetUpdateLifecycleStateEnumStringValues() []string {
	return []string{
		"AVAILABLE",
		"SUCCESS",
		"IN_PROGRESS",
		"FAILED",
	}
}

// GetMappingUpdateLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateLifecycleStateEnum(val string) (UpdateLifecycleStateEnum, bool) {
	enum, ok := mappingUpdateLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
