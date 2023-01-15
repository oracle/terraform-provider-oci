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

// Patch The representation of Patch
type Patch struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the patch.
	Id *string `mandatory:"true" json:"id"`

	// The text describing this patch package.
	Description *string `mandatory:"true" json:"description"`

	// The date and time that the patch was released.
	TimeReleased *common.SDKTime `mandatory:"true" json:"timeReleased"`

	// The version of this patch package.
	Version *string `mandatory:"true" json:"version"`

	// Action that is currently being performed or was completed last.
	LastAction PatchLastActionEnum `mandatory:"false" json:"lastAction,omitempty"`

	// Actions that can possibly be performed using this patch.
	AvailableActions []PatchAvailableActionsEnum `mandatory:"false" json:"availableActions,omitempty"`

	// A descriptive text associated with the lifecycleState.
	// Typically can contain additional displayable text.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The current state of the patch as a result of lastAction.
	LifecycleState PatchLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m Patch) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Patch) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingPatchLastActionEnum(string(m.LastAction)); !ok && m.LastAction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LastAction: %s. Supported values are: %s.", m.LastAction, strings.Join(GetPatchLastActionEnumStringValues(), ",")))
	}
	for _, val := range m.AvailableActions {
		if _, ok := GetMappingPatchAvailableActionsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AvailableActions: %s. Supported values are: %s.", val, strings.Join(GetPatchAvailableActionsEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingPatchLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetPatchLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PatchLastActionEnum Enum with underlying type: string
type PatchLastActionEnum string

// Set of constants representing the allowable values for PatchLastActionEnum
const (
	PatchLastActionApply    PatchLastActionEnum = "APPLY"
	PatchLastActionPrecheck PatchLastActionEnum = "PRECHECK"
)

var mappingPatchLastActionEnum = map[string]PatchLastActionEnum{
	"APPLY":    PatchLastActionApply,
	"PRECHECK": PatchLastActionPrecheck,
}

var mappingPatchLastActionEnumLowerCase = map[string]PatchLastActionEnum{
	"apply":    PatchLastActionApply,
	"precheck": PatchLastActionPrecheck,
}

// GetPatchLastActionEnumValues Enumerates the set of values for PatchLastActionEnum
func GetPatchLastActionEnumValues() []PatchLastActionEnum {
	values := make([]PatchLastActionEnum, 0)
	for _, v := range mappingPatchLastActionEnum {
		values = append(values, v)
	}
	return values
}

// GetPatchLastActionEnumStringValues Enumerates the set of values in String for PatchLastActionEnum
func GetPatchLastActionEnumStringValues() []string {
	return []string{
		"APPLY",
		"PRECHECK",
	}
}

// GetMappingPatchLastActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPatchLastActionEnum(val string) (PatchLastActionEnum, bool) {
	enum, ok := mappingPatchLastActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// PatchAvailableActionsEnum Enum with underlying type: string
type PatchAvailableActionsEnum string

// Set of constants representing the allowable values for PatchAvailableActionsEnum
const (
	PatchAvailableActionsApply    PatchAvailableActionsEnum = "APPLY"
	PatchAvailableActionsPrecheck PatchAvailableActionsEnum = "PRECHECK"
)

var mappingPatchAvailableActionsEnum = map[string]PatchAvailableActionsEnum{
	"APPLY":    PatchAvailableActionsApply,
	"PRECHECK": PatchAvailableActionsPrecheck,
}

var mappingPatchAvailableActionsEnumLowerCase = map[string]PatchAvailableActionsEnum{
	"apply":    PatchAvailableActionsApply,
	"precheck": PatchAvailableActionsPrecheck,
}

// GetPatchAvailableActionsEnumValues Enumerates the set of values for PatchAvailableActionsEnum
func GetPatchAvailableActionsEnumValues() []PatchAvailableActionsEnum {
	values := make([]PatchAvailableActionsEnum, 0)
	for _, v := range mappingPatchAvailableActionsEnum {
		values = append(values, v)
	}
	return values
}

// GetPatchAvailableActionsEnumStringValues Enumerates the set of values in String for PatchAvailableActionsEnum
func GetPatchAvailableActionsEnumStringValues() []string {
	return []string{
		"APPLY",
		"PRECHECK",
	}
}

// GetMappingPatchAvailableActionsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPatchAvailableActionsEnum(val string) (PatchAvailableActionsEnum, bool) {
	enum, ok := mappingPatchAvailableActionsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// PatchLifecycleStateEnum Enum with underlying type: string
type PatchLifecycleStateEnum string

// Set of constants representing the allowable values for PatchLifecycleStateEnum
const (
	PatchLifecycleStateAvailable  PatchLifecycleStateEnum = "AVAILABLE"
	PatchLifecycleStateSuccess    PatchLifecycleStateEnum = "SUCCESS"
	PatchLifecycleStateInProgress PatchLifecycleStateEnum = "IN_PROGRESS"
	PatchLifecycleStateFailed     PatchLifecycleStateEnum = "FAILED"
)

var mappingPatchLifecycleStateEnum = map[string]PatchLifecycleStateEnum{
	"AVAILABLE":   PatchLifecycleStateAvailable,
	"SUCCESS":     PatchLifecycleStateSuccess,
	"IN_PROGRESS": PatchLifecycleStateInProgress,
	"FAILED":      PatchLifecycleStateFailed,
}

var mappingPatchLifecycleStateEnumLowerCase = map[string]PatchLifecycleStateEnum{
	"available":   PatchLifecycleStateAvailable,
	"success":     PatchLifecycleStateSuccess,
	"in_progress": PatchLifecycleStateInProgress,
	"failed":      PatchLifecycleStateFailed,
}

// GetPatchLifecycleStateEnumValues Enumerates the set of values for PatchLifecycleStateEnum
func GetPatchLifecycleStateEnumValues() []PatchLifecycleStateEnum {
	values := make([]PatchLifecycleStateEnum, 0)
	for _, v := range mappingPatchLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetPatchLifecycleStateEnumStringValues Enumerates the set of values in String for PatchLifecycleStateEnum
func GetPatchLifecycleStateEnumStringValues() []string {
	return []string{
		"AVAILABLE",
		"SUCCESS",
		"IN_PROGRESS",
		"FAILED",
	}
}

// GetMappingPatchLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPatchLifecycleStateEnum(val string) (PatchLifecycleStateEnum, bool) {
	enum, ok := mappingPatchLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
