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

// PatchHistoryEntry The representation of PatchHistoryEntry
type PatchHistoryEntry struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the patch history entry.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the patch.
	PatchId *string `mandatory:"true" json:"patchId"`

	// The current state of the action.
	LifecycleState PatchHistoryEntryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time when the patch action started.
	TimeStarted *common.SDKTime `mandatory:"true" json:"timeStarted"`

	// The action being performed or was completed.
	Action PatchHistoryEntryActionEnum `mandatory:"false" json:"action,omitempty"`

	// A descriptive text associated with the lifecycleState.
	// Typically contains additional displayable text.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The date and time when the patch action completed
	TimeEnded *common.SDKTime `mandatory:"false" json:"timeEnded"`
}

func (m PatchHistoryEntry) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PatchHistoryEntry) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPatchHistoryEntryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetPatchHistoryEntryLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingPatchHistoryEntryActionEnum(string(m.Action)); !ok && m.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", m.Action, strings.Join(GetPatchHistoryEntryActionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PatchHistoryEntryActionEnum Enum with underlying type: string
type PatchHistoryEntryActionEnum string

// Set of constants representing the allowable values for PatchHistoryEntryActionEnum
const (
	PatchHistoryEntryActionApply    PatchHistoryEntryActionEnum = "APPLY"
	PatchHistoryEntryActionPrecheck PatchHistoryEntryActionEnum = "PRECHECK"
)

var mappingPatchHistoryEntryActionEnum = map[string]PatchHistoryEntryActionEnum{
	"APPLY":    PatchHistoryEntryActionApply,
	"PRECHECK": PatchHistoryEntryActionPrecheck,
}

var mappingPatchHistoryEntryActionEnumLowerCase = map[string]PatchHistoryEntryActionEnum{
	"apply":    PatchHistoryEntryActionApply,
	"precheck": PatchHistoryEntryActionPrecheck,
}

// GetPatchHistoryEntryActionEnumValues Enumerates the set of values for PatchHistoryEntryActionEnum
func GetPatchHistoryEntryActionEnumValues() []PatchHistoryEntryActionEnum {
	values := make([]PatchHistoryEntryActionEnum, 0)
	for _, v := range mappingPatchHistoryEntryActionEnum {
		values = append(values, v)
	}
	return values
}

// GetPatchHistoryEntryActionEnumStringValues Enumerates the set of values in String for PatchHistoryEntryActionEnum
func GetPatchHistoryEntryActionEnumStringValues() []string {
	return []string{
		"APPLY",
		"PRECHECK",
	}
}

// GetMappingPatchHistoryEntryActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPatchHistoryEntryActionEnum(val string) (PatchHistoryEntryActionEnum, bool) {
	enum, ok := mappingPatchHistoryEntryActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// PatchHistoryEntryLifecycleStateEnum Enum with underlying type: string
type PatchHistoryEntryLifecycleStateEnum string

// Set of constants representing the allowable values for PatchHistoryEntryLifecycleStateEnum
const (
	PatchHistoryEntryLifecycleStateInProgress PatchHistoryEntryLifecycleStateEnum = "IN_PROGRESS"
	PatchHistoryEntryLifecycleStateSucceeded  PatchHistoryEntryLifecycleStateEnum = "SUCCEEDED"
	PatchHistoryEntryLifecycleStateFailed     PatchHistoryEntryLifecycleStateEnum = "FAILED"
)

var mappingPatchHistoryEntryLifecycleStateEnum = map[string]PatchHistoryEntryLifecycleStateEnum{
	"IN_PROGRESS": PatchHistoryEntryLifecycleStateInProgress,
	"SUCCEEDED":   PatchHistoryEntryLifecycleStateSucceeded,
	"FAILED":      PatchHistoryEntryLifecycleStateFailed,
}

var mappingPatchHistoryEntryLifecycleStateEnumLowerCase = map[string]PatchHistoryEntryLifecycleStateEnum{
	"in_progress": PatchHistoryEntryLifecycleStateInProgress,
	"succeeded":   PatchHistoryEntryLifecycleStateSucceeded,
	"failed":      PatchHistoryEntryLifecycleStateFailed,
}

// GetPatchHistoryEntryLifecycleStateEnumValues Enumerates the set of values for PatchHistoryEntryLifecycleStateEnum
func GetPatchHistoryEntryLifecycleStateEnumValues() []PatchHistoryEntryLifecycleStateEnum {
	values := make([]PatchHistoryEntryLifecycleStateEnum, 0)
	for _, v := range mappingPatchHistoryEntryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetPatchHistoryEntryLifecycleStateEnumStringValues Enumerates the set of values in String for PatchHistoryEntryLifecycleStateEnum
func GetPatchHistoryEntryLifecycleStateEnumStringValues() []string {
	return []string{
		"IN_PROGRESS",
		"SUCCEEDED",
		"FAILED",
	}
}

// GetMappingPatchHistoryEntryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPatchHistoryEntryLifecycleStateEnum(val string) (PatchHistoryEntryLifecycleStateEnum, bool) {
	enum, ok := mappingPatchHistoryEntryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
