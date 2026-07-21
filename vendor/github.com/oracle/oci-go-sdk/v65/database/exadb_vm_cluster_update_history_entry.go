// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExadbVmClusterUpdateHistoryEntry The record of an maintenance update action on a specified Exadata VM cluster on Exascale Infrastructure.
type ExadbVmClusterUpdateHistoryEntry struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the maintenance update history entry.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the maintenance update.
	UpdateId *string `mandatory:"true" json:"updateId"`

	// The type of cloud VM cluster maintenance update.
	UpdateType ExadbVmClusterUpdateHistoryEntryUpdateTypeEnum `mandatory:"true" json:"updateType"`

	// The current lifecycle state of the maintenance update operation.
	LifecycleState ExadbVmClusterUpdateHistoryEntryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time when the maintenance update action started.
	TimeStarted *common.SDKTime `mandatory:"true" json:"timeStarted"`

	// The update action.
	UpdateAction ExadbVmClusterUpdateHistoryEntryUpdateActionEnum `mandatory:"false" json:"updateAction,omitempty"`

	// The OS update mode performed using this maintenance update.
	UpdateMode ExadbVmClusterUpdateHistoryEntryUpdateModeEnum `mandatory:"false" json:"updateMode,omitempty"`

	// Descriptive text providing additional details about the lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The date and time when the maintenance update action completed.
	TimeCompleted *common.SDKTime `mandatory:"false" json:"timeCompleted"`

	// The version of the maintenance update package.
	Version *string `mandatory:"false" json:"version"`
}

