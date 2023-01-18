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

// VmClusterUpdate A maintenance update for a VM cluster. Applies to Exadata Cloud@Customer instances only.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access,
// see Getting Started with Policies (https://docs.cloud.oracle.com/Content/Identity/Concepts/policygetstarted.htm).
type VmClusterUpdate struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the maintenance update.
	Id *string `mandatory:"true" json:"id"`

	// Details of the maintenance update package.
	Description *string `mandatory:"true" json:"description"`

	// The type of VM cluster maintenance update.
	UpdateType VmClusterUpdateUpdateTypeEnum `mandatory:"true" json:"updateType"`

	// The date and time the maintenance update was released.
	TimeReleased *common.SDKTime `mandatory:"true" json:"timeReleased"`

	// The version of the maintenance update package.
	Version *string `mandatory:"true" json:"version"`

	// The update action performed most recently using this maintenance update.
	LastAction VmClusterUpdateLastActionEnum `mandatory:"false" json:"lastAction,omitempty"`

	// The possible actions that can be performed using this maintenance update.
	AvailableActions []VmClusterUpdateAvailableActionsEnum `mandatory:"false" json:"availableActions,omitempty"`

	// Descriptive text providing additional details about the lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The current state of the maintenance update. Dependent on value of `lastAction`.
	LifecycleState VmClusterUpdateLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m VmClusterUpdate) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VmClusterUpdate) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingVmClusterUpdateUpdateTypeEnum(string(m.UpdateType)); !ok && m.UpdateType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpdateType: %s. Supported values are: %s.", m.UpdateType, strings.Join(GetVmClusterUpdateUpdateTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingVmClusterUpdateLastActionEnum(string(m.LastAction)); !ok && m.LastAction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LastAction: %s. Supported values are: %s.", m.LastAction, strings.Join(GetVmClusterUpdateLastActionEnumStringValues(), ",")))
	}
	for _, val := range m.AvailableActions {
		if _, ok := GetMappingVmClusterUpdateAvailableActionsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AvailableActions: %s. Supported values are: %s.", val, strings.Join(GetVmClusterUpdateAvailableActionsEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingVmClusterUpdateLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetVmClusterUpdateLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// VmClusterUpdateLastActionEnum Enum with underlying type: string
type VmClusterUpdateLastActionEnum string

// Set of constants representing the allowable values for VmClusterUpdateLastActionEnum
const (
	VmClusterUpdateLastActionRollingApply VmClusterUpdateLastActionEnum = "ROLLING_APPLY"
	VmClusterUpdateLastActionPrecheck     VmClusterUpdateLastActionEnum = "PRECHECK"
	VmClusterUpdateLastActionRollback     VmClusterUpdateLastActionEnum = "ROLLBACK"
)

var mappingVmClusterUpdateLastActionEnum = map[string]VmClusterUpdateLastActionEnum{
	"ROLLING_APPLY": VmClusterUpdateLastActionRollingApply,
	"PRECHECK":      VmClusterUpdateLastActionPrecheck,
	"ROLLBACK":      VmClusterUpdateLastActionRollback,
}

var mappingVmClusterUpdateLastActionEnumLowerCase = map[string]VmClusterUpdateLastActionEnum{
	"rolling_apply": VmClusterUpdateLastActionRollingApply,
	"precheck":      VmClusterUpdateLastActionPrecheck,
	"rollback":      VmClusterUpdateLastActionRollback,
}

// GetVmClusterUpdateLastActionEnumValues Enumerates the set of values for VmClusterUpdateLastActionEnum
func GetVmClusterUpdateLastActionEnumValues() []VmClusterUpdateLastActionEnum {
	values := make([]VmClusterUpdateLastActionEnum, 0)
	for _, v := range mappingVmClusterUpdateLastActionEnum {
		values = append(values, v)
	}
	return values
}

// GetVmClusterUpdateLastActionEnumStringValues Enumerates the set of values in String for VmClusterUpdateLastActionEnum
func GetVmClusterUpdateLastActionEnumStringValues() []string {
	return []string{
		"ROLLING_APPLY",
		"PRECHECK",
		"ROLLBACK",
	}
}

// GetMappingVmClusterUpdateLastActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVmClusterUpdateLastActionEnum(val string) (VmClusterUpdateLastActionEnum, bool) {
	enum, ok := mappingVmClusterUpdateLastActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// VmClusterUpdateAvailableActionsEnum Enum with underlying type: string
type VmClusterUpdateAvailableActionsEnum string

// Set of constants representing the allowable values for VmClusterUpdateAvailableActionsEnum
const (
	VmClusterUpdateAvailableActionsRollingApply VmClusterUpdateAvailableActionsEnum = "ROLLING_APPLY"
	VmClusterUpdateAvailableActionsPrecheck     VmClusterUpdateAvailableActionsEnum = "PRECHECK"
	VmClusterUpdateAvailableActionsRollback     VmClusterUpdateAvailableActionsEnum = "ROLLBACK"
)

var mappingVmClusterUpdateAvailableActionsEnum = map[string]VmClusterUpdateAvailableActionsEnum{
	"ROLLING_APPLY": VmClusterUpdateAvailableActionsRollingApply,
	"PRECHECK":      VmClusterUpdateAvailableActionsPrecheck,
	"ROLLBACK":      VmClusterUpdateAvailableActionsRollback,
}

var mappingVmClusterUpdateAvailableActionsEnumLowerCase = map[string]VmClusterUpdateAvailableActionsEnum{
	"rolling_apply": VmClusterUpdateAvailableActionsRollingApply,
	"precheck":      VmClusterUpdateAvailableActionsPrecheck,
	"rollback":      VmClusterUpdateAvailableActionsRollback,
}

// GetVmClusterUpdateAvailableActionsEnumValues Enumerates the set of values for VmClusterUpdateAvailableActionsEnum
func GetVmClusterUpdateAvailableActionsEnumValues() []VmClusterUpdateAvailableActionsEnum {
	values := make([]VmClusterUpdateAvailableActionsEnum, 0)
	for _, v := range mappingVmClusterUpdateAvailableActionsEnum {
		values = append(values, v)
	}
	return values
}

// GetVmClusterUpdateAvailableActionsEnumStringValues Enumerates the set of values in String for VmClusterUpdateAvailableActionsEnum
func GetVmClusterUpdateAvailableActionsEnumStringValues() []string {
	return []string{
		"ROLLING_APPLY",
		"PRECHECK",
		"ROLLBACK",
	}
}

// GetMappingVmClusterUpdateAvailableActionsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVmClusterUpdateAvailableActionsEnum(val string) (VmClusterUpdateAvailableActionsEnum, bool) {
	enum, ok := mappingVmClusterUpdateAvailableActionsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// VmClusterUpdateUpdateTypeEnum Enum with underlying type: string
type VmClusterUpdateUpdateTypeEnum string

// Set of constants representing the allowable values for VmClusterUpdateUpdateTypeEnum
const (
	VmClusterUpdateUpdateTypeGiUpgrade VmClusterUpdateUpdateTypeEnum = "GI_UPGRADE"
	VmClusterUpdateUpdateTypeGiPatch   VmClusterUpdateUpdateTypeEnum = "GI_PATCH"
	VmClusterUpdateUpdateTypeOsUpdate  VmClusterUpdateUpdateTypeEnum = "OS_UPDATE"
)

var mappingVmClusterUpdateUpdateTypeEnum = map[string]VmClusterUpdateUpdateTypeEnum{
	"GI_UPGRADE": VmClusterUpdateUpdateTypeGiUpgrade,
	"GI_PATCH":   VmClusterUpdateUpdateTypeGiPatch,
	"OS_UPDATE":  VmClusterUpdateUpdateTypeOsUpdate,
}

var mappingVmClusterUpdateUpdateTypeEnumLowerCase = map[string]VmClusterUpdateUpdateTypeEnum{
	"gi_upgrade": VmClusterUpdateUpdateTypeGiUpgrade,
	"gi_patch":   VmClusterUpdateUpdateTypeGiPatch,
	"os_update":  VmClusterUpdateUpdateTypeOsUpdate,
}

// GetVmClusterUpdateUpdateTypeEnumValues Enumerates the set of values for VmClusterUpdateUpdateTypeEnum
func GetVmClusterUpdateUpdateTypeEnumValues() []VmClusterUpdateUpdateTypeEnum {
	values := make([]VmClusterUpdateUpdateTypeEnum, 0)
	for _, v := range mappingVmClusterUpdateUpdateTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetVmClusterUpdateUpdateTypeEnumStringValues Enumerates the set of values in String for VmClusterUpdateUpdateTypeEnum
func GetVmClusterUpdateUpdateTypeEnumStringValues() []string {
	return []string{
		"GI_UPGRADE",
		"GI_PATCH",
		"OS_UPDATE",
	}
}

// GetMappingVmClusterUpdateUpdateTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVmClusterUpdateUpdateTypeEnum(val string) (VmClusterUpdateUpdateTypeEnum, bool) {
	enum, ok := mappingVmClusterUpdateUpdateTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// VmClusterUpdateLifecycleStateEnum Enum with underlying type: string
type VmClusterUpdateLifecycleStateEnum string

// Set of constants representing the allowable values for VmClusterUpdateLifecycleStateEnum
const (
	VmClusterUpdateLifecycleStateAvailable  VmClusterUpdateLifecycleStateEnum = "AVAILABLE"
	VmClusterUpdateLifecycleStateSuccess    VmClusterUpdateLifecycleStateEnum = "SUCCESS"
	VmClusterUpdateLifecycleStateInProgress VmClusterUpdateLifecycleStateEnum = "IN_PROGRESS"
	VmClusterUpdateLifecycleStateFailed     VmClusterUpdateLifecycleStateEnum = "FAILED"
)

var mappingVmClusterUpdateLifecycleStateEnum = map[string]VmClusterUpdateLifecycleStateEnum{
	"AVAILABLE":   VmClusterUpdateLifecycleStateAvailable,
	"SUCCESS":     VmClusterUpdateLifecycleStateSuccess,
	"IN_PROGRESS": VmClusterUpdateLifecycleStateInProgress,
	"FAILED":      VmClusterUpdateLifecycleStateFailed,
}

var mappingVmClusterUpdateLifecycleStateEnumLowerCase = map[string]VmClusterUpdateLifecycleStateEnum{
	"available":   VmClusterUpdateLifecycleStateAvailable,
	"success":     VmClusterUpdateLifecycleStateSuccess,
	"in_progress": VmClusterUpdateLifecycleStateInProgress,
	"failed":      VmClusterUpdateLifecycleStateFailed,
}

// GetVmClusterUpdateLifecycleStateEnumValues Enumerates the set of values for VmClusterUpdateLifecycleStateEnum
func GetVmClusterUpdateLifecycleStateEnumValues() []VmClusterUpdateLifecycleStateEnum {
	values := make([]VmClusterUpdateLifecycleStateEnum, 0)
	for _, v := range mappingVmClusterUpdateLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetVmClusterUpdateLifecycleStateEnumStringValues Enumerates the set of values in String for VmClusterUpdateLifecycleStateEnum
func GetVmClusterUpdateLifecycleStateEnumStringValues() []string {
	return []string{
		"AVAILABLE",
		"SUCCESS",
		"IN_PROGRESS",
		"FAILED",
	}
}

// GetMappingVmClusterUpdateLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVmClusterUpdateLifecycleStateEnum(val string) (VmClusterUpdateLifecycleStateEnum, bool) {
	enum, ok := mappingVmClusterUpdateLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
