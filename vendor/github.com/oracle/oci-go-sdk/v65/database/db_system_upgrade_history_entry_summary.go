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

// DbSystemUpgradeHistoryEntrySummary The summary for the record of an OS upgrade action on a DB system.
type DbSystemUpgradeHistoryEntrySummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the upgrade history entry.
	Id *string `mandatory:"true" json:"id"`

	// The operating system upgrade action.
	Action DbSystemUpgradeHistoryEntrySummaryActionEnum `mandatory:"true" json:"action"`

	// The retention period, in days, for the snapshot that allows you to perform a rollback of the upgrade operation. After this number of days passes, you cannot roll back the upgrade.
	SnapshotRetentionPeriodInDays *int `mandatory:"true" json:"snapshotRetentionPeriodInDays"`

	// The current state of the action.
	LifecycleState DbSystemUpgradeHistoryEntrySummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time when the upgrade action started.
	TimeStarted *common.SDKTime `mandatory:"true" json:"timeStarted"`

	// A valid Oracle Grid Infrastructure (GI) software version.
	NewGiVersion *string `mandatory:"false" json:"newGiVersion"`

	// A valid Oracle Grid Infrastructure (GI) software version.
	OldGiVersion *string `mandatory:"false" json:"oldGiVersion"`

	// A valid Oracle Software (OS) version eg. Oracle Linux Server release 8
	OldOsVersion *string `mandatory:"false" json:"oldOsVersion"`

	// A valid Oracle Software (OS) version eg. Oracle Linux Server release 8
	NewOsVersion *string `mandatory:"false" json:"newOsVersion"`

	// A descriptive text associated with the lifecycleState.
	// Typically contains additional displayable text.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The date and time when the upgrade action completed
	TimeEnded *common.SDKTime `mandatory:"false" json:"timeEnded"`
}

