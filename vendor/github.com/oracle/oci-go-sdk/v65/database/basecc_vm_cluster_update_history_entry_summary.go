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

// BaseccVmClusterUpdateHistoryEntrySummary The record of a maintenance update action on a specified BaseDB-C@C VM cluster.
type BaseccVmClusterUpdateHistoryEntrySummary struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the maintenance update history entry.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the maintenance update.
	UpdateId *string `mandatory:"true" json:"updateId"`

	// The type of Base Database Service on Cloud@Customer (BaseDB-C@C) VM cluster maintenance update.
	UpdateType BaseccVmClusterUpdateHistoryEntrySummaryUpdateTypeEnum `mandatory:"true" json:"updateType"`

	// The current lifecycle state of the maintenance update operation.
	LifecycleState BaseccVmClusterUpdateHistoryEntrySummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time when the maintenance update action started.
	TimeStarted *common.SDKTime `mandatory:"true" json:"timeStarted"`

	// The update action.
	UpdateAction BaseccVmClusterUpdateHistoryEntrySummaryUpdateActionEnum `mandatory:"false" json:"updateAction,omitempty"`

	// Descriptive text providing additional details about the lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The date and time when the maintenance update action completed.
	TimeCompleted *common.SDKTime `mandatory:"false" json:"timeCompleted"`

	// The version of the maintenance update package.
	Version *string `mandatory:"false" json:"version"`
}

