// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateHistoryEntrySummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingUpdateHistoryEntrySummaryUpdateTypeEnum(string(m.UpdateType)); !ok && m.UpdateType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpdateType: %s. Supported values are: %s.", m.UpdateType, strings.Join(GetUpdateHistoryEntrySummaryUpdateTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUpdateHistoryEntrySummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetUpdateHistoryEntrySummaryLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingUpdateHistoryEntrySummaryUpdateActionEnum(string(m.UpdateAction)); !ok && m.UpdateAction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpdateAction: %s. Supported values are: %s.", m.UpdateAction, strings.Join(GetUpdateHistoryEntrySummaryUpdateActionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingUpdateHistoryEntrySummaryUpdateActionEnum = map[string]UpdateHistoryEntrySummaryUpdateActionEnum{
	"ROLLING_APPLY":     UpdateHistoryEntrySummaryUpdateActionRollingApply,
	"NON_ROLLING_APPLY": UpdateHistoryEntrySummaryUpdateActionNonRollingApply,
	"PRECHECK":          UpdateHistoryEntrySummaryUpdateActionPrecheck,
	"ROLLBACK":          UpdateHistoryEntrySummaryUpdateActionRollback,
}

var mappingUpdateHistoryEntrySummaryUpdateActionEnumLowerCase = map[string]UpdateHistoryEntrySummaryUpdateActionEnum{
	"rolling_apply":     UpdateHistoryEntrySummaryUpdateActionRollingApply,
	"non_rolling_apply": UpdateHistoryEntrySummaryUpdateActionNonRollingApply,
	"precheck":          UpdateHistoryEntrySummaryUpdateActionPrecheck,
	"rollback":          UpdateHistoryEntrySummaryUpdateActionRollback,
}

