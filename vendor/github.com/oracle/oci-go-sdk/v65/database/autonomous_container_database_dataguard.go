// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// AutonomousContainerDatabaseDataguard The properties that define Autonomous Container Databases Dataguard.
type AutonomousContainerDatabaseDataguard struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous Container Database that has a relationship with the peer Autonomous Container Database. Used only by Autonomous Database on Dedicated Exadata Infrastructure.
	AutonomousContainerDatabaseId *string `mandatory:"true" json:"autonomousContainerDatabaseId"`

	// The Data Guard role of the Autonomous Container Database or Autonomous Database, if Autonomous Data Guard is enabled.
	Role AutonomousContainerDatabaseDataguardRoleEnum `mandatory:"true" json:"role"`

	// The current state of Autonomous Data Guard.
	LifecycleState AutonomousContainerDatabaseDataguardLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Additional information about the current lifecycleState, if available.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The protection mode of this Autonomous Data Guard association. For more information, see
	// Oracle Data Guard Protection Modes (http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-protection-modes.htm#SBYDB02000)
	// in the Oracle Data Guard documentation.
	ProtectionMode AutonomousContainerDatabaseDataguardProtectionModeEnum `mandatory:"false" json:"protectionMode,omitempty"`

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

	// The domain of the Autonomous Container Database
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// Timestamp when the lags were last calculated for a standby.
	TimeLagRefreshedOn *common.SDKTime `mandatory:"false" json:"timeLagRefreshedOn"`

	// Automatically selected by backend based on the protection mode.
	RedoTransportMode *string `mandatory:"false" json:"redoTransportMode"`

	// Automatically selected by backend when observer is enabled.
	AutomaticFailoverTarget *string `mandatory:"false" json:"automaticFailoverTarget"`
}

