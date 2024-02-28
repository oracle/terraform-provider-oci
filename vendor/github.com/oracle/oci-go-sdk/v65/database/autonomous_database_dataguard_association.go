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

// AutonomousDatabaseDataguardAssociation The properties that define dataguard association between two different Autonomous Databases.
// Note that Autonomous Databases inherit DataGuard association from parent Autonomous Container Database.
// No actions can be taken on AutonomousDatabaseDataguardAssociation, usage is strictly informational.
type AutonomousDatabaseDataguardAssociation struct {

	// The OCID of the Autonomous Dataguard created for Autonomous Container Database where given Autonomous Database resides in.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Autonomous Database that has a relationship with the peer Autonomous Database.
	AutonomousDatabaseId *string `mandatory:"true" json:"autonomousDatabaseId"`

	// The Data Guard role of the Autonomous Container Database or Autonomous Database, if Autonomous Data Guard is enabled.
	Role AutonomousDatabaseDataguardAssociationRoleEnum `mandatory:"true" json:"role"`

	// The current state of Autonomous Data Guard.
	LifecycleState AutonomousDatabaseDataguardAssociationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The Data Guard role of the Autonomous Container Database or Autonomous Database, if Autonomous Data Guard is enabled.
	PeerRole AutonomousDatabaseDataguardAssociationPeerRoleEnum `mandatory:"true" json:"peerRole"`

	// Additional information about the current lifecycleState, if available.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the peer Autonomous Database.
	PeerAutonomousDatabaseId *string `mandatory:"false" json:"peerAutonomousDatabaseId"`

	// The current state of the Autonomous Database.
	PeerAutonomousDatabaseLifeCycleState AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateEnum `mandatory:"false" json:"peerAutonomousDatabaseLifeCycleState,omitempty"`

	// The protection mode of this Autonomous Data Guard association. For more information, see
	// Oracle Data Guard Protection Modes (http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-protection-modes.htm#SBYDB02000)
	// in the Oracle Data Guard documentation.
	ProtectionMode AutonomousDatabaseDataguardAssociationProtectionModeEnum `mandatory:"false" json:"protectionMode,omitempty"`

	// The lag time between updates to the primary database and application of the redo data on the standby database,
	// as computed by the reporting database.
	// Example: `9 seconds`
	ApplyLag *string `mandatory:"false" json:"applyLag"`

	// The rate at which redo logs are synced between the associated databases.
	// Example: `180 Mb per second`
	ApplyRate *string `mandatory:"false" json:"applyRate"`

	// Indicates whether Automatic Failover is enabled for Autonomous Container Database Dataguard Association
	IsAutomaticFailoverEnabled *bool `mandatory:"false" json:"isAutomaticFailoverEnabled"`

	// The approximate number of seconds of redo data not yet available on the standby Autonomous Container Database,
	// as computed by the reporting database.
	// Example: `7 seconds`
	TransportLag *string `mandatory:"false" json:"transportLag"`

	// The date and time of the last update to the apply lag, apply rate, and transport lag values.
	TimeLastSynced *common.SDKTime `mandatory:"false" json:"timeLastSynced"`

	// The date and time the Data Guard association was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time when the last role change action happened.
	TimeLastRoleChanged *common.SDKTime `mandatory:"false" json:"timeLastRoleChanged"`
}

