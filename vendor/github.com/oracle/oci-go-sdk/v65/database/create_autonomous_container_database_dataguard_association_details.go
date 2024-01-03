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

// CreateAutonomousContainerDatabaseDataguardAssociationDetails Create Autonomous Dataguard Association to an existing Autonomous Container Database
type CreateAutonomousContainerDatabaseDataguardAssociationDetails struct {

	// The display name for the peer Autonomous Container Database.
	PeerAutonomousContainerDatabaseDisplayName *string `mandatory:"true" json:"peerAutonomousContainerDatabaseDisplayName"`

	// The protection mode of this Autonomous Data Guard association. For more information, see
	// Oracle Data Guard Protection Modes (http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-protection-modes.htm#SBYDB02000)
	// in the Oracle Data Guard documentation.
	ProtectionMode CreateAutonomousContainerDatabaseDataguardAssociationDetailsProtectionModeEnum `mandatory:"true" json:"protectionMode"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment where the standby Autonomous Container Database
	// will be created.
	PeerAutonomousContainerDatabaseCompartmentId *string `mandatory:"false" json:"peerAutonomousContainerDatabaseCompartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the peer cloud Autonomous Exadata VM Cluster.
	PeerCloudAutonomousVmClusterId *string `mandatory:"false" json:"peerCloudAutonomousVmClusterId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the peer Autonomous Exadata VM Cluster.
	PeerAutonomousVmClusterId *string `mandatory:"false" json:"peerAutonomousVmClusterId"`

	// Specifies the `DB_UNIQUE_NAME` of the peer database to be created.
	PeerDbUniqueName *string `mandatory:"false" json:"peerDbUniqueName"`

	PeerAutonomousContainerDatabaseBackupConfig *PeerAutonomousContainerDatabaseBackupConfig `mandatory:"false" json:"peerAutonomousContainerDatabaseBackupConfig"`

	// Indicates whether Automatic Failover is enabled for Autonomous Container Database Dataguard Association
	IsAutomaticFailoverEnabled *bool `mandatory:"false" json:"isAutomaticFailoverEnabled"`

	// The lag time for my preference based on data loss tolerance in seconds.
	FastStartFailOverLagLimitInSeconds *int `mandatory:"false" json:"fastStartFailOverLagLimitInSeconds"`

	// The scheduling detail for the quarterly maintenance window of the standby Autonomous Container Database.
	// This value represents the number of days before scheduled maintenance of the primary database.
	StandbyMaintenanceBufferInDays *int `mandatory:"false" json:"standbyMaintenanceBufferInDays"`
}

func (m CreateAutonomousContainerDatabaseDataguardAssociationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateAutonomousContainerDatabaseDataguardAssociationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateAutonomousContainerDatabaseDataguardAssociationDetailsProtectionModeEnum(string(m.ProtectionMode)); !ok && m.ProtectionMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ProtectionMode: %s. Supported values are: %s.", m.ProtectionMode, strings.Join(GetCreateAutonomousContainerDatabaseDataguardAssociationDetailsProtectionModeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateAutonomousContainerDatabaseDataguardAssociationDetailsProtectionModeEnum Enum with underlying type: string
type CreateAutonomousContainerDatabaseDataguardAssociationDetailsProtectionModeEnum string

// Set of constants representing the allowable values for CreateAutonomousContainerDatabaseDataguardAssociationDetailsProtectionModeEnum
const (
	CreateAutonomousContainerDatabaseDataguardAssociationDetailsProtectionModeAvailability CreateAutonomousContainerDatabaseDataguardAssociationDetailsProtectionModeEnum = "MAXIMUM_AVAILABILITY"
	CreateAutonomousContainerDatabaseDataguardAssociationDetailsProtectionModePerformance  CreateAutonomousContainerDatabaseDataguardAssociationDetailsProtectionModeEnum = "MAXIMUM_PERFORMANCE"
)

var mappingCreateAutonomousContainerDatabaseDataguardAssociationDetailsProtectionModeEnum = map[string]CreateAutonomousContainerDatabaseDataguardAssociationDetailsProtectionModeEnum{
	"MAXIMUM_AVAILABILITY": CreateAutonomousContainerDatabaseDataguardAssociationDetailsProtectionModeAvailability,
	"MAXIMUM_PERFORMANCE":  CreateAutonomousContainerDatabaseDataguardAssociationDetailsProtectionModePerformance,
}

var mappingCreateAutonomousContainerDatabaseDataguardAssociationDetailsProtectionModeEnumLowerCase = map[string]CreateAutonomousContainerDatabaseDataguardAssociationDetailsProtectionModeEnum{
	"maximum_availability": CreateAutonomousContainerDatabaseDataguardAssociationDetailsProtectionModeAvailability,
	"maximum_performance":  CreateAutonomousContainerDatabaseDataguardAssociationDetailsProtectionModePerformance,
}

// GetCreateAutonomousContainerDatabaseDataguardAssociationDetailsProtectionModeEnumValues Enumerates the set of values for CreateAutonomousContainerDatabaseDataguardAssociationDetailsProtectionModeEnum
func GetCreateAutonomousContainerDatabaseDataguardAssociationDetailsProtectionModeEnumValues() []CreateAutonomousContainerDatabaseDataguardAssociationDetailsProtectionModeEnum {
	values := make([]CreateAutonomousContainerDatabaseDataguardAssociationDetailsProtectionModeEnum, 0)
	for _, v := range mappingCreateAutonomousContainerDatabaseDataguardAssociationDetailsProtectionModeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateAutonomousContainerDatabaseDataguardAssociationDetailsProtectionModeEnumStringValues Enumerates the set of values in String for CreateAutonomousContainerDatabaseDataguardAssociationDetailsProtectionModeEnum
func GetCreateAutonomousContainerDatabaseDataguardAssociationDetailsProtectionModeEnumStringValues() []string {
	return []string{
		"MAXIMUM_AVAILABILITY",
		"MAXIMUM_PERFORMANCE",
	}
}

// GetMappingCreateAutonomousContainerDatabaseDataguardAssociationDetailsProtectionModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateAutonomousContainerDatabaseDataguardAssociationDetailsProtectionModeEnum(val string) (CreateAutonomousContainerDatabaseDataguardAssociationDetailsProtectionModeEnum, bool) {
	enum, ok := mappingCreateAutonomousContainerDatabaseDataguardAssociationDetailsProtectionModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
