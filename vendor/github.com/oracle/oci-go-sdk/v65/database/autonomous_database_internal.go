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

// AutonomousDatabaseInternal Details of Autonomous AI Database for reconciliation.
type AutonomousDatabaseInternal struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous AI Database.
	Id *string `mandatory:"true" json:"id"`

	// The current state of the Autonomous AI Database.
	LifecycleState AutonomousDatabaseInternalLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The database name.
	DbName *string `mandatory:"false" json:"dbName"`

	// The included compute amount (CPUs) available for FAW provisioned database.
	IncludedCompute *float32 `mandatory:"false" json:"includedCompute"`

	// The included storage value for a FAW provisioned database, in terabytes.
	IncludedDataStorageInGBs *int `mandatory:"false" json:"includedDataStorageInGBs"`
}

func (m AutonomousDatabaseInternal) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AutonomousDatabaseInternal) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAutonomousDatabaseInternalLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAutonomousDatabaseInternalLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AutonomousDatabaseInternalLifecycleStateEnum Enum with underlying type: string
type AutonomousDatabaseInternalLifecycleStateEnum string

// Set of constants representing the allowable values for AutonomousDatabaseInternalLifecycleStateEnum
const (
	AutonomousDatabaseInternalLifecycleStateProvisioning            AutonomousDatabaseInternalLifecycleStateEnum = "PROVISIONING"
	AutonomousDatabaseInternalLifecycleStateAvailable               AutonomousDatabaseInternalLifecycleStateEnum = "AVAILABLE"
	AutonomousDatabaseInternalLifecycleStateStopping                AutonomousDatabaseInternalLifecycleStateEnum = "STOPPING"
	AutonomousDatabaseInternalLifecycleStateStopped                 AutonomousDatabaseInternalLifecycleStateEnum = "STOPPED"
	AutonomousDatabaseInternalLifecycleStateStarting                AutonomousDatabaseInternalLifecycleStateEnum = "STARTING"
	AutonomousDatabaseInternalLifecycleStateTerminating             AutonomousDatabaseInternalLifecycleStateEnum = "TERMINATING"
	AutonomousDatabaseInternalLifecycleStateTerminated              AutonomousDatabaseInternalLifecycleStateEnum = "TERMINATED"
	AutonomousDatabaseInternalLifecycleStateUnavailable             AutonomousDatabaseInternalLifecycleStateEnum = "UNAVAILABLE"
	AutonomousDatabaseInternalLifecycleStateRestoreInProgress       AutonomousDatabaseInternalLifecycleStateEnum = "RESTORE_IN_PROGRESS"
	AutonomousDatabaseInternalLifecycleStateRestoreFailed           AutonomousDatabaseInternalLifecycleStateEnum = "RESTORE_FAILED"
	AutonomousDatabaseInternalLifecycleStateBackupInProgress        AutonomousDatabaseInternalLifecycleStateEnum = "BACKUP_IN_PROGRESS"
	AutonomousDatabaseInternalLifecycleStateScaleInProgress         AutonomousDatabaseInternalLifecycleStateEnum = "SCALE_IN_PROGRESS"
	AutonomousDatabaseInternalLifecycleStateAvailableNeedsAttention AutonomousDatabaseInternalLifecycleStateEnum = "AVAILABLE_NEEDS_ATTENTION"
	AutonomousDatabaseInternalLifecycleStateUpdating                AutonomousDatabaseInternalLifecycleStateEnum = "UPDATING"
	AutonomousDatabaseInternalLifecycleStateMaintenanceInProgress   AutonomousDatabaseInternalLifecycleStateEnum = "MAINTENANCE_IN_PROGRESS"
	AutonomousDatabaseInternalLifecycleStateRestarting              AutonomousDatabaseInternalLifecycleStateEnum = "RESTARTING"
	AutonomousDatabaseInternalLifecycleStateRecreating              AutonomousDatabaseInternalLifecycleStateEnum = "RECREATING"
	AutonomousDatabaseInternalLifecycleStateRoleChangeInProgress    AutonomousDatabaseInternalLifecycleStateEnum = "ROLE_CHANGE_IN_PROGRESS"
	AutonomousDatabaseInternalLifecycleStateUpgrading               AutonomousDatabaseInternalLifecycleStateEnum = "UPGRADING"
	AutonomousDatabaseInternalLifecycleStateInaccessible            AutonomousDatabaseInternalLifecycleStateEnum = "INACCESSIBLE"
	AutonomousDatabaseInternalLifecycleStateStandby                 AutonomousDatabaseInternalLifecycleStateEnum = "STANDBY"
	AutonomousDatabaseInternalLifecycleStateTransporting            AutonomousDatabaseInternalLifecycleStateEnum = "TRANSPORTING"
	AutonomousDatabaseInternalLifecycleStateArchived                AutonomousDatabaseInternalLifecycleStateEnum = "ARCHIVED"
)

