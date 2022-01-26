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

// VmClusterUpdate A maintenance update for a VM cluster. Applies to Exadata Cloud@Customer instances only.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access,
// see Getting Started with Policies (https://docs.cloud.oracle.com/Content/Identity/Concepts/policygetstarted.htm).
type VmClusterUpdate struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the maintenance update.
	Id *string `mandatory:"true" json:"id"`

	// Details of the maintenance update package.
	Description *string `mandatory:"true" json:"description"`

	// The type of VM cluster maintenance update.
	UpdateType VmClusterUpdateUpdateTypeEnum `mandatory:"true" json:"updateType"`

	// The date and time the maintenance update was released.
	TimeReleased *common.SDKTime `mandatory:"true" json:"timeReleased"`

	// The version of the maintenance update package.
	Version *string `mandatory:"true" json:"version"`

	// The update action performed most recently using this maintenance update.
	LastAction VmClusterUpdateLastActionEnum `mandatory:"false" json:"lastAction,omitempty"`

	// The possible actions that can be performed using this maintenance update.
	AvailableActions []VmClusterUpdateAvailableActionsEnum `mandatory:"false" json:"availableActions,omitempty"`

	// Descriptive text providing additional details about the lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The current state of the maintenance update. Dependent on value of `lastAction`.
	LifecycleState VmClusterUpdateLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m VmClusterUpdate) String() string {
	return common.PointerString(m)
}

// VmClusterUpdateLastActionEnum Enum with underlying type: string
type VmClusterUpdateLastActionEnum string

// Set of constants representing the allowable values for VmClusterUpdateLastActionEnum
const (
	VmClusterUpdateLastActionRollingApply VmClusterUpdateLastActionEnum = "ROLLING_APPLY"
	VmClusterUpdateLastActionPrecheck     VmClusterUpdateLastActionEnum = "PRECHECK"
	VmClusterUpdateLastActionRollback     VmClusterUpdateLastActionEnum = "ROLLBACK"
)

var mappingVmClusterUpdateLastAction = map[string]VmClusterUpdateLastActionEnum{
	"ROLLING_APPLY": VmClusterUpdateLastActionRollingApply,
	"PRECHECK":      VmClusterUpdateLastActionPrecheck,
	"ROLLBACK":      VmClusterUpdateLastActionRollback,
}

// GetVmClusterUpdateLastActionEnumValues Enumerates the set of values for VmClusterUpdateLastActionEnum
func GetVmClusterUpdateLastActionEnumValues() []VmClusterUpdateLastActionEnum {
	values := make([]VmClusterUpdateLastActionEnum, 0)
	for _, v := range mappingVmClusterUpdateLastAction {
		values = append(values, v)
	}
	return values
}

// VmClusterUpdateAvailableActionsEnum Enum with underlying type: string
type VmClusterUpdateAvailableActionsEnum string

// Set of constants representing the allowable values for VmClusterUpdateAvailableActionsEnum
const (
	VmClusterUpdateAvailableActionsRollingApply VmClusterUpdateAvailableActionsEnum = "ROLLING_APPLY"
	VmClusterUpdateAvailableActionsPrecheck     VmClusterUpdateAvailableActionsEnum = "PRECHECK"
	VmClusterUpdateAvailableActionsRollback     VmClusterUpdateAvailableActionsEnum = "ROLLBACK"
)

var mappingVmClusterUpdateAvailableActions = map[string]VmClusterUpdateAvailableActionsEnum{
	"ROLLING_APPLY": VmClusterUpdateAvailableActionsRollingApply,
	"PRECHECK":      VmClusterUpdateAvailableActionsPrecheck,
	"ROLLBACK":      VmClusterUpdateAvailableActionsRollback,
}

// GetVmClusterUpdateAvailableActionsEnumValues Enumerates the set of values for VmClusterUpdateAvailableActionsEnum
func GetVmClusterUpdateAvailableActionsEnumValues() []VmClusterUpdateAvailableActionsEnum {
	values := make([]VmClusterUpdateAvailableActionsEnum, 0)
	for _, v := range mappingVmClusterUpdateAvailableActions {
		values = append(values, v)
	}
	return values
}

// VmClusterUpdateUpdateTypeEnum Enum with underlying type: string
type VmClusterUpdateUpdateTypeEnum string

// Set of constants representing the allowable values for VmClusterUpdateUpdateTypeEnum
const (
	VmClusterUpdateUpdateTypeGiUpgrade VmClusterUpdateUpdateTypeEnum = "GI_UPGRADE"
	VmClusterUpdateUpdateTypeGiPatch   VmClusterUpdateUpdateTypeEnum = "GI_PATCH"
	VmClusterUpdateUpdateTypeOsUpdate  VmClusterUpdateUpdateTypeEnum = "OS_UPDATE"
)

var mappingVmClusterUpdateUpdateType = map[string]VmClusterUpdateUpdateTypeEnum{
	"GI_UPGRADE": VmClusterUpdateUpdateTypeGiUpgrade,
	"GI_PATCH":   VmClusterUpdateUpdateTypeGiPatch,
	"OS_UPDATE":  VmClusterUpdateUpdateTypeOsUpdate,
}

// GetVmClusterUpdateUpdateTypeEnumValues Enumerates the set of values for VmClusterUpdateUpdateTypeEnum
func GetVmClusterUpdateUpdateTypeEnumValues() []VmClusterUpdateUpdateTypeEnum {
	values := make([]VmClusterUpdateUpdateTypeEnum, 0)
	for _, v := range mappingVmClusterUpdateUpdateType {
		values = append(values, v)
	}
	return values
}

// VmClusterUpdateLifecycleStateEnum Enum with underlying type: string
type VmClusterUpdateLifecycleStateEnum string

// Set of constants representing the allowable values for VmClusterUpdateLifecycleStateEnum
const (
	VmClusterUpdateLifecycleStateAvailable  VmClusterUpdateLifecycleStateEnum = "AVAILABLE"
	VmClusterUpdateLifecycleStateSuccess    VmClusterUpdateLifecycleStateEnum = "SUCCESS"
	VmClusterUpdateLifecycleStateInProgress VmClusterUpdateLifecycleStateEnum = "IN_PROGRESS"
	VmClusterUpdateLifecycleStateFailed     VmClusterUpdateLifecycleStateEnum = "FAILED"
)

var mappingVmClusterUpdateLifecycleState = map[string]VmClusterUpdateLifecycleStateEnum{
	"AVAILABLE":   VmClusterUpdateLifecycleStateAvailable,
	"SUCCESS":     VmClusterUpdateLifecycleStateSuccess,
	"IN_PROGRESS": VmClusterUpdateLifecycleStateInProgress,
	"FAILED":      VmClusterUpdateLifecycleStateFailed,
}

// GetVmClusterUpdateLifecycleStateEnumValues Enumerates the set of values for VmClusterUpdateLifecycleStateEnum
func GetVmClusterUpdateLifecycleStateEnumValues() []VmClusterUpdateLifecycleStateEnum {
	values := make([]VmClusterUpdateLifecycleStateEnum, 0)
	for _, v := range mappingVmClusterUpdateLifecycleState {
		values = append(values, v)
	}
	return values
}
