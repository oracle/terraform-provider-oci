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

// UpdateHistoryEntry The representation of UpdateHistoryEntry
type UpdateHistoryEntry struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the maintenance update history entry.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the maintenance update.
	UpdateId *string `mandatory:"true" json:"updateId"`

	// The type of cloud VM cluster maintenance update.
	UpdateType UpdateHistoryEntryUpdateTypeEnum `mandatory:"true" json:"updateType"`

	// The current lifecycle state of the maintenance update operation.
	LifecycleState UpdateHistoryEntryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time when the maintenance update action started.
	TimeStarted *common.SDKTime `mandatory:"true" json:"timeStarted"`

	// The update action.
	UpdateAction UpdateHistoryEntryUpdateActionEnum `mandatory:"false" json:"updateAction,omitempty"`

	// Descriptive text providing additional details about the lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The date and time when the maintenance update action completed.
	TimeCompleted *common.SDKTime `mandatory:"false" json:"timeCompleted"`
}

func (m UpdateHistoryEntry) String() string {
	return common.PointerString(m)
}

// UpdateHistoryEntryUpdateActionEnum Enum with underlying type: string
type UpdateHistoryEntryUpdateActionEnum string

// Set of constants representing the allowable values for UpdateHistoryEntryUpdateActionEnum
const (
	UpdateHistoryEntryUpdateActionRollingApply    UpdateHistoryEntryUpdateActionEnum = "ROLLING_APPLY"
	UpdateHistoryEntryUpdateActionNonRollingApply UpdateHistoryEntryUpdateActionEnum = "NON_ROLLING_APPLY"
	UpdateHistoryEntryUpdateActionPrecheck        UpdateHistoryEntryUpdateActionEnum = "PRECHECK"
	UpdateHistoryEntryUpdateActionRollback        UpdateHistoryEntryUpdateActionEnum = "ROLLBACK"
)

var mappingUpdateHistoryEntryUpdateAction = map[string]UpdateHistoryEntryUpdateActionEnum{
	"ROLLING_APPLY":     UpdateHistoryEntryUpdateActionRollingApply,
	"NON_ROLLING_APPLY": UpdateHistoryEntryUpdateActionNonRollingApply,
	"PRECHECK":          UpdateHistoryEntryUpdateActionPrecheck,
	"ROLLBACK":          UpdateHistoryEntryUpdateActionRollback,
}

// GetUpdateHistoryEntryUpdateActionEnumValues Enumerates the set of values for UpdateHistoryEntryUpdateActionEnum
func GetUpdateHistoryEntryUpdateActionEnumValues() []UpdateHistoryEntryUpdateActionEnum {
	values := make([]UpdateHistoryEntryUpdateActionEnum, 0)
	for _, v := range mappingUpdateHistoryEntryUpdateAction {
		values = append(values, v)
	}
	return values
}

// UpdateHistoryEntryUpdateTypeEnum Enum with underlying type: string
type UpdateHistoryEntryUpdateTypeEnum string

// Set of constants representing the allowable values for UpdateHistoryEntryUpdateTypeEnum
const (
	UpdateHistoryEntryUpdateTypeGiUpgrade UpdateHistoryEntryUpdateTypeEnum = "GI_UPGRADE"
	UpdateHistoryEntryUpdateTypeGiPatch   UpdateHistoryEntryUpdateTypeEnum = "GI_PATCH"
	UpdateHistoryEntryUpdateTypeOsUpdate  UpdateHistoryEntryUpdateTypeEnum = "OS_UPDATE"
)

var mappingUpdateHistoryEntryUpdateType = map[string]UpdateHistoryEntryUpdateTypeEnum{
	"GI_UPGRADE": UpdateHistoryEntryUpdateTypeGiUpgrade,
	"GI_PATCH":   UpdateHistoryEntryUpdateTypeGiPatch,
	"OS_UPDATE":  UpdateHistoryEntryUpdateTypeOsUpdate,
}

// GetUpdateHistoryEntryUpdateTypeEnumValues Enumerates the set of values for UpdateHistoryEntryUpdateTypeEnum
func GetUpdateHistoryEntryUpdateTypeEnumValues() []UpdateHistoryEntryUpdateTypeEnum {
	values := make([]UpdateHistoryEntryUpdateTypeEnum, 0)
	for _, v := range mappingUpdateHistoryEntryUpdateType {
		values = append(values, v)
	}
	return values
}

// UpdateHistoryEntryLifecycleStateEnum Enum with underlying type: string
type UpdateHistoryEntryLifecycleStateEnum string

// Set of constants representing the allowable values for UpdateHistoryEntryLifecycleStateEnum
const (
	UpdateHistoryEntryLifecycleStateInProgress UpdateHistoryEntryLifecycleStateEnum = "IN_PROGRESS"
	UpdateHistoryEntryLifecycleStateSucceeded  UpdateHistoryEntryLifecycleStateEnum = "SUCCEEDED"
	UpdateHistoryEntryLifecycleStateFailed     UpdateHistoryEntryLifecycleStateEnum = "FAILED"
)

var mappingUpdateHistoryEntryLifecycleState = map[string]UpdateHistoryEntryLifecycleStateEnum{
	"IN_PROGRESS": UpdateHistoryEntryLifecycleStateInProgress,
	"SUCCEEDED":   UpdateHistoryEntryLifecycleStateSucceeded,
	"FAILED":      UpdateHistoryEntryLifecycleStateFailed,
}

// GetUpdateHistoryEntryLifecycleStateEnumValues Enumerates the set of values for UpdateHistoryEntryLifecycleStateEnum
func GetUpdateHistoryEntryLifecycleStateEnumValues() []UpdateHistoryEntryLifecycleStateEnum {
	values := make([]UpdateHistoryEntryLifecycleStateEnum, 0)
	for _, v := range mappingUpdateHistoryEntryLifecycleState {
		values = append(values, v)
	}
	return values
}
