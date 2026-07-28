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

// BaseccVmClusterUpdate Maintenance update details for a BaseDB-C@C VM cluster.
type BaseccVmClusterUpdate struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the maintenance update.
	Id *string `mandatory:"true" json:"id"`

	// Details of the maintenance update package.
	Description *string `mandatory:"true" json:"description"`

	// The type of Base Database Service on Cloud@Customer (BaseDB-C@C) VM cluster maintenance update.
	UpdateType BaseccVmClusterUpdateUpdateTypeEnum `mandatory:"true" json:"updateType"`

	// The date and time the maintenance update was released.
	TimeReleased *common.SDKTime `mandatory:"true" json:"timeReleased"`

	// The version of the maintenance update package.
	Version *string `mandatory:"true" json:"version"`

	// The previous update action performed.
	LastAction BaseccVmClusterUpdateLastActionEnum `mandatory:"false" json:"lastAction,omitempty"`

	// The possible actions performed by the update operation on the infrastructure components.
	AvailableActions []BaseccVmClusterUpdateAvailableActionsEnum `mandatory:"false" json:"availableActions,omitempty"`

	// Descriptive text providing additional details about the lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The current state of the maintenance update. Dependent on value of `lastAction`.
	LifecycleState BaseccVmClusterUpdateLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m BaseccVmClusterUpdate) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BaseccVmClusterUpdate) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBaseccVmClusterUpdateUpdateTypeEnum(string(m.UpdateType)); !ok && m.UpdateType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpdateType: %s. Supported values are: %s.", m.UpdateType, strings.Join(GetBaseccVmClusterUpdateUpdateTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingBaseccVmClusterUpdateLastActionEnum(string(m.LastAction)); !ok && m.LastAction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LastAction: %s. Supported values are: %s.", m.LastAction, strings.Join(GetBaseccVmClusterUpdateLastActionEnumStringValues(), ",")))
	}
	for _, val := range m.AvailableActions {
		if _, ok := GetMappingBaseccVmClusterUpdateAvailableActionsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AvailableActions: %s. Supported values are: %s.", val, strings.Join(GetBaseccVmClusterUpdateAvailableActionsEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingBaseccVmClusterUpdateLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetBaseccVmClusterUpdateLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BaseccVmClusterUpdateLastActionEnum Enum with underlying type: string
type BaseccVmClusterUpdateLastActionEnum string

// Set of constants representing the allowable values for BaseccVmClusterUpdateLastActionEnum
const (
	BaseccVmClusterUpdateLastActionRollingApply BaseccVmClusterUpdateLastActionEnum = "ROLLING_APPLY"
	BaseccVmClusterUpdateLastActionPrecheck     BaseccVmClusterUpdateLastActionEnum = "PRECHECK"
	BaseccVmClusterUpdateLastActionRollback     BaseccVmClusterUpdateLastActionEnum = "ROLLBACK"
)

var mappingBaseccVmClusterUpdateLastActionEnum = map[string]BaseccVmClusterUpdateLastActionEnum{
	"ROLLING_APPLY": BaseccVmClusterUpdateLastActionRollingApply,
	"PRECHECK":      BaseccVmClusterUpdateLastActionPrecheck,
	"ROLLBACK":      BaseccVmClusterUpdateLastActionRollback,
}

var mappingBaseccVmClusterUpdateLastActionEnumLowerCase = map[string]BaseccVmClusterUpdateLastActionEnum{
	"rolling_apply": BaseccVmClusterUpdateLastActionRollingApply,
	"precheck":      BaseccVmClusterUpdateLastActionPrecheck,
	"rollback":      BaseccVmClusterUpdateLastActionRollback,
}

// GetBaseccVmClusterUpdateLastActionEnumValues Enumerates the set of values for BaseccVmClusterUpdateLastActionEnum
func GetBaseccVmClusterUpdateLastActionEnumValues() []BaseccVmClusterUpdateLastActionEnum {
	values := make([]BaseccVmClusterUpdateLastActionEnum, 0)
	for _, v := range mappingBaseccVmClusterUpdateLastActionEnum {
		values = append(values, v)
	}
	return values
}

// GetBaseccVmClusterUpdateLastActionEnumStringValues Enumerates the set of values in String for BaseccVmClusterUpdateLastActionEnum
func GetBaseccVmClusterUpdateLastActionEnumStringValues() []string {
	return []string{
		"ROLLING_APPLY",
		"PRECHECK",
		"ROLLBACK",
	}
}

// GetMappingBaseccVmClusterUpdateLastActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBaseccVmClusterUpdateLastActionEnum(val string) (BaseccVmClusterUpdateLastActionEnum, bool) {
	enum, ok := mappingBaseccVmClusterUpdateLastActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// BaseccVmClusterUpdateAvailableActionsEnum Enum with underlying type: string
type BaseccVmClusterUpdateAvailableActionsEnum string

// Set of constants representing the allowable values for BaseccVmClusterUpdateAvailableActionsEnum
const (
	BaseccVmClusterUpdateAvailableActionsRollingApply BaseccVmClusterUpdateAvailableActionsEnum = "ROLLING_APPLY"
	BaseccVmClusterUpdateAvailableActionsPrecheck     BaseccVmClusterUpdateAvailableActionsEnum = "PRECHECK"
	BaseccVmClusterUpdateAvailableActionsRollback     BaseccVmClusterUpdateAvailableActionsEnum = "ROLLBACK"
)

var mappingBaseccVmClusterUpdateAvailableActionsEnum = map[string]BaseccVmClusterUpdateAvailableActionsEnum{
	"ROLLING_APPLY": BaseccVmClusterUpdateAvailableActionsRollingApply,
	"PRECHECK":      BaseccVmClusterUpdateAvailableActionsPrecheck,
	"ROLLBACK":      BaseccVmClusterUpdateAvailableActionsRollback,
}

var mappingBaseccVmClusterUpdateAvailableActionsEnumLowerCase = map[string]BaseccVmClusterUpdateAvailableActionsEnum{
	"rolling_apply": BaseccVmClusterUpdateAvailableActionsRollingApply,
	"precheck":      BaseccVmClusterUpdateAvailableActionsPrecheck,
	"rollback":      BaseccVmClusterUpdateAvailableActionsRollback,
}

// GetBaseccVmClusterUpdateAvailableActionsEnumValues Enumerates the set of values for BaseccVmClusterUpdateAvailableActionsEnum
func GetBaseccVmClusterUpdateAvailableActionsEnumValues() []BaseccVmClusterUpdateAvailableActionsEnum {
	values := make([]BaseccVmClusterUpdateAvailableActionsEnum, 0)
	for _, v := range mappingBaseccVmClusterUpdateAvailableActionsEnum {
		values = append(values, v)
	}
	return values
}

// GetBaseccVmClusterUpdateAvailableActionsEnumStringValues Enumerates the set of values in String for BaseccVmClusterUpdateAvailableActionsEnum
func GetBaseccVmClusterUpdateAvailableActionsEnumStringValues() []string {
	return []string{
		"ROLLING_APPLY",
		"PRECHECK",
		"ROLLBACK",
	}
}

// GetMappingBaseccVmClusterUpdateAvailableActionsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBaseccVmClusterUpdateAvailableActionsEnum(val string) (BaseccVmClusterUpdateAvailableActionsEnum, bool) {
	enum, ok := mappingBaseccVmClusterUpdateAvailableActionsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// BaseccVmClusterUpdateUpdateTypeEnum Enum with underlying type: string
type BaseccVmClusterUpdateUpdateTypeEnum string

// Set of constants representing the allowable values for BaseccVmClusterUpdateUpdateTypeEnum
const (
	BaseccVmClusterUpdateUpdateTypeUpgrade BaseccVmClusterUpdateUpdateTypeEnum = "GI_UPGRADE"
	BaseccVmClusterUpdateUpdateTypePatch   BaseccVmClusterUpdateUpdateTypeEnum = "GI_PATCH"
)

var mappingBaseccVmClusterUpdateUpdateTypeEnum = map[string]BaseccVmClusterUpdateUpdateTypeEnum{
	"GI_UPGRADE": BaseccVmClusterUpdateUpdateTypeUpgrade,
	"GI_PATCH":   BaseccVmClusterUpdateUpdateTypePatch,
}

var mappingBaseccVmClusterUpdateUpdateTypeEnumLowerCase = map[string]BaseccVmClusterUpdateUpdateTypeEnum{
	"gi_upgrade": BaseccVmClusterUpdateUpdateTypeUpgrade,
	"gi_patch":   BaseccVmClusterUpdateUpdateTypePatch,
}

// GetBaseccVmClusterUpdateUpdateTypeEnumValues Enumerates the set of values for BaseccVmClusterUpdateUpdateTypeEnum
func GetBaseccVmClusterUpdateUpdateTypeEnumValues() []BaseccVmClusterUpdateUpdateTypeEnum {
	values := make([]BaseccVmClusterUpdateUpdateTypeEnum, 0)
	for _, v := range mappingBaseccVmClusterUpdateUpdateTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetBaseccVmClusterUpdateUpdateTypeEnumStringValues Enumerates the set of values in String for BaseccVmClusterUpdateUpdateTypeEnum
func GetBaseccVmClusterUpdateUpdateTypeEnumStringValues() []string {
	return []string{
		"GI_UPGRADE",
		"GI_PATCH",
	}
}

// GetMappingBaseccVmClusterUpdateUpdateTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBaseccVmClusterUpdateUpdateTypeEnum(val string) (BaseccVmClusterUpdateUpdateTypeEnum, bool) {
	enum, ok := mappingBaseccVmClusterUpdateUpdateTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// BaseccVmClusterUpdateLifecycleStateEnum Enum with underlying type: string
type BaseccVmClusterUpdateLifecycleStateEnum string

// Set of constants representing the allowable values for BaseccVmClusterUpdateLifecycleStateEnum
const (
	BaseccVmClusterUpdateLifecycleStateAvailable  BaseccVmClusterUpdateLifecycleStateEnum = "AVAILABLE"
	BaseccVmClusterUpdateLifecycleStateSuccess    BaseccVmClusterUpdateLifecycleStateEnum = "SUCCESS"
	BaseccVmClusterUpdateLifecycleStateInProgress BaseccVmClusterUpdateLifecycleStateEnum = "IN_PROGRESS"
	BaseccVmClusterUpdateLifecycleStateFailed     BaseccVmClusterUpdateLifecycleStateEnum = "FAILED"
)

var mappingBaseccVmClusterUpdateLifecycleStateEnum = map[string]BaseccVmClusterUpdateLifecycleStateEnum{
	"AVAILABLE":   BaseccVmClusterUpdateLifecycleStateAvailable,
	"SUCCESS":     BaseccVmClusterUpdateLifecycleStateSuccess,
	"IN_PROGRESS": BaseccVmClusterUpdateLifecycleStateInProgress,
	"FAILED":      BaseccVmClusterUpdateLifecycleStateFailed,
}

var mappingBaseccVmClusterUpdateLifecycleStateEnumLowerCase = map[string]BaseccVmClusterUpdateLifecycleStateEnum{
	"available":   BaseccVmClusterUpdateLifecycleStateAvailable,
	"success":     BaseccVmClusterUpdateLifecycleStateSuccess,
	"in_progress": BaseccVmClusterUpdateLifecycleStateInProgress,
	"failed":      BaseccVmClusterUpdateLifecycleStateFailed,
}

// GetBaseccVmClusterUpdateLifecycleStateEnumValues Enumerates the set of values for BaseccVmClusterUpdateLifecycleStateEnum
func GetBaseccVmClusterUpdateLifecycleStateEnumValues() []BaseccVmClusterUpdateLifecycleStateEnum {
	values := make([]BaseccVmClusterUpdateLifecycleStateEnum, 0)
	for _, v := range mappingBaseccVmClusterUpdateLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetBaseccVmClusterUpdateLifecycleStateEnumStringValues Enumerates the set of values in String for BaseccVmClusterUpdateLifecycleStateEnum
func GetBaseccVmClusterUpdateLifecycleStateEnumStringValues() []string {
	return []string{
		"AVAILABLE",
		"SUCCESS",
		"IN_PROGRESS",
		"FAILED",
	}
}

// GetMappingBaseccVmClusterUpdateLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBaseccVmClusterUpdateLifecycleStateEnum(val string) (BaseccVmClusterUpdateLifecycleStateEnum, bool) {
	enum, ok := mappingBaseccVmClusterUpdateLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
