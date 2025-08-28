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

// EditAutonomousContainerDatabaseDataguardDetails The configuration details for updating a Autonomous Container DatabaseData Guard for a Autonomous Container Database.
type EditAutonomousContainerDatabaseDataguardDetails struct {

	// Indicates whether Automatic Failover is enabled for Autonomous Container Database Dataguard Association
	IsAutomaticFailoverEnabled *bool `mandatory:"false" json:"isAutomaticFailoverEnabled"`

	// The protection mode of this Autonomous Data Guard association. For more information, see
	// Oracle Data Guard Protection Modes (http://docs.oracle.com/database/122/SBYDB/oracle-data-guard-protection-modes.htm#SBYDB02000)
	// in the Oracle Data Guard documentation.
	ProtectionMode EditAutonomousContainerDatabaseDataguardDetailsProtectionModeEnum `mandatory:"false" json:"protectionMode,omitempty"`

	// The lag time for my preference based on data loss tolerance in seconds.
	FastStartFailOverLagLimitInSeconds *int `mandatory:"false" json:"fastStartFailOverLagLimitInSeconds"`
}

func (m EditAutonomousContainerDatabaseDataguardDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EditAutonomousContainerDatabaseDataguardDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingEditAutonomousContainerDatabaseDataguardDetailsProtectionModeEnum(string(m.ProtectionMode)); !ok && m.ProtectionMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ProtectionMode: %s. Supported values are: %s.", m.ProtectionMode, strings.Join(GetEditAutonomousContainerDatabaseDataguardDetailsProtectionModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// EditAutonomousContainerDatabaseDataguardDetailsProtectionModeEnum Enum with underlying type: string
type EditAutonomousContainerDatabaseDataguardDetailsProtectionModeEnum string

// Set of constants representing the allowable values for EditAutonomousContainerDatabaseDataguardDetailsProtectionModeEnum
const (
	EditAutonomousContainerDatabaseDataguardDetailsProtectionModeAvailability EditAutonomousContainerDatabaseDataguardDetailsProtectionModeEnum = "MAXIMUM_AVAILABILITY"
	EditAutonomousContainerDatabaseDataguardDetailsProtectionModePerformance  EditAutonomousContainerDatabaseDataguardDetailsProtectionModeEnum = "MAXIMUM_PERFORMANCE"
)

var mappingEditAutonomousContainerDatabaseDataguardDetailsProtectionModeEnum = map[string]EditAutonomousContainerDatabaseDataguardDetailsProtectionModeEnum{
	"MAXIMUM_AVAILABILITY": EditAutonomousContainerDatabaseDataguardDetailsProtectionModeAvailability,
	"MAXIMUM_PERFORMANCE":  EditAutonomousContainerDatabaseDataguardDetailsProtectionModePerformance,
}

var mappingEditAutonomousContainerDatabaseDataguardDetailsProtectionModeEnumLowerCase = map[string]EditAutonomousContainerDatabaseDataguardDetailsProtectionModeEnum{
	"maximum_availability": EditAutonomousContainerDatabaseDataguardDetailsProtectionModeAvailability,
	"maximum_performance":  EditAutonomousContainerDatabaseDataguardDetailsProtectionModePerformance,
}

// GetEditAutonomousContainerDatabaseDataguardDetailsProtectionModeEnumValues Enumerates the set of values for EditAutonomousContainerDatabaseDataguardDetailsProtectionModeEnum
func GetEditAutonomousContainerDatabaseDataguardDetailsProtectionModeEnumValues() []EditAutonomousContainerDatabaseDataguardDetailsProtectionModeEnum {
	values := make([]EditAutonomousContainerDatabaseDataguardDetailsProtectionModeEnum, 0)
	for _, v := range mappingEditAutonomousContainerDatabaseDataguardDetailsProtectionModeEnum {
		values = append(values, v)
	}
	return values
}

// GetEditAutonomousContainerDatabaseDataguardDetailsProtectionModeEnumStringValues Enumerates the set of values in String for EditAutonomousContainerDatabaseDataguardDetailsProtectionModeEnum
func GetEditAutonomousContainerDatabaseDataguardDetailsProtectionModeEnumStringValues() []string {
	return []string{
		"MAXIMUM_AVAILABILITY",
		"MAXIMUM_PERFORMANCE",
	}
}

// GetMappingEditAutonomousContainerDatabaseDataguardDetailsProtectionModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEditAutonomousContainerDatabaseDataguardDetailsProtectionModeEnum(val string) (EditAutonomousContainerDatabaseDataguardDetailsProtectionModeEnum, bool) {
	enum, ok := mappingEditAutonomousContainerDatabaseDataguardDetailsProtectionModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
