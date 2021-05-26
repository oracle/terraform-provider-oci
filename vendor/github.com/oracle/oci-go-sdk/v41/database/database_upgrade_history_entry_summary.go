// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/v41/common"
)

// DatabaseUpgradeHistoryEntrySummary The Database service supports the upgrade history of databases.
// To use any of the API operations, you must be authorized in an IAM policy. If you are not authorized, talk to an administrator. If you are an administrator who needs to write policies to give users access, see Getting Started with Policies (https://docs.cloud.oracle.com/Content/Identity/Concepts/policygetstarted.htm).
// For information about access control and compartments, see
// Overview of the Identity Service (https://docs.cloud.oracle.com/Content/Identity/Concepts/overview.htm).
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type DatabaseUpgradeHistoryEntrySummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the database upgrade history.
	Id *string `mandatory:"true" json:"id"`

	// The database upgrade action.
	Action DatabaseUpgradeHistoryEntrySummaryActionEnum `mandatory:"true" json:"action"`

	// Status of database upgrade history SUCCEEDED|IN_PROGRESS|FAILED.
	LifecycleState DatabaseUpgradeHistoryEntrySummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time when the database upgrade started.
	TimeStarted *common.SDKTime `mandatory:"true" json:"timeStarted"`

	// The source of the Oracle Database software to be used for the upgrade.
	//  - Use `DB_HOME` to specify an existing Database Home to upgrade the database. The database is moved to the target Database Home and makes use of the Oracle Database software version of the target Database Home.
	//  - Use `DB_VERSION` to specify a generally-available Oracle Database software version to upgrade the database.
	//  - Use `DB_SOFTWARE_IMAGE` to specify a database software image (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databasesoftwareimage.htm) to upgrade the database.
	Source DatabaseUpgradeHistoryEntrySummarySourceEnum `mandatory:"false" json:"source,omitempty"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// A valid Oracle Database version. To get a list of supported versions, use the ListDbVersions operation.
	TargetDBVersion *string `mandatory:"false" json:"targetDBVersion"`

	// the database software image used for upgrading database.
	TargetDatabaseSoftwareImageId *string `mandatory:"false" json:"targetDatabaseSoftwareImageId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Database Home.
	TargetDbHomeId *string `mandatory:"false" json:"targetDbHomeId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Database Home.
	SourceDbHomeId *string `mandatory:"false" json:"sourceDbHomeId"`

	// The date and time when the database upgrade ended.
	TimeEnded *common.SDKTime `mandatory:"false" json:"timeEnded"`

	// Additional upgrade options supported by DBUA(Database Upgrade Assistant).
	// Example: "-upgradeTimezone false -keepEvents"
	Options *string `mandatory:"false" json:"options"`
}

func (m DatabaseUpgradeHistoryEntrySummary) String() string {
	return common.PointerString(m)
}

// DatabaseUpgradeHistoryEntrySummaryActionEnum Enum with underlying type: string
type DatabaseUpgradeHistoryEntrySummaryActionEnum string

// Set of constants representing the allowable values for DatabaseUpgradeHistoryEntrySummaryActionEnum
const (
	DatabaseUpgradeHistoryEntrySummaryActionPrecheck DatabaseUpgradeHistoryEntrySummaryActionEnum = "PRECHECK"
	DatabaseUpgradeHistoryEntrySummaryActionUpgrade  DatabaseUpgradeHistoryEntrySummaryActionEnum = "UPGRADE"
	DatabaseUpgradeHistoryEntrySummaryActionRollback DatabaseUpgradeHistoryEntrySummaryActionEnum = "ROLLBACK"
)

var mappingDatabaseUpgradeHistoryEntrySummaryAction = map[string]DatabaseUpgradeHistoryEntrySummaryActionEnum{
	"PRECHECK": DatabaseUpgradeHistoryEntrySummaryActionPrecheck,
	"UPGRADE":  DatabaseUpgradeHistoryEntrySummaryActionUpgrade,
	"ROLLBACK": DatabaseUpgradeHistoryEntrySummaryActionRollback,
}

// GetDatabaseUpgradeHistoryEntrySummaryActionEnumValues Enumerates the set of values for DatabaseUpgradeHistoryEntrySummaryActionEnum
func GetDatabaseUpgradeHistoryEntrySummaryActionEnumValues() []DatabaseUpgradeHistoryEntrySummaryActionEnum {
	values := make([]DatabaseUpgradeHistoryEntrySummaryActionEnum, 0)
	for _, v := range mappingDatabaseUpgradeHistoryEntrySummaryAction {
		values = append(values, v)
	}
	return values
}

// DatabaseUpgradeHistoryEntrySummarySourceEnum Enum with underlying type: string
type DatabaseUpgradeHistoryEntrySummarySourceEnum string

// Set of constants representing the allowable values for DatabaseUpgradeHistoryEntrySummarySourceEnum
const (
	DatabaseUpgradeHistoryEntrySummarySourceHome          DatabaseUpgradeHistoryEntrySummarySourceEnum = "DB_HOME"
	DatabaseUpgradeHistoryEntrySummarySourceVersion       DatabaseUpgradeHistoryEntrySummarySourceEnum = "DB_VERSION"
	DatabaseUpgradeHistoryEntrySummarySourceSoftwareImage DatabaseUpgradeHistoryEntrySummarySourceEnum = "DB_SOFTWARE_IMAGE"
)

var mappingDatabaseUpgradeHistoryEntrySummarySource = map[string]DatabaseUpgradeHistoryEntrySummarySourceEnum{
	"DB_HOME":           DatabaseUpgradeHistoryEntrySummarySourceHome,
	"DB_VERSION":        DatabaseUpgradeHistoryEntrySummarySourceVersion,
	"DB_SOFTWARE_IMAGE": DatabaseUpgradeHistoryEntrySummarySourceSoftwareImage,
}

// GetDatabaseUpgradeHistoryEntrySummarySourceEnumValues Enumerates the set of values for DatabaseUpgradeHistoryEntrySummarySourceEnum
func GetDatabaseUpgradeHistoryEntrySummarySourceEnumValues() []DatabaseUpgradeHistoryEntrySummarySourceEnum {
	values := make([]DatabaseUpgradeHistoryEntrySummarySourceEnum, 0)
	for _, v := range mappingDatabaseUpgradeHistoryEntrySummarySource {
		values = append(values, v)
	}
	return values
}

// DatabaseUpgradeHistoryEntrySummaryLifecycleStateEnum Enum with underlying type: string
type DatabaseUpgradeHistoryEntrySummaryLifecycleStateEnum string

// Set of constants representing the allowable values for DatabaseUpgradeHistoryEntrySummaryLifecycleStateEnum
const (
	DatabaseUpgradeHistoryEntrySummaryLifecycleStateSucceeded  DatabaseUpgradeHistoryEntrySummaryLifecycleStateEnum = "SUCCEEDED"
	DatabaseUpgradeHistoryEntrySummaryLifecycleStateFailed     DatabaseUpgradeHistoryEntrySummaryLifecycleStateEnum = "FAILED"
	DatabaseUpgradeHistoryEntrySummaryLifecycleStateInProgress DatabaseUpgradeHistoryEntrySummaryLifecycleStateEnum = "IN_PROGRESS"
)

var mappingDatabaseUpgradeHistoryEntrySummaryLifecycleState = map[string]DatabaseUpgradeHistoryEntrySummaryLifecycleStateEnum{
	"SUCCEEDED":   DatabaseUpgradeHistoryEntrySummaryLifecycleStateSucceeded,
	"FAILED":      DatabaseUpgradeHistoryEntrySummaryLifecycleStateFailed,
	"IN_PROGRESS": DatabaseUpgradeHistoryEntrySummaryLifecycleStateInProgress,
}

// GetDatabaseUpgradeHistoryEntrySummaryLifecycleStateEnumValues Enumerates the set of values for DatabaseUpgradeHistoryEntrySummaryLifecycleStateEnum
func GetDatabaseUpgradeHistoryEntrySummaryLifecycleStateEnumValues() []DatabaseUpgradeHistoryEntrySummaryLifecycleStateEnum {
	values := make([]DatabaseUpgradeHistoryEntrySummaryLifecycleStateEnum, 0)
	for _, v := range mappingDatabaseUpgradeHistoryEntrySummaryLifecycleState {
		values = append(values, v)
	}
	return values
}
