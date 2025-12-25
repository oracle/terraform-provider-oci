// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// BaseccVmClusterUpdateHistoryEntry The record of a maintenance update action on a specified BaseDB-C@C VM cluster.
type BaseccVmClusterUpdateHistoryEntry struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the maintenance update history entry.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the maintenance update.
	UpdateId *string `mandatory:"true" json:"updateId"`

	// The type of Base Database Service on Cloud@Customer (BaseDB-C@C) VM cluster maintenance update.
	UpdateType BaseccVmClusterUpdateHistoryEntryUpdateTypeEnum `mandatory:"true" json:"updateType"`

	// The current lifecycle state of the maintenance update operation.
	LifecycleState BaseccVmClusterUpdateHistoryEntryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time when the maintenance update action started.
	TimeStarted *common.SDKTime `mandatory:"true" json:"timeStarted"`

	// The update action.
	UpdateAction BaseccVmClusterUpdateHistoryEntryUpdateActionEnum `mandatory:"false" json:"updateAction,omitempty"`

	// Descriptive text providing additional details about the lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The date and time when the maintenance update action completed.
	TimeCompleted *common.SDKTime `mandatory:"false" json:"timeCompleted"`

	// The version of the maintenance update package.
	Version *string `mandatory:"false" json:"version"`
}

