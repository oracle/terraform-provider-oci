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

// ExadbVmClusterUpdateSummary A maintenance update details for an Exadata VM cluster on Exascale Infrastructure.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access,
// see Getting Started with Policies (https://docs.cloud.oracle.com/Content/Identity/Concepts/policygetstarted.htm).
type ExadbVmClusterUpdateSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the maintenance update.
	Id *string `mandatory:"true" json:"id"`

	// Details of the maintenance update package.
	Description *string `mandatory:"true" json:"description"`

	// The type of cloud VM cluster maintenance update.
	UpdateType ExadbVmClusterUpdateSummaryUpdateTypeEnum `mandatory:"true" json:"updateType"`

	// The date and time the maintenance update was released.
	TimeReleased *common.SDKTime `mandatory:"true" json:"timeReleased"`

	// The version of the maintenance update package.
	Version *string `mandatory:"true" json:"version"`

	// The previous update action performed.
	LastAction ExadbVmClusterUpdateSummaryLastActionEnum `mandatory:"false" json:"lastAction,omitempty"`

	// The possible actions performed by the update operation on the infrastructure components.
	AvailableActions []ExadbVmClusterUpdateSummaryAvailableActionsEnum `mandatory:"false" json:"availableActions,omitempty"`

	// Descriptive text providing additional details about the lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The current state of the maintenance update. Dependent on value of `lastAction`.
	LifecycleState ExadbVmClusterUpdateSummaryLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m ExadbVmClusterUpdateSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExadbVmClusterUpdateSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExadbVmClusterUpdateSummaryUpdateTypeEnum(string(m.UpdateType)); !ok && m.UpdateType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpdateType: %s. Supported values are: %s.", m.UpdateType, strings.Join(GetExadbVmClusterUpdateSummaryUpdateTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingExadbVmClusterUpdateSummaryLastActionEnum(string(m.LastAction)); !ok && m.LastAction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LastAction: %s. Supported values are: %s.", m.LastAction, strings.Join(GetExadbVmClusterUpdateSummaryLastActionEnumStringValues(), ",")))
	}
	for _, val := range m.AvailableActions {
		if _, ok := GetMappingExadbVmClusterUpdateSummaryAvailableActionsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AvailableActions: %s. Supported values are: %s.", val, strings.Join(GetExadbVmClusterUpdateSummaryAvailableActionsEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingExadbVmClusterUpdateSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetExadbVmClusterUpdateSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExadbVmClusterUpdateSummaryLastActionEnum Enum with underlying type: string
type ExadbVmClusterUpdateSummaryLastActionEnum string

// Set of constants representing the allowable values for ExadbVmClusterUpdateSummaryLastActionEnum
const (
	ExadbVmClusterUpdateSummaryLastActionRollingApply    ExadbVmClusterUpdateSummaryLastActionEnum = "ROLLING_APPLY"
	ExadbVmClusterUpdateSummaryLastActionNonRollingApply ExadbVmClusterUpdateSummaryLastActionEnum = "NON_ROLLING_APPLY"
	ExadbVmClusterUpdateSummaryLastActionPrecheck        ExadbVmClusterUpdateSummaryLastActionEnum = "PRECHECK"
	ExadbVmClusterUpdateSummaryLastActionRollback        ExadbVmClusterUpdateSummaryLastActionEnum = "ROLLBACK"
)

var mappingExadbVmClusterUpdateSummaryLastActionEnum = map[string]ExadbVmClusterUpdateSummaryLastActionEnum{
	"ROLLING_APPLY":     ExadbVmClusterUpdateSummaryLastActionRollingApply,
	"NON_ROLLING_APPLY": ExadbVmClusterUpdateSummaryLastActionNonRollingApply,
	"PRECHECK":          ExadbVmClusterUpdateSummaryLastActionPrecheck,
	"ROLLBACK":          ExadbVmClusterUpdateSummaryLastActionRollback,
}

var mappingExadbVmClusterUpdateSummaryLastActionEnumLowerCase = map[string]ExadbVmClusterUpdateSummaryLastActionEnum{
	"rolling_apply":     ExadbVmClusterUpdateSummaryLastActionRollingApply,
	"non_rolling_apply": ExadbVmClusterUpdateSummaryLastActionNonRollingApply,
	"precheck":          ExadbVmClusterUpdateSummaryLastActionPrecheck,
	"rollback":          ExadbVmClusterUpdateSummaryLastActionRollback,
}

// GetExadbVmClusterUpdateSummaryLastActionEnumValues Enumerates the set of values for ExadbVmClusterUpdateSummaryLastActionEnum
func GetExadbVmClusterUpdateSummaryLastActionEnumValues() []ExadbVmClusterUpdateSummaryLastActionEnum {
	values := make([]ExadbVmClusterUpdateSummaryLastActionEnum, 0)
	for _, v := range mappingExadbVmClusterUpdateSummaryLastActionEnum {
		values = append(values, v)
	}
	return values
}

// GetExadbVmClusterUpdateSummaryLastActionEnumStringValues Enumerates the set of values in String for ExadbVmClusterUpdateSummaryLastActionEnum
func GetExadbVmClusterUpdateSummaryLastActionEnumStringValues() []string {
	return []string{
		"ROLLING_APPLY",
		"NON_ROLLING_APPLY",
		"PRECHECK",
		"ROLLBACK",
	}
}

// GetMappingExadbVmClusterUpdateSummaryLastActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExadbVmClusterUpdateSummaryLastActionEnum(val string) (ExadbVmClusterUpdateSummaryLastActionEnum, bool) {
	enum, ok := mappingExadbVmClusterUpdateSummaryLastActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ExadbVmClusterUpdateSummaryAvailableActionsEnum Enum with underlying type: string
type ExadbVmClusterUpdateSummaryAvailableActionsEnum string

// Set of constants representing the allowable values for ExadbVmClusterUpdateSummaryAvailableActionsEnum
const (
	ExadbVmClusterUpdateSummaryAvailableActionsRollingApply    ExadbVmClusterUpdateSummaryAvailableActionsEnum = "ROLLING_APPLY"
	ExadbVmClusterUpdateSummaryAvailableActionsNonRollingApply ExadbVmClusterUpdateSummaryAvailableActionsEnum = "NON_ROLLING_APPLY"
	ExadbVmClusterUpdateSummaryAvailableActionsPrecheck        ExadbVmClusterUpdateSummaryAvailableActionsEnum = "PRECHECK"
	ExadbVmClusterUpdateSummaryAvailableActionsRollback        ExadbVmClusterUpdateSummaryAvailableActionsEnum = "ROLLBACK"
)

var mappingExadbVmClusterUpdateSummaryAvailableActionsEnum = map[string]ExadbVmClusterUpdateSummaryAvailableActionsEnum{
	"ROLLING_APPLY":     ExadbVmClusterUpdateSummaryAvailableActionsRollingApply,
	"NON_ROLLING_APPLY": ExadbVmClusterUpdateSummaryAvailableActionsNonRollingApply,
	"PRECHECK":          ExadbVmClusterUpdateSummaryAvailableActionsPrecheck,
	"ROLLBACK":          ExadbVmClusterUpdateSummaryAvailableActionsRollback,
}

var mappingExadbVmClusterUpdateSummaryAvailableActionsEnumLowerCase = map[string]ExadbVmClusterUpdateSummaryAvailableActionsEnum{
	"rolling_apply":     ExadbVmClusterUpdateSummaryAvailableActionsRollingApply,
	"non_rolling_apply": ExadbVmClusterUpdateSummaryAvailableActionsNonRollingApply,
	"precheck":          ExadbVmClusterUpdateSummaryAvailableActionsPrecheck,
	"rollback":          ExadbVmClusterUpdateSummaryAvailableActionsRollback,
}

// GetExadbVmClusterUpdateSummaryAvailableActionsEnumValues Enumerates the set of values for ExadbVmClusterUpdateSummaryAvailableActionsEnum
func GetExadbVmClusterUpdateSummaryAvailableActionsEnumValues() []ExadbVmClusterUpdateSummaryAvailableActionsEnum {
	values := make([]ExadbVmClusterUpdateSummaryAvailableActionsEnum, 0)
	for _, v := range mappingExadbVmClusterUpdateSummaryAvailableActionsEnum {
		values = append(values, v)
	}
	return values
}

// GetExadbVmClusterUpdateSummaryAvailableActionsEnumStringValues Enumerates the set of values in String for ExadbVmClusterUpdateSummaryAvailableActionsEnum
func GetExadbVmClusterUpdateSummaryAvailableActionsEnumStringValues() []string {
	return []string{
		"ROLLING_APPLY",
		"NON_ROLLING_APPLY",
		"PRECHECK",
		"ROLLBACK",
	}
}

// GetMappingExadbVmClusterUpdateSummaryAvailableActionsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExadbVmClusterUpdateSummaryAvailableActionsEnum(val string) (ExadbVmClusterUpdateSummaryAvailableActionsEnum, bool) {
	enum, ok := mappingExadbVmClusterUpdateSummaryAvailableActionsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ExadbVmClusterUpdateSummaryUpdateTypeEnum Enum with underlying type: string
type ExadbVmClusterUpdateSummaryUpdateTypeEnum string

// Set of constants representing the allowable values for ExadbVmClusterUpdateSummaryUpdateTypeEnum
const (
	ExadbVmClusterUpdateSummaryUpdateTypeGiUpgrade ExadbVmClusterUpdateSummaryUpdateTypeEnum = "GI_UPGRADE"
	ExadbVmClusterUpdateSummaryUpdateTypeGiPatch   ExadbVmClusterUpdateSummaryUpdateTypeEnum = "GI_PATCH"
	ExadbVmClusterUpdateSummaryUpdateTypeOsUpdate  ExadbVmClusterUpdateSummaryUpdateTypeEnum = "OS_UPDATE"
)

var mappingExadbVmClusterUpdateSummaryUpdateTypeEnum = map[string]ExadbVmClusterUpdateSummaryUpdateTypeEnum{
	"GI_UPGRADE": ExadbVmClusterUpdateSummaryUpdateTypeGiUpgrade,
	"GI_PATCH":   ExadbVmClusterUpdateSummaryUpdateTypeGiPatch,
	"OS_UPDATE":  ExadbVmClusterUpdateSummaryUpdateTypeOsUpdate,
}

var mappingExadbVmClusterUpdateSummaryUpdateTypeEnumLowerCase = map[string]ExadbVmClusterUpdateSummaryUpdateTypeEnum{
	"gi_upgrade": ExadbVmClusterUpdateSummaryUpdateTypeGiUpgrade,
	"gi_patch":   ExadbVmClusterUpdateSummaryUpdateTypeGiPatch,
	"os_update":  ExadbVmClusterUpdateSummaryUpdateTypeOsUpdate,
}

// GetExadbVmClusterUpdateSummaryUpdateTypeEnumValues Enumerates the set of values for ExadbVmClusterUpdateSummaryUpdateTypeEnum
func GetExadbVmClusterUpdateSummaryUpdateTypeEnumValues() []ExadbVmClusterUpdateSummaryUpdateTypeEnum {
	values := make([]ExadbVmClusterUpdateSummaryUpdateTypeEnum, 0)
	for _, v := range mappingExadbVmClusterUpdateSummaryUpdateTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetExadbVmClusterUpdateSummaryUpdateTypeEnumStringValues Enumerates the set of values in String for ExadbVmClusterUpdateSummaryUpdateTypeEnum
func GetExadbVmClusterUpdateSummaryUpdateTypeEnumStringValues() []string {
	return []string{
		"GI_UPGRADE",
		"GI_PATCH",
		"OS_UPDATE",
	}
}

// GetMappingExadbVmClusterUpdateSummaryUpdateTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExadbVmClusterUpdateSummaryUpdateTypeEnum(val string) (ExadbVmClusterUpdateSummaryUpdateTypeEnum, bool) {
	enum, ok := mappingExadbVmClusterUpdateSummaryUpdateTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ExadbVmClusterUpdateSummaryLifecycleStateEnum Enum with underlying type: string
type ExadbVmClusterUpdateSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for ExadbVmClusterUpdateSummaryLifecycleStateEnum
const (
	ExadbVmClusterUpdateSummaryLifecycleStateAvailable  ExadbVmClusterUpdateSummaryLifecycleStateEnum = "AVAILABLE"
	ExadbVmClusterUpdateSummaryLifecycleStateSuccess    ExadbVmClusterUpdateSummaryLifecycleStateEnum = "SUCCESS"
	ExadbVmClusterUpdateSummaryLifecycleStateInProgress ExadbVmClusterUpdateSummaryLifecycleStateEnum = "IN_PROGRESS"
	ExadbVmClusterUpdateSummaryLifecycleStateFailed     ExadbVmClusterUpdateSummaryLifecycleStateEnum = "FAILED"
)

var mappingExadbVmClusterUpdateSummaryLifecycleStateEnum = map[string]ExadbVmClusterUpdateSummaryLifecycleStateEnum{
	"AVAILABLE":   ExadbVmClusterUpdateSummaryLifecycleStateAvailable,
	"SUCCESS":     ExadbVmClusterUpdateSummaryLifecycleStateSuccess,
	"IN_PROGRESS": ExadbVmClusterUpdateSummaryLifecycleStateInProgress,
	"FAILED":      ExadbVmClusterUpdateSummaryLifecycleStateFailed,
}

var mappingExadbVmClusterUpdateSummaryLifecycleStateEnumLowerCase = map[string]ExadbVmClusterUpdateSummaryLifecycleStateEnum{
	"available":   ExadbVmClusterUpdateSummaryLifecycleStateAvailable,
	"success":     ExadbVmClusterUpdateSummaryLifecycleStateSuccess,
	"in_progress": ExadbVmClusterUpdateSummaryLifecycleStateInProgress,
	"failed":      ExadbVmClusterUpdateSummaryLifecycleStateFailed,
}

// GetExadbVmClusterUpdateSummaryLifecycleStateEnumValues Enumerates the set of values for ExadbVmClusterUpdateSummaryLifecycleStateEnum
func GetExadbVmClusterUpdateSummaryLifecycleStateEnumValues() []ExadbVmClusterUpdateSummaryLifecycleStateEnum {
	values := make([]ExadbVmClusterUpdateSummaryLifecycleStateEnum, 0)
	for _, v := range mappingExadbVmClusterUpdateSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetExadbVmClusterUpdateSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for ExadbVmClusterUpdateSummaryLifecycleStateEnum
func GetExadbVmClusterUpdateSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"AVAILABLE",
		"SUCCESS",
		"IN_PROGRESS",
		"FAILED",
	}
}

// GetMappingExadbVmClusterUpdateSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExadbVmClusterUpdateSummaryLifecycleStateEnum(val string) (ExadbVmClusterUpdateSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingExadbVmClusterUpdateSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
