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

// AutonomousDatabaseInBackup Details of Autonomous AI Database in Autonomous Container Database
type AutonomousDatabaseInBackup struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The user-friendly name for the Autonomous AI Database. The name does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The current state of the Autonomous AI Database.
	LifecycleState AutonomousDatabaseInBackupLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m AutonomousDatabaseInBackup) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AutonomousDatabaseInBackup) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingAutonomousDatabaseInBackupLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAutonomousDatabaseInBackupLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AutonomousDatabaseInBackupLifecycleStateEnum Enum with underlying type: string
type AutonomousDatabaseInBackupLifecycleStateEnum string

// Set of constants representing the allowable values for AutonomousDatabaseInBackupLifecycleStateEnum
const (
	AutonomousDatabaseInBackupLifecycleStateProvisioning            AutonomousDatabaseInBackupLifecycleStateEnum = "PROVISIONING"
	AutonomousDatabaseInBackupLifecycleStateAvailable               AutonomousDatabaseInBackupLifecycleStateEnum = "AVAILABLE"
	AutonomousDatabaseInBackupLifecycleStateStopping                AutonomousDatabaseInBackupLifecycleStateEnum = "STOPPING"
	AutonomousDatabaseInBackupLifecycleStateStopped                 AutonomousDatabaseInBackupLifecycleStateEnum = "STOPPED"
	AutonomousDatabaseInBackupLifecycleStateStarting                AutonomousDatabaseInBackupLifecycleStateEnum = "STARTING"
	AutonomousDatabaseInBackupLifecycleStateTerminating             AutonomousDatabaseInBackupLifecycleStateEnum = "TERMINATING"
	AutonomousDatabaseInBackupLifecycleStateTerminated              AutonomousDatabaseInBackupLifecycleStateEnum = "TERMINATED"
	AutonomousDatabaseInBackupLifecycleStateUnavailable             AutonomousDatabaseInBackupLifecycleStateEnum = "UNAVAILABLE"
	AutonomousDatabaseInBackupLifecycleStateRestoreInProgress       AutonomousDatabaseInBackupLifecycleStateEnum = "RESTORE_IN_PROGRESS"
	AutonomousDatabaseInBackupLifecycleStateRestoreFailed           AutonomousDatabaseInBackupLifecycleStateEnum = "RESTORE_FAILED"
	AutonomousDatabaseInBackupLifecycleStateBackupInProgress        AutonomousDatabaseInBackupLifecycleStateEnum = "BACKUP_IN_PROGRESS"
	AutonomousDatabaseInBackupLifecycleStateScaleInProgress         AutonomousDatabaseInBackupLifecycleStateEnum = "SCALE_IN_PROGRESS"
	AutonomousDatabaseInBackupLifecycleStateAvailableNeedsAttention AutonomousDatabaseInBackupLifecycleStateEnum = "AVAILABLE_NEEDS_ATTENTION"
	AutonomousDatabaseInBackupLifecycleStateUpdating                AutonomousDatabaseInBackupLifecycleStateEnum = "UPDATING"
	AutonomousDatabaseInBackupLifecycleStateMaintenanceInProgress   AutonomousDatabaseInBackupLifecycleStateEnum = "MAINTENANCE_IN_PROGRESS"
	AutonomousDatabaseInBackupLifecycleStateRestarting              AutonomousDatabaseInBackupLifecycleStateEnum = "RESTARTING"
	AutonomousDatabaseInBackupLifecycleStateRecreating              AutonomousDatabaseInBackupLifecycleStateEnum = "RECREATING"
	AutonomousDatabaseInBackupLifecycleStateRoleChangeInProgress    AutonomousDatabaseInBackupLifecycleStateEnum = "ROLE_CHANGE_IN_PROGRESS"
	AutonomousDatabaseInBackupLifecycleStateUpgrading               AutonomousDatabaseInBackupLifecycleStateEnum = "UPGRADING"
	AutonomousDatabaseInBackupLifecycleStateInaccessible            AutonomousDatabaseInBackupLifecycleStateEnum = "INACCESSIBLE"
	AutonomousDatabaseInBackupLifecycleStateStandby                 AutonomousDatabaseInBackupLifecycleStateEnum = "STANDBY"
	AutonomousDatabaseInBackupLifecycleStateTransporting            AutonomousDatabaseInBackupLifecycleStateEnum = "TRANSPORTING"
	AutonomousDatabaseInBackupLifecycleStateArchived                AutonomousDatabaseInBackupLifecycleStateEnum = "ARCHIVED"
)

