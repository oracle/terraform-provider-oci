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

// PatchSummary A Patch for a DB system or DB Home.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access,
// see Getting Started with Policies (https://docs.cloud.oracle.com/Content/Identity/Concepts/policygetstarted.htm).
type PatchSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the patch.
	Id *string `mandatory:"true" json:"id"`

	// The text describing this patch package.
	Description *string `mandatory:"true" json:"description"`

	// The date and time that the patch was released.
	TimeReleased *common.SDKTime `mandatory:"true" json:"timeReleased"`

	// The version of this patch package.
	Version *string `mandatory:"true" json:"version"`

	// Action that is currently being performed or was completed last.
	LastAction PatchSummaryLastActionEnum `mandatory:"false" json:"lastAction,omitempty"`

	// Actions that can possibly be performed using this patch.
	AvailableActions []PatchSummaryAvailableActionsEnum `mandatory:"false" json:"availableActions,omitempty"`

	// A descriptive text associated with the lifecycleState.
	// Typically can contain additional displayable text.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The current state of the patch as a result of lastAction.
	LifecycleState PatchSummaryLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m PatchSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PatchSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingPatchSummaryLastActionEnum(string(m.LastAction)); !ok && m.LastAction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LastAction: %s. Supported values are: %s.", m.LastAction, strings.Join(GetPatchSummaryLastActionEnumStringValues(), ",")))
	}
	for _, val := range m.AvailableActions {
		if _, ok := GetMappingPatchSummaryAvailableActionsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AvailableActions: %s. Supported values are: %s.", val, strings.Join(GetPatchSummaryAvailableActionsEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingPatchSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetPatchSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PatchSummaryLastActionEnum Enum with underlying type: string
type PatchSummaryLastActionEnum string

// Set of constants representing the allowable values for PatchSummaryLastActionEnum
const (
	PatchSummaryLastActionApply    PatchSummaryLastActionEnum = "APPLY"
	PatchSummaryLastActionPrecheck PatchSummaryLastActionEnum = "PRECHECK"
)

var mappingPatchSummaryLastActionEnum = map[string]PatchSummaryLastActionEnum{
	"APPLY":    PatchSummaryLastActionApply,
	"PRECHECK": PatchSummaryLastActionPrecheck,
}

var mappingPatchSummaryLastActionEnumLowerCase = map[string]PatchSummaryLastActionEnum{
	"apply":    PatchSummaryLastActionApply,
	"precheck": PatchSummaryLastActionPrecheck,
}

// GetPatchSummaryLastActionEnumValues Enumerates the set of values for PatchSummaryLastActionEnum
func GetPatchSummaryLastActionEnumValues() []PatchSummaryLastActionEnum {
	values := make([]PatchSummaryLastActionEnum, 0)
	for _, v := range mappingPatchSummaryLastActionEnum {
		values = append(values, v)
	}
	return values
}

// GetPatchSummaryLastActionEnumStringValues Enumerates the set of values in String for PatchSummaryLastActionEnum
func GetPatchSummaryLastActionEnumStringValues() []string {
	return []string{
		"APPLY",
		"PRECHECK",
	}
}

// GetMappingPatchSummaryLastActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPatchSummaryLastActionEnum(val string) (PatchSummaryLastActionEnum, bool) {
	enum, ok := mappingPatchSummaryLastActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// PatchSummaryAvailableActionsEnum Enum with underlying type: string
type PatchSummaryAvailableActionsEnum string

// Set of constants representing the allowable values for PatchSummaryAvailableActionsEnum
const (
	PatchSummaryAvailableActionsApply    PatchSummaryAvailableActionsEnum = "APPLY"
	PatchSummaryAvailableActionsPrecheck PatchSummaryAvailableActionsEnum = "PRECHECK"
)

var mappingPatchSummaryAvailableActionsEnum = map[string]PatchSummaryAvailableActionsEnum{
	"APPLY":    PatchSummaryAvailableActionsApply,
	"PRECHECK": PatchSummaryAvailableActionsPrecheck,
}

var mappingPatchSummaryAvailableActionsEnumLowerCase = map[string]PatchSummaryAvailableActionsEnum{
	"apply":    PatchSummaryAvailableActionsApply,
	"precheck": PatchSummaryAvailableActionsPrecheck,
}

// GetPatchSummaryAvailableActionsEnumValues Enumerates the set of values for PatchSummaryAvailableActionsEnum
func GetPatchSummaryAvailableActionsEnumValues() []PatchSummaryAvailableActionsEnum {
	values := make([]PatchSummaryAvailableActionsEnum, 0)
	for _, v := range mappingPatchSummaryAvailableActionsEnum {
		values = append(values, v)
	}
	return values
}

// GetPatchSummaryAvailableActionsEnumStringValues Enumerates the set of values in String for PatchSummaryAvailableActionsEnum
func GetPatchSummaryAvailableActionsEnumStringValues() []string {
	return []string{
		"APPLY",
		"PRECHECK",
	}
}

// GetMappingPatchSummaryAvailableActionsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPatchSummaryAvailableActionsEnum(val string) (PatchSummaryAvailableActionsEnum, bool) {
	enum, ok := mappingPatchSummaryAvailableActionsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// PatchSummaryLifecycleStateEnum Enum with underlying type: string
type PatchSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for PatchSummaryLifecycleStateEnum
const (
	PatchSummaryLifecycleStateAvailable  PatchSummaryLifecycleStateEnum = "AVAILABLE"
	PatchSummaryLifecycleStateSuccess    PatchSummaryLifecycleStateEnum = "SUCCESS"
	PatchSummaryLifecycleStateInProgress PatchSummaryLifecycleStateEnum = "IN_PROGRESS"
	PatchSummaryLifecycleStateFailed     PatchSummaryLifecycleStateEnum = "FAILED"
)

var mappingPatchSummaryLifecycleStateEnum = map[string]PatchSummaryLifecycleStateEnum{
	"AVAILABLE":   PatchSummaryLifecycleStateAvailable,
	"SUCCESS":     PatchSummaryLifecycleStateSuccess,
	"IN_PROGRESS": PatchSummaryLifecycleStateInProgress,
	"FAILED":      PatchSummaryLifecycleStateFailed,
}

var mappingPatchSummaryLifecycleStateEnumLowerCase = map[string]PatchSummaryLifecycleStateEnum{
	"available":   PatchSummaryLifecycleStateAvailable,
	"success":     PatchSummaryLifecycleStateSuccess,
	"in_progress": PatchSummaryLifecycleStateInProgress,
	"failed":      PatchSummaryLifecycleStateFailed,
}

// GetPatchSummaryLifecycleStateEnumValues Enumerates the set of values for PatchSummaryLifecycleStateEnum
func GetPatchSummaryLifecycleStateEnumValues() []PatchSummaryLifecycleStateEnum {
	values := make([]PatchSummaryLifecycleStateEnum, 0)
	for _, v := range mappingPatchSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetPatchSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for PatchSummaryLifecycleStateEnum
func GetPatchSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"AVAILABLE",
		"SUCCESS",
		"IN_PROGRESS",
		"FAILED",
	}
}

// GetMappingPatchSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPatchSummaryLifecycleStateEnum(val string) (PatchSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingPatchSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
