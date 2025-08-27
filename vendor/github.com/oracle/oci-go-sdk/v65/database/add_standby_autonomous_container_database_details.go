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

// AddStandbyAutonomousContainerDatabaseDetails Create Standby Autonomous Container Database to an existing Autonomous Container Database
type AddStandbyAutonomousContainerDatabaseDetails struct {

	// The lag time for my preference based on data loss tolerance in seconds.
	FastStartFailOverLagLimitInSeconds *int `mandatory:"false" json:"fastStartFailOverLagLimitInSeconds"`

	// Indicates whether Automatic Failover is enabled for Autonomous Container Database Dataguard Association
	IsAutomaticFailoverEnabled *bool `mandatory:"false" json:"isAutomaticFailoverEnabled"`

	PeerAutonomousContainerDatabaseBackupConfig *PeerAutonomousContainerDatabaseBackupConfig `mandatory:"false" json:"peerAutonomousContainerDatabaseBackupConfig"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where the standby Autonomous Container Database
	// will be created.
	PeerAutonomousContainerDatabaseCompartmentId *string `mandatory:"false" json:"peerAutonomousContainerDatabaseCompartmentId"`

	// The display name for the peer Autonomous Container Database.
	PeerAutonomousContainerDatabaseDisplayName *string `mandatory:"false" json:"peerAutonomousContainerDatabaseDisplayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the peer cloud Autonomous Exadata VM Cluster.
	PeerCloudAutonomousVmClusterId *string `mandatory:"false" json:"peerCloudAutonomousVmClusterId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the peer Autonomous Exadata VM Cluster.
	PeerAutonomousVmClusterId *string `mandatory:"false" json:"peerAutonomousVmClusterId"`

	// Specifies the `DB_UNIQUE_NAME` of the peer database to be created.
	PeerDbUniqueName *string `mandatory:"false" json:"peerDbUniqueName"`

	// The protection mode of this Autonomous Data Guard association. For more information, see
	// Oracle Data Guard Protection Modes (http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-protection-modes.htm#SBYDB02000)
	// in the Oracle Data Guard documentation.
	ProtectionMode AddStandbyAutonomousContainerDatabaseDetailsProtectionModeEnum `mandatory:"false" json:"protectionMode,omitempty"`

	// The scheduling detail for the quarterly maintenance window of the standby Autonomous Container Database.
	// This value represents the number of days before scheduled maintenance of the primary database.
	StandbyMaintenanceBufferInDays *int `mandatory:"false" json:"standbyMaintenanceBufferInDays"`
}

func (m AddStandbyAutonomousContainerDatabaseDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AddStandbyAutonomousContainerDatabaseDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingAddStandbyAutonomousContainerDatabaseDetailsProtectionModeEnum(string(m.ProtectionMode)); !ok && m.ProtectionMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ProtectionMode: %s. Supported values are: %s.", m.ProtectionMode, strings.Join(GetAddStandbyAutonomousContainerDatabaseDetailsProtectionModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AddStandbyAutonomousContainerDatabaseDetailsProtectionModeEnum Enum with underlying type: string
type AddStandbyAutonomousContainerDatabaseDetailsProtectionModeEnum string

// Set of constants representing the allowable values for AddStandbyAutonomousContainerDatabaseDetailsProtectionModeEnum
const (
	AddStandbyAutonomousContainerDatabaseDetailsProtectionModeAvailability AddStandbyAutonomousContainerDatabaseDetailsProtectionModeEnum = "MAXIMUM_AVAILABILITY"
	AddStandbyAutonomousContainerDatabaseDetailsProtectionModePerformance  AddStandbyAutonomousContainerDatabaseDetailsProtectionModeEnum = "MAXIMUM_PERFORMANCE"
)

var mappingAddStandbyAutonomousContainerDatabaseDetailsProtectionModeEnum = map[string]AddStandbyAutonomousContainerDatabaseDetailsProtectionModeEnum{
	"MAXIMUM_AVAILABILITY": AddStandbyAutonomousContainerDatabaseDetailsProtectionModeAvailability,
	"MAXIMUM_PERFORMANCE":  AddStandbyAutonomousContainerDatabaseDetailsProtectionModePerformance,
}

var mappingAddStandbyAutonomousContainerDatabaseDetailsProtectionModeEnumLowerCase = map[string]AddStandbyAutonomousContainerDatabaseDetailsProtectionModeEnum{
	"maximum_availability": AddStandbyAutonomousContainerDatabaseDetailsProtectionModeAvailability,
	"maximum_performance":  AddStandbyAutonomousContainerDatabaseDetailsProtectionModePerformance,
}

// GetAddStandbyAutonomousContainerDatabaseDetailsProtectionModeEnumValues Enumerates the set of values for AddStandbyAutonomousContainerDatabaseDetailsProtectionModeEnum
func GetAddStandbyAutonomousContainerDatabaseDetailsProtectionModeEnumValues() []AddStandbyAutonomousContainerDatabaseDetailsProtectionModeEnum {
	values := make([]AddStandbyAutonomousContainerDatabaseDetailsProtectionModeEnum, 0)
	for _, v := range mappingAddStandbyAutonomousContainerDatabaseDetailsProtectionModeEnum {
		values = append(values, v)
	}
	return values
}

// GetAddStandbyAutonomousContainerDatabaseDetailsProtectionModeEnumStringValues Enumerates the set of values in String for AddStandbyAutonomousContainerDatabaseDetailsProtectionModeEnum
func GetAddStandbyAutonomousContainerDatabaseDetailsProtectionModeEnumStringValues() []string {
	return []string{
		"MAXIMUM_AVAILABILITY",
		"MAXIMUM_PERFORMANCE",
	}
}

// GetMappingAddStandbyAutonomousContainerDatabaseDetailsProtectionModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAddStandbyAutonomousContainerDatabaseDetailsProtectionModeEnum(val string) (AddStandbyAutonomousContainerDatabaseDetailsProtectionModeEnum, bool) {
	enum, ok := mappingAddStandbyAutonomousContainerDatabaseDetailsProtectionModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