func (m AutonomousDatabaseDataguardAssociation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AutonomousDatabaseDataguardAssociation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAutonomousDatabaseDataguardAssociationRoleEnum(string(m.Role)); !ok && m.Role != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Role: %s. Supported values are: %s.", m.Role, strings.Join(GetAutonomousDatabaseDataguardAssociationRoleEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousDatabaseDataguardAssociationLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAutonomousDatabaseDataguardAssociationLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousDatabaseDataguardAssociationPeerRoleEnum(string(m.PeerRole)); !ok && m.PeerRole != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PeerRole: %s. Supported values are: %s.", m.PeerRole, strings.Join(GetAutonomousDatabaseDataguardAssociationPeerRoleEnumStringValues(), ",")))
	}

	if _, ok := GetMappingAutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateEnum(string(m.PeerAutonomousDatabaseLifeCycleState)); !ok && m.PeerAutonomousDatabaseLifeCycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PeerAutonomousDatabaseLifeCycleState: %s. Supported values are: %s.", m.PeerAutonomousDatabaseLifeCycleState, strings.Join(GetAutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousDatabaseDataguardAssociationProtectionModeEnum(string(m.ProtectionMode)); !ok && m.ProtectionMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ProtectionMode: %s. Supported values are: %s.", m.ProtectionMode, strings.Join(GetAutonomousDatabaseDataguardAssociationProtectionModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AutonomousDatabaseDataguardAssociationRoleEnum Enum with underlying type: string
type AutonomousDatabaseDataguardAssociationRoleEnum string

// Set of constants representing the allowable values for AutonomousDatabaseDataguardAssociationRoleEnum
const (
	AutonomousDatabaseDataguardAssociationRolePrimary         AutonomousDatabaseDataguardAssociationRoleEnum = "PRIMARY"
	AutonomousDatabaseDataguardAssociationRoleStandby         AutonomousDatabaseDataguardAssociationRoleEnum = "STANDBY"
	AutonomousDatabaseDataguardAssociationRoleDisabledStandby AutonomousDatabaseDataguardAssociationRoleEnum = "DISABLED_STANDBY"
	AutonomousDatabaseDataguardAssociationRoleBackupCopy      AutonomousDatabaseDataguardAssociationRoleEnum = "BACKUP_COPY"
	AutonomousDatabaseDataguardAssociationRoleSnapshotStandby AutonomousDatabaseDataguardAssociationRoleEnum = "SNAPSHOT_STANDBY"
)

var mappingAutonomousDatabaseDataguardAssociationRoleEnum = map[string]AutonomousDatabaseDataguardAssociationRoleEnum{
	"PRIMARY":          AutonomousDatabaseDataguardAssociationRolePrimary,
	"STANDBY":          AutonomousDatabaseDataguardAssociationRoleStandby,
	"DISABLED_STANDBY": AutonomousDatabaseDataguardAssociationRoleDisabledStandby,
	"BACKUP_COPY":      AutonomousDatabaseDataguardAssociationRoleBackupCopy,
	"SNAPSHOT_STANDBY": AutonomousDatabaseDataguardAssociationRoleSnapshotStandby,
}

var mappingAutonomousDatabaseDataguardAssociationRoleEnumLowerCase = map[string]AutonomousDatabaseDataguardAssociationRoleEnum{
	"primary":          AutonomousDatabaseDataguardAssociationRolePrimary,
	"standby":          AutonomousDatabaseDataguardAssociationRoleStandby,
	"disabled_standby": AutonomousDatabaseDataguardAssociationRoleDisabledStandby,
	"backup_copy":      AutonomousDatabaseDataguardAssociationRoleBackupCopy,
	"snapshot_standby": AutonomousDatabaseDataguardAssociationRoleSnapshotStandby,
}

// GetAutonomousDatabaseDataguardAssociationRoleEnumValues Enumerates the set of values for AutonomousDatabaseDataguardAssociationRoleEnum
func GetAutonomousDatabaseDataguardAssociationRoleEnumValues() []AutonomousDatabaseDataguardAssociationRoleEnum {
	values := make([]AutonomousDatabaseDataguardAssociationRoleEnum, 0)
	for _, v := range mappingAutonomousDatabaseDataguardAssociationRoleEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDatabaseDataguardAssociationRoleEnumStringValues Enumerates the set of values in String for AutonomousDatabaseDataguardAssociationRoleEnum
func GetAutonomousDatabaseDataguardAssociationRoleEnumStringValues() []string {
	return []string{
		"PRIMARY",
		"STANDBY",
		"DISABLED_STANDBY",
		"BACKUP_COPY",
		"SNAPSHOT_STANDBY",
	}
}

// GetMappingAutonomousDatabaseDataguardAssociationRoleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDatabaseDataguardAssociationRoleEnum(val string) (AutonomousDatabaseDataguardAssociationRoleEnum, bool) {
	enum, ok := mappingAutonomousDatabaseDataguardAssociationRoleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousDatabaseDataguardAssociationLifecycleStateEnum Enum with underlying type: string
type AutonomousDatabaseDataguardAssociationLifecycleStateEnum string

// Set of constants representing the allowable values for AutonomousDatabaseDataguardAssociationLifecycleStateEnum
const (
	AutonomousDatabaseDataguardAssociationLifecycleStateProvisioning         AutonomousDatabaseDataguardAssociationLifecycleStateEnum = "PROVISIONING"
	AutonomousDatabaseDataguardAssociationLifecycleStateAvailable            AutonomousDatabaseDataguardAssociationLifecycleStateEnum = "AVAILABLE"
	AutonomousDatabaseDataguardAssociationLifecycleStateRoleChangeInProgress AutonomousDatabaseDataguardAssociationLifecycleStateEnum = "ROLE_CHANGE_IN_PROGRESS"
	AutonomousDatabaseDataguardAssociationLifecycleStateTerminating          AutonomousDatabaseDataguardAssociationLifecycleStateEnum = "TERMINATING"
	AutonomousDatabaseDataguardAssociationLifecycleStateTerminated           AutonomousDatabaseDataguardAssociationLifecycleStateEnum = "TERMINATED"
	AutonomousDatabaseDataguardAssociationLifecycleStateFailed               AutonomousDatabaseDataguardAssociationLifecycleStateEnum = "FAILED"
	AutonomousDatabaseDataguardAssociationLifecycleStateUnavailable          AutonomousDatabaseDataguardAssociationLifecycleStateEnum = "UNAVAILABLE"
	AutonomousDatabaseDataguardAssociationLifecycleStateUpdating             AutonomousDatabaseDataguardAssociationLifecycleStateEnum = "UPDATING"
)

var mappingAutonomousDatabaseDataguardAssociationLifecycleStateEnum = map[string]AutonomousDatabaseDataguardAssociationLifecycleStateEnum{
	"PROVISIONING":            AutonomousDatabaseDataguardAssociationLifecycleStateProvisioning,
	"AVAILABLE":               AutonomousDatabaseDataguardAssociationLifecycleStateAvailable,
	"ROLE_CHANGE_IN_PROGRESS": AutonomousDatabaseDataguardAssociationLifecycleStateRoleChangeInProgress,
	"TERMINATING":             AutonomousDatabaseDataguardAssociationLifecycleStateTerminating,
	"TERMINATED":              AutonomousDatabaseDataguardAssociationLifecycleStateTerminated,
	"FAILED":                  AutonomousDatabaseDataguardAssociationLifecycleStateFailed,
	"UNAVAILABLE":             AutonomousDatabaseDataguardAssociationLifecycleStateUnavailable,
	"UPDATING":                AutonomousDatabaseDataguardAssociationLifecycleStateUpdating,
}

var mappingAutonomousDatabaseDataguardAssociationLifecycleStateEnumLowerCase = map[string]AutonomousDatabaseDataguardAssociationLifecycleStateEnum{
	"provisioning":            AutonomousDatabaseDataguardAssociationLifecycleStateProvisioning,
	"available":               AutonomousDatabaseDataguardAssociationLifecycleStateAvailable,
	"role_change_in_progress": AutonomousDatabaseDataguardAssociationLifecycleStateRoleChangeInProgress,
	"terminating":             AutonomousDatabaseDataguardAssociationLifecycleStateTerminating,
	"terminated":              AutonomousDatabaseDataguardAssociationLifecycleStateTerminated,
	"failed":                  AutonomousDatabaseDataguardAssociationLifecycleStateFailed,
	"unavailable":             AutonomousDatabaseDataguardAssociationLifecycleStateUnavailable,
	"updating":                AutonomousDatabaseDataguardAssociationLifecycleStateUpdating,
}

// GetAutonomousDatabaseDataguardAssociationLifecycleStateEnumValues Enumerates the set of values for AutonomousDatabaseDataguardAssociationLifecycleStateEnum
func GetAutonomousDatabaseDataguardAssociationLifecycleStateEnumValues() []AutonomousDatabaseDataguardAssociationLifecycleStateEnum {
	values := make([]AutonomousDatabaseDataguardAssociationLifecycleStateEnum, 0)
	for _, v := range mappingAutonomousDatabaseDataguardAssociationLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDatabaseDataguardAssociationLifecycleStateEnumStringValues Enumerates the set of values in String for AutonomousDatabaseDataguardAssociationLifecycleStateEnum
func GetAutonomousDatabaseDataguardAssociationLifecycleStateEnumStringValues() []string {
	return []string{
		"PROVISIONING",
		"AVAILABLE",
		"ROLE_CHANGE_IN_PROGRESS",
		"TERMINATING",
		"TERMINATED",
		"FAILED",
		"UNAVAILABLE",
		"UPDATING",
	}
}

// GetMappingAutonomousDatabaseDataguardAssociationLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDatabaseDataguardAssociationLifecycleStateEnum(val string) (AutonomousDatabaseDataguardAssociationLifecycleStateEnum, bool) {
	enum, ok := mappingAutonomousDatabaseDataguardAssociationLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousDatabaseDataguardAssociationPeerRoleEnum Enum with underlying type: string
type AutonomousDatabaseDataguardAssociationPeerRoleEnum string

// Set of constants representing the allowable values for AutonomousDatabaseDataguardAssociationPeerRoleEnum
const (
	AutonomousDatabaseDataguardAssociationPeerRolePrimary         AutonomousDatabaseDataguardAssociationPeerRoleEnum = "PRIMARY"
	AutonomousDatabaseDataguardAssociationPeerRoleStandby         AutonomousDatabaseDataguardAssociationPeerRoleEnum = "STANDBY"
	AutonomousDatabaseDataguardAssociationPeerRoleDisabledStandby AutonomousDatabaseDataguardAssociationPeerRoleEnum = "DISABLED_STANDBY"
	AutonomousDatabaseDataguardAssociationPeerRoleBackupCopy      AutonomousDatabaseDataguardAssociationPeerRoleEnum = "BACKUP_COPY"
	AutonomousDatabaseDataguardAssociationPeerRoleSnapshotStandby AutonomousDatabaseDataguardAssociationPeerRoleEnum = "SNAPSHOT_STANDBY"
)

var mappingAutonomousDatabaseDataguardAssociationPeerRoleEnum = map[string]AutonomousDatabaseDataguardAssociationPeerRoleEnum{
	"PRIMARY":          AutonomousDatabaseDataguardAssociationPeerRolePrimary,
	"STANDBY":          AutonomousDatabaseDataguardAssociationPeerRoleStandby,
	"DISABLED_STANDBY": AutonomousDatabaseDataguardAssociationPeerRoleDisabledStandby,
	"BACKUP_COPY":      AutonomousDatabaseDataguardAssociationPeerRoleBackupCopy,
	"SNAPSHOT_STANDBY": AutonomousDatabaseDataguardAssociationPeerRoleSnapshotStandby,
}

var mappingAutonomousDatabaseDataguardAssociationPeerRoleEnumLowerCase = map[string]AutonomousDatabaseDataguardAssociationPeerRoleEnum{
	"primary":          AutonomousDatabaseDataguardAssociationPeerRolePrimary,
	"standby":          AutonomousDatabaseDataguardAssociationPeerRoleStandby,
	"disabled_standby": AutonomousDatabaseDataguardAssociationPeerRoleDisabledStandby,
	"backup_copy":      AutonomousDatabaseDataguardAssociationPeerRoleBackupCopy,
	"snapshot_standby": AutonomousDatabaseDataguardAssociationPeerRoleSnapshotStandby,
}

// GetAutonomousDatabaseDataguardAssociationPeerRoleEnumValues Enumerates the set of values for AutonomousDatabaseDataguardAssociationPeerRoleEnum
func GetAutonomousDatabaseDataguardAssociationPeerRoleEnumValues() []AutonomousDatabaseDataguardAssociationPeerRoleEnum {
	values := make([]AutonomousDatabaseDataguardAssociationPeerRoleEnum, 0)
	for _, v := range mappingAutonomousDatabaseDataguardAssociationPeerRoleEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDatabaseDataguardAssociationPeerRoleEnumStringValues Enumerates the set of values in String for AutonomousDatabaseDataguardAssociationPeerRoleEnum
func GetAutonomousDatabaseDataguardAssociationPeerRoleEnumStringValues() []string {
	return []string{
		"PRIMARY",
		"STANDBY",
		"DISABLED_STANDBY",
		"BACKUP_COPY",
		"SNAPSHOT_STANDBY",
	}
}

// GetMappingAutonomousDatabaseDataguardAssociationPeerRoleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDatabaseDataguardAssociationPeerRoleEnum(val string) (AutonomousDatabaseDataguardAssociationPeerRoleEnum, bool) {
	enum, ok := mappingAutonomousDatabaseDataguardAssociationPeerRoleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateEnum Enum with underlying type: string
type AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateEnum string

// Set of constants representing the allowable values for AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateEnum
const (
	AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateProvisioning            AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateEnum = "PROVISIONING"
	AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateAvailable               AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateEnum = "AVAILABLE"
	AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateStopping                AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateEnum = "STOPPING"
	AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateStopped                 AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateEnum = "STOPPED"
	AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateStarting                AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateEnum = "STARTING"
	AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateTerminating             AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateEnum = "TERMINATING"
	AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateTerminated              AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateEnum = "TERMINATED"
	AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateUnavailable             AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateEnum = "UNAVAILABLE"
	AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateRestoreInProgress       AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateEnum = "RESTORE_IN_PROGRESS"
	AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateRestoreFailed           AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateEnum = "RESTORE_FAILED"
	AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateBackupInProgress        AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateEnum = "BACKUP_IN_PROGRESS"
	AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateScaleInProgress         AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateEnum = "SCALE_IN_PROGRESS"
	AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateAvailableNeedsAttention AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateEnum = "AVAILABLE_NEEDS_ATTENTION"
	AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateUpdating                AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateEnum = "UPDATING"
	AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateMaintenanceInProgress   AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateEnum = "MAINTENANCE_IN_PROGRESS"
	AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateRestarting              AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateEnum = "RESTARTING"
	AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateRecreating              AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateEnum = "RECREATING"
	AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateRoleChangeInProgress    AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateEnum = "ROLE_CHANGE_IN_PROGRESS"
	AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateUpgrading               AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateEnum = "UPGRADING"
	AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateInaccessible            AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateEnum = "INACCESSIBLE"
	AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateStandby                 AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateEnum = "STANDBY"
)

var mappingAutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateEnum = map[string]AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateEnum{
	"PROVISIONING":              AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateProvisioning,
	"AVAILABLE":                 AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateAvailable,
	"STOPPING":                  AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateStopping,
	"STOPPED":                   AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateStopped,
	"STARTING":                  AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateStarting,
	"TERMINATING":               AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateTerminating,
	"TERMINATED":                AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateTerminated,
	"UNAVAILABLE":               AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateUnavailable,
	"RESTORE_IN_PROGRESS":       AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateRestoreInProgress,
	"RESTORE_FAILED":            AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateRestoreFailed,
	"BACKUP_IN_PROGRESS":        AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateBackupInProgress,
	"SCALE_IN_PROGRESS":         AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateScaleInProgress,
	"AVAILABLE_NEEDS_ATTENTION": AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateAvailableNeedsAttention,
	"UPDATING":                  AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateUpdating,
	"MAINTENANCE_IN_PROGRESS":   AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateMaintenanceInProgress,
	"RESTARTING":                AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateRestarting,
	"RECREATING":                AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateRecreating,
	"ROLE_CHANGE_IN_PROGRESS":   AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateRoleChangeInProgress,
	"UPGRADING":                 AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateUpgrading,
	"INACCESSIBLE":              AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateInaccessible,
	"STANDBY":                   AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateStandby,
}

var mappingAutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateEnumLowerCase = map[string]AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateEnum{
	"provisioning":              AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateProvisioning,
	"available":                 AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateAvailable,
	"stopping":                  AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateStopping,
	"stopped":                   AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateStopped,
	"starting":                  AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateStarting,
	"terminating":               AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateTerminating,
	"terminated":                AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateTerminated,
	"unavailable":               AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateUnavailable,
	"restore_in_progress":       AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateRestoreInProgress,
	"restore_failed":            AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateRestoreFailed,
	"backup_in_progress":        AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateBackupInProgress,
	"scale_in_progress":         AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateScaleInProgress,
	"available_needs_attention": AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateAvailableNeedsAttention,
	"updating":                  AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateUpdating,
	"maintenance_in_progress":   AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateMaintenanceInProgress,
	"restarting":                AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateRestarting,
	"recreating":                AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateRecreating,
	"role_change_in_progress":   AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateRoleChangeInProgress,
	"upgrading":                 AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateUpgrading,
	"inaccessible":              AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateInaccessible,
	"standby":                   AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateStandby,
}

// GetAutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateEnumValues Enumerates the set of values for AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateEnum
func GetAutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateEnumValues() []AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateEnum {
	values := make([]AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateEnum, 0)
	for _, v := range mappingAutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateEnumStringValues Enumerates the set of values in String for AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateEnum
func GetAutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateEnumStringValues() []string {
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

// GetMappingAutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateEnum(val string) (AutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateEnum, bool) {
	enum, ok := mappingAutonomousDatabaseDataguardAssociationPeerAutonomousDatabaseLifeCycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousDatabaseDataguardAssociationProtectionModeEnum Enum with underlying type: string
type AutonomousDatabaseDataguardAssociationProtectionModeEnum string

// Set of constants representing the allowable values for AutonomousDatabaseDataguardAssociationProtectionModeEnum
const (
	AutonomousDatabaseDataguardAssociationProtectionModeAvailability AutonomousDatabaseDataguardAssociationProtectionModeEnum = "MAXIMUM_AVAILABILITY"
	AutonomousDatabaseDataguardAssociationProtectionModePerformance  AutonomousDatabaseDataguardAssociationProtectionModeEnum = "MAXIMUM_PERFORMANCE"
)

var mappingAutonomousDatabaseDataguardAssociationProtectionModeEnum = map[string]AutonomousDatabaseDataguardAssociationProtectionModeEnum{
	"MAXIMUM_AVAILABILITY": AutonomousDatabaseDataguardAssociationProtectionModeAvailability,
	"MAXIMUM_PERFORMANCE":  AutonomousDatabaseDataguardAssociationProtectionModePerformance,
}

var mappingAutonomousDatabaseDataguardAssociationProtectionModeEnumLowerCase = map[string]AutonomousDatabaseDataguardAssociationProtectionModeEnum{
	"maximum_availability": AutonomousDatabaseDataguardAssociationProtectionModeAvailability,
	"maximum_performance":  AutonomousDatabaseDataguardAssociationProtectionModePerformance,
}

// GetAutonomousDatabaseDataguardAssociationProtectionModeEnumValues Enumerates the set of values for AutonomousDatabaseDataguardAssociationProtectionModeEnum
func GetAutonomousDatabaseDataguardAssociationProtectionModeEnumValues() []AutonomousDatabaseDataguardAssociationProtectionModeEnum {
	values := make([]AutonomousDatabaseDataguardAssociationProtectionModeEnum, 0)
	for _, v := range mappingAutonomousDatabaseDataguardAssociationProtectionModeEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousDatabaseDataguardAssociationProtectionModeEnumStringValues Enumerates the set of values in String for AutonomousDatabaseDataguardAssociationProtectionModeEnum
func GetAutonomousDatabaseDataguardAssociationProtectionModeEnumStringValues() []string {
	return []string{
		"MAXIMUM_AVAILABILITY",
		"MAXIMUM_PERFORMANCE",
	}
}

// GetMappingAutonomousDatabaseDataguardAssociationProtectionModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousDatabaseDataguardAssociationProtectionModeEnum(val string) (AutonomousDatabaseDataguardAssociationProtectionModeEnum, bool) {
	enum, ok := mappingAutonomousDatabaseDataguardAssociationProtectionModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
