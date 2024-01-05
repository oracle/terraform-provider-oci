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

// PdbConversionHistoryEntrySummary Details of operations performed to convert a non-container database to pluggable database.
type PdbConversionHistoryEntrySummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the database conversion history.
	Id *string `mandatory:"true" json:"id"`

	// The operations used to convert a non-container database to a pluggable database.
	// - Use `PRECHECK` to run a pre-check operation on non-container database prior to converting it into a pluggable database.
	// - Use `CONVERT` to convert a non-container database into a pluggable database.
	// - Use `SYNC` if the non-container database was manually converted into a pluggable database using the dbcli command-line utility. Databases may need to be converted manually if the CONVERT action fails when converting a non-container database using the API.
	// - Use `SYNC_ROLLBACK` if the conversion of a non-container database into a pluggable database was manually rolled back using the dbcli command line utility. Conversions may need to be manually rolled back if the CONVERT action fails when converting a non-container database using the API.
	Action PdbConversionHistoryEntrySummaryActionEnum `mandatory:"true" json:"action"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the database.
	SourceDatabaseId *string `mandatory:"true" json:"sourceDatabaseId"`

	// The database name. The name must begin with an alphabetic character and can contain a maximum of 8 alphanumeric characters. Special characters are not permitted. The database name must be unique in the tenancy.
	CdbName *string `mandatory:"true" json:"cdbName"`

	// Status of an operation performed during the conversion of a non-container database to a pluggable database.
	LifecycleState PdbConversionHistoryEntrySummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time when the database conversion operation started.
	TimeStarted *common.SDKTime `mandatory:"true" json:"timeStarted"`

	// The target container database of the pluggable database created by the database conversion operation. Currently, the database conversion operation only supports creating the pluggable database in a new container database.
	//  - Use `NEW_DATABASE` to specify that the pluggable database be created within a new container database in the same database home.
	Target PdbConversionHistoryEntrySummaryTargetEnum `mandatory:"false" json:"target,omitempty"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the database.
	TargetDatabaseId *string `mandatory:"false" json:"targetDatabaseId"`

	// Additional information about the current lifecycle state for the conversion operation.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The date and time when the database conversion operation ended.
	TimeEnded *common.SDKTime `mandatory:"false" json:"timeEnded"`

	// Additional container database parameter.
	AdditionalCdbParams *string `mandatory:"false" json:"additionalCdbParams"`
}

func (m PdbConversionHistoryEntrySummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PdbConversionHistoryEntrySummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPdbConversionHistoryEntrySummaryActionEnum(string(m.Action)); !ok && m.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", m.Action, strings.Join(GetPdbConversionHistoryEntrySummaryActionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPdbConversionHistoryEntrySummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetPdbConversionHistoryEntrySummaryLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingPdbConversionHistoryEntrySummaryTargetEnum(string(m.Target)); !ok && m.Target != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Target: %s. Supported values are: %s.", m.Target, strings.Join(GetPdbConversionHistoryEntrySummaryTargetEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PdbConversionHistoryEntrySummaryActionEnum Enum with underlying type: string
type PdbConversionHistoryEntrySummaryActionEnum string

// Set of constants representing the allowable values for PdbConversionHistoryEntrySummaryActionEnum
const (
	PdbConversionHistoryEntrySummaryActionPrecheck     PdbConversionHistoryEntrySummaryActionEnum = "PRECHECK"
	PdbConversionHistoryEntrySummaryActionConvert      PdbConversionHistoryEntrySummaryActionEnum = "CONVERT"
	PdbConversionHistoryEntrySummaryActionSync         PdbConversionHistoryEntrySummaryActionEnum = "SYNC"
	PdbConversionHistoryEntrySummaryActionSyncRollback PdbConversionHistoryEntrySummaryActionEnum = "SYNC_ROLLBACK"
)

var mappingPdbConversionHistoryEntrySummaryActionEnum = map[string]PdbConversionHistoryEntrySummaryActionEnum{
	"PRECHECK":      PdbConversionHistoryEntrySummaryActionPrecheck,
	"CONVERT":       PdbConversionHistoryEntrySummaryActionConvert,
	"SYNC":          PdbConversionHistoryEntrySummaryActionSync,
	"SYNC_ROLLBACK": PdbConversionHistoryEntrySummaryActionSyncRollback,
}

var mappingPdbConversionHistoryEntrySummaryActionEnumLowerCase = map[string]PdbConversionHistoryEntrySummaryActionEnum{
	"precheck":      PdbConversionHistoryEntrySummaryActionPrecheck,
	"convert":       PdbConversionHistoryEntrySummaryActionConvert,
	"sync":          PdbConversionHistoryEntrySummaryActionSync,
	"sync_rollback": PdbConversionHistoryEntrySummaryActionSyncRollback,
}

// GetPdbConversionHistoryEntrySummaryActionEnumValues Enumerates the set of values for PdbConversionHistoryEntrySummaryActionEnum
func GetPdbConversionHistoryEntrySummaryActionEnumValues() []PdbConversionHistoryEntrySummaryActionEnum {
	values := make([]PdbConversionHistoryEntrySummaryActionEnum, 0)
	for _, v := range mappingPdbConversionHistoryEntrySummaryActionEnum {
		values = append(values, v)
	}
	return values
}

// GetPdbConversionHistoryEntrySummaryActionEnumStringValues Enumerates the set of values in String for PdbConversionHistoryEntrySummaryActionEnum
func GetPdbConversionHistoryEntrySummaryActionEnumStringValues() []string {
	return []string{
		"PRECHECK",
		"CONVERT",
		"SYNC",
		"SYNC_ROLLBACK",
	}
}

// GetMappingPdbConversionHistoryEntrySummaryActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPdbConversionHistoryEntrySummaryActionEnum(val string) (PdbConversionHistoryEntrySummaryActionEnum, bool) {
	enum, ok := mappingPdbConversionHistoryEntrySummaryActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// PdbConversionHistoryEntrySummaryTargetEnum Enum with underlying type: string
type PdbConversionHistoryEntrySummaryTargetEnum string

// Set of constants representing the allowable values for PdbConversionHistoryEntrySummaryTargetEnum
const (
	PdbConversionHistoryEntrySummaryTargetNewDatabase PdbConversionHistoryEntrySummaryTargetEnum = "NEW_DATABASE"
)

var mappingPdbConversionHistoryEntrySummaryTargetEnum = map[string]PdbConversionHistoryEntrySummaryTargetEnum{
	"NEW_DATABASE": PdbConversionHistoryEntrySummaryTargetNewDatabase,
}

var mappingPdbConversionHistoryEntrySummaryTargetEnumLowerCase = map[string]PdbConversionHistoryEntrySummaryTargetEnum{
	"new_database": PdbConversionHistoryEntrySummaryTargetNewDatabase,
}

// GetPdbConversionHistoryEntrySummaryTargetEnumValues Enumerates the set of values for PdbConversionHistoryEntrySummaryTargetEnum
func GetPdbConversionHistoryEntrySummaryTargetEnumValues() []PdbConversionHistoryEntrySummaryTargetEnum {
	values := make([]PdbConversionHistoryEntrySummaryTargetEnum, 0)
	for _, v := range mappingPdbConversionHistoryEntrySummaryTargetEnum {
		values = append(values, v)
	}
	return values
}

// GetPdbConversionHistoryEntrySummaryTargetEnumStringValues Enumerates the set of values in String for PdbConversionHistoryEntrySummaryTargetEnum
func GetPdbConversionHistoryEntrySummaryTargetEnumStringValues() []string {
	return []string{
		"NEW_DATABASE",
	}
}

// GetMappingPdbConversionHistoryEntrySummaryTargetEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPdbConversionHistoryEntrySummaryTargetEnum(val string) (PdbConversionHistoryEntrySummaryTargetEnum, bool) {
	enum, ok := mappingPdbConversionHistoryEntrySummaryTargetEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// PdbConversionHistoryEntrySummaryLifecycleStateEnum Enum with underlying type: string
type PdbConversionHistoryEntrySummaryLifecycleStateEnum string

// Set of constants representing the allowable values for PdbConversionHistoryEntrySummaryLifecycleStateEnum
const (
	PdbConversionHistoryEntrySummaryLifecycleStateSucceeded  PdbConversionHistoryEntrySummaryLifecycleStateEnum = "SUCCEEDED"
	PdbConversionHistoryEntrySummaryLifecycleStateFailed     PdbConversionHistoryEntrySummaryLifecycleStateEnum = "FAILED"
	PdbConversionHistoryEntrySummaryLifecycleStateInProgress PdbConversionHistoryEntrySummaryLifecycleStateEnum = "IN_PROGRESS"
)

var mappingPdbConversionHistoryEntrySummaryLifecycleStateEnum = map[string]PdbConversionHistoryEntrySummaryLifecycleStateEnum{
	"SUCCEEDED":   PdbConversionHistoryEntrySummaryLifecycleStateSucceeded,
	"FAILED":      PdbConversionHistoryEntrySummaryLifecycleStateFailed,
	"IN_PROGRESS": PdbConversionHistoryEntrySummaryLifecycleStateInProgress,
}

var mappingPdbConversionHistoryEntrySummaryLifecycleStateEnumLowerCase = map[string]PdbConversionHistoryEntrySummaryLifecycleStateEnum{
	"succeeded":   PdbConversionHistoryEntrySummaryLifecycleStateSucceeded,
	"failed":      PdbConversionHistoryEntrySummaryLifecycleStateFailed,
	"in_progress": PdbConversionHistoryEntrySummaryLifecycleStateInProgress,
}

// GetPdbConversionHistoryEntrySummaryLifecycleStateEnumValues Enumerates the set of values for PdbConversionHistoryEntrySummaryLifecycleStateEnum
func GetPdbConversionHistoryEntrySummaryLifecycleStateEnumValues() []PdbConversionHistoryEntrySummaryLifecycleStateEnum {
	values := make([]PdbConversionHistoryEntrySummaryLifecycleStateEnum, 0)
	for _, v := range mappingPdbConversionHistoryEntrySummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetPdbConversionHistoryEntrySummaryLifecycleStateEnumStringValues Enumerates the set of values in String for PdbConversionHistoryEntrySummaryLifecycleStateEnum
func GetPdbConversionHistoryEntrySummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"SUCCEEDED",
		"FAILED",
		"IN_PROGRESS",
	}
}

// GetMappingPdbConversionHistoryEntrySummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPdbConversionHistoryEntrySummaryLifecycleStateEnum(val string) (PdbConversionHistoryEntrySummaryLifecycleStateEnum, bool) {
	enum, ok := mappingPdbConversionHistoryEntrySummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