func (m BaseccVmClusterUpdateHistoryEntrySummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BaseccVmClusterUpdateHistoryEntrySummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBaseccVmClusterUpdateHistoryEntrySummaryUpdateTypeEnum(string(m.UpdateType)); !ok && m.UpdateType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpdateType: %s. Supported values are: %s.", m.UpdateType, strings.Join(GetBaseccVmClusterUpdateHistoryEntrySummaryUpdateTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingBaseccVmClusterUpdateHistoryEntrySummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetBaseccVmClusterUpdateHistoryEntrySummaryLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingBaseccVmClusterUpdateHistoryEntrySummaryUpdateActionEnum(string(m.UpdateAction)); !ok && m.UpdateAction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpdateAction: %s. Supported values are: %s.", m.UpdateAction, strings.Join(GetBaseccVmClusterUpdateHistoryEntrySummaryUpdateActionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BaseccVmClusterUpdateHistoryEntrySummaryUpdateActionEnum Enum with underlying type: string
type BaseccVmClusterUpdateHistoryEntrySummaryUpdateActionEnum string

// Set of constants representing the allowable values for BaseccVmClusterUpdateHistoryEntrySummaryUpdateActionEnum
const (
	BaseccVmClusterUpdateHistoryEntrySummaryUpdateActionRollingApply BaseccVmClusterUpdateHistoryEntrySummaryUpdateActionEnum = "ROLLING_APPLY"
	BaseccVmClusterUpdateHistoryEntrySummaryUpdateActionPrecheck     BaseccVmClusterUpdateHistoryEntrySummaryUpdateActionEnum = "PRECHECK"
	BaseccVmClusterUpdateHistoryEntrySummaryUpdateActionRollback     BaseccVmClusterUpdateHistoryEntrySummaryUpdateActionEnum = "ROLLBACK"
)

var mappingBaseccVmClusterUpdateHistoryEntrySummaryUpdateActionEnum = map[string]BaseccVmClusterUpdateHistoryEntrySummaryUpdateActionEnum{
	"ROLLING_APPLY": BaseccVmClusterUpdateHistoryEntrySummaryUpdateActionRollingApply,
	"PRECHECK":      BaseccVmClusterUpdateHistoryEntrySummaryUpdateActionPrecheck,
	"ROLLBACK":      BaseccVmClusterUpdateHistoryEntrySummaryUpdateActionRollback,
}

var mappingBaseccVmClusterUpdateHistoryEntrySummaryUpdateActionEnumLowerCase = map[string]BaseccVmClusterUpdateHistoryEntrySummaryUpdateActionEnum{
	"rolling_apply": BaseccVmClusterUpdateHistoryEntrySummaryUpdateActionRollingApply,
	"precheck":      BaseccVmClusterUpdateHistoryEntrySummaryUpdateActionPrecheck,
	"rollback":      BaseccVmClusterUpdateHistoryEntrySummaryUpdateActionRollback,
}

// GetBaseccVmClusterUpdateHistoryEntrySummaryUpdateActionEnumValues Enumerates the set of values for BaseccVmClusterUpdateHistoryEntrySummaryUpdateActionEnum
func GetBaseccVmClusterUpdateHistoryEntrySummaryUpdateActionEnumValues() []BaseccVmClusterUpdateHistoryEntrySummaryUpdateActionEnum {
	values := make([]BaseccVmClusterUpdateHistoryEntrySummaryUpdateActionEnum, 0)
	for _, v := range mappingBaseccVmClusterUpdateHistoryEntrySummaryUpdateActionEnum {
		values = append(values, v)
	}
	return values
}

// GetBaseccVmClusterUpdateHistoryEntrySummaryUpdateActionEnumStringValues Enumerates the set of values in String for BaseccVmClusterUpdateHistoryEntrySummaryUpdateActionEnum
func GetBaseccVmClusterUpdateHistoryEntrySummaryUpdateActionEnumStringValues() []string {
	return []string{
		"ROLLING_APPLY",
		"PRECHECK",
		"ROLLBACK",
	}
}

// GetMappingBaseccVmClusterUpdateHistoryEntrySummaryUpdateActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBaseccVmClusterUpdateHistoryEntrySummaryUpdateActionEnum(val string) (BaseccVmClusterUpdateHistoryEntrySummaryUpdateActionEnum, bool) {
	enum, ok := mappingBaseccVmClusterUpdateHistoryEntrySummaryUpdateActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// BaseccVmClusterUpdateHistoryEntrySummaryUpdateTypeEnum Enum with underlying type: string
type BaseccVmClusterUpdateHistoryEntrySummaryUpdateTypeEnum string

// Set of constants representing the allowable values for BaseccVmClusterUpdateHistoryEntrySummaryUpdateTypeEnum
const (
	BaseccVmClusterUpdateHistoryEntrySummaryUpdateTypeUpgrade BaseccVmClusterUpdateHistoryEntrySummaryUpdateTypeEnum = "GI_UPGRADE"
	BaseccVmClusterUpdateHistoryEntrySummaryUpdateTypePatch   BaseccVmClusterUpdateHistoryEntrySummaryUpdateTypeEnum = "GI_PATCH"
)

var mappingBaseccVmClusterUpdateHistoryEntrySummaryUpdateTypeEnum = map[string]BaseccVmClusterUpdateHistoryEntrySummaryUpdateTypeEnum{
	"GI_UPGRADE": BaseccVmClusterUpdateHistoryEntrySummaryUpdateTypeUpgrade,
	"GI_PATCH":   BaseccVmClusterUpdateHistoryEntrySummaryUpdateTypePatch,
}

var mappingBaseccVmClusterUpdateHistoryEntrySummaryUpdateTypeEnumLowerCase = map[string]BaseccVmClusterUpdateHistoryEntrySummaryUpdateTypeEnum{
	"gi_upgrade": BaseccVmClusterUpdateHistoryEntrySummaryUpdateTypeUpgrade,
	"gi_patch":   BaseccVmClusterUpdateHistoryEntrySummaryUpdateTypePatch,
}

// GetBaseccVmClusterUpdateHistoryEntrySummaryUpdateTypeEnumValues Enumerates the set of values for BaseccVmClusterUpdateHistoryEntrySummaryUpdateTypeEnum
func GetBaseccVmClusterUpdateHistoryEntrySummaryUpdateTypeEnumValues() []BaseccVmClusterUpdateHistoryEntrySummaryUpdateTypeEnum {
	values := make([]BaseccVmClusterUpdateHistoryEntrySummaryUpdateTypeEnum, 0)
	for _, v := range mappingBaseccVmClusterUpdateHistoryEntrySummaryUpdateTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetBaseccVmClusterUpdateHistoryEntrySummaryUpdateTypeEnumStringValues Enumerates the set of values in String for BaseccVmClusterUpdateHistoryEntrySummaryUpdateTypeEnum
func GetBaseccVmClusterUpdateHistoryEntrySummaryUpdateTypeEnumStringValues() []string {
	return []string{
		"GI_UPGRADE",
		"GI_PATCH",
	}
}

// GetMappingBaseccVmClusterUpdateHistoryEntrySummaryUpdateTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBaseccVmClusterUpdateHistoryEntrySummaryUpdateTypeEnum(val string) (BaseccVmClusterUpdateHistoryEntrySummaryUpdateTypeEnum, bool) {
	enum, ok := mappingBaseccVmClusterUpdateHistoryEntrySummaryUpdateTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// BaseccVmClusterUpdateHistoryEntrySummaryLifecycleStateEnum Enum with underlying type: string
type BaseccVmClusterUpdateHistoryEntrySummaryLifecycleStateEnum string

// Set of constants representing the allowable values for BaseccVmClusterUpdateHistoryEntrySummaryLifecycleStateEnum
const (
	BaseccVmClusterUpdateHistoryEntrySummaryLifecycleStateInProgress BaseccVmClusterUpdateHistoryEntrySummaryLifecycleStateEnum = "IN_PROGRESS"
	BaseccVmClusterUpdateHistoryEntrySummaryLifecycleStateSucceeded  BaseccVmClusterUpdateHistoryEntrySummaryLifecycleStateEnum = "SUCCEEDED"
	BaseccVmClusterUpdateHistoryEntrySummaryLifecycleStateFailed     BaseccVmClusterUpdateHistoryEntrySummaryLifecycleStateEnum = "FAILED"
)

var mappingBaseccVmClusterUpdateHistoryEntrySummaryLifecycleStateEnum = map[string]BaseccVmClusterUpdateHistoryEntrySummaryLifecycleStateEnum{
	"IN_PROGRESS": BaseccVmClusterUpdateHistoryEntrySummaryLifecycleStateInProgress,
	"SUCCEEDED":   BaseccVmClusterUpdateHistoryEntrySummaryLifecycleStateSucceeded,
	"FAILED":      BaseccVmClusterUpdateHistoryEntrySummaryLifecycleStateFailed,
}

var mappingBaseccVmClusterUpdateHistoryEntrySummaryLifecycleStateEnumLowerCase = map[string]BaseccVmClusterUpdateHistoryEntrySummaryLifecycleStateEnum{
	"in_progress": BaseccVmClusterUpdateHistoryEntrySummaryLifecycleStateInProgress,
	"succeeded":   BaseccVmClusterUpdateHistoryEntrySummaryLifecycleStateSucceeded,
	"failed":      BaseccVmClusterUpdateHistoryEntrySummaryLifecycleStateFailed,
}

// GetBaseccVmClusterUpdateHistoryEntrySummaryLifecycleStateEnumValues Enumerates the set of values for BaseccVmClusterUpdateHistoryEntrySummaryLifecycleStateEnum
func GetBaseccVmClusterUpdateHistoryEntrySummaryLifecycleStateEnumValues() []BaseccVmClusterUpdateHistoryEntrySummaryLifecycleStateEnum {
	values := make([]BaseccVmClusterUpdateHistoryEntrySummaryLifecycleStateEnum, 0)
	for _, v := range mappingBaseccVmClusterUpdateHistoryEntrySummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetBaseccVmClusterUpdateHistoryEntrySummaryLifecycleStateEnumStringValues Enumerates the set of values in String for BaseccVmClusterUpdateHistoryEntrySummaryLifecycleStateEnum
func GetBaseccVmClusterUpdateHistoryEntrySummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"IN_PROGRESS",
		"SUCCEEDED",
		"FAILED",
	}
}

// GetMappingBaseccVmClusterUpdateHistoryEntrySummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBaseccVmClusterUpdateHistoryEntrySummaryLifecycleStateEnum(val string) (BaseccVmClusterUpdateHistoryEntrySummaryLifecycleStateEnum, bool) {
	enum, ok := mappingBaseccVmClusterUpdateHistoryEntrySummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
