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

// ExadbVmClusterUpdateHistoryEntrySummary The record of an maintenance update action on a specified Exadata VM cluster on Exascale Infrastructure.
type ExadbVmClusterUpdateHistoryEntrySummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the maintenance update history entry.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the maintenance update.
	UpdateId *string `mandatory:"true" json:"updateId"`

	// The type of cloud VM cluster maintenance update.
	UpdateType ExadbVmClusterUpdateHistoryEntrySummaryUpdateTypeEnum `mandatory:"true" json:"updateType"`

	// The current lifecycle state of the maintenance update operation.
	LifecycleState ExadbVmClusterUpdateHistoryEntrySummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time when the maintenance update action started.
	TimeStarted *common.SDKTime `mandatory:"true" json:"timeStarted"`

	// The update action.
	UpdateAction ExadbVmClusterUpdateHistoryEntrySummaryUpdateActionEnum `mandatory:"false" json:"updateAction,omitempty"`

	// Descriptive text providing additional details about the lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The date and time when the maintenance update action completed.
	TimeCompleted *common.SDKTime `mandatory:"false" json:"timeCompleted"`

	// The version of the maintenance update package.
	Version *string `mandatory:"false" json:"version"`
}

func (m ExadbVmClusterUpdateHistoryEntrySummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExadbVmClusterUpdateHistoryEntrySummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExadbVmClusterUpdateHistoryEntrySummaryUpdateTypeEnum(string(m.UpdateType)); !ok && m.UpdateType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpdateType: %s. Supported values are: %s.", m.UpdateType, strings.Join(GetExadbVmClusterUpdateHistoryEntrySummaryUpdateTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingExadbVmClusterUpdateHistoryEntrySummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetExadbVmClusterUpdateHistoryEntrySummaryLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingExadbVmClusterUpdateHistoryEntrySummaryUpdateActionEnum(string(m.UpdateAction)); !ok && m.UpdateAction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpdateAction: %s. Supported values are: %s.", m.UpdateAction, strings.Join(GetExadbVmClusterUpdateHistoryEntrySummaryUpdateActionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExadbVmClusterUpdateHistoryEntrySummaryUpdateActionEnum Enum with underlying type: string
type ExadbVmClusterUpdateHistoryEntrySummaryUpdateActionEnum string

// Set of constants representing the allowable values for ExadbVmClusterUpdateHistoryEntrySummaryUpdateActionEnum
const (
	ExadbVmClusterUpdateHistoryEntrySummaryUpdateActionRollingApply    ExadbVmClusterUpdateHistoryEntrySummaryUpdateActionEnum = "ROLLING_APPLY"
	ExadbVmClusterUpdateHistoryEntrySummaryUpdateActionNonRollingApply ExadbVmClusterUpdateHistoryEntrySummaryUpdateActionEnum = "NON_ROLLING_APPLY"
	ExadbVmClusterUpdateHistoryEntrySummaryUpdateActionPrecheck        ExadbVmClusterUpdateHistoryEntrySummaryUpdateActionEnum = "PRECHECK"
	ExadbVmClusterUpdateHistoryEntrySummaryUpdateActionRollback        ExadbVmClusterUpdateHistoryEntrySummaryUpdateActionEnum = "ROLLBACK"
)

var mappingExadbVmClusterUpdateHistoryEntrySummaryUpdateActionEnum = map[string]ExadbVmClusterUpdateHistoryEntrySummaryUpdateActionEnum{
	"ROLLING_APPLY":     ExadbVmClusterUpdateHistoryEntrySummaryUpdateActionRollingApply,
	"NON_ROLLING_APPLY": ExadbVmClusterUpdateHistoryEntrySummaryUpdateActionNonRollingApply,
	"PRECHECK":          ExadbVmClusterUpdateHistoryEntrySummaryUpdateActionPrecheck,
	"ROLLBACK":          ExadbVmClusterUpdateHistoryEntrySummaryUpdateActionRollback,
}

var mappingExadbVmClusterUpdateHistoryEntrySummaryUpdateActionEnumLowerCase = map[string]ExadbVmClusterUpdateHistoryEntrySummaryUpdateActionEnum{
	"rolling_apply":     ExadbVmClusterUpdateHistoryEntrySummaryUpdateActionRollingApply,
	"non_rolling_apply": ExadbVmClusterUpdateHistoryEntrySummaryUpdateActionNonRollingApply,
	"precheck":          ExadbVmClusterUpdateHistoryEntrySummaryUpdateActionPrecheck,
	"rollback":          ExadbVmClusterUpdateHistoryEntrySummaryUpdateActionRollback,
}

// GetExadbVmClusterUpdateHistoryEntrySummaryUpdateActionEnumValues Enumerates the set of values for ExadbVmClusterUpdateHistoryEntrySummaryUpdateActionEnum
func GetExadbVmClusterUpdateHistoryEntrySummaryUpdateActionEnumValues() []ExadbVmClusterUpdateHistoryEntrySummaryUpdateActionEnum {
	values := make([]ExadbVmClusterUpdateHistoryEntrySummaryUpdateActionEnum, 0)
	for _, v := range mappingExadbVmClusterUpdateHistoryEntrySummaryUpdateActionEnum {
		values = append(values, v)
	}
	return values
}

// GetExadbVmClusterUpdateHistoryEntrySummaryUpdateActionEnumStringValues Enumerates the set of values in String for ExadbVmClusterUpdateHistoryEntrySummaryUpdateActionEnum
func GetExadbVmClusterUpdateHistoryEntrySummaryUpdateActionEnumStringValues() []string {
	return []string{
		"ROLLING_APPLY",
		"NON_ROLLING_APPLY",
		"PRECHECK",
		"ROLLBACK",
	}
}

// GetMappingExadbVmClusterUpdateHistoryEntrySummaryUpdateActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExadbVmClusterUpdateHistoryEntrySummaryUpdateActionEnum(val string) (ExadbVmClusterUpdateHistoryEntrySummaryUpdateActionEnum, bool) {
	enum, ok := mappingExadbVmClusterUpdateHistoryEntrySummaryUpdateActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ExadbVmClusterUpdateHistoryEntrySummaryUpdateTypeEnum Enum with underlying type: string
type ExadbVmClusterUpdateHistoryEntrySummaryUpdateTypeEnum string

// Set of constants representing the allowable values for ExadbVmClusterUpdateHistoryEntrySummaryUpdateTypeEnum
const (
	ExadbVmClusterUpdateHistoryEntrySummaryUpdateTypeGiUpgrade ExadbVmClusterUpdateHistoryEntrySummaryUpdateTypeEnum = "GI_UPGRADE"
	ExadbVmClusterUpdateHistoryEntrySummaryUpdateTypeGiPatch   ExadbVmClusterUpdateHistoryEntrySummaryUpdateTypeEnum = "GI_PATCH"
	ExadbVmClusterUpdateHistoryEntrySummaryUpdateTypeOsUpdate  ExadbVmClusterUpdateHistoryEntrySummaryUpdateTypeEnum = "OS_UPDATE"
)

var mappingExadbVmClusterUpdateHistoryEntrySummaryUpdateTypeEnum = map[string]ExadbVmClusterUpdateHistoryEntrySummaryUpdateTypeEnum{
	"GI_UPGRADE": ExadbVmClusterUpdateHistoryEntrySummaryUpdateTypeGiUpgrade,
	"GI_PATCH":   ExadbVmClusterUpdateHistoryEntrySummaryUpdateTypeGiPatch,
	"OS_UPDATE":  ExadbVmClusterUpdateHistoryEntrySummaryUpdateTypeOsUpdate,
}

var mappingExadbVmClusterUpdateHistoryEntrySummaryUpdateTypeEnumLowerCase = map[string]ExadbVmClusterUpdateHistoryEntrySummaryUpdateTypeEnum{
	"gi_upgrade": ExadbVmClusterUpdateHistoryEntrySummaryUpdateTypeGiUpgrade,
	"gi_patch":   ExadbVmClusterUpdateHistoryEntrySummaryUpdateTypeGiPatch,
	"os_update":  ExadbVmClusterUpdateHistoryEntrySummaryUpdateTypeOsUpdate,
}

// GetExadbVmClusterUpdateHistoryEntrySummaryUpdateTypeEnumValues Enumerates the set of values for ExadbVmClusterUpdateHistoryEntrySummaryUpdateTypeEnum
func GetExadbVmClusterUpdateHistoryEntrySummaryUpdateTypeEnumValues() []ExadbVmClusterUpdateHistoryEntrySummaryUpdateTypeEnum {
	values := make([]ExadbVmClusterUpdateHistoryEntrySummaryUpdateTypeEnum, 0)
	for _, v := range mappingExadbVmClusterUpdateHistoryEntrySummaryUpdateTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetExadbVmClusterUpdateHistoryEntrySummaryUpdateTypeEnumStringValues Enumerates the set of values in String for ExadbVmClusterUpdateHistoryEntrySummaryUpdateTypeEnum
func GetExadbVmClusterUpdateHistoryEntrySummaryUpdateTypeEnumStringValues() []string {
	return []string{
		"GI_UPGRADE",
		"GI_PATCH",
		"OS_UPDATE",
	}
}

// GetMappingExadbVmClusterUpdateHistoryEntrySummaryUpdateTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExadbVmClusterUpdateHistoryEntrySummaryUpdateTypeEnum(val string) (ExadbVmClusterUpdateHistoryEntrySummaryUpdateTypeEnum, bool) {
	enum, ok := mappingExadbVmClusterUpdateHistoryEntrySummaryUpdateTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ExadbVmClusterUpdateHistoryEntrySummaryLifecycleStateEnum Enum with underlying type: string
type ExadbVmClusterUpdateHistoryEntrySummaryLifecycleStateEnum string

// Set of constants representing the allowable values for ExadbVmClusterUpdateHistoryEntrySummaryLifecycleStateEnum
const (
	ExadbVmClusterUpdateHistoryEntrySummaryLifecycleStateInProgress ExadbVmClusterUpdateHistoryEntrySummaryLifecycleStateEnum = "IN_PROGRESS"
	ExadbVmClusterUpdateHistoryEntrySummaryLifecycleStateSucceeded  ExadbVmClusterUpdateHistoryEntrySummaryLifecycleStateEnum = "SUCCEEDED"
	ExadbVmClusterUpdateHistoryEntrySummaryLifecycleStateFailed     ExadbVmClusterUpdateHistoryEntrySummaryLifecycleStateEnum = "FAILED"
)

var mappingExadbVmClusterUpdateHistoryEntrySummaryLifecycleStateEnum = map[string]ExadbVmClusterUpdateHistoryEntrySummaryLifecycleStateEnum{
	"IN_PROGRESS": ExadbVmClusterUpdateHistoryEntrySummaryLifecycleStateInProgress,
	"SUCCEEDED":   ExadbVmClusterUpdateHistoryEntrySummaryLifecycleStateSucceeded,
	"FAILED":      ExadbVmClusterUpdateHistoryEntrySummaryLifecycleStateFailed,
}

var mappingExadbVmClusterUpdateHistoryEntrySummaryLifecycleStateEnumLowerCase = map[string]ExadbVmClusterUpdateHistoryEntrySummaryLifecycleStateEnum{
	"in_progress": ExadbVmClusterUpdateHistoryEntrySummaryLifecycleStateInProgress,
	"succeeded":   ExadbVmClusterUpdateHistoryEntrySummaryLifecycleStateSucceeded,
	"failed":      ExadbVmClusterUpdateHistoryEntrySummaryLifecycleStateFailed,
}

// GetExadbVmClusterUpdateHistoryEntrySummaryLifecycleStateEnumValues Enumerates the set of values for ExadbVmClusterUpdateHistoryEntrySummaryLifecycleStateEnum
func GetExadbVmClusterUpdateHistoryEntrySummaryLifecycleStateEnumValues() []ExadbVmClusterUpdateHistoryEntrySummaryLifecycleStateEnum {
	values := make([]ExadbVmClusterUpdateHistoryEntrySummaryLifecycleStateEnum, 0)
	for _, v := range mappingExadbVmClusterUpdateHistoryEntrySummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetExadbVmClusterUpdateHistoryEntrySummaryLifecycleStateEnumStringValues Enumerates the set of values in String for ExadbVmClusterUpdateHistoryEntrySummaryLifecycleStateEnum
func GetExadbVmClusterUpdateHistoryEntrySummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"IN_PROGRESS",
		"SUCCEEDED",
		"FAILED",
	}
}

// GetMappingExadbVmClusterUpdateHistoryEntrySummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExadbVmClusterUpdateHistoryEntrySummaryLifecycleStateEnum(val string) (ExadbVmClusterUpdateHistoryEntrySummaryLifecycleStateEnum, bool) {
	enum, ok := mappingExadbVmClusterUpdateHistoryEntrySummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