func (m AutonomousContainerDatabaseDataguard) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AutonomousContainerDatabaseDataguard) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAutonomousContainerDatabaseDataguardRoleEnum(string(m.Role)); !ok && m.Role != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Role: %s. Supported values are: %s.", m.Role, strings.Join(GetAutonomousContainerDatabaseDataguardRoleEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAutonomousContainerDatabaseDataguardLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAutonomousContainerDatabaseDataguardLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingAutonomousContainerDatabaseDataguardProtectionModeEnum(string(m.ProtectionMode)); !ok && m.ProtectionMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ProtectionMode: %s. Supported values are: %s.", m.ProtectionMode, strings.Join(GetAutonomousContainerDatabaseDataguardProtectionModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AutonomousContainerDatabaseDataguardRoleEnum Enum with underlying type: string
type AutonomousContainerDatabaseDataguardRoleEnum string

// Set of constants representing the allowable values for AutonomousContainerDatabaseDataguardRoleEnum
const (
	AutonomousContainerDatabaseDataguardRolePrimary         AutonomousContainerDatabaseDataguardRoleEnum = "PRIMARY"
	AutonomousContainerDatabaseDataguardRoleStandby         AutonomousContainerDatabaseDataguardRoleEnum = "STANDBY"
	AutonomousContainerDatabaseDataguardRoleDisabledStandby AutonomousContainerDatabaseDataguardRoleEnum = "DISABLED_STANDBY"
	AutonomousContainerDatabaseDataguardRoleBackupCopy      AutonomousContainerDatabaseDataguardRoleEnum = "BACKUP_COPY"
	AutonomousContainerDatabaseDataguardRoleSnapshotStandby AutonomousContainerDatabaseDataguardRoleEnum = "SNAPSHOT_STANDBY"
)

var mappingAutonomousContainerDatabaseDataguardRoleEnum = map[string]AutonomousContainerDatabaseDataguardRoleEnum{
	"PRIMARY":          AutonomousContainerDatabaseDataguardRolePrimary,
	"STANDBY":          AutonomousContainerDatabaseDataguardRoleStandby,
	"DISABLED_STANDBY": AutonomousContainerDatabaseDataguardRoleDisabledStandby,
	"BACKUP_COPY":      AutonomousContainerDatabaseDataguardRoleBackupCopy,
	"SNAPSHOT_STANDBY": AutonomousContainerDatabaseDataguardRoleSnapshotStandby,
}

var mappingAutonomousContainerDatabaseDataguardRoleEnumLowerCase = map[string]AutonomousContainerDatabaseDataguardRoleEnum{
	"primary":          AutonomousContainerDatabaseDataguardRolePrimary,
	"standby":          AutonomousContainerDatabaseDataguardRoleStandby,
	"disabled_standby": AutonomousContainerDatabaseDataguardRoleDisabledStandby,
	"backup_copy":      AutonomousContainerDatabaseDataguardRoleBackupCopy,
	"snapshot_standby": AutonomousContainerDatabaseDataguardRoleSnapshotStandby,
}

// GetAutonomousContainerDatabaseDataguardRoleEnumValues Enumerates the set of values for AutonomousContainerDatabaseDataguardRoleEnum
func GetAutonomousContainerDatabaseDataguardRoleEnumValues() []AutonomousContainerDatabaseDataguardRoleEnum {
	values := make([]AutonomousContainerDatabaseDataguardRoleEnum, 0)
	for _, v := range mappingAutonomousContainerDatabaseDataguardRoleEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousContainerDatabaseDataguardRoleEnumStringValues Enumerates the set of values in String for AutonomousContainerDatabaseDataguardRoleEnum
func GetAutonomousContainerDatabaseDataguardRoleEnumStringValues() []string {
	return []string{
		"PRIMARY",
		"STANDBY",
		"DISABLED_STANDBY",
		"BACKUP_COPY",
		"SNAPSHOT_STANDBY",
	}
}

// GetMappingAutonomousContainerDatabaseDataguardRoleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousContainerDatabaseDataguardRoleEnum(val string) (AutonomousContainerDatabaseDataguardRoleEnum, bool) {
	enum, ok := mappingAutonomousContainerDatabaseDataguardRoleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousContainerDatabaseDataguardLifecycleStateEnum Enum with underlying type: string
type AutonomousContainerDatabaseDataguardLifecycleStateEnum string

// Set of constants representing the allowable values for AutonomousContainerDatabaseDataguardLifecycleStateEnum
const (
	AutonomousContainerDatabaseDataguardLifecycleStateProvisioning         AutonomousContainerDatabaseDataguardLifecycleStateEnum = "PROVISIONING"
	AutonomousContainerDatabaseDataguardLifecycleStateAvailable            AutonomousContainerDatabaseDataguardLifecycleStateEnum = "AVAILABLE"
	AutonomousContainerDatabaseDataguardLifecycleStateRoleChangeInProgress AutonomousContainerDatabaseDataguardLifecycleStateEnum = "ROLE_CHANGE_IN_PROGRESS"
	AutonomousContainerDatabaseDataguardLifecycleStateTerminating          AutonomousContainerDatabaseDataguardLifecycleStateEnum = "TERMINATING"
	AutonomousContainerDatabaseDataguardLifecycleStateTerminated           AutonomousContainerDatabaseDataguardLifecycleStateEnum = "TERMINATED"
	AutonomousContainerDatabaseDataguardLifecycleStateFailed               AutonomousContainerDatabaseDataguardLifecycleStateEnum = "FAILED"
	AutonomousContainerDatabaseDataguardLifecycleStateUnavailable          AutonomousContainerDatabaseDataguardLifecycleStateEnum = "UNAVAILABLE"
	AutonomousContainerDatabaseDataguardLifecycleStateUpdating             AutonomousContainerDatabaseDataguardLifecycleStateEnum = "UPDATING"
)

var mappingAutonomousContainerDatabaseDataguardLifecycleStateEnum = map[string]AutonomousContainerDatabaseDataguardLifecycleStateEnum{
	"PROVISIONING":            AutonomousContainerDatabaseDataguardLifecycleStateProvisioning,
	"AVAILABLE":               AutonomousContainerDatabaseDataguardLifecycleStateAvailable,
	"ROLE_CHANGE_IN_PROGRESS": AutonomousContainerDatabaseDataguardLifecycleStateRoleChangeInProgress,
	"TERMINATING":             AutonomousContainerDatabaseDataguardLifecycleStateTerminating,
	"TERMINATED":              AutonomousContainerDatabaseDataguardLifecycleStateTerminated,
	"FAILED":                  AutonomousContainerDatabaseDataguardLifecycleStateFailed,
	"UNAVAILABLE":             AutonomousContainerDatabaseDataguardLifecycleStateUnavailable,
	"UPDATING":                AutonomousContainerDatabaseDataguardLifecycleStateUpdating,
}

var mappingAutonomousContainerDatabaseDataguardLifecycleStateEnumLowerCase = map[string]AutonomousContainerDatabaseDataguardLifecycleStateEnum{
	"provisioning":            AutonomousContainerDatabaseDataguardLifecycleStateProvisioning,
	"available":               AutonomousContainerDatabaseDataguardLifecycleStateAvailable,
	"role_change_in_progress": AutonomousContainerDatabaseDataguardLifecycleStateRoleChangeInProgress,
	"terminating":             AutonomousContainerDatabaseDataguardLifecycleStateTerminating,
	"terminated":              AutonomousContainerDatabaseDataguardLifecycleStateTerminated,
	"failed":                  AutonomousContainerDatabaseDataguardLifecycleStateFailed,
	"unavailable":             AutonomousContainerDatabaseDataguardLifecycleStateUnavailable,
	"updating":                AutonomousContainerDatabaseDataguardLifecycleStateUpdating,
}

// GetAutonomousContainerDatabaseDataguardLifecycleStateEnumValues Enumerates the set of values for AutonomousContainerDatabaseDataguardLifecycleStateEnum
func GetAutonomousContainerDatabaseDataguardLifecycleStateEnumValues() []AutonomousContainerDatabaseDataguardLifecycleStateEnum {
	values := make([]AutonomousContainerDatabaseDataguardLifecycleStateEnum, 0)
	for _, v := range mappingAutonomousContainerDatabaseDataguardLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousContainerDatabaseDataguardLifecycleStateEnumStringValues Enumerates the set of values in String for AutonomousContainerDatabaseDataguardLifecycleStateEnum
func GetAutonomousContainerDatabaseDataguardLifecycleStateEnumStringValues() []string {
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

// GetMappingAutonomousContainerDatabaseDataguardLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousContainerDatabaseDataguardLifecycleStateEnum(val string) (AutonomousContainerDatabaseDataguardLifecycleStateEnum, bool) {
	enum, ok := mappingAutonomousContainerDatabaseDataguardLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AutonomousContainerDatabaseDataguardProtectionModeEnum Enum with underlying type: string
type AutonomousContainerDatabaseDataguardProtectionModeEnum string

// Set of constants representing the allowable values for AutonomousContainerDatabaseDataguardProtectionModeEnum
const (
	AutonomousContainerDatabaseDataguardProtectionModeAvailability AutonomousContainerDatabaseDataguardProtectionModeEnum = "MAXIMUM_AVAILABILITY"
	AutonomousContainerDatabaseDataguardProtectionModePerformance  AutonomousContainerDatabaseDataguardProtectionModeEnum = "MAXIMUM_PERFORMANCE"
)

var mappingAutonomousContainerDatabaseDataguardProtectionModeEnum = map[string]AutonomousContainerDatabaseDataguardProtectionModeEnum{
	"MAXIMUM_AVAILABILITY": AutonomousContainerDatabaseDataguardProtectionModeAvailability,
	"MAXIMUM_PERFORMANCE":  AutonomousContainerDatabaseDataguardProtectionModePerformance,
}

var mappingAutonomousContainerDatabaseDataguardProtectionModeEnumLowerCase = map[string]AutonomousContainerDatabaseDataguardProtectionModeEnum{
	"maximum_availability": AutonomousContainerDatabaseDataguardProtectionModeAvailability,
	"maximum_performance":  AutonomousContainerDatabaseDataguardProtectionModePerformance,
}

// GetAutonomousContainerDatabaseDataguardProtectionModeEnumValues Enumerates the set of values for AutonomousContainerDatabaseDataguardProtectionModeEnum
func GetAutonomousContainerDatabaseDataguardProtectionModeEnumValues() []AutonomousContainerDatabaseDataguardProtectionModeEnum {
	values := make([]AutonomousContainerDatabaseDataguardProtectionModeEnum, 0)
	for _, v := range mappingAutonomousContainerDatabaseDataguardProtectionModeEnum {
		values = append(values, v)
	}
	return values
}

// GetAutonomousContainerDatabaseDataguardProtectionModeEnumStringValues Enumerates the set of values in String for AutonomousContainerDatabaseDataguardProtectionModeEnum
func GetAutonomousContainerDatabaseDataguardProtectionModeEnumStringValues() []string {
	return []string{
		"MAXIMUM_AVAILABILITY",
		"MAXIMUM_PERFORMANCE",
	}
}

// GetMappingAutonomousContainerDatabaseDataguardProtectionModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutonomousContainerDatabaseDataguardProtectionModeEnum(val string) (AutonomousContainerDatabaseDataguardProtectionModeEnum, bool) {
	enum, ok := mappingAutonomousContainerDatabaseDataguardProtectionModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
