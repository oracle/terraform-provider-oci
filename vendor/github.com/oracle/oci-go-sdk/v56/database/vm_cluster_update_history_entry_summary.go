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

// VmClusterUpdateHistoryEntrySummary The record of a maintenance update action performed on a specified VM cluster. Applies to Exadata Cloud@Customer instances only.
type VmClusterUpdateHistoryEntrySummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the maintenance update history entry.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the maintenance update.
	UpdateId *string `mandatory:"true" json:"updateId"`

	// The type of VM cluster maintenance update.
	UpdateType VmClusterUpdateHistoryEntrySummaryUpdateTypeEnum `mandatory:"true" json:"updateType"`

	// The current lifecycle state of the maintenance update operation.
	LifecycleState VmClusterUpdateHistoryEntrySummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time when the maintenance update action started.
	TimeStarted *common.SDKTime `mandatory:"true" json:"timeStarted"`

	// The update action performed using this maintenance update.
	UpdateAction VmClusterUpdateHistoryEntrySummaryUpdateActionEnum `mandatory:"false" json:"updateAction,omitempty"`

	// Descriptive text providing additional details about the lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The date and time when the maintenance update action completed.
	TimeCompleted *common.SDKTime `mandatory:"false" json:"timeCompleted"`
}

func (m VmClusterUpdateHistoryEntrySummary) String() string {
	return common.PointerString(m)
}

// VmClusterUpdateHistoryEntrySummaryUpdateActionEnum Enum with underlying type: string
type VmClusterUpdateHistoryEntrySummaryUpdateActionEnum string

// Set of constants representing the allowable values for VmClusterUpdateHistoryEntrySummaryUpdateActionEnum
const (
	VmClusterUpdateHistoryEntrySummaryUpdateActionRollingApply VmClusterUpdateHistoryEntrySummaryUpdateActionEnum = "ROLLING_APPLY"
	VmClusterUpdateHistoryEntrySummaryUpdateActionPrecheck     VmClusterUpdateHistoryEntrySummaryUpdateActionEnum = "PRECHECK"
	VmClusterUpdateHistoryEntrySummaryUpdateActionRollback     VmClusterUpdateHistoryEntrySummaryUpdateActionEnum = "ROLLBACK"
)

var mappingVmClusterUpdateHistoryEntrySummaryUpdateAction = map[string]VmClusterUpdateHistoryEntrySummaryUpdateActionEnum{
	"ROLLING_APPLY": VmClusterUpdateHistoryEntrySummaryUpdateActionRollingApply,
	"PRECHECK":      VmClusterUpdateHistoryEntrySummaryUpdateActionPrecheck,
	"ROLLBACK":      VmClusterUpdateHistoryEntrySummaryUpdateActionRollback,
}

// GetVmClusterUpdateHistoryEntrySummaryUpdateActionEnumValues Enumerates the set of values for VmClusterUpdateHistoryEntrySummaryUpdateActionEnum
func GetVmClusterUpdateHistoryEntrySummaryUpdateActionEnumValues() []VmClusterUpdateHistoryEntrySummaryUpdateActionEnum {
	values := make([]VmClusterUpdateHistoryEntrySummaryUpdateActionEnum, 0)
	for _, v := range mappingVmClusterUpdateHistoryEntrySummaryUpdateAction {
		values = append(values, v)
	}
	return values
}

// VmClusterUpdateHistoryEntrySummaryUpdateTypeEnum Enum with underlying type: string
type VmClusterUpdateHistoryEntrySummaryUpdateTypeEnum string

// Set of constants representing the allowable values for VmClusterUpdateHistoryEntrySummaryUpdateTypeEnum
const (
	VmClusterUpdateHistoryEntrySummaryUpdateTypeGiUpgrade VmClusterUpdateHistoryEntrySummaryUpdateTypeEnum = "GI_UPGRADE"
	VmClusterUpdateHistoryEntrySummaryUpdateTypeGiPatch   VmClusterUpdateHistoryEntrySummaryUpdateTypeEnum = "GI_PATCH"
	VmClusterUpdateHistoryEntrySummaryUpdateTypeOsUpdate  VmClusterUpdateHistoryEntrySummaryUpdateTypeEnum = "OS_UPDATE"
)

var mappingVmClusterUpdateHistoryEntrySummaryUpdateType = map[string]VmClusterUpdateHistoryEntrySummaryUpdateTypeEnum{
	"GI_UPGRADE": VmClusterUpdateHistoryEntrySummaryUpdateTypeGiUpgrade,
	"GI_PATCH":   VmClusterUpdateHistoryEntrySummaryUpdateTypeGiPatch,
	"OS_UPDATE":  VmClusterUpdateHistoryEntrySummaryUpdateTypeOsUpdate,
}

// GetVmClusterUpdateHistoryEntrySummaryUpdateTypeEnumValues Enumerates the set of values for VmClusterUpdateHistoryEntrySummaryUpdateTypeEnum
func GetVmClusterUpdateHistoryEntrySummaryUpdateTypeEnumValues() []VmClusterUpdateHistoryEntrySummaryUpdateTypeEnum {
	values := make([]VmClusterUpdateHistoryEntrySummaryUpdateTypeEnum, 0)
	for _, v := range mappingVmClusterUpdateHistoryEntrySummaryUpdateType {
		values = append(values, v)
	}
	return values
}

// VmClusterUpdateHistoryEntrySummaryLifecycleStateEnum Enum with underlying type: string
type VmClusterUpdateHistoryEntrySummaryLifecycleStateEnum string

// Set of constants representing the allowable values for VmClusterUpdateHistoryEntrySummaryLifecycleStateEnum
const (
	VmClusterUpdateHistoryEntrySummaryLifecycleStateInProgress VmClusterUpdateHistoryEntrySummaryLifecycleStateEnum = "IN_PROGRESS"
	VmClusterUpdateHistoryEntrySummaryLifecycleStateSucceeded  VmClusterUpdateHistoryEntrySummaryLifecycleStateEnum = "SUCCEEDED"
	VmClusterUpdateHistoryEntrySummaryLifecycleStateFailed     VmClusterUpdateHistoryEntrySummaryLifecycleStateEnum = "FAILED"
)

var mappingVmClusterUpdateHistoryEntrySummaryLifecycleState = map[string]VmClusterUpdateHistoryEntrySummaryLifecycleStateEnum{
	"IN_PROGRESS": VmClusterUpdateHistoryEntrySummaryLifecycleStateInProgress,
	"SUCCEEDED":   VmClusterUpdateHistoryEntrySummaryLifecycleStateSucceeded,
	"FAILED":      VmClusterUpdateHistoryEntrySummaryLifecycleStateFailed,
}

// GetVmClusterUpdateHistoryEntrySummaryLifecycleStateEnumValues Enumerates the set of values for VmClusterUpdateHistoryEntrySummaryLifecycleStateEnum
func GetVmClusterUpdateHistoryEntrySummaryLifecycleStateEnumValues() []VmClusterUpdateHistoryEntrySummaryLifecycleStateEnum {
	values := make([]VmClusterUpdateHistoryEntrySummaryLifecycleStateEnum, 0)
	for _, v := range mappingVmClusterUpdateHistoryEntrySummaryLifecycleState {
		values = append(values, v)
	}
	return values
}
