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

// DbSystemOsPatchHistoryEntrySummary The summary of an OS patch action history on a DB system.
type DbSystemOsPatchHistoryEntrySummary struct {

	// The action being performed or was completed.
	Action DbSystemOsPatchHistoryEntrySummaryActionEnum `mandatory:"true" json:"action"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DB system.
	DbSystemId *string `mandatory:"true" json:"dbSystemId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the OS patch history entry.
	Id *string `mandatory:"true" json:"id"`

	// The current state of the action.
	LifecycleState DbSystemOsPatchHistoryEntrySummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	OsPatchDetails *DbSystemOsPatchDetailsCollection `mandatory:"true" json:"osPatchDetails"`

	// The date and time when the patch action started.
	TimeStarted *common.SDKTime `mandatory:"true" json:"timeStarted"`

	// A descriptive text associated with the lifecycleState.
	// Typically contains additional displayable text.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The date and time when the patch action completed
	TimeEnded *common.SDKTime `mandatory:"false" json:"timeEnded"`
}

func (m DbSystemOsPatchHistoryEntrySummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DbSystemOsPatchHistoryEntrySummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDbSystemOsPatchHistoryEntrySummaryActionEnum(string(m.Action)); !ok && m.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", m.Action, strings.Join(GetDbSystemOsPatchHistoryEntrySummaryActionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDbSystemOsPatchHistoryEntrySummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDbSystemOsPatchHistoryEntrySummaryLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DbSystemOsPatchHistoryEntrySummaryActionEnum Enum with underlying type: string
type DbSystemOsPatchHistoryEntrySummaryActionEnum string

// Set of constants representing the allowable values for DbSystemOsPatchHistoryEntrySummaryActionEnum
const (
	DbSystemOsPatchHistoryEntrySummaryActionPrecheck DbSystemOsPatchHistoryEntrySummaryActionEnum = "PRECHECK"
	DbSystemOsPatchHistoryEntrySummaryActionApply    DbSystemOsPatchHistoryEntrySummaryActionEnum = "APPLY"
)

var mappingDbSystemOsPatchHistoryEntrySummaryActionEnum = map[string]DbSystemOsPatchHistoryEntrySummaryActionEnum{
	"PRECHECK": DbSystemOsPatchHistoryEntrySummaryActionPrecheck,
	"APPLY":    DbSystemOsPatchHistoryEntrySummaryActionApply,
}

var mappingDbSystemOsPatchHistoryEntrySummaryActionEnumLowerCase = map[string]DbSystemOsPatchHistoryEntrySummaryActionEnum{
	"precheck": DbSystemOsPatchHistoryEntrySummaryActionPrecheck,
	"apply":    DbSystemOsPatchHistoryEntrySummaryActionApply,
}

// GetDbSystemOsPatchHistoryEntrySummaryActionEnumValues Enumerates the set of values for DbSystemOsPatchHistoryEntrySummaryActionEnum
func GetDbSystemOsPatchHistoryEntrySummaryActionEnumValues() []DbSystemOsPatchHistoryEntrySummaryActionEnum {
	values := make([]DbSystemOsPatchHistoryEntrySummaryActionEnum, 0)
	for _, v := range mappingDbSystemOsPatchHistoryEntrySummaryActionEnum {
		values = append(values, v)
	}
	return values
}

// GetDbSystemOsPatchHistoryEntrySummaryActionEnumStringValues Enumerates the set of values in String for DbSystemOsPatchHistoryEntrySummaryActionEnum
func GetDbSystemOsPatchHistoryEntrySummaryActionEnumStringValues() []string {
	return []string{
		"PRECHECK",
		"APPLY",
	}
}

// GetMappingDbSystemOsPatchHistoryEntrySummaryActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDbSystemOsPatchHistoryEntrySummaryActionEnum(val string) (DbSystemOsPatchHistoryEntrySummaryActionEnum, bool) {
	enum, ok := mappingDbSystemOsPatchHistoryEntrySummaryActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DbSystemOsPatchHistoryEntrySummaryLifecycleStateEnum Enum with underlying type: string
type DbSystemOsPatchHistoryEntrySummaryLifecycleStateEnum string

// Set of constants representing the allowable values for DbSystemOsPatchHistoryEntrySummaryLifecycleStateEnum
const (
	DbSystemOsPatchHistoryEntrySummaryLifecycleStateInProgress DbSystemOsPatchHistoryEntrySummaryLifecycleStateEnum = "IN_PROGRESS"
	DbSystemOsPatchHistoryEntrySummaryLifecycleStateSucceeded  DbSystemOsPatchHistoryEntrySummaryLifecycleStateEnum = "SUCCEEDED"
	DbSystemOsPatchHistoryEntrySummaryLifecycleStateFailed     DbSystemOsPatchHistoryEntrySummaryLifecycleStateEnum = "FAILED"
)

var mappingDbSystemOsPatchHistoryEntrySummaryLifecycleStateEnum = map[string]DbSystemOsPatchHistoryEntrySummaryLifecycleStateEnum{
	"IN_PROGRESS": DbSystemOsPatchHistoryEntrySummaryLifecycleStateInProgress,
	"SUCCEEDED":   DbSystemOsPatchHistoryEntrySummaryLifecycleStateSucceeded,
	"FAILED":      DbSystemOsPatchHistoryEntrySummaryLifecycleStateFailed,
}

var mappingDbSystemOsPatchHistoryEntrySummaryLifecycleStateEnumLowerCase = map[string]DbSystemOsPatchHistoryEntrySummaryLifecycleStateEnum{
	"in_progress": DbSystemOsPatchHistoryEntrySummaryLifecycleStateInProgress,
	"succeeded":   DbSystemOsPatchHistoryEntrySummaryLifecycleStateSucceeded,
	"failed":      DbSystemOsPatchHistoryEntrySummaryLifecycleStateFailed,
}

// GetDbSystemOsPatchHistoryEntrySummaryLifecycleStateEnumValues Enumerates the set of values for DbSystemOsPatchHistoryEntrySummaryLifecycleStateEnum
func GetDbSystemOsPatchHistoryEntrySummaryLifecycleStateEnumValues() []DbSystemOsPatchHistoryEntrySummaryLifecycleStateEnum {
	values := make([]DbSystemOsPatchHistoryEntrySummaryLifecycleStateEnum, 0)
	for _, v := range mappingDbSystemOsPatchHistoryEntrySummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDbSystemOsPatchHistoryEntrySummaryLifecycleStateEnumStringValues Enumerates the set of values in String for DbSystemOsPatchHistoryEntrySummaryLifecycleStateEnum
func GetDbSystemOsPatchHistoryEntrySummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"IN_PROGRESS",
		"SUCCEEDED",
		"FAILED",
	}
}

// GetMappingDbSystemOsPatchHistoryEntrySummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDbSystemOsPatchHistoryEntrySummaryLifecycleStateEnum(val string) (DbSystemOsPatchHistoryEntrySummaryLifecycleStateEnum, bool) {
	enum, ok := mappingDbSystemOsPatchHistoryEntrySummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
