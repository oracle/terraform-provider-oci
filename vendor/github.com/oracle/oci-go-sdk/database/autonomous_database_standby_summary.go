// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/common"
)

// AutonomousDatabaseStandbySummary Autonomous Data Guard standby database details.
type AutonomousDatabaseStandbySummary struct {

	// The amount of time, in seconds, that the data of the standby database lags the data of the primary database. Can be used to determine the potential data loss in the event of a failover.
	LagTimeInSeconds *int `mandatory:"false" json:"lagTimeInSeconds"`

	// The current state of the Autonomous Database.
	LifecycleState AutonomousDatabaseStandbySummaryLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Additional information about the current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`
}

func (m AutonomousDatabaseStandbySummary) String() string {
	return common.PointerString(m)
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
)

var mappingAutonomousDatabaseStandbySummaryLifecycleState = map[string]AutonomousDatabaseStandbySummaryLifecycleStateEnum{
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
}

// GetAutonomousDatabaseStandbySummaryLifecycleStateEnumValues Enumerates the set of values for AutonomousDatabaseStandbySummaryLifecycleStateEnum
func GetAutonomousDatabaseStandbySummaryLifecycleStateEnumValues() []AutonomousDatabaseStandbySummaryLifecycleStateEnum {
	values := make([]AutonomousDatabaseStandbySummaryLifecycleStateEnum, 0)
	for _, v := range mappingAutonomousDatabaseStandbySummaryLifecycleState {
		values = append(values, v)
	}
	return values
}
