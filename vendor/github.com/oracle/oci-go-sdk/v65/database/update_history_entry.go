// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateHistoryEntry) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingUpdateHistoryEntryUpdateTypeEnum(string(m.UpdateType)); !ok && m.UpdateType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpdateType: %s. Supported values are: %s.", m.UpdateType, strings.Join(GetUpdateHistoryEntryUpdateTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUpdateHistoryEntryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetUpdateHistoryEntryLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingUpdateHistoryEntryUpdateActionEnum(string(m.UpdateAction)); !ok && m.UpdateAction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpdateAction: %s. Supported values are: %s.", m.UpdateAction, strings.Join(GetUpdateHistoryEntryUpdateActionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingUpdateHistoryEntryUpdateActionEnum = map[string]UpdateHistoryEntryUpdateActionEnum{
	"ROLLING_APPLY":     UpdateHistoryEntryUpdateActionRollingApply,
	"NON_ROLLING_APPLY": UpdateHistoryEntryUpdateActionNonRollingApply,
	"PRECHECK":          UpdateHistoryEntryUpdateActionPrecheck,
	"ROLLBACK":          UpdateHistoryEntryUpdateActionRollback,
}

var mappingUpdateHistoryEntryUpdateActionEnumLowerCase = map[string]UpdateHistoryEntryUpdateActionEnum{
	"rolling_apply":     UpdateHistoryEntryUpdateActionRollingApply,
	"non_rolling_apply": UpdateHistoryEntryUpdateActionNonRollingApply,
	"precheck":          UpdateHistoryEntryUpdateActionPrecheck,
	"rollback":          UpdateHistoryEntryUpdateActionRollback,
}

// GetUpdateHistoryEntryUpdateActionEnumValues Enumerates the set of values for UpdateHistoryEntryUpdateActionEnum
func GetUpdateHistoryEntryUpdateActionEnumValues() []UpdateHistoryEntryUpdateActionEnum {
	values := make([]UpdateHistoryEntryUpdateActionEnum, 0)
	for _, v := range mappingUpdateHistoryEntryUpdateActionEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateHistoryEntryUpdateActionEnumStringValues Enumerates the set of values in String for UpdateHistoryEntryUpdateActionEnum
func GetUpdateHistoryEntryUpdateActionEnumStringValues() []string {
	return []string{
		"ROLLING_APPLY",
		"NON_ROLLING_APPLY",
		"PRECHECK",
		"ROLLBACK",
	}
}

// GetMappingUpdateHistoryEntryUpdateActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateHistoryEntryUpdateActionEnum(val string) (UpdateHistoryEntryUpdateActionEnum, bool) {
	enum, ok := mappingUpdateHistoryEntryUpdateActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// UpdateHistoryEntryUpdateTypeEnum Enum with underlying type: string
type UpdateHistoryEntryUpdateTypeEnum string

// Set of constants representing the allowable values for UpdateHistoryEntryUpdateTypeEnum
const (
	UpdateHistoryEntryUpdateTypeGiUpgrade UpdateHistoryEntryUpdateTypeEnum = "GI_UPGRADE"
	UpdateHistoryEntryUpdateTypeGiPatch   UpdateHistoryEntryUpdateTypeEnum = "GI_PATCH"
	UpdateHistoryEntryUpdateTypeOsUpdate  UpdateHistoryEntryUpdateTypeEnum = "OS_UPDATE"
)

var mappingUpdateHistoryEntryUpdateTypeEnum = map[string]UpdateHistoryEntryUpdateTypeEnum{
	"GI_UPGRADE": UpdateHistoryEntryUpdateTypeGiUpgrade,
	"GI_PATCH":   UpdateHistoryEntryUpdateTypeGiPatch,
	"OS_UPDATE":  UpdateHistoryEntryUpdateTypeOsUpdate,
}

var mappingUpdateHistoryEntryUpdateTypeEnumLowerCase = map[string]UpdateHistoryEntryUpdateTypeEnum{
	"gi_upgrade": UpdateHistoryEntryUpdateTypeGiUpgrade,
	"gi_patch":   UpdateHistoryEntryUpdateTypeGiPatch,
	"os_update":  UpdateHistoryEntryUpdateTypeOsUpdate,
}

// GetUpdateHistoryEntryUpdateTypeEnumValues Enumerates the set of values for UpdateHistoryEntryUpdateTypeEnum
func GetUpdateHistoryEntryUpdateTypeEnumValues() []UpdateHistoryEntryUpdateTypeEnum {
	values := make([]UpdateHistoryEntryUpdateTypeEnum, 0)
	for _, v := range mappingUpdateHistoryEntryUpdateTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateHistoryEntryUpdateTypeEnumStringValues Enumerates the set of values in String for UpdateHistoryEntryUpdateTypeEnum
func GetUpdateHistoryEntryUpdateTypeEnumStringValues() []string {
	return []string{
		"GI_UPGRADE",
		"GI_PATCH",
		"OS_UPDATE",
	}
}

// GetMappingUpdateHistoryEntryUpdateTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateHistoryEntryUpdateTypeEnum(val string) (UpdateHistoryEntryUpdateTypeEnum, bool) {
	enum, ok := mappingUpdateHistoryEntryUpdateTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// UpdateHistoryEntryLifecycleStateEnum Enum with underlying type: string
type UpdateHistoryEntryLifecycleStateEnum string

// Set of constants representing the allowable values for UpdateHistoryEntryLifecycleStateEnum
const (
	UpdateHistoryEntryLifecycleStateInProgress UpdateHistoryEntryLifecycleStateEnum = "IN_PROGRESS"
	UpdateHistoryEntryLifecycleStateSucceeded  UpdateHistoryEntryLifecycleStateEnum = "SUCCEEDED"
	UpdateHistoryEntryLifecycleStateFailed     UpdateHistoryEntryLifecycleStateEnum = "FAILED"
)

var mappingUpdateHistoryEntryLifecycleStateEnum = map[string]UpdateHistoryEntryLifecycleStateEnum{
	"IN_PROGRESS": UpdateHistoryEntryLifecycleStateInProgress,
	"SUCCEEDED":   UpdateHistoryEntryLifecycleStateSucceeded,
	"FAILED":      UpdateHistoryEntryLifecycleStateFailed,
}

var mappingUpdateHistoryEntryLifecycleStateEnumLowerCase = map[string]UpdateHistoryEntryLifecycleStateEnum{
	"in_progress": UpdateHistoryEntryLifecycleStateInProgress,
	"succeeded":   UpdateHistoryEntryLifecycleStateSucceeded,
	"failed":      UpdateHistoryEntryLifecycleStateFailed,
}

// GetUpdateHistoryEntryLifecycleStateEnumValues Enumerates the set of values for UpdateHistoryEntryLifecycleStateEnum
func GetUpdateHistoryEntryLifecycleStateEnumValues() []UpdateHistoryEntryLifecycleStateEnum {
	values := make([]UpdateHistoryEntryLifecycleStateEnum, 0)
	for _, v := range mappingUpdateHistoryEntryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateHistoryEntryLifecycleStateEnumStringValues Enumerates the set of values in String for UpdateHistoryEntryLifecycleStateEnum
func GetUpdateHistoryEntryLifecycleStateEnumStringValues() []string {
	return []string{
		"IN_PROGRESS",
		"SUCCEEDED",
		"FAILED",
	}
}

// GetMappingUpdateHistoryEntryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateHistoryEntryLifecycleStateEnum(val string) (UpdateHistoryEntryLifecycleStateEnum, bool) {
	enum, ok := mappingUpdateHistoryEntryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