func (m DbSystemUpgradeHistoryEntrySummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DbSystemUpgradeHistoryEntrySummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDbSystemUpgradeHistoryEntrySummaryActionEnum(string(m.Action)); !ok && m.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", m.Action, strings.Join(GetDbSystemUpgradeHistoryEntrySummaryActionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDbSystemUpgradeHistoryEntrySummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDbSystemUpgradeHistoryEntrySummaryLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DbSystemUpgradeHistoryEntrySummaryActionEnum Enum with underlying type: string
type DbSystemUpgradeHistoryEntrySummaryActionEnum string

// Set of constants representing the allowable values for DbSystemUpgradeHistoryEntrySummaryActionEnum
const (
	DbSystemUpgradeHistoryEntrySummaryActionPrecheck                    DbSystemUpgradeHistoryEntrySummaryActionEnum = "PRECHECK"
	DbSystemUpgradeHistoryEntrySummaryActionRollback                    DbSystemUpgradeHistoryEntrySummaryActionEnum = "ROLLBACK"
	DbSystemUpgradeHistoryEntrySummaryActionUpdateSnapshotRetentionDays DbSystemUpgradeHistoryEntrySummaryActionEnum = "UPDATE_SNAPSHOT_RETENTION_DAYS"
	DbSystemUpgradeHistoryEntrySummaryActionUpgrade                     DbSystemUpgradeHistoryEntrySummaryActionEnum = "UPGRADE"
)

var mappingDbSystemUpgradeHistoryEntrySummaryActionEnum = map[string]DbSystemUpgradeHistoryEntrySummaryActionEnum{
	"PRECHECK":                       DbSystemUpgradeHistoryEntrySummaryActionPrecheck,
	"ROLLBACK":                       DbSystemUpgradeHistoryEntrySummaryActionRollback,
	"UPDATE_SNAPSHOT_RETENTION_DAYS": DbSystemUpgradeHistoryEntrySummaryActionUpdateSnapshotRetentionDays,
	"UPGRADE":                        DbSystemUpgradeHistoryEntrySummaryActionUpgrade,
}

var mappingDbSystemUpgradeHistoryEntrySummaryActionEnumLowerCase = map[string]DbSystemUpgradeHistoryEntrySummaryActionEnum{
	"precheck":                       DbSystemUpgradeHistoryEntrySummaryActionPrecheck,
	"rollback":                       DbSystemUpgradeHistoryEntrySummaryActionRollback,
	"update_snapshot_retention_days": DbSystemUpgradeHistoryEntrySummaryActionUpdateSnapshotRetentionDays,
	"upgrade":                        DbSystemUpgradeHistoryEntrySummaryActionUpgrade,
}

// GetDbSystemUpgradeHistoryEntrySummaryActionEnumValues Enumerates the set of values for DbSystemUpgradeHistoryEntrySummaryActionEnum
func GetDbSystemUpgradeHistoryEntrySummaryActionEnumValues() []DbSystemUpgradeHistoryEntrySummaryActionEnum {
	values := make([]DbSystemUpgradeHistoryEntrySummaryActionEnum, 0)
	for _, v := range mappingDbSystemUpgradeHistoryEntrySummaryActionEnum {
		values = append(values, v)
	}
	return values
}

// GetDbSystemUpgradeHistoryEntrySummaryActionEnumStringValues Enumerates the set of values in String for DbSystemUpgradeHistoryEntrySummaryActionEnum
func GetDbSystemUpgradeHistoryEntrySummaryActionEnumStringValues() []string {
	return []string{
		"PRECHECK",
		"ROLLBACK",
		"UPDATE_SNAPSHOT_RETENTION_DAYS",
		"UPGRADE",
	}
}

// GetMappingDbSystemUpgradeHistoryEntrySummaryActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDbSystemUpgradeHistoryEntrySummaryActionEnum(val string) (DbSystemUpgradeHistoryEntrySummaryActionEnum, bool) {
	enum, ok := mappingDbSystemUpgradeHistoryEntrySummaryActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DbSystemUpgradeHistoryEntrySummaryLifecycleStateEnum Enum with underlying type: string
type DbSystemUpgradeHistoryEntrySummaryLifecycleStateEnum string

// Set of constants representing the allowable values for DbSystemUpgradeHistoryEntrySummaryLifecycleStateEnum
const (
	DbSystemUpgradeHistoryEntrySummaryLifecycleStateInProgress     DbSystemUpgradeHistoryEntrySummaryLifecycleStateEnum = "IN_PROGRESS"
	DbSystemUpgradeHistoryEntrySummaryLifecycleStateSucceeded      DbSystemUpgradeHistoryEntrySummaryLifecycleStateEnum = "SUCCEEDED"
	DbSystemUpgradeHistoryEntrySummaryLifecycleStateFailed         DbSystemUpgradeHistoryEntrySummaryLifecycleStateEnum = "FAILED"
	DbSystemUpgradeHistoryEntrySummaryLifecycleStateNeedsAttention DbSystemUpgradeHistoryEntrySummaryLifecycleStateEnum = "NEEDS_ATTENTION"
)

var mappingDbSystemUpgradeHistoryEntrySummaryLifecycleStateEnum = map[string]DbSystemUpgradeHistoryEntrySummaryLifecycleStateEnum{
	"IN_PROGRESS":     DbSystemUpgradeHistoryEntrySummaryLifecycleStateInProgress,
	"SUCCEEDED":       DbSystemUpgradeHistoryEntrySummaryLifecycleStateSucceeded,
	"FAILED":          DbSystemUpgradeHistoryEntrySummaryLifecycleStateFailed,
	"NEEDS_ATTENTION": DbSystemUpgradeHistoryEntrySummaryLifecycleStateNeedsAttention,
}

var mappingDbSystemUpgradeHistoryEntrySummaryLifecycleStateEnumLowerCase = map[string]DbSystemUpgradeHistoryEntrySummaryLifecycleStateEnum{
	"in_progress":     DbSystemUpgradeHistoryEntrySummaryLifecycleStateInProgress,
	"succeeded":       DbSystemUpgradeHistoryEntrySummaryLifecycleStateSucceeded,
	"failed":          DbSystemUpgradeHistoryEntrySummaryLifecycleStateFailed,
	"needs_attention": DbSystemUpgradeHistoryEntrySummaryLifecycleStateNeedsAttention,
}

// GetDbSystemUpgradeHistoryEntrySummaryLifecycleStateEnumValues Enumerates the set of values for DbSystemUpgradeHistoryEntrySummaryLifecycleStateEnum
func GetDbSystemUpgradeHistoryEntrySummaryLifecycleStateEnumValues() []DbSystemUpgradeHistoryEntrySummaryLifecycleStateEnum {
	values := make([]DbSystemUpgradeHistoryEntrySummaryLifecycleStateEnum, 0)
	for _, v := range mappingDbSystemUpgradeHistoryEntrySummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDbSystemUpgradeHistoryEntrySummaryLifecycleStateEnumStringValues Enumerates the set of values in String for DbSystemUpgradeHistoryEntrySummaryLifecycleStateEnum
func GetDbSystemUpgradeHistoryEntrySummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"IN_PROGRESS",
		"SUCCEEDED",
		"FAILED",
		"NEEDS_ATTENTION",
	}
}

// GetMappingDbSystemUpgradeHistoryEntrySummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDbSystemUpgradeHistoryEntrySummaryLifecycleStateEnum(val string) (DbSystemUpgradeHistoryEntrySummaryLifecycleStateEnum, bool) {
	enum, ok := mappingDbSystemUpgradeHistoryEntrySummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
