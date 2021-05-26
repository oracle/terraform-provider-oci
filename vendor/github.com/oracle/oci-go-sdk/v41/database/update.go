// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/v41/common"
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

// UpdateLastActionEnum Enum with underlying type: string
type UpdateLastActionEnum string

// Set of constants representing the allowable values for UpdateLastActionEnum
const (
	UpdateLastActionRollingApply    UpdateLastActionEnum = "ROLLING_APPLY"
	UpdateLastActionNonRollingApply UpdateLastActionEnum = "NON_ROLLING_APPLY"
	UpdateLastActionPrecheck        UpdateLastActionEnum = "PRECHECK"
	UpdateLastActionRollback        UpdateLastActionEnum = "ROLLBACK"
)

var mappingUpdateLastAction = map[string]UpdateLastActionEnum{
	"ROLLING_APPLY":     UpdateLastActionRollingApply,
	"NON_ROLLING_APPLY": UpdateLastActionNonRollingApply,
	"PRECHECK":          UpdateLastActionPrecheck,
	"ROLLBACK":          UpdateLastActionRollback,
}

// GetUpdateLastActionEnumValues Enumerates the set of values for UpdateLastActionEnum
func GetUpdateLastActionEnumValues() []UpdateLastActionEnum {
	values := make([]UpdateLastActionEnum, 0)
	for _, v := range mappingUpdateLastAction {
		values = append(values, v)
	}
	return values
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

var mappingUpdateAvailableActions = map[string]UpdateAvailableActionsEnum{
	"ROLLING_APPLY":     UpdateAvailableActionsRollingApply,
	"NON_ROLLING_APPLY": UpdateAvailableActionsNonRollingApply,
	"PRECHECK":          UpdateAvailableActionsPrecheck,
	"ROLLBACK":          UpdateAvailableActionsRollback,
}

// GetUpdateAvailableActionsEnumValues Enumerates the set of values for UpdateAvailableActionsEnum
func GetUpdateAvailableActionsEnumValues() []UpdateAvailableActionsEnum {
	values := make([]UpdateAvailableActionsEnum, 0)
	for _, v := range mappingUpdateAvailableActions {
		values = append(values, v)
	}
	return values
}

// UpdateUpdateTypeEnum Enum with underlying type: string
type UpdateUpdateTypeEnum string

// Set of constants representing the allowable values for UpdateUpdateTypeEnum
const (
	UpdateUpdateTypeGiUpgrade UpdateUpdateTypeEnum = "GI_UPGRADE"
	UpdateUpdateTypeGiPatch   UpdateUpdateTypeEnum = "GI_PATCH"
	UpdateUpdateTypeOsUpdate  UpdateUpdateTypeEnum = "OS_UPDATE"
)

var mappingUpdateUpdateType = map[string]UpdateUpdateTypeEnum{
	"GI_UPGRADE": UpdateUpdateTypeGiUpgrade,
	"GI_PATCH":   UpdateUpdateTypeGiPatch,
	"OS_UPDATE":  UpdateUpdateTypeOsUpdate,
}

// GetUpdateUpdateTypeEnumValues Enumerates the set of values for UpdateUpdateTypeEnum
func GetUpdateUpdateTypeEnumValues() []UpdateUpdateTypeEnum {
	values := make([]UpdateUpdateTypeEnum, 0)
	for _, v := range mappingUpdateUpdateType {
		values = append(values, v)
	}
	return values
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

var mappingUpdateLifecycleState = map[string]UpdateLifecycleStateEnum{
	"AVAILABLE":   UpdateLifecycleStateAvailable,
	"SUCCESS":     UpdateLifecycleStateSuccess,
	"IN_PROGRESS": UpdateLifecycleStateInProgress,
	"FAILED":      UpdateLifecycleStateFailed,
}

// GetUpdateLifecycleStateEnumValues Enumerates the set of values for UpdateLifecycleStateEnum
func GetUpdateLifecycleStateEnumValues() []UpdateLifecycleStateEnum {
	values := make([]UpdateLifecycleStateEnum, 0)
	for _, v := range mappingUpdateLifecycleState {
		values = append(values, v)
	}
	return values
}
