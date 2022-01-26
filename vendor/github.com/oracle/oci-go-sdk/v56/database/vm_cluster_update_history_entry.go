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

// VmClusterUpdateHistoryEntry The record of a maintenance update action performed on a specified VM cluster. Applies to Exadata Cloud@Customer instances only.
type VmClusterUpdateHistoryEntry struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the maintenance update history entry.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the maintenance update.
	UpdateId *string `mandatory:"true" json:"updateId"`

	// The type of VM cluster maintenance update.
	UpdateType VmClusterUpdateHistoryEntryUpdateTypeEnum `mandatory:"true" json:"updateType"`

	// The current lifecycle state of the maintenance update operation.
	LifecycleState VmClusterUpdateHistoryEntryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time when the maintenance update action started.
	TimeStarted *common.SDKTime `mandatory:"true" json:"timeStarted"`

	// The update action performed using this maintenance update.
	UpdateAction VmClusterUpdateHistoryEntryUpdateActionEnum `mandatory:"false" json:"updateAction,omitempty"`

	// Descriptive text providing additional details about the lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The date and time when the maintenance update action completed.
	TimeCompleted *common.SDKTime `mandatory:"false" json:"timeCompleted"`
}

func (m VmClusterUpdateHistoryEntry) String() string {
	return common.PointerString(m)
}

// VmClusterUpdateHistoryEntryUpdateActionEnum Enum with underlying type: string
type VmClusterUpdateHistoryEntryUpdateActionEnum string

// Set of constants representing the allowable values for VmClusterUpdateHistoryEntryUpdateActionEnum
const (
	VmClusterUpdateHistoryEntryUpdateActionRollingApply VmClusterUpdateHistoryEntryUpdateActionEnum = "ROLLING_APPLY"
	VmClusterUpdateHistoryEntryUpdateActionPrecheck     VmClusterUpdateHistoryEntryUpdateActionEnum = "PRECHECK"
	VmClusterUpdateHistoryEntryUpdateActionRollback     VmClusterUpdateHistoryEntryUpdateActionEnum = "ROLLBACK"
)

var mappingVmClusterUpdateHistoryEntryUpdateAction = map[string]VmClusterUpdateHistoryEntryUpdateActionEnum{
	"ROLLING_APPLY": VmClusterUpdateHistoryEntryUpdateActionRollingApply,
	"PRECHECK":      VmClusterUpdateHistoryEntryUpdateActionPrecheck,
	"ROLLBACK":      VmClusterUpdateHistoryEntryUpdateActionRollback,
}

// GetVmClusterUpdateHistoryEntryUpdateActionEnumValues Enumerates the set of values for VmClusterUpdateHistoryEntryUpdateActionEnum
func GetVmClusterUpdateHistoryEntryUpdateActionEnumValues() []VmClusterUpdateHistoryEntryUpdateActionEnum {
	values := make([]VmClusterUpdateHistoryEntryUpdateActionEnum, 0)
	for _, v := range mappingVmClusterUpdateHistoryEntryUpdateAction {
		values = append(values, v)
	}
	return values
}

// VmClusterUpdateHistoryEntryUpdateTypeEnum Enum with underlying type: string
type VmClusterUpdateHistoryEntryUpdateTypeEnum string

// Set of constants representing the allowable values for VmClusterUpdateHistoryEntryUpdateTypeEnum
const (
	VmClusterUpdateHistoryEntryUpdateTypeGiUpgrade VmClusterUpdateHistoryEntryUpdateTypeEnum = "GI_UPGRADE"
	VmClusterUpdateHistoryEntryUpdateTypeGiPatch   VmClusterUpdateHistoryEntryUpdateTypeEnum = "GI_PATCH"
	VmClusterUpdateHistoryEntryUpdateTypeOsUpdate  VmClusterUpdateHistoryEntryUpdateTypeEnum = "OS_UPDATE"
)

var mappingVmClusterUpdateHistoryEntryUpdateType = map[string]VmClusterUpdateHistoryEntryUpdateTypeEnum{
	"GI_UPGRADE": VmClusterUpdateHistoryEntryUpdateTypeGiUpgrade,
	"GI_PATCH":   VmClusterUpdateHistoryEntryUpdateTypeGiPatch,
	"OS_UPDATE":  VmClusterUpdateHistoryEntryUpdateTypeOsUpdate,
}

// GetVmClusterUpdateHistoryEntryUpdateTypeEnumValues Enumerates the set of values for VmClusterUpdateHistoryEntryUpdateTypeEnum
func GetVmClusterUpdateHistoryEntryUpdateTypeEnumValues() []VmClusterUpdateHistoryEntryUpdateTypeEnum {
	values := make([]VmClusterUpdateHistoryEntryUpdateTypeEnum, 0)
	for _, v := range mappingVmClusterUpdateHistoryEntryUpdateType {
		values = append(values, v)
	}
	return values
}

// VmClusterUpdateHistoryEntryLifecycleStateEnum Enum with underlying type: string
type VmClusterUpdateHistoryEntryLifecycleStateEnum string

// Set of constants representing the allowable values for VmClusterUpdateHistoryEntryLifecycleStateEnum
const (
	VmClusterUpdateHistoryEntryLifecycleStateInProgress VmClusterUpdateHistoryEntryLifecycleStateEnum = "IN_PROGRESS"
	VmClusterUpdateHistoryEntryLifecycleStateSucceeded  VmClusterUpdateHistoryEntryLifecycleStateEnum = "SUCCEEDED"
	VmClusterUpdateHistoryEntryLifecycleStateFailed     VmClusterUpdateHistoryEntryLifecycleStateEnum = "FAILED"
)

var mappingVmClusterUpdateHistoryEntryLifecycleState = map[string]VmClusterUpdateHistoryEntryLifecycleStateEnum{
	"IN_PROGRESS": VmClusterUpdateHistoryEntryLifecycleStateInProgress,
	"SUCCEEDED":   VmClusterUpdateHistoryEntryLifecycleStateSucceeded,
	"FAILED":      VmClusterUpdateHistoryEntryLifecycleStateFailed,
}

// GetVmClusterUpdateHistoryEntryLifecycleStateEnumValues Enumerates the set of values for VmClusterUpdateHistoryEntryLifecycleStateEnum
func GetVmClusterUpdateHistoryEntryLifecycleStateEnumValues() []VmClusterUpdateHistoryEntryLifecycleStateEnum {
	values := make([]VmClusterUpdateHistoryEntryLifecycleStateEnum, 0)
	for _, v := range mappingVmClusterUpdateHistoryEntryLifecycleState {
		values = append(values, v)
	}
	return values
}
