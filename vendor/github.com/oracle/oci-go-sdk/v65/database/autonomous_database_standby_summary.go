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
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AutonomousDatabaseStandbySummary Autonomous Data Guard standby database details.
type AutonomousDatabaseStandbySummary struct {

	// The amount of time, in seconds, that the data of the standby database lags the data of the primary database. Can be used to determine the potential data loss in the event of a failover.
	LagTimeInSeconds *int `mandatory:"false" json:"lagTimeInSeconds"`

	// The current state of the Autonomous Database.
	LifecycleState AutonomousDatabaseStandbySummaryLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The date and time the Autonomous Data Guard role was switched for the standby Autonomous Database.
	TimeDataGuardRoleChanged *common.SDKTime `mandatory:"false" json:"timeDataGuardRoleChanged"`
}

func (m AutonomousDatabaseStandbySummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AutonomousDatabaseStandbySummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingAutonomousDatabaseStandbySummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAutonomousDatabaseStandbySummaryLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AutonomousDatabaseStandbySummaryLifecycleStateEnum Enum with underlying type: string
type AutonomousDatabaseStandbySummaryLifecycleStateEnum string

// Set of constants representing the allowable values for AutonomousDatabaseStandbySummaryLifecycleStateEnum
const (
	AutonomousDatabaseStandbySummaryLifecycleStateProvisioning            AutonomousDatabaseStandbySummaryLifecycleStateEnum = "PROVISIONING"
	AutonomousDatabaseStandbySummaryLifecycleStateAvailable               AutonomousDatabaseStandbySummaryLifecycleStateEnum = "AVAILABLE"
	AutonomousDatabaseStandbySummaryLifecycleStateStopping                AutonomousDatabaseStandbySummaryLifecycleStateEnum = "STOPPING"
	AutonomousDatabaseStandbySummaryLifecycleStateStopped                 AutonomousDatabaseStandbySummaryLifecycleStateEnum = "STOPPED"
	AutonomousDatabaseStandbySummaryLifecycleStateStarting                AutonomousDatabaseStandbySummaryLifecycleStateEnum = "STARTING"
	AutonomousDatabaseStandbySummaryLifecycleStateTerminating             AutonomousDatabaseStandbySummaryLifecycleStateEnum = "TERMINATING"
	AutonomousDatabaseStandbySummaryLifecycleStateTerminated              AutonomousDatabaseStandbySummaryLifecycleStateEnum = "TERMINATED"
	AutonomousDatabaseStandbySummaryLifecycleStateUnavailable             AutonomousDatabaseStandbySummaryLifecycleStateEnum = "UNAVAILABLE"
	AutonomousDatabaseStandbySummaryLifecycleStateRestoreInProgress       AutonomousDatabaseStandbySummaryLifecycleStateEnum = "RESTORE_IN_PROGRESS"
	AutonomousDatabaseStandbySummaryLifecycleStateRestoreFailed           AutonomousDatabaseStandbySummaryLifecycleStateEnum = "RESTORE_FAILED"
	AutonomousDatabaseStandbySummaryLifecycleStateBackupInProgress        AutonomousDatabaseStandbySummaryLifecycleStateEnum = "BACKUP_IN_PROGRESS"
	AutonomousDatabaseStandbySummaryLifecycleStateScaleInProgress         AutonomousDatabaseStandbySummaryLifecycleStateEnum = "SCALE_IN_PROGRESS"
	AutonomousDatabaseStandbySummaryLifecycleStateAvailableNeedsAttention AutonomousDatabaseStandbySummaryLifecycleStateEnum = "AVAILABLE_NEEDS_ATTENTION"
	AutonomousDatabaseStandbySummaryLifecycleStateUpdating                AutonomousDatabaseStandbySummaryLifecycleStateEnum = "UPDATING"
	AutonomousDatabaseStandbySummaryLifecycleStateMaintenanceInProgress   AutonomousDatabaseStandbySummaryLifecycleStateEnum = "MAINTENANCE_IN_PROGRESS"
	AutonomousDatabaseStandbySummaryLifecycleStateRestarting              AutonomousDatabaseStandbySummaryLifecycleStateEnum = "RESTARTING"
	AutonomousDatabaseStandbySummaryLifecycleStateRecreating              AutonomousDatabaseStandbySummaryLifecycleStateEnum = "RECREATING"
	AutonomousDatabaseStandbySummaryLifecycleStateRoleChangeInProgress    AutonomousDatabaseStandbySummaryLifecycleStateEnum = "ROLE_CHANGE_IN_PROGRESS"
	AutonomousDatabaseStandbySummaryLifecycleStateUpgrading               AutonomousDatabaseStandbySummaryLifecycleStateEnum = "UPGRADING"
	AutonomousDatabaseStandbySummaryLifecycleStateInaccessible            AutonomousDatabaseStandbySummaryLifecycleStateEnum = "INACCESSIBLE"
	AutonomousDatabaseStandbySummaryLifecycleStateStandby                 AutonomousDatabaseStandbySummaryLifecycleStateEnum = "STANDBY"
)

var mappingAutonomousDatabaseStandbySummaryLifecycleStateEnum = map[string]AutonomousDatabaseStandbySummaryLifecycleStateEnum{
	"PROVISIONING":              AutonomousDatabaseStandbySummaryLifecycleStateProvisioning,
	"AVAILABLE":                 AutonomousDatabaseStandbySummaryLifecycleStateAvailable,
	"STOPPING":                  AutonomousDatabaseStandbySummaryLifecycleStateStopping,
	"STOPPED":                   AutonomousDatabaseStandbySummaryLifecycleStateStopped,
	"STARTING":                  AutonomousDatabaseStandbySummaryLifecycleStateStarting,
	"TERMINATING":               AutonomousDatabaseStandbySummaryLifecycleStateTerminating,
	"TERMINATED":                AutonomousDatabaseStandbySummaryLifecycleStateTerminated,
	"UNAVAILABLE":               AutonomousDatabaseStandbySummaryLifecycleStateUnavailable,
	"RESTORE_IN_PROGRESS":       AutonomousDatabaseStandbySummaryLifecycleStateRestoreInProgress,
	"RESTORE_FAILED":            AutonomousDatabaseStandbySummaryLifecycleStateRestoreFailed,
	"BACKUP_IN_PROGRESS":        AutonomousDatabaseStandbySummaryLifecycleStateBackupInProgress,
	"SCALE_IN_PROGRESS":         AutonomousDatabaseStandbySummaryLifecycleStateScaleInProgress,
	"AVAILABLE_NEEDS_ATTENTION": AutonomousDatabaseStandbySummaryLifecycleStateAvailableNeedsAttention,
	"UPDATING":                  AutonomousDatabaseStandbySummaryLifecycleStateUpdating,
	"MAINTENANCE_IN_PROGRESS":   AutonomousDatabaseStandbySummaryLifecycleStateMaintenanceInProgress,
	"RESTARTING":                AutonomousDatabaseStandbySummaryLifecycleStateRestarting,
	"RECREATING":                AutonomousDatabaseStandbySummaryLifecycleStateRecreating,
	"ROLE_CHANGE_IN_PROGRESS":   AutonomousDatabaseStandbySummaryLifecycleStateRoleChangeInProgress,
	"UPGRADING":                 AutonomousDatabaseStandbySummaryLifecycleStateUpgrading,
	"INACCESSIBLE":              AutonomousDatabaseStandbySummaryLifecycleStateInaccessible,
	"STANDBY":                   AutonomousDatabaseStandbySummaryLifecycleStateStandby,
}

var mappingAutonomousDatabaseStandbySummaryLifecycleStateEnumLowerCase = map[string]AutonomousDatabaseStandbySummaryLifecycleStateEnum{
	"provisioning":              AutonomousDatabaseStandbySummaryLifecycleStateProvisioning,
	"available":                 AutonomousDatabaseStandbySummaryLifecycleStateAvailable,
	"stopping":                  AutonomousDatabaseStandbySummaryLifecycleStateStopping,
	"stopped":                   AutonomousDatabaseStandbySummaryLifecycleStateStopped,
	"starting":                  AutonomousDatabaseStandbySummaryLifecycleStateStarting,
	"terminating":               AutonomousDatabaseStandbySummaryLifecycleStateTerminating,
	"terminated":                AutonomousDatabaseStandbySummaryLifecycleStateTerminated,
	"unavailable":               AutonomousDatabaseStandbySummaryLifecycleStateUnavailable,
	"restore_in_progress":       AutonomousDatabaseStandbySummaryLifecycleStateRestoreInProgress,
	"restore_failed":            AutonomousDatabaseStandbySummaryLifecycleStateRestoreFailed,
	"backup_in_progress":        AutonomousDatabaseStandbySummaryLifecycleStateBackupInProgress,
	"scale_in_progress":         AutonomousDatabaseStandbySummaryLifecycleStateScaleInProgress,
	"available_needs_attention": AutonomousDatabaseStandbySummaryLifecycleStateAvailableNeedsAttention,
	"updating":                  AutonomousDatabaseStandbySummaryLifecycleStateUpdating,
	"maintenance_in_progress":   AutonomousDatabaseStandbySummaryLifecycleStateMaintenanceInProgress,
	"restarting":                AutonomousDatabaseStandbySummaryLifecycleStateRestarting,
	"recreating":                AutonomousDatabaseStandbySummaryLifecycleStateRecreating,
	"role_change_in_progress":   AutonomousDatabaseStandbySummaryLifecycleStateRoleChangeInProgress,
	"upgrading":                 AutonomousDatabaseStandbySummaryLifecycleStateUpgrading,
	"inaccessible":              AutonomousDatabaseStandbySummaryLifecycleStateInaccessible,
	"standby":                   AutonomousDatabaseStandbySummaryLifecycleStateStandby,
}

// GetAutonomousDatabaseStandbySummaryLifecycleStateEnumValues Enumerates the set of values for AutonomousDatabaseStandbySummaryLifecycleStateEnum
func GetAutonomousDatabaseStandbySummaryLifecycleStateEnumValues() []AutonomousDatabaseStandbySummaryLifecycleStateEnum {
	values := make([]AutonomousDatabaseStandbySummaryLifecycleStateEnum, 0)
	for _, v := range mappingAutonomousDatabaseStandbySummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDatabaseStandbySummaryLifecycleStateEnumStringValues Enumerates the set of values in String for AutonomousDatabaseStandbySummaryLifecycleStateEnum
func GetAutonomousDatabaseStandbySummaryLifecycleStateEnumStringValues() []string {
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
	}
}

// GetMappingAutonomousDatabaseStandbySummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDatabaseStandbySummaryLifecycleStateEnum(val string) (AutonomousDatabaseStandbySummaryLifecycleStateEnum, bool) {
	enum, ok := mappingAutonomousDatabaseStandbySummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
