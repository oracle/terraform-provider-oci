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

// UpdateSummaryLastActionEnum Enum with underlying type: string
type UpdateSummaryLastActionEnum string

// Set of constants representing the allowable values for UpdateSummaryLastActionEnum
const (
	UpdateSummaryLastActionRollingApply    UpdateSummaryLastActionEnum = "ROLLING_APPLY"
	UpdateSummaryLastActionNonRollingApply UpdateSummaryLastActionEnum = "NON_ROLLING_APPLY"
	UpdateSummaryLastActionPrecheck        UpdateSummaryLastActionEnum = "PRECHECK"
	UpdateSummaryLastActionRollback        UpdateSummaryLastActionEnum = "ROLLBACK"
)

var mappingUpdateSummaryLastAction = map[string]UpdateSummaryLastActionEnum{
	"ROLLING_APPLY":     UpdateSummaryLastActionRollingApply,
	"NON_ROLLING_APPLY": UpdateSummaryLastActionNonRollingApply,
	"PRECHECK":          UpdateSummaryLastActionPrecheck,
	"ROLLBACK":          UpdateSummaryLastActionRollback,
}

// GetUpdateSummaryLastActionEnumValues Enumerates the set of values for UpdateSummaryLastActionEnum
func GetUpdateSummaryLastActionEnumValues() []UpdateSummaryLastActionEnum {
	values := make([]UpdateSummaryLastActionEnum, 0)
	for _, v := range mappingUpdateSummaryLastAction {
		values = append(values, v)
	}
	return values
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

var mappingUpdateSummaryAvailableActions = map[string]UpdateSummaryAvailableActionsEnum{
	"ROLLING_APPLY":     UpdateSummaryAvailableActionsRollingApply,
	"NON_ROLLING_APPLY": UpdateSummaryAvailableActionsNonRollingApply,
	"PRECHECK":          UpdateSummaryAvailableActionsPrecheck,
	"ROLLBACK":          UpdateSummaryAvailableActionsRollback,
}

// GetUpdateSummaryAvailableActionsEnumValues Enumerates the set of values for UpdateSummaryAvailableActionsEnum
func GetUpdateSummaryAvailableActionsEnumValues() []UpdateSummaryAvailableActionsEnum {
	values := make([]UpdateSummaryAvailableActionsEnum, 0)
	for _, v := range mappingUpdateSummaryAvailableActions {
		values = append(values, v)
	}
	return values
}

// UpdateSummaryUpdateTypeEnum Enum with underlying type: string
type UpdateSummaryUpdateTypeEnum string

// Set of constants representing the allowable values for UpdateSummaryUpdateTypeEnum
const (
	UpdateSummaryUpdateTypeGiUpgrade UpdateSummaryUpdateTypeEnum = "GI_UPGRADE"
	UpdateSummaryUpdateTypeGiPatch   UpdateSummaryUpdateTypeEnum = "GI_PATCH"
	UpdateSummaryUpdateTypeOsUpdate  UpdateSummaryUpdateTypeEnum = "OS_UPDATE"
)

var mappingUpdateSummaryUpdateType = map[string]UpdateSummaryUpdateTypeEnum{
	"GI_UPGRADE": UpdateSummaryUpdateTypeGiUpgrade,
	"GI_PATCH":   UpdateSummaryUpdateTypeGiPatch,
	"OS_UPDATE":  UpdateSummaryUpdateTypeOsUpdate,
}

// GetUpdateSummaryUpdateTypeEnumValues Enumerates the set of values for UpdateSummaryUpdateTypeEnum
func GetUpdateSummaryUpdateTypeEnumValues() []UpdateSummaryUpdateTypeEnum {
	values := make([]UpdateSummaryUpdateTypeEnum, 0)
	for _, v := range mappingUpdateSummaryUpdateType {
		values = append(values, v)
	}
	return values
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

var mappingUpdateSummaryLifecycleState = map[string]UpdateSummaryLifecycleStateEnum{
	"AVAILABLE":   UpdateSummaryLifecycleStateAvailable,
	"SUCCESS":     UpdateSummaryLifecycleStateSuccess,
	"IN_PROGRESS": UpdateSummaryLifecycleStateInProgress,
	"FAILED":      UpdateSummaryLifecycleStateFailed,
}

// GetUpdateSummaryLifecycleStateEnumValues Enumerates the set of values for UpdateSummaryLifecycleStateEnum
func GetUpdateSummaryLifecycleStateEnumValues() []UpdateSummaryLifecycleStateEnum {
	values := make([]UpdateSummaryLifecycleStateEnum, 0)
	for _, v := range mappingUpdateSummaryLifecycleState {
		values = append(values, v)
	}
	return values
}
