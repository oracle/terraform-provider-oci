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

// VmClusterUpdateSummary A maintenance update for a VM cluster. Applies to Exadata Cloud@Customer instances only.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access,
// see Getting Started with Policies (https://docs.oracle.com/iaas/Content/Identity/Concepts/policygetstarted.htm).
type VmClusterUpdateSummary struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the maintenance update.
	Id *string `mandatory:"true" json:"id"`

	// Details of the maintenance update package.
	Description *string `mandatory:"true" json:"description"`

	// The type of VM cluster maintenance update.
	UpdateType VmClusterUpdateSummaryUpdateTypeEnum `mandatory:"true" json:"updateType"`

	// The date and time the maintenance update was released.
	TimeReleased *common.SDKTime `mandatory:"true" json:"timeReleased"`

	// The version of the maintenance update package.
	Version *string `mandatory:"true" json:"version"`

	// The update action performed most recently using this maintenance update.
	LastAction VmClusterUpdateSummaryLastActionEnum `mandatory:"false" json:"lastAction,omitempty"`

	// The update mode performed most recently using this maintenance update (only valid for OS Update).
	LastUpdateMode VmClusterUpdateSummaryLastUpdateModeEnum `mandatory:"false" json:"lastUpdateMode,omitempty"`

	// The possible update options that can be performed using this maintenance update (only valid for OS Update).
	AvailableUpdateModes []VmClusterUpdateSummaryAvailableUpdateModesEnum `mandatory:"false" json:"availableUpdateModes,omitempty"`

	// Oracle Linux version for the respective Exadata Image.
	OracleLinuxVersion *string `mandatory:"false" json:"oracleLinuxVersion"`

	// The possible actions that can be performed using this maintenance update.
	AvailableActions []VmClusterUpdateSummaryAvailableActionsEnum `mandatory:"false" json:"availableActions,omitempty"`

	// Descriptive text providing additional details about the lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The current state of the maintenance update. Dependent on value of `lastAction`.
	LifecycleState VmClusterUpdateSummaryLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m VmClusterUpdateSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VmClusterUpdateSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingVmClusterUpdateSummaryUpdateTypeEnum(string(m.UpdateType)); !ok && m.UpdateType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpdateType: %s. Supported values are: %s.", m.UpdateType, strings.Join(GetVmClusterUpdateSummaryUpdateTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingVmClusterUpdateSummaryLastActionEnum(string(m.LastAction)); !ok && m.LastAction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LastAction: %s. Supported values are: %s.", m.LastAction, strings.Join(GetVmClusterUpdateSummaryLastActionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingVmClusterUpdateSummaryLastUpdateModeEnum(string(m.LastUpdateMode)); !ok && m.LastUpdateMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LastUpdateMode: %s. Supported values are: %s.", m.LastUpdateMode, strings.Join(GetVmClusterUpdateSummaryLastUpdateModeEnumStringValues(), ",")))
	}
	for _, val := range m.AvailableUpdateModes {
		if _, ok := GetMappingVmClusterUpdateSummaryAvailableUpdateModesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AvailableUpdateModes: %s. Supported values are: %s.", val, strings.Join(GetVmClusterUpdateSummaryAvailableUpdateModesEnumStringValues(), ",")))
		}
	}

	for _, val := range m.AvailableActions {
		if _, ok := GetMappingVmClusterUpdateSummaryAvailableActionsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AvailableActions: %s. Supported values are: %s.", val, strings.Join(GetVmClusterUpdateSummaryAvailableActionsEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingVmClusterUpdateSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetVmClusterUpdateSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// VmClusterUpdateSummaryLastActionEnum Enum with underlying type: string
type VmClusterUpdateSummaryLastActionEnum string

// Set of constants representing the allowable values for VmClusterUpdateSummaryLastActionEnum
const (
	VmClusterUpdateSummaryLastActionRollingApply VmClusterUpdateSummaryLastActionEnum = "ROLLING_APPLY"
	VmClusterUpdateSummaryLastActionPrecheck     VmClusterUpdateSummaryLastActionEnum = "PRECHECK"
	VmClusterUpdateSummaryLastActionRollback     VmClusterUpdateSummaryLastActionEnum = "ROLLBACK"
)

var mappingVmClusterUpdateSummaryLastActionEnum = map[string]VmClusterUpdateSummaryLastActionEnum{
	"ROLLING_APPLY": VmClusterUpdateSummaryLastActionRollingApply,
	"PRECHECK":      VmClusterUpdateSummaryLastActionPrecheck,
	"ROLLBACK":      VmClusterUpdateSummaryLastActionRollback,
}

var mappingVmClusterUpdateSummaryLastActionEnumLowerCase = map[string]VmClusterUpdateSummaryLastActionEnum{
	"rolling_apply": VmClusterUpdateSummaryLastActionRollingApply,
	"precheck":      VmClusterUpdateSummaryLastActionPrecheck,
	"rollback":      VmClusterUpdateSummaryLastActionRollback,
}

// GetVmClusterUpdateSummaryLastActionEnumValues Enumerates the set of values for VmClusterUpdateSummaryLastActionEnum
func GetVmClusterUpdateSummaryLastActionEnumValues() []VmClusterUpdateSummaryLastActionEnum {
	values := make([]VmClusterUpdateSummaryLastActionEnum, 0)
	for _, v := range mappingVmClusterUpdateSummaryLastActionEnum {
		values = append(values, v)
	}
	return values
}

// GetVmClusterUpdateSummaryLastActionEnumStringValues Enumerates the set of values in String for VmClusterUpdateSummaryLastActionEnum
func GetVmClusterUpdateSummaryLastActionEnumStringValues() []string {
	return []string{
		"ROLLING_APPLY",
		"PRECHECK",
		"ROLLBACK",
	}
}

// GetMappingVmClusterUpdateSummaryLastActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVmClusterUpdateSummaryLastActionEnum(val string) (VmClusterUpdateSummaryLastActionEnum, bool) {
	enum, ok := mappingVmClusterUpdateSummaryLastActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// VmClusterUpdateSummaryLastUpdateModeEnum Enum with underlying type: string
type VmClusterUpdateSummaryLastUpdateModeEnum string

// Set of constants representing the allowable values for VmClusterUpdateSummaryLastUpdateModeEnum
const (
	VmClusterUpdateSummaryLastUpdateModeOnlineHighcvss   VmClusterUpdateSummaryLastUpdateModeEnum = "ONLINE_HIGHCVSS"
	VmClusterUpdateSummaryLastUpdateModeOnlineAllcvss    VmClusterUpdateSummaryLastUpdateModeEnum = "ONLINE_ALLCVSS"
	VmClusterUpdateSummaryLastUpdateModeOnlineAllUpdates VmClusterUpdateSummaryLastUpdateModeEnum = "ONLINE_ALL_UPDATES"
	VmClusterUpdateSummaryLastUpdateModePendingUpdates   VmClusterUpdateSummaryLastUpdateModeEnum = "PENDING_UPDATES"
	VmClusterUpdateSummaryLastUpdateModeFullUpdate       VmClusterUpdateSummaryLastUpdateModeEnum = "FULL_UPDATE"
)

var mappingVmClusterUpdateSummaryLastUpdateModeEnum = map[string]VmClusterUpdateSummaryLastUpdateModeEnum{
	"ONLINE_HIGHCVSS":    VmClusterUpdateSummaryLastUpdateModeOnlineHighcvss,
	"ONLINE_ALLCVSS":     VmClusterUpdateSummaryLastUpdateModeOnlineAllcvss,
	"ONLINE_ALL_UPDATES": VmClusterUpdateSummaryLastUpdateModeOnlineAllUpdates,
	"PENDING_UPDATES":    VmClusterUpdateSummaryLastUpdateModePendingUpdates,
	"FULL_UPDATE":        VmClusterUpdateSummaryLastUpdateModeFullUpdate,
}

var mappingVmClusterUpdateSummaryLastUpdateModeEnumLowerCase = map[string]VmClusterUpdateSummaryLastUpdateModeEnum{
	"online_highcvss":    VmClusterUpdateSummaryLastUpdateModeOnlineHighcvss,
	"online_allcvss":     VmClusterUpdateSummaryLastUpdateModeOnlineAllcvss,
	"online_all_updates": VmClusterUpdateSummaryLastUpdateModeOnlineAllUpdates,
	"pending_updates":    VmClusterUpdateSummaryLastUpdateModePendingUpdates,
	"full_update":        VmClusterUpdateSummaryLastUpdateModeFullUpdate,
}

// GetVmClusterUpdateSummaryLastUpdateModeEnumValues Enumerates the set of values for VmClusterUpdateSummaryLastUpdateModeEnum
func GetVmClusterUpdateSummaryLastUpdateModeEnumValues() []VmClusterUpdateSummaryLastUpdateModeEnum {
	values := make([]VmClusterUpdateSummaryLastUpdateModeEnum, 0)
	for _, v := range mappingVmClusterUpdateSummaryLastUpdateModeEnum {
		values = append(values, v)
	}
	return values
}

// GetVmClusterUpdateSummaryLastUpdateModeEnumStringValues Enumerates the set of values in String for VmClusterUpdateSummaryLastUpdateModeEnum
func GetVmClusterUpdateSummaryLastUpdateModeEnumStringValues() []string {
	return []string{
		"ONLINE_HIGHCVSS",
		"ONLINE_ALLCVSS",
		"ONLINE_ALL_UPDATES",
		"PENDING_UPDATES",
		"FULL_UPDATE",
	}
}

// GetMappingVmClusterUpdateSummaryLastUpdateModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVmClusterUpdateSummaryLastUpdateModeEnum(val string) (VmClusterUpdateSummaryLastUpdateModeEnum, bool) {
	enum, ok := mappingVmClusterUpdateSummaryLastUpdateModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// VmClusterUpdateSummaryAvailableUpdateModesEnum Enum with underlying type: string
type VmClusterUpdateSummaryAvailableUpdateModesEnum string

// Set of constants representing the allowable values for VmClusterUpdateSummaryAvailableUpdateModesEnum
const (
	VmClusterUpdateSummaryAvailableUpdateModesOnlineHighcvss   VmClusterUpdateSummaryAvailableUpdateModesEnum = "ONLINE_HIGHCVSS"
	VmClusterUpdateSummaryAvailableUpdateModesOnlineAllcvss    VmClusterUpdateSummaryAvailableUpdateModesEnum = "ONLINE_ALLCVSS"
	VmClusterUpdateSummaryAvailableUpdateModesOnlineAllUpdates VmClusterUpdateSummaryAvailableUpdateModesEnum = "ONLINE_ALL_UPDATES"
	VmClusterUpdateSummaryAvailableUpdateModesPendingUpdates   VmClusterUpdateSummaryAvailableUpdateModesEnum = "PENDING_UPDATES"
	VmClusterUpdateSummaryAvailableUpdateModesFullUpdate       VmClusterUpdateSummaryAvailableUpdateModesEnum = "FULL_UPDATE"
)

var mappingVmClusterUpdateSummaryAvailableUpdateModesEnum = map[string]VmClusterUpdateSummaryAvailableUpdateModesEnum{
	"ONLINE_HIGHCVSS":    VmClusterUpdateSummaryAvailableUpdateModesOnlineHighcvss,
	"ONLINE_ALLCVSS":     VmClusterUpdateSummaryAvailableUpdateModesOnlineAllcvss,
	"ONLINE_ALL_UPDATES": VmClusterUpdateSummaryAvailableUpdateModesOnlineAllUpdates,
	"PENDING_UPDATES":    VmClusterUpdateSummaryAvailableUpdateModesPendingUpdates,
	"FULL_UPDATE":        VmClusterUpdateSummaryAvailableUpdateModesFullUpdate,
}

var mappingVmClusterUpdateSummaryAvailableUpdateModesEnumLowerCase = map[string]VmClusterUpdateSummaryAvailableUpdateModesEnum{
	"online_highcvss":    VmClusterUpdateSummaryAvailableUpdateModesOnlineHighcvss,
	"online_allcvss":     VmClusterUpdateSummaryAvailableUpdateModesOnlineAllcvss,
	"online_all_updates": VmClusterUpdateSummaryAvailableUpdateModesOnlineAllUpdates,
	"pending_updates":    VmClusterUpdateSummaryAvailableUpdateModesPendingUpdates,
	"full_update":        VmClusterUpdateSummaryAvailableUpdateModesFullUpdate,
}

// GetVmClusterUpdateSummaryAvailableUpdateModesEnumValues Enumerates the set of values for VmClusterUpdateSummaryAvailableUpdateModesEnum
func GetVmClusterUpdateSummaryAvailableUpdateModesEnumValues() []VmClusterUpdateSummaryAvailableUpdateModesEnum {
	values := make([]VmClusterUpdateSummaryAvailableUpdateModesEnum, 0)
	for _, v := range mappingVmClusterUpdateSummaryAvailableUpdateModesEnum {
		values = append(values, v)
	}
	return values
}

// GetVmClusterUpdateSummaryAvailableUpdateModesEnumStringValues Enumerates the set of values in String for VmClusterUpdateSummaryAvailableUpdateModesEnum
func GetVmClusterUpdateSummaryAvailableUpdateModesEnumStringValues() []string {
	return []string{
		"ONLINE_HIGHCVSS",
		"ONLINE_ALLCVSS",
		"ONLINE_ALL_UPDATES",
		"PENDING_UPDATES",
		"FULL_UPDATE",
	}
}

// GetMappingVmClusterUpdateSummaryAvailableUpdateModesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVmClusterUpdateSummaryAvailableUpdateModesEnum(val string) (VmClusterUpdateSummaryAvailableUpdateModesEnum, bool) {
	enum, ok := mappingVmClusterUpdateSummaryAvailableUpdateModesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// VmClusterUpdateSummaryAvailableActionsEnum Enum with underlying type: string
type VmClusterUpdateSummaryAvailableActionsEnum string

// Set of constants representing the allowable values for VmClusterUpdateSummaryAvailableActionsEnum
const (
	VmClusterUpdateSummaryAvailableActionsRollingApply VmClusterUpdateSummaryAvailableActionsEnum = "ROLLING_APPLY"
	VmClusterUpdateSummaryAvailableActionsPrecheck     VmClusterUpdateSummaryAvailableActionsEnum = "PRECHECK"
	VmClusterUpdateSummaryAvailableActionsRollback     VmClusterUpdateSummaryAvailableActionsEnum = "ROLLBACK"
)

var mappingVmClusterUpdateSummaryAvailableActionsEnum = map[string]VmClusterUpdateSummaryAvailableActionsEnum{
	"ROLLING_APPLY": VmClusterUpdateSummaryAvailableActionsRollingApply,
	"PRECHECK":      VmClusterUpdateSummaryAvailableActionsPrecheck,
	"ROLLBACK":      VmClusterUpdateSummaryAvailableActionsRollback,
}

var mappingVmClusterUpdateSummaryAvailableActionsEnumLowerCase = map[string]VmClusterUpdateSummaryAvailableActionsEnum{
	"rolling_apply": VmClusterUpdateSummaryAvailableActionsRollingApply,
	"precheck":      VmClusterUpdateSummaryAvailableActionsPrecheck,
	"rollback":      VmClusterUpdateSummaryAvailableActionsRollback,
}

// GetVmClusterUpdateSummaryAvailableActionsEnumValues Enumerates the set of values for VmClusterUpdateSummaryAvailableActionsEnum
func GetVmClusterUpdateSummaryAvailableActionsEnumValues() []VmClusterUpdateSummaryAvailableActionsEnum {
	values := make([]VmClusterUpdateSummaryAvailableActionsEnum, 0)
	for _, v := range mappingVmClusterUpdateSummaryAvailableActionsEnum {
		values = append(values, v)
	}
	return values
}

// GetVmClusterUpdateSummaryAvailableActionsEnumStringValues Enumerates the set of values in String for VmClusterUpdateSummaryAvailableActionsEnum
func GetVmClusterUpdateSummaryAvailableActionsEnumStringValues() []string {
	return []string{
		"ROLLING_APPLY",
		"PRECHECK",
		"ROLLBACK",
	}
}

// GetMappingVmClusterUpdateSummaryAvailableActionsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVmClusterUpdateSummaryAvailableActionsEnum(val string) (VmClusterUpdateSummaryAvailableActionsEnum, bool) {
	enum, ok := mappingVmClusterUpdateSummaryAvailableActionsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// VmClusterUpdateSummaryUpdateTypeEnum Enum with underlying type: string
type VmClusterUpdateSummaryUpdateTypeEnum string

// Set of constants representing the allowable values for VmClusterUpdateSummaryUpdateTypeEnum
const (
	VmClusterUpdateSummaryUpdateTypeGiUpgrade VmClusterUpdateSummaryUpdateTypeEnum = "GI_UPGRADE"
	VmClusterUpdateSummaryUpdateTypeGiPatch   VmClusterUpdateSummaryUpdateTypeEnum = "GI_PATCH"
	VmClusterUpdateSummaryUpdateTypeOsUpdate  VmClusterUpdateSummaryUpdateTypeEnum = "OS_UPDATE"
)

var mappingVmClusterUpdateSummaryUpdateTypeEnum = map[string]VmClusterUpdateSummaryUpdateTypeEnum{
	"GI_UPGRADE": VmClusterUpdateSummaryUpdateTypeGiUpgrade,
	"GI_PATCH":   VmClusterUpdateSummaryUpdateTypeGiPatch,
	"OS_UPDATE":  VmClusterUpdateSummaryUpdateTypeOsUpdate,
}

var mappingVmClusterUpdateSummaryUpdateTypeEnumLowerCase = map[string]VmClusterUpdateSummaryUpdateTypeEnum{
	"gi_upgrade": VmClusterUpdateSummaryUpdateTypeGiUpgrade,
	"gi_patch":   VmClusterUpdateSummaryUpdateTypeGiPatch,
	"os_update":  VmClusterUpdateSummaryUpdateTypeOsUpdate,
}

// GetVmClusterUpdateSummaryUpdateTypeEnumValues Enumerates the set of values for VmClusterUpdateSummaryUpdateTypeEnum
func GetVmClusterUpdateSummaryUpdateTypeEnumValues() []VmClusterUpdateSummaryUpdateTypeEnum {
	values := make([]VmClusterUpdateSummaryUpdateTypeEnum, 0)
	for _, v := range mappingVmClusterUpdateSummaryUpdateTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetVmClusterUpdateSummaryUpdateTypeEnumStringValues Enumerates the set of values in String for VmClusterUpdateSummaryUpdateTypeEnum
func GetVmClusterUpdateSummaryUpdateTypeEnumStringValues() []string {
	return []string{
		"GI_UPGRADE",
		"GI_PATCH",
		"OS_UPDATE",
	}
}

// GetMappingVmClusterUpdateSummaryUpdateTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVmClusterUpdateSummaryUpdateTypeEnum(val string) (VmClusterUpdateSummaryUpdateTypeEnum, bool) {
	enum, ok := mappingVmClusterUpdateSummaryUpdateTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// VmClusterUpdateSummaryLifecycleStateEnum Enum with underlying type: string
type VmClusterUpdateSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for VmClusterUpdateSummaryLifecycleStateEnum
const (
	VmClusterUpdateSummaryLifecycleStateAvailable  VmClusterUpdateSummaryLifecycleStateEnum = "AVAILABLE"
	VmClusterUpdateSummaryLifecycleStateSuccess    VmClusterUpdateSummaryLifecycleStateEnum = "SUCCESS"
	VmClusterUpdateSummaryLifecycleStateInProgress VmClusterUpdateSummaryLifecycleStateEnum = "IN_PROGRESS"
	VmClusterUpdateSummaryLifecycleStateFailed     VmClusterUpdateSummaryLifecycleStateEnum = "FAILED"
)

var mappingVmClusterUpdateSummaryLifecycleStateEnum = map[string]VmClusterUpdateSummaryLifecycleStateEnum{
	"AVAILABLE":   VmClusterUpdateSummaryLifecycleStateAvailable,
	"SUCCESS":     VmClusterUpdateSummaryLifecycleStateSuccess,
	"IN_PROGRESS": VmClusterUpdateSummaryLifecycleStateInProgress,
	"FAILED":      VmClusterUpdateSummaryLifecycleStateFailed,
}

var mappingVmClusterUpdateSummaryLifecycleStateEnumLowerCase = map[string]VmClusterUpdateSummaryLifecycleStateEnum{
	"available":   VmClusterUpdateSummaryLifecycleStateAvailable,
	"success":     VmClusterUpdateSummaryLifecycleStateSuccess,
	"in_progress": VmClusterUpdateSummaryLifecycleStateInProgress,
	"failed":      VmClusterUpdateSummaryLifecycleStateFailed,
}

// GetVmClusterUpdateSummaryLifecycleStateEnumValues Enumerates the set of values for VmClusterUpdateSummaryLifecycleStateEnum
func GetVmClusterUpdateSummaryLifecycleStateEnumValues() []VmClusterUpdateSummaryLifecycleStateEnum {
	values := make([]VmClusterUpdateSummaryLifecycleStateEnum, 0)
	for _, v := range mappingVmClusterUpdateSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetVmClusterUpdateSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for VmClusterUpdateSummaryLifecycleStateEnum
func GetVmClusterUpdateSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"AVAILABLE",
		"SUCCESS",
		"IN_PROGRESS",
		"FAILED",
	}
}

// GetMappingVmClusterUpdateSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVmClusterUpdateSummaryLifecycleStateEnum(val string) (VmClusterUpdateSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingVmClusterUpdateSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