var mappingAutonomousDatabaseInternalLifecycleStateEnum = map[string]AutonomousDatabaseInternalLifecycleStateEnum{
	"PROVISIONING":              AutonomousDatabaseInternalLifecycleStateProvisioning,
	"AVAILABLE":                 AutonomousDatabaseInternalLifecycleStateAvailable,
	"STOPPING":                  AutonomousDatabaseInternalLifecycleStateStopping,
	"STOPPED":                   AutonomousDatabaseInternalLifecycleStateStopped,
	"STARTING":                  AutonomousDatabaseInternalLifecycleStateStarting,
	"TERMINATING":               AutonomousDatabaseInternalLifecycleStateTerminating,
	"TERMINATED":                AutonomousDatabaseInternalLifecycleStateTerminated,
	"UNAVAILABLE":               AutonomousDatabaseInternalLifecycleStateUnavailable,
	"RESTORE_IN_PROGRESS":       AutonomousDatabaseInternalLifecycleStateRestoreInProgress,
	"RESTORE_FAILED":            AutonomousDatabaseInternalLifecycleStateRestoreFailed,
	"BACKUP_IN_PROGRESS":        AutonomousDatabaseInternalLifecycleStateBackupInProgress,
	"SCALE_IN_PROGRESS":         AutonomousDatabaseInternalLifecycleStateScaleInProgress,
	"AVAILABLE_NEEDS_ATTENTION": AutonomousDatabaseInternalLifecycleStateAvailableNeedsAttention,
	"UPDATING":                  AutonomousDatabaseInternalLifecycleStateUpdating,
	"MAINTENANCE_IN_PROGRESS":   AutonomousDatabaseInternalLifecycleStateMaintenanceInProgress,
	"RESTARTING":                AutonomousDatabaseInternalLifecycleStateRestarting,
	"RECREATING":                AutonomousDatabaseInternalLifecycleStateRecreating,
	"ROLE_CHANGE_IN_PROGRESS":   AutonomousDatabaseInternalLifecycleStateRoleChangeInProgress,
	"UPGRADING":                 AutonomousDatabaseInternalLifecycleStateUpgrading,
	"INACCESSIBLE":              AutonomousDatabaseInternalLifecycleStateInaccessible,
	"STANDBY":                   AutonomousDatabaseInternalLifecycleStateStandby,
	"TRANSPORTING":              AutonomousDatabaseInternalLifecycleStateTransporting,
	"ARCHIVED":                  AutonomousDatabaseInternalLifecycleStateArchived,
}

var mappingAutonomousDatabaseInternalLifecycleStateEnumLowerCase = map[string]AutonomousDatabaseInternalLifecycleStateEnum{
	"provisioning":              AutonomousDatabaseInternalLifecycleStateProvisioning,
	"available":                 AutonomousDatabaseInternalLifecycleStateAvailable,
	"stopping":                  AutonomousDatabaseInternalLifecycleStateStopping,
	"stopped":                   AutonomousDatabaseInternalLifecycleStateStopped,
	"starting":                  AutonomousDatabaseInternalLifecycleStateStarting,
	"terminating":               AutonomousDatabaseInternalLifecycleStateTerminating,
	"terminated":                AutonomousDatabaseInternalLifecycleStateTerminated,
	"unavailable":               AutonomousDatabaseInternalLifecycleStateUnavailable,
	"restore_in_progress":       AutonomousDatabaseInternalLifecycleStateRestoreInProgress,
	"restore_failed":            AutonomousDatabaseInternalLifecycleStateRestoreFailed,
	"backup_in_progress":        AutonomousDatabaseInternalLifecycleStateBackupInProgress,
	"scale_in_progress":         AutonomousDatabaseInternalLifecycleStateScaleInProgress,
	"available_needs_attention": AutonomousDatabaseInternalLifecycleStateAvailableNeedsAttention,
	"updating":                  AutonomousDatabaseInternalLifecycleStateUpdating,
	"maintenance_in_progress":   AutonomousDatabaseInternalLifecycleStateMaintenanceInProgress,
	"restarting":                AutonomousDatabaseInternalLifecycleStateRestarting,
	"recreating":                AutonomousDatabaseInternalLifecycleStateRecreating,
	"role_change_in_progress":   AutonomousDatabaseInternalLifecycleStateRoleChangeInProgress,
	"upgrading":                 AutonomousDatabaseInternalLifecycleStateUpgrading,
	"inaccessible":              AutonomousDatabaseInternalLifecycleStateInaccessible,
	"standby":                   AutonomousDatabaseInternalLifecycleStateStandby,
	"transporting":              AutonomousDatabaseInternalLifecycleStateTransporting,
	"archived":                  AutonomousDatabaseInternalLifecycleStateArchived,
}

// GetAutonomousDatabaseInternalLifecycleStateEnumValues Enumerates the set of values for AutonomousDatabaseInternalLifecycleStateEnum
func GetAutonomousDatabaseInternalLifecycleStateEnumValues() []AutonomousDatabaseInternalLifecycleStateEnum {
	values := make([]AutonomousDatabaseInternalLifecycleStateEnum, 0)
	for _, v := range mappingAutonomousDatabaseInternalLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDatabaseInternalLifecycleStateEnumStringValues Enumerates the set of values in String for AutonomousDatabaseInternalLifecycleStateEnum
func GetAutonomousDatabaseInternalLifecycleStateEnumStringValues() []string {
	return []string{
		"PROVISIONING",
		"AVAILABLE",
		"STOPPING",
		"STOPPED",
		"STARTING",
		"TERMINATING",
		"TERMINATED",
		"UNAVAILABLE",
		"RESTORE_IN_PROGRESS",
		"RESTORE_FAILED",
		"BACKUP_IN_PROGRESS",
		"SCALE_IN_PROGRESS",
		"AVAILABLE_NEEDS_ATTENTION",
		"UPDATING",
		"MAINTENANCE_IN_PROGRESS",
		"RESTARTING",
		"RECREATING",
		"ROLE_CHANGE_IN_PROGRESS",
		"UPGRADING",
		"INACCESSIBLE",
		"STANDBY",
		"TRANSPORTING",
		"ARCHIVED",
	}
}

// GetMappingAutonomousDatabaseInternalLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDatabaseInternalLifecycleStateEnum(val string) (AutonomousDatabaseInternalLifecycleStateEnum, bool) {
	enum, ok := mappingAutonomousDatabaseInternalLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
