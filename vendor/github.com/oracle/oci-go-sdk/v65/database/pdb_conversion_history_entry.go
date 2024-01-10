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

// PdbConversionHistoryEntry Details of operations performed to convert a non-container database to pluggable database.
type PdbConversionHistoryEntry struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the database conversion history.
	Id *string `mandatory:"true" json:"id"`

	// The operations used to convert a non-container database to a pluggable database.
	// - Use `PRECHECK` to run a pre-check operation on non-container database prior to converting it into a pluggable database.
	// - Use `CONVERT` to convert a non-container database into a pluggable database.
	// - Use `SYNC` if the non-container database was manually converted into a pluggable database using the dbcli command-line utility. Databases may need to be converted manually if the CONVERT action fails when converting a non-container database using the API.
	// - Use `SYNC_ROLLBACK` if the conversion of a non-container database into a pluggable database was manually rolled back using the dbcli command line utility. Conversions may need to be manually rolled back if the CONVERT action fails when converting a non-container database using the API.
	Action PdbConversionHistoryEntryActionEnum `mandatory:"true" json:"action"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the database.
	SourceDatabaseId *string `mandatory:"true" json:"sourceDatabaseId"`

	// The database name. The name must begin with an alphabetic character and can contain a maximum of 8 alphanumeric characters. Special characters are not permitted. The database name must be unique in the tenancy.
	CdbName *string `mandatory:"true" json:"cdbName"`

	// Status of an operation performed during the conversion of a non-container database to a pluggable database.
	LifecycleState PdbConversionHistoryEntryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time when the database conversion operation started.
	TimeStarted *common.SDKTime `mandatory:"true" json:"timeStarted"`

	// The target container database of the pluggable database created by the database conversion operation. Currently, the database conversion operation only supports creating the pluggable database in a new container database.
	//  - Use `NEW_DATABASE` to specify that the pluggable database be created within a new container database in the same database home.
	Target PdbConversionHistoryEntryTargetEnum `mandatory:"false" json:"target,omitempty"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the database.
	TargetDatabaseId *string `mandatory:"false" json:"targetDatabaseId"`

	// Additional information about the current lifecycle state for the conversion operation.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The date and time when the database conversion operation ended.
	TimeEnded *common.SDKTime `mandatory:"false" json:"timeEnded"`

	// Additional container database parameter.
	AdditionalCdbParams *string `mandatory:"false" json:"additionalCdbParams"`
}

