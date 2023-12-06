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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VmClusterUpdateHistoryEntrySummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingVmClusterUpdateHistoryEntrySummaryUpdateTypeEnum(string(m.UpdateType)); !ok && m.UpdateType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpdateType: %s. Supported values are: %s.", m.UpdateType, strings.Join(GetVmClusterUpdateHistoryEntrySummaryUpdateTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingVmClusterUpdateHistoryEntrySummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetVmClusterUpdateHistoryEntrySummaryLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingVmClusterUpdateHistoryEntrySummaryUpdateActionEnum(string(m.UpdateAction)); !ok && m.UpdateAction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpdateAction: %s. Supported values are: %s.", m.UpdateAction, strings.Join(GetVmClusterUpdateHistoryEntrySummaryUpdateActionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// VmClusterUpdateHistoryEntrySummaryUpdateActionEnum Enum with underlying type: string
type VmClusterUpdateHistoryEntrySummaryUpdateActionEnum string

// Set of constants representing the allowable values for VmClusterUpdateHistoryEntrySummaryUpdateActionEnum
const (
	VmClusterUpdateHistoryEntrySummaryUpdateActionRollingApply VmClusterUpdateHistoryEntrySummaryUpdateActionEnum = "ROLLING_APPLY"
	VmClusterUpdateHistoryEntrySummaryUpdateActionPrecheck     VmClusterUpdateHistoryEntrySummaryUpdateActionEnum = "PRECHECK"
	VmClusterUpdateHistoryEntrySummaryUpdateActionRollback     VmClusterUpdateHistoryEntrySummaryUpdateActionEnum = "ROLLBACK"
)

var mappingVmClusterUpdateHistoryEntrySummaryUpdateActionEnum = map[string]VmClusterUpdateHistoryEntrySummaryUpdateActionEnum{
	"ROLLING_APPLY": VmClusterUpdateHistoryEntrySummaryUpdateActionRollingApply,
	"PRECHECK":      VmClusterUpdateHistoryEntrySummaryUpdateActionPrecheck,
	"ROLLBACK":      VmClusterUpdateHistoryEntrySummaryUpdateActionRollback,
}

var mappingVmClusterUpdateHistoryEntrySummaryUpdateActionEnumLowerCase = map[string]VmClusterUpdateHistoryEntrySummaryUpdateActionEnum{
	"rolling_apply": VmClusterUpdateHistoryEntrySummaryUpdateActionRollingApply,
	"precheck":      VmClusterUpdateHistoryEntrySummaryUpdateActionPrecheck,
	"rollback":      VmClusterUpdateHistoryEntrySummaryUpdateActionRollback,
}

// GetVmClusterUpdateHistoryEntrySummaryUpdateActionEnumValues Enumerates the set of values for VmClusterUpdateHistoryEntrySummaryUpdateActionEnum
func GetVmClusterUpdateHistoryEntrySummaryUpdateActionEnumValues() []VmClusterUpdateHistoryEntrySummaryUpdateActionEnum {
	values := make([]VmClusterUpdateHistoryEntrySummaryUpdateActionEnum, 0)
	for _, v := range mappingVmClusterUpdateHistoryEntrySummaryUpdateActionEnum {
		values = append(values, v)
	}
	return values
}

// GetVmClusterUpdateHistoryEntrySummaryUpdateActionEnumStringValues Enumerates the set of values in String for VmClusterUpdateHistoryEntrySummaryUpdateActionEnum
func GetVmClusterUpdateHistoryEntrySummaryUpdateActionEnumStringValues() []string {
	return []string{
		"ROLLING_APPLY",
		"PRECHECK",
		"ROLLBACK",
	}
}

// GetMappingVmClusterUpdateHistoryEntrySummaryUpdateActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVmClusterUpdateHistoryEntrySummaryUpdateActionEnum(val string) (VmClusterUpdateHistoryEntrySummaryUpdateActionEnum, bool) {
	enum, ok := mappingVmClusterUpdateHistoryEntrySummaryUpdateActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// VmClusterUpdateHistoryEntrySummaryUpdateTypeEnum Enum with underlying type: string
type VmClusterUpdateHistoryEntrySummaryUpdateTypeEnum string

// Set of constants representing the allowable values for VmClusterUpdateHistoryEntrySummaryUpdateTypeEnum
const (
	VmClusterUpdateHistoryEntrySummaryUpdateTypeGiUpgrade VmClusterUpdateHistoryEntrySummaryUpdateTypeEnum = "GI_UPGRADE"
	VmClusterUpdateHistoryEntrySummaryUpdateTypeGiPatch   VmClusterUpdateHistoryEntrySummaryUpdateTypeEnum = "GI_PATCH"
	VmClusterUpdateHistoryEntrySummaryUpdateTypeOsUpdate  VmClusterUpdateHistoryEntrySummaryUpdateTypeEnum = "OS_UPDATE"
)

var mappingVmClusterUpdateHistoryEntrySummaryUpdateTypeEnum = map[string]VmClusterUpdateHistoryEntrySummaryUpdateTypeEnum{
	"GI_UPGRADE": VmClusterUpdateHistoryEntrySummaryUpdateTypeGiUpgrade,
	"GI_PATCH":   VmClusterUpdateHistoryEntrySummaryUpdateTypeGiPatch,
	"OS_UPDATE":  VmClusterUpdateHistoryEntrySummaryUpdateTypeOsUpdate,
}

var mappingVmClusterUpdateHistoryEntrySummaryUpdateTypeEnumLowerCase = map[string]VmClusterUpdateHistoryEntrySummaryUpdateTypeEnum{
	"gi_upgrade": VmClusterUpdateHistoryEntrySummaryUpdateTypeGiUpgrade,
	"gi_patch":   VmClusterUpdateHistoryEntrySummaryUpdateTypeGiPatch,
	"os_update":  VmClusterUpdateHistoryEntrySummaryUpdateTypeOsUpdate,
}

// GetVmClusterUpdateHistoryEntrySummaryUpdateTypeEnumValues Enumerates the set of values for VmClusterUpdateHistoryEntrySummaryUpdateTypeEnum
func GetVmClusterUpdateHistoryEntrySummaryUpdateTypeEnumValues() []VmClusterUpdateHistoryEntrySummaryUpdateTypeEnum {
	values := make([]VmClusterUpdateHistoryEntrySummaryUpdateTypeEnum, 0)
	for _, v := range mappingVmClusterUpdateHistoryEntrySummaryUpdateTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetVmClusterUpdateHistoryEntrySummaryUpdateTypeEnumStringValues Enumerates the set of values in String for VmClusterUpdateHistoryEntrySummaryUpdateTypeEnum
func GetVmClusterUpdateHistoryEntrySummaryUpdateTypeEnumStringValues() []string {
	return []string{
		"GI_UPGRADE",
		"GI_PATCH",
		"OS_UPDATE",
	}
}

// GetMappingVmClusterUpdateHistoryEntrySummaryUpdateTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVmClusterUpdateHistoryEntrySummaryUpdateTypeEnum(val string) (VmClusterUpdateHistoryEntrySummaryUpdateTypeEnum, bool) {
	enum, ok := mappingVmClusterUpdateHistoryEntrySummaryUpdateTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// VmClusterUpdateHistoryEntrySummaryLifecycleStateEnum Enum with underlying type: string
type VmClusterUpdateHistoryEntrySummaryLifecycleStateEnum string

// Set of constants representing the allowable values for VmClusterUpdateHistoryEntrySummaryLifecycleStateEnum
const (
	VmClusterUpdateHistoryEntrySummaryLifecycleStateInProgress VmClusterUpdateHistoryEntrySummaryLifecycleStateEnum = "IN_PROGRESS"
	VmClusterUpdateHistoryEntrySummaryLifecycleStateSucceeded  VmClusterUpdateHistoryEntrySummaryLifecycleStateEnum = "SUCCEEDED"
	VmClusterUpdateHistoryEntrySummaryLifecycleStateFailed     VmClusterUpdateHistoryEntrySummaryLifecycleStateEnum = "FAILED"
)

var mappingVmClusterUpdateHistoryEntrySummaryLifecycleStateEnum = map[string]VmClusterUpdateHistoryEntrySummaryLifecycleStateEnum{
	"IN_PROGRESS": VmClusterUpdateHistoryEntrySummaryLifecycleStateInProgress,
	"SUCCEEDED":   VmClusterUpdateHistoryEntrySummaryLifecycleStateSucceeded,
	"FAILED":      VmClusterUpdateHistoryEntrySummaryLifecycleStateFailed,
}

var mappingVmClusterUpdateHistoryEntrySummaryLifecycleStateEnumLowerCase = map[string]VmClusterUpdateHistoryEntrySummaryLifecycleStateEnum{
	"in_progress": VmClusterUpdateHistoryEntrySummaryLifecycleStateInProgress,
	"succeeded":   VmClusterUpdateHistoryEntrySummaryLifecycleStateSucceeded,
	"failed":      VmClusterUpdateHistoryEntrySummaryLifecycleStateFailed,
}

// GetVmClusterUpdateHistoryEntrySummaryLifecycleStateEnumValues Enumerates the set of values for VmClusterUpdateHistoryEntrySummaryLifecycleStateEnum
func GetVmClusterUpdateHistoryEntrySummaryLifecycleStateEnumValues() []VmClusterUpdateHistoryEntrySummaryLifecycleStateEnum {
	values := make([]VmClusterUpdateHistoryEntrySummaryLifecycleStateEnum, 0)
	for _, v := range mappingVmClusterUpdateHistoryEntrySummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetVmClusterUpdateHistoryEntrySummaryLifecycleStateEnumStringValues Enumerates the set of values in String for VmClusterUpdateHistoryEntrySummaryLifecycleStateEnum
func GetVmClusterUpdateHistoryEntrySummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"IN_PROGRESS",
		"SUCCEEDED",
		"FAILED",
	}
}

// GetMappingVmClusterUpdateHistoryEntrySummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVmClusterUpdateHistoryEntrySummaryLifecycleStateEnum(val string) (VmClusterUpdateHistoryEntrySummaryLifecycleStateEnum, bool) {
	enum, ok := mappingVmClusterUpdateHistoryEntrySummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
