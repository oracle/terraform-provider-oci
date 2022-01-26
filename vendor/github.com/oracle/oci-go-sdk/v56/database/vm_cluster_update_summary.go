// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// VmClusterUpdateSummary A maintenance update for a VM cluster. Applies to Exadata Cloud@Customer instances only.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access,
// see Getting Started with Policies (https://docs.cloud.oracle.com/Content/Identity/Concepts/policygetstarted.htm).
type VmClusterUpdateSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the maintenance update.
	Id *string `mandatory:"true" json:"id"`

	// Details of the maintenance update package.
	Description *string `mandatory:"true" json:"description"`

	// The type of VM cluster maintenance update.
	UpdateType VmClusterUpdateSummaryUpdateTypeEnum `mandatory:"true" json:"updateType"`

	// The date and time the maintenance update was released.
	TimeReleased *common.SDKTime `mandatory:"true" json:"timeReleased"`

	// The version of the maintenance update package.
	Version *string `mandatory:"true" json:"version"`

	// The update action performed most recently using this maintenance update.
	LastAction VmClusterUpdateSummaryLastActionEnum `mandatory:"false" json:"lastAction,omitempty"`

	// The possible actions that can be performed using this maintenance update.
	AvailableActions []VmClusterUpdateSummaryAvailableActionsEnum `mandatory:"false" json:"availableActions,omitempty"`

	// Descriptive text providing additional details about the lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The current state of the maintenance update. Dependent on value of `lastAction`.
	LifecycleState VmClusterUpdateSummaryLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m VmClusterUpdateSummary) String() string {
	return common.PointerString(m)
}

// VmClusterUpdateSummaryLastActionEnum Enum with underlying type: string
type VmClusterUpdateSummaryLastActionEnum string

// Set of constants representing the allowable values for VmClusterUpdateSummaryLastActionEnum
const (
	VmClusterUpdateSummaryLastActionRollingApply VmClusterUpdateSummaryLastActionEnum = "ROLLING_APPLY"
	VmClusterUpdateSummaryLastActionPrecheck     VmClusterUpdateSummaryLastActionEnum = "PRECHECK"
	VmClusterUpdateSummaryLastActionRollback     VmClusterUpdateSummaryLastActionEnum = "ROLLBACK"
)

var mappingVmClusterUpdateSummaryLastAction = map[string]VmClusterUpdateSummaryLastActionEnum{
	"ROLLING_APPLY": VmClusterUpdateSummaryLastActionRollingApply,
	"PRECHECK":      VmClusterUpdateSummaryLastActionPrecheck,
	"ROLLBACK":      VmClusterUpdateSummaryLastActionRollback,
}

// GetVmClusterUpdateSummaryLastActionEnumValues Enumerates the set of values for VmClusterUpdateSummaryLastActionEnum
func GetVmClusterUpdateSummaryLastActionEnumValues() []VmClusterUpdateSummaryLastActionEnum {
	values := make([]VmClusterUpdateSummaryLastActionEnum, 0)
	for _, v := range mappingVmClusterUpdateSummaryLastAction {
		values = append(values, v)
	}
	return values
}

// VmClusterUpdateSummaryAvailableActionsEnum Enum with underlying type: string
type VmClusterUpdateSummaryAvailableActionsEnum string

// Set of constants representing the allowable values for VmClusterUpdateSummaryAvailableActionsEnum
const (
	VmClusterUpdateSummaryAvailableActionsRollingApply VmClusterUpdateSummaryAvailableActionsEnum = "ROLLING_APPLY"
	VmClusterUpdateSummaryAvailableActionsPrecheck     VmClusterUpdateSummaryAvailableActionsEnum = "PRECHECK"
	VmClusterUpdateSummaryAvailableActionsRollback     VmClusterUpdateSummaryAvailableActionsEnum = "ROLLBACK"
)

var mappingVmClusterUpdateSummaryAvailableActions = map[string]VmClusterUpdateSummaryAvailableActionsEnum{
	"ROLLING_APPLY": VmClusterUpdateSummaryAvailableActionsRollingApply,
	"PRECHECK":      VmClusterUpdateSummaryAvailableActionsPrecheck,
	"ROLLBACK":      VmClusterUpdateSummaryAvailableActionsRollback,
}

