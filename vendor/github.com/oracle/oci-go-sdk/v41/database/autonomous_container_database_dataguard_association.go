// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"github.com/oracle/oci-go-sdk/v41/common"
)

// AutonomousContainerDatabaseDataguardAssociation The properties that define Autonomous Data Guard association between two different Autonomous Container Databases.
type AutonomousContainerDatabaseDataguardAssociation struct {

	// The OCID of the Autonomous Data Guard created for a given Autonomous Container Database.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Autonomous Container Database that has a relationship with the peer Autonomous Container Database.
	AutonomousContainerDatabaseId *string `mandatory:"true" json:"autonomousContainerDatabaseId"`

	// The Data Guard role of the Autonomous Container Database, if Autonomous Data Guard is enabled.
	Role AutonomousContainerDatabaseDataguardAssociationRoleEnum `mandatory:"true" json:"role"`

	// The current state of Autonomous Data Guard.
	LifecycleState AutonomousContainerDatabaseDataguardAssociationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The Data Guard role of the Autonomous Container Database, if Autonomous Data Guard is enabled.
	PeerRole AutonomousContainerDatabaseDataguardAssociationPeerRoleEnum `mandatory:"true" json:"peerRole"`

	// Additional information about the current lifecycleState, if available.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The OCID of the peer Autonomous Container Database-Autonomous Data Guard association.
	PeerAutonomousContainerDatabaseDataguardAssociationId *string `mandatory:"false" json:"peerAutonomousContainerDatabaseDataguardAssociationId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the peer Autonomous Container Database.
	PeerAutonomousContainerDatabaseId *string `mandatory:"false" json:"peerAutonomousContainerDatabaseId"`

	// The current state of Autonomous Data Guard.
	PeerLifecycleState AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateEnum `mandatory:"false" json:"peerLifecycleState,omitempty"`

	// The protection mode of this Autonomous Data Guard association. For more information, see
	// Oracle Data Guard Protection Modes (http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-protection-modes.htm#SBYDB02000)
	// in the Oracle Data Guard documentation.
	ProtectionMode AutonomousContainerDatabaseDataguardAssociationProtectionModeEnum `mandatory:"false" json:"protectionMode,omitempty"`

	// The lag time between updates to the primary Autonomous Container Database and application of the redo data on the standby Autonomous Container Database,
	// as computed by the reporting database.
	// Example: `9 seconds`
	ApplyLag *string `mandatory:"false" json:"applyLag"`

	// The rate at which redo logs are synchronized between the associated Autonomous Container Databases.
	// Example: `180 Mb per second`
	ApplyRate *string `mandatory:"false" json:"applyRate"`

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

// AutonomousContainerDatabaseDataguardAssociationRoleEnum Enum with underlying type: string
type AutonomousContainerDatabaseDataguardAssociationRoleEnum string

// Set of constants representing the allowable values for AutonomousContainerDatabaseDataguardAssociationRoleEnum
const (
	AutonomousContainerDatabaseDataguardAssociationRolePrimary         AutonomousContainerDatabaseDataguardAssociationRoleEnum = "PRIMARY"
	AutonomousContainerDatabaseDataguardAssociationRoleStandby         AutonomousContainerDatabaseDataguardAssociationRoleEnum = "STANDBY"
	AutonomousContainerDatabaseDataguardAssociationRoleDisabledStandby AutonomousContainerDatabaseDataguardAssociationRoleEnum = "DISABLED_STANDBY"
)

var mappingAutonomousContainerDatabaseDataguardAssociationRole = map[string]AutonomousContainerDatabaseDataguardAssociationRoleEnum{
	"PRIMARY":          AutonomousContainerDatabaseDataguardAssociationRolePrimary,
	"STANDBY":          AutonomousContainerDatabaseDataguardAssociationRoleStandby,
	"DISABLED_STANDBY": AutonomousContainerDatabaseDataguardAssociationRoleDisabledStandby,
}

// GetAutonomousContainerDatabaseDataguardAssociationRoleEnumValues Enumerates the set of values for AutonomousContainerDatabaseDataguardAssociationRoleEnum
func GetAutonomousContainerDatabaseDataguardAssociationRoleEnumValues() []AutonomousContainerDatabaseDataguardAssociationRoleEnum {
	values := make([]AutonomousContainerDatabaseDataguardAssociationRoleEnum, 0)
	for _, v := range mappingAutonomousContainerDatabaseDataguardAssociationRole {
		values = append(values, v)
	}
	return values
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
)

var mappingAutonomousContainerDatabaseDataguardAssociationLifecycleState = map[string]AutonomousContainerDatabaseDataguardAssociationLifecycleStateEnum{
	"PROVISIONING":            AutonomousContainerDatabaseDataguardAssociationLifecycleStateProvisioning,
	"AVAILABLE":               AutonomousContainerDatabaseDataguardAssociationLifecycleStateAvailable,
	"ROLE_CHANGE_IN_PROGRESS": AutonomousContainerDatabaseDataguardAssociationLifecycleStateRoleChangeInProgress,
	"TERMINATING":             AutonomousContainerDatabaseDataguardAssociationLifecycleStateTerminating,
	"TERMINATED":              AutonomousContainerDatabaseDataguardAssociationLifecycleStateTerminated,
	"FAILED":                  AutonomousContainerDatabaseDataguardAssociationLifecycleStateFailed,
	"UNAVAILABLE":             AutonomousContainerDatabaseDataguardAssociationLifecycleStateUnavailable,
}

// GetAutonomousContainerDatabaseDataguardAssociationLifecycleStateEnumValues Enumerates the set of values for AutonomousContainerDatabaseDataguardAssociationLifecycleStateEnum
func GetAutonomousContainerDatabaseDataguardAssociationLifecycleStateEnumValues() []AutonomousContainerDatabaseDataguardAssociationLifecycleStateEnum {
	values := make([]AutonomousContainerDatabaseDataguardAssociationLifecycleStateEnum, 0)
	for _, v := range mappingAutonomousContainerDatabaseDataguardAssociationLifecycleState {
		values = append(values, v)
	}
	return values
}

// AutonomousContainerDatabaseDataguardAssociationPeerRoleEnum Enum with underlying type: string
type AutonomousContainerDatabaseDataguardAssociationPeerRoleEnum string

// Set of constants representing the allowable values for AutonomousContainerDatabaseDataguardAssociationPeerRoleEnum
const (
	AutonomousContainerDatabaseDataguardAssociationPeerRolePrimary         AutonomousContainerDatabaseDataguardAssociationPeerRoleEnum = "PRIMARY"
	AutonomousContainerDatabaseDataguardAssociationPeerRoleStandby         AutonomousContainerDatabaseDataguardAssociationPeerRoleEnum = "STANDBY"
	AutonomousContainerDatabaseDataguardAssociationPeerRoleDisabledStandby AutonomousContainerDatabaseDataguardAssociationPeerRoleEnum = "DISABLED_STANDBY"
)

var mappingAutonomousContainerDatabaseDataguardAssociationPeerRole = map[string]AutonomousContainerDatabaseDataguardAssociationPeerRoleEnum{
	"PRIMARY":          AutonomousContainerDatabaseDataguardAssociationPeerRolePrimary,
	"STANDBY":          AutonomousContainerDatabaseDataguardAssociationPeerRoleStandby,
	"DISABLED_STANDBY": AutonomousContainerDatabaseDataguardAssociationPeerRoleDisabledStandby,
}

// GetAutonomousContainerDatabaseDataguardAssociationPeerRoleEnumValues Enumerates the set of values for AutonomousContainerDatabaseDataguardAssociationPeerRoleEnum
func GetAutonomousContainerDatabaseDataguardAssociationPeerRoleEnumValues() []AutonomousContainerDatabaseDataguardAssociationPeerRoleEnum {
	values := make([]AutonomousContainerDatabaseDataguardAssociationPeerRoleEnum, 0)
	for _, v := range mappingAutonomousContainerDatabaseDataguardAssociationPeerRole {
		values = append(values, v)
	}
	return values
}

// AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateEnum Enum with underlying type: string
type AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateEnum string

// Set of constants representing the allowable values for AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateEnum
const (
	AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateProvisioning         AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateEnum = "PROVISIONING"
	AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateAvailable            AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateEnum = "AVAILABLE"
	AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateRoleChangeInProgress AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateEnum = "ROLE_CHANGE_IN_PROGRESS"
	AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateTerminating          AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateEnum = "TERMINATING"
	AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateTerminated           AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateEnum = "TERMINATED"
	AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateFailed               AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateEnum = "FAILED"
	AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateUnavailable          AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateEnum = "UNAVAILABLE"
)

var mappingAutonomousContainerDatabaseDataguardAssociationPeerLifecycleState = map[string]AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateEnum{
	"PROVISIONING":            AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateProvisioning,
	"AVAILABLE":               AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateAvailable,
	"ROLE_CHANGE_IN_PROGRESS": AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateRoleChangeInProgress,
	"TERMINATING":             AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateTerminating,
	"TERMINATED":              AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateTerminated,
	"FAILED":                  AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateFailed,
	"UNAVAILABLE":             AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateUnavailable,
}

// GetAutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateEnumValues Enumerates the set of values for AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateEnum
func GetAutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateEnumValues() []AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateEnum {
	values := make([]AutonomousContainerDatabaseDataguardAssociationPeerLifecycleStateEnum, 0)
	for _, v := range mappingAutonomousContainerDatabaseDataguardAssociationPeerLifecycleState {
		values = append(values, v)
	}
	return values
}

// AutonomousContainerDatabaseDataguardAssociationProtectionModeEnum Enum with underlying type: string
type AutonomousContainerDatabaseDataguardAssociationProtectionModeEnum string

// Set of constants representing the allowable values for AutonomousContainerDatabaseDataguardAssociationProtectionModeEnum
const (
	AutonomousContainerDatabaseDataguardAssociationProtectionModeAvailability AutonomousContainerDatabaseDataguardAssociationProtectionModeEnum = "MAXIMUM_AVAILABILITY"
	AutonomousContainerDatabaseDataguardAssociationProtectionModePerformance  AutonomousContainerDatabaseDataguardAssociationProtectionModeEnum = "MAXIMUM_PERFORMANCE"
)

var mappingAutonomousContainerDatabaseDataguardAssociationProtectionMode = map[string]AutonomousContainerDatabaseDataguardAssociationProtectionModeEnum{
	"MAXIMUM_AVAILABILITY": AutonomousContainerDatabaseDataguardAssociationProtectionModeAvailability,
	"MAXIMUM_PERFORMANCE":  AutonomousContainerDatabaseDataguardAssociationProtectionModePerformance,
}

// GetAutonomousContainerDatabaseDataguardAssociationProtectionModeEnumValues Enumerates the set of values for AutonomousContainerDatabaseDataguardAssociationProtectionModeEnum
func GetAutonomousContainerDatabaseDataguardAssociationProtectionModeEnumValues() []AutonomousContainerDatabaseDataguardAssociationProtectionModeEnum {
	values := make([]AutonomousContainerDatabaseDataguardAssociationProtectionModeEnum, 0)
	for _, v := range mappingAutonomousContainerDatabaseDataguardAssociationProtectionMode {
		values = append(values, v)
	}
	return values
}
