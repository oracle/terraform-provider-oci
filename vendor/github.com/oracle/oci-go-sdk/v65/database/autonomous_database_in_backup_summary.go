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

// AutonomousDatabaseInBackupSummary The summary of the autonomous database that are in the Autonomous Container Database Backup associated.
type AutonomousDatabaseInBackupSummary struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The user-friendly name for the Autonomous AI Database. The name does not have to be unique.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The current state of the Autonomous AI Database.
	LifecycleState AutonomousDatabaseInBackupSummaryLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m AutonomousDatabaseInBackupSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AutonomousDatabaseInBackupSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingAutonomousDatabaseInBackupSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAutonomousDatabaseInBackupSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AutonomousDatabaseInBackupSummaryLifecycleStateEnum Enum with underlying type: string
type AutonomousDatabaseInBackupSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for AutonomousDatabaseInBackupSummaryLifecycleStateEnum
const (
	AutonomousDatabaseInBackupSummaryLifecycleStateProvisioning            AutonomousDatabaseInBackupSummaryLifecycleStateEnum = "PROVISIONING"
	AutonomousDatabaseInBackupSummaryLifecycleStateAvailable               AutonomousDatabaseInBackupSummaryLifecycleStateEnum = "AVAILABLE"
	AutonomousDatabaseInBackupSummaryLifecycleStateStopping                AutonomousDatabaseInBackupSummaryLifecycleStateEnum = "STOPPING"
	AutonomousDatabaseInBackupSummaryLifecycleStateStopped                 AutonomousDatabaseInBackupSummaryLifecycleStateEnum = "STOPPED"
	AutonomousDatabaseInBackupSummaryLifecycleStateStarting                AutonomousDatabaseInBackupSummaryLifecycleStateEnum = "STARTING"
	AutonomousDatabaseInBackupSummaryLifecycleStateTerminating             AutonomousDatabaseInBackupSummaryLifecycleStateEnum = "TERMINATING"
	AutonomousDatabaseInBackupSummaryLifecycleStateTerminated              AutonomousDatabaseInBackupSummaryLifecycleStateEnum = "TERMINATED"
	AutonomousDatabaseInBackupSummaryLifecycleStateUnavailable             AutonomousDatabaseInBackupSummaryLifecycleStateEnum = "UNAVAILABLE"
	AutonomousDatabaseInBackupSummaryLifecycleStateRestoreInProgress       AutonomousDatabaseInBackupSummaryLifecycleStateEnum = "RESTORE_IN_PROGRESS"
	AutonomousDatabaseInBackupSummaryLifecycleStateRestoreFailed           AutonomousDatabaseInBackupSummaryLifecycleStateEnum = "RESTORE_FAILED"
	AutonomousDatabaseInBackupSummaryLifecycleStateBackupInProgress        AutonomousDatabaseInBackupSummaryLifecycleStateEnum = "BACKUP_IN_PROGRESS"
	AutonomousDatabaseInBackupSummaryLifecycleStateScaleInProgress         AutonomousDatabaseInBackupSummaryLifecycleStateEnum = "SCALE_IN_PROGRESS"
	AutonomousDatabaseInBackupSummaryLifecycleStateAvailableNeedsAttention AutonomousDatabaseInBackupSummaryLifecycleStateEnum = "AVAILABLE_NEEDS_ATTENTION"
	AutonomousDatabaseInBackupSummaryLifecycleStateUpdating                AutonomousDatabaseInBackupSummaryLifecycleStateEnum = "UPDATING"
	AutonomousDatabaseInBackupSummaryLifecycleStateMaintenanceInProgress   AutonomousDatabaseInBackupSummaryLifecycleStateEnum = "MAINTENANCE_IN_PROGRESS"
	AutonomousDatabaseInBackupSummaryLifecycleStateRestarting              AutonomousDatabaseInBackupSummaryLifecycleStateEnum = "RESTARTING"
	AutonomousDatabaseInBackupSummaryLifecycleStateRecreating              AutonomousDatabaseInBackupSummaryLifecycleStateEnum = "RECREATING"
	AutonomousDatabaseInBackupSummaryLifecycleStateRoleChangeInProgress    AutonomousDatabaseInBackupSummaryLifecycleStateEnum = "ROLE_CHANGE_IN_PROGRESS"
	AutonomousDatabaseInBackupSummaryLifecycleStateUpgrading               AutonomousDatabaseInBackupSummaryLifecycleStateEnum = "UPGRADING"
	AutonomousDatabaseInBackupSummaryLifecycleStateInaccessible            AutonomousDatabaseInBackupSummaryLifecycleStateEnum = "INACCESSIBLE"
	AutonomousDatabaseInBackupSummaryLifecycleStateStandby                 AutonomousDatabaseInBackupSummaryLifecycleStateEnum = "STANDBY"
	AutonomousDatabaseInBackupSummaryLifecycleStateTransporting            AutonomousDatabaseInBackupSummaryLifecycleStateEnum = "TRANSPORTING"
)

