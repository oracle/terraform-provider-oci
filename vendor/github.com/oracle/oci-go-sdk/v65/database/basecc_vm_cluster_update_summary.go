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

// BaseccVmClusterUpdateSummary Maintenance update details for a BaseDB-C@C VM cluster.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access,
// see Getting Started with Policies (https://docs.oracle.com/iaas/Content/Identity/Concepts/policygetstarted.htm).
type BaseccVmClusterUpdateSummary struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the maintenance update.
	Id *string `mandatory:"true" json:"id"`

	// Details of the maintenance update package.
	Description *string `mandatory:"true" json:"description"`

	// The type of Base Database Service on Cloud@Customer (BaseDB-C@C) VM cluster maintenance update.
	UpdateType BaseccVmClusterUpdateSummaryUpdateTypeEnum `mandatory:"true" json:"updateType"`

	// The date and time the maintenance update was released.
	TimeReleased *common.SDKTime `mandatory:"true" json:"timeReleased"`

	// The version of the maintenance update package.
	Version *string `mandatory:"true" json:"version"`

	// The previous update action performed.
	LastAction BaseccVmClusterUpdateSummaryLastActionEnum `mandatory:"false" json:"lastAction,omitempty"`

	// The possible actions performed by the update operation on the infrastructure components.
	AvailableActions []BaseccVmClusterUpdateSummaryAvailableActionsEnum `mandatory:"false" json:"availableActions,omitempty"`

	// Descriptive text providing additional details about the lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The current state of the maintenance update. Dependent on value of `lastAction`.
	LifecycleState BaseccVmClusterUpdateSummaryLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m BaseccVmClusterUpdateSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BaseccVmClusterUpdateSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBaseccVmClusterUpdateSummaryUpdateTypeEnum(string(m.UpdateType)); !ok && m.UpdateType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpdateType: %s. Supported values are: %s.", m.UpdateType, strings.Join(GetBaseccVmClusterUpdateSummaryUpdateTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingBaseccVmClusterUpdateSummaryLastActionEnum(string(m.LastAction)); !ok && m.LastAction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LastAction: %s. Supported values are: %s.", m.LastAction, strings.Join(GetBaseccVmClusterUpdateSummaryLastActionEnumStringValues(), ",")))
	}
	for _, val := range m.AvailableActions {
		if _, ok := GetMappingBaseccVmClusterUpdateSummaryAvailableActionsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AvailableActions: %s. Supported values are: %s.", val, strings.Join(GetBaseccVmClusterUpdateSummaryAvailableActionsEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingBaseccVmClusterUpdateSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetBaseccVmClusterUpdateSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BaseccVmClusterUpdateSummaryLastActionEnum Enum with underlying type: string
type BaseccVmClusterUpdateSummaryLastActionEnum string

// Set of constants representing the allowable values for BaseccVmClusterUpdateSummaryLastActionEnum
const (
	BaseccVmClusterUpdateSummaryLastActionRollingApply BaseccVmClusterUpdateSummaryLastActionEnum = "ROLLING_APPLY"
	BaseccVmClusterUpdateSummaryLastActionPrecheck     BaseccVmClusterUpdateSummaryLastActionEnum = "PRECHECK"
	BaseccVmClusterUpdateSummaryLastActionRollback     BaseccVmClusterUpdateSummaryLastActionEnum = "ROLLBACK"
)

var mappingBaseccVmClusterUpdateSummaryLastActionEnum = map[string]BaseccVmClusterUpdateSummaryLastActionEnum{
	"ROLLING_APPLY": BaseccVmClusterUpdateSummaryLastActionRollingApply,
	"PRECHECK":      BaseccVmClusterUpdateSummaryLastActionPrecheck,
	"ROLLBACK":      BaseccVmClusterUpdateSummaryLastActionRollback,
}

var mappingBaseccVmClusterUpdateSummaryLastActionEnumLowerCase = map[string]BaseccVmClusterUpdateSummaryLastActionEnum{
	"rolling_apply": BaseccVmClusterUpdateSummaryLastActionRollingApply,
	"precheck":      BaseccVmClusterUpdateSummaryLastActionPrecheck,
	"rollback":      BaseccVmClusterUpdateSummaryLastActionRollback,
}

// GetBaseccVmClusterUpdateSummaryLastActionEnumValues Enumerates the set of values for BaseccVmClusterUpdateSummaryLastActionEnum
func GetBaseccVmClusterUpdateSummaryLastActionEnumValues() []BaseccVmClusterUpdateSummaryLastActionEnum {
	values := make([]BaseccVmClusterUpdateSummaryLastActionEnum, 0)
	for _, v := range mappingBaseccVmClusterUpdateSummaryLastActionEnum {
		values = append(values, v)
	}
	return values
}

// GetBaseccVmClusterUpdateSummaryLastActionEnumStringValues Enumerates the set of values in String for BaseccVmClusterUpdateSummaryLastActionEnum
func GetBaseccVmClusterUpdateSummaryLastActionEnumStringValues() []string {
	return []string{
		"ROLLING_APPLY",
		"PRECHECK",
		"ROLLBACK",
	}
}

// GetMappingBaseccVmClusterUpdateSummaryLastActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBaseccVmClusterUpdateSummaryLastActionEnum(val string) (BaseccVmClusterUpdateSummaryLastActionEnum, bool) {
	enum, ok := mappingBaseccVmClusterUpdateSummaryLastActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// BaseccVmClusterUpdateSummaryAvailableActionsEnum Enum with underlying type: string
type BaseccVmClusterUpdateSummaryAvailableActionsEnum string

// Set of constants representing the allowable values for BaseccVmClusterUpdateSummaryAvailableActionsEnum
const (
	BaseccVmClusterUpdateSummaryAvailableActionsRollingApply BaseccVmClusterUpdateSummaryAvailableActionsEnum = "ROLLING_APPLY"
	BaseccVmClusterUpdateSummaryAvailableActionsPrecheck     BaseccVmClusterUpdateSummaryAvailableActionsEnum = "PRECHECK"
	BaseccVmClusterUpdateSummaryAvailableActionsRollback     BaseccVmClusterUpdateSummaryAvailableActionsEnum = "ROLLBACK"
)

var mappingBaseccVmClusterUpdateSummaryAvailableActionsEnum = map[string]BaseccVmClusterUpdateSummaryAvailableActionsEnum{
	"ROLLING_APPLY": BaseccVmClusterUpdateSummaryAvailableActionsRollingApply,
	"PRECHECK":      BaseccVmClusterUpdateSummaryAvailableActionsPrecheck,
	"ROLLBACK":      BaseccVmClusterUpdateSummaryAvailableActionsRollback,
}

var mappingBaseccVmClusterUpdateSummaryAvailableActionsEnumLowerCase = map[string]BaseccVmClusterUpdateSummaryAvailableActionsEnum{
	"rolling_apply": BaseccVmClusterUpdateSummaryAvailableActionsRollingApply,
	"precheck":      BaseccVmClusterUpdateSummaryAvailableActionsPrecheck,
	"rollback":      BaseccVmClusterUpdateSummaryAvailableActionsRollback,
}

// GetBaseccVmClusterUpdateSummaryAvailableActionsEnumValues Enumerates the set of values for BaseccVmClusterUpdateSummaryAvailableActionsEnum
func GetBaseccVmClusterUpdateSummaryAvailableActionsEnumValues() []BaseccVmClusterUpdateSummaryAvailableActionsEnum {
	values := make([]BaseccVmClusterUpdateSummaryAvailableActionsEnum, 0)
	for _, v := range mappingBaseccVmClusterUpdateSummaryAvailableActionsEnum {
		values = append(values, v)
	}
	return values
}

// GetBaseccVmClusterUpdateSummaryAvailableActionsEnumStringValues Enumerates the set of values in String for BaseccVmClusterUpdateSummaryAvailableActionsEnum
func GetBaseccVmClusterUpdateSummaryAvailableActionsEnumStringValues() []string {
	return []string{
		"ROLLING_APPLY",
		"PRECHECK",
		"ROLLBACK",
	}
}

// GetMappingBaseccVmClusterUpdateSummaryAvailableActionsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBaseccVmClusterUpdateSummaryAvailableActionsEnum(val string) (BaseccVmClusterUpdateSummaryAvailableActionsEnum, bool) {
	enum, ok := mappingBaseccVmClusterUpdateSummaryAvailableActionsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// BaseccVmClusterUpdateSummaryUpdateTypeEnum Enum with underlying type: string
type BaseccVmClusterUpdateSummaryUpdateTypeEnum string

// Set of constants representing the allowable values for BaseccVmClusterUpdateSummaryUpdateTypeEnum
const (
	BaseccVmClusterUpdateSummaryUpdateTypeUpgrade BaseccVmClusterUpdateSummaryUpdateTypeEnum = "GI_UPGRADE"
	BaseccVmClusterUpdateSummaryUpdateTypePatch   BaseccVmClusterUpdateSummaryUpdateTypeEnum = "GI_PATCH"
)

var mappingBaseccVmClusterUpdateSummaryUpdateTypeEnum = map[string]BaseccVmClusterUpdateSummaryUpdateTypeEnum{
	"GI_UPGRADE": BaseccVmClusterUpdateSummaryUpdateTypeUpgrade,
	"GI_PATCH":   BaseccVmClusterUpdateSummaryUpdateTypePatch,
}

var mappingBaseccVmClusterUpdateSummaryUpdateTypeEnumLowerCase = map[string]BaseccVmClusterUpdateSummaryUpdateTypeEnum{
	"gi_upgrade": BaseccVmClusterUpdateSummaryUpdateTypeUpgrade,
	"gi_patch":   BaseccVmClusterUpdateSummaryUpdateTypePatch,
}

// GetBaseccVmClusterUpdateSummaryUpdateTypeEnumValues Enumerates the set of values for BaseccVmClusterUpdateSummaryUpdateTypeEnum
func GetBaseccVmClusterUpdateSummaryUpdateTypeEnumValues() []BaseccVmClusterUpdateSummaryUpdateTypeEnum {
	values := make([]BaseccVmClusterUpdateSummaryUpdateTypeEnum, 0)
	for _, v := range mappingBaseccVmClusterUpdateSummaryUpdateTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetBaseccVmClusterUpdateSummaryUpdateTypeEnumStringValues Enumerates the set of values in String for BaseccVmClusterUpdateSummaryUpdateTypeEnum
func GetBaseccVmClusterUpdateSummaryUpdateTypeEnumStringValues() []string {
	return []string{
		"GI_UPGRADE",
		"GI_PATCH",
	}
}

// GetMappingBaseccVmClusterUpdateSummaryUpdateTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBaseccVmClusterUpdateSummaryUpdateTypeEnum(val string) (BaseccVmClusterUpdateSummaryUpdateTypeEnum, bool) {
	enum, ok := mappingBaseccVmClusterUpdateSummaryUpdateTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// BaseccVmClusterUpdateSummaryLifecycleStateEnum Enum with underlying type: string
type BaseccVmClusterUpdateSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for BaseccVmClusterUpdateSummaryLifecycleStateEnum
const (
	BaseccVmClusterUpdateSummaryLifecycleStateAvailable  BaseccVmClusterUpdateSummaryLifecycleStateEnum = "AVAILABLE"
	BaseccVmClusterUpdateSummaryLifecycleStateSuccess    BaseccVmClusterUpdateSummaryLifecycleStateEnum = "SUCCESS"
	BaseccVmClusterUpdateSummaryLifecycleStateInProgress BaseccVmClusterUpdateSummaryLifecycleStateEnum = "IN_PROGRESS"
	BaseccVmClusterUpdateSummaryLifecycleStateFailed     BaseccVmClusterUpdateSummaryLifecycleStateEnum = "FAILED"
)

var mappingBaseccVmClusterUpdateSummaryLifecycleStateEnum = map[string]BaseccVmClusterUpdateSummaryLifecycleStateEnum{
	"AVAILABLE":   BaseccVmClusterUpdateSummaryLifecycleStateAvailable,
	"SUCCESS":     BaseccVmClusterUpdateSummaryLifecycleStateSuccess,
	"IN_PROGRESS": BaseccVmClusterUpdateSummaryLifecycleStateInProgress,
	"FAILED":      BaseccVmClusterUpdateSummaryLifecycleStateFailed,
}

var mappingBaseccVmClusterUpdateSummaryLifecycleStateEnumLowerCase = map[string]BaseccVmClusterUpdateSummaryLifecycleStateEnum{
	"available":   BaseccVmClusterUpdateSummaryLifecycleStateAvailable,
	"success":     BaseccVmClusterUpdateSummaryLifecycleStateSuccess,
	"in_progress": BaseccVmClusterUpdateSummaryLifecycleStateInProgress,
	"failed":      BaseccVmClusterUpdateSummaryLifecycleStateFailed,
}

// GetBaseccVmClusterUpdateSummaryLifecycleStateEnumValues Enumerates the set of values for BaseccVmClusterUpdateSummaryLifecycleStateEnum
func GetBaseccVmClusterUpdateSummaryLifecycleStateEnumValues() []BaseccVmClusterUpdateSummaryLifecycleStateEnum {
	values := make([]BaseccVmClusterUpdateSummaryLifecycleStateEnum, 0)
	for _, v := range mappingBaseccVmClusterUpdateSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetBaseccVmClusterUpdateSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for BaseccVmClusterUpdateSummaryLifecycleStateEnum
func GetBaseccVmClusterUpdateSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"AVAILABLE",
		"SUCCESS",
		"IN_PROGRESS",
		"FAILED",
	}
}

// GetMappingBaseccVmClusterUpdateSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBaseccVmClusterUpdateSummaryLifecycleStateEnum(val string) (BaseccVmClusterUpdateSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingBaseccVmClusterUpdateSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