var mappingAutonomousDatabaseInBackupLifecycleStateEnum = map[string]AutonomousDatabaseInBackupLifecycleStateEnum{
	"PROVISIONING":              AutonomousDatabaseInBackupLifecycleStateProvisioning,
	"AVAILABLE":                 AutonomousDatabaseInBackupLifecycleStateAvailable,
	"STOPPING":                  AutonomousDatabaseInBackupLifecycleStateStopping,
	"STOPPED":                   AutonomousDatabaseInBackupLifecycleStateStopped,
	"STARTING":                  AutonomousDatabaseInBackupLifecycleStateStarting,
	"TERMINATING":               AutonomousDatabaseInBackupLifecycleStateTerminating,
	"TERMINATED":                AutonomousDatabaseInBackupLifecycleStateTerminated,
	"UNAVAILABLE":               AutonomousDatabaseInBackupLifecycleStateUnavailable,
	"RESTORE_IN_PROGRESS":       AutonomousDatabaseInBackupLifecycleStateRestoreInProgress,
	"RESTORE_FAILED":            AutonomousDatabaseInBackupLifecycleStateRestoreFailed,
	"BACKUP_IN_PROGRESS":        AutonomousDatabaseInBackupLifecycleStateBackupInProgress,
	"SCALE_IN_PROGRESS":         AutonomousDatabaseInBackupLifecycleStateScaleInProgress,
	"AVAILABLE_NEEDS_ATTENTION": AutonomousDatabaseInBackupLifecycleStateAvailableNeedsAttention,
	"UPDATING":                  AutonomousDatabaseInBackupLifecycleStateUpdating,
	"MAINTENANCE_IN_PROGRESS":   AutonomousDatabaseInBackupLifecycleStateMaintenanceInProgress,
	"RESTARTING":                AutonomousDatabaseInBackupLifecycleStateRestarting,
	"RECREATING":                AutonomousDatabaseInBackupLifecycleStateRecreating,
	"ROLE_CHANGE_IN_PROGRESS":   AutonomousDatabaseInBackupLifecycleStateRoleChangeInProgress,
	"UPGRADING":                 AutonomousDatabaseInBackupLifecycleStateUpgrading,
	"INACCESSIBLE":              AutonomousDatabaseInBackupLifecycleStateInaccessible,
	"STANDBY":                   AutonomousDatabaseInBackupLifecycleStateStandby,
	"TRANSPORTING":              AutonomousDatabaseInBackupLifecycleStateTransporting,
	"ARCHIVED":                  AutonomousDatabaseInBackupLifecycleStateArchived,
}

var mappingAutonomousDatabaseInBackupLifecycleStateEnumLowerCase = map[string]AutonomousDatabaseInBackupLifecycleStateEnum{
	"provisioning":              AutonomousDatabaseInBackupLifecycleStateProvisioning,
	"available":                 AutonomousDatabaseInBackupLifecycleStateAvailable,
	"stopping":                  AutonomousDatabaseInBackupLifecycleStateStopping,
	"stopped":                   AutonomousDatabaseInBackupLifecycleStateStopped,
	"starting":                  AutonomousDatabaseInBackupLifecycleStateStarting,
	"terminating":               AutonomousDatabaseInBackupLifecycleStateTerminating,
	"terminated":                AutonomousDatabaseInBackupLifecycleStateTerminated,
	"unavailable":               AutonomousDatabaseInBackupLifecycleStateUnavailable,
	"restore_in_progress":       AutonomousDatabaseInBackupLifecycleStateRestoreInProgress,
	"restore_failed":            AutonomousDatabaseInBackupLifecycleStateRestoreFailed,
	"backup_in_progress":        AutonomousDatabaseInBackupLifecycleStateBackupInProgress,
	"scale_in_progress":         AutonomousDatabaseInBackupLifecycleStateScaleInProgress,
	"available_needs_attention": AutonomousDatabaseInBackupLifecycleStateAvailableNeedsAttention,
	"updating":                  AutonomousDatabaseInBackupLifecycleStateUpdating,
	"maintenance_in_progress":   AutonomousDatabaseInBackupLifecycleStateMaintenanceInProgress,
	"restarting":                AutonomousDatabaseInBackupLifecycleStateRestarting,
	"recreating":                AutonomousDatabaseInBackupLifecycleStateRecreating,
	"role_change_in_progress":   AutonomousDatabaseInBackupLifecycleStateRoleChangeInProgress,
	"upgrading":                 AutonomousDatabaseInBackupLifecycleStateUpgrading,
	"inaccessible":              AutonomousDatabaseInBackupLifecycleStateInaccessible,
	"standby":                   AutonomousDatabaseInBackupLifecycleStateStandby,
	"transporting":              AutonomousDatabaseInBackupLifecycleStateTransporting,
	"archived":                  AutonomousDatabaseInBackupLifecycleStateArchived,
}

// GetAutonomousDatabaseInBackupLifecycleStateEnumValues Enumerates the set of values for AutonomousDatabaseInBackupLifecycleStateEnum
func GetAutonomousDatabaseInBackupLifecycleStateEnumValues() []AutonomousDatabaseInBackupLifecycleStateEnum {
	values := make([]AutonomousDatabaseInBackupLifecycleStateEnum, 0)
	for _, v := range mappingAutonomousDatabaseInBackupLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDatabaseInBackupLifecycleStateEnumStringValues Enumerates the set of values in String for AutonomousDatabaseInBackupLifecycleStateEnum
func GetAutonomousDatabaseInBackupLifecycleStateEnumStringValues() []string {
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

// GetMappingAutonomousDatabaseInBackupLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDatabaseInBackupLifecycleStateEnum(val string) (AutonomousDatabaseInBackupLifecycleStateEnum, bool) {
	enum, ok := mappingAutonomousDatabaseInBackupLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
