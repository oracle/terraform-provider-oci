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

// UpdateAutonomousContainerDatabaseDataGuardAssociationDetails The configuration details for updating a Autonomous Container DatabaseData Guard association for a Autonomous Container Database.
type UpdateAutonomousContainerDatabaseDataGuardAssociationDetails struct {

	// Indicates whether Automatic Failover is enabled for Autonomous Container Database Dataguard Association
	IsAutomaticFailoverEnabled *bool `mandatory:"false" json:"isAutomaticFailoverEnabled"`

	// The protection mode of this Autonomous Data Guard association. For more information, see
	// Oracle Data Guard Protection Modes (http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-protection-modes.htm#SBYDB02000)
	// in the Oracle Data Guard documentation.
	ProtectionMode UpdateAutonomousContainerDatabaseDataGuardAssociationDetailsProtectionModeEnum `mandatory:"false" json:"protectionMode,omitempty"`

	// The lag time for my preference based on data loss tolerance in seconds.
	FastStartFailOverLagLimitInSeconds *int `mandatory:"false" json:"fastStartFailOverLagLimitInSeconds"`
}

func (m UpdateAutonomousContainerDatabaseDataGuardAssociationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateAutonomousContainerDatabaseDataGuardAssociationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateAutonomousContainerDatabaseDataGuardAssociationDetailsProtectionModeEnum(string(m.ProtectionMode)); !ok && m.ProtectionMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ProtectionMode: %s. Supported values are: %s.", m.ProtectionMode, strings.Join(GetUpdateAutonomousContainerDatabaseDataGuardAssociationDetailsProtectionModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateAutonomousContainerDatabaseDataGuardAssociationDetailsProtectionModeEnum Enum with underlying type: string
type UpdateAutonomousContainerDatabaseDataGuardAssociationDetailsProtectionModeEnum string

// Set of constants representing the allowable values for UpdateAutonomousContainerDatabaseDataGuardAssociationDetailsProtectionModeEnum
const (
	UpdateAutonomousContainerDatabaseDataGuardAssociationDetailsProtectionModeAvailability UpdateAutonomousContainerDatabaseDataGuardAssociationDetailsProtectionModeEnum = "MAXIMUM_AVAILABILITY"
	UpdateAutonomousContainerDatabaseDataGuardAssociationDetailsProtectionModePerformance  UpdateAutonomousContainerDatabaseDataGuardAssociationDetailsProtectionModeEnum = "MAXIMUM_PERFORMANCE"
)

var mappingUpdateAutonomousContainerDatabaseDataGuardAssociationDetailsProtectionModeEnum = map[string]UpdateAutonomousContainerDatabaseDataGuardAssociationDetailsProtectionModeEnum{
	"MAXIMUM_AVAILABILITY": UpdateAutonomousContainerDatabaseDataGuardAssociationDetailsProtectionModeAvailability,
	"MAXIMUM_PERFORMANCE":  UpdateAutonomousContainerDatabaseDataGuardAssociationDetailsProtectionModePerformance,
}

var mappingUpdateAutonomousContainerDatabaseDataGuardAssociationDetailsProtectionModeEnumLowerCase = map[string]UpdateAutonomousContainerDatabaseDataGuardAssociationDetailsProtectionModeEnum{
	"maximum_availability": UpdateAutonomousContainerDatabaseDataGuardAssociationDetailsProtectionModeAvailability,
	"maximum_performance":  UpdateAutonomousContainerDatabaseDataGuardAssociationDetailsProtectionModePerformance,
}

// GetUpdateAutonomousContainerDatabaseDataGuardAssociationDetailsProtectionModeEnumValues Enumerates the set of values for UpdateAutonomousContainerDatabaseDataGuardAssociationDetailsProtectionModeEnum
func GetUpdateAutonomousContainerDatabaseDataGuardAssociationDetailsProtectionModeEnumValues() []UpdateAutonomousContainerDatabaseDataGuardAssociationDetailsProtectionModeEnum {
	values := make([]UpdateAutonomousContainerDatabaseDataGuardAssociationDetailsProtectionModeEnum, 0)
	for _, v := range mappingUpdateAutonomousContainerDatabaseDataGuardAssociationDetailsProtectionModeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateAutonomousContainerDatabaseDataGuardAssociationDetailsProtectionModeEnumStringValues Enumerates the set of values in String for UpdateAutonomousContainerDatabaseDataGuardAssociationDetailsProtectionModeEnum
func GetUpdateAutonomousContainerDatabaseDataGuardAssociationDetailsProtectionModeEnumStringValues() []string {
	return []string{
		"MAXIMUM_AVAILABILITY",
		"MAXIMUM_PERFORMANCE",
	}
}

// GetMappingUpdateAutonomousContainerDatabaseDataGuardAssociationDetailsProtectionModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateAutonomousContainerDatabaseDataGuardAssociationDetailsProtectionModeEnum(val string) (UpdateAutonomousContainerDatabaseDataGuardAssociationDetailsProtectionModeEnum, bool) {
	enum, ok := mappingUpdateAutonomousContainerDatabaseDataGuardAssociationDetailsProtectionModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
