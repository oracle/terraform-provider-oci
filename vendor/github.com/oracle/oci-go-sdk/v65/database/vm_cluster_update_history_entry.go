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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VmClusterUpdateHistoryEntry) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingVmClusterUpdateHistoryEntryUpdateTypeEnum(string(m.UpdateType)); !ok && m.UpdateType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpdateType: %s. Supported values are: %s.", m.UpdateType, strings.Join(GetVmClusterUpdateHistoryEntryUpdateTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingVmClusterUpdateHistoryEntryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetVmClusterUpdateHistoryEntryLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingVmClusterUpdateHistoryEntryUpdateActionEnum(string(m.UpdateAction)); !ok && m.UpdateAction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpdateAction: %s. Supported values are: %s.", m.UpdateAction, strings.Join(GetVmClusterUpdateHistoryEntryUpdateActionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// VmClusterUpdateHistoryEntryUpdateActionEnum Enum with underlying type: string
type VmClusterUpdateHistoryEntryUpdateActionEnum string

// Set of constants representing the allowable values for VmClusterUpdateHistoryEntryUpdateActionEnum
const (
	VmClusterUpdateHistoryEntryUpdateActionRollingApply VmClusterUpdateHistoryEntryUpdateActionEnum = "ROLLING_APPLY"
	VmClusterUpdateHistoryEntryUpdateActionPrecheck     VmClusterUpdateHistoryEntryUpdateActionEnum = "PRECHECK"
	VmClusterUpdateHistoryEntryUpdateActionRollback     VmClusterUpdateHistoryEntryUpdateActionEnum = "ROLLBACK"
)

var mappingVmClusterUpdateHistoryEntryUpdateActionEnum = map[string]VmClusterUpdateHistoryEntryUpdateActionEnum{
	"ROLLING_APPLY": VmClusterUpdateHistoryEntryUpdateActionRollingApply,
	"PRECHECK":      VmClusterUpdateHistoryEntryUpdateActionPrecheck,
	"ROLLBACK":      VmClusterUpdateHistoryEntryUpdateActionRollback,
}

var mappingVmClusterUpdateHistoryEntryUpdateActionEnumLowerCase = map[string]VmClusterUpdateHistoryEntryUpdateActionEnum{
	"rolling_apply": VmClusterUpdateHistoryEntryUpdateActionRollingApply,
	"precheck":      VmClusterUpdateHistoryEntryUpdateActionPrecheck,
	"rollback":      VmClusterUpdateHistoryEntryUpdateActionRollback,
}

// GetVmClusterUpdateHistoryEntryUpdateActionEnumValues Enumerates the set of values for VmClusterUpdateHistoryEntryUpdateActionEnum
func GetVmClusterUpdateHistoryEntryUpdateActionEnumValues() []VmClusterUpdateHistoryEntryUpdateActionEnum {
	values := make([]VmClusterUpdateHistoryEntryUpdateActionEnum, 0)
	for _, v := range mappingVmClusterUpdateHistoryEntryUpdateActionEnum {
		values = append(values, v)
	}
	return values
}

// GetVmClusterUpdateHistoryEntryUpdateActionEnumStringValues Enumerates the set of values in String for VmClusterUpdateHistoryEntryUpdateActionEnum
func GetVmClusterUpdateHistoryEntryUpdateActionEnumStringValues() []string {
	return []string{
		"ROLLING_APPLY",
		"PRECHECK",
		"ROLLBACK",
	}
}

// GetMappingVmClusterUpdateHistoryEntryUpdateActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVmClusterUpdateHistoryEntryUpdateActionEnum(val string) (VmClusterUpdateHistoryEntryUpdateActionEnum, bool) {
	enum, ok := mappingVmClusterUpdateHistoryEntryUpdateActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// VmClusterUpdateHistoryEntryUpdateTypeEnum Enum with underlying type: string
type VmClusterUpdateHistoryEntryUpdateTypeEnum string

// Set of constants representing the allowable values for VmClusterUpdateHistoryEntryUpdateTypeEnum
const (
	VmClusterUpdateHistoryEntryUpdateTypeGiUpgrade VmClusterUpdateHistoryEntryUpdateTypeEnum = "GI_UPGRADE"
	VmClusterUpdateHistoryEntryUpdateTypeGiPatch   VmClusterUpdateHistoryEntryUpdateTypeEnum = "GI_PATCH"
	VmClusterUpdateHistoryEntryUpdateTypeOsUpdate  VmClusterUpdateHistoryEntryUpdateTypeEnum = "OS_UPDATE"
)

var mappingVmClusterUpdateHistoryEntryUpdateTypeEnum = map[string]VmClusterUpdateHistoryEntryUpdateTypeEnum{
	"GI_UPGRADE": VmClusterUpdateHistoryEntryUpdateTypeGiUpgrade,
	"GI_PATCH":   VmClusterUpdateHistoryEntryUpdateTypeGiPatch,
	"OS_UPDATE":  VmClusterUpdateHistoryEntryUpdateTypeOsUpdate,
}

var mappingVmClusterUpdateHistoryEntryUpdateTypeEnumLowerCase = map[string]VmClusterUpdateHistoryEntryUpdateTypeEnum{
	"gi_upgrade": VmClusterUpdateHistoryEntryUpdateTypeGiUpgrade,
	"gi_patch":   VmClusterUpdateHistoryEntryUpdateTypeGiPatch,
	"os_update":  VmClusterUpdateHistoryEntryUpdateTypeOsUpdate,
}

// GetVmClusterUpdateHistoryEntryUpdateTypeEnumValues Enumerates the set of values for VmClusterUpdateHistoryEntryUpdateTypeEnum
func GetVmClusterUpdateHistoryEntryUpdateTypeEnumValues() []VmClusterUpdateHistoryEntryUpdateTypeEnum {
	values := make([]VmClusterUpdateHistoryEntryUpdateTypeEnum, 0)
	for _, v := range mappingVmClusterUpdateHistoryEntryUpdateTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetVmClusterUpdateHistoryEntryUpdateTypeEnumStringValues Enumerates the set of values in String for VmClusterUpdateHistoryEntryUpdateTypeEnum
func GetVmClusterUpdateHistoryEntryUpdateTypeEnumStringValues() []string {
	return []string{
		"GI_UPGRADE",
		"GI_PATCH",
		"OS_UPDATE",
	}
}

// GetMappingVmClusterUpdateHistoryEntryUpdateTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVmClusterUpdateHistoryEntryUpdateTypeEnum(val string) (VmClusterUpdateHistoryEntryUpdateTypeEnum, bool) {
	enum, ok := mappingVmClusterUpdateHistoryEntryUpdateTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// VmClusterUpdateHistoryEntryLifecycleStateEnum Enum with underlying type: string
type VmClusterUpdateHistoryEntryLifecycleStateEnum string

// Set of constants representing the allowable values for VmClusterUpdateHistoryEntryLifecycleStateEnum
const (
	VmClusterUpdateHistoryEntryLifecycleStateInProgress VmClusterUpdateHistoryEntryLifecycleStateEnum = "IN_PROGRESS"
	VmClusterUpdateHistoryEntryLifecycleStateSucceeded  VmClusterUpdateHistoryEntryLifecycleStateEnum = "SUCCEEDED"
	VmClusterUpdateHistoryEntryLifecycleStateFailed     VmClusterUpdateHistoryEntryLifecycleStateEnum = "FAILED"
)

var mappingVmClusterUpdateHistoryEntryLifecycleStateEnum = map[string]VmClusterUpdateHistoryEntryLifecycleStateEnum{
	"IN_PROGRESS": VmClusterUpdateHistoryEntryLifecycleStateInProgress,
	"SUCCEEDED":   VmClusterUpdateHistoryEntryLifecycleStateSucceeded,
	"FAILED":      VmClusterUpdateHistoryEntryLifecycleStateFailed,
}

var mappingVmClusterUpdateHistoryEntryLifecycleStateEnumLowerCase = map[string]VmClusterUpdateHistoryEntryLifecycleStateEnum{
	"in_progress": VmClusterUpdateHistoryEntryLifecycleStateInProgress,
	"succeeded":   VmClusterUpdateHistoryEntryLifecycleStateSucceeded,
	"failed":      VmClusterUpdateHistoryEntryLifecycleStateFailed,
}

// GetVmClusterUpdateHistoryEntryLifecycleStateEnumValues Enumerates the set of values for VmClusterUpdateHistoryEntryLifecycleStateEnum
func GetVmClusterUpdateHistoryEntryLifecycleStateEnumValues() []VmClusterUpdateHistoryEntryLifecycleStateEnum {
	values := make([]VmClusterUpdateHistoryEntryLifecycleStateEnum, 0)
	for _, v := range mappingVmClusterUpdateHistoryEntryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetVmClusterUpdateHistoryEntryLifecycleStateEnumStringValues Enumerates the set of values in String for VmClusterUpdateHistoryEntryLifecycleStateEnum
func GetVmClusterUpdateHistoryEntryLifecycleStateEnumStringValues() []string {
	return []string{
		"IN_PROGRESS",
		"SUCCEEDED",
		"FAILED",
	}
}

// GetMappingVmClusterUpdateHistoryEntryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVmClusterUpdateHistoryEntryLifecycleStateEnum(val string) (VmClusterUpdateHistoryEntryLifecycleStateEnum, bool) {
	enum, ok := mappingVmClusterUpdateHistoryEntryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