func (m ExadbVmClusterUpdateHistoryEntry) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExadbVmClusterUpdateHistoryEntry) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExadbVmClusterUpdateHistoryEntryUpdateTypeEnum(string(m.UpdateType)); !ok && m.UpdateType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpdateType: %s. Supported values are: %s.", m.UpdateType, strings.Join(GetExadbVmClusterUpdateHistoryEntryUpdateTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingExadbVmClusterUpdateHistoryEntryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetExadbVmClusterUpdateHistoryEntryLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingExadbVmClusterUpdateHistoryEntryUpdateActionEnum(string(m.UpdateAction)); !ok && m.UpdateAction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpdateAction: %s. Supported values are: %s.", m.UpdateAction, strings.Join(GetExadbVmClusterUpdateHistoryEntryUpdateActionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingExadbVmClusterUpdateHistoryEntryUpdateModeEnum(string(m.UpdateMode)); !ok && m.UpdateMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpdateMode: %s. Supported values are: %s.", m.UpdateMode, strings.Join(GetExadbVmClusterUpdateHistoryEntryUpdateModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExadbVmClusterUpdateHistoryEntryUpdateActionEnum Enum with underlying type: string
type ExadbVmClusterUpdateHistoryEntryUpdateActionEnum string

// Set of constants representing the allowable values for ExadbVmClusterUpdateHistoryEntryUpdateActionEnum
const (
	ExadbVmClusterUpdateHistoryEntryUpdateActionRollingApply    ExadbVmClusterUpdateHistoryEntryUpdateActionEnum = "ROLLING_APPLY"
	ExadbVmClusterUpdateHistoryEntryUpdateActionNonRollingApply ExadbVmClusterUpdateHistoryEntryUpdateActionEnum = "NON_ROLLING_APPLY"
	ExadbVmClusterUpdateHistoryEntryUpdateActionPrecheck        ExadbVmClusterUpdateHistoryEntryUpdateActionEnum = "PRECHECK"
	ExadbVmClusterUpdateHistoryEntryUpdateActionRollback        ExadbVmClusterUpdateHistoryEntryUpdateActionEnum = "ROLLBACK"
)

var mappingExadbVmClusterUpdateHistoryEntryUpdateActionEnum = map[string]ExadbVmClusterUpdateHistoryEntryUpdateActionEnum{
	"ROLLING_APPLY":     ExadbVmClusterUpdateHistoryEntryUpdateActionRollingApply,
	"NON_ROLLING_APPLY": ExadbVmClusterUpdateHistoryEntryUpdateActionNonRollingApply,
	"PRECHECK":          ExadbVmClusterUpdateHistoryEntryUpdateActionPrecheck,
	"ROLLBACK":          ExadbVmClusterUpdateHistoryEntryUpdateActionRollback,
}

var mappingExadbVmClusterUpdateHistoryEntryUpdateActionEnumLowerCase = map[string]ExadbVmClusterUpdateHistoryEntryUpdateActionEnum{
	"rolling_apply":     ExadbVmClusterUpdateHistoryEntryUpdateActionRollingApply,
	"non_rolling_apply": ExadbVmClusterUpdateHistoryEntryUpdateActionNonRollingApply,
	"precheck":          ExadbVmClusterUpdateHistoryEntryUpdateActionPrecheck,
	"rollback":          ExadbVmClusterUpdateHistoryEntryUpdateActionRollback,
}

// GetExadbVmClusterUpdateHistoryEntryUpdateActionEnumValues Enumerates the set of values for ExadbVmClusterUpdateHistoryEntryUpdateActionEnum
func GetExadbVmClusterUpdateHistoryEntryUpdateActionEnumValues() []ExadbVmClusterUpdateHistoryEntryUpdateActionEnum {
	values := make([]ExadbVmClusterUpdateHistoryEntryUpdateActionEnum, 0)
	for _, v := range mappingExadbVmClusterUpdateHistoryEntryUpdateActionEnum {
		values = append(values, v)
	}
	return values
}

// GetExadbVmClusterUpdateHistoryEntryUpdateActionEnumStringValues Enumerates the set of values in String for ExadbVmClusterUpdateHistoryEntryUpdateActionEnum
func GetExadbVmClusterUpdateHistoryEntryUpdateActionEnumStringValues() []string {
	return []string{
		"ROLLING_APPLY",
		"NON_ROLLING_APPLY",
		"PRECHECK",
		"ROLLBACK",
	}
}

// GetMappingExadbVmClusterUpdateHistoryEntryUpdateActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExadbVmClusterUpdateHistoryEntryUpdateActionEnum(val string) (ExadbVmClusterUpdateHistoryEntryUpdateActionEnum, bool) {
	enum, ok := mappingExadbVmClusterUpdateHistoryEntryUpdateActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ExadbVmClusterUpdateHistoryEntryUpdateModeEnum Enum with underlying type: string
type ExadbVmClusterUpdateHistoryEntryUpdateModeEnum string

// Set of constants representing the allowable values for ExadbVmClusterUpdateHistoryEntryUpdateModeEnum
const (
	ExadbVmClusterUpdateHistoryEntryUpdateModeOnlineHighcvss         ExadbVmClusterUpdateHistoryEntryUpdateModeEnum = "ONLINE_HIGHCVSS"
	ExadbVmClusterUpdateHistoryEntryUpdateModeOnlineAllcvss          ExadbVmClusterUpdateHistoryEntryUpdateModeEnum = "ONLINE_ALLCVSS"
	ExadbVmClusterUpdateHistoryEntryUpdateModeOnlineAllUpdates       ExadbVmClusterUpdateHistoryEntryUpdateModeEnum = "ONLINE_ALL_UPDATES"
	ExadbVmClusterUpdateHistoryEntryUpdateModePendingUpdatesHighcvss ExadbVmClusterUpdateHistoryEntryUpdateModeEnum = "PENDING_UPDATES_HIGHCVSS"
	ExadbVmClusterUpdateHistoryEntryUpdateModePendingUpdatesAllcvss  ExadbVmClusterUpdateHistoryEntryUpdateModeEnum = "PENDING_UPDATES_ALLCVSS"
	ExadbVmClusterUpdateHistoryEntryUpdateModePendingAllUpdates      ExadbVmClusterUpdateHistoryEntryUpdateModeEnum = "PENDING_ALL_UPDATES"
	ExadbVmClusterUpdateHistoryEntryUpdateModeFullUpdate             ExadbVmClusterUpdateHistoryEntryUpdateModeEnum = "FULL_UPDATE"
)

var mappingExadbVmClusterUpdateHistoryEntryUpdateModeEnum = map[string]ExadbVmClusterUpdateHistoryEntryUpdateModeEnum{
	"ONLINE_HIGHCVSS":          ExadbVmClusterUpdateHistoryEntryUpdateModeOnlineHighcvss,
	"ONLINE_ALLCVSS":           ExadbVmClusterUpdateHistoryEntryUpdateModeOnlineAllcvss,
	"ONLINE_ALL_UPDATES":       ExadbVmClusterUpdateHistoryEntryUpdateModeOnlineAllUpdates,
	"PENDING_UPDATES_HIGHCVSS": ExadbVmClusterUpdateHistoryEntryUpdateModePendingUpdatesHighcvss,
	"PENDING_UPDATES_ALLCVSS":  ExadbVmClusterUpdateHistoryEntryUpdateModePendingUpdatesAllcvss,
	"PENDING_ALL_UPDATES":      ExadbVmClusterUpdateHistoryEntryUpdateModePendingAllUpdates,
	"FULL_UPDATE":              ExadbVmClusterUpdateHistoryEntryUpdateModeFullUpdate,
}

var mappingExadbVmClusterUpdateHistoryEntryUpdateModeEnumLowerCase = map[string]ExadbVmClusterUpdateHistoryEntryUpdateModeEnum{
	"online_highcvss":          ExadbVmClusterUpdateHistoryEntryUpdateModeOnlineHighcvss,
	"online_allcvss":           ExadbVmClusterUpdateHistoryEntryUpdateModeOnlineAllcvss,
	"online_all_updates":       ExadbVmClusterUpdateHistoryEntryUpdateModeOnlineAllUpdates,
	"pending_updates_highcvss": ExadbVmClusterUpdateHistoryEntryUpdateModePendingUpdatesHighcvss,
	"pending_updates_allcvss":  ExadbVmClusterUpdateHistoryEntryUpdateModePendingUpdatesAllcvss,
	"pending_all_updates":      ExadbVmClusterUpdateHistoryEntryUpdateModePendingAllUpdates,
	"full_update":              ExadbVmClusterUpdateHistoryEntryUpdateModeFullUpdate,
}

// GetExadbVmClusterUpdateHistoryEntryUpdateModeEnumValues Enumerates the set of values for ExadbVmClusterUpdateHistoryEntryUpdateModeEnum
func GetExadbVmClusterUpdateHistoryEntryUpdateModeEnumValues() []ExadbVmClusterUpdateHistoryEntryUpdateModeEnum {
	values := make([]ExadbVmClusterUpdateHistoryEntryUpdateModeEnum, 0)
	for _, v := range mappingExadbVmClusterUpdateHistoryEntryUpdateModeEnum {
		values = append(values, v)
	}
	return values
}

// GetExadbVmClusterUpdateHistoryEntryUpdateModeEnumStringValues Enumerates the set of values in String for ExadbVmClusterUpdateHistoryEntryUpdateModeEnum
func GetExadbVmClusterUpdateHistoryEntryUpdateModeEnumStringValues() []string {
	return []string{
		"ONLINE_HIGHCVSS",
		"ONLINE_ALLCVSS",
		"ONLINE_ALL_UPDATES",
		"PENDING_UPDATES_HIGHCVSS",
		"PENDING_UPDATES_ALLCVSS",
		"PENDING_ALL_UPDATES",
		"FULL_UPDATE",
	}
}

// GetMappingExadbVmClusterUpdateHistoryEntryUpdateModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExadbVmClusterUpdateHistoryEntryUpdateModeEnum(val string) (ExadbVmClusterUpdateHistoryEntryUpdateModeEnum, bool) {
	enum, ok := mappingExadbVmClusterUpdateHistoryEntryUpdateModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ExadbVmClusterUpdateHistoryEntryUpdateTypeEnum Enum with underlying type: string
type ExadbVmClusterUpdateHistoryEntryUpdateTypeEnum string

// Set of constants representing the allowable values for ExadbVmClusterUpdateHistoryEntryUpdateTypeEnum
const (
	ExadbVmClusterUpdateHistoryEntryUpdateTypeGiUpgrade ExadbVmClusterUpdateHistoryEntryUpdateTypeEnum = "GI_UPGRADE"
	ExadbVmClusterUpdateHistoryEntryUpdateTypeGiPatch   ExadbVmClusterUpdateHistoryEntryUpdateTypeEnum = "GI_PATCH"
	ExadbVmClusterUpdateHistoryEntryUpdateTypeOsUpdate  ExadbVmClusterUpdateHistoryEntryUpdateTypeEnum = "OS_UPDATE"
)

var mappingExadbVmClusterUpdateHistoryEntryUpdateTypeEnum = map[string]ExadbVmClusterUpdateHistoryEntryUpdateTypeEnum{
	"GI_UPGRADE": ExadbVmClusterUpdateHistoryEntryUpdateTypeGiUpgrade,
	"GI_PATCH":   ExadbVmClusterUpdateHistoryEntryUpdateTypeGiPatch,
	"OS_UPDATE":  ExadbVmClusterUpdateHistoryEntryUpdateTypeOsUpdate,
}

var mappingExadbVmClusterUpdateHistoryEntryUpdateTypeEnumLowerCase = map[string]ExadbVmClusterUpdateHistoryEntryUpdateTypeEnum{
	"gi_upgrade": ExadbVmClusterUpdateHistoryEntryUpdateTypeGiUpgrade,
	"gi_patch":   ExadbVmClusterUpdateHistoryEntryUpdateTypeGiPatch,
	"os_update":  ExadbVmClusterUpdateHistoryEntryUpdateTypeOsUpdate,
}

// GetExadbVmClusterUpdateHistoryEntryUpdateTypeEnumValues Enumerates the set of values for ExadbVmClusterUpdateHistoryEntryUpdateTypeEnum
func GetExadbVmClusterUpdateHistoryEntryUpdateTypeEnumValues() []ExadbVmClusterUpdateHistoryEntryUpdateTypeEnum {
	values := make([]ExadbVmClusterUpdateHistoryEntryUpdateTypeEnum, 0)
	for _, v := range mappingExadbVmClusterUpdateHistoryEntryUpdateTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetExadbVmClusterUpdateHistoryEntryUpdateTypeEnumStringValues Enumerates the set of values in String for ExadbVmClusterUpdateHistoryEntryUpdateTypeEnum
func GetExadbVmClusterUpdateHistoryEntryUpdateTypeEnumStringValues() []string {
	return []string{
		"GI_UPGRADE",
		"GI_PATCH",
		"OS_UPDATE",
	}
}

// GetMappingExadbVmClusterUpdateHistoryEntryUpdateTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExadbVmClusterUpdateHistoryEntryUpdateTypeEnum(val string) (ExadbVmClusterUpdateHistoryEntryUpdateTypeEnum, bool) {
	enum, ok := mappingExadbVmClusterUpdateHistoryEntryUpdateTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ExadbVmClusterUpdateHistoryEntryLifecycleStateEnum Enum with underlying type: string
type ExadbVmClusterUpdateHistoryEntryLifecycleStateEnum string

// Set of constants representing the allowable values for ExadbVmClusterUpdateHistoryEntryLifecycleStateEnum
const (
	ExadbVmClusterUpdateHistoryEntryLifecycleStateInProgress ExadbVmClusterUpdateHistoryEntryLifecycleStateEnum = "IN_PROGRESS"
	ExadbVmClusterUpdateHistoryEntryLifecycleStateSucceeded  ExadbVmClusterUpdateHistoryEntryLifecycleStateEnum = "SUCCEEDED"
	ExadbVmClusterUpdateHistoryEntryLifecycleStateFailed     ExadbVmClusterUpdateHistoryEntryLifecycleStateEnum = "FAILED"
)

var mappingExadbVmClusterUpdateHistoryEntryLifecycleStateEnum = map[string]ExadbVmClusterUpdateHistoryEntryLifecycleStateEnum{
	"IN_PROGRESS": ExadbVmClusterUpdateHistoryEntryLifecycleStateInProgress,
	"SUCCEEDED":   ExadbVmClusterUpdateHistoryEntryLifecycleStateSucceeded,
	"FAILED":      ExadbVmClusterUpdateHistoryEntryLifecycleStateFailed,
}

var mappingExadbVmClusterUpdateHistoryEntryLifecycleStateEnumLowerCase = map[string]ExadbVmClusterUpdateHistoryEntryLifecycleStateEnum{
	"in_progress": ExadbVmClusterUpdateHistoryEntryLifecycleStateInProgress,
	"succeeded":   ExadbVmClusterUpdateHistoryEntryLifecycleStateSucceeded,
	"failed":      ExadbVmClusterUpdateHistoryEntryLifecycleStateFailed,
}

// GetExadbVmClusterUpdateHistoryEntryLifecycleStateEnumValues Enumerates the set of values for ExadbVmClusterUpdateHistoryEntryLifecycleStateEnum
func GetExadbVmClusterUpdateHistoryEntryLifecycleStateEnumValues() []ExadbVmClusterUpdateHistoryEntryLifecycleStateEnum {
	values := make([]ExadbVmClusterUpdateHistoryEntryLifecycleStateEnum, 0)
	for _, v := range mappingExadbVmClusterUpdateHistoryEntryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetExadbVmClusterUpdateHistoryEntryLifecycleStateEnumStringValues Enumerates the set of values in String for ExadbVmClusterUpdateHistoryEntryLifecycleStateEnum
func GetExadbVmClusterUpdateHistoryEntryLifecycleStateEnumStringValues() []string {
	return []string{
		"IN_PROGRESS",
		"SUCCEEDED",
		"FAILED",
	}
}

// GetMappingExadbVmClusterUpdateHistoryEntryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExadbVmClusterUpdateHistoryEntryLifecycleStateEnum(val string) (ExadbVmClusterUpdateHistoryEntryLifecycleStateEnum, bool) {
	enum, ok := mappingExadbVmClusterUpdateHistoryEntryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
