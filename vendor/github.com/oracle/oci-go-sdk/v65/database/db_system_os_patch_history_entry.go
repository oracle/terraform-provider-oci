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

// DbSystemOsPatchHistoryEntry The record of an OS patch action on a DB system.
type DbSystemOsPatchHistoryEntry struct {

	// The action being performed or was completed.
	Action DbSystemOsPatchHistoryEntryActionEnum `mandatory:"true" json:"action"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DB system.
	DbSystemId *string `mandatory:"true" json:"dbSystemId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the OS patch history entry.
	Id *string `mandatory:"true" json:"id"`

	// The current state of the action.
	LifecycleState DbSystemOsPatchHistoryEntryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	OsPatchDetails *DbSystemOsPatchDetailsCollection `mandatory:"true" json:"osPatchDetails"`

	// The date and time when the patch action started.
	TimeStarted *common.SDKTime `mandatory:"true" json:"timeStarted"`

	// A descriptive text associated with the lifecycleState.
	// Typically contains additional displayable text.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The date and time when the patch action completed
	TimeEnded *common.SDKTime `mandatory:"false" json:"timeEnded"`
}

func (m DbSystemOsPatchHistoryEntry) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DbSystemOsPatchHistoryEntry) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDbSystemOsPatchHistoryEntryActionEnum(string(m.Action)); !ok && m.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", m.Action, strings.Join(GetDbSystemOsPatchHistoryEntryActionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDbSystemOsPatchHistoryEntryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDbSystemOsPatchHistoryEntryLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DbSystemOsPatchHistoryEntryActionEnum Enum with underlying type: string
type DbSystemOsPatchHistoryEntryActionEnum string

// Set of constants representing the allowable values for DbSystemOsPatchHistoryEntryActionEnum
const (
	DbSystemOsPatchHistoryEntryActionPrecheck DbSystemOsPatchHistoryEntryActionEnum = "PRECHECK"
	DbSystemOsPatchHistoryEntryActionApply    DbSystemOsPatchHistoryEntryActionEnum = "APPLY"
)

var mappingDbSystemOsPatchHistoryEntryActionEnum = map[string]DbSystemOsPatchHistoryEntryActionEnum{
	"PRECHECK": DbSystemOsPatchHistoryEntryActionPrecheck,
	"APPLY":    DbSystemOsPatchHistoryEntryActionApply,
}

var mappingDbSystemOsPatchHistoryEntryActionEnumLowerCase = map[string]DbSystemOsPatchHistoryEntryActionEnum{
	"precheck": DbSystemOsPatchHistoryEntryActionPrecheck,
	"apply":    DbSystemOsPatchHistoryEntryActionApply,
}

// GetDbSystemOsPatchHistoryEntryActionEnumValues Enumerates the set of values for DbSystemOsPatchHistoryEntryActionEnum
func GetDbSystemOsPatchHistoryEntryActionEnumValues() []DbSystemOsPatchHistoryEntryActionEnum {
	values := make([]DbSystemOsPatchHistoryEntryActionEnum, 0)
	for _, v := range mappingDbSystemOsPatchHistoryEntryActionEnum {
		values = append(values, v)
	}
	return values
}

// GetDbSystemOsPatchHistoryEntryActionEnumStringValues Enumerates the set of values in String for DbSystemOsPatchHistoryEntryActionEnum
func GetDbSystemOsPatchHistoryEntryActionEnumStringValues() []string {
	return []string{
		"PRECHECK",
		"APPLY",
	}
}

// GetMappingDbSystemOsPatchHistoryEntryActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDbSystemOsPatchHistoryEntryActionEnum(val string) (DbSystemOsPatchHistoryEntryActionEnum, bool) {
	enum, ok := mappingDbSystemOsPatchHistoryEntryActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DbSystemOsPatchHistoryEntryLifecycleStateEnum Enum with underlying type: string
type DbSystemOsPatchHistoryEntryLifecycleStateEnum string

// Set of constants representing the allowable values for DbSystemOsPatchHistoryEntryLifecycleStateEnum
const (
	DbSystemOsPatchHistoryEntryLifecycleStateInProgress DbSystemOsPatchHistoryEntryLifecycleStateEnum = "IN_PROGRESS"
	DbSystemOsPatchHistoryEntryLifecycleStateSucceeded  DbSystemOsPatchHistoryEntryLifecycleStateEnum = "SUCCEEDED"
	DbSystemOsPatchHistoryEntryLifecycleStateFailed     DbSystemOsPatchHistoryEntryLifecycleStateEnum = "FAILED"
)

var mappingDbSystemOsPatchHistoryEntryLifecycleStateEnum = map[string]DbSystemOsPatchHistoryEntryLifecycleStateEnum{
	"IN_PROGRESS": DbSystemOsPatchHistoryEntryLifecycleStateInProgress,
	"SUCCEEDED":   DbSystemOsPatchHistoryEntryLifecycleStateSucceeded,
	"FAILED":      DbSystemOsPatchHistoryEntryLifecycleStateFailed,
}

var mappingDbSystemOsPatchHistoryEntryLifecycleStateEnumLowerCase = map[string]DbSystemOsPatchHistoryEntryLifecycleStateEnum{
	"in_progress": DbSystemOsPatchHistoryEntryLifecycleStateInProgress,
	"succeeded":   DbSystemOsPatchHistoryEntryLifecycleStateSucceeded,
	"failed":      DbSystemOsPatchHistoryEntryLifecycleStateFailed,
}

// GetDbSystemOsPatchHistoryEntryLifecycleStateEnumValues Enumerates the set of values for DbSystemOsPatchHistoryEntryLifecycleStateEnum
func GetDbSystemOsPatchHistoryEntryLifecycleStateEnumValues() []DbSystemOsPatchHistoryEntryLifecycleStateEnum {
	values := make([]DbSystemOsPatchHistoryEntryLifecycleStateEnum, 0)
	for _, v := range mappingDbSystemOsPatchHistoryEntryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDbSystemOsPatchHistoryEntryLifecycleStateEnumStringValues Enumerates the set of values in String for DbSystemOsPatchHistoryEntryLifecycleStateEnum
func GetDbSystemOsPatchHistoryEntryLifecycleStateEnumStringValues() []string {
	return []string{
		"IN_PROGRESS",
		"SUCCEEDED",
		"FAILED",
	}
}

// GetMappingDbSystemOsPatchHistoryEntryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDbSystemOsPatchHistoryEntryLifecycleStateEnum(val string) (DbSystemOsPatchHistoryEntryLifecycleStateEnum, bool) {
	enum, ok := mappingDbSystemOsPatchHistoryEntryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
