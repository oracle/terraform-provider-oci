// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// AutonomousDatabaseInternalSummary Details of Autonomous database for reconciliation.
type AutonomousDatabaseInternalSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Autonomous Database.
	Id *string `mandatory:"true" json:"id"`

	// The current state of the Autonomous Database.
	LifecycleState AutonomousDatabaseInternalSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The database name.
	DbName *string `mandatory:"false" json:"dbName"`

	// The included compute amount (CPUs) available for FAW provisioned database.
	IncludedCompute *float32 `mandatory:"false" json:"includedCompute"`

	// The included storage value for a FAW provisioned database, in terabytes.
	IncludedDataStorageInGBs *int `mandatory:"false" json:"includedDataStorageInGBs"`
}

func (m AutonomousDatabaseInternalSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AutonomousDatabaseInternalSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAutonomousDatabaseInternalSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAutonomousDatabaseInternalSummaryLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AutonomousDatabaseInternalSummaryLifecycleStateEnum Enum with underlying type: string
type AutonomousDatabaseInternalSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for AutonomousDatabaseInternalSummaryLifecycleStateEnum
const (
	AutonomousDatabaseInternalSummaryLifecycleStateProvisioning            AutonomousDatabaseInternalSummaryLifecycleStateEnum = "PROVISIONING"
	AutonomousDatabaseInternalSummaryLifecycleStateAvailable               AutonomousDatabaseInternalSummaryLifecycleStateEnum = "AVAILABLE"
	AutonomousDatabaseInternalSummaryLifecycleStateStopping                AutonomousDatabaseInternalSummaryLifecycleStateEnum = "STOPPING"
	AutonomousDatabaseInternalSummaryLifecycleStateStopped                 AutonomousDatabaseInternalSummaryLifecycleStateEnum = "STOPPED"
	AutonomousDatabaseInternalSummaryLifecycleStateStarting                AutonomousDatabaseInternalSummaryLifecycleStateEnum = "STARTING"
	AutonomousDatabaseInternalSummaryLifecycleStateTerminating             AutonomousDatabaseInternalSummaryLifecycleStateEnum = "TERMINATING"
	AutonomousDatabaseInternalSummaryLifecycleStateTerminated              AutonomousDatabaseInternalSummaryLifecycleStateEnum = "TERMINATED"
	AutonomousDatabaseInternalSummaryLifecycleStateUnavailable             AutonomousDatabaseInternalSummaryLifecycleStateEnum = "UNAVAILABLE"
	AutonomousDatabaseInternalSummaryLifecycleStateRestoreInProgress       AutonomousDatabaseInternalSummaryLifecycleStateEnum = "RESTORE_IN_PROGRESS"
	AutonomousDatabaseInternalSummaryLifecycleStateRestoreFailed           AutonomousDatabaseInternalSummaryLifecycleStateEnum = "RESTORE_FAILED"
	AutonomousDatabaseInternalSummaryLifecycleStateBackupInProgress        AutonomousDatabaseInternalSummaryLifecycleStateEnum = "BACKUP_IN_PROGRESS"
	AutonomousDatabaseInternalSummaryLifecycleStateScaleInProgress         AutonomousDatabaseInternalSummaryLifecycleStateEnum = "SCALE_IN_PROGRESS"
	AutonomousDatabaseInternalSummaryLifecycleStateAvailableNeedsAttention AutonomousDatabaseInternalSummaryLifecycleStateEnum = "AVAILABLE_NEEDS_ATTENTION"
	AutonomousDatabaseInternalSummaryLifecycleStateUpdating                AutonomousDatabaseInternalSummaryLifecycleStateEnum = "UPDATING"
	AutonomousDatabaseInternalSummaryLifecycleStateMaintenanceInProgress   AutonomousDatabaseInternalSummaryLifecycleStateEnum = "MAINTENANCE_IN_PROGRESS"
	AutonomousDatabaseInternalSummaryLifecycleStateRestarting              AutonomousDatabaseInternalSummaryLifecycleStateEnum = "RESTARTING"
	AutonomousDatabaseInternalSummaryLifecycleStateRecreating              AutonomousDatabaseInternalSummaryLifecycleStateEnum = "RECREATING"
	AutonomousDatabaseInternalSummaryLifecycleStateRoleChangeInProgress    AutonomousDatabaseInternalSummaryLifecycleStateEnum = "ROLE_CHANGE_IN_PROGRESS"
	AutonomousDatabaseInternalSummaryLifecycleStateUpgrading               AutonomousDatabaseInternalSummaryLifecycleStateEnum = "UPGRADING"
	AutonomousDatabaseInternalSummaryLifecycleStateInaccessible            AutonomousDatabaseInternalSummaryLifecycleStateEnum = "INACCESSIBLE"
	AutonomousDatabaseInternalSummaryLifecycleStateStandby                 AutonomousDatabaseInternalSummaryLifecycleStateEnum = "STANDBY"
	AutonomousDatabaseInternalSummaryLifecycleStateTransporting            AutonomousDatabaseInternalSummaryLifecycleStateEnum = "TRANSPORTING"
)

var mappingAutonomousDatabaseInternalSummaryLifecycleStateEnum = map[string]AutonomousDatabaseInternalSummaryLifecycleStateEnum{
	"PROVISIONING":              AutonomousDatabaseInternalSummaryLifecycleStateProvisioning,
	"AVAILABLE":                 AutonomousDatabaseInternalSummaryLifecycleStateAvailable,
	"STOPPING":                  AutonomousDatabaseInternalSummaryLifecycleStateStopping,
	"STOPPED":                   AutonomousDatabaseInternalSummaryLifecycleStateStopped,
	"STARTING":                  AutonomousDatabaseInternalSummaryLifecycleStateStarting,
	"TERMINATING":               AutonomousDatabaseInternalSummaryLifecycleStateTerminating,
	"TERMINATED":                AutonomousDatabaseInternalSummaryLifecycleStateTerminated,
	"UNAVAILABLE":               AutonomousDatabaseInternalSummaryLifecycleStateUnavailable,
	"RESTORE_IN_PROGRESS":       AutonomousDatabaseInternalSummaryLifecycleStateRestoreInProgress,
	"RESTORE_FAILED":            AutonomousDatabaseInternalSummaryLifecycleStateRestoreFailed,
	"BACKUP_IN_PROGRESS":        AutonomousDatabaseInternalSummaryLifecycleStateBackupInProgress,
	"SCALE_IN_PROGRESS":         AutonomousDatabaseInternalSummaryLifecycleStateScaleInProgress,
	"AVAILABLE_NEEDS_ATTENTION": AutonomousDatabaseInternalSummaryLifecycleStateAvailableNeedsAttention,
	"UPDATING":                  AutonomousDatabaseInternalSummaryLifecycleStateUpdating,
	"MAINTENANCE_IN_PROGRESS":   AutonomousDatabaseInternalSummaryLifecycleStateMaintenanceInProgress,
	"RESTARTING":                AutonomousDatabaseInternalSummaryLifecycleStateRestarting,
	"RECREATING":                AutonomousDatabaseInternalSummaryLifecycleStateRecreating,
	"ROLE_CHANGE_IN_PROGRESS":   AutonomousDatabaseInternalSummaryLifecycleStateRoleChangeInProgress,
	"UPGRADING":                 AutonomousDatabaseInternalSummaryLifecycleStateUpgrading,
	"INACCESSIBLE":              AutonomousDatabaseInternalSummaryLifecycleStateInaccessible,
	"STANDBY":                   AutonomousDatabaseInternalSummaryLifecycleStateStandby,
	"TRANSPORTING":              AutonomousDatabaseInternalSummaryLifecycleStateTransporting,
}

var mappingAutonomousDatabaseInternalSummaryLifecycleStateEnumLowerCase = map[string]AutonomousDatabaseInternalSummaryLifecycleStateEnum{
	"provisioning":              AutonomousDatabaseInternalSummaryLifecycleStateProvisioning,
	"available":                 AutonomousDatabaseInternalSummaryLifecycleStateAvailable,
	"stopping":                  AutonomousDatabaseInternalSummaryLifecycleStateStopping,
	"stopped":                   AutonomousDatabaseInternalSummaryLifecycleStateStopped,
	"starting":                  AutonomousDatabaseInternalSummaryLifecycleStateStarting,
	"terminating":               AutonomousDatabaseInternalSummaryLifecycleStateTerminating,
	"terminated":                AutonomousDatabaseInternalSummaryLifecycleStateTerminated,
	"unavailable":               AutonomousDatabaseInternalSummaryLifecycleStateUnavailable,
	"restore_in_progress":       AutonomousDatabaseInternalSummaryLifecycleStateRestoreInProgress,
	"restore_failed":            AutonomousDatabaseInternalSummaryLifecycleStateRestoreFailed,
	"backup_in_progress":        AutonomousDatabaseInternalSummaryLifecycleStateBackupInProgress,
	"scale_in_progress":         AutonomousDatabaseInternalSummaryLifecycleStateScaleInProgress,
	"available_needs_attention": AutonomousDatabaseInternalSummaryLifecycleStateAvailableNeedsAttention,
	"updating":                  AutonomousDatabaseInternalSummaryLifecycleStateUpdating,
	"maintenance_in_progress":   AutonomousDatabaseInternalSummaryLifecycleStateMaintenanceInProgress,
	"restarting":                AutonomousDatabaseInternalSummaryLifecycleStateRestarting,
	"recreating":                AutonomousDatabaseInternalSummaryLifecycleStateRecreating,
	"role_change_in_progress":   AutonomousDatabaseInternalSummaryLifecycleStateRoleChangeInProgress,
	"upgrading":                 AutonomousDatabaseInternalSummaryLifecycleStateUpgrading,
	"inaccessible":              AutonomousDatabaseInternalSummaryLifecycleStateInaccessible,
	"standby":                   AutonomousDatabaseInternalSummaryLifecycleStateStandby,
	"transporting":              AutonomousDatabaseInternalSummaryLifecycleStateTransporting,
}

// GetAutonomousDatabaseInternalSummaryLifecycleStateEnumValues Enumerates the set of values for AutonomousDatabaseInternalSummaryLifecycleStateEnum
func GetAutonomousDatabaseInternalSummaryLifecycleStateEnumValues() []AutonomousDatabaseInternalSummaryLifecycleStateEnum {
	values := make([]AutonomousDatabaseInternalSummaryLifecycleStateEnum, 0)
	for _, v := range mappingAutonomousDatabaseInternalSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDatabaseInternalSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for AutonomousDatabaseInternalSummaryLifecycleStateEnum
func GetAutonomousDatabaseInternalSummaryLifecycleStateEnumStringValues() []string {
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
	}
}

// GetMappingAutonomousDatabaseInternalSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDatabaseInternalSummaryLifecycleStateEnum(val string) (AutonomousDatabaseInternalSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingAutonomousDatabaseInternalSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