// GetUpdateHistoryEntrySummaryUpdateActionEnumValues Enumerates the set of values for UpdateHistoryEntrySummaryUpdateActionEnum
func GetUpdateHistoryEntrySummaryUpdateActionEnumValues() []UpdateHistoryEntrySummaryUpdateActionEnum {
	values := make([]UpdateHistoryEntrySummaryUpdateActionEnum, 0)
	for _, v := range mappingUpdateHistoryEntrySummaryUpdateActionEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateHistoryEntrySummaryUpdateActionEnumStringValues Enumerates the set of values in String for UpdateHistoryEntrySummaryUpdateActionEnum
func GetUpdateHistoryEntrySummaryUpdateActionEnumStringValues() []string {
	return []string{
		"ROLLING_APPLY",
		"NON_ROLLING_APPLY",
		"PRECHECK",
		"ROLLBACK",
	}
}

// GetMappingUpdateHistoryEntrySummaryUpdateActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateHistoryEntrySummaryUpdateActionEnum(val string) (UpdateHistoryEntrySummaryUpdateActionEnum, bool) {
	enum, ok := mappingUpdateHistoryEntrySummaryUpdateActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// UpdateHistoryEntrySummaryUpdateTypeEnum Enum with underlying type: string
type UpdateHistoryEntrySummaryUpdateTypeEnum string

// Set of constants representing the allowable values for UpdateHistoryEntrySummaryUpdateTypeEnum
const (
	UpdateHistoryEntrySummaryUpdateTypeGiUpgrade UpdateHistoryEntrySummaryUpdateTypeEnum = "GI_UPGRADE"
	UpdateHistoryEntrySummaryUpdateTypeGiPatch   UpdateHistoryEntrySummaryUpdateTypeEnum = "GI_PATCH"
	UpdateHistoryEntrySummaryUpdateTypeOsUpdate  UpdateHistoryEntrySummaryUpdateTypeEnum = "OS_UPDATE"
)

var mappingUpdateHistoryEntrySummaryUpdateTypeEnum = map[string]UpdateHistoryEntrySummaryUpdateTypeEnum{
	"GI_UPGRADE": UpdateHistoryEntrySummaryUpdateTypeGiUpgrade,
	"GI_PATCH":   UpdateHistoryEntrySummaryUpdateTypeGiPatch,
	"OS_UPDATE":  UpdateHistoryEntrySummaryUpdateTypeOsUpdate,
}

var mappingUpdateHistoryEntrySummaryUpdateTypeEnumLowerCase = map[string]UpdateHistoryEntrySummaryUpdateTypeEnum{
	"gi_upgrade": UpdateHistoryEntrySummaryUpdateTypeGiUpgrade,
	"gi_patch":   UpdateHistoryEntrySummaryUpdateTypeGiPatch,
	"os_update":  UpdateHistoryEntrySummaryUpdateTypeOsUpdate,
}

// GetUpdateHistoryEntrySummaryUpdateTypeEnumValues Enumerates the set of values for UpdateHistoryEntrySummaryUpdateTypeEnum
func GetUpdateHistoryEntrySummaryUpdateTypeEnumValues() []UpdateHistoryEntrySummaryUpdateTypeEnum {
	values := make([]UpdateHistoryEntrySummaryUpdateTypeEnum, 0)
	for _, v := range mappingUpdateHistoryEntrySummaryUpdateTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateHistoryEntrySummaryUpdateTypeEnumStringValues Enumerates the set of values in String for UpdateHistoryEntrySummaryUpdateTypeEnum
func GetUpdateHistoryEntrySummaryUpdateTypeEnumStringValues() []string {
	return []string{
		"GI_UPGRADE",
		"GI_PATCH",
		"OS_UPDATE",
	}
}

// GetMappingUpdateHistoryEntrySummaryUpdateTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateHistoryEntrySummaryUpdateTypeEnum(val string) (UpdateHistoryEntrySummaryUpdateTypeEnum, bool) {
	enum, ok := mappingUpdateHistoryEntrySummaryUpdateTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// UpdateHistoryEntrySummaryLifecycleStateEnum Enum with underlying type: string
type UpdateHistoryEntrySummaryLifecycleStateEnum string

// Set of constants representing the allowable values for UpdateHistoryEntrySummaryLifecycleStateEnum
const (
	UpdateHistoryEntrySummaryLifecycleStateInProgress UpdateHistoryEntrySummaryLifecycleStateEnum = "IN_PROGRESS"
	UpdateHistoryEntrySummaryLifecycleStateSucceeded  UpdateHistoryEntrySummaryLifecycleStateEnum = "SUCCEEDED"
	UpdateHistoryEntrySummaryLifecycleStateFailed     UpdateHistoryEntrySummaryLifecycleStateEnum = "FAILED"
)

var mappingUpdateHistoryEntrySummaryLifecycleStateEnum = map[string]UpdateHistoryEntrySummaryLifecycleStateEnum{
	"IN_PROGRESS": UpdateHistoryEntrySummaryLifecycleStateInProgress,
	"SUCCEEDED":   UpdateHistoryEntrySummaryLifecycleStateSucceeded,
	"FAILED":      UpdateHistoryEntrySummaryLifecycleStateFailed,
}

var mappingUpdateHistoryEntrySummaryLifecycleStateEnumLowerCase = map[string]UpdateHistoryEntrySummaryLifecycleStateEnum{
	"in_progress": UpdateHistoryEntrySummaryLifecycleStateInProgress,
	"succeeded":   UpdateHistoryEntrySummaryLifecycleStateSucceeded,
	"failed":      UpdateHistoryEntrySummaryLifecycleStateFailed,
}

// GetUpdateHistoryEntrySummaryLifecycleStateEnumValues Enumerates the set of values for UpdateHistoryEntrySummaryLifecycleStateEnum
func GetUpdateHistoryEntrySummaryLifecycleStateEnumValues() []UpdateHistoryEntrySummaryLifecycleStateEnum {
	values := make([]UpdateHistoryEntrySummaryLifecycleStateEnum, 0)
	for _, v := range mappingUpdateHistoryEntrySummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateHistoryEntrySummaryLifecycleStateEnumStringValues Enumerates the set of values in String for UpdateHistoryEntrySummaryLifecycleStateEnum
func GetUpdateHistoryEntrySummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"IN_PROGRESS",
		"SUCCEEDED",
		"FAILED",
	}
}

// GetMappingUpdateHistoryEntrySummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateHistoryEntrySummaryLifecycleStateEnum(val string) (UpdateHistoryEntrySummaryLifecycleStateEnum, bool) {
	enum, ok := mappingUpdateHistoryEntrySummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