func (m BaseccVmClusterUpdateHistoryEntry) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BaseccVmClusterUpdateHistoryEntry) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBaseccVmClusterUpdateHistoryEntryUpdateTypeEnum(string(m.UpdateType)); !ok && m.UpdateType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpdateType: %s. Supported values are: %s.", m.UpdateType, strings.Join(GetBaseccVmClusterUpdateHistoryEntryUpdateTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingBaseccVmClusterUpdateHistoryEntryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetBaseccVmClusterUpdateHistoryEntryLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingBaseccVmClusterUpdateHistoryEntryUpdateActionEnum(string(m.UpdateAction)); !ok && m.UpdateAction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpdateAction: %s. Supported values are: %s.", m.UpdateAction, strings.Join(GetBaseccVmClusterUpdateHistoryEntryUpdateActionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BaseccVmClusterUpdateHistoryEntryUpdateActionEnum Enum with underlying type: string
type BaseccVmClusterUpdateHistoryEntryUpdateActionEnum string

// Set of constants representing the allowable values for BaseccVmClusterUpdateHistoryEntryUpdateActionEnum
const (
	BaseccVmClusterUpdateHistoryEntryUpdateActionRollingApply BaseccVmClusterUpdateHistoryEntryUpdateActionEnum = "ROLLING_APPLY"
	BaseccVmClusterUpdateHistoryEntryUpdateActionPrecheck     BaseccVmClusterUpdateHistoryEntryUpdateActionEnum = "PRECHECK"
	BaseccVmClusterUpdateHistoryEntryUpdateActionRollback     BaseccVmClusterUpdateHistoryEntryUpdateActionEnum = "ROLLBACK"
)

var mappingBaseccVmClusterUpdateHistoryEntryUpdateActionEnum = map[string]BaseccVmClusterUpdateHistoryEntryUpdateActionEnum{
	"ROLLING_APPLY": BaseccVmClusterUpdateHistoryEntryUpdateActionRollingApply,
	"PRECHECK":      BaseccVmClusterUpdateHistoryEntryUpdateActionPrecheck,
	"ROLLBACK":      BaseccVmClusterUpdateHistoryEntryUpdateActionRollback,
}

var mappingBaseccVmClusterUpdateHistoryEntryUpdateActionEnumLowerCase = map[string]BaseccVmClusterUpdateHistoryEntryUpdateActionEnum{
	"rolling_apply": BaseccVmClusterUpdateHistoryEntryUpdateActionRollingApply,
	"precheck":      BaseccVmClusterUpdateHistoryEntryUpdateActionPrecheck,
	"rollback":      BaseccVmClusterUpdateHistoryEntryUpdateActionRollback,
}

// GetBaseccVmClusterUpdateHistoryEntryUpdateActionEnumValues Enumerates the set of values for BaseccVmClusterUpdateHistoryEntryUpdateActionEnum
func GetBaseccVmClusterUpdateHistoryEntryUpdateActionEnumValues() []BaseccVmClusterUpdateHistoryEntryUpdateActionEnum {
	values := make([]BaseccVmClusterUpdateHistoryEntryUpdateActionEnum, 0)
	for _, v := range mappingBaseccVmClusterUpdateHistoryEntryUpdateActionEnum {
		values = append(values, v)
	}
	return values
}

// GetBaseccVmClusterUpdateHistoryEntryUpdateActionEnumStringValues Enumerates the set of values in String for BaseccVmClusterUpdateHistoryEntryUpdateActionEnum
func GetBaseccVmClusterUpdateHistoryEntryUpdateActionEnumStringValues() []string {
	return []string{
		"ROLLING_APPLY",
		"PRECHECK",
		"ROLLBACK",
	}
}

// GetMappingBaseccVmClusterUpdateHistoryEntryUpdateActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBaseccVmClusterUpdateHistoryEntryUpdateActionEnum(val string) (BaseccVmClusterUpdateHistoryEntryUpdateActionEnum, bool) {
	enum, ok := mappingBaseccVmClusterUpdateHistoryEntryUpdateActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// BaseccVmClusterUpdateHistoryEntryUpdateTypeEnum Enum with underlying type: string
type BaseccVmClusterUpdateHistoryEntryUpdateTypeEnum string

// Set of constants representing the allowable values for BaseccVmClusterUpdateHistoryEntryUpdateTypeEnum
const (
	BaseccVmClusterUpdateHistoryEntryUpdateTypeUpgrade BaseccVmClusterUpdateHistoryEntryUpdateTypeEnum = "GI_UPGRADE"
	BaseccVmClusterUpdateHistoryEntryUpdateTypePatch   BaseccVmClusterUpdateHistoryEntryUpdateTypeEnum = "GI_PATCH"
)

var mappingBaseccVmClusterUpdateHistoryEntryUpdateTypeEnum = map[string]BaseccVmClusterUpdateHistoryEntryUpdateTypeEnum{
	"GI_UPGRADE": BaseccVmClusterUpdateHistoryEntryUpdateTypeUpgrade,
	"GI_PATCH":   BaseccVmClusterUpdateHistoryEntryUpdateTypePatch,
}

var mappingBaseccVmClusterUpdateHistoryEntryUpdateTypeEnumLowerCase = map[string]BaseccVmClusterUpdateHistoryEntryUpdateTypeEnum{
	"gi_upgrade": BaseccVmClusterUpdateHistoryEntryUpdateTypeUpgrade,
	"gi_patch":   BaseccVmClusterUpdateHistoryEntryUpdateTypePatch,
}

// GetBaseccVmClusterUpdateHistoryEntryUpdateTypeEnumValues Enumerates the set of values for BaseccVmClusterUpdateHistoryEntryUpdateTypeEnum
func GetBaseccVmClusterUpdateHistoryEntryUpdateTypeEnumValues() []BaseccVmClusterUpdateHistoryEntryUpdateTypeEnum {
	values := make([]BaseccVmClusterUpdateHistoryEntryUpdateTypeEnum, 0)
	for _, v := range mappingBaseccVmClusterUpdateHistoryEntryUpdateTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetBaseccVmClusterUpdateHistoryEntryUpdateTypeEnumStringValues Enumerates the set of values in String for BaseccVmClusterUpdateHistoryEntryUpdateTypeEnum
func GetBaseccVmClusterUpdateHistoryEntryUpdateTypeEnumStringValues() []string {
	return []string{
		"GI_UPGRADE",
		"GI_PATCH",
	}
}

// GetMappingBaseccVmClusterUpdateHistoryEntryUpdateTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBaseccVmClusterUpdateHistoryEntryUpdateTypeEnum(val string) (BaseccVmClusterUpdateHistoryEntryUpdateTypeEnum, bool) {
	enum, ok := mappingBaseccVmClusterUpdateHistoryEntryUpdateTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// BaseccVmClusterUpdateHistoryEntryLifecycleStateEnum Enum with underlying type: string
type BaseccVmClusterUpdateHistoryEntryLifecycleStateEnum string

// Set of constants representing the allowable values for BaseccVmClusterUpdateHistoryEntryLifecycleStateEnum
const (
	BaseccVmClusterUpdateHistoryEntryLifecycleStateInProgress BaseccVmClusterUpdateHistoryEntryLifecycleStateEnum = "IN_PROGRESS"
	BaseccVmClusterUpdateHistoryEntryLifecycleStateSucceeded  BaseccVmClusterUpdateHistoryEntryLifecycleStateEnum = "SUCCEEDED"
	BaseccVmClusterUpdateHistoryEntryLifecycleStateFailed     BaseccVmClusterUpdateHistoryEntryLifecycleStateEnum = "FAILED"
)

var mappingBaseccVmClusterUpdateHistoryEntryLifecycleStateEnum = map[string]BaseccVmClusterUpdateHistoryEntryLifecycleStateEnum{
	"IN_PROGRESS": BaseccVmClusterUpdateHistoryEntryLifecycleStateInProgress,
	"SUCCEEDED":   BaseccVmClusterUpdateHistoryEntryLifecycleStateSucceeded,
	"FAILED":      BaseccVmClusterUpdateHistoryEntryLifecycleStateFailed,
}

var mappingBaseccVmClusterUpdateHistoryEntryLifecycleStateEnumLowerCase = map[string]BaseccVmClusterUpdateHistoryEntryLifecycleStateEnum{
	"in_progress": BaseccVmClusterUpdateHistoryEntryLifecycleStateInProgress,
	"succeeded":   BaseccVmClusterUpdateHistoryEntryLifecycleStateSucceeded,
	"failed":      BaseccVmClusterUpdateHistoryEntryLifecycleStateFailed,
}

// GetBaseccVmClusterUpdateHistoryEntryLifecycleStateEnumValues Enumerates the set of values for BaseccVmClusterUpdateHistoryEntryLifecycleStateEnum
func GetBaseccVmClusterUpdateHistoryEntryLifecycleStateEnumValues() []BaseccVmClusterUpdateHistoryEntryLifecycleStateEnum {
	values := make([]BaseccVmClusterUpdateHistoryEntryLifecycleStateEnum, 0)
	for _, v := range mappingBaseccVmClusterUpdateHistoryEntryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetBaseccVmClusterUpdateHistoryEntryLifecycleStateEnumStringValues Enumerates the set of values in String for BaseccVmClusterUpdateHistoryEntryLifecycleStateEnum
func GetBaseccVmClusterUpdateHistoryEntryLifecycleStateEnumStringValues() []string {
	return []string{
		"IN_PROGRESS",
		"SUCCEEDED",
		"FAILED",
	}
}

// GetMappingBaseccVmClusterUpdateHistoryEntryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBaseccVmClusterUpdateHistoryEntryLifecycleStateEnum(val string) (BaseccVmClusterUpdateHistoryEntryLifecycleStateEnum, bool) {
	enum, ok := mappingBaseccVmClusterUpdateHistoryEntryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
