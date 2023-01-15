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

// PatchHistoryEntrySummary The record of a patch action on a specified target.
type PatchHistoryEntrySummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the patch history entry.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the patch.
	PatchId *string `mandatory:"true" json:"patchId"`

	// The current state of the action.
	LifecycleState PatchHistoryEntrySummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time when the patch action started.
	TimeStarted *common.SDKTime `mandatory:"true" json:"timeStarted"`

	// The action being performed or was completed.
	Action PatchHistoryEntrySummaryActionEnum `mandatory:"false" json:"action,omitempty"`

	// A descriptive text associated with the lifecycleState.
	// Typically contains additional displayable text.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The date and time when the patch action completed
	TimeEnded *common.SDKTime `mandatory:"false" json:"timeEnded"`
}

func (m PatchHistoryEntrySummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PatchHistoryEntrySummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPatchHistoryEntrySummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetPatchHistoryEntrySummaryLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingPatchHistoryEntrySummaryActionEnum(string(m.Action)); !ok && m.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", m.Action, strings.Join(GetPatchHistoryEntrySummaryActionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PatchHistoryEntrySummaryActionEnum Enum with underlying type: string
type PatchHistoryEntrySummaryActionEnum string

// Set of constants representing the allowable values for PatchHistoryEntrySummaryActionEnum
const (
	PatchHistoryEntrySummaryActionApply    PatchHistoryEntrySummaryActionEnum = "APPLY"
	PatchHistoryEntrySummaryActionPrecheck PatchHistoryEntrySummaryActionEnum = "PRECHECK"
)

var mappingPatchHistoryEntrySummaryActionEnum = map[string]PatchHistoryEntrySummaryActionEnum{
	"APPLY":    PatchHistoryEntrySummaryActionApply,
	"PRECHECK": PatchHistoryEntrySummaryActionPrecheck,
}

var mappingPatchHistoryEntrySummaryActionEnumLowerCase = map[string]PatchHistoryEntrySummaryActionEnum{
	"apply":    PatchHistoryEntrySummaryActionApply,
	"precheck": PatchHistoryEntrySummaryActionPrecheck,
}

// GetPatchHistoryEntrySummaryActionEnumValues Enumerates the set of values for PatchHistoryEntrySummaryActionEnum
func GetPatchHistoryEntrySummaryActionEnumValues() []PatchHistoryEntrySummaryActionEnum {
	values := make([]PatchHistoryEntrySummaryActionEnum, 0)
	for _, v := range mappingPatchHistoryEntrySummaryActionEnum {
		values = append(values, v)
	}
	return values
}

// GetPatchHistoryEntrySummaryActionEnumStringValues Enumerates the set of values in String for PatchHistoryEntrySummaryActionEnum
func GetPatchHistoryEntrySummaryActionEnumStringValues() []string {
	return []string{
		"APPLY",
		"PRECHECK",
	}
}

// GetMappingPatchHistoryEntrySummaryActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPatchHistoryEntrySummaryActionEnum(val string) (PatchHistoryEntrySummaryActionEnum, bool) {
	enum, ok := mappingPatchHistoryEntrySummaryActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// PatchHistoryEntrySummaryLifecycleStateEnum Enum with underlying type: string
type PatchHistoryEntrySummaryLifecycleStateEnum string

// Set of constants representing the allowable values for PatchHistoryEntrySummaryLifecycleStateEnum
const (
	PatchHistoryEntrySummaryLifecycleStateInProgress PatchHistoryEntrySummaryLifecycleStateEnum = "IN_PROGRESS"
	PatchHistoryEntrySummaryLifecycleStateSucceeded  PatchHistoryEntrySummaryLifecycleStateEnum = "SUCCEEDED"
	PatchHistoryEntrySummaryLifecycleStateFailed     PatchHistoryEntrySummaryLifecycleStateEnum = "FAILED"
)

var mappingPatchHistoryEntrySummaryLifecycleStateEnum = map[string]PatchHistoryEntrySummaryLifecycleStateEnum{
	"IN_PROGRESS": PatchHistoryEntrySummaryLifecycleStateInProgress,
	"SUCCEEDED":   PatchHistoryEntrySummaryLifecycleStateSucceeded,
	"FAILED":      PatchHistoryEntrySummaryLifecycleStateFailed,
}

var mappingPatchHistoryEntrySummaryLifecycleStateEnumLowerCase = map[string]PatchHistoryEntrySummaryLifecycleStateEnum{
	"in_progress": PatchHistoryEntrySummaryLifecycleStateInProgress,
	"succeeded":   PatchHistoryEntrySummaryLifecycleStateSucceeded,
	"failed":      PatchHistoryEntrySummaryLifecycleStateFailed,
}

// GetPatchHistoryEntrySummaryLifecycleStateEnumValues Enumerates the set of values for PatchHistoryEntrySummaryLifecycleStateEnum
func GetPatchHistoryEntrySummaryLifecycleStateEnumValues() []PatchHistoryEntrySummaryLifecycleStateEnum {
	values := make([]PatchHistoryEntrySummaryLifecycleStateEnum, 0)
	for _, v := range mappingPatchHistoryEntrySummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetPatchHistoryEntrySummaryLifecycleStateEnumStringValues Enumerates the set of values in String for PatchHistoryEntrySummaryLifecycleStateEnum
func GetPatchHistoryEntrySummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"IN_PROGRESS",
		"SUCCEEDED",
		"FAILED",
	}
}

// GetMappingPatchHistoryEntrySummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPatchHistoryEntrySummaryLifecycleStateEnum(val string) (PatchHistoryEntrySummaryLifecycleStateEnum, bool) {
	enum, ok := mappingPatchHistoryEntrySummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