var mappingAutonomousDatabaseInBackupSummaryLifecycleStateEnum = map[string]AutonomousDatabaseInBackupSummaryLifecycleStateEnum{
	"PROVISIONING":              AutonomousDatabaseInBackupSummaryLifecycleStateProvisioning,
	"AVAILABLE":                 AutonomousDatabaseInBackupSummaryLifecycleStateAvailable,
	"STOPPING":                  AutonomousDatabaseInBackupSummaryLifecycleStateStopping,
	"STOPPED":                   AutonomousDatabaseInBackupSummaryLifecycleStateStopped,
	"STARTING":                  AutonomousDatabaseInBackupSummaryLifecycleStateStarting,
	"TERMINATING":               AutonomousDatabaseInBackupSummaryLifecycleStateTerminating,
	"TERMINATED":                AutonomousDatabaseInBackupSummaryLifecycleStateTerminated,
	"UNAVAILABLE":               AutonomousDatabaseInBackupSummaryLifecycleStateUnavailable,
	"RESTORE_IN_PROGRESS":       AutonomousDatabaseInBackupSummaryLifecycleStateRestoreInProgress,
	"RESTORE_FAILED":            AutonomousDatabaseInBackupSummaryLifecycleStateRestoreFailed,
	"BACKUP_IN_PROGRESS":        AutonomousDatabaseInBackupSummaryLifecycleStateBackupInProgress,
	"SCALE_IN_PROGRESS":         AutonomousDatabaseInBackupSummaryLifecycleStateScaleInProgress,
	"AVAILABLE_NEEDS_ATTENTION": AutonomousDatabaseInBackupSummaryLifecycleStateAvailableNeedsAttention,
	"UPDATING":                  AutonomousDatabaseInBackupSummaryLifecycleStateUpdating,
	"MAINTENANCE_IN_PROGRESS":   AutonomousDatabaseInBackupSummaryLifecycleStateMaintenanceInProgress,
	"RESTARTING":                AutonomousDatabaseInBackupSummaryLifecycleStateRestarting,
	"RECREATING":                AutonomousDatabaseInBackupSummaryLifecycleStateRecreating,
	"ROLE_CHANGE_IN_PROGRESS":   AutonomousDatabaseInBackupSummaryLifecycleStateRoleChangeInProgress,
	"UPGRADING":                 AutonomousDatabaseInBackupSummaryLifecycleStateUpgrading,
	"INACCESSIBLE":              AutonomousDatabaseInBackupSummaryLifecycleStateInaccessible,
	"STANDBY":                   AutonomousDatabaseInBackupSummaryLifecycleStateStandby,
	"TRANSPORTING":              AutonomousDatabaseInBackupSummaryLifecycleStateTransporting,
}

var mappingAutonomousDatabaseInBackupSummaryLifecycleStateEnumLowerCase = map[string]AutonomousDatabaseInBackupSummaryLifecycleStateEnum{
	"provisioning":              AutonomousDatabaseInBackupSummaryLifecycleStateProvisioning,
	"available":                 AutonomousDatabaseInBackupSummaryLifecycleStateAvailable,
	"stopping":                  AutonomousDatabaseInBackupSummaryLifecycleStateStopping,
	"stopped":                   AutonomousDatabaseInBackupSummaryLifecycleStateStopped,
	"starting":                  AutonomousDatabaseInBackupSummaryLifecycleStateStarting,
	"terminating":               AutonomousDatabaseInBackupSummaryLifecycleStateTerminating,
	"terminated":                AutonomousDatabaseInBackupSummaryLifecycleStateTerminated,
	"unavailable":               AutonomousDatabaseInBackupSummaryLifecycleStateUnavailable,
	"restore_in_progress":       AutonomousDatabaseInBackupSummaryLifecycleStateRestoreInProgress,
	"restore_failed":            AutonomousDatabaseInBackupSummaryLifecycleStateRestoreFailed,
	"backup_in_progress":        AutonomousDatabaseInBackupSummaryLifecycleStateBackupInProgress,
	"scale_in_progress":         AutonomousDatabaseInBackupSummaryLifecycleStateScaleInProgress,
	"available_needs_attention": AutonomousDatabaseInBackupSummaryLifecycleStateAvailableNeedsAttention,
	"updating":                  AutonomousDatabaseInBackupSummaryLifecycleStateUpdating,
	"maintenance_in_progress":   AutonomousDatabaseInBackupSummaryLifecycleStateMaintenanceInProgress,
	"restarting":                AutonomousDatabaseInBackupSummaryLifecycleStateRestarting,
	"recreating":                AutonomousDatabaseInBackupSummaryLifecycleStateRecreating,
	"role_change_in_progress":   AutonomousDatabaseInBackupSummaryLifecycleStateRoleChangeInProgress,
	"upgrading":                 AutonomousDatabaseInBackupSummaryLifecycleStateUpgrading,
	"inaccessible":              AutonomousDatabaseInBackupSummaryLifecycleStateInaccessible,
	"standby":                   AutonomousDatabaseInBackupSummaryLifecycleStateStandby,
	"transporting":              AutonomousDatabaseInBackupSummaryLifecycleStateTransporting,
}

// GetAutonomousDatabaseInBackupSummaryLifecycleStateEnumValues Enumerates the set of values for AutonomousDatabaseInBackupSummaryLifecycleStateEnum
func GetAutonomousDatabaseInBackupSummaryLifecycleStateEnumValues() []AutonomousDatabaseInBackupSummaryLifecycleStateEnum {
	values := make([]AutonomousDatabaseInBackupSummaryLifecycleStateEnum, 0)
	for _, v := range mappingAutonomousDatabaseInBackupSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDatabaseInBackupSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for AutonomousDatabaseInBackupSummaryLifecycleStateEnum
func GetAutonomousDatabaseInBackupSummaryLifecycleStateEnumStringValues() []string {
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

// GetMappingAutonomousDatabaseInBackupSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDatabaseInBackupSummaryLifecycleStateEnum(val string) (AutonomousDatabaseInBackupSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingAutonomousDatabaseInBackupSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
