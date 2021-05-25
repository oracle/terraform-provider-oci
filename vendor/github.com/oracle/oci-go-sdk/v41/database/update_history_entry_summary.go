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

// UpdateHistoryEntrySummary The record of an maintenance update action on a specified cloud VM cluster. Applies to Exadata Cloud Service instances only.
type UpdateHistoryEntrySummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the maintenance update history entry.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the maintenance update.
	UpdateId *string `mandatory:"true" json:"updateId"`

	// The type of cloud VM cluster maintenance update.
	UpdateType UpdateHistoryEntrySummaryUpdateTypeEnum `mandatory:"true" json:"updateType"`

	// The current lifecycle state of the maintenance update operation.
	LifecycleState UpdateHistoryEntrySummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time when the maintenance update action started.
	TimeStarted *common.SDKTime `mandatory:"true" json:"timeStarted"`

	// The update action.
	UpdateAction UpdateHistoryEntrySummaryUpdateActionEnum `mandatory:"false" json:"updateAction,omitempty"`

	// Descriptive text providing additional details about the lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The date and time when the maintenance update action completed.
	TimeCompleted *common.SDKTime `mandatory:"false" json:"timeCompleted"`
}

func (m UpdateHistoryEntrySummary) String() string {
	return common.PointerString(m)
}

// UpdateHistoryEntrySummaryUpdateActionEnum Enum with underlying type: string
type UpdateHistoryEntrySummaryUpdateActionEnum string

// Set of constants representing the allowable values for UpdateHistoryEntrySummaryUpdateActionEnum
const (
	UpdateHistoryEntrySummaryUpdateActionRollingApply    UpdateHistoryEntrySummaryUpdateActionEnum = "ROLLING_APPLY"
	UpdateHistoryEntrySummaryUpdateActionNonRollingApply UpdateHistoryEntrySummaryUpdateActionEnum = "NON_ROLLING_APPLY"
	UpdateHistoryEntrySummaryUpdateActionPrecheck        UpdateHistoryEntrySummaryUpdateActionEnum = "PRECHECK"
	UpdateHistoryEntrySummaryUpdateActionRollback        UpdateHistoryEntrySummaryUpdateActionEnum = "ROLLBACK"
)

var mappingUpdateHistoryEntrySummaryUpdateAction = map[string]UpdateHistoryEntrySummaryUpdateActionEnum{
	"ROLLING_APPLY":     UpdateHistoryEntrySummaryUpdateActionRollingApply,
	"NON_ROLLING_APPLY": UpdateHistoryEntrySummaryUpdateActionNonRollingApply,
	"PRECHECK":          UpdateHistoryEntrySummaryUpdateActionPrecheck,
	"ROLLBACK":          UpdateHistoryEntrySummaryUpdateActionRollback,
}

// GetUpdateHistoryEntrySummaryUpdateActionEnumValues Enumerates the set of values for UpdateHistoryEntrySummaryUpdateActionEnum
func GetUpdateHistoryEntrySummaryUpdateActionEnumValues() []UpdateHistoryEntrySummaryUpdateActionEnum {
	values := make([]UpdateHistoryEntrySummaryUpdateActionEnum, 0)
	for _, v := range mappingUpdateHistoryEntrySummaryUpdateAction {
		values = append(values, v)
	}
	return values
}

// UpdateHistoryEntrySummaryUpdateTypeEnum Enum with underlying type: string
type UpdateHistoryEntrySummaryUpdateTypeEnum string

// Set of constants representing the allowable values for UpdateHistoryEntrySummaryUpdateTypeEnum
const (
	UpdateHistoryEntrySummaryUpdateTypeGiUpgrade UpdateHistoryEntrySummaryUpdateTypeEnum = "GI_UPGRADE"
	UpdateHistoryEntrySummaryUpdateTypeGiPatch   UpdateHistoryEntrySummaryUpdateTypeEnum = "GI_PATCH"
	UpdateHistoryEntrySummaryUpdateTypeOsUpdate  UpdateHistoryEntrySummaryUpdateTypeEnum = "OS_UPDATE"
)

var mappingUpdateHistoryEntrySummaryUpdateType = map[string]UpdateHistoryEntrySummaryUpdateTypeEnum{
	"GI_UPGRADE": UpdateHistoryEntrySummaryUpdateTypeGiUpgrade,
	"GI_PATCH":   UpdateHistoryEntrySummaryUpdateTypeGiPatch,
	"OS_UPDATE":  UpdateHistoryEntrySummaryUpdateTypeOsUpdate,
}

// GetUpdateHistoryEntrySummaryUpdateTypeEnumValues Enumerates the set of values for UpdateHistoryEntrySummaryUpdateTypeEnum
func GetUpdateHistoryEntrySummaryUpdateTypeEnumValues() []UpdateHistoryEntrySummaryUpdateTypeEnum {
	values := make([]UpdateHistoryEntrySummaryUpdateTypeEnum, 0)
	for _, v := range mappingUpdateHistoryEntrySummaryUpdateType {
		values = append(values, v)
	}
	return values
}

// UpdateHistoryEntrySummaryLifecycleStateEnum Enum with underlying type: string
type UpdateHistoryEntrySummaryLifecycleStateEnum string

// Set of constants representing the allowable values for UpdateHistoryEntrySummaryLifecycleStateEnum
const (
	UpdateHistoryEntrySummaryLifecycleStateInProgress UpdateHistoryEntrySummaryLifecycleStateEnum = "IN_PROGRESS"
	UpdateHistoryEntrySummaryLifecycleStateSucceeded  UpdateHistoryEntrySummaryLifecycleStateEnum = "SUCCEEDED"
	UpdateHistoryEntrySummaryLifecycleStateFailed     UpdateHistoryEntrySummaryLifecycleStateEnum = "FAILED"
)

var mappingUpdateHistoryEntrySummaryLifecycleState = map[string]UpdateHistoryEntrySummaryLifecycleStateEnum{
	"IN_PROGRESS": UpdateHistoryEntrySummaryLifecycleStateInProgress,
	"SUCCEEDED":   UpdateHistoryEntrySummaryLifecycleStateSucceeded,
	"FAILED":      UpdateHistoryEntrySummaryLifecycleStateFailed,
}

// GetUpdateHistoryEntrySummaryLifecycleStateEnumValues Enumerates the set of values for UpdateHistoryEntrySummaryLifecycleStateEnum
func GetUpdateHistoryEntrySummaryLifecycleStateEnumValues() []UpdateHistoryEntrySummaryLifecycleStateEnum {
	values := make([]UpdateHistoryEntrySummaryLifecycleStateEnum, 0)
	for _, v := range mappingUpdateHistoryEntrySummaryLifecycleState {
		values = append(values, v)
	}
	return values
}
