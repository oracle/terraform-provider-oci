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

// ExadbVmClusterUpdate A maintenance update details for an Exadata VM cluster on Exascale Infrastructure.
type ExadbVmClusterUpdate struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the maintenance update.
	Id *string `mandatory:"true" json:"id"`

	// Details of the maintenance update package.
	Description *string `mandatory:"true" json:"description"`

	// The type of cloud VM cluster maintenance update.
	UpdateType ExadbVmClusterUpdateUpdateTypeEnum `mandatory:"true" json:"updateType"`

	// The date and time the maintenance update was released.
	TimeReleased *common.SDKTime `mandatory:"true" json:"timeReleased"`

	// The version of the maintenance update package.
	Version *string `mandatory:"true" json:"version"`

	// The previous update action performed.
	LastAction ExadbVmClusterUpdateLastActionEnum `mandatory:"false" json:"lastAction,omitempty"`

	// The possible actions performed by the update operation on the infrastructure components.
	AvailableActions []ExadbVmClusterUpdateAvailableActionsEnum `mandatory:"false" json:"availableActions,omitempty"`

	// Descriptive text providing additional details about the lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The current state of the maintenance update. Dependent on value of `lastAction`.
	LifecycleState ExadbVmClusterUpdateLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m ExadbVmClusterUpdate) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExadbVmClusterUpdate) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExadbVmClusterUpdateUpdateTypeEnum(string(m.UpdateType)); !ok && m.UpdateType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpdateType: %s. Supported values are: %s.", m.UpdateType, strings.Join(GetExadbVmClusterUpdateUpdateTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingExadbVmClusterUpdateLastActionEnum(string(m.LastAction)); !ok && m.LastAction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LastAction: %s. Supported values are: %s.", m.LastAction, strings.Join(GetExadbVmClusterUpdateLastActionEnumStringValues(), ",")))
	}
	for _, val := range m.AvailableActions {
		if _, ok := GetMappingExadbVmClusterUpdateAvailableActionsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AvailableActions: %s. Supported values are: %s.", val, strings.Join(GetExadbVmClusterUpdateAvailableActionsEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingExadbVmClusterUpdateLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetExadbVmClusterUpdateLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExadbVmClusterUpdateLastActionEnum Enum with underlying type: string
type ExadbVmClusterUpdateLastActionEnum string

// Set of constants representing the allowable values for ExadbVmClusterUpdateLastActionEnum
const (
	ExadbVmClusterUpdateLastActionRollingApply    ExadbVmClusterUpdateLastActionEnum = "ROLLING_APPLY"
	ExadbVmClusterUpdateLastActionNonRollingApply ExadbVmClusterUpdateLastActionEnum = "NON_ROLLING_APPLY"
	ExadbVmClusterUpdateLastActionPrecheck        ExadbVmClusterUpdateLastActionEnum = "PRECHECK"
	ExadbVmClusterUpdateLastActionRollback        ExadbVmClusterUpdateLastActionEnum = "ROLLBACK"
)

var mappingExadbVmClusterUpdateLastActionEnum = map[string]ExadbVmClusterUpdateLastActionEnum{
	"ROLLING_APPLY":     ExadbVmClusterUpdateLastActionRollingApply,
	"NON_ROLLING_APPLY": ExadbVmClusterUpdateLastActionNonRollingApply,
	"PRECHECK":          ExadbVmClusterUpdateLastActionPrecheck,
	"ROLLBACK":          ExadbVmClusterUpdateLastActionRollback,
}

var mappingExadbVmClusterUpdateLastActionEnumLowerCase = map[string]ExadbVmClusterUpdateLastActionEnum{
	"rolling_apply":     ExadbVmClusterUpdateLastActionRollingApply,
	"non_rolling_apply": ExadbVmClusterUpdateLastActionNonRollingApply,
	"precheck":          ExadbVmClusterUpdateLastActionPrecheck,
	"rollback":          ExadbVmClusterUpdateLastActionRollback,
}

// GetExadbVmClusterUpdateLastActionEnumValues Enumerates the set of values for ExadbVmClusterUpdateLastActionEnum
func GetExadbVmClusterUpdateLastActionEnumValues() []ExadbVmClusterUpdateLastActionEnum {
	values := make([]ExadbVmClusterUpdateLastActionEnum, 0)
	for _, v := range mappingExadbVmClusterUpdateLastActionEnum {
		values = append(values, v)
	}
	return values
}

// GetExadbVmClusterUpdateLastActionEnumStringValues Enumerates the set of values in String for ExadbVmClusterUpdateLastActionEnum
func GetExadbVmClusterUpdateLastActionEnumStringValues() []string {
	return []string{
		"ROLLING_APPLY",
		"NON_ROLLING_APPLY",
		"PRECHECK",
		"ROLLBACK",
	}
}

// GetMappingExadbVmClusterUpdateLastActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExadbVmClusterUpdateLastActionEnum(val string) (ExadbVmClusterUpdateLastActionEnum, bool) {
	enum, ok := mappingExadbVmClusterUpdateLastActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ExadbVmClusterUpdateAvailableActionsEnum Enum with underlying type: string
type ExadbVmClusterUpdateAvailableActionsEnum string

// Set of constants representing the allowable values for ExadbVmClusterUpdateAvailableActionsEnum
const (
	ExadbVmClusterUpdateAvailableActionsRollingApply    ExadbVmClusterUpdateAvailableActionsEnum = "ROLLING_APPLY"
	ExadbVmClusterUpdateAvailableActionsNonRollingApply ExadbVmClusterUpdateAvailableActionsEnum = "NON_ROLLING_APPLY"
	ExadbVmClusterUpdateAvailableActionsPrecheck        ExadbVmClusterUpdateAvailableActionsEnum = "PRECHECK"
	ExadbVmClusterUpdateAvailableActionsRollback        ExadbVmClusterUpdateAvailableActionsEnum = "ROLLBACK"
)

var mappingExadbVmClusterUpdateAvailableActionsEnum = map[string]ExadbVmClusterUpdateAvailableActionsEnum{
	"ROLLING_APPLY":     ExadbVmClusterUpdateAvailableActionsRollingApply,
	"NON_ROLLING_APPLY": ExadbVmClusterUpdateAvailableActionsNonRollingApply,
	"PRECHECK":          ExadbVmClusterUpdateAvailableActionsPrecheck,
	"ROLLBACK":          ExadbVmClusterUpdateAvailableActionsRollback,
}

var mappingExadbVmClusterUpdateAvailableActionsEnumLowerCase = map[string]ExadbVmClusterUpdateAvailableActionsEnum{
	"rolling_apply":     ExadbVmClusterUpdateAvailableActionsRollingApply,
	"non_rolling_apply": ExadbVmClusterUpdateAvailableActionsNonRollingApply,
	"precheck":          ExadbVmClusterUpdateAvailableActionsPrecheck,
	"rollback":          ExadbVmClusterUpdateAvailableActionsRollback,
}

// GetExadbVmClusterUpdateAvailableActionsEnumValues Enumerates the set of values for ExadbVmClusterUpdateAvailableActionsEnum
func GetExadbVmClusterUpdateAvailableActionsEnumValues() []ExadbVmClusterUpdateAvailableActionsEnum {
	values := make([]ExadbVmClusterUpdateAvailableActionsEnum, 0)
	for _, v := range mappingExadbVmClusterUpdateAvailableActionsEnum {
		values = append(values, v)
	}
	return values
}

// GetExadbVmClusterUpdateAvailableActionsEnumStringValues Enumerates the set of values in String for ExadbVmClusterUpdateAvailableActionsEnum
func GetExadbVmClusterUpdateAvailableActionsEnumStringValues() []string {
	return []string{
		"ROLLING_APPLY",
		"NON_ROLLING_APPLY",
		"PRECHECK",
		"ROLLBACK",
	}
}

// GetMappingExadbVmClusterUpdateAvailableActionsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExadbVmClusterUpdateAvailableActionsEnum(val string) (ExadbVmClusterUpdateAvailableActionsEnum, bool) {
	enum, ok := mappingExadbVmClusterUpdateAvailableActionsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ExadbVmClusterUpdateUpdateTypeEnum Enum with underlying type: string
type ExadbVmClusterUpdateUpdateTypeEnum string

// Set of constants representing the allowable values for ExadbVmClusterUpdateUpdateTypeEnum
const (
	ExadbVmClusterUpdateUpdateTypeGiUpgrade ExadbVmClusterUpdateUpdateTypeEnum = "GI_UPGRADE"
	ExadbVmClusterUpdateUpdateTypeGiPatch   ExadbVmClusterUpdateUpdateTypeEnum = "GI_PATCH"
	ExadbVmClusterUpdateUpdateTypeOsUpdate  ExadbVmClusterUpdateUpdateTypeEnum = "OS_UPDATE"
)

var mappingExadbVmClusterUpdateUpdateTypeEnum = map[string]ExadbVmClusterUpdateUpdateTypeEnum{
	"GI_UPGRADE": ExadbVmClusterUpdateUpdateTypeGiUpgrade,
	"GI_PATCH":   ExadbVmClusterUpdateUpdateTypeGiPatch,
	"OS_UPDATE":  ExadbVmClusterUpdateUpdateTypeOsUpdate,
}

var mappingExadbVmClusterUpdateUpdateTypeEnumLowerCase = map[string]ExadbVmClusterUpdateUpdateTypeEnum{
	"gi_upgrade": ExadbVmClusterUpdateUpdateTypeGiUpgrade,
	"gi_patch":   ExadbVmClusterUpdateUpdateTypeGiPatch,
	"os_update":  ExadbVmClusterUpdateUpdateTypeOsUpdate,
}

// GetExadbVmClusterUpdateUpdateTypeEnumValues Enumerates the set of values for ExadbVmClusterUpdateUpdateTypeEnum
func GetExadbVmClusterUpdateUpdateTypeEnumValues() []ExadbVmClusterUpdateUpdateTypeEnum {
	values := make([]ExadbVmClusterUpdateUpdateTypeEnum, 0)
	for _, v := range mappingExadbVmClusterUpdateUpdateTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetExadbVmClusterUpdateUpdateTypeEnumStringValues Enumerates the set of values in String for ExadbVmClusterUpdateUpdateTypeEnum
func GetExadbVmClusterUpdateUpdateTypeEnumStringValues() []string {
	return []string{
		"GI_UPGRADE",
		"GI_PATCH",
		"OS_UPDATE",
	}
}

// GetMappingExadbVmClusterUpdateUpdateTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExadbVmClusterUpdateUpdateTypeEnum(val string) (ExadbVmClusterUpdateUpdateTypeEnum, bool) {
	enum, ok := mappingExadbVmClusterUpdateUpdateTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ExadbVmClusterUpdateLifecycleStateEnum Enum with underlying type: string
type ExadbVmClusterUpdateLifecycleStateEnum string

// Set of constants representing the allowable values for ExadbVmClusterUpdateLifecycleStateEnum
const (
	ExadbVmClusterUpdateLifecycleStateAvailable  ExadbVmClusterUpdateLifecycleStateEnum = "AVAILABLE"
	ExadbVmClusterUpdateLifecycleStateSuccess    ExadbVmClusterUpdateLifecycleStateEnum = "SUCCESS"
	ExadbVmClusterUpdateLifecycleStateInProgress ExadbVmClusterUpdateLifecycleStateEnum = "IN_PROGRESS"
	ExadbVmClusterUpdateLifecycleStateFailed     ExadbVmClusterUpdateLifecycleStateEnum = "FAILED"
)

var mappingExadbVmClusterUpdateLifecycleStateEnum = map[string]ExadbVmClusterUpdateLifecycleStateEnum{
	"AVAILABLE":   ExadbVmClusterUpdateLifecycleStateAvailable,
	"SUCCESS":     ExadbVmClusterUpdateLifecycleStateSuccess,
	"IN_PROGRESS": ExadbVmClusterUpdateLifecycleStateInProgress,
	"FAILED":      ExadbVmClusterUpdateLifecycleStateFailed,
}

var mappingExadbVmClusterUpdateLifecycleStateEnumLowerCase = map[string]ExadbVmClusterUpdateLifecycleStateEnum{
	"available":   ExadbVmClusterUpdateLifecycleStateAvailable,
	"success":     ExadbVmClusterUpdateLifecycleStateSuccess,
	"in_progress": ExadbVmClusterUpdateLifecycleStateInProgress,
	"failed":      ExadbVmClusterUpdateLifecycleStateFailed,
}

// GetExadbVmClusterUpdateLifecycleStateEnumValues Enumerates the set of values for ExadbVmClusterUpdateLifecycleStateEnum
func GetExadbVmClusterUpdateLifecycleStateEnumValues() []ExadbVmClusterUpdateLifecycleStateEnum {
	values := make([]ExadbVmClusterUpdateLifecycleStateEnum, 0)
	for _, v := range mappingExadbVmClusterUpdateLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetExadbVmClusterUpdateLifecycleStateEnumStringValues Enumerates the set of values in String for ExadbVmClusterUpdateLifecycleStateEnum
func GetExadbVmClusterUpdateLifecycleStateEnumStringValues() []string {
	return []string{
		"AVAILABLE",
		"SUCCESS",
		"IN_PROGRESS",
		"FAILED",
	}
}

// GetMappingExadbVmClusterUpdateLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExadbVmClusterUpdateLifecycleStateEnum(val string) (ExadbVmClusterUpdateLifecycleStateEnum, bool) {
	enum, ok := mappingExadbVmClusterUpdateLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