// GetVmClusterUpdateSummaryAvailableActionsEnumValues Enumerates the set of values for VmClusterUpdateSummaryAvailableActionsEnum
func GetVmClusterUpdateSummaryAvailableActionsEnumValues() []VmClusterUpdateSummaryAvailableActionsEnum {
	values := make([]VmClusterUpdateSummaryAvailableActionsEnum, 0)
	for _, v := range mappingVmClusterUpdateSummaryAvailableActions {
		values = append(values, v)
	}
	return values
}

// VmClusterUpdateSummaryUpdateTypeEnum Enum with underlying type: string
type VmClusterUpdateSummaryUpdateTypeEnum string

// Set of constants representing the allowable values for VmClusterUpdateSummaryUpdateTypeEnum
const (
	VmClusterUpdateSummaryUpdateTypeGiUpgrade VmClusterUpdateSummaryUpdateTypeEnum = "GI_UPGRADE"
	VmClusterUpdateSummaryUpdateTypeGiPatch   VmClusterUpdateSummaryUpdateTypeEnum = "GI_PATCH"
	VmClusterUpdateSummaryUpdateTypeOsUpdate  VmClusterUpdateSummaryUpdateTypeEnum = "OS_UPDATE"
)

var mappingVmClusterUpdateSummaryUpdateType = map[string]VmClusterUpdateSummaryUpdateTypeEnum{
	"GI_UPGRADE": VmClusterUpdateSummaryUpdateTypeGiUpgrade,
	"GI_PATCH":   VmClusterUpdateSummaryUpdateTypeGiPatch,
	"OS_UPDATE":  VmClusterUpdateSummaryUpdateTypeOsUpdate,
}

// GetVmClusterUpdateSummaryUpdateTypeEnumValues Enumerates the set of values for VmClusterUpdateSummaryUpdateTypeEnum
func GetVmClusterUpdateSummaryUpdateTypeEnumValues() []VmClusterUpdateSummaryUpdateTypeEnum {
	values := make([]VmClusterUpdateSummaryUpdateTypeEnum, 0)
	for _, v := range mappingVmClusterUpdateSummaryUpdateType {
		values = append(values, v)
	}
	return values
}

// VmClusterUpdateSummaryLifecycleStateEnum Enum with underlying type: string
type VmClusterUpdateSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for VmClusterUpdateSummaryLifecycleStateEnum
const (
	VmClusterUpdateSummaryLifecycleStateAvailable  VmClusterUpdateSummaryLifecycleStateEnum = "AVAILABLE"
	VmClusterUpdateSummaryLifecycleStateSuccess    VmClusterUpdateSummaryLifecycleStateEnum = "SUCCESS"
	VmClusterUpdateSummaryLifecycleStateInProgress VmClusterUpdateSummaryLifecycleStateEnum = "IN_PROGRESS"
	VmClusterUpdateSummaryLifecycleStateFailed     VmClusterUpdateSummaryLifecycleStateEnum = "FAILED"
)

var mappingVmClusterUpdateSummaryLifecycleState = map[string]VmClusterUpdateSummaryLifecycleStateEnum{
	"AVAILABLE":   VmClusterUpdateSummaryLifecycleStateAvailable,
	"SUCCESS":     VmClusterUpdateSummaryLifecycleStateSuccess,
	"IN_PROGRESS": VmClusterUpdateSummaryLifecycleStateInProgress,
	"FAILED":      VmClusterUpdateSummaryLifecycleStateFailed,
}

// GetVmClusterUpdateSummaryLifecycleStateEnumValues Enumerates the set of values for VmClusterUpdateSummaryLifecycleStateEnum
func GetVmClusterUpdateSummaryLifecycleStateEnumValues() []VmClusterUpdateSummaryLifecycleStateEnum {
	values := make([]VmClusterUpdateSummaryLifecycleStateEnum, 0)
	for _, v := range mappingVmClusterUpdateSummaryLifecycleState {
		values = append(values, v)
	}
	return values
}
