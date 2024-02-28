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

// AutonomousContainerDatabaseDataguardAssociation The properties that define Autonomous Data Guard association between two different Autonomous Container Databases.
type AutonomousContainerDatabaseDataguardAssociation struct {

	// The OCID of the Autonomous Data Guard created for a given Autonomous Container Database.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Autonomous Container Database that has a relationship with the peer Autonomous Container Database. Used only by Autonomous Database on Dedicated Exadata Infrastructure.
	AutonomousContainerDatabaseId *string `mandatory:"true" json:"autonomousContainerDatabaseId"`

	// The Data Guard role of the Autonomous Container Database or Autonomous Database, if Autonomous Data Guard is enabled.
	Role AutonomousContainerDatabaseDataguardAssociationRoleEnum `mandatory:"true" json:"role"`

	// The current state of Autonomous Data Guard.
	LifecycleState AutonomousContainerDatabaseDataguardAssociationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The Data Guard role of the Autonomous Container Database or Autonomous Database, if Autonomous Data Guard is enabled.
	PeerRole AutonomousContainerDatabaseDataguardAssociationPeerRoleEnum `mandatory:"true" json:"peerRole"`

	// Additional information about the current lifecycleState, if available.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The OCID of the peer Autonomous Container Database-Autonomous Data Guard association.
	PeerAutonomousContainerDatabaseDataguardAssociationId *string `mandatory:"false" json:"peerAutonomousContainerDatabaseDataguardAssociationId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the peer Autonomous Container Database.
	PeerAutonomousContainerDatabaseId *string `mandatory:"false" json:"peerAutonomousContainerDatabaseId"`

	// The current state of the Autonomous Container Database.
	PeerLifecycleState AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateEnum `mandatory:"false" json:"peerLifecycleState,omitempty"`

	// The protection mode of this Autonomous Data Guard association. For more information, see
	// Oracle Data Guard Protection Modes (http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-protection-modes.htm#SBYDB02000)
	// in the Oracle Data Guard documentation.
	ProtectionMode AutonomousContainerDatabaseDataguardAssociationProtectionModeEnum `mandatory:"false" json:"protectionMode,omitempty"`

	// The lag time for my preference based on data loss tolerance in seconds.
	FastStartFailOverLagLimitInSeconds *int `mandatory:"false" json:"fastStartFailOverLagLimitInSeconds"`

	// The lag time between updates to the primary Autonomous Container Database and application of the redo data on the standby Autonomous Container Database,
	// as computed by the reporting database.
	// Example: `9 seconds`
	ApplyLag *string `mandatory:"false" json:"applyLag"`

	// The rate at which redo logs are synchronized between the associated Autonomous Container Databases.
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

	// The date and time the Autonomous DataGuard association was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time when the last role change action happened.
	TimeLastRoleChanged *common.SDKTime `mandatory:"false" json:"timeLastRoleChanged"`
}

func (m AutonomousContainerDatabaseDataguardAssociation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AutonomousContainerDatabaseDataguardAssociation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAutonomousContainerDatabaseDataguardAssociationRoleEnum(string(m.Role)); !ok && m.Role != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Role: %s. Supported values are: %s.", m.Role, strings.Join(GetAutonomousContainerDatabaseDataguardAssociationRoleEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousContainerDatabaseDataguardAssociationLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAutonomousContainerDatabaseDataguardAssociationLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousContainerDatabaseDataguardAssociationPeerRoleEnum(string(m.PeerRole)); !ok && m.PeerRole != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PeerRole: %s. Supported values are: %s.", m.PeerRole, strings.Join(GetAutonomousContainerDatabaseDataguardAssociationPeerRoleEnumStringValues(), ",")))
	}

	if _, ok := GetMappingAutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateEnum(string(m.PeerLifecycleState)); !ok && m.PeerLifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PeerLifecycleState: %s. Supported values are: %s.", m.PeerLifecycleState, strings.Join(GetAutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousContainerDatabaseDataguardAssociationProtectionModeEnum(string(m.ProtectionMode)); !ok && m.ProtectionMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ProtectionMode: %s. Supported values are: %s.", m.ProtectionMode, strings.Join(GetAutonomousContainerDatabaseDataguardAssociationProtectionModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AutonomousContainerDatabaseDataguardAssociationRoleEnum Enum with underlying type: string
type AutonomousContainerDatabaseDataguardAssociationRoleEnum string

// Set of constants representing the allowable values for AutonomousContainerDatabaseDataguardAssociationRoleEnum
const (
	AutonomousContainerDatabaseDataguardAssociationRolePrimary         AutonomousContainerDatabaseDataguardAssociationRoleEnum = "PRIMARY"
	AutonomousContainerDatabaseDataguardAssociationRoleStandby         AutonomousContainerDatabaseDataguardAssociationRoleEnum = "STANDBY"
	AutonomousContainerDatabaseDataguardAssociationRoleDisabledStandby AutonomousContainerDatabaseDataguardAssociationRoleEnum = "DISABLED_STANDBY"
	AutonomousContainerDatabaseDataguardAssociationRoleBackupCopy      AutonomousContainerDatabaseDataguardAssociationRoleEnum = "BACKUP_COPY"
	AutonomousContainerDatabaseDataguardAssociationRoleSnapshotStandby AutonomousContainerDatabaseDataguardAssociationRoleEnum = "SNAPSHOT_STANDBY"
)

var mappingAutonomousContainerDatabaseDataguardAssociationRoleEnum = map[string]AutonomousContainerDatabaseDataguardAssociationRoleEnum{
	"PRIMARY":          AutonomousContainerDatabaseDataguardAssociationRolePrimary,
	"STANDBY":          AutonomousContainerDatabaseDataguardAssociationRoleStandby,
	"DISABLED_STANDBY": AutonomousContainerDatabaseDataguardAssociationRoleDisabledStandby,
	"BACKUP_COPY":      AutonomousContainerDatabaseDataguardAssociationRoleBackupCopy,
	"SNAPSHOT_STANDBY": AutonomousContainerDatabaseDataguardAssociationRoleSnapshotStandby,
}

var mappingAutonomousContainerDatabaseDataguardAssociationRoleEnumLowerCase = map[string]AutonomousContainerDatabaseDataguardAssociationRoleEnum{
	"primary":          AutonomousContainerDatabaseDataguardAssociationRolePrimary,
	"standby":          AutonomousContainerDatabaseDataguardAssociationRoleStandby,
	"disabled_standby": AutonomousContainerDatabaseDataguardAssociationRoleDisabledStandby,
	"backup_copy":      AutonomousContainerDatabaseDataguardAssociationRoleBackupCopy,
	"snapshot_standby": AutonomousContainerDatabaseDataguardAssociationRoleSnapshotStandby,
}

// GetAutonomousContainerDatabaseDataguardAssociationRoleEnumValues Enumerates the set of values for AutonomousContainerDatabaseDataguardAssociationRoleEnum
func GetAutonomousContainerDatabaseDataguardAssociationRoleEnumValues() []AutonomousContainerDatabaseDataguardAssociationRoleEnum {
	values := make([]AutonomousContainerDatabaseDataguardAssociationRoleEnum, 0)
	for _, v := range mappingAutonomousContainerDatabaseDataguardAssociationRoleEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousContainerDatabaseDataguardAssociationRoleEnumStringValues Enumerates the set of values in String for AutonomousContainerDatabaseDataguardAssociationRoleEnum
func GetAutonomousContainerDatabaseDataguardAssociationRoleEnumStringValues() []string {
	return []string{
		"PRIMARY",
		"STANDBY",
		"DISABLED_STANDBY",
		"BACKUP_COPY",
		"SNAPSHOT_STANDBY",
	}
}

// GetMappingAutonomousContainerDatabaseDataguardAssociationRoleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousContainerDatabaseDataguardAssociationRoleEnum(val string) (AutonomousContainerDatabaseDataguardAssociationRoleEnum, bool) {
	enum, ok := mappingAutonomousContainerDatabaseDataguardAssociationRoleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousContainerDatabaseDataguardAssociationLifecycleStateEnum Enum with underlying type: string
type AutonomousContainerDatabaseDataguardAssociationLifecycleStateEnum string

// Set of constants representing the allowable values for AutonomousContainerDatabaseDataguardAssociationLifecycleStateEnum
const (
	AutonomousContainerDatabaseDataguardAssociationLifecycleStateProvisioning         AutonomousContainerDatabaseDataguardAssociationLifecycleStateEnum = "PROVISIONING"
	AutonomousContainerDatabaseDataguardAssociationLifecycleStateAvailable            AutonomousContainerDatabaseDataguardAssociationLifecycleStateEnum = "AVAILABLE"
	AutonomousContainerDatabaseDataguardAssociationLifecycleStateRoleChangeInProgress AutonomousContainerDatabaseDataguardAssociationLifecycleStateEnum = "ROLE_CHANGE_IN_PROGRESS"
	AutonomousContainerDatabaseDataguardAssociationLifecycleStateTerminating          AutonomousContainerDatabaseDataguardAssociationLifecycleStateEnum = "TERMINATING"
	AutonomousContainerDatabaseDataguardAssociationLifecycleStateTerminated           AutonomousContainerDatabaseDataguardAssociationLifecycleStateEnum = "TERMINATED"
	AutonomousContainerDatabaseDataguardAssociationLifecycleStateFailed               AutonomousContainerDatabaseDataguardAssociationLifecycleStateEnum = "FAILED"
	AutonomousContainerDatabaseDataguardAssociationLifecycleStateUnavailable          AutonomousContainerDatabaseDataguardAssociationLifecycleStateEnum = "UNAVAILABLE"
	AutonomousContainerDatabaseDataguardAssociationLifecycleStateUpdating             AutonomousContainerDatabaseDataguardAssociationLifecycleStateEnum = "UPDATING"
)

var mappingAutonomousContainerDatabaseDataguardAssociationLifecycleStateEnum = map[string]AutonomousContainerDatabaseDataguardAssociationLifecycleStateEnum{
	"PROVISIONING":            AutonomousContainerDatabaseDataguardAssociationLifecycleStateProvisioning,
	"AVAILABLE":               AutonomousContainerDatabaseDataguardAssociationLifecycleStateAvailable,
	"ROLE_CHANGE_IN_PROGRESS": AutonomousContainerDatabaseDataguardAssociationLifecycleStateRoleChangeInProgress,
	"TERMINATING":             AutonomousContainerDatabaseDataguardAssociationLifecycleStateTerminating,
	"TERMINATED":              AutonomousContainerDatabaseDataguardAssociationLifecycleStateTerminated,
	"FAILED":                  AutonomousContainerDatabaseDataguardAssociationLifecycleStateFailed,
	"UNAVAILABLE":             AutonomousContainerDatabaseDataguardAssociationLifecycleStateUnavailable,
	"UPDATING":                AutonomousContainerDatabaseDataguardAssociationLifecycleStateUpdating,
}

var mappingAutonomousContainerDatabaseDataguardAssociationLifecycleStateEnumLowerCase = map[string]AutonomousContainerDatabaseDataguardAssociationLifecycleStateEnum{
	"provisioning":            AutonomousContainerDatabaseDataguardAssociationLifecycleStateProvisioning,
	"available":               AutonomousContainerDatabaseDataguardAssociationLifecycleStateAvailable,
	"role_change_in_progress": AutonomousContainerDatabaseDataguardAssociationLifecycleStateRoleChangeInProgress,
	"terminating":             AutonomousContainerDatabaseDataguardAssociationLifecycleStateTerminating,
	"terminated":              AutonomousContainerDatabaseDataguardAssociationLifecycleStateTerminated,
	"failed":                  AutonomousContainerDatabaseDataguardAssociationLifecycleStateFailed,
	"unavailable":             AutonomousContainerDatabaseDataguardAssociationLifecycleStateUnavailable,
	"updating":                AutonomousContainerDatabaseDataguardAssociationLifecycleStateUpdating,
}

// GetAutonomousContainerDatabaseDataguardAssociationLifecycleStateEnumValues Enumerates the set of values for AutonomousContainerDatabaseDataguardAssociationLifecycleStateEnum
func GetAutonomousContainerDatabaseDataguardAssociationLifecycleStateEnumValues() []AutonomousContainerDatabaseDataguardAssociationLifecycleStateEnum {
	values := make([]AutonomousContainerDatabaseDataguardAssociationLifecycleStateEnum, 0)
	for _, v := range mappingAutonomousContainerDatabaseDataguardAssociationLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousContainerDatabaseDataguardAssociationLifecycleStateEnumStringValues Enumerates the set of values in String for AutonomousContainerDatabaseDataguardAssociationLifecycleStateEnum
func GetAutonomousContainerDatabaseDataguardAssociationLifecycleStateEnumStringValues() []string {
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

// GetMappingAutonomousContainerDatabaseDataguardAssociationLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousContainerDatabaseDataguardAssociationLifecycleStateEnum(val string) (AutonomousContainerDatabaseDataguardAssociationLifecycleStateEnum, bool) {
	enum, ok := mappingAutonomousContainerDatabaseDataguardAssociationLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousContainerDatabaseDataguardAssociationPeerRoleEnum Enum with underlying type: string
type AutonomousContainerDatabaseDataguardAssociationPeerRoleEnum string

// Set of constants representing the allowable values for AutonomousContainerDatabaseDataguardAssociationPeerRoleEnum
const (
	AutonomousContainerDatabaseDataguardAssociationPeerRolePrimary         AutonomousContainerDatabaseDataguardAssociationPeerRoleEnum = "PRIMARY"
	AutonomousContainerDatabaseDataguardAssociationPeerRoleStandby         AutonomousContainerDatabaseDataguardAssociationPeerRoleEnum = "STANDBY"
	AutonomousContainerDatabaseDataguardAssociationPeerRoleDisabledStandby AutonomousContainerDatabaseDataguardAssociationPeerRoleEnum = "DISABLED_STANDBY"
	AutonomousContainerDatabaseDataguardAssociationPeerRoleBackupCopy      AutonomousContainerDatabaseDataguardAssociationPeerRoleEnum = "BACKUP_COPY"
	AutonomousContainerDatabaseDataguardAssociationPeerRoleSnapshotStandby AutonomousContainerDatabaseDataguardAssociationPeerRoleEnum = "SNAPSHOT_STANDBY"
)

var mappingAutonomousContainerDatabaseDataguardAssociationPeerRoleEnum = map[string]AutonomousContainerDatabaseDataguardAssociationPeerRoleEnum{
	"PRIMARY":          AutonomousContainerDatabaseDataguardAssociationPeerRolePrimary,
	"STANDBY":          AutonomousContainerDatabaseDataguardAssociationPeerRoleStandby,
	"DISABLED_STANDBY": AutonomousContainerDatabaseDataguardAssociationPeerRoleDisabledStandby,
	"BACKUP_COPY":      AutonomousContainerDatabaseDataguardAssociationPeerRoleBackupCopy,
	"SNAPSHOT_STANDBY": AutonomousContainerDatabaseDataguardAssociationPeerRoleSnapshotStandby,
}

var mappingAutonomousContainerDatabaseDataguardAssociationPeerRoleEnumLowerCase = map[string]AutonomousContainerDatabaseDataguardAssociationPeerRoleEnum{
	"primary":          AutonomousContainerDatabaseDataguardAssociationPeerRolePrimary,
	"standby":          AutonomousContainerDatabaseDataguardAssociationPeerRoleStandby,
	"disabled_standby": AutonomousContainerDatabaseDataguardAssociationPeerRoleDisabledStandby,
	"backup_copy":      AutonomousContainerDatabaseDataguardAssociationPeerRoleBackupCopy,
	"snapshot_standby": AutonomousContainerDatabaseDataguardAssociationPeerRoleSnapshotStandby,
}

// GetAutonomousContainerDatabaseDataguardAssociationPeerRoleEnumValues Enumerates the set of values for AutonomousContainerDatabaseDataguardAssociationPeerRoleEnum
func GetAutonomousContainerDatabaseDataguardAssociationPeerRoleEnumValues() []AutonomousContainerDatabaseDataguardAssociationPeerRoleEnum {
	values := make([]AutonomousContainerDatabaseDataguardAssociationPeerRoleEnum, 0)
	for _, v := range mappingAutonomousContainerDatabaseDataguardAssociationPeerRoleEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousContainerDatabaseDataguardAssociationPeerRoleEnumStringValues Enumerates the set of values in String for AutonomousContainerDatabaseDataguardAssociationPeerRoleEnum
func GetAutonomousContainerDatabaseDataguardAssociationPeerRoleEnumStringValues() []string {
	return []string{
		"PRIMARY",
		"STANDBY",
		"DISABLED_STANDBY",
		"BACKUP_COPY",
		"SNAPSHOT_STANDBY",
	}
}

// GetMappingAutonomousContainerDatabaseDataguardAssociationPeerRoleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousContainerDatabaseDataguardAssociationPeerRoleEnum(val string) (AutonomousContainerDatabaseDataguardAssociationPeerRoleEnum, bool) {
	enum, ok := mappingAutonomousContainerDatabaseDataguardAssociationPeerRoleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateEnum Enum with underlying type: string
type AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateEnum string

// Set of constants representing the allowable values for AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateEnum
const (
	AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateProvisioning                AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateEnum = "PROVISIONING"
	AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateAvailable                   AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateEnum = "AVAILABLE"
	AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateUpdating                    AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateEnum = "UPDATING"
	AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateTerminating                 AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateEnum = "TERMINATING"
	AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateTerminated                  AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateEnum = "TERMINATED"
	AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateFailed                      AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateEnum = "FAILED"
	AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateBackupInProgress            AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateEnum = "BACKUP_IN_PROGRESS"
	AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateRestoring                   AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateEnum = "RESTORING"
	AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateRestoreFailed               AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateEnum = "RESTORE_FAILED"
	AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateRestarting                  AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateEnum = "RESTARTING"
	AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateMaintenanceInProgress       AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateEnum = "MAINTENANCE_IN_PROGRESS"
	AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateRoleChangeInProgress        AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateEnum = "ROLE_CHANGE_IN_PROGRESS"
	AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateEnablingAutonomousDataGuard AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateEnum = "ENABLING_AUTONOMOUS_DATA_GUARD"
	AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateUnavailable                 AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateEnum = "UNAVAILABLE"
)

var mappingAutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateEnum = map[string]AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateEnum{
	"PROVISIONING":                   AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateProvisioning,
	"AVAILABLE":                      AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateAvailable,
	"UPDATING":                       AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateUpdating,
	"TERMINATING":                    AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateTerminating,
	"TERMINATED":                     AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateTerminated,
	"FAILED":                         AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateFailed,
	"BACKUP_IN_PROGRESS":             AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateBackupInProgress,
	"RESTORING":                      AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateRestoring,
	"RESTORE_FAILED":                 AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateRestoreFailed,
	"RESTARTING":                     AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateRestarting,
	"MAINTENANCE_IN_PROGRESS":        AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateMaintenanceInProgress,
	"ROLE_CHANGE_IN_PROGRESS":        AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateRoleChangeInProgress,
	"ENABLING_AUTONOMOUS_DATA_GUARD": AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateEnablingAutonomousDataGuard,
	"UNAVAILABLE":                    AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateUnavailable,
}

var mappingAutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateEnumLowerCase = map[string]AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateEnum{
	"provisioning":                   AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateProvisioning,
	"available":                      AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateAvailable,
	"updating":                       AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateUpdating,
	"terminating":                    AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateTerminating,
	"terminated":                     AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateTerminated,
	"failed":                         AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateFailed,
	"backup_in_progress":             AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateBackupInProgress,
	"restoring":                      AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateRestoring,
	"restore_failed":                 AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateRestoreFailed,
	"restarting":                     AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateRestarting,
	"maintenance_in_progress":        AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateMaintenanceInProgress,
	"role_change_in_progress":        AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateRoleChangeInProgress,
	"enabling_autonomous_data_guard": AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateEnablingAutonomousDataGuard,
	"unavailable":                    AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateUnavailable,
}

// GetAutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateEnumValues Enumerates the set of values for AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateEnum
func GetAutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateEnumValues() []AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateEnum {
	values := make([]AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateEnum, 0)
	for _, v := range mappingAutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateEnumStringValues Enumerates the set of values in String for AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateEnum
func GetAutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateEnumStringValues() []string {
	return []string{
		"PROVISIONING",
		"AVAILABLE",
		"UPDATING",
		"TERMINATING",
		"TERMINATED",
		"FAILED",
		"BACKUP_IN_PROGRESS",
		"RESTORING",
		"RESTORE_FAILED",
		"RESTARTING",
		"MAINTENANCE_IN_PROGRESS",
		"ROLE_CHANGE_IN_PROGRESS",
		"ENABLING_AUTONOMOUS_DATA_GUARD",
		"UNAVAILABLE",
	}
}

// GetMappingAutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateEnum(val string) (AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateEnum, bool) {
	enum, ok := mappingAutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousContainerDatabaseDataguardAssociationProtectionModeEnum Enum with underlying type: string
type AutonomousContainerDatabaseDataguardAssociationProtectionModeEnum string

// Set of constants representing the allowable values for AutonomousContainerDatabaseDataguardAssociationProtectionModeEnum
const (
	AutonomousContainerDatabaseDataguardAssociationProtectionModeAvailability AutonomousContainerDatabaseDataguardAssociationProtectionModeEnum = "MAXIMUM_AVAILABILITY"
	AutonomousContainerDatabaseDataguardAssociationProtectionModePerformance  AutonomousContainerDatabaseDataguardAssociationProtectionModeEnum = "MAXIMUM_PERFORMANCE"
)

var mappingAutonomousContainerDatabaseDataguardAssociationProtectionModeEnum = map[string]AutonomousContainerDatabaseDataguardAssociationProtectionModeEnum{
	"MAXIMUM_AVAILABILITY": AutonomousContainerDatabaseDataguardAssociationProtectionModeAvailability,
	"MAXIMUM_PERFORMANCE":  AutonomousContainerDatabaseDataguardAssociationProtectionModePerformance,
}

var mappingAutonomousContainerDatabaseDataguardAssociationProtectionModeEnumLowerCase = map[string]AutonomousContainerDatabaseDataguardAssociationProtectionModeEnum{
	"maximum_availability": AutonomousContainerDatabaseDataguardAssociationProtectionModeAvailability,
	"maximum_performance":  AutonomousContainerDatabaseDataguardAssociationProtectionModePerformance,
}

// GetAutonomousContainerDatabaseDataguardAssociationProtectionModeEnumValues Enumerates the set of values for AutonomousContainerDatabaseDataguardAssociationProtectionModeEnum
func GetAutonomousContainerDatabaseDataguardAssociationProtectionModeEnumValues() []AutonomousContainerDatabaseDataguardAssociationProtectionModeEnum {
	values := make([]AutonomousContainerDatabaseDataguardAssociationProtectionModeEnum, 0)
	for _, v := range mappingAutonomousContainerDatabaseDataguardAssociationProtectionModeEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousContainerDatabaseDataguardAssociationProtectionModeEnumStringValues Enumerates the set of values in String for AutonomousContainerDatabaseDataguardAssociationProtectionModeEnum
func GetAutonomousContainerDatabaseDataguardAssociationProtectionModeEnumStringValues() []string {
	return []string{
		"MAXIMUM_AVAILABILITY",
		"MAXIMUM_PERFORMANCE",
	}
}

// GetMappingAutonomousContainerDatabaseDataguardAssociationProtectionModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousContainerDatabaseDataguardAssociationProtectionModeEnum(val string) (AutonomousContainerDatabaseDataguardAssociationProtectionModeEnum, bool) {
	enum, ok := mappingAutonomousContainerDatabaseDataguardAssociationProtectionModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