func (m PdbConversionHistoryEntry) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PdbConversionHistoryEntry) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPdbConversionHistoryEntryActionEnum(string(m.Action)); !ok && m.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", m.Action, strings.Join(GetPdbConversionHistoryEntryActionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPdbConversionHistoryEntryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetPdbConversionHistoryEntryLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingPdbConversionHistoryEntryTargetEnum(string(m.Target)); !ok && m.Target != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Target: %s. Supported values are: %s.", m.Target, strings.Join(GetPdbConversionHistoryEntryTargetEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PdbConversionHistoryEntryActionEnum Enum with underlying type: string
type PdbConversionHistoryEntryActionEnum string

// Set of constants representing the allowable values for PdbConversionHistoryEntryActionEnum
const (
	PdbConversionHistoryEntryActionPrecheck     PdbConversionHistoryEntryActionEnum = "PRECHECK"
	PdbConversionHistoryEntryActionConvert      PdbConversionHistoryEntryActionEnum = "CONVERT"
	PdbConversionHistoryEntryActionSync         PdbConversionHistoryEntryActionEnum = "SYNC"
	PdbConversionHistoryEntryActionSyncRollback PdbConversionHistoryEntryActionEnum = "SYNC_ROLLBACK"
)

var mappingPdbConversionHistoryEntryActionEnum = map[string]PdbConversionHistoryEntryActionEnum{
	"PRECHECK":      PdbConversionHistoryEntryActionPrecheck,
	"CONVERT":       PdbConversionHistoryEntryActionConvert,
	"SYNC":          PdbConversionHistoryEntryActionSync,
	"SYNC_ROLLBACK": PdbConversionHistoryEntryActionSyncRollback,
}

var mappingPdbConversionHistoryEntryActionEnumLowerCase = map[string]PdbConversionHistoryEntryActionEnum{
	"precheck":      PdbConversionHistoryEntryActionPrecheck,
	"convert":       PdbConversionHistoryEntryActionConvert,
	"sync":          PdbConversionHistoryEntryActionSync,
	"sync_rollback": PdbConversionHistoryEntryActionSyncRollback,
}

// GetPdbConversionHistoryEntryActionEnumValues Enumerates the set of values for PdbConversionHistoryEntryActionEnum
func GetPdbConversionHistoryEntryActionEnumValues() []PdbConversionHistoryEntryActionEnum {
	values := make([]PdbConversionHistoryEntryActionEnum, 0)
	for _, v := range mappingPdbConversionHistoryEntryActionEnum {
		values = append(values, v)
	}
	return values
}

// GetPdbConversionHistoryEntryActionEnumStringValues Enumerates the set of values in String for PdbConversionHistoryEntryActionEnum
func GetPdbConversionHistoryEntryActionEnumStringValues() []string {
	return []string{
		"PRECHECK",
		"CONVERT",
		"SYNC",
		"SYNC_ROLLBACK",
	}
}

// GetMappingPdbConversionHistoryEntryActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPdbConversionHistoryEntryActionEnum(val string) (PdbConversionHistoryEntryActionEnum, bool) {
	enum, ok := mappingPdbConversionHistoryEntryActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// PdbConversionHistoryEntryTargetEnum Enum with underlying type: string
type PdbConversionHistoryEntryTargetEnum string

// Set of constants representing the allowable values for PdbConversionHistoryEntryTargetEnum
const (
	PdbConversionHistoryEntryTargetNewDatabase PdbConversionHistoryEntryTargetEnum = "NEW_DATABASE"
)

var mappingPdbConversionHistoryEntryTargetEnum = map[string]PdbConversionHistoryEntryTargetEnum{
	"NEW_DATABASE": PdbConversionHistoryEntryTargetNewDatabase,
}

var mappingPdbConversionHistoryEntryTargetEnumLowerCase = map[string]PdbConversionHistoryEntryTargetEnum{
	"new_database": PdbConversionHistoryEntryTargetNewDatabase,
}

// GetPdbConversionHistoryEntryTargetEnumValues Enumerates the set of values for PdbConversionHistoryEntryTargetEnum
func GetPdbConversionHistoryEntryTargetEnumValues() []PdbConversionHistoryEntryTargetEnum {
	values := make([]PdbConversionHistoryEntryTargetEnum, 0)
	for _, v := range mappingPdbConversionHistoryEntryTargetEnum {
		values = append(values, v)
	}
	return values
}

// GetPdbConversionHistoryEntryTargetEnumStringValues Enumerates the set of values in String for PdbConversionHistoryEntryTargetEnum
func GetPdbConversionHistoryEntryTargetEnumStringValues() []string {
	return []string{
		"NEW_DATABASE",
	}
}

// GetMappingPdbConversionHistoryEntryTargetEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPdbConversionHistoryEntryTargetEnum(val string) (PdbConversionHistoryEntryTargetEnum, bool) {
	enum, ok := mappingPdbConversionHistoryEntryTargetEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// PdbConversionHistoryEntryLifecycleStateEnum Enum with underlying type: string
type PdbConversionHistoryEntryLifecycleStateEnum string

// Set of constants representing the allowable values for PdbConversionHistoryEntryLifecycleStateEnum
const (
	PdbConversionHistoryEntryLifecycleStateSucceeded  PdbConversionHistoryEntryLifecycleStateEnum = "SUCCEEDED"
	PdbConversionHistoryEntryLifecycleStateFailed     PdbConversionHistoryEntryLifecycleStateEnum = "FAILED"
	PdbConversionHistoryEntryLifecycleStateInProgress PdbConversionHistoryEntryLifecycleStateEnum = "IN_PROGRESS"
)

var mappingPdbConversionHistoryEntryLifecycleStateEnum = map[string]PdbConversionHistoryEntryLifecycleStateEnum{
	"SUCCEEDED":   PdbConversionHistoryEntryLifecycleStateSucceeded,
	"FAILED":      PdbConversionHistoryEntryLifecycleStateFailed,
	"IN_PROGRESS": PdbConversionHistoryEntryLifecycleStateInProgress,
}

var mappingPdbConversionHistoryEntryLifecycleStateEnumLowerCase = map[string]PdbConversionHistoryEntryLifecycleStateEnum{
	"succeeded":   PdbConversionHistoryEntryLifecycleStateSucceeded,
	"failed":      PdbConversionHistoryEntryLifecycleStateFailed,
	"in_progress": PdbConversionHistoryEntryLifecycleStateInProgress,
}

// GetPdbConversionHistoryEntryLifecycleStateEnumValues Enumerates the set of values for PdbConversionHistoryEntryLifecycleStateEnum
func GetPdbConversionHistoryEntryLifecycleStateEnumValues() []PdbConversionHistoryEntryLifecycleStateEnum {
	values := make([]PdbConversionHistoryEntryLifecycleStateEnum, 0)
	for _, v := range mappingPdbConversionHistoryEntryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetPdbConversionHistoryEntryLifecycleStateEnumStringValues Enumerates the set of values in String for PdbConversionHistoryEntryLifecycleStateEnum
func GetPdbConversionHistoryEntryLifecycleStateEnumStringValues() []string {
	return []string{
		"SUCCEEDED",
		"FAILED",
		"IN_PROGRESS",
	}
}

// GetMappingPdbConversionHistoryEntryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPdbConversionHistoryEntryLifecycleStateEnum(val string) (PdbConversionHistoryEntryLifecycleStateEnum, bool) {
	enum, ok := mappingPdbConversionHistoryEntryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
