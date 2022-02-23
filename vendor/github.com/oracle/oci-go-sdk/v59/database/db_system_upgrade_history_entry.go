// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v59/common"
	"strings"
)

// DbSystemUpgradeHistoryEntry The record of an OS upgrade action on a DB system.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type DbSystemUpgradeHistoryEntry struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the upgrade history entry.
	Id *string `mandatory:"true" json:"id"`

	// The operating system upgrade action.
	Action DbSystemUpgradeHistoryEntryActionEnum `mandatory:"true" json:"action"`

	// A valid Oracle Grid Infrastructure (GI) software version.
	NewGiVersion *string `mandatory:"true" json:"newGiVersion"`

	// A valid Oracle Grid Infrastructure (GI) software version.
	OldGiVersion *string `mandatory:"true" json:"oldGiVersion"`

	// The retention period, in days, for the snapshot that allows you to perform a rollback of the upgrade operation. After this number of days passes, you cannot roll back the upgrade.
	SnapshotRetentionPeriodInDays *int `mandatory:"true" json:"snapshotRetentionPeriodInDays"`

	// The current state of the action.
	LifecycleState DbSystemUpgradeHistoryEntryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time when the upgrade action started.
	TimeStarted *common.SDKTime `mandatory:"true" json:"timeStarted"`

	// A descriptive text associated with the lifecycleState.
	// Typically contains additional displayable text.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The date and time when the upgrade action completed
	TimeEnded *common.SDKTime `mandatory:"false" json:"timeEnded"`
}

func (m DbSystemUpgradeHistoryEntry) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DbSystemUpgradeHistoryEntry) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := mappingDbSystemUpgradeHistoryEntryActionEnum[string(m.Action)]; !ok && m.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", m.Action, strings.Join(GetDbSystemUpgradeHistoryEntryActionEnumStringValues(), ",")))
	}
	if _, ok := mappingDbSystemUpgradeHistoryEntryLifecycleStateEnum[string(m.LifecycleState)]; !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDbSystemUpgradeHistoryEntryLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DbSystemUpgradeHistoryEntryActionEnum Enum with underlying type: string
type DbSystemUpgradeHistoryEntryActionEnum string

// Set of constants representing the allowable values for DbSystemUpgradeHistoryEntryActionEnum
const (
	DbSystemUpgradeHistoryEntryActionPrecheck                    DbSystemUpgradeHistoryEntryActionEnum = "PRECHECK"
	DbSystemUpgradeHistoryEntryActionRollback                    DbSystemUpgradeHistoryEntryActionEnum = "ROLLBACK"
	DbSystemUpgradeHistoryEntryActionUpdateSnapshotRetentionDays DbSystemUpgradeHistoryEntryActionEnum = "UPDATE_SNAPSHOT_RETENTION_DAYS"
	DbSystemUpgradeHistoryEntryActionUpgrade                     DbSystemUpgradeHistoryEntryActionEnum = "UPGRADE"
)

var mappingDbSystemUpgradeHistoryEntryActionEnum = map[string]DbSystemUpgradeHistoryEntryActionEnum{
	"PRECHECK":                       DbSystemUpgradeHistoryEntryActionPrecheck,
	"ROLLBACK":                       DbSystemUpgradeHistoryEntryActionRollback,
	"UPDATE_SNAPSHOT_RETENTION_DAYS": DbSystemUpgradeHistoryEntryActionUpdateSnapshotRetentionDays,
	"UPGRADE":                        DbSystemUpgradeHistoryEntryActionUpgrade,
}

// GetDbSystemUpgradeHistoryEntryActionEnumValues Enumerates the set of values for DbSystemUpgradeHistoryEntryActionEnum
func GetDbSystemUpgradeHistoryEntryActionEnumValues() []DbSystemUpgradeHistoryEntryActionEnum {
	values := make([]DbSystemUpgradeHistoryEntryActionEnum, 0)
	for _, v := range mappingDbSystemUpgradeHistoryEntryActionEnum {
		values = append(values, v)
	}
	return values
}

// GetDbSystemUpgradeHistoryEntryActionEnumStringValues Enumerates the set of values in String for DbSystemUpgradeHistoryEntryActionEnum
func GetDbSystemUpgradeHistoryEntryActionEnumStringValues() []string {
	return []string{
		"PRECHECK",
		"ROLLBACK",
		"UPDATE_SNAPSHOT_RETENTION_DAYS",
		"UPGRADE",
	}
}

// DbSystemUpgradeHistoryEntryLifecycleStateEnum Enum with underlying type: string
type DbSystemUpgradeHistoryEntryLifecycleStateEnum string

// Set of constants representing the allowable values for DbSystemUpgradeHistoryEntryLifecycleStateEnum
const (
	DbSystemUpgradeHistoryEntryLifecycleStateInProgress     DbSystemUpgradeHistoryEntryLifecycleStateEnum = "IN_PROGRESS"
	DbSystemUpgradeHistoryEntryLifecycleStateSucceeded      DbSystemUpgradeHistoryEntryLifecycleStateEnum = "SUCCEEDED"
	DbSystemUpgradeHistoryEntryLifecycleStateFailed         DbSystemUpgradeHistoryEntryLifecycleStateEnum = "FAILED"
	DbSystemUpgradeHistoryEntryLifecycleStateNeedsAttention DbSystemUpgradeHistoryEntryLifecycleStateEnum = "NEEDS_ATTENTION"
)

var mappingDbSystemUpgradeHistoryEntryLifecycleStateEnum = map[string]DbSystemUpgradeHistoryEntryLifecycleStateEnum{
	"IN_PROGRESS":     DbSystemUpgradeHistoryEntryLifecycleStateInProgress,
	"SUCCEEDED":       DbSystemUpgradeHistoryEntryLifecycleStateSucceeded,
	"FAILED":          DbSystemUpgradeHistoryEntryLifecycleStateFailed,
	"NEEDS_ATTENTION": DbSystemUpgradeHistoryEntryLifecycleStateNeedsAttention,
}

// GetDbSystemUpgradeHistoryEntryLifecycleStateEnumValues Enumerates the set of values for DbSystemUpgradeHistoryEntryLifecycleStateEnum
func GetDbSystemUpgradeHistoryEntryLifecycleStateEnumValues() []DbSystemUpgradeHistoryEntryLifecycleStateEnum {
	values := make([]DbSystemUpgradeHistoryEntryLifecycleStateEnum, 0)
	for _, v := range mappingDbSystemUpgradeHistoryEntryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDbSystemUpgradeHistoryEntryLifecycleStateEnumStringValues Enumerates the set of values in String for DbSystemUpgradeHistoryEntryLifecycleStateEnum
func GetDbSystemUpgradeHistoryEntryLifecycleStateEnumStringValues() []string {
	return []string{
		"IN_PROGRESS",
		"SUCCEEDED",
		"FAILED",
		"NEEDS_ATTENTION",
	}
}
